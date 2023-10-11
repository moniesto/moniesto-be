package config

type AppEnv string

const AppEnvLocal AppEnv = "LOCAL"
const AppEnvAlpha AppEnv = "ALPHA"
const AppEnvProd AppEnv = "PROD"

var SupportedAppEnv []AppEnv = []AppEnv{
	AppEnvLocal,
	AppEnvAlpha,
	AppEnvProd,
}
