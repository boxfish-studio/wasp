/*
Wasp API

REST API for the Wasp node

API version: 0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
)

// checks if the JSONDict type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JSONDict{}

// JSONDict struct for JSONDict
type JSONDict struct {
	// items
	Items []Item `json:"items"`
}

// NewJSONDict instantiates a new JSONDict object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSONDict(items []Item) *JSONDict {
	this := JSONDict{}
	this.Items = items
	return &this
}

// NewJSONDictWithDefaults instantiates a new JSONDict object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSONDictWithDefaults() *JSONDict {
	this := JSONDict{}
	return &this
}

// GetItems returns the Items field value
func (o *JSONDict) GetItems() []Item {
	if o == nil {
		var ret []Item
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *JSONDict) GetItemsOk() ([]Item, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *JSONDict) SetItems(v []Item) {
	o.Items = v
}

func (o JSONDict) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JSONDict) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableJSONDict struct {
	value *JSONDict
	isSet bool
}

func (v NullableJSONDict) Get() *JSONDict {
	return v.value
}

func (v *NullableJSONDict) Set(val *JSONDict) {
	v.value = val
	v.isSet = true
}

func (v NullableJSONDict) IsSet() bool {
	return v.isSet
}

func (v *NullableJSONDict) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSONDict(val *JSONDict) *NullableJSONDict {
	return &NullableJSONDict{value: val, isSet: true}
}

func (v NullableJSONDict) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSONDict) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


