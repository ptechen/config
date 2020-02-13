# Parse the configuration file. Currently only yml files and toml files are supported.
example:

    package main

    import (
	    "fmt"
	    "github.com/ptechen/config"
    )

    type YmlParams struct {
	    User string `json:"user"`
    }

    func main()  {
	    con := config.New().SetEnv("test").SetConfigFilepathDir("config")
	    u := &YmlParams{}
	    con.ParseFile(u)
	    fmt.Printf("%#v", u.User)
    }