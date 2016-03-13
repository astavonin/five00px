// Package main provides ...
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
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

func new500px(c *Config, cnfgPath string) *five00px.Five00px {

	f00 := five00px.New(c.ConsumerKey, c.ConsumerSecret, logrus.New())

	if c.Token == nil {
		t, err := f00.Auth()
		if err != nil {
			log.Fatal(err)
		}
		c.Token = t
		storeConfig(cnfgPath, c)
	} else {
		f00.Restore(c.Token)
	}
	return &f00
}

func getPhotos(f00 *five00px.Five00px) *five00px.Photos {

	stream := five00px.StreamCriterias{
		Feature: five00px.FeaturePopular,
		Only: five00px.Categories{
			five00px.CategoryNude,
			five00px.CategoryFashion,
			five00px.CategoryPeople,
		},
	}
	photos, err := f00.ListPhotos(stream, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Top", len(photos.Photos), "photos:")
	for _, photo := range photos.Photos {
		author := &photo.User
		fmt.Printf("(%f/%d) %s by %s %s\n", photo.Rating, photo.ID, photo.Name,
			author.Firstname, author.Lastname)
	}
	return photos
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
	f00 := new500px(c, *configPath)

	curUser, err := f00.GetUserByID(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current user is (%d) %s %s\n", curUser.ID, curUser.Firstname,
		curUser.Lastname)

	// Lets select top 20 photos in Nude, Fashion and People categories
	photos := getPhotos(f00)

	// we need 3 photos from TOP20
	r := rand.NewSource(time.Now().Unix())

	// select only photos we didn't like before
	idxs := []int{}
	s := len(photos.Photos)
	for i, j := 0, 0; i < 1 && j < s; j++ {
		photo := photos.Photos[int(r.Int63())%s]
		if !photo.Voted {
			idxs = append(idxs, photo.ID)
			i++
		}
	}

	// we will add likes and comments to selected photos
	for _, idx := range idxs {
		//err = f00.AddVote(idx, true)
		//if err != nil {
		//fmt.Printf("Cannot add vote to %s, error: %s\n", idx, err)
		//}
		err = f00.AddComment(idx, "amazing work")
		if err != nil {
			fmt.Printf("Cannot add comment to %d, error: %s\n", idx, err)
		}
	}

	// check that likes and comments are in place and download images
	phInfo := five00px.PhotoInfo{
		ImageSize: five00px.Size900l,
		Comments:  true,
	}
	for _, idx := range idxs {
		photo, err := f00.GetPhotoByID(idx, &phInfo)
		if err != nil {
			fmt.Printf("Cannot get photo with ID=%derror: %s\n", idx, err)
			continue
		}
		fmt.Println("We liked the photo", photo.ID, photo.Liked)
		for _, comment := range photo.Comments {
			if comment.UserID == curUser.ID {
				fmt.Printf("And this is our comment \"%s\"\n", comment.Body)
			}
		}
		//resp, err := http.Get(photo.ImageURL)
		//defer resp.Body.Close()
		//if err != nil {
		//fmt.Println(err)
		//continue
		//}
		//file, err := os.Create(strconv.Itoa(idx) + ".jpg")
		//defer file.Close()
		//if err != nil {
		//fmt.Println(err)
		//continue
		//}
		//io.Copy(file, resp.Body)
	}

	//photo, err := f00.GetPhotoByID(photoID, nil)

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

	//fr, err := f00.ListFriends(9091479, &five00px.Page{1, 1})

	//for _, u := range fr.Users {
	//fmt.Println(u.Avatars.Default)
	//}
	//u, err := f00.GetUserByID(9091479)
	//if err != nil {
	//log.Fatal(err)
	//}
	//fmt.Println(u.City)

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
