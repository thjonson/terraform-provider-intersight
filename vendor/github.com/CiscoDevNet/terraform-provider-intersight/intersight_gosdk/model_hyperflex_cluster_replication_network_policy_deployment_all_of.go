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

// HyperflexClusterReplicationNetworkPolicyDeploymentAllOf Definition of the list of properties defined in 'hyperflex.ClusterReplicationNetworkPolicyDeployment', excluding properties defined in parent classes.
type HyperflexClusterReplicationNetworkPolicyDeploymentAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Uuid of the HyperFlex cluster.
	ClusterUuid *string `json:"ClusterUuid,omitempty"`
	// Description from corresponding ClusterReplicationNetworkPolicy.
	Description *string `json:"Description,omitempty"`
	// True if record created by discovery on HyperFlex cluster.
	Discovered *bool `json:"Discovered,omitempty"`
	// Name from corresponding ClusterReplicationNetworkPolicy.
	Name *string `json:"Name,omitempty"`
	// Deployed network policy moid.
	PolicyMoid *string `json:"PolicyMoid,omitempty"`
	// Deployed cluster profile moid.
	ProfileMoid *string `json:"ProfileMoid,omitempty"`
	// Bandwidth for the Replication network in Mbps.
	ReplicationBandwidthMbps *int64                 `json:"ReplicationBandwidthMbps,omitempty"`
	ReplicationIpranges      []HyperflexIpAddrRange `json:"ReplicationIpranges,omitempty"`
	// MTU for the Replication network.
	ReplicationMtu  *int64                     `json:"ReplicationMtu,omitempty"`
	ReplicationVlan NullableHyperflexNamedVlan `json:"ReplicationVlan,omitempty"`
	// Unique request ID allowing retry of the same logical request following a transient communication failure.
	RequestId            *string                               `json:"RequestId,omitempty"`
	Cluster              *HyperflexClusterRelationship         `json:"Cluster,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexClusterReplicationNetworkPolicyDeploymentAllOf HyperflexClusterReplicationNetworkPolicyDeploymentAllOf

// NewHyperflexClusterReplicationNetworkPolicyDeploymentAllOf instantiates a new HyperflexClusterReplicationNetworkPolicyDeploymentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexClusterReplicationNetworkPolicyDeploymentAllOf(classId string, objectType string) *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf {
	this := HyperflexClusterReplicationNetworkPolicyDeploymentAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexClusterReplicationNetworkPolicyDeploymentAllOfWithDefaults instantiates a new HyperflexClusterReplicationNetworkPolicyDeploymentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexClusterReplicationNetworkPolicyDeploymentAllOfWithDefaults() *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf {
	this := HyperflexClusterReplicationNetworkPolicyDeploymentAllOf{}
	var classId string = "hyperflex.ClusterReplicationNetworkPolicyDeployment"
	this.ClassId = classId
	var objectType string = "hyperflex.ClusterReplicationNetworkPolicyDeployment"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterUuid returns the ClusterUuid field value if set, zero value otherwise.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetClusterUuid() string {
	if o == nil || o.ClusterUuid == nil {
		var ret string
		return ret
	}
	return *o.ClusterUuid
}

// GetClusterUuidOk returns a tuple with the ClusterUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetClusterUuidOk() (*string, bool) {
	if o == nil || o.ClusterUuid == nil {
		return nil, false
	}
	return o.ClusterUuid, true
}

// HasClusterUuid returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasClusterUuid() bool {
	if o != nil && o.ClusterUuid != nil {
		return true
	}

	return false
}

// SetClusterUuid gets a reference to the given string and assigns it to the ClusterUuid field.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetClusterUuid(v string) {
	o.ClusterUuid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetDiscovered returns the Discovered field value if set, zero value otherwise.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetDiscovered() bool {
	if o == nil || o.Discovered == nil {
		var ret bool
		return ret
	}
	return *o.Discovered
}

// GetDiscoveredOk returns a tuple with the Discovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetDiscoveredOk() (*bool, bool) {
	if o == nil || o.Discovered == nil {
		return nil, false
	}
	return o.Discovered, true
}

// HasDiscovered returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasDiscovered() bool {
	if o != nil && o.Discovered != nil {
		return true
	}

	return false
}

// SetDiscovered gets a reference to the given bool and assigns it to the Discovered field.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetDiscovered(v bool) {
	o.Discovered = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetName(v string) {
	o.Name = &v
}

// GetPolicyMoid returns the PolicyMoid field value if set, zero value otherwise.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetPolicyMoid() string {
	if o == nil || o.PolicyMoid == nil {
		var ret string
		return ret
	}
	return *o.PolicyMoid
}

// GetPolicyMoidOk returns a tuple with the PolicyMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetPolicyMoidOk() (*string, bool) {
	if o == nil || o.PolicyMoid == nil {
		return nil, false
	}
	return o.PolicyMoid, true
}

// HasPolicyMoid returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasPolicyMoid() bool {
	if o != nil && o.PolicyMoid != nil {
		return true
	}

	return false
}

// SetPolicyMoid gets a reference to the given string and assigns it to the PolicyMoid field.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetPolicyMoid(v string) {
	o.PolicyMoid = &v
}

// GetProfileMoid returns the ProfileMoid field value if set, zero value otherwise.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetProfileMoid() string {
	if o == nil || o.ProfileMoid == nil {
		var ret string
		return ret
	}
	return *o.ProfileMoid
}

// GetProfileMoidOk returns a tuple with the ProfileMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetProfileMoidOk() (*string, bool) {
	if o == nil || o.ProfileMoid == nil {
		return nil, false
	}
	return o.ProfileMoid, true
}

// HasProfileMoid returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasProfileMoid() bool {
	if o != nil && o.ProfileMoid != nil {
		return true
	}

	return false
}

// SetProfileMoid gets a reference to the given string and assigns it to the ProfileMoid field.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetProfileMoid(v string) {
	o.ProfileMoid = &v
}

// GetReplicationBandwidthMbps returns the ReplicationBandwidthMbps field value if set, zero value otherwise.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetReplicationBandwidthMbps() int64 {
	if o == nil || o.ReplicationBandwidthMbps == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationBandwidthMbps
}

// GetReplicationBandwidthMbpsOk returns a tuple with the ReplicationBandwidthMbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetReplicationBandwidthMbpsOk() (*int64, bool) {
	if o == nil || o.ReplicationBandwidthMbps == nil {
		return nil, false
	}
	return o.ReplicationBandwidthMbps, true
}

// HasReplicationBandwidthMbps returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasReplicationBandwidthMbps() bool {
	if o != nil && o.ReplicationBandwidthMbps != nil {
		return true
	}

	return false
}

// SetReplicationBandwidthMbps gets a reference to the given int64 and assigns it to the ReplicationBandwidthMbps field.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetReplicationBandwidthMbps(v int64) {
	o.ReplicationBandwidthMbps = &v
}

// GetReplicationIpranges returns the ReplicationIpranges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetReplicationIpranges() []HyperflexIpAddrRange {
	if o == nil {
		var ret []HyperflexIpAddrRange
		return ret
	}
	return o.ReplicationIpranges
}

// GetReplicationIprangesOk returns a tuple with the ReplicationIpranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetReplicationIprangesOk() (*[]HyperflexIpAddrRange, bool) {
	if o == nil || o.ReplicationIpranges == nil {
		return nil, false
	}
	return &o.ReplicationIpranges, true
}

// HasReplicationIpranges returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasReplicationIpranges() bool {
	if o != nil && o.ReplicationIpranges != nil {
		return true
	}

	return false
}

// SetReplicationIpranges gets a reference to the given []HyperflexIpAddrRange and assigns it to the ReplicationIpranges field.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetReplicationIpranges(v []HyperflexIpAddrRange) {
	o.ReplicationIpranges = v
}

// GetReplicationMtu returns the ReplicationMtu field value if set, zero value otherwise.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetReplicationMtu() int64 {
	if o == nil || o.ReplicationMtu == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationMtu
}

// GetReplicationMtuOk returns a tuple with the ReplicationMtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetReplicationMtuOk() (*int64, bool) {
	if o == nil || o.ReplicationMtu == nil {
		return nil, false
	}
	return o.ReplicationMtu, true
}

// HasReplicationMtu returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasReplicationMtu() bool {
	if o != nil && o.ReplicationMtu != nil {
		return true
	}

	return false
}

// SetReplicationMtu gets a reference to the given int64 and assigns it to the ReplicationMtu field.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetReplicationMtu(v int64) {
	o.ReplicationMtu = &v
}

// GetReplicationVlan returns the ReplicationVlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetReplicationVlan() HyperflexNamedVlan {
	if o == nil || o.ReplicationVlan.Get() == nil {
		var ret HyperflexNamedVlan
		return ret
	}
	return *o.ReplicationVlan.Get()
}

// GetReplicationVlanOk returns a tuple with the ReplicationVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetReplicationVlanOk() (*HyperflexNamedVlan, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReplicationVlan.Get(), o.ReplicationVlan.IsSet()
}

// HasReplicationVlan returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasReplicationVlan() bool {
	if o != nil && o.ReplicationVlan.IsSet() {
		return true
	}

	return false
}

// SetReplicationVlan gets a reference to the given NullableHyperflexNamedVlan and assigns it to the ReplicationVlan field.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetReplicationVlan(v HyperflexNamedVlan) {
	o.ReplicationVlan.Set(&v)
}

// SetReplicationVlanNil sets the value for ReplicationVlan to be an explicit nil
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetReplicationVlanNil() {
	o.ReplicationVlan.Set(nil)
}

// UnsetReplicationVlan ensures that no value is present for ReplicationVlan, not even an explicit nil
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) UnsetReplicationVlan() {
	o.ReplicationVlan.Unset()
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetRequestId(v string) {
	o.RequestId = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetCluster() HyperflexClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClusterUuid != nil {
		toSerialize["ClusterUuid"] = o.ClusterUuid
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Discovered != nil {
		toSerialize["Discovered"] = o.Discovered
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PolicyMoid != nil {
		toSerialize["PolicyMoid"] = o.PolicyMoid
	}
	if o.ProfileMoid != nil {
		toSerialize["ProfileMoid"] = o.ProfileMoid
	}
	if o.ReplicationBandwidthMbps != nil {
		toSerialize["ReplicationBandwidthMbps"] = o.ReplicationBandwidthMbps
	}
	if o.ReplicationIpranges != nil {
		toSerialize["ReplicationIpranges"] = o.ReplicationIpranges
	}
	if o.ReplicationMtu != nil {
		toSerialize["ReplicationMtu"] = o.ReplicationMtu
	}
	if o.ReplicationVlan.IsSet() {
		toSerialize["ReplicationVlan"] = o.ReplicationVlan.Get()
	}
	if o.RequestId != nil {
		toSerialize["RequestId"] = o.RequestId
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexClusterReplicationNetworkPolicyDeploymentAllOf := _HyperflexClusterReplicationNetworkPolicyDeploymentAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexClusterReplicationNetworkPolicyDeploymentAllOf); err == nil {
		*o = HyperflexClusterReplicationNetworkPolicyDeploymentAllOf(varHyperflexClusterReplicationNetworkPolicyDeploymentAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterUuid")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Discovered")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PolicyMoid")
		delete(additionalProperties, "ProfileMoid")
		delete(additionalProperties, "ReplicationBandwidthMbps")
		delete(additionalProperties, "ReplicationIpranges")
		delete(additionalProperties, "ReplicationMtu")
		delete(additionalProperties, "ReplicationVlan")
		delete(additionalProperties, "RequestId")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexClusterReplicationNetworkPolicyDeploymentAllOf struct {
	value *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf
	isSet bool
}

func (v NullableHyperflexClusterReplicationNetworkPolicyDeploymentAllOf) Get() *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf {
	return v.value
}

func (v *NullableHyperflexClusterReplicationNetworkPolicyDeploymentAllOf) Set(val *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexClusterReplicationNetworkPolicyDeploymentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexClusterReplicationNetworkPolicyDeploymentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexClusterReplicationNetworkPolicyDeploymentAllOf(val *HyperflexClusterReplicationNetworkPolicyDeploymentAllOf) *NullableHyperflexClusterReplicationNetworkPolicyDeploymentAllOf {
	return &NullableHyperflexClusterReplicationNetworkPolicyDeploymentAllOf{value: val, isSet: true}
}

func (v NullableHyperflexClusterReplicationNetworkPolicyDeploymentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexClusterReplicationNetworkPolicyDeploymentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
