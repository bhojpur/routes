package territories

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

func TestIntegrationAddAvoidanceZone(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	zone := &AvoidanceZone{
		Name:  "John" + strconv.Itoa(rand.Int()),
		Color: "beeeee",
		Territory: TerritoryShape{
			Type: Circle,
			Data: []string{"37.569752822786455,-77.47833251953125", "5000"},
		},
	}
	_, err := service.AddAvoidanceZone(zone)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationGetAvoidanceZones(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}

	zones, err := service.GetAvoidanceZones(&Query{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(zones) < 1 {
		t.Skip("Not enough avoidance zones to test get 1.")
	}
	_, err = service.GetAvoidanceZone(&Query{ID: zones[0].ID})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationRemoveAvoidanceZone(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	zones, err := service.GetAvoidanceZones(&Query{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(zones) < 1 {
		t.Skip("Not enough avoidance zones to test remove.")
	}
	err = service.DeleteAvoidanceZone(&Query{ID: zones[0].ID})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationUpdateAvoidanceZone(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	zones, err := service.GetAvoidanceZones(&Query{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(zones) < 1 {
		t.Skip("Not enough avoidance zones to test remove.")
	}
	zones[0].Name = "Johny" + strconv.Itoa(rand.Int())
	_, err = service.UpdateAvoidanceZone(&zones[0])
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationAddTerritory(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	zone := &Territory{
		Name:  "John" + strconv.Itoa(rand.Int()),
		Color: "beeeee",
		Territory: TerritoryShape{
			Type: Circle,
			Data: []string{"37.569752822786455,-77.47833251953125", "5000"},
		},
	}
	_, err := service.AddTerritory(zone)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationGetTerritories(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}

	zones, err := service.GetTerritories(&Query{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(zones) < 1 {
		t.Skip("Not enough territories to test get 1.")
	}
	_, err = service.GetTerritory(&Query{ID: zones[0].ID})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationRemoveTerritories(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	zones, err := service.GetTerritories(&Query{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(zones) < 1 {
		t.Skip("Not enough territories to test remove.")
	}
	err = service.DeleteTerritory(&Query{ID: zones[0].ID})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationUpdateTerritory(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	zones, err := service.GetTerritories(&Query{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(zones) < 1 {
		t.Skip("Not enough territories to test remove.")
	}
	zones[0].Name = "Johny" + strconv.Itoa(rand.Int())
	_, err = service.UpdateTerritory(&zones[0])
	if err != nil {
		t.Error(err)
		return
	}
}
