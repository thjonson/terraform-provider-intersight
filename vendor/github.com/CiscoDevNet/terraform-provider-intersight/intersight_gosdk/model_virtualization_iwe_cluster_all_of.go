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

// VirtualizationIweClusterAllOf Definition of the list of properties defined in 'virtualization.IweCluster', excluding properties defined in parent classes.
type VirtualizationIweClusterAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of days remaining before the cluster's storage utilization reaches the recommended capacity limit of 76%. Default value is math.MaxInt32 to indicate that the capacity runway is \"Unknown\" for a cluster that is not connected or with not sufficient data.
	CapacityRunway *int64 `json:"CapacityRunway,omitempty"`
	// The name of this HyperFlex cluster.
	ClusterName *string `json:"ClusterName,omitempty"`
	// The number of compute nodes that belong to this cluster.
	ComputeNodeCount *int64 `json:"ComputeNodeCount,omitempty"`
	// CPU oversubscription factor configured on the cluster.
	ConfiguredCpuOverSubFactor *float64 `json:"ConfiguredCpuOverSubFactor,omitempty"`
	// The number of converged nodes that belong to this cluster.
	ConvergedNodeCount *int64                              `json:"ConvergedNodeCount,omitempty"`
	CpuAllocation      NullableVirtualizationCpuAllocation `json:"CpuAllocation,omitempty"`
	// Current oversubscription factor of the cluster.
	CurrentCpuOverSubFactor *float64 `json:"CurrentCpuOverSubFactor,omitempty"`
	// Datacenter to which the cluster belongs.
	DatacenterName *string `json:"DatacenterName,omitempty"`
	// The deployment type of the HyperFlex cluster. The cluster can have one of the following configurations: 1. Datacenter: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes on a single site. 2. Stretched Cluster: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes distributed across multiple sites. 3. Edge: The HyperFlex cluster consists of 2-4 standalone nodes. If the cluster is running a HyperFlex Data Platform version less than 4.0 or if the deployment type cannot be determined, the deployment type is set as 'NA' (not available). * `NA` - The deployment type of the cluster is not available. * `Datacenter` - The deployment type of a cluster consisting of UCS Fabric Interconnect-attached nodes on the same site. * `Stretched Cluster` - The deployment type of a cluster consisting of UCS Fabric Interconnect-attached nodes across different sites. * `Edge` - The deployment type of a cluster consisting of 2 or more standalone nodes.
	DeploymentType *string `json:"DeploymentType,omitempty"`
	// The type of the drives used for storage in this cluster. * `NA` - The drive type of the cluster is not available. * `All-Flash` - Indicates that this cluster contains flash drives only. * `Hybrid` - Indicates that this cluster contains both flash and hard disk drives.
	DriveType *string `json:"DriveType,omitempty"`
	// Reason for the failure when cluster is in failed state.
	FailureReason *string `json:"FailureReason,omitempty"`
	// The HyperFlex Data or Application Platform version of this cluster.
	HxVersion *string `json:"HxVersion,omitempty"`
	// Hypervisor version of HyperFlex compute cluster along with build number.
	HypervisorBuild *string `json:"HypervisorBuild,omitempty"`
	// Management IP Address of the cluster.
	ManagementIpAddress *string                                `json:"ManagementIpAddress,omitempty"`
	MemoryAllocation    NullableVirtualizationMemoryAllocation `json:"MemoryAllocation,omitempty"`
	// The storage capacity in this cluster.
	StorageCapacity *int64 `json:"StorageCapacity,omitempty"`
	// The number of storage nodes that belong to this cluster.
	StorageNodeCount *int64 `json:"StorageNodeCount,omitempty"`
	// The storage utilization is computed based on total capacity and current capacity utilization.
	StorageUtilization *float32 `json:"StorageUtilization,omitempty"`
	// The storage utilization percentage is computed based on total capacity and current capacity utilization.
	UtilizationPercentage *float32 `json:"UtilizationPercentage,omitempty"`
	// The storage utilization trend percentage represents the trend in percentage computed using the first and last point from historical data.
	UtilizationTrendPercentage *float32                             `json:"UtilizationTrendPercentage,omitempty"`
	AssociatedProfile          *PolicyAbstractProfileRelationship   `json:"AssociatedProfile,omitempty"`
	HxCluster                  *StorageBaseClusterRelationship      `json:"HxCluster,omitempty"`
	RegisteredDevice           *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _VirtualizationIweClusterAllOf VirtualizationIweClusterAllOf

