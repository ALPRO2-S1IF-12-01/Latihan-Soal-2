package main

import "fmt"

func kali(n, m int) int {
	if m == 0 {
		return 0
	}
	return n + kali(n, m-1)
}

func main() {
	var n, m int
	fmt.Println("NAMA: MULIA AKBAR NANDA PRATAMA\nNIM: 103112400034")
	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan bilangan m: ")
	fmt.Scan(&m)

	hasil := kali(n, m)
	fmt.Printf("Hasil dari %d x %d = %d\n", n, m, hasil)
}
