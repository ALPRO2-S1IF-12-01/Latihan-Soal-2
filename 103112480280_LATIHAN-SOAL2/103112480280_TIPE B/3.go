// Nama: Anggun Wahyu Widiyana
// NIM : 103112480280
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func jumlahKelipatan4Rekursif(numbers_103112480280 []int, idx int, acc int) int {
	if idx == len(numbers_103112480280) {
		return acc
	}
	if numbers_103112480280[idx] > 0 && numbers_103112480280[idx]%4 == 0 {
		acc += numbers_103112480280[idx]
	}
	return jumlahKelipatan4Rekursif(numbers_103112480280, idx+1, acc)
}

func main() {
	fmt.Println("Masukkan bilangan (negatif untuk berhenti):")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	fields := strings.FieldsFunc(input, func(r rune) bool {
		return r < '0' || r > '9' && r != '-'
	})

	var numbers_103112480280 []int
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err == nil {
			if num < 0 {
				break
			}
			numbers_103112480280 = append(numbers_103112480280, num)
		}
	}

	result := jumlahKelipatan4Rekursif(numbers_103112480280, 0, 0)
	fmt.Printf("Jumlah bilangan kelipatan 4: %d\n", result)
}