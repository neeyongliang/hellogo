package exercises

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func MyServerAPI() {
	db := database{"Go T-Shirt": 25, "Go Jacket": 55}
	log.Fatal(http.ListenAndServe("localhost: 8080", db))
}