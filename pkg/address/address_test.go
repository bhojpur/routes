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
	"math/rand"
	"strconv"
	"testing"

	engine "github.com/bhojpur/routes/pkg/engine"
)

var client = engine.NewClient("11111111111111111111111111111111")
var service = &Service{Client: client}

func TestIntegrationGet(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	query := &Query{
		Limit:  10,
		Offset: 0,
	}

	_, _, err := service.Get(query)
	if err != nil {
		t.Error(err)
	}
}

func TestIntegrationAdd(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	contact := &Contact{
		FirstName:    "Pramila" + strconv.Itoa(rand.Int()),
		Alias:        "pramila" + strconv.Itoa(rand.Int()),
		Address1:     "Some address" + strconv.Itoa(rand.Int()),
		CachedLat:    38.024654,
		CachedLong:   -77.338814,
		Email:        "pramila@bhojpur.net",
		PhoneNumber:  "000-000-000",
		StateID:      "5",
		CountryID:    "3",
		City:         "City",
		PIN:          "000-000",
		CurbsideLat:  38.024654,
		CurbsideLong: -77.338814,
		Color:        "fffeee",
	}

	_, err := service.Add(contact)
	if err != nil {
		t.Error(err)
	}
}

func TestIntegrationRemove(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	query := &Query{
		Limit:  1,
		Offset: 0,
	}
	contacts, _, err := service.Get(query)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = service.Delete([]string{strconv.FormatUint(contacts[0].ID, 10)})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	query := &Query{
		Limit:  1,
		Offset: 0,
	}
	contacts, _, err := service.Get(query)
	if err != nil {
		t.Error(err)
		return
	}
	contacts[0].FirstName = "EditedName"
	_, err = service.Update(&contacts[0])
	if err != nil {
		t.Error(err)
		return
	}
}
