package database

import (
	"fmt"
	"os"
	"time"

	"github.com/rafaelbreno/work-at-olist/cmd/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Conn       *gorm.DB
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbName     string
	dbConnStr  string
}

func (pg *Postgres) SetPostgres() {
	pg.setCredentials()

	pg.setConn()
}

func (pg *Postgres) setCredentials() {
	pg.dbName = os.Getenv("POSTGRES_DB")
	pg.dbHost = os.Getenv("POSTGRES_HOST")
	pg.dbPort = os.Getenv("POSTGRES_PORT")
	pg.dbUser = os.Getenv("POSTGRES_USER")
	pg.dbPassword = os.Getenv("POSTGRES_PASSWORD")
	pg.dbConnStr = fmt.Sprintf(`host=%s 
								port=%s 
								user=%s 
								password=%s 
								dbname=%s 
								sslmode=disable`,
		pg.dbHost, pg.dbPort, pg.dbUser, pg.dbPassword, pg.dbName)
}

func (pg *Postgres) setConn() {
	var conn *gorm.DB
	var err error
	for {
		conn, err = gorm.Open(postgres.Open(pg.dbConnStr), &gorm.Config{})
		if err == nil {
			break
		}
		logger.Error(fmt.Sprintf("Error: %s\nRetrying in 5 seconds", err.Error()))
		time.Sleep(time.Second * 5)
	}
	pg.Conn = conn
}
