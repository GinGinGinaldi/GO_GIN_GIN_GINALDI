package main

import (
	"fmt"
	"sort"
)

func main() {
	dataSatu := []float64{1, 1, 2, 5, 6, 8, 12, 4, 5, 5, 5, 8, 9}

	fmt.Printf("Jumlah: %.2f\n", jumlah(dataSatu))       // 71.00
	fmt.Printf("Rata-rata: %.2f\n", rataRata(dataSatu))  // 5.46
	fmt.Printf("Median: %.2f\n", hitungMedian(dataSatu)) // 5.00
	fmt.Printf("Modus: %.2f\n", hitungModus(dataSatu))   // 5.00

	dataDua := []float64{6, 7, 1, 11, 8, 12, 6, 12, 1, 7, 8, 10, 14}

	fmt.Printf("Jumlah: %.2f\n", jumlah(dataDua))       // 103.00
	fmt.Printf("Rata-rata: %.2f\n", rataRata(dataDua))  // 7.92
	fmt.Printf("Median: %.2f\n", hitungMedian(dataDua)) // 8.00
	fmt.Printf("Modus: %.2f\n", hitungModus(dataDua))   // 1.00
}

func jumlah(data []float64) float64 {
	total := 0.0
	for _, angka := range data {
		total += angka
	}
	return total
}

func rataRata(data []float64) float64 {
	return jumlah(data) / float64(len(data))
}

func hitungMedian(data []float64) float64 {
	sort.Float64s(data)
	panjang := len(data)
	if panjang%2 == 1 {
		return data[panjang/2]
	}
	return (data[panjang/2-1] + data[panjang/2]) / 2
}

func hitungModus(data []float64) float64 {
	frekuensi := make(map[float64]int)
	frekuensiTertinggi := 0
	modus := data[0]

	for _, angka := range data {
		frekuensi[angka]++
		if frekuensi[angka] > frekuensiTertinggi {
			frekuensiTertinggi = frekuensi[angka]
			modus = angka
		}
	}
	return modus
}
