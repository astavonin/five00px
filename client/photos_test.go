// Package five00px provides ...
package five00px

import (
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"testing"
)

func handleComments(id string) (string, int) {
	if id == "142795351" {
		return "comments.json", http.StatusOK
	} else if id == "42" {
		return "404.json", http.StatusNotFound
	}
	return "", http.StatusNotFound
}

func handleVote(id, method string) (string, int) {
	if id == "101" && method == http.MethodPost {
		return "vote.json", http.StatusOK
	} else if id == "101" && method == http.MethodDelete {
		return "vote_del.json", http.StatusOK
	} else if id == "42" && method == http.MethodPost {
		return "404.json", http.StatusNotFound
	} else if id == "100" && method == http.MethodPost {
		return "403.json", http.StatusForbidden
	}
	return "", http.StatusInternalServerError
}

func handleVotes(id string) (string, int) {
	if id == "142795351" {
		return "photo_votes.json", http.StatusOK
	} else if id == "42" {
		return "404.json", http.StatusNotFound
	} else if id == "100" {
		return "403.json", http.StatusForbidden
	}
	return "", http.StatusNotFound
}

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

	rePhotoById := regexp.MustCompile(`/photos/(\w+)$`)
	rePhotoVote := regexp.MustCompile(`/photos/(\w+)/vote$`)
	rePhotoVotes := regexp.MustCompile(`/photos/(\w+)/votes`)
	rePhotoComments := regexp.MustCompile(`/photos/(\w+)/comments`)

	if u.Path == "/photos" && r.Method == http.MethodGet {
		return "photos.json", http.StatusOK
	} else if u.Path == "/photos" && r.Method == http.MethodPost {
		return "upload.json", http.StatusOK
	} else if u.Path == "/photos/search" {
		return "photos_search.json", http.StatusOK
	} else if res := rePhotoComments.FindStringSubmatch(u.Path); len(res) > 0 {
		return handleComments(res[1])
	} else if res := rePhotoVote.FindStringSubmatch(u.Path); len(res) > 0 {
		return handleVote(res[1], r.Method)
	} else if res := rePhotoVotes.FindStringSubmatch(u.Path); len(res) > 0 {
		return handleVotes(res[1])
	} else if res := rePhotoById.FindStringSubmatch(u.Path); len(res) > 0 {
		return handleById(res[1])
	}
	return "404.json", http.StatusNotFound
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
		t.Errorf("Photo: %t, Comments: %t", p == nil, p.Comments == nil)
	}

	if p.ID != 142795351 || p.UserID != 8264807 {
		t.Error("Unexpected Photo")
	}
	if len(p.Comments) != 2 || p.CommentsCount != 2 || p.Comments[0].ID != 274841309 {
		t.Error("Unexpected Comments")
	}
}

func TestVotes(t *testing.T) {
	f00 := NewTest500px()

	_, err := f00.PhotoVotes(42, nil)
	if err != ErrPhotoNotFound {
		t.Errorf("Expecting \"%s\" but found \"%s\"", ErrPhotoNotFound, err)
	}

	_, err = f00.PhotoVotes(100, nil)
	if err != ErrPhotoNotAvailable {
		t.Errorf("Expecting \"%s\" but found \"%s\"", ErrPhotoNotAvailable, err)
	}

	v, err := f00.PhotoVotes(142795351, nil)
	if err != nil {
		t.Fatal(err)
	}
	if v == nil || v.Users == nil {
		t.Errorf("Votes: %t, Users: %t", v == nil, v.Users == nil)
	}

	if v.TotalItems != 12 || len(v.Users) != 12 {
		t.Error("Unexpected votes")
	}
}

func TestComments(t *testing.T) {
	f00 := NewTest500px()

	_, err := f00.PhotoComments(42, nil)
	if err != ErrPhotoNotFound {
		t.Errorf("Expecting \"%s\" but found \"%s\"", ErrPhotoNotFound, err)
	}

	c, err := f00.PhotoComments(142795351, nil)
	if err != nil {
		t.Fatal(err)
	}
	if c == nil || c.Comments == nil {
		t.Errorf("Votes: %t, Users: %t", c == nil, c.Comments == nil)
	}

	if c.TotalItems != 2 || len(c.Comments) != 2 {
		t.Error("Unexpected votes")
	}
}

func TestVote(t *testing.T) {
	f00 := NewTest500px()

	err := f00.Vote(42, true)
	if err != ErrPhotoNotFound {
		t.Errorf("Expecting \"%s\" but found \"%s\"", ErrPhotoNotFound, err)
	}

	err = f00.Vote(100, true)
	if err != ErrVoteRejected {
		t.Errorf("Expecting \"%s\" but found \"%s\"", ErrVoteRejected, err)
	}

	err = f00.Vote(101, true)
	if err != nil {
		t.Error(err)
	}
	err = f00.Vote(101, false)
	if err != nil {
		t.Error(err)
	}
}

func TestUpload(t *testing.T) {
	f00 := NewTest500px()
	info := UploadInfo{}

	if info.Valid() {
		t.Error("Expecting invalid")
	}

	info.Name = "test name"
	info.Description = "test description"
	info.Category = CategoryBW
	info.Photo = strings.NewReader("")
	info.Tags = []string{"tag1", "tag2", "tag N"}
	if !info.Valid() {
		t.Error("Should be valid here")
	}

	err := f00.Upload(info)
	if err != nil {
		t.Fatal(err)
	}

}
