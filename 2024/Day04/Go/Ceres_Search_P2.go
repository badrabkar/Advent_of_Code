/*
--- Day 4: Ceres Search ---
--- Part Two ---

The Elf looks quizzically at you. Did you misunderstand the assignment?

Looking for the instructions, you flip over the word search to find that this isn't actually an XMAS puzzle; it's an X-MAS puzzle in which you're supposed to find two MAS in the shape of an X. One way to achieve that is like this:

M.S
.A.
M.S

Irrelevant characters have again been replaced with . in the above diagram. Within the X, each MAS can be written forwards or backwards.

Here's the same example from before, but this time all of the X-MASes have been kept instead:

.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........

In this example, an X-MAS appears 9 times.

Flip the word search from the instructions back over to the word search side and try again. How many times does an X-MAS appear?

*/

package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	var match int
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {

			if x - 1 >= 0 && y - 1 >= 0 && y + 1 < len(input) && x + 1 < len(input[y]) && input[y][x] == byte('A') && ((input[y - 1][x + 1] == byte('M') && input[y + 1][x - 1] == byte('S')) || (input[y - 1][x + 1] == byte('S') && input[y + 1][x - 1] == byte('M'))) && ((input[y - 1][x - 1] == byte('M') && input[y + 1][x + 1] == byte('S')) || (input[y - 1][x - 1] == byte('S') && input[y + 1][x + 1] == byte('M'))) {
				match++
			}
		}
	}
	fmt.Println(match)
}
