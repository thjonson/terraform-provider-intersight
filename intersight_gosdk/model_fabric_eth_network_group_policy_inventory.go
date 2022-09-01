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

// FabricEthNetworkGroupPolicyInventory The allowed VLAN/s on an interface.
type FabricEthNetworkGroupPolicyInventory struct {
	PolicyAbstractPolicyInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                     `json:"ObjectType"`
	VlanSettings         NullableFabricVlanSettings `json:"VlanSettings,omitempty"`
	TargetMo             *MoBaseMoRelationship      `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricEthNetworkGroupPolicyInventory FabricEthNetworkGroupPolicyInventory

// NewFabricEthNetworkGroupPolicyInventory instantiates a new FabricEthNetworkGroupPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricEthNetworkGroupPolicyInventory(classId string, objectType string) *FabricEthNetworkGroupPolicyInventory {
	this := FabricEthNetworkGroupPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricEthNetworkGroupPolicyInventoryWithDefaults instantiates a new FabricEthNetworkGroupPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricEthNetworkGroupPolicyInventoryWithDefaults() *FabricEthNetworkGroupPolicyInventory {
	this := FabricEthNetworkGroupPolicyInventory{}
	var classId string = "fabric.EthNetworkGroupPolicyInventory"
	this.ClassId = classId
	var objectType string = "fabric.EthNetworkGroupPolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricEthNetworkGroupPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkGroupPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricEthNetworkGroupPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricEthNetworkGroupPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkGroupPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricEthNetworkGroupPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVlanSettings returns the VlanSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricEthNetworkGroupPolicyInventory) GetVlanSettings() FabricVlanSettings {
	if o == nil || o.VlanSettings.Get() == nil {
		var ret FabricVlanSettings
		return ret
	}
	return *o.VlanSettings.Get()
}

// GetVlanSettingsOk returns a tuple with the VlanSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricEthNetworkGroupPolicyInventory) GetVlanSettingsOk() (*FabricVlanSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.VlanSettings.Get(), o.VlanSettings.IsSet()
}

// HasVlanSettings returns a boolean if a field has been set.
func (o *FabricEthNetworkGroupPolicyInventory) HasVlanSettings() bool {
	if o != nil && o.VlanSettings.IsSet() {
		return true
	}

	return false
}

// SetVlanSettings gets a reference to the given NullableFabricVlanSettings and assigns it to the VlanSettings field.
func (o *FabricEthNetworkGroupPolicyInventory) SetVlanSettings(v FabricVlanSettings) {
	o.VlanSettings.Set(&v)
}

// SetVlanSettingsNil sets the value for VlanSettings to be an explicit nil
func (o *FabricEthNetworkGroupPolicyInventory) SetVlanSettingsNil() {
	o.VlanSettings.Set(nil)
}

// UnsetVlanSettings ensures that no value is present for VlanSettings, not even an explicit nil
func (o *FabricEthNetworkGroupPolicyInventory) UnsetVlanSettings() {
	o.VlanSettings.Unset()
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *FabricEthNetworkGroupPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkGroupPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *FabricEthNetworkGroupPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *FabricEthNetworkGroupPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o FabricEthNetworkGroupPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicyInventory, errPolicyAbstractPolicyInventory := json.Marshal(o.PolicyAbstractPolicyInventory)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
	errPolicyAbstractPolicyInventory = json.Unmarshal([]byte(serializedPolicyAbstractPolicyInventory), &toSerialize)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.VlanSettings.IsSet() {
		toSerialize["VlanSettings"] = o.VlanSettings.Get()
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricEthNetworkGroupPolicyInventory) UnmarshalJSON(bytes []byte) (err error) {
	type FabricEthNetworkGroupPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string                     `json:"ObjectType"`
		VlanSettings NullableFabricVlanSettings `json:"VlanSettings,omitempty"`
		TargetMo     *MoBaseMoRelationship      `json:"TargetMo,omitempty"`
	}

	varFabricEthNetworkGroupPolicyInventoryWithoutEmbeddedStruct := FabricEthNetworkGroupPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricEthNetworkGroupPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varFabricEthNetworkGroupPolicyInventory := _FabricEthNetworkGroupPolicyInventory{}
		varFabricEthNetworkGroupPolicyInventory.ClassId = varFabricEthNetworkGroupPolicyInventoryWithoutEmbeddedStruct.ClassId
		varFabricEthNetworkGroupPolicyInventory.ObjectType = varFabricEthNetworkGroupPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varFabricEthNetworkGroupPolicyInventory.VlanSettings = varFabricEthNetworkGroupPolicyInventoryWithoutEmbeddedStruct.VlanSettings
		varFabricEthNetworkGroupPolicyInventory.TargetMo = varFabricEthNetworkGroupPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = FabricEthNetworkGroupPolicyInventory(varFabricEthNetworkGroupPolicyInventory)
	} else {
		return err
	}

	varFabricEthNetworkGroupPolicyInventory := _FabricEthNetworkGroupPolicyInventory{}

	err = json.Unmarshal(bytes, &varFabricEthNetworkGroupPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varFabricEthNetworkGroupPolicyInventory.PolicyAbstractPolicyInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "VlanSettings")
		delete(additionalProperties, "TargetMo")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicyInventory := reflect.ValueOf(o.PolicyAbstractPolicyInventory)
		for i := 0; i < reflectPolicyAbstractPolicyInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicyInventory.Type().Field(i)

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

type NullableFabricEthNetworkGroupPolicyInventory struct {
	value *FabricEthNetworkGroupPolicyInventory
	isSet bool
}

func (v NullableFabricEthNetworkGroupPolicyInventory) Get() *FabricEthNetworkGroupPolicyInventory {
	return v.value
}

func (v *NullableFabricEthNetworkGroupPolicyInventory) Set(val *FabricEthNetworkGroupPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricEthNetworkGroupPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricEthNetworkGroupPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricEthNetworkGroupPolicyInventory(val *FabricEthNetworkGroupPolicyInventory) *NullableFabricEthNetworkGroupPolicyInventory {
	return &NullableFabricEthNetworkGroupPolicyInventory{value: val, isSet: true}
}

func (v NullableFabricEthNetworkGroupPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricEthNetworkGroupPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
