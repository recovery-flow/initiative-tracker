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

// checks if the InitiativeData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InitiativeData{}

// InitiativeData struct for InitiativeData
type InitiativeData struct {
	// initiative id
	Id string `json:"id"`
	Type string `json:"type"`
	Attributes InitiativeDataAttributes `json:"attributes"`
	Links Object `json:"links"`
	Relationships InitiativeDataRelationships `json:"relationships"`
}

type _InitiativeData InitiativeData

// NewInitiativeData instantiates a new InitiativeData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitiativeData(id string, type_ string, attributes InitiativeDataAttributes, links Object, relationships InitiativeDataRelationships) *InitiativeData {
	this := InitiativeData{}
	this.Id = id
	this.Type = type_
	this.Attributes = attributes
	this.Links = links
	this.Relationships = relationships
	return &this
}

// NewInitiativeDataWithDefaults instantiates a new InitiativeData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitiativeDataWithDefaults() *InitiativeData {
	this := InitiativeData{}
	return &this
}

// GetId returns the Id field value
func (o *InitiativeData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InitiativeData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InitiativeData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *InitiativeData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InitiativeData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InitiativeData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *InitiativeData) GetAttributes() InitiativeDataAttributes {
	if o == nil {
		var ret InitiativeDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *InitiativeData) GetAttributesOk() (*InitiativeDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *InitiativeData) SetAttributes(v InitiativeDataAttributes) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *InitiativeData) GetLinks() Object {
	if o == nil {
		var ret Object
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *InitiativeData) GetLinksOk() (*Object, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *InitiativeData) SetLinks(v Object) {
	o.Links = v
}

// GetRelationships returns the Relationships field value
func (o *InitiativeData) GetRelationships() InitiativeDataRelationships {
	if o == nil {
		var ret InitiativeDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *InitiativeData) GetRelationshipsOk() (*InitiativeDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *InitiativeData) SetRelationships(v InitiativeDataRelationships) {
	o.Relationships = v
}

func (o InitiativeData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InitiativeData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["links"] = o.Links
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *InitiativeData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"attributes",
		"links",
		"relationships",
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

	varInitiativeData := _InitiativeData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInitiativeData)

	if err != nil {
		return err
	}

	*o = InitiativeData(varInitiativeData)

	return err
}

type NullableInitiativeData struct {
	value *InitiativeData
	isSet bool
}

func (v NullableInitiativeData) Get() *InitiativeData {
	return v.value
}

func (v *NullableInitiativeData) Set(val *InitiativeData) {
	v.value = val
	v.isSet = true
}

func (v NullableInitiativeData) IsSet() bool {
	return v.isSet
}

func (v *NullableInitiativeData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitiativeData(val *InitiativeData) *NullableInitiativeData {
	return &NullableInitiativeData{value: val, isSet: true}
}

func (v NullableInitiativeData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitiativeData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


