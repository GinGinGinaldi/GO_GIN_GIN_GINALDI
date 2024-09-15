package main

import (
	"fmt"
)

func main() {
	var budget, waktuPengerjaan, tingkatKesulitan, skorBudget, skorWaktu, skorKesulitan, totalSkor int
	var prioritas string

	// minta input buat budget, waktu pengerjaan, sama tingkat kesulitan
	fmt.Print("Masukkan budget proyek: ")
	fmt.Scan(&budget)
	fmt.Print("Masukkan waktu pengerjaan (hari): ")
	fmt.Scan(&waktuPengerjaan)
	fmt.Print("Masukkan tingkat kesulitan (0-10): ")
	fmt.Scan(&tingkatKesulitan)

	// nentuin skor buat budget
	if budget >= 0 && budget <= 50 {
		skorBudget = 4
	} else if budget >= 51 && budget <= 80 {
		skorBudget = 3
	} else if budget >= 81 && budget <= 100 {
		skorBudget = 2
	} else {
		skorBudget = 1
	}

	// hitungin skor buat waktu pengerjaan
	if waktuPengerjaan >= 0 && waktuPengerjaan <= 20 {
		skorWaktu = 10
	} else if waktuPengerjaan >= 21 && waktuPengerjaan <= 30 {
		skorWaktu = 5
	} else if waktuPengerjaan >= 31 && waktuPengerjaan <= 50 {
		skorWaktu = 2
	} else {
		skorWaktu = 1
	}

	// hitungin skor buat tingkat kesulitan
	if tingkatKesulitan >= 0 && tingkatKesulitan <= 3 {
		skorKesulitan = 10
	} else if tingkatKesulitan >= 4 && tingkatKesulitan <= 6 {
		skorKesulitan = 5
	} else if tingkatKesulitan >= 7 && tingkatKesulitan <= 10 {
		skorKesulitan = 1
	} else {
		skorKesulitan = 0
	}

	// hitung total skor
	totalSkor = skorBudget + skorWaktu + skorKesulitan

	// tentuin prioritas berdasarkan total skor
	if totalSkor >= 17 && totalSkor <= 24 {
		prioritas = "High"
	} else if totalSkor >= 10 && totalSkor <= 16 {
		prioritas = "Medium"
	} else if totalSkor >= 3 && totalSkor <= 9 {
		prioritas = "Low"
	} else {
		prioritas = "Impossible"
	}

	// tampilin hasil prioritas proyek
	fmt.Println("Prioritas proyek: ", prioritas)
}
