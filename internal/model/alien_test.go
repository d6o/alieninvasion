package model

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestAlien_City(t *testing.T) {
	city := &City{
		name: "Test",
	}

	type fields struct {
		id   int
		city *City
	}
	tests := []struct {
		name   string
		fields fields
		want   *City
	}{
		{
			name: "Returns the city the alien is allocated to",
			fields: fields{
				id:   1,
				city: city,
			},
			want: city,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Alien{
				id:   tt.fields.id,
				city: tt.fields.city,
			}
			if got := a.City(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("City() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlien_Id(t *testing.T) {
	type fields struct {
		id int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Returns Alien's ID",
			fields: fields{
				id: 100,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Alien{
				id: tt.fields.id,
			}
			if got := a.ID(); got != tt.want {
				t.Errorf("ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlien_IsTrapped(t *testing.T) {
	cityA, err := NewCity("CityA")
	assert.Nil(t, err)

	cityB, err := NewCity("CityB")
	assert.Nil(t, err)

	island, err := NewCity("Island")
	assert.Nil(t, err)

	cityA.SetCity(cityB, North)

	type fields struct {
		id   int
		city *City
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Alien is on a city with no links",
			fields: fields{
				id:   1,
				city: island,
			},
			want: true,
		},
		{
			name: "Alien is on a city with links",
			fields: fields{
				id:   1,
				city: cityA,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Alien{
				id:   tt.fields.id,
				city: tt.fields.city,
			}
			if got := a.IsTrapped(); got != tt.want {
				t.Errorf("IsTrapped() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlien_SetCity(t *testing.T) {
	alien := NewAlien(1)
	city, err := NewCity("test")
	assert.Nil(t, err)

	alien.SetCity(city)

	assert.Equal(t, alien.City(), city)
}

func TestNewAlien(t *testing.T) {
	alien := NewAlien(5)
	assert.Equal(t, alien.ID(), 5)
}
