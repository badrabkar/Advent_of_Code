/*
--- Part Two ---
As you scan through the corrupted memory, you notice that some of the conditional statements are also still intact. If you handle some of the uncorrupted conditional statements in the program, you might be able to get an even more accurate result.

There are two new instructions you'll need to handle:

The do() instruction enables future mul instructions.
The don't() instruction disables future mul instructions.
Only the most recent do() or don't() instruction applies. At the beginning of the program, mul instructions are enabled.

For example:

xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))
This corrupted memory is similar to the example from before, but this time the mul(5,5) and mul(11,8) instructions are disabled because there is a don't() instruction before them. The other mul instructions function normally, including the one at the end that gets re-enabled by a do() instruction.

This time, the sum of the results is 48 (2*4 + 8*5).

Handle the new instructions; what do you get if you add up all of the results of just the enabled multiplications?
*/

package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"unicode"
	"strconv"
)

func CalculateProduct(line string) (int) {
	factors := strings.Split(line,",")
	if len(factors) != 2 {
		return 0
	}
	lfactor, _ := strconv.Atoi(factors[0])
	rfactor, _ := strconv.Atoi(factors[1])
	return lfactor * rfactor 
}

func Mull_It_Over(line string , product int) (int) {
	index := strings.Index(line,"mul(")
	rparenth, lparenth := index + 4, index + 3

	if index == -1 {
		return product 
	}
	
	for rparenth < len(line) && (line[rparenth] == ',' || unicode.IsDigit(rune(line[rparenth]))) {
		rparenth++;
	}

	if line[rparenth] == ')' {
		product += CalculateProduct(line[lparenth + 1:rparenth])
	}
	return Mull_It_Over(line[rparenth + 1:len(line)], product) 
}

func ParseEnabledInstructions(line string) (int) {
	product := 0 
	lines := strings.Split(line, "do()")

	index := 0
	for index < len(lines) {
		line = lines[index]
		disable := strings.Split(line,"don't()")
		product += Mull_It_Over(disable[0], 0)
		index++
	}

	return product
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result	:= 0
	var sb strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		sb.WriteString(line)
	}
	finalString := sb.String()
	result = ParseEnabledInstructions(finalString)	
	fmt.Println("result: ", result)
}












