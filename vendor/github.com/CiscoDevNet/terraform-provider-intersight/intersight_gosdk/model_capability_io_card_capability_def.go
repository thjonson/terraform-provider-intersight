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

// CapabilityIoCardCapabilityDef Chassis Iocard module capabilities.
type CapabilityIoCardCapabilityDef struct {
	CapabilityCapability
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Device connector support on Iocard.
	DcSupported          *bool `json:"DcSupported,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityIoCardCapabilityDef CapabilityIoCardCapabilityDef

// NewCapabilityIoCardCapabilityDef instantiates a new CapabilityIoCardCapabilityDef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityIoCardCapabilityDef(classId string, objectType string) *CapabilityIoCardCapabilityDef {
	this := CapabilityIoCardCapabilityDef{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityIoCardCapabilityDefWithDefaults instantiates a new CapabilityIoCardCapabilityDef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityIoCardCapabilityDefWithDefaults() *CapabilityIoCardCapabilityDef {
	this := CapabilityIoCardCapabilityDef{}
	var classId string = "capability.IoCardCapabilityDef"
	this.ClassId = classId
	var objectType string = "capability.IoCardCapabilityDef"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityIoCardCapabilityDef) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityIoCardCapabilityDef) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityIoCardCapabilityDef) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityIoCardCapabilityDef) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityIoCardCapabilityDef) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityIoCardCapabilityDef) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDcSupported returns the DcSupported field value if set, zero value otherwise.
func (o *CapabilityIoCardCapabilityDef) GetDcSupported() bool {
	if o == nil || o.DcSupported == nil {
		var ret bool
		return ret
	}
	return *o.DcSupported
}

// GetDcSupportedOk returns a tuple with the DcSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityIoCardCapabilityDef) GetDcSupportedOk() (*bool, bool) {
	if o == nil || o.DcSupported == nil {
		return nil, false
	}
	return o.DcSupported, true
}

// HasDcSupported returns a boolean if a field has been set.
func (o *CapabilityIoCardCapabilityDef) HasDcSupported() bool {
	if o != nil && o.DcSupported != nil {
		return true
	}

	return false
}

// SetDcSupported gets a reference to the given bool and assigns it to the DcSupported field.
func (o *CapabilityIoCardCapabilityDef) SetDcSupported(v bool) {
	o.DcSupported = &v
}

func (o CapabilityIoCardCapabilityDef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityCapability, errCapabilityCapability := json.Marshal(o.CapabilityCapability)
	if errCapabilityCapability != nil {
		return []byte{}, errCapabilityCapability
	}
	errCapabilityCapability = json.Unmarshal([]byte(serializedCapabilityCapability), &toSerialize)
	if errCapabilityCapability != nil {
		return []byte{}, errCapabilityCapability
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DcSupported != nil {
		toSerialize["DcSupported"] = o.DcSupported
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilityIoCardCapabilityDef) UnmarshalJSON(bytes []byte) (err error) {
	type CapabilityIoCardCapabilityDefWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Device connector support on Iocard.
		DcSupported *bool `json:"DcSupported,omitempty"`
	}

	varCapabilityIoCardCapabilityDefWithoutEmbeddedStruct := CapabilityIoCardCapabilityDefWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCapabilityIoCardCapabilityDefWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityIoCardCapabilityDef := _CapabilityIoCardCapabilityDef{}
		varCapabilityIoCardCapabilityDef.ClassId = varCapabilityIoCardCapabilityDefWithoutEmbeddedStruct.ClassId
		varCapabilityIoCardCapabilityDef.ObjectType = varCapabilityIoCardCapabilityDefWithoutEmbeddedStruct.ObjectType
		varCapabilityIoCardCapabilityDef.DcSupported = varCapabilityIoCardCapabilityDefWithoutEmbeddedStruct.DcSupported
		*o = CapabilityIoCardCapabilityDef(varCapabilityIoCardCapabilityDef)
	} else {
		return err
	}

	varCapabilityIoCardCapabilityDef := _CapabilityIoCardCapabilityDef{}

	err = json.Unmarshal(bytes, &varCapabilityIoCardCapabilityDef)
	if err == nil {
		o.CapabilityCapability = varCapabilityIoCardCapabilityDef.CapabilityCapability
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DcSupported")

		// remove fields from embedded structs
		reflectCapabilityCapability := reflect.ValueOf(o.CapabilityCapability)
		for i := 0; i < reflectCapabilityCapability.Type().NumField(); i++ {
			t := reflectCapabilityCapability.Type().Field(i)

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

type NullableCapabilityIoCardCapabilityDef struct {
	value *CapabilityIoCardCapabilityDef
	isSet bool
}

func (v NullableCapabilityIoCardCapabilityDef) Get() *CapabilityIoCardCapabilityDef {
	return v.value
}

func (v *NullableCapabilityIoCardCapabilityDef) Set(val *CapabilityIoCardCapabilityDef) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityIoCardCapabilityDef) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityIoCardCapabilityDef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityIoCardCapabilityDef(val *CapabilityIoCardCapabilityDef) *NullableCapabilityIoCardCapabilityDef {
	return &NullableCapabilityIoCardCapabilityDef{value: val, isSet: true}
}

func (v NullableCapabilityIoCardCapabilityDef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityIoCardCapabilityDef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
