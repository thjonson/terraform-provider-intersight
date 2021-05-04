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

// TamPsirtSeverity < Severity level of a security advisory. Each advisory type (psirt/field notice etc.) may have a different method of calculating the severity. For e.g. a PSIRT severity may depend upon the corresponding CVSS score where as field notices are always critical in nature.
type TamPsirtSeverity struct {
	TamSeverity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Severity level associated with the security advisory. * `critical` - < If applicable, it may expose users to critical failures and it needs to be addressed immediately. For e.g. a PSIRT advisory with a corresponding CVSS score of above 9.0. * `high` - < If applicable, it may expose the users to critical failure and it needs to be addressed immediately. For e.g. a PSIRT advisory with a corresponding CVSS score between 7.0-8.9. These may be the vulnerabilities that are more difficult to exploit than the ones deemed critical but once exploited, the will cause critical failures. * `medium` - < If applicable, it may expose the users to failure of certain functions. for e.g. a PSIRT advisory with a corresponding CVSS score between 4.0-6.9. These may be the vulnerabilities that are mitigated to a large extent but that may still be exploited in certain restricted cases. * `info` - < If applicable, it may have some minimal impact for certain functions in the user environment. For e.g. a PSIRT advisory with a corresponding CVSS score below 4.0.
	Level                *string `json:"Level,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamPsirtSeverity TamPsirtSeverity

// NewTamPsirtSeverity instantiates a new TamPsirtSeverity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamPsirtSeverity(classId string, objectType string) *TamPsirtSeverity {
	this := TamPsirtSeverity{}
	this.ClassId = classId
	this.ObjectType = objectType
	var level string = "critical"
	this.Level = &level
	return &this
}

// NewTamPsirtSeverityWithDefaults instantiates a new TamPsirtSeverity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamPsirtSeverityWithDefaults() *TamPsirtSeverity {
	this := TamPsirtSeverity{}
	var classId string = "tam.PsirtSeverity"
	this.ClassId = classId
	var objectType string = "tam.PsirtSeverity"
	this.ObjectType = objectType
	var level string = "critical"
	this.Level = &level
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamPsirtSeverity) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamPsirtSeverity) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamPsirtSeverity) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TamPsirtSeverity) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamPsirtSeverity) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamPsirtSeverity) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *TamPsirtSeverity) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamPsirtSeverity) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *TamPsirtSeverity) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *TamPsirtSeverity) SetLevel(v string) {
	o.Level = &v
}

func (o TamPsirtSeverity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedTamSeverity, errTamSeverity := json.Marshal(o.TamSeverity)
	if errTamSeverity != nil {
		return []byte{}, errTamSeverity
	}
	errTamSeverity = json.Unmarshal([]byte(serializedTamSeverity), &toSerialize)
	if errTamSeverity != nil {
		return []byte{}, errTamSeverity
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Level != nil {
		toSerialize["Level"] = o.Level
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamPsirtSeverity) UnmarshalJSON(bytes []byte) (err error) {
	type TamPsirtSeverityWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Severity level associated with the security advisory. * `critical` - < If applicable, it may expose users to critical failures and it needs to be addressed immediately. For e.g. a PSIRT advisory with a corresponding CVSS score of above 9.0. * `high` - < If applicable, it may expose the users to critical failure and it needs to be addressed immediately. For e.g. a PSIRT advisory with a corresponding CVSS score between 7.0-8.9. These may be the vulnerabilities that are more difficult to exploit than the ones deemed critical but once exploited, the will cause critical failures. * `medium` - < If applicable, it may expose the users to failure of certain functions. for e.g. a PSIRT advisory with a corresponding CVSS score between 4.0-6.9. These may be the vulnerabilities that are mitigated to a large extent but that may still be exploited in certain restricted cases. * `info` - < If applicable, it may have some minimal impact for certain functions in the user environment. For e.g. a PSIRT advisory with a corresponding CVSS score below 4.0.
		Level *string `json:"Level,omitempty"`
	}

	varTamPsirtSeverityWithoutEmbeddedStruct := TamPsirtSeverityWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTamPsirtSeverityWithoutEmbeddedStruct)
	if err == nil {
		varTamPsirtSeverity := _TamPsirtSeverity{}
		varTamPsirtSeverity.ClassId = varTamPsirtSeverityWithoutEmbeddedStruct.ClassId
		varTamPsirtSeverity.ObjectType = varTamPsirtSeverityWithoutEmbeddedStruct.ObjectType
		varTamPsirtSeverity.Level = varTamPsirtSeverityWithoutEmbeddedStruct.Level
		*o = TamPsirtSeverity(varTamPsirtSeverity)
	} else {
		return err
	}

	varTamPsirtSeverity := _TamPsirtSeverity{}

	err = json.Unmarshal(bytes, &varTamPsirtSeverity)
	if err == nil {
		o.TamSeverity = varTamPsirtSeverity.TamSeverity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Level")

		// remove fields from embedded structs
		reflectTamSeverity := reflect.ValueOf(o.TamSeverity)
		for i := 0; i < reflectTamSeverity.Type().NumField(); i++ {
			t := reflectTamSeverity.Type().Field(i)

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

type NullableTamPsirtSeverity struct {
	value *TamPsirtSeverity
	isSet bool
}

func (v NullableTamPsirtSeverity) Get() *TamPsirtSeverity {
	return v.value
}

func (v *NullableTamPsirtSeverity) Set(val *TamPsirtSeverity) {
	v.value = val
	v.isSet = true
}

func (v NullableTamPsirtSeverity) IsSet() bool {
	return v.isSet
}

func (v *NullableTamPsirtSeverity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamPsirtSeverity(val *TamPsirtSeverity) *NullableTamPsirtSeverity {
	return &NullableTamPsirtSeverity{value: val, isSet: true}
}

func (v NullableTamPsirtSeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamPsirtSeverity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
