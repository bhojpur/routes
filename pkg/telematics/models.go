package telematics

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

import engine "github.com/bhojpur/routes/pkg/engine"

type VendorSize string

const (
	VendorSizeGlobal   VendorSize = "global"
	VendorSizeRegional VendorSize = "regional"
	VendorSizeLocal    VendorSize = "local"
)

type VendorQuery struct {
	Size       VendorSize  `http:"size"`
	Integrated engine.Bool `http:"is_integrated"`
	Feature    string      `http:"feature"`
	Country    string      `http:"country"`
	Search     string      `http:"search"`
	Page       uint        `http:"page"`
	PerPage    uint        `http:"per_page"`
}

type Vendor struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Title        string     `json:"title"`
	Slug         string     `json:"slug"`
	Description  string     `json:"description"`
	LogoURL      string     `json:"logo_url"`
	WebsiteURL   string     `json:"website_url"`
	APIDocsURL   string     `json:"api_docs_url"`
	IsIntegrated string     `json:"is_integrated"`
	Size         VendorSize `json:"size"`
	Features     []Feature  `json:"features"`
	Countries    []Country  `json:"countries"`
}

type Country struct {
	ID          string `json:"id"`
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
}

type Feature struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	FeatureGroup string `json:"feature_group"`
}
