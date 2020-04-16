package external

import (
	"crypto/tls"
	"encoding/json"

	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

type Option struct {
	Timeout   int
	KeepAlive int
	MaxIdle   int
}

type Request struct {
	client *http.Client
}

func (request *Request) Init(opt *Option) error {
	cli := http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				// timeout for tcp connection
				// Timeout:   time.Duration(opt.Timeout) * time.Second,
				KeepAlive: time.Duration(opt.KeepAlive) * time.Second,
			}).Dial,
			MaxIdleConns:        opt.MaxIdle,
			MaxIdleConnsPerHost: opt.MaxIdle / 3,
			IdleConnTimeout:     time.Duration(opt.KeepAlive*3) * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			// ResponseHeaderTimeout: time.Duration(opt.Timeout) * time.Second,
			// TLSHandshakeTimeout:   time.Duration(opt.Timeout) * time.Second,
		},

		// timeout for http response
		// tcp connection+TLSHandshakeTimeout+ResponseHeaderTimeout+BodyTimeout
		Timeout: time.Duration(opt.Timeout) * time.Second,
	}
	request.client = &cli

	return nil
}

func (requst *Request) NewRequest(method string, uri string,
	headers map[string]string, payload []byte) (*http.Request, error) {

	req, err := http.NewRequest(method, uri, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	if headers != nil {
		if _, ok := headers["User-Agent"]; !ok {
			req.Header.Add("User-Agent", "DD01-Agent/0.1")
		}

		for key, val := range headers {
			req.Header.Add(key, val)
		}
	} else {
		req.Header.Add("User-Agent", "DD01-Agent/0.1")
	}

	return req, nil
}

func CheckStatus(r *http.Response) error {
	if code := r.StatusCode; 200 <= code && code <= 299 {
		return nil
	}

	return errors.New(fmt.Sprintf("%v %v: %d",
		r.Request.Method, r.Request.URL, r.StatusCode))
}

// Do sends an API request and returns the API response.
// The API response is JSON decoded and stored in the value pointed by v,
// or returned as an error if an API error has occurred.
// If v implements the io.Writer interface, the raw response body will
// be written to v,
// without attempting to decode it.
func (request *Request) Do(method string, uri string,
	headers map[string]string, payload []byte, v interface{}) error {

	req, err := request.NewRequest(method, uri, headers, payload)
	if err != nil {
		return err
	}

	res, err := request.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = CheckStatus(res)
	if err != nil {
		return err
	}

	//buf := new(bytes.Buffer)
	//buf.ReadFrom(res.Body)
	//s := buf.String()
	//log.Println("Beginning")
	//log.Println(s)
	//log.Println("Ending")

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, res.Body)
		} else {
			err = json.NewDecoder(res.Body).Decode(v)
		}
	}

	return err
}

func (request *Request) Get(uri string, headers map[string]string,
	v interface{}) error {

	return request.Do("GET", uri, headers, nil, v)
}

func (request *Request) Post(uri string, headers map[string]string,
	payload []byte, v interface{}) error {

	return request.Do("POST", uri, headers, payload, v)
}

func (request *Request) Put(uri string, headers map[string]string,
	payload []byte, v interface{}) error {

	return request.Do("PUT", uri, headers, payload, v)
}

func (request *Request) Delete(uri string, headers map[string]string,
	v interface{}) error {

	return request.Do("DELETE", uri, headers, nil, nil)
}
