package service

import (
	"math/rand"

	"github.com/d6o/alieninvasion/internal/model"
)

type (
	RandomCityFromCity struct {
		rand *rand.Rand
	}
)

func NewRandomCityFromCity(seed int64) *RandomCityFromCity {
	return &RandomCityFromCity{
		rand: rand.New(rand.NewSource(seed)),
	}
}

func (r RandomCityFromCity) CityFromCity(city *model.City) *model.City {
	directions := city.Directions()
	return city.Destinations()[directions[r.rand.Intn(len(directions))]]
}
