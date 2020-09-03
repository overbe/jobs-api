package main

import (
	"jobs/server"
	"jobs/server/config"
	"log"
	"strings"

	"github.com/spf13/viper"
)

// init reads in config file and ENV variables if set.
func initViper() {
	viper.SetEnvPrefix("JOBS") // all jobs environment variables must be prefixed with JOBS
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetConfigType("yml")

	env := viper.GetString("ENV")
	if env != "" {
		viper.SetConfigName(env + "-" + "jobs.yml")
	} else {
		viper.SetConfigName("jobs.yml")
	}

	viper.AddConfigPath("../config/")

	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}

	log.Println("Used config is: " + viper.ConfigFileUsed())
}

func main() {
	initViper()
	builder := config.Init(viper.GetViper())
	err := server.Run(builder)
	if err != nil {
		log.Panic(err)
	}
}
