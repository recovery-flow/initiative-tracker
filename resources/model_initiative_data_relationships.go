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

// checks if the InitiativeDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InitiativeDataRelationships{}

// InitiativeDataRelationships struct for InitiativeDataRelationships
type InitiativeDataRelationships struct {
	Chat LinksDirect `json:"chat"`
	Likes LinksDirect `json:"likes"`
	Reposts LinksDirect `json:"reposts"`
	Reports LinksDirect `json:"reports"`
}

type _InitiativeDataRelationships InitiativeDataRelationships

// NewInitiativeDataRelationships instantiates a new InitiativeDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitiativeDataRelationships(chat LinksDirect, likes LinksDirect, reposts LinksDirect, reports LinksDirect) *InitiativeDataRelationships {
	this := InitiativeDataRelationships{}
	this.Chat = chat
	this.Likes = likes
	this.Reposts = reposts
	this.Reports = reports
	return &this
}

// NewInitiativeDataRelationshipsWithDefaults instantiates a new InitiativeDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitiativeDataRelationshipsWithDefaults() *InitiativeDataRelationships {
	this := InitiativeDataRelationships{}
	return &this
}

// GetChat returns the Chat field value
func (o *InitiativeDataRelationships) GetChat() LinksDirect {
	if o == nil {
		var ret LinksDirect
		return ret
	}

	return o.Chat
}

// GetChatOk returns a tuple with the Chat field value
// and a boolean to check if the value has been set.
func (o *InitiativeDataRelationships) GetChatOk() (*LinksDirect, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chat, true
}

// SetChat sets field value
func (o *InitiativeDataRelationships) SetChat(v LinksDirect) {
	o.Chat = v
}

// GetLikes returns the Likes field value
func (o *InitiativeDataRelationships) GetLikes() LinksDirect {
	if o == nil {
		var ret LinksDirect
		return ret
	}

	return o.Likes
}

// GetLikesOk returns a tuple with the Likes field value
// and a boolean to check if the value has been set.
func (o *InitiativeDataRelationships) GetLikesOk() (*LinksDirect, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Likes, true
}

// SetLikes sets field value
func (o *InitiativeDataRelationships) SetLikes(v LinksDirect) {
	o.Likes = v
}

// GetReposts returns the Reposts field value
func (o *InitiativeDataRelationships) GetReposts() LinksDirect {
	if o == nil {
		var ret LinksDirect
		return ret
	}

	return o.Reposts
}

// GetRepostsOk returns a tuple with the Reposts field value
// and a boolean to check if the value has been set.
func (o *InitiativeDataRelationships) GetRepostsOk() (*LinksDirect, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reposts, true
}

// SetReposts sets field value
func (o *InitiativeDataRelationships) SetReposts(v LinksDirect) {
	o.Reposts = v
}

// GetReports returns the Reports field value
func (o *InitiativeDataRelationships) GetReports() LinksDirect {
	if o == nil {
		var ret LinksDirect
		return ret
	}

	return o.Reports
}

// GetReportsOk returns a tuple with the Reports field value
// and a boolean to check if the value has been set.
func (o *InitiativeDataRelationships) GetReportsOk() (*LinksDirect, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reports, true
}

// SetReports sets field value
func (o *InitiativeDataRelationships) SetReports(v LinksDirect) {
	o.Reports = v
}

func (o InitiativeDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InitiativeDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat"] = o.Chat
	toSerialize["likes"] = o.Likes
	toSerialize["reposts"] = o.Reposts
	toSerialize["reports"] = o.Reports
	return toSerialize, nil
}

func (o *InitiativeDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat",
		"likes",
		"reposts",
		"reports",
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

	varInitiativeDataRelationships := _InitiativeDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInitiativeDataRelationships)

	if err != nil {
		return err
	}

	*o = InitiativeDataRelationships(varInitiativeDataRelationships)

	return err
}

type NullableInitiativeDataRelationships struct {
	value *InitiativeDataRelationships
	isSet bool
}

func (v NullableInitiativeDataRelationships) Get() *InitiativeDataRelationships {
	return v.value
}

func (v *NullableInitiativeDataRelationships) Set(val *InitiativeDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableInitiativeDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableInitiativeDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitiativeDataRelationships(val *InitiativeDataRelationships) *NullableInitiativeDataRelationships {
	return &NullableInitiativeDataRelationships{value: val, isSet: true}
}

func (v NullableInitiativeDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitiativeDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


