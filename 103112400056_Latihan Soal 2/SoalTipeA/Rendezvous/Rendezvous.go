// Felix Pedrosa V

package main

import "fmt"

func main() {
    var x_103112400056, y int

    // Input values for x and y
    fmt.Print("Masukkan nilai x: ")
    fmt.Scan(&x_103112400056)
    fmt.Print("Masukkan nilai y: ")
    fmt.Scan(&y)

    // Counting meetings days
    jumlahPertemuan := 0
    // Iterate from 1 to 365 to find meeting days
    for hari := 1; hari <= 365; hari++ {
        if hari%x_103112400056 == 0 && hari%y != 0 {
            jumlahPertemuan++
        }
    }

    // Display result
    fmt.Printf("Jumlah pertemuan dalam setahun: %d\n", jumlahPertemuan)
}