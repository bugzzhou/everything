package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Map [4][4]int

func main() {
	inMap := [4][4]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	genNewOne(&inMap)
	display(&inMap)

	for {
		var action string
		fmt.Scanln(&action)
		if !canMove(&inMap, action) {
			// fmt.Println("Invalid move")
			continue
		}
		act(&inMap, action)
		display(&inMap)
	}
}

func act(inMap *[4][4]int, action string) {
	move(inMap, action)
	genNewOne(inMap)
}

func move(inMap *[4][4]int, action string) {
	switch action {
	case "w", "W":
		moveUp(inMap)
	case "s", "S":
		moveDown(inMap)
	case "a", "A":
		moveLeft(inMap)
	case "d", "D":
		moveRight(inMap)
	}
}

func moveUp(inMap *[4][4]int) {
	for j := 0; j < 4; j++ {
		raw := getColumn(inMap, j)
		merged := merge(raw)
		setColumn(inMap, j, merged)
	}
}

func moveDown(inMap *[4][4]int) {
	for j := 0; j < 4; j++ {
		raw := reverse(getColumn(inMap, j))
		merged := reverse(merge(raw))
		setColumn(inMap, j, merged)
	}
}

func moveLeft(inMap *[4][4]int) {
	for i := 0; i < 4; i++ {
		raw := inMap[i][:]
		merged := merge(raw)
		for j := 0; j < 4; j++ {
			inMap[i][j] = merged[j]
		}
	}
}

func moveRight(inMap *[4][4]int) {
	for i := 0; i < 4; i++ {
		raw := reverse(inMap[i][:])
		merged := reverse(merge(raw))
		for j := 0; j < 4; j++ {
			inMap[i][j] = merged[j]
		}
	}
}

func getColumn(inMap *[4][4]int, j int) []int {
	column := make([]int, 4)
	for i := 0; i < 4; i++ {
		column[i] = inMap[i][j]
	}
	return column
}

func setColumn(inMap *[4][4]int, j int, column []int) {
	for i := 0; i < 4; i++ {
		inMap[i][j] = column[i]
	}
}

func reverse(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func merge(row []int) []int {
	raw := make([]int, 0)
	for _, v := range row {
		if v != 0 {
			raw = append(raw, v)
		}
	}

	merged := make([]int, 4)
	t := 0
	for i := 0; i < len(raw); i++ {
		if i < len(raw)-1 && raw[i] == raw[i+1] {
			merged[t] = raw[i] * 2
			i++
		} else {
			merged[t] = raw[i]
		}
		t++
	}
	return merged
}

func genNewOne(inMap *[4][4]int) {
	emptySpaces := make([][2]int, 0)
	for i, row := range inMap {
		for j, val := range row {
			if val == 0 {
				emptySpaces = append(emptySpaces, [2]int{i, j})
			}
		}
	}

	if len(emptySpaces) > 0 {
		rand.Seed(time.Now().UnixNano())
		choice := emptySpaces[rand.Intn(len(emptySpaces))]
		inMap[choice[0]][choice[1]] = 2
	}
}

func display(inMap *[4][4]int) {
	for _, v := range inMap {
		fmt.Println(v)
	}
}

func canMove(inMap *[4][4]int, action string) bool {
	testMap := *inMap
	move(&testMap, action)
	for i := range inMap {
		for j := range inMap[i] {
			if inMap[i][j] != testMap[i][j] {
				return true
			}
		}
	}
	return false
}
