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

// StorageBaseArrayDiskAllOf Definition of the list of properties defined in 'storage.BaseArrayDisk', excluding properties defined in parent classes.
type StorageBaseArrayDiskAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Disk name available in storage array.
	Name *string `json:"Name,omitempty"`
	// Storage disk part number.
	PartNumber *string `json:"PartNumber,omitempty"`
	// Storage protocol used in disk for communication. Possible values are SAS, SATA and NVMe. * `Unknown` - Disk protocol is unknown. * `SAS` - Serial Attached SCSI protocol (SAS) used in disk. * `NVMe` - Non-volatile memory express (NVMe) protocol used in disk. * `SATA` - Serial Advanced Technology Attachment (SATA) used in disk.
	Protocol *string `json:"Protocol,omitempty"`
	// Disk speed for read or write operation measured in rpm.
	Speed *int64 `json:"Speed,omitempty"`
	// Storage disk health status. * `Unknown` - Component status is not available. * `Ok` - Component is healthy and no issues found. * `Degraded` - Functioning, but not at full capability due to a non-fatal failure. * `Critical` - Not functioning or requiring immediate attention. * `Offline` - Component is installed, but powered off. * `Identifying` - Component is in initialization process. * `NotAvailable` - Component is not installed or not available. * `Updating` - Software update is in progress. * `Unrecognized` - Component is not recognized. It may be because the component is not installed properly or it is not supported.
	Status             *string                     `json:"Status,omitempty"`
	StorageUtilization NullableStorageBaseCapacity `json:"StorageUtilization,omitempty"`
	// Storage disk type - it can be SSD, HDD, NVRAM. * `Unknown` - Default unknown disk type. * `SSD` - Storage disk with Solid state disk. * `HDD` - Storage disk with Hard disk drive. * `NVRAM` - Storage disk with Non-volatile random-access memory type. * `SATA` - Disk drive implementation with Serial Advanced Technology Attachment (SATA). * `BSAS` - Bridged SAS-SATA disks with added hardware to enable them to be plugged into a SAS-connected storage shelf. * `FC` - Storage disk with Fiber Channel. * `FSAS` - Near Line SAS. NL-SAS drives are enterprise SATA drives with a SAS interface, head, media, androtational speed of traditional enterprise-class SATA drives with the fully capable SAS interfacetypical for classic SAS drives. * `LUN` - Logical Unit Number refers to a logical disk. * `MSATA` - SATA disk in multi-disk carrier storage shelf. * `SAS` - Storage disk with serial attached SCSI. * `VMDISK` - Virtual machine Data Disk.
	Type *string `json:"Type,omitempty"`
	// Storage disk version number.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseArrayDiskAllOf StorageBaseArrayDiskAllOf

