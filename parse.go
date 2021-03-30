package rover

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type plateau struct {
	maxX int
	maxY int
}

func ParsePlateauSize(rawInput string) (plateau, error) {
	parsedCoords := strings.Split(rawInput, " ")

	// do we either have too many or too little arguments
	if len(parsedCoords) != 2 {
		return plateau{}, errors.New(fmt.Sprintf("invalid plateau input"))
	}

	// parse x to int
	coordX, err := strconv.Atoi(parsedCoords[0])
	if err != nil {
		return plateau{}, fmt.Errorf("invalid X coordinate %s", parsedCoords[0])
	}

	// parse y to int
	coordY, err := strconv.Atoi(parsedCoords[1])
	if err != nil {
		return plateau{}, fmt.Errorf("invalid Y coordinate %s", parsedCoords[1])
	}

	// check if plateau size is negative, also check we have a plateau (i.e. coordX and coordY isnt 0)
	if coordX < 1 && coordY < 1 {
		return plateau{}, fmt.Errorf("plateau coordinates must be positive")
	}

	return plateau{maxX: coordX, maxY: coordY}, nil
}

func ParseRoverInputs(planet plateau, inputLines []string) ([]rover, error) {
	// inputLines is the input, minus the first line (plateau size)
	roverConfiguration := make([]string, 0)
	parsed := make([][]string, 0)
	rovers := make([]rover, 0)

	for _, line := range inputLines {
		// if roverConfiguration len == 2, append to output & reset
		if len(roverConfiguration) == 2 {
			parsed = append(parsed, roverConfiguration)
			roverConfiguration = make([]string, 0)
		}

		// append to roverConfiguration to build
		roverConfiguration = append(roverConfiguration, line)
	}

	// double check if roverConfiguration has another set of instructions before we return
	if len(roverConfiguration) == 2 {
		parsed = append(parsed, roverConfiguration)
	}

	// parse rover configurations to rovers
	for _, r := range parsed {
		roverInstructionsRaw := r[1] // r[1] is our 2nd input line of rover, e.g. MMMMRMMMMLMMMM
		x, y, dir, err := ParseRoverLocation(r[0])
		if err != nil {
			return []rover{}, err
		}

		// check to see if rover starts outside of plateau
		if x > planet.maxX || y > planet.maxY || x < 0 || y < 0 {
			return []rover{}, errors.New("rover would start outside of plateau")
		}

		roverInstructions := strings.Split(roverInstructionsRaw, "")

		rovers = append(rovers, rover{
			planet:       &planet,
			positionX:    x,
			positionY:    y,
			direction:    dir,
			instructions: roverInstructions,
		})

	}

	return rovers, nil
}

func ParseRoverLocation(input string) (x int, y int, dir Direction, err error) {
	splitInput := strings.Split(input, " ")

	// rover split inputs should have 3 elements, 'X Y D' e.g. 0 0 N
	if len(splitInput) != 3 {
		return 0, 0, 0, errors.New("invalid rover position")
	}

	x, err = strconv.Atoi(splitInput[0])
	if err != nil {
		return 0, 0, 0, errors.New("invalid rover x position")
	}

	y, err = strconv.Atoi(splitInput[1])
	if err != nil {
		return 0, 0, 0, errors.New("invalid rover y position")
	}

	dir, err = ParseStringToDirection(splitInput[2])
	if err != nil {
		return 0, 0, 0, errors.New("invalid rover direction")
	}

	return x, y, dir, nil
}
