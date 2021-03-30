/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-03-27T10:08:12Z.
 *
 * API version: 1.0.9-4136
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"time"
)

// HyperflexDevicePackageDownloadStateAllOf Definition of the list of properties defined in 'hyperflex.DevicePackageDownloadState', excluding properties defined in parent classes.
type HyperflexDevicePackageDownloadStateAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Checksum of HyperFlex health check Debian package installed on the HyperFlex Device.
	Checksum *string `json:"Checksum,omitempty"`
	// HyperFlex Device Name for which the package download state is tracked.
	HxDeviceName *string  `json:"HxDeviceName,omitempty"`
	HxNodes      []string `json:"HxNodes,omitempty"`
	// Timestamp of the last health check Debian package installation on the HyperFlex Device.
	Timestamp            *time.Time                           `json:"Timestamp,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexDevicePackageDownloadStateAllOf HyperflexDevicePackageDownloadStateAllOf

// NewHyperflexDevicePackageDownloadStateAllOf instantiates a new HyperflexDevicePackageDownloadStateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexDevicePackageDownloadStateAllOf(classId string, objectType string) *HyperflexDevicePackageDownloadStateAllOf {
	this := HyperflexDevicePackageDownloadStateAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexDevicePackageDownloadStateAllOfWithDefaults instantiates a new HyperflexDevicePackageDownloadStateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexDevicePackageDownloadStateAllOfWithDefaults() *HyperflexDevicePackageDownloadStateAllOf {
	this := HyperflexDevicePackageDownloadStateAllOf{}
	var classId string = "hyperflex.DevicePackageDownloadState"
	this.ClassId = classId
	var objectType string = "hyperflex.DevicePackageDownloadState"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexDevicePackageDownloadStateAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexDevicePackageDownloadStateAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexDevicePackageDownloadStateAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexDevicePackageDownloadStateAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexDevicePackageDownloadStateAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexDevicePackageDownloadStateAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *HyperflexDevicePackageDownloadStateAllOf) GetChecksum() string {
	if o == nil || o.Checksum == nil {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDevicePackageDownloadStateAllOf) GetChecksumOk() (*string, bool) {
	if o == nil || o.Checksum == nil {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *HyperflexDevicePackageDownloadStateAllOf) HasChecksum() bool {
	if o != nil && o.Checksum != nil {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *HyperflexDevicePackageDownloadStateAllOf) SetChecksum(v string) {
	o.Checksum = &v
}

// GetHxDeviceName returns the HxDeviceName field value if set, zero value otherwise.
func (o *HyperflexDevicePackageDownloadStateAllOf) GetHxDeviceName() string {
	if o == nil || o.HxDeviceName == nil {
		var ret string
		return ret
	}
	return *o.HxDeviceName
}

// GetHxDeviceNameOk returns a tuple with the HxDeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDevicePackageDownloadStateAllOf) GetHxDeviceNameOk() (*string, bool) {
	if o == nil || o.HxDeviceName == nil {
		return nil, false
	}
	return o.HxDeviceName, true
}

// HasHxDeviceName returns a boolean if a field has been set.
func (o *HyperflexDevicePackageDownloadStateAllOf) HasHxDeviceName() bool {
	if o != nil && o.HxDeviceName != nil {
		return true
	}

	return false
}

// SetHxDeviceName gets a reference to the given string and assigns it to the HxDeviceName field.
func (o *HyperflexDevicePackageDownloadStateAllOf) SetHxDeviceName(v string) {
	o.HxDeviceName = &v
}

// GetHxNodes returns the HxNodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexDevicePackageDownloadStateAllOf) GetHxNodes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.HxNodes
}

// GetHxNodesOk returns a tuple with the HxNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexDevicePackageDownloadStateAllOf) GetHxNodesOk() (*[]string, bool) {
	if o == nil || o.HxNodes == nil {
		return nil, false
	}
	return &o.HxNodes, true
}

// HasHxNodes returns a boolean if a field has been set.
func (o *HyperflexDevicePackageDownloadStateAllOf) HasHxNodes() bool {
	if o != nil && o.HxNodes != nil {
		return true
	}

	return false
}

// SetHxNodes gets a reference to the given []string and assigns it to the HxNodes field.
func (o *HyperflexDevicePackageDownloadStateAllOf) SetHxNodes(v []string) {
	o.HxNodes = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *HyperflexDevicePackageDownloadStateAllOf) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDevicePackageDownloadStateAllOf) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *HyperflexDevicePackageDownloadStateAllOf) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *HyperflexDevicePackageDownloadStateAllOf) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *HyperflexDevicePackageDownloadStateAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDevicePackageDownloadStateAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *HyperflexDevicePackageDownloadStateAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *HyperflexDevicePackageDownloadStateAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o HyperflexDevicePackageDownloadStateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Checksum != nil {
		toSerialize["Checksum"] = o.Checksum
	}
	if o.HxDeviceName != nil {
		toSerialize["HxDeviceName"] = o.HxDeviceName
	}
	if o.HxNodes != nil {
		toSerialize["HxNodes"] = o.HxNodes
	}
	if o.Timestamp != nil {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexDevicePackageDownloadStateAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexDevicePackageDownloadStateAllOf := _HyperflexDevicePackageDownloadStateAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexDevicePackageDownloadStateAllOf); err == nil {
		*o = HyperflexDevicePackageDownloadStateAllOf(varHyperflexDevicePackageDownloadStateAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Checksum")
		delete(additionalProperties, "HxDeviceName")
		delete(additionalProperties, "HxNodes")
		delete(additionalProperties, "Timestamp")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexDevicePackageDownloadStateAllOf struct {
	value *HyperflexDevicePackageDownloadStateAllOf
	isSet bool
}

func (v NullableHyperflexDevicePackageDownloadStateAllOf) Get() *HyperflexDevicePackageDownloadStateAllOf {
	return v.value
}

func (v *NullableHyperflexDevicePackageDownloadStateAllOf) Set(val *HyperflexDevicePackageDownloadStateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexDevicePackageDownloadStateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexDevicePackageDownloadStateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexDevicePackageDownloadStateAllOf(val *HyperflexDevicePackageDownloadStateAllOf) *NullableHyperflexDevicePackageDownloadStateAllOf {
	return &NullableHyperflexDevicePackageDownloadStateAllOf{value: val, isSet: true}
}

func (v NullableHyperflexDevicePackageDownloadStateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexDevicePackageDownloadStateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}