// MUHAMMAD GAMEL AL GHIFARI
// 103112400028
package main

import "fmt"

func main() {
	var intervalA, intervalB int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&intervalA)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&intervalB)

	totalMeetings := calculateSecretMeetings(intervalA, intervalB)

	fmt.Printf("Jumlah pertemuan rahasia dalam setahun: %d\n", totalMeetings)
}

func calculateSecretMeetings(a, b int) int {
	meetingCount := 0
	for day := 1; day <= 365; day++ {
		if meetsCriteria(day, a, b) {
			meetingCount++
		}
	}
	return meetingCount
}

func meetsCriteria(day, x, y int) bool {
	return day%x == 0 && day%y != 0
}
