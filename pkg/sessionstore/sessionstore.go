package sessionstore

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"log"
	"net/http"
)

type SessionStore struct {
	*session.Store
}

func New() *SessionStore {
	return &SessionStore{session.New()}
}

func (s *SessionStore) Set(ctx *fiber.Ctx, id, role string) error {
	sess, err := s.Get(ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	sess.Set("id", id)
	sess.Set("role", role)

	if err = sess.Save(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *SessionStore) Check(expected string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		sess, err := s.Get(ctx)
		if err != nil {
			return fiber.NewError(http.StatusUnauthorized,
				"произошла ошибка при получении сеанса")
		}

		role, ok := sess.Get("role").(string)
		if !ok || role == "" {
			return fiber.NewError(http.StatusUnauthorized,
				"произошла ошибка при получении сеанса")
		}

		switch role {
		case "seller":
			return ctx.Next()
		case "buyer":
			if expected == "seller" {
				return fiber.NewError(http.StatusForbidden,
					"недостаточно прав")
			}

			return ctx.Next()
		default:
			return fiber.NewError(http.StatusUnauthorized,
				"неизвестная роль")
		}
	}
}

func (s *SessionStore) ID(ctx *fiber.Ctx) (string, error) {
	sess, err := s.Get(ctx)
	if err != nil {
		return "", err
	}

	return sess.Get("id").(string), nil
}
