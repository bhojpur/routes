package engine

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
	"testing"
)

func TestBoolUnmarshalling(t *testing.T) {
	str := &struct {
		FromStringInteger Bool `json:"fsi"`
		FromInteger       Bool `json:"fi"`
		FromBoolean       Bool `json:"fb"`
		FromStringBoolean Bool `json:"fsb"`
	}{}
	err := json.Unmarshal([]byte(`{"fsi":"1","fi":1,"fb": true,"fsb":"true"}`), str)
	if err != nil {
		t.Error("Could not unmarshall integer to boolean: ", err)
	}
	if str.FromBoolean != true || str.FromInteger != true || str.FromStringBoolean != true || str.FromStringInteger != true {
		t.Error("Error occured while unmarshalling bool")
	}
	err = json.Unmarshal([]byte(`{"fsi":"0","fi":0,"fb": false,"fsb":"false"}`), str)
	if err != nil {
		t.Error("Could not unmarshall integer to boolean: ", err)
	}
	if str.FromBoolean != false || str.FromInteger != false || str.FromStringBoolean != false || str.FromStringInteger != false {
		t.Error("Error occured while unmarshalling bool")
	}
	err = json.Unmarshal([]byte(`{"fsi":"10"}`), str)
	if err == nil {
		t.Error("Marshalling undefined string should have errored")
	}
}

func TestBoolMarshalling(t *testing.T) {
	str := &struct {
		True  Bool `json:"true"`
		False Bool `json:"false"`
	}{True: true, False: false}
	byt, err := json.Marshal(str)
	if err != nil {
		t.Error("Error occured while marshalling bool: ", err)
	}
	if string(byt) != `{"true":true,"false":false}` {
		t.Error("Error occured while marshalling bool.")
	}
}

func TestStringBoolUnmarshalling(t *testing.T) {
	str := &struct {
		FromStringInteger StringBool `json:"fsi"`
		FromInteger       StringBool `json:"fi"`
		FromBoolean       StringBool `json:"fb"`
		FromStringBoolean StringBool `json:"fsb"`
	}{}
	err := json.Unmarshal([]byte(`{"fsi":"1","fi":1,"fb": true,"fsb":"true"}`), str)
	if err != nil {
		t.Error("Could not unmarshall integer to boolean: ", err)
	}
	if str.FromBoolean != true || str.FromInteger != true || str.FromStringBoolean != true || str.FromStringInteger != true {
		t.Error("Error occured while unmarshalling bool")
	}
	err = json.Unmarshal([]byte(`{"fsi":"0","fi":0,"fb": false,"fsb":"false"}`), str)
	if err != nil {
		t.Error("Could not unmarshall integer to boolean: ", err)
	}
	if str.FromBoolean != false || str.FromInteger != false || str.FromStringBoolean != false || str.FromStringInteger != false {
		t.Error("Error occured while unmarshalling bool")
	}
	err = json.Unmarshal([]byte(`{"fsi":"10"}`), str)
	if err == nil {
		t.Error("Marshalling undefined string should have errored")
	}
}

func TestStringBoolMarshalling(t *testing.T) {
	str := &struct {
		True  StringBool `json:"true"`
		False StringBool `json:"false"`
	}{True: true, False: false}
	byt, err := json.Marshal(str)
	if err != nil {
		t.Error("Error occured while marshalling bool: ", err)
	}
	if string(byt) != `{"true":"TRUE","false":"FALSE"}` {
		t.Error("Error occured while marshalling bool.")
	}
}
