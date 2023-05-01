package service

import (
	"context"
	"errors"
	"github.com/d6o/alieninvasion/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifyFight(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	alienRepo := model.NewMockAlienRepository(ctrl)
	cityRemover := NewMockcityRemover(ctrl)
	printer := NewMockprinter(ctrl)

	verifyFight := NewVerifyFight(alienRepo, cityRemover, printer)

	alien1 := model.NewAlien(1)
	alien2 := model.NewAlien(2)

	cityWithNoAliens, err := model.NewCity("CityWithNoAliens")
	assert.NoError(t, err)

	cityWithOneAlien, err := model.NewCity("CityWithOneAlien")
	assert.NoError(t, err)

	cityWithOneAlien.SetAlien(alien1)

	cityWithTwoAliens, err := model.NewCity("CityWithTwoAliens")
	assert.NoError(t, err)

	cityWithTwoAliens.SetAlien(alien1)
	cityWithTwoAliens.SetAlien(alien2)

	t.Run("no fight, no aliens ", func(t *testing.T) {
		err := verifyFight.VerifyFight(ctx, cityWithNoAliens)
		if err != nil {
			t.Errorf("VerifyFight failed: %v", err)
		}
	})

	t.Run("no fight, sinlge aliens ", func(t *testing.T) {
		err := verifyFight.VerifyFight(ctx, cityWithOneAlien)
		if err != nil {
			t.Errorf("VerifyFight failed: %v", err)
		}
	})

	t.Run("fight_occurs", func(t *testing.T) {
		printer.EXPECT().CityDestroyed(cityWithTwoAliens, cityWithTwoAliens.Aliens()).Return(nil)
		alienRepo.EXPECT().Remove(cityWithTwoAliens.Aliens())
		cityRemover.EXPECT().Remove(ctx, cityWithTwoAliens)

		err := verifyFight.VerifyFight(ctx, cityWithTwoAliens)
		if err != nil {
			t.Errorf("VerifyFight failed: %v", err)
		}
	})

	t.Run("failed_print", func(t *testing.T) {
		printer.EXPECT().CityDestroyed(cityWithTwoAliens, cityWithTwoAliens.Aliens()).Return(errors.New("print error"))

		err := verifyFight.VerifyFight(ctx, cityWithTwoAliens)
		if err == nil || err.Error() != "city destroyed but failed to print the message: print error" {
			t.Errorf("VerifyFight failed: %v", err)
		}
	})
}
