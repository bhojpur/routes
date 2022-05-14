package telematics

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
	"time"

	engine "github.com/bhojpur/routes/pkg/engine"
)

const (
	TelematicsBaseURL = "https://telematics.bhojpur.net"
	vendorsEndpoint   = "/v1/vendors"
)

type TelematicsService struct {
	client *engine.Client
}

func NewTelematicsService(APIKey string) *TelematicsService {
	return &TelematicsService{
		client: engine.NewClientWithOptions(APIKey, engine.DefaultTimeout, TelematicsBaseURL),
	}
}

func NewTelematicsServiceWithOptions(APIKey string, timeout time.Duration, baseURL string) *TelematicsService {
	return &TelematicsService{
		client: engine.NewClientWithOptions(APIKey, timeout, baseURL),
	}
}

type getVendorsResponse struct {
	Vendors []Vendor `json:"vendors"`
}

func (s *TelematicsService) GetVendors() ([]Vendor, error) {
	response := &getVendorsResponse{}
	return response.Vendors, s.client.Do(http.MethodGet, vendorsEndpoint, nil, &response)
}

type getVendorRequest struct {
	VendorID int `http:"vendor_id"`
}

type getVendorResponse struct {
	Vendor *Vendor `json:"vendor"`
}

func (s *TelematicsService) GetVendor(vendor int) (*Vendor, error) {
	response := &getVendorResponse{}
	return response.Vendor, s.client.Do(http.MethodGet, vendorsEndpoint, &getVendorRequest{VendorID: vendor}, response)
}

func (s *TelematicsService) SearchVendors(query *VendorQuery) ([]Vendor, error) {
	response := &getVendorsResponse{}
	return response.Vendors, s.client.Do(http.MethodGet, vendorsEndpoint, nil, &response)
}

type compareVendorsRequest struct {
	Vendors []int `json:"vendors`
}

func (s *TelematicsService) CompareVendors(vendors ...int) ([]Vendor, error) {
	response := &getVendorsResponse{}
	return response.Vendors, s.client.Do(http.MethodGet, vendorsEndpoint, &compareVendorsRequest{Vendors: vendors}, response)
}
