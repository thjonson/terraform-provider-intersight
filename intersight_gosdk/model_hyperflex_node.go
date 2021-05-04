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

// HyperflexNode A host participating in the cluster. The host consists of a hypervisor installed on a node that manages virtual machines.
type HyperflexNode struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The build number of the hypervisor running on the host.
	BuildNumber *string `json:"BuildNumber,omitempty"`
	// The user-friendly string representation of the hypervisor version of the host.
	DisplayVersion *string `json:"DisplayVersion,omitempty"`
	// The hostname configured for the hypervisor running on the host.
	HostName *string `json:"HostName,omitempty"`
	// The type of hypervisor running on the host.
	Hypervisor *string                             `json:"Hypervisor,omitempty"`
	Identity   NullableHyperflexHxUuIdDt           `json:"Identity,omitempty"`
	Ip         NullableHyperflexHxNetworkAddressDt `json:"Ip,omitempty"`
	// The admin state of lockdown mode on the host. If 'true', lockdown mode is enabled.
	Lockdown *bool `json:"Lockdown,omitempty"`
	// The model of the host server.
	ModelNumber *string `json:"ModelNumber,omitempty"`
	// The role of the host in the HyperFlex cluster. Specifies whether this host is used for compute or for both compute and storage. * `UNKNOWN` - The role of the HyperFlex cluster node is not known. * `STORAGE` - The HyperFlex cluster node provides both storage and compute resources for the cluster. * `COMPUTE` - The HyperFlex cluster node provides compute resources for the cluster.
	Role *string `json:"Role,omitempty"`
	// The serial of the host server.
	SerialNumber *string `json:"SerialNumber,omitempty"`
	// The status of the host. Indicates whether the hypervisor is online. * `UNKNOWN` - The host status cannot be determined. * `ONLINE` - The host is online and operational. * `OFFLINE` - The host is offline and is currently not participating in the HyperFlex cluster. * `INMAINTENANCE` - The host is not participating in the HyperFlex cluster because of a maintenance operation, such as firmware or data platform upgrade. * `DEGRADED` - The host is degraded and may not be performing in its full operational capacity.
	Status *string `json:"Status,omitempty"`
	// The version of the hypervisor running on the host.
	Version              *string                         `json:"Version,omitempty"`
	Cluster              *HyperflexClusterRelationship   `json:"Cluster,omitempty"`
	ClusterMember        *AssetClusterMemberRelationship `json:"ClusterMember,omitempty"`
	PhysicalServer       *ComputePhysicalRelationship    `json:"PhysicalServer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexNode HyperflexNode

// NewHyperflexNode instantiates a new HyperflexNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexNode(classId string, objectType string) *HyperflexNode {
	this := HyperflexNode{}
	this.ClassId = classId
	this.ObjectType = objectType
	var role string = "UNKNOWN"
	this.Role = &role
	var status string = "UNKNOWN"
	this.Status = &status
	return &this
}

// NewHyperflexNodeWithDefaults instantiates a new HyperflexNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexNodeWithDefaults() *HyperflexNode {
	this := HyperflexNode{}
	var classId string = "hyperflex.Node"
	this.ClassId = classId
	var objectType string = "hyperflex.Node"
	this.ObjectType = objectType
	var role string = "UNKNOWN"
	this.Role = &role
	var status string = "UNKNOWN"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexNode) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexNode) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexNode) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexNode) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBuildNumber returns the BuildNumber field value if set, zero value otherwise.
func (o *HyperflexNode) GetBuildNumber() string {
	if o == nil || o.BuildNumber == nil {
		var ret string
		return ret
	}
	return *o.BuildNumber
}

// GetBuildNumberOk returns a tuple with the BuildNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetBuildNumberOk() (*string, bool) {
	if o == nil || o.BuildNumber == nil {
		return nil, false
	}
	return o.BuildNumber, true
}

// HasBuildNumber returns a boolean if a field has been set.
func (o *HyperflexNode) HasBuildNumber() bool {
	if o != nil && o.BuildNumber != nil {
		return true
	}

	return false
}

// SetBuildNumber gets a reference to the given string and assigns it to the BuildNumber field.
func (o *HyperflexNode) SetBuildNumber(v string) {
	o.BuildNumber = &v
}

// GetDisplayVersion returns the DisplayVersion field value if set, zero value otherwise.
func (o *HyperflexNode) GetDisplayVersion() string {
	if o == nil || o.DisplayVersion == nil {
		var ret string
		return ret
	}
	return *o.DisplayVersion
}

// GetDisplayVersionOk returns a tuple with the DisplayVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetDisplayVersionOk() (*string, bool) {
	if o == nil || o.DisplayVersion == nil {
		return nil, false
	}
	return o.DisplayVersion, true
}

// HasDisplayVersion returns a boolean if a field has been set.
func (o *HyperflexNode) HasDisplayVersion() bool {
	if o != nil && o.DisplayVersion != nil {
		return true
	}

	return false
}

// SetDisplayVersion gets a reference to the given string and assigns it to the DisplayVersion field.
func (o *HyperflexNode) SetDisplayVersion(v string) {
	o.DisplayVersion = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *HyperflexNode) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *HyperflexNode) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *HyperflexNode) SetHostName(v string) {
	o.HostName = &v
}

// GetHypervisor returns the Hypervisor field value if set, zero value otherwise.
func (o *HyperflexNode) GetHypervisor() string {
	if o == nil || o.Hypervisor == nil {
		var ret string
		return ret
	}
	return *o.Hypervisor
}

// GetHypervisorOk returns a tuple with the Hypervisor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetHypervisorOk() (*string, bool) {
	if o == nil || o.Hypervisor == nil {
		return nil, false
	}
	return o.Hypervisor, true
}

// HasHypervisor returns a boolean if a field has been set.
func (o *HyperflexNode) HasHypervisor() bool {
	if o != nil && o.Hypervisor != nil {
		return true
	}

	return false
}

// SetHypervisor gets a reference to the given string and assigns it to the Hypervisor field.
func (o *HyperflexNode) SetHypervisor(v string) {
	o.Hypervisor = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexNode) GetIdentity() HyperflexHxUuIdDt {
	if o == nil || o.Identity.Get() == nil {
		var ret HyperflexHxUuIdDt
		return ret
	}
	return *o.Identity.Get()
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexNode) GetIdentityOk() (*HyperflexHxUuIdDt, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identity.Get(), o.Identity.IsSet()
}

// HasIdentity returns a boolean if a field has been set.
func (o *HyperflexNode) HasIdentity() bool {
	if o != nil && o.Identity.IsSet() {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given NullableHyperflexHxUuIdDt and assigns it to the Identity field.
func (o *HyperflexNode) SetIdentity(v HyperflexHxUuIdDt) {
	o.Identity.Set(&v)
}

// SetIdentityNil sets the value for Identity to be an explicit nil
func (o *HyperflexNode) SetIdentityNil() {
	o.Identity.Set(nil)
}

// UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
func (o *HyperflexNode) UnsetIdentity() {
	o.Identity.Unset()
}

// GetIp returns the Ip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexNode) GetIp() HyperflexHxNetworkAddressDt {
	if o == nil || o.Ip.Get() == nil {
		var ret HyperflexHxNetworkAddressDt
		return ret
	}
	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexNode) GetIpOk() (*HyperflexHxNetworkAddressDt, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// HasIp returns a boolean if a field has been set.
func (o *HyperflexNode) HasIp() bool {
	if o != nil && o.Ip.IsSet() {
		return true
	}

	return false
}

// SetIp gets a reference to the given NullableHyperflexHxNetworkAddressDt and assigns it to the Ip field.
func (o *HyperflexNode) SetIp(v HyperflexHxNetworkAddressDt) {
	o.Ip.Set(&v)
}

// SetIpNil sets the value for Ip to be an explicit nil
func (o *HyperflexNode) SetIpNil() {
	o.Ip.Set(nil)
}

// UnsetIp ensures that no value is present for Ip, not even an explicit nil
func (o *HyperflexNode) UnsetIp() {
	o.Ip.Unset()
}

// GetLockdown returns the Lockdown field value if set, zero value otherwise.
func (o *HyperflexNode) GetLockdown() bool {
	if o == nil || o.Lockdown == nil {
		var ret bool
		return ret
	}
	return *o.Lockdown
}

// GetLockdownOk returns a tuple with the Lockdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetLockdownOk() (*bool, bool) {
	if o == nil || o.Lockdown == nil {
		return nil, false
	}
	return o.Lockdown, true
}

// HasLockdown returns a boolean if a field has been set.
func (o *HyperflexNode) HasLockdown() bool {
	if o != nil && o.Lockdown != nil {
		return true
	}

	return false
}

// SetLockdown gets a reference to the given bool and assigns it to the Lockdown field.
func (o *HyperflexNode) SetLockdown(v bool) {
	o.Lockdown = &v
}

// GetModelNumber returns the ModelNumber field value if set, zero value otherwise.
func (o *HyperflexNode) GetModelNumber() string {
	if o == nil || o.ModelNumber == nil {
		var ret string
		return ret
	}
	return *o.ModelNumber
}

// GetModelNumberOk returns a tuple with the ModelNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetModelNumberOk() (*string, bool) {
	if o == nil || o.ModelNumber == nil {
		return nil, false
	}
	return o.ModelNumber, true
}

// HasModelNumber returns a boolean if a field has been set.
func (o *HyperflexNode) HasModelNumber() bool {
	if o != nil && o.ModelNumber != nil {
		return true
	}

	return false
}

// SetModelNumber gets a reference to the given string and assigns it to the ModelNumber field.
func (o *HyperflexNode) SetModelNumber(v string) {
	o.ModelNumber = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *HyperflexNode) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *HyperflexNode) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *HyperflexNode) SetRole(v string) {
	o.Role = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *HyperflexNode) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *HyperflexNode) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *HyperflexNode) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HyperflexNode) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HyperflexNode) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HyperflexNode) SetStatus(v string) {
	o.Status = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexNode) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexNode) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexNode) SetVersion(v string) {
	o.Version = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexNode) GetCluster() HyperflexClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexNode) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexNode) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster = &v
}

// GetClusterMember returns the ClusterMember field value if set, zero value otherwise.
func (o *HyperflexNode) GetClusterMember() AssetClusterMemberRelationship {
	if o == nil || o.ClusterMember == nil {
		var ret AssetClusterMemberRelationship
		return ret
	}
	return *o.ClusterMember
}

// GetClusterMemberOk returns a tuple with the ClusterMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetClusterMemberOk() (*AssetClusterMemberRelationship, bool) {
	if o == nil || o.ClusterMember == nil {
		return nil, false
	}
	return o.ClusterMember, true
}

// HasClusterMember returns a boolean if a field has been set.
func (o *HyperflexNode) HasClusterMember() bool {
	if o != nil && o.ClusterMember != nil {
		return true
	}

	return false
}

// SetClusterMember gets a reference to the given AssetClusterMemberRelationship and assigns it to the ClusterMember field.
func (o *HyperflexNode) SetClusterMember(v AssetClusterMemberRelationship) {
	o.ClusterMember = &v
}

// GetPhysicalServer returns the PhysicalServer field value if set, zero value otherwise.
func (o *HyperflexNode) GetPhysicalServer() ComputePhysicalRelationship {
	if o == nil || o.PhysicalServer == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.PhysicalServer
}

// GetPhysicalServerOk returns a tuple with the PhysicalServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNode) GetPhysicalServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.PhysicalServer == nil {
		return nil, false
	}
	return o.PhysicalServer, true
}

// HasPhysicalServer returns a boolean if a field has been set.
func (o *HyperflexNode) HasPhysicalServer() bool {
	if o != nil && o.PhysicalServer != nil {
		return true
	}

	return false
}

// SetPhysicalServer gets a reference to the given ComputePhysicalRelationship and assigns it to the PhysicalServer field.
func (o *HyperflexNode) SetPhysicalServer(v ComputePhysicalRelationship) {
	o.PhysicalServer = &v
}

func (o HyperflexNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BuildNumber != nil {
		toSerialize["BuildNumber"] = o.BuildNumber
	}
	if o.DisplayVersion != nil {
		toSerialize["DisplayVersion"] = o.DisplayVersion
	}
	if o.HostName != nil {
		toSerialize["HostName"] = o.HostName
	}
	if o.Hypervisor != nil {
		toSerialize["Hypervisor"] = o.Hypervisor
	}
	if o.Identity.IsSet() {
		toSerialize["Identity"] = o.Identity.Get()
	}
	if o.Ip.IsSet() {
		toSerialize["Ip"] = o.Ip.Get()
	}
	if o.Lockdown != nil {
		toSerialize["Lockdown"] = o.Lockdown
	}
	if o.ModelNumber != nil {
		toSerialize["ModelNumber"] = o.ModelNumber
	}
	if o.Role != nil {
		toSerialize["Role"] = o.Role
	}
	if o.SerialNumber != nil {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.ClusterMember != nil {
		toSerialize["ClusterMember"] = o.ClusterMember
	}
	if o.PhysicalServer != nil {
		toSerialize["PhysicalServer"] = o.PhysicalServer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexNode) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexNodeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The build number of the hypervisor running on the host.
		BuildNumber *string `json:"BuildNumber,omitempty"`
		// The user-friendly string representation of the hypervisor version of the host.
		DisplayVersion *string `json:"DisplayVersion,omitempty"`
		// The hostname configured for the hypervisor running on the host.
		HostName *string `json:"HostName,omitempty"`
		// The type of hypervisor running on the host.
		Hypervisor *string                             `json:"Hypervisor,omitempty"`
		Identity   NullableHyperflexHxUuIdDt           `json:"Identity,omitempty"`
		Ip         NullableHyperflexHxNetworkAddressDt `json:"Ip,omitempty"`
		// The admin state of lockdown mode on the host. If 'true', lockdown mode is enabled.
		Lockdown *bool `json:"Lockdown,omitempty"`
		// The model of the host server.
		ModelNumber *string `json:"ModelNumber,omitempty"`
		// The role of the host in the HyperFlex cluster. Specifies whether this host is used for compute or for both compute and storage. * `UNKNOWN` - The role of the HyperFlex cluster node is not known. * `STORAGE` - The HyperFlex cluster node provides both storage and compute resources for the cluster. * `COMPUTE` - The HyperFlex cluster node provides compute resources for the cluster.
		Role *string `json:"Role,omitempty"`
		// The serial of the host server.
		SerialNumber *string `json:"SerialNumber,omitempty"`
		// The status of the host. Indicates whether the hypervisor is online. * `UNKNOWN` - The host status cannot be determined. * `ONLINE` - The host is online and operational. * `OFFLINE` - The host is offline and is currently not participating in the HyperFlex cluster. * `INMAINTENANCE` - The host is not participating in the HyperFlex cluster because of a maintenance operation, such as firmware or data platform upgrade. * `DEGRADED` - The host is degraded and may not be performing in its full operational capacity.
		Status *string `json:"Status,omitempty"`
		// The version of the hypervisor running on the host.
		Version        *string                         `json:"Version,omitempty"`
		Cluster        *HyperflexClusterRelationship   `json:"Cluster,omitempty"`
		ClusterMember  *AssetClusterMemberRelationship `json:"ClusterMember,omitempty"`
		PhysicalServer *ComputePhysicalRelationship    `json:"PhysicalServer,omitempty"`
	}

	varHyperflexNodeWithoutEmbeddedStruct := HyperflexNodeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexNodeWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexNode := _HyperflexNode{}
		varHyperflexNode.ClassId = varHyperflexNodeWithoutEmbeddedStruct.ClassId
		varHyperflexNode.ObjectType = varHyperflexNodeWithoutEmbeddedStruct.ObjectType
		varHyperflexNode.BuildNumber = varHyperflexNodeWithoutEmbeddedStruct.BuildNumber
		varHyperflexNode.DisplayVersion = varHyperflexNodeWithoutEmbeddedStruct.DisplayVersion
		varHyperflexNode.HostName = varHyperflexNodeWithoutEmbeddedStruct.HostName
		varHyperflexNode.Hypervisor = varHyperflexNodeWithoutEmbeddedStruct.Hypervisor
		varHyperflexNode.Identity = varHyperflexNodeWithoutEmbeddedStruct.Identity
		varHyperflexNode.Ip = varHyperflexNodeWithoutEmbeddedStruct.Ip
		varHyperflexNode.Lockdown = varHyperflexNodeWithoutEmbeddedStruct.Lockdown
		varHyperflexNode.ModelNumber = varHyperflexNodeWithoutEmbeddedStruct.ModelNumber
		varHyperflexNode.Role = varHyperflexNodeWithoutEmbeddedStruct.Role
		varHyperflexNode.SerialNumber = varHyperflexNodeWithoutEmbeddedStruct.SerialNumber
		varHyperflexNode.Status = varHyperflexNodeWithoutEmbeddedStruct.Status
		varHyperflexNode.Version = varHyperflexNodeWithoutEmbeddedStruct.Version
		varHyperflexNode.Cluster = varHyperflexNodeWithoutEmbeddedStruct.Cluster
		varHyperflexNode.ClusterMember = varHyperflexNodeWithoutEmbeddedStruct.ClusterMember
		varHyperflexNode.PhysicalServer = varHyperflexNodeWithoutEmbeddedStruct.PhysicalServer
		*o = HyperflexNode(varHyperflexNode)
	} else {
		return err
	}

	varHyperflexNode := _HyperflexNode{}

	err = json.Unmarshal(bytes, &varHyperflexNode)
	if err == nil {
		o.MoBaseMo = varHyperflexNode.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BuildNumber")
		delete(additionalProperties, "DisplayVersion")
		delete(additionalProperties, "HostName")
		delete(additionalProperties, "Hypervisor")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "Ip")
		delete(additionalProperties, "Lockdown")
		delete(additionalProperties, "ModelNumber")
		delete(additionalProperties, "Role")
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "ClusterMember")
		delete(additionalProperties, "PhysicalServer")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableHyperflexNode struct {
	value *HyperflexNode
	isSet bool
}

func (v NullableHyperflexNode) Get() *HyperflexNode {
	return v.value
}

func (v *NullableHyperflexNode) Set(val *HyperflexNode) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexNode) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexNode(val *HyperflexNode) *NullableHyperflexNode {
	return &NullableHyperflexNode{value: val, isSet: true}
}

func (v NullableHyperflexNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
