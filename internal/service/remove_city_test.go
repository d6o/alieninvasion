package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/d6o/alieninvasion/internal/model"
	"github.com/golang/mock/gomock"
)

func TestRemoveCity(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCityRepo := model.NewMockCityRepository(ctrl)

	city1, err := model.NewCity("City1")
	assert.NoError(t, err)
	city2, err := model.NewCity("City2")
	assert.NoError(t, err)

	city1.SetCity(city2, model.North)

	removeCitySvc := NewRemoveCity(mockCityRepo)

	// Expect the city repository to remove the city "A"
	mockCityRepo.EXPECT().Remove(city1.Name())

	removeCitySvc.Remove(ctx, city1)

	if _, ok := city1.Destinations()[model.North]; ok {
		t.Errorf("City 1 still has the North direction")
	}

	if _, ok := city2.Destinations()[model.South]; ok {
		t.Errorf("City 2 still has the South direction")
	}
}
