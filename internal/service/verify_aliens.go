package service

import (
	"context"

	"github.com/d6o/alieninvasion/internal/appcontext"
	"github.com/d6o/alieninvasion/internal/model"
	"go.uber.org/zap"
)

const (
	minAliensAlive = 1
)

type (
	VerifyAliens struct {
		alienRepository model.AlienRepository
	}
)

func NewVerifyAliens(alienRepository model.AlienRepository) *VerifyAliens {
	return &VerifyAliens{alienRepository: alienRepository}
}

func (v VerifyAliens) AreEnoughAliensAround(ctx context.Context) bool {
	aliens := v.alienRepository.All()
	aliensAlive := len(aliens)

	logger := appcontext.Logger(ctx)
	logger.With(zap.Int("aliens_alive", aliensAlive)).Debug("Verifying if there are enough aliens alive")

	if aliensAlive < minAliensAlive {
		logger.With(
			zap.Int("aliens_alive", aliensAlive),
			zap.Int("min_aliens_alive", minAliensAlive),
		).Debug("Not enough aliens are alive")
		return false
	}

	logger.Debug("Verifying if there are aliens capable of moving")
	for _, alien := range aliens {
		if !alien.IsTrapped() {
			logger.With(zap.Int("alien_id", alien.ID())).Debug("Alien can move")
			return true
		}
		logger.With(zap.Int("alien_id", alien.ID())).Debug("Alien is trapped")
	}

	logger.Debug("All aliens are trapped")
	return false
}
