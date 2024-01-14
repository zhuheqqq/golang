# go学习笔记
- [go学习笔记](#go学习笔记)
  - [go基本语法](#go基本语法)
    - [内置打印函数详解](#内置打印函数详解)
    - [一些别的](#一些别的)



## go基本语法

### 声明方式

```go
var--声明变量

const--声明常量

type--声明类型

func--声明函数
```



### go变量类型

#### 值类型

`bool、int、uint、float、string、complex、array（固定长度的数组）`

#### 指针类型

```go
slice	--序列数组
map		--映射
chan	--管道
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