// NewVirtualizationIweClusterAllOf instantiates a new VirtualizationIweClusterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationIweClusterAllOf(classId string, objectType string) *VirtualizationIweClusterAllOf {
	this := VirtualizationIweClusterAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationIweClusterAllOfWithDefaults instantiates a new VirtualizationIweClusterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationIweClusterAllOfWithDefaults() *VirtualizationIweClusterAllOf {
	this := VirtualizationIweClusterAllOf{}
	var classId string = "virtualization.IweCluster"
	this.ClassId = classId
	var objectType string = "virtualization.IweCluster"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationIweClusterAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationIweClusterAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationIweClusterAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationIweClusterAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCapacityRunway returns the CapacityRunway field value if set, zero value otherwise.
func (o *VirtualizationIweClusterAllOf) GetCapacityRunway() int64 {
	if o == nil || o.CapacityRunway == nil {
		var ret int64
		return ret
	}
	return *o.CapacityRunway
}

// GetCapacityRunwayOk returns a tuple with the CapacityRunway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetCapacityRunwayOk() (*int64, bool) {
	if o == nil || o.CapacityRunway == nil {
		return nil, false
	}
	return o.CapacityRunway, true
}

// HasCapacityRunway returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasCapacityRunway() bool {
	if o != nil && o.CapacityRunway != nil {
		return true
	}

	return false
}

// SetCapacityRunway gets a reference to the given int64 and assigns it to the CapacityRunway field.
func (o *VirtualizationIweClusterAllOf) SetCapacityRunway(v int64) {
	o.CapacityRunway = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *VirtualizationIweClusterAllOf) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *VirtualizationIweClusterAllOf) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetComputeNodeCount returns the ComputeNodeCount field value if set, zero value otherwise.
func (o *VirtualizationIweClusterAllOf) GetComputeNodeCount() int64 {
	if o == nil || o.ComputeNodeCount == nil {
		var ret int64
		return ret
	}
	return *o.ComputeNodeCount
}

// GetComputeNodeCountOk returns a tuple with the ComputeNodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetComputeNodeCountOk() (*int64, bool) {
	if o == nil || o.ComputeNodeCount == nil {
		return nil, false
	}
	return o.ComputeNodeCount, true
}

// HasComputeNodeCount returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasComputeNodeCount() bool {
	if o != nil && o.ComputeNodeCount != nil {
		return true
	}

	return false
}

// SetComputeNodeCount gets a reference to the given int64 and assigns it to the ComputeNodeCount field.
func (o *VirtualizationIweClusterAllOf) SetComputeNodeCount(v int64) {
	o.ComputeNodeCount = &v
}

// GetConfiguredCpuOverSubFactor returns the ConfiguredCpuOverSubFactor field value if set, zero value otherwise.
func (o *VirtualizationIweClusterAllOf) GetConfiguredCpuOverSubFactor() float64 {
	if o == nil || o.ConfiguredCpuOverSubFactor == nil {
		var ret float64
		return ret
	}
	return *o.ConfiguredCpuOverSubFactor
}

// GetConfiguredCpuOverSubFactorOk returns a tuple with the ConfiguredCpuOverSubFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetConfiguredCpuOverSubFactorOk() (*float64, bool) {
	if o == nil || o.ConfiguredCpuOverSubFactor == nil {
		return nil, false
	}
	return o.ConfiguredCpuOverSubFactor, true
}

