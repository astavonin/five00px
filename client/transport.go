// Package five00px provides ...
package five00px

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"

	"github.com/Sirupsen/logrus"
)

type httpError struct {
	ErrorCode int
	Response  string
}

func (h httpError) Error() string {
	return fmt.Sprintf("HTTP error %d. Message: %s", h.ErrorCode, h.Response)
}

func buildQuery(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	_, _ = buf.WriteString("?")
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		prefix := k + "="
		for _, v := range vs {
			if buf.Len() > 1 {
				_ = buf.WriteByte('&')
			}
			_, _ = buf.WriteString(urlEncode(prefix))
			_, _ = buf.WriteString(urlEncode(v))
		}
	}
	return buf.String()
}

func doCommand(c *http.Client, dstPoint, method string, vals url.Values) ([]byte, error) {

	dstPoint += buildQuery(vals)

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
