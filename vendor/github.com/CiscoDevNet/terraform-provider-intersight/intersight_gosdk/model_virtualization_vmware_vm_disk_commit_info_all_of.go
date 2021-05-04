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

// VirtualizationVmwareVmDiskCommitInfoAllOf Definition of the list of properties defined in 'virtualization.VmwareVmDiskCommitInfo', excluding properties defined in parent classes.
type VirtualizationVmwareVmDiskCommitInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Disk committed in bytes on this virtual machine (disk space used up).
	CommittedDisk *int64 `json:"CommittedDisk,omitempty"`
	// Total uncommitted disk space that is available for use (in bytes).
	UnCommittedDisk *int64 `json:"UnCommittedDisk,omitempty"`
	// Total unshared disk space (in bytes).
	UnsharedDisk         *int64 `json:"UnsharedDisk,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareVmDiskCommitInfoAllOf VirtualizationVmwareVmDiskCommitInfoAllOf

// NewVirtualizationVmwareVmDiskCommitInfoAllOf instantiates a new VirtualizationVmwareVmDiskCommitInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareVmDiskCommitInfoAllOf(classId string, objectType string) *VirtualizationVmwareVmDiskCommitInfoAllOf {
	this := VirtualizationVmwareVmDiskCommitInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwareVmDiskCommitInfoAllOfWithDefaults instantiates a new VirtualizationVmwareVmDiskCommitInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareVmDiskCommitInfoAllOfWithDefaults() *VirtualizationVmwareVmDiskCommitInfoAllOf {
	this := VirtualizationVmwareVmDiskCommitInfoAllOf{}
	var classId string = "virtualization.VmwareVmDiskCommitInfo"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareVmDiskCommitInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareVmDiskCommitInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmDiskCommitInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareVmDiskCommitInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareVmDiskCommitInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmDiskCommitInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareVmDiskCommitInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCommittedDisk returns the CommittedDisk field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmDiskCommitInfoAllOf) GetCommittedDisk() int64 {
	if o == nil || o.CommittedDisk == nil {
		var ret int64
		return ret
	}
	return *o.CommittedDisk
}

// GetCommittedDiskOk returns a tuple with the CommittedDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmDiskCommitInfoAllOf) GetCommittedDiskOk() (*int64, bool) {
	if o == nil || o.CommittedDisk == nil {
		return nil, false
	}
	return o.CommittedDisk, true
}

// HasCommittedDisk returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmDiskCommitInfoAllOf) HasCommittedDisk() bool {
	if o != nil && o.CommittedDisk != nil {
		return true
	}

	return false
}

// SetCommittedDisk gets a reference to the given int64 and assigns it to the CommittedDisk field.
func (o *VirtualizationVmwareVmDiskCommitInfoAllOf) SetCommittedDisk(v int64) {
	o.CommittedDisk = &v
}

// GetUnCommittedDisk returns the UnCommittedDisk field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmDiskCommitInfoAllOf) GetUnCommittedDisk() int64 {
	if o == nil || o.UnCommittedDisk == nil {
		var ret int64
		return ret
	}
	return *o.UnCommittedDisk
}

// GetUnCommittedDiskOk returns a tuple with the UnCommittedDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmDiskCommitInfoAllOf) GetUnCommittedDiskOk() (*int64, bool) {
	if o == nil || o.UnCommittedDisk == nil {
		return nil, false
	}
	return o.UnCommittedDisk, true
}

// HasUnCommittedDisk returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmDiskCommitInfoAllOf) HasUnCommittedDisk() bool {
	if o != nil && o.UnCommittedDisk != nil {
		return true
	}

	return false
}

// SetUnCommittedDisk gets a reference to the given int64 and assigns it to the UnCommittedDisk field.
func (o *VirtualizationVmwareVmDiskCommitInfoAllOf) SetUnCommittedDisk(v int64) {
	o.UnCommittedDisk = &v
}

// GetUnsharedDisk returns the UnsharedDisk field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmDiskCommitInfoAllOf) GetUnsharedDisk() int64 {
	if o == nil || o.UnsharedDisk == nil {
		var ret int64
		return ret
	}
	return *o.UnsharedDisk
}

// GetUnsharedDiskOk returns a tuple with the UnsharedDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmDiskCommitInfoAllOf) GetUnsharedDiskOk() (*int64, bool) {
	if o == nil || o.UnsharedDisk == nil {
		return nil, false
	}
	return o.UnsharedDisk, true
}

// HasUnsharedDisk returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmDiskCommitInfoAllOf) HasUnsharedDisk() bool {
	if o != nil && o.UnsharedDisk != nil {
		return true
	}

	return false
}

// SetUnsharedDisk gets a reference to the given int64 and assigns it to the UnsharedDisk field.
func (o *VirtualizationVmwareVmDiskCommitInfoAllOf) SetUnsharedDisk(v int64) {
	o.UnsharedDisk = &v
}

func (o VirtualizationVmwareVmDiskCommitInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CommittedDisk != nil {
		toSerialize["CommittedDisk"] = o.CommittedDisk
	}
	if o.UnCommittedDisk != nil {
		toSerialize["UnCommittedDisk"] = o.UnCommittedDisk
	}
	if o.UnsharedDisk != nil {
		toSerialize["UnsharedDisk"] = o.UnsharedDisk
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareVmDiskCommitInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationVmwareVmDiskCommitInfoAllOf := _VirtualizationVmwareVmDiskCommitInfoAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationVmwareVmDiskCommitInfoAllOf); err == nil {
		*o = VirtualizationVmwareVmDiskCommitInfoAllOf(varVirtualizationVmwareVmDiskCommitInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CommittedDisk")
		delete(additionalProperties, "UnCommittedDisk")
		delete(additionalProperties, "UnsharedDisk")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVmwareVmDiskCommitInfoAllOf struct {
	value *VirtualizationVmwareVmDiskCommitInfoAllOf
	isSet bool
}

func (v NullableVirtualizationVmwareVmDiskCommitInfoAllOf) Get() *VirtualizationVmwareVmDiskCommitInfoAllOf {
	return v.value
}

func (v *NullableVirtualizationVmwareVmDiskCommitInfoAllOf) Set(val *VirtualizationVmwareVmDiskCommitInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareVmDiskCommitInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareVmDiskCommitInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareVmDiskCommitInfoAllOf(val *VirtualizationVmwareVmDiskCommitInfoAllOf) *NullableVirtualizationVmwareVmDiskCommitInfoAllOf {
	return &NullableVirtualizationVmwareVmDiskCommitInfoAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareVmDiskCommitInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareVmDiskCommitInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
