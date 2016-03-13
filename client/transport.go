// Package five00px provides ...
package five00px

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"sort"
	"strings"

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

func doUpload(c *http.Client, dstPoint, fName string, f io.Reader, vals url.Values) ([]byte, error) {

	dstPoint += buildQuery(vals)
	log := logger.WithFields(logrus.Fields{
		"context": "doUpload",
		"path":    dstPoint,
	})

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fw, err := w.CreateFormFile("file", fName)
	if err != nil {
		log.WithError(err).Error("Cannot create Writer")
		return nil, err
	}
	if _, err = io.Copy(fw, f); err != nil {
		log.WithError(err).Error("Cannot read file")
		return nil, err
	}

	_ = w.Close()

	return do(c, mainAPIUrl+dstPoint, http.MethodPost, w.FormDataContentType(), &b, log)
}

func doCommand(c *http.Client, dstPoint, method string, vals url.Values, body url.Values) ([]byte, error) {
	dstPoint += buildQuery(vals)

	log := logger.WithFields(logrus.Fields{
		"context": "HTTP " + method,
		"path":    dstPoint,
		"values":  vals,
	})

	var (
		b     io.Reader = nil
		ctype           = ""
	)
	if body != nil {
		ctype = "application/x-www-form-urlencoded"
		b = strings.NewReader(body.Encode())
	}

	return do(c, mainAPIUrl+dstPoint, method, ctype, b, log)
}

func do(c *http.Client, url, method, ctype string, body io.Reader, log *logrus.Entry) ([]byte, error) {

	log = log.WithField("host", mainAPIUrl)
	log.Info("Initiating request")

	req, err := http.NewRequest(method, url, body)
	if ctype != "" {
		req.Header.Set("Content-Type", ctype)
	}
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
