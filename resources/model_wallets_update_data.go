/*
test

example

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the WalletsUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletsUpdateData{}

// WalletsUpdateData struct for WalletsUpdateData
type WalletsUpdateData struct {
	Type string `json:"type"`
	Attributes Wallets `json:"attributes"`
}

type _WalletsUpdateData WalletsUpdateData

// NewWalletsUpdateData instantiates a new WalletsUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletsUpdateData(type_ string, attributes Wallets) *WalletsUpdateData {
	this := WalletsUpdateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewWalletsUpdateDataWithDefaults instantiates a new WalletsUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletsUpdateDataWithDefaults() *WalletsUpdateData {
	this := WalletsUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *WalletsUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WalletsUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WalletsUpdateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *WalletsUpdateData) GetAttributes() Wallets {
	if o == nil {
		var ret Wallets
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *WalletsUpdateData) GetAttributesOk() (*Wallets, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *WalletsUpdateData) SetAttributes(v Wallets) {
	o.Attributes = v
}

func (o WalletsUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletsUpdateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *WalletsUpdateData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"attributes",
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

	varWalletsUpdateData := _WalletsUpdateData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWalletsUpdateData)

	if err != nil {
		return err
	}

	*o = WalletsUpdateData(varWalletsUpdateData)

	return err
}

type NullableWalletsUpdateData struct {
	value *WalletsUpdateData
	isSet bool
}

func (v NullableWalletsUpdateData) Get() *WalletsUpdateData {
	return v.value
}

func (v *NullableWalletsUpdateData) Set(val *WalletsUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletsUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletsUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletsUpdateData(val *WalletsUpdateData) *NullableWalletsUpdateData {
	return &NullableWalletsUpdateData{value: val, isSet: true}
}

func (v NullableWalletsUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletsUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


