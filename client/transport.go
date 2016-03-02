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

func doCommand(c *http.Client, dstPoint, method string, vals url.Values) ([]byte, error) {
	if len(vals) > 0 {
		dstPoint += "?" + vals.Encode()
	}
	log := logrus.WithFields(logrus.Fields{
		"context": "HTTP " + method,
		"host":    mainAPIUrl,
		"path":    dstPoint,
		"values":  vals,
	})

	log.Info("Initiating request")

	req, err := http.NewRequest(method, mainAPIUrl+dstPoint, nil)
	r, err := c.Do(req)
	defer func() {
		_ = r.Body.Close()
	}()
	if err != nil {
		log.WithError(err).Error("Request failed")
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
		return b, httpError{r.StatusCode, d}
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithError(err).Error("Reading error")
		return b, err
	}
	log.Info("Done")
	return b, nil
}
