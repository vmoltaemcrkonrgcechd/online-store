package repo

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/entities"
	"github.com/vmoltaemcrkonrgcechd/online-store/pkg/pg"
	"log"
	"net/http"
	"strings"
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
			"category_id", "unit_price", "units_in_stock", "description").
		Values(product.Name, product.UserID, product.ColorID,
			product.CategoryID, product.UnitPrice, product.UnitsInStock, product.Description).
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

func (r ProductRepo) All(params entities.AllProductsQP) (result entities.AllProductsDTO, err error) {
	var (
		subQuery, query string
		args            []any
		rows            *sql.Rows
	)

	subBuilder := r.Sq.Select("count(*)").From("product").Prefix("(").Suffix(")")

	subBuilder = r.addFilter(subBuilder, params)

	if subQuery, _, err = subBuilder.ToSql(); err != nil {
		log.Println(err)
		return result, fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при получении всех продуктов")
	}

	builder := r.Sq.Select("product_id", "product_name", "unit_price", "units_in_stock", "description",
		"color_id", "color_name", "category_id", "category_name",
		"user_id", "username", "u.image_path", "role", "city_id", "city_name", "pi.image_path",
		subQuery).From("product").Join("\"user\" u USING (user_id)").LeftJoin("city USING (city_id)").
		LeftJoin("color USING (color_id)").LeftJoin("category USING (category_id)").
		LeftJoin("product_image pi USING (product_id)")

	if params.OrderBy != nil {
		builder = builder.OrderBy(strings.Join(params.OrderBy, " "))
	}

	if params.Limit != 0 {
		builder = builder.Limit(params.Limit)
	}

	if params.Offset != 0 {
		builder = builder.Offset(params.Offset)
	}

	builder = r.addFilter(builder, params)

	if query, args, err = builder.ToSql(); err != nil {
		log.Println(err)
		return result, fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при получении всех продуктов")
	}

	if rows, err = r.DB.Query(query, args...); err != nil {
		log.Println(err)
		return result, fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при получении всех продуктов")
	}

	return r.scanRows(rows)
}

func (r ProductRepo) addFilter(builder sq.SelectBuilder, params entities.AllProductsQP) sq.SelectBuilder {
	if len(params.Colors) != 0 {
		builder = builder.Where(sq.Eq{"color_id": params.Colors})
	}

	if len(params.Categories) != 0 {
		builder = builder.Where(sq.Eq{"category_id": params.Categories})
	}

	return builder
}

func (r ProductRepo) scanRows(rows *sql.Rows) (result entities.AllProductsDTO, err error) {
	var (
		color, category, city struct{ ID, Name *string }
		imagePath             *string
		productsMap           = make(map[string]*entities.Product)
		product               entities.Product
	)

	for rows.Next() {
		if err = rows.Scan(&product.ID, &product.Name, &product.UnitPrice,
			&product.UnitsInStock, &product.Description, &color.ID, &color.Name, &category.ID,
			&category.Name, &product.User.ID, &product.User.Username,
			&product.User.ImagePath, &product.User.Role, &city.ID,
			&city.Name, &imagePath, &result.Quantity); err != nil {
			log.Println(err)
			return result, fiber.NewError(http.StatusInternalServerError,
				"произошла ошибка при получении всех продуктов")
		}

		if category.ID != nil && category.Name != nil {
			product.Category = &entities.Category{*category.ID, *category.Name}
		}

		if color.ID != nil && color.Name != nil {
			product.Color = &entities.Color{*color.ID, *color.Name}
		}

		if city.ID != nil && city.Name != nil {
			product.User.City = &entities.City{*city.ID, *city.Name}
		}

		if _, ok := productsMap[product.ID]; !ok {
			product.ImagePaths = make([]string, 0)
			result.Products = append(result.Products, product)
			productsMap[product.ID] = &result.Products[len(result.Products)-1]
		}

		if imagePath != nil {
			productsMap[product.ID].ImagePaths =
				append(productsMap[product.ID].ImagePaths, *imagePath)
		}
	}

	return result, nil
}
