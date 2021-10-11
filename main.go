package main

import (
	"encoding/base64"
	"fmt"
)

type Person struct {
	First string `json:"first"`
}

func main() {
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))
}
