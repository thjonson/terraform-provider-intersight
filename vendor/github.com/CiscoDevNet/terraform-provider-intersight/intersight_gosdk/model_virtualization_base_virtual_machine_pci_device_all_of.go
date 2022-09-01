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

// VirtualizationBaseVirtualMachinePciDeviceAllOf Definition of the list of properties defined in 'virtualization.BaseVirtualMachinePciDevice', excluding properties defined in parent classes.
type VirtualizationBaseVirtualMachinePciDeviceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The backing physical host PCI device Id for this device.
	BackingPciId *string `json:"BackingPciId,omitempty"`
	// Name of this virtual machine PCI device.
	Name *string `json:"Name,omitempty"`
	// Indicates if this virtual machine PCI device is enabled via passthrough from the host.
	Passthrough          *bool                                         `json:"Passthrough,omitempty"`
	BackingPciDevice     *VirtualizationBaseHostPciDeviceRelationship  `json:"BackingPciDevice,omitempty"`
	VirtualMachine       *VirtualizationBaseVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationBaseVirtualMachinePciDeviceAllOf VirtualizationBaseVirtualMachinePciDeviceAllOf

// NewVirtualizationBaseVirtualMachinePciDeviceAllOf instantiates a new VirtualizationBaseVirtualMachinePciDeviceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationBaseVirtualMachinePciDeviceAllOf(classId string, objectType string) *VirtualizationBaseVirtualMachinePciDeviceAllOf {
	this := VirtualizationBaseVirtualMachinePciDeviceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationBaseVirtualMachinePciDeviceAllOfWithDefaults instantiates a new VirtualizationBaseVirtualMachinePciDeviceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationBaseVirtualMachinePciDeviceAllOfWithDefaults() *VirtualizationBaseVirtualMachinePciDeviceAllOf {
	this := VirtualizationBaseVirtualMachinePciDeviceAllOf{}
	var classId string = "virtualization.VmwareVirtualMachineGpu"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareVirtualMachineGpu"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBackingPciId returns the BackingPciId field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) GetBackingPciId() string {
	if o == nil || o.BackingPciId == nil {
		var ret string
		return ret
	}
	return *o.BackingPciId
}

// GetBackingPciIdOk returns a tuple with the BackingPciId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) GetBackingPciIdOk() (*string, bool) {
	if o == nil || o.BackingPciId == nil {
		return nil, false
	}
	return o.BackingPciId, true
}

// HasBackingPciId returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) HasBackingPciId() bool {
	if o != nil && o.BackingPciId != nil {
		return true
	}

	return false
}

// SetBackingPciId gets a reference to the given string and assigns it to the BackingPciId field.
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) SetBackingPciId(v string) {
	o.BackingPciId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) SetName(v string) {
	o.Name = &v
}

// GetPassthrough returns the Passthrough field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) GetPassthrough() bool {
	if o == nil || o.Passthrough == nil {
		var ret bool
		return ret
	}
	return *o.Passthrough
}

// GetPassthroughOk returns a tuple with the Passthrough field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) GetPassthroughOk() (*bool, bool) {
	if o == nil || o.Passthrough == nil {
		return nil, false
	}
	return o.Passthrough, true
}

// HasPassthrough returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) HasPassthrough() bool {
	if o != nil && o.Passthrough != nil {
		return true
	}

	return false
}

// SetPassthrough gets a reference to the given bool and assigns it to the Passthrough field.
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) SetPassthrough(v bool) {
	o.Passthrough = &v
}

// GetBackingPciDevice returns the BackingPciDevice field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) GetBackingPciDevice() VirtualizationBaseHostPciDeviceRelationship {
	if o == nil || o.BackingPciDevice == nil {
		var ret VirtualizationBaseHostPciDeviceRelationship
		return ret
	}
	return *o.BackingPciDevice
}

// GetBackingPciDeviceOk returns a tuple with the BackingPciDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) GetBackingPciDeviceOk() (*VirtualizationBaseHostPciDeviceRelationship, bool) {
	if o == nil || o.BackingPciDevice == nil {
		return nil, false
	}
	return o.BackingPciDevice, true
}

// HasBackingPciDevice returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) HasBackingPciDevice() bool {
	if o != nil && o.BackingPciDevice != nil {
		return true
	}

	return false
}

// SetBackingPciDevice gets a reference to the given VirtualizationBaseHostPciDeviceRelationship and assigns it to the BackingPciDevice field.
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) SetBackingPciDevice(v VirtualizationBaseHostPciDeviceRelationship) {
	o.BackingPciDevice = &v
}

// GetVirtualMachine returns the VirtualMachine field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) GetVirtualMachine() VirtualizationBaseVirtualMachineRelationship {
	if o == nil || o.VirtualMachine == nil {
		var ret VirtualizationBaseVirtualMachineRelationship
		return ret
	}
	return *o.VirtualMachine
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) GetVirtualMachineOk() (*VirtualizationBaseVirtualMachineRelationship, bool) {
	if o == nil || o.VirtualMachine == nil {
		return nil, false
	}
	return o.VirtualMachine, true
}

// HasVirtualMachine returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) HasVirtualMachine() bool {
	if o != nil && o.VirtualMachine != nil {
		return true
	}

	return false
}

// SetVirtualMachine gets a reference to the given VirtualizationBaseVirtualMachineRelationship and assigns it to the VirtualMachine field.
func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) SetVirtualMachine(v VirtualizationBaseVirtualMachineRelationship) {
	o.VirtualMachine = &v
}

func (o VirtualizationBaseVirtualMachinePciDeviceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BackingPciId != nil {
		toSerialize["BackingPciId"] = o.BackingPciId
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Passthrough != nil {
		toSerialize["Passthrough"] = o.Passthrough
	}
	if o.BackingPciDevice != nil {
		toSerialize["BackingPciDevice"] = o.BackingPciDevice
	}
	if o.VirtualMachine != nil {
		toSerialize["VirtualMachine"] = o.VirtualMachine
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationBaseVirtualMachinePciDeviceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationBaseVirtualMachinePciDeviceAllOf := _VirtualizationBaseVirtualMachinePciDeviceAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationBaseVirtualMachinePciDeviceAllOf); err == nil {
		*o = VirtualizationBaseVirtualMachinePciDeviceAllOf(varVirtualizationBaseVirtualMachinePciDeviceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BackingPciId")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Passthrough")
		delete(additionalProperties, "BackingPciDevice")
		delete(additionalProperties, "VirtualMachine")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationBaseVirtualMachinePciDeviceAllOf struct {
	value *VirtualizationBaseVirtualMachinePciDeviceAllOf
	isSet bool
}

func (v NullableVirtualizationBaseVirtualMachinePciDeviceAllOf) Get() *VirtualizationBaseVirtualMachinePciDeviceAllOf {
	return v.value
}

func (v *NullableVirtualizationBaseVirtualMachinePciDeviceAllOf) Set(val *VirtualizationBaseVirtualMachinePciDeviceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationBaseVirtualMachinePciDeviceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationBaseVirtualMachinePciDeviceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationBaseVirtualMachinePciDeviceAllOf(val *VirtualizationBaseVirtualMachinePciDeviceAllOf) *NullableVirtualizationBaseVirtualMachinePciDeviceAllOf {
	return &NullableVirtualizationBaseVirtualMachinePciDeviceAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationBaseVirtualMachinePciDeviceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationBaseVirtualMachinePciDeviceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
