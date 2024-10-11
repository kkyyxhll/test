package router

import (
	//"net/http"

	"exchangeapp/controllers"
	"exchangeapp/middlewares"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		fmt.Println("")
		fmt.Println("controllers.Login")
		auth.POST("/register", controllers.Register)
		fmt.Println("controllers.Register")
	}
	api := r.Group("/api")
	api.GET("exchangeRates", controllers.GetExchangeRates)
	api.Use(middlewares.AuthMiddleWare())
	{
		api.POST("/exchangeRates", controllers.CreateExchangeRate)
	}
	return r
}
