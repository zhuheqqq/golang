//找出数组中和为给定值的两个元素的下标

package main

import "fmt"

func myTest(a [5]int, target int) {
	//遍历数组
	for i := 0; i < len(a); i++ {
		other := target - a[i]
		//继续遍历
		for j := i + 1; j < len(a); j++ {
			if a[j] == other {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}

	}
}

func main() {
	b := [5]int{1, 3, 5, 8, 7}
	myTest(b, 8)
}
