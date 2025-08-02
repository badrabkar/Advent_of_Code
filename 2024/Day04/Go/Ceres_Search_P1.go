/*
--- Day 4: Ceres Search ---
--- Part One ---
"Looks like the Chief's not here. Next!" One of The Historians pulls out a device and pushes the only button on it. After a brief flash, you recognize the interior of the Ceres monitoring station!

As the search for the Chief continues, a small Elf who lives on the station tugs on your shirt; she'd like to know if you could help her with her word search (your puzzle input). She only has to find one word: X
MAS.

This word search allows words to be horizontal, vertical, diagonal, written backwards, or even overlapping other words. It's a little unusual, though, as you don't merely need to find one instance of XMAS - you
need to find all of them. Here are a few ways XMAS might appear, where irrelevant characters have been replaced with .:


..X...
.SAMX.
.A..A.
XMAS.S
.X....
The actual word search will be full of letters instead. For example:

MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
In this word search, XMAS occurs a total of 18 times; here's the same word search again, but where letters not involved in any XMAS have been replaced with .:

....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX
Take a look at the little Elf's word search. How many times does XMAS appear?


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

			if x + 3 < len(input[y]) && input[y][x] == byte('X') &&  input[y][x + 1] == byte('M') && input[y][x + 2] == byte('A') && input[y][x + 3] == byte('S') {
					match++
			}
			if x - 3 >= 0 && input[y][x] == byte('X') &&  input[y][x - 1] == byte('M') && input[y][x - 2] == byte('A') && input[y][x - 3] == byte('S') {
					match++
			}
			if  y - 3 >= 0 && input[y][x] == byte('X') &&  input[y - 1][x] == byte('M') && input[y -2][x] == byte('A') && input[y - 3][x] == byte('S') {
					match++
			}
			if  y + 3 < len(input) && input[y][x] == byte('X') &&  input[y + 1][x] == byte('M') && input[y + 2][x] == byte('A') && input[y + 3][x] == byte('S') {
					match++
			}

			if  y - 3 >= 0 && x + 3 < len(input[y]) && input[y][x] == byte('X') &&  input[y - 1][x + 1] == byte('M') && input[y - 2][x + 2] == byte('A') && input[y - 3][x + 3] == byte('S') {
					match++
			}

			if  y - 3 >= 0 && x - 3 >= 0  && input[y][x] == byte('X') &&  input[y - 1][x - 1] == byte('M') && input[y - 2][x - 2] == byte('A') && input[y - 3][x - 3] == byte('S') {
					match++
			}


			if  y + 3 < len(input) && x + 3 < len(input[y]) && input[y][x] == byte('X') &&  input[y + 1][x + 1] == byte('M') && input[y + 2][x + 2] == byte('A') && input[y + 3][x + 3] == byte('S') {
					match++
			}

			if  y + 3 < len(input) && x - 3 >= 0 && input[y][x] == byte('X') &&  input[y + 1][x - 1] == byte('M') && input[y + 2][x - 2] == byte('A') && input[y + 3][x - 3] == byte('S') {
					match++
			}

		}
	}
	fmt.Println(match)
}
