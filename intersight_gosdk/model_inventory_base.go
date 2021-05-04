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

// InventoryBase Base class for all inventory MOs.
type InventoryBase struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The database identifier of the registered device of an object.
	DeviceMoId *string `json:"DeviceMoId,omitempty"`
	// The Distinguished Name unambiguously identifies an object in the system.
	Dn *string `json:"Dn,omitempty"`
	// The Relative Name uniquely identifies an object within a given context.
	Rn                   *string `json:"Rn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryBase InventoryBase

// NewInventoryBase instantiates a new InventoryBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryBase(classId string, objectType string) *InventoryBase {
	this := InventoryBase{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewInventoryBaseWithDefaults instantiates a new InventoryBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryBaseWithDefaults() *InventoryBase {
	this := InventoryBase{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *InventoryBase) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *InventoryBase) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *InventoryBase) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *InventoryBase) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *InventoryBase) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *InventoryBase) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeviceMoId returns the DeviceMoId field value if set, zero value otherwise.
func (o *InventoryBase) GetDeviceMoId() string {
	if o == nil || o.DeviceMoId == nil {
		var ret string
		return ret
	}
	return *o.DeviceMoId
}

// GetDeviceMoIdOk returns a tuple with the DeviceMoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryBase) GetDeviceMoIdOk() (*string, bool) {
	if o == nil || o.DeviceMoId == nil {
		return nil, false
	}
	return o.DeviceMoId, true
}

// HasDeviceMoId returns a boolean if a field has been set.
func (o *InventoryBase) HasDeviceMoId() bool {
	if o != nil && o.DeviceMoId != nil {
		return true
	}

	return false
}

// SetDeviceMoId gets a reference to the given string and assigns it to the DeviceMoId field.
func (o *InventoryBase) SetDeviceMoId(v string) {
	o.DeviceMoId = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *InventoryBase) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryBase) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *InventoryBase) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *InventoryBase) SetDn(v string) {
	o.Dn = &v
}

// GetRn returns the Rn field value if set, zero value otherwise.
func (o *InventoryBase) GetRn() string {
	if o == nil || o.Rn == nil {
		var ret string
		return ret
	}
	return *o.Rn
}

// GetRnOk returns a tuple with the Rn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryBase) GetRnOk() (*string, bool) {
	if o == nil || o.Rn == nil {
		return nil, false
	}
	return o.Rn, true
}

// HasRn returns a boolean if a field has been set.
func (o *InventoryBase) HasRn() bool {
	if o != nil && o.Rn != nil {
		return true
	}

	return false
}

// SetRn gets a reference to the given string and assigns it to the Rn field.
func (o *InventoryBase) SetRn(v string) {
	o.Rn = &v
}

func (o InventoryBase) MarshalJSON() ([]byte, error) {
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
	if o.DeviceMoId != nil {
		toSerialize["DeviceMoId"] = o.DeviceMoId
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.Rn != nil {
		toSerialize["Rn"] = o.Rn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InventoryBase) UnmarshalJSON(bytes []byte) (err error) {
	type InventoryBaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The database identifier of the registered device of an object.
		DeviceMoId *string `json:"DeviceMoId,omitempty"`
		// The Distinguished Name unambiguously identifies an object in the system.
		Dn *string `json:"Dn,omitempty"`
		// The Relative Name uniquely identifies an object within a given context.
		Rn *string `json:"Rn,omitempty"`
	}

	varInventoryBaseWithoutEmbeddedStruct := InventoryBaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varInventoryBaseWithoutEmbeddedStruct)
	if err == nil {
		varInventoryBase := _InventoryBase{}
		varInventoryBase.ClassId = varInventoryBaseWithoutEmbeddedStruct.ClassId
		varInventoryBase.ObjectType = varInventoryBaseWithoutEmbeddedStruct.ObjectType
		varInventoryBase.DeviceMoId = varInventoryBaseWithoutEmbeddedStruct.DeviceMoId
		varInventoryBase.Dn = varInventoryBaseWithoutEmbeddedStruct.Dn
		varInventoryBase.Rn = varInventoryBaseWithoutEmbeddedStruct.Rn
		*o = InventoryBase(varInventoryBase)
	} else {
		return err
	}

	varInventoryBase := _InventoryBase{}

	err = json.Unmarshal(bytes, &varInventoryBase)
	if err == nil {
		o.MoBaseMo = varInventoryBase.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceMoId")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "Rn")

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

type NullableInventoryBase struct {
	value *InventoryBase
	isSet bool
}

func (v NullableInventoryBase) Get() *InventoryBase {
	return v.value
}

func (v *NullableInventoryBase) Set(val *InventoryBase) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryBase) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryBase(val *InventoryBase) *NullableInventoryBase {
	return &NullableInventoryBase{value: val, isSet: true}
}

func (v NullableInventoryBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
