package error_handle

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type MyError struct {
}

func (e *MyError) Error() string {
	return "this is my error occur..."
}

func testRead() error {
	_, err := os.Open("aaa")
	if err != nil {
		// test MyError
		return errors.Wrap(err, "read failed")
	}
	return nil
}

func Example() {
	err := testRead()
	switch errors.Cause(err).(type) {
	case *MyError:
		fmt.Println("handle my error")
	default:
		fmt.Println("handle default error")
	}
}
