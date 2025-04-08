package main

import "fmt"

func perkalianIteratif(n, m int) int {
	hasil := 0
	for i := 0; i < m; i++ {
		hasil += n
	}
	return hasil
}

func perkalianRekursif(n, m int) int {
	return perkalianRekursifHelper(n, m, 0)
}

func perkalianRekursifHelper(n, m, akumulator int) int {

	if m == 0 {
		return akumulator
	}

	return perkalianRekursifHelper(n, m-1, akumulator+n)
}

func main() {
	var n, m int

	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan bilangan m: ")
	fmt.Scan(&m)

	hasil := perkalianRekursif(n, m)

	fmt.Printf("Hasil dari %d x %d = %d\n", n, m, hasil)
}
