package tetris

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

const (
	width  = 10
	height = 20
)

type Point struct {
	x, y int
}

type Tetromino struct {
	shape [][]int
	pos   Point
}

var tetrominos = [][][]int{
	{
		{1, 1, 1, 1},
	},
	{
		{1, 1},
		{1, 1},
	},
	{
		{0, 1, 0},
		{1, 1, 1},
	},
	{
		{1, 0, 0},
		{1, 1, 1},
	},
	{
		{0, 0, 1},
		{1, 1, 1},
	},
}

var board [height][width]int
var current Tetromino

func initGame() {
	rand.Seed(time.Now().UnixNano())
	newTetromino()
}

func newTetromino() {
	shape := tetrominos[rand.Intn(len(tetrominos))]
	current = Tetromino{shape: shape, pos: Point{x: width/2 - len(shape[0])/2, y: 0}}
}

func canMove(newPos Point, newShape [][]int) bool {
	for y, row := range newShape {
		for x, cell := range row {
			if cell != 0 {
				newX, newY := newPos.x+x, newPos.y+y
				if newX < 0 || newX >= width || newY < 0 || newY >= height || board[newY][newX] != 0 {
					return false
				}
			}
		}
	}
	return true
}

func rotate(shape [][]int) [][]int {
	newShape := make([][]int, len(shape[0]))
	for i := range newShape {
		newShape[i] = make([]int, len(shape))
	}
	for y, row := range shape {
		for x, cell := range row {
			newShape[x][len(shape)-y-1] = cell
		}
	}
	return newShape
}

func merge() {
	for y, row := range current.shape {
		for x, cell := range row {
			if cell != 0 {
				board[current.pos.y+y][current.pos.x+x] = cell
			}
		}
	}
	clearLines()
	newTetromino()
}

func clearLines() {
	for y := height - 1; y >= 0; y-- {
		full := true
		for x := 0; x < width; x++ {
			if board[y][x] == 0 {
				full = false
				break
			}
		}
		if full {
			for ny := y; ny > 0; ny-- {
				for nx := 0; nx < width; nx++ {
					board[ny][nx] = board[ny-1][nx]
				}
			}
			for nx := 0; nx < width; nx++ {
				board[0][nx] = 0
			}
			y++
		}
	}
}

func moveDown() bool {
	newPos := Point{current.pos.x, current.pos.y + 1}
	if canMove(newPos, current.shape) {
		current.pos = newPos
		return true
	}
	merge()
	return false
}

func moveLeft() {
	newPos := Point{current.pos.x - 1, current.pos.y}
	if canMove(newPos, current.shape) {
		current.pos = newPos
	}
}

func moveRight() {
	newPos := Point{current.pos.x + 1, current.pos.y}
	if canMove(newPos, current.shape) {
		current.pos = newPos
	}
}

func drop() {
	for moveDown() {
	}
}

func draw() {
	fmt.Print("\033[H\033[2J")
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if board[y][x] != 0 {
				fmt.Print("[]")
			} else {
				empty := true
				for dy, row := range current.shape {
					for dx, cell := range row {
						if cell != 0 && current.pos.y+dy == y && current.pos.x+dx == x {
							fmt.Print("[]")
							empty = false
						}
					}
				}
				if empty {
					fmt.Print("  ")
				}
			}
		}
		fmt.Println()
	}
	fmt.Println("Controls: a - left, d - right, s - down, w - rotate, space - drop, q - quit")
}

func Run() {
	initGame()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	go func() {
		for range ticker.C {
			moveDown()
			draw()
		}
	}()
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		switch char {
		case 'a':
			moveLeft()
		case 'd':
			moveRight()
		case 's':
			moveDown()
		case 'w':
			newShape := rotate(current.shape)
			if canMove(current.pos, newShape) {
				current.shape = newShape
			}
		case ' ':
			drop()
		case 'q':
			return
		}
		if key == keyboard.KeyEsc {
			return
		}
		draw()
	}
}

func T1() {
	for i := 0; i != 10; i = i + 1 {
		fmt.Fprintf(os.Stdout, "result is %d\r", i)
		time.Sleep(time.Second * 1)
		fmt.Println("over")
	}
}
