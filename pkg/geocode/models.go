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

type Address struct {
	PinCode string `json:"pincode,omitempty"`
	Name    string `json:"street_name,omitempty"`
}

type Coordinates struct {
	Lat  float64 `json:"lat,omitempty"`
	Long float64 `json:"long,omitempty"`
}

type Geocode struct {
	BBox          []Coordinates `json:"bbox"`
	Confidence    string        `json:"confidence,omitempty"`
	CountryRegion string        `json:"countryRegion,omitempty"`
	Coordinates   []Coordinates `json:"curbside_coordinates,omitempty"`
	Key           string        `json:"key,omitempty"`
	Lat           float64       `json:"lat"`
	Long          float64       `json:"long"`
	Name          string        `json:"name,omitempty"`
	PostalCode    string        `json:"postalCode,omitempty"`
	Type          string        `json:"type,omitempty"`
}

type Row struct {
	Address   string `json:"address"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Username  string `json:"username"`
	Website   string `json:"web-site"`
}

type BulkResponse struct {
	OptimizationProblemID string `json:"optimization_problem_id"`
	AddressCount          int    `json:"address_count"`
	Status                bool   `json:"status"`
}
