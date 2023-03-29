package usecase

import (
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/entities"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/usecase/repo"
)

type ColorUseCase struct {
	repo repo.ColorRepo
}

func NewColorUseCase(repo repo.ColorRepo) ColorUseCase {
	return ColorUseCase{repo}
}

func (uc ColorUseCase) All() ([]entities.Color, error) {
	colors, err := uc.repo.All()
	if err != nil {
		return nil, err
	}

	if colors == nil {
		colors = make([]entities.Color, 0, 0)
	}

	return colors, nil
}
