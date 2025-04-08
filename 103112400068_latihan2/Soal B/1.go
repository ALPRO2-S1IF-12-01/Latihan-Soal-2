package main

import "fmt"

func main() {
	var r, n int
	fmt.Print("masukkan nilai A:")
	fmt.Scanln(&r)
	fmt.Print("masukkan nilai B:")
	fmt.Scanln(&n)
	if r > n {
		r, n = n, r
	}
	hasil := 0
	for i := r; i <= n; i++ {
		if i%2 != 0 {
			hasil++
		}
	}
	fmt.Println("Banyaknya angka ganjil", hasil)
}
