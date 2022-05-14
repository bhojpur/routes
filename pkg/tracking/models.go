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

import "github.com/bhojpur/routes/pkg/routing"

type TimePeriod string

const (
	TimePeriodToday     TimePeriod = "today"
	TimePeriodYesterday TimePeriod = "yesterday"
	TimePeriodThisMonth TimePeriod = "thismonth"
	TimePeriod7Days     TimePeriod = "7days"
	TimePeriod14Days    TimePeriod = "14days"
	TimePeriod30Days    TimePeriod = "30days"
	TimePeriod60Days    TimePeriod = "60days"
	TimePeriod90Days    TimePeriod = "90days"
	TimePeriodAllTime   TimePeriod = "all_time"
	TimePeriodCustom    TimePeriod = "custom"
)

type TrackingHistoryQuery struct {
	EndDate      uint64     `http:"end_date,omitempty"`
	LastPosition bool       `http:"last_position,omitempty"`
	RouteID      string     `http:"route_id"`
	StartDate    uint64     `http:"start_date,omitempty"`
	TimePeriod   TimePeriod `http:"time_period,omitempty"`
}

type GPS struct {
	Format          string             `http:"format"`
	MemberID        int                `http:"member_id"`
	RouteID         string             `http:"route_id"`
	TxID            string             `http:"tx_id"`
	VehicleID       int                `http:"vehicle_id"`
	Course          int                `http:"course"`
	Speed           float64            `http:"speed"`
	Latitude        float64            `http:"lat"`
	Longitude       float64            `http:"long"`
	Altitude        float64            `http:"altitude"`
	DeviceType      routing.DeviceType `http:"device_type"`
	DeviceGUID      string             `http:"device_guid"`
	DeviceTimestamp string             `http:"device_timestamp"`
	AppVersion      string             `http:"app_version"`
}
