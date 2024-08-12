// main.go

package main

import (
	"rest-api/database"
	"rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.DbIn() // Initialize the database connection
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
