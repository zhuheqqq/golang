/*
建立与服务端的连接
进行数据收发
关闭连接
*/
package main

import (
	"bufio"
	"fmt"
	"net"
	"os" //提供与操作系统交互的功能
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer conn.Close()                       //关闭连接
	inputReader := bufio.NewReader(os.Stdin) //创建一个带缓冲的读取器，用于从标准输入读取用户输入
	for {
		input, _ := inputReader.ReadString('\n') //读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { //输入q就退出
			return
		}
		_, err = conn.Write([]byte(inputInfo)) //发送给服务端
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:]) //读取服务器发送的数据
		if err != nil {
			fmt.Println("recv failed,err", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
