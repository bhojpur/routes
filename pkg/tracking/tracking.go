package tracking

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
	"github.com/bhojpur/routes/pkg/routing"
)

const (
	setEndpoint            = "/v1/track/set"
	lastLocationEndpoint   = "/v1/route"
	deviceLocationEndpoint = "/v1/track/get_device_location"
	statusEndpoint         = "/v1/status"
)

type Service struct {
	Client *engine.Client
}

func (s *Service) SetGPS(data *GPS) (string, error) {
	byt, err := s.Client.DoNoDecode(http.MethodGet, setEndpoint, data)
	return string(byt), err
}

type getLastLocationRequest struct {
	RouteID               string `http:"route_id"`
	DeviceTrackingHistory int    `http:"device_tracking_history"`
}

func (s *Service) GetLastLocation(routeID string) (*routing.DataObject, error) {
	request := &getLastLocationRequest{RouteID: routeID, DeviceTrackingHistory: 1}
	resp := &routing.DataObject{}
	return resp, s.Client.Do(http.MethodGet, lastLocationEndpoint, request, resp)
}

type getDeviceLocationHistoryResponse struct {
	Data []routing.TrackingHistory `json:"data"`
}

func (s *Service) GetDeviceLocationHistory(query *TrackingHistoryQuery) ([]routing.TrackingHistory, error) {
	resp := &getDeviceLocationHistoryResponse{}
	return resp.Data, s.Client.Do(http.MethodGet, deviceLocationEndpoint, query, resp)
}

func (s *Service) TrackAssets(tracking string) (*AssetTracking, error) {
	resp := &AssetTracking{}
	return resp, s.Client.Do(http.MethodGet, statusEndpoint, &struct {
		Tracking string `http:"tracking"`
	}{tracking}, resp)
}
