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
	"reflect"
	"strings"
)

// VirtualizationVmwareVirtualNetworkInterface Details of VMware virtual network interface.
type VirtualizationVmwareVirtualNetworkInterface struct {
	VirtualizationBaseVirtualNetworkInterface
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of virtual ethernet adapter for virtual network interface.
	AdapterType *string `json:"AdapterType,omitempty"`
	// Connect or not to connect the device when the virtual machine starts.
	ConnectAtPowerOn *bool `json:"ConnectAtPowerOn,omitempty"`
	// Device is currently connected or not. Valid only while the virtual machine is running.
	Connected *bool `json:"Connected,omitempty"`
	// The internally assigned key of this virtual network interface. This entity is not manipulated by users.
	Key *int64 `json:"Key,omitempty"`
	// MAC address assigned to virtual network interface.
	MacAddress *string `json:"MacAddress,omitempty"`
	// MAC address type for the mac address assigned to virtual network interface. * `manual` - Statically assigned MAC address. * `generated` - Automatically generated MAC address. * `assigned` - MAC address assigned by VCenter to the virtual network interface card.
	MacAddressType *string `json:"MacAddressType,omitempty"`
	// Type of network for virtual network interface. It can be either standard or distributed.
	NetworkType *string `json:"NetworkType,omitempty"`
	// Identity of the virtual machine where the virtual network interface is created.
	VmIdentity           *string                                         `json:"VmIdentity,omitempty"`
	Network              *VirtualizationBaseNetworkRelationship          `json:"Network,omitempty"`
	VirtualMachine       *VirtualizationVmwareVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareVirtualNetworkInterface VirtualizationVmwareVirtualNetworkInterface

// NewVirtualizationVmwareVirtualNetworkInterface instantiates a new VirtualizationVmwareVirtualNetworkInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareVirtualNetworkInterface(classId string, objectType string) *VirtualizationVmwareVirtualNetworkInterface {
	this := VirtualizationVmwareVirtualNetworkInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	var macAddressType string = "manual"
	this.MacAddressType = &macAddressType
	return &this
}

// NewVirtualizationVmwareVirtualNetworkInterfaceWithDefaults instantiates a new VirtualizationVmwareVirtualNetworkInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareVirtualNetworkInterfaceWithDefaults() *VirtualizationVmwareVirtualNetworkInterface {
	this := VirtualizationVmwareVirtualNetworkInterface{}
	var classId string = "virtualization.VmwareVirtualNetworkInterface"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareVirtualNetworkInterface"
	this.ObjectType = objectType
	var macAddressType string = "manual"
	this.MacAddressType = &macAddressType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareVirtualNetworkInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualNetworkInterface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareVirtualNetworkInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareVirtualNetworkInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualNetworkInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareVirtualNetworkInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdapterType returns the AdapterType field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualNetworkInterface) GetAdapterType() string {
	if o == nil || o.AdapterType == nil {
		var ret string
		return ret
	}
	return *o.AdapterType
}

// GetAdapterTypeOk returns a tuple with the AdapterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualNetworkInterface) GetAdapterTypeOk() (*string, bool) {
	if o == nil || o.AdapterType == nil {
		return nil, false
	}
	return o.AdapterType, true
}

// HasAdapterType returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualNetworkInterface) HasAdapterType() bool {
	if o != nil && o.AdapterType != nil {
		return true
	}

	return false
}

// SetAdapterType gets a reference to the given string and assigns it to the AdapterType field.
func (o *VirtualizationVmwareVirtualNetworkInterface) SetAdapterType(v string) {
	o.AdapterType = &v
}

// GetConnectAtPowerOn returns the ConnectAtPowerOn field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualNetworkInterface) GetConnectAtPowerOn() bool {
	if o == nil || o.ConnectAtPowerOn == nil {
		var ret bool
		return ret
	}
	return *o.ConnectAtPowerOn
}

// GetConnectAtPowerOnOk returns a tuple with the ConnectAtPowerOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualNetworkInterface) GetConnectAtPowerOnOk() (*bool, bool) {
	if o == nil || o.ConnectAtPowerOn == nil {
		return nil, false
	}
	return o.ConnectAtPowerOn, true
}

// HasConnectAtPowerOn returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualNetworkInterface) HasConnectAtPowerOn() bool {
	if o != nil && o.ConnectAtPowerOn != nil {
		return true
	}

	return false
}

