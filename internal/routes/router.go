package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/vmoltaemcrkonrgcechd/online-store/docs"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/usecase"
)

func WithRouter(app *fiber.App, cityUC usecase.CityUseCase) {
	app.Get("/swagger-ui/*", swagger.New(swagger.ConfigDefault))

	withCityRoutes(app, cityUC)
}
