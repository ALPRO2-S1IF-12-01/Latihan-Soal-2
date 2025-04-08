package main

import "fmt"

func hitungGanjil(a, b int) int {
	count := 0
	for i := a; i <= b; i++ {
		if i%2 != 0 {
			count++
		}
	}
	return count
}
func main() {
	fmt.Println("Nama: Raja Muhammad Lufhti\nNim : 103112400027")
	var a, b int
	fmt.Print("masukkan nilai a :")
	fmt.Scan(&a)
	fmt.Print("masukkan nilai b :")
	fmt.Scan(&b)

	fmt.Printf("banyak angka ganjil :%d", hitungGanjil(a, b))
}