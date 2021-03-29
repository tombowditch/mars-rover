package rover

import (
	"fmt"
)

type rover struct {
	planet    *plateau
	positionX int
	positionY int
	direction Direction
}

func (r *rover) Move() error {
	switch r.direction {
	case DirNorth:
		r.positionY = max(r.positionY+1, r.planet.maxY)
		break
	case DirEast:
		r.positionX = max(r.positionX+1, r.planet.maxX)
		break
	case DirSouth:
		r.positionY = min(r.positionY-1, 0)
		break
	case DirWest:
		r.positionX = min(r.positionX-1, 0)
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
