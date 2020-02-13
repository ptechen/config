package parse

import (
	"fmt"
	"testing"
	"time"
)

type YmlParams struct {
	User string `json:"user"`
}

type TomlParams struct {
	Owner struct{
		User string    `json:"user"`
		Dob  time.Time `json:"dob"`
	} `json:"owner"`

}

func TestYAML(t *testing.T) {
	con := New().SetEnv("test").SetConfigFilepathDir("../config")
	u := &YmlParams{}
	con.ParseFile(u)
	fmt.Printf("%#v", u.User)
}

func TestYAML1(t *testing.T) {
	con := New().SetEnv("test").SetConfigFilepathDir("../config").SetConfigFiletype("yaml")
	u := &YmlParams{}
	con.ParseFile(u)
	fmt.Printf("%#v", u.User)
}

func TestTOML(t *testing.T) {
	con := New().SetEnv("test").SetConfigFilepathDir("../config").SetConfigFiletype("toml")
	u := &TomlParams{}
	con.ParseFile(u)
	fmt.Printf("%#v", u.Owner.User)
}
