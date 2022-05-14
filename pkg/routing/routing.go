package routing

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
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	engine "github.com/bhojpur/routes/pkg/engine"
	"github.com/bhojpur/routes/pkg/utils"
)

const (
	addressEndpoint              = "/v1/address"
	routeEndpoint                = "/v1/route"
	optimizationEndpoint         = "/v1/optimization_problem"
	reoptimizeRouteEndpoint      = "/v1/route/reoptimize"
	duplicateRouteEndpoint       = "/v1/actions/duplicate_route"
	notesEndpoint                = "/v1/actions/addRouteNotes"
	moveRouteDestinationEndpoint = "/v1/actions/route/move_route_destination"
	mergeRoutesEndpoint          = "/v1/actions/merge_routes"
	shareRouteEndpoint           = "/v1/actions/route/share_route"
	updateAddressVisitedEndpoint = "/v1/api/route/mark_address_visited"
	markAddressDepartedEndpoint  = "/v1/api/route/mark_address_departed"
)

// codebeat:disable[TOO_MANY_FUNCTIONS]
type Service struct {
	Client *engine.Client
}

//Optimization
func (s *Service) GetOptimization(parameters *OptimizationParameters) (*DataObject, error) {
	resp := &DataObject{}
	return resp, s.Client.Do(http.MethodGet, optimizationEndpoint, parameters, resp)
}

type getOptimizationsResponse struct {
	Optimizations []DataObject `json:"optimizations"`
}

func (s *Service) GetOptimizations(parameters *RouteQuery) ([]DataObject, error) {
	resp := &getOptimizationsResponse{}
	return resp.Optimizations, s.Client.Do(http.MethodGet, optimizationEndpoint, parameters, resp)
}

func (s *Service) RunOptimization(parameters *OptimizationParameters) (*DataObject, error) {
	resp := &DataObject{}
	return resp, s.Client.Do(http.MethodPost, optimizationEndpoint, parameters, resp)
}

func (s *Service) UpdateOptimization(parameters *OptimizationParameters) (*DataObject, error) {
	resp := &DataObject{}
	return resp, s.Client.Do(http.MethodPut, optimizationEndpoint, parameters, resp)
}

type deleteOptimizationRequest struct {
	OptimizationProblemIDs []string `json:"optimization_problem_ids"`
}

func (s *Service) DeleteOptimization(optimizationProblemID string) error {
	return s.DeleteOptimizations(optimizationProblemID)
}

func (s *Service) DeleteOptimizations(optimizationProblemID ...string) error {
	return s.Client.Do(http.MethodDelete, optimizationEndpoint, &deleteOptimizationRequest{OptimizationProblemIDs: optimizationProblemID}, nil)
}

//Addresses
func (s *Service) GetAddress(query *AddressQuery) (*Address, error) {
	resp := &Address{}
	return resp, s.Client.Do(http.MethodGet, addressEndpoint, query, resp)
}

func (s *Service) UpdateAddress(data *Address) (*Address, error) {
	resp := &Address{}
	return resp, s.Client.Do(http.MethodPut, addressEndpoint, data, resp)
}

type deleteAddressRequest struct {
	OptimizationProblemID string      `http:"optimization_problem_id"`
	RouteDestinationID    json.Number `http:"route_destination_id"`
}

type deleteAddressResponse struct {
	Deleted            bool         `json:"deleted"`
	RouteDestinationID *json.Number `json:"route_destination_id,omitempty"`
}

// DeleteAddress removes a destination (an address) with specified route_destination_id from an optimization problem with specified optimization_problem_id.
func (s *Service) DeleteAddress(optimizationID string, routeDestinationID string) (*json.Number, error) {
	req := &deleteAddressRequest{
		OptimizationProblemID: optimizationID,
		RouteDestinationID:    json.Number(routeDestinationID),
	}
	resp := &deleteAddressResponse{}
	err := s.Client.Do(http.MethodDelete, addressEndpoint, req, resp)
	if err == nil && !resp.Deleted {
		return nil, utils.ErrOperationFailed
	}
	return resp.RouteDestinationID, err
}

//Routes
func (s *Service) GetRoute(query *RouteQuery) (*Route, error) {
	resp := &Route{}
	return resp, s.Client.Do(http.MethodGet, routeEndpoint, query, resp)
}

