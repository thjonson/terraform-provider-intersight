/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4663
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// InventoryGenericInventoryAllOf Definition of the list of properties defined in 'inventory.GenericInventory', excluding properties defined in parent classes.
type InventoryGenericInventoryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Key of inventory data for Generic Inventory data set.
	Key *string `json:"Key,omitempty"`
	// Type of inventory data for Generic Inventory data set.
	Type *string `json:"Type,omitempty"`
	// Value of inventory data for Generic Inventory data set.
	Value                           *string                                      `json:"Value,omitempty"`
	InventoryDeviceInfo             *InventoryDeviceInfoRelationship             `json:"InventoryDeviceInfo,omitempty"`
	InventoryGenericInventoryHolder *InventoryGenericInventoryHolderRelationship `json:"InventoryGenericInventoryHolder,omitempty"`
	RegisteredDevice                *AssetDeviceRegistrationRelationship         `json:"RegisteredDevice,omitempty"`
	AdditionalProperties            map[string]interface{}
}

type _InventoryGenericInventoryAllOf InventoryGenericInventoryAllOf

// NewInventoryGenericInventoryAllOf instantiates a new InventoryGenericInventoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryGenericInventoryAllOf(classId string, objectType string) *InventoryGenericInventoryAllOf {
	this := InventoryGenericInventoryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewInventoryGenericInventoryAllOfWithDefaults instantiates a new InventoryGenericInventoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryGenericInventoryAllOfWithDefaults() *InventoryGenericInventoryAllOf {
	this := InventoryGenericInventoryAllOf{}
	var classId string = "inventory.GenericInventory"
	this.ClassId = classId
	var objectType string = "inventory.GenericInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *InventoryGenericInventoryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *InventoryGenericInventoryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *InventoryGenericInventoryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *InventoryGenericInventoryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *InventoryGenericInventoryAllOf) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryAllOf) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *InventoryGenericInventoryAllOf) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *InventoryGenericInventoryAllOf) SetKey(v string) {
	o.Key = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InventoryGenericInventoryAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InventoryGenericInventoryAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InventoryGenericInventoryAllOf) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InventoryGenericInventoryAllOf) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryAllOf) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InventoryGenericInventoryAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InventoryGenericInventoryAllOf) SetValue(v string) {
	o.Value = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *InventoryGenericInventoryAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *InventoryGenericInventoryAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *InventoryGenericInventoryAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetInventoryGenericInventoryHolder returns the InventoryGenericInventoryHolder field value if set, zero value otherwise.
func (o *InventoryGenericInventoryAllOf) GetInventoryGenericInventoryHolder() InventoryGenericInventoryHolderRelationship {
	if o == nil || o.InventoryGenericInventoryHolder == nil {
		var ret InventoryGenericInventoryHolderRelationship
		return ret
	}
	return *o.InventoryGenericInventoryHolder
}

// GetInventoryGenericInventoryHolderOk returns a tuple with the InventoryGenericInventoryHolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryAllOf) GetInventoryGenericInventoryHolderOk() (*InventoryGenericInventoryHolderRelationship, bool) {
	if o == nil || o.InventoryGenericInventoryHolder == nil {
		return nil, false
	}
	return o.InventoryGenericInventoryHolder, true
}

// HasInventoryGenericInventoryHolder returns a boolean if a field has been set.
func (o *InventoryGenericInventoryAllOf) HasInventoryGenericInventoryHolder() bool {
	if o != nil && o.InventoryGenericInventoryHolder != nil {
		return true
	}

	return false
}

// SetInventoryGenericInventoryHolder gets a reference to the given InventoryGenericInventoryHolderRelationship and assigns it to the InventoryGenericInventoryHolder field.
func (o *InventoryGenericInventoryAllOf) SetInventoryGenericInventoryHolder(v InventoryGenericInventoryHolderRelationship) {
	o.InventoryGenericInventoryHolder = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *InventoryGenericInventoryAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *InventoryGenericInventoryAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *InventoryGenericInventoryAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o InventoryGenericInventoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.InventoryGenericInventoryHolder != nil {
		toSerialize["InventoryGenericInventoryHolder"] = o.InventoryGenericInventoryHolder
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InventoryGenericInventoryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varInventoryGenericInventoryAllOf := _InventoryGenericInventoryAllOf{}

	if err = json.Unmarshal(bytes, &varInventoryGenericInventoryAllOf); err == nil {
		*o = InventoryGenericInventoryAllOf(varInventoryGenericInventoryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Value")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "InventoryGenericInventoryHolder")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInventoryGenericInventoryAllOf struct {
	value *InventoryGenericInventoryAllOf
	isSet bool
}

func (v NullableInventoryGenericInventoryAllOf) Get() *InventoryGenericInventoryAllOf {
	return v.value
}

func (v *NullableInventoryGenericInventoryAllOf) Set(val *InventoryGenericInventoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryGenericInventoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryGenericInventoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryGenericInventoryAllOf(val *InventoryGenericInventoryAllOf) *NullableInventoryGenericInventoryAllOf {
	return &NullableInventoryGenericInventoryAllOf{value: val, isSet: true}
}

func (v NullableInventoryGenericInventoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryGenericInventoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
