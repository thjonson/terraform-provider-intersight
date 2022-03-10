/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5517
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// NiatelemetrySnmpSrc Object to capture SNMP Src details.
type NiatelemetrySnmpSrc struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Parent common policy for SNMP Source in APIC.
	CommonPolicy *string `json:"CommonPolicy,omitempty"`
	// Dn of the SNMP Source in APIC.
	Dn *string `json:"Dn,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// List of SNMP trap destination for the above src.
	SnmpTrapDest *string `json:"SnmpTrapDest,omitempty"`
	// SNMP trap destination group for the above src.
	SnmpTrapDestGrp      *string                              `json:"SnmpTrapDestGrp,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetrySnmpSrc NiatelemetrySnmpSrc

// NewNiatelemetrySnmpSrc instantiates a new NiatelemetrySnmpSrc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetrySnmpSrc(classId string, objectType string) *NiatelemetrySnmpSrc {
	this := NiatelemetrySnmpSrc{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetrySnmpSrcWithDefaults instantiates a new NiatelemetrySnmpSrc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetrySnmpSrcWithDefaults() *NiatelemetrySnmpSrc {
	this := NiatelemetrySnmpSrc{}
	var classId string = "niatelemetry.SnmpSrc"
	this.ClassId = classId
	var objectType string = "niatelemetry.SnmpSrc"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetrySnmpSrc) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySnmpSrc) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetrySnmpSrc) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetrySnmpSrc) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySnmpSrc) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetrySnmpSrc) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCommonPolicy returns the CommonPolicy field value if set, zero value otherwise.
func (o *NiatelemetrySnmpSrc) GetCommonPolicy() string {
	if o == nil || o.CommonPolicy == nil {
		var ret string
		return ret
	}
	return *o.CommonPolicy
}

// GetCommonPolicyOk returns a tuple with the CommonPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySnmpSrc) GetCommonPolicyOk() (*string, bool) {
	if o == nil || o.CommonPolicy == nil {
		return nil, false
	}
	return o.CommonPolicy, true
}

// HasCommonPolicy returns a boolean if a field has been set.
func (o *NiatelemetrySnmpSrc) HasCommonPolicy() bool {
	if o != nil && o.CommonPolicy != nil {
		return true
	}

	return false
}

// SetCommonPolicy gets a reference to the given string and assigns it to the CommonPolicy field.
func (o *NiatelemetrySnmpSrc) SetCommonPolicy(v string) {
	o.CommonPolicy = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetrySnmpSrc) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySnmpSrc) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetrySnmpSrc) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetrySnmpSrc) SetDn(v string) {
	o.Dn = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetrySnmpSrc) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySnmpSrc) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetrySnmpSrc) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetrySnmpSrc) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetrySnmpSrc) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySnmpSrc) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetrySnmpSrc) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetrySnmpSrc) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetrySnmpSrc) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySnmpSrc) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetrySnmpSrc) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetrySnmpSrc) SetSiteName(v string) {
	o.SiteName = &v
}

// GetSnmpTrapDest returns the SnmpTrapDest field value if set, zero value otherwise.
func (o *NiatelemetrySnmpSrc) GetSnmpTrapDest() string {
	if o == nil || o.SnmpTrapDest == nil {
		var ret string
		return ret
	}
	return *o.SnmpTrapDest
}

// GetSnmpTrapDestOk returns a tuple with the SnmpTrapDest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySnmpSrc) GetSnmpTrapDestOk() (*string, bool) {
	if o == nil || o.SnmpTrapDest == nil {
		return nil, false
	}
	return o.SnmpTrapDest, true
}

// HasSnmpTrapDest returns a boolean if a field has been set.
func (o *NiatelemetrySnmpSrc) HasSnmpTrapDest() bool {
	if o != nil && o.SnmpTrapDest != nil {
		return true
	}

	return false
}

// SetSnmpTrapDest gets a reference to the given string and assigns it to the SnmpTrapDest field.
func (o *NiatelemetrySnmpSrc) SetSnmpTrapDest(v string) {
	o.SnmpTrapDest = &v
}

// GetSnmpTrapDestGrp returns the SnmpTrapDestGrp field value if set, zero value otherwise.
func (o *NiatelemetrySnmpSrc) GetSnmpTrapDestGrp() string {
	if o == nil || o.SnmpTrapDestGrp == nil {
		var ret string
		return ret
	}
	return *o.SnmpTrapDestGrp
}

