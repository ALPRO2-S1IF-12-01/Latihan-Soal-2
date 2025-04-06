package main

import "fmt"

func perkalian(n, m, hasil int) int {
    if m == 0 {
        return hasil
    }
    return perkalian(n, m-1, hasil+n)
}

func main() {
    var n_103112400073, m int

    fmt.Print("Masukkan bilangan n: ")
    fmt.Scan(&n_103112400073)
    fmt.Print("Masukkan bilangan m: ")
    fmt.Scan(&m)

    negatif := false
    if m < 0 {
        m = -m
        negatif = true
    }

    hasil := perkalian(n_103112400073, m, 0)

    if negatif {
        hasil = -hasil
    }

    fmt.Printf("Hasil dari %d x %d = %d\n", n_103112400073, m, hasil)
}