// HasConfiguredCpuOverSubFactor returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasConfiguredCpuOverSubFactor() bool {
	if o != nil && o.ConfiguredCpuOverSubFactor != nil {
		return true
	}

	return false
}

// SetConfiguredCpuOverSubFactor gets a reference to the given float64 and assigns it to the ConfiguredCpuOverSubFactor field.
func (o *VirtualizationIweClusterAllOf) SetConfiguredCpuOverSubFactor(v float64) {
	o.ConfiguredCpuOverSubFactor = &v
}

// GetConvergedNodeCount returns the ConvergedNodeCount field value if set, zero value otherwise.
func (o *VirtualizationIweClusterAllOf) GetConvergedNodeCount() int64 {
	if o == nil || o.ConvergedNodeCount == nil {
		var ret int64
		return ret
	}
	return *o.ConvergedNodeCount
}

// GetConvergedNodeCountOk returns a tuple with the ConvergedNodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetConvergedNodeCountOk() (*int64, bool) {
	if o == nil || o.ConvergedNodeCount == nil {
		return nil, false
	}
	return o.ConvergedNodeCount, true
}

// HasConvergedNodeCount returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasConvergedNodeCount() bool {
	if o != nil && o.ConvergedNodeCount != nil {
		return true
	}

	return false
}

// SetConvergedNodeCount gets a reference to the given int64 and assigns it to the ConvergedNodeCount field.
func (o *VirtualizationIweClusterAllOf) SetConvergedNodeCount(v int64) {
	o.ConvergedNodeCount = &v
}

// GetCpuAllocation returns the CpuAllocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweClusterAllOf) GetCpuAllocation() VirtualizationCpuAllocation {
	if o == nil || o.CpuAllocation.Get() == nil {
		var ret VirtualizationCpuAllocation
		return ret
	}
	return *o.CpuAllocation.Get()
}

// GetCpuAllocationOk returns a tuple with the CpuAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweClusterAllOf) GetCpuAllocationOk() (*VirtualizationCpuAllocation, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuAllocation.Get(), o.CpuAllocation.IsSet()
}

// HasCpuAllocation returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasCpuAllocation() bool {
	if o != nil && o.CpuAllocation.IsSet() {
		return true
	}

	return false
}

// SetCpuAllocation gets a reference to the given NullableVirtualizationCpuAllocation and assigns it to the CpuAllocation field.
func (o *VirtualizationIweClusterAllOf) SetCpuAllocation(v VirtualizationCpuAllocation) {
	o.CpuAllocation.Set(&v)
}

// SetCpuAllocationNil sets the value for CpuAllocation to be an explicit nil
func (o *VirtualizationIweClusterAllOf) SetCpuAllocationNil() {
	o.CpuAllocation.Set(nil)
}

// UnsetCpuAllocation ensures that no value is present for CpuAllocation, not even an explicit nil
func (o *VirtualizationIweClusterAllOf) UnsetCpuAllocation() {
	o.CpuAllocation.Unset()
}

// GetCurrentCpuOverSubFactor returns the CurrentCpuOverSubFactor field value if set, zero value otherwise.
func (o *VirtualizationIweClusterAllOf) GetCurrentCpuOverSubFactor() float64 {
	if o == nil || o.CurrentCpuOverSubFactor == nil {
		var ret float64
		return ret
	}
	return *o.CurrentCpuOverSubFactor
}

// GetCurrentCpuOverSubFactorOk returns a tuple with the CurrentCpuOverSubFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetCurrentCpuOverSubFactorOk() (*float64, bool) {
	if o == nil || o.CurrentCpuOverSubFactor == nil {
		return nil, false
	}
	return o.CurrentCpuOverSubFactor, true
}

