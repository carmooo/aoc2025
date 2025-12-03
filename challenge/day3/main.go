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
	file, err := os.Open("challenge/day3/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	res := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var biggest, big uint8
		var idx int
		// first pass
		for i := 0; i < len(line)-1; i++ {
			if line[i] > biggest {
				biggest = line[i]
				idx = i
			}
			if line[i]-'0' == '9' {
				break
			}
		}
		// second pass
		for i := idx + 1; i < len(line); i++ {
			if line[i] > big {
				big = line[i]
			}
		}
		res += int((biggest-'0')*10 + big - '0')
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func part2() {
	file, err := os.Open("challenge/day3/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	res := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		idx := -1
		var n int
		for rem := 11; rem >= 0; rem-- {
			var biggest uint8
			for i := idx + 1; i < len(line)-rem; i++ {
				if line[i] > biggest {
					biggest = line[i]
					idx = i
				}
				if line[i]-'0' == '9' {
					break
				}
			}
			n = n*10 + int(biggest-'0')
		}

		res += n
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(res)
}
