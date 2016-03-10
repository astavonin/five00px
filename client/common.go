// Package five00px provides ...
package five00px

import (
	"net/url"
	"strconv"
)

var uploadAPIUrl = "http://upload.500px.com/v1/upload"
var mainAPIUrl = "https://api.500px.com/v1/"

// ValsConverter converts type to url.Values
type ValsConverter interface {
	Vals() url.Values
}

// Page structure stores RPP(Results Per Page) and Page(Return the specified
// page of the resource) values.
type Page struct {
	Rpp  int
	Page int
}

// NewPage call returns Page with 500px default values `Results Per Page` 20
// and `Page` 1
func NewPage() Page {
	return Page{20, 1}
}

// Vals converts Page to url.Values
func (p *Page) Vals() url.Values {
	vals := url.Values{}
	if p != nil {
		vals.Add("page", strconv.Itoa(p.Page))
		vals.Add("rpp", strconv.Itoa(p.Rpp))
	}
	return vals
}

func urlEncode(str string) string {
	u, err := url.Parse(str)
	if err != nil {
		return str
	}
	return u.String()
}
