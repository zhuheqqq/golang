package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 求元素和
func sumarr(a [10]int) int {
	var sum int = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum

}

func main() {
	rand.Seed(time.Now().Unix()) //生成随机数
	var b [10]int
	for i := 0; i < len(b); i++ {
		//产生随机数
		b[i] = rand.Intn(1000)
	}
	sum := sumarr(b)
	fmt.Println("sum=%d\n", sum)

}
