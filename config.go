package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

const globalConfigurationKeyword = "~"

type Config struct {
	Env           string `json:"env"` // dev、test、prod
	ConfigFileDir string `json:"config_file_dir"` // config
	FileType string `json:"file_type"` // yml、toml
	FileName string `json:"file_name"` // config
}


func New() *Config {
	return &Config{
		Env: "dev",  // 默认dev
		ConfigFileDir: "config", // 默认 config
		FileType: "yml", // 默认 yml
		FileName: "config", // 默认 config
	}
}

func (p *Config) SetEnv(env string) *Config {
	p.Env = env
	return p
}

func (p *Config) SetConfigFilepathDir(filepathDir string) *Config {
	p.ConfigFileDir = filepathDir
	return p
}

func (p *Config) SetConfigFiletype(fileType string) *Config {
	p.FileType = fileType
	return p
}

func (p *Config) SetConfigFileName(fileType string) *Config {
	p.FileType = fileType
	return p
}

func (p *Config) ParseFile(res interface{}){
	filename := p.filename()
	if p.FileType == "yml" || p.FileType == "yaml" {
		p.YAML(filename, res)
	} else if p.FileType == "toml" {
		p.TOML(filename, res)
	} else {
		panic(fmt.Errorf("Currently only yml and toml file types are supported."))
	}
}

func (p *Config)YAML(filename string, res interface{}) {

	// check for globe configuration file and use that, otherwise
	// return the default configuration if file doesn't exist.
	if filename == globalConfigurationKeyword {
		filename = homeConfigurationFilename(".yml")
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			panic("default configuration file '" + filename + "' does not exist")
		}
	}

	err := parseYAML(filename, res)
	if err != nil {
		panic(err)
	}
}

// TOML reads Configuration from a toml-compatible document file.
// Read more about toml's implementation at:
// https://github.com/toml-lang/toml
//
//
// Accepts the absolute path of the configuration file.
// An error will be shown to the user via panic with the error message.
// Error may occur when the file doesn't exists or is not formatted correctly.
//


func (p *Config)TOML(filename string, res interface{}){

	if filename == globalConfigurationKeyword {
		filename = homeConfigurationFilename(".tml")
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			panic(err)
		}
	}

	// get the abs
	// which will try to find the 'filename' from current workind dir too.
	tomlAbsPath, err := filepath.Abs(filename)
	if err != nil {
		panic(err)
	}

	// read the raw contents of the file
	data, err := ioutil.ReadFile(tomlAbsPath)
	if err != nil {
		panic(err)
	}
	// put the file's contents as toml to the default configuration(c)
	if _, err := toml.Decode(string(data), res); err != nil {
		panic(err)
	}

}

func (p *Config) filename() string {
	if p.FileType == "yml" || p.FileType == "yaml"{
		return fmt.Sprintf("%s/%s/%s.%s", p.ConfigFileDir, p.Env, p.FileName, p.FileType)
	} else if p.FileType == "toml" {
		return fmt.Sprintf("%s/%s/%s.%s", p.ConfigFileDir, p.Env, p.FileName, p.FileType)
	} else {
		panic(fmt.Errorf("Currently only yml and toml file types are supported."))
	}
}

func homeConfigurationFilename(ext string) string {
	return filepath.Join(homeDir(), ext)
}

func homeDir() (home string) {
	u, err := user.Current()
	if u != nil && err == nil {
		home = u.HomeDir
	}

	if home == "" {
		home = os.Getenv("HOME")
	}

	if home == "" {
		if runtime.GOOS == "plan9" {
			home = os.Getenv("home")
		} else if runtime.GOOS == "windows" {
			home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
			if home == "" {
				home = os.Getenv("USERPROFILE")
			}
		}
	}

	return
}

func parseYAML(filename string, res interface{}) error {
	// get the abs
	// which will try to find the 'filename' from current workind dir too.
	yamlAbsPath, err := filepath.Abs(filename)
	if err != nil {
		return err
	}

	// read the raw contents of the file
	data, err := ioutil.ReadFile(yamlAbsPath)
	if err != nil {
		return err
	}

	// put the file's contents as yaml to the default configuration(c)
	if err := yaml.Unmarshal(data, res); err != nil {
		return err
	}
	return nil
}

// YAML reads Configuration from a configuration.yml file.
//
// Accepts the absolute path of the cfg.yml.
// An error will be shown to the user via panic with the error message.
// Error may occur when the cfg.yml doesn't exists or is not formatted correctly.
//