// SetConnectAtPowerOn gets a reference to the given bool and assigns it to the ConnectAtPowerOn field.
func (o *VirtualizationVmwareVirtualNetworkInterface) SetConnectAtPowerOn(v bool) {
	o.ConnectAtPowerOn = &v
}

// GetConnected returns the Connected field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualNetworkInterface) GetConnected() bool {
	if o == nil || o.Connected == nil {
		var ret bool
		return ret
	}
	return *o.Connected
}

// GetConnectedOk returns a tuple with the Connected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualNetworkInterface) GetConnectedOk() (*bool, bool) {
	if o == nil || o.Connected == nil {
		return nil, false
	}
	return o.Connected, true
}

// HasConnected returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualNetworkInterface) HasConnected() bool {
	if o != nil && o.Connected != nil {
		return true
	}

	return false
}

// SetConnected gets a reference to the given bool and assigns it to the Connected field.
func (o *VirtualizationVmwareVirtualNetworkInterface) SetConnected(v bool) {
	o.Connected = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualNetworkInterface) GetKey() int64 {
	if o == nil || o.Key == nil {
		var ret int64
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualNetworkInterface) GetKeyOk() (*int64, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualNetworkInterface) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given int64 and assigns it to the Key field.
func (o *VirtualizationVmwareVirtualNetworkInterface) SetKey(v int64) {
	o.Key = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualNetworkInterface) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualNetworkInterface) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualNetworkInterface) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *VirtualizationVmwareVirtualNetworkInterface) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetMacAddressType returns the MacAddressType field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualNetworkInterface) GetMacAddressType() string {
	if o == nil || o.MacAddressType == nil {
		var ret string
		return ret
	}
	return *o.MacAddressType
}

// GetMacAddressTypeOk returns a tuple with the MacAddressType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualNetworkInterface) GetMacAddressTypeOk() (*string, bool) {
	if o == nil || o.MacAddressType == nil {
		return nil, false
	}
	return o.MacAddressType, true
}

// HasMacAddressType returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualNetworkInterface) HasMacAddressType() bool {
	if o != nil && o.MacAddressType != nil {
		return true
	}

	return false
}

// SetMacAddressType gets a reference to the given string and assigns it to the MacAddressType field.
func (o *VirtualizationVmwareVirtualNetworkInterface) SetMacAddressType(v string) {
	o.MacAddressType = &v
}

// GetNetworkType returns the NetworkType field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualNetworkInterface) GetNetworkType() string {
	if o == nil || o.NetworkType == nil {
		var ret string
		return ret
	}
	return *o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualNetworkInterface) GetNetworkTypeOk() (*string, bool) {
	if o == nil || o.NetworkType == nil {
		return nil, false
	}
	return o.NetworkType, true
}

// HasNetworkType returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualNetworkInterface) HasNetworkType() bool {
	if o != nil && o.NetworkType != nil {
		return true
	}

	return false
}

// SetNetworkType gets a reference to the given string and assigns it to the NetworkType field.
func (o *VirtualizationVmwareVirtualNetworkInterface) SetNetworkType(v string) {
	o.NetworkType = &v
}

// GetVmIdentity returns the VmIdentity field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualNetworkInterface) GetVmIdentity() string {
	if o == nil || o.VmIdentity == nil {
		var ret string
		return ret
	}
	return *o.VmIdentity
}

// GetVmIdentityOk returns a tuple with the VmIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualNetworkInterface) GetVmIdentityOk() (*string, bool) {
	if o == nil || o.VmIdentity == nil {
		return nil, false
	}
	return o.VmIdentity, true
}

// HasVmIdentity returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualNetworkInterface) HasVmIdentity() bool {
	if o != nil && o.VmIdentity != nil {
		return true
	}

	return false
}

