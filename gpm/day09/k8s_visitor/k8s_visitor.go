package k8s_visitor

import "fmt"

type Info struct {
	Namespace   string
	Name        string
	OtherThings string
}

type VisitorFunc func(*Info, error) error
type Visitor interface {
	Visit(VisitorFunc) error
}

// Info 实现了 Visit 方法，可以算作是 Visitor 类型
func (info *Info) Visit(fn VisitorFunc) error {
	return fn(info, nil)
}

type DecoratedVisitor struct {
	visitor    Visitor
	decorators []VisitorFunc
}

// DecoratedVisitor 实现了 Visit 方法，所以可以算是 Visitor 类型
// 用一个 DecoratedVisitor 的结构来存放所有的 VistorFunc 函数
// NewDecoratedVisitor 可以把所有的 VisitorFunc 转给它，构造 DecoratedVisitor 对象。
// DecoratedVisitor 实现了 Visit() 方法，里面就是来做一个 for-loop，顺着调用所有的 VisitorFunc
func (v DecoratedVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		if err != nil {
			return err
		}
		if err := fn(info, nil); err != nil {
			return err
		}
		for i := range v.decorators {
			if err := v.decorators[i](info, nil); err != nil {
				return err
			}
		}
		return nil
	})
}

func NewDecoratedVisitor(v Visitor, fn ...VisitorFunc) Visitor {
	if len(fn) == 0 {
		return v
	}
	return DecoratedVisitor{v, fn}
}

func NameVisitorFun(info *Info, err error) error {
	fmt.Printf("==> Name=%s, NameSpace=%s\n", info.Name, info.Namespace)
	return err
}

func OtherThingsVisitorFun(info *Info, err error) error {
	fmt.Printf("==> OtherThings=%s\n", info.OtherThings)
	return err
}

func Example() {
	fmt.Println("===> k8s visitor example...")
	info := Info{}
	var v Visitor = &info
	v = NewDecoratedVisitor(v, NameVisitorFun, OtherThingsVisitorFun)

	loadFile := func(info *Info, err error) error {
		info.Name = "Hao Chen"
		info.Namespace = "MegaEase"
		info.OtherThings = "We are running as remote team."
		return nil
	}
	v.Visit(loadFile)
}
