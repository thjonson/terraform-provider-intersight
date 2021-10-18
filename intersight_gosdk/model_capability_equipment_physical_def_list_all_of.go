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
)

// CapabilityEquipmentPhysicalDefListAllOf struct for CapabilityEquipmentPhysicalDefListAllOf
type CapabilityEquipmentPhysicalDefListAllOf struct {
	// The total number of 'capability.EquipmentPhysicalDef' resources matching the request, accross all pages. The 'Count' attribute is included when the HTTP GET request includes the '$inlinecount' parameter.
	Count *int32 `json:"Count,omitempty"`
	// The array of 'capability.EquipmentPhysicalDef' resources matching the request.
	Results              []CapabilityEquipmentPhysicalDef `json:"Results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityEquipmentPhysicalDefListAllOf CapabilityEquipmentPhysicalDefListAllOf

// NewCapabilityEquipmentPhysicalDefListAllOf instantiates a new CapabilityEquipmentPhysicalDefListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityEquipmentPhysicalDefListAllOf() *CapabilityEquipmentPhysicalDefListAllOf {
	this := CapabilityEquipmentPhysicalDefListAllOf{}
	return &this
}

// NewCapabilityEquipmentPhysicalDefListAllOfWithDefaults instantiates a new CapabilityEquipmentPhysicalDefListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityEquipmentPhysicalDefListAllOfWithDefaults() *CapabilityEquipmentPhysicalDefListAllOf {
	this := CapabilityEquipmentPhysicalDefListAllOf{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CapabilityEquipmentPhysicalDefListAllOf) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentPhysicalDefListAllOf) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CapabilityEquipmentPhysicalDefListAllOf) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CapabilityEquipmentPhysicalDefListAllOf) SetCount(v int32) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityEquipmentPhysicalDefListAllOf) GetResults() []CapabilityEquipmentPhysicalDef {
	if o == nil {
		var ret []CapabilityEquipmentPhysicalDef
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityEquipmentPhysicalDefListAllOf) GetResultsOk() (*[]CapabilityEquipmentPhysicalDef, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return &o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *CapabilityEquipmentPhysicalDefListAllOf) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []CapabilityEquipmentPhysicalDef and assigns it to the Results field.
func (o *CapabilityEquipmentPhysicalDefListAllOf) SetResults(v []CapabilityEquipmentPhysicalDef) {
	o.Results = v
}

func (o CapabilityEquipmentPhysicalDefListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.Results != nil {
		toSerialize["Results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilityEquipmentPhysicalDefListAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilityEquipmentPhysicalDefListAllOf := _CapabilityEquipmentPhysicalDefListAllOf{}

	if err = json.Unmarshal(bytes, &varCapabilityEquipmentPhysicalDefListAllOf); err == nil {
		*o = CapabilityEquipmentPhysicalDefListAllOf(varCapabilityEquipmentPhysicalDefListAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Count")
		delete(additionalProperties, "Results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilityEquipmentPhysicalDefListAllOf struct {
	value *CapabilityEquipmentPhysicalDefListAllOf
	isSet bool
}

func (v NullableCapabilityEquipmentPhysicalDefListAllOf) Get() *CapabilityEquipmentPhysicalDefListAllOf {
	return v.value
}

func (v *NullableCapabilityEquipmentPhysicalDefListAllOf) Set(val *CapabilityEquipmentPhysicalDefListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityEquipmentPhysicalDefListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityEquipmentPhysicalDefListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityEquipmentPhysicalDefListAllOf(val *CapabilityEquipmentPhysicalDefListAllOf) *NullableCapabilityEquipmentPhysicalDefListAllOf {
	return &NullableCapabilityEquipmentPhysicalDefListAllOf{value: val, isSet: true}
}

func (v NullableCapabilityEquipmentPhysicalDefListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityEquipmentPhysicalDefListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
