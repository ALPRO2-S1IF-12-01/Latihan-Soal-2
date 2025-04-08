//Pratama Bintang Daniswara 103112400051
package main

import "fmt"

func main() {
	var n, m int
	fmt.Print("Masukkan bilangan n: ")
	fmt.Scanln(&n)
	fmt.Print("Masukkan bilangan m: ")
	fmt.Scanln(&m)

	hasil := kali(n, m)
	fmt.Println("Hasil dari", n, "x", m, "=", hasil)
}

func kali(n, m int) int {
	if m == 0 {
		return 0
	}
	return n + kali(n, m-1)
}
