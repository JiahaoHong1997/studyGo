package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// 我们可以自己定义一个协议，比如数据包的前4个字节为包头，里面存储的是发送的数据的长度。

// 消息编码
func Encoder(message string) ([]byte, error) {
	// 读取消息长度，转换成int32类型(4字节)
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	/*
	bytes.Buffer 是 Golang 标准库中的缓冲区，具有读写方法和可变大小的字节存储功能。缓冲区的零值是一个待使用的空缓冲区。定义如下：

	type Buffer struct {
		buf      []byte // contents are the bytes buf[off : len(buf)]
		off      int    // read at &buf[off], write at &buf[len(buf)]
		lastRead readOp // last read operation, so that Unread* can work correctly.
	}
	 */

	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}


// 消息解码
func Decoder(reader *bufio.Reader) (string, error) {
	// 读取消息的长度
	lengthByte, _ := reader.Peek(4)  // 读取前4个字节的数据
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	// 读取真正的消息数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
