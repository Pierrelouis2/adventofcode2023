package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("04/problem1/inputs/inputs.txt")
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
	total_count := 0
	for id, line := range lines {
		s := strings.Split(line, ":")[1]
		win_number := strings.Split(s, "|")[0]
		arr_win_number := convert_to_int(strings.Split(win_number, " "))
		my_number := strings.Split(s, "|")[1]
		arr_my_number := convert_to_int(strings.Split(my_number, " "))
		fmt.Println(arr_win_number)
		fmt.Println(arr_my_number)
		card_count := 0
		for _, num := range arr_my_number {
			if slices.Contains(arr_win_number, num) {
				if card_count == 0 {
					card_count++
				} else {
					card_count *= 2
				}
				fmt.Println(num)
			}
		}
		total_count += card_count
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("card_count", card_count, "card number", id)
	}
	fmt.Println(total_count)
}

func convert_to_int(arr []string) []int {
	intSlice := make([]int, 0, len(arr))

	for _, s := range arr {
		num, err := strconv.Atoi(s)
		if err != nil {
			if err.Error() == `strconv.Atoi: parsing "": invalid syntax` {
				continue
			}
			fmt.Println(err)
			continue
		}
		intSlice = append(intSlice, num)
	}
	return intSlice
}
