//
// Copyright (c) 2019 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package util

import (
	"reflect"
	"strings"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/models"

	"github.com/stretchr/testify/assert"
)

func TestSplitComma(t *testing.T) {
	commaDelimmited := "Hel lo,, ,Test,Hi, "
	results := strings.FieldsFunc(commaDelimmited, SplitComma)
	// Should have 4 elements (space counts as an element)
	assert.Equal(t, 5, len(results))
	assert.Equal(t, results[0], "Hel lo")
	assert.Equal(t, results[1], " ")
	assert.Equal(t, results[2], "Test")
	assert.Equal(t, results[3], "Hi")
	assert.Equal(t, results[4], " ")
}
func TestSplitCommaEmpty(t *testing.T) {
	commaDelimmited := ""
	results := strings.FieldsFunc(commaDelimmited, SplitComma)
	// Should have 4 elements (space counts as an element)
	assert.Equal(t, 0, len(results))
}

func TestDeleteEmptyAndTrim(t *testing.T) {
	strings := []string{" Hel lo", "test ", " "}
	results := DeleteEmptyAndTrim(strings)
	// Should have 4 elements (space counts as an element)
	assert.Equal(t, 2, len(results))
	assert.Equal(t, results[0], "Hel lo")
	assert.Equal(t, results[1], "test")
}

func TestCoerceTypeStringToByteArray(t *testing.T) {
	myData := "" //string
	var expectedType []byte
	result, err := CoerceType(myData)
	assert.NoError(t, err)
	assert.IsType(t, reflect.TypeOf(expectedType), reflect.TypeOf(result))
}

func TestCoerceTypeByteArrayToByteArray(t *testing.T) {
	var myData []byte //string
	var expectedType []byte
	result, err := CoerceType(myData)
	assert.NoError(t, err)
	assert.IsType(t, reflect.TypeOf(expectedType), reflect.TypeOf(result))
}
func TestCoerceTypeJSONMarshalerToByteArray(t *testing.T) {
	myData := models.Event{
		Device: "deviceId",
	}
	var expectedType []byte
	result, err := CoerceType(myData)
	assert.NoError(t, err)
	assert.IsType(t, reflect.TypeOf(expectedType), reflect.TypeOf(result))
}
func TestCoerceTypeNotSupportedToByteArray(t *testing.T) {
	myData := 25
	var expectedType []byte
	result, err := CoerceType(myData)
	assert.Error(t, err)
	assert.IsType(t, reflect.TypeOf(expectedType), reflect.TypeOf(result))
}
