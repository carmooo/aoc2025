package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open("challenge/day2/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var text string
	for scanner.Scan() {
		text = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	rangesStrings := strings.Split(text, ",")
	ranges := [][]int{}
	for _, s := range rangesStrings {
		splits := strings.Split(s, "-")
		low, _ := strconv.Atoi(splits[0])
		high, _ := strconv.Atoi(splits[1])
		ranges = append(ranges, []int{low, high})
	}

	res := 0
	for _, rng := range ranges {
		cur := rng[1]
		for cur >= rng[0] {
			if cur < 10 {
				break
			}
			digits := int(math.Log10(float64(cur))) + 1
			div := int(math.Pow10(digits / 2))
			if cur/div == cur%div {
				res += cur
			}
			cur--
		}
	}
	fmt.Println(res)
}

func part2() {
	file, err := os.Open("challenge/day2/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var text string
	for scanner.Scan() {
		text = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	rangesStrings := strings.Split(text, ",")
	ranges := [][]int{}
	for _, s := range rangesStrings {
		splits := strings.Split(s, "-")
		low, _ := strconv.Atoi(splits[0])
		high, _ := strconv.Atoi(splits[1])
		ranges = append(ranges, []int{low, high})
	}

	res := 0
	for _, rng := range ranges {
		cur := rng[1]
		for cur >= rng[0] {
			if cur < 10 {
				break
			}
			digits := int(math.Log10(float64(cur))) + 1
			for n := 2; n <= digits; n++ {
				if digits%n != 0 {
					continue
				}
				div := int(math.Pow10(digits / n))
				tmp := cur
				cmp := tmp % div
				add := true
				for tmp > 0 {
					if tmp%div != cmp {
						add = false
						break
					}
					tmp /= div
				}

				if add {
					res += cur
					break
				}
			}
			cur--
		}
	}
	fmt.Println(res)
}
