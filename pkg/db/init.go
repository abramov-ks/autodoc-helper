package db

import "github.com/go-pg/pg/v10"

func GetConnection(db *DatabaseConfig) *pg.DB {

	dbConnection := pg.Connect(&pg.Options{
		Addr:     db.Url,
		User:     db.Username,
		Password: db.Password,
		Database: db.Database,
	})

	return dbConnection
}
