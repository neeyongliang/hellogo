package exercises

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type GithubResponse []struct {
	FullName string `json:"full_name"`
}

type customWriter struct{}

func (w customWriter) Write(p []byte) (n int, err error) {
	var resp GithubResponse
	json.Unmarshal(p, &resp)
	for _, r := range resp {
		fmt.Println(r.FullName)
	}
	return len(p), nil
}

func GithubResp() {
	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	// io.Copy(os.Stdout, resp.Body)
	writer := customWriter{}
	io.Copy(writer, resp.Body)
}
