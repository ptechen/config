package config

import (
	"fmt"
	"testing"
)

type YmlParams struct {
	User string `json:"user"`
}

type TomlParams struct {
	Owner struct {
		User string    `json:"user"`
	} `json:"owner"`
}

func TestNew(t *testing.T) {
	conf := New()
	if con.Env == "dev" && con.ConfigFileName == "config" && con.ConfigFileDir == "config" && con.ConfigFileType == "yml" {
	} else {
		t.Errorf("%#v", conf)
	}
}

func TestFlag(t *testing.T) {
	conf := Flag().SetEnv("test")
	u := &YmlParams{}
	conf.ParseFile(u)
	if u.User != "taochen" {
		t.Errorf("%#v", u)
	}
}

func TestConfig_SetConfigFileDir(t *testing.T) {
	conf := New()
	conf.SetConfigFileDir("test")
	if con.ConfigFileDir != "test" {
		t.Errorf("%#v", con)
	}
}

func TestConfig_SetConfigFileName(t *testing.T) {
	conf := New()
	conf.SetConfigFileName("config")
	if con.ConfigFileName != "config" {
		t.Errorf("%#v", con)
	}
}

func TestConfig_SetConfigFileType(t *testing.T) {
	conf := New()
	conf.SetConfigFileType("yml")
	if con.ConfigFileType != "yml" {
		t.Errorf("%#v", con)
	}
}

func TestConfig_SetEnv(t *testing.T) {
	conf := New()
	conf.SetEnv("test")
	if con.Env != "test" {
		t.Errorf("%#v", con)
	}
}

func TestConfig_YAML(t *testing.T) {
	conf := New().SetEnv("test").SetConfigFileDir("config")
	u := &YmlParams{}
	conf.ParseFile(u)
	if u.User != "taochen" {
		t.Errorf("%#v", u)
	}
}

func TestConfig_TOML(t *testing.T) {
	conf := New().SetEnv("test").SetConfigFileDir("config").SetConfigFileType("toml")
	u := &TomlParams{}
	conf.ParseFile(u)
	if u.Owner.User != "taochen" {
		t.Errorf("%#v", u)
	}
}

func TestConfig_ParseFile(t *testing.T) {
	conf := New().SetEnv("test").SetConfigFileType("toml")
	u := &TomlParams{}
	conf.ParseFile(u)
	if u.Owner.User != "taochen" {
		t.Errorf("%#v", u)
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		New()
	}
}

func BenchmarkConfig_SetEnv(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		New().SetEnv("test")
	}
}

func BenchmarkConfig_SetConfigFileDir(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		New().SetConfigFileDir("config")
	}
}

func BenchmarkConfig_SetConfigFileName(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		New().SetConfigFileName("config")
	}
}

func BenchmarkConfig_SetConfigFileType(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		New().SetConfigFileType("yml")
	}
}

func BenchmarkConfig_ParseFile(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		u := &YmlParams{}
		conf := New().SetEnv("test")
		conf.ParseFile(u)
	}
}

func BenchmarkConfig_YAML(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		u := &YmlParams{}
		conf := New().SetEnv("test")
		conf.YAML("config/test/config.yml", u)
	}
}

func BenchmarkConfig_TOML(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		u := &TomlParams{}
		conf := New().SetEnv("test")
		conf.TOML("config/test/config.toml", u)
	}
}

func ExampleNew() {
	conf := New()
	fmt.Printf("%#v", conf)
}

func ExampleConfig_SetEnv() {
	conf := New().SetEnv("test")
	fmt.Printf("%#v", conf)
	// Output: &config.Config{Env:"test", ConfigFileDir:"config", ConfigFileType:"toml", ConfigFileName:"config"}
}

func ExampleConfig_SetConfigFileDir() {
	conf := New().SetConfigFileDir("config")
	fmt.Printf("%#v", conf)
	// Output: &config.Config{Env:"test", ConfigFileDir:"config", ConfigFileType:"toml", ConfigFileName:"config"}
}

func ExampleConfig_SetConfigFileType() {
	conf := New().SetConfigFileType("yml")
	fmt.Printf("%#v", conf)
	// Output: &config.Config{Env:"test", ConfigFileDir:"config", ConfigFileType:"yml", ConfigFileName:"config"}
}

func ExampleConfig_SetConfigFileName() {
	conf := New().SetConfigFileName("config")
	fmt.Printf("%#v", conf)
	// Output: &config.Config{Env:"test", ConfigFileDir:"config", ConfigFileType:"yml", ConfigFileName:"config"}
}

func ExampleConfig_ParseFile() {
	conf := New().SetEnv("test")
	u := &YmlParams{}
	conf.ParseFile(u)
	fmt.Printf("%#v", u)
	// Output: &config.YmlParams{User:"taochen"}
}

func ExampleConfig_ParseFile2() {
	conf := New().SetEnv("test").SetConfigFileType("toml")
	u := &TomlParams{}
	conf.ParseFile(u)
	fmt.Printf("%#v", u)
	// Output: &config.TomlParams{Owner:struct { User string "json:\"user\"" }{User:"taochen"}}
	}