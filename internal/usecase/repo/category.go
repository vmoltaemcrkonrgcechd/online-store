package repo

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/entities"
	"github.com/vmoltaemcrkonrgcechd/online-store/pkg/pg"
	"log"
	"net/http"
)

type CategoryRepo struct {
	*pg.PG
}

func NewCategoryRepo(pq *pg.PG) CategoryRepo {
	return CategoryRepo{pq}
}

func (r CategoryRepo) All() (categories []entities.Category, err error) {
	var rows *sql.Rows
	if rows, err = r.Sq.Select("category_id", "category_name").
		From("category").Query(); err != nil {
		log.Println(err)
		return nil, fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при получении всех категорий")
	}

	var category entities.Category
	for rows.Next() {
		if err = rows.Scan(&category.ID, &category.Name); err != nil {
			log.Println(err)
			return nil, fiber.NewError(http.StatusInternalServerError,
				"произошла ошибка при получении всех категорий")
		}

		categories = append(categories, category)
	}

	return categories, nil
}
