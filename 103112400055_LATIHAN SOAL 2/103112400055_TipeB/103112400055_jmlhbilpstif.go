package main

import "fmt"

// Fungsi rekursif untuk menjumlahkan bilangan kelipatan 4
func jumlahKelipatanEmpat(numbers []int, index int) int {
	if index == len(numbers) {
		return 0
	}

	if numbers[index] > 0 && numbers[index]%4 == 0 {
		return numbers[index] + jumlahKelipatanEmpat(numbers, index+1)
	}

	return jumlahKelipatanEmpat(numbers, index+1)
}

func main() {
	var numbers []int
	var input_103112400055 int

	fmt.Println("Masukkan bilangan (negatif untuk berhenti): ")

	for {
		fmt.Scan(&input_103112400055)
		if input_103112400055 < 0 {
			break
		}
		numbers = append(numbers, input_103112400055)
	}

	total := jumlahKelipatanEmpat(numbers, 0)

	fmt.Printf("Jumlah bilangan kelipatan 4: %d\n", total)
}
