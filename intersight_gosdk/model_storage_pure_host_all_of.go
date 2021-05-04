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

// StoragePureHostAllOf Definition of the list of properties defined in 'storage.PureHost', excluding properties defined in parent classes.
type StoragePureHostAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of host group where the host belongs to. Empty if host is not part of any HostGroup.
	HostGroupName        *string                                 `json:"HostGroupName,omitempty"`
	Array                *StoragePureArrayRelationship           `json:"Array,omitempty"`
	HostGroup            *StoragePureHostGroupRelationship       `json:"HostGroup,omitempty"`
	ProtectionGroup      *StoragePureProtectionGroupRelationship `json:"ProtectionGroup,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoragePureHostAllOf StoragePureHostAllOf

// NewStoragePureHostAllOf instantiates a new StoragePureHostAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureHostAllOf(classId string, objectType string) *StoragePureHostAllOf {
	this := StoragePureHostAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStoragePureHostAllOfWithDefaults instantiates a new StoragePureHostAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureHostAllOfWithDefaults() *StoragePureHostAllOf {
	this := StoragePureHostAllOf{}
	var classId string = "storage.PureHost"
	this.ClassId = classId
	var objectType string = "storage.PureHost"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StoragePureHostAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StoragePureHostAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StoragePureHostAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StoragePureHostAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StoragePureHostAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StoragePureHostAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHostGroupName returns the HostGroupName field value if set, zero value otherwise.
func (o *StoragePureHostAllOf) GetHostGroupName() string {
	if o == nil || o.HostGroupName == nil {
		var ret string
		return ret
	}
	return *o.HostGroupName
}

// GetHostGroupNameOk returns a tuple with the HostGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostAllOf) GetHostGroupNameOk() (*string, bool) {
	if o == nil || o.HostGroupName == nil {
		return nil, false
	}
	return o.HostGroupName, true
}

// HasHostGroupName returns a boolean if a field has been set.
func (o *StoragePureHostAllOf) HasHostGroupName() bool {
	if o != nil && o.HostGroupName != nil {
		return true
	}

	return false
}

// SetHostGroupName gets a reference to the given string and assigns it to the HostGroupName field.
func (o *StoragePureHostAllOf) SetHostGroupName(v string) {
	o.HostGroupName = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StoragePureHostAllOf) GetArray() StoragePureArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StoragePureArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostAllOf) GetArrayOk() (*StoragePureArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StoragePureHostAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StoragePureArrayRelationship and assigns it to the Array field.
func (o *StoragePureHostAllOf) SetArray(v StoragePureArrayRelationship) {
	o.Array = &v
}

// GetHostGroup returns the HostGroup field value if set, zero value otherwise.
func (o *StoragePureHostAllOf) GetHostGroup() StoragePureHostGroupRelationship {
	if o == nil || o.HostGroup == nil {
		var ret StoragePureHostGroupRelationship
		return ret
	}
	return *o.HostGroup
}

// GetHostGroupOk returns a tuple with the HostGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostAllOf) GetHostGroupOk() (*StoragePureHostGroupRelationship, bool) {
	if o == nil || o.HostGroup == nil {
		return nil, false
	}
	return o.HostGroup, true
}

// HasHostGroup returns a boolean if a field has been set.
func (o *StoragePureHostAllOf) HasHostGroup() bool {
	if o != nil && o.HostGroup != nil {
		return true
	}

	return false
}

// SetHostGroup gets a reference to the given StoragePureHostGroupRelationship and assigns it to the HostGroup field.
func (o *StoragePureHostAllOf) SetHostGroup(v StoragePureHostGroupRelationship) {
	o.HostGroup = &v
}

// GetProtectionGroup returns the ProtectionGroup field value if set, zero value otherwise.
func (o *StoragePureHostAllOf) GetProtectionGroup() StoragePureProtectionGroupRelationship {
	if o == nil || o.ProtectionGroup == nil {
		var ret StoragePureProtectionGroupRelationship
		return ret
	}
	return *o.ProtectionGroup
}

// GetProtectionGroupOk returns a tuple with the ProtectionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostAllOf) GetProtectionGroupOk() (*StoragePureProtectionGroupRelationship, bool) {
	if o == nil || o.ProtectionGroup == nil {
		return nil, false
	}
	return o.ProtectionGroup, true
}

// HasProtectionGroup returns a boolean if a field has been set.
func (o *StoragePureHostAllOf) HasProtectionGroup() bool {
	if o != nil && o.ProtectionGroup != nil {
		return true
	}

	return false
}

// SetProtectionGroup gets a reference to the given StoragePureProtectionGroupRelationship and assigns it to the ProtectionGroup field.
func (o *StoragePureHostAllOf) SetProtectionGroup(v StoragePureProtectionGroupRelationship) {
	o.ProtectionGroup = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePureHostAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePureHostAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePureHostAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StoragePureHostAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.HostGroupName != nil {
		toSerialize["HostGroupName"] = o.HostGroupName
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.HostGroup != nil {
		toSerialize["HostGroup"] = o.HostGroup
	}
	if o.ProtectionGroup != nil {
		toSerialize["ProtectionGroup"] = o.ProtectionGroup
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoragePureHostAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStoragePureHostAllOf := _StoragePureHostAllOf{}

	if err = json.Unmarshal(bytes, &varStoragePureHostAllOf); err == nil {
		*o = StoragePureHostAllOf(varStoragePureHostAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "HostGroupName")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "HostGroup")
		delete(additionalProperties, "ProtectionGroup")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStoragePureHostAllOf struct {
	value *StoragePureHostAllOf
	isSet bool
}

func (v NullableStoragePureHostAllOf) Get() *StoragePureHostAllOf {
	return v.value
}

func (v *NullableStoragePureHostAllOf) Set(val *StoragePureHostAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureHostAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureHostAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureHostAllOf(val *StoragePureHostAllOf) *NullableStoragePureHostAllOf {
	return &NullableStoragePureHostAllOf{value: val, isSet: true}
}

func (v NullableStoragePureHostAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureHostAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
