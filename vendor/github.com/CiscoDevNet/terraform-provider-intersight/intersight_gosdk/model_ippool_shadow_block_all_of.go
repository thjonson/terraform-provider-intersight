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

// IppoolShadowBlockAllOf Definition of the list of properties defined in 'ippool.ShadowBlock', excluding properties defined in parent classes.
type IppoolShadowBlockAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of this IP addresses blocks. * `IPv4` - IP V4 address type requested. * `IPv6` - IP V6 address type requested.
	IpType               *string                       `json:"IpType,omitempty"`
	IpV4Block            *IppoolIpV4Block              `json:"IpV4Block,omitempty"`
	IpV6Block            *IppoolIpV6Block              `json:"IpV6Block,omitempty"`
	Pool                 *IppoolShadowPoolRelationship `json:"Pool,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IppoolShadowBlockAllOf IppoolShadowBlockAllOf

// NewIppoolShadowBlockAllOf instantiates a new IppoolShadowBlockAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolShadowBlockAllOf(classId string, objectType string) *IppoolShadowBlockAllOf {
	this := IppoolShadowBlockAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var ipType string = "IPv4"
	this.IpType = &ipType
	return &this
}

// NewIppoolShadowBlockAllOfWithDefaults instantiates a new IppoolShadowBlockAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolShadowBlockAllOfWithDefaults() *IppoolShadowBlockAllOf {
	this := IppoolShadowBlockAllOf{}
	var classId string = "ippool.ShadowBlock"
	this.ClassId = classId
	var objectType string = "ippool.ShadowBlock"
	this.ObjectType = objectType
	var ipType string = "IPv4"
	this.IpType = &ipType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IppoolShadowBlockAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IppoolShadowBlockAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IppoolShadowBlockAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IppoolShadowBlockAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IppoolShadowBlockAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IppoolShadowBlockAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIpType returns the IpType field value if set, zero value otherwise.
func (o *IppoolShadowBlockAllOf) GetIpType() string {
	if o == nil || o.IpType == nil {
		var ret string
		return ret
	}
	return *o.IpType
}

// GetIpTypeOk returns a tuple with the IpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowBlockAllOf) GetIpTypeOk() (*string, bool) {
	if o == nil || o.IpType == nil {
		return nil, false
	}
	return o.IpType, true
}

// HasIpType returns a boolean if a field has been set.
func (o *IppoolShadowBlockAllOf) HasIpType() bool {
	if o != nil && o.IpType != nil {
		return true
	}

	return false
}

// SetIpType gets a reference to the given string and assigns it to the IpType field.
func (o *IppoolShadowBlockAllOf) SetIpType(v string) {
	o.IpType = &v
}

// GetIpV4Block returns the IpV4Block field value if set, zero value otherwise.
func (o *IppoolShadowBlockAllOf) GetIpV4Block() IppoolIpV4Block {
	if o == nil || o.IpV4Block == nil {
		var ret IppoolIpV4Block
		return ret
	}
	return *o.IpV4Block
}

// GetIpV4BlockOk returns a tuple with the IpV4Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowBlockAllOf) GetIpV4BlockOk() (*IppoolIpV4Block, bool) {
	if o == nil || o.IpV4Block == nil {
		return nil, false
	}
	return o.IpV4Block, true
}

// HasIpV4Block returns a boolean if a field has been set.
func (o *IppoolShadowBlockAllOf) HasIpV4Block() bool {
	if o != nil && o.IpV4Block != nil {
		return true
	}

	return false
}

// SetIpV4Block gets a reference to the given IppoolIpV4Block and assigns it to the IpV4Block field.
func (o *IppoolShadowBlockAllOf) SetIpV4Block(v IppoolIpV4Block) {
	o.IpV4Block = &v
}

// GetIpV6Block returns the IpV6Block field value if set, zero value otherwise.
func (o *IppoolShadowBlockAllOf) GetIpV6Block() IppoolIpV6Block {
	if o == nil || o.IpV6Block == nil {
		var ret IppoolIpV6Block
		return ret
	}
	return *o.IpV6Block
}

// GetIpV6BlockOk returns a tuple with the IpV6Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowBlockAllOf) GetIpV6BlockOk() (*IppoolIpV6Block, bool) {
	if o == nil || o.IpV6Block == nil {
		return nil, false
	}
	return o.IpV6Block, true
}

// HasIpV6Block returns a boolean if a field has been set.
func (o *IppoolShadowBlockAllOf) HasIpV6Block() bool {
	if o != nil && o.IpV6Block != nil {
		return true
	}

	return false
}

// SetIpV6Block gets a reference to the given IppoolIpV6Block and assigns it to the IpV6Block field.
func (o *IppoolShadowBlockAllOf) SetIpV6Block(v IppoolIpV6Block) {
	o.IpV6Block = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *IppoolShadowBlockAllOf) GetPool() IppoolShadowPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret IppoolShadowPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowBlockAllOf) GetPoolOk() (*IppoolShadowPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *IppoolShadowBlockAllOf) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given IppoolShadowPoolRelationship and assigns it to the Pool field.
func (o *IppoolShadowBlockAllOf) SetPool(v IppoolShadowPoolRelationship) {
	o.Pool = &v
}

func (o IppoolShadowBlockAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IpType != nil {
		toSerialize["IpType"] = o.IpType
	}
	if o.IpV4Block != nil {
		toSerialize["IpV4Block"] = o.IpV4Block
	}
	if o.IpV6Block != nil {
		toSerialize["IpV6Block"] = o.IpV6Block
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IppoolShadowBlockAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIppoolShadowBlockAllOf := _IppoolShadowBlockAllOf{}

	if err = json.Unmarshal(bytes, &varIppoolShadowBlockAllOf); err == nil {
		*o = IppoolShadowBlockAllOf(varIppoolShadowBlockAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IpType")
		delete(additionalProperties, "IpV4Block")
		delete(additionalProperties, "IpV6Block")
		delete(additionalProperties, "Pool")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIppoolShadowBlockAllOf struct {
	value *IppoolShadowBlockAllOf
	isSet bool
}

func (v NullableIppoolShadowBlockAllOf) Get() *IppoolShadowBlockAllOf {
	return v.value
}

func (v *NullableIppoolShadowBlockAllOf) Set(val *IppoolShadowBlockAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolShadowBlockAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolShadowBlockAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolShadowBlockAllOf(val *IppoolShadowBlockAllOf) *NullableIppoolShadowBlockAllOf {
	return &NullableIppoolShadowBlockAllOf{value: val, isSet: true}
}

func (v NullableIppoolShadowBlockAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolShadowBlockAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
