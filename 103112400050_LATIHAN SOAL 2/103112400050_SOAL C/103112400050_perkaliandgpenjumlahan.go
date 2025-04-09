package main

import "fmt"

//103112400050
func main() {
	var n, m int
	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan bilangan m: ")
	fmt.Scan(&m)

	hasil := perkalianRekursif(n, m, 0)

	fmt.Printf("Hasil dari %d x %d = %d\n", n, m, hasil)
}

func perkalianRekursif(n, m, total int) int {
	if n == 0 || m == 0 {
		return 0
	}
	if n < 0 {
		return -perkalianRekursif(-n, m, 0)
	}
	if n == 1 {
		return total + m
	}
	return perkalianRekursif(n-1, m, total+m)
}

func perkalianIteratif(n, m int) int {
	if n == 0 || m == 0 {
		return 0
	}
	negatif := false
	if n < 0 {
		n = -n
		negatif = true
	}
	total := 0
	for i := 0; i < n; i++ {
		total += m
	}
	if negatif {
		return -total
	}
	return total
}
