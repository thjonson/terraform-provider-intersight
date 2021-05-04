/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-04-28T13:03:38Z.
 *
 * API version: 1.0.9-4267
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// TamBaseAdvisoryDetails Base type for specifying advisory type (Field notice, PSIRT) specific details.
type TamBaseAdvisoryDetails struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Brief description of details specified for an advisory type.
	Description          *string `json:"Description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamBaseAdvisoryDetails TamBaseAdvisoryDetails

// NewTamBaseAdvisoryDetails instantiates a new TamBaseAdvisoryDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamBaseAdvisoryDetails(classId string, objectType string) *TamBaseAdvisoryDetails {
	this := TamBaseAdvisoryDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTamBaseAdvisoryDetailsWithDefaults instantiates a new TamBaseAdvisoryDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamBaseAdvisoryDetailsWithDefaults() *TamBaseAdvisoryDetails {
	this := TamBaseAdvisoryDetails{}
	var classId string = "tam.SecurityAdvisoryDetails"
	this.ClassId = classId
	var objectType string = "tam.SecurityAdvisoryDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamBaseAdvisoryDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamBaseAdvisoryDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamBaseAdvisoryDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TamBaseAdvisoryDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamBaseAdvisoryDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamBaseAdvisoryDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TamBaseAdvisoryDetails) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamBaseAdvisoryDetails) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TamBaseAdvisoryDetails) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TamBaseAdvisoryDetails) SetDescription(v string) {
	o.Description = &v
}

func (o TamBaseAdvisoryDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamBaseAdvisoryDetails) UnmarshalJSON(bytes []byte) (err error) {
	type TamBaseAdvisoryDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Brief description of details specified for an advisory type.
		Description *string `json:"Description,omitempty"`
	}

	varTamBaseAdvisoryDetailsWithoutEmbeddedStruct := TamBaseAdvisoryDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTamBaseAdvisoryDetailsWithoutEmbeddedStruct)
	if err == nil {
		varTamBaseAdvisoryDetails := _TamBaseAdvisoryDetails{}
		varTamBaseAdvisoryDetails.ClassId = varTamBaseAdvisoryDetailsWithoutEmbeddedStruct.ClassId
		varTamBaseAdvisoryDetails.ObjectType = varTamBaseAdvisoryDetailsWithoutEmbeddedStruct.ObjectType
		varTamBaseAdvisoryDetails.Description = varTamBaseAdvisoryDetailsWithoutEmbeddedStruct.Description
		*o = TamBaseAdvisoryDetails(varTamBaseAdvisoryDetails)
	} else {
		return err
	}

	varTamBaseAdvisoryDetails := _TamBaseAdvisoryDetails{}

	err = json.Unmarshal(bytes, &varTamBaseAdvisoryDetails)
	if err == nil {
		o.MoBaseComplexType = varTamBaseAdvisoryDetails.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableTamBaseAdvisoryDetails struct {
	value *TamBaseAdvisoryDetails
	isSet bool
}

func (v NullableTamBaseAdvisoryDetails) Get() *TamBaseAdvisoryDetails {
	return v.value
}

func (v *NullableTamBaseAdvisoryDetails) Set(val *TamBaseAdvisoryDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTamBaseAdvisoryDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTamBaseAdvisoryDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamBaseAdvisoryDetails(val *TamBaseAdvisoryDetails) *NullableTamBaseAdvisoryDetails {
	return &NullableTamBaseAdvisoryDetails{value: val, isSet: true}
}

func (v NullableTamBaseAdvisoryDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamBaseAdvisoryDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
