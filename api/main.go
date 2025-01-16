package main

import (
	//"crossx/database"
	"crossx/routes"
)

// @title		CrossX API
// @version		1.0
// @description	La classe cette API
// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	r := routes.New()
	router := r.SetupRouter()
	router.Run(":8080")

	//end
	//	database.CloseAllDBs()
}