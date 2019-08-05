package polandNotation

import (
	"bytes"
	"dsa/linkedList/cal"
	"dsa/linkedList/genericStack"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const sep string = ","

func Notaction(expression string) {
	s1 := genericStack.InitGenericStack()
	s2 := genericStack.InitGenericStack()
	array := []rune(expression)
	cursor := 0
	length := len(array)
	for cursor <= length-1 {
		if !cal.IsOper(array[cursor]) && !cal.IsBrackets(array[cursor]) { //如果是数字要考虑多位数
			if cursor == length-1 { //切片的最后一位肯定不是多位数
				genericStack.Push(s2, array[cursor])
			} else {
				var buff bytes.Buffer
				buff.Grow(length)
				buff.WriteRune(array[cursor])
				for !cal.IsOper(array[cursor+1]) && !cal.IsBrackets(array[cursor+1]) {
					cursor++
					buff.WriteRune(array[cursor])
				}
				genericStack.Push(s2, buff.String())
			}

		} else if cal.IsOper(array[cursor]) { //如果是操作符
			if genericStack.IsEmpty(s1) { //栈空入栈
				genericStack.Push(s1, array[cursor])
			} else {
				_, tempIter := genericStack.Peek(s1)
				temp := tempIter.(rune)
				if cal.IsBrackets(array[cursor]) { //如果是左括号直接入栈
					if cal.IsLeftBrackets(array[cursor]) {
						genericStack.Push(s1, array[cursor])
					} else { //从s1弹出直到碰到左括号为止 入s2 括号不入
						for !genericStack.IsEmpty(s1) {
							_, tempIter := genericStack.Pop(s1)
							temp := tempIter.(rune)
							if cal.IsLeftBrackets(temp) {
								break
							}
							genericStack.Push(s2, temp)
						}
					}

				} else {
					//比较优先级
					if cal.Priority(array[cursor]) > cal.Priority(temp) { //如果比栈顶的优先级高站街入栈
						genericStack.Push(s1, array[cursor])
					} else { //从s1弹出入s2
						_, operIter := genericStack.Pop(s1)
						oper := operIter.(rune)
						genericStack.Push(s2, oper)
					}
				}
			}

		}
		cursor++

	}
	//把s1内容转到s2中
	for !genericStack.IsEmpty(s1) {
		_, tempIter := genericStack.Pop(s1)
		temp := tempIter.(rune)
		genericStack.Push(s2, temp)
	}
	//颠倒s2
	for !genericStack.IsEmpty(s2) {
		_, tempIter := genericStack.Pop(s2)
		temp := tempIter.(rune)
		genericStack.Push(s1, temp)
	}

}
func CalStack(gs *genericStack.GenericStack) int {
	for !genericStack.IsEmpty(gs) {
		_, tempIter := genericStack.Pop(gs)
		temp := tempIter.(rune)
		fmt.Println(temp)
	}
	return 0
}

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
