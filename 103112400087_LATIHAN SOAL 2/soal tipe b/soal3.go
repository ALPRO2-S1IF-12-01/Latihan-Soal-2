package main

import "fmt"

func main() {
	fmt.Println("Masukkan bilangan (negatif untuk berhenti):")

	var nums []int
	for {
		var num int
		fmt.Scan(&num)
		if num < 0 {
			break
		}
		nums = append(nums, num)
	}

	total := jumlahKelipatan4(nums, 0, 0)
	fmt.Printf("Jumlah bilangan kelipatan 4: %d\n", total)
}

func jumlahKelipatan4(data []int, index_103112400087 int, total int) int {
	if index_103112400087 >= len(data) {
		return total
	}
	if data[index_103112400087]%4 == 0 {
		total += data[index_103112400087]
	}
	return jumlahKelipatan4(data, index_103112400087+1, total)
}
