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
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const apiKey = "11111111111111111111111111111111"

func TestNewClient(t *testing.T) {
	cli := NewClient(apiKey)
	if cli.Client == nil {
		t.Error("Client has not been initialized.")
	}
	if cli.APIKey != apiKey {
		t.Error("APIKey has not been assigned.")
	}
}

func testClient(code int, body string) (*httptest.Server, *Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		read, err := ioutil.ReadAll(r.Body)
		if err != nil || len(read) == 0 {
			fmt.Fprint(w, body)
		} else {
			w.Write(read)
		}
	}))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	httpClient := &http.Client{Transport: transport}
	client := &Client{Client: httpClient, APIKey: apiKey, BaseURL: "http://bhojpur.net"}
	return server, client
}

type response struct {
	Number int `json:"json"`
}

func TestDecodingGet(t *testing.T) {
	server, cli := testClient(200, `{"json":42}`)
	defer server.Close()
	resp := &response{}
	err := cli.Do(http.MethodGet, "/whatever/", &struct{}{}, resp)
	if err != nil {
		t.Error(err)
	}
	if resp.Number != 42 {
		t.Error("Unmarshalling went wrong")
	}
}

func TestDecodingPost(t *testing.T) {
	server, cli := testClient(200, `{"json":42}`)
	defer server.Close()
	resp := &response{}
	err := cli.Do(http.MethodPost, "/whatever/", &struct {
		Number int `json:"json"`
	}{Number: 152}, resp)
	if err != nil {
		t.Error(err)
	}
	if resp.Number != 152 {
		t.Error("Error occured during unmarshalling")
	}
}

func TestDecodingErrors(t *testing.T) {
	server, cli := testClient(500, `{"errors":["error#1","error#2"]}`)
	defer server.Close()
	resp := &response{}
	err := cli.Do(http.MethodGet, "/", &struct{}{}, resp)
	if err == nil || err.Error() != "error#1,error#2" {
		t.Error("Expecting error 'error#1,error#2', got: ", err.Error())
	}
}
