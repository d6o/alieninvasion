package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type (
	Config struct {
		Input   io.Reader
		Output  io.Writer
		Verbose bool
		Aliens  int
	}
)

const (
	filePermissions = 0o666
	numMinAliens    = 1
	defaultAliens   = 3
)

func Parse() (*Config, error) {
	input := flag.String("input", "", "Input file path. Defaults to stdin")
	output := flag.String("output", "", "Output file path. Defaults to stdout")
	aliens := flag.Int("aliens", defaultAliens, "Defines the number of aliens to be spawn")
	verbose := flag.Bool("v", false, "Verbose mode - enable logs")

	flag.Parse()

	cfg := &Config{
		Aliens:  *aliens,
		Input:   os.Stdin,
		Output:  os.Stdout,
		Verbose: *verbose,
	}

	if cfg.Aliens < numMinAliens {
		return nil, fmt.Errorf("invalid number of aliens provided - got %d but need at least %d", cfg.Aliens, numMinAliens)
	}

	var err error

	if *input != "" {
		cfg.Input, err = os.OpenFile(*input, os.O_RDONLY, filePermissions)
		if err != nil {
			return nil, err
		}
	}

	if *output != "" {
		cfg.Output, err = os.OpenFile(*output, os.O_RDWR|os.O_CREATE|os.O_TRUNC, filePermissions)
		if err != nil {
			return nil, err
		}
	}

	if cfg.Input == os.Stdin {
		return cfg, validateDataOnStdin()
	}

	return cfg, nil
}

func validateDataOnStdin() error {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return err
	}

	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return errors.New("there's no data being piped into stdin")
	}

	return nil
}
