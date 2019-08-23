/*
 * (C) Datadog, Inc. 2019
 * All rights reserved
 * Licensed under a 3-clause BSD style license (see LICENSE)
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog_v1

import (
	"encoding/json"
	"errors"
)

type ServiceLevelObjectivesBulkDeletedErrors struct {
	// The ID of the service level objective object associated with this error.
	Id *string `json:"id,omitempty"`

	// The error message
	Message *string `json:"message,omitempty"`

	// The timeframe of the threshold associated with this error or \"all\" if all thresholds are affected.
	Timeframe *string `json:"timeframe,omitempty"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *ServiceLevelObjectivesBulkDeletedErrors) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjectivesBulkDeletedErrors) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceLevelObjectivesBulkDeletedErrors) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceLevelObjectivesBulkDeletedErrors) SetId(v string) {
	o.Id = &v
}

// GetMessage returns the Message field if non-nil, zero value otherwise.
func (o *ServiceLevelObjectivesBulkDeletedErrors) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjectivesBulkDeletedErrors) GetMessageOk() (string, bool) {
	if o == nil || o.Message == nil {
		var ret string
		return ret, false
	}
	return *o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ServiceLevelObjectivesBulkDeletedErrors) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ServiceLevelObjectivesBulkDeletedErrors) SetMessage(v string) {
	o.Message = &v
}

// GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.
func (o *ServiceLevelObjectivesBulkDeletedErrors) GetTimeframe() string {
	if o == nil || o.Timeframe == nil {
		var ret string
		return ret
	}
	return *o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjectivesBulkDeletedErrors) GetTimeframeOk() (string, bool) {
	if o == nil || o.Timeframe == nil {
		var ret string
		return ret, false
	}
	return *o.Timeframe, true
}

// HasTimeframe returns a boolean if a field has been set.
func (o *ServiceLevelObjectivesBulkDeletedErrors) HasTimeframe() bool {
	if o != nil && o.Timeframe != nil {
		return true
	}

	return false
}

// SetTimeframe gets a reference to the given string and assigns it to the Timeframe field.
func (o *ServiceLevelObjectivesBulkDeletedErrors) SetTimeframe(v string) {
	o.Timeframe = &v
}

func (o ServiceLevelObjectivesBulkDeletedErrors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id == nil {
		return nil, errors.New("Id is required and not nullable, but was not set on ServiceLevelObjectivesBulkDeletedErrors")
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Message == nil {
		return nil, errors.New("Message is required and not nullable, but was not set on ServiceLevelObjectivesBulkDeletedErrors")
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Timeframe == nil {
		return nil, errors.New("Timeframe is required and not nullable, but was not set on ServiceLevelObjectivesBulkDeletedErrors")
	}
	if o.Timeframe != nil {
		toSerialize["timeframe"] = o.Timeframe
	}
	return json.Marshal(toSerialize)
}
