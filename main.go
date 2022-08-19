package main

import (
	config "articles/config"
	migration "articles/databases"
)

func main() {
	// DB connection databases
	dbConn := config.InitDB()
	// Migration tables
	migration.Migration(dbConn)

}
