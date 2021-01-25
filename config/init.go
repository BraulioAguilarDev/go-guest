package config

import (
	"github.com/spf13/viper"
)

// Config var
var Config Configuration

// Configuration struct
type Configuration struct {
	Port         string
	Dialect      string
	DatabaseURI  string
	FirebaseJSON string
}

// Init func
func Init() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Config); err != nil {
		return err
	}

	return nil
}
