package main

import "fmt"

func main() {
	var bilangan int
	fmt.Print("masukkan bilangan bulat positif (>10) : ")
	fmt.Scanln(&bilangan)

	if bilangan <= 10 {
		fmt.Println("bilangan harus lebih dari 10.")
		return
	}

	potongBilangan(bilangan)
}

func potongBilangan(bilangan2 int) {
	var jumlahDigit, kiri, kanan, hasil, pemisah int
	salinanBilangan := bilangan2

	for salinanBilangan > 0 {
		jumlahDigit++
		salinanBilangan = salinanBilangan / 10
	}

	jumlahKiri := (jumlahDigit + 1) / 2

	pemisah = 1
	for i := 0; i < jumlahDigit-jumlahKiri; i++ {
		pemisah = pemisah * 10
	}

	kiri = bilangan2 / pemisah
	kanan = bilangan2 % pemisah
	hasil = kiri + kanan

	fmt.Println(kiri, kanan)
	fmt.Println(hasil)
	fmt.Print("Sheila Stephanie A [103112400086]")
}
