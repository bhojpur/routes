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

import "github.com/bhojpur/routes/pkg/geocode"

type Direction struct {
	Location *Location `json:"location,omitempty"`
	Steps    []Steps   `json:"steps,omitempty"`
}

type Location struct {
	DirectionsError string  `json:"directions_error,omitempty"`
	EndLocation     string  `json:"end_location,omitempty"`
	ErrorCode       int     `json:"error_code,omitempty"`
	Name            string  `json:"name,omitempty"`
	SegmentDistance float64 `json:"segment_distance,omitempty"`
	StartLocation   string  `json:"start_location,omitempty"`
	Time            int     `json:"time,omitempty"`
}

type Steps struct {
	CompassDirection string               `json:"compass_direction,omitempty"`
	Direction        string               `json:"direction,omitempty"`
	Directions       string               `json:"directions,omitempty"`
	Distance         float64              `json:"distance,omitempty"`
	DistanceUnit     string               `json:"distance_unit,omitempty"`
	DurationSec      int                  `json:"duration_sec,omitempty"`
	ManeuverPoint    *geocode.Coordinates `json:"maneuverPoint,omitempty"`
	ManeuverType     string               `json:"maneuverType,omitempty"`
}
