package server

import (
	"encoding/json"
	"fmt"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/cmd/server/factories/controllers"
	inMemoryDB "github.com/elissonalvesilva/interview-hash/my-ecommerce/infrastructure/db/in-memory"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/infrastructure/logger/logrus"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/shared/file"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
	"time"
)

type App struct {
	httpServer *http.Server
}

func (a *App) Run(application string) {
	fmt.Println(application+" Is running in port", a.httpServer.Addr)
	err := a.httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func NewApp(port int) *App {
	db := file.ReadDBFile(os.Getenv("DB_IN_MEMORY_PATHNAME"))
	var blackFridayDate time.Time

	database := inMemoryDB.NewInMemoryDatabase(db)

	blackFridayDay, _ := strconv.Atoi(os.Getenv("BLACK_FRIDAY_DAY"))
	blackFridayMonth, _ := strconv.Atoi(os.Getenv("BLACK_FRIDAY_MONTH"))
	blackFridayDate = time.Date(time.Now().Year(), time.Month(blackFridayMonth), blackFridayDay, 00, 00, 00, 00, time.UTC)

	router := mux.NewRouter()

	logger := logrus.NewLogger()
	controller := controllers.MakeCheckoutController(database, blackFridayDate)

	router.HandleFunc("/checkout", controller.CheckoutProductsController).Methods("POST")
	router.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		json.NewEncoder(writer).Encode("ok")
		return
	}).Methods("GET")

	return &App{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%v", port),
			Handler: logger.WithLogging(router),
		},
	}
}
