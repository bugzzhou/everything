package sudoku

import (
	"testing"
)

// func TestTrans(t *testing.T) {
// 	matrix := [][]int{
// 		{5, 3, 4, 6, 7, 8, 9, 1, 2},
// 		{6, 7, 2, 1, 9, 5, 3, 4, 8},
// 		{1, 9, 8, 3, 4, 2, 5, 6, 7},
// 		{8, 5, 9, 7, 6, 1, 4, 2, 3},
// 		{4, 2, 6, 8, 5, 3, 7, 9, 1},
// 		{7, 1, 3, 9, 2, 4, 8, 5, 6},
// 		{9, 6, 1, 5, 3, 7, 2, 8, 4},
// 		{2, 8, 7, 4, 1, 9, 6, 3, 5},
// 		{3, 4, 5, 2, 8, 6, 1, 7, 9},
// 	}
// 	fmt.Printf("before: %v\n", IsValidSudoku(matrix))
// 	n, m := 1, 1
// 	matrix = Trans(matrix, n, m)
// 	fmt.Printf("after: %v\n", IsValidSudoku(matrix))
// 	fmt.Println("变换后的矩阵:")
// 	for _, row := range matrix {
// 		fmt.Println(row)
// 	}
// }

func TestTrans(t *testing.T) {
	var sudo = Sudo{}
	sudo.Gen()
	sudo.Fill()
	sudo.Display()
}
