package service_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/d6o/alieninvasion/internal/model"
	"github.com/d6o/alieninvasion/internal/service"
	"github.com/golang/mock/gomock"
)

func TestAreEnoughAliensAround(t *testing.T) {
	ctx := context.Background()

	alien1 := model.NewAlien(1)
	alien2 := model.NewAlien(2)

	city1, err := model.NewCity("City1")
	assert.NoError(t, err)
	city2, err := model.NewCity("City2")
	assert.NoError(t, err)

	alien1.SetCity(city1)
	alien2.SetCity(city2)

	t.Run("NotEnoughAliensAlive", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockAlienRepo := model.NewMockAlienRepository(ctrl)
		mockAlienRepo.EXPECT().All().Return(map[int]*model.Alien{})

		verifyAliens := service.NewVerifyAliens(mockAlienRepo)
		assert.False(t, verifyAliens.AreEnoughAliensAround(ctx))
	})

	t.Run("AliensTrapped", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockAlienRepo := model.NewMockAlienRepository(ctrl)
		mockAlienRepo.EXPECT().All().Return(map[int]*model.Alien{
			alien1.ID(): alien1,
			alien2.ID(): alien2,
		})

		verifyAliens := service.NewVerifyAliens(mockAlienRepo)
		assert.False(t, verifyAliens.AreEnoughAliensAround(ctx))
	})

	city1.SetCity(city2, model.North)

	t.Run("AliensCanMove", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockAlienRepo := model.NewMockAlienRepository(ctrl)
		mockAlienRepo.EXPECT().All().Return(map[int]*model.Alien{
			alien1.ID(): alien1,
			alien2.ID(): alien2,
		})

		verifyAliens := service.NewVerifyAliens(mockAlienRepo)
		assert.True(t, verifyAliens.AreEnoughAliensAround(ctx))
	})
}
