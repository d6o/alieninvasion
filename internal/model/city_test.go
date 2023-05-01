package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCity_Aliens(t *testing.T) {
	type fields struct {
		name   string
		aliens []*Alien
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "No Aliens",
			fields: fields{
				name:   "TestCity",
				aliens: []*Alien{},
			},
			want: 0,
		},
		{
			name: "One Alien",
			fields: fields{
				name: "TestCity",
				aliens: []*Alien{
					NewAlien(1),
				},
			},
			want: 1,
		},
		{
			name: "Three Alien",
			fields: fields{
				name: "TestCity",
				aliens: []*Alien{
					NewAlien(1),
					NewAlien(2),
					NewAlien(3),
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewCity(tt.fields.name)
			assert.Nil(t, err)

			for _, alien := range tt.fields.aliens {
				alien.SetCity(c)
			}

			assert.Equalf(t, tt.want, len(c.Aliens()), "Aliens()")
		})
	}
}

func TestCity_Destinations(t *testing.T) {
	cityA, err := NewCity("CityA")
	assert.Nil(t, err)

	cityB, err := NewCity("CityB")
	assert.Nil(t, err)

	cityC, err := NewCity("CityC")
	assert.Nil(t, err)

	island, err := NewCity("Island")
	assert.Nil(t, err)

	cityA.SetCity(cityB, North)
	cityA.SetCity(cityC, East)

	type fields struct {
		city *City
	}
	tests := []struct {
		name   string
		fields fields
		want   map[Direction]*City
	}{
		{
			name: "City without destinations",
			fields: fields{
				city: island,
			},
			want: map[Direction]*City{},
		},
		{
			name: "City with one destination",
			fields: fields{
				city: cityB,
			},
			want: map[Direction]*City{
				South: cityA,
			},
		},
		{
			name: "City with multiple destinations",
			fields: fields{
				city: cityA,
			},
			want: map[Direction]*City{
				North: cityB,
				East:  cityC,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.fields.city.Destinations(), "Destinations()")
		})
	}
}

func TestCity_Directions(t *testing.T) {
	cityA, err := NewCity("CityA")
	assert.Nil(t, err)

	cityB, err := NewCity("CityB")
	assert.Nil(t, err)

	cityC, err := NewCity("CityC")
	assert.Nil(t, err)

	island, err := NewCity("Island")
	assert.Nil(t, err)

	cityA.SetCity(cityB, North)
	cityA.SetCity(cityC, East)

	type fields struct {
		city *City
	}
	tests := []struct {
		name   string
		fields fields
		want   []Direction
	}{
		{
			name: "City with no directions",
			fields: fields{
				city: island,
			},
			want: nil,
		},
		{
			name: "City with one direction",
			fields: fields{
				city: cityB,
			},
			want: []Direction{South},
		},
		{
			name: "City with multiple directions",
			fields: fields{
				city: cityA,
			},
			want: []Direction{North, East},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, tt.fields.city.Directions(), "Directions()")
		})
	}
}

func TestCity_HasLinks(t *testing.T) {
	cityA, err := NewCity("CityA")
	assert.Nil(t, err)

	cityB, err := NewCity("CityB")
	assert.Nil(t, err)

	island, err := NewCity("Island")
	assert.Nil(t, err)

	cityA.SetCity(cityB, North)

	type fields struct {
		city *City
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "City doesn't have any links",
			fields: fields{
				city: island,
			},
			want: false,
		},
		{
			name: "City have any links",
			fields: fields{
				city: cityA,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.fields.city.HasLinks(), "HasLinks()")
		})
	}
}

func TestCity_HaveMultipleAliens(t *testing.T) {
	type fields struct {
		name   string
		aliens []*Alien
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "City has no aliens",
			fields: fields{
				name:   "TestCity",
				aliens: []*Alien{},
			},
			want: false,
		},
		{
			name: "City has one alien",
			fields: fields{
				name: "TestCity",
				aliens: []*Alien{
					NewAlien(1),
				},
			},
			want: false,
		},
		{
			name: "City has two aliens",
			fields: fields{
				name: "TestCity",
				aliens: []*Alien{
					NewAlien(1),
					NewAlien(2),
				},
			},
			want: true,
		},
		{
			name: "City has three aliens",
			fields: fields{
				name: "TestCity",
				aliens: []*Alien{
					NewAlien(1),
					NewAlien(2),
					NewAlien(3),
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewCity(tt.fields.name)
			assert.Nil(t, err)

			for _, alien := range tt.fields.aliens {
				alien.SetCity(c)
			}

			assert.Equalf(t, tt.want, c.HaveMultipleAliens(), "HaveMultipleAliens()")
		})
	}
}

func TestCity_Name(t *testing.T) {
	type fields struct {
		name       string
		directions map[Direction]*City
		aliens     map[int]*Alien
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &City{
				name:       tt.fields.name,
				directions: tt.fields.directions,
				aliens:     tt.fields.aliens,
			}
			assert.Equalf(t, tt.want, c.Name(), "Name()")
		})
	}
}

func TestCity_RemoveAlien(t *testing.T) {
	type fields struct {
		name       string
		directions map[Direction]*City
		aliens     map[int]*Alien
	}
	type args struct {
		alien *Alien
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &City{
				name:       tt.fields.name,
				directions: tt.fields.directions,
				aliens:     tt.fields.aliens,
			}
			c.RemoveAlien(tt.args.alien)
		})
	}
}

func TestCity_RemoveDirection(t *testing.T) {
	type fields struct {
		name       string
		directions map[Direction]*City
		aliens     map[int]*Alien
	}
	type args struct {
		direction Direction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &City{
				name:       tt.fields.name,
				directions: tt.fields.directions,
				aliens:     tt.fields.aliens,
			}
			c.RemoveDirection(tt.args.direction)
		})
	}
}

func TestCity_SetAlien(t *testing.T) {
	type fields struct {
		name       string
		directions map[Direction]*City
		aliens     map[int]*Alien
	}
	type args struct {
		alien *Alien
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &City{
				name:       tt.fields.name,
				directions: tt.fields.directions,
				aliens:     tt.fields.aliens,
			}
			c.SetAlien(tt.args.alien)
		})
	}
}

func TestCity_SetCity(t *testing.T) {
	type fields struct {
		name       string
		directions map[Direction]*City
		aliens     map[int]*Alien
	}
	type args struct {
		otherCity *City
		direction Direction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &City{
				name:       tt.fields.name,
				directions: tt.fields.directions,
				aliens:     tt.fields.aliens,
			}
			c.SetCity(tt.args.otherCity, tt.args.direction)
		})
	}
}

func TestCity_String(t *testing.T) {
	type fields struct {
		name       string
		directions map[Direction]*City
		aliens     map[int]*Alien
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := City{
				name:       tt.fields.name,
				directions: tt.fields.directions,
				aliens:     tt.fields.aliens,
			}
			assert.Equalf(t, tt.want, c.String(), "String()")
		})
	}
}

func TestNewCity(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *City
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCity(tt.args.name)
			if !tt.wantErr(t, err, fmt.Sprintf("NewCity(%v)", tt.args.name)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewCity(%v)", tt.args.name)
		})
	}
}
