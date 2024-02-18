package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// Run запускает приложение
func (app *Application) Run() {
    r := gin.Default()

    // эндпоинт получения всех монет
    r.GET("/api/v1/currencies", app.Handler.GetAllCoins)
    // эндпоинт получения монеты по валюте
    r.GET("/api/v1/currencies/:currency", app.Handler.GetCoinByCurrency)

    addr := fmt.Sprintf("%s:%d", app.Config.ServiceHost, app.Config.ServicePort)
    r.Run(addr)
    log.Println("Server down")
}
