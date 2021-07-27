package server

import (
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/infrastructure/env/dotenv"
	"os"
	"strconv"
)

func Run() {
	dotenv.InitEnvironment()
	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))
	name := os.Getenv("APP_NAME")

	NewApp(port).Run(name)
}