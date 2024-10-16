package hook

type Config struct {
	Enable      bool     `json:"enable,omitempty"`
	MinInterval int      `json:"minInterval,omitempty"`
	MaxInterval int      `json:"maxInterval,omitempty"`
	Sqls        []string `json:"sqls,omitempty"`
}
