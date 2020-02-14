package main

import (
	"fmt"
	"github.com/ptechen/config"
)
type YmlParams struct {
	User string `json:"user"`
}

func main() {
	con := config.Flag()
	u := &YmlParams{}
	con.ParseFile(u)
	fmt.Println(u.User)
}