package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/entities"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/usecase"
	"github.com/vmoltaemcrkonrgcechd/online-store/pkg/sessionstore"
	"net/http"
)

type productRoutes struct {
	uc usecase.ProductUseCase
}

func withProductRoutes(app *fiber.App,
	uc usecase.ProductUseCase, store *sessionstore.SessionStore) {
	r := productRoutes{uc}
	app.Post("/products", store.Check("seller"), r.add)
}

//	@tags	продукты
//	@param	product	body	entities.ProductDTO	true	"продукт"
//	@router	/products [post]
func (r productRoutes) add(ctx *fiber.Ctx) error {
	var product entities.ProductDTO

	err := ctx.BodyParser(&product)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "неверный json")
	}

	var id string
	if id, err = r.uc.Add(product); err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).SendString(id)
}
