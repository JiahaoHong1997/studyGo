package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/* *
 * 一个TCP客户端进行TCP通信的流程如下：

 * 建立与服务端的链接
 * 进行数据收发
 * 关闭链接
*/

func main() {
	// 1.与服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("dial fail,err:%v\n", err)
		return
	}
	// 2.利用该连接进行数据的发送和接收
	input := bufio.NewReader(os.Stdin)  // 从终端读取用户的输入
	for {
		s, _ := input.ReadString('\n') // 遇到'\n'则结束
		s = strings.TrimSpace(s) // 把空格去掉
		if strings.ToUpper(s) == "Q" { // 输入内容为Q表示退出
			return
		}
		// 给服务端发消息
		_, err := conn.Write([]byte(s))
		if err != nil {
			fmt.Printf("send failed,err:%v\n", err)
			return
		}
		// 从服务端接收回复的消息
		var buf [1024]byte
		n, err := conn.Read(buf[:])  // Read返回当初的字节数
		if err != nil {
			fmt.Printf("read failed, err:%v\n", err)
		}
		fmt.Println("收到服务端回复", string(buf[:n]))
	}
}
