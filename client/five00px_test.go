// Package five00px provides ...
package five00px

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"strings"
	"testing"
)

// TestHandlerFunc returns file name  with responce content
type TestHandlerFunc func(http.ResponseWriter, *http.Request) (string, int)

const testDataDir string = "test_data"

var (
	TestURL  = ""
	handlers map[string]TestHandlerFunc
)

func registerHandlers() {
	handlers = make(map[string]TestHandlerFunc)

	handlers["users"] = userHandler
	handlers["photos"] = photosHandler
}

func handler(w http.ResponseWriter, r *http.Request) {
	u, _ := url.Parse(r.URL.String())
	hName := strings.Split(u.Path, "/")[1]
	h, ok := handlers[hName]
	if !ok {
		http.Error(w, fmt.Sprintf("Cannot find \"%s\" handler", hName),
			http.StatusInternalServerError)
		return
	}

	fName, status := h(w, r)
	if fName == "" {
		http.NotFound(w, r)
		return
	}
	b, err := ioutil.ReadFile(path.Join(testDataDir, fName))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if status != 200 {
		http.Error(w, string(b), status)
		return
	}
	_, _ = w.Write(b)
}

func NewTest500px() Five00px {
	c := http.Client{}
	return Five00px{c: &c}
}

func TestMain(m *testing.M) {
	registerHandlers()
	ts := httptest.NewServer(http.HandlerFunc(handler))
	mainAPIUrl = ts.URL + "/"
	defer ts.Close()

	flag.Parse()
	os.Exit(m.Run())
}
