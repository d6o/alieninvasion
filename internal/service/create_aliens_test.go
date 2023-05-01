package service

import (
	"context"
	"errors"
	"testing"

	"github.com/d6o/alieninvasion/internal/appcontext"
	"github.com/d6o/alieninvasion/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNewCreateAliens(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlienRepo := model.NewMockAlienRepository(ctrl)
	mockCityRandomizer := NewMockcityRandomizer(ctrl)
	mockVerifyFight := NewMockverifyFight(ctrl)

	ca := NewCreateAliens(mockAlienRepo, mockCityRandomizer, mockVerifyFight)

	assert.NotNil(t, ca)
}

func TestCreateAliens_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlienRepo := model.NewMockAlienRepository(ctrl)
	mockCityRandomizer := NewMockcityRandomizer(ctrl)
	mockVerifyFight := NewMockverifyFight(ctrl)

	ca := &CreateAliens{
		alienRepository: mockAlienRepo,
		cityRandomizer:  mockCityRandomizer,
		verifyFight:     mockVerifyFight,
	}

	ctx := appcontext.WithLogger(context.Background(), zap.NewNop())

	mockCity, err := model.NewCity("City1")
	assert.NoError(t, err)

	mockAlienRepo.EXPECT().NextID().Return(1).Times(1)
	mockAlienRepo.EXPECT().Add(gomock.Any()).Times(1)

	mockCityRandomizer.EXPECT().RandomCity().Return(mockCity, nil).Times(1)

	mockVerifyFight.EXPECT().VerifyFight(ctx, mockCity).Return(nil).Times(1)

	err = ca.Create(ctx, 1)

	assert.NoError(t, err)
}

func TestCreateAliens_Create_ErrorRandomCity(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlienRepo := model.NewMockAlienRepository(ctrl)
	mockCityRandomizer := NewMockcityRandomizer(ctrl)
	mockVerifyFight := NewMockverifyFight(ctrl)

	ca := &CreateAliens{
		alienRepository: mockAlienRepo,
		cityRandomizer:  mockCityRandomizer,
		verifyFight:     mockVerifyFight,
	}

	ctx := appcontext.WithLogger(context.Background(), zap.NewNop())

	mockCityRandomizer.EXPECT().RandomCity().Return(nil, errors.New("random city error")).Times(1)

	err := ca.Create(ctx, 1)

	assert.Error(t, err)
}

func TestCreateAliens_Create_ErrorVerifyFight(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlienRepo := model.NewMockAlienRepository(ctrl)
	mockCityRandomizer := NewMockcityRandomizer(ctrl)
	mockVerifyFight := NewMockverifyFight(ctrl)

	ca := &CreateAliens{
		alienRepository: mockAlienRepo,
		cityRandomizer:  mockCityRandomizer,
		verifyFight:     mockVerifyFight,
	}

	ctx := appcontext.WithLogger(context.Background(), zap.NewNop())

	mockCity, err := model.NewCity("City1")
	assert.NoError(t, err)

	mockAlienRepo.EXPECT().NextID().Return(1).Times(1)
	mockAlienRepo.EXPECT().Add(gomock.Any()).Times(1)

	mockCityRandomizer.EXPECT().RandomCity().Return(mockCity, nil).Times(1)
	mockVerifyFight.EXPECT().VerifyFight(ctx, mockCity).Return(errors.New("verify fight error")).Times(1)

	err = ca.Create(ctx, 1)

	assert.Error(t, err)
}
