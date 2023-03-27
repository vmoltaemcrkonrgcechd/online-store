package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/vmoltaemcrkonrgcechd/online-store/docs"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/usecase"
	"github.com/vmoltaemcrkonrgcechd/online-store/pkg/sessionstore"
)

func WithRouter(app *fiber.App,
	cityUC usecase.CityUseCase,
	imageUC usecase.ImageUseCase,
	userUC usecase.UserUseCase,
	store *sessionstore.SessionStore,
) {
	app.Get("/swagger-ui/*", swagger.New(swagger.ConfigDefault))

	withCityRoutes(app, cityUC)
	withImageRoutes(app, imageUC)
	withUserRoutes(app, userUC, store)
}
