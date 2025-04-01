package main

import (
	"fmt"
	"os"

	"github.com/CloudyKit/jet/v6"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lassejlv/go-app-railway/utils"
)

func main() {
	godotenv.Load()

	PORT := os.Getenv("PORT")

	app := echo.New()

	views := jet.NewSet(
		jet.NewOSFileSystemLoader("views"),
		jet.InDevelopmentMode(),
	)

	app.Renderer = utils.NewRenderer(views)

	app.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "${time_rfc3339} ${method} ${uri} ${status} ${latency_human}\n",
		},
	))

	app.GET("/", func(c echo.Context) error {
		return c.Render(200, "home.jet", map[string]interface{ any }{})
	})

	app.GET("/register", func(c echo.Context) error {
		return c.Render(200, "register.jet", map[string]interface{ any }{})
	})

	app.Static("/", "./public")

	app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", PORT)))
}
