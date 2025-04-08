package main

import "fmt"

func sumKelipatan4(arr []int, index int, total int) int {

	if index == len(arr) {
		return total
	}

	if arr[index] < 0 {
		return total
	}

	if arr[index]%4 == 0 {
		total += arr[index]
	}

	return sumKelipatan4(arr, index+1, total)
}

func main() {
	fmt.Println("Nama: Raja Muhammad Lufhti\nNim : 103112400027")
	var bilangan []int
	var input int

	fmt.Println("Masukkan bilangan (negatif untuk berhenti):")
	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		bilangan = append(bilangan, input)
	}

	hasil := sumKelipatan4(bilangan, 0, 0)

	fmt.Printf("Jumlah bilangan kelipatan 4: %d\n", hasil)
}
