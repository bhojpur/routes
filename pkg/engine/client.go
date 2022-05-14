package engine

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/bhojpur/routes/pkg/utils"
)

const (
	DefaultTimeout time.Duration = time.Minute * 30
	BaseURL                      = "https://route.bhojpur.net"
)

var InvalidStatusCode = errors.New("Invalid status code")

type Client struct {
	APIKey  string
	Client  *http.Client
	BaseURL string
}

func NewClientWithOptions(APIKey string, timeout time.Duration, baseURL string) *Client {
	return &Client{
		APIKey:  APIKey,
		Client:  &http.Client{Timeout: timeout},
		BaseURL: baseURL,
	}
}

// NewClient creates a Bhojpur Routes client
func NewClient(APIKey string) *Client {
	return NewClientWithOptions(APIKey, DefaultTimeout, BaseURL)
}

func (c *Client) constructBody(data interface{}) (contentType string, reader bytes.Buffer, err error) {
	//Check if the data struct has any postform data to pass to the body
	params := utils.StructToURLValues("form", data)
	// if there are no form parameters, it's likely the request is a json
	if len(params) == 0 {
		if err = json.NewEncoder(&reader).Encode(data); err != nil {
			return
		}
		contentType = "application/json"
		return
	}

	//otherwise we encode the form as a multipart form
	w := multipart.NewWriter(&reader)
	defer w.Close()
	for key, vals := range params {
		for _, v := range vals {
			err = w.WriteField(key, v)
			if err != nil {
				return
			}
		}
	}
	contentType = w.FormDataContentType()
	return
}

func (c *Client) DoNoDecode(method string, endpoint string, data interface{}) (response []byte, err error) {
	var requestBody bytes.Buffer
	var contentType string
	//We might change this to == for better accuracy
	if method != http.MethodGet && method != http.MethodOptions {
		if contentType, requestBody, err = c.constructBody(data); err != nil {
			return response, err
		}
	}

	request, err := http.NewRequest(method, c.BaseURL+endpoint, &requestBody)
	if err != nil {
		return response, err
	}

	request.Header.Set("Content-Type", contentType)
	params := url.Values{}

	if data != nil {
		//Prepare query string
		params = utils.StructToURLValues("http", data)
	}
	params.Add("api_key", c.APIKey)
	params.Add("format", "json")
	request.URL.RawQuery = params.Encode()

	// b, err := httputil.DumpRequestOut(request, true)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(b))

	resp, err := c.Client.Do(request)

	// b, err = httputil.DumpResponse(resp, true)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(b))
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	response, err = ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return response, InvalidStatusCode
	}
	return response, err
}

func (c *Client) Do(method string, endpoint string, data interface{}, out interface{}) error {
	read, err := c.DoNoDecode(method, endpoint, data)
	//Error handling is a bit weird
	if err != nil {
		return c.parseErrors(read, err)
	}
	if out == nil {
		return err
	}
	err = json.Unmarshal(read, out)
	return err
}

func (c *Client) parseErrors(data []byte, err error) error {
	//Check if invalid status code - errors:[] response is returned only when statuscode is not 200
	if err == InvalidStatusCode {
		errs := &ErrorResponse{}
		//Try to parse to ErrorResponse
		unmerr := json.Unmarshal(data, errs)
		//Sometimes (status code: 500,404) errors:[] might not be returned, we return the err from the request when it happens
		if unmerr != nil || len(errs.Errors) == 0 {
			return err
		}
		//Join all errors in the ErrorResponse
		return errors.New(strings.Join(errs.Errors, ","))
	}
	return err
}
