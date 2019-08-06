package maze

import (
	"fmt"
)

func Maze() {
	var mazeMap [8][7]int
	for index := 0; index < 7; index++ {
		mazeMap[0][index], mazeMap[7][index] = 1, 1
	}
	for index := 0; index < 8; index++ {
		mazeMap[index][0], mazeMap[index][6] = 1, 1
	}
	mazeMap[3][1], mazeMap[3][2] = 1, 1
	GoMaze(&mazeMap, 1, 1)
	Iter(&mazeMap)

}
func GoMaze(mm *[8][7]int, i int, j int) bool {
	if (*mm)[6][5] == 2 {
		return true
	}
	if (*mm)[i][j] == 0 {
		(*mm)[i][j] = 2 //假设能走通
		//制定策略down->right->up->left
		if GoMaze(mm, i+1, j) {
			return true
		} else if GoMaze(mm, i, j+1) {
			return true
		} else if GoMaze(mm, i-1, j) {
			return true
		} else if GoMaze(mm, i, j-1) {
			return true
		} else {
			(*mm)[i][j] = 3
			return false
		}
	} else {
		return false
	}
}

func Iter(pm *[8][7]int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Printf("%d  ", (*pm)[i][j])
		}
		fmt.Println()
	}
}
