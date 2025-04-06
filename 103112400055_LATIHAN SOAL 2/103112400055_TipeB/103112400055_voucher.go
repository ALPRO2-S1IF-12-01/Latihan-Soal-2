package main

import "fmt"

func main() {
	var a_103112400055, b int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a_103112400055)

	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	count := 0
	for i := a_103112400055; i <= b; i++ {
		if i%2 != 0 {
			count++
		}
	}

	fmt.Println("Banyaknya angka ganjil:", count)
}
