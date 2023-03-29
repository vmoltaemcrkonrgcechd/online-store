package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/usecase"
	"github.com/vmoltaemcrkonrgcechd/online-store/pkg/sessionstore"
)

type categoryRoutes struct {
	uc usecase.CategoryUseCase
}

func withCategoryRoutes(app *fiber.App,
	uc usecase.CategoryUseCase, store *sessionstore.SessionStore) {
	r := categoryRoutes{uc}
	app.Get("/categories", store.Check("buyer"), r.all)
}

//@tags категории
//@success 200 {array} entities.Category
//@router /categories [get]
func (r categoryRoutes) all(ctx *fiber.Ctx) error {
	categories, err := r.uc.All()
	if err != nil {
		return err
	}

	return ctx.JSON(categories)
}
