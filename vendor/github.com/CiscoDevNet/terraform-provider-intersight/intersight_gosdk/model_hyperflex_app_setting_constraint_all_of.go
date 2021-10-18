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

// HyperflexAppSettingConstraintAllOf Definition of the list of properties defined in 'hyperflex.AppSettingConstraint', excluding properties defined in parent classes.
type HyperflexAppSettingConstraintAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The deployment type of the cluster. * `NA` - The deployment type of the HyperFlex cluster is not available. * `Datacenter` - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes on the same site. * `Stretched Cluster` - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes across different sites. * `Edge` - The deployment type of a HyperFlex cluster consisting of 2 or more standalone nodes.
	DeploymentType *string `json:"DeploymentType,omitempty"`
	// The supported HyperFlex Data Platform version in regex format.
	HxdpVersion *string `json:"HxdpVersion,omitempty"`
	// The hypervisor type for the HyperFlex cluster. * `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * `HyperFlexAp` - The hypervisor of the virtualization platform is Cisco HyperFlex Application Platform. * `IWE` - The hypervisor of the virtualization platform is Cisco Intersight Workload Engine. * `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * `Unknown` - The hypervisor running on the HyperFlex cluster is not known.
	HypervisorType *string `json:"HypervisorType,omitempty"`
	// The supported management platform for the HyperFlex Cluster. * `FI` - The host servers used in the cluster deployment are managed by a UCS Fabric Interconnect. * `EDGE` - The host servers used in the cluster deployment are standalone severs.
	MgmtPlatform *string `json:"MgmtPlatform,omitempty"`
	// The supported server models in regex format.
	ServerModel          *string `json:"ServerModel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexAppSettingConstraintAllOf HyperflexAppSettingConstraintAllOf

// NewHyperflexAppSettingConstraintAllOf instantiates a new HyperflexAppSettingConstraintAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexAppSettingConstraintAllOf(classId string, objectType string) *HyperflexAppSettingConstraintAllOf {
	this := HyperflexAppSettingConstraintAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var deploymentType string = "NA"
	this.DeploymentType = &deploymentType
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	var mgmtPlatform string = "FI"
	this.MgmtPlatform = &mgmtPlatform
	return &this
}

// NewHyperflexAppSettingConstraintAllOfWithDefaults instantiates a new HyperflexAppSettingConstraintAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexAppSettingConstraintAllOfWithDefaults() *HyperflexAppSettingConstraintAllOf {
	this := HyperflexAppSettingConstraintAllOf{}
	var classId string = "hyperflex.AppSettingConstraint"
	this.ClassId = classId
	var objectType string = "hyperflex.AppSettingConstraint"
	this.ObjectType = objectType
	var deploymentType string = "NA"
	this.DeploymentType = &deploymentType
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	var mgmtPlatform string = "FI"
	this.MgmtPlatform = &mgmtPlatform
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexAppSettingConstraintAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexAppSettingConstraintAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexAppSettingConstraintAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexAppSettingConstraintAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexAppSettingConstraintAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexAppSettingConstraintAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeploymentType returns the DeploymentType field value if set, zero value otherwise.
func (o *HyperflexAppSettingConstraintAllOf) GetDeploymentType() string {
	if o == nil || o.DeploymentType == nil {
		var ret string
		return ret
	}
	return *o.DeploymentType
}

// GetDeploymentTypeOk returns a tuple with the DeploymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppSettingConstraintAllOf) GetDeploymentTypeOk() (*string, bool) {
	if o == nil || o.DeploymentType == nil {
		return nil, false
	}
	return o.DeploymentType, true
}

// HasDeploymentType returns a boolean if a field has been set.
func (o *HyperflexAppSettingConstraintAllOf) HasDeploymentType() bool {
	if o != nil && o.DeploymentType != nil {
		return true
	}

	return false
}

// SetDeploymentType gets a reference to the given string and assigns it to the DeploymentType field.
func (o *HyperflexAppSettingConstraintAllOf) SetDeploymentType(v string) {
	o.DeploymentType = &v
}

// GetHxdpVersion returns the HxdpVersion field value if set, zero value otherwise.
func (o *HyperflexAppSettingConstraintAllOf) GetHxdpVersion() string {
	if o == nil || o.HxdpVersion == nil {
		var ret string
		return ret
	}
	return *o.HxdpVersion
}

// GetHxdpVersionOk returns a tuple with the HxdpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppSettingConstraintAllOf) GetHxdpVersionOk() (*string, bool) {
	if o == nil || o.HxdpVersion == nil {
		return nil, false
	}
	return o.HxdpVersion, true
}

// HasHxdpVersion returns a boolean if a field has been set.
func (o *HyperflexAppSettingConstraintAllOf) HasHxdpVersion() bool {
	if o != nil && o.HxdpVersion != nil {
		return true
	}

	return false
}

// SetHxdpVersion gets a reference to the given string and assigns it to the HxdpVersion field.
func (o *HyperflexAppSettingConstraintAllOf) SetHxdpVersion(v string) {
	o.HxdpVersion = &v
}

// GetHypervisorType returns the HypervisorType field value if set, zero value otherwise.
func (o *HyperflexAppSettingConstraintAllOf) GetHypervisorType() string {
	if o == nil || o.HypervisorType == nil {
		var ret string
		return ret
	}
	return *o.HypervisorType
}

// GetHypervisorTypeOk returns a tuple with the HypervisorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppSettingConstraintAllOf) GetHypervisorTypeOk() (*string, bool) {
	if o == nil || o.HypervisorType == nil {
		return nil, false
	}
	return o.HypervisorType, true
}

// HasHypervisorType returns a boolean if a field has been set.
func (o *HyperflexAppSettingConstraintAllOf) HasHypervisorType() bool {
	if o != nil && o.HypervisorType != nil {
		return true
	}

	return false
}

// SetHypervisorType gets a reference to the given string and assigns it to the HypervisorType field.
func (o *HyperflexAppSettingConstraintAllOf) SetHypervisorType(v string) {
	o.HypervisorType = &v
}

// GetMgmtPlatform returns the MgmtPlatform field value if set, zero value otherwise.
func (o *HyperflexAppSettingConstraintAllOf) GetMgmtPlatform() string {
	if o == nil || o.MgmtPlatform == nil {
		var ret string
		return ret
	}
	return *o.MgmtPlatform
}

// GetMgmtPlatformOk returns a tuple with the MgmtPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppSettingConstraintAllOf) GetMgmtPlatformOk() (*string, bool) {
	if o == nil || o.MgmtPlatform == nil {
		return nil, false
	}
	return o.MgmtPlatform, true
}

// HasMgmtPlatform returns a boolean if a field has been set.
func (o *HyperflexAppSettingConstraintAllOf) HasMgmtPlatform() bool {
	if o != nil && o.MgmtPlatform != nil {
		return true
	}

	return false
}

// SetMgmtPlatform gets a reference to the given string and assigns it to the MgmtPlatform field.
func (o *HyperflexAppSettingConstraintAllOf) SetMgmtPlatform(v string) {
	o.MgmtPlatform = &v
}

// GetServerModel returns the ServerModel field value if set, zero value otherwise.
func (o *HyperflexAppSettingConstraintAllOf) GetServerModel() string {
	if o == nil || o.ServerModel == nil {
		var ret string
		return ret
	}
	return *o.ServerModel
}

// GetServerModelOk returns a tuple with the ServerModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppSettingConstraintAllOf) GetServerModelOk() (*string, bool) {
	if o == nil || o.ServerModel == nil {
		return nil, false
	}
	return o.ServerModel, true
}

// HasServerModel returns a boolean if a field has been set.
func (o *HyperflexAppSettingConstraintAllOf) HasServerModel() bool {
	if o != nil && o.ServerModel != nil {
		return true
	}

	return false
}

// SetServerModel gets a reference to the given string and assigns it to the ServerModel field.
func (o *HyperflexAppSettingConstraintAllOf) SetServerModel(v string) {
	o.ServerModel = &v
}

func (o HyperflexAppSettingConstraintAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DeploymentType != nil {
		toSerialize["DeploymentType"] = o.DeploymentType
	}
	if o.HxdpVersion != nil {
		toSerialize["HxdpVersion"] = o.HxdpVersion
	}
	if o.HypervisorType != nil {
		toSerialize["HypervisorType"] = o.HypervisorType
	}
	if o.MgmtPlatform != nil {
		toSerialize["MgmtPlatform"] = o.MgmtPlatform
	}
	if o.ServerModel != nil {
		toSerialize["ServerModel"] = o.ServerModel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexAppSettingConstraintAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexAppSettingConstraintAllOf := _HyperflexAppSettingConstraintAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexAppSettingConstraintAllOf); err == nil {
		*o = HyperflexAppSettingConstraintAllOf(varHyperflexAppSettingConstraintAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeploymentType")
		delete(additionalProperties, "HxdpVersion")
		delete(additionalProperties, "HypervisorType")
		delete(additionalProperties, "MgmtPlatform")
		delete(additionalProperties, "ServerModel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexAppSettingConstraintAllOf struct {
	value *HyperflexAppSettingConstraintAllOf
	isSet bool
}

func (v NullableHyperflexAppSettingConstraintAllOf) Get() *HyperflexAppSettingConstraintAllOf {
	return v.value
}

func (v *NullableHyperflexAppSettingConstraintAllOf) Set(val *HyperflexAppSettingConstraintAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexAppSettingConstraintAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexAppSettingConstraintAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexAppSettingConstraintAllOf(val *HyperflexAppSettingConstraintAllOf) *NullableHyperflexAppSettingConstraintAllOf {
	return &NullableHyperflexAppSettingConstraintAllOf{value: val, isSet: true}
}

func (v NullableHyperflexAppSettingConstraintAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexAppSettingConstraintAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
