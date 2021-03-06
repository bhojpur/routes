package utils

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
	"reflect"
	"testing"
)

type testStruct struct {
	String    string  `http:"string"`
	Int       int     `http:"int"`
	Int16     int32   `http:"int16"`
	Int32     int32   `http:"int32"`
	Int64     int64   `http:"int64"`
	UInt      uint    `http:"uint"`
	UInt16    int32   `http:"uint16"`
	UInt32    int32   `http:"uint32"`
	UInt64    int64   `http:"uint64"`
	Float32   float32 `http:"float32"`
	Float64   float64 `http:"float64"`
	Float64E  float64 `http:"float64e"`
	Boolean   bool    `http:"boolean"`
	NoConvert string
}

func TestStructToURLValues(t *testing.T) {
	test := &testStruct{
		String:    "foo",
		Int:       423,
		Int16:     552,
		Int32:     -12555,
		Int64:     -423424,
		UInt:      45,
		UInt16:    52,
		UInt32:    155,
		UInt64:    4242,
		Float64:   -4242,
		Float32:   15,
		Boolean:   true,
		NoConvert: "hey",
	}
	urlValues := StructToURLValues("http", test)
	if urlValues.Get("NoConvert") != "" {
		t.Error("Value that should not be converted somehow ended up being.")
	}
	if urlValues.Get("float64") != "-4242.0000" {
		t.Error("Wrong float64 value.")
	}
	if urlValues.Get("uint") != "45" {
		t.Error("Wrong uint value.")
	}
	if urlValues.Get("string") != "foo" {
		t.Error("Wrong string value.")
	}
	if urlValues.Get("float32") != "15.0000" {
		t.Error("Wrong float32 value.")
	}
	if urlValues.Get("float64e") != "" {
		t.Error("Empty value marshalled")
	}
}

func TestIsEmptyValue(t *testing.T) {
	if isEmptyValue(reflect.ValueOf(true)) || !isEmptyValue(reflect.ValueOf(false)) {
		t.Error("Boolean empty value check failed")
	}
	if isEmptyValue(reflect.ValueOf("e")) || !isEmptyValue(reflect.ValueOf("")) {
		t.Error("String empty value check failed")
	}
	if isEmptyValue(reflect.ValueOf(5)) || isEmptyValue(reflect.ValueOf(-2.5)) || !isEmptyValue(reflect.ValueOf(-0)) {
		t.Error("Numbers empty value check failed")
	}
	if !isEmptyValue(reflect.ValueOf((*testStruct)(nil))) || isEmptyValue(reflect.ValueOf(&testStruct{String: "abc"})) {
		t.Error("Interface empty value check failed")
	}
}
