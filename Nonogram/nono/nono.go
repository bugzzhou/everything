package nono

import (
	"fmt"
	"math/rand"
	"time"
)

type NonogramInterface interface {
	Gen(level string) // 用于生成一个随机数织
	FillRowsAndCols() //用于填充行、列的数字展示
	Display()         //展示数组的二维数组
	Check() bool      //验证是否符合规则： 1、二维数组符合部分条件  2、行、列数字展示与二维数组一致
}

type Nonogram struct {
	Rows [][]int `json:"rows"`
	Cols [][]int `json:"cols"`
	Grid [][]int `json:"grid"`
}

// 声明的接口类
// 1
// 1
// 1
// 1
// 1
// 1
// 1
// 1
// 1
// 1
// 1
// 1
// 1
// 具体的实现

var _ NonogramInterface = (*Nonogram)(nil)

func (N *Nonogram) Gen(level string) {
	n, is := getLevel(level)

	N.Grid = genGrid(n, is)
	N.FillRowsAndCols()
}

// TODO buggzhou
// 根据二维数组，填充  行列展示数字
func (N *Nonogram) FillRowsAndCols() {
	rows := make([][]int, 0)
	cols := make([][]int, 0)
	N.Rows = rows
	N.Cols = cols
}

func (N *Nonogram) Display() {
	for _, v := range N.Grid {
		for _, vv := range v {
			fmt.Printf("%d ", vv)
		}
		fmt.Printf("\n")
	}
}

// TODO buggzhou
// 验证：
// 1、二维数组是否符合某些条件
// 2、二维数组和行列是否匹配
func (N *Nonogram) Check() bool {
	return true
}

func genGrid(n, is int) [][]int {
	// 使用随机数源创建一个新的随机数生成器
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
		res[i][i] = 1
	}
	//为了保证每一行、每一列都有至少一个数据，所以需要先在对角线增加1，再进行行、列变换 以打乱顺序
	res = trans(res, n, n)

	var count = 0
	count += n
	for count < is {
		// 随机选择行和列
		row := r.Intn(n)
		col := r.Intn(n)
		if res[row][col] == 0 {
			res[row][col] = 1
			count++
		}
	}
	return res
}

func getLevel(level string) (int, int) {
	switch level {
	case "simple":
		return 5, 13
	case "normal":
		return 10, 50
	case "hard":
		return 15, 112
	default:
		return 5, 13
	}
}

func trans(matrix [][]int, n int, m int) [][]int {
	rand.Seed(time.Now().UnixNano())

	// 进行n次行变换
	for i := 0; i < n; i++ {
		swapRows(matrix)
	}

	// 进行m次列变换
	for i := 0; i < m; i++ {
		swapCols(matrix)
	}

	return matrix
}

func swapRows(matrix [][]int) {
	rows := len(matrix)
	if rows < 2 {
		return
	}
	r1 := rand.Intn(rows)
	r2 := rand.Intn(rows)
	// 确保交换的两行不相同
	for r2 == r1 {
		r2 = rand.Intn(rows)
	}
	matrix[r1], matrix[r2] = matrix[r2], matrix[r1]
}

// swapCols 随机交换二维数组中的两列
func swapCols(matrix [][]int) {
	cols := len(matrix[0])
	if cols < 2 {
		return
	}
	c1 := rand.Intn(cols)
	c2 := rand.Intn(cols)
	// 确保交换的两列不相同
	for c2 == c1 {
		c2 = rand.Intn(cols)
	}
	// 交换列
	for i := 0; i < len(matrix); i++ {
		matrix[i][c1], matrix[i][c2] = matrix[i][c2], matrix[i][c1]
	}
}
