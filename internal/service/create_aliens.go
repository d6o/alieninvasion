package service

import (
	"context"

	"github.com/d6o/alieninvasion/internal/appcontext"
	"github.com/d6o/alieninvasion/internal/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

//go:generate mockgen -source create_aliens.go -destination mock_create_aliens.go -package service

type (
	CreateAliens struct {
		alienRepository model.AlienRepository
		cityRandomizer  cityRandomizer
		verifyFight     verifyFight
	}

	cityRandomizer interface {
		RandomCity() (*model.City, error)
	}
)

func NewCreateAliens(
	alienRepository model.AlienRepository,
	cityRandomizer cityRandomizer,
	verifyFight verifyFight,
) *CreateAliens {
	return &CreateAliens{
		alienRepository: alienRepository,
		cityRandomizer:  cityRandomizer,
		verifyFight:     verifyFight,
	}
}

func (c CreateAliens) Create(ctx context.Context, num int) error {
	logger := appcontext.Logger(ctx)
	logger.With(zap.Int("num_aliens", num)).Debug("Creating aliens")
	for i := 0; i < num; i++ {
		city, err := c.cityRandomizer.RandomCity()
		if err != nil {
			return errors.Wrap(err, "can't get a city to allocate the alien")
		}

		alien := model.NewAlien(c.alienRepository.NextID())

		alien.SetCity(city)

		logger.With(zap.Int("alien_id", alien.ID()), zap.String("alien_city", city.Name())).Debug("Alien created")

		c.alienRepository.Add(alien)

		if err := c.verifyFight.VerifyFight(ctx, city); err != nil {
			return errors.Wrap(err, "alien created, but failed to verify if a fight is happening")
		}
	}

	logger.Debug("Aliens created")
	return nil
}
