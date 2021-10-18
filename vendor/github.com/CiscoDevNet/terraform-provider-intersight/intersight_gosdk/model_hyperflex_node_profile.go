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
	"reflect"
	"strings"
)

// HyperflexNodeProfile A configuration profile per node in the HyperFlex cluster. It defines node settings such as IP address configuration for hypervisor management network, storage data network, HyperFlex management network, and the assigned physical server.
type HyperflexNodeProfile struct {
	PolicyAbstractConfigProfile
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IP address for storage data network (Controller VM interface).
	HxdpDataIp *string `json:"HxdpDataIp,omitempty"`
	// IP address for HyperFlex management network.
	HxdpMgmtIp *string `json:"HxdpMgmtIp,omitempty"`
	// IP address for storage client network (Controller VM interface).
	HxdpStorageClientIp *string `json:"HxdpStorageClientIp,omitempty"`
	// IP address for hypervisor control such as VM migration or pod management.
	HypervisorControlIp *string `json:"HypervisorControlIp,omitempty"`
	// IP address for storage data network (Hypervisor interface).
	HypervisorDataIp *string `json:"HypervisorDataIp,omitempty"`
	// IP address for Hypervisor management network.
	HypervisorMgmtIp *string `json:"HypervisorMgmtIp,omitempty"`
	// The role that this node performs in the HyperFlex cluster. * `Storage` - Cluster of storage nodes used to persist data. * `Compute` - Cluster of compute nodes used to execute business logic. * `Unknown` - This cluster type is Unknown. Expect Compute or Storage as valid values.
	NodeRole             *string                              `json:"NodeRole,omitempty"`
	AssignedServer       *ComputePhysicalRelationship         `json:"AssignedServer,omitempty"`
	ClusterProfile       *HyperflexClusterProfileRelationship `json:"ClusterProfile,omitempty"`
	Node                 *HyperflexNodeRelationship           `json:"Node,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexNodeProfile HyperflexNodeProfile

// NewHyperflexNodeProfile instantiates a new HyperflexNodeProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexNodeProfile(classId string, objectType string) *HyperflexNodeProfile {
	this := HyperflexNodeProfile{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "instance"
	this.Type = &type_
	var action string = "No-op"
	this.Action = &action
	return &this
}

// NewHyperflexNodeProfileWithDefaults instantiates a new HyperflexNodeProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexNodeProfileWithDefaults() *HyperflexNodeProfile {
	this := HyperflexNodeProfile{}
	var classId string = "hyperflex.NodeProfile"
	this.ClassId = classId
	var objectType string = "hyperflex.NodeProfile"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexNodeProfile) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexNodeProfile) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexNodeProfile) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexNodeProfile) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexNodeProfile) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexNodeProfile) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHxdpDataIp returns the HxdpDataIp field value if set, zero value otherwise.
func (o *HyperflexNodeProfile) GetHxdpDataIp() string {
	if o == nil || o.HxdpDataIp == nil {
		var ret string
		return ret
	}
	return *o.HxdpDataIp
}

// GetHxdpDataIpOk returns a tuple with the HxdpDataIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeProfile) GetHxdpDataIpOk() (*string, bool) {
	if o == nil || o.HxdpDataIp == nil {
		return nil, false
	}
	return o.HxdpDataIp, true
}

// HasHxdpDataIp returns a boolean if a field has been set.
func (o *HyperflexNodeProfile) HasHxdpDataIp() bool {
	if o != nil && o.HxdpDataIp != nil {
		return true
	}

	return false
}

// SetHxdpDataIp gets a reference to the given string and assigns it to the HxdpDataIp field.
func (o *HyperflexNodeProfile) SetHxdpDataIp(v string) {
	o.HxdpDataIp = &v
}

// GetHxdpMgmtIp returns the HxdpMgmtIp field value if set, zero value otherwise.
func (o *HyperflexNodeProfile) GetHxdpMgmtIp() string {
	if o == nil || o.HxdpMgmtIp == nil {
		var ret string
		return ret
	}
	return *o.HxdpMgmtIp
}

// GetHxdpMgmtIpOk returns a tuple with the HxdpMgmtIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeProfile) GetHxdpMgmtIpOk() (*string, bool) {
	if o == nil || o.HxdpMgmtIp == nil {
		return nil, false
	}
	return o.HxdpMgmtIp, true
}

// HasHxdpMgmtIp returns a boolean if a field has been set.
func (o *HyperflexNodeProfile) HasHxdpMgmtIp() bool {
	if o != nil && o.HxdpMgmtIp != nil {
		return true
	}

	return false
}

// SetHxdpMgmtIp gets a reference to the given string and assigns it to the HxdpMgmtIp field.
func (o *HyperflexNodeProfile) SetHxdpMgmtIp(v string) {
	o.HxdpMgmtIp = &v
}

// GetHxdpStorageClientIp returns the HxdpStorageClientIp field value if set, zero value otherwise.
func (o *HyperflexNodeProfile) GetHxdpStorageClientIp() string {
	if o == nil || o.HxdpStorageClientIp == nil {
		var ret string
		return ret
	}
	return *o.HxdpStorageClientIp
}

// GetHxdpStorageClientIpOk returns a tuple with the HxdpStorageClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeProfile) GetHxdpStorageClientIpOk() (*string, bool) {
	if o == nil || o.HxdpStorageClientIp == nil {
		return nil, false
	}
	return o.HxdpStorageClientIp, true
}

// HasHxdpStorageClientIp returns a boolean if a field has been set.
func (o *HyperflexNodeProfile) HasHxdpStorageClientIp() bool {
	if o != nil && o.HxdpStorageClientIp != nil {
		return true
	}

	return false
}

// SetHxdpStorageClientIp gets a reference to the given string and assigns it to the HxdpStorageClientIp field.
func (o *HyperflexNodeProfile) SetHxdpStorageClientIp(v string) {
	o.HxdpStorageClientIp = &v
}

// GetHypervisorControlIp returns the HypervisorControlIp field value if set, zero value otherwise.
func (o *HyperflexNodeProfile) GetHypervisorControlIp() string {
	if o == nil || o.HypervisorControlIp == nil {
		var ret string
		return ret
	}
	return *o.HypervisorControlIp
}

// GetHypervisorControlIpOk returns a tuple with the HypervisorControlIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeProfile) GetHypervisorControlIpOk() (*string, bool) {
	if o == nil || o.HypervisorControlIp == nil {
		return nil, false
	}
	return o.HypervisorControlIp, true
}

// HasHypervisorControlIp returns a boolean if a field has been set.
func (o *HyperflexNodeProfile) HasHypervisorControlIp() bool {
	if o != nil && o.HypervisorControlIp != nil {
		return true
	}

	return false
}

// SetHypervisorControlIp gets a reference to the given string and assigns it to the HypervisorControlIp field.
func (o *HyperflexNodeProfile) SetHypervisorControlIp(v string) {
	o.HypervisorControlIp = &v
}

// GetHypervisorDataIp returns the HypervisorDataIp field value if set, zero value otherwise.
func (o *HyperflexNodeProfile) GetHypervisorDataIp() string {
	if o == nil || o.HypervisorDataIp == nil {
		var ret string
		return ret
	}
	return *o.HypervisorDataIp
}

// GetHypervisorDataIpOk returns a tuple with the HypervisorDataIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeProfile) GetHypervisorDataIpOk() (*string, bool) {
	if o == nil || o.HypervisorDataIp == nil {
		return nil, false
	}
	return o.HypervisorDataIp, true
}

// HasHypervisorDataIp returns a boolean if a field has been set.
func (o *HyperflexNodeProfile) HasHypervisorDataIp() bool {
	if o != nil && o.HypervisorDataIp != nil {
		return true
	}

	return false
}

// SetHypervisorDataIp gets a reference to the given string and assigns it to the HypervisorDataIp field.
func (o *HyperflexNodeProfile) SetHypervisorDataIp(v string) {
	o.HypervisorDataIp = &v
}

// GetHypervisorMgmtIp returns the HypervisorMgmtIp field value if set, zero value otherwise.
func (o *HyperflexNodeProfile) GetHypervisorMgmtIp() string {
	if o == nil || o.HypervisorMgmtIp == nil {
		var ret string
		return ret
	}
	return *o.HypervisorMgmtIp
}

// GetHypervisorMgmtIpOk returns a tuple with the HypervisorMgmtIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeProfile) GetHypervisorMgmtIpOk() (*string, bool) {
	if o == nil || o.HypervisorMgmtIp == nil {
		return nil, false
	}
	return o.HypervisorMgmtIp, true
}

// HasHypervisorMgmtIp returns a boolean if a field has been set.
func (o *HyperflexNodeProfile) HasHypervisorMgmtIp() bool {
	if o != nil && o.HypervisorMgmtIp != nil {
		return true
	}

	return false
}

// SetHypervisorMgmtIp gets a reference to the given string and assigns it to the HypervisorMgmtIp field.
func (o *HyperflexNodeProfile) SetHypervisorMgmtIp(v string) {
	o.HypervisorMgmtIp = &v
}

// GetNodeRole returns the NodeRole field value if set, zero value otherwise.
func (o *HyperflexNodeProfile) GetNodeRole() string {
	if o == nil || o.NodeRole == nil {
		var ret string
		return ret
	}
	return *o.NodeRole
}

// GetNodeRoleOk returns a tuple with the NodeRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeProfile) GetNodeRoleOk() (*string, bool) {
	if o == nil || o.NodeRole == nil {
		return nil, false
	}
	return o.NodeRole, true
}

// HasNodeRole returns a boolean if a field has been set.
func (o *HyperflexNodeProfile) HasNodeRole() bool {
	if o != nil && o.NodeRole != nil {
		return true
	}

	return false
}

// SetNodeRole gets a reference to the given string and assigns it to the NodeRole field.
func (o *HyperflexNodeProfile) SetNodeRole(v string) {
	o.NodeRole = &v
}

// GetAssignedServer returns the AssignedServer field value if set, zero value otherwise.
func (o *HyperflexNodeProfile) GetAssignedServer() ComputePhysicalRelationship {
	if o == nil || o.AssignedServer == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.AssignedServer
}

// GetAssignedServerOk returns a tuple with the AssignedServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeProfile) GetAssignedServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.AssignedServer == nil {
		return nil, false
	}
	return o.AssignedServer, true
}

// HasAssignedServer returns a boolean if a field has been set.
func (o *HyperflexNodeProfile) HasAssignedServer() bool {
	if o != nil && o.AssignedServer != nil {
		return true
	}

	return false
}

// SetAssignedServer gets a reference to the given ComputePhysicalRelationship and assigns it to the AssignedServer field.
func (o *HyperflexNodeProfile) SetAssignedServer(v ComputePhysicalRelationship) {
	o.AssignedServer = &v
}

// GetClusterProfile returns the ClusterProfile field value if set, zero value otherwise.
func (o *HyperflexNodeProfile) GetClusterProfile() HyperflexClusterProfileRelationship {
	if o == nil || o.ClusterProfile == nil {
		var ret HyperflexClusterProfileRelationship
		return ret
	}
	return *o.ClusterProfile
}

// GetClusterProfileOk returns a tuple with the ClusterProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeProfile) GetClusterProfileOk() (*HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfile == nil {
		return nil, false
	}
	return o.ClusterProfile, true
}

// HasClusterProfile returns a boolean if a field has been set.
func (o *HyperflexNodeProfile) HasClusterProfile() bool {
	if o != nil && o.ClusterProfile != nil {
		return true
	}

	return false
}

// SetClusterProfile gets a reference to the given HyperflexClusterProfileRelationship and assigns it to the ClusterProfile field.
func (o *HyperflexNodeProfile) SetClusterProfile(v HyperflexClusterProfileRelationship) {
	o.ClusterProfile = &v
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *HyperflexNodeProfile) GetNode() HyperflexNodeRelationship {
	if o == nil || o.Node == nil {
		var ret HyperflexNodeRelationship
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeProfile) GetNodeOk() (*HyperflexNodeRelationship, bool) {
	if o == nil || o.Node == nil {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *HyperflexNodeProfile) HasNode() bool {
	if o != nil && o.Node != nil {
		return true
	}

	return false
}

// SetNode gets a reference to the given HyperflexNodeRelationship and assigns it to the Node field.
func (o *HyperflexNodeProfile) SetNode(v HyperflexNodeRelationship) {
	o.Node = &v
}

func (o HyperflexNodeProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigProfile, errPolicyAbstractConfigProfile := json.Marshal(o.PolicyAbstractConfigProfile)
	if errPolicyAbstractConfigProfile != nil {
		return []byte{}, errPolicyAbstractConfigProfile
	}
	errPolicyAbstractConfigProfile = json.Unmarshal([]byte(serializedPolicyAbstractConfigProfile), &toSerialize)
	if errPolicyAbstractConfigProfile != nil {
		return []byte{}, errPolicyAbstractConfigProfile
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.HxdpDataIp != nil {
		toSerialize["HxdpDataIp"] = o.HxdpDataIp
	}
	if o.HxdpMgmtIp != nil {
		toSerialize["HxdpMgmtIp"] = o.HxdpMgmtIp
	}
	if o.HxdpStorageClientIp != nil {
		toSerialize["HxdpStorageClientIp"] = o.HxdpStorageClientIp
	}
	if o.HypervisorControlIp != nil {
		toSerialize["HypervisorControlIp"] = o.HypervisorControlIp
	}
	if o.HypervisorDataIp != nil {
		toSerialize["HypervisorDataIp"] = o.HypervisorDataIp
	}
	if o.HypervisorMgmtIp != nil {
		toSerialize["HypervisorMgmtIp"] = o.HypervisorMgmtIp
	}
	if o.NodeRole != nil {
		toSerialize["NodeRole"] = o.NodeRole
	}
	if o.AssignedServer != nil {
		toSerialize["AssignedServer"] = o.AssignedServer
	}
	if o.ClusterProfile != nil {
		toSerialize["ClusterProfile"] = o.ClusterProfile
	}
	if o.Node != nil {
		toSerialize["Node"] = o.Node
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexNodeProfile) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexNodeProfileWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// IP address for storage data network (Controller VM interface).
		HxdpDataIp *string `json:"HxdpDataIp,omitempty"`
		// IP address for HyperFlex management network.
		HxdpMgmtIp *string `json:"HxdpMgmtIp,omitempty"`
		// IP address for storage client network (Controller VM interface).
		HxdpStorageClientIp *string `json:"HxdpStorageClientIp,omitempty"`
		// IP address for hypervisor control such as VM migration or pod management.
		HypervisorControlIp *string `json:"HypervisorControlIp,omitempty"`
		// IP address for storage data network (Hypervisor interface).
		HypervisorDataIp *string `json:"HypervisorDataIp,omitempty"`
		// IP address for Hypervisor management network.
		HypervisorMgmtIp *string `json:"HypervisorMgmtIp,omitempty"`
		// The role that this node performs in the HyperFlex cluster. * `Storage` - Cluster of storage nodes used to persist data. * `Compute` - Cluster of compute nodes used to execute business logic. * `Unknown` - This cluster type is Unknown. Expect Compute or Storage as valid values.
		NodeRole       *string                              `json:"NodeRole,omitempty"`
		AssignedServer *ComputePhysicalRelationship         `json:"AssignedServer,omitempty"`
		ClusterProfile *HyperflexClusterProfileRelationship `json:"ClusterProfile,omitempty"`
		Node           *HyperflexNodeRelationship           `json:"Node,omitempty"`
	}

	varHyperflexNodeProfileWithoutEmbeddedStruct := HyperflexNodeProfileWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexNodeProfileWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexNodeProfile := _HyperflexNodeProfile{}
		varHyperflexNodeProfile.ClassId = varHyperflexNodeProfileWithoutEmbeddedStruct.ClassId
		varHyperflexNodeProfile.ObjectType = varHyperflexNodeProfileWithoutEmbeddedStruct.ObjectType
		varHyperflexNodeProfile.HxdpDataIp = varHyperflexNodeProfileWithoutEmbeddedStruct.HxdpDataIp
		varHyperflexNodeProfile.HxdpMgmtIp = varHyperflexNodeProfileWithoutEmbeddedStruct.HxdpMgmtIp
		varHyperflexNodeProfile.HxdpStorageClientIp = varHyperflexNodeProfileWithoutEmbeddedStruct.HxdpStorageClientIp
		varHyperflexNodeProfile.HypervisorControlIp = varHyperflexNodeProfileWithoutEmbeddedStruct.HypervisorControlIp
		varHyperflexNodeProfile.HypervisorDataIp = varHyperflexNodeProfileWithoutEmbeddedStruct.HypervisorDataIp
		varHyperflexNodeProfile.HypervisorMgmtIp = varHyperflexNodeProfileWithoutEmbeddedStruct.HypervisorMgmtIp
		varHyperflexNodeProfile.NodeRole = varHyperflexNodeProfileWithoutEmbeddedStruct.NodeRole
		varHyperflexNodeProfile.AssignedServer = varHyperflexNodeProfileWithoutEmbeddedStruct.AssignedServer
		varHyperflexNodeProfile.ClusterProfile = varHyperflexNodeProfileWithoutEmbeddedStruct.ClusterProfile
		varHyperflexNodeProfile.Node = varHyperflexNodeProfileWithoutEmbeddedStruct.Node
		*o = HyperflexNodeProfile(varHyperflexNodeProfile)
	} else {
		return err
	}

	varHyperflexNodeProfile := _HyperflexNodeProfile{}

	err = json.Unmarshal(bytes, &varHyperflexNodeProfile)
	if err == nil {
		o.PolicyAbstractConfigProfile = varHyperflexNodeProfile.PolicyAbstractConfigProfile
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "HxdpDataIp")
		delete(additionalProperties, "HxdpMgmtIp")
		delete(additionalProperties, "HxdpStorageClientIp")
		delete(additionalProperties, "HypervisorControlIp")
		delete(additionalProperties, "HypervisorDataIp")
		delete(additionalProperties, "HypervisorMgmtIp")
		delete(additionalProperties, "NodeRole")
		delete(additionalProperties, "AssignedServer")
		delete(additionalProperties, "ClusterProfile")
		delete(additionalProperties, "Node")

		// remove fields from embedded structs
		reflectPolicyAbstractConfigProfile := reflect.ValueOf(o.PolicyAbstractConfigProfile)
		for i := 0; i < reflectPolicyAbstractConfigProfile.Type().NumField(); i++ {
			t := reflectPolicyAbstractConfigProfile.Type().Field(i)

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

type NullableHyperflexNodeProfile struct {
	value *HyperflexNodeProfile
	isSet bool
}

func (v NullableHyperflexNodeProfile) Get() *HyperflexNodeProfile {
	return v.value
}

func (v *NullableHyperflexNodeProfile) Set(val *HyperflexNodeProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexNodeProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexNodeProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexNodeProfile(val *HyperflexNodeProfile) *NullableHyperflexNodeProfile {
	return &NullableHyperflexNodeProfile{value: val, isSet: true}
}

func (v NullableHyperflexNodeProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexNodeProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
