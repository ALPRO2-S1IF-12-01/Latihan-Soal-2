// Felix Pedrosa V 

package main

import "fmt"

// Fungsi rekursif dengan tail-end recursion untuk menghitung perkalian
func perkalianRekursif(n, m, hasil int) int {
	if m == 0 {
		return hasil
	}
	return perkalianRekursif(n, m-1, hasil+n)
}

func main() {
	var n, m_103112400056 int
	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan bilangan m: ")
	fmt.Scan(&m_103112400056)

	hasil := perkalianRekursif(n, m_103112400056, 0)
	fmt.Printf("Hasil dari %d x %d = %d\n", n, m_103112400056, hasil)
}
