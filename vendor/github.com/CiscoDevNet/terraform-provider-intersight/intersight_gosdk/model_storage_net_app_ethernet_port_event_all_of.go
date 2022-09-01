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

// StorageNetAppEthernetPortEventAllOf Definition of the list of properties defined in 'storage.NetAppEthernetPortEvent', excluding properties defined in parent classes.
type StorageNetAppEthernetPortEventAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                 `json:"ObjectType"`
	EthernetPort         *StorageNetAppEthernetPortRelationship `json:"EthernetPort,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppEthernetPortEventAllOf StorageNetAppEthernetPortEventAllOf

// NewStorageNetAppEthernetPortEventAllOf instantiates a new StorageNetAppEthernetPortEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppEthernetPortEventAllOf(classId string, objectType string) *StorageNetAppEthernetPortEventAllOf {
	this := StorageNetAppEthernetPortEventAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppEthernetPortEventAllOfWithDefaults instantiates a new StorageNetAppEthernetPortEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppEthernetPortEventAllOfWithDefaults() *StorageNetAppEthernetPortEventAllOf {
	this := StorageNetAppEthernetPortEventAllOf{}
	var classId string = "storage.NetAppEthernetPortEvent"
	this.ClassId = classId
	var objectType string = "storage.NetAppEthernetPortEvent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppEthernetPortEventAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortEventAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppEthernetPortEventAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppEthernetPortEventAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortEventAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppEthernetPortEventAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEthernetPort returns the EthernetPort field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPortEventAllOf) GetEthernetPort() StorageNetAppEthernetPortRelationship {
	if o == nil || o.EthernetPort == nil {
		var ret StorageNetAppEthernetPortRelationship
		return ret
	}
	return *o.EthernetPort
}

// GetEthernetPortOk returns a tuple with the EthernetPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortEventAllOf) GetEthernetPortOk() (*StorageNetAppEthernetPortRelationship, bool) {
	if o == nil || o.EthernetPort == nil {
		return nil, false
	}
	return o.EthernetPort, true
}

// HasEthernetPort returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPortEventAllOf) HasEthernetPort() bool {
	if o != nil && o.EthernetPort != nil {
		return true
	}

	return false
}

// SetEthernetPort gets a reference to the given StorageNetAppEthernetPortRelationship and assigns it to the EthernetPort field.
func (o *StorageNetAppEthernetPortEventAllOf) SetEthernetPort(v StorageNetAppEthernetPortRelationship) {
	o.EthernetPort = &v
}

func (o StorageNetAppEthernetPortEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EthernetPort != nil {
		toSerialize["EthernetPort"] = o.EthernetPort
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppEthernetPortEventAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppEthernetPortEventAllOf := _StorageNetAppEthernetPortEventAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppEthernetPortEventAllOf); err == nil {
		*o = StorageNetAppEthernetPortEventAllOf(varStorageNetAppEthernetPortEventAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EthernetPort")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppEthernetPortEventAllOf struct {
	value *StorageNetAppEthernetPortEventAllOf
	isSet bool
}

func (v NullableStorageNetAppEthernetPortEventAllOf) Get() *StorageNetAppEthernetPortEventAllOf {
	return v.value
}

func (v *NullableStorageNetAppEthernetPortEventAllOf) Set(val *StorageNetAppEthernetPortEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppEthernetPortEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppEthernetPortEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppEthernetPortEventAllOf(val *StorageNetAppEthernetPortEventAllOf) *NullableStorageNetAppEthernetPortEventAllOf {
	return &NullableStorageNetAppEthernetPortEventAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppEthernetPortEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppEthernetPortEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
