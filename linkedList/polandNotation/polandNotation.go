package polandNotation

import (
	"bytes"
	"dsa/linkedList/cal"
	"dsa/linkedList/genericStack"
	_ "dsa/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const sep string = ","

func Cal(str string) int {
	str = Notaction(str)
	return CalStr(str)
}

func Notaction(expression string) string {
	s1 := genericStack.InitGenericStack()
	s2 := genericStack.InitGenericStack()
	array := []rune(expression)
	cursor := 0
	length := len(array)
	for cursor <= length-1 {
		if !cal.IsOper(array[cursor]) && !cal.IsBrackets(array[cursor]) { //如果是数字要考虑多位数
			if cursor == length-1 { //切片的最后一位肯定不是多位数
				genericStack.Push(s2, string(array[cursor]))
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
				genericStack.Push(s1, string(array[cursor]))
			} else {

				for !genericStack.IsEmpty(s1) { //循环与s1(运算符栈）栈顶比较
					_, tempIter := genericStack.Peek(s1)
					temp := tempIter.(string)
					tempArray := []rune(temp)
					//比较优先级
					if cal.Priority(array[cursor]) > cal.Priority(tempArray[0]) || cal.IsLeftBrackets(tempArray[0]) { //如果比栈顶的优先级高站街入栈并退出循环
						break
					} else { //循环与栈顶元素比较 并从s1弹出入s2 新的运算符入s1栈
						_, operIter := genericStack.Pop(s1)
						oper := operIter.(string)
						genericStack.Push(s2, oper)
					}
				}
				genericStack.Push(s1, string(array[cursor]))
			}

		} else if cal.IsBrackets(array[cursor]) { //如果是左括号直接入栈
			if cal.IsLeftBrackets(array[cursor]) {
				genericStack.Push(s1, string(array[cursor]))
			} else { //从s1弹出直到碰到左括号为止 入s2 括号不入
				for !genericStack.IsEmpty(s1) {
					_, tempIter := genericStack.Pop(s1)
					temp := tempIter.(string)
					tempArray := []rune(temp)
					if cal.IsLeftBrackets(tempArray[0]) {
						break
					}
					genericStack.Push(s2, temp)
				}
			}
		}
		cursor++
	}
	//把s1内容转到s2中
	for !genericStack.IsEmpty(s1) {
		_, tempIter := genericStack.Pop(s1)
		temp := tempIter.(string)
		genericStack.Push(s2, temp)
	}
	//颠倒s2
	for !genericStack.IsEmpty(s2) {
		_, tempIter := genericStack.Pop(s2)
		temp := tempIter.(string)
		genericStack.Push(s1, temp)
	}
	return getStr(s1)
}
func getStr(gs *genericStack.GenericStack) string {
	var buff bytes.Buffer
	for !genericStack.IsEmpty(gs) {
		_, tempIter := genericStack.Pop(gs)
		tempStr := tempIter.(string)
		buff.WriteString(tempStr)
		buff.WriteString(",")
	}
	res := buff.String()[0 : len(buff.String())-1]
	fmt.Println(res)
	return res
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
