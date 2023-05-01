package memory_test

import (
	"testing"

	"github.com/d6o/alieninvasion/internal/infrastructure/storage/memory"
	"github.com/d6o/alieninvasion/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestNewCityRepository(t *testing.T) {
	repo := memory.NewCityRepository()
	assert.NotNil(t, repo)
	assert.NotNil(t, repo.All())
}

func TestAll(t *testing.T) {
	repo := memory.NewCityRepository()

	assert.Equal(t, 0, len(repo.All()))

	city, _ := model.NewCity("TestCity")
	repo.Add(city)

	assert.Equal(t, 1, len(repo.All()))
}

func TestAllNames(t *testing.T) {
	repo := memory.NewCityRepository()
	city, _ := model.NewCity("TestCity")
	repo.Add(city)

	names := repo.AllNames()
	assert.Equal(t, 1, len(names))
	assert.Equal(t, "TestCity", names[0])
}

func TestRemove(t *testing.T) {
	repo := memory.NewCityRepository()
	city, _ := model.NewCity("TestCity")
	repo.Add(city)

	repo.Remove("TestCity")
	all := repo.All()
	assert.Equal(t, 0, len(all))
}

func TestGetOrAdd(t *testing.T) {
	repo := memory.NewCityRepository()

	// Test adding a new city
	city, err := repo.GetOrAdd("TestCity")
	assert.Nil(t, err)
	assert.NotNil(t, city)
	assert.Equal(t, "TestCity", city.Name())

	// Test getting an existing city
	existingCity, err := repo.GetOrAdd("TestCity")
	assert.Nil(t, err)
	assert.Equal(t, city, existingCity)
}

func TestGet(t *testing.T) {
	repo := memory.NewCityRepository()
	city, _ := model.NewCity("TestCity")
	repo.Add(city)

	// Test getting an existing city
	existingCity, err := repo.Get("TestCity")
	assert.Nil(t, err)
	assert.Equal(t, city, existingCity)

	// Test getting a non-existing city
	_, err = repo.Get("NonExistentCity")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "city doesn't exist")
}

func TestAdd(t *testing.T) {
	repo := memory.NewCityRepository()
	city, _ := model.NewCity("TestCity")

	repo.Add(city)
	all := repo.All()
	assert.Equal(t, 1, len(all))
	assert.Equal(t, city, all["TestCity"])
}
