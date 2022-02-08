package main

import (
	"FinalProjectGolangH8/config"
	"FinalProjectGolangH8/routes"
	"log"

	"github.com/joho/godotenv" // swagger embed files
	// gin-swagger middleware
	"github.com/swaggo/swag/example/basic/docs"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	docs.SwaggerInfo.Title = "MyGram"
	docs.SwaggerInfo.Description = "Persyaratan untuk menyelesaikan hacktive 8 course Golang, serta menjadi Portofolio pribadi"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	r := routes.SetupRouter(db)
	r.Run()
}
