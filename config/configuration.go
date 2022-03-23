package config

import "os"

const (
	LOCAL       = "local"
	DEVELOPMENT = "development"
)

// ENVIRONMENT
const ENVIRONMENT string = LOCAL //LOCAL, DEVELOPMENT, PRODUCTION

var env = map[string]map[string]string{
	// LOCAL ENVIRONMENT CONFIGURATION
	"local": {
		"APP_NAME":   "FSE",
		"SECRET_KEY": "SECRET",
	},
}

// CONFIG : BLOBAL CONFIGURATION
var CONFIG = env[ENVIRONMENT]

// GetEnv : function for ENVIRONMENT LOOKUP
func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func InitConfig() {
	for key := range CONFIG {
		CONFIG[key] = GetEnv(key, CONFIG[key])
		os.Setenv(key, CONFIG[key])
	}
}
