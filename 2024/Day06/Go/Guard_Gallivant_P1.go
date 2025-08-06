package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type guard struct {
	x int
	y int
}

type location struct {
	x int
	y int
}

func find_the_path(matrix []string, guard guard, width int, height int) {

	direction := 0
	var steps int
	y := guard.y
	x := guard.x
	locations := make(map[location]bool)
	location := location{x:guard.x,y:guard.y}
	locations[location] = true
	fmt.Println( x, y)
	for (x < width && x > 0) || (y < height && y > 0) {
		location.x = x
		location.y = y
		locations[location] = true
		if (matrix[y][x - 1] == '#' && direction == 3) || (matrix[y][x + 1] == '#' && direction == 1) || (matrix[y - 1][x] == '#' && direction == 0) || (matrix[y + 1][x] == '#' && direction == 2) {
			direction = (direction + 1) % 4
		} 
		if direction == 0 {
			y--
			if y == 0  {
				break
			}
		} else if direction == 1 {
			x++
			if  x == width - 1 {
				break
			}
		} else if direction == 2 {
			y++
			if y == height - 1 {
				break;
			}
		} else {
			x--
			if  x == 0 {
				break
			}
		}
		
		fmt.Println("x - y", x, y)
		steps++
	}
	location.x = x
	location.y = y
	locations[location] = true
	fmt.Println("steps: ", steps)
	fmt.Println(len(locations))
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	var matrix []string
	guard := guard{x:0, y:0}
	var x,y int

	for scanner.Scan() {
		line := scanner.Text()
		x = strings.IndexRune(line, '^')
		if  x != -1 {
			guard.x = x
			guard.y = y
		}
		matrix = append(matrix, line)
		y++
	}
	x = len(matrix[0])
	find_the_path(matrix, guard, x, y)
}
