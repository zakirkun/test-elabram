package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/zakirkun/test-elabram/infrastructure"
	"github.com/zakirkun/test-elabram/internal/pkg/config"
	"github.com/zakirkun/test-elabram/internal/pkg/database"
	"github.com/zakirkun/test-elabram/internal/pkg/webserver"
	"github.com/zakirkun/test-elabram/router"
)

var configFile *string

func init() {
	// Read Config
	configFile = flag.String("c", "config.toml", "configuration file")
	flag.Parse()
}

func main() {
	setConfig()

	infra := infrastructure.NewInfrastructure(SetDatabase(), SetWebServer())
	infra.Database()
	infra.WebServer()
}

func setConfig() {
	cfg := config.NewConfig(*configFile)
	if err := cfg.Initialize(); err != nil {
		log.Fatalf("Error reading config : %v", err)
		os.Exit(1)
	}
}

func SetDatabase() database.DBModel {
	return database.DBModel{
		ServerMode:   config.GetString("server.mode"),
		Driver:       config.GetString("database.db_driver"),
		Host:         config.GetString("database.db_host"),
		Port:         config.GetString("database.db_port"),
		Name:         config.GetString("database.db_name"),
		Username:     config.GetString("database.db_username"),
		Password:     config.GetString("database.db_password"),
		MaxIdleConn:  config.GetInt("pool.conn_idle"),
		MaxOpenConn:  config.GetInt("pool.conn_max"),
		ConnLifeTime: config.GetInt("pool.conn_lifetime"),
	}
}

func SetWebServer() webserver.ServerContext {
	return webserver.ServerContext{
		Host:         ":" + config.GetString("server.port"),
		Handler:      router.NewRouter(),
		ReadTimeout:  time.Duration(config.GetInt("server.http_timeout")),
		WriteTimeout: time.Duration(config.GetInt("server.http_timeout")),
	}
}
