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
)

// HyperflexIpAddrRangeAllOf Definition of the list of properties defined in 'hyperflex.IpAddrRange', excluding properties defined in parent classes.
type HyperflexIpAddrRangeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The end IPv4 address of the range.
	EndAddr *string `json:"EndAddr,omitempty"`
	// The default gateway for the start and end IPv4 addresses.
	Gateway *string `json:"Gateway,omitempty"`
	// The netmask specified in dot decimal notation. The start address, end address, and gateway must all be within the network specified by this netmask.
	Netmask *string `json:"Netmask,omitempty"`
	// The start IPv4 address of the range.
	StartAddr            *string `json:"StartAddr,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexIpAddrRangeAllOf HyperflexIpAddrRangeAllOf

// NewHyperflexIpAddrRangeAllOf instantiates a new HyperflexIpAddrRangeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexIpAddrRangeAllOf(classId string, objectType string) *HyperflexIpAddrRangeAllOf {
	this := HyperflexIpAddrRangeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexIpAddrRangeAllOfWithDefaults instantiates a new HyperflexIpAddrRangeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexIpAddrRangeAllOfWithDefaults() *HyperflexIpAddrRangeAllOf {
	this := HyperflexIpAddrRangeAllOf{}
	var classId string = "hyperflex.IpAddrRange"
	this.ClassId = classId
	var objectType string = "hyperflex.IpAddrRange"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexIpAddrRangeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexIpAddrRangeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexIpAddrRangeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexIpAddrRangeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexIpAddrRangeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexIpAddrRangeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEndAddr returns the EndAddr field value if set, zero value otherwise.
func (o *HyperflexIpAddrRangeAllOf) GetEndAddr() string {
	if o == nil || o.EndAddr == nil {
		var ret string
		return ret
	}
	return *o.EndAddr
}

// GetEndAddrOk returns a tuple with the EndAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexIpAddrRangeAllOf) GetEndAddrOk() (*string, bool) {
	if o == nil || o.EndAddr == nil {
		return nil, false
	}
	return o.EndAddr, true
}

// HasEndAddr returns a boolean if a field has been set.
func (o *HyperflexIpAddrRangeAllOf) HasEndAddr() bool {
	if o != nil && o.EndAddr != nil {
		return true
	}

	return false
}

// SetEndAddr gets a reference to the given string and assigns it to the EndAddr field.
func (o *HyperflexIpAddrRangeAllOf) SetEndAddr(v string) {
	o.EndAddr = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *HyperflexIpAddrRangeAllOf) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexIpAddrRangeAllOf) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *HyperflexIpAddrRangeAllOf) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *HyperflexIpAddrRangeAllOf) SetGateway(v string) {
	o.Gateway = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *HyperflexIpAddrRangeAllOf) GetNetmask() string {
	if o == nil || o.Netmask == nil {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexIpAddrRangeAllOf) GetNetmaskOk() (*string, bool) {
	if o == nil || o.Netmask == nil {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *HyperflexIpAddrRangeAllOf) HasNetmask() bool {
	if o != nil && o.Netmask != nil {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *HyperflexIpAddrRangeAllOf) SetNetmask(v string) {
	o.Netmask = &v
}

// GetStartAddr returns the StartAddr field value if set, zero value otherwise.
func (o *HyperflexIpAddrRangeAllOf) GetStartAddr() string {
	if o == nil || o.StartAddr == nil {
		var ret string
		return ret
	}
	return *o.StartAddr
}

// GetStartAddrOk returns a tuple with the StartAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexIpAddrRangeAllOf) GetStartAddrOk() (*string, bool) {
	if o == nil || o.StartAddr == nil {
		return nil, false
	}
	return o.StartAddr, true
}

// HasStartAddr returns a boolean if a field has been set.
func (o *HyperflexIpAddrRangeAllOf) HasStartAddr() bool {
	if o != nil && o.StartAddr != nil {
		return true
	}

	return false
}

// SetStartAddr gets a reference to the given string and assigns it to the StartAddr field.
func (o *HyperflexIpAddrRangeAllOf) SetStartAddr(v string) {
	o.StartAddr = &v
}

func (o HyperflexIpAddrRangeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EndAddr != nil {
		toSerialize["EndAddr"] = o.EndAddr
	}
	if o.Gateway != nil {
		toSerialize["Gateway"] = o.Gateway
	}
	if o.Netmask != nil {
		toSerialize["Netmask"] = o.Netmask
	}
	if o.StartAddr != nil {
		toSerialize["StartAddr"] = o.StartAddr
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexIpAddrRangeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexIpAddrRangeAllOf := _HyperflexIpAddrRangeAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexIpAddrRangeAllOf); err == nil {
		*o = HyperflexIpAddrRangeAllOf(varHyperflexIpAddrRangeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EndAddr")
		delete(additionalProperties, "Gateway")
		delete(additionalProperties, "Netmask")
		delete(additionalProperties, "StartAddr")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexIpAddrRangeAllOf struct {
	value *HyperflexIpAddrRangeAllOf
	isSet bool
}

func (v NullableHyperflexIpAddrRangeAllOf) Get() *HyperflexIpAddrRangeAllOf {
	return v.value
}

func (v *NullableHyperflexIpAddrRangeAllOf) Set(val *HyperflexIpAddrRangeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexIpAddrRangeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexIpAddrRangeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexIpAddrRangeAllOf(val *HyperflexIpAddrRangeAllOf) *NullableHyperflexIpAddrRangeAllOf {
	return &NullableHyperflexIpAddrRangeAllOf{value: val, isSet: true}
}

func (v NullableHyperflexIpAddrRangeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexIpAddrRangeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
