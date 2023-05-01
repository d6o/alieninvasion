package model

//go:generate mockgen -source alien.go -destination mock_alien.go -package model

type (
	Alien struct {
		id   int
		city *City
	}

	AlienRepository interface {
		NextID() int
		All() map[int]*Alien
		Remove(aliens map[int]*Alien)
		Add(alien *Alien)
	}
)

func NewAlien(id int) *Alien {
	return &Alien{id: id}
}

func (a *Alien) ID() int {
	return a.id
}

func (a *Alien) City() *City {
	return a.city
}

func (a *Alien) SetCity(city *City) {
	if a.city != nil {
		a.city.RemoveAlien(a)
	}

	a.city = city
	a.city.SetAlien(a)
}

func (a *Alien) IsTrapped() bool {
	return !a.city.HasLinks()
}
