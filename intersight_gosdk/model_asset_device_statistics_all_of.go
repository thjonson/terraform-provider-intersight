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

// AssetDeviceStatisticsAllOf Definition of the list of properties defined in 'asset.DeviceStatistics', excluding properties defined in parent classes.
type AssetDeviceStatisticsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Deployment type of HyperFlex cluster.
	ClusterDeploymentType *string `json:"ClusterDeploymentType,omitempty"`
	// Reference to HyperFlex cluster target device ID.
	ClusterDeviceMoid *string `json:"ClusterDeviceMoid,omitempty"`
	// Name of the cluster. It is specified only for HyperFlex based devices.
	ClusterName *string `json:"ClusterName,omitempty"`
	// Data replication factor of HyperFlex cluster.
	ClusterReplicationFactor *int64 `json:"ClusterReplicationFactor,omitempty"`
	// The status of the persistent connection between the device connector and Intersight, for HyperFlex or UCS device. 1 represents being connected and 0 represents being disconnected.
	Connected *int64 `json:"Connected,omitempty"`
	// Defines the average proportion of resources used by the device within the cluster. example in a cluster having 3 nodes, the membershipRatio of each node is 1/3 or 0.33. It is specified only for HyperFlex based devices.
	MembershipRatio *float32 `json:"MembershipRatio,omitempty"`
	// Memory Reliability, availability and serviceability (RAS) factor.
	MemoryMirroringFactor *float32            `json:"MemoryMirroringFactor,omitempty"`
	VmHost                NullableAssetVmHost `json:"VmHost,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _AssetDeviceStatisticsAllOf AssetDeviceStatisticsAllOf

// NewAssetDeviceStatisticsAllOf instantiates a new AssetDeviceStatisticsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDeviceStatisticsAllOf(classId string, objectType string) *AssetDeviceStatisticsAllOf {
	this := AssetDeviceStatisticsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetDeviceStatisticsAllOfWithDefaults instantiates a new AssetDeviceStatisticsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDeviceStatisticsAllOfWithDefaults() *AssetDeviceStatisticsAllOf {
	this := AssetDeviceStatisticsAllOf{}
	var classId string = "asset.DeviceStatistics"
	this.ClassId = classId
	var objectType string = "asset.DeviceStatistics"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetDeviceStatisticsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatisticsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetDeviceStatisticsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetDeviceStatisticsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatisticsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetDeviceStatisticsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterDeploymentType returns the ClusterDeploymentType field value if set, zero value otherwise.
func (o *AssetDeviceStatisticsAllOf) GetClusterDeploymentType() string {
	if o == nil || o.ClusterDeploymentType == nil {
		var ret string
		return ret
	}
	return *o.ClusterDeploymentType
}

// GetClusterDeploymentTypeOk returns a tuple with the ClusterDeploymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatisticsAllOf) GetClusterDeploymentTypeOk() (*string, bool) {
	if o == nil || o.ClusterDeploymentType == nil {
		return nil, false
	}
	return o.ClusterDeploymentType, true
}

// HasClusterDeploymentType returns a boolean if a field has been set.
func (o *AssetDeviceStatisticsAllOf) HasClusterDeploymentType() bool {
	if o != nil && o.ClusterDeploymentType != nil {
		return true
	}

	return false
}

// SetClusterDeploymentType gets a reference to the given string and assigns it to the ClusterDeploymentType field.
func (o *AssetDeviceStatisticsAllOf) SetClusterDeploymentType(v string) {
	o.ClusterDeploymentType = &v
}

// GetClusterDeviceMoid returns the ClusterDeviceMoid field value if set, zero value otherwise.
func (o *AssetDeviceStatisticsAllOf) GetClusterDeviceMoid() string {
	if o == nil || o.ClusterDeviceMoid == nil {
		var ret string
		return ret
	}
	return *o.ClusterDeviceMoid
}

// GetClusterDeviceMoidOk returns a tuple with the ClusterDeviceMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatisticsAllOf) GetClusterDeviceMoidOk() (*string, bool) {
	if o == nil || o.ClusterDeviceMoid == nil {
		return nil, false
	}
	return o.ClusterDeviceMoid, true
}

// HasClusterDeviceMoid returns a boolean if a field has been set.
func (o *AssetDeviceStatisticsAllOf) HasClusterDeviceMoid() bool {
	if o != nil && o.ClusterDeviceMoid != nil {
		return true
	}

	return false
}

// SetClusterDeviceMoid gets a reference to the given string and assigns it to the ClusterDeviceMoid field.
func (o *AssetDeviceStatisticsAllOf) SetClusterDeviceMoid(v string) {
	o.ClusterDeviceMoid = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *AssetDeviceStatisticsAllOf) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatisticsAllOf) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *AssetDeviceStatisticsAllOf) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *AssetDeviceStatisticsAllOf) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetClusterReplicationFactor returns the ClusterReplicationFactor field value if set, zero value otherwise.
func (o *AssetDeviceStatisticsAllOf) GetClusterReplicationFactor() int64 {
	if o == nil || o.ClusterReplicationFactor == nil {
		var ret int64
		return ret
	}
	return *o.ClusterReplicationFactor
}

// GetClusterReplicationFactorOk returns a tuple with the ClusterReplicationFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatisticsAllOf) GetClusterReplicationFactorOk() (*int64, bool) {
	if o == nil || o.ClusterReplicationFactor == nil {
		return nil, false
	}
	return o.ClusterReplicationFactor, true
}

// HasClusterReplicationFactor returns a boolean if a field has been set.
func (o *AssetDeviceStatisticsAllOf) HasClusterReplicationFactor() bool {
	if o != nil && o.ClusterReplicationFactor != nil {
		return true
	}

	return false
}

// SetClusterReplicationFactor gets a reference to the given int64 and assigns it to the ClusterReplicationFactor field.
func (o *AssetDeviceStatisticsAllOf) SetClusterReplicationFactor(v int64) {
	o.ClusterReplicationFactor = &v
}

// GetConnected returns the Connected field value if set, zero value otherwise.
func (o *AssetDeviceStatisticsAllOf) GetConnected() int64 {
	if o == nil || o.Connected == nil {
		var ret int64
		return ret
	}
	return *o.Connected
}

// GetConnectedOk returns a tuple with the Connected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatisticsAllOf) GetConnectedOk() (*int64, bool) {
	if o == nil || o.Connected == nil {
		return nil, false
	}
	return o.Connected, true
}

// HasConnected returns a boolean if a field has been set.
func (o *AssetDeviceStatisticsAllOf) HasConnected() bool {
	if o != nil && o.Connected != nil {
		return true
	}

	return false
}

// SetConnected gets a reference to the given int64 and assigns it to the Connected field.
func (o *AssetDeviceStatisticsAllOf) SetConnected(v int64) {
	o.Connected = &v
}

// GetMembershipRatio returns the MembershipRatio field value if set, zero value otherwise.
func (o *AssetDeviceStatisticsAllOf) GetMembershipRatio() float32 {
	if o == nil || o.MembershipRatio == nil {
		var ret float32
		return ret
	}
	return *o.MembershipRatio
}

// GetMembershipRatioOk returns a tuple with the MembershipRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatisticsAllOf) GetMembershipRatioOk() (*float32, bool) {
	if o == nil || o.MembershipRatio == nil {
		return nil, false
	}
	return o.MembershipRatio, true
}

// HasMembershipRatio returns a boolean if a field has been set.
func (o *AssetDeviceStatisticsAllOf) HasMembershipRatio() bool {
	if o != nil && o.MembershipRatio != nil {
		return true
	}

	return false
}

// SetMembershipRatio gets a reference to the given float32 and assigns it to the MembershipRatio field.
func (o *AssetDeviceStatisticsAllOf) SetMembershipRatio(v float32) {
	o.MembershipRatio = &v
}

// GetMemoryMirroringFactor returns the MemoryMirroringFactor field value if set, zero value otherwise.
func (o *AssetDeviceStatisticsAllOf) GetMemoryMirroringFactor() float32 {
	if o == nil || o.MemoryMirroringFactor == nil {
		var ret float32
		return ret
	}
	return *o.MemoryMirroringFactor
}

// GetMemoryMirroringFactorOk returns a tuple with the MemoryMirroringFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatisticsAllOf) GetMemoryMirroringFactorOk() (*float32, bool) {
	if o == nil || o.MemoryMirroringFactor == nil {
		return nil, false
	}
	return o.MemoryMirroringFactor, true
}

// HasMemoryMirroringFactor returns a boolean if a field has been set.
func (o *AssetDeviceStatisticsAllOf) HasMemoryMirroringFactor() bool {
	if o != nil && o.MemoryMirroringFactor != nil {
		return true
	}

	return false
}

// SetMemoryMirroringFactor gets a reference to the given float32 and assigns it to the MemoryMirroringFactor field.
func (o *AssetDeviceStatisticsAllOf) SetMemoryMirroringFactor(v float32) {
	o.MemoryMirroringFactor = &v
}

// GetVmHost returns the VmHost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeviceStatisticsAllOf) GetVmHost() AssetVmHost {
	if o == nil || o.VmHost.Get() == nil {
		var ret AssetVmHost
		return ret
	}
	return *o.VmHost.Get()
}

// GetVmHostOk returns a tuple with the VmHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeviceStatisticsAllOf) GetVmHostOk() (*AssetVmHost, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmHost.Get(), o.VmHost.IsSet()
}

// HasVmHost returns a boolean if a field has been set.
func (o *AssetDeviceStatisticsAllOf) HasVmHost() bool {
	if o != nil && o.VmHost.IsSet() {
		return true
	}

	return false
}

// SetVmHost gets a reference to the given NullableAssetVmHost and assigns it to the VmHost field.
func (o *AssetDeviceStatisticsAllOf) SetVmHost(v AssetVmHost) {
	o.VmHost.Set(&v)
}

// SetVmHostNil sets the value for VmHost to be an explicit nil
func (o *AssetDeviceStatisticsAllOf) SetVmHostNil() {
	o.VmHost.Set(nil)
}

// UnsetVmHost ensures that no value is present for VmHost, not even an explicit nil
func (o *AssetDeviceStatisticsAllOf) UnsetVmHost() {
	o.VmHost.Unset()
}

func (o AssetDeviceStatisticsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClusterDeploymentType != nil {
		toSerialize["ClusterDeploymentType"] = o.ClusterDeploymentType
	}
	if o.ClusterDeviceMoid != nil {
		toSerialize["ClusterDeviceMoid"] = o.ClusterDeviceMoid
	}
	if o.ClusterName != nil {
		toSerialize["ClusterName"] = o.ClusterName
	}
	if o.ClusterReplicationFactor != nil {
		toSerialize["ClusterReplicationFactor"] = o.ClusterReplicationFactor
	}
	if o.Connected != nil {
		toSerialize["Connected"] = o.Connected
	}
	if o.MembershipRatio != nil {
		toSerialize["MembershipRatio"] = o.MembershipRatio
	}
	if o.MemoryMirroringFactor != nil {
		toSerialize["MemoryMirroringFactor"] = o.MemoryMirroringFactor
	}
	if o.VmHost.IsSet() {
		toSerialize["VmHost"] = o.VmHost.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetDeviceStatisticsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetDeviceStatisticsAllOf := _AssetDeviceStatisticsAllOf{}

	if err = json.Unmarshal(bytes, &varAssetDeviceStatisticsAllOf); err == nil {
		*o = AssetDeviceStatisticsAllOf(varAssetDeviceStatisticsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterDeploymentType")
		delete(additionalProperties, "ClusterDeviceMoid")
		delete(additionalProperties, "ClusterName")
		delete(additionalProperties, "ClusterReplicationFactor")
		delete(additionalProperties, "Connected")
		delete(additionalProperties, "MembershipRatio")
		delete(additionalProperties, "MemoryMirroringFactor")
		delete(additionalProperties, "VmHost")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetDeviceStatisticsAllOf struct {
	value *AssetDeviceStatisticsAllOf
	isSet bool
}

func (v NullableAssetDeviceStatisticsAllOf) Get() *AssetDeviceStatisticsAllOf {
	return v.value
}

func (v *NullableAssetDeviceStatisticsAllOf) Set(val *AssetDeviceStatisticsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDeviceStatisticsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDeviceStatisticsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDeviceStatisticsAllOf(val *AssetDeviceStatisticsAllOf) *NullableAssetDeviceStatisticsAllOf {
	return &NullableAssetDeviceStatisticsAllOf{value: val, isSet: true}
}

func (v NullableAssetDeviceStatisticsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDeviceStatisticsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
