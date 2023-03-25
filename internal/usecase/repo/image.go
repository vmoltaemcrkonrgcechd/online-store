package repo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmoltaemcrkonrgcechd/online-store/pkg/pg"
	"log"
	"net/http"
)

type ImageRepo struct {
	*pg.PG
}

func NewImageRepo(pq *pg.PG) ImageRepo {
	return ImageRepo{pq}
}

func (r ImageRepo) SaveImagePaths(imagePaths []string) error {
	query := r.Sq.Insert("image").
		Columns("image_path")

	for _, path := range imagePaths {
		query = query.Values(path)
	}

	if _, err := query.Exec(); err != nil {
		log.Println(err)
		return fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при сохранении путей к изображениям")
	}

	return nil
}
