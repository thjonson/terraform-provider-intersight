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

// NiatelemetryHttpsAclContractDetails Object to capture the HTTPS ACL contract details in APIC.
type NiatelemetryHttpsAclContractDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Consumer Dn of the HTTPS ACL contract children MOs for APIC.
	ConsumerDn *string `json:"ConsumerDn,omitempty"`
	// Name of HTTPS ACL contract for APIC.
	ContractName *string `json:"ContractName,omitempty"`
	// Mgmt Inst Dn of the HTTPS ACL contract children MOs for APIC.
	MgmtInstpDn *string `json:"MgmtInstpDn,omitempty"`
	// Mgmt subnet address of the HTTPS ACL contract children MOs for APIC.
	MgmtSubnetAddresses *string `json:"MgmtSubnetAddresses,omitempty"`
	// Provider dn of the HTTPS ACL contract children MOs for APIC.
	ProviderDn *string `json:"ProviderDn,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName             *string                              `json:"SiteName,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryHttpsAclContractDetails NiatelemetryHttpsAclContractDetails

// NewNiatelemetryHttpsAclContractDetails instantiates a new NiatelemetryHttpsAclContractDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryHttpsAclContractDetails(classId string, objectType string) *NiatelemetryHttpsAclContractDetails {
	this := NiatelemetryHttpsAclContractDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryHttpsAclContractDetailsWithDefaults instantiates a new NiatelemetryHttpsAclContractDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryHttpsAclContractDetailsWithDefaults() *NiatelemetryHttpsAclContractDetails {
	this := NiatelemetryHttpsAclContractDetails{}
	var classId string = "niatelemetry.HttpsAclContractDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.HttpsAclContractDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryHttpsAclContractDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryHttpsAclContractDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryHttpsAclContractDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryHttpsAclContractDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConsumerDn returns the ConsumerDn field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractDetails) GetConsumerDn() string {
	if o == nil || o.ConsumerDn == nil {
		var ret string
		return ret
	}
	return *o.ConsumerDn
}

// GetConsumerDnOk returns a tuple with the ConsumerDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractDetails) GetConsumerDnOk() (*string, bool) {
	if o == nil || o.ConsumerDn == nil {
		return nil, false
	}
	return o.ConsumerDn, true
}

// HasConsumerDn returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractDetails) HasConsumerDn() bool {
	if o != nil && o.ConsumerDn != nil {
		return true
	}

	return false
}

// SetConsumerDn gets a reference to the given string and assigns it to the ConsumerDn field.
func (o *NiatelemetryHttpsAclContractDetails) SetConsumerDn(v string) {
	o.ConsumerDn = &v
}

// GetContractName returns the ContractName field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractDetails) GetContractName() string {
	if o == nil || o.ContractName == nil {
		var ret string
		return ret
	}
	return *o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractDetails) GetContractNameOk() (*string, bool) {
	if o == nil || o.ContractName == nil {
		return nil, false
	}
	return o.ContractName, true
}

// HasContractName returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractDetails) HasContractName() bool {
	if o != nil && o.ContractName != nil {
		return true
	}

	return false
}

// SetContractName gets a reference to the given string and assigns it to the ContractName field.
func (o *NiatelemetryHttpsAclContractDetails) SetContractName(v string) {
	o.ContractName = &v
}

// GetMgmtInstpDn returns the MgmtInstpDn field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractDetails) GetMgmtInstpDn() string {
	if o == nil || o.MgmtInstpDn == nil {
		var ret string
		return ret
	}
	return *o.MgmtInstpDn
}

// GetMgmtInstpDnOk returns a tuple with the MgmtInstpDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractDetails) GetMgmtInstpDnOk() (*string, bool) {
	if o == nil || o.MgmtInstpDn == nil {
		return nil, false
	}
	return o.MgmtInstpDn, true
}

// HasMgmtInstpDn returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractDetails) HasMgmtInstpDn() bool {
	if o != nil && o.MgmtInstpDn != nil {
		return true
	}

	return false
}

// SetMgmtInstpDn gets a reference to the given string and assigns it to the MgmtInstpDn field.
func (o *NiatelemetryHttpsAclContractDetails) SetMgmtInstpDn(v string) {
	o.MgmtInstpDn = &v
}

// GetMgmtSubnetAddresses returns the MgmtSubnetAddresses field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractDetails) GetMgmtSubnetAddresses() string {
	if o == nil || o.MgmtSubnetAddresses == nil {
		var ret string
		return ret
	}
	return *o.MgmtSubnetAddresses
}

// GetMgmtSubnetAddressesOk returns a tuple with the MgmtSubnetAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractDetails) GetMgmtSubnetAddressesOk() (*string, bool) {
	if o == nil || o.MgmtSubnetAddresses == nil {
		return nil, false
	}
	return o.MgmtSubnetAddresses, true
}

// HasMgmtSubnetAddresses returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractDetails) HasMgmtSubnetAddresses() bool {
	if o != nil && o.MgmtSubnetAddresses != nil {
		return true
	}

	return false
}

// SetMgmtSubnetAddresses gets a reference to the given string and assigns it to the MgmtSubnetAddresses field.
func (o *NiatelemetryHttpsAclContractDetails) SetMgmtSubnetAddresses(v string) {
	o.MgmtSubnetAddresses = &v
}

