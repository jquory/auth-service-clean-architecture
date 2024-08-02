package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
	"sync"
)

type (
	Config struct {
		Server *Server `mapstructure:"server"`
		DB     *DB     `mapstructure:"database"`
		Jwt    *Jwt    `mapstructure:"jwt"`
	}

	Server struct {
		Port int `mapstructure:"port"`
	}

	DB struct {
		Host              string `mapstructure:"host"`
		Port              int    `mapstructure:"port"`
		User              string `mapstructure:"user"`
		Password          string `mapstructure:"password"`
		DBName            string `mapstructure:"dbname"`
		SslMode           string `mapstructure:"sslmode"`
		MaxConnectionPool int    `mapstructure:"max_connection_pool"`
		MaxIdleTime       int    `mapstructure:"max_idle_time"`
		MaxLifetimePool   int    `mapstructure:"max_lifetime_pool"`
	}

	Jwt struct {
		JwtSecret string `mapstructure:"jwt_secret"`
		JwtExpire int    `mapstructure:"jwt_expire_in_minutes"`
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Sprintf("Error read config %v", err))
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(fmt.Sprintf("Error unmarshall %v", err))
		}
	})

	return configInstance
}
