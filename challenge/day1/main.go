package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open("challenge/day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	cur := 50
	res := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		dir := line[0:1]
		step, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		switch dir {
		case "R":
			cur += step
			for cur > 99 {
				cur -= 100
			}
		case "L":
			cur -= step
			for cur < 0 {
				cur += 100
			}
		}

		if cur == 0 {
			res++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func part2() {
	file, err := os.Open("challenge/day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	cur := 50
	res := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		dir := line[0:1]
		step, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		switch dir {
		case "R":
			for step > 0 {
				step--
				cur++
				if cur == 100 {
					cur = 0
					res++
				}
			}
		case "L":
			for step > 0 {
				step--
				cur--
				if cur == 0 {
					res++
				}
				if cur == -1 {
					cur = 99
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(res)
}
