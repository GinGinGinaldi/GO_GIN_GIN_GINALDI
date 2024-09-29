package main

import "fmt"

// Struktur untuk menyimpan data siswa
type Siswa struct {
	nama       string
	nilaiMat   int
	nilaiSains int
	nilaiIng   int
}

// Fungsi untuk mencari siswa dengan nilai terbaik
func nilaiTerbaik(siswa []Siswa) (Siswa, float64) {
	terbaik := siswa[0]
	totalRata := 0.0
	maxRata := 0.0

	// Mencari nilai terbaik dan rata-rata
	for _, s := range siswa {
		rata := (float64(s.nilaiMat) + float64(s.nilaiSains) + float64(s.nilaiIng)) / 3
		totalRata += rata
		if rata > maxRata {
			maxRata = rata
			terbaik = s
		}
	}

	return terbaik, totalRata / float64(len(siswa)) // Mengembalikan siswa terbaik dan rata-rata
}

func main() {
	siswa := []Siswa{
		{"Ali", 85, 90, 80},
		{"Budi", 90, 85, 88},
		{"Cici", 75, 70, 80},
	}

	terbaik, rata := nilaiTerbaik(siswa)
	fmt.Printf("Siswa dengan nilai terbaik: %s\n", terbaik.nama)
	fmt.Printf("Nilai rata-rata seluruh siswa: %.2f\n", rata)
}
