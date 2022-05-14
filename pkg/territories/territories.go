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
	"net/http"

	engine "github.com/bhojpur/routes/pkg/engine"
	"github.com/bhojpur/routes/pkg/utils"
)

const (
	avoidanceEndpoint = "/v1/avoidance"
	territoryEndpoint = "/v1/territory"
)

type Service struct {
	Client *engine.Client
}

func (s *Service) GetAvoidanceZone(query *Query) (*AvoidanceZone, error) {
	resp := &AvoidanceZone{}
	return resp, s.Client.Do(http.MethodGet, avoidanceEndpoint, query, resp)
}

func (s *Service) GetAvoidanceZones(query *Query) ([]AvoidanceZone, error) {
	resp := []AvoidanceZone{}
	return resp, s.Client.Do(http.MethodGet, avoidanceEndpoint, query, &resp)
}

func (s *Service) AddAvoidanceZone(data *AvoidanceZone) (*AvoidanceZone, error) {
	resp := &AvoidanceZone{}
	return resp, s.Client.Do(http.MethodPost, avoidanceEndpoint, data, resp)
}

func (s *Service) UpdateAvoidanceZone(data *AvoidanceZone) (*AvoidanceZone, error) {
	resp := &AvoidanceZone{}
	return resp, s.Client.Do(http.MethodPut, avoidanceEndpoint, data, resp)
}

func (s *Service) DeleteAvoidanceZone(data *Query) error {
	resp := &utils.StatusResponse{}
	err := s.Client.Do(http.MethodDelete, avoidanceEndpoint, data, resp)
	if err == nil && !resp.Status {
		return utils.ErrOperationFailed
	}
	return err
}

func (s *Service) GetTerritory(query *Query) (*Territory, error) {
	resp := &Territory{}
	return resp, s.Client.Do(http.MethodGet, territoryEndpoint, query, resp)
}

func (s *Service) GetTerritories(query *Query) ([]Territory, error) {
	resp := []Territory{}
	return resp, s.Client.Do(http.MethodGet, territoryEndpoint, query, &resp)
}
func (s *Service) AddTerritory(data *Territory) (*Territory, error) {
	resp := &Territory{}
	return resp, s.Client.Do(http.MethodPost, territoryEndpoint, data, resp)
}

func (s *Service) UpdateTerritory(data *Territory) (*Territory, error) {
	resp := &Territory{}
	return resp, s.Client.Do(http.MethodPut, territoryEndpoint, data, resp)
}

func (s *Service) DeleteTerritory(data *Query) error {
	resp := &utils.StatusResponse{}
	err := s.Client.Do(http.MethodDelete, territoryEndpoint, data, resp)
	if err == nil && !resp.Status {
		return utils.ErrOperationFailed
	}
	return err
}
