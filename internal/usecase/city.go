package usecase

import (
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/entities"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/usecase/repo"
)

type CityUseCase struct {
	repo repo.CityRepo
}

func NewCityUseCase(repo repo.CityRepo) CityUseCase {
	return CityUseCase{repo}
}

func (uc CityUseCase) All() ([]entities.City, error) {
	cities, err := uc.repo.All()
	if err != nil {
		return nil, err
	}

	if cities == nil {
		cities = make([]entities.City, 0, 0)
	}

	return cities, nil
}
