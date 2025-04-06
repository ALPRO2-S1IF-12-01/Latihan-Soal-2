//Anastasia Adinda N.I
package main

import (
	"fmt"
)

func perkalianRekursif(n_103112400085, m, hasil int) int {
	if n_103112400085 == 0 {
		return hasil
	}
	return perkalianRekursif(n_103112400085-1, m, hasil+m)
}

func main() {
	var n_103112400085, m int

	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n_103112400085)

	fmt.Print("Masukkan bilangan m: ")
	fmt.Scan(&m)

	hasil := perkalianRekursif(n_103112400085, m, 0)

	fmt.Printf("Hasil dari %d x %d = %d\n", n_103112400085, m, hasil)
}
