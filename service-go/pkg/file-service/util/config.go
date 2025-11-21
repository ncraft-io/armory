package util

import (
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/config"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
)

type Config struct {
	RootUrl              string                `json:"rootUrl"`
	StaticProviderConfig *StaticProviderConfig `json:"staticProvider"`
}

type StaticProviderConfig struct {
	Root string `json:"root"`
}

func GetConfig() (*Config, error) {
	conf := &Config{}
	err := config.ScanFrom(&conf, "file")
	if err != nil {
		logs.Warnw("failed to get the file config from config file", "error", err.Error())
		return nil, err
	}
	return conf, nil
}
