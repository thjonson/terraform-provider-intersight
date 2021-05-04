/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-04-28T13:03:38Z.
 *
 * API version: 1.0.9-4267
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// IqnpoolBlockAllOf Definition of the list of properties defined in 'iqnpool.Block', excluding properties defined in parent classes.
type IqnpoolBlockAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                   `json:"ObjectType"`
	IqnSuffixBlock       *IqnpoolIqnSuffixBlock   `json:"IqnSuffixBlock,omitempty"`
	Pool                 *IqnpoolPoolRelationship `json:"Pool,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IqnpoolBlockAllOf IqnpoolBlockAllOf

// NewIqnpoolBlockAllOf instantiates a new IqnpoolBlockAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIqnpoolBlockAllOf(classId string, objectType string) *IqnpoolBlockAllOf {
	this := IqnpoolBlockAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIqnpoolBlockAllOfWithDefaults instantiates a new IqnpoolBlockAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIqnpoolBlockAllOfWithDefaults() *IqnpoolBlockAllOf {
	this := IqnpoolBlockAllOf{}
	var classId string = "iqnpool.Block"
	this.ClassId = classId
	var objectType string = "iqnpool.Block"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IqnpoolBlockAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IqnpoolBlockAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IqnpoolBlockAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IqnpoolBlockAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IqnpoolBlockAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IqnpoolBlockAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIqnSuffixBlock returns the IqnSuffixBlock field value if set, zero value otherwise.
func (o *IqnpoolBlockAllOf) GetIqnSuffixBlock() IqnpoolIqnSuffixBlock {
	if o == nil || o.IqnSuffixBlock == nil {
		var ret IqnpoolIqnSuffixBlock
		return ret
	}
	return *o.IqnSuffixBlock
}

// GetIqnSuffixBlockOk returns a tuple with the IqnSuffixBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolBlockAllOf) GetIqnSuffixBlockOk() (*IqnpoolIqnSuffixBlock, bool) {
	if o == nil || o.IqnSuffixBlock == nil {
		return nil, false
	}
	return o.IqnSuffixBlock, true
}

// HasIqnSuffixBlock returns a boolean if a field has been set.
func (o *IqnpoolBlockAllOf) HasIqnSuffixBlock() bool {
	if o != nil && o.IqnSuffixBlock != nil {
		return true
	}

	return false
}

// SetIqnSuffixBlock gets a reference to the given IqnpoolIqnSuffixBlock and assigns it to the IqnSuffixBlock field.
func (o *IqnpoolBlockAllOf) SetIqnSuffixBlock(v IqnpoolIqnSuffixBlock) {
	o.IqnSuffixBlock = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *IqnpoolBlockAllOf) GetPool() IqnpoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret IqnpoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolBlockAllOf) GetPoolOk() (*IqnpoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *IqnpoolBlockAllOf) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given IqnpoolPoolRelationship and assigns it to the Pool field.
func (o *IqnpoolBlockAllOf) SetPool(v IqnpoolPoolRelationship) {
	o.Pool = &v
}

func (o IqnpoolBlockAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IqnSuffixBlock != nil {
		toSerialize["IqnSuffixBlock"] = o.IqnSuffixBlock
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IqnpoolBlockAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIqnpoolBlockAllOf := _IqnpoolBlockAllOf{}

	if err = json.Unmarshal(bytes, &varIqnpoolBlockAllOf); err == nil {
		*o = IqnpoolBlockAllOf(varIqnpoolBlockAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IqnSuffixBlock")
		delete(additionalProperties, "Pool")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIqnpoolBlockAllOf struct {
	value *IqnpoolBlockAllOf
	isSet bool
}

func (v NullableIqnpoolBlockAllOf) Get() *IqnpoolBlockAllOf {
	return v.value
}

func (v *NullableIqnpoolBlockAllOf) Set(val *IqnpoolBlockAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIqnpoolBlockAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIqnpoolBlockAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIqnpoolBlockAllOf(val *IqnpoolBlockAllOf) *NullableIqnpoolBlockAllOf {
	return &NullableIqnpoolBlockAllOf{value: val, isSet: true}
}

func (v NullableIqnpoolBlockAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIqnpoolBlockAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
