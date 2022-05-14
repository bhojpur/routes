package activity

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
	"time"

	"github.com/bhojpur/routes/pkg/users"
)

type Activity struct {
	ID        string       `json:"activity_id,omitempty"`
	Type      ActivityType `json:"activity_type"`
	Timestamp uint64       `json:"activity_timestamp,omitempty"`
	Message   string       `json:"activity_message,omitempty"`
	RouteID   string       `json:"route_destination_id,omitempty"`
	RouteName string       `json:"route_name,omitempty"`

	NoteID       *string `json:"note_id,omitempty"`
	NoteType     *string `json:"note_type,omitempty"`
	NoteContents *string `json:"note_contents,omitempty"`
	NoteFile     *string `json:"note_file,omitempty"`

	Member *users.Member `json:"member,omitempty"`
}

type ActivityType string

const (
	DeleteDestination       ActivityType = "delete-destination"
	InsertDestination       ActivityType = "insert-destination"
	MarkDestinationDeparted ActivityType = "mark-destination-departed"
	MarkDestinationVisited  ActivityType = "mark-destination-visited"
	MemberCreated           ActivityType = "member-created"
	MemberDeleted           ActivityType = "member-deleted"
	MemberModified          ActivityType = "member-modified"
	MoveDestination         ActivityType = "move-destination"
	NoteInsert              ActivityType = "note-insert"
	RouteDelete             ActivityType = "route-delete"
	RouteOptimized          ActivityType = "route-optimized"
	RouteOwnerChanged       ActivityType = "route-owner-changed"
	UpdateDestinations      ActivityType = "update-destinations"
	AreaAdded               ActivityType = "area-added"
	AreaRemoved             ActivityType = "area-removed"
	AreaUpdated             ActivityType = "area-updated"
	DestinationOutSequence  ActivityType = "destination-out-sequence"
	DriverArrivedEarly      ActivityType = "driver-arrived-early"
	DriverArrivedOnTime     ActivityType = "driver-arrived-on-time"
	DriverArrivedLate       ActivityType = "driver-arrived-late"
	GeofenceEntered         ActivityType = "geofence-entered"
	GeofenceLeft            ActivityType = "geofence-left"
	UserMessage             ActivityType = "user_message"
)

type Query struct {
	RouteID  string       `http:"route_id"`
	DeviceID string       `http:"device_id"`
	Type     ActivityType `http:"activity_type"`
	MemberID int          `http:"member_id"`
	Team     bool         `http:"team"`
	Limit    uint         `http:"limit"`
	Offset   uint         `http:"offset"`
	Start    time.Time    `http:"start"`
	End      time.Time    `http:"end"`
}
