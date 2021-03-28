package rover

import (
	"fmt"
	"testing"
)

type RoverTest struct {
	Name           string
	Input          string
	ExpectedOutput string
}

func TestRover(t *testing.T) {
	tests := []RoverTest{
		{
			Name:           "provided test case",
			Input:          "5 5\n1 2 N\nLMLMLMLMM\n3 3 E\nMMRMMRMRRM",
			ExpectedOutput: "1 3 N\n5 1 E",
		},
		{
			Name:           "invalid y position",
			Input:          "5 5\n1 2 N\nLMLMLMLMM\n3 3e E\nMMRMMRMRRM\n1 1 W\nRRMMMLLLMMMMM",
			ExpectedOutput: "invalid rover y position",
		},
		{
			Name:           "invalid x position",
			Input:          "5 5\n1 2 N\nLMLMLMLMM\ne 3 E\nMMRMMRMRRM\n1 1 W\nRRMMMLLLMMMMM",
			ExpectedOutput: "invalid rover x position",
		},
		{
			Name:           "bad plateau input",
			Input:          "5 5 7\n1 2 N\nLMLMLMLMM\n3 3 E\nMMRMMRMRRM",
			ExpectedOutput: "invalid plateau input",
		},
		{
			Name:           "not enough lines",
			Input:          "5 5 7\n1 2 N",
			ExpectedOutput: "atleast 3 input lines required to function",
		},
		{
			Name:           "attempt to go off edge of map",
			Input:          "5 5\n0 0 N\nMMMMMMMMMRMMMMMMMMMM",
			ExpectedOutput: "5 5 E",
		},
		{
			Name:           "attempt to go off edge of map negative",
			Input:          "5 5\n0 0 N\nRRMMMMMMRMMMMMMMM",
			ExpectedOutput: "0 0 W",
		},
		{
			Name:           "full rotation clockwise",
			Input:          "5 5\n0 0 N\nRRRR",
			ExpectedOutput: "0 0 N",
		},
		{
			Name:           "full rotation anti-clockwise",
			Input:          "5 5\n0 0 N\nLLLL",
			ExpectedOutput: "0 0 N",
		},
		{
			Name:           "invalid rover instruction",
			Input:          "5 5\n0 0 N\nMMMQ",
			ExpectedOutput: "unknown instruction Q",
		},
		{
			Name:           "start rover outside of plateau (X)",
			Input:          "5 5\n6 0 N\nMMM",
			ExpectedOutput: "rover would start outside of plateau",
		},
		{
			Name:           "start rover outside of plateau (Y)",
			Input:          "5 5\n0 6 N\nMMM",
			ExpectedOutput: "rover would start outside of plateau",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			fmt.Printf("=====================\nRunning test: %s\nINPUT:\n%s\nEXPECTED OUTPUT:\n%s\n", test.Name, test.Input, test.ExpectedOutput)
			out, err := Start(test.Input)
			if err != nil {
				// set out to the error in this case for output & comparison below
				out = err.Error()
			}

			fmt.Printf("OUTPUT:\n%s\n", out)

			if out != test.ExpectedOutput {
				t.Errorf("got %s, wanted %s", out, test.ExpectedOutput)
			}
		})
	}
}