// HasCurrentCpuOverSubFactor returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasCurrentCpuOverSubFactor() bool {
	if o != nil && o.CurrentCpuOverSubFactor != nil {
		return true
	}

	return false
}

// SetCurrentCpuOverSubFactor gets a reference to the given float64 and assigns it to the CurrentCpuOverSubFactor field.
func (o *VirtualizationIweClusterAllOf) SetCurrentCpuOverSubFactor(v float64) {
	o.CurrentCpuOverSubFactor = &v
}

// GetDatacenterName returns the DatacenterName field value if set, zero value otherwise.
func (o *VirtualizationIweClusterAllOf) GetDatacenterName() string {
	if o == nil || o.DatacenterName == nil {
		var ret string
		return ret
	}
	return *o.DatacenterName
}

// GetDatacenterNameOk returns a tuple with the DatacenterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetDatacenterNameOk() (*string, bool) {
	if o == nil || o.DatacenterName == nil {
		return nil, false
	}
	return o.DatacenterName, true
}

// HasDatacenterName returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasDatacenterName() bool {
	if o != nil && o.DatacenterName != nil {
		return true
	}

	return false
}

// SetDatacenterName gets a reference to the given string and assigns it to the DatacenterName field.
func (o *VirtualizationIweClusterAllOf) SetDatacenterName(v string) {
	o.DatacenterName = &v
}

// GetDeploymentType returns the DeploymentType field value if set, zero value otherwise.
func (o *VirtualizationIweClusterAllOf) GetDeploymentType() string {
	if o == nil || o.DeploymentType == nil {
		var ret string
		return ret
	}
	return *o.DeploymentType
}

// GetDeploymentTypeOk returns a tuple with the DeploymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetDeploymentTypeOk() (*string, bool) {
	if o == nil || o.DeploymentType == nil {
		return nil, false
	}
	return o.DeploymentType, true
}

// HasDeploymentType returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasDeploymentType() bool {
	if o != nil && o.DeploymentType != nil {
		return true
	}

	return false
}

// SetDeploymentType gets a reference to the given string and assigns it to the DeploymentType field.
func (o *VirtualizationIweClusterAllOf) SetDeploymentType(v string) {
	o.DeploymentType = &v
}

// GetDriveType returns the DriveType field value if set, zero value otherwise.
func (o *VirtualizationIweClusterAllOf) GetDriveType() string {
	if o == nil || o.DriveType == nil {
		var ret string
		return ret
	}
	return *o.DriveType
}

// GetDriveTypeOk returns a tuple with the DriveType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetDriveTypeOk() (*string, bool) {
	if o == nil || o.DriveType == nil {
		return nil, false
	}
	return o.DriveType, true
}

// HasDriveType returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasDriveType() bool {
	if o != nil && o.DriveType != nil {
		return true
	}

	return false
}

