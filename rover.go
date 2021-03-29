package rover

import (
	"fmt"
)

type rover struct {
	planet    *plateau
	positionX int
	positionY int
	direction Direction
	actions   []string
}

func (r *rover) Move() error {
	switch r.direction {
	case DirNorth:
		r.positionY = boundsCheck(r.positionY+1, 0, r.planet.maxY)
		break
	case DirEast:
		r.positionX = boundsCheck(r.positionX+1, 0, r.planet.maxX)
		break
	case DirSouth:
		r.positionY = boundsCheck(r.positionY-1, 0, r.planet.maxY)
		break
	case DirWest:
		r.positionX = boundsCheck(r.positionX-1, 0, r.planet.maxX)
		break
	default:
		return fmt.Errorf("unknown direction %d", r.direction)
	}

	return nil
}

// executes a M,L,R instruction, returns error if instruction invalid
func (r *rover) ExecuteInstruction(instruction string) error {
	switch instruction {
	case "M":
		return r.Move()
	case "L":
		r.direction = r.direction.TurnLeft()
		break
	case "R":
		r.direction = r.direction.TurnRight()
		break
	default:
		return fmt.Errorf("unknown instruction %s", instruction)
	}

	return nil
}
