package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open("challenge/day4/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	dirs := [][]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
	res := 0
	coords := [][]int{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] != '@' {
				continue
			}
			nNeighbours := 0
			for _, dir := range dirs {
				newY, newX := y+dir[0], x+dir[1]
				if newY < 0 || newY >= len(grid) {
					continue
				}
				if newX < 0 || newX >= len(grid[newY]) {
					continue
				}
				if grid[newY][newX] == '@' {
					nNeighbours++
					if nNeighbours > 4 {
						break
					}
				}
			}
			if nNeighbours < 4 {
				coords = append(coords, []int{y, x})
				res++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func part2() {
	file, err := os.Open("challenge/day4/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	var grid, replacement []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
		replacement = append(replacement, line)
	}

	dirs := [][]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
	total := 0
	coords := [][]int{}
	for {
		sub := 0
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				if grid[y][x] != '@' {
					continue
				}
				nNeighbours := 0
				for _, dir := range dirs {
					newY, newX := y+dir[0], x+dir[1]
					if newY < 0 || newY >= len(grid) {
						continue
					}
					if newX < 0 || newX >= len(grid[newY]) {
						continue
					}
					if grid[newY][newX] == '@' {
						nNeighbours++
						if nNeighbours > 4 {
							break
						}
					}
				}
				if nNeighbours < 4 {
					coords = append(coords, []int{y, x})
					sub++
					replacement[y] = replacement[y][:x] + "." + replacement[y][x+1:]
				}
			}
		}
		grid = replacement
		total += sub
		if sub == 0 {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(total)
}
