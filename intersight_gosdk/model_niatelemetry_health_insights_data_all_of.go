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

// NiatelemetryHealthInsightsDataAllOf Definition of the list of properties defined in 'niatelemetry.HealthInsightsData', excluding properties defined in parent classes.
type NiatelemetryHealthInsightsDataAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Flt eqpt flash minor alarm fault count for the APIC.
	FltEqptFlashMinorAlarmFaultCount *int64 `json:"FltEqptFlashMinorAlarmFaultCount,omitempty"`
	// Flt eqpt flash worn out fault count for the APIC.
	FltEqptFlashWornOutFaultCount *int64 `json:"FltEqptFlashWornOutFaultCount,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected..
	SiteName *string `json:"SiteName,omitempty"`
	// Ssd fault monitoring count for the APIC.
	SsdMonitoringFaultCount *int64                               `json:"SsdMonitoringFaultCount,omitempty"`
	RegisteredDevice        *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _NiatelemetryHealthInsightsDataAllOf NiatelemetryHealthInsightsDataAllOf

// NewNiatelemetryHealthInsightsDataAllOf instantiates a new NiatelemetryHealthInsightsDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryHealthInsightsDataAllOf(classId string, objectType string) *NiatelemetryHealthInsightsDataAllOf {
	this := NiatelemetryHealthInsightsDataAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryHealthInsightsDataAllOfWithDefaults instantiates a new NiatelemetryHealthInsightsDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryHealthInsightsDataAllOfWithDefaults() *NiatelemetryHealthInsightsDataAllOf {
	this := NiatelemetryHealthInsightsDataAllOf{}
	var classId string = "niatelemetry.HealthInsightsData"
	this.ClassId = classId
	var objectType string = "niatelemetry.HealthInsightsData"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryHealthInsightsDataAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryHealthInsightsDataAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryHealthInsightsDataAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryHealthInsightsDataAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryHealthInsightsDataAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryHealthInsightsDataAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFltEqptFlashMinorAlarmFaultCount returns the FltEqptFlashMinorAlarmFaultCount field value if set, zero value otherwise.
func (o *NiatelemetryHealthInsightsDataAllOf) GetFltEqptFlashMinorAlarmFaultCount() int64 {
	if o == nil || o.FltEqptFlashMinorAlarmFaultCount == nil {
		var ret int64
		return ret
	}
	return *o.FltEqptFlashMinorAlarmFaultCount
}

// GetFltEqptFlashMinorAlarmFaultCountOk returns a tuple with the FltEqptFlashMinorAlarmFaultCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHealthInsightsDataAllOf) GetFltEqptFlashMinorAlarmFaultCountOk() (*int64, bool) {
	if o == nil || o.FltEqptFlashMinorAlarmFaultCount == nil {
		return nil, false
	}
	return o.FltEqptFlashMinorAlarmFaultCount, true
}

// HasFltEqptFlashMinorAlarmFaultCount returns a boolean if a field has been set.
func (o *NiatelemetryHealthInsightsDataAllOf) HasFltEqptFlashMinorAlarmFaultCount() bool {
	if o != nil && o.FltEqptFlashMinorAlarmFaultCount != nil {
		return true
	}

	return false
}

// SetFltEqptFlashMinorAlarmFaultCount gets a reference to the given int64 and assigns it to the FltEqptFlashMinorAlarmFaultCount field.
func (o *NiatelemetryHealthInsightsDataAllOf) SetFltEqptFlashMinorAlarmFaultCount(v int64) {
	o.FltEqptFlashMinorAlarmFaultCount = &v
}

// GetFltEqptFlashWornOutFaultCount returns the FltEqptFlashWornOutFaultCount field value if set, zero value otherwise.
func (o *NiatelemetryHealthInsightsDataAllOf) GetFltEqptFlashWornOutFaultCount() int64 {
	if o == nil || o.FltEqptFlashWornOutFaultCount == nil {
		var ret int64
		return ret
	}
	return *o.FltEqptFlashWornOutFaultCount
}

// GetFltEqptFlashWornOutFaultCountOk returns a tuple with the FltEqptFlashWornOutFaultCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHealthInsightsDataAllOf) GetFltEqptFlashWornOutFaultCountOk() (*int64, bool) {
	if o == nil || o.FltEqptFlashWornOutFaultCount == nil {
		return nil, false
	}
	return o.FltEqptFlashWornOutFaultCount, true
}

// HasFltEqptFlashWornOutFaultCount returns a boolean if a field has been set.
func (o *NiatelemetryHealthInsightsDataAllOf) HasFltEqptFlashWornOutFaultCount() bool {
	if o != nil && o.FltEqptFlashWornOutFaultCount != nil {
		return true
	}

	return false
}

// SetFltEqptFlashWornOutFaultCount gets a reference to the given int64 and assigns it to the FltEqptFlashWornOutFaultCount field.
func (o *NiatelemetryHealthInsightsDataAllOf) SetFltEqptFlashWornOutFaultCount(v int64) {
	o.FltEqptFlashWornOutFaultCount = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryHealthInsightsDataAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHealthInsightsDataAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryHealthInsightsDataAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryHealthInsightsDataAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryHealthInsightsDataAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHealthInsightsDataAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryHealthInsightsDataAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryHealthInsightsDataAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryHealthInsightsDataAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHealthInsightsDataAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryHealthInsightsDataAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryHealthInsightsDataAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetSsdMonitoringFaultCount returns the SsdMonitoringFaultCount field value if set, zero value otherwise.
func (o *NiatelemetryHealthInsightsDataAllOf) GetSsdMonitoringFaultCount() int64 {
	if o == nil || o.SsdMonitoringFaultCount == nil {
		var ret int64
		return ret
	}
	return *o.SsdMonitoringFaultCount
}

// GetSsdMonitoringFaultCountOk returns a tuple with the SsdMonitoringFaultCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHealthInsightsDataAllOf) GetSsdMonitoringFaultCountOk() (*int64, bool) {
	if o == nil || o.SsdMonitoringFaultCount == nil {
		return nil, false
	}
	return o.SsdMonitoringFaultCount, true
}

// HasSsdMonitoringFaultCount returns a boolean if a field has been set.
func (o *NiatelemetryHealthInsightsDataAllOf) HasSsdMonitoringFaultCount() bool {
	if o != nil && o.SsdMonitoringFaultCount != nil {
		return true
	}

	return false
}

// SetSsdMonitoringFaultCount gets a reference to the given int64 and assigns it to the SsdMonitoringFaultCount field.
func (o *NiatelemetryHealthInsightsDataAllOf) SetSsdMonitoringFaultCount(v int64) {
	o.SsdMonitoringFaultCount = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryHealthInsightsDataAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHealthInsightsDataAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryHealthInsightsDataAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryHealthInsightsDataAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryHealthInsightsDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FltEqptFlashMinorAlarmFaultCount != nil {
		toSerialize["FltEqptFlashMinorAlarmFaultCount"] = o.FltEqptFlashMinorAlarmFaultCount
	}
	if o.FltEqptFlashWornOutFaultCount != nil {
		toSerialize["FltEqptFlashWornOutFaultCount"] = o.FltEqptFlashWornOutFaultCount
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
	if o.SsdMonitoringFaultCount != nil {
		toSerialize["SsdMonitoringFaultCount"] = o.SsdMonitoringFaultCount
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryHealthInsightsDataAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryHealthInsightsDataAllOf := _NiatelemetryHealthInsightsDataAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryHealthInsightsDataAllOf); err == nil {
		*o = NiatelemetryHealthInsightsDataAllOf(varNiatelemetryHealthInsightsDataAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FltEqptFlashMinorAlarmFaultCount")
		delete(additionalProperties, "FltEqptFlashWornOutFaultCount")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "SsdMonitoringFaultCount")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryHealthInsightsDataAllOf struct {
	value *NiatelemetryHealthInsightsDataAllOf
	isSet bool
}

func (v NullableNiatelemetryHealthInsightsDataAllOf) Get() *NiatelemetryHealthInsightsDataAllOf {
	return v.value
}

func (v *NullableNiatelemetryHealthInsightsDataAllOf) Set(val *NiatelemetryHealthInsightsDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryHealthInsightsDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryHealthInsightsDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryHealthInsightsDataAllOf(val *NiatelemetryHealthInsightsDataAllOf) *NullableNiatelemetryHealthInsightsDataAllOf {
	return &NullableNiatelemetryHealthInsightsDataAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryHealthInsightsDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryHealthInsightsDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
