package util

import (
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/config"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/storage"
)

type Config struct {
	RootUrl              string                `json:"rootUrl"`
	Provider             string                `json:"provider"`
	StaticProviderConfig *StaticProviderConfig `json:"staticProvider"`
	S3ProviderConfig     *S3ProviderConfig     `json:"s3Provider"`
}

// S3ProviderConfig shares the ncraft object storage configuration.
type S3ProviderConfig = storage.Config

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
