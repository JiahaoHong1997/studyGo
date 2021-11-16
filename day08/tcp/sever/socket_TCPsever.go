package main

import (
	"bufio"
	"fmt"
	"net"
)

/*
 * Socket是应用层与TCP/IP协议族通信的中间软件抽象层。在设计模式中，Socket其实就是一个门面模式，它把复杂的TCP/IP协议族隐藏在Socket后面，
 * 对用户来说只需要调用Socket规定的相关函数，让Socket去组织符合指定的协议数据然后进行通信。
 * 一个TCP服务端可以同时连接很多个客户端，例如世界各地的用户使用自己电脑上的浏览器访问淘宝网。
 * 因为Go语言中创建多个goroutine实现并发非常方便和高效，所以我们可以每建立一次链接就创建一个goroutine去处理。

 * TCP服务端程序的处理流程：
 * 监听端口
 * 接收客户端请求建立链接
 * 创建goroutine处理链接。
 */

func process(conn net.Conn){
	defer conn.Close() // 处理完之后关闭这个连接
	// 针对当前的连接做数据的发送和接收
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Printf("read from conn faile, err:%v\n", err)
			break
		}
		recv := string(buf[:n])
		fmt.Println("接收到的数据：", recv)
		conn.Write([]byte("ok"))  // 通知客户端收到数据，返回"ok"
	}
}



func main() {
	// 1.开启服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000") // 监听端口
	if err != nil {
		fmt.Println("listen failed,err:%v\n", err)
		return
	}
	for {
		// 2.等待客户端来建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		// 3.启动一个单独的goroutine去处理连接
		go process(conn)
	}
}