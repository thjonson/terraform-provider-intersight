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

// FabricVlanSetAllOf Definition of the list of properties defined in 'fabric.VlanSet', excluding properties defined in parent classes.
type FabricVlanSetAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Used to determine whether this VLAN will be allowed on all uplink ports and PCs in this FI.
	AutoAllowOnUplinks *bool `json:"AutoAllowOnUplinks,omitempty"`
	// Used to define whether this VLAN is to be classified as 'native' for traffic in this FI.
	IsNative *bool `json:"IsNative,omitempty"`
	// The 'name' used to identify this VLAN.
	Name *string `json:"Name,omitempty"`
	// Set of VLANs defined by VLAN object with identical configuration.
	Vlans                *string                             `json:"Vlans,omitempty"`
	EthNetworkPolicy     *FabricEthNetworkPolicyRelationship `json:"EthNetworkPolicy,omitempty"`
	MulticastPolicy      *FabricMulticastPolicyRelationship  `json:"MulticastPolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricVlanSetAllOf FabricVlanSetAllOf

// NewFabricVlanSetAllOf instantiates a new FabricVlanSetAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricVlanSetAllOf(classId string, objectType string) *FabricVlanSetAllOf {
	this := FabricVlanSetAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricVlanSetAllOfWithDefaults instantiates a new FabricVlanSetAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricVlanSetAllOfWithDefaults() *FabricVlanSetAllOf {
	this := FabricVlanSetAllOf{}
	var classId string = "fabric.VlanSet"
	this.ClassId = classId
	var objectType string = "fabric.VlanSet"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricVlanSetAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricVlanSetAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricVlanSetAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricVlanSetAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricVlanSetAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricVlanSetAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAutoAllowOnUplinks returns the AutoAllowOnUplinks field value if set, zero value otherwise.
func (o *FabricVlanSetAllOf) GetAutoAllowOnUplinks() bool {
	if o == nil || o.AutoAllowOnUplinks == nil {
		var ret bool
		return ret
	}
	return *o.AutoAllowOnUplinks
}

// GetAutoAllowOnUplinksOk returns a tuple with the AutoAllowOnUplinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlanSetAllOf) GetAutoAllowOnUplinksOk() (*bool, bool) {
	if o == nil || o.AutoAllowOnUplinks == nil {
		return nil, false
	}
	return o.AutoAllowOnUplinks, true
}

// HasAutoAllowOnUplinks returns a boolean if a field has been set.
func (o *FabricVlanSetAllOf) HasAutoAllowOnUplinks() bool {
	if o != nil && o.AutoAllowOnUplinks != nil {
		return true
	}

	return false
}

// SetAutoAllowOnUplinks gets a reference to the given bool and assigns it to the AutoAllowOnUplinks field.
func (o *FabricVlanSetAllOf) SetAutoAllowOnUplinks(v bool) {
	o.AutoAllowOnUplinks = &v
}

// GetIsNative returns the IsNative field value if set, zero value otherwise.
func (o *FabricVlanSetAllOf) GetIsNative() bool {
	if o == nil || o.IsNative == nil {
		var ret bool
		return ret
	}
	return *o.IsNative
}

// GetIsNativeOk returns a tuple with the IsNative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlanSetAllOf) GetIsNativeOk() (*bool, bool) {
	if o == nil || o.IsNative == nil {
		return nil, false
	}
	return o.IsNative, true
}

// HasIsNative returns a boolean if a field has been set.
func (o *FabricVlanSetAllOf) HasIsNative() bool {
	if o != nil && o.IsNative != nil {
		return true
	}

	return false
}

// SetIsNative gets a reference to the given bool and assigns it to the IsNative field.
func (o *FabricVlanSetAllOf) SetIsNative(v bool) {
	o.IsNative = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FabricVlanSetAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlanSetAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FabricVlanSetAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FabricVlanSetAllOf) SetName(v string) {
	o.Name = &v
}

// GetVlans returns the Vlans field value if set, zero value otherwise.
func (o *FabricVlanSetAllOf) GetVlans() string {
	if o == nil || o.Vlans == nil {
		var ret string
		return ret
	}
	return *o.Vlans
}

// GetVlansOk returns a tuple with the Vlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlanSetAllOf) GetVlansOk() (*string, bool) {
	if o == nil || o.Vlans == nil {
		return nil, false
	}
	return o.Vlans, true
}

// HasVlans returns a boolean if a field has been set.
func (o *FabricVlanSetAllOf) HasVlans() bool {
	if o != nil && o.Vlans != nil {
		return true
	}

	return false
}

// SetVlans gets a reference to the given string and assigns it to the Vlans field.
func (o *FabricVlanSetAllOf) SetVlans(v string) {
	o.Vlans = &v
}

// GetEthNetworkPolicy returns the EthNetworkPolicy field value if set, zero value otherwise.
func (o *FabricVlanSetAllOf) GetEthNetworkPolicy() FabricEthNetworkPolicyRelationship {
	if o == nil || o.EthNetworkPolicy == nil {
		var ret FabricEthNetworkPolicyRelationship
		return ret
	}
	return *o.EthNetworkPolicy
}

// GetEthNetworkPolicyOk returns a tuple with the EthNetworkPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlanSetAllOf) GetEthNetworkPolicyOk() (*FabricEthNetworkPolicyRelationship, bool) {
	if o == nil || o.EthNetworkPolicy == nil {
		return nil, false
	}
	return o.EthNetworkPolicy, true
}

// HasEthNetworkPolicy returns a boolean if a field has been set.
func (o *FabricVlanSetAllOf) HasEthNetworkPolicy() bool {
	if o != nil && o.EthNetworkPolicy != nil {
		return true
	}

	return false
}

// SetEthNetworkPolicy gets a reference to the given FabricEthNetworkPolicyRelationship and assigns it to the EthNetworkPolicy field.
func (o *FabricVlanSetAllOf) SetEthNetworkPolicy(v FabricEthNetworkPolicyRelationship) {
	o.EthNetworkPolicy = &v
}

// GetMulticastPolicy returns the MulticastPolicy field value if set, zero value otherwise.
func (o *FabricVlanSetAllOf) GetMulticastPolicy() FabricMulticastPolicyRelationship {
	if o == nil || o.MulticastPolicy == nil {
		var ret FabricMulticastPolicyRelationship
		return ret
	}
	return *o.MulticastPolicy
}

// GetMulticastPolicyOk returns a tuple with the MulticastPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlanSetAllOf) GetMulticastPolicyOk() (*FabricMulticastPolicyRelationship, bool) {
	if o == nil || o.MulticastPolicy == nil {
		return nil, false
	}
	return o.MulticastPolicy, true
}

// HasMulticastPolicy returns a boolean if a field has been set.
func (o *FabricVlanSetAllOf) HasMulticastPolicy() bool {
	if o != nil && o.MulticastPolicy != nil {
		return true
	}

	return false
}

// SetMulticastPolicy gets a reference to the given FabricMulticastPolicyRelationship and assigns it to the MulticastPolicy field.
func (o *FabricVlanSetAllOf) SetMulticastPolicy(v FabricMulticastPolicyRelationship) {
	o.MulticastPolicy = &v
}

func (o FabricVlanSetAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AutoAllowOnUplinks != nil {
		toSerialize["AutoAllowOnUplinks"] = o.AutoAllowOnUplinks
	}
	if o.IsNative != nil {
		toSerialize["IsNative"] = o.IsNative
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Vlans != nil {
		toSerialize["Vlans"] = o.Vlans
	}
	if o.EthNetworkPolicy != nil {
		toSerialize["EthNetworkPolicy"] = o.EthNetworkPolicy
	}
	if o.MulticastPolicy != nil {
		toSerialize["MulticastPolicy"] = o.MulticastPolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricVlanSetAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricVlanSetAllOf := _FabricVlanSetAllOf{}

	if err = json.Unmarshal(bytes, &varFabricVlanSetAllOf); err == nil {
		*o = FabricVlanSetAllOf(varFabricVlanSetAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AutoAllowOnUplinks")
		delete(additionalProperties, "IsNative")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Vlans")
		delete(additionalProperties, "EthNetworkPolicy")
		delete(additionalProperties, "MulticastPolicy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricVlanSetAllOf struct {
	value *FabricVlanSetAllOf
	isSet bool
}

func (v NullableFabricVlanSetAllOf) Get() *FabricVlanSetAllOf {
	return v.value
}

func (v *NullableFabricVlanSetAllOf) Set(val *FabricVlanSetAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricVlanSetAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricVlanSetAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricVlanSetAllOf(val *FabricVlanSetAllOf) *NullableFabricVlanSetAllOf {
	return &NullableFabricVlanSetAllOf{value: val, isSet: true}
}

func (v NullableFabricVlanSetAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricVlanSetAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
