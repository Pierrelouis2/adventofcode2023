package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"unicode"
)

func main() {
	file, err := os.Open("03/problem1/inputs/inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// scanner.Split(bufio.ScanRunes)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fmt.Println(lines)
	type span struct {
		x [2]int // Droite et gauche
		y [2]int // Haut et bas
	}
	var cur_span span
	SYMBOLS := [10]string{"#", "$", "%", "&", "*", "-", "/", "=", "@", "+"}

	for line_number, line := range lines {
		// fmt.Println(line)
		for char_number, char := range line {
			if unicode.IsNumber(rune(char)) { // if number then define span
				if line_number == 0 {
					if char_number == 0 {
						cur_span := span{x: [2]int{0, 1}, y: [2]int{0, 1}}
					}
					if char_number == len(line) {
						cur_span := span{x: [2]int{-1, 0}, y: [2]int{0, 1}}
					} else {
						cur_span := span{x: [2]int{-1, 1}, y: [2]int{0, 1}}
					}
				}
				if line_number == len(lines) {
					if char_number == 0 {
						cur_span := span{x: [2]int{0, 1}, y: [2]int{-1, 0}}
					}
					if char_number == len(line) {
						cur_span := span{x: [2]int{-1, 0}, y: [2]int{-1, 0}}
					} else {
						cur_span := span{x: [2]int{-1, 1}, y: [2]int{-1, 0}}
					}
				}
				if char_number == 0 {
					cur_span := span{x: [2]int{0, 1}, y: [2]int{-1, 1}}
				}
				if char_number == len(line) {
					cur_span := span{x: [2]int{-1, 0}, y: [2]int{-1, 1}}
				} else {
					cur_span := span{x: [2]int{-1, 1}, y: [2]int{-1, 1}}
				}
			}
			for _, x_range := range cur_span.x {
				if x_range != 0 {
					for _, y_range := range cur_span.y {
						if y_range != 0 {
							if slices.Contains(SYMBOLS, string(lines[line_number+y_range][char_number+x_range])) {
								fmt.Println("Symbol found")
							}
						}
					}
				}

			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

	}
}
