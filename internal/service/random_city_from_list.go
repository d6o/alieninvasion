package service

import (
	"math/rand"

	"github.com/d6o/alieninvasion/internal/model"
	"github.com/pkg/errors"
)

type (
	RandomCityFromRepository struct {
		cityRepository model.CityRepository
		rand           *rand.Rand
	}
)

func NewRandomCityFromList(seed int64, cityRepository model.CityRepository) *RandomCityFromRepository {
	return &RandomCityFromRepository{
		cityRepository: cityRepository,
		rand:           rand.New(rand.NewSource(seed)),
	}
}

func (r RandomCityFromRepository) RandomCity() (*model.City, error) {
	list := r.cityRepository.AllNames()

	if len(list) == 0 {
		return nil, errors.New("there are no cities available")
	}

	city := list[r.rand.Intn(len(list))]
	return r.cityRepository.Get(city)
}
