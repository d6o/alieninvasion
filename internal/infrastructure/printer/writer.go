package printer

import (
	"fmt"
	"io"

	"github.com/d6o/alieninvasion/internal/model"
	"github.com/pkg/errors"
)

type Writer struct {
	writer         io.Writer
	cityRepository model.CityRepository
}

func NewWriter(writer io.Writer, cityRepository model.CityRepository) *Writer {
	return &Writer{writer: writer, cityRepository: cityRepository}
}

func (w Writer) CityDestroyed(city *model.City, aliens map[int]*model.Alien) error {
	_, err := fmt.Fprintf(w.writer, "%s has been destroyed by Aliens: ", city.Name())
	if err != nil {
		return errors.Wrap(err, "can't write message to output writer")
	}

	ids := make([]int, 0, len(aliens))
	for id := range aliens {
		ids = append(ids, id)
	}

	_, err = fmt.Fprintln(w.writer, ids)
	return errors.Wrap(err, "can't write message to output writer")
}

func (w Writer) WorldMap() error {
	for _, city := range w.cityRepository.All() {
		_, err := fmt.Fprintln(w.writer, city.String())
		if err != nil {
			return err
		}
	}

	return nil
}
