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

// KubernetesAciCniApicAllOf Definition of the list of properties defined in 'kubernetes.AciCniApic', excluding properties defined in parent classes.
type KubernetesAciCniApicAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The Moid of the apic device under the asset service.
	AssetApicMoid *string `json:"AssetApicMoid,omitempty"`
	// The number of ACI CNI profiles configured for this APIC.
	NumAciCniProfiles *int64 `json:"NumAciCniProfiles,omitempty"`
	// An array of relationships to kubernetesAciCniProfile resources.
	AciCniProfiles       []KubernetesAciCniProfileRelationship `json:"AciCniProfiles,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship  `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesAciCniApicAllOf KubernetesAciCniApicAllOf

// NewKubernetesAciCniApicAllOf instantiates a new KubernetesAciCniApicAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAciCniApicAllOf(classId string, objectType string) *KubernetesAciCniApicAllOf {
	this := KubernetesAciCniApicAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesAciCniApicAllOfWithDefaults instantiates a new KubernetesAciCniApicAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAciCniApicAllOfWithDefaults() *KubernetesAciCniApicAllOf {
	this := KubernetesAciCniApicAllOf{}
	var classId string = "kubernetes.AciCniApic"
	this.ClassId = classId
	var objectType string = "kubernetes.AciCniApic"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesAciCniApicAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniApicAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesAciCniApicAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesAciCniApicAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniApicAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesAciCniApicAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAssetApicMoid returns the AssetApicMoid field value if set, zero value otherwise.
func (o *KubernetesAciCniApicAllOf) GetAssetApicMoid() string {
	if o == nil || o.AssetApicMoid == nil {
		var ret string
		return ret
	}
	return *o.AssetApicMoid
}

// GetAssetApicMoidOk returns a tuple with the AssetApicMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniApicAllOf) GetAssetApicMoidOk() (*string, bool) {
	if o == nil || o.AssetApicMoid == nil {
		return nil, false
	}
	return o.AssetApicMoid, true
}

// HasAssetApicMoid returns a boolean if a field has been set.
func (o *KubernetesAciCniApicAllOf) HasAssetApicMoid() bool {
	if o != nil && o.AssetApicMoid != nil {
		return true
	}

	return false
}

// SetAssetApicMoid gets a reference to the given string and assigns it to the AssetApicMoid field.
func (o *KubernetesAciCniApicAllOf) SetAssetApicMoid(v string) {
	o.AssetApicMoid = &v
}

// GetNumAciCniProfiles returns the NumAciCniProfiles field value if set, zero value otherwise.
func (o *KubernetesAciCniApicAllOf) GetNumAciCniProfiles() int64 {
	if o == nil || o.NumAciCniProfiles == nil {
		var ret int64
		return ret
	}
	return *o.NumAciCniProfiles
}

// GetNumAciCniProfilesOk returns a tuple with the NumAciCniProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniApicAllOf) GetNumAciCniProfilesOk() (*int64, bool) {
	if o == nil || o.NumAciCniProfiles == nil {
		return nil, false
	}
	return o.NumAciCniProfiles, true
}

// HasNumAciCniProfiles returns a boolean if a field has been set.
func (o *KubernetesAciCniApicAllOf) HasNumAciCniProfiles() bool {
	if o != nil && o.NumAciCniProfiles != nil {
		return true
	}

	return false
}

// SetNumAciCniProfiles gets a reference to the given int64 and assigns it to the NumAciCniProfiles field.
func (o *KubernetesAciCniApicAllOf) SetNumAciCniProfiles(v int64) {
	o.NumAciCniProfiles = &v
}

// GetAciCniProfiles returns the AciCniProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAciCniApicAllOf) GetAciCniProfiles() []KubernetesAciCniProfileRelationship {
	if o == nil {
		var ret []KubernetesAciCniProfileRelationship
		return ret
	}
	return o.AciCniProfiles
}

// GetAciCniProfilesOk returns a tuple with the AciCniProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAciCniApicAllOf) GetAciCniProfilesOk() ([]KubernetesAciCniProfileRelationship, bool) {
	if o == nil || o.AciCniProfiles == nil {
		return nil, false
	}
	return o.AciCniProfiles, true
}

// HasAciCniProfiles returns a boolean if a field has been set.
func (o *KubernetesAciCniApicAllOf) HasAciCniProfiles() bool {
	if o != nil && o.AciCniProfiles != nil {
		return true
	}

	return false
}

// SetAciCniProfiles gets a reference to the given []KubernetesAciCniProfileRelationship and assigns it to the AciCniProfiles field.
func (o *KubernetesAciCniApicAllOf) SetAciCniProfiles(v []KubernetesAciCniProfileRelationship) {
	o.AciCniProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesAciCniApicAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniApicAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesAciCniApicAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesAciCniApicAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *KubernetesAciCniApicAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniApicAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *KubernetesAciCniApicAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *KubernetesAciCniApicAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o KubernetesAciCniApicAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AssetApicMoid != nil {
		toSerialize["AssetApicMoid"] = o.AssetApicMoid
	}
	if o.NumAciCniProfiles != nil {
		toSerialize["NumAciCniProfiles"] = o.NumAciCniProfiles
	}
	if o.AciCniProfiles != nil {
		toSerialize["AciCniProfiles"] = o.AciCniProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesAciCniApicAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesAciCniApicAllOf := _KubernetesAciCniApicAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesAciCniApicAllOf); err == nil {
		*o = KubernetesAciCniApicAllOf(varKubernetesAciCniApicAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AssetApicMoid")
		delete(additionalProperties, "NumAciCniProfiles")
		delete(additionalProperties, "AciCniProfiles")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesAciCniApicAllOf struct {
	value *KubernetesAciCniApicAllOf
	isSet bool
}

func (v NullableKubernetesAciCniApicAllOf) Get() *KubernetesAciCniApicAllOf {
	return v.value
}

func (v *NullableKubernetesAciCniApicAllOf) Set(val *KubernetesAciCniApicAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAciCniApicAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAciCniApicAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAciCniApicAllOf(val *KubernetesAciCniApicAllOf) *NullableKubernetesAciCniApicAllOf {
	return &NullableKubernetesAciCniApicAllOf{value: val, isSet: true}
}

func (v NullableKubernetesAciCniApicAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAciCniApicAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
