package rover

import (
	"errors"
	"fmt"
	"strings"
)

func Start(input string) (string, error) {
	out := ""
	inputLines := strings.Split(input, "\n")

	// atleast 3 lines required to parse a proper input
	if len(inputLines) < 3 {
		return "", errors.New("atleast 3 input lines required to function")
	}

	// take first line of input and parse plateau size with it
	planet, err := ParsePlateauSize(inputLines[0])
	if err != nil {
		return "", err
	}

	// take the rest of the input (minus line 1 which is plateau size) and parse rover inputs from it
	roverInputs := ParseRoverInput(inputLines[1:])

	// for every rover input...
	for _, roverInput := range roverInputs {
		roverLocation := roverInput[0]
		roverInstructions := roverInput[1]

		// parse rover starting location
		x, y, dir, err := ParseRoverLocation(roverLocation)
		if err != nil {
			return "", err
		}

		// make rover struct with all info
		roverPlanet := rover{
			planet:    &planet,
			positionX: x,
			positionY: y,
			direction: dir,
		}

		// for every instruction...
		for _, instruction := range strings.Split(roverInstructions, "") {
			// execute instruction
			err := roverPlanet.ExecuteInstruction(instruction)
			if err != nil {
				return "", err
			}
		}

		// get the ending direction
		roverDir, err := roverPlanet.direction.ConvertToString()
		if err != nil {
			return "", err
		}

		// add it to the final output
		out += fmt.Sprintf("%d %d %s\n", roverPlanet.positionX, roverPlanet.positionY, roverDir)
	}

	// return the final output, minus the last line (trailing \n)
	return out[:len(out)-1], nil
}
