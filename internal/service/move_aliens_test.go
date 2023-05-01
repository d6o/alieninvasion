package service

import (
	"context"
	"errors"
	"testing"

	"github.com/d6o/alieninvasion/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewMoveAliens(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlienRepo := model.NewMockAlienRepository(ctrl)
	mockCityNearbyRandomizer := NewMockcityNearbyRandomizer(ctrl)
	mockVerifyFight := NewMockverifyFight(ctrl)

	moveAliens := NewMoveAliens(mockAlienRepo, mockCityNearbyRandomizer, mockVerifyFight)

	assert.NotNil(t, moveAliens)
}

func TestMoveAliens_Move(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlienRepo := model.NewMockAlienRepository(ctrl)
	mockCityNearbyRandomizer := NewMockcityNearbyRandomizer(ctrl)
	mockVerifyFight := NewMockverifyFight(ctrl)

	moveAliens := &MoveAliens{
		alienRepository:      mockAlienRepo,
		cityNearbyRandomizer: mockCityNearbyRandomizer,
		verifyFight:          mockVerifyFight,
	}

	ctx := context.Background()

	// Create some sample data.
	city1, err := model.NewCity("City1")
	assert.NoError(t, err)
	city2, err := model.NewCity("City2")
	assert.NoError(t, err)

	city1.SetCity(city2, model.North)

	alien1 := model.NewAlien(1)
	alien2 := model.NewAlien(2)

	alien1.SetCity(city1)
	alien2.SetCity(city2)

	// Set up mock expectations.
	mockAlienRepo.EXPECT().All().Return(map[int]*model.Alien{
		alien1.ID(): alien1,
		alien2.ID(): alien2,
	})
	mockCityNearbyRandomizer.EXPECT().CityFromCity(city1).Return(city2)
	mockCityNearbyRandomizer.EXPECT().CityFromCity(city2).Return(city1)
	mockVerifyFight.EXPECT().VerifyFight(ctx, city2).Return(nil)
	mockVerifyFight.EXPECT().VerifyFight(ctx, city1).Return(nil)

	// Run the Move function.
	err = moveAliens.Move(ctx)

	assert.NoError(t, err)
}

func TestMoveAliens_Move_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlienRepo := model.NewMockAlienRepository(ctrl)
	mockCityNearbyRandomizer := NewMockcityNearbyRandomizer(ctrl)
	mockVerifyFight := NewMockverifyFight(ctrl)

	moveAliens := &MoveAliens{
		alienRepository:      mockAlienRepo,
		cityNearbyRandomizer: mockCityNearbyRandomizer,
		verifyFight:          mockVerifyFight,
	}

	ctx := context.Background()

	// Create some sample data.
	city1, err := model.NewCity("City1")
	assert.NoError(t, err)
	city2, err := model.NewCity("City2")
	assert.NoError(t, err)

	city1.SetCity(city2, model.North)

	alien1 := model.NewAlien(1)
	alien2 := model.NewAlien(2)

	alien1.SetCity(city1)
	alien2.SetCity(city2)

	// Set up mock expectations.
	mockAlienRepo.EXPECT().All().Return(map[int]*model.Alien{
		alien1.ID(): alien1,
	})
	mockCityNearbyRandomizer.EXPECT().CityFromCity(city1).Return(city2)
	mockVerifyFight.EXPECT().VerifyFight(ctx, city2).Return(errors.New("fake error"))
	// Run the Move function.
	err = moveAliens.Move(ctx)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "alien moved, but failed to verify if a fight is happening")
}
