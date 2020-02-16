# Parse the configuration file. Currently only yml files and toml files are supported.

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/ptechen/config)
[![Build Status](https://travis-ci.com/ptechen/config.svg?branch=master)](https://travis-ci.com/ptechen/config)

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
        // config.Flge()
	    con := config.New().SetEnv("test").SetConfigFileDir("config")
	    u := &YmlParams{}
	    con.ParseFile(u)
	    fmt.Printf("%#v", u.User)
    }