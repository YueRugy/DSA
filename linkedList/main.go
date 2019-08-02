package main

import (
	"dsa/linkedList/linked"
	_ "fmt"
)

func main() {
	ll := linked.InitList()
	linked.AddFive(ll)
	linked.Iter(ll)
	linked.Sort(ll)
	linked.Iter(ll)
	//linked.Reverse(&ll)
	//linked.Iter(ll)
	linked.ReverseShow(ll)
	linked.Iter(ll)
	linked.Reverse(&ll)
	linked.Iter(ll)

}
