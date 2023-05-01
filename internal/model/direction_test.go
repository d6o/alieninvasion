package model_test

import (
	"github.com/d6o/alieninvasion/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToDirection(t *testing.T) {
	tt := []struct {
		name          string
		input         string
		expected      model.Direction
		expectedError bool
	}{
		{"North", "north", model.North, false},
		{"South", "south", model.South, false},
		{"East", "east", model.East, false},
		{"West", "west", model.West, false},
		{"Invalid", "invalid", model.North, true},
		{"MixedCase", "EaSt", model.East, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := model.ToDirection(tc.input)
			assert.Equal(t, tc.expected, actual)
			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestOpposite(t *testing.T) {
	tt := []struct {
		name     string
		input    model.Direction
		expected model.Direction
	}{
		{"North", model.North, model.South},
		{"South", model.South, model.North},
		{"East", model.East, model.West},
		{"West", model.West, model.East},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.input.Opposite()
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestString(t *testing.T) {
	tt := []struct {
		name     string
		input    model.Direction
		expected string
	}{
		{"North", model.North, "north"},
		{"South", model.South, "south"},
		{"East", model.East, "east"},
		{"West", model.West, "west"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.input.String()
			assert.Equal(t, tc.expected, actual)
		})
	}
}
