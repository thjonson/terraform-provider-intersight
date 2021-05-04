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

// KubernetesVersion A policy capturing software versions for a Kubernetes cluster.
type KubernetesVersion struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Desired Kubernetes version.
	KubernetesVersion *string `json:"KubernetesVersion,omitempty"`
	// The name of this IKS kubernetes version.
	Name                        *string                                    `json:"Name,omitempty"`
	BootIso                     *SoftwareSolutionDistributableRelationship `json:"BootIso,omitempty"`
	Catalog                     *KubernetesCatalogRelationship             `json:"Catalog,omitempty"`
	OvaImageTemplate            *SoftwareSolutionDistributableRelationship `json:"OvaImageTemplate,omitempty"`
	Qcow2NodeTemplate           *SoftwareSolutionDistributableRelationship `json:"Qcow2NodeTemplate,omitempty"`
	Qcow2VirtualMachineTemplate *SoftwareSolutionDistributableRelationship `json:"Qcow2VirtualMachineTemplate,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _KubernetesVersion KubernetesVersion

// NewKubernetesVersion instantiates a new KubernetesVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesVersion(classId string, objectType string) *KubernetesVersion {
	this := KubernetesVersion{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesVersionWithDefaults instantiates a new KubernetesVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesVersionWithDefaults() *KubernetesVersion {
	this := KubernetesVersion{}
	var classId string = "kubernetes.Version"
	this.ClassId = classId
	var objectType string = "kubernetes.Version"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesVersion) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesVersion) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesVersion) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesVersion) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesVersion) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesVersion) SetObjectType(v string) {
	o.ObjectType = v
}

// GetKubernetesVersion returns the KubernetesVersion field value if set, zero value otherwise.
func (o *KubernetesVersion) GetKubernetesVersion() string {
	if o == nil || o.KubernetesVersion == nil {
		var ret string
		return ret
	}
	return *o.KubernetesVersion
}

// GetKubernetesVersionOk returns a tuple with the KubernetesVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVersion) GetKubernetesVersionOk() (*string, bool) {
	if o == nil || o.KubernetesVersion == nil {
		return nil, false
	}
	return o.KubernetesVersion, true
}

// HasKubernetesVersion returns a boolean if a field has been set.
func (o *KubernetesVersion) HasKubernetesVersion() bool {
	if o != nil && o.KubernetesVersion != nil {
		return true
	}

	return false
}

// SetKubernetesVersion gets a reference to the given string and assigns it to the KubernetesVersion field.
func (o *KubernetesVersion) SetKubernetesVersion(v string) {
	o.KubernetesVersion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesVersion) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVersion) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesVersion) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesVersion) SetName(v string) {
	o.Name = &v
}

// GetBootIso returns the BootIso field value if set, zero value otherwise.
func (o *KubernetesVersion) GetBootIso() SoftwareSolutionDistributableRelationship {
	if o == nil || o.BootIso == nil {
		var ret SoftwareSolutionDistributableRelationship
		return ret
	}
	return *o.BootIso
}

// GetBootIsoOk returns a tuple with the BootIso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVersion) GetBootIsoOk() (*SoftwareSolutionDistributableRelationship, bool) {
	if o == nil || o.BootIso == nil {
		return nil, false
	}
	return o.BootIso, true
}

// HasBootIso returns a boolean if a field has been set.
func (o *KubernetesVersion) HasBootIso() bool {
	if o != nil && o.BootIso != nil {
		return true
	}

	return false
}

// SetBootIso gets a reference to the given SoftwareSolutionDistributableRelationship and assigns it to the BootIso field.
func (o *KubernetesVersion) SetBootIso(v SoftwareSolutionDistributableRelationship) {
	o.BootIso = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *KubernetesVersion) GetCatalog() KubernetesCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret KubernetesCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVersion) GetCatalogOk() (*KubernetesCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *KubernetesVersion) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given KubernetesCatalogRelationship and assigns it to the Catalog field.
func (o *KubernetesVersion) SetCatalog(v KubernetesCatalogRelationship) {
	o.Catalog = &v
}

// GetOvaImageTemplate returns the OvaImageTemplate field value if set, zero value otherwise.
func (o *KubernetesVersion) GetOvaImageTemplate() SoftwareSolutionDistributableRelationship {
	if o == nil || o.OvaImageTemplate == nil {
		var ret SoftwareSolutionDistributableRelationship
		return ret
	}
	return *o.OvaImageTemplate
}

// GetOvaImageTemplateOk returns a tuple with the OvaImageTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVersion) GetOvaImageTemplateOk() (*SoftwareSolutionDistributableRelationship, bool) {
	if o == nil || o.OvaImageTemplate == nil {
		return nil, false
	}
	return o.OvaImageTemplate, true
}

// HasOvaImageTemplate returns a boolean if a field has been set.
func (o *KubernetesVersion) HasOvaImageTemplate() bool {
	if o != nil && o.OvaImageTemplate != nil {
		return true
	}

	return false
}

// SetOvaImageTemplate gets a reference to the given SoftwareSolutionDistributableRelationship and assigns it to the OvaImageTemplate field.
func (o *KubernetesVersion) SetOvaImageTemplate(v SoftwareSolutionDistributableRelationship) {
	o.OvaImageTemplate = &v
}

// GetQcow2NodeTemplate returns the Qcow2NodeTemplate field value if set, zero value otherwise.
func (o *KubernetesVersion) GetQcow2NodeTemplate() SoftwareSolutionDistributableRelationship {
	if o == nil || o.Qcow2NodeTemplate == nil {
		var ret SoftwareSolutionDistributableRelationship
		return ret
	}
	return *o.Qcow2NodeTemplate
}

// GetQcow2NodeTemplateOk returns a tuple with the Qcow2NodeTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVersion) GetQcow2NodeTemplateOk() (*SoftwareSolutionDistributableRelationship, bool) {
	if o == nil || o.Qcow2NodeTemplate == nil {
		return nil, false
	}
	return o.Qcow2NodeTemplate, true
}

// HasQcow2NodeTemplate returns a boolean if a field has been set.
func (o *KubernetesVersion) HasQcow2NodeTemplate() bool {
	if o != nil && o.Qcow2NodeTemplate != nil {
		return true
	}

	return false
}

// SetQcow2NodeTemplate gets a reference to the given SoftwareSolutionDistributableRelationship and assigns it to the Qcow2NodeTemplate field.
func (o *KubernetesVersion) SetQcow2NodeTemplate(v SoftwareSolutionDistributableRelationship) {
	o.Qcow2NodeTemplate = &v
}

// GetQcow2VirtualMachineTemplate returns the Qcow2VirtualMachineTemplate field value if set, zero value otherwise.
func (o *KubernetesVersion) GetQcow2VirtualMachineTemplate() SoftwareSolutionDistributableRelationship {
	if o == nil || o.Qcow2VirtualMachineTemplate == nil {
		var ret SoftwareSolutionDistributableRelationship
		return ret
	}
	return *o.Qcow2VirtualMachineTemplate
}

// GetQcow2VirtualMachineTemplateOk returns a tuple with the Qcow2VirtualMachineTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVersion) GetQcow2VirtualMachineTemplateOk() (*SoftwareSolutionDistributableRelationship, bool) {
	if o == nil || o.Qcow2VirtualMachineTemplate == nil {
		return nil, false
	}
	return o.Qcow2VirtualMachineTemplate, true
}

// HasQcow2VirtualMachineTemplate returns a boolean if a field has been set.
func (o *KubernetesVersion) HasQcow2VirtualMachineTemplate() bool {
	if o != nil && o.Qcow2VirtualMachineTemplate != nil {
		return true
	}

	return false
}

// SetQcow2VirtualMachineTemplate gets a reference to the given SoftwareSolutionDistributableRelationship and assigns it to the Qcow2VirtualMachineTemplate field.
func (o *KubernetesVersion) SetQcow2VirtualMachineTemplate(v SoftwareSolutionDistributableRelationship) {
	o.Qcow2VirtualMachineTemplate = &v
}

func (o KubernetesVersion) MarshalJSON() ([]byte, error) {
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
	if o.KubernetesVersion != nil {
		toSerialize["KubernetesVersion"] = o.KubernetesVersion
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.BootIso != nil {
		toSerialize["BootIso"] = o.BootIso
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}
	if o.OvaImageTemplate != nil {
		toSerialize["OvaImageTemplate"] = o.OvaImageTemplate
	}
	if o.Qcow2NodeTemplate != nil {
		toSerialize["Qcow2NodeTemplate"] = o.Qcow2NodeTemplate
	}
	if o.Qcow2VirtualMachineTemplate != nil {
		toSerialize["Qcow2VirtualMachineTemplate"] = o.Qcow2VirtualMachineTemplate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesVersion) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesVersionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Desired Kubernetes version.
		KubernetesVersion *string `json:"KubernetesVersion,omitempty"`
		// The name of this IKS kubernetes version.
		Name                        *string                                    `json:"Name,omitempty"`
		BootIso                     *SoftwareSolutionDistributableRelationship `json:"BootIso,omitempty"`
		Catalog                     *KubernetesCatalogRelationship             `json:"Catalog,omitempty"`
		OvaImageTemplate            *SoftwareSolutionDistributableRelationship `json:"OvaImageTemplate,omitempty"`
		Qcow2NodeTemplate           *SoftwareSolutionDistributableRelationship `json:"Qcow2NodeTemplate,omitempty"`
		Qcow2VirtualMachineTemplate *SoftwareSolutionDistributableRelationship `json:"Qcow2VirtualMachineTemplate,omitempty"`
	}

	varKubernetesVersionWithoutEmbeddedStruct := KubernetesVersionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesVersionWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesVersion := _KubernetesVersion{}
		varKubernetesVersion.ClassId = varKubernetesVersionWithoutEmbeddedStruct.ClassId
		varKubernetesVersion.ObjectType = varKubernetesVersionWithoutEmbeddedStruct.ObjectType
		varKubernetesVersion.KubernetesVersion = varKubernetesVersionWithoutEmbeddedStruct.KubernetesVersion
		varKubernetesVersion.Name = varKubernetesVersionWithoutEmbeddedStruct.Name
		varKubernetesVersion.BootIso = varKubernetesVersionWithoutEmbeddedStruct.BootIso
		varKubernetesVersion.Catalog = varKubernetesVersionWithoutEmbeddedStruct.Catalog
		varKubernetesVersion.OvaImageTemplate = varKubernetesVersionWithoutEmbeddedStruct.OvaImageTemplate
		varKubernetesVersion.Qcow2NodeTemplate = varKubernetesVersionWithoutEmbeddedStruct.Qcow2NodeTemplate
		varKubernetesVersion.Qcow2VirtualMachineTemplate = varKubernetesVersionWithoutEmbeddedStruct.Qcow2VirtualMachineTemplate
		*o = KubernetesVersion(varKubernetesVersion)
	} else {
		return err
	}

	varKubernetesVersion := _KubernetesVersion{}

	err = json.Unmarshal(bytes, &varKubernetesVersion)
	if err == nil {
		o.MoBaseMo = varKubernetesVersion.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "KubernetesVersion")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "BootIso")
		delete(additionalProperties, "Catalog")
		delete(additionalProperties, "OvaImageTemplate")
		delete(additionalProperties, "Qcow2NodeTemplate")
		delete(additionalProperties, "Qcow2VirtualMachineTemplate")

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

type NullableKubernetesVersion struct {
	value *KubernetesVersion
	isSet bool
}

func (v NullableKubernetesVersion) Get() *KubernetesVersion {
	return v.value
}

func (v *NullableKubernetesVersion) Set(val *KubernetesVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesVersion(val *KubernetesVersion) *NullableKubernetesVersion {
	return &NullableKubernetesVersion{value: val, isSet: true}
}

func (v NullableKubernetesVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
