package main

import (
	factory "articles/app"
	middlewares "articles/app/middlewares"
	config "articles/config"
	migration "articles/databases"
	routes "articles/routes"
)

func main() {
	// DB connection databases
	dbConn := config.InitDB()
	// Migration tables
	migration.Migration(dbConn)
	//routes
	presenter := factory.InitFactory(dbConn)
	e := routes.New(presenter)

	middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}
