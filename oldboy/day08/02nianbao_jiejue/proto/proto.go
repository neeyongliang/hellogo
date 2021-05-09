package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// 出现”粘包”的关键在于接收方不确定将要传输的数据包的大小，因此我们可以对数据包进行封包和拆包的操作。

// Encode 消息编码
func Encode(message string) ([]byte, error) {
	// 将读取消息的长度，转换为 int32（占4个字节）
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 小端写入
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}

	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode 消息解码
func Decode(reader *bufio.Reader) (string, error) {
	// 读取 4 个字节
	lengthByte, _ := reader.Peek(4)
	lengthBuffer := bytes.NewBuffer(lengthByte)
	var length int32

	err := binary.Read(lengthBuffer, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}

	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}

	return string(pack[4:]), nil
}
