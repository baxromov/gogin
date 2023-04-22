package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		PORT string `mapstructure:"PORT"`
	}
	Database struct {
		DSN string `mapstructure:"DSN"`
	}
}

func ReadConfig() (*Config, error) {
	viper.SetConfigName(".env")
    viper.AddConfigPath(".")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Failed to read configuration file: %v", err)
    }

    // Unmarshal the configuration file into a struct
    if err := viper.Unmarshal(&Config); err != nil {
        log.Fatalf("Failed to unmarshal configuration file: %v", err)
    }

    return cfg, nil
}
