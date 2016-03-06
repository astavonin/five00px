// Package five00px provides ...
package five00px

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"testing"
)

func handleById(id string) (string, int) {
	if id == "142795351" {
		return "photo.json", http.StatusOK
	} else if id == "42" {
		return "404.json", http.StatusNotFound
	} else if id == "100" {
		return "403.json", http.StatusForbidden
	}
	return "", http.StatusNotFound
}

func photosHandler(w http.ResponseWriter, r *http.Request) (string, int) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		return "", http.StatusInternalServerError
	}

	rePhotoById := regexp.MustCompile(`/photos/(\w+)`)

	if u.Path == "/photos" {
		return "photos.json", http.StatusOK
	} else if u.Path == "/photos/search" {
		return "photos_search.json", http.StatusOK
	} else if res := rePhotoById.FindStringSubmatch(u.Path); len(res) > 0 {
		return handleById(res[1])
	}
	return "", http.StatusNotFound
}

func TestCategory(t *testing.T) {
	cats := Categories{1, 2, 42}
	if cats.Valid() {
		t.Error("There is no category 42")
	}

	cats = Categories{CategoryTravel, CategoryBW, CategoryArch}
	cExp := "Travel,Black and White,City and Architecture"
	cRes := cats.String()
	if cExp != cRes {
		t.Errorf("Expecting \"%s\" found \"%s\"", cExp, cRes)
	}
}

func TestPhotos(t *testing.T) {
	// photos?feature=fresh_today&only=Abstract&sort=created_at&rpp=3&image_size=4&include_store=store_download&include_states=voted&tags=1
	f00 := NewTest500px()

	// Valid request
	s := StreamCriterias{
		Feature: FeaturePopular,
	}

	f, err := f00.Photos(s, nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(f.Photos) != 3 || f.CurrentPage != 1 || f.TotalPages != 282 || f.TotalItems != 844 {
		t.Error("Unexpected data")
	}

	fmt.Println("Category:", f.Photos[0].Category)
	// Unexpected feature
	s.Feature = "asdfg"
	f, err = f00.Photos(s, nil)
	if err != ErrInvalidInput {
		t.Errorf("Expecting \"%s\" but found \"%s\"", ErrInvalidInput, err)
	}
}

func TestPhotosSearch(t *testing.T) {
	f00 := NewTest500px()
	c := SearchCriterias{
		Term:        "test",
		Tag:         "best",
		LicenseType: LicAll,
	}

	cRes := buildQuery(c.Vals())
	cExp := "?tag=best&tags=1&term=test"
	if cRes != cExp {
		t.Errorf("Expecting:\t%s\nFound:\t%s", cExp, cRes)
	}

	p, err := f00.PhotosSearch(c, nil)
	if err != nil {
		t.Fatal(err)
	}

	if p.CurrentPage != 1 || p.TotalItems != 84 || p.TotalPages != 28 ||
		len(p.Photos) != 3 {
		t.Error("Unexpected data")
	}
}

func TestPhotoById(t *testing.T) {
	f00 := NewTest500px()

	_, err := f00.PhotoById(42, nil)
	if err != ErrPhotoNotFound {
		t.Errorf("Expecting \"%s\" but found \"%s\"", ErrPhotoNotFound, err)
	}

	_, err = f00.PhotoById(100, nil)
	if err != ErrPhotoNotAvailable {
		t.Errorf("Expecting \"%s\" but found \"%s\"", ErrPhotoNotAvailable, err)
	}

	info := PhotoInfo{Comments: true}
	p, err := f00.PhotoById(142795351, &info)
	if err != nil {
		t.Fatal(err)
	}
	if p == nil || p.Comments == nil {
		t.Error("Photo: %t, Comments: %t", p == nil, p.Comments == nil)
	}

	fmt.Println(p.ID, p.UserID)
	if p.ID != 142795351 || p.UserID != 8264807 {
		t.Error("Unexpected Photo")
	}
	if len(p.Comments) != 2 || p.CommentsCount != 2 || p.Comments[0].ID != 274841309 {
		t.Error("Unexpected Comments")
	}
}
