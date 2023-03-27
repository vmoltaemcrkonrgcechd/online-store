package usecase

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gofiber/fiber/v2"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/entities"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/usecase/repo"
	"net/http"
)

type UserUseCase struct {
	repo repo.UserRepo
}

func NewUserUseCase(repo repo.UserRepo) UserUseCase {
	return UserUseCase{repo}
}

func (uc UserUseCase) SingUp(user entities.UserDTO) (string, error) {
	// TODO добавить проверку пользователя.

	//проверить, есть ли пользователь с таким именем.
	ok, err := uc.repo.UserExists(user.Username)
	if err != nil {
		return "", err
	}

	if ok {
		return "", fiber.NewError(http.StatusBadRequest,
			"пользователь с таким именем уже существует")
	}

	//получить хэш пароля.
	user.Password = uc.hash(user.Password)

	//сохранить пользователя в базу данных.
	var id string
	if id, err = uc.repo.SaveUser(user); err != nil {
		return "", err
	}

	return id, nil
}

func (uc UserUseCase) hash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}
