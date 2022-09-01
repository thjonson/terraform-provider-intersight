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
)

// StoragePhysicalDiskExtensionAllOf Definition of the list of properties defined in 'storage.PhysicalDiskExtension', excluding properties defined in parent classes.
type StoragePhysicalDiskExtensionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The whether disk is bootable or not.
	Bootable *string `json:"Bootable,omitempty"`
	// The distinguished name of the Physical drive.
	DiskDn *string `json:"DiskDn,omitempty"`
	// The storage Enclosure slotId.
	DiskId *int64 `json:"DiskId,omitempty"`
	// The current drive state of disk.
	DiskState *string `json:"DiskState,omitempty"`
	// The current drive state of disk.
	Health               *string                              `json:"Health,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	PhysicalDisk         *StoragePhysicalDiskRelationship     `json:"PhysicalDisk,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	StorageController    *StorageControllerRelationship       `json:"StorageController,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoragePhysicalDiskExtensionAllOf StoragePhysicalDiskExtensionAllOf

// NewStoragePhysicalDiskExtensionAllOf instantiates a new StoragePhysicalDiskExtensionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePhysicalDiskExtensionAllOf(classId string, objectType string) *StoragePhysicalDiskExtensionAllOf {
	this := StoragePhysicalDiskExtensionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStoragePhysicalDiskExtensionAllOfWithDefaults instantiates a new StoragePhysicalDiskExtensionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePhysicalDiskExtensionAllOfWithDefaults() *StoragePhysicalDiskExtensionAllOf {
	this := StoragePhysicalDiskExtensionAllOf{}
	var classId string = "storage.PhysicalDiskExtension"
	this.ClassId = classId
	var objectType string = "storage.PhysicalDiskExtension"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StoragePhysicalDiskExtensionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskExtensionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StoragePhysicalDiskExtensionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StoragePhysicalDiskExtensionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskExtensionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StoragePhysicalDiskExtensionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBootable returns the Bootable field value if set, zero value otherwise.
func (o *StoragePhysicalDiskExtensionAllOf) GetBootable() string {
	if o == nil || o.Bootable == nil {
		var ret string
		return ret
	}
	return *o.Bootable
}

// GetBootableOk returns a tuple with the Bootable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskExtensionAllOf) GetBootableOk() (*string, bool) {
	if o == nil || o.Bootable == nil {
		return nil, false
	}
	return o.Bootable, true
}

// HasBootable returns a boolean if a field has been set.
func (o *StoragePhysicalDiskExtensionAllOf) HasBootable() bool {
	if o != nil && o.Bootable != nil {
		return true
	}

	return false
}

// SetBootable gets a reference to the given string and assigns it to the Bootable field.
func (o *StoragePhysicalDiskExtensionAllOf) SetBootable(v string) {
	o.Bootable = &v
}

// GetDiskDn returns the DiskDn field value if set, zero value otherwise.
func (o *StoragePhysicalDiskExtensionAllOf) GetDiskDn() string {
	if o == nil || o.DiskDn == nil {
		var ret string
		return ret
	}
	return *o.DiskDn
}

// GetDiskDnOk returns a tuple with the DiskDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskExtensionAllOf) GetDiskDnOk() (*string, bool) {
	if o == nil || o.DiskDn == nil {
		return nil, false
	}
	return o.DiskDn, true
}

// HasDiskDn returns a boolean if a field has been set.
func (o *StoragePhysicalDiskExtensionAllOf) HasDiskDn() bool {
	if o != nil && o.DiskDn != nil {
		return true
	}

	return false
}

// SetDiskDn gets a reference to the given string and assigns it to the DiskDn field.
func (o *StoragePhysicalDiskExtensionAllOf) SetDiskDn(v string) {
	o.DiskDn = &v
}

// GetDiskId returns the DiskId field value if set, zero value otherwise.
func (o *StoragePhysicalDiskExtensionAllOf) GetDiskId() int64 {
	if o == nil || o.DiskId == nil {
		var ret int64
		return ret
	}
	return *o.DiskId
}

// GetDiskIdOk returns a tuple with the DiskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskExtensionAllOf) GetDiskIdOk() (*int64, bool) {
	if o == nil || o.DiskId == nil {
		return nil, false
	}
	return o.DiskId, true
}

// HasDiskId returns a boolean if a field has been set.
func (o *StoragePhysicalDiskExtensionAllOf) HasDiskId() bool {
	if o != nil && o.DiskId != nil {
		return true
	}

	return false
}

// SetDiskId gets a reference to the given int64 and assigns it to the DiskId field.
func (o *StoragePhysicalDiskExtensionAllOf) SetDiskId(v int64) {
	o.DiskId = &v
}

// GetDiskState returns the DiskState field value if set, zero value otherwise.
func (o *StoragePhysicalDiskExtensionAllOf) GetDiskState() string {
	if o == nil || o.DiskState == nil {
		var ret string
		return ret
	}
	return *o.DiskState
}

// GetDiskStateOk returns a tuple with the DiskState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskExtensionAllOf) GetDiskStateOk() (*string, bool) {
	if o == nil || o.DiskState == nil {
		return nil, false
	}
	return o.DiskState, true
}

