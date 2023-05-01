package service

import (
	"context"

	"github.com/d6o/alieninvasion/internal/appcontext"
	"github.com/d6o/alieninvasion/internal/model"
	"go.uber.org/zap"
)

type RemoveCity struct {
	cityRepository model.CityRepository
}

func NewRemoveCity(cityRepository model.CityRepository) *RemoveCity {
	return &RemoveCity{cityRepository: cityRepository}
}

func (r RemoveCity) Remove(ctx context.Context, city *model.City) {
	logger := appcontext.Logger(ctx).With(zap.String("city", city.Name()))
	logger.Debug("Removing city")

	for direction, destinationCity := range city.Destinations() {
		destDirection := direction.Opposite()
		logger.With(
			zap.String("dest_city", destinationCity.Name()),
			zap.String("direction", destDirection.String()),
		).Debug("Removing link between cities")

		destinationCity.RemoveDirection(direction.Opposite())
		city.RemoveDirection(direction)
	}

	r.cityRepository.Remove(city.Name())

	logger.Debug("City removed")
}
