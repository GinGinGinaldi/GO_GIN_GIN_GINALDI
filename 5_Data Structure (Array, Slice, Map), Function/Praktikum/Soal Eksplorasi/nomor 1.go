package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Course struct {
	Name                       string
	MentorSatisfactionScores   []float64
	LearningSatisfactionScores []float64
}

func main() {
	// Membuka file CSV
	file, err := os.Open("C:/Users/ADITT/Documents/GO_GIN GIN GINALDI/5_Data Structure (Array, Slice, Map), Function/Praktikum/Soal Eksplorasi/survey.csv")
	if err != nil {
		fmt.Println("Error membuka file:", err)
		return
	}
	defer file.Close()

	// Membaca file CSV
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error membaca file CSV:", err)
		return
	}

	// Map untuk menyimpan data kursus berdasarkan nama kursus
	courses := make(map[string]*Course)

	// Mengambil header
	for i, record := range records {
		// Lewati baris header
		if i == 0 {
			continue
		}

		courseName := record[0]
		mentorSatisfaction, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			fmt.Println("Error parsing mentor satisfaction score:", err)
			continue
		}

		learningSatisfaction, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			fmt.Println("Error parsing learning satisfaction score:", err)
			continue
		}

		// Jika kursus belum ada dalam map, inisialisasi kursus baru
		if _, exists := courses[courseName]; !exists {
			courses[courseName] = &Course{
				Name:                       courseName,
				MentorSatisfactionScores:   []float64{},
				LearningSatisfactionScores: []float64{},
			}
		}

		// Tambahkan skor ke kursus
		courses[courseName].MentorSatisfactionScores = append(courses[courseName].MentorSatisfactionScores, mentorSatisfaction)
		courses[courseName].LearningSatisfactionScores = append(courses[courseName].LearningSatisfactionScores, learningSatisfaction)
	}

	// Menghitung rata-rata kepuasan mentor dan kepuasan belajar secara keseluruhan
	var totalMentorSatisfaction, totalLearningSatisfaction float64
	var totalMentorCount, totalLearningCount int
	var highestMentorSatisfaction float64
	var highestMentorSatisfactionCourse string

	for _, course := range courses {
		var mentorSum float64

		// Hitung total untuk setiap kursus
		for _, score := range course.MentorSatisfactionScores {
			mentorSum += score
		}
		for _, score := range course.LearningSatisfactionScores {
			totalLearningSatisfaction += score
			totalLearningCount++
		}

		// Hitung rata-rata untuk setiap kursus
		mentorAverage := mentorSum / float64(len(course.MentorSatisfactionScores))

		// Tambahkan ke total keseluruhan
		totalMentorSatisfaction += mentorSum
		totalMentorCount += len(course.MentorSatisfactionScores)

		// Cek jika kursus ini memiliki rata-rata kepuasan mentor tertinggi
		if mentorAverage > highestMentorSatisfaction {
			highestMentorSatisfaction = mentorAverage
			highestMentorSatisfactionCourse = course.Name
		}
	}

	// Hitung rata-rata keseluruhan
	overallMentorAverage := totalMentorSatisfaction / float64(totalMentorCount)
	overallLearningAverage := totalLearningSatisfaction / float64(totalLearningCount)

	// Cetak hasil
	fmt.Printf("Rata-rata kepuasan mentor secara keseluruhan: %.2f\n", overallMentorAverage)
	fmt.Printf("Rata-rata kepuasan belajar secara keseluruhan: %.2f\n", overallLearningAverage)
	fmt.Printf("Kursus dengan rata-rata kepuasan mentor tertinggi: %s (%.2f)\n", highestMentorSatisfactionCourse, highestMentorSatisfaction)
}
