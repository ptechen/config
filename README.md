# Parse the configuration file. 
[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/ptechen/config)
[![Build Status](https://travis-ci.com/ptechen/config.svg?branch=master)](https://travis-ci.com/ptechen/config)
[![Go Report Card](https://goreportcard.com/badge/github.com/ptechen/config)](https://goreportcard.com/report/github.com/ptechen/config)
[![codecov](https://codecov.io/gh/ptechen/config/branch/master/graph/badge.svg)](https://codecov.io/gh/ptechen/config)
[![Coverage Status](https://coveralls.io/repos/github/ptechen/config/badge.svg)](https://coveralls.io/github/ptechen/config)

Currently only yml files and toml files are supported.

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
        // config.Flag()
	    con := config.New().SetEnv("test").SetConfigFileDir("config")
	    u := &YmlParams{}
	    con.ParseFile(u)
	    fmt.Printf("%#v", u.User)
    }
