package main

import (
	"github.com/boltdb/bolt"

	service "github.com/dan-compton/gofigure/service"
	_ "github.com/lib/pq"
	log "github.com/opsee/logrus"
	"github.com/spf13/viper"
)

func main() {
	viper.SetEnvPrefix("gofigure_")
	viper.AutomaticEnv()

	viper.SetDefault("grpc_addr", ":9118")
	viper.SetDefault("db_addr", "gofigure.db")
	viper.SetDefault("log_level", "debug")
	logLevelStr := viper.GetString("log_level")
	logLevel, err := log.ParseLevel(logLevelStr)
	if err != nil {
		log.WithError(err).Error("Could not parse log level, using default.")
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)

	db, err := bolt.Open(viper.GetString("db_addr"), 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Info("connected to boltdb")
	service.New(db).Start(viper.GetString("grpc_addr"))
}
