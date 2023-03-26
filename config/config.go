package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v7"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	IsProduction bool          `env:"PRODUCTION"`
	Hosts        []string      `env:"HOSTS" envSeparator:":"`
	Duration     time.Duration `env:"DURATION"`
}

func InitConfig(logger log.Entry) Config {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", err)
	return cfg

}
