package main

import (
	"fmt"
	"log"
	"os"
	"test/database"
	_ "test/docs"
	"test/pkg/mysql"
	"test/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go + Gin SWAG
// @version 1.0

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:5000
// @BasePath /api/v1
// @query.collection.format multi
func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No ENV file found")
	}
	// mysql database init
	mysql.DatabaseInit()
	// migration
	database.RunMigrate()
	r := gin.Default()

	// sub route API
	api := r.Group("/api/v1")
	routes.RoutesInit(api)

	// cors
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"X-Requested-With", "Content-Type", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	r.Use(cors.New(config))

	var port = os.Getenv("PORT")

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fmt.Println("test Running on Port", port)
	r.Run(":" + port)
}