// SetDriveType gets a reference to the given string and assigns it to the DriveType field.
func (o *VirtualizationIweClusterAllOf) SetDriveType(v string) {
	o.DriveType = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *VirtualizationIweClusterAllOf) GetFailureReason() string {
	if o == nil || o.FailureReason == nil {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetFailureReasonOk() (*string, bool) {
	if o == nil || o.FailureReason == nil {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasFailureReason() bool {
	if o != nil && o.FailureReason != nil {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *VirtualizationIweClusterAllOf) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetHxVersion returns the HxVersion field value if set, zero value otherwise.
func (o *VirtualizationIweClusterAllOf) GetHxVersion() string {
	if o == nil || o.HxVersion == nil {
		var ret string
		return ret
	}
	return *o.HxVersion
}

// GetHxVersionOk returns a tuple with the HxVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetHxVersionOk() (*string, bool) {
	if o == nil || o.HxVersion == nil {
		return nil, false
	}
	return o.HxVersion, true
}

// HasHxVersion returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasHxVersion() bool {
	if o != nil && o.HxVersion != nil {
		return true
	}

	return false
}

// SetHxVersion gets a reference to the given string and assigns it to the HxVersion field.
func (o *VirtualizationIweClusterAllOf) SetHxVersion(v string) {
	o.HxVersion = &v
}

// GetHypervisorBuild returns the HypervisorBuild field value if set, zero value otherwise.
func (o *VirtualizationIweClusterAllOf) GetHypervisorBuild() string {
	if o == nil || o.HypervisorBuild == nil {
		var ret string
		return ret
	}
	return *o.HypervisorBuild
}

// GetHypervisorBuildOk returns a tuple with the HypervisorBuild field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetHypervisorBuildOk() (*string, bool) {
	if o == nil || o.HypervisorBuild == nil {
		return nil, false
	}
	return o.HypervisorBuild, true
}

// HasHypervisorBuild returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasHypervisorBuild() bool {
	if o != nil && o.HypervisorBuild != nil {
		return true
	}

	return false
}

// SetHypervisorBuild gets a reference to the given string and assigns it to the HypervisorBuild field.
func (o *VirtualizationIweClusterAllOf) SetHypervisorBuild(v string) {
	o.HypervisorBuild = &v
}

// GetManagementIpAddress returns the ManagementIpAddress field value if set, zero value otherwise.
func (o *VirtualizationIweClusterAllOf) GetManagementIpAddress() string {
	if o == nil || o.ManagementIpAddress == nil {
		var ret string
		return ret
	}
	return *o.ManagementIpAddress
}

// GetManagementIpAddressOk returns a tuple with the ManagementIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetManagementIpAddressOk() (*string, bool) {
	if o == nil || o.ManagementIpAddress == nil {
		return nil, false
	}
	return o.ManagementIpAddress, true
}

// HasManagementIpAddress returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasManagementIpAddress() bool {
	if o != nil && o.ManagementIpAddress != nil {
		return true
	}

	return false
}

// SetManagementIpAddress gets a reference to the given string and assigns it to the ManagementIpAddress field.
func (o *VirtualizationIweClusterAllOf) SetManagementIpAddress(v string) {
	o.ManagementIpAddress = &v
}

// GetMemoryAllocation returns the MemoryAllocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweClusterAllOf) GetMemoryAllocation() VirtualizationMemoryAllocation {
	if o == nil || o.MemoryAllocation.Get() == nil {
		var ret VirtualizationMemoryAllocation
		return ret
	}
	return *o.MemoryAllocation.Get()
}

// GetMemoryAllocationOk returns a tuple with the MemoryAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweClusterAllOf) GetMemoryAllocationOk() (*VirtualizationMemoryAllocation, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryAllocation.Get(), o.MemoryAllocation.IsSet()
}

// HasMemoryAllocation returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasMemoryAllocation() bool {
	if o != nil && o.MemoryAllocation.IsSet() {
		return true
	}

	return false
}

// SetMemoryAllocation gets a reference to the given NullableVirtualizationMemoryAllocation and assigns it to the MemoryAllocation field.
func (o *VirtualizationIweClusterAllOf) SetMemoryAllocation(v VirtualizationMemoryAllocation) {
	o.MemoryAllocation.Set(&v)
}

// SetMemoryAllocationNil sets the value for MemoryAllocation to be an explicit nil
func (o *VirtualizationIweClusterAllOf) SetMemoryAllocationNil() {
	o.MemoryAllocation.Set(nil)
}

// UnsetMemoryAllocation ensures that no value is present for MemoryAllocation, not even an explicit nil
func (o *VirtualizationIweClusterAllOf) UnsetMemoryAllocation() {
	o.MemoryAllocation.Unset()
}

// GetStorageCapacity returns the StorageCapacity field value if set, zero value otherwise.
func (o *VirtualizationIweClusterAllOf) GetStorageCapacity() int64 {
	if o == nil || o.StorageCapacity == nil {
		var ret int64
		return ret
	}
	return *o.StorageCapacity
}