// SetVmIdentity gets a reference to the given string and assigns it to the VmIdentity field.
func (o *VirtualizationVmwareVirtualNetworkInterface) SetVmIdentity(v string) {
	o.VmIdentity = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualNetworkInterface) GetNetwork() VirtualizationBaseNetworkRelationship {
	if o == nil || o.Network == nil {
		var ret VirtualizationBaseNetworkRelationship
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualNetworkInterface) GetNetworkOk() (*VirtualizationBaseNetworkRelationship, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualNetworkInterface) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given VirtualizationBaseNetworkRelationship and assigns it to the Network field.
func (o *VirtualizationVmwareVirtualNetworkInterface) SetNetwork(v VirtualizationBaseNetworkRelationship) {
	o.Network = &v
}

// GetVirtualMachine returns the VirtualMachine field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualNetworkInterface) GetVirtualMachine() VirtualizationVmwareVirtualMachineRelationship {
	if o == nil || o.VirtualMachine == nil {
		var ret VirtualizationVmwareVirtualMachineRelationship
		return ret
	}
	return *o.VirtualMachine
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualNetworkInterface) GetVirtualMachineOk() (*VirtualizationVmwareVirtualMachineRelationship, bool) {
	if o == nil || o.VirtualMachine == nil {
		return nil, false
	}
	return o.VirtualMachine, true
}

// HasVirtualMachine returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualNetworkInterface) HasVirtualMachine() bool {
	if o != nil && o.VirtualMachine != nil {
		return true
	}

	return false
}

// SetVirtualMachine gets a reference to the given VirtualizationVmwareVirtualMachineRelationship and assigns it to the VirtualMachine field.
func (o *VirtualizationVmwareVirtualNetworkInterface) SetVirtualMachine(v VirtualizationVmwareVirtualMachineRelationship) {
	o.VirtualMachine = &v
}

