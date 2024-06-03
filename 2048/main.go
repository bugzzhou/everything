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
		fmt.Println(action)
		act(&inMap, action)
		display(&inMap)
	}
	//moveUp(&inMap)
	//moveDown(&inMap)
	//display(&inMap)
	//fmt.Println(inMap)
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
		raw := make([]int, 5)
		// 去掉0
		var t int
		for i := 0; i < 4; i++ {
			if inMap[i][j] != 0 {
				raw[t] = inMap[i][j]
				t++
			}
		}

		fmt.Println("raw is:", raw)

		// 合并  自上而下
		merged := make([]int, 4)

		t = 0
		for i := 0; i < len(raw)-1; i++ {
			if raw[i] == raw[i+1] {
				merged[t] = raw[i] * 2
				t++
				i++
			} else {
				merged[t] = raw[i]
				t++
			}
		}
		//merged[t] = raw[3]
		fmt.Println("merged is:", merged)

		for i := 0; i < 4; i++ {
			inMap[i][j] = merged[i]
		}
	}
}

func moveDown(inMap *[4][4]int) {
	for j := 0; j < 4; j++ {
		raw := make([]int, 5)
		// 去掉0
		var t int
		for i := 3; i >= 0; i-- {
			if inMap[i][j] != 0 {
				raw[t] = inMap[i][j]
				t++
			}
		}
		fmt.Println("raw is:", raw)

		// 合并  自下而上
		merged := make([]int, 4)

		t = 0
		for i := 0; i < len(raw)-1; i++ {
			if raw[i] == raw[i+1] {
				merged[t] = raw[i] * 2
				t++
				i++
			} else {
				merged[t] = raw[i]
				t++
			}
		}
		fmt.Println("merged is:", merged)

		for i := 0; i < 4; i++ {
			inMap[3-i][j] = merged[i]
		}
	}
}

func moveLeft(inMap *[4][4]int) {
	for i := 0; i < 4; i++ {
		raw := make([]int, 5)
		// 去掉0
		var t int
		for j := 0; j < 4; j++ {
			if inMap[i][j] != 0 {
				raw[t] = inMap[i][j]
				t++
			}
		}
		fmt.Println("raw is:", raw)

		// 合并  自上而下
		merged := make([]int, 4)

		t = 0
		for i := 0; i < len(raw)-1; i++ {
			if raw[i] == raw[i+1] {
				merged[t] = raw[i] * 2
				t++
				i++
			} else {
				merged[t] = raw[i]
				t++
			}
		}
		fmt.Println("merged is:", merged)

		for j := 0; j < 4; j++ {
			inMap[i][j] = merged[j]
		}
	}
}

func moveRight(inMap *[4][4]int) {
	for i := 0; i < 4; i++ {
		raw := make([]int, 5)
		// 去掉0
		var t int
		for j := 3; j >= 0; j-- {
			if inMap[i][j] != 0 {
				raw[t] = inMap[i][j]
				t++
			}
		}
		fmt.Println("raw is:", raw)

		// 合并  自下而上
		merged := make([]int, 4)

		t = 0
		for i := 0; i < len(raw)-1; i++ {
			if raw[i] == raw[i+1] {
				merged[t] = raw[i] * 2
				t++
				i++
			} else {
				merged[t] = raw[i]
				t++
			}
		}
		fmt.Println("merged is:", merged)

		for j := 0; j < 4; j++ {
			inMap[i][3-j] = merged[j]
		}
	}
}

func genNewOne(inMap *[4][4]int) {
	noZeroCount := 0
	for _, v1 := range inMap {
		for _, v2 := range v1 {
			if v2 == 0 {
				noZeroCount++
			}
		}
	}

	rand.Seed(time.Now().Unix())
	randCount := rand.Intn(noZeroCount + 1)

	tmpCount := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if inMap[i][j] == 0 {
				tmpCount++
				if tmpCount == randCount {
					inMap[i][j] = 2
				}
			}
		}
	}

}

func display(inMap *[4][4]int) {
	for _, v := range inMap {
		fmt.Println(v)
	}
}
