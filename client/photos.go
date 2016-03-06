// Package five00px provides ...
package five00px

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	return &photos, err
}

func (f00 *Five00px) PhotoById(id int, info *PhotoInfo) (*Photo, error) {
	log := logrus.WithFields(logrus.Fields{
		"context": "PhotoById",
		"id":      id,
		"info":    info,
	})

	vals := info.Vals()
	b, err := doCommand(f00.c, "photos/"+strconv.Itoa(id), http.MethodGet, vals)
	if err != nil {
		return nil, processError(log, b)
	}

	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(b, &objmap)

	if err != nil {
		log.WithError(err).WithField("data", string(b)).
			Error("Unable to unmarshall data")
		return nil, err
	}

	var photo Photo

	err = json.Unmarshal(*objmap["photo"], &photo)
	log.WithError(err).Info("Done")
	return &photo, err
}

func processError(log *logrus.Entry, b []byte) error {
	var e00 five00Error
	err := json.Unmarshal(b, &e00)
	if err != nil {
		log.WithError(err).WithField("data", string(b)).
			Error("Unable to unmarshall data")
		return ErrInternal
	}
	log.WithField("status", strconv.Itoa(e00.Status)).
		Info("server returns error")
	switch e00.Status {
	case http.StatusNotFound:
		return ErrPhotoNotFound
	case http.StatusForbidden:
		return ErrPhotoNotAvailable
	}
	return nil
}

func (f00 *Five00px) PhotoVotes(id int, p *Page) (*Votes, error) {
	log := logrus.WithFields(logrus.Fields{
		"context": "PhotoVotes",
		"id":      id,
		"page":    p,
	})

	b, err := doCommand(f00.c, "photos/"+strconv.Itoa(id)+"/votes", http.MethodGet, p.Vals())
	if err != nil {
		return nil, processError(log, b)
	}

	var votes Votes

	err = json.Unmarshal(b, &votes)
	log.WithError(err).Info("Done")
	return &votes, err
}
