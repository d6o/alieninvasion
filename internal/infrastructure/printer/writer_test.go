package printer

import (
	"bytes"
	"github.com/d6o/alieninvasion/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCityDestroyed(t *testing.T) {
	var buf bytes.Buffer
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := model.NewMockCityRepository(ctrl)

	writer := NewWriter(&buf, mockRepo)

	city, err := model.NewCity("TestCity")
	assert.NoError(t, err)

	alien1 := model.NewAlien(1)
	alien2 := model.NewAlien(2)

	aliens := map[int]*model.Alien{
		alien1.ID(): alien1,
		alien2.ID(): alien2,
	}

	err = writer.CityDestroyed(city, aliens)
	assert.Nil(t, err)
	assert.Contains(t, buf.String(), "TestCity has been destroyed by Aliens:")
	assert.Contains(t, buf.String(), "1")
	assert.Contains(t, buf.String(), "2")
}

func TestWorldMap(t *testing.T) {
	var buf bytes.Buffer
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := model.NewMockCityRepository(ctrl)
	writer := NewWriter(&buf, mockRepo)

	city1, err := model.NewCity("City1")
	assert.NoError(t, err)
	city2, err := model.NewCity("City2")
	assert.NoError(t, err)

	mockRepo.EXPECT().All().Return(map[string]*model.City{
		city1.Name(): city1,
		city2.Name(): city2,
	})

	err = writer.WorldMap()
	assert.Nil(t, err)
	assert.Contains(t, buf.String(), "City1")
	assert.Contains(t, buf.String(), "City2")
}
