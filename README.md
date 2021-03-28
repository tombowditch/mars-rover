# mars-rover

`mars-rover` takes a set of instructions (inside `planet_test.go`) to navigate the plateau (Mars). Rovers instructions are ran one by one, in supplied order.
`mars-rover` takes a minimum input of three lines, of which the first one is the plateau size (`X Y`, e.g. `5 5` for a 5x5 grid). The second and third can be repeated indefinitely, of which they are the rover starting location: `X Y D` e.g. `3 3 N` and the instructions for the rover to take: a string of characters either `M`, `L` or `R` e.g. `MMMRRMMMLLMM`

There is checks to make sure the rover can not go out of bounds of the plateau and checks to make sure you cannot start the rover outside of the plateau.
Future improvements could include output logging at each stage to see exactly where and how the rover is moving and also an interface to display this. Additionally, it may be helpful to have warning logs if the rover tries to go off the edge of the plateau.

I chose to write the mars rover in Go because of the excellent unit testing package and the overall readability of the code.

## Example input/output

Input:

```
5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
```

Output:

```
1 3 N
5 1 E
```

## Running

To run unit tests defined in `planet_test.go`, run `go test`

## Testing

Automated tests with Linux on Go 1.16.x are ran every commit with Github Actions, available [here](https://github.com/tombowditch/mars-rover/actions)
