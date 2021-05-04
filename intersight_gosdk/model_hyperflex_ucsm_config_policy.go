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
	"reflect"
	"strings"
)

// HyperflexUcsmConfigPolicy A policy specifying UCS Manager settings such as service profile org, KVM IP addresses, and MAC prefix for server configuration in Fabric Interconnect environment.
type HyperflexUcsmConfigPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType     string                              `json:"ObjectType"`
	KvmIpRange     NullableHyperflexIpAddrRange        `json:"KvmIpRange,omitempty"`
	MacPrefixRange NullableHyperflexMacAddrPrefixRange `json:"MacPrefixRange,omitempty"`
	// The server firmware bundle version used for server components such as CIMC, adapters, BIOS, etc.
	ServerFirmwareVersion *string `json:"ServerFirmwareVersion,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles      []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexUcsmConfigPolicy HyperflexUcsmConfigPolicy

// NewHyperflexUcsmConfigPolicy instantiates a new HyperflexUcsmConfigPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexUcsmConfigPolicy(classId string, objectType string) *HyperflexUcsmConfigPolicy {
	this := HyperflexUcsmConfigPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexUcsmConfigPolicyWithDefaults instantiates a new HyperflexUcsmConfigPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexUcsmConfigPolicyWithDefaults() *HyperflexUcsmConfigPolicy {
	this := HyperflexUcsmConfigPolicy{}
	var classId string = "hyperflex.UcsmConfigPolicy"
	this.ClassId = classId
	var objectType string = "hyperflex.UcsmConfigPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexUcsmConfigPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexUcsmConfigPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexUcsmConfigPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexUcsmConfigPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexUcsmConfigPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexUcsmConfigPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetKvmIpRange returns the KvmIpRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexUcsmConfigPolicy) GetKvmIpRange() HyperflexIpAddrRange {
	if o == nil || o.KvmIpRange.Get() == nil {
		var ret HyperflexIpAddrRange
		return ret
	}
	return *o.KvmIpRange.Get()
}

// GetKvmIpRangeOk returns a tuple with the KvmIpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexUcsmConfigPolicy) GetKvmIpRangeOk() (*HyperflexIpAddrRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.KvmIpRange.Get(), o.KvmIpRange.IsSet()
}

// HasKvmIpRange returns a boolean if a field has been set.
func (o *HyperflexUcsmConfigPolicy) HasKvmIpRange() bool {
	if o != nil && o.KvmIpRange.IsSet() {
		return true
	}

	return false
}

// SetKvmIpRange gets a reference to the given NullableHyperflexIpAddrRange and assigns it to the KvmIpRange field.
func (o *HyperflexUcsmConfigPolicy) SetKvmIpRange(v HyperflexIpAddrRange) {
	o.KvmIpRange.Set(&v)
}

// SetKvmIpRangeNil sets the value for KvmIpRange to be an explicit nil
func (o *HyperflexUcsmConfigPolicy) SetKvmIpRangeNil() {
	o.KvmIpRange.Set(nil)
}

// UnsetKvmIpRange ensures that no value is present for KvmIpRange, not even an explicit nil
func (o *HyperflexUcsmConfigPolicy) UnsetKvmIpRange() {
	o.KvmIpRange.Unset()
}

// GetMacPrefixRange returns the MacPrefixRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexUcsmConfigPolicy) GetMacPrefixRange() HyperflexMacAddrPrefixRange {
	if o == nil || o.MacPrefixRange.Get() == nil {
		var ret HyperflexMacAddrPrefixRange
		return ret
	}
	return *o.MacPrefixRange.Get()
}

// GetMacPrefixRangeOk returns a tuple with the MacPrefixRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexUcsmConfigPolicy) GetMacPrefixRangeOk() (*HyperflexMacAddrPrefixRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.MacPrefixRange.Get(), o.MacPrefixRange.IsSet()
}

// HasMacPrefixRange returns a boolean if a field has been set.
func (o *HyperflexUcsmConfigPolicy) HasMacPrefixRange() bool {
	if o != nil && o.MacPrefixRange.IsSet() {
		return true
	}

	return false
}

// SetMacPrefixRange gets a reference to the given NullableHyperflexMacAddrPrefixRange and assigns it to the MacPrefixRange field.
func (o *HyperflexUcsmConfigPolicy) SetMacPrefixRange(v HyperflexMacAddrPrefixRange) {
	o.MacPrefixRange.Set(&v)
}

// SetMacPrefixRangeNil sets the value for MacPrefixRange to be an explicit nil
func (o *HyperflexUcsmConfigPolicy) SetMacPrefixRangeNil() {
	o.MacPrefixRange.Set(nil)
}

// UnsetMacPrefixRange ensures that no value is present for MacPrefixRange, not even an explicit nil
func (o *HyperflexUcsmConfigPolicy) UnsetMacPrefixRange() {
	o.MacPrefixRange.Unset()
}

// GetServerFirmwareVersion returns the ServerFirmwareVersion field value if set, zero value otherwise.
func (o *HyperflexUcsmConfigPolicy) GetServerFirmwareVersion() string {
	if o == nil || o.ServerFirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.ServerFirmwareVersion
}

// GetServerFirmwareVersionOk returns a tuple with the ServerFirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexUcsmConfigPolicy) GetServerFirmwareVersionOk() (*string, bool) {
	if o == nil || o.ServerFirmwareVersion == nil {
		return nil, false
	}
	return o.ServerFirmwareVersion, true
}

// HasServerFirmwareVersion returns a boolean if a field has been set.
func (o *HyperflexUcsmConfigPolicy) HasServerFirmwareVersion() bool {
	if o != nil && o.ServerFirmwareVersion != nil {
		return true
	}

	return false
}

// SetServerFirmwareVersion gets a reference to the given string and assigns it to the ServerFirmwareVersion field.
func (o *HyperflexUcsmConfigPolicy) SetServerFirmwareVersion(v string) {
	o.ServerFirmwareVersion = &v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexUcsmConfigPolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexUcsmConfigPolicy) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return &o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexUcsmConfigPolicy) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexUcsmConfigPolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexUcsmConfigPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexUcsmConfigPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexUcsmConfigPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexUcsmConfigPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o HyperflexUcsmConfigPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.KvmIpRange.IsSet() {
		toSerialize["KvmIpRange"] = o.KvmIpRange.Get()
	}
	if o.MacPrefixRange.IsSet() {
		toSerialize["MacPrefixRange"] = o.MacPrefixRange.Get()
	}
	if o.ServerFirmwareVersion != nil {
		toSerialize["ServerFirmwareVersion"] = o.ServerFirmwareVersion
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexUcsmConfigPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexUcsmConfigPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType     string                              `json:"ObjectType"`
		KvmIpRange     NullableHyperflexIpAddrRange        `json:"KvmIpRange,omitempty"`
		MacPrefixRange NullableHyperflexMacAddrPrefixRange `json:"MacPrefixRange,omitempty"`
		// The server firmware bundle version used for server components such as CIMC, adapters, BIOS, etc.
		ServerFirmwareVersion *string `json:"ServerFirmwareVersion,omitempty"`
		// An array of relationships to hyperflexClusterProfile resources.
		ClusterProfiles []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
		Organization    *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varHyperflexUcsmConfigPolicyWithoutEmbeddedStruct := HyperflexUcsmConfigPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexUcsmConfigPolicyWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexUcsmConfigPolicy := _HyperflexUcsmConfigPolicy{}
		varHyperflexUcsmConfigPolicy.ClassId = varHyperflexUcsmConfigPolicyWithoutEmbeddedStruct.ClassId
		varHyperflexUcsmConfigPolicy.ObjectType = varHyperflexUcsmConfigPolicyWithoutEmbeddedStruct.ObjectType
		varHyperflexUcsmConfigPolicy.KvmIpRange = varHyperflexUcsmConfigPolicyWithoutEmbeddedStruct.KvmIpRange
		varHyperflexUcsmConfigPolicy.MacPrefixRange = varHyperflexUcsmConfigPolicyWithoutEmbeddedStruct.MacPrefixRange
		varHyperflexUcsmConfigPolicy.ServerFirmwareVersion = varHyperflexUcsmConfigPolicyWithoutEmbeddedStruct.ServerFirmwareVersion
		varHyperflexUcsmConfigPolicy.ClusterProfiles = varHyperflexUcsmConfigPolicyWithoutEmbeddedStruct.ClusterProfiles
		varHyperflexUcsmConfigPolicy.Organization = varHyperflexUcsmConfigPolicyWithoutEmbeddedStruct.Organization
		*o = HyperflexUcsmConfigPolicy(varHyperflexUcsmConfigPolicy)
	} else {
		return err
	}

	varHyperflexUcsmConfigPolicy := _HyperflexUcsmConfigPolicy{}

	err = json.Unmarshal(bytes, &varHyperflexUcsmConfigPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varHyperflexUcsmConfigPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "KvmIpRange")
		delete(additionalProperties, "MacPrefixRange")
		delete(additionalProperties, "ServerFirmwareVersion")
		delete(additionalProperties, "ClusterProfiles")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableHyperflexUcsmConfigPolicy struct {
	value *HyperflexUcsmConfigPolicy
	isSet bool
}

func (v NullableHyperflexUcsmConfigPolicy) Get() *HyperflexUcsmConfigPolicy {
	return v.value
}

func (v *NullableHyperflexUcsmConfigPolicy) Set(val *HyperflexUcsmConfigPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexUcsmConfigPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexUcsmConfigPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexUcsmConfigPolicy(val *HyperflexUcsmConfigPolicy) *NullableHyperflexUcsmConfigPolicy {
	return &NullableHyperflexUcsmConfigPolicy{value: val, isSet: true}
}

func (v NullableHyperflexUcsmConfigPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexUcsmConfigPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
