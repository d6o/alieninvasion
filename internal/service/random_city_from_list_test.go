package service

import (
	"errors"
	"github.com/d6o/alieninvasion/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestNewRandomCityFromList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCityRepository := model.NewMockCityRepository(ctrl)

	randomCityFromRepo := NewRandomCityFromList(1, mockCityRepository)
	assert.NotNil(t, randomCityFromRepo)
	assert.Equal(t, mockCityRepository, randomCityFromRepo.cityRepository)
}

func TestRandomCityFromRepository_RandomCity(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCityRepository := model.NewMockCityRepository(ctrl)

	city1, err := model.NewCity("City1")
	assert.NoError(t, err)
	city2, err := model.NewCity("City2")
	assert.NoError(t, err)
	city3, err := model.NewCity("City3")
	assert.NoError(t, err)

	city1.SetCity(city2, model.North)
	city2.SetCity(city3, model.West)

	t.Run("no cities available", func(t *testing.T) {
		mockCityRepository.EXPECT().AllNames().Return([]string{})

		randomCityFromRepo := NewRandomCityFromList(1, mockCityRepository)
		city, err := randomCityFromRepo.RandomCity()

		assert.Nil(t, city)
		assert.EqualError(t, err, "there are no cities available")
	})

	t.Run("city available", func(t *testing.T) {
		rand.Seed(1)
		mockCityRepository.EXPECT().AllNames().Return([]string{"City1", "City2", "City3"})
		mockCityRepository.EXPECT().Get("City3").Return(city3, nil)

		randomCityFromRepo := NewRandomCityFromList(1, mockCityRepository)
		city, err := randomCityFromRepo.RandomCity()

		assert.NoError(t, err)
		assert.NotNil(t, city)
		assert.Equal(t, city3.Name(), city.Name())
	})

	t.Run("city available, but fails to get from repo", func(t *testing.T) {
		rand.Seed(1)
		mockCityRepository.EXPECT().AllNames().Return([]string{"City1", "City2", "City3"})
		mockCityRepository.EXPECT().Get("City3").Return(nil, errors.New("fake error"))

		randomCityFromRepo := NewRandomCityFromList(1, mockCityRepository)
		city, err := randomCityFromRepo.RandomCity()

		assert.Error(t, err)
		assert.Nil(t, city)
	})
}
