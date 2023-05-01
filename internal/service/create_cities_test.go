package service_test

import (
	"bytes"
	"context"
	"errors"
	"github.com/d6o/alieninvasion/internal/model"
	"github.com/d6o/alieninvasion/internal/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFromReader(t *testing.T) {
	ctx := context.Background()

	cityA, err := model.NewCity("CityA")
	assert.NoError(t, err)
	cityB, err := model.NewCity("CityB")
	assert.NoError(t, err)
	cityC, err := model.NewCity("CityC")
	assert.NoError(t, err)
	cityD, err := model.NewCity("CityD")
	assert.NoError(t, err)
	cityE, err := model.NewCity("CityE")
	assert.NoError(t, err)

	tests := []struct {
		name           string
		input          string
		mockRepoCalls  func(repo *model.MockCityRepository)
		wantErr        bool
		expectedCities map[string]map[model.Direction]string
	}{
		{
			name:  "Example",
			input: "CityA north=CityB west=CityC south=CityD\nCityB south=CityA west=CityE",
			mockRepoCalls: func(repo *model.MockCityRepository) {
				repo.EXPECT().GetOrAdd(cityA.Name()).Return(cityA, nil).AnyTimes()
				repo.EXPECT().GetOrAdd(cityB.Name()).Return(cityB, nil).AnyTimes()
				repo.EXPECT().GetOrAdd(cityC.Name()).Return(cityC, nil).AnyTimes()
				repo.EXPECT().GetOrAdd(cityD.Name()).Return(cityD, nil).AnyTimes()
				repo.EXPECT().GetOrAdd(cityE.Name()).Return(cityE, nil).AnyTimes()
				repo.EXPECT().All().Return(map[string]*model.City{
					cityA.Name(): cityA,
					cityB.Name(): cityB,
					cityC.Name(): cityC,
					cityD.Name(): cityD,
					cityE.Name(): cityE,
				})
			},
			wantErr: false,
			expectedCities: map[string]map[model.Direction]string{
				"CityA": {
					model.North: "CityB",
					model.West:  "CityC",
					model.South: "CityD",
				},
				"CityB": {
					model.South: "CityA",
					model.West:  "CityE",
				},
				"CityC": {
					model.East: "CityA",
				},
				"CityD": {
					model.North: "CityA",
				},
				"CityE": {
					model.East: "CityB",
				},
			},
		},
		{
			name:  "valid input",
			input: "CityA east=CityB north=CityC\nCityB west=CityA\nCityC south=CityA",
			mockRepoCalls: func(repo *model.MockCityRepository) {
				repo.EXPECT().GetOrAdd(cityA.Name()).Return(cityA, nil).AnyTimes()
				repo.EXPECT().GetOrAdd(cityB.Name()).Return(cityB, nil).AnyTimes()
				repo.EXPECT().GetOrAdd(cityC.Name()).Return(cityC, nil).AnyTimes()
				repo.EXPECT().All().Return(map[string]*model.City{
					cityA.Name(): cityA,
					cityB.Name(): cityB,
					cityC.Name(): cityC,
				})
			},
			wantErr: false,
			expectedCities: map[string]map[model.Direction]string{
				"CityA": {
					model.East:  "CityB",
					model.North: "CityC",
				},
				"CityB": {
					model.West: "CityA",
				},
				"CityC": {
					model.South: "CityA",
				},
			},
		},
		{
			name:          "not enough cities",
			input:         "",
			mockRepoCalls: func(repo *model.MockCityRepository) {},
			wantErr:       true,
		},
		{
			name:  "invalid direction",
			input: "CityA east=CityB north=CityC\nCityB invalid=CityA",
			mockRepoCalls: func(repo *model.MockCityRepository) {
				repo.EXPECT().GetOrAdd(cityA.Name()).Return(cityA, nil).AnyTimes()
				repo.EXPECT().GetOrAdd(cityB.Name()).Return(cityB, nil).AnyTimes()
				repo.EXPECT().GetOrAdd(cityC.Name()).Return(cityC, nil).AnyTimes()
				repo.EXPECT().All().Return(map[string]*model.City{
					cityA.Name(): cityA,
					cityB.Name(): cityB,
					cityC.Name(): cityC,
				})
			},
			wantErr: false,
		},
		{
			name:  "repository error",
			input: "CityA east=CityB north=CityC\nCityB west=CityA",
			mockRepoCalls: func(repo *model.MockCityRepository) {
				repo.EXPECT().GetOrAdd(cityA.Name()).Return(cityA, nil).AnyTimes().Return(nil, errors.New("repository error"))
				repo.EXPECT().GetOrAdd(cityB.Name()).Return(cityB, nil).AnyTimes()
				repo.EXPECT().GetOrAdd(cityC.Name()).Return(cityC, nil).AnyTimes()
				repo.EXPECT().All().Return(map[string]*model.City{
					cityB.Name(): cityB,
					cityC.Name(): cityC,
				})
			},
			wantErr: false,
		},
		{
			name:  "invalid position format",
			input: "CityA eastCityB northCityC\nCityB westCityA",
			mockRepoCalls: func(repo *model.MockCityRepository) {
				repo.EXPECT().GetOrAdd(cityA.Name()).Return(cityA, nil).AnyTimes()
				repo.EXPECT().GetOrAdd(cityB.Name()).Return(cityB, nil).AnyTimes()
				repo.EXPECT().GetOrAdd(cityC.Name()).Return(cityC, nil).AnyTimes()
				repo.EXPECT().All().Return(map[string]*model.City{
					cityB.Name(): cityB,
					cityC.Name(): cityC,
				})
			},
			wantErr: false,
		},
		{
			name:          "empty city name",
			input:         " east=CityB north=CityC\nCityB west=CityA",
			mockRepoCalls: func(repo *model.MockCityRepository) {},
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockRepo := model.NewMockCityRepository(ctrl)
			tt.mockRepoCalls(mockRepo)

			c := service.NewCreateCities(mockRepo)

			reader := bytes.NewReader([]byte(tt.input))
			err := c.ParseFromReader(ctx, reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFromReader() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Validate city connections
			for cityName, expectedConnections := range tt.expectedCities {
				city, err := mockRepo.GetOrAdd(cityName)
				if err != nil {
					t.Errorf("Error getting city %s: %v", cityName, err)
				}

				for direction, expectedDestination := range expectedConnections {
					actualDestination := city.Destinations()[direction]
					if actualDestination == nil || actualDestination.Name() != expectedDestination {
						t.Errorf("City %s, direction %s: expected destination %s, got %s", cityName, direction, expectedDestination, actualDestination)
					}
				}
			}

			ctrl.Finish()
		})
	}
}
