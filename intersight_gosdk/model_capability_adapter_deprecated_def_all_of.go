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

// CapabilityAdapterDeprecatedDefAllOf Definition of the list of properties defined in 'capability.AdapterDeprecatedDef', excluding properties defined in parent classes.
type CapabilityAdapterDeprecatedDefAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Model of the unsupported adapter.
	Model *string `json:"Model,omitempty"`
	// Vendor of the unsupported adapter.
	Vendor               *string `json:"Vendor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityAdapterDeprecatedDefAllOf CapabilityAdapterDeprecatedDefAllOf

// NewCapabilityAdapterDeprecatedDefAllOf instantiates a new CapabilityAdapterDeprecatedDefAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityAdapterDeprecatedDefAllOf(classId string, objectType string) *CapabilityAdapterDeprecatedDefAllOf {
	this := CapabilityAdapterDeprecatedDefAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityAdapterDeprecatedDefAllOfWithDefaults instantiates a new CapabilityAdapterDeprecatedDefAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityAdapterDeprecatedDefAllOfWithDefaults() *CapabilityAdapterDeprecatedDefAllOf {
	this := CapabilityAdapterDeprecatedDefAllOf{}
	var classId string = "capability.AdapterDeprecatedDef"
	this.ClassId = classId
	var objectType string = "capability.AdapterDeprecatedDef"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityAdapterDeprecatedDefAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterDeprecatedDefAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityAdapterDeprecatedDefAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityAdapterDeprecatedDefAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterDeprecatedDefAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityAdapterDeprecatedDefAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *CapabilityAdapterDeprecatedDefAllOf) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterDeprecatedDefAllOf) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *CapabilityAdapterDeprecatedDefAllOf) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *CapabilityAdapterDeprecatedDefAllOf) SetModel(v string) {
	o.Model = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *CapabilityAdapterDeprecatedDefAllOf) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterDeprecatedDefAllOf) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *CapabilityAdapterDeprecatedDefAllOf) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *CapabilityAdapterDeprecatedDefAllOf) SetVendor(v string) {
	o.Vendor = &v
}

func (o CapabilityAdapterDeprecatedDefAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilityAdapterDeprecatedDefAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilityAdapterDeprecatedDefAllOf := _CapabilityAdapterDeprecatedDefAllOf{}

	if err = json.Unmarshal(bytes, &varCapabilityAdapterDeprecatedDefAllOf); err == nil {
		*o = CapabilityAdapterDeprecatedDefAllOf(varCapabilityAdapterDeprecatedDefAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "Vendor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilityAdapterDeprecatedDefAllOf struct {
	value *CapabilityAdapterDeprecatedDefAllOf
	isSet bool
}

func (v NullableCapabilityAdapterDeprecatedDefAllOf) Get() *CapabilityAdapterDeprecatedDefAllOf {
	return v.value
}

func (v *NullableCapabilityAdapterDeprecatedDefAllOf) Set(val *CapabilityAdapterDeprecatedDefAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityAdapterDeprecatedDefAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityAdapterDeprecatedDefAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityAdapterDeprecatedDefAllOf(val *CapabilityAdapterDeprecatedDefAllOf) *NullableCapabilityAdapterDeprecatedDefAllOf {
	return &NullableCapabilityAdapterDeprecatedDefAllOf{value: val, isSet: true}
}

func (v NullableCapabilityAdapterDeprecatedDefAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityAdapterDeprecatedDefAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
