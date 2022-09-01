/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-7766
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// ChassisConfigResult The profile configuration (deploy, validation) results with the overall state and detailed result messages.
type ChassisConfigResult struct {
	PolicyAbstractConfigResult
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                         `json:"ObjectType"`
	IomProfile *ChassisIomProfileRelationship `json:"IomProfile,omitempty"`
	Profile    *ChassisProfileRelationship    `json:"Profile,omitempty"`
	// An array of relationships to chassisConfigResultEntry resources.
	ResultEntries        []ChassisConfigResultEntryRelationship `json:"ResultEntries,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChassisConfigResult ChassisConfigResult

// NewChassisConfigResult instantiates a new ChassisConfigResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChassisConfigResult(classId string, objectType string) *ChassisConfigResult {
	this := ChassisConfigResult{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewChassisConfigResultWithDefaults instantiates a new ChassisConfigResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChassisConfigResultWithDefaults() *ChassisConfigResult {
	this := ChassisConfigResult{}
	var classId string = "chassis.ConfigResult"
	this.ClassId = classId
	var objectType string = "chassis.ConfigResult"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ChassisConfigResult) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ChassisConfigResult) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ChassisConfigResult) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ChassisConfigResult) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ChassisConfigResult) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ChassisConfigResult) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIomProfile returns the IomProfile field value if set, zero value otherwise.
func (o *ChassisConfigResult) GetIomProfile() ChassisIomProfileRelationship {
	if o == nil || o.IomProfile == nil {
		var ret ChassisIomProfileRelationship
		return ret
	}
	return *o.IomProfile
}

// GetIomProfileOk returns a tuple with the IomProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisConfigResult) GetIomProfileOk() (*ChassisIomProfileRelationship, bool) {
	if o == nil || o.IomProfile == nil {
		return nil, false
	}
	return o.IomProfile, true
}

// HasIomProfile returns a boolean if a field has been set.
func (o *ChassisConfigResult) HasIomProfile() bool {
	if o != nil && o.IomProfile != nil {
		return true
	}

	return false
}

// SetIomProfile gets a reference to the given ChassisIomProfileRelationship and assigns it to the IomProfile field.
func (o *ChassisConfigResult) SetIomProfile(v ChassisIomProfileRelationship) {
	o.IomProfile = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *ChassisConfigResult) GetProfile() ChassisProfileRelationship {
	if o == nil || o.Profile == nil {
		var ret ChassisProfileRelationship
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisConfigResult) GetProfileOk() (*ChassisProfileRelationship, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *ChassisConfigResult) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given ChassisProfileRelationship and assigns it to the Profile field.
func (o *ChassisConfigResult) SetProfile(v ChassisProfileRelationship) {
	o.Profile = &v
}

// GetResultEntries returns the ResultEntries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisConfigResult) GetResultEntries() []ChassisConfigResultEntryRelationship {
	if o == nil {
		var ret []ChassisConfigResultEntryRelationship
		return ret
	}
	return o.ResultEntries
}

// GetResultEntriesOk returns a tuple with the ResultEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisConfigResult) GetResultEntriesOk() ([]ChassisConfigResultEntryRelationship, bool) {
	if o == nil || o.ResultEntries == nil {
		return nil, false
	}
	return o.ResultEntries, true
}

// HasResultEntries returns a boolean if a field has been set.
func (o *ChassisConfigResult) HasResultEntries() bool {
	if o != nil && o.ResultEntries != nil {
		return true
	}

	return false
}

// SetResultEntries gets a reference to the given []ChassisConfigResultEntryRelationship and assigns it to the ResultEntries field.
func (o *ChassisConfigResult) SetResultEntries(v []ChassisConfigResultEntryRelationship) {
	o.ResultEntries = v
}

func (o ChassisConfigResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigResult, errPolicyAbstractConfigResult := json.Marshal(o.PolicyAbstractConfigResult)
	if errPolicyAbstractConfigResult != nil {
		return []byte{}, errPolicyAbstractConfigResult
	}
	errPolicyAbstractConfigResult = json.Unmarshal([]byte(serializedPolicyAbstractConfigResult), &toSerialize)
	if errPolicyAbstractConfigResult != nil {
		return []byte{}, errPolicyAbstractConfigResult
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IomProfile != nil {
		toSerialize["IomProfile"] = o.IomProfile
	}
	if o.Profile != nil {
		toSerialize["Profile"] = o.Profile
	}
	if o.ResultEntries != nil {
		toSerialize["ResultEntries"] = o.ResultEntries
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChassisConfigResult) UnmarshalJSON(bytes []byte) (err error) {
	type ChassisConfigResultWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                         `json:"ObjectType"`
		IomProfile *ChassisIomProfileRelationship `json:"IomProfile,omitempty"`
		Profile    *ChassisProfileRelationship    `json:"Profile,omitempty"`
		// An array of relationships to chassisConfigResultEntry resources.
		ResultEntries []ChassisConfigResultEntryRelationship `json:"ResultEntries,omitempty"`
	}

	varChassisConfigResultWithoutEmbeddedStruct := ChassisConfigResultWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varChassisConfigResultWithoutEmbeddedStruct)
	if err == nil {
		varChassisConfigResult := _ChassisConfigResult{}
		varChassisConfigResult.ClassId = varChassisConfigResultWithoutEmbeddedStruct.ClassId
		varChassisConfigResult.ObjectType = varChassisConfigResultWithoutEmbeddedStruct.ObjectType
		varChassisConfigResult.IomProfile = varChassisConfigResultWithoutEmbeddedStruct.IomProfile
		varChassisConfigResult.Profile = varChassisConfigResultWithoutEmbeddedStruct.Profile
		varChassisConfigResult.ResultEntries = varChassisConfigResultWithoutEmbeddedStruct.ResultEntries
		*o = ChassisConfigResult(varChassisConfigResult)
	} else {
		return err
	}

	varChassisConfigResult := _ChassisConfigResult{}

	err = json.Unmarshal(bytes, &varChassisConfigResult)
	if err == nil {
		o.PolicyAbstractConfigResult = varChassisConfigResult.PolicyAbstractConfigResult
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IomProfile")
		delete(additionalProperties, "Profile")
		delete(additionalProperties, "ResultEntries")

		// remove fields from embedded structs
		reflectPolicyAbstractConfigResult := reflect.ValueOf(o.PolicyAbstractConfigResult)
		for i := 0; i < reflectPolicyAbstractConfigResult.Type().NumField(); i++ {
			t := reflectPolicyAbstractConfigResult.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChassisConfigResult struct {
	value *ChassisConfigResult
	isSet bool
}

func (v NullableChassisConfigResult) Get() *ChassisConfigResult {
	return v.value
}

func (v *NullableChassisConfigResult) Set(val *ChassisConfigResult) {
	v.value = val
	v.isSet = true
}

func (v NullableChassisConfigResult) IsSet() bool {
	return v.isSet
}

func (v *NullableChassisConfigResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChassisConfigResult(val *ChassisConfigResult) *NullableChassisConfigResult {
	return &NullableChassisConfigResult{value: val, isSet: true}
}

func (v NullableChassisConfigResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChassisConfigResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
