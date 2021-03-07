package database

// Setting Database connections

var PGConn Postgres

func init() {
	PGConn.SetPostgres()
}