// GetStorageCapacityOk returns a tuple with the StorageCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetStorageCapacityOk() (*int64, bool) {
	if o == nil || o.StorageCapacity == nil {
		return nil, false
	}
	return o.StorageCapacity, true
}

// HasStorageCapacity returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasStorageCapacity() bool {
	if o != nil && o.StorageCapacity != nil {
		return true
	}

	return false
}

// SetStorageCapacity gets a reference to the given int64 and assigns it to the StorageCapacity field.
func (o *VirtualizationIweClusterAllOf) SetStorageCapacity(v int64) {
	o.StorageCapacity = &v
}

// GetStorageNodeCount returns the StorageNodeCount field value if set, zero value otherwise.
func (o *VirtualizationIweClusterAllOf) GetStorageNodeCount() int64 {
	if o == nil || o.StorageNodeCount == nil {
		var ret int64
		return ret
	}
	return *o.StorageNodeCount
}

// GetStorageNodeCountOk returns a tuple with the StorageNodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetStorageNodeCountOk() (*int64, bool) {
	if o == nil || o.StorageNodeCount == nil {
		return nil, false
	}
	return o.StorageNodeCount, true
}

// HasStorageNodeCount returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasStorageNodeCount() bool {
	if o != nil && o.StorageNodeCount != nil {
		return true
	}

	return false
}

// SetStorageNodeCount gets a reference to the given int64 and assigns it to the StorageNodeCount field.
func (o *VirtualizationIweClusterAllOf) SetStorageNodeCount(v int64) {
	o.StorageNodeCount = &v
}

// GetStorageUtilization returns the StorageUtilization field value if set, zero value otherwise.
func (o *VirtualizationIweClusterAllOf) GetStorageUtilization() float32 {
	if o == nil || o.StorageUtilization == nil {
		var ret float32
		return ret
	}
	return *o.StorageUtilization
}

// GetStorageUtilizationOk returns a tuple with the StorageUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetStorageUtilizationOk() (*float32, bool) {
	if o == nil || o.StorageUtilization == nil {
		return nil, false
	}
	return o.StorageUtilization, true
}

// HasStorageUtilization returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasStorageUtilization() bool {
	if o != nil && o.StorageUtilization != nil {
		return true
	}

	return false
}

// SetStorageUtilization gets a reference to the given float32 and assigns it to the StorageUtilization field.
func (o *VirtualizationIweClusterAllOf) SetStorageUtilization(v float32) {
	o.StorageUtilization = &v
}

// GetUtilizationPercentage returns the UtilizationPercentage field value if set, zero value otherwise.
func (o *VirtualizationIweClusterAllOf) GetUtilizationPercentage() float32 {
	if o == nil || o.UtilizationPercentage == nil {
		var ret float32
		return ret
	}
	return *o.UtilizationPercentage
}

// GetUtilizationPercentageOk returns a tuple with the UtilizationPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetUtilizationPercentageOk() (*float32, bool) {
	if o == nil || o.UtilizationPercentage == nil {
		return nil, false
	}
	return o.UtilizationPercentage, true
}

// HasUtilizationPercentage returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasUtilizationPercentage() bool {
	if o != nil && o.UtilizationPercentage != nil {
		return true
	}

	return false
}

// SetUtilizationPercentage gets a reference to the given float32 and assigns it to the UtilizationPercentage field.
func (o *VirtualizationIweClusterAllOf) SetUtilizationPercentage(v float32) {
	o.UtilizationPercentage = &v
}

// GetUtilizationTrendPercentage returns the UtilizationTrendPercentage field value if set, zero value otherwise.
func (o *VirtualizationIweClusterAllOf) GetUtilizationTrendPercentage() float32 {
	if o == nil || o.UtilizationTrendPercentage == nil {
		var ret float32
		return ret
	}
	return *o.UtilizationTrendPercentage
}

