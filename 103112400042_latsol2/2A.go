// M.HANIF AL FAIZ
// 103112400042
package main

import "fmt"

func main() {
	var start, end int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&start)

	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&end)

	if start > end {
		fmt.Println("Error: a harus lebih kecil atau sama dengan b")
		return
	}

	if start < 0 || end < 0 {
		fmt.Println("Error: Masukkan bilangan bulat positif")
		return
	}

	oddCount := countOddNumbers(start, end)
	fmt.Printf("Banyaknya angka ganjil: %d\n", oddCount)
}
func countOddNumbers(start, end int) int {

	if start%2 == 0 {
		start++
	}
	if end%2 == 0 {
		end--
	}
	if start > end {
		return 0
	}
	return (end-start)/2 + 1
}
