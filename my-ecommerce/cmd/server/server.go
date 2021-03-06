package server

import (
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/infrastructure/env/dotenv"
	"log"
	"os"
	"strconv"
	"time"
)

func Run() {
	dotenv.InitEnvironment()

	location, errToGetTimezone := time.LoadLocation(os.Getenv("TIMEZONE"))
	if errToGetTimezone != nil {
		log.Fatal(errToGetTimezone)
		return
	}
	time.Local = location
	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))
	name := os.Getenv("APP_NAME")

	NewApp(port).Run(name)
}