type getRouteIDRequest struct {
	ProblemID         string `http:"optimization_problem_id"`
	WaitForFinalState int    `http:"wait_for_final_state,string"`
}

func (s *Service) GetRouteID(problemID string) (string, error) {
	request := &getRouteIDRequest{ProblemID: problemID, WaitForFinalState: 1}
	response := &DataObject{}
	err := s.Client.Do(http.MethodGet, optimizationEndpoint, request, response)
	if err != nil {
		return "", err
	}
	if len(response.Routes) > 0 {
		return response.Routes[0].ID, nil
	}
	return "", errors.New("Could not find requested route")
}

type markAddress struct {
	RouteID    string `http:"route_id" json:"-"`
	AddressID  string `http:"address_id" json:"-"`
	MemberID   int    `http:"member_id" json:"-"`
	IsVisited  bool   `http:"is_visited" json:"-"`
	IsDeparted bool   `http:"is_departed" json:"-"`
}

func (s *Service) MarkAddressAsVisited(address *Address) (bool, error) {
	req := &markAddress{
		RouteID:   address.RouteID,
		AddressID: address.RouteDestinationID.String(),
		MemberID:  address.MemberID,
		IsVisited: address.IsVisited,
	}
	resp := &utils.StatusResponse{}
	return resp.Status, s.Client.Do(http.MethodPut, updateAddressVisitedEndpoint, req, resp)
}

func (s *Service) MarkAddressAsDeparted(address *Address) (bool, error) {
	req := &markAddress{
		RouteID:    address.RouteID,
		AddressID:  address.RouteDestinationID.String(),
		MemberID:   address.MemberID,
		IsDeparted: address.IsDeparted,
	}
	resp := &utils.StatusResponse{}
	return resp.Status, s.Client.Do(http.MethodPut, markAddressDepartedEndpoint, req, resp)
}

type duplicateRouteResponse struct {
	ProblemID string `json:"optimization_problem_id,omitempty"`
	Success   bool   `json:"success"`
}

type duplicateRouteRequest struct {
	RouteID string `http:"route_id"`
	To      string `http:"to"`
}

func (s *Service) DuplicateRoute(routeID string) (string, error) {
	request := &duplicateRouteRequest{RouteID: routeID, To: "none"}
	response := &duplicateRouteResponse{}
	err := s.Client.Do(http.MethodGet, duplicateRouteEndpoint, request, response)
	if err != nil {
		return "", err
	}
	if response.Success && response.ProblemID != "" {
		return s.GetRouteID(response.ProblemID)
	}
	return "", errors.New("Could not find requested route")
}

func (s *Service) GetTeamRoutes(query *RouteQuery) ([]Route, error) {
	resp := []Route{}
	return resp, s.Client.Do(http.MethodGet, routeEndpoint, query, &resp)
}

func (s *Service) UpdateRoute(route *Route) (*Route, error) {
	resp := &Route{}
	return resp, s.Client.Do(http.MethodPut, routeEndpoint, route, resp)
}

type deleteRequest struct {
	RouteID string `http:"route_id"`
}

type DeletedRoutes struct {
	Deleted  bool     `json:"deleted"`
	RouteID  string   `json:"route_id,omitempty"`
	RouteIDs []string `json:"route_ids,omitempty"`
}

func (s *Service) DeleteRoutes(routeIDs ...string) (DeletedRoutes, error) {
	request := &deleteRequest{
		RouteID: strings.Join(routeIDs, ","),
	}
	resp := DeletedRoutes{}
	return resp, s.Client.Do(http.MethodDelete, routeEndpoint, request, &resp)
}

type MergeRequest struct {
	RouteIDs       string  `form:"route_ids"`
	DepotAddress   string  `form:"depot_address"`
	RemoveOrigin   bool    `form:"remove_origin"`
	DepotLatitude  float64 `form:"depot_lat"`
	DepotLongitude float64 `form:"depot_long"`
}

func (s *Service) MergeRoutes(request *MergeRequest) error {
	resp := &utils.StatusResponse{}
	err := s.Client.Do(http.MethodPost, mergeRoutesEndpoint, request, resp)
	if err == nil && !resp.Status {
		return utils.ErrOperationFailed
	}
	return err
}

