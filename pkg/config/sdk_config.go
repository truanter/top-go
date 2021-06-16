package config

type SDKConfig struct {
	Config
	IsSandBox     bool
	IsHttp        bool
	IsDebug       bool
	DebugFilePath string
}
