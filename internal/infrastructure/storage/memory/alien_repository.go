package memory

import (
	"github.com/d6o/alieninvasion/internal/model"
)

type AlienRepository struct {
	data   map[int]*model.Alien
	nextID int
}

func NewAlienRepository() *AlienRepository {
	return &AlienRepository{
		data:   map[int]*model.Alien{},
		nextID: 0,
	}
}

func (a *AlienRepository) All() map[int]*model.Alien {
	return a.data
}

func (a *AlienRepository) Remove(aliens map[int]*model.Alien) {
	for id := range aliens {
		delete(a.data, id)
	}
}

func (a *AlienRepository) Add(alien *model.Alien) {
	a.data[alien.ID()] = alien
}

func (a *AlienRepository) NextID() int {
	a.nextID++
	return a.nextID
}
