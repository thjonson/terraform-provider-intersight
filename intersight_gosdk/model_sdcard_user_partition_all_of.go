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
)

// SdcardUserPartitionAllOf Definition of the list of properties defined in 'sdcard.UserPartition', excluding properties defined in parent classes.
type SdcardUserPartitionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of virtual drive for user partition.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SdcardUserPartitionAllOf SdcardUserPartitionAllOf

// NewSdcardUserPartitionAllOf instantiates a new SdcardUserPartitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdcardUserPartitionAllOf(classId string, objectType string) *SdcardUserPartitionAllOf {
	this := SdcardUserPartitionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var name string = "UserPartition"
	this.Name = &name
	return &this
}

// NewSdcardUserPartitionAllOfWithDefaults instantiates a new SdcardUserPartitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdcardUserPartitionAllOfWithDefaults() *SdcardUserPartitionAllOf {
	this := SdcardUserPartitionAllOf{}
	var classId string = "sdcard.UserPartition"
	this.ClassId = classId
	var objectType string = "sdcard.UserPartition"
	this.ObjectType = objectType
	var name string = "UserPartition"
	this.Name = &name
	return &this
}

// GetClassId returns the ClassId field value
func (o *SdcardUserPartitionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SdcardUserPartitionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SdcardUserPartitionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SdcardUserPartitionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SdcardUserPartitionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SdcardUserPartitionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SdcardUserPartitionAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdcardUserPartitionAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SdcardUserPartitionAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SdcardUserPartitionAllOf) SetName(v string) {
	o.Name = &v
}

func (o SdcardUserPartitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SdcardUserPartitionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSdcardUserPartitionAllOf := _SdcardUserPartitionAllOf{}

	if err = json.Unmarshal(bytes, &varSdcardUserPartitionAllOf); err == nil {
		*o = SdcardUserPartitionAllOf(varSdcardUserPartitionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSdcardUserPartitionAllOf struct {
	value *SdcardUserPartitionAllOf
	isSet bool
}

func (v NullableSdcardUserPartitionAllOf) Get() *SdcardUserPartitionAllOf {
	return v.value
}

func (v *NullableSdcardUserPartitionAllOf) Set(val *SdcardUserPartitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSdcardUserPartitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSdcardUserPartitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdcardUserPartitionAllOf(val *SdcardUserPartitionAllOf) *NullableSdcardUserPartitionAllOf {
	return &NullableSdcardUserPartitionAllOf{value: val, isSet: true}
}

func (v NullableSdcardUserPartitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdcardUserPartitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
