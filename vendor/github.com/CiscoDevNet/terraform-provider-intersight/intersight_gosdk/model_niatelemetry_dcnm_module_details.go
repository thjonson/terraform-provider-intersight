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

// NiatelemetryDcnmModuleDetails Inventory Object available for DCNM hardware modules.
type NiatelemetryDcnmModuleDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the hardware module in the fabric inventory.
	Name *string `json:"Name,omitempty"`
	// Product ID of the hardware module in the fabric inventory.
	ProductId *string `json:"ProductId,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Serial number of the hardware module in the fabric inventory.
	SerialNumber *string `json:"SerialNumber,omitempty"`
	// Vendor Id of the hardware module in the fabric inventory.
	VendorId             *string                              `json:"VendorId,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryDcnmModuleDetails NiatelemetryDcnmModuleDetails

// NewNiatelemetryDcnmModuleDetails instantiates a new NiatelemetryDcnmModuleDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryDcnmModuleDetails(classId string, objectType string) *NiatelemetryDcnmModuleDetails {
	this := NiatelemetryDcnmModuleDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryDcnmModuleDetailsWithDefaults instantiates a new NiatelemetryDcnmModuleDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryDcnmModuleDetailsWithDefaults() *NiatelemetryDcnmModuleDetails {
	this := NiatelemetryDcnmModuleDetails{}
	var classId string = "niatelemetry.DcnmModuleDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.DcnmModuleDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryDcnmModuleDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmModuleDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryDcnmModuleDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryDcnmModuleDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmModuleDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryDcnmModuleDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NiatelemetryDcnmModuleDetails) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmModuleDetails) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NiatelemetryDcnmModuleDetails) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NiatelemetryDcnmModuleDetails) SetName(v string) {
	o.Name = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *NiatelemetryDcnmModuleDetails) GetProductId() string {
	if o == nil || o.ProductId == nil {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmModuleDetails) GetProductIdOk() (*string, bool) {
	if o == nil || o.ProductId == nil {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *NiatelemetryDcnmModuleDetails) HasProductId() bool {
	if o != nil && o.ProductId != nil {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *NiatelemetryDcnmModuleDetails) SetProductId(v string) {
	o.ProductId = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryDcnmModuleDetails) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmModuleDetails) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryDcnmModuleDetails) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryDcnmModuleDetails) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryDcnmModuleDetails) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmModuleDetails) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryDcnmModuleDetails) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryDcnmModuleDetails) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *NiatelemetryDcnmModuleDetails) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmModuleDetails) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *NiatelemetryDcnmModuleDetails) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *NiatelemetryDcnmModuleDetails) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetVendorId returns the VendorId field value if set, zero value otherwise.
func (o *NiatelemetryDcnmModuleDetails) GetVendorId() string {
	if o == nil || o.VendorId == nil {
		var ret string
		return ret
	}
	return *o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmModuleDetails) GetVendorIdOk() (*string, bool) {
	if o == nil || o.VendorId == nil {
		return nil, false
	}
	return o.VendorId, true
}

// HasVendorId returns a boolean if a field has been set.
func (o *NiatelemetryDcnmModuleDetails) HasVendorId() bool {
	if o != nil && o.VendorId != nil {
		return true
	}

	return false
}

// SetVendorId gets a reference to the given string and assigns it to the VendorId field.
func (o *NiatelemetryDcnmModuleDetails) SetVendorId(v string) {
	o.VendorId = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryDcnmModuleDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmModuleDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryDcnmModuleDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryDcnmModuleDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryDcnmModuleDetails) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ProductId != nil {
		toSerialize["ProductId"] = o.ProductId
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.RecordVersion != nil {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if o.SerialNumber != nil {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if o.VendorId != nil {
		toSerialize["VendorId"] = o.VendorId
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryDcnmModuleDetails) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryDcnmModuleDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the hardware module in the fabric inventory.
		Name *string `json:"Name,omitempty"`
		// Product ID of the hardware module in the fabric inventory.
		ProductId *string `json:"ProductId,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Serial number of the hardware module in the fabric inventory.
		SerialNumber *string `json:"SerialNumber,omitempty"`
		// Vendor Id of the hardware module in the fabric inventory.
		VendorId         *string                              `json:"VendorId,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryDcnmModuleDetailsWithoutEmbeddedStruct := NiatelemetryDcnmModuleDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryDcnmModuleDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryDcnmModuleDetails := _NiatelemetryDcnmModuleDetails{}
		varNiatelemetryDcnmModuleDetails.ClassId = varNiatelemetryDcnmModuleDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryDcnmModuleDetails.ObjectType = varNiatelemetryDcnmModuleDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryDcnmModuleDetails.Name = varNiatelemetryDcnmModuleDetailsWithoutEmbeddedStruct.Name
		varNiatelemetryDcnmModuleDetails.ProductId = varNiatelemetryDcnmModuleDetailsWithoutEmbeddedStruct.ProductId
		varNiatelemetryDcnmModuleDetails.RecordType = varNiatelemetryDcnmModuleDetailsWithoutEmbeddedStruct.RecordType
		varNiatelemetryDcnmModuleDetails.RecordVersion = varNiatelemetryDcnmModuleDetailsWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryDcnmModuleDetails.SerialNumber = varNiatelemetryDcnmModuleDetailsWithoutEmbeddedStruct.SerialNumber
		varNiatelemetryDcnmModuleDetails.VendorId = varNiatelemetryDcnmModuleDetailsWithoutEmbeddedStruct.VendorId
		varNiatelemetryDcnmModuleDetails.RegisteredDevice = varNiatelemetryDcnmModuleDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryDcnmModuleDetails(varNiatelemetryDcnmModuleDetails)
	} else {
		return err
	}

	varNiatelemetryDcnmModuleDetails := _NiatelemetryDcnmModuleDetails{}

	err = json.Unmarshal(bytes, &varNiatelemetryDcnmModuleDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryDcnmModuleDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ProductId")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "VendorId")
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

type NullableNiatelemetryDcnmModuleDetails struct {
	value *NiatelemetryDcnmModuleDetails
	isSet bool
}

func (v NullableNiatelemetryDcnmModuleDetails) Get() *NiatelemetryDcnmModuleDetails {
	return v.value
}

func (v *NullableNiatelemetryDcnmModuleDetails) Set(val *NiatelemetryDcnmModuleDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryDcnmModuleDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryDcnmModuleDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryDcnmModuleDetails(val *NiatelemetryDcnmModuleDetails) *NullableNiatelemetryDcnmModuleDetails {
	return &NullableNiatelemetryDcnmModuleDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryDcnmModuleDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryDcnmModuleDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
