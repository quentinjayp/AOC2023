package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	var sum int

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		d, err := get_digits(fileScanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		sum = sum + d
	}
	fmt.Printf("The sum is %d\n", sum)

	readFile.Close()
}

func get_digits(line string) (found_num int, err error) {
	line = strings.ToLower(line)

	line = strings.ReplaceAll(line, "oneight", "oneeight")
	line = strings.ReplaceAll(line, "twone", "twoone")
	line = strings.ReplaceAll(line, "threeight", "threeeight")
	line = strings.ReplaceAll(line, "fiveight", "fiveeight")
	line = strings.ReplaceAll(line, "sevenine", "sevennine")
	line = strings.ReplaceAll(line, "eightwo", "eighttwo")
	line = strings.ReplaceAll(line, "eighthree", "eightthree")
	line = strings.ReplaceAll(line, "nineight", "nineeight")
	re := regexp.MustCompile("([0-9]|one|two|three|four|five|six|seven|eight|nine)")
	found_digits := re.FindAllString(line, -1)
	first := word_to_digit(found_digits[0])
	last := word_to_digit(found_digits[len(found_digits)-1])
	fmt.Println(first + last)
	found_num, err = strconv.Atoi(strings.Join([]string{first, last}, ""))
	return
}

func word_to_digit(word string) (digit string) {
	m := make(map[string]string)
	m["one"] = "1"
	m["two"] = "2"
	m["three"] = "3"
	m["four"] = "4"
	m["five"] = "5"
	m["six"] = "6"
	m["seven"] = "7"
	m["eight"] = "8"
	m["nine"] = "9"

	digit, ok := m[word]
	if ok {
		return digit
	}
	digit = word
	return digit
}
