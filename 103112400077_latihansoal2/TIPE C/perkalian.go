package main

import "fmt"

func dikali(n, m, hasil int) int {
    if m == 0 {
        return hasil
    }
    return dikali(n, m-1, hasil+n)
}

func main() {
    var n_103112400077, m int

    fmt.Print("Masukkan bilangan n: ")
    fmt.Scan(&n_103112400077)
    fmt.Print("Masukkan bilangan m: ")
    fmt.Scan(&m)

    if m < 0 {
        m = -m
        n_103112400077 = -n_103112400077
    }

    hasil := dikali(n_103112400077, m, 0)

    fmt.Printf("Hasil dari %d x %d = %d", n_103112400077, m, hasil)
}
