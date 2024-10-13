package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var projectName, projectType string

	// Memasukkan nama project dari user
	fmt.Print("Enter project name: ")
	fmt.Scanln(&projectName)

	// Validasi nama project tidak boleh ada spasi
	if strings.Contains(projectName, " ") {
		fmt.Println("Project name should not contain spaces.")
		return
	}

	// Memilih tipe project: simple atau api
	fmt.Print("Choose project type (simple, api): ")
	fmt.Scanln(&projectType)

	// Jika tipe project tidak valid
	if projectType != "simple" && projectType != "api" {
		fmt.Println("creating", projectName, "project...")
		fmt.Println("invalid project type")
		return
	}

	// Membuat folder project
	fmt.Println("creating", projectName, "project...")
	err := os.Mkdir(projectName, 0755)
	if err != nil {
		fmt.Println("Error creating project folder:", err)
		return
	}

	// Inisialisasi Go module
	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Dir = projectName
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error initializing Go module:", err)
		return
	}

	// Menangani pembuatan project berdasarkan tipe yang dipilih
	switch projectType {
	case "simple":
		// Membuat file main.go untuk project simple
		mainGoContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello, this is a simple project!")
}`
		err = os.WriteFile(fmt.Sprintf("%s/main.go", projectName), []byte(mainGoContent), 0644)
		if err != nil {
			fmt.Println("Error creating main.go:", err)
			return
		}

	case "api":
		// Membuat struktur folder untuk project api
		dirs := []string{"routes", "models", "repositories", "services", "configs", "controllers"}
		for _, dir := range dirs {
			err = os.Mkdir(fmt.Sprintf("%s/%s", projectName, dir), 0755)
			if err != nil {
				fmt.Println("Error creating directory:", err)
				return
			}
		}

		// Membuat file main.go untuk project api
		mainGoContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello, this is an API project!")
}`
		err = os.WriteFile(fmt.Sprintf("%s/main.go", projectName), []byte(mainGoContent), 0644)
		if err != nil {
			fmt.Println("Error creating main.go:", err)
			return
		}
	}

	fmt.Printf("Project %s created successfully!\n", projectName)
}
