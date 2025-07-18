package main

import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"os"
)

func getSafety(numbers []int) (int) {
	var number  int
	var current int
	var diff		int

	number = numbers[0]
	for i := 1; i < len(numbers); i++ {
		current = numbers[i]
		diff = current - number
		if diff > 3 || diff <= 0 {
				return 0
		}
		number = current
	}
	return 1;
}

func get_numbers_array(splited []string)  ([]int) {
	size := len(splited)
	begin, _ := strconv.Atoi(splited[0])
	last, _ := strconv.Atoi(splited[size - 1])
	numbers := []int{}

	if (begin > last) {
		for i := size - 1; i >= 0; i-- {
			temp, _ := strconv.Atoi(splited[i])
			numbers = append(numbers, temp)
		}
	} else {
		for i := 0; i < size; i++ {
			temp, _ := strconv.Atoi(splited[i])
			numbers = append(numbers, temp)
		}
	}
	return numbers
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
		numbers = get_numbers_array(splited) 
		result += getSafety(numbers)
	}
	fmt.Println(result)
}
