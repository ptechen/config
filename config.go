package config

import (
	"errors"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"sync"
)

const globalConfigurationKeyword = "~"

// Config contains the parameters needed to parse the configuration file.
//
type Config struct {
	Env            string `json:"env"`             // dev、test、prod
	ConfigFileDir  string `json:"config_file_dir"` // config
	ConfigFileType string `json:"file_type"`       // yml、toml
	ConfigFileName string `json:"file_name"`       // config
}

var con *Config
var once sync.Once
var onceFlag sync.Once
var env string
var configFileDir string
var configFileType string
var configFileName string
// Flag are arguments passed from outside the program.
//
func Flag() *Config {
	onceFlag.Do(func() {
		flag.StringVar(&env, "e", "dev", "Operating environment")
		flag.StringVar(&configFileDir, "cfd", "config", "Configuration file directory")
		flag.StringVar(&configFileType, "cft", "yml", "Configuration file suffix")
		flag.StringVar(&configFileName, "cfn", "config", "Configuration file name")
		flag.Parse()
		con = New()
		con.Env = env
		con.ConfigFileDir = configFileDir
		con.ConfigFileType = configFileType
		con.ConfigFileName = configFileName
	})
	return con
}

// New the Config structure and set the default values.
//
func New() *Config {
	once.Do(func() {
		con = &Config{
			Env:            "dev",    // 默认dev
			ConfigFileDir:  "config", // 默认 config
			ConfigFileType: "yml",    // 默认 yml
			ConfigFileName: "config", // 默认 config
		}
	})
	return con
}

// SetEnv the value of Config.Env.
//
func (p *Config) SetEnv(env string) *Config {
	p.Env = env
	return p
}

// SetConfigFileDir the value of Config.ConfigFileDir.
//
func (p *Config) SetConfigFileDir(fileDir string) *Config {
	p.ConfigFileDir = fileDir
	return p
}

// SetConfigFileType the value of Config.ConfigFileType.
//
func (p *Config) SetConfigFileType(fileType string) *Config {
	p.ConfigFileType = fileType
	return p
}

// SetConfigFileName the value of Config.ConfigFileName.
//
func (p *Config) SetConfigFileName(fileName string) *Config {
	p.ConfigFileName = fileName
	return p
}

// ParseFile is the method to parse the configuration file.
//
func (p *Config) ParseFile(res interface{}) {
	filename := p.filename()
	if p.ConfigFileType == "yml" || p.ConfigFileType == "yaml" {
		p.YAML(filename, res)
	} else if p.ConfigFileType == "toml" {
		p.TOML(filename, res)
	} else {
		panic(errors.New("Currently only yml and toml file types are supported."))
	}
}

// YAML reads Configuration from a configuration.yml file.
//
// Accepts the absolute path of the cfg.yml.
// An error will be shown to the user via panic with the error message.
// Error may occur when the cfg.yml doesn't exists or is not formatted correctly.
//
func (p *Config) YAML(filename string, res interface{}) {

	// check for globe configuration file and use that, otherwise
	// return the default configuration if file doesn't exist.
	if filename == globalConfigurationKeyword {
		filename = homeConfigurationFilename(".yml")
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			panic(errors.New("default configuration file '" + filename + "' does not exist"))
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
// Accepts the absolute path of the configuration file.
// An error will be shown to the user via panic with the error message.
// Error may occur when the file doesn't exists or is not formatted correctly.
//
func (p *Config) TOML(filename string, res interface{}) {

	if filename == globalConfigurationKeyword {
		filename = homeConfigurationFilename(".tml")
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			panic(errors.New("default configuration file '" + filename + "' does not exist"))
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
	if p.ConfigFileType == "yml" || p.ConfigFileType == "yaml" {
		return fmt.Sprintf("%s/%s/%s.%s", p.ConfigFileDir, p.Env, p.ConfigFileName, p.ConfigFileType)
	} else if p.ConfigFileType == "toml" {
		return fmt.Sprintf("%s/%s/%s.%s", p.ConfigFileDir, p.Env, p.ConfigFileName, p.ConfigFileType)
	} else {
		panic(errors.New("Currently only yml and toml file types are supported."))
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
