package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	columnNumbers := GenerateColumnNumbers()
	selectedLayout := GenerateRandLayouts()
	generate := GenerateTicket(columnNumbers, selectedLayout)
	PrintTicket(generate)
}

func GenerateTicket(columnNumbers [9][]int, selectedLayout [3][5]int) [3][9]int {
	generate := [3][9]int{}
	for i := 0; i < len(selectedLayout); i++ {
		for j := 0; j < len(selectedLayout[i]); j++ {
			q := selectedLayout[i][j]
			r := rand.Intn(len(columnNumbers[q]))
			sayi := columnNumbers[q][r]
			columnNumbers[q] = append(columnNumbers[q][:r], columnNumbers[q][r+1:]...)
			generate[i][q] = sayi
		}
	}
	return generate
}

func GenerateRandLayouts() [3][5]int {
	types := rand.Intn(100)
	if types%2 == 0 {
		return [3][5]int{
			{0, 2, 4, 6, 7},
			{1, 3, 5, 7, 8},
			{0, 2, 4, 6, 8},
		}
	}
	return [3][5]int{
		{0, 1, 3, 5, 7},
		{1, 2, 4, 6, 8},
		{0, 2, 3, 5, 7},
	}
}

func GenerateColumnNumbers() [9][]int {
	numbers := [9][]int{}
	for i := 0; i < 9; i++ {
		sayi := 10 * i
		for j := 0; j < 10; j++ {
			if i == 0 && j == 0 {
				continue
			}
			numbers[i] = append(numbers[i], sayi+j)
		}
		if i == 8 {
			numbers[i] = append(numbers[i], 90)
		}
	}
	return numbers
}

func PrintTicket(data [3][9]int) {
	for i := 0; i < len(data); i++ {
		jlen := len(data[i])
		for j := 0; j < jlen; j++ {
			q := data[i][j]
			if q == 0 {
				fmt.Print("[  ] ")
			} else {
				fmt.Printf("[%2.d] ", q)
			}
		}
		fmt.Println()
	}
}
