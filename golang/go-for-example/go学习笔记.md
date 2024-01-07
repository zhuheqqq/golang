# go学习笔记

[TOC]



## go基本语法

### 内置打印函数详解

Printf函数:

```go
          Printf(format string, v ...) (n int, errno os.Error)

```

其中"..."表示数目可变参数，和C语言中"stdarg.h"中的宏类似。不过Go中，可变参数是通道 一个空接口（"interface {}"）和反射（reflection）库实现的。

printf方法需要加参数表示后面打印的变量是什么，在go中最简单的方法是用"%v"标志，它可以以适当的格式输出任意的类型（包括数组和结构）。

Print/Println函数：

使用这两个函数甚至不需要格式化字符串，他两就相当于封装好的printf,会自动格式化

Print函数默认将每个参数以%v格式输出

Println函数则是在Print函数基础上增加一个换行符

### 一些别的

调用Println函数自动拥有换行符并且每个参数打印出来中间自带空格  Print函数并不自带换行符

go没有三目运算符

import导入多个应该为（）