package model

import (
	"errors"
	"strings"
)

type Direction string

func ToDirection(s string) (Direction, error) {
	switch strings.ToLower(s) {
	case string(North):
		return North, nil
	case string(South):
		return South, nil
	case string(East):
		return East, nil
	case string(West):
		return West, nil
	}

	return North, errors.New("invalid Direction")
}

const (
	North Direction = "north"
	South Direction = "south"
	East  Direction = "east"
	West  Direction = "west"
)

var oppositeDirection = map[Direction]Direction{ //nolint:gochecknoglobals // no need to recreate on every object.
	North: South,
	South: North,
	East:  West,
	West:  East,
}

func (d Direction) Opposite() Direction {
	return oppositeDirection[d]
}

func (d Direction) String() string {
	return string(d)
}
