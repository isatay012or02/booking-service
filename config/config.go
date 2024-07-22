package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Http      HttpCfg
	DB        DBSettings
	SwaggerUI SwaggerUIConfig
}

type HttpCfg struct {
	Port int `json:"port"`
	Gin  struct {
		ReleaseMode bool `json:"ReleaseMode"`
		UseLogger   bool `json:"UseLogger"`
		UseRecovery bool `json:"UseRecovery"`
	}
	ProfilingEnabled bool `json:"ProfilingEnabled"`
	StopTimeout      int  `json:"StopTimeout"`
}

type DBSettings struct {
	ConnectionString string
	LogMode          bool
	MaxOpenConns     int
	MaxIdleConns     int
}

type SwaggerUIConfig struct {
	PageTitle   string
	Host        string
	Description string
}

func Init(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(file)
	cfg := new(Config)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
