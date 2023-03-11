package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application
// the values are read by viper from a config file
type Config struct {
	// APP SERVER CONFIG
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBSource          string `mapstructure:"DB_SOURCE"`
	MigrationURL      string `mapstructure:"MIGRATION_URL"`
	DBSourceTest      string `mapstructure:"DB_SOURCE_TEST"`
	ServerAddress     string `mapstructure:"SERVER_ADDRESS"`
	SmtpHost          string `mapstructure:"SMTP_HOST"`
	SmtpPort          string `mapstructure:"SMTP_PORT"`
	ScoringServiceURL string `mapstructure:"SCORING_SERVICE_URL"`

	// APP LOGIC CONFIG
	AccessTokenDuration                time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	AccessTokenDurationTest            time.Duration `mapstructure:"ACCESS_TOKEN_DURATION_TEST"`
	ClientURL                          string        `mapstructure:"CLIENT_URL"`
	PasswordResetTokenDuration         time.Duration `mapstructure:"PASSWORD_RESET_TOKEN_DURATION"`
	PasswordResetTokenDurationTest     time.Duration `mapstructure:"PASSWORD_RESET_TOKEN_DURATION_TEST"`
	EmailVerificationTokenDuration     time.Duration `mapstructure:"EMAIL_VERIFICATION_TOKEN_DURATION"`
	EmailVerificationTokenDurationTest time.Duration `mapstructure:"EMAIL_VERIFICATION_TOKEN_DURATION_TEST"`
	MinFee                             float64       `mapstructure:"MIN_FEE"`
	MaxBioLenght                       int           `mapstructure:"MAX_BIO_LENGTH"`
	MaxDescriptionLength               int           `mapstructure:"MAX_DESCRIPTION_LENGTH"`
	MaxSubscriptionMessageLength       int           `mapstructure:"MAX_SUBSCRIPTION_MESSAGE_LENGTH"`

	// CREDENTIALS
	TokenKey        string `mapstructure:"TOKEN_KEY"`
	TokenKeyTest    string `mapstructure:"TOKEN_KEY_TEST"`
	NoReplyEmail    string `mapstructure:"NO_REPLY_EMAIL"`
	NoReplyPassword string `mapstructure:"NO_REPLY_PASSWORD"`
	CloudinaryURL   string `mapstructure:"CLOUDINARY_URL"`
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
