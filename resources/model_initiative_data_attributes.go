/*
test

example

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the InitiativeDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InitiativeDataAttributes{}

// InitiativeDataAttributes struct for InitiativeDataAttributes
type InitiativeDataAttributes struct {
	// name of initiative
	Name string `json:"name"`
	// description of initiative
	Desc string `json:"desc"`
	// goal of initiative
	Goal string `json:"goal"`
	// verified status
	Verified bool `json:"verified"`
	// location of initiative
	Location *string `json:"location,omitempty"`
	// status of initiative
	Status string `json:"status"`
	// likes of initiative
	Likes int64 `json:"likes"`
	// reposts of initiative
	Reposts int64 `json:"reposts"`
	// reports of initiative
	Reports int64 `json:"reports"`
	// Initiative created at
	CreatedAt time.Time `json:"created_at"`
	// Initiative updated at
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Initiative closed at
	ClosedAt *time.Time `json:"closed_at,omitempty"`
}

type _InitiativeDataAttributes InitiativeDataAttributes

// NewInitiativeDataAttributes instantiates a new InitiativeDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitiativeDataAttributes(name string, desc string, goal string, verified bool, status string, likes int64, reposts int64, reports int64, createdAt time.Time) *InitiativeDataAttributes {
	this := InitiativeDataAttributes{}
	this.Name = name
	this.Desc = desc
	this.Goal = goal
	this.Verified = verified
	this.Status = status
	this.Likes = likes
	this.Reposts = reposts
	this.Reports = reports
	this.CreatedAt = createdAt
	return &this
}

// NewInitiativeDataAttributesWithDefaults instantiates a new InitiativeDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitiativeDataAttributesWithDefaults() *InitiativeDataAttributes {
	this := InitiativeDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *InitiativeDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InitiativeDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InitiativeDataAttributes) SetName(v string) {
	o.Name = v
}

// GetDesc returns the Desc field value
func (o *InitiativeDataAttributes) GetDesc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Desc
}

// GetDescOk returns a tuple with the Desc field value
// and a boolean to check if the value has been set.
func (o *InitiativeDataAttributes) GetDescOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Desc, true
}

// SetDesc sets field value
func (o *InitiativeDataAttributes) SetDesc(v string) {
	o.Desc = v
}

// GetGoal returns the Goal field value
func (o *InitiativeDataAttributes) GetGoal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Goal
}

// GetGoalOk returns a tuple with the Goal field value
// and a boolean to check if the value has been set.
func (o *InitiativeDataAttributes) GetGoalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Goal, true
}

// SetGoal sets field value
func (o *InitiativeDataAttributes) SetGoal(v string) {
	o.Goal = v
}

// GetVerified returns the Verified field value
func (o *InitiativeDataAttributes) GetVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value
// and a boolean to check if the value has been set.
func (o *InitiativeDataAttributes) GetVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verified, true
}

// SetVerified sets field value
func (o *InitiativeDataAttributes) SetVerified(v bool) {
	o.Verified = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *InitiativeDataAttributes) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitiativeDataAttributes) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *InitiativeDataAttributes) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *InitiativeDataAttributes) SetLocation(v string) {
	o.Location = &v
}

// GetStatus returns the Status field value
func (o *InitiativeDataAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InitiativeDataAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InitiativeDataAttributes) SetStatus(v string) {
	o.Status = v
}

// GetLikes returns the Likes field value
func (o *InitiativeDataAttributes) GetLikes() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Likes
}

// GetLikesOk returns a tuple with the Likes field value
// and a boolean to check if the value has been set.
func (o *InitiativeDataAttributes) GetLikesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Likes, true
}

// SetLikes sets field value
func (o *InitiativeDataAttributes) SetLikes(v int64) {
	o.Likes = v
}

// GetReposts returns the Reposts field value
func (o *InitiativeDataAttributes) GetReposts() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Reposts
}

// GetRepostsOk returns a tuple with the Reposts field value
// and a boolean to check if the value has been set.
func (o *InitiativeDataAttributes) GetRepostsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reposts, true
}

// SetReposts sets field value
func (o *InitiativeDataAttributes) SetReposts(v int64) {
	o.Reposts = v
}

// GetReports returns the Reports field value
func (o *InitiativeDataAttributes) GetReports() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Reports
}

// GetReportsOk returns a tuple with the Reports field value
// and a boolean to check if the value has been set.
func (o *InitiativeDataAttributes) GetReportsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reports, true
}

// SetReports sets field value
func (o *InitiativeDataAttributes) SetReports(v int64) {
	o.Reports = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *InitiativeDataAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *InitiativeDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *InitiativeDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *InitiativeDataAttributes) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitiativeDataAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InitiativeDataAttributes) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *InitiativeDataAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetClosedAt returns the ClosedAt field value if set, zero value otherwise.
func (o *InitiativeDataAttributes) GetClosedAt() time.Time {
	if o == nil || IsNil(o.ClosedAt) {
		var ret time.Time
		return ret
	}
	return *o.ClosedAt
}

// GetClosedAtOk returns a tuple with the ClosedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitiativeDataAttributes) GetClosedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ClosedAt) {
		return nil, false
	}
	return o.ClosedAt, true
}

// HasClosedAt returns a boolean if a field has been set.
func (o *InitiativeDataAttributes) HasClosedAt() bool {
	if o != nil && !IsNil(o.ClosedAt) {
		return true
	}

	return false
}

// SetClosedAt gets a reference to the given time.Time and assigns it to the ClosedAt field.
func (o *InitiativeDataAttributes) SetClosedAt(v time.Time) {
	o.ClosedAt = &v
}

func (o InitiativeDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InitiativeDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["desc"] = o.Desc
	toSerialize["goal"] = o.Goal
	toSerialize["verified"] = o.Verified
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	toSerialize["status"] = o.Status
	toSerialize["likes"] = o.Likes
	toSerialize["reposts"] = o.Reposts
	toSerialize["reports"] = o.Reports
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.ClosedAt) {
		toSerialize["closed_at"] = o.ClosedAt
	}
	return toSerialize, nil
}

func (o *InitiativeDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"desc",
		"goal",
		"verified",
		"status",
		"likes",
		"reposts",
		"reports",
		"created_at",
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

	varInitiativeDataAttributes := _InitiativeDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInitiativeDataAttributes)

	if err != nil {
		return err
	}

	*o = InitiativeDataAttributes(varInitiativeDataAttributes)

	return err
}

type NullableInitiativeDataAttributes struct {
	value *InitiativeDataAttributes
	isSet bool
}

func (v NullableInitiativeDataAttributes) Get() *InitiativeDataAttributes {
	return v.value
}

func (v *NullableInitiativeDataAttributes) Set(val *InitiativeDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableInitiativeDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableInitiativeDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitiativeDataAttributes(val *InitiativeDataAttributes) *NullableInitiativeDataAttributes {
	return &NullableInitiativeDataAttributes{value: val, isSet: true}
}

func (v NullableInitiativeDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitiativeDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


