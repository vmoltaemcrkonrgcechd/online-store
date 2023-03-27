package repo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/entities"
	"github.com/vmoltaemcrkonrgcechd/online-store/pkg/pg"
	"log"
	"net/http"
)

type UserRepo struct {
	*pg.PG
}

func NewUserRepo(pq *pg.PG) UserRepo {
	return UserRepo{pq}
}

func (r UserRepo) UserExists(username string) (ok bool, err error) {
	if err = r.Sq.Select("count(*) > 0").
		From("\"user\"").
		Where("username = ?", username).QueryRow().Scan(&ok); err != nil {
		log.Println(err)
		return false, fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при проверке присутствия пользователя")
	}

	return ok, nil
}

func (r UserRepo) SaveUser(user entities.UserDTO) (id string, err error) {
	if err = r.Sq.Insert("\"user\"").
		Columns("username", "password", "city_id", "image_path", "role").
		Values(user.Username, user.Password, user.CityID, user.ImagePath, user.Role).
		Suffix("RETURNING user_id").QueryRow().Scan(&id); err != nil {
		log.Println(err)
		return "", fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при сохранении пользователя в базе данных")
	}

	return id, nil
}
