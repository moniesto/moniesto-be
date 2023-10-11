package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application
// the values are read by viper from a config file
type Config struct {
	// APP SERVER CONFIG
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	MigrationURL  string `mapstructure:"MIGRATION_URL"`
	DBSourceTest  string `mapstructure:"DB_SOURCE_TEST"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	SmtpHost      string `mapstructure:"SMTP_HOST"`
	SmtpPort      string `mapstructure:"SMTP_PORT"`

	// APP LOGIC CONFIG
	AppEnv                             string        `mapstructure:"APP_ENV"`
	MaintenanceMode                    bool          `mapstructure:"MAINTENANCE_MODE"`
	AccessTokenDuration                time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	AccessTokenDurationTest            time.Duration `mapstructure:"ACCESS_TOKEN_DURATION_TEST"`
	ClientURL                          string        `mapstructure:"CLIENT_URL"`
	PasswordResetTokenDuration         time.Duration `mapstructure:"PASSWORD_RESET_TOKEN_DURATION"`
	PasswordResetTokenDurationTest     time.Duration `mapstructure:"PASSWORD_RESET_TOKEN_DURATION_TEST"`
	EmailVerificationTokenDuration     time.Duration `mapstructure:"EMAIL_VERIFICATION_TOKEN_DURATION"`
	EmailVerificationTokenDurationTest time.Duration `mapstructure:"EMAIL_VERIFICATION_TOKEN_DURATION_TEST"`
	MinFee                             float64       `mapstructure:"MIN_FEE"`
	MaxFee                             float64       `mapstructure:"MAX_FEE"`
	OperationFeePercentage             float64       `mapstructure:"OPERATION_FEE_PERCENTAGE"`

	// CREDENTIALS
	TokenKey        string `mapstructure:"TOKEN_KEY"`
	TokenKeyTest    string `mapstructure:"TOKEN_KEY_TEST"`
	NoReplyEmail    string `mapstructure:"NO_REPLY_EMAIL"`
	NoReplyPassword string `mapstructure:"NO_REPLY_PASSWORD"`
	CloudinaryURL   string `mapstructure:"CLOUDINARY_URL"`

	BinanceApiKey    string `mapstructure:"BINANCE_API_KEY"`
	BinanceSecretKey string `mapstructure:"BINANCE_SECRET_KEY"`
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
	if err != nil {
		return Config{}, err
	}

	// check configs is valid
	err = config.Valid()
	if err != nil {
		return Config{}, err
	}

	config.Enhance()

	return
}
