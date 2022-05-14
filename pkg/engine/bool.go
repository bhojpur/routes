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
	"errors"
	"strings"
)

func UnmarshalJSON(bytes []byte) (bool, error) {
	str := string(bytes)
	if strings.HasPrefix(str, `"`) && strings.HasSuffix(str, `"`) {
		str = str[1 : len(str)-1]
	}
	if strings.ToLower(str) == "true" || str == "1" {
		return true, nil
	} else if strings.ToLower(str) == "false" || str == "0" || str == "null" || str == "" {
		return false, nil
	} else {
		return false, errors.New("Can't unmarshall unknown format to boolean " + str)
	}
}

type Bool bool

func (b *Bool) UnmarshalJSON(bytes []byte) error {
	res, err := UnmarshalJSON(bytes)
	if err != nil {
		return err
	}
	*b = Bool(res)
	return nil
}

func (b *Bool) MarshalJSON() ([]byte, error) {
	if *b == true {
		return []byte("true"), nil
	}
	return []byte("false"), nil
}

type StringBool Bool

func (b *StringBool) UnmarshalJSON(bytes []byte) error {
	res, err := UnmarshalJSON(bytes)
	if err != nil {
		return err
	}
	*b = StringBool(res)
	return nil
}

func (b *StringBool) MarshalJSON() ([]byte, error) {
	if *b == true {
		return []byte("\"TRUE\""), nil
	}
	return []byte("\"FALSE\""), nil
}
