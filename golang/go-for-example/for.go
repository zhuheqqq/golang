package main

import "fmt"

func main() {

	i := 1 //声明并初始化变量i
	for i <= 3 {
		fmt.Println(i) //println函数自动添加换行符
		i = i + 1
	}
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