type shareRequest struct {
	RecipientEmail string `form:"recipient_email"`
	RouteID        string `http:"route_id"`
	ResponseFormat string `http:"response_format"`
}

func (s *Service) ShareRoute(routeID string, email string) error {
	request := &shareRequest{
		RecipientEmail: email,
		RouteID:        routeID,
		ResponseFormat: "json",
	}
	resp := &utils.StatusResponse{}
	err := s.Client.Do(http.MethodPost, shareRouteEndpoint, request, resp)
	if err == nil && !resp.Status {
		return utils.ErrOperationFailed
	}
	return err
}

func (s *Service) GetAddressNotes(query *NoteQuery) ([]Note, error) {
	addressQuery := &AddressQuery{
		RouteID:            query.RouteID,
		RouteDestinationID: query.AddressID,
		Notes:              true,
	}
	addr, err := s.GetAddress(addressQuery)
	return addr.Notes, err
}

type addAddressNoteRequest struct {
	*NoteQuery
	NoteContents string `form:"strNoteContents"`
	UpdateType   string `form:"strUpdateType"`
}

type addAddressNoteResponse struct {
	Status bool  `json:"status"`
	Note   *Note `json:"note,omitempty"`
}

func (s *Service) AddAddressNote(query *NoteQuery, noteContents string) (*Note, error) {
	strUpdateType := "unclassified"
	if query.ActivityType != "" {
		strUpdateType = string(query.ActivityType)
	}

	request := &addAddressNoteRequest{
		NoteQuery:    query,
		UpdateType:   strUpdateType,
		NoteContents: noteContents,
	}
	resp := &addAddressNoteResponse{}
	err := s.Client.Do(http.MethodPost, notesEndpoint, request, resp)
	// always returns false (needs clarification)
	// if err == nil && !resp.Status {
	// 	return nil, utils.ErrOperationFailed
	// }
	return resp.Note, err
}

type addRouteDestinationsRequest struct {
	RouteID   string    `http:"route_id"`
	Addresses []Address `json:"addresses"`
}

func (s *Service) AddRouteDestinations(routeID string, addresses []Address) (*DataObject, error) {
	request := &addRouteDestinationsRequest{
		RouteID:   routeID,
		Addresses: addresses,
	}
	resp := &DataObject{}
	return resp, s.Client.Do(http.MethodPut, routeEndpoint, request, resp)
}

type removeRouteDestinationRequest struct {
	RouteID            string `http:"route_id"`
	RouteDestinationID string `http:"route_destination_id"`
}

type removeRouteDestinationResposne struct {
	Deleted bool `json:"deleted"`
}

func (s *Service) RemoveRouteDestination(routeID string, destinationID string) (bool, error) {
	request := &removeRouteDestinationRequest{
		RouteID:            routeID,
		RouteDestinationID: destinationID,
	}
	resp := &removeRouteDestinationResposne{}
	return resp.Deleted, s.Client.Do(http.MethodDelete, routeEndpoint, request, resp)
}

type DestinationMoveRequest struct {
	ToRouteID          string `form:"to_route_id"`
	RouteDestinationID string `form:"route_destination_id"`
	AfterDestinationID string `form:"after_destination_id"`
}

type moveDestinationToRouteResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

func (s *Service) MoveDestinationToRoute(query *DestinationMoveRequest) error {
	resp := &moveDestinationToRouteResponse{}
	err := s.Client.Do(http.MethodPost, moveRouteDestinationEndpoint, query, resp)
	if err == nil && !resp.Success {
		return utils.ErrOperationFailed
	}
	return err
}

type resequenceRouteRequest struct {
	RouteID             string      `http:"route_id"`
	DisableOptimization engine.Bool `http:"disable_optimization"`
	Optimize            string      `http:"optimize"`
}

func (s *Service) ResequenceRoute(routeID string) (bool, error) {
	resp := &utils.StatusResponse{}
	return resp.Status, s.Client.Do(http.MethodGet, reoptimizeRouteEndpoint, &resequenceRouteRequest{RouteID: routeID, DisableOptimization: false, Optimize: "Distance;"}, resp)
}

// codebeat:enable[TOO_MANY_FUNCTIONS]
