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
	rovers, err := ParseRoverInputsToRovers(planet, inputLines[1:])
	if err != nil {
		return "", err
	}

	// for every rover input...
	for _, r := range rovers {
		////////////
		// for every instruction...
		for _, instruction := range r.actions {
			// execute instruction
			err := r.ExecuteInstruction(instruction)
			if err != nil {
				return "", err
			}
		}

		// get the ending direction
		roverDir, err := r.direction.ConvertToString()
		if err != nil {
			return "", err
		}

		// add it to the final output
		out += fmt.Sprintf("%d %d %s\n", r.positionX, r.positionY, roverDir)
	}

	// return the final output, minus the last line (trailing \n)
	return out[:len(out)-1], nil
}
