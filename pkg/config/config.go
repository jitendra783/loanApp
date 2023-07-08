package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var (
	config *viper.Viper
)

const (
	USERID        = "userId"
	REQUESTID     = "requestID"
	UCC           = "ucc"
	AUTHORIZATION = "Authorization"
	XLENGTH       = "X-Length"
	SCOPE         = "scope"
)

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Load(env string, configPaths ...string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	config.AddConfigPath("../app/")
	config.AddConfigPath(".")
	if len(configPaths) != 0 {
		for _, path := range configPaths {
			config.AddConfigPath(path)
		}
	}
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file", err)
		return
	}
	if env == "server" {
		for _, v := range config.AllKeys() {
			if strings.ToLower(v) == "version" {
				continue // skipping first line
			}
			key := config.GetString(v)
			key = strings.ReplaceAll(key, "$", "")
			if ev, ok := os.LookupEnv(key); ok {
				config.Set(v, ev)
			} else {
				log.Fatal("Env value for key [", key, "] is missing")
			}

		}
		log.Println("application running with server.yaml")
	} else {
		log.Println("application running locally")
	}

}

func GetConfig() *viper.Viper {
	return config
}
