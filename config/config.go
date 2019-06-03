package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

//struct of db connection
type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}

type GatewayConfig struct {
	Grpcport int
	Host     string
	Port     int
}

//config
type Config struct {
	Datasource  DBConfig
	RestGateway GatewayConfig
}

//load
func Load() {
	//configPath := flag.String("config", "./config.yaml", "path to the configuration file")
	//flag.Parse()
	viper.SetConfigFile("./config.yaml")
	// Searches for config file in given paths and read it
	if err := viper.ReadInConfig(); err != nil {
		log.Panic("Error reading config file: ", err)
	}
}

//db
func Db() DBConfig {
	// Instead of reading keys one by one we can extract sub-tree
	// using Sub and decode it in struct using Unmarshal
	sub := viper.Sub("database")
	var db DBConfig

	err := sub.Unmarshal(&db)
	if err != nil {
		log.Panic("unable to decode into struct: ", err)
	}
	return db
}

func Gateway() GatewayConfig {
	sub := viper.Sub("restgateway")
	var gw GatewayConfig

	err := sub.Unmarshal(&gw)
	if err != nil {
		log.Panic("unable to decode into struct: ", err)
	}
	return gw
}
