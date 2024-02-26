package main

import (
	"pbi-final-task-go-api/database"
	"pbi-final-task-go-api/initializers"
	"pbi-final-task-go-api/router"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoanEnvVariables()
	database.ConnectToDB()
	database.Migrate()
}

func main() {
	r := gin.Default()
	router.UserRoutes(r)
	router.PhotoRoutes(r)
	r.Run()
}
