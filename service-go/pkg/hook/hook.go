package hook

import (
	"context"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/config"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"sync"
	"time"
)

type Hook struct {
	Config *Config

	SQLRunners []*SQLRunner
}

type Event struct {
	Time   time.Time
	Action bool
}

var hook *Hook
var hookOnce sync.Once

var eventChannel = make(chan *Event, 1000)

func New() *Hook {
	cfg := &Config{}
	_ = config.ScanFrom(cfg, "hook")

	h := &Hook{
		Config: cfg,
	}

	//TODO should make some sql check
	for _, s := range h.Config.Sqls {
		h.SQLRunners = append(h.SQLRunners, &SQLRunner{SQL: s})
	}

	return h
}

func GetHook() *Hook {
	hookOnce.Do(func() {
		hook = New()
	})
	return hook
}

func init() {
	if !GetHook().Config.Enable {
		return
	}

	// timer
	timer := time.NewTimer(GetHook().MinInterval())
	maxTimer := time.NewTimer(GetHook().MaxInterval()) // max interval to invoke the hooks

	cycle := 0

	var flags sync.Map // 1: created, 2: doing, 3: done

	// make sure run the hooks within at least 1 min interval.
	go func() {
		for {
			select {
			case <-eventChannel:
				if s, ok := flags.Load(cycle); ok {
					if s == 2 {
						flags.Store(cycle+1, 1)
						logs.Infow("received the hook event, but the event is processing now, then put to next")
					} else {
						logs.Infow("received the duplicated hook event, skip it ")
					}
				} else {
					flags.Store(cycle, 1)
					logs.Infow("received the hook event")
				}
			case <-timer.C:
				if s, ok := flags.LoadOrStore(cycle, 2); ok && s == 1 {
					logs.Infow("run the hook")

					for _, r := range GetHook().SQLRunners {
						r.Run(context.Background())
					}

					flags.Delete(cycle)
				}

				cycle++
				timer.Reset(GetHook().MinInterval())
			case <-maxTimer.C:
				flags.Store(cycle, 1)
				maxTimer.Reset(GetHook().MaxInterval())
				logs.Infow("send the max interval hook event")
			}
		}
	}()
}

func (h *Hook) Run(ctx context.Context) {
	_ = ctx
	if h.Config.Enable {
		eventChannel <- &Event{
			Time:   time.Now(),
			Action: true,
		}
		logs.Infow("sent the hook event to the channel")
	}
}

func (h *Hook) MinInterval() time.Duration {
	if h != nil && h.Config != nil && h.Config.MinInterval > 0 {
		return time.Duration(h.Config.MinInterval) * time.Second
	}
	return 60 * time.Second
}

func (h *Hook) MaxInterval() time.Duration {
	if h != nil && h.Config != nil && h.Config.MaxInterval > 0 {
		return time.Duration(h.Config.MaxInterval) * time.Second
	}
	return 60 * 60 * time.Second
}
