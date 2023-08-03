package main

import (
	"fmt"
	"log"
	"os"
	"typing/controllers"
	"typing/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	errEnv := godotenv.Load()

	if errEnv != nil {
		panic(errEnv.Error())
	}
	database.Open()
	defer database.Close()

	gin.SetMode(os.Getenv("GIN_MODE"))
	router := controllers.InitRouter()

	if err := router.SetTrustedProxies(nil); err != nil {
		panic(err.Error())
	}
	if err := router.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		log.Fatal(err)
	}
}
