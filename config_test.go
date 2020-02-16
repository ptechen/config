package config

import (
	"testing"
	"time"
)

type YmlParams struct {
	User string `json:"user"`
}

type TomlParams struct {
	Owner struct {
		User string    `json:"user"`
		Dob  time.Time `json:"dob"`
	} `json:"owner"`
}

func TestNew(t *testing.T) {
	con := New()
	if con.Env == "dev" && con.ConfigFileName == "config" && con.ConfigFileDir == "config" && con.ConfigFileType == "yml" {
	} else {
		t.Errorf("%#v", con)
	}
}

func TestFlag(t *testing.T) {
	con := Flag().SetEnv("test")
	u := &YmlParams{}
	con.ParseFile(u)
	if u.User != "taochen" {
		t.Errorf("%#v", u)
	}
}

func TestConfig_SetConfigFileDir(t *testing.T) {
	con := New()
	con.SetConfigFileDir("test")
	if con.ConfigFileDir != "test" {
		t.Errorf("%#v", con)
	}
}

func TestConfig_SetConfigFileName(t *testing.T) {
	con := New()
	con.SetConfigFileName("config")
	if con.ConfigFileName != "config" {
		t.Errorf("%#v", con)
	}
}

func TestConfig_SetConfigFileType(t *testing.T) {
	con := New()
	con.SetConfigFileType("yml")
	if con.ConfigFileType != "yml" {
		t.Errorf("%#v", con)
	}
}

func TestConfig_SetEnv(t *testing.T) {
	con := New()
	con.SetEnv("test")
	if con.Env != "test" {
		t.Errorf("%#v", con)
	}
}

func TestConfig_YAML(t *testing.T) {
	con := New().SetEnv("test").SetConfigFileDir("config")
	u := &YmlParams{}
	con.ParseFile(u)
	if u.User != "taochen" {
		t.Errorf("%#v", u)
	}
}

func TestConfig_TOML(t *testing.T) {
	con := New().SetEnv("test").SetConfigFileDir("config").SetConfigFileType("toml")
	u := &TomlParams{}
	con.ParseFile(u)
	if u.Owner.User != "taochen" {
		t.Errorf("%#v", u)
	}
}

func TestConfig_ParseFile(t *testing.T) {
	con := New().SetEnv("test").SetConfigFileType("toml")
	u := &TomlParams{}
	con.ParseFile(u)
	if u.Owner.User != "taochen" {
		t.Errorf("%#v", u)
	}
}

