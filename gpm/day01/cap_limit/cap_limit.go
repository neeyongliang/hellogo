package cap_limit

import (
	"bytes"
	"fmt"
)

func no_limit() {
	path := []byte("AAAA/BBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	dir1 := path[:sepIndex]
	dir2 := path[sepIndex:]

	fmt.Println("dir1 =>", string(dir1))
	fmt.Println("dir2 =>", string(dir2))

	dir1 = append(dir1, "suffix"...)
	fmt.Println("dir1 ==>>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 ==>>", string(dir2)) //prints: dir2 => uffixBBBB
}

func limit() {
	path := []byte("AAAA/BBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	// 此处使用了 Limited Capacity，后续 append 会重新分配内存
	dir1 := path[:sepIndex:sepIndex]
	dir2 := path[sepIndex:]

	fmt.Println("dir1 =>", string(dir1))
	fmt.Println("dir2 =>", string(dir2))

	dir1 = append(dir1, "suffix"...)
	fmt.Println("dir1 ==>>", string(dir1))
	fmt.Println("dir2 ==>>", string(dir2))
}

func Example() {
	// 当长度不够时候，切片会扩大内存
	// 切片共享内存
	no_limit()
	limit()
}
