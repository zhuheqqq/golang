# go学习笔记
- [go学习笔记](#go学习笔记)
  - [go基本语法](#go基本语法)
    - [声明方式](#声明方式)
      - [var声明方式](#var声明方式)
      - [const声明方式](#const声明方式)
    - [go变量类型](#go变量类型)
      - [值类型](#值类型)
      - [指针类型](#指针类型)
      - [byte和rune类型](#byte和rune类型)
      - [修改字符串](#修改字符串)
      - [类型转换](#类型转换)
      - [数组](#数组)
      - [切片slice](#切片slice)
        - [创建切片](#创建切片)
        - [slice底层实现](#slice底层实现)
        - [slice操作](#slice操作)
        - [通过make创建切片](#通过make创建切片)
    - [init函数和main函数](#init函数和main函数)
    - [内置打印函数详解](#内置打印函数详解)
    - [下划线](#下划线)
    - [一些别的](#一些别的)



## go基本语法

### 声明方式

```go
var--声明变量

const--声明常量

type--声明类型

func--声明函数
```

#### var声明方式

```go
var 变量名 变量类型
```

```go
//批量声明
var(
	a string
	b int
)
//简写 短变量声明
n:=2
```

#### const声明方式

```go
const a=1
const(
	a=1
    b=2
)

```

其中const中有一个iota常量计数器

```go
const(
    n1=iota//0
    n2	//1
    n3	//2
    n4	//3
)

//常用的iota实例
//使用_跳过某些值
const(
    n1=iota	//0
    _		//1
    n2		//2
)

//iota声明中间插队
const(
    n1=iota	//0
    n2=100
    n3=iota	//2
    n4		//3
)

const (
		n1 = iota//0
		n2 = 100
		n3		//100
		n4		//100
	)
```



### go变量类型

#### 值类型

`bool、int、uint、float、string、complex(复数）、array（固定长度的数组）`

#### 指针类型

```go
slice	--序列数组
map		--映射
chan	--管道
```

#### byte和rune类型

uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。

rune类型，代表一个 UTF-8字符。

```go
// 遍历字符串
    func traversalString() {
        s := "pprof.cn博客"
        for i := 0; i < len(s); i++ { //byte
            fmt.Printf("%v(%c) ", s[i], s[i])
        }
        fmt.Println()
        for _, r := range s { //rune
            fmt.Printf("%v(%c) ", r, r)
        }
        fmt.Println()
    }
```

其输出如下

```go
 112(p) 112(p) 114(r) 111(o) 102(f) 46(.) 99(c) 110(n) 229(å) 141() 154() 229(å) 174(®) 162(¢)
    112(p) 112(p) 114(r) 111(o) 102(f) 46(.) 99(c) 110(n) 21338(博) 23458(客)
```

UTF8编码下一个中文汉字由`3~4`个字节组成，不能简单的按照字节去遍历一个包含中文的字符串，否则就会出现上面输出中第一行的结果。

字符串底层是一个byte数组，可以和[]byte类型相互转换。

字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度。 rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。

#### 修改字符串

要修改字符串，需要先将其转换为[]rune或者[]byte,完成后在转换为string。

每次都会重新分配内存，并复制字节数组

```go
func changeString() {
        s1 := "hello"
        // 强制类型转换
        byteS1 := []byte(s1)
        byteS1[0] = 'H'
        fmt.Println(string(byteS1))

        s2 := "博客"
        runeS2 := []rune(s2)
        runeS2[0] = '狗'
        fmt.Println(string(runeS2))
    }
```

#### 类型转换

go语言只有强制类型转换，转换基本语法为：

```go
T（表达式）

var a float
var c int
c=int(a)
```

#### 数组

数组定义，**数组长度一旦定义就不能更改**

```go
var a [len]int
```

多维数组

```go
var arr [len1][len2]int
```

1. **长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。**len()和cap()都可以得到数组长度
2. **数组是值类型，赋值和传参会复制整个数组而不是指针，改变副本的值不会改变本身的值。**
3. **数组进行值拷贝会造成性能问题，通常使用slice和数组指针。**
4. **支持“==”和“!=”操作符**

#### 切片slice

##### 创建切片

```go
package main

import "fmt"

func main() {
   //1.声明切片
   var s1 []int
   if s1 == nil {
      fmt.Println("是空")
   } else {
      fmt.Println("不是空")
   }
   // 2.:=
   s2 := []int{}
   // 3.make()
   var s3 []int = make([]int, 0)
   fmt.Println(s1, s2, s3)
   // 4.初始化赋值
   var s4 []int = make([]int, 0, 0)
   fmt.Println(s4)
   s5 := []int{1, 2, 3}
   fmt.Println(s5)
   // 5.从数组切片
   arr := [5]int{1, 2, 3, 4, 5}
   var s6 []int
   // 前包后不包
   s6 = arr[1:4]
   fmt.Println(s6)
}
```

##### slice底层原理

**读写操作实际目标是底层数组**

![Alt text](image.png)
![Alt text](image-1.png)

##### slice操作

```go
len(s)	//获取切片的长度
cap(s)	//测量切片最长可以达到多少

package main

import "fmt"

func main() {
   var numbers = make([]int,3,5)

   printSlice(numbers)
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```

输出结果为：

```go
len=3 cap=5 slice=[0 0 0]
```

##### 通过make创建切片

```go
 var slice []type = make([]type, len)
    slice  := make([]type, len)//省略cap则默认cap=len
    slice  := make([]type, len, cap)
```



### init函数和main函数

**init函数：**

> init函数可以用于包的初始化，init函数不能被其他函数调用，在main函数之前自动被调用
>
> 每个包可以拥有多个init函数，每个源文件也可以拥有多个init函数，多个init函数执行顺序没有明确定义

**main函数：**

main函数是程序的入口

```go
func main(){
    
}
```

**main函数和init函数的异同：**

**相同点：**

- 在定义时都不能有任何的参数和返回值有程序自动调用

**不同点：**

- init可以应用于任意包中，main函数只能用于main包中
- init函数可以重复定义多个，main函数只能定义一个



### 内置打印函数详解



**Printf函数:**

```go
          Printf(format string, v ...) (n int, errno os.Error)

```

其中"..."表示数目可变参数，和C语言中"stdarg.h"中的宏类似。不过Go中，可变参数是通道 一个空接口（"interface {}"）和反射（reflection）库实现的。

printf方法需要加参数表示后面打印的变量是什么，在go中最简单的方法是用"%v"标志，它可以以适当的格式输出任意的类型（包括数组和结构）。

**Print/Println函数：**

使用这两个函数甚至不需要格式化字符串，他两就相当于封装好的printf,会自动格式化

Print函数默认将每个参数以%v格式输出

Println函数则是在Print函数基础上增加一个换行符

### 下划线

“_”是特殊标识符，用来忽略结果。

**在import中**

```go
import_ "./hello"	//代表仅仅调用了该包的init函数，无法通过包名调用里面其他函数
```

**在代码中**

在代码中是占位符，意思是那个位置本应赋给某个值，但是咱们不需要这个值。所以就把该值赋给下划线，意思是丢掉不要。

```go
 import "database/sql"
    import _ "github.com/go-sql-driver/mysql"
```



### 一些别的

调用Println函数自动拥有换行符并且每个参数打印出来中间自带空格  Print函数并不自带换行符

go没有三目运算符

import导入多个应该为（）

