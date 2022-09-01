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

// ConvergedinfraStorageComplianceDetailsAllOf Definition of the list of properties defined in 'convergedinfra.StorageComplianceDetails', excluding properties defined in parent classes.
type ConvergedinfraStorageComplianceDetailsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The operating system name and version (e.g. NetApp ONTAP 9.10) running on the storage array for which the compliance is getting evaluated.
	Os *string `json:"Os,omitempty"`
	// The protocol configured for the communication between the switch and the storage array.
	Protocol *string `json:"Protocol,omitempty"`
	// The reference device (e.g. adapter, fabric interconnect) against which the storage compliance is getting evaluated. * `Server` - The component type for a server in a converged infrastructure pod. * `Adapter` - The component type for an adapter on a server in a converged infrastructure pod. * `FabricInterconnect` - The component type for a fabric interconnect in a converged infrastructure pod. * `Nexus` - The component type for a nexus switch in a converged infrastructure pod. * `Storage` - The component type for a storage array in a converged infrastructure pod.
	RefDevice            *string                                             `json:"RefDevice,omitempty"`
	AdapterCompliance    *ConvergedinfraAdapterComplianceDetailsRelationship `json:"AdapterCompliance,omitempty"`
	PodCompliance        *ConvergedinfraPodComplianceInfoRelationship        `json:"PodCompliance,omitempty"`
	StorageArray         *StorageBaseArrayRelationship                       `json:"StorageArray,omitempty"`
	SwitchCompliance     *ConvergedinfraSwitchComplianceDetailsRelationship  `json:"SwitchCompliance,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConvergedinfraStorageComplianceDetailsAllOf ConvergedinfraStorageComplianceDetailsAllOf

// NewConvergedinfraStorageComplianceDetailsAllOf instantiates a new ConvergedinfraStorageComplianceDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvergedinfraStorageComplianceDetailsAllOf(classId string, objectType string) *ConvergedinfraStorageComplianceDetailsAllOf {
	this := ConvergedinfraStorageComplianceDetailsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConvergedinfraStorageComplianceDetailsAllOfWithDefaults instantiates a new ConvergedinfraStorageComplianceDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvergedinfraStorageComplianceDetailsAllOfWithDefaults() *ConvergedinfraStorageComplianceDetailsAllOf {
	this := ConvergedinfraStorageComplianceDetailsAllOf{}
	var classId string = "convergedinfra.StorageComplianceDetails"
	this.ClassId = classId
	var objectType string = "convergedinfra.StorageComplianceDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConvergedinfraStorageComplianceDetailsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConvergedinfraStorageComplianceDetailsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConvergedinfraStorageComplianceDetailsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConvergedinfraStorageComplianceDetailsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) GetOs() string {
	if o == nil || o.Os == nil {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) GetOsOk() (*string, bool) {
	if o == nil || o.Os == nil {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) HasOs() bool {
	if o != nil && o.Os != nil {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) SetOs(v string) {
	o.Os = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) SetProtocol(v string) {
	o.Protocol = &v
}

// GetRefDevice returns the RefDevice field value if set, zero value otherwise.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) GetRefDevice() string {
	if o == nil || o.RefDevice == nil {
		var ret string
		return ret
	}
	return *o.RefDevice
}

// GetRefDeviceOk returns a tuple with the RefDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) GetRefDeviceOk() (*string, bool) {
	if o == nil || o.RefDevice == nil {
		return nil, false
	}
	return o.RefDevice, true
}

// HasRefDevice returns a boolean if a field has been set.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) HasRefDevice() bool {
	if o != nil && o.RefDevice != nil {
		return true
	}

	return false
}

// SetRefDevice gets a reference to the given string and assigns it to the RefDevice field.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) SetRefDevice(v string) {
	o.RefDevice = &v
}

// GetAdapterCompliance returns the AdapterCompliance field value if set, zero value otherwise.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) GetAdapterCompliance() ConvergedinfraAdapterComplianceDetailsRelationship {
	if o == nil || o.AdapterCompliance == nil {
		var ret ConvergedinfraAdapterComplianceDetailsRelationship
		return ret
	}
	return *o.AdapterCompliance
}

// GetAdapterComplianceOk returns a tuple with the AdapterCompliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) GetAdapterComplianceOk() (*ConvergedinfraAdapterComplianceDetailsRelationship, bool) {
	if o == nil || o.AdapterCompliance == nil {
		return nil, false
	}
	return o.AdapterCompliance, true
}

// HasAdapterCompliance returns a boolean if a field has been set.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) HasAdapterCompliance() bool {
	if o != nil && o.AdapterCompliance != nil {
		return true
	}

	return false
}

// SetAdapterCompliance gets a reference to the given ConvergedinfraAdapterComplianceDetailsRelationship and assigns it to the AdapterCompliance field.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) SetAdapterCompliance(v ConvergedinfraAdapterComplianceDetailsRelationship) {
	o.AdapterCompliance = &v
}

// GetPodCompliance returns the PodCompliance field value if set, zero value otherwise.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) GetPodCompliance() ConvergedinfraPodComplianceInfoRelationship {
	if o == nil || o.PodCompliance == nil {
		var ret ConvergedinfraPodComplianceInfoRelationship
		return ret
	}
	return *o.PodCompliance
}

// GetPodComplianceOk returns a tuple with the PodCompliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) GetPodComplianceOk() (*ConvergedinfraPodComplianceInfoRelationship, bool) {
	if o == nil || o.PodCompliance == nil {
		return nil, false
	}
	return o.PodCompliance, true
}

// HasPodCompliance returns a boolean if a field has been set.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) HasPodCompliance() bool {
	if o != nil && o.PodCompliance != nil {
		return true
	}

	return false
}

// SetPodCompliance gets a reference to the given ConvergedinfraPodComplianceInfoRelationship and assigns it to the PodCompliance field.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) SetPodCompliance(v ConvergedinfraPodComplianceInfoRelationship) {
	o.PodCompliance = &v
}

// GetStorageArray returns the StorageArray field value if set, zero value otherwise.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) GetStorageArray() StorageBaseArrayRelationship {
	if o == nil || o.StorageArray == nil {
		var ret StorageBaseArrayRelationship
		return ret
	}
	return *o.StorageArray
}

// GetStorageArrayOk returns a tuple with the StorageArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) GetStorageArrayOk() (*StorageBaseArrayRelationship, bool) {
	if o == nil || o.StorageArray == nil {
		return nil, false
	}
	return o.StorageArray, true
}

// HasStorageArray returns a boolean if a field has been set.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) HasStorageArray() bool {
	if o != nil && o.StorageArray != nil {
		return true
	}

	return false
}

// SetStorageArray gets a reference to the given StorageBaseArrayRelationship and assigns it to the StorageArray field.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) SetStorageArray(v StorageBaseArrayRelationship) {
	o.StorageArray = &v
}

// GetSwitchCompliance returns the SwitchCompliance field value if set, zero value otherwise.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) GetSwitchCompliance() ConvergedinfraSwitchComplianceDetailsRelationship {
	if o == nil || o.SwitchCompliance == nil {
		var ret ConvergedinfraSwitchComplianceDetailsRelationship
		return ret
	}
	return *o.SwitchCompliance
}

// GetSwitchComplianceOk returns a tuple with the SwitchCompliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) GetSwitchComplianceOk() (*ConvergedinfraSwitchComplianceDetailsRelationship, bool) {
	if o == nil || o.SwitchCompliance == nil {
		return nil, false
	}
	return o.SwitchCompliance, true
}

// HasSwitchCompliance returns a boolean if a field has been set.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) HasSwitchCompliance() bool {
	if o != nil && o.SwitchCompliance != nil {
		return true
	}

	return false
}

// SetSwitchCompliance gets a reference to the given ConvergedinfraSwitchComplianceDetailsRelationship and assigns it to the SwitchCompliance field.
func (o *ConvergedinfraStorageComplianceDetailsAllOf) SetSwitchCompliance(v ConvergedinfraSwitchComplianceDetailsRelationship) {
	o.SwitchCompliance = &v
}

func (o ConvergedinfraStorageComplianceDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Os != nil {
		toSerialize["Os"] = o.Os
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.RefDevice != nil {
		toSerialize["RefDevice"] = o.RefDevice
	}
	if o.AdapterCompliance != nil {
		toSerialize["AdapterCompliance"] = o.AdapterCompliance
	}
	if o.PodCompliance != nil {
		toSerialize["PodCompliance"] = o.PodCompliance
	}
	if o.StorageArray != nil {
		toSerialize["StorageArray"] = o.StorageArray
	}
	if o.SwitchCompliance != nil {
		toSerialize["SwitchCompliance"] = o.SwitchCompliance
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConvergedinfraStorageComplianceDetailsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConvergedinfraStorageComplianceDetailsAllOf := _ConvergedinfraStorageComplianceDetailsAllOf{}

	if err = json.Unmarshal(bytes, &varConvergedinfraStorageComplianceDetailsAllOf); err == nil {
		*o = ConvergedinfraStorageComplianceDetailsAllOf(varConvergedinfraStorageComplianceDetailsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Os")
		delete(additionalProperties, "Protocol")
		delete(additionalProperties, "RefDevice")
		delete(additionalProperties, "AdapterCompliance")
		delete(additionalProperties, "PodCompliance")
		delete(additionalProperties, "StorageArray")
		delete(additionalProperties, "SwitchCompliance")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConvergedinfraStorageComplianceDetailsAllOf struct {
	value *ConvergedinfraStorageComplianceDetailsAllOf
	isSet bool
}

func (v NullableConvergedinfraStorageComplianceDetailsAllOf) Get() *ConvergedinfraStorageComplianceDetailsAllOf {
	return v.value
}

func (v *NullableConvergedinfraStorageComplianceDetailsAllOf) Set(val *ConvergedinfraStorageComplianceDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConvergedinfraStorageComplianceDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConvergedinfraStorageComplianceDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvergedinfraStorageComplianceDetailsAllOf(val *ConvergedinfraStorageComplianceDetailsAllOf) *NullableConvergedinfraStorageComplianceDetailsAllOf {
	return &NullableConvergedinfraStorageComplianceDetailsAllOf{value: val, isSet: true}
}

func (v NullableConvergedinfraStorageComplianceDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvergedinfraStorageComplianceDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
