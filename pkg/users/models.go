package users

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
	engine "github.com/bhojpur/routes/pkg/engine"
)

type MemberType string

const (
	PrimaryAccount            MemberType = "PRIMARY_ACCOUNT"
	SubAccountAdmin           MemberType = "SUB_ACCOUNT_ADMIN"
	SubAccountRegionalManager MemberType = "SUB_ACCOUNT_REGIONAL_MANAGER"
	SubAccountDispatcher      MemberType = "SUB_ACCOUNT_DISPATCHER"
	SubAccountDriver          MemberType = "SUB_ACCOUNT_DRIVER"
)

type MemberBase struct {
	FirstName string     `json:"member_first_name,omitempty"`
	LastName  string     `json:"member_last_name,omitempty"`
	Phone     string     `json:"member_phone,omitempty"`
	Type      MemberType `json:"member_type,omitempty"`
	Email     string     `json:"member_email,omitempty"`
	Password  string     `json:"member_password,omitempty"`
	PINCode   string     `json:"member_pincode,omitempty"`

	PreferredUnits    string `json:"preferred_units,omitempty"`
	PreferredLanguage string `json:"preferred_language,omitempty"`
	Timezone          string `json:"timezone,omitempty"`

	RegionCountryID string `json:"user_reg_country_id,omitempty"`
	RegionStateID   string `json:"user_reg_state_id,omitempty"`

	HideRouteAddresses   engine.StringBool `json:"HIDE_ROUTED_ADDRESSES,omitempty"`
	HideVisitedAddresses engine.StringBool `json:"HIDE_VISITED_ADDRESSES,omitempty"`
	HideNonFutureRoutes  engine.StringBool `json:"HIDE_NONFUTURE_ROUTES,omitempty"`
	ReadOnly             engine.StringBool `json:"READONLY_USER,omitempty"`
	ShowAllDrivers       engine.StringBool `json:"SHOW_ALL_DRIVERS,omitempty"`
	ShowAllVehicles      engine.StringBool `json:"SHOW_ALLVehicles,omitempty"`
	DateOfBirth          string            `json:"date_of_birth,omitempty"`
}

type Member struct {
	MemberBase
	ID      int64 `json:"member_id,string,omitempty"`
	OwnerID int64 `json:"OWNER_MEMBER_ID,string,omitempty"`
}

type Session struct {
	Status                        bool        `json:"status"`
	Error                         string      `json:"error,omitempty"`
	GeocodeService                string      `json:"geocode_service"`
	SessionID                     int64       `json:"session_id"`
	SessionGUID                   string      `json:"session_guid"`
	MemberID                      string      `json:"member_id"`
	APIKey                        string      `json:"api_key"`
	TrackingTTL                   int         `json:"tracking_ttl"`
	GeofencePolygonShape          string      `json:"geofence_polygon_shape"`
	GeofencePolygonSize           int         `json:"geofence_polygon_size"`
	GeofenceTimeOnsiteTriggerSecs int         `json:"geofence_time_onsite_trigger_secs"`
	GeofenceMinimumTriggerSpeed   int         `json:"geofence_minimum_trigger_speed"`
	IsSubscriptionPastDue         engine.Bool `json:"is_subscription_past_due"`
	VisitedDepartedEnabled        engine.Bool `json:"visited_departed_enabled"`
	LongPressEnabled              engine.Bool `json:"long_press_enabled"`
	AccountTypeID                 string      `json:"account_type_id"`
	AccountTypeAlias              string      `json:"account_type_alias"`
	MemberType                    string      `json:"member_type"`
	MaxStopsPerRoute              string      `json:"max_stops_per_route"`
	MaxRoutes                     string      `json:"max_routes"`
	RoutesPlanned                 string      `json:"routes_planned"`

	HideRouteAddresses   engine.Bool `json:"HIDE_ROUTED_ADDRESSES"`
	HideVisitedAddresses engine.Bool `json:"HIDE_VISITED_ADDRESSES"`
	HideNonFutureRoutes  engine.Bool `json:"HIDE_NONFUTURE_ROUTES"`
	ReadOnly             engine.Bool `json:"READONLY_USER"`
	AutoLogoutTs         int         `json:"auto_logout_ts"`
}

type KeyValue struct {
	MemberID int    `json:"member_id,omitempty"`
	Key      string `json:"config_key,omitempty"`
	Value    string `json:"config_value,omitempty"`
}

type WebinarRegistration struct {
	EmailAddress string `json:"email_address"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	PhoneNumber  string `json:"phone_number"`
	CompanyName  string `json:"company_name"`
	MemberID     string `json:"member_id"`
	WebiinarDate string `json:"webiinar_date"`
}
