package main

import "fmt"

func multiply(n int, m_103112400087 int) int {
	if m_103112400087 == 0 {
		return 0
	}
	return n + multiply(n, m_103112400087-1)
}

func main() {
	var n, m_103112400087 int
	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan bilangan m: ")
	fmt.Scan(&m_103112400087)

	result := multiply(n, m_103112400087)
	fmt.Printf("Hasil dari %d x %d = %d\n", n, m_103112400087, result)
}
