package service

import (
	"context"
	"io"
	"strings"

	"github.com/d6o/alieninvasion/internal/appcontext"
	"github.com/d6o/alieninvasion/internal/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	newLine     = "\n"
	positionSep = "="
	lineSep     = " "
	minCities   = 1
)

type (
	CreateCities struct {
		cityRepository model.CityRepository
	}
)

func NewCreateCities(cityRepository model.CityRepository) *CreateCities {
	return &CreateCities{cityRepository: cityRepository}
}

const (
	validPositionsPart = 2
)

func (c CreateCities) ParseFromReader(ctx context.Context, reader io.Reader) error {
	logger := appcontext.Logger(ctx)
	logger.Debug("Creating cities from provided string")

	input, err := io.ReadAll(reader)
	if err != nil {
		return errors.Wrap(err, "can't read world map from input")
	}

	worldMap := strings.Trim(string(input), newLine)

	for _, line := range strings.Split(worldMap, newLine) {
		parts := strings.Split(line, lineSep)

		name := strings.TrimSpace(parts[0])
		positions := parts[1:]
		logger.With(zap.String("city_name", name), zap.Any("city_positions", positions)).Debug("Processing city")

		if len(name) == 0 {
			return errors.New("a city needs to have a name")
		}

		city, err := c.cityRepository.GetOrAdd(name)
		if err != nil {
			logger.With(zap.Error(err), zap.String("city_name", name)).Debug("Can't create city")
			continue
		}

		for _, position := range positions {
			parts = strings.Split(position, positionSep)

			if len(parts) != validPositionsPart {
				logger.With(zap.String("city_name", name), zap.Any("position", parts)).Debug("City position is invalid, discarding")
				continue
			}

			cityDirection := strings.TrimSpace(parts[0])
			cityDestination := strings.TrimSpace(parts[1])

			otherCity, err := c.cityRepository.GetOrAdd(cityDestination)
			if err != nil {
				logger.With(zap.Error(err), zap.String("city_name", name)).Debug("Can't create city")
				continue
			}

			logger.With(
				zap.String("city", name),
				zap.String("destination_city", otherCity.Name()),
				zap.String("direction", cityDirection),
			).Debug("Linking cities")

			direction, err := model.ToDirection(cityDirection)
			if err != nil {
				logger.With(zap.Error(err), zap.String("city_direction", cityDirection)).Debug("Can't parse direction")
				continue
			}

			city.SetCity(otherCity, direction)
		}
	}

	qtyCities := len(c.cityRepository.All())
	logger.With(zap.Int("qty_cities", qtyCities)).Debug("Cities created")

	if qtyCities < minCities {
		return errors.New("not enough cities created - bad map file")
	}

	return nil
}
