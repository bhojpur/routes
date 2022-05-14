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
	"encoding/json"
)

type Arrival struct {
	FromUnixTimestamp int `json:"from_unix_timestamp,omitempty"`
	ToUnixTimestamp   int `json:"to_unix_timestamp,omitempty"`
}

type AssetTracking struct {
	Arrival                  []Arrival        `json:"arrival,omitempty"`
	CustomData               *json.RawMessage `json:"custom_data,omitempty"`
	Delivered                bool             `json:"delivered,omitempty"`
	DestinationAddress1      string           `json:"destination_address_1,omitempty"`
	DestinationAddress2      string           `json:"destination_address_2,omitempty"`
	DriverName               string           `json:"driver_name,omitempty"`
	DriverPhone              string           `json:"driver_phone,omitempty"`
	DriverPicture            string           `json:"driver_picture,omitempty"`
	LargeLogoAlignment       string           `json:"large_logo_alignment,omitempty"`
	LargeLogoURI             string           `json:"large_logo_uri,omitempty"`
	Locations                []Locations      `json:"locations,omitempty"`
	MapColor                 string           `json:"map_color,omitempty"`
	MobileLogoAlignment      string           `json:"mobile_logo_alignment,omitempty"`
	MobileLogoURI            string           `json:"mobile_logo_uri,omitempty"`
	ShowMapZoomControls      bool             `json:"show_map_zoom_controls,omitempty"`
	StatusHistory            []StatusHistory  `json:"status_history,omitempty"`
	TimestampGeofenceVisited int              `json:"timestamp_geofence_visited,omitempty"`
	TimestampLastVisited     int              `json:"timestamp_last_visited,omitempty"`
	TrackingNumber           string           `json:"tracking_number,omitempty"`
	TrackingPageSubheadline  string           `json:"tracking_page_subheadline,omitempty"`
}

type Locations struct {
	Anchor      []int   `json:"anchor,omitempty"`
	Angle       int     `json:"angle,omitempty"`
	Icon        string  `json:"icon,omitempty"`
	Info        string  `json:"info,omitempty"`
	Lat         float64 `json:"lat,omitempty"`
	Long        float64 `json:"long,omitempty"`
	PopupAnchor []int   `json:"popupAnchor,omitempty"`
	Size        int     `json:"size,omitempty"`
}

type StatusHistory struct {
	Info          string `json:"info,omitempty"`
	UnixTimestamp int    `json:"unix_timestamp,omitempty"`
}
