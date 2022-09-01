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

// NiatelemetryPodTimeServerPoliciesAllOf Definition of the list of properties defined in 'niatelemetry.PodTimeServerPolicies', excluding properties defined in parent classes.
type NiatelemetryPodTimeServerPoliciesAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Dn of the Time server Pol in APIC.
	PolDn *string `json:"PolDn,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// Admin State of the time server Pol in APIC.
	TimeServerEnableSt *string `json:"TimeServerEnableSt,omitempty"`
	// Time server of the time server Pol in APIC.
	TimeServers          *string                              `json:"TimeServers,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryPodTimeServerPoliciesAllOf NiatelemetryPodTimeServerPoliciesAllOf

// NewNiatelemetryPodTimeServerPoliciesAllOf instantiates a new NiatelemetryPodTimeServerPoliciesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryPodTimeServerPoliciesAllOf(classId string, objectType string) *NiatelemetryPodTimeServerPoliciesAllOf {
	this := NiatelemetryPodTimeServerPoliciesAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryPodTimeServerPoliciesAllOfWithDefaults instantiates a new NiatelemetryPodTimeServerPoliciesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryPodTimeServerPoliciesAllOfWithDefaults() *NiatelemetryPodTimeServerPoliciesAllOf {
	this := NiatelemetryPodTimeServerPoliciesAllOf{}
	var classId string = "niatelemetry.PodTimeServerPolicies"
	this.ClassId = classId
	var objectType string = "niatelemetry.PodTimeServerPolicies"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryPodTimeServerPoliciesAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryPodTimeServerPoliciesAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPolDn returns the PolDn field value if set, zero value otherwise.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetPolDn() string {
	if o == nil || o.PolDn == nil {
		var ret string
		return ret
	}
	return *o.PolDn
}

// GetPolDnOk returns a tuple with the PolDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetPolDnOk() (*string, bool) {
	if o == nil || o.PolDn == nil {
		return nil, false
	}
	return o.PolDn, true
}

// HasPolDn returns a boolean if a field has been set.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) HasPolDn() bool {
	if o != nil && o.PolDn != nil {
		return true
	}

	return false
}

// SetPolDn gets a reference to the given string and assigns it to the PolDn field.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) SetPolDn(v string) {
	o.PolDn = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetTimeServerEnableSt returns the TimeServerEnableSt field value if set, zero value otherwise.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetTimeServerEnableSt() string {
	if o == nil || o.TimeServerEnableSt == nil {
		var ret string
		return ret
	}
	return *o.TimeServerEnableSt
}

// GetTimeServerEnableStOk returns a tuple with the TimeServerEnableSt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetTimeServerEnableStOk() (*string, bool) {
	if o == nil || o.TimeServerEnableSt == nil {
		return nil, false
	}
	return o.TimeServerEnableSt, true
}

// HasTimeServerEnableSt returns a boolean if a field has been set.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) HasTimeServerEnableSt() bool {
	if o != nil && o.TimeServerEnableSt != nil {
		return true
	}

	return false
}

// SetTimeServerEnableSt gets a reference to the given string and assigns it to the TimeServerEnableSt field.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) SetTimeServerEnableSt(v string) {
	o.TimeServerEnableSt = &v
}

// GetTimeServers returns the TimeServers field value if set, zero value otherwise.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetTimeServers() string {
	if o == nil || o.TimeServers == nil {
		var ret string
		return ret
	}
	return *o.TimeServers
}

// GetTimeServersOk returns a tuple with the TimeServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetTimeServersOk() (*string, bool) {
	if o == nil || o.TimeServers == nil {
		return nil, false
	}
	return o.TimeServers, true
}

// HasTimeServers returns a boolean if a field has been set.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) HasTimeServers() bool {
	if o != nil && o.TimeServers != nil {
		return true
	}

	return false
}

// SetTimeServers gets a reference to the given string and assigns it to the TimeServers field.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) SetTimeServers(v string) {
	o.TimeServers = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryPodTimeServerPoliciesAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryPodTimeServerPoliciesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.PolDn != nil {
		toSerialize["PolDn"] = o.PolDn
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.RecordVersion != nil {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if o.SiteName != nil {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.TimeServerEnableSt != nil {
		toSerialize["TimeServerEnableSt"] = o.TimeServerEnableSt
	}
	if o.TimeServers != nil {
		toSerialize["TimeServers"] = o.TimeServers
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryPodTimeServerPoliciesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryPodTimeServerPoliciesAllOf := _NiatelemetryPodTimeServerPoliciesAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryPodTimeServerPoliciesAllOf); err == nil {
		*o = NiatelemetryPodTimeServerPoliciesAllOf(varNiatelemetryPodTimeServerPoliciesAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PolDn")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "TimeServerEnableSt")
		delete(additionalProperties, "TimeServers")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryPodTimeServerPoliciesAllOf struct {
	value *NiatelemetryPodTimeServerPoliciesAllOf
	isSet bool
}

func (v NullableNiatelemetryPodTimeServerPoliciesAllOf) Get() *NiatelemetryPodTimeServerPoliciesAllOf {
	return v.value
}

func (v *NullableNiatelemetryPodTimeServerPoliciesAllOf) Set(val *NiatelemetryPodTimeServerPoliciesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryPodTimeServerPoliciesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryPodTimeServerPoliciesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryPodTimeServerPoliciesAllOf(val *NiatelemetryPodTimeServerPoliciesAllOf) *NullableNiatelemetryPodTimeServerPoliciesAllOf {
	return &NullableNiatelemetryPodTimeServerPoliciesAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryPodTimeServerPoliciesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryPodTimeServerPoliciesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
