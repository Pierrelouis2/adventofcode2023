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
	file, err := os.Open("04/problem2/inputs/inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// create a register of cards and populate with 1
	var card_register = make([]int, len(lines))
	for i := range card_register {
		card_register[i] = 1
	}

	total_count := 0
	for id, line := range lines {
		// create a map of card number and count
		s := strings.Split(line, ":")[1]
		win_number := strings.Split(s, "|")[0]
		arr_win_number := convert_to_int(strings.Split(win_number, " "))
		my_number := strings.Split(s, "|")[1]
		arr_my_number := convert_to_int(strings.Split(my_number, " "))

		// fmt.Println(arr_win_number)
		// fmt.Println(arr_my_number)
		card_score := 0

		// loop to sum up the card count
		for _, num := range arr_my_number {
			if slices.Contains(arr_win_number, num) {
				card_score++
			}
		}

		// second loop to distribute score over the other cards
		if card_register[id] == 0 {
			for i := 1; i <= card_score; i++ {
				card_register[id+i]++
			}
		} else {
			for j := 1; j <= card_register[id]; j++ {
				for i := 1; i <= card_score; i++ {
					card_register[id+i]++
				}
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		// fmt.Println("card_score", card_score, "card number", id, card_register)
	}
	for _, v := range card_register {
		total_count += v
	}
	fmt.Println("total_count", total_count)

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