// GetUtilizationTrendPercentageOk returns a tuple with the UtilizationTrendPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetUtilizationTrendPercentageOk() (*float32, bool) {
	if o == nil || o.UtilizationTrendPercentage == nil {
		return nil, false
	}
	return o.UtilizationTrendPercentage, true
}

// HasUtilizationTrendPercentage returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasUtilizationTrendPercentage() bool {
	if o != nil && o.UtilizationTrendPercentage != nil {
		return true
	}

	return false
}

// SetUtilizationTrendPercentage gets a reference to the given float32 and assigns it to the UtilizationTrendPercentage field.
func (o *VirtualizationIweClusterAllOf) SetUtilizationTrendPercentage(v float32) {
	o.UtilizationTrendPercentage = &v
}

// GetAssociatedProfile returns the AssociatedProfile field value if set, zero value otherwise.
func (o *VirtualizationIweClusterAllOf) GetAssociatedProfile() PolicyAbstractProfileRelationship {
	if o == nil || o.AssociatedProfile == nil {
		var ret PolicyAbstractProfileRelationship
		return ret
	}
	return *o.AssociatedProfile
}

// GetAssociatedProfileOk returns a tuple with the AssociatedProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetAssociatedProfileOk() (*PolicyAbstractProfileRelationship, bool) {
	if o == nil || o.AssociatedProfile == nil {
		return nil, false
	}
	return o.AssociatedProfile, true
}

// HasAssociatedProfile returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasAssociatedProfile() bool {
	if o != nil && o.AssociatedProfile != nil {
		return true
	}

	return false
}

// SetAssociatedProfile gets a reference to the given PolicyAbstractProfileRelationship and assigns it to the AssociatedProfile field.
func (o *VirtualizationIweClusterAllOf) SetAssociatedProfile(v PolicyAbstractProfileRelationship) {
	o.AssociatedProfile = &v
}

// GetHxCluster returns the HxCluster field value if set, zero value otherwise.
func (o *VirtualizationIweClusterAllOf) GetHxCluster() StorageBaseClusterRelationship {
	if o == nil || o.HxCluster == nil {
		var ret StorageBaseClusterRelationship
		return ret
	}
	return *o.HxCluster
}

// GetHxClusterOk returns a tuple with the HxCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetHxClusterOk() (*StorageBaseClusterRelationship, bool) {
	if o == nil || o.HxCluster == nil {
		return nil, false
	}
	return o.HxCluster, true
}

// HasHxCluster returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasHxCluster() bool {
	if o != nil && o.HxCluster != nil {
		return true
	}

	return false
}

