package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	lines := []string{"two1nine", "eightwo", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}

	for i := 0; i < len(lines); i++ {
		//fmt.Println(lines[i])
		line := lines[i]
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
		fmt.Println(found_digits)
	}
}
