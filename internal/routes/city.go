package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/usecase"
)

type cityRoutes struct {
	uc usecase.CityUseCase
}

func withCityRoutes(app *fiber.App, uc usecase.CityUseCase) {
	r := cityRoutes{uc}
	app.Get("/cities", r.all)
}

//	@tags		города
//	@success	200	{array}	entities.City
//	@router		/cities [get]
func (r cityRoutes) all(ctx *fiber.Ctx) error {
	cities, err := r.uc.All()
	if err != nil {
		return err
	}

	return ctx.JSON(cities)
}
