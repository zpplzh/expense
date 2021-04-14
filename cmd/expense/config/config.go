package configs

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type myConfig struct {
	PortServer string `env:"PORT"`
	DBName     string `env:"DBNAME"`
	DBUsername string `env:"DBUSER"`
	DBPass     string `env:"DBPASS"`
}

func GetEnv() *myConfig {
	var conf myConfig
	ctx := context.Background()

	if err := envconfig.Process(ctx, &conf); err != nil {
		log.Fatal(err)
	}

	return &conf

}
