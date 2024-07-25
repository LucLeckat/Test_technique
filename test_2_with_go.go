package main

import (
	"fmt"
)

func generateDiagonalTable(width int, height int) {
	// Create an empty table
	table := make([][]int, height)
	for i := range table {
		table[i] = make([]int, width)
	}

	value := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == j {
				table[i][j] = value
				value++
			}
		}
	}

	// Print the table in a nice format
	for _, row := range table {
		for _, cell := range row {
			fmt.Printf("%d ", cell)
		}
		fmt.Println()
	}
}

func main() {
	var width, height int
	fmt.Print("Entrer la largeur de la table: ")
	fmt.Scan(&width)
	fmt.Print("Entrer la hauteur de la table: ")
	fmt.Scan(&height)
	generateDiagonalTable(width, height)
}
