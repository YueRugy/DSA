package polandNotation

import (
	"dsa/linkedList/genericStack"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const sep string = ","

func CalStr(expression string) int {
	array := strings.Split(expression, sep)
	numStack := genericStack.InitGenericStack()
	for _, temp := range array {
		flag, _ := regexp.MatchString(`\d+`, temp)
		if flag {
			num, _ := strconv.Atoi(temp)
			genericStack.Push(numStack, num)
		} else {

			_, num1Inter := genericStack.Pop(numStack)
			_, num2Inter := genericStack.Pop(numStack)
			num1 := num1Inter.(int)
			num2 := num2Inter.(int)
			res := CalOper(num1, num2, temp)
			genericStack.Push(numStack, res)
		}
	}
	_, resIter := genericStack.Pop(numStack)
	res := resIter.(int)
	return res
}
func CalOper(num1 int, num2 int, oper string) int {
	res := 0
	switch oper {
	case "+":
		res = num1 + num2
	case "-":
		res = num2 - num1
	case "*":
		res = num1 * num2
	case "/":
		res = num2 / num1
	default:
		fmt.Println("运算符输入有误")
	}
	return res
}