// GetProviderDn returns the ProviderDn field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractDetails) GetProviderDn() string {
	if o == nil || o.ProviderDn == nil {
		var ret string
		return ret
	}
	return *o.ProviderDn
}

// GetProviderDnOk returns a tuple with the ProviderDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractDetails) GetProviderDnOk() (*string, bool) {
	if o == nil || o.ProviderDn == nil {
		return nil, false
	}
	return o.ProviderDn, true
}

// HasProviderDn returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractDetails) HasProviderDn() bool {
	if o != nil && o.ProviderDn != nil {
		return true
	}

	return false
}

// SetProviderDn gets a reference to the given string and assigns it to the ProviderDn field.
func (o *NiatelemetryHttpsAclContractDetails) SetProviderDn(v string) {
	o.ProviderDn = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractDetails) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractDetails) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractDetails) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryHttpsAclContractDetails) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractDetails) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractDetails) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractDetails) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryHttpsAclContractDetails) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractDetails) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractDetails) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractDetails) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryHttpsAclContractDetails) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclContractDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclContractDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclContractDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryHttpsAclContractDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryHttpsAclContractDetails) MarshalJSON() ([]byte, error) {
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
	if o.ConsumerDn != nil {
		toSerialize["ConsumerDn"] = o.ConsumerDn
	}
	if o.ContractName != nil {
		toSerialize["ContractName"] = o.ContractName
	}
	if o.MgmtInstpDn != nil {
		toSerialize["MgmtInstpDn"] = o.MgmtInstpDn
	}
	if o.MgmtSubnetAddresses != nil {
		toSerialize["MgmtSubnetAddresses"] = o.MgmtSubnetAddresses
	}
	if o.ProviderDn != nil {
		toSerialize["ProviderDn"] = o.ProviderDn
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

func (o *NiatelemetryHttpsAclContractDetails) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Consumer Dn of the HTTPS ACL contract children MOs for APIC.
		ConsumerDn *string `json:"ConsumerDn,omitempty"`
		// Name of HTTPS ACL contract for APIC.
		ContractName *string `json:"ContractName,omitempty"`
		// Mgmt Inst Dn of the HTTPS ACL contract children MOs for APIC.
		MgmtInstpDn *string `json:"MgmtInstpDn,omitempty"`
		// Mgmt subnet address of the HTTPS ACL contract children MOs for APIC.
		MgmtSubnetAddresses *string `json:"MgmtSubnetAddresses,omitempty"`
		// Provider dn of the HTTPS ACL contract children MOs for APIC.
		ProviderDn *string `json:"ProviderDn,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName         *string                              `json:"SiteName,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct := NiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryHttpsAclContractDetails := _NiatelemetryHttpsAclContractDetails{}
		varNiatelemetryHttpsAclContractDetails.ClassId = varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryHttpsAclContractDetails.ObjectType = varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryHttpsAclContractDetails.ConsumerDn = varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct.ConsumerDn
		varNiatelemetryHttpsAclContractDetails.ContractName = varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct.ContractName
		varNiatelemetryHttpsAclContractDetails.MgmtInstpDn = varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct.MgmtInstpDn
		varNiatelemetryHttpsAclContractDetails.MgmtSubnetAddresses = varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct.MgmtSubnetAddresses
		varNiatelemetryHttpsAclContractDetails.ProviderDn = varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct.ProviderDn
		varNiatelemetryHttpsAclContractDetails.RecordType = varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct.RecordType
		varNiatelemetryHttpsAclContractDetails.RecordVersion = varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryHttpsAclContractDetails.SiteName = varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct.SiteName
		varNiatelemetryHttpsAclContractDetails.RegisteredDevice = varNiatelemetryHttpsAclContractDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryHttpsAclContractDetails(varNiatelemetryHttpsAclContractDetails)
	} else {
		return err
	}

	varNiatelemetryHttpsAclContractDetails := _NiatelemetryHttpsAclContractDetails{}

	err = json.Unmarshal(bytes, &varNiatelemetryHttpsAclContractDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryHttpsAclContractDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConsumerDn")
		delete(additionalProperties, "ContractName")
		delete(additionalProperties, "MgmtInstpDn")
		delete(additionalProperties, "MgmtSubnetAddresses")
		delete(additionalProperties, "ProviderDn")
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

type NullableNiatelemetryHttpsAclContractDetails struct {
	value *NiatelemetryHttpsAclContractDetails
	isSet bool
}

func (v NullableNiatelemetryHttpsAclContractDetails) Get() *NiatelemetryHttpsAclContractDetails {
	return v.value
}

func (v *NullableNiatelemetryHttpsAclContractDetails) Set(val *NiatelemetryHttpsAclContractDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryHttpsAclContractDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryHttpsAclContractDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryHttpsAclContractDetails(val *NiatelemetryHttpsAclContractDetails) *NullableNiatelemetryHttpsAclContractDetails {
	return &NullableNiatelemetryHttpsAclContractDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryHttpsAclContractDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryHttpsAclContractDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
