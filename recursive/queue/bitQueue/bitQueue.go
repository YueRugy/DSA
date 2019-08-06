package bitQueue

import (
	"fmt"
)

var count = 0

func Queue() {
	goQueue(0, 0, 0)
	fmt.Println(count)
}

func goQueue(l, c, r uint8) {
	if c == 255 { //如果列的位中全是1表明已经全部完成
		count++
		return
	}
	// |运算有1为1
	cond := ^(l | c | r) //找出全部这一行符合条件的位置，0符合1 不符合取反后1符合0不符合
	for cond != 0 {
		//&运算有0为0
		bit := cond & (^cond + 1) //找出最右边的1
		cond -= bit
		goQueue((l|bit)<<1, cond|bit, (r|bit)>>1)
	}
}
