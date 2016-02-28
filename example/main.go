// Package main provides ...
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/astavonin/five00px/client"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Print("go run main.go")
	fmt.Println("  --config <json configuration file>")
	fmt.Println("Configuration file format:")
	fmt.Println("{\n\t\"ConsumerKey\": \"see https://github.com/500px/api-documentation\",")
	fmt.Println("\t\"ConsumerSecret\": \"see https://github.com/500px/api-documentation\",")
	fmt.Println("\t\"Token\": {}\n}")
}

type Config struct {
	ConsumerKey    string
	ConsumerSecret string
	Token          *five00px.AccessToken
}

func main() {

	var configPath = flag.String(
		"config",
		"",
		"PatD9D4B9h to configuration file")
	flag.Parse()
	if len(*configPath) == 0 {
		usage()
		os.Exit(1)
	}

	b, err := ioutil.ReadFile(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	var c Config
	err = json.Unmarshal(b, &c)
	if err != nil {
		log.Fatal(err)
	}

	f00 := five00px.New(c.ConsumerKey, c.ConsumerSecret)

	if c.Token.Token == "" || c.Token.Secret == "" {
		t, err := f00.Auth()
		if err != nil {
			log.Fatal(err)
		}
		c.Token = t
	} else {
		f00.Restore(c.Token)
	}

	b, err = json.MarshalIndent(c, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	err = ioutil.WriteFile(*configPath, b, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	f, err := f00.Friends(9091479, &five00px.Page{1, 1})

	for _, u := range f.Users {
		fmt.Println(u.Avatars.Default)
	}
	u, err := f00.UserByID(9091479)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.City)
}
