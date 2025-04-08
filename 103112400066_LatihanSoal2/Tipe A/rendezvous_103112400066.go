// DWI OKTA SURYANINGRUM / 103112400066

package main

import "fmt"

func main() {
    var x, y int

    // Meminta input dari user
    fmt.Print("Masukkan nilai x: ")
    fmt.Scan(&x)
    
    fmt.Print("Masukkan nilai y: ")
    fmt.Scan(&y)

    // Validasi input
    if x <= 0 || y <= 0 {
        fmt.Println("Error: Nilai x dan y harus bilangan bulat positif")
        return
    }

    // Menghitung jumlah pertemuan
    jumlahPertemuan := hitungPertemuan(x, y)
    
    fmt.Println("Jumlah pertemuan dalam setahun:", jumlahPertemuan)
}

// Fungsi untuk menghitung jumlah pertemuan
func hitungPertemuan(x, y int) int {
    count := 0
    for hari := 1; hari <= 365; hari++ {
        // Cek apakah hari kelipatan x tetapi bukan kelipatan y
        if hari%x == 0 && hari%y != 0 {
            count++
        }
    }
    return count
}