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
	"testing"

	engine "github.com/bhojpur/routes/pkg/engine"
	"github.com/bhojpur/routes/pkg/routing"
)

var client = engine.NewClient("11111111111111111111111111111111")
var service = &Service{Client: client}

func TestIntegrationSetGPS(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	oService := &routing.Service{Client: client}
	opts, err := oService.GetOptimizations(&routing.RouteQuery{
		Limit: 1,
	})
	if err != nil {
		t.Error("Error occured in external service:", err)
		return
	}
	if len(opts) < 1 {
		t.Skip("Not enough routes to test activity stream")
	}
	opt, err := oService.GetOptimization(&routing.OptimizationParameters{
		ProblemID: opts[0].ProblemID,
	})
	if len(opt.Routes) < 1 {
		t.Skip("Not enough routes to test activity stream")
	}
	query := &GPS{
		RouteID:         opt.Routes[0].ID,
		Latitude:        33.14384,
		Longitude:       -83.22466,
		Course:          1,
		Speed:           120,
		DeviceType:      routing.IPad,
		MemberID:        1,
		DeviceGUID:      "TEST_GPS",
		DeviceTimestamp: "2014-06-14 17:43:35",
	}
	_, err = service.SetGPS(query)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationGetLastLocation(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	oService := &routing.Service{Client: client}
	opts, err := oService.GetOptimizations(&routing.RouteQuery{
		Limit: 1,
	})
	if err != nil {
		t.Error("Error occured in external service:", err)
		return
	}
	if len(opts) < 1 {
		t.Skip("Not enough routes to test activity stream")
	}
	opt, err := oService.GetOptimization(&routing.OptimizationParameters{
		ProblemID: opts[0].ProblemID,
	})
	if len(opt.Routes) < 1 {
		t.Skip("Not enough routes to test activity stream")
	}
	_, err = service.GetLastLocation(opt.Routes[0].ID)
	if err != nil {
		t.Error(err)
	}
}

func TestIntegrationGetDeviceLocationHistory(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	oService := &routing.Service{Client: client}
	opts, err := oService.GetOptimizations(&routing.RouteQuery{
		Limit: 1,
	})
	if err != nil {
		t.Error("Error occured in external service:", err)
		return
	}
	if len(opts) < 1 {
		t.Skip("Not enough routes to test activity stream")
	}
	opt, err := oService.GetOptimization(&routing.OptimizationParameters{
		ProblemID: opts[0].ProblemID,
	})
	if len(opt.Routes) < 1 {
		t.Skip("Not enough routes to test activity stream")
	}
	_, err = service.GetDeviceLocationHistory(&TrackingHistoryQuery{RouteID: opt.Routes[0].ID})
	if err != nil {
		t.Error(err)
	}
}

func TestIntegrationAssetTracking(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if _, err := service.TrackAssets("R7A9P1N9"); err != nil {
		t.Error(err)
	}
}
