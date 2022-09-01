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

// AssetOrchestrationHsmVmwareVcenterOptionsAllOf Definition of the list of properties defined in 'asset.OrchestrationHsmVmwareVcenterOptions', excluding properties defined in parent classes.
type AssetOrchestrationHsmVmwareVcenterOptionsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// HsmEnabled controls whether Hardware Support Manager is enabled or not. vCenter Server version 7.0 or later.
	HsmEnabled *bool `json:"HsmEnabled,omitempty"`
	// Indicates whether the value of the 'clientSecret' property has been set.
	IsClientSecretSet    *bool `json:"IsClientSecretSet,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetOrchestrationHsmVmwareVcenterOptionsAllOf AssetOrchestrationHsmVmwareVcenterOptionsAllOf

// NewAssetOrchestrationHsmVmwareVcenterOptionsAllOf instantiates a new AssetOrchestrationHsmVmwareVcenterOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetOrchestrationHsmVmwareVcenterOptionsAllOf(classId string, objectType string) *AssetOrchestrationHsmVmwareVcenterOptionsAllOf {
	this := AssetOrchestrationHsmVmwareVcenterOptionsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetOrchestrationHsmVmwareVcenterOptionsAllOfWithDefaults instantiates a new AssetOrchestrationHsmVmwareVcenterOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetOrchestrationHsmVmwareVcenterOptionsAllOfWithDefaults() *AssetOrchestrationHsmVmwareVcenterOptionsAllOf {
	this := AssetOrchestrationHsmVmwareVcenterOptionsAllOf{}
	var classId string = "asset.OrchestrationHsmVmwareVcenterOptions"
	this.ClassId = classId
	var objectType string = "asset.OrchestrationHsmVmwareVcenterOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHsmEnabled returns the HsmEnabled field value if set, zero value otherwise.
func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) GetHsmEnabled() bool {
	if o == nil || o.HsmEnabled == nil {
		var ret bool
		return ret
	}
	return *o.HsmEnabled
}

// GetHsmEnabledOk returns a tuple with the HsmEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) GetHsmEnabledOk() (*bool, bool) {
	if o == nil || o.HsmEnabled == nil {
		return nil, false
	}
	return o.HsmEnabled, true
}

// HasHsmEnabled returns a boolean if a field has been set.
func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) HasHsmEnabled() bool {
	if o != nil && o.HsmEnabled != nil {
		return true
	}

	return false
}

// SetHsmEnabled gets a reference to the given bool and assigns it to the HsmEnabled field.
func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) SetHsmEnabled(v bool) {
	o.HsmEnabled = &v
}

// GetIsClientSecretSet returns the IsClientSecretSet field value if set, zero value otherwise.
func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) GetIsClientSecretSet() bool {
	if o == nil || o.IsClientSecretSet == nil {
		var ret bool
		return ret
	}
	return *o.IsClientSecretSet
}

// GetIsClientSecretSetOk returns a tuple with the IsClientSecretSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) GetIsClientSecretSetOk() (*bool, bool) {
	if o == nil || o.IsClientSecretSet == nil {
		return nil, false
	}
	return o.IsClientSecretSet, true
}

// HasIsClientSecretSet returns a boolean if a field has been set.
func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) HasIsClientSecretSet() bool {
	if o != nil && o.IsClientSecretSet != nil {
		return true
	}

	return false
}

// SetIsClientSecretSet gets a reference to the given bool and assigns it to the IsClientSecretSet field.
func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) SetIsClientSecretSet(v bool) {
	o.IsClientSecretSet = &v
}

func (o AssetOrchestrationHsmVmwareVcenterOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.HsmEnabled != nil {
		toSerialize["HsmEnabled"] = o.HsmEnabled
	}
	if o.IsClientSecretSet != nil {
		toSerialize["IsClientSecretSet"] = o.IsClientSecretSet
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetOrchestrationHsmVmwareVcenterOptionsAllOf := _AssetOrchestrationHsmVmwareVcenterOptionsAllOf{}

	if err = json.Unmarshal(bytes, &varAssetOrchestrationHsmVmwareVcenterOptionsAllOf); err == nil {
		*o = AssetOrchestrationHsmVmwareVcenterOptionsAllOf(varAssetOrchestrationHsmVmwareVcenterOptionsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "HsmEnabled")
		delete(additionalProperties, "IsClientSecretSet")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetOrchestrationHsmVmwareVcenterOptionsAllOf struct {
	value *AssetOrchestrationHsmVmwareVcenterOptionsAllOf
	isSet bool
}

func (v NullableAssetOrchestrationHsmVmwareVcenterOptionsAllOf) Get() *AssetOrchestrationHsmVmwareVcenterOptionsAllOf {
	return v.value
}

func (v *NullableAssetOrchestrationHsmVmwareVcenterOptionsAllOf) Set(val *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetOrchestrationHsmVmwareVcenterOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetOrchestrationHsmVmwareVcenterOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetOrchestrationHsmVmwareVcenterOptionsAllOf(val *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) *NullableAssetOrchestrationHsmVmwareVcenterOptionsAllOf {
	return &NullableAssetOrchestrationHsmVmwareVcenterOptionsAllOf{value: val, isSet: true}
}

func (v NullableAssetOrchestrationHsmVmwareVcenterOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetOrchestrationHsmVmwareVcenterOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
