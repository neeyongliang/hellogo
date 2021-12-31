package with_interface

import (
	"fmt"
)

type Country struct {
	Name string
}

type City struct {
	Name string
}

type Stringable interface {
	ToString() string
}

// 分别实现 Stringable 接口

func (c Country) ToString() string {
	return "Country = " + c.Name
}

func (c City) ToString() string {
	return "City = " + c.Name
}

// 接收 Stringable 类型参数
func PrintStr(p Stringable) {
	fmt.Println(p.ToString())
}

func Example() {
	d1 := Country{"China"}
	d2 := City{"Shanghai"}

	// 强类型检测，如果 City 没有实现 Stringable 就会报错
	var _ Stringable = (*City)(nil)
	PrintStr(d1)
	PrintStr(d2)
}
