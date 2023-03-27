package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/entities"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/usecase"
	"github.com/vmoltaemcrkonrgcechd/online-store/pkg/sessionstore"
	"net/http"
)

type userRoutes struct {
	uc    usecase.UserUseCase
	store *sessionstore.SessionStore
}

//@tags пользователи
//@param user body entities.UserDTO true "пользователь"
//@router /sign-up [post]
func withUserRoutes(app *fiber.App,
	uc usecase.UserUseCase,
	store *sessionstore.SessionStore) {
	r := userRoutes{uc, store}
	app.Post("/sign-up", r.signUp)
}

func (r userRoutes) signUp(ctx *fiber.Ctx) error {
	var user entities.UserDTO

	err := ctx.BodyParser(&user)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "неверный json")
	}

	var id string
	if id, err = r.uc.SingUp(user); err != nil {
		return err
	}

	if err = r.store.Set(ctx, id, user.Role); err != nil {
		return fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при настройке сеанса")
	}

	return ctx.Status(http.StatusCreated).SendString(id)
}