// HasDiskState returns a boolean if a field has been set.
func (o *StoragePhysicalDiskExtensionAllOf) HasDiskState() bool {
	if o != nil && o.DiskState != nil {
		return true
	}

	return false
}

// SetDiskState gets a reference to the given string and assigns it to the DiskState field.
func (o *StoragePhysicalDiskExtensionAllOf) SetDiskState(v string) {
	o.DiskState = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *StoragePhysicalDiskExtensionAllOf) GetHealth() string {
	if o == nil || o.Health == nil {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskExtensionAllOf) GetHealthOk() (*string, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *StoragePhysicalDiskExtensionAllOf) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *StoragePhysicalDiskExtensionAllOf) SetHealth(v string) {
	o.Health = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StoragePhysicalDiskExtensionAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskExtensionAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StoragePhysicalDiskExtensionAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StoragePhysicalDiskExtensionAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetPhysicalDisk returns the PhysicalDisk field value if set, zero value otherwise.
func (o *StoragePhysicalDiskExtensionAllOf) GetPhysicalDisk() StoragePhysicalDiskRelationship {
	if o == nil || o.PhysicalDisk == nil {
		var ret StoragePhysicalDiskRelationship
		return ret
	}
	return *o.PhysicalDisk
}

// GetPhysicalDiskOk returns a tuple with the PhysicalDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskExtensionAllOf) GetPhysicalDiskOk() (*StoragePhysicalDiskRelationship, bool) {
	if o == nil || o.PhysicalDisk == nil {
		return nil, false
	}
	return o.PhysicalDisk, true
}

// HasPhysicalDisk returns a boolean if a field has been set.
func (o *StoragePhysicalDiskExtensionAllOf) HasPhysicalDisk() bool {
	if o != nil && o.PhysicalDisk != nil {
		return true
	}

	return false
}

// SetPhysicalDisk gets a reference to the given StoragePhysicalDiskRelationship and assigns it to the PhysicalDisk field.
func (o *StoragePhysicalDiskExtensionAllOf) SetPhysicalDisk(v StoragePhysicalDiskRelationship) {
	o.PhysicalDisk = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePhysicalDiskExtensionAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskExtensionAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePhysicalDiskExtensionAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePhysicalDiskExtensionAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageController returns the StorageController field value if set, zero value otherwise.
func (o *StoragePhysicalDiskExtensionAllOf) GetStorageController() StorageControllerRelationship {
	if o == nil || o.StorageController == nil {
		var ret StorageControllerRelationship
		return ret
	}
	return *o.StorageController
}

// GetStorageControllerOk returns a tuple with the StorageController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskExtensionAllOf) GetStorageControllerOk() (*StorageControllerRelationship, bool) {
	if o == nil || o.StorageController == nil {
		return nil, false
	}
	return o.StorageController, true
}

// HasStorageController returns a boolean if a field has been set.
func (o *StoragePhysicalDiskExtensionAllOf) HasStorageController() bool {
	if o != nil && o.StorageController != nil {
		return true
	}

	return false
}

// SetStorageController gets a reference to the given StorageControllerRelationship and assigns it to the StorageController field.
func (o *StoragePhysicalDiskExtensionAllOf) SetStorageController(v StorageControllerRelationship) {
	o.StorageController = &v
}

func (o StoragePhysicalDiskExtensionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Bootable != nil {
		toSerialize["Bootable"] = o.Bootable
	}
	if o.DiskDn != nil {
		toSerialize["DiskDn"] = o.DiskDn
	}
	if o.DiskId != nil {
		toSerialize["DiskId"] = o.DiskId
	}
	if o.DiskState != nil {
		toSerialize["DiskState"] = o.DiskState
	}
	if o.Health != nil {
		toSerialize["Health"] = o.Health
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.PhysicalDisk != nil {
		toSerialize["PhysicalDisk"] = o.PhysicalDisk
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageController != nil {
		toSerialize["StorageController"] = o.StorageController
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoragePhysicalDiskExtensionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStoragePhysicalDiskExtensionAllOf := _StoragePhysicalDiskExtensionAllOf{}

	if err = json.Unmarshal(bytes, &varStoragePhysicalDiskExtensionAllOf); err == nil {
		*o = StoragePhysicalDiskExtensionAllOf(varStoragePhysicalDiskExtensionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Bootable")
		delete(additionalProperties, "DiskDn")
		delete(additionalProperties, "DiskId")
		delete(additionalProperties, "DiskState")
		delete(additionalProperties, "Health")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "PhysicalDisk")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageController")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStoragePhysicalDiskExtensionAllOf struct {
	value *StoragePhysicalDiskExtensionAllOf
	isSet bool
}

func (v NullableStoragePhysicalDiskExtensionAllOf) Get() *StoragePhysicalDiskExtensionAllOf {
	return v.value
}

func (v *NullableStoragePhysicalDiskExtensionAllOf) Set(val *StoragePhysicalDiskExtensionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePhysicalDiskExtensionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePhysicalDiskExtensionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePhysicalDiskExtensionAllOf(val *StoragePhysicalDiskExtensionAllOf) *NullableStoragePhysicalDiskExtensionAllOf {
	return &NullableStoragePhysicalDiskExtensionAllOf{value: val, isSet: true}
}

func (v NullableStoragePhysicalDiskExtensionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePhysicalDiskExtensionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
