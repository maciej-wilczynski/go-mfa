package argparser

import (
	"errors"
	"os"
)

type ArgParser struct {
	Pin           string
	TargetProfile string
}

func New(args []string) (ArgParser, error) {
	a := ArgParser{}

	if len(os.Args) < 3 {
		return a, errors.New("not enough arguments")
	}

	a.Pin = args[2]
	a.TargetProfile = args[1]

	return a, nil
}
