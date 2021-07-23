/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// SyntheticsCheckType Type of assertion to apply in an API test.
type SyntheticsCheckType string

// List of SyntheticsCheckType
const (
	SYNTHETICSCHECKTYPE_EQUALS          SyntheticsCheckType = "equals"
	SYNTHETICSCHECKTYPE_NOT_EQUALS      SyntheticsCheckType = "notEquals"
	SYNTHETICSCHECKTYPE_CONTAINS        SyntheticsCheckType = "contains"
	SYNTHETICSCHECKTYPE_NOT_CONTAINS    SyntheticsCheckType = "notContains"
	SYNTHETICSCHECKTYPE_STARTS_WITH     SyntheticsCheckType = "startsWith"
	SYNTHETICSCHECKTYPE_NOT_STARTS_WITH SyntheticsCheckType = "notStartsWith"
	SYNTHETICSCHECKTYPE_GREATER         SyntheticsCheckType = "greater"
	SYNTHETICSCHECKTYPE_LOWER           SyntheticsCheckType = "lower"
	SYNTHETICSCHECKTYPE_GREATER_EQUALS  SyntheticsCheckType = "greaterEquals"
	SYNTHETICSCHECKTYPE_LOWER_EQUALS    SyntheticsCheckType = "lowerEquals"
	SYNTHETICSCHECKTYPE_MATCH_REGEX     SyntheticsCheckType = "matchRegex"
	SYNTHETICSCHECKTYPE_BETWEEN         SyntheticsCheckType = "between"
	SYNTHETICSCHECKTYPE_IS_EMPTY        SyntheticsCheckType = "isEmpty"
	SYNTHETICSCHECKTYPE_NOT_IS_EMPTY    SyntheticsCheckType = "notIsEmpty"
)

var allowedSyntheticsCheckTypeEnumValues = []SyntheticsCheckType{
	"equals",
	"notEquals",
	"contains",
	"notContains",
	"startsWith",
	"notStartsWith",
	"greater",
	"lower",
	"greaterEquals",
	"lowerEquals",
	"matchRegex",
	"between",
	"isEmpty",
	"notIsEmpty",
}

func (w *SyntheticsCheckType) GetAllowedValues() []SyntheticsCheckType {
	return allowedSyntheticsCheckTypeEnumValues
}

func (v *SyntheticsCheckType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsCheckType(value)
	return nil
}

// NewSyntheticsCheckTypeFromValue returns a pointer to a valid SyntheticsCheckType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSyntheticsCheckTypeFromValue(v string) (*SyntheticsCheckType, error) {
	ev := SyntheticsCheckType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SyntheticsCheckType: valid values are %v", v, allowedSyntheticsCheckTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SyntheticsCheckType) IsValid() bool {
	for _, existing := range allowedSyntheticsCheckTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsCheckType value
func (v SyntheticsCheckType) Ptr() *SyntheticsCheckType {
	return &v
}

type NullableSyntheticsCheckType struct {
	value *SyntheticsCheckType
	isSet bool
}

func (v NullableSyntheticsCheckType) Get() *SyntheticsCheckType {
	return v.value
}

func (v *NullableSyntheticsCheckType) Set(val *SyntheticsCheckType) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsCheckType) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsCheckType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsCheckType(val *SyntheticsCheckType) *NullableSyntheticsCheckType {
	return &NullableSyntheticsCheckType{value: val, isSet: true}
}

func (v NullableSyntheticsCheckType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsCheckType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
