// Nama: Anggun Wahyu Widiyana
// NIM : 103112480280
package main

import "fmt"

func perfect(angka int) bool {
    if angka <= 1 {
        return false
    }
    jumlah := 0
    for i := 1; i <= angka/2; i++ {
        if angka%i == 0 {
            jumlah += i
        }
    }
    return jumlah == angka
}

func main() {
    var awal_103112480280, akhir int
    fmt.Print("Masukkan nilai a: ")
    fmt.Scan(&awal_103112480280)
    fmt.Print("Masukkan nilai b: ")
    fmt.Scan(&akhir)

    fmt.Printf("Perfect numbers antara %d dan %d: ", awal_103112480280, akhir)
    ketemu := false
    for num := awal_103112480280; num <= akhir; num++ {
        if perfect(num) {
            fmt.Printf("%d ", num)
            ketemu = true
        }
    }
    if !ketemu {
        fmt.Print("Tidak ada")
    }
    fmt.Println()
}