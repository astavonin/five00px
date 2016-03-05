// Package five00px provides ...
package five00px

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
)

//func doPhotosReq(c *http.Client, mathod, context string, vals url.Values) (*Photos, error) {
//return nil, ErrInvalidInput
//}

// Photos call returns a list of photos for specified phot stream
func (f00 *Five00px) Photos(c StreamCriterias, p *Page) (*Photos, error) {
	log := logrus.WithFields(logrus.Fields{
		"context":   "Photos",
		"criterias": c,
		"page":      p,
	})

	if !c.Valid() {
		log.Error(ErrInvalidInput)
		return nil, ErrInvalidInput
	}

	vals := c.Vals()
	for k, v := range p.Vals() {
		vals[k] = v
	}
	b, err := doCommand(f00.c, "photos", http.MethodGet, vals)
	if err != nil {
		log.WithError(err).Error("Failed to get data")
		return nil, ErrInternal
	}

	var photos Photos

	err = json.Unmarshal(b, &photos)
	log.WithError(err).Info("Done")
	if err != nil {
		fmt.Println(string(b))
	}
	return &photos, err
}

// PhotosSearch searches for specific photos
func (f00 *Five00px) PhotosSearch(c SearchCriterias, p *Page) (*Photos, error) {
	log := logrus.WithFields(logrus.Fields{
		"context":   "PhotosSearch",
		"criterias": c,
		"page":      p,
	})

	if !c.Valid() {
		log.Error(ErrInvalidInput)
		return nil, ErrInvalidInput
	}

	vals := c.Vals()
	for k, v := range p.Vals() {
		vals[k] = v
	}
	b, err := doCommand(f00.c, "photos/search", http.MethodGet, vals)
	if err != nil {
		log.WithError(err).Error("Failed to get data")
		return nil, ErrInternal
	}

	var photos Photos

	err = json.Unmarshal(b, &photos)
	log.WithError(err).Info("Done")
	if err != nil {
		fmt.Println(string(b))
	}
	return &photos, err
}
