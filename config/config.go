package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application
// the values are read by viper from a config file
type Config struct {
	DBDriver                       string        `mapstructure:"DB_DRIVER"`
	DBSource                       string        `mapstructure:"DB_SOURCE"`
	DBSourceTest                   string        `mapstructure:"DB_SOURCE_TEST"`
	ServerAddress                  string        `mapstructure:"SERVER_ADDRESS"`
	TokenKey                       string        `mapstructure:"TOKEN_KEY"`
	TokenKeyTest                   string        `mapstructure:"TOKEN_KEY_TEST"`
	AccessTokenDuration            time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	AccessTokenDurationTest        time.Duration `mapstructure:"ACCESS_TOKEN_DURATION_TEST"`
	ClientURL                      string        `mapstructure:"CLIENT_URL"`
	PasswordResetTokenDuration     time.Duration `mapstructure:"PASSWORD_RESER_TOKEN_DURATION"`
	PasswordResetTokenDurationTest time.Duration `mapstructure:"PASSWORD_RESER_TOKEN_DURATION_TEST"`
	NoReplyEmail                   string        `mapstructure:"NO_REPLY_EMAIL"`
	NoReplyPassword                string        `mapstructure:"NO_REPLY_PASSWORD"`
	SmtpHost                       string        `mapstructure:"SMTP_HOST"`
	SmtpPort                       string        `mapstructure:"SMTP_PORT"`
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
