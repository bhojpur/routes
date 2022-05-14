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

type Contact struct {
	ID           uint64  `json:"address_id,omitempty"`
	Group        string  `json:"address_group,omitempty"`
	Alias        string  `json:"address_alias,omitempty"`
	Address1     string  `json:"address_1"`
	Address2     string  `json:"address_2,omitempty"`
	FirstName    string  `json:"first_name,omitempty"`
	LastName     string  `json:"last_name,omitempty"`
	Email        string  `json:"address_email,omitempty"`
	PhoneNumber  string  `json:"address_phone_number,omitempty"`
	City         string  `json:"address_city,omitempty"`
	StateID      string  `json:"address_state_id,omitempty"`
	CountryID    string  `json:"address_country_id,omitempty"`
	PIN          string  `json:"address_pin,omitempty"`
	CachedLat    float64 `json:"cached_lat"`
	CachedLong   float64 `json:"cached_long"`
	Icon         string  `json:"address_icon,omitempty"`
	Color        string  `json:"color,omitempty"`
	CurbsideLat  float64 `json:"curbside_lat,omitempty"`
	CurbsideLong float64 `json:"curbside_long,omitempty"`
}

type Query struct {
	//Comma separated list of ids
	AddressID string `http:"address_id"`
	Limit     uint   `http:"limit"`
	Offset    uint   `http:"offset"`
	Start     uint   `http:"start"`
	Query     string `http:"query"`
	Fields    string `http:"fields"`
	Display   string `http:"display"`
}
