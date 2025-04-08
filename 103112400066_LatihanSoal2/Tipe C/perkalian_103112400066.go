// DWI OKTA SURYANINGRUM / 103112400066

package main

import "fmt"

// Fungsi rekursif untuk perkalian dengan penjumlahan
func perkalianRekursif(n, m int) int {
    // Basis: jika n = 0, hasilnya 0
    if n == 0 {
        return 0
    }
    // Rekursi: tambahkan m dengan hasil perkalian (n-1) x m
    return m + perkalianRekursif(n-1, m)
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