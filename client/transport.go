// Package five00px provides ...
package five00px

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/Sirupsen/logrus"
)

type httpError struct {
	ErrorCode int
	Response  string
}

func (h httpError) Error() string {
	return fmt.Sprintf("HTTP error %d. Message: %s", h.ErrorCode, h.Response)
}

func doGet(c *http.Client, dstPoint string, vals *url.Values) ([]byte, error) {
	if vals != nil {
		dstPoint += "?" + vals.Encode()
	}
	log := logrus.WithFields(logrus.Fields{
		"context": "HTTP GET",
		"host":    mainAPIUrl,
		"path":    dstPoint,
	})

	log.Info("Getting data")

	r, err := c.Get(mainAPIUrl + dstPoint)
	defer func() {
		_ = r.Body.Close()
	}()
	if err != nil {
		log.WithError(err)
		return nil, err
	} else if r.StatusCode != 200 {
		b, err := ioutil.ReadAll(r.Body)
		d := ""
		if err == nil {
			d = string(b)
		}

		log.WithFields(logrus.Fields{
			"StatusCode": r.StatusCode,
			"data":       d,
		}).Warn("Server returns error")
		return nil, httpError{r.StatusCode, d}
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithError(err)
		return nil, err
	}
	log.Info("Succeed")
	return b, nil
}
