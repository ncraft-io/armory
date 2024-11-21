package hook

import (
	"context"
	"github.com/ncraft-io/armory/service-go/pkg/synchro"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
)

type SQLRunner struct {
	SQL string
}

func (r *SQLRunner) Run(ctx context.Context) {
	if r != nil {
		err := synchro.GetDataDB().Exec(r.SQL).Error
		if err != nil {
			logs.ErrLogw("failed to run the sql", "sql", r.SQL, "err", err)
		} else {
			logs.Infow("success to run the sql", "sql", r.SQL)
		}
	}
}
