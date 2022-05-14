package vehicles

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

import "encoding/json"

type Vehicle struct {
	ID           string      `json:"vehicle_id"`
	CreatedTime  string      `json:"created_time"`
	MemberID     int64       `json:"member_id"`
	Alias        string      `json:"vehicle_alias,omitempty"`
	VIN          string      `json:"vehicle_vin,omitempty"`
	RegState     string      `json:"vehicle_reg_state,omitempty"`
	RegStateID   string      `json:"vehicle_reg_state_id,omitempty"`
	RegCountry   string      `json:"vehicle_reg_country,omitempty"`
	RegCountryID string      `json:"vehicle_reg_country_id,omitempty"`
	LicensePlate string      `json:"vehicle_license_plate,omitempty"`
	Make         string      `json:"vehicle_make,omitempty"`
	ModelYear    string      `json:"vehicle_model_year,omitempty"`
	Model        string      `json:"vehicle_model,omitempty"`
	YearAcquired string      `json:"vehicle_year_acquired,omitempty"`
	AxleCount    json.Number `json:"vehicle_axle_count,omitempty"`
	KmplCity     json.Number `json:"kmpl_city,omitempty"`
	KmplHighway  json.Number `json:"kmpl_highway,omitempty"`
	FuelType     string      `json:"fuel_type,omitempty"`
	HeightInches json.Number `json:"height_inches,omitempty"`
	WeightLB     json.Number `json:"weight_lb,omitempty"`
}
