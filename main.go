package main

import (
	"log"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	mongodb "github.com/danisbagus/simple-go-ddd/modules/mongodb"

	repo "github.com/danisbagus/simple-go-ddd/infrastructure/persistence"

	app "github.com/danisbagus/simple-go-ddd/application"

	handler "github.com/danisbagus/simple-go-ddd/interfaces"

	"github.com/danisbagus/simple-go-ddd/utils/config"
)

func main() {

	// get app config
	appConfig := config.GetAPPConfig()

	// init modules
	mongodb.Init(appConfig.MongoDatabase)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.Use(ServiceRequestTime)

	v1Router := e.Group("/v1")

	mongoDB, err := mongodb.GetDatabase()
	if err != nil {
		log.Fatal(err)
	}

	productRepo := repo.NewProductRepo(mongoDB)

	producApp := app.NewProductApp(productRepo)

	productHandler := handler.NewProductHandler(producApp)

	productkRoute := v1Router.Group("/product")
	productkRoute.POST("", productHandler.Insert)
	productkRoute.GET("", productHandler.List)
	productkRoute.GET("/:id", productHandler.View)
	productkRoute.PUT("/:id", productHandler.Update)
	productkRoute.DELETE("/:id", productHandler.Delete)

	//

	port := "5000"
	if err := e.Start(":" + port); err != nil {
		e.Logger.Info("Shutting down the server")
	}
}

func ServiceRequestTime(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request().Header.Set("X-App-RequestTime", time.Now().Format(time.RFC3339))
		return next(c)
	}
}
