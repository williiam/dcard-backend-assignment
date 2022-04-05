package main

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/williiam/dcard-backend-assignment/internal/app/handler"
	db "github.com/williiam/dcard-backend-assignment/internal/app/DB"
	dao "github.com/williiam/dcard-backend-assignment/internal/app/dao"
	service "github.com/williiam/dcard-backend-assignment/internal/app/service"
)

var port string = "8080"

func main() {
	fmt.Println("Server is running on port " + port)
	// sleep for a while to wait for the database to be ready
	time.Sleep(time.Second * 3)

	sqlURLRepo := dao.SQLURLRepo{DB: db.InitDB()}
	redisService := service.RedisServiceClient{RedisClient: service.InitService() }
	h := handler.URLHandler{
		Port:    port,
		URLRepo: sqlURLRepo,
		RedisService: redisService,
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return h.Home(c)
	})

	e.POST("/shorten", func(c echo.Context) error {
		return h.CreateShortURL(c)
	})

	e.GET("/:shortURL", func(c echo.Context) error {
		return h.HandleShortURLRedirect(c)
	})

	e.GET("/url/:shortURL", func(c echo.Context) error {
		return h.HandleShortURLDetail(c)
	})

	e.Logger.Fatal(e.Start(":" + port))
}
