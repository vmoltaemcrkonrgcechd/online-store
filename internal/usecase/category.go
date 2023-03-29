package usecase

import (
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/entities"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/usecase/repo"
)

type CategoryUseCase struct {
	repo repo.CategoryRepo
}

func NewCategoryUseCase(repo repo.CategoryRepo) CategoryUseCase {
	return CategoryUseCase{repo}
}

func (uc CategoryUseCase) All() ([]entities.Category, error) {
	categories, err := uc.repo.All()
	if err != nil {
		return nil, err
	}

	if categories == nil {
		categories = make([]entities.Category, 0, 0)
	}

	return categories, nil
}
