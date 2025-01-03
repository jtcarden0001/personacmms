/*
PersonaCMMS API

This is the Personal Computer Maintenance Management System REST API.

API version: 1.0
Contact: greenrivercodelabs@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TypesTool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TypesTool{}

// TypesTool struct for TypesTool
type TypesTool struct {
	Title string `json:"title"`
}

type _TypesTool TypesTool

// NewTypesTool instantiates a new TypesTool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypesTool(title string) *TypesTool {
	this := TypesTool{}
	this.Title = title
	return &this
}

// NewTypesToolWithDefaults instantiates a new TypesTool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypesToolWithDefaults() *TypesTool {
	this := TypesTool{}
	return &this
}

// GetTitle returns the Title field value
func (o *TypesTool) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *TypesTool) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *TypesTool) SetTitle(v string) {
	o.Title = v
}

func (o TypesTool) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TypesTool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	return toSerialize, nil
}

func (o *TypesTool) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTypesTool := _TypesTool{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTypesTool)

	if err != nil {
		return err
	}

	*o = TypesTool(varTypesTool)

	return err
}

type NullableTypesTool struct {
	value *TypesTool
	isSet bool
}

func (v NullableTypesTool) Get() *TypesTool {
	return v.value
}

func (v *NullableTypesTool) Set(val *TypesTool) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesTool) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesTool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesTool(val *TypesTool) *NullableTypesTool {
	return &NullableTypesTool{value: val, isSet: true}
}

func (v NullableTypesTool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesTool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


