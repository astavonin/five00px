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

func loadConfig(path string) *Config {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var c Config
	err = json.Unmarshal(b, &c)
	if err != nil {
		log.Fatal(err)
	}
	return &c
}

func storeConfig(path string, c *Config) {
	b, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(path, b, 0666)
	if err != nil {
		log.Panicln(err)
	}
}

func main() {

	var configPath = flag.String(
		"config",
		"",
		"Path to configuration file")
	flag.Parse()
	if len(*configPath) == 0 {
		usage()
		os.Exit(1)
	}

	c := loadConfig(*configPath)

	f00 := five00px.New(c.ConsumerKey, c.ConsumerSecret, nil)

	if c.Token == nil {
		t, err := f00.Auth()
		if err != nil {
			log.Fatal(err)
		}
		c.Token = t
		storeConfig(*configPath, c)
	} else {
		f00.Restore(c.Token)
	}

	//f, err := os.Open("../client/test_data/test_img.jpg")
	//if err != nil {
	//log.Fatal(err)
	//}
	//defer f.Close()

	//var upInfo = five00px.UploadInfo{
	//Category:    five00px.CategoryBW,
	//Description: "test description",
	//Name:        "test name",
	//PhotoStream: f,
	//}
	//photo, err := f00.AddPhoto(upInfo)
	//if err != nil {
	//log.Fatal(err)
	//}
	//fmt.Println(photo)

	fr, err := f00.ListFriends(9091479, &five00px.Page{1, 1})

	for _, u := range fr.Users {
		fmt.Println(u.Avatars.Default)
	}
	u, err := f00.GetUserByID(9091479)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.City)

	//fl, err := f00.Followers(9091479, nil)
	//if err != nil {
	//log.Fatal(err)
	//}
	//fmt.Println(fl.FollowersCount)

	//u, err = f00.AddFriend(42)
	//fmt.Println(err)

	//u, err = f00.DelFriend(42)
	//fmt.Println(err)

	//s := five00px.StreamCriterias{
	//Feature: five00px.FeaturePopular,
	//}
	//p := five00px.Page{
	//Rpp: 3,
	//}
	//photos, err := f00.Photos(s, &p)
	//fmt.Println("", len(photos.Photos), photos.TotalPages, photos.TotalItems)

	//sCrit := five00px.SearchCriterias{
	//Term:        "test",
	//Tag:         "best",
	//LicenseType: five00px.LicAll,
	//}

	//photos, err = f00.PhotosSearch(sCrit, &p)

	//if err != nil || photos.CurrentPage != 1 || photos.TotalItems != 84 ||
	//photos.TotalPages != 28 || len(photos.Photos) != 3 {
	//log.Fatal(err)
	//}

}
