package memory

import (
	"github.com/d6o/alieninvasion/internal/model"
	"reflect"
	"testing"
)

func TestAlienRepository_Add(t *testing.T) {
	firstAlien := model.NewAlien(1)
	secondAlien := model.NewAlien(2)

	type fields struct {
		data map[int]*model.Alien
	}
	type args struct {
		alien *model.Alien
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[int]*model.Alien
	}{
		{
			name: "Adds the first alien",
			fields: fields{
				data: map[int]*model.Alien{},
			},
			args: args{
				alien: firstAlien,
			},
			want: map[int]*model.Alien{
				1: firstAlien,
			},
		},
		{
			name: "Adds one alien, but one already exists",
			fields: fields{
				data: map[int]*model.Alien{
					1: firstAlien,
				},
			},
			args: args{
				alien: secondAlien,
			},
			want: map[int]*model.Alien{
				1: firstAlien,
				2: secondAlien,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AlienRepository{
				data: tt.fields.data,
			}
			a.Add(tt.args.alien)
			if got := a.data; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlienRepository_All(t *testing.T) {
	firstAlien := model.NewAlien(1)
	secondAlien := model.NewAlien(2)

	tests := []struct {
		name string
		data []*model.Alien
		want map[int]*model.Alien
	}{
		{
			name: "There are no aliens in the repository",
			data: []*model.Alien{},
			want: map[int]*model.Alien{},
		},
		{
			name: "There is one alien in the repository",
			data: []*model.Alien{firstAlien},
			want: map[int]*model.Alien{
				1: firstAlien,
			},
		},
		{
			name: "There are two aliens in the repository",
			data: []*model.Alien{firstAlien, secondAlien},
			want: map[int]*model.Alien{
				1: firstAlien,
				2: secondAlien,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAlienRepository()

			for _, alien := range tt.data {
				a.Add(alien)
			}

			if got := a.All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlienRepository_NextID(t *testing.T) {
	firstAlien := model.NewAlien(1)
	secondAlien := model.NewAlien(2)

	tests := []struct {
		name string
		data []*model.Alien
		want int
	}{
		{
			name: "No aliens in the repo",
			data: nil,
			want: 1,
		},
		{
			name: "One alien in the repo",
			data: []*model.Alien{
				firstAlien,
			},
			want: 2,
		},
		{
			name: "Two aliens in the repo",
			data: []*model.Alien{
				firstAlien,
				secondAlien,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAlienRepository()
			for _, alien := range tt.data {
				a.Add(alien)
				a.NextID()
			}

			if got := a.NextID(); got != tt.want {
				t.Errorf("NextID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlienRepository_Remove(t *testing.T) {
	firstAlien := model.NewAlien(1)
	secondAlien := model.NewAlien(2)
	thirdAlien := model.NewAlien(3)

	type fields struct {
		data map[int]*model.Alien
	}
	type args struct {
		aliens map[int]*model.Alien
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[int]*model.Alien
	}{
		{
			name: "Has one, removes one",
			fields: fields{
				data: map[int]*model.Alien{
					1: firstAlien,
				},
			},
			args: args{
				aliens: map[int]*model.Alien{
					1: firstAlien,
				},
			},
			want: map[int]*model.Alien{},
		},
		{
			name: "Has two, removes one",
			fields: fields{
				data: map[int]*model.Alien{
					1: firstAlien,
					2: secondAlien,
				},
			},
			args: args{
				aliens: map[int]*model.Alien{
					2: secondAlien,
				},
			},
			want: map[int]*model.Alien{
				1: firstAlien,
			},
		},
		{
			name: "Tries to deletes an alien that doesn't exist in the repo",
			fields: fields{
				data: map[int]*model.Alien{
					2: secondAlien,
					3: thirdAlien,
				},
			},
			args: args{
				aliens: map[int]*model.Alien{
					1: firstAlien,
				},
			},
			want: map[int]*model.Alien{
				2: secondAlien,
				3: thirdAlien,
			},
		},
		{
			name: "Tries to deletes an alien from an empty repo",
			fields: fields{
				data: map[int]*model.Alien{},
			},
			args: args{
				aliens: map[int]*model.Alien{
					1: firstAlien,
				},
			},
			want: map[int]*model.Alien{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AlienRepository{
				data: tt.fields.data,
			}
			a.Remove(tt.args.aliens)
			if got := a.data; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAlienRepository(t *testing.T) {
	dataWant := map[int]*model.Alien{}

	alienRepo := NewAlienRepository()
	if !reflect.DeepEqual(alienRepo.data, dataWant) {
		t.Errorf("NewAlienRepository().data = %v, want %v", alienRepo.data, dataWant)
	}
}
