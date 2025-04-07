// RYAN AKEYLA NOVIANTO WIDODO
// 103112400081

package main

import "fmt"

func perkalianRekursif(n, m, hasil int) int {
	if m == 0 {
		return hasil
	}
	return perkalianRekursif(n, m-1, hasil+n)
}

func main() {
	var n, m_103112400081 int
	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan bilangan m: ")
	fmt.Scan(&m_103112400081)

	hasil := perkalianRekursif(n, m_103112400081, 0)
	fmt.Printf("Hasil dari %d x % d = %d\n", n, m_103112400081, hasil)
}
