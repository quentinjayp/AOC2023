package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var bag = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	part_one()
	part_two()
}

func part_one() {
	fmt.Println("part one")
	readFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		id, err := process_text(fileScanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		sum = sum + id
	}
	fmt.Println(sum)
	readFile.Close()
}

func part_two() {
	fmt.Println("part two")

	readFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		power := get_powers(fileScanner.Text())

		sum = sum + power
	}
	fmt.Println(sum)
	readFile.Close()
}

func process_text(line string) (id int, err error) {
	line = strings.TrimLeft(line, "Game ")
	tokens := strings.Split(line, ":")
	id, err = strconv.Atoi(tokens[0])
	if err != nil {
		return
	}
	rounds := strings.Split(tokens[1], ";")
	valid := validate_game(rounds)

	if valid {
		return
	} else {
		id = 0
	}
	return
}

func validate_game(rounds []string) (valid bool) {
	for i := 0; i < len(rounds); i++ {
		colours := strings.Split(rounds[i], ",")
		for j := 0; j < len(colours); j++ {
			colour_result := strings.TrimSpace(colours[j])
			result := strings.Split(colour_result, " ")
			num, err := strconv.Atoi(result[0])
			if err != nil {
				log.Fatal(err)
			}

			colour := result[1]
			if num > bag[colour] {
				return false
			}
		}
	}
	return true
}

func get_powers(line string) (power int) {
	line = strings.TrimLeft(line, "Game ")
	tokens := strings.Split(line, ":")
	rounds := strings.Split(tokens[1], ";")
	red, green, blue := get_max_rounds(rounds)
	power = (red * green * blue)
	return
}

func get_max_rounds(rounds []string) (red int, green int, blue int) {
	var current_max = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for i := 0; i < len(rounds); i++ {
		colours := strings.Split(rounds[i], ",")
		for j := 0; j < len(colours); j++ {
			colour_result := strings.TrimSpace(colours[j])
			result := strings.Split(colour_result, " ")
			num, err := strconv.Atoi(result[0])
			if err != nil {
				log.Fatal(err)
			}

			colour := result[1]
			if num > current_max[colour] {
				current_max[colour] = num
			}
		}
	}
	return current_max["red"], current_max["green"], current_max["blue"]
}
