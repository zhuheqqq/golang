# go语言流程控制

此篇笔记主要是针对go语言与c、cpp流程控制不同之处，相同之处不再赘述。

## switch

switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。

```go
package main

import "fmt"

func main() {
    var x interface{}//声明了一个空接口变量 x，这意味着 x 可以存储任意类型的值。
    //写法一：
    switch i := x.(type) { // 带初始化语句
    case nil:
        fmt.Printf(" x 的类型 :%T\r\n", i)
    case int:
        fmt.Printf("x 是 int 型")
    case float64:
        fmt.Printf("x 是 float64 型")
    case func(int) float64:
        fmt.Printf("x 是 func(int) 型")
    case bool, string:
        fmt.Printf("x 是 bool 或 string 型")
    default:
        fmt.Printf("未知型")
    }
    //写法二
    var j = 0
    switch j {
    case 0:
    case 1:
        fmt.Println("1")
    case 2:
        fmt.Println("2")
    default:
        fmt.Println("def")
    }
    //写法三
    var k = 0
    switch k {
    case 0:
        println("fallthrough")
        fallthrough
        /*
            Go的switch非常灵活，表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；
            而如果switch没有表达式，它会匹配true。
            Go里面switch默认相当于每个case最后带有break，
            匹配成功后不会自动向下执行其他case，而是跳出整个switch,
            但是可以使用fallthrough强制执行后面的case代码。
        */
    case 1:
        fmt.Println("1")
    case 2:
        fmt.Println("2")
    default:
        fmt.Println("def")
    }
    //写法三
    var m = 0
    switch m {
    case 0, 1:
        fmt.Println("1")
    case 2:
        fmt.Println("2")
    default:
        fmt.Println("def")
    }
    //写法四
    var n = 0
    switch { //省略条件表达式，可当 if...else if...else
    case n > 0 && n < 10:
        fmt.Println("i > 0 and i < 10")
    case n > 10 && n < 20:
        fmt.Println("i > 10 and i < 20")
    default:
        fmt.Println("def")
    }
}
```

## selete

selete主要用于处理异步IO操作，语法类似switch语句，但是每个case必须是一个通信操作，要么是发送要么是接收。

**selete随机执行一个可运行的case，如果没有case可运行将会阻塞直到有case可运行。**

语法如下：

```go
    select { //不停的在这里检测
    case <-chanl : //检测有没有数据可以读
    //如果chanl成功读取到数据，则进行该case处理语句
    case chan2 <- 1 : //检测有没有可以写
    //如果成功向chan2写入数据，则进行该case处理语句


    //假如没有default，那么在以上两个条件都不成立的情况下，就会在此阻塞//一般default会不写在里面，select中的default子句总是可运行的，因为会很消耗CPU资源
    default:
    //如果以上都没有符合条件，那么则进行default处理流程
    }
```

> 每个case都必须是一个通信
> 所有channel表达式都会被求值
> 所有被发送的表达式都会被求值
> 如果任意某个通信可以进行，它就执行；其他被忽略。
> 如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
> 否则：
> 如果有default子句，则执行该语句。
> 如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。

```go
package main

import "fmt"

func main() {
   var c1, c2, c3 chan int//声明三个整型通道变量
   var i1, i2 int
   select {
      case i1 = <-c1://如果可以从通道c1中接收到数据，则执行该case分支
         fmt.Printf("received ", i1, " from c1\n")
      case c2 <- i2:
         fmt.Printf("sent ", i2, " to c2\n")
      case i3, ok := (<-c3):  // same as: i3, ok := <-c3
         if ok {
            fmt.Printf("received ", i3, " from c3\n")
         } else {
            fmt.Printf("c3 is closed\n")
         }
      default:
         fmt.Printf("no communication\n")
   }    
}
```

在上述代码中，隐式声明变量i3,并且`ok` 是一个布尔类型的变量，通常用于接收通道操作的第二个返回值，用于指示通道操作是否成功。

**selete可以监听channel的读写操作，当case中channel读写操作为非阻塞状态时，会出发相应的动作。**

### 超时判断

```go
//比如在下面的场景中，使用全局resChan来接受response，如果时间超过3S,resChan中还没有数据返回，则第二条case将执行
var resChan = make(chan int)//声明并初始化了一个全局的整型通道 resChan
// do request
func test() {
    select {
    case data := <-resChan:
        doData(data)
    case <-time.After(time.Second * 3):
        fmt.Println("request time out")
    }
}

func doData(data int) {
    //...
}
```

### 退出

```go
//主线程（协程）中如下：
var shouldQuit=make(chan struct{})
fun main(){
    {
        //loop
    }
    //...out of the loop
    select {
        case <-c.shouldQuit:
            cleanUp()
            return
        default:
        }
    //...
}

//再另外一个协程中，如果运行遇到非法操作或不可处理的错误，就向shouldQuit发送数据通知程序停止运行
close(shouldQuit)
```

### 判断channel是否阻塞

```go
//在某些情况下是存在不希望channel缓存满了的需求的，可以用如下方法判断
ch := make (chan int, 5)
//...
data：=0
select {
case ch <- data:
default:
    //做相应操作，比如丢弃data。视需求而定
}
```

## range

range类似迭代器操作，返回 (索引, 值) 或 (键, 值)。

```go
package main

func main() {
    s := "abc"
    // 忽略 2nd value，支持 string/array/slice/map。
    for i := range s {
        println(s[i])
    }
    // 忽略 index。
    for _, c := range s {
        println(c)
    }
    // 忽略全部返回值，仅迭代。
    for range s {

    }

    m := map[string]int{"a": 1, "b": 2}
    // 返回 (key, value)。
    for k, v := range m {
        println(k, v)
    }
}
```

**注意，range 会复制对象。**改成引用类型，底层数据不会被复制。