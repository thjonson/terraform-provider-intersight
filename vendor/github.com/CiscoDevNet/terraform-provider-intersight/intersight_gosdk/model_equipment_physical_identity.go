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

// EquipmentPhysicalIdentity An Abstract Identity object that uniquely represents a server object under a DR.
type EquipmentPhysicalIdentity struct {
	EquipmentIdentity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType                 string                               `json:"ObjectType"`
	PhysicalDeviceRegistration *AssetDeviceRegistrationRelationship `json:"PhysicalDeviceRegistration,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _EquipmentPhysicalIdentity EquipmentPhysicalIdentity

// NewEquipmentPhysicalIdentity instantiates a new EquipmentPhysicalIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentPhysicalIdentity(classId string, objectType string) *EquipmentPhysicalIdentity {
	this := EquipmentPhysicalIdentity{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentPhysicalIdentityWithDefaults instantiates a new EquipmentPhysicalIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentPhysicalIdentityWithDefaults() *EquipmentPhysicalIdentity {
	this := EquipmentPhysicalIdentity{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentPhysicalIdentity) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentPhysicalIdentity) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentPhysicalIdentity) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentPhysicalIdentity) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentPhysicalIdentity) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentPhysicalIdentity) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPhysicalDeviceRegistration returns the PhysicalDeviceRegistration field value if set, zero value otherwise.
func (o *EquipmentPhysicalIdentity) GetPhysicalDeviceRegistration() AssetDeviceRegistrationRelationship {
	if o == nil || o.PhysicalDeviceRegistration == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.PhysicalDeviceRegistration
}

// GetPhysicalDeviceRegistrationOk returns a tuple with the PhysicalDeviceRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPhysicalIdentity) GetPhysicalDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.PhysicalDeviceRegistration == nil {
		return nil, false
	}
	return o.PhysicalDeviceRegistration, true
}

// HasPhysicalDeviceRegistration returns a boolean if a field has been set.
func (o *EquipmentPhysicalIdentity) HasPhysicalDeviceRegistration() bool {
	if o != nil && o.PhysicalDeviceRegistration != nil {
		return true
	}

	return false
}

// SetPhysicalDeviceRegistration gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the PhysicalDeviceRegistration field.
func (o *EquipmentPhysicalIdentity) SetPhysicalDeviceRegistration(v AssetDeviceRegistrationRelationship) {
	o.PhysicalDeviceRegistration = &v
}

func (o EquipmentPhysicalIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentIdentity, errEquipmentIdentity := json.Marshal(o.EquipmentIdentity)
	if errEquipmentIdentity != nil {
		return []byte{}, errEquipmentIdentity
	}
	errEquipmentIdentity = json.Unmarshal([]byte(serializedEquipmentIdentity), &toSerialize)
	if errEquipmentIdentity != nil {
		return []byte{}, errEquipmentIdentity
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.PhysicalDeviceRegistration != nil {
		toSerialize["PhysicalDeviceRegistration"] = o.PhysicalDeviceRegistration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentPhysicalIdentity) UnmarshalJSON(bytes []byte) (err error) {
	type EquipmentPhysicalIdentityWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType                 string                               `json:"ObjectType"`
		PhysicalDeviceRegistration *AssetDeviceRegistrationRelationship `json:"PhysicalDeviceRegistration,omitempty"`
	}

	varEquipmentPhysicalIdentityWithoutEmbeddedStruct := EquipmentPhysicalIdentityWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEquipmentPhysicalIdentityWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentPhysicalIdentity := _EquipmentPhysicalIdentity{}
		varEquipmentPhysicalIdentity.ClassId = varEquipmentPhysicalIdentityWithoutEmbeddedStruct.ClassId
		varEquipmentPhysicalIdentity.ObjectType = varEquipmentPhysicalIdentityWithoutEmbeddedStruct.ObjectType
		varEquipmentPhysicalIdentity.PhysicalDeviceRegistration = varEquipmentPhysicalIdentityWithoutEmbeddedStruct.PhysicalDeviceRegistration
		*o = EquipmentPhysicalIdentity(varEquipmentPhysicalIdentity)
	} else {
		return err
	}

	varEquipmentPhysicalIdentity := _EquipmentPhysicalIdentity{}

	err = json.Unmarshal(bytes, &varEquipmentPhysicalIdentity)
	if err == nil {
		o.EquipmentIdentity = varEquipmentPhysicalIdentity.EquipmentIdentity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PhysicalDeviceRegistration")

		// remove fields from embedded structs
		reflectEquipmentIdentity := reflect.ValueOf(o.EquipmentIdentity)
		for i := 0; i < reflectEquipmentIdentity.Type().NumField(); i++ {
			t := reflectEquipmentIdentity.Type().Field(i)

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

type NullableEquipmentPhysicalIdentity struct {
	value *EquipmentPhysicalIdentity
	isSet bool
}

func (v NullableEquipmentPhysicalIdentity) Get() *EquipmentPhysicalIdentity {
	return v.value
}

func (v *NullableEquipmentPhysicalIdentity) Set(val *EquipmentPhysicalIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentPhysicalIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentPhysicalIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentPhysicalIdentity(val *EquipmentPhysicalIdentity) *NullableEquipmentPhysicalIdentity {
	return &NullableEquipmentPhysicalIdentity{value: val, isSet: true}
}

func (v NullableEquipmentPhysicalIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentPhysicalIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
