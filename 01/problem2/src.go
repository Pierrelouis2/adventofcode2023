package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("01/problem2/inputs/inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var convert_dict = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	scanner := bufio.NewScanner(file)

	// scanner.Split(bufio.ScanRunes)
	var i = 0
	var count int
	for scanner.Scan() {
		var lst_num = []string{}
		var line_number string
		for num := range scanner.Text() {
			for k, v := range convert_dict {
				if num+len(k) <= len(scanner.Text()) {
					if string(scanner.Text()[num:num+len(k)]) == k {
						fmt.Println("V=", v)
						lst_num = append(lst_num, v)
					}
				}
			}
			if unicode.IsDigit(rune(scanner.Text()[num])) {
				fmt.Println("Num=", string(scanner.Text()[num]))
				lst_num = append(lst_num, string(scanner.Text()[num]))
			}
		}

		if len(lst_num) == 1 {
			line_number = lst_num[0] + lst_num[0]
		}
		if len(lst_num) > 1 {
			line_number = lst_num[0] + lst_num[len(lst_num)-1]
		}
		fmt.Println("Line number =", line_number)
		if line_number != "" {
			// string to int
			num, err := strconv.Atoi(line_number)
			if err != nil {
				// ... handle error
				panic(err)
			}
			count += num
		}

		fmt.Println(lst_num)
		i++
		fmt.Println("I =", i)
	}
	fmt.Println("Count =", count)
}
