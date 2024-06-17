package sudoku

import (
	"fmt"
	"math/rand"
	"time"
)

type SudokuInterface interface {
	Gen()
	Tips()
	Fill()
	Check() bool
	Display()
}

type Sudo struct {
	Grid         [][]int `json:"grid"`         //使用行列变换获得的原始数独
	ToFillGrid   [][]int `json:"toFillGrid"`   //随机删除部分数字后的数独
	AutoFillGrid [][]int `json:"autoFillGrid"` //自动填充后完整的数独
}

//interface接口类
//-----
//
//
//
//
//
//
//
//
//-----
//具体实现

var _ SudokuInterface = (*Sudo)(nil)

func (S *Sudo) Gen() {
	S.Grid = copyMatrix(Raw)
	Trans(S.Grid, 10, 10)

	S.ToFillGrid = copyMatrix(S.Grid)
	exclude(S.ToFillGrid, 20)
}

// TODO bugzzhou
// 用于提示哪些地方可以填写
func (S *Sudo) Tips() {}

// 用于自动生成正确的数独组合
func (S *Sudo) Fill() {
	S.AutoFillGrid = copyMatrix(S.ToFillGrid)
	solveSudoku(S.AutoFillGrid)
}

func (S *Sudo) Check() bool {
	return IsValidSudoku(S.Grid)
}

func (S *Sudo) Display() {
	fmt.Printf("grid is: \n")
	for _, v := range S.Grid {
		for _, vv := range v {
			fmt.Printf("%d ", vv)
		}
		fmt.Printf("\n")
	}

	fmt.Printf("toFillGrid is: \n")
	for _, v := range S.ToFillGrid {
		for _, vv := range v {
			if vv == 0 {
				fmt.Printf("  ")
			} else {
				fmt.Printf("%d ", vv)
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("autoFillGrid is: \n")
	for _, v := range S.AutoFillGrid {
		for _, vv := range v {
			if vv == 0 {
				fmt.Printf("  ")
			} else {
				fmt.Printf("%d ", vv)
			}
		}
		fmt.Printf("\n")
	}
}

// 检查在给定位置填入数字是否有效
func isValid(matrix [][]int, row, col, num int) bool {
	// 检查行
	for i := 0; i < 9; i++ {
		if matrix[row][i] == num {
			return false
		}
	}

	// 检查列
	for i := 0; i < 9; i++ {
		if matrix[i][col] == num {
			return false
		}
	}

	// 检查3x3的九宫格
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if matrix[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true
}

// 填充数独
func solveSudoku(matrix [][]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			// 找到一个空位置
			if matrix[row][col] == 0 {
				// 尝试填入数字1到9
				for num := 1; num <= 9; num++ {
					if isValid(matrix, row, col, num) {
						matrix[row][col] = num
						// 递归解决剩余的数独
						if solveSudoku(matrix) {
							return true
						}
						// 如果失败，回溯
						matrix[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func exclude(matrix [][]int, empty int) {
	n := Length
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var count = 0
	for count < empty {
		// 随机选择行和列
		row := r.Intn(n)
		col := r.Intn(n)
		if matrix[row][col] != 0 {
			matrix[row][col] = 0
			count++
		}
	}
}

func Trans(matrix [][]int, n int, m int) {
	rand.Seed(time.Now().UnixNano())

	// 进行n次行变换
	for i := 0; i < n; i++ {
		swapRows(matrix)
	}

	// 进行m次列变换
	for i := 0; i < m; i++ {
		swapCols(matrix)
	}
}

// 只在123行、456行、789行中随机交换两行
func swapRows(matrix [][]int) {
	// 确保是9*9的矩阵
	if len(matrix) != 9 || len(matrix[0]) != 9 {
		return
	}

	// 随机选择一个区块（0-2对应123行，3-5对应456行，6-8对应789行）
	block := rand.Intn(3)
	r1 := block*3 + rand.Intn(3)
	r2 := block*3 + rand.Intn(3)

	// 确保交换的两行不相同
	for r2 == r1 {
		r2 = block*3 + rand.Intn(3)
	}
	matrix[r1], matrix[r2] = matrix[r2], matrix[r1]
}

// 只在123列、456列、789列中随机交换两列
func swapCols(matrix [][]int) {
	// 确保是9*9的矩阵
	if len(matrix) != 9 || len(matrix[0]) != 9 {
		return
	}

	// 随机选择一个区块（0-2对应123列，3-5对应456列，6-8对应789列）
	block := rand.Intn(3)
	c1 := block*3 + rand.Intn(3)
	c2 := block*3 + rand.Intn(3)

	// 确保交换的两列不相同
	for c2 == c1 {
		c2 = block*3 + rand.Intn(3)
	}

	// 交换列
	for i := 0; i < len(matrix); i++ {
		matrix[i][c1], matrix[i][c2] = matrix[i][c2], matrix[i][c1]
	}
}

// 判断一个二维数组是否符合数独的条件
func IsValidSudoku(matrix [][]int) bool {
	// 检查行
	for i := 0; i < 9; i++ {
		if !isValidSet(matrix[i]) {
			return false
		}
	}

	// 检查列
	for j := 0; j < 9; j++ {
		col := make([]int, 9)
		for i := 0; i < 9; i++ {
			col[i] = matrix[i][j]
		}
		if !isValidSet(col) {
			return false
		}
	}

	// 检查每个3x3的九宫格
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if !isValidBlock(matrix, i, j) {
				return false
			}
		}
	}

	return true
}

// 判断一个数组是否包含1到9且不重复
func isValidSet(arr []int) bool {
	check := make(map[int]bool)
	for _, num := range arr {
		if num < 1 || num > 9 || check[num] {
			return false
		}
		check[num] = true
	}
	return true
}

// 判断一个3x3的九宫格是否包含1到9且不重复
func isValidBlock(matrix [][]int, startRow int, startCol int) bool {
	check := make(map[int]bool)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := matrix[startRow+i][startCol+j]
			if num < 1 || num > 9 || check[num] {
				return false
			}
			check[num] = true
		}
	}
	return true
}

func copyMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	newMatrix := make([][]int, n)
	for i := range matrix {
		newMatrix[i] = make([]int, len(matrix[i]))
		copy(newMatrix[i], matrix[i])
	}
	return newMatrix
}
