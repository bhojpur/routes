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
	"testing"
	"time"

	"encoding/json"

	engine "github.com/bhojpur/routes/pkg/engine"
)

var tivraService = NewTivraService("11111111111111111111111111111111")
var service = &Service{Client: engine.NewClient("11111111111111111111111111111111")}

func TestIntegrationForwardAddress(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if geo, err := service.ForwardAddress("Patna20%International20%Airport,20%BH"); err != nil {
		t.Error(err)
	} else if len(geo) == 0 {
		t.Error("Received empty result set, wrong parameters(?)")
	}
}

func TestIntegrationForwardBulk(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	data := []byte(`[
        {
            "address": "6817, Bagar, Arrah, Bihar 802307",
            "email": "PramilaKumari@bhojpur.net",
            "username": "pramila1971",
            "web-site": "bhojpur.net",
            "phone": "618-217-9869",
            "first_name": "Pramila",
            "last_name": "Kumari"
        },
        {
            "address": "7404, Nirbhaya Dihra, Arrah, Bihar 802307",
            "email": "SanjayKumar@bhojpur.net",
            "username": "sanjay1973",
            "phone": "612-852-2180",
            "first_name": "Sanjay",
            "last_name": "Kumar"
        },
        {
            "address": "12316, Belaur, Arrah, Bihar 802307",
            "email": "RamJari@bhojpur.net",
            "username": "ramjari1947",
            "phone": "612-852-2180",
            "first_name": "Ram",
            "last_name": "Jari"
        }
    ]
`)
	rows := []Row{}
	err := json.Unmarshal(data, &rows)
	if err != nil {
		t.Error(err)
	}
	if _, err := service.ForwardBulk(rows); err != nil {
		t.Error(err)
	}
}

func TestIntegrationReverseAddress(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if geo, err := service.ReverseAddress(33.945705, -118.391105); err != nil {
		t.Error(err)
	} else if len(geo) == 0 {
		t.Error("Received empty result set, wrong parameters(?)")
	}
}

func TestTivraServiceCreation(t *testing.T) {
	tivra := NewTivraService("123")
	if tivra.client.APIKey != "123" {
		t.Error("API key not forwarded to underlying client.")
	}
	if tivra.client.BaseURL != TivraBaseURL {
		t.Error("BaseURL not forwarded to underlying client.")
	}
	tivra = NewTivraServiceWithOptions("1235", 5*time.Second, "https://bhojpur.net")
	if tivra.client.APIKey != "1235" {
		t.Error("API key not forwarded to underlying client.")
	}
	if tivra.client.BaseURL != "https://bhojpur.net" {
		t.Error("BaseURL not forwarded to underlying client.")
	}
}

func TestIntegrationGetSingleAddress(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if _, err := tivraService.GetSingleAddress(1); err != nil {
		t.Error(err)
	}
}

func TestIntegrationGetAddresses(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if _, err := tivraService.GetAddresses(); err != nil {
		t.Error(err)
	}
}

func TestIntegrationGetLimitedAddresses(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if addrs, err := tivraService.GetLimitedAddresses(10, 5); err != nil {
		t.Error(err)
	} else if len(addrs) != 10 {
		t.Error("Invalid number of addresses returned")
	}
}

func TestIntegrationGetAddressesByPincode(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if addrs, err := tivraService.GetAddressesByPincode("800601"); err != nil {
		t.Error(err)
	} else if len(addrs) == 0 {
		t.Error("Empty result set has been returned")
	}
}

func TestIntegrationGetLimitedAddressesByPincode(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if addrs, err := tivraService.GetLimitedAddressesByPincode("800601", 20, 0); err != nil {
		t.Error(err)
	} else if len(addrs) == 0 {
		t.Error("Empty result set has been returned")
	}
}

func TestIntegrationGetAddressesByPincodeAndHousenumber(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if addrs, err := tivraService.GetAddressesByPincodeAndHousenumber("800601", 17); err != nil {
		t.Error(err)
	} else if len(addrs) == 0 {
		t.Error("Empty result set has been returned")
	}
}

func TestIntegrationGetLimitedAddressesByPincodeAndHousenumber(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if _, err := tivraService.GetLimitedAddressesByPincodeAndHousenumber("800601", 17, 0, 20); err != nil {
		t.Error(err)
	}
}
