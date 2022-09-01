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
	"reflect"
	"strings"
)

// NiatelemetryHttpsAclContractFilterMap Object to capture the HTTPS ACL contract filter map in APIC.
type NiatelemetryHttpsAclContractFilterMap struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of HTTPS ACL contract for APIC.
	ContractName *string `json:"ContractName,omitempty"`
	// Dn of the HTTPS ACL EPGs subject filter relation MO for APIC.
	Dn *string `json:"Dn,omitempty"`
	// Name of HTTPS ACL filter for APIC.
	FilterName *string `json:"FilterName,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName             *string                              `json:"SiteName,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryHttpsAclContractFilterMap NiatelemetryHttpsAclContractFilterMap

// NewNiatelemetryHttpsAclContractFilterMap instantiates a new NiatelemetryHttpsAclContractFilterMap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryHttpsAclContractFilterMap(classId string, objectType string) *NiatelemetryHttpsAclContractFilterMap {
	this := NiatelemetryHttpsAclContractFilterMap{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryHttpsAclContractFilterMapWithDefaults instantiates a new NiatelemetryHttpsAclContractFilterMap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryHttpsAclContractFilterMapWithDefaults() *NiatelemetryHttpsAclContractFilterMap {
	this := NiatelemetryHttpsAclContractFilterMap{}
	var classId string = "niatelemetry.HttpsAclContractFilterMap"
	this.ClassId = classId
	var objectType string = "niatelemetry.HttpsAclContractFilterMap"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryHttpsAclContractFilterMap) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryHttpsAclContractFilterMap) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryHttpsAclContractFilterMap) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryHttpsAclContractFilterMap) SetObjectType(v string) {
	o.ObjectType = v
}

// GetContractName returns the ContractName field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractFilterMap) GetContractName() string {
	if o == nil || o.ContractName == nil {
		var ret string
		return ret
	}
	return *o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) GetContractNameOk() (*string, bool) {
	if o == nil || o.ContractName == nil {
		return nil, false
	}
	return o.ContractName, true
}

// HasContractName returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) HasContractName() bool {
	if o != nil && o.ContractName != nil {
		return true
	}

	return false
}

// SetContractName gets a reference to the given string and assigns it to the ContractName field.
func (o *NiatelemetryHttpsAclContractFilterMap) SetContractName(v string) {
	o.ContractName = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractFilterMap) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryHttpsAclContractFilterMap) SetDn(v string) {
	o.Dn = &v
}

// GetFilterName returns the FilterName field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractFilterMap) GetFilterName() string {
	if o == nil || o.FilterName == nil {
		var ret string
		return ret
	}
	return *o.FilterName
}

// GetFilterNameOk returns a tuple with the FilterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) GetFilterNameOk() (*string, bool) {
	if o == nil || o.FilterName == nil {
		return nil, false
	}
	return o.FilterName, true
}

// HasFilterName returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) HasFilterName() bool {
	if o != nil && o.FilterName != nil {
		return true
	}

	return false
}

// SetFilterName gets a reference to the given string and assigns it to the FilterName field.
func (o *NiatelemetryHttpsAclContractFilterMap) SetFilterName(v string) {
	o.FilterName = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractFilterMap) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryHttpsAclContractFilterMap) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractFilterMap) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryHttpsAclContractFilterMap) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractFilterMap) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryHttpsAclContractFilterMap) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractFilterMap) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractFilterMap) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryHttpsAclContractFilterMap) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryHttpsAclContractFilterMap) MarshalJSON() ([]byte, error) {
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
	if o.ContractName != nil {
		toSerialize["ContractName"] = o.ContractName
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.FilterName != nil {
		toSerialize["FilterName"] = o.FilterName
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
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryHttpsAclContractFilterMap) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of HTTPS ACL contract for APIC.
		ContractName *string `json:"ContractName,omitempty"`
		// Dn of the HTTPS ACL EPGs subject filter relation MO for APIC.
		Dn *string `json:"Dn,omitempty"`
		// Name of HTTPS ACL filter for APIC.
		FilterName *string `json:"FilterName,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName         *string                              `json:"SiteName,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct := NiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryHttpsAclContractFilterMap := _NiatelemetryHttpsAclContractFilterMap{}
		varNiatelemetryHttpsAclContractFilterMap.ClassId = varNiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct.ClassId
		varNiatelemetryHttpsAclContractFilterMap.ObjectType = varNiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct.ObjectType
		varNiatelemetryHttpsAclContractFilterMap.ContractName = varNiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct.ContractName
		varNiatelemetryHttpsAclContractFilterMap.Dn = varNiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct.Dn
		varNiatelemetryHttpsAclContractFilterMap.FilterName = varNiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct.FilterName
		varNiatelemetryHttpsAclContractFilterMap.RecordType = varNiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct.RecordType
		varNiatelemetryHttpsAclContractFilterMap.RecordVersion = varNiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryHttpsAclContractFilterMap.SiteName = varNiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct.SiteName
		varNiatelemetryHttpsAclContractFilterMap.RegisteredDevice = varNiatelemetryHttpsAclContractFilterMapWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryHttpsAclContractFilterMap(varNiatelemetryHttpsAclContractFilterMap)
	} else {
		return err
	}

	varNiatelemetryHttpsAclContractFilterMap := _NiatelemetryHttpsAclContractFilterMap{}

	err = json.Unmarshal(bytes, &varNiatelemetryHttpsAclContractFilterMap)
	if err == nil {
		o.MoBaseMo = varNiatelemetryHttpsAclContractFilterMap.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ContractName")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "FilterName")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
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

type NullableNiatelemetryHttpsAclContractFilterMap struct {
	value *NiatelemetryHttpsAclContractFilterMap
	isSet bool
}

func (v NullableNiatelemetryHttpsAclContractFilterMap) Get() *NiatelemetryHttpsAclContractFilterMap {
	return v.value
}

func (v *NullableNiatelemetryHttpsAclContractFilterMap) Set(val *NiatelemetryHttpsAclContractFilterMap) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryHttpsAclContractFilterMap) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryHttpsAclContractFilterMap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryHttpsAclContractFilterMap(val *NiatelemetryHttpsAclContractFilterMap) *NullableNiatelemetryHttpsAclContractFilterMap {
	return &NullableNiatelemetryHttpsAclContractFilterMap{value: val, isSet: true}
}

func (v NullableNiatelemetryHttpsAclContractFilterMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryHttpsAclContractFilterMap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
