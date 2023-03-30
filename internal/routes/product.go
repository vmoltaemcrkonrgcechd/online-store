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
	app.Get("/products", store.Check("buyer"), r.all)
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

//	@tags		продукты
//@param limit query int false "ограничение"
//@param offset query int false "смещение"
//@param orderBy query string false "сортировка"
//	@success	200	{object}	entities.AllProductsDTO
//	@router		/products [get]
func (r productRoutes) all(ctx *fiber.Ctx) (err error) {
	var (
		qp     entities.AllProductsQP
		result entities.AllProductsDTO
	)

	if err = ctx.QueryParser(&qp); err != nil {
		return fiber.NewError(http.StatusBadRequest,
			"недействительные параметры")
	}

	if result, err = r.uc.All(qp); err != nil {
		return err
	}

	return ctx.JSON(result)
}
