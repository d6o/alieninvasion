package memory

import (
	"errors"

	"github.com/d6o/alieninvasion/internal/model"
	"github.com/d6o/alieninvasion/pkg/maps"
)

type CityRepository struct {
	data map[string]*model.City
}

func NewCityRepository() *CityRepository {
	return &CityRepository{data: map[string]*model.City{}}
}

func (c *CityRepository) All() map[string]*model.City {
	return c.data
}

func (c *CityRepository) AllNames() []string {
	return maps.Keys(c.data)
}

func (c *CityRepository) Remove(name string) {
	delete(c.data, name)
}

func (c *CityRepository) GetOrAdd(name string) (*model.City, error) {
	if city, err := c.Get(name); err == nil {
		return city, nil
	}

	city, err := model.NewCity(name)
	if err != nil {
		return nil, err
	}

	c.Add(city)

	return city, nil
}

func (c *CityRepository) Get(name string) (*model.City, error) {
	city, ok := c.data[name]
	if !ok {
		return nil, errors.New("city doesn't exist")
	}

	return city, nil
}

func (c *CityRepository) Add(city *model.City) {
	c.data[city.Name()] = city
}
