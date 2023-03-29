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

func withUserRoutes(app *fiber.App,
	uc usecase.UserUseCase,
	store *sessionstore.SessionStore) {
	r := userRoutes{uc, store}
	app.Post("/sign-up", r.signUp)
	app.Post("/sign-in", r.signIn)
	app.Get("/users/:id", store.Check("buyer"), r.userByID)
	app.Get("/meDetails", store.Check("buyer"), r.meDetails)
}

//@tags пользователи
//@param user body entities.UserDTO true "пользователь"
//@router /sign-up [post]
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

//@tags пользователи
//@param credentials body entities.Credentials true "реквизиты для входа"
//@router /sign-in [post]
func (r userRoutes) signIn(ctx *fiber.Ctx) error {
	var credentials entities.Credentials

	err := ctx.BodyParser(&credentials)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "неверный json")
	}

	var user entities.User
	if user, err = r.uc.SignIn(credentials); err != nil {
		return err
	}

	if err = r.store.Set(ctx, user.ID, user.Role); err != nil {
		return fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при настройке сеанса")
	}

	return ctx.JSON(user)
}

//@tags пользователи
//@param id path string true "идентификатор"
//@router /users/{id} [get]
func (r userRoutes) userByID(ctx *fiber.Ctx) error {
	user, err := r.uc.UserByID(ctx.Params("id"))
	if err != nil {
		return err
	}

	return ctx.JSON(user)
}

//@tags пользователи
//@router /meDetails [get]
func (r userRoutes) meDetails(ctx *fiber.Ctx) error {
	id, _ := r.store.ID(ctx)

	user, err := r.uc.UserByID(id)
	if err != nil {
		return err
	}

	return ctx.JSON(user)
}
