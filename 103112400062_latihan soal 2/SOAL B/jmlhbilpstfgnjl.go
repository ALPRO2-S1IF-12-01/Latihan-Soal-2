// M. DAVI ILYAS RENALDO/103112400062 
package main

import (
	"fmt"
)

func jumlahKelipatan4(arr []int, index int) int {
	if index >= len(arr) {
		return 0
	}
	if arr[index] > 0 && arr[index]%4 == 0 {
		return arr[index] + jumlahKelipatan4(arr, index+1)
	}
	return jumlahKelipatan4(arr, index+1)
}

func main() {
	var input int
	var data []int

	fmt.Println("Masukkan bilangan (negatif untuk berhenti):")
	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		data = append(data, input)
	}

	hasil := jumlahKelipatan4(data, 0)
	fmt.Println("Jumlah bilangan kelipatan 4:", hasil)
}
