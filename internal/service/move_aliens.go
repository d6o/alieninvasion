package service

import (
	"context"

	"github.com/d6o/alieninvasion/internal/appcontext"
	"github.com/d6o/alieninvasion/internal/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

//go:generate mockgen -source move_aliens.go -destination mock_move_aliens.go -package service

type (
	MoveAliens struct {
		alienRepository      model.AlienRepository
		cityNearbyRandomizer cityNearbyRandomizer
		verifyFight          verifyFight
	}

	cityNearbyRandomizer interface {
		CityFromCity(city *model.City) *model.City
	}

	verifyFight interface {
		VerifyFight(ctx context.Context, destinationCity *model.City) error
	}
)

func NewMoveAliens(
	alienRepository model.AlienRepository,
	cityNearbyRandomizer cityNearbyRandomizer,
	verifyFight verifyFight,
) *MoveAliens {
	return &MoveAliens{
		alienRepository:      alienRepository,
		cityNearbyRandomizer: cityNearbyRandomizer,
		verifyFight:          verifyFight,
	}
}

func (m *MoveAliens) Move(ctx context.Context) error {
	logger := appcontext.Logger(ctx)
	logger.Debug("Moving aliens")

	for _, alien := range m.alienRepository.All() {
		if alien.IsTrapped() {
			logger.With(zap.Int("alien_id", alien.ID()), zap.String("alien_city", alien.City().Name())).Debug("Alien is trapped")
			continue
		}

		destinationCity := m.cityNearbyRandomizer.CityFromCity(alien.City())
		logger.With(
			zap.Int("alien_id", alien.ID()),
			zap.String("origin_city", alien.City().Name()),
			zap.String("dest_city", destinationCity.Name()),
		).Debug("Moving alien to a new city")
		alien.SetCity(destinationCity)

		if err := m.verifyFight.VerifyFight(ctx, destinationCity); err != nil {
			return errors.Wrap(err, "alien moved, but failed to verify if a fight is happening")
		}
	}

	return nil
}
