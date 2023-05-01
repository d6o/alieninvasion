package service

import (
	"context"

	"github.com/d6o/alieninvasion/internal/appcontext"
	"github.com/d6o/alieninvasion/internal/model"
	"github.com/d6o/alieninvasion/pkg/maps"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

//go:generate mockgen -source fight_aliens.go -destination mock_fight_aliens.go -package service

type (
	VerifyFight struct {
		alienRepository model.AlienRepository
		cityRemover     cityRemover
		printer         printer
	}

	cityRemover interface {
		Remove(ctx context.Context, city *model.City)
	}

	printer interface {
		CityDestroyed(city *model.City, aliens map[int]*model.Alien) error
	}
)

func NewVerifyFight(alienRepository model.AlienRepository, cityRemover cityRemover, printer printer) *VerifyFight {
	return &VerifyFight{alienRepository: alienRepository, cityRemover: cityRemover, printer: printer}
}

func (f VerifyFight) VerifyFight(ctx context.Context, city *model.City) error {
	logger := appcontext.Logger(ctx).With(zap.String("city", city.Name()))
	logger.Debug("Verifying fight")

	if !city.HaveMultipleAliens() {
		logger.Debug("City doesn't have multiple aliens")
		return nil
	}

	aliensInCity := city.Aliens()
	logger = logger.With(zap.Any("aliens_in_city", maps.Keys(aliensInCity)))
	logger.Debug("City has multiple aliens")

	if err := f.printer.CityDestroyed(city, aliensInCity); err != nil {
		return errors.Wrap(err, "city destroyed but failed to print the message")
	}

	logger.Debug("Removing aliens")
	f.alienRepository.Remove(aliensInCity)

	f.cityRemover.Remove(ctx, city)

	logger.Debug("Fight processed")
	return nil
}
