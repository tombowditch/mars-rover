package rover

import (
	"errors"
	"fmt"
)

const (
	DirNorth = iota
	DirEast
	DirSouth
	DirWest
)

type Direction int

// converts Direction (int) to string
// 0 = N, 1 = E, 2 = S, 3 = W
func (d Direction) ConvertToString() (string, error) {
	switch d {
	case DirNorth:
		return "N", nil
	case DirEast:
		return "E", nil
	case DirSouth:
		return "S", nil
	case DirWest:
		return "W", nil
	default:
		return "", errors.New(fmt.Sprintf("unknown direction %d", d))
	}
}

func (d Direction) TurnRight() Direction {
	d += 1

	// if direction was 3 and is now 4, set to 0 (north) because there is no direction '4'
	if d > 3 {
		d = DirNorth
	}

	return d
}

func (d Direction) TurnLeft() Direction {
	d -= 1

	// if direction was 0 and is now -1, set to 3 (west) because there is no direction '-1'
	if d < 0 {
		d = DirWest
	}

	return d
}

// parses N,E,S,W to direction int 0,1,2,3 for reading rover starting location
func ParseStringToDirection(input string) (dir Direction, err error) {
	switch input {
	case "N":
		return DirNorth, nil
	case "E":
		return DirEast, nil
	case "S":
		return DirSouth, nil
	case "W":
		return DirWest, nil
	default:
		return 0, errors.New(fmt.Sprintf("unknown direction %s", input))
	}
}
