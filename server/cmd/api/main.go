package main

import(
	"log"
	"github.com/code-guy21/PingUp/server/config"
	"github.com/code-guy21/PingUp/server/handlers"
	"github.com/code-guy21/PingUp/server/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main(){
	//Load environment variables
	config.LoadConfig()	

	 // Print environment variables for debugging
	 log.Println("DB_HOST:", config.GetEnv("DB_HOST", "localhost"))
	 log.Println("DB_USER:", config.GetEnv("DB_USER", "user"))
	 log.Println("DB_PASSWORD:", config.GetEnv("DB_PASSWORD", "password"))
	 log.Println("DB_NAME:", config.GetEnv("DB_NAME", "pingup"))

	dsn := "host=" + config.GetEnv("DB_HOST", "localhost") +
		" user=" + config.GetEnv("DB_USER", "postgres") +
		" password=" + config.GetEnv("DB_PASSWORD", "") +
		" dbname=" + config.GetEnv("DB_NAME", "pingupDB") +
		" port=5432 sslmode=disable"

	log.Println("DSN", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	repositories.DB = db

	r := gin.Default()

	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)

	r.Run(":8080")
}
