package repo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/entities"
	"github.com/vmoltaemcrkonrgcechd/online-store/pkg/pg"
	"log"
	"net/http"
)

type ProductRepo struct {
	*pg.PG
}

func NewProductRepo(pq *pg.PG) ProductRepo {
	return ProductRepo{pq}
}

func (r ProductRepo) Add(product entities.ProductDTO) (id string, err error) {
	if err = r.Sq.Insert("product").
		Columns("product_name", "user_id", "color_id",
			"category_id", "unit_price", "units_in_stock").
		Values(product.Name, product.UserID, product.ColorID,
			product.CategoryID, product.UnitPrice, product.UnitsInStock).
		Suffix("RETURNING product_id").QueryRow().Scan(&id); err != nil {
		log.Println(err)
		return "", fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при добавлении продукта")
	}

	return id, nil
}

func (r ProductRepo) AddProductImages(id string, imagePaths []string) error {
	query := r.Sq.Insert("product_image").
		Columns("product_id", "image_path")

	for _, path := range imagePaths {
		query = query.Values(id, path)
	}

	if _, err := query.Exec(); err != nil {
		log.Println(err)
		return fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при добавлении изображений к продукту")
	}

	return nil
}
