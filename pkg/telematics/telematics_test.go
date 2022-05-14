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
	"testing"
	"time"
)

var service = NewTelematicsService("11111111111111111111111111111111")

func TestTelematicsServiceCreation(t *testing.T) {
	Telematics := NewTelematicsService("123")
	if Telematics.client.APIKey != "123" {
		t.Error("API key not forwarded to underlying client.")
	}
	if Telematics.client.BaseURL != TelematicsBaseURL {
		t.Error("BaseURL not forwarded to underlying client.")
	}
	Telematics = NewTelematicsServiceWithOptions("1235", 5*time.Second, "https://bhojpur.net")
	if Telematics.client.APIKey != "1235" {
		t.Error("API key not forwarded to underlying client.")
	}
	if Telematics.client.BaseURL != "https://bhojpur.net" {
		t.Error("BaseURL not forwarded to underlying client.")
	}
}

func TestIntegrationGetVendors(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if _, err := service.GetVendors(); err != nil {
		t.Error(err)
	}
}

func TestIntegrationGetVendor(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if _, err := service.GetVendor(153); err != nil {
		t.Error(err)
	}
}

func TestIntegrationSearchVendors(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if _, err := service.SearchVendors(&VendorQuery{
		Integrated: true,
		Feature:    "Satellite",
		Country:    "GB",
	}); err != nil {
		t.Error(err)
	}
}

func TestIngegrationCompareVendors(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if _, err := service.CompareVendors(52, 53); err != nil {
		t.Error(err)
	}
}
