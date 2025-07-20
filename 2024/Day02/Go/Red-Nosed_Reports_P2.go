/*
--- Part Two ---
The engineers are surprised by the low number of safe reports until they realize they forgot to tell you about the Problem Dampener.

The Problem Dampener is a reactor-mounted module that lets the reactor safety systems tolerate a single bad level in what would otherwise be a safe report. It's like the bad level never happened!

Now, the same rules apply as before, except if removing a single level from an unsafe report would make it safe, the report instead counts as safe.

More of the above example's reports are now safe:

7 6 4 2 1: Safe without removing any level.
1 2 7 8 9: Unsafe regardless of which level is removed.
9 7 6 2 1: Unsafe regardless of which level is removed.
1 3 2 4 5: Safe by removing the second level, 3.
8 6 4 4 1: Safe by removing the third level, 4.
1 3 6 7 9: Safe without removing any level.
Thanks to the Problem Dampener, 4 reports are actually safe!

Update your analysis by handling situations where the Problem Dampener can remove a single level from unsafe reports. How many reports are now safe?
*/
package main

import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"os"
)

func Abs(number int) (int) {
	if (number < 0) {
		return -number
	}
	return number
}

func isSafe(numbers []int) (int) {
	var diff		int
	var incr bool
	var decr bool

	for i := 0; i < len(numbers) - 1; i++ {

		diff = numbers[i + 1] - numbers[i] 

		if (diff > 0) {
			incr = true
		} else if (diff < 0) {
			decr = true
		}

		if !(Abs(diff) > 0 && Abs(diff) <= 3) || (incr && decr) {
				return 0
		}
	}
	return 1;
}

func StringsToInts(splited []string)  ([]int) {
	size := len(splited)
	numbers := []int{}

	for i := 0; i < size; i++ {
		temp, _ := strconv.Atoi(splited[i])
		numbers = append(numbers, temp)
	}

	return numbers
}


func brute_force(numbers []int) (int)  {

	for i := 0; i < len(numbers); i++ {
		number := make([]int, len(numbers))
		copy(number, numbers)
		if isSafe(append(number[:i], number[i+1:]...)) == 1 {
			return 1
		}
	}
	return 0
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result int
	var splited []string
	var numbers []int

	for scanner.Scan() {
		splited = strings.Split(scanner.Text(), " ")
		numbers = StringsToInts(splited) 
		result += brute_force(numbers)
	}
	fmt.Println(result)
}
