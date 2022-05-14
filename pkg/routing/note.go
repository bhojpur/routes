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

type Note struct {
	//NoteID
	ID                 int     `json:"note_it,omitempty"`
	RouteID            string  `json:"route_id,omitempty"`
	RouteDestinationID int     `json:"route_destination_id,omitempty"`
	UploadID           string  `json:"upload_id"`
	TimestampAdded     uint64  `json:"ts_added,omitempty"`
	Latitude           float64 `json:"lat"`
	Longitude          float64 `json:"long"`
	ActivityType       string  `json:"activity_type"`
	Contents           string  `json:"contents"`
	UploadType         string  `json:"upload_type"`
	UploadURL          string  `json:"upload_url"`
	UploadExtension    string  `json:"upload_extension"`
	DeviceType         string  `json:"device_type"`
}

type NoteQuery struct {
	RouteID      string           `http:"route_id"`
	AddressID    string           `http:"address_id"`
	Latitude     float64          `http:"dev_lat"`
	Longitude    float64          `http:"dev_long"`
	DeviceType   string           `http:"device_type"`
	ActivityType StatusUpdateType `http:"strUpdateType"`
}

type StatusUpdateType string

const (
	Pickup                      StatusUpdateType = "pickup"
	DropOff                     StatusUpdateType = "dropoff"
	NoAnswer                    StatusUpdateType = "noanswer"
	NotFound                    StatusUpdateType = "notfound"
	NotPaid                     StatusUpdateType = "notpaid"
	Paid                        StatusUpdateType = "paid"
	WrongDelivery               StatusUpdateType = "wrongdelivery"
	WrongAddressRecipient       StatusUpdateType = "wrongaddressrecipient"
	NotPresent                  StatusUpdateType = "notpresent"
	PartsMissing                StatusUpdateType = "parts_missing"
	ServiceRendered             StatusUpdateType = "service_rendered"
	FollowUp                    StatusUpdateType = "follow_up"
	LeftInformation             StatusUpdateType = "left_information"
	SpokeWithDecisionMaker      StatusUpdateType = "spoke_with_decision_maker"
	SpokeWithDecisionInfluencer StatusUpdateType = "spoke_with_decision_influencer"
	CompetitiveAccount          StatusUpdateType = "competitive_account"
	ScheduledFollowUpMeeting    StatusUpdateType = "scheduled_follow_up_meeting"
	ScheduledLunch              StatusUpdateType = "scheduled_lunch"
	ScheduledProductDemo        StatusUpdateType = "scheduled_product_demo"
	ScheduledClinicalDemo       StatusUpdateType = "scheduled_clinical_demo"
	NoOpportunity               StatusUpdateType = "no_opportunity"
)
