/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-03-27T10:08:12Z.
 *
 * API version: 1.0.9-4136
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// KubernetesNode Node inventory represents the Kubernetes Node. One node represent one Kubernetes worker or master. A Kubernetes node is a worker machine in Kubernetes. A node may be a virtual machine or physical machine. Each node contains the services necessary to run pods and is managed by the master components.
type KubernetesNode struct {
	KubernetesAbstractNode
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                               `json:"ObjectType"`
	NodeAddresses        []KubernetesNodeAddress              `json:"NodeAddresses,omitempty"`
	NodeInfo             NullableKubernetesNodeInfo           `json:"NodeInfo,omitempty"`
	NodeSpec             NullableKubernetesNodeSpec           `json:"NodeSpec,omitempty"`
	NodeStatuses         []KubernetesNodeStatus               `json:"NodeStatuses,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesNode KubernetesNode

// NewKubernetesNode instantiates a new KubernetesNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNode(classId string, objectType string) *KubernetesNode {
	this := KubernetesNode{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesNodeWithDefaults instantiates a new KubernetesNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNodeWithDefaults() *KubernetesNode {
	this := KubernetesNode{}
	var classId string = "kubernetes.Node"
	this.ClassId = classId
	var objectType string = "kubernetes.Node"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesNode) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesNode) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesNode) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesNode) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesNode) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesNode) SetObjectType(v string) {
	o.ObjectType = v
}

// GetNodeAddresses returns the NodeAddresses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNode) GetNodeAddresses() []KubernetesNodeAddress {
	if o == nil {
		var ret []KubernetesNodeAddress
		return ret
	}
	return o.NodeAddresses
}

// GetNodeAddressesOk returns a tuple with the NodeAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNode) GetNodeAddressesOk() (*[]KubernetesNodeAddress, bool) {
	if o == nil || o.NodeAddresses == nil {
		return nil, false
	}
	return &o.NodeAddresses, true
}

// HasNodeAddresses returns a boolean if a field has been set.
func (o *KubernetesNode) HasNodeAddresses() bool {
	if o != nil && o.NodeAddresses != nil {
		return true
	}

	return false
}

// SetNodeAddresses gets a reference to the given []KubernetesNodeAddress and assigns it to the NodeAddresses field.
func (o *KubernetesNode) SetNodeAddresses(v []KubernetesNodeAddress) {
	o.NodeAddresses = v
}

// GetNodeInfo returns the NodeInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNode) GetNodeInfo() KubernetesNodeInfo {
	if o == nil || o.NodeInfo.Get() == nil {
		var ret KubernetesNodeInfo
		return ret
	}
	return *o.NodeInfo.Get()
}

// GetNodeInfoOk returns a tuple with the NodeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNode) GetNodeInfoOk() (*KubernetesNodeInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeInfo.Get(), o.NodeInfo.IsSet()
}

// HasNodeInfo returns a boolean if a field has been set.
func (o *KubernetesNode) HasNodeInfo() bool {
	if o != nil && o.NodeInfo.IsSet() {
		return true
	}

	return false
}

// SetNodeInfo gets a reference to the given NullableKubernetesNodeInfo and assigns it to the NodeInfo field.
func (o *KubernetesNode) SetNodeInfo(v KubernetesNodeInfo) {
	o.NodeInfo.Set(&v)
}

// SetNodeInfoNil sets the value for NodeInfo to be an explicit nil
func (o *KubernetesNode) SetNodeInfoNil() {
	o.NodeInfo.Set(nil)
}

// UnsetNodeInfo ensures that no value is present for NodeInfo, not even an explicit nil
func (o *KubernetesNode) UnsetNodeInfo() {
	o.NodeInfo.Unset()
}

// GetNodeSpec returns the NodeSpec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNode) GetNodeSpec() KubernetesNodeSpec {
	if o == nil || o.NodeSpec.Get() == nil {
		var ret KubernetesNodeSpec
		return ret
	}
	return *o.NodeSpec.Get()
}

// GetNodeSpecOk returns a tuple with the NodeSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNode) GetNodeSpecOk() (*KubernetesNodeSpec, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeSpec.Get(), o.NodeSpec.IsSet()
}

// HasNodeSpec returns a boolean if a field has been set.
func (o *KubernetesNode) HasNodeSpec() bool {
	if o != nil && o.NodeSpec.IsSet() {
		return true
	}

	return false
}

// SetNodeSpec gets a reference to the given NullableKubernetesNodeSpec and assigns it to the NodeSpec field.
func (o *KubernetesNode) SetNodeSpec(v KubernetesNodeSpec) {
	o.NodeSpec.Set(&v)
}

// SetNodeSpecNil sets the value for NodeSpec to be an explicit nil
func (o *KubernetesNode) SetNodeSpecNil() {
	o.NodeSpec.Set(nil)
}

// UnsetNodeSpec ensures that no value is present for NodeSpec, not even an explicit nil
func (o *KubernetesNode) UnsetNodeSpec() {
	o.NodeSpec.Unset()
}

// GetNodeStatuses returns the NodeStatuses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNode) GetNodeStatuses() []KubernetesNodeStatus {
	if o == nil {
		var ret []KubernetesNodeStatus
		return ret
	}
	return o.NodeStatuses
}

// GetNodeStatusesOk returns a tuple with the NodeStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNode) GetNodeStatusesOk() (*[]KubernetesNodeStatus, bool) {
	if o == nil || o.NodeStatuses == nil {
		return nil, false
	}
	return &o.NodeStatuses, true
}

// HasNodeStatuses returns a boolean if a field has been set.
func (o *KubernetesNode) HasNodeStatuses() bool {
	if o != nil && o.NodeStatuses != nil {
		return true
	}

	return false
}

// SetNodeStatuses gets a reference to the given []KubernetesNodeStatus and assigns it to the NodeStatuses field.
func (o *KubernetesNode) SetNodeStatuses(v []KubernetesNodeStatus) {
	o.NodeStatuses = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *KubernetesNode) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNode) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *KubernetesNode) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *KubernetesNode) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o KubernetesNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKubernetesAbstractNode, errKubernetesAbstractNode := json.Marshal(o.KubernetesAbstractNode)
	if errKubernetesAbstractNode != nil {
		return []byte{}, errKubernetesAbstractNode
	}
	errKubernetesAbstractNode = json.Unmarshal([]byte(serializedKubernetesAbstractNode), &toSerialize)
	if errKubernetesAbstractNode != nil {
		return []byte{}, errKubernetesAbstractNode
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.NodeAddresses != nil {
		toSerialize["NodeAddresses"] = o.NodeAddresses
	}
	if o.NodeInfo.IsSet() {
		toSerialize["NodeInfo"] = o.NodeInfo.Get()
	}
	if o.NodeSpec.IsSet() {
		toSerialize["NodeSpec"] = o.NodeSpec.Get()
	}
	if o.NodeStatuses != nil {
		toSerialize["NodeStatuses"] = o.NodeStatuses
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesNode) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesNodeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType       string                               `json:"ObjectType"`
		NodeAddresses    []KubernetesNodeAddress              `json:"NodeAddresses,omitempty"`
		NodeInfo         NullableKubernetesNodeInfo           `json:"NodeInfo,omitempty"`
		NodeSpec         NullableKubernetesNodeSpec           `json:"NodeSpec,omitempty"`
		NodeStatuses     []KubernetesNodeStatus               `json:"NodeStatuses,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varKubernetesNodeWithoutEmbeddedStruct := KubernetesNodeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesNodeWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesNode := _KubernetesNode{}
		varKubernetesNode.ClassId = varKubernetesNodeWithoutEmbeddedStruct.ClassId
		varKubernetesNode.ObjectType = varKubernetesNodeWithoutEmbeddedStruct.ObjectType
		varKubernetesNode.NodeAddresses = varKubernetesNodeWithoutEmbeddedStruct.NodeAddresses
		varKubernetesNode.NodeInfo = varKubernetesNodeWithoutEmbeddedStruct.NodeInfo
		varKubernetesNode.NodeSpec = varKubernetesNodeWithoutEmbeddedStruct.NodeSpec
		varKubernetesNode.NodeStatuses = varKubernetesNodeWithoutEmbeddedStruct.NodeStatuses
		varKubernetesNode.RegisteredDevice = varKubernetesNodeWithoutEmbeddedStruct.RegisteredDevice
		*o = KubernetesNode(varKubernetesNode)
	} else {
		return err
	}

	varKubernetesNode := _KubernetesNode{}

	err = json.Unmarshal(bytes, &varKubernetesNode)
	if err == nil {
		o.KubernetesAbstractNode = varKubernetesNode.KubernetesAbstractNode
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "NodeAddresses")
		delete(additionalProperties, "NodeInfo")
		delete(additionalProperties, "NodeSpec")
		delete(additionalProperties, "NodeStatuses")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectKubernetesAbstractNode := reflect.ValueOf(o.KubernetesAbstractNode)
		for i := 0; i < reflectKubernetesAbstractNode.Type().NumField(); i++ {
			t := reflectKubernetesAbstractNode.Type().Field(i)

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

type NullableKubernetesNode struct {
	value *KubernetesNode
	isSet bool
}

func (v NullableKubernetesNode) Get() *KubernetesNode {
	return v.value
}

func (v *NullableKubernetesNode) Set(val *KubernetesNode) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNode) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNode(val *KubernetesNode) *NullableKubernetesNode {
	return &NullableKubernetesNode{value: val, isSet: true}
}

func (v NullableKubernetesNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}