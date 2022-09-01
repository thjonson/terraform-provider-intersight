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

// KubernetesVirtualMachineNodeProfile A profile specifying configuration settings for a Virtual Machine node. It is auto-generated based on the NodeGroupProfile and VirtualMachineNodePolicy configuration. Users can do operations like Drain, Cordon, Rebuild on a node.
type KubernetesVirtualMachineNodeProfile struct {
	KubernetesNodeProfile
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string               `json:"ObjectType"`
	Interfaces []KubernetesEthernet `json:"Interfaces,omitempty"`
	// An array of relationships to ippoolIpLease resources.
	// Deprecated
	IpAddresses          []IppoolIpLeaseRelationship               `json:"IpAddresses,omitempty"`
	NodeIp               *IppoolIpLeaseRelationship                `json:"NodeIp,omitempty"`
	VirtualMachine       *VirtualizationVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesVirtualMachineNodeProfile KubernetesVirtualMachineNodeProfile

// NewKubernetesVirtualMachineNodeProfile instantiates a new KubernetesVirtualMachineNodeProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesVirtualMachineNodeProfile(classId string, objectType string) *KubernetesVirtualMachineNodeProfile {
	this := KubernetesVirtualMachineNodeProfile{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "instance"
	this.Type = &type_
	var action string = "No-op"
	this.Action = &action
	var cloudProvider string = "noProvider"
	this.CloudProvider = &cloudProvider
	return &this
}

// NewKubernetesVirtualMachineNodeProfileWithDefaults instantiates a new KubernetesVirtualMachineNodeProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesVirtualMachineNodeProfileWithDefaults() *KubernetesVirtualMachineNodeProfile {
	this := KubernetesVirtualMachineNodeProfile{}
	var classId string = "kubernetes.VirtualMachineNodeProfile"
	this.ClassId = classId
	var objectType string = "kubernetes.VirtualMachineNodeProfile"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesVirtualMachineNodeProfile) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineNodeProfile) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesVirtualMachineNodeProfile) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesVirtualMachineNodeProfile) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineNodeProfile) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesVirtualMachineNodeProfile) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesVirtualMachineNodeProfile) GetInterfaces() []KubernetesEthernet {
	if o == nil {
		var ret []KubernetesEthernet
		return ret
	}
	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesVirtualMachineNodeProfile) GetInterfacesOk() ([]KubernetesEthernet, bool) {
	if o == nil || o.Interfaces == nil {
		return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineNodeProfile) HasInterfaces() bool {
	if o != nil && o.Interfaces != nil {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []KubernetesEthernet and assigns it to the Interfaces field.
func (o *KubernetesVirtualMachineNodeProfile) SetInterfaces(v []KubernetesEthernet) {
	o.Interfaces = v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *KubernetesVirtualMachineNodeProfile) GetIpAddresses() []IppoolIpLeaseRelationship {
	if o == nil {
		var ret []IppoolIpLeaseRelationship
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *KubernetesVirtualMachineNodeProfile) GetIpAddressesOk() ([]IppoolIpLeaseRelationship, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineNodeProfile) HasIpAddresses() bool {
	if o != nil && o.IpAddresses != nil {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []IppoolIpLeaseRelationship and assigns it to the IpAddresses field.
// Deprecated
func (o *KubernetesVirtualMachineNodeProfile) SetIpAddresses(v []IppoolIpLeaseRelationship) {
	o.IpAddresses = v
}

// GetNodeIp returns the NodeIp field value if set, zero value otherwise.
func (o *KubernetesVirtualMachineNodeProfile) GetNodeIp() IppoolIpLeaseRelationship {
	if o == nil || o.NodeIp == nil {
		var ret IppoolIpLeaseRelationship
		return ret
	}
	return *o.NodeIp
}

// GetNodeIpOk returns a tuple with the NodeIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineNodeProfile) GetNodeIpOk() (*IppoolIpLeaseRelationship, bool) {
	if o == nil || o.NodeIp == nil {
		return nil, false
	}
	return o.NodeIp, true
}

// HasNodeIp returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineNodeProfile) HasNodeIp() bool {
	if o != nil && o.NodeIp != nil {
		return true
	}

	return false
}

// SetNodeIp gets a reference to the given IppoolIpLeaseRelationship and assigns it to the NodeIp field.
func (o *KubernetesVirtualMachineNodeProfile) SetNodeIp(v IppoolIpLeaseRelationship) {
	o.NodeIp = &v
}

// GetVirtualMachine returns the VirtualMachine field value if set, zero value otherwise.
func (o *KubernetesVirtualMachineNodeProfile) GetVirtualMachine() VirtualizationVirtualMachineRelationship {
	if o == nil || o.VirtualMachine == nil {
		var ret VirtualizationVirtualMachineRelationship
		return ret
	}
	return *o.VirtualMachine
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineNodeProfile) GetVirtualMachineOk() (*VirtualizationVirtualMachineRelationship, bool) {
	if o == nil || o.VirtualMachine == nil {
		return nil, false
	}
	return o.VirtualMachine, true
}

// HasVirtualMachine returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineNodeProfile) HasVirtualMachine() bool {
	if o != nil && o.VirtualMachine != nil {
		return true
	}

	return false
}

// SetVirtualMachine gets a reference to the given VirtualizationVirtualMachineRelationship and assigns it to the VirtualMachine field.
func (o *KubernetesVirtualMachineNodeProfile) SetVirtualMachine(v VirtualizationVirtualMachineRelationship) {
	o.VirtualMachine = &v
}

func (o KubernetesVirtualMachineNodeProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKubernetesNodeProfile, errKubernetesNodeProfile := json.Marshal(o.KubernetesNodeProfile)
	if errKubernetesNodeProfile != nil {
		return []byte{}, errKubernetesNodeProfile
	}
	errKubernetesNodeProfile = json.Unmarshal([]byte(serializedKubernetesNodeProfile), &toSerialize)
	if errKubernetesNodeProfile != nil {
		return []byte{}, errKubernetesNodeProfile
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Interfaces != nil {
		toSerialize["Interfaces"] = o.Interfaces
	}
	if o.IpAddresses != nil {
		toSerialize["IpAddresses"] = o.IpAddresses
	}
	if o.NodeIp != nil {
		toSerialize["NodeIp"] = o.NodeIp
	}
	if o.VirtualMachine != nil {
		toSerialize["VirtualMachine"] = o.VirtualMachine
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesVirtualMachineNodeProfile) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesVirtualMachineNodeProfileWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string               `json:"ObjectType"`
		Interfaces []KubernetesEthernet `json:"Interfaces,omitempty"`
		// An array of relationships to ippoolIpLease resources.
		// Deprecated
		IpAddresses    []IppoolIpLeaseRelationship               `json:"IpAddresses,omitempty"`
		NodeIp         *IppoolIpLeaseRelationship                `json:"NodeIp,omitempty"`
		VirtualMachine *VirtualizationVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	}

	varKubernetesVirtualMachineNodeProfileWithoutEmbeddedStruct := KubernetesVirtualMachineNodeProfileWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesVirtualMachineNodeProfileWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesVirtualMachineNodeProfile := _KubernetesVirtualMachineNodeProfile{}
		varKubernetesVirtualMachineNodeProfile.ClassId = varKubernetesVirtualMachineNodeProfileWithoutEmbeddedStruct.ClassId
		varKubernetesVirtualMachineNodeProfile.ObjectType = varKubernetesVirtualMachineNodeProfileWithoutEmbeddedStruct.ObjectType
		varKubernetesVirtualMachineNodeProfile.Interfaces = varKubernetesVirtualMachineNodeProfileWithoutEmbeddedStruct.Interfaces
		varKubernetesVirtualMachineNodeProfile.IpAddresses = varKubernetesVirtualMachineNodeProfileWithoutEmbeddedStruct.IpAddresses
		varKubernetesVirtualMachineNodeProfile.NodeIp = varKubernetesVirtualMachineNodeProfileWithoutEmbeddedStruct.NodeIp
		varKubernetesVirtualMachineNodeProfile.VirtualMachine = varKubernetesVirtualMachineNodeProfileWithoutEmbeddedStruct.VirtualMachine
		*o = KubernetesVirtualMachineNodeProfile(varKubernetesVirtualMachineNodeProfile)
	} else {
		return err
	}

	varKubernetesVirtualMachineNodeProfile := _KubernetesVirtualMachineNodeProfile{}

	err = json.Unmarshal(bytes, &varKubernetesVirtualMachineNodeProfile)
	if err == nil {
		o.KubernetesNodeProfile = varKubernetesVirtualMachineNodeProfile.KubernetesNodeProfile
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Interfaces")
		delete(additionalProperties, "IpAddresses")
		delete(additionalProperties, "NodeIp")
		delete(additionalProperties, "VirtualMachine")

		// remove fields from embedded structs
		reflectKubernetesNodeProfile := reflect.ValueOf(o.KubernetesNodeProfile)
		for i := 0; i < reflectKubernetesNodeProfile.Type().NumField(); i++ {
			t := reflectKubernetesNodeProfile.Type().Field(i)

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

type NullableKubernetesVirtualMachineNodeProfile struct {
	value *KubernetesVirtualMachineNodeProfile
	isSet bool
}

func (v NullableKubernetesVirtualMachineNodeProfile) Get() *KubernetesVirtualMachineNodeProfile {
	return v.value
}

func (v *NullableKubernetesVirtualMachineNodeProfile) Set(val *KubernetesVirtualMachineNodeProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesVirtualMachineNodeProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesVirtualMachineNodeProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesVirtualMachineNodeProfile(val *KubernetesVirtualMachineNodeProfile) *NullableKubernetesVirtualMachineNodeProfile {
	return &NullableKubernetesVirtualMachineNodeProfile{value: val, isSet: true}
}

func (v NullableKubernetesVirtualMachineNodeProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesVirtualMachineNodeProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
