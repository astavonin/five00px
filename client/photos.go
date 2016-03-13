// Package five00px provides ...
package five00px

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Sirupsen/logrus"
)

//func doPhotosReq(c *http.Client, mathod, context string, vals url.Values) (*Photos, error) {
//return nil, ErrInvalidInput
//}

// ListPhotos call returns a list of photos for specified phot stream
func (f00 *Five00px) ListPhotos(c StreamCriterias, p *Page) (*Photos, error) {
	log := logger.WithFields(logrus.Fields{
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

// SearchPhoto searches for specific photos
func (f00 *Five00px) SearchPhoto(c SearchCriterias, p *Page) (*Photos, error) {
	log := logger.WithFields(logrus.Fields{
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

// GetPhotoByID returns photo by its ID. info may contains additional photo criterias
func (f00 *Five00px) GetPhotoByID(id int, info *PhotoInfo) (*Photo, error) {
	log := logger.WithFields(logrus.Fields{
		"context": "PhotoById",
		"id":      id,
		"info":    info,
	})

	vals := info.Vals()
	b, err := doCommand(f00.c, "photos/"+strconv.Itoa(id), http.MethodGet, vals)
	if err != nil {
		return nil, processError(log, b, errorTable{
			http.StatusNotFound:  ErrPhotoNotFound,
			http.StatusForbidden: ErrPhotoNotAvailable,
		})
	}

	photo, err := extractPhoto(b, log)
	log.WithError(err).Info("Done")

	return photo, err
}

type errorTable map[int]error

func processError(log *logrus.Entry, b []byte, errTbl errorTable) error {
	var e00 five00Error
	err := json.Unmarshal(b, &e00)
	if err != nil {
		log.WithError(err).WithField("data", string(b)).
			Error("Unable to unmarshall data")
		return ErrInternal
	}
	if e00.Status == 200 {
		return nil
	}
	log.WithField("status", strconv.Itoa(e00.Status)).
		Info("server returns error")
	ret := errTbl[e00.Status]
	if ret == nil {
		return ErrInternal
	}
	return ret
}

// ListComments call returns cammonts for selected photo.
func (f00 *Five00px) ListComments(id int, p *Page) (*Comments, error) {
	log := logger.WithFields(logrus.Fields{
		"context": "PhotoComments",
		"id":      id,
		"page":    p,
	})

	b, err := doCommand(f00.c, "photos/"+strconv.Itoa(id)+"/comments", http.MethodGet, p.Vals())
	if err != nil {
		return nil, processError(log, b, errorTable{
			http.StatusNotFound:  ErrPhotoNotFound,
			http.StatusForbidden: ErrPhotoNotAvailable,
		})
	}

	var c Comments

	err = json.Unmarshal(b, &c)
	log.WithError(err).Info("Done")
	return &c, err
}

// ListVotes call returns votes fpr selected photo.
func (f00 *Five00px) ListVotes(id int, p *Page) (*Votes, error) {
	log := logger.WithFields(logrus.Fields{
		"context": "PhotoVotes",
		"id":      id,
		"page":    p,
	})

	b, err := doCommand(f00.c, "photos/"+strconv.Itoa(id)+"/votes", http.MethodGet, p.Vals())
	if err != nil {
		return nil, processError(log, b, errorTable{
			http.StatusNotFound:  ErrPhotoNotFound,
			http.StatusForbidden: ErrPhotoNotAvailable,
		})
	}

	var votes Votes

	err = json.Unmarshal(b, &votes)
	log.WithError(err).Info("Done")
	return &votes, err
}

// AddVote adds/removes vote from selected photo
func (f00 *Five00px) AddVote(id int, like bool) error {
	log := logger.WithFields(logrus.Fields{
		"context": "Vote",
		"id":      id,
		"like":    like,
	})

	method := http.MethodPost
	vals := url.Values{}
	if !like {
		method = http.MethodDelete
	} else {
		vals.Add("vote", "1")
	}
	b, err := doCommand(f00.c, "photos/"+strconv.Itoa(id)+"/vote", method, vals)
	if err != nil {
		return processError(log, b, errorTable{
			http.StatusNotFound:   ErrPhotoNotFound,
			http.StatusForbidden:  ErrVoteRejected,
			http.StatusBadRequest: ErrInvalidInput,
		})
	}

	return nil
}

// AddComment call adds new comment to selected photo
func (f00 *Five00px) AddComment(id int, comment string) error {
	log := logger.WithFields(logrus.Fields{
		"context": "AddComment",
		"id":      id,
		"comment": comment,
	})

	vals := url.Values{"body": {comment}}
	b, err := doCommand(f00.c, "photos/"+strconv.Itoa(id)+"/commens", http.MethodPost, vals)
	if err != nil {
		return processError(log, b, errorTable{
			http.StatusNotFound:   ErrPhotoNotFound,
			http.StatusBadRequest: ErrBadComment,
		})
	}

	return nil
}

func extractPhoto(buf []byte, log *logrus.Entry) (*Photo, error) {
	var objmap map[string]*json.RawMessage
	err := json.Unmarshal(buf, &objmap)

	if err != nil {
		log.WithError(err).WithField("data", string(buf)).
			Error("Unable to unmarshall data")
		return nil, err
	}

	var photo Photo

	err = json.Unmarshal(*objmap["photo"], &photo)

	return &photo, err
}

// AddPhoto call uploads new photo. Information about uploaded photo or
// upload error will be returned
func (f00 *Five00px) AddPhoto(info UploadInfo) (*Photo, error) {
	log := logger.WithFields(logrus.Fields{
		"context": "Upload",
		"info":    info,
	})

	if !info.Valid() {
		log.Error("Invalid input")
		return nil, ErrInvalidInput
	}
	b, err := doUpload(f00.c, "photos/upload", info.FileName, info.PhotoStream, info.Vals())
	if err != nil {
		return nil, processError(log, b, errorTable{
			422: ErrUnprocessableEntity,
		})
	}

	photo, err := extractPhoto(b, log)
	log.WithError(err).Info("Done")

	return photo, err
}

// DelPhoto removes photo by id
func (f00 *Five00px) DelPhoto(id int) error {
	log := logger.WithFields(logrus.Fields{
		"context": "PhotoDelete",
		"id":      id,
	})

	b, err := doCommand(f00.c, "photos/"+strconv.Itoa(id),
		http.MethodDelete, nil)

	if err != nil {
		return processError(log, b, errorTable{
			http.StatusNotFound:  ErrPhotoNotFound,
			http.StatusForbidden: ErrPhotoNotAvailable,
		})
	}

	return nil
}
