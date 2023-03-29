package repo

import (
	"database/sql"
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

func (r UserRepo) UserByCredentials(credentials entities.Credentials) (
	user entities.User, err error) {
	if err = r.Sq.Select("user_id", "username", "city_id", "city_name", "image_path", "role").
		From("\"user\"").
		Join("city USING (city_id)").
		Where("username = ? AND password = ?", credentials.Username, credentials.Password).
		QueryRow().Scan(&user.ID, &user.Username, &user.City.ID,
		&user.City.Name, &user.ImagePath, &user.Role); err != nil {
		if err == sql.ErrNoRows {
			return user, fiber.NewError(http.StatusBadRequest,
				"неправильный логин или пароль")
		}

		log.Println(err)
		return user, fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при входе в систему")
	}

	return user, nil
}

func (r UserRepo) UserByID(id string) (user entities.User, err error) {
	if err = r.Sq.Select("user_id", "username", "city_id", "city_name", "image_path", "role").
		From("\"user\"").
		Join("city USING (city_id)").
		Where("user_id = ?", id).
		QueryRow().Scan(&user.ID, &user.Username, &user.City.ID,
		&user.City.Name, &user.ImagePath, &user.Role); err != nil {
		log.Println(err)

		if err == sql.ErrNoRows {
			return user, fiber.NewError(http.StatusBadRequest,
				"такого пользователя не существует")
		}

		return user, fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при получении пользователя")
	}

	return user, nil
}
