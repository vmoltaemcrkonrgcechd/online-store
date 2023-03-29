package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/usecase"
	"github.com/vmoltaemcrkonrgcechd/online-store/pkg/sessionstore"
)

type colorRoutes struct {
	uc usecase.ColorUseCase
}

func withColorRoutes(app *fiber.App, uc usecase.ColorUseCase, store *sessionstore.SessionStore) {
	r := colorRoutes{uc}
	app.Get("/colors", store.Check("buyer"), r.all)
}

//@tags цвета
//@success 200 {array} entities.Color
//@router /colors [get]
func (r colorRoutes) all(ctx *fiber.Ctx) error {
	colors, err := r.uc.All()
	if err != nil {
		return err
	}

	return ctx.JSON(colors)
}
