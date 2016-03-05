// Package five00px provides ...
package five00px

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func photosHandler(w http.ResponseWriter, r *http.Request) (string, int) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		return "", http.StatusInternalServerError
	}

	if u.Path == "/photos" {
		return "photos.json", http.StatusOK
	} else if u.Path == "/photos/search" {
		return "photos_search.json", http.StatusOK
	}
	return "", http.StatusInternalServerError
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
