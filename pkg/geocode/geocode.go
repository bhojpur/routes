package geocode

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
	"net/http"
	"strconv"
	"time"

	engine "github.com/bhojpur/routes/pkg/engine"
)

const (
	TivraBaseURL        = "https://tivra.bhojpur.net"
	streetDataEndpoint  = "/street_data/"
	geocoderEndpoint    = "/v1/geocoder"
	bulkGeocodeEndpoint = "/v1/actions/upload/json-geocode"
)

type Service struct {
	Client *engine.Client
}

type geocodeRequest struct {
	Addresses string `http:"addresses" json:"-"`
	Format    string `http:"format" json:"-"`
}

func (s *Service) ForwardAddress(address string) ([]Geocode, error) {
	geo := []Geocode{}
	return geo, s.Client.Do(http.MethodPost, geocoderEndpoint, &geocodeRequest{Addresses: address, Format: "json"}, &geo)
}

type forwardBulkRequest struct {
	Rows []Row `json:"rows"`
}

func (s *Service) ForwardBulk(addresses []Row) (*BulkResponse, error) {
	resp := &BulkResponse{}
	return resp, s.Client.Do(http.MethodPost, bulkGeocodeEndpoint, &forwardBulkRequest{Rows: addresses}, resp)
}

func (s *Service) ReverseAddress(latitude float64, longitude float64) ([]Geocode, error) {
	geo := []Geocode{}
	return geo, s.Client.Do(http.MethodPost, geocoderEndpoint, &geocodeRequest{Addresses: fmt.Sprintf("%f,%f", latitude, longitude), Format: "json"}, &geo)
}

type TivraService struct {
	client *engine.Client
}

func NewTivraService(APIKey string) *TivraService {
	return &TivraService{
		client: engine.NewClientWithOptions(APIKey, engine.DefaultTimeout, TivraBaseURL),
	}
}

func NewTivraServiceWithOptions(APIKey string, timeout time.Duration, baseURL string) *TivraService {
	return &TivraService{
		client: engine.NewClientWithOptions(APIKey, timeout, baseURL),
	}
}

func (s *TivraService) GetAddresses() ([]Address, error) {
	response := []Address{}
	return response, s.client.Do(http.MethodGet, streetDataEndpoint, nil, &response)
}

func (s *TivraService) GetSingleAddress(pk int) (*Address, error) {
	response := &Address{}
	return response, s.client.Do(http.MethodGet, streetDataEndpoint+strconv.Itoa(pk), nil, response)
}

func (s *TivraService) GetLimitedAddresses(limit int, offset int) ([]Address, error) {
	response := []Address{}
	return response, s.client.Do(http.MethodGet, fmt.Sprintf("%s%d/%d/", streetDataEndpoint, offset, limit), nil, &response)
}

func (s *TivraService) GetAddressesByPincode(pincode string) ([]Address, error) {
	response := []Address{}
	return response, s.client.Do(http.MethodGet, streetDataEndpoint+"pincode/"+pincode, nil, &response)
}

func (s *TivraService) GetLimitedAddressesByPincode(pincode string, limit int, offset int) ([]Address, error) {
	response := []Address{}
	return response, s.client.Do(http.MethodGet, fmt.Sprintf("%s%s/%s/%d/%d/", streetDataEndpoint, "pincode", pincode, offset, limit), nil, &response)
}

func (s *TivraService) GetAddressesByPincodeAndHousenumber(pincode string, housenumber int) ([]Address, error) {
	response := []Address{}
	return response, s.client.Do(http.MethodGet, fmt.Sprintf("%s%s/%s/%d/", streetDataEndpoint, "service", pincode, housenumber), nil, &response)
}

func (s *TivraService) GetLimitedAddressesByPincodeAndHousenumber(pincode string, housenumber int, limit int, offset int) ([]Address, error) {
	response := []Address{}
	return response, s.client.Do(http.MethodGet, fmt.Sprintf("%s%s/%s/%d/%d/%d/", streetDataEndpoint, "service", pincode, housenumber, offset, limit), nil, &response)
}
