package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/entities"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/usecase/repo"
	"net/http"
)

type ProductUseCase struct {
	repo repo.ProductRepo
}

func NewProductUseCase(repo repo.ProductRepo) ProductUseCase {
	return ProductUseCase{repo}
}

func (uc ProductUseCase) Add(product entities.ProductDTO) (string, error) {
	// TODO добавить проверку.

	id, err := uc.repo.Add(product)
	if err != nil {
		return "", err
	}

	if len(product.ImagePaths) != 0 {
		if err = uc.repo.AddProductImages(id, product.ImagePaths); err != nil {
			return "", fiber.NewError(http.StatusInternalServerError,
				"продукт был добавлен но: "+err.Error())
		}
	}

	return id, nil
}

func (uc ProductUseCase) All() (entities.AllProductsDTO, error) {
	result, err := uc.repo.All()
	if err != nil {
		return result, err
	}

	if result.Products == nil {
		result.Products = make([]entities.Product, 0, 0)
	}

	return result, nil
}
