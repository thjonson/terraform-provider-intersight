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

// AssetVirtualizationCloudOptionsAllOf Definition of the list of properties defined in 'asset.VirtualizationCloudOptions', excluding properties defined in parent classes.
type AssetVirtualizationCloudOptionsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType     string    `json:"ObjectType"`
	ManagedRegions []MoMoRef `json:"ManagedRegions,omitempty"`
	// The name of the region group to which the managedRegions belong. Populated from the group property in cloud.Regions.
	RegionGroup          *string `json:"RegionGroup,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetVirtualizationCloudOptionsAllOf AssetVirtualizationCloudOptionsAllOf

// NewAssetVirtualizationCloudOptionsAllOf instantiates a new AssetVirtualizationCloudOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetVirtualizationCloudOptionsAllOf(classId string, objectType string) *AssetVirtualizationCloudOptionsAllOf {
	this := AssetVirtualizationCloudOptionsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetVirtualizationCloudOptionsAllOfWithDefaults instantiates a new AssetVirtualizationCloudOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetVirtualizationCloudOptionsAllOfWithDefaults() *AssetVirtualizationCloudOptionsAllOf {
	this := AssetVirtualizationCloudOptionsAllOf{}
	var classId string = "asset.VirtualizationAmazonWebServiceOptions"
	this.ClassId = classId
	var objectType string = "asset.VirtualizationAmazonWebServiceOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetVirtualizationCloudOptionsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetVirtualizationCloudOptionsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetVirtualizationCloudOptionsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetVirtualizationCloudOptionsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetVirtualizationCloudOptionsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetVirtualizationCloudOptionsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetManagedRegions returns the ManagedRegions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetVirtualizationCloudOptionsAllOf) GetManagedRegions() []MoMoRef {
	if o == nil {
		var ret []MoMoRef
		return ret
	}
	return o.ManagedRegions
}

// GetManagedRegionsOk returns a tuple with the ManagedRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetVirtualizationCloudOptionsAllOf) GetManagedRegionsOk() ([]MoMoRef, bool) {
	if o == nil || o.ManagedRegions == nil {
		return nil, false
	}
	return o.ManagedRegions, true
}

// HasManagedRegions returns a boolean if a field has been set.
func (o *AssetVirtualizationCloudOptionsAllOf) HasManagedRegions() bool {
	if o != nil && o.ManagedRegions != nil {
		return true
	}

	return false
}

// SetManagedRegions gets a reference to the given []MoMoRef and assigns it to the ManagedRegions field.
func (o *AssetVirtualizationCloudOptionsAllOf) SetManagedRegions(v []MoMoRef) {
	o.ManagedRegions = v
}

// GetRegionGroup returns the RegionGroup field value if set, zero value otherwise.
func (o *AssetVirtualizationCloudOptionsAllOf) GetRegionGroup() string {
	if o == nil || o.RegionGroup == nil {
		var ret string
		return ret
	}
	return *o.RegionGroup
}

// GetRegionGroupOk returns a tuple with the RegionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetVirtualizationCloudOptionsAllOf) GetRegionGroupOk() (*string, bool) {
	if o == nil || o.RegionGroup == nil {
		return nil, false
	}
	return o.RegionGroup, true
}

// HasRegionGroup returns a boolean if a field has been set.
func (o *AssetVirtualizationCloudOptionsAllOf) HasRegionGroup() bool {
	if o != nil && o.RegionGroup != nil {
		return true
	}

	return false
}

// SetRegionGroup gets a reference to the given string and assigns it to the RegionGroup field.
func (o *AssetVirtualizationCloudOptionsAllOf) SetRegionGroup(v string) {
	o.RegionGroup = &v
}

func (o AssetVirtualizationCloudOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ManagedRegions != nil {
		toSerialize["ManagedRegions"] = o.ManagedRegions
	}
	if o.RegionGroup != nil {
		toSerialize["RegionGroup"] = o.RegionGroup
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetVirtualizationCloudOptionsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetVirtualizationCloudOptionsAllOf := _AssetVirtualizationCloudOptionsAllOf{}

	if err = json.Unmarshal(bytes, &varAssetVirtualizationCloudOptionsAllOf); err == nil {
		*o = AssetVirtualizationCloudOptionsAllOf(varAssetVirtualizationCloudOptionsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ManagedRegions")
		delete(additionalProperties, "RegionGroup")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetVirtualizationCloudOptionsAllOf struct {
	value *AssetVirtualizationCloudOptionsAllOf
	isSet bool
}

func (v NullableAssetVirtualizationCloudOptionsAllOf) Get() *AssetVirtualizationCloudOptionsAllOf {
	return v.value
}

func (v *NullableAssetVirtualizationCloudOptionsAllOf) Set(val *AssetVirtualizationCloudOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetVirtualizationCloudOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetVirtualizationCloudOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetVirtualizationCloudOptionsAllOf(val *AssetVirtualizationCloudOptionsAllOf) *NullableAssetVirtualizationCloudOptionsAllOf {
	return &NullableAssetVirtualizationCloudOptionsAllOf{value: val, isSet: true}
}

func (v NullableAssetVirtualizationCloudOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetVirtualizationCloudOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
