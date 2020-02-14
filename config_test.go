package config

import (
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
	con := Flag().SetEnv("test").SetConfigFilepathDir("config")
	u := &YmlParams{}
	con.ParseFile(u)
	if u.User != "taochen" {
		t.Errorf("%#v", u)
	}
}

func TestYAML1(t *testing.T) {
	con := New().SetEnv("test").SetConfigFilepathDir("config").SetConfigFiletype("yaml")
	u := &YmlParams{}
	con.ParseFile(u)
	if u.User != "taochen" {
		t.Errorf("%#v", u)
	}
}

func TestYAML2(t *testing.T) {
	con := Flag()
	u := &YmlParams{}
	con.ParseFile(u)
	if u.User != "taochen" {
		t.Errorf("%#v", u)
	}
}

func TestTOML(t *testing.T) {
	con := New().SetEnv("test").SetConfigFilepathDir("config").SetConfigFiletype("toml")
	u := &TomlParams{}
	con.ParseFile(u)
	if u.Owner.User != "taochen" {
		t.Errorf("%#v", u)
	}
}
