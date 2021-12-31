package resource_clean

import (
	"io"
	"log"
	"os"
)

func Open(name string) (*os.File, error) {
	return os.Open(name)
}

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Println(err)
	}
}

func Example() {
	r, err := Open("a")
	if err != nil {
		// log.Fatalf("error opening 'a'\n")
		log.Printf("error opening 'a'\n")
	}

	defer Close(r)

	r, err = Open("b")
	if err != nil {
		log.Printf("error opening 'b'\n")
	}

	defer Close(r)
}