// GetSnmpTrapDestGrpOk returns a tuple with the SnmpTrapDestGrp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySnmpSrc) GetSnmpTrapDestGrpOk() (*string, bool) {
	if o == nil || o.SnmpTrapDestGrp == nil {
		return nil, false
	}
	return o.SnmpTrapDestGrp, true
}

// HasSnmpTrapDestGrp returns a boolean if a field has been set.
func (o *NiatelemetrySnmpSrc) HasSnmpTrapDestGrp() bool {
	if o != nil && o.SnmpTrapDestGrp != nil {
		return true
	}

	return false
}

// SetSnmpTrapDestGrp gets a reference to the given string and assigns it to the SnmpTrapDestGrp field.
func (o *NiatelemetrySnmpSrc) SetSnmpTrapDestGrp(v string) {
	o.SnmpTrapDestGrp = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetrySnmpSrc) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySnmpSrc) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetrySnmpSrc) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetrySnmpSrc) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetrySnmpSrc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CommonPolicy != nil {
		toSerialize["CommonPolicy"] = o.CommonPolicy
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
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
	if o.SnmpTrapDest != nil {
		toSerialize["SnmpTrapDest"] = o.SnmpTrapDest
	}
	if o.SnmpTrapDestGrp != nil {
		toSerialize["SnmpTrapDestGrp"] = o.SnmpTrapDestGrp
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetrySnmpSrc) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetrySnmpSrcWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Parent common policy for SNMP Source in APIC.
		CommonPolicy *string `json:"CommonPolicy,omitempty"`
		// Dn of the SNMP Source in APIC.
		Dn *string `json:"Dn,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName *string `json:"SiteName,omitempty"`
		// List of SNMP trap destination for the above src.
		SnmpTrapDest *string `json:"SnmpTrapDest,omitempty"`
		// SNMP trap destination group for the above src.
		SnmpTrapDestGrp  *string                              `json:"SnmpTrapDestGrp,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetrySnmpSrcWithoutEmbeddedStruct := NiatelemetrySnmpSrcWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetrySnmpSrcWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetrySnmpSrc := _NiatelemetrySnmpSrc{}
		varNiatelemetrySnmpSrc.ClassId = varNiatelemetrySnmpSrcWithoutEmbeddedStruct.ClassId
		varNiatelemetrySnmpSrc.ObjectType = varNiatelemetrySnmpSrcWithoutEmbeddedStruct.ObjectType
		varNiatelemetrySnmpSrc.CommonPolicy = varNiatelemetrySnmpSrcWithoutEmbeddedStruct.CommonPolicy
		varNiatelemetrySnmpSrc.Dn = varNiatelemetrySnmpSrcWithoutEmbeddedStruct.Dn
		varNiatelemetrySnmpSrc.RecordType = varNiatelemetrySnmpSrcWithoutEmbeddedStruct.RecordType
		varNiatelemetrySnmpSrc.RecordVersion = varNiatelemetrySnmpSrcWithoutEmbeddedStruct.RecordVersion
		varNiatelemetrySnmpSrc.SiteName = varNiatelemetrySnmpSrcWithoutEmbeddedStruct.SiteName
		varNiatelemetrySnmpSrc.SnmpTrapDest = varNiatelemetrySnmpSrcWithoutEmbeddedStruct.SnmpTrapDest
		varNiatelemetrySnmpSrc.SnmpTrapDestGrp = varNiatelemetrySnmpSrcWithoutEmbeddedStruct.SnmpTrapDestGrp
		varNiatelemetrySnmpSrc.RegisteredDevice = varNiatelemetrySnmpSrcWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetrySnmpSrc(varNiatelemetrySnmpSrc)
	} else {
		return err
	}

	varNiatelemetrySnmpSrc := _NiatelemetrySnmpSrc{}

	err = json.Unmarshal(bytes, &varNiatelemetrySnmpSrc)
	if err == nil {
		o.MoBaseMo = varNiatelemetrySnmpSrc.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CommonPolicy")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "SnmpTrapDest")
		delete(additionalProperties, "SnmpTrapDestGrp")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableNiatelemetrySnmpSrc struct {
	value *NiatelemetrySnmpSrc
	isSet bool
}

func (v NullableNiatelemetrySnmpSrc) Get() *NiatelemetrySnmpSrc {
	return v.value
}

func (v *NullableNiatelemetrySnmpSrc) Set(val *NiatelemetrySnmpSrc) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetrySnmpSrc) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetrySnmpSrc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetrySnmpSrc(val *NiatelemetrySnmpSrc) *NullableNiatelemetrySnmpSrc {
	return &NullableNiatelemetrySnmpSrc{value: val, isSet: true}
}

func (v NullableNiatelemetrySnmpSrc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetrySnmpSrc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}