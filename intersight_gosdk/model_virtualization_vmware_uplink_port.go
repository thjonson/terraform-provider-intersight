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

// VirtualizationVmwareUplinkPort The VMware uplink port object is represented here.
type VirtualizationVmwareUplinkPort struct {
	VirtualizationBaseSourceDevice
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The VMware managed object reference as a string.
	Identity *string `json:"Identity,omitempty"`
	// The internally assigned key of this uplink port object. This entity is not manipulated by users.
	Key *string `json:"Key,omitempty"`
	// User-provided name to identify the uplink port object.
	Name                     *string                                                   `json:"Name,omitempty"`
	DistributedNetwork       *VirtualizationVmwareDistributedNetworkRelationship       `json:"DistributedNetwork,omitempty"`
	Host                     *VirtualizationVmwareHostRelationship                     `json:"Host,omitempty"`
	PhysicalNetworkInterface *VirtualizationVmwarePhysicalNetworkInterfaceRelationship `json:"PhysicalNetworkInterface,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _VirtualizationVmwareUplinkPort VirtualizationVmwareUplinkPort

// NewVirtualizationVmwareUplinkPort instantiates a new VirtualizationVmwareUplinkPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareUplinkPort(classId string, objectType string) *VirtualizationVmwareUplinkPort {
	this := VirtualizationVmwareUplinkPort{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwareUplinkPortWithDefaults instantiates a new VirtualizationVmwareUplinkPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareUplinkPortWithDefaults() *VirtualizationVmwareUplinkPort {
	this := VirtualizationVmwareUplinkPort{}
	var classId string = "virtualization.VmwareUplinkPort"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareUplinkPort"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareUplinkPort) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareUplinkPort) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareUplinkPort) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareUplinkPort) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareUplinkPort) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareUplinkPort) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *VirtualizationVmwareUplinkPort) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareUplinkPort) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *VirtualizationVmwareUplinkPort) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *VirtualizationVmwareUplinkPort) SetIdentity(v string) {
	o.Identity = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *VirtualizationVmwareUplinkPort) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareUplinkPort) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *VirtualizationVmwareUplinkPort) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *VirtualizationVmwareUplinkPort) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationVmwareUplinkPort) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareUplinkPort) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationVmwareUplinkPort) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationVmwareUplinkPort) SetName(v string) {
	o.Name = &v
}

// GetDistributedNetwork returns the DistributedNetwork field value if set, zero value otherwise.
func (o *VirtualizationVmwareUplinkPort) GetDistributedNetwork() VirtualizationVmwareDistributedNetworkRelationship {
	if o == nil || o.DistributedNetwork == nil {
		var ret VirtualizationVmwareDistributedNetworkRelationship
		return ret
	}
	return *o.DistributedNetwork
}

// GetDistributedNetworkOk returns a tuple with the DistributedNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareUplinkPort) GetDistributedNetworkOk() (*VirtualizationVmwareDistributedNetworkRelationship, bool) {
	if o == nil || o.DistributedNetwork == nil {
		return nil, false
	}
	return o.DistributedNetwork, true
}

// HasDistributedNetwork returns a boolean if a field has been set.
func (o *VirtualizationVmwareUplinkPort) HasDistributedNetwork() bool {
	if o != nil && o.DistributedNetwork != nil {
		return true
	}

	return false
}

// SetDistributedNetwork gets a reference to the given VirtualizationVmwareDistributedNetworkRelationship and assigns it to the DistributedNetwork field.
func (o *VirtualizationVmwareUplinkPort) SetDistributedNetwork(v VirtualizationVmwareDistributedNetworkRelationship) {
	o.DistributedNetwork = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *VirtualizationVmwareUplinkPort) GetHost() VirtualizationVmwareHostRelationship {
	if o == nil || o.Host == nil {
		var ret VirtualizationVmwareHostRelationship
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareUplinkPort) GetHostOk() (*VirtualizationVmwareHostRelationship, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *VirtualizationVmwareUplinkPort) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given VirtualizationVmwareHostRelationship and assigns it to the Host field.
func (o *VirtualizationVmwareUplinkPort) SetHost(v VirtualizationVmwareHostRelationship) {
	o.Host = &v
}

// GetPhysicalNetworkInterface returns the PhysicalNetworkInterface field value if set, zero value otherwise.
func (o *VirtualizationVmwareUplinkPort) GetPhysicalNetworkInterface() VirtualizationVmwarePhysicalNetworkInterfaceRelationship {
	if o == nil || o.PhysicalNetworkInterface == nil {
		var ret VirtualizationVmwarePhysicalNetworkInterfaceRelationship
		return ret
	}
	return *o.PhysicalNetworkInterface
}

// GetPhysicalNetworkInterfaceOk returns a tuple with the PhysicalNetworkInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareUplinkPort) GetPhysicalNetworkInterfaceOk() (*VirtualizationVmwarePhysicalNetworkInterfaceRelationship, bool) {
	if o == nil || o.PhysicalNetworkInterface == nil {
		return nil, false
	}
	return o.PhysicalNetworkInterface, true
}

// HasPhysicalNetworkInterface returns a boolean if a field has been set.
func (o *VirtualizationVmwareUplinkPort) HasPhysicalNetworkInterface() bool {
	if o != nil && o.PhysicalNetworkInterface != nil {
		return true
	}

	return false
}

// SetPhysicalNetworkInterface gets a reference to the given VirtualizationVmwarePhysicalNetworkInterfaceRelationship and assigns it to the PhysicalNetworkInterface field.
func (o *VirtualizationVmwareUplinkPort) SetPhysicalNetworkInterface(v VirtualizationVmwarePhysicalNetworkInterfaceRelationship) {
	o.PhysicalNetworkInterface = &v
}

func (o VirtualizationVmwareUplinkPort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseSourceDevice, errVirtualizationBaseSourceDevice := json.Marshal(o.VirtualizationBaseSourceDevice)
	if errVirtualizationBaseSourceDevice != nil {
		return []byte{}, errVirtualizationBaseSourceDevice
	}
	errVirtualizationBaseSourceDevice = json.Unmarshal([]byte(serializedVirtualizationBaseSourceDevice), &toSerialize)
	if errVirtualizationBaseSourceDevice != nil {
		return []byte{}, errVirtualizationBaseSourceDevice
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.DistributedNetwork != nil {
		toSerialize["DistributedNetwork"] = o.DistributedNetwork
	}
	if o.Host != nil {
		toSerialize["Host"] = o.Host
	}
	if o.PhysicalNetworkInterface != nil {
		toSerialize["PhysicalNetworkInterface"] = o.PhysicalNetworkInterface
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareUplinkPort) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationVmwareUplinkPortWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The VMware managed object reference as a string.
		Identity *string `json:"Identity,omitempty"`
		// The internally assigned key of this uplink port object. This entity is not manipulated by users.
		Key *string `json:"Key,omitempty"`
		// User-provided name to identify the uplink port object.
		Name                     *string                                                   `json:"Name,omitempty"`
		DistributedNetwork       *VirtualizationVmwareDistributedNetworkRelationship       `json:"DistributedNetwork,omitempty"`
		Host                     *VirtualizationVmwareHostRelationship                     `json:"Host,omitempty"`
		PhysicalNetworkInterface *VirtualizationVmwarePhysicalNetworkInterfaceRelationship `json:"PhysicalNetworkInterface,omitempty"`
	}

	varVirtualizationVmwareUplinkPortWithoutEmbeddedStruct := VirtualizationVmwareUplinkPortWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareUplinkPortWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareUplinkPort := _VirtualizationVmwareUplinkPort{}
		varVirtualizationVmwareUplinkPort.ClassId = varVirtualizationVmwareUplinkPortWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareUplinkPort.ObjectType = varVirtualizationVmwareUplinkPortWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareUplinkPort.Identity = varVirtualizationVmwareUplinkPortWithoutEmbeddedStruct.Identity
		varVirtualizationVmwareUplinkPort.Key = varVirtualizationVmwareUplinkPortWithoutEmbeddedStruct.Key
		varVirtualizationVmwareUplinkPort.Name = varVirtualizationVmwareUplinkPortWithoutEmbeddedStruct.Name
		varVirtualizationVmwareUplinkPort.DistributedNetwork = varVirtualizationVmwareUplinkPortWithoutEmbeddedStruct.DistributedNetwork
		varVirtualizationVmwareUplinkPort.Host = varVirtualizationVmwareUplinkPortWithoutEmbeddedStruct.Host
		varVirtualizationVmwareUplinkPort.PhysicalNetworkInterface = varVirtualizationVmwareUplinkPortWithoutEmbeddedStruct.PhysicalNetworkInterface
		*o = VirtualizationVmwareUplinkPort(varVirtualizationVmwareUplinkPort)
	} else {
		return err
	}

	varVirtualizationVmwareUplinkPort := _VirtualizationVmwareUplinkPort{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareUplinkPort)
	if err == nil {
		o.VirtualizationBaseSourceDevice = varVirtualizationVmwareUplinkPort.VirtualizationBaseSourceDevice
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "DistributedNetwork")
		delete(additionalProperties, "Host")
		delete(additionalProperties, "PhysicalNetworkInterface")

		// remove fields from embedded structs
		reflectVirtualizationBaseSourceDevice := reflect.ValueOf(o.VirtualizationBaseSourceDevice)
		for i := 0; i < reflectVirtualizationBaseSourceDevice.Type().NumField(); i++ {
			t := reflectVirtualizationBaseSourceDevice.Type().Field(i)

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

type NullableVirtualizationVmwareUplinkPort struct {
	value *VirtualizationVmwareUplinkPort
	isSet bool
}

func (v NullableVirtualizationVmwareUplinkPort) Get() *VirtualizationVmwareUplinkPort {
	return v.value
}

func (v *NullableVirtualizationVmwareUplinkPort) Set(val *VirtualizationVmwareUplinkPort) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareUplinkPort) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareUplinkPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareUplinkPort(val *VirtualizationVmwareUplinkPort) *NullableVirtualizationVmwareUplinkPort {
	return &NullableVirtualizationVmwareUplinkPort{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareUplinkPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareUplinkPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
