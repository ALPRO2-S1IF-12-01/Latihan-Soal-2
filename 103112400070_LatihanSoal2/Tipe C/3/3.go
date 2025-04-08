// Achmad Zulvan Nur Hakim 103112400070
package main

import "fmt"


func kali(n, m int) int {
	hasil := 0
	for i := 0; i < m; i++ {
		hasil += n
	}
	return hasil
}

func main() {
	var n, m int
	fmt.Print("Bilangan n: ")
	fmt.Scan(&n)
	fmt.Print("Bilangan m: ")
	fmt.Scan(&m)

	hasil := kali(n, m)
	fmt.Printf("Hasil dari %d x %d = %d\n", n, m, hasil)
}