func (o VirtualizationVmwareVirtualNetworkInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseVirtualNetworkInterface, errVirtualizationBaseVirtualNetworkInterface := json.Marshal(o.VirtualizationBaseVirtualNetworkInterface)
	if errVirtualizationBaseVirtualNetworkInterface != nil {
		return []byte{}, errVirtualizationBaseVirtualNetworkInterface
	}
	errVirtualizationBaseVirtualNetworkInterface = json.Unmarshal([]byte(serializedVirtualizationBaseVirtualNetworkInterface), &toSerialize)
	if errVirtualizationBaseVirtualNetworkInterface != nil {
		return []byte{}, errVirtualizationBaseVirtualNetworkInterface
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdapterType != nil {
		toSerialize["AdapterType"] = o.AdapterType
	}
	if o.ConnectAtPowerOn != nil {
		toSerialize["ConnectAtPowerOn"] = o.ConnectAtPowerOn
	}
	if o.Connected != nil {
		toSerialize["Connected"] = o.Connected
	}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.MacAddressType != nil {
		toSerialize["MacAddressType"] = o.MacAddressType
	}
	if o.NetworkType != nil {
		toSerialize["NetworkType"] = o.NetworkType
	}
	if o.VmIdentity != nil {
		toSerialize["VmIdentity"] = o.VmIdentity
	}
	if o.Network != nil {
		toSerialize["Network"] = o.Network
	}
	if o.VirtualMachine != nil {
		toSerialize["VirtualMachine"] = o.VirtualMachine
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareVirtualNetworkInterface) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationVmwareVirtualNetworkInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Type of virtual ethernet adapter for virtual network interface.
		AdapterType *string `json:"AdapterType,omitempty"`
		// Connect or not to connect the device when the virtual machine starts.
		ConnectAtPowerOn *bool `json:"ConnectAtPowerOn,omitempty"`
		// Device is currently connected or not. Valid only while the virtual machine is running.
		Connected *bool `json:"Connected,omitempty"`
		// The internally assigned key of this virtual network interface. This entity is not manipulated by users.
		Key *int64 `json:"Key,omitempty"`
		// MAC address assigned to virtual network interface.
		MacAddress *string `json:"MacAddress,omitempty"`
		// MAC address type for the mac address assigned to virtual network interface. * `manual` - Statically assigned MAC address. * `generated` - Automatically generated MAC address. * `assigned` - MAC address assigned by VCenter to the virtual network interface card.
		MacAddressType *string `json:"MacAddressType,omitempty"`
		// Type of network for virtual network interface. It can be either standard or distributed.
		NetworkType *string `json:"NetworkType,omitempty"`
		// Identity of the virtual machine where the virtual network interface is created.
		VmIdentity     *string                                         `json:"VmIdentity,omitempty"`
		Network        *VirtualizationBaseNetworkRelationship          `json:"Network,omitempty"`
		VirtualMachine *VirtualizationVmwareVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	}

	varVirtualizationVmwareVirtualNetworkInterfaceWithoutEmbeddedStruct := VirtualizationVmwareVirtualNetworkInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareVirtualNetworkInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareVirtualNetworkInterface := _VirtualizationVmwareVirtualNetworkInterface{}
		varVirtualizationVmwareVirtualNetworkInterface.ClassId = varVirtualizationVmwareVirtualNetworkInterfaceWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareVirtualNetworkInterface.ObjectType = varVirtualizationVmwareVirtualNetworkInterfaceWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareVirtualNetworkInterface.AdapterType = varVirtualizationVmwareVirtualNetworkInterfaceWithoutEmbeddedStruct.AdapterType
		varVirtualizationVmwareVirtualNetworkInterface.ConnectAtPowerOn = varVirtualizationVmwareVirtualNetworkInterfaceWithoutEmbeddedStruct.ConnectAtPowerOn
		varVirtualizationVmwareVirtualNetworkInterface.Connected = varVirtualizationVmwareVirtualNetworkInterfaceWithoutEmbeddedStruct.Connected
		varVirtualizationVmwareVirtualNetworkInterface.Key = varVirtualizationVmwareVirtualNetworkInterfaceWithoutEmbeddedStruct.Key
		varVirtualizationVmwareVirtualNetworkInterface.MacAddress = varVirtualizationVmwareVirtualNetworkInterfaceWithoutEmbeddedStruct.MacAddress
		varVirtualizationVmwareVirtualNetworkInterface.MacAddressType = varVirtualizationVmwareVirtualNetworkInterfaceWithoutEmbeddedStruct.MacAddressType
		varVirtualizationVmwareVirtualNetworkInterface.NetworkType = varVirtualizationVmwareVirtualNetworkInterfaceWithoutEmbeddedStruct.NetworkType
		varVirtualizationVmwareVirtualNetworkInterface.VmIdentity = varVirtualizationVmwareVirtualNetworkInterfaceWithoutEmbeddedStruct.VmIdentity
		varVirtualizationVmwareVirtualNetworkInterface.Network = varVirtualizationVmwareVirtualNetworkInterfaceWithoutEmbeddedStruct.Network
		varVirtualizationVmwareVirtualNetworkInterface.VirtualMachine = varVirtualizationVmwareVirtualNetworkInterfaceWithoutEmbeddedStruct.VirtualMachine
		*o = VirtualizationVmwareVirtualNetworkInterface(varVirtualizationVmwareVirtualNetworkInterface)
	} else {
		return err
	}

	varVirtualizationVmwareVirtualNetworkInterface := _VirtualizationVmwareVirtualNetworkInterface{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareVirtualNetworkInterface)
	if err == nil {
		o.VirtualizationBaseVirtualNetworkInterface = varVirtualizationVmwareVirtualNetworkInterface.VirtualizationBaseVirtualNetworkInterface
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdapterType")
		delete(additionalProperties, "ConnectAtPowerOn")
		delete(additionalProperties, "Connected")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "MacAddressType")
		delete(additionalProperties, "NetworkType")
		delete(additionalProperties, "VmIdentity")
		delete(additionalProperties, "Network")
		delete(additionalProperties, "VirtualMachine")

		// remove fields from embedded structs
		reflectVirtualizationBaseVirtualNetworkInterface := reflect.ValueOf(o.VirtualizationBaseVirtualNetworkInterface)
		for i := 0; i < reflectVirtualizationBaseVirtualNetworkInterface.Type().NumField(); i++ {
			t := reflectVirtualizationBaseVirtualNetworkInterface.Type().Field(i)

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

type NullableVirtualizationVmwareVirtualNetworkInterface struct {
	value *VirtualizationVmwareVirtualNetworkInterface
	isSet bool
}

func (v NullableVirtualizationVmwareVirtualNetworkInterface) Get() *VirtualizationVmwareVirtualNetworkInterface {
	return v.value
}

func (v *NullableVirtualizationVmwareVirtualNetworkInterface) Set(val *VirtualizationVmwareVirtualNetworkInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareVirtualNetworkInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareVirtualNetworkInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareVirtualNetworkInterface(val *VirtualizationVmwareVirtualNetworkInterface) *NullableVirtualizationVmwareVirtualNetworkInterface {
	return &NullableVirtualizationVmwareVirtualNetworkInterface{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareVirtualNetworkInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareVirtualNetworkInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
