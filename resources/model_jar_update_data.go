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

// checks if the JarUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JarUpdateData{}

// JarUpdateData struct for JarUpdateData
type JarUpdateData struct {
	Type string `json:"type"`
	Attributes JarAttributes `json:"attributes"`
}

type _JarUpdateData JarUpdateData

// NewJarUpdateData instantiates a new JarUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJarUpdateData(type_ string, attributes JarAttributes) *JarUpdateData {
	this := JarUpdateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewJarUpdateDataWithDefaults instantiates a new JarUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJarUpdateDataWithDefaults() *JarUpdateData {
	this := JarUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *JarUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *JarUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *JarUpdateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *JarUpdateData) GetAttributes() JarAttributes {
	if o == nil {
		var ret JarAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *JarUpdateData) GetAttributesOk() (*JarAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *JarUpdateData) SetAttributes(v JarAttributes) {
	o.Attributes = v
}

func (o JarUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JarUpdateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *JarUpdateData) UnmarshalJSON(data []byte) (err error) {
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

	varJarUpdateData := _JarUpdateData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJarUpdateData)

	if err != nil {
		return err
	}

	*o = JarUpdateData(varJarUpdateData)

	return err
}

type NullableJarUpdateData struct {
	value *JarUpdateData
	isSet bool
}

func (v NullableJarUpdateData) Get() *JarUpdateData {
	return v.value
}

func (v *NullableJarUpdateData) Set(val *JarUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableJarUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableJarUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJarUpdateData(val *JarUpdateData) *NullableJarUpdateData {
	return &NullableJarUpdateData{value: val, isSet: true}
}

func (v NullableJarUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJarUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


