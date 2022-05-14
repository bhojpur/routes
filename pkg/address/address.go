package address

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
	"net/http"

	engine "github.com/bhojpur/routes/pkg/engine"
	"github.com/bhojpur/routes/pkg/utils"
)

const endpoint = "/v1/address_book"

type Service struct {
	Client *engine.Client
}

type getResponse struct {
	Results []Contact `json:"results"`
	Total   int       `json:"total"`
}

func (s *Service) Get(query *Query) ([]Contact, int, error) {
	resp := &getResponse{}
	return resp.Results, resp.Total, s.Client.Do(http.MethodGet, endpoint, query, resp)
}

func (s *Service) Add(data *Contact) (*Contact, error) {
	resp := &Contact{}
	err := s.Client.Do(http.MethodPost, endpoint, data, resp)
	return resp, err
}

func (s *Service) Update(data *Contact) (*Contact, error) {
	resp := &Contact{}
	return resp, s.Client.Do(http.MethodPut, endpoint, data, resp)
}

type deleteRequest struct {
	AddressIDs []string `json:"address_ids"`
}

func (s *Service) Delete(ids []string) (bool, error) {
	request := &deleteRequest{AddressIDs: ids}
	resp := &utils.StatusResponse{}
	return resp.Status, s.Client.Do(http.MethodDelete, endpoint, request, resp)
}
