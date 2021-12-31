package error_switch

import (
	"encoding/json"
	"fmt"
	"errors"
)

type ZeroDivisonError struct {
}

type NullPointerError struct {
}

// 自定义错误需要实现 Error 方法，接口编程
func (err *ZeroDivisonError) Error() string {
	return "zero divison occur"
}

func (err *NullPointerError) Error() string {
	return "null pointer error"
}

func Example() {
	myerr := errors.New("my error")
	if myerr != nil {
		switch myerr.(type) {
		case *json.SyntaxError:
			fmt.Println("json syntax error...")
		case *ZeroDivisonError:
			fmt.Println("zero divison error...")
		case *NullPointerError:
			fmt.Println("null pointer error...")
		default:
			fmt.Println("default error...")
		}
	}
}
