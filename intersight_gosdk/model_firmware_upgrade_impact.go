/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-07-21T16:37:30Z.
 *
 * API version: 1.0.9-4403
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// FirmwareUpgradeImpact Before submitting firmware upgrade operation, the estimate impact helps to know the list of components be impacted and require host reboot. This cannot be used for network share based upgrade.
type FirmwareUpgradeImpact struct {
	FirmwareUpgradeImpactBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// An array of relationships to equipmentChassis resources.
	Chassis []EquipmentChassisRelationship `json:"Chassis,omitempty"`
	// An array of relationships to assetDeviceRegistration resources.
	Device        []AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	Distributable *FirmwareDistributableRelationship    `json:"Distributable,omitempty"`
	// An array of relationships to networkElement resources.
	NetworkElements []NetworkElementRelationship           `json:"NetworkElements,omitempty"`
	Release         *SoftwarerepositoryReleaseRelationship `json:"Release,omitempty"`
	// An array of relationships to computePhysical resources.
	Server               []ComputePhysicalRelationship `json:"Server,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareUpgradeImpact FirmwareUpgradeImpact

// NewFirmwareUpgradeImpact instantiates a new FirmwareUpgradeImpact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareUpgradeImpact(classId string, objectType string) *FirmwareUpgradeImpact {
	this := FirmwareUpgradeImpact{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwareUpgradeImpactWithDefaults instantiates a new FirmwareUpgradeImpact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareUpgradeImpactWithDefaults() *FirmwareUpgradeImpact {
	this := FirmwareUpgradeImpact{}
	var classId string = "firmware.UpgradeImpact"
	this.ClassId = classId
	var objectType string = "firmware.UpgradeImpact"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareUpgradeImpact) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeImpact) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareUpgradeImpact) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareUpgradeImpact) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeImpact) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareUpgradeImpact) SetObjectType(v string) {
	o.ObjectType = v
}

// GetChassis returns the Chassis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareUpgradeImpact) GetChassis() []EquipmentChassisRelationship {
	if o == nil {
		var ret []EquipmentChassisRelationship
		return ret
	}
	return o.Chassis
}

// GetChassisOk returns a tuple with the Chassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareUpgradeImpact) GetChassisOk() (*[]EquipmentChassisRelationship, bool) {
	if o == nil || o.Chassis == nil {
		return nil, false
	}
	return &o.Chassis, true
}

// HasChassis returns a boolean if a field has been set.
func (o *FirmwareUpgradeImpact) HasChassis() bool {
	if o != nil && o.Chassis != nil {
		return true
	}

	return false
}

// SetChassis gets a reference to the given []EquipmentChassisRelationship and assigns it to the Chassis field.
func (o *FirmwareUpgradeImpact) SetChassis(v []EquipmentChassisRelationship) {
	o.Chassis = v
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareUpgradeImpact) GetDevice() []AssetDeviceRegistrationRelationship {
	if o == nil {
		var ret []AssetDeviceRegistrationRelationship
		return ret
	}
	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareUpgradeImpact) GetDeviceOk() (*[]AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return &o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *FirmwareUpgradeImpact) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given []AssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *FirmwareUpgradeImpact) SetDevice(v []AssetDeviceRegistrationRelationship) {
	o.Device = v
}

// GetDistributable returns the Distributable field value if set, zero value otherwise.
func (o *FirmwareUpgradeImpact) GetDistributable() FirmwareDistributableRelationship {
	if o == nil || o.Distributable == nil {
		var ret FirmwareDistributableRelationship
		return ret
	}
	return *o.Distributable
}

// GetDistributableOk returns a tuple with the Distributable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeImpact) GetDistributableOk() (*FirmwareDistributableRelationship, bool) {
	if o == nil || o.Distributable == nil {
		return nil, false
	}
	return o.Distributable, true
}

// HasDistributable returns a boolean if a field has been set.
func (o *FirmwareUpgradeImpact) HasDistributable() bool {
	if o != nil && o.Distributable != nil {
		return true
	}

	return false
}

// SetDistributable gets a reference to the given FirmwareDistributableRelationship and assigns it to the Distributable field.
func (o *FirmwareUpgradeImpact) SetDistributable(v FirmwareDistributableRelationship) {
	o.Distributable = &v
}

// GetNetworkElements returns the NetworkElements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareUpgradeImpact) GetNetworkElements() []NetworkElementRelationship {
	if o == nil {
		var ret []NetworkElementRelationship
		return ret
	}
	return o.NetworkElements
}

// GetNetworkElementsOk returns a tuple with the NetworkElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareUpgradeImpact) GetNetworkElementsOk() (*[]NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElements == nil {
		return nil, false
	}
	return &o.NetworkElements, true
}

// HasNetworkElements returns a boolean if a field has been set.
func (o *FirmwareUpgradeImpact) HasNetworkElements() bool {
	if o != nil && o.NetworkElements != nil {
		return true
	}

	return false
}

// SetNetworkElements gets a reference to the given []NetworkElementRelationship and assigns it to the NetworkElements field.
func (o *FirmwareUpgradeImpact) SetNetworkElements(v []NetworkElementRelationship) {
	o.NetworkElements = v
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *FirmwareUpgradeImpact) GetRelease() SoftwarerepositoryReleaseRelationship {
	if o == nil || o.Release == nil {
		var ret SoftwarerepositoryReleaseRelationship
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeImpact) GetReleaseOk() (*SoftwarerepositoryReleaseRelationship, bool) {
	if o == nil || o.Release == nil {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *FirmwareUpgradeImpact) HasRelease() bool {
	if o != nil && o.Release != nil {
		return true
	}

	return false
}

// SetRelease gets a reference to the given SoftwarerepositoryReleaseRelationship and assigns it to the Release field.
func (o *FirmwareUpgradeImpact) SetRelease(v SoftwarerepositoryReleaseRelationship) {
	o.Release = &v
}

// GetServer returns the Server field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareUpgradeImpact) GetServer() []ComputePhysicalRelationship {
	if o == nil {
		var ret []ComputePhysicalRelationship
		return ret
	}
	return o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareUpgradeImpact) GetServerOk() (*[]ComputePhysicalRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return &o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *FirmwareUpgradeImpact) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given []ComputePhysicalRelationship and assigns it to the Server field.
func (o *FirmwareUpgradeImpact) SetServer(v []ComputePhysicalRelationship) {
	o.Server = v
}

func (o FirmwareUpgradeImpact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFirmwareUpgradeImpactBase, errFirmwareUpgradeImpactBase := json.Marshal(o.FirmwareUpgradeImpactBase)
	if errFirmwareUpgradeImpactBase != nil {
		return []byte{}, errFirmwareUpgradeImpactBase
	}
	errFirmwareUpgradeImpactBase = json.Unmarshal([]byte(serializedFirmwareUpgradeImpactBase), &toSerialize)
	if errFirmwareUpgradeImpactBase != nil {
		return []byte{}, errFirmwareUpgradeImpactBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Chassis != nil {
		toSerialize["Chassis"] = o.Chassis
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}
	if o.Distributable != nil {
		toSerialize["Distributable"] = o.Distributable
	}
	if o.NetworkElements != nil {
		toSerialize["NetworkElements"] = o.NetworkElements
	}
	if o.Release != nil {
		toSerialize["Release"] = o.Release
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareUpgradeImpact) UnmarshalJSON(bytes []byte) (err error) {
	type FirmwareUpgradeImpactWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// An array of relationships to equipmentChassis resources.
		Chassis []EquipmentChassisRelationship `json:"Chassis,omitempty"`
		// An array of relationships to assetDeviceRegistration resources.
		Device        []AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
		Distributable *FirmwareDistributableRelationship    `json:"Distributable,omitempty"`
		// An array of relationships to networkElement resources.
		NetworkElements []NetworkElementRelationship           `json:"NetworkElements,omitempty"`
		Release         *SoftwarerepositoryReleaseRelationship `json:"Release,omitempty"`
		// An array of relationships to computePhysical resources.
		Server []ComputePhysicalRelationship `json:"Server,omitempty"`
	}

	varFirmwareUpgradeImpactWithoutEmbeddedStruct := FirmwareUpgradeImpactWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFirmwareUpgradeImpactWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareUpgradeImpact := _FirmwareUpgradeImpact{}
		varFirmwareUpgradeImpact.ClassId = varFirmwareUpgradeImpactWithoutEmbeddedStruct.ClassId
		varFirmwareUpgradeImpact.ObjectType = varFirmwareUpgradeImpactWithoutEmbeddedStruct.ObjectType
		varFirmwareUpgradeImpact.Chassis = varFirmwareUpgradeImpactWithoutEmbeddedStruct.Chassis
		varFirmwareUpgradeImpact.Device = varFirmwareUpgradeImpactWithoutEmbeddedStruct.Device
		varFirmwareUpgradeImpact.Distributable = varFirmwareUpgradeImpactWithoutEmbeddedStruct.Distributable
		varFirmwareUpgradeImpact.NetworkElements = varFirmwareUpgradeImpactWithoutEmbeddedStruct.NetworkElements
		varFirmwareUpgradeImpact.Release = varFirmwareUpgradeImpactWithoutEmbeddedStruct.Release
		varFirmwareUpgradeImpact.Server = varFirmwareUpgradeImpactWithoutEmbeddedStruct.Server
		*o = FirmwareUpgradeImpact(varFirmwareUpgradeImpact)
	} else {
		return err
	}

	varFirmwareUpgradeImpact := _FirmwareUpgradeImpact{}

	err = json.Unmarshal(bytes, &varFirmwareUpgradeImpact)
	if err == nil {
		o.FirmwareUpgradeImpactBase = varFirmwareUpgradeImpact.FirmwareUpgradeImpactBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Chassis")
		delete(additionalProperties, "Device")
		delete(additionalProperties, "Distributable")
		delete(additionalProperties, "NetworkElements")
		delete(additionalProperties, "Release")
		delete(additionalProperties, "Server")

		// remove fields from embedded structs
		reflectFirmwareUpgradeImpactBase := reflect.ValueOf(o.FirmwareUpgradeImpactBase)
		for i := 0; i < reflectFirmwareUpgradeImpactBase.Type().NumField(); i++ {
			t := reflectFirmwareUpgradeImpactBase.Type().Field(i)

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

type NullableFirmwareUpgradeImpact struct {
	value *FirmwareUpgradeImpact
	isSet bool
}

func (v NullableFirmwareUpgradeImpact) Get() *FirmwareUpgradeImpact {
	return v.value
}

func (v *NullableFirmwareUpgradeImpact) Set(val *FirmwareUpgradeImpact) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareUpgradeImpact) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareUpgradeImpact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareUpgradeImpact(val *FirmwareUpgradeImpact) *NullableFirmwareUpgradeImpact {
	return &NullableFirmwareUpgradeImpact{value: val, isSet: true}
}

func (v NullableFirmwareUpgradeImpact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareUpgradeImpact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
