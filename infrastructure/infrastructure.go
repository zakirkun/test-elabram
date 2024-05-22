package infrastructure

import (
	"log"
	"os"

	"github.com/zakirkun/test-elabram/app/domain/model"
	"github.com/zakirkun/test-elabram/internal/globals"
	"github.com/zakirkun/test-elabram/internal/pkg/database"
	"github.com/zakirkun/test-elabram/internal/pkg/webserver"
)

type Infrastructure interface {
	Database()
	WebServer()
}

type infrastructureContext struct {
	database database.DBModel
	server   webserver.ServerContext
}

func NewInfrastructure(
	database database.DBModel,
	server webserver.ServerContext,
) Infrastructure {
	return infrastructureContext{
		database: database,
		server:   server,
	}
}

func (i infrastructureContext) Database() {
	db, err := i.database.OpenDB()
	if err != nil {
		log.Fatalf("Error Init Database: %v", err)
		os.Exit(1)
	}

	globals.DB = &i.database

	migrateErrr := db.AutoMigrate(&model.Company{}, &model.Admin{}, &model.Employe{}, &model.RequestClaim{})
	if migrateErrr != nil {
		log.Printf("LOG Migrate Table: %v", err)
	}
}

func (i infrastructureContext) WebServer() {
	i.server.Run()
}
