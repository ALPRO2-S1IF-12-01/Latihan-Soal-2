// DWI OKTA SURYANINGRUM / 103112400066
package main

import "fmt"

func main() {
    var a, b int

    // Minta user untuk input dua angka
    fmt.Print("Masukkan nilai a: ")
    fmt.Scan(&a)

    fmt.Print("Masukkan nilai b: ")
    fmt.Scan(&b)

    // Cek apakah input valid, a harus lebih kecil atau sama dengan b
    if a > b {
        fmt.Println("Error: Nilai a harus lebih kecil atau sama dengan b")
        return
    }

    // Cari angka-angka perfect di antara a dan b
    perfectNumbers := findPerfectNumbers(a, b)

    fmt.Print("Perfect numbers antara ", a, " dan ", b, ": ")
    if len(perfectNumbers) == 0 {
        fmt.Println("Tidak ada") // Kalau nggak ketemu, kasih tahu user
    } else {
        // Tampilkan semua perfect number yang ditemukan
        for i, num := range perfectNumbers {
            if i > 0 {
                fmt.Print(", ") // Tambah koma kalau bukan angka pertama
            }
            fmt.Print(num)
        }
        fmt.Println() // Biar barisnya rapi
    }
}

// Fungsi ini buat nyari semua perfect number dalam rentang yang diberikan
func findPerfectNumbers(a, b int) []int {
    var perfectNumbers []int

    // Cek satu per satu angka dari a sampai b
    for num := a; num <= b; num++ {
        if isPerfectNumber(num) {
            perfectNumbers = append(perfectNumbers, num) // Tambahin ke list kalau perfect
        }
    }

    return perfectNumbers
}

// Fungsi ini buat ngecek apakah suatu angka termasuk perfect number
func isPerfectNumber(num int) bool {
    if num <= 1 {
        return false // Angka 1 ke bawah nggak dihitung perfect
    }

    sum := 1 // 1 itu selalu faktor dari angka berapapun
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            sum += i
            otherFactor := num / i
            if otherFactor != i {
                sum += otherFactor // Tambahin pasangan faktornya juga
            }
        }
    }

    return sum == num // Kalau total faktornya sama dengan angkanya, berarti perfect!
}