- [go函数](#go函数)
  - [引用传递](#引用传递)
  - [返回值](#返回值)
  - [匿名函[go进阶.md](go%E8%BF%9B%E9%98%B6.md)数](#匿名函数)

## go函数

使用关键字func定义函数，左大括号依旧不能另起一行。如果函数没有返回值，则返回列表可以省略

```go
func test(x, y int, s string) (int, string) {
    // 类型相同的相邻参数，参数类型可合并。 多返回值必须用括号。
    n := x + y          
    return n, fmt.Sprintf(s, n)
}
//返回值类型为int string
```

### 引用传递

在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

```go
package main

import (
    "fmt"
)

/* 定义相互交换值的函数 */
func swap(x, y *int) {
    var temp int

    temp = *x /* 保存 x 的值 */
    *x = *y   /* 将 y 值赋给 x */
    *y = temp /* 将 temp 值赋给 y*/

}

func main() {
    var a, b int = 1, 2
    /*
        调用 swap() 函数
        &a 指向 a 指针，a 变量的地址
        &b 指向 b 指针，b 变量的地址
    */
    swap(&a, &b)

    fmt.Println(a, b)
}
```

默认情况下，go语言使用值传递

**无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝。引用传递是地址的拷贝，一般来说，地址拷贝更为高效。**

**map、slice、chan、指针、interface默认以引用的方式传递。**

### 返回值

Golang返回值不能用容器对象接收多返回值。只能用多个变量，或 `"_"` 忽略。

### 匿名函数

匿名函数由一个不带函数名的函数声明和函数体组成。

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    getSqrt := func(a float64) float64 {
        return math.Sqrt(a)
    }
    fmt.Println(getSqrt(4))
}
```

定义了一个名为getSqrt 的变量，初始化该变量时和之前的变量初始化有些不同，使用了func，func是定义函数的，可是这个函数和上面说的函数最大不同就是没有函数名，也就是匿名函数。