// SetHxCluster gets a reference to the given StorageBaseClusterRelationship and assigns it to the HxCluster field.
func (o *VirtualizationIweClusterAllOf) SetHxCluster(v StorageBaseClusterRelationship) {
	o.HxCluster = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *VirtualizationIweClusterAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweClusterAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *VirtualizationIweClusterAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *VirtualizationIweClusterAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o VirtualizationIweClusterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CapacityRunway != nil {
		toSerialize["CapacityRunway"] = o.CapacityRunway
	}
	if o.ClusterName != nil {
		toSerialize["ClusterName"] = o.ClusterName
	}
	if o.ComputeNodeCount != nil {
		toSerialize["ComputeNodeCount"] = o.ComputeNodeCount
	}
	if o.ConfiguredCpuOverSubFactor != nil {
		toSerialize["ConfiguredCpuOverSubFactor"] = o.ConfiguredCpuOverSubFactor
	}
	if o.ConvergedNodeCount != nil {
		toSerialize["ConvergedNodeCount"] = o.ConvergedNodeCount
	}
	if o.CpuAllocation.IsSet() {
		toSerialize["CpuAllocation"] = o.CpuAllocation.Get()
	}
	if o.CurrentCpuOverSubFactor != nil {
		toSerialize["CurrentCpuOverSubFactor"] = o.CurrentCpuOverSubFactor
	}
	if o.DatacenterName != nil {
		toSerialize["DatacenterName"] = o.DatacenterName
	}
	if o.DeploymentType != nil {
		toSerialize["DeploymentType"] = o.DeploymentType
	}
	if o.DriveType != nil {
		toSerialize["DriveType"] = o.DriveType
	}
	if o.FailureReason != nil {
		toSerialize["FailureReason"] = o.FailureReason
	}
	if o.HxVersion != nil {
		toSerialize["HxVersion"] = o.HxVersion
	}
	if o.HypervisorBuild != nil {
		toSerialize["HypervisorBuild"] = o.HypervisorBuild
	}
	if o.ManagementIpAddress != nil {
		toSerialize["ManagementIpAddress"] = o.ManagementIpAddress
	}
	if o.MemoryAllocation.IsSet() {
		toSerialize["MemoryAllocation"] = o.MemoryAllocation.Get()
	}
	if o.StorageCapacity != nil {
		toSerialize["StorageCapacity"] = o.StorageCapacity
	}
	if o.StorageNodeCount != nil {
		toSerialize["StorageNodeCount"] = o.StorageNodeCount
	}
	if o.StorageUtilization != nil {
		toSerialize["StorageUtilization"] = o.StorageUtilization
	}
	if o.UtilizationPercentage != nil {
		toSerialize["UtilizationPercentage"] = o.UtilizationPercentage
	}
	if o.UtilizationTrendPercentage != nil {
		toSerialize["UtilizationTrendPercentage"] = o.UtilizationTrendPercentage
	}
	if o.AssociatedProfile != nil {
		toSerialize["AssociatedProfile"] = o.AssociatedProfile
	}
	if o.HxCluster != nil {
		toSerialize["HxCluster"] = o.HxCluster
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationIweClusterAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationIweClusterAllOf := _VirtualizationIweClusterAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationIweClusterAllOf); err == nil {
		*o = VirtualizationIweClusterAllOf(varVirtualizationIweClusterAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CapacityRunway")
		delete(additionalProperties, "ClusterName")
		delete(additionalProperties, "ComputeNodeCount")
		delete(additionalProperties, "ConfiguredCpuOverSubFactor")
		delete(additionalProperties, "ConvergedNodeCount")
		delete(additionalProperties, "CpuAllocation")
		delete(additionalProperties, "CurrentCpuOverSubFactor")
		delete(additionalProperties, "DatacenterName")
		delete(additionalProperties, "DeploymentType")
		delete(additionalProperties, "DriveType")
		delete(additionalProperties, "FailureReason")
		delete(additionalProperties, "HxVersion")
		delete(additionalProperties, "HypervisorBuild")
		delete(additionalProperties, "ManagementIpAddress")
		delete(additionalProperties, "MemoryAllocation")
		delete(additionalProperties, "StorageCapacity")
		delete(additionalProperties, "StorageNodeCount")
		delete(additionalProperties, "StorageUtilization")
		delete(additionalProperties, "UtilizationPercentage")
		delete(additionalProperties, "UtilizationTrendPercentage")
		delete(additionalProperties, "AssociatedProfile")
		delete(additionalProperties, "HxCluster")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationIweClusterAllOf struct {
	value *VirtualizationIweClusterAllOf
	isSet bool
}

func (v NullableVirtualizationIweClusterAllOf) Get() *VirtualizationIweClusterAllOf {
	return v.value
}

func (v *NullableVirtualizationIweClusterAllOf) Set(val *VirtualizationIweClusterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationIweClusterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationIweClusterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationIweClusterAllOf(val *VirtualizationIweClusterAllOf) *NullableVirtualizationIweClusterAllOf {
	return &NullableVirtualizationIweClusterAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationIweClusterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationIweClusterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
