package app

import (
	"context"
	"io"
	"time"

	"github.com/d6o/alieninvasion/internal/appcontext"
	"github.com/d6o/alieninvasion/internal/infrastructure/log"
	"github.com/d6o/alieninvasion/internal/infrastructure/printer"
	"github.com/d6o/alieninvasion/internal/infrastructure/storage/memory"
	"github.com/d6o/alieninvasion/internal/service"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type App struct {
	aliens  int
	input   io.Reader
	output  io.Writer
	verbose bool
}

func NewApp(aliens int, input io.Reader, output io.Writer, verbose bool) *App {
	return &App{aliens: aliens, input: input, output: output, verbose: verbose}
}

const (
	roundsMax = 10000
)

func (a *App) Run(ctx context.Context) error {
	baseLogger, err := log.NewZap(a.verbose)
	if err != nil {
		return errors.Wrap(err, "can't create a baseLogger")
	}

	defer func(baseLogger *zap.Logger) {
		_ = baseLogger.Sync()
	}(baseLogger)

	ctx = appcontext.WithLogger(ctx, baseLogger)

	cityRepository := memory.NewCityRepository()
	alienRepository := memory.NewAlienRepository()

	outputPrinter := printer.NewWriter(a.output, cityRepository)

	randSeed := time.Now().UnixNano()

	cityRandomizerService := service.NewRandomCityFromList(randSeed, cityRepository)
	cityNearbyRandomizerService := service.NewRandomCityFromCity(randSeed)
	cityRemoverService := service.NewRemoveCity(cityRepository)
	verifyFight := service.NewVerifyFight(alienRepository, cityRemoverService, outputPrinter)
	createCitiesService := service.NewCreateCities(cityRepository)
	createAliensService := service.NewCreateAliens(alienRepository, cityRandomizerService, verifyFight)
	alienVerifyService := service.NewVerifyAliens(alienRepository)
	moveAliensService := service.NewMoveAliens(alienRepository, cityNearbyRandomizerService, verifyFight)

	if err := createCitiesService.ParseFromReader(ctx, a.input); err != nil {
		return errors.Wrap(err, "can't create cities from the provided input")
	}

	if err := createAliensService.Create(ctx, a.aliens); err != nil {
		return errors.Wrap(err, "can't create aliens")
	}

	for i := 1; i <= roundsMax; i++ {
		logger := baseLogger.With(zap.Int("round", i))
		ctx := appcontext.WithLogger(ctx, logger)

		logger.Debug("Starting new round")

		if !alienVerifyService.AreEnoughAliensAround(ctx) {
			return outputPrinter.WorldMap()
		}

		if err := moveAliensService.Move(ctx); err != nil {
			return errors.Wrap(err, "error while moving aliens")
		}

		logger.Debug("Round finished")
	}

	baseLogger.Debug("Max rounds reached")
	return outputPrinter.WorldMap()
}
