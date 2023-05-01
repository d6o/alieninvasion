package service_test

import (
	"github.com/d6o/alieninvasion/internal/model"
	"github.com/d6o/alieninvasion/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRandomCityFromCity(t *testing.T) {
	r := service.NewRandomCityFromCity(1)

	if r == nil {
		t.Error("NewRandomCityFromCity() should not return nil")
	}
}

func TestCityFromCity(t *testing.T) {
	r := service.NewRandomCityFromCity(1)

	// Test case with a single destination
	city1, err := model.NewCity("City1")
	assert.NoError(t, err)
	city2, err := model.NewCity("City2")
	assert.NoError(t, err)

	city1.SetCity(city2, model.North)

	result := r.CityFromCity(city1)
	if result.Name() != "City2" {
		t.Errorf("Expected City2, got %s", result.Name())
	}

	// Test case with multiple destinations
	city3, err := model.NewCity("City3")
	assert.NoError(t, err)

	city2.SetCity(city3, model.West)

	// Run test multiple times to ensure randomness
	for i := 0; i < 10; i++ {
		result := r.CityFromCity(city2)
		if result.Name() != "City1" && result.Name() != "City3" {
			t.Errorf("Expected CityA or CityC, got %s", result.Name())
		}
	}
}
