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

// KubernetesEsxiVirtualMachineInfraConfigAllOf Definition of the list of properties defined in 'kubernetes.EsxiVirtualMachineInfraConfig', excluding properties defined in parent classes.
type KubernetesEsxiVirtualMachineInfraConfigAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the vSphere cluster on which the virtual machines are created.
	Cluster *string `json:"Cluster,omitempty"`
	// Name of the datasore on which the virtual machine disks are created.
	Datastore *string `json:"Datastore,omitempty"`
	// Indicates whether the value of the 'passphrase' property has been set.
	IsPassphraseSet *bool `json:"IsPassphraseSet,omitempty"`
	// Passphrase for the vcenter user.
	Passphrase *string `json:"Passphrase,omitempty"`
	// Name of the vSphere resource pool on which the virtual machines are created.
	ResourcePool         *string `json:"ResourcePool,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesEsxiVirtualMachineInfraConfigAllOf KubernetesEsxiVirtualMachineInfraConfigAllOf

// NewKubernetesEsxiVirtualMachineInfraConfigAllOf instantiates a new KubernetesEsxiVirtualMachineInfraConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesEsxiVirtualMachineInfraConfigAllOf(classId string, objectType string) *KubernetesEsxiVirtualMachineInfraConfigAllOf {
	this := KubernetesEsxiVirtualMachineInfraConfigAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesEsxiVirtualMachineInfraConfigAllOfWithDefaults instantiates a new KubernetesEsxiVirtualMachineInfraConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesEsxiVirtualMachineInfraConfigAllOfWithDefaults() *KubernetesEsxiVirtualMachineInfraConfigAllOf {
	this := KubernetesEsxiVirtualMachineInfraConfigAllOf{}
	var classId string = "kubernetes.EsxiVirtualMachineInfraConfig"
	this.ClassId = classId
	var objectType string = "kubernetes.EsxiVirtualMachineInfraConfig"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetCluster() string {
	if o == nil || o.Cluster == nil {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetClusterOk() (*string, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) SetCluster(v string) {
	o.Cluster = &v
}

// GetDatastore returns the Datastore field value if set, zero value otherwise.
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetDatastore() string {
	if o == nil || o.Datastore == nil {
		var ret string
		return ret
	}
	return *o.Datastore
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetDatastoreOk() (*string, bool) {
	if o == nil || o.Datastore == nil {
		return nil, false
	}
	return o.Datastore, true
}

// HasDatastore returns a boolean if a field has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) HasDatastore() bool {
	if o != nil && o.Datastore != nil {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given string and assigns it to the Datastore field.
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) SetDatastore(v string) {
	o.Datastore = &v
}

// GetIsPassphraseSet returns the IsPassphraseSet field value if set, zero value otherwise.
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetIsPassphraseSet() bool {
	if o == nil || o.IsPassphraseSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPassphraseSet
}

// GetIsPassphraseSetOk returns a tuple with the IsPassphraseSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetIsPassphraseSetOk() (*bool, bool) {
	if o == nil || o.IsPassphraseSet == nil {
		return nil, false
	}
	return o.IsPassphraseSet, true
}

// HasIsPassphraseSet returns a boolean if a field has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) HasIsPassphraseSet() bool {
	if o != nil && o.IsPassphraseSet != nil {
		return true
	}

	return false
}

// SetIsPassphraseSet gets a reference to the given bool and assigns it to the IsPassphraseSet field.
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) SetIsPassphraseSet(v bool) {
	o.IsPassphraseSet = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetPassphrase() string {
	if o == nil || o.Passphrase == nil {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetPassphraseOk() (*string, bool) {
	if o == nil || o.Passphrase == nil {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) HasPassphrase() bool {
	if o != nil && o.Passphrase != nil {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) SetPassphrase(v string) {
	o.Passphrase = &v
}

// GetResourcePool returns the ResourcePool field value if set, zero value otherwise.
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetResourcePool() string {
	if o == nil || o.ResourcePool == nil {
		var ret string
		return ret
	}
	return *o.ResourcePool
}

// GetResourcePoolOk returns a tuple with the ResourcePool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) GetResourcePoolOk() (*string, bool) {
	if o == nil || o.ResourcePool == nil {
		return nil, false
	}
	return o.ResourcePool, true
}

// HasResourcePool returns a boolean if a field has been set.
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) HasResourcePool() bool {
	if o != nil && o.ResourcePool != nil {
		return true
	}

	return false
}

// SetResourcePool gets a reference to the given string and assigns it to the ResourcePool field.
func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) SetResourcePool(v string) {
	o.ResourcePool = &v
}

func (o KubernetesEsxiVirtualMachineInfraConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.Datastore != nil {
		toSerialize["Datastore"] = o.Datastore
	}
	if o.IsPassphraseSet != nil {
		toSerialize["IsPassphraseSet"] = o.IsPassphraseSet
	}
	if o.Passphrase != nil {
		toSerialize["Passphrase"] = o.Passphrase
	}
	if o.ResourcePool != nil {
		toSerialize["ResourcePool"] = o.ResourcePool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesEsxiVirtualMachineInfraConfigAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesEsxiVirtualMachineInfraConfigAllOf := _KubernetesEsxiVirtualMachineInfraConfigAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesEsxiVirtualMachineInfraConfigAllOf); err == nil {
		*o = KubernetesEsxiVirtualMachineInfraConfigAllOf(varKubernetesEsxiVirtualMachineInfraConfigAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "Datastore")
		delete(additionalProperties, "IsPassphraseSet")
		delete(additionalProperties, "Passphrase")
		delete(additionalProperties, "ResourcePool")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesEsxiVirtualMachineInfraConfigAllOf struct {
	value *KubernetesEsxiVirtualMachineInfraConfigAllOf
	isSet bool
}

func (v NullableKubernetesEsxiVirtualMachineInfraConfigAllOf) Get() *KubernetesEsxiVirtualMachineInfraConfigAllOf {
	return v.value
}

func (v *NullableKubernetesEsxiVirtualMachineInfraConfigAllOf) Set(val *KubernetesEsxiVirtualMachineInfraConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesEsxiVirtualMachineInfraConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesEsxiVirtualMachineInfraConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesEsxiVirtualMachineInfraConfigAllOf(val *KubernetesEsxiVirtualMachineInfraConfigAllOf) *NullableKubernetesEsxiVirtualMachineInfraConfigAllOf {
	return &NullableKubernetesEsxiVirtualMachineInfraConfigAllOf{value: val, isSet: true}
}

func (v NullableKubernetesEsxiVirtualMachineInfraConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesEsxiVirtualMachineInfraConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
