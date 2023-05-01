package model

import (
	"errors"
	"fmt"

	"github.com/d6o/alieninvasion/pkg/maps"
)

//go:generate mockgen -source city.go -destination mock_city.go -package model

type (
	City struct {
		name       string
		directions map[Direction]*City
		aliens     map[int]*Alien
	}

	CityRepository interface {
		All() map[string]*City
		AllNames() []string
		Remove(name string)
		GetOrAdd(name string) (*City, error)
		Get(name string) (*City, error)
	}
)

const (
	minLengthName = 2
)

func NewCity(name string) (*City, error) {
	if len(name) < minLengthName {
		return nil, errors.New("invalid city name")
	}

	return &City{
		name:       name,
		directions: map[Direction]*City{},
		aliens:     map[int]*Alien{},
	}, nil
}

func (c *City) Destinations() map[Direction]*City {
	return c.directions
}

func (c *City) Directions() []Direction {
	return maps.Keys(c.directions)
}

func (c *City) HasLinks() bool {
	return len(c.directions) > 0
}

func (c *City) Aliens() map[int]*Alien {
	return c.aliens
}

func (c *City) HaveMultipleAliens() bool {
	return len(c.aliens) > 1
}

func (c *City) Name() string {
	return c.name
}

func (c *City) SetAlien(alien *Alien) {
	c.aliens[alien.id] = alien
}

func (c *City) SetCity(otherCity *City, direction Direction) {
	if c.directions[direction] == otherCity {
		return
	}

	c.directions[direction] = otherCity

	otherCity.SetCity(c, direction.Opposite())
}

func (c *City) String() string {
	var dir string

	for direction, city := range c.directions {
		dir += fmt.Sprintf("%s=%s ", direction, city.name)
	}

	return fmt.Sprintf("%s %s", c.name, dir)
}

func (c *City) RemoveDirection(direction Direction) {
	delete(c.directions, direction)
}

func (c *City) RemoveAlien(alien *Alien) {
	delete(c.aliens, alien.id)
}
