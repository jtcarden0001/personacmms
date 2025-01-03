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

// checks if the TypesCategory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TypesCategory{}

// TypesCategory struct for TypesCategory
type TypesCategory struct {
	Description *string `json:"description,omitempty"`
	Title string `json:"title"`
}

type _TypesCategory TypesCategory

// NewTypesCategory instantiates a new TypesCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypesCategory(title string) *TypesCategory {
	this := TypesCategory{}
	this.Title = title
	return &this
}

// NewTypesCategoryWithDefaults instantiates a new TypesCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypesCategoryWithDefaults() *TypesCategory {
	this := TypesCategory{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TypesCategory) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesCategory) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TypesCategory) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TypesCategory) SetDescription(v string) {
	o.Description = &v
}

// GetTitle returns the Title field value
func (o *TypesCategory) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *TypesCategory) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *TypesCategory) SetTitle(v string) {
	o.Title = v
}

func (o TypesCategory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TypesCategory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["title"] = o.Title
	return toSerialize, nil
}

func (o *TypesCategory) UnmarshalJSON(data []byte) (err error) {
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

	varTypesCategory := _TypesCategory{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTypesCategory)

	if err != nil {
		return err
	}

	*o = TypesCategory(varTypesCategory)

	return err
}

type NullableTypesCategory struct {
	value *TypesCategory
	isSet bool
}

func (v NullableTypesCategory) Get() *TypesCategory {
	return v.value
}

func (v *NullableTypesCategory) Set(val *TypesCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesCategory(val *TypesCategory) *NullableTypesCategory {
	return &NullableTypesCategory{value: val, isSet: true}
}

func (v NullableTypesCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

