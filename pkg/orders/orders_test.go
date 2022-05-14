package orders

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

func TestIntegrationAdd(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	order := &Order{CachedLatitude: 48.335991, CachedLongitude: 31.18287, Address1: "258, Nirbhaya Dihra, Arrah, BH 802307, IN", AddressAlias: "Auto test address"}
	_, err := service.Add(order)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationGet(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}

	orders, _, err := service.GetAll(&Query{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(orders) < 1 {
		t.Skip("Not enough orders to test get 1.")
	}
	_, err = service.Get(orders[0].ID)
	if err != nil {
		t.Error(err)
		return
	}

}

func TestIntegrationRemove(t *testing.T) {
	//t.Skip("Skipping Removal integration test. Looks like the endpoint is broken.")
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	orders, _, err := service.GetAll(&Query{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(orders) < 1 {
		t.Skip("Not enough orders to test remove.")
	}
	success, err := service.Delete([]uint64{orders[0].ID})
	if err != nil {
		t.Error(err)
		return
	}
	if !success {
		t.Error("Deleting order failed")
		return
	}
}

func TestIntegrationUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	orders, _, err := service.GetAll(&Query{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(orders) < 1 {
		t.Skip("Not enough avoidance zones to test remove.")
	}
	orders[0].Address1 = "Some random" + strconv.Itoa(rand.Int())
	_, err = service.Update(&orders[0])
	if err != nil {
		t.Error(err)
		return
	}
}
