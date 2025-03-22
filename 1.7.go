package main

import "fmt"

//func main() {
//	matrix := [][]int{
//		{1, 2, 3, 4},
//		{5, 6, 7, 8},
//		{9, 10, 11, 12},
//		{13, 14, 15, 16},
//	}
//	res := rotateMatrix(matrix)
//	printMatrix(res)
////}

func rotateMatrix(matrix [][]int) [][]int {
	rows := len(matrix)
	if rows == 0 || rows != len(matrix[0]) {
		fmt.Println("Matrix must be square for in-place rotation.")
		return matrix
	}

	for layer := 0; layer < rows/2; layer++ {
		first := layer
		last := rows - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			top := matrix[first][i]

			// Left -> Top
			matrix[first][i] = matrix[last-offset][first]

			// Bottom -> Left
			matrix[last-offset][first] = matrix[last][last-offset]

			// Right -> Bottom
			matrix[last][last-offset] = matrix[i][last]

			// Top -> Right
			matrix[i][last] = top
		}
	}
	return matrix
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}
