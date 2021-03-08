package domain

import "github.com/rafaelbreno/work-at-olist/cmd/database"

func init() {
	// Migration
	database.
		PGConn.
		Conn.
		AutoMigrate(&Author{})
}
