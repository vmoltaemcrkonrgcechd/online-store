package repo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/entities"
	"github.com/vmoltaemcrkonrgcechd/online-store/pkg/pg"
	"log"
	"net/http"
)

type CityRepo struct {
	*pg.PG
}

func NewCityRepo(pq *pg.PG) CityRepo {
	return CityRepo{pq}
}

func (r CityRepo) All() ([]entities.City, error) {
	userFriendlyErr := fiber.NewError(http.StatusInternalServerError,
		"произошла ошибка при получении городов")

	rows, err := r.Sq.
		Select("city_id", "city_name").
		From("city").
		Query()
	if err != nil {
		log.Println(err)
		return nil, userFriendlyErr
	}

	var (
		city   entities.City
		cities []entities.City
	)
	for rows.Next() {
		if err = rows.Scan(&city.ID, &city.Name); err != nil {
			log.Println(err)
			return nil, userFriendlyErr
		}

		cities = append(cities, city)
	}

	return cities, nil
}
