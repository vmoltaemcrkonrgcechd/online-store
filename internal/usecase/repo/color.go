package repo

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/entities"
	"github.com/vmoltaemcrkonrgcechd/online-store/pkg/pg"
	"log"
	"net/http"
)

type ColorRepo struct {
	*pg.PG
}

func NewColorRepo(pq *pg.PG) ColorRepo {
	return ColorRepo{pq}
}

func (r ColorRepo) All() (colors []entities.Color, err error) {
	var rows *sql.Rows
	if rows, err = r.Sq.Select("color_id", "color_name").
		From("color").Query(); err != nil {
		log.Println(err)
		return nil, fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при получении всех цветов")
	}

	var color entities.Color
	for rows.Next() {
		if err = rows.Scan(&color.ID, &color.Name); err != nil {
			log.Println(err)
			return nil, fiber.NewError(http.StatusInternalServerError,
				"произошла ошибка при получении всех цветов")
		}

		colors = append(colors, color)
	}

	return colors, nil
}
