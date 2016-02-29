// Package five00px provides ...
package five00px

import (
	"net/url"
	"strconv"
)

func pageToVals(page *Page) url.Values {
	vals := url.Values{}
	if page != nil {
		vals.Add("page", strconv.Itoa(page.Page))
		vals.Add("rpp", strconv.Itoa(page.Rpp))
	}
	return vals
}
