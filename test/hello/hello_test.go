package hello

import "testing"

func TestSayHello(t *testing.T) {
	const helloGolang = "Hello, Golang!"
	res := SayHello()
	if res != helloGolang {
		t.Errorf("result not say Hello, Golang!")
	}
}
