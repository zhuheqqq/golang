/*
监听端口
接收客户端请求建立连接
创建goroutine处理连接
*/
package main

import (
	"bufio" //提供带缓冲的I/O操作，用于读取数据
	"fmt"
	"net" //提供用于网络编程的基本功能
)

// TCP server端
func process(conn net.Conn) {
	defer conn.Close() //函数退出时延迟关闭连接

	for {
		reader := bufio.NewReader(conn) //创建一个带缓冲的读取器，用于从连接中读取数据
		var buf [128]byte
		n, err := reader.Read(buf[:]) //读取数据
		if err != nil {
			fmt.Println("read from client failed,err:", err)
			break
		}
		recvStr := string(buf[:n]) //将接收的字节数据转换成字符串
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr)) //将接收到的数据作为响应发送回客户端
	}

}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("accept failed,err:", err)
		return
	}
	for {
		conn, err := listen.Accept() //建立连接
		if err != nil {
			fmt.Println("accept failed,err:", err)
			continue
		}
		go process(conn) //启动一个goroutine处理连接
	}

}
