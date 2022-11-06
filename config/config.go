package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application
// the values are read by viper from a config file
type Config struct {
	DBDriver                string        `mapstructure:"DB_DRIVER"`
	DBSource                string        `mapstructure:"DB_SOURCE"`
	DBSourceTest            string        `mapstructure:"DB_SOURCE_TEST"`
	ServerAddress           string        `mapstructure:"SERVER_ADDRESS"`
	TokenKey                string        `mapstructure:"TOKEN_KEY"`
	TokenKeyTest            string        `mapstructure:"TOKEN_KEY_TEST"`
	AccessTokenDuration     time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	AccessTokenDurationTest time.Duration `mapstructure:"ACCESS_TOKEN_DURATION_TEST"`
}

// LoadConfig reads configuration from file
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
