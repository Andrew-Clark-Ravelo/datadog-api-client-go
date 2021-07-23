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

// RelationshipToUser Relationship to user.
type RelationshipToUser struct {
	Data RelationshipToUserData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewRelationshipToUser instantiates a new RelationshipToUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipToUser(data RelationshipToUserData) *RelationshipToUser {
	this := RelationshipToUser{}
	this.Data = data
	return &this
}

// NewRelationshipToUserWithDefaults instantiates a new RelationshipToUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipToUserWithDefaults() *RelationshipToUser {
	this := RelationshipToUser{}
	return &this
}

// GetData returns the Data field value
func (o *RelationshipToUser) GetData() RelationshipToUserData {
	if o == nil {
		var ret RelationshipToUserData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RelationshipToUser) GetDataOk() (*RelationshipToUserData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *RelationshipToUser) SetData(v RelationshipToUserData) {
	o.Data = v
}

func (o RelationshipToUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

func (o *RelationshipToUser) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Data *RelationshipToUserData `json:"data"`
	}{}
	all := struct {
		Data RelationshipToUserData `json:"data"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Data == nil {
		return fmt.Errorf("Required field data missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Data = all.Data
	return nil
}

type NullableRelationshipToUser struct {
	value *RelationshipToUser
	isSet bool
}

func (v NullableRelationshipToUser) Get() *RelationshipToUser {
	return v.value
}

func (v *NullableRelationshipToUser) Set(val *RelationshipToUser) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationshipToUser) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationshipToUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationshipToUser(val *RelationshipToUser) *NullableRelationshipToUser {
	return &NullableRelationshipToUser{value: val, isSet: true}
}

func (v NullableRelationshipToUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationshipToUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
