package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Chess struct {
	row int
	col int
	val int
}

func main() {
	//构建原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2
	//显示原始数组
	//for _, v := range chessMap {
	//	for _, v1 := range v {
	//		fmt.Print(v1)
	//	}
	//	fmt.Println()
	//}
	//构建稀疏数组
	var sparseSlice []Chess
	//原始数组的行列默认值
	c1 := Chess{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseSlice = append(sparseSlice, c1)
	//稀疏数组数据填充
	for i, v := range chessMap {
		for j, v1 := range v {
			if v1 != 0 {
				c := Chess{
					row: i,
					col: j,
					val: v1,
				}
				sparseSlice = append(sparseSlice, c)
			}
		}
	}
	//显示稀疏数组数据
	//for _, v := range sparseSlice {
	//	fmt.Println(v)
	//}
	//写入磁盘
	f, err := os.Create("chess.txt")
	if err != nil {
		fmt.Println("create file error")
	}
	defer f.Close()
	write := bufio.NewWriter(f)
	for _, v := range sparseSlice {
		str := strconv.Itoa(v.row) + "\t" + strconv.Itoa(v.col) + "\t" + strconv.Itoa(v.val) + "\r\n"
		write.WriteString(str)
	}
	write.Flush()
	//稀疏数组转换成原始数组不在从文件中取
	//切片的第一个元素创建原始数组
	//ch := sparseSlice[0]
	//row := ch.row
	//col := ch.col
	//fmt.Println(row, col)
	var data [11][11]int
	for i, v := range sparseSlice {
		if i != 0 {
			data[v.row][v.col] = v.val
		}
	}
	for _, v := range data {
		for _, v1 := range v {
			fmt.Print(v1)
		}
		fmt.Println()
	}

}
