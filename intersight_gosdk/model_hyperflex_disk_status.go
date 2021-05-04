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

// HyperflexDiskStatus Status details of virtual disk.
type HyperflexDiskStatus struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Percentage of download completed.
	DownloadPercentage *string `json:"DownloadPercentage,omitempty"`
	// Current state of the virtual disk. * `Unknown` - No details available on the disk state. * `Succeeded` - Last operation on the disk has been successful. * `ImportInProgress` - Import operation on the disk is in progress. * `ImportFailed` - Import operation on the disk has failed. * `CloneInProgress` - Disk clone operation on the disk is in progress. * `CloneFailed` - Clone operation on the disk has failed. * `CloneScheduled` - Clone operation on the disk has been scheduled. * `ImportScheduled` - Import operation on the disk has been scheduled. * `Pending` - Submitted operation on the disk is currently pending. * `` - Disk state is not available.
	State *string `json:"State,omitempty"`
	// Identity of the Volume associated with virtual machine disk.
	VolumeHandle *string `json:"VolumeHandle,omitempty"`
	// Name of the Volume associated with virtual machine disk.
	VolumeName           *string `json:"VolumeName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexDiskStatus HyperflexDiskStatus

// NewHyperflexDiskStatus instantiates a new HyperflexDiskStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexDiskStatus(classId string, objectType string) *HyperflexDiskStatus {
	this := HyperflexDiskStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	var state string = "Unknown"
	this.State = &state
	return &this
}

// NewHyperflexDiskStatusWithDefaults instantiates a new HyperflexDiskStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexDiskStatusWithDefaults() *HyperflexDiskStatus {
	this := HyperflexDiskStatus{}
	var classId string = "hyperflex.DiskStatus"
	this.ClassId = classId
	var objectType string = "hyperflex.DiskStatus"
	this.ObjectType = objectType
	var state string = "Unknown"
	this.State = &state
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexDiskStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexDiskStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexDiskStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexDiskStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexDiskStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexDiskStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDownloadPercentage returns the DownloadPercentage field value if set, zero value otherwise.
func (o *HyperflexDiskStatus) GetDownloadPercentage() string {
	if o == nil || o.DownloadPercentage == nil {
		var ret string
		return ret
	}
	return *o.DownloadPercentage
}

// GetDownloadPercentageOk returns a tuple with the DownloadPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDiskStatus) GetDownloadPercentageOk() (*string, bool) {
	if o == nil || o.DownloadPercentage == nil {
		return nil, false
	}
	return o.DownloadPercentage, true
}

// HasDownloadPercentage returns a boolean if a field has been set.
func (o *HyperflexDiskStatus) HasDownloadPercentage() bool {
	if o != nil && o.DownloadPercentage != nil {
		return true
	}

	return false
}

// SetDownloadPercentage gets a reference to the given string and assigns it to the DownloadPercentage field.
func (o *HyperflexDiskStatus) SetDownloadPercentage(v string) {
	o.DownloadPercentage = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *HyperflexDiskStatus) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDiskStatus) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *HyperflexDiskStatus) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *HyperflexDiskStatus) SetState(v string) {
	o.State = &v
}

// GetVolumeHandle returns the VolumeHandle field value if set, zero value otherwise.
func (o *HyperflexDiskStatus) GetVolumeHandle() string {
	if o == nil || o.VolumeHandle == nil {
		var ret string
		return ret
	}
	return *o.VolumeHandle
}

// GetVolumeHandleOk returns a tuple with the VolumeHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDiskStatus) GetVolumeHandleOk() (*string, bool) {
	if o == nil || o.VolumeHandle == nil {
		return nil, false
	}
	return o.VolumeHandle, true
}

// HasVolumeHandle returns a boolean if a field has been set.
func (o *HyperflexDiskStatus) HasVolumeHandle() bool {
	if o != nil && o.VolumeHandle != nil {
		return true
	}

	return false
}

// SetVolumeHandle gets a reference to the given string and assigns it to the VolumeHandle field.
func (o *HyperflexDiskStatus) SetVolumeHandle(v string) {
	o.VolumeHandle = &v
}

// GetVolumeName returns the VolumeName field value if set, zero value otherwise.
func (o *HyperflexDiskStatus) GetVolumeName() string {
	if o == nil || o.VolumeName == nil {
		var ret string
		return ret
	}
	return *o.VolumeName
}

// GetVolumeNameOk returns a tuple with the VolumeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDiskStatus) GetVolumeNameOk() (*string, bool) {
	if o == nil || o.VolumeName == nil {
		return nil, false
	}
	return o.VolumeName, true
}

// HasVolumeName returns a boolean if a field has been set.
func (o *HyperflexDiskStatus) HasVolumeName() bool {
	if o != nil && o.VolumeName != nil {
		return true
	}

	return false
}

// SetVolumeName gets a reference to the given string and assigns it to the VolumeName field.
func (o *HyperflexDiskStatus) SetVolumeName(v string) {
	o.VolumeName = &v
}

func (o HyperflexDiskStatus) MarshalJSON() ([]byte, error) {
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
	if o.DownloadPercentage != nil {
		toSerialize["DownloadPercentage"] = o.DownloadPercentage
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.VolumeHandle != nil {
		toSerialize["VolumeHandle"] = o.VolumeHandle
	}
	if o.VolumeName != nil {
		toSerialize["VolumeName"] = o.VolumeName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexDiskStatus) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexDiskStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Percentage of download completed.
		DownloadPercentage *string `json:"DownloadPercentage,omitempty"`
		// Current state of the virtual disk. * `Unknown` - No details available on the disk state. * `Succeeded` - Last operation on the disk has been successful. * `ImportInProgress` - Import operation on the disk is in progress. * `ImportFailed` - Import operation on the disk has failed. * `CloneInProgress` - Disk clone operation on the disk is in progress. * `CloneFailed` - Clone operation on the disk has failed. * `CloneScheduled` - Clone operation on the disk has been scheduled. * `ImportScheduled` - Import operation on the disk has been scheduled. * `Pending` - Submitted operation on the disk is currently pending. * `` - Disk state is not available.
		State *string `json:"State,omitempty"`
		// Identity of the Volume associated with virtual machine disk.
		VolumeHandle *string `json:"VolumeHandle,omitempty"`
		// Name of the Volume associated with virtual machine disk.
		VolumeName *string `json:"VolumeName,omitempty"`
	}

	varHyperflexDiskStatusWithoutEmbeddedStruct := HyperflexDiskStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexDiskStatusWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexDiskStatus := _HyperflexDiskStatus{}
		varHyperflexDiskStatus.ClassId = varHyperflexDiskStatusWithoutEmbeddedStruct.ClassId
		varHyperflexDiskStatus.ObjectType = varHyperflexDiskStatusWithoutEmbeddedStruct.ObjectType
		varHyperflexDiskStatus.DownloadPercentage = varHyperflexDiskStatusWithoutEmbeddedStruct.DownloadPercentage
		varHyperflexDiskStatus.State = varHyperflexDiskStatusWithoutEmbeddedStruct.State
		varHyperflexDiskStatus.VolumeHandle = varHyperflexDiskStatusWithoutEmbeddedStruct.VolumeHandle
		varHyperflexDiskStatus.VolumeName = varHyperflexDiskStatusWithoutEmbeddedStruct.VolumeName
		*o = HyperflexDiskStatus(varHyperflexDiskStatus)
	} else {
		return err
	}

	varHyperflexDiskStatus := _HyperflexDiskStatus{}

	err = json.Unmarshal(bytes, &varHyperflexDiskStatus)
	if err == nil {
		o.MoBaseComplexType = varHyperflexDiskStatus.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DownloadPercentage")
		delete(additionalProperties, "State")
		delete(additionalProperties, "VolumeHandle")
		delete(additionalProperties, "VolumeName")

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

type NullableHyperflexDiskStatus struct {
	value *HyperflexDiskStatus
	isSet bool
}

func (v NullableHyperflexDiskStatus) Get() *HyperflexDiskStatus {
	return v.value
}

func (v *NullableHyperflexDiskStatus) Set(val *HyperflexDiskStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexDiskStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexDiskStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexDiskStatus(val *HyperflexDiskStatus) *NullableHyperflexDiskStatus {
	return &NullableHyperflexDiskStatus{value: val, isSet: true}
}

func (v NullableHyperflexDiskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexDiskStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