// NewStorageBaseArrayDiskAllOf instantiates a new StorageBaseArrayDiskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseArrayDiskAllOf(classId string, objectType string) *StorageBaseArrayDiskAllOf {
	this := StorageBaseArrayDiskAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageBaseArrayDiskAllOfWithDefaults instantiates a new StorageBaseArrayDiskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseArrayDiskAllOfWithDefaults() *StorageBaseArrayDiskAllOf {
	this := StorageBaseArrayDiskAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseArrayDiskAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseArrayDiskAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseArrayDiskAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseArrayDiskAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseArrayDiskAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseArrayDiskAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBaseArrayDiskAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseArrayDiskAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBaseArrayDiskAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBaseArrayDiskAllOf) SetName(v string) {
	o.Name = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *StorageBaseArrayDiskAllOf) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseArrayDiskAllOf) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *StorageBaseArrayDiskAllOf) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *StorageBaseArrayDiskAllOf) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *StorageBaseArrayDiskAllOf) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseArrayDiskAllOf) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *StorageBaseArrayDiskAllOf) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *StorageBaseArrayDiskAllOf) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *StorageBaseArrayDiskAllOf) GetSpeed() int64 {
	if o == nil || o.Speed == nil {
		var ret int64
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseArrayDiskAllOf) GetSpeedOk() (*int64, bool) {
	if o == nil || o.Speed == nil {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *StorageBaseArrayDiskAllOf) HasSpeed() bool {
	if o != nil && o.Speed != nil {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int64 and assigns it to the Speed field.
func (o *StorageBaseArrayDiskAllOf) SetSpeed(v int64) {
	o.Speed = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StorageBaseArrayDiskAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseArrayDiskAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StorageBaseArrayDiskAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StorageBaseArrayDiskAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetStorageUtilization returns the StorageUtilization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageBaseArrayDiskAllOf) GetStorageUtilization() StorageBaseCapacity {
	if o == nil || o.StorageUtilization.Get() == nil {
		var ret StorageBaseCapacity
		return ret
	}
	return *o.StorageUtilization.Get()
}

// GetStorageUtilizationOk returns a tuple with the StorageUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageBaseArrayDiskAllOf) GetStorageUtilizationOk() (*StorageBaseCapacity, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageUtilization.Get(), o.StorageUtilization.IsSet()
}

// HasStorageUtilization returns a boolean if a field has been set.
func (o *StorageBaseArrayDiskAllOf) HasStorageUtilization() bool {
	if o != nil && o.StorageUtilization.IsSet() {
		return true
	}

	return false
}

// SetStorageUtilization gets a reference to the given NullableStorageBaseCapacity and assigns it to the StorageUtilization field.
func (o *StorageBaseArrayDiskAllOf) SetStorageUtilization(v StorageBaseCapacity) {
	o.StorageUtilization.Set(&v)
}

// SetStorageUtilizationNil sets the value for StorageUtilization to be an explicit nil
func (o *StorageBaseArrayDiskAllOf) SetStorageUtilizationNil() {
	o.StorageUtilization.Set(nil)
}

// UnsetStorageUtilization ensures that no value is present for StorageUtilization, not even an explicit nil
func (o *StorageBaseArrayDiskAllOf) UnsetStorageUtilization() {
	o.StorageUtilization.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StorageBaseArrayDiskAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseArrayDiskAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StorageBaseArrayDiskAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StorageBaseArrayDiskAllOf) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *StorageBaseArrayDiskAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseArrayDiskAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *StorageBaseArrayDiskAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *StorageBaseArrayDiskAllOf) SetVersion(v string) {
	o.Version = &v
}

func (o StorageBaseArrayDiskAllOf) MarshalJSON() ([]byte, error) {
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
	if o.PartNumber != nil {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.Speed != nil {
		toSerialize["Speed"] = o.Speed
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StorageUtilization.IsSet() {
		toSerialize["StorageUtilization"] = o.StorageUtilization.Get()
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseArrayDiskAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageBaseArrayDiskAllOf := _StorageBaseArrayDiskAllOf{}

	if err = json.Unmarshal(bytes, &varStorageBaseArrayDiskAllOf); err == nil {
		*o = StorageBaseArrayDiskAllOf(varStorageBaseArrayDiskAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "Protocol")
		delete(additionalProperties, "Speed")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StorageUtilization")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageBaseArrayDiskAllOf struct {
	value *StorageBaseArrayDiskAllOf
	isSet bool
}

func (v NullableStorageBaseArrayDiskAllOf) Get() *StorageBaseArrayDiskAllOf {
	return v.value
}

func (v *NullableStorageBaseArrayDiskAllOf) Set(val *StorageBaseArrayDiskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseArrayDiskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseArrayDiskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseArrayDiskAllOf(val *StorageBaseArrayDiskAllOf) *NullableStorageBaseArrayDiskAllOf {
	return &NullableStorageBaseArrayDiskAllOf{value: val, isSet: true}
}

func (v NullableStorageBaseArrayDiskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseArrayDiskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
