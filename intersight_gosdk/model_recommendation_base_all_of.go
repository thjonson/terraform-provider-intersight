/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4663
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"time"
)

// RecommendationBaseAllOf Definition of the list of properties defined in 'recommendation.Base', excluding properties defined in parent classes.
type RecommendationBaseAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The time when the recommendation was last updated.
	LastUpdatedTime *time.Time `json:"LastUpdatedTime,omitempty"`
	// The name of the recommendation.
	Name *string `json:"Name,omitempty"`
	// Indicates if the recommendation requirement is met by the existing setup by adding hardware components to it or it needs expansion. For example if the recommendation is to add 16 disks to a HyperFlex cluster but the cluster is already alost full and only 8 disks can be added, then this property is set to false.
	RequirementMet       *bool `json:"RequirementMet,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecommendationBaseAllOf RecommendationBaseAllOf

// NewRecommendationBaseAllOf instantiates a new RecommendationBaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationBaseAllOf(classId string, objectType string) *RecommendationBaseAllOf {
	this := RecommendationBaseAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewRecommendationBaseAllOfWithDefaults instantiates a new RecommendationBaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationBaseAllOfWithDefaults() *RecommendationBaseAllOf {
	this := RecommendationBaseAllOf{}
	var classId string = "recommendation.CapacityRunway"
	this.ClassId = classId
	var objectType string = "recommendation.CapacityRunway"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecommendationBaseAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecommendationBaseAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecommendationBaseAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecommendationBaseAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecommendationBaseAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecommendationBaseAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *RecommendationBaseAllOf) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationBaseAllOf) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *RecommendationBaseAllOf) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *RecommendationBaseAllOf) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RecommendationBaseAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationBaseAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RecommendationBaseAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RecommendationBaseAllOf) SetName(v string) {
	o.Name = &v
}

// GetRequirementMet returns the RequirementMet field value if set, zero value otherwise.
func (o *RecommendationBaseAllOf) GetRequirementMet() bool {
	if o == nil || o.RequirementMet == nil {
		var ret bool
		return ret
	}
	return *o.RequirementMet
}

// GetRequirementMetOk returns a tuple with the RequirementMet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationBaseAllOf) GetRequirementMetOk() (*bool, bool) {
	if o == nil || o.RequirementMet == nil {
		return nil, false
	}
	return o.RequirementMet, true
}

// HasRequirementMet returns a boolean if a field has been set.
func (o *RecommendationBaseAllOf) HasRequirementMet() bool {
	if o != nil && o.RequirementMet != nil {
		return true
	}

	return false
}

// SetRequirementMet gets a reference to the given bool and assigns it to the RequirementMet field.
func (o *RecommendationBaseAllOf) SetRequirementMet(v bool) {
	o.RequirementMet = &v
}

func (o RecommendationBaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.LastUpdatedTime != nil {
		toSerialize["LastUpdatedTime"] = o.LastUpdatedTime
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.RequirementMet != nil {
		toSerialize["RequirementMet"] = o.RequirementMet
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecommendationBaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varRecommendationBaseAllOf := _RecommendationBaseAllOf{}

	if err = json.Unmarshal(bytes, &varRecommendationBaseAllOf); err == nil {
		*o = RecommendationBaseAllOf(varRecommendationBaseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LastUpdatedTime")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "RequirementMet")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecommendationBaseAllOf struct {
	value *RecommendationBaseAllOf
	isSet bool
}

func (v NullableRecommendationBaseAllOf) Get() *RecommendationBaseAllOf {
	return v.value
}

func (v *NullableRecommendationBaseAllOf) Set(val *RecommendationBaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationBaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationBaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationBaseAllOf(val *RecommendationBaseAllOf) *NullableRecommendationBaseAllOf {
	return &NullableRecommendationBaseAllOf{value: val, isSet: true}
}

func (v NullableRecommendationBaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationBaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
