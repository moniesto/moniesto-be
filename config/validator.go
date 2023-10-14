package config

import (
	"fmt"

	"github.com/moniesto/moniesto-be/util"
)

func (config *Config) Valid() error {
	// STEP: check app env
	if contains := util.Contains[AppEnv](SupportedAppEnv, AppEnv(config.AppEnv)); !contains {
		return fmt.Errorf("unsupported app env %s", config.AppEnv)
	}

	return nil
}

func (config *Config) IsProd() bool {
	return config.AppEnv == string(AppEnvProd)
}

func (config *Config) IsLocal() bool {
	return config.AppEnv == string(AppEnvLocal)
}

func (config *Config) Enhance() {

	// update the token key based on the env
	config.TokenKey = config.TokenKey + config.AppEnv
}
