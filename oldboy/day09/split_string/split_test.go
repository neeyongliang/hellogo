package split_string

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	abc := "a:b:c"
	r := Split(abc, ":")
	fmt.Println(r)
	e := []string{"a", "b", "c"}
	if !reflect.DeepEqual(r, e) {
		t.Error()
	}
}
