// Nama: Anggun Wahyu Widiyana
// NIM : 103112480280
package main

import "fmt"

func main() {
    var x_103112480280, y int
    fmt.Print("Masukkan nilai x: ")
    fmt.Scan(&x_103112480280)
    fmt.Print("Masukkan nilai y: ")
    fmt.Scan(&y)

    pertemuan := 0
    for hari := 1; hari <= 365; hari++ {
        if hari%x_103112480280 == 0 && hari%y != 0 {
            pertemuan++
        }
    }

    fmt.Printf("Jumlah pertemuan dalam setahun: %d\n", pertemuan)
}