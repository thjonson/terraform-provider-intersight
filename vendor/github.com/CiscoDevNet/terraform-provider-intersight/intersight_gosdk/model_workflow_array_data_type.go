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

// WorkflowArrayDataType This data type represents an array of a given type. It can be an array of primitive data or of custom data.
type WorkflowArrayDataType struct {
	WorkflowBaseDataType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType    string                    `json:"ObjectType"`
	ArrayItemType NullableWorkflowArrayItem `json:"ArrayItemType,omitempty"`
	// Specify the maximum value of the array.
	Max *int64 `json:"Max,omitempty"`
	// Specify the minimum value of the array.
	Min                  *int64 `json:"Min,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowArrayDataType WorkflowArrayDataType

// NewWorkflowArrayDataType instantiates a new WorkflowArrayDataType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowArrayDataType(classId string, objectType string) *WorkflowArrayDataType {
	this := WorkflowArrayDataType{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowArrayDataTypeWithDefaults instantiates a new WorkflowArrayDataType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowArrayDataTypeWithDefaults() *WorkflowArrayDataType {
	this := WorkflowArrayDataType{}
	var classId string = "workflow.ArrayDataType"
	this.ClassId = classId
	var objectType string = "workflow.ArrayDataType"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowArrayDataType) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowArrayDataType) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowArrayDataType) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowArrayDataType) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowArrayDataType) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowArrayDataType) SetObjectType(v string) {
	o.ObjectType = v
}

// GetArrayItemType returns the ArrayItemType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowArrayDataType) GetArrayItemType() WorkflowArrayItem {
	if o == nil || o.ArrayItemType.Get() == nil {
		var ret WorkflowArrayItem
		return ret
	}
	return *o.ArrayItemType.Get()
}

// GetArrayItemTypeOk returns a tuple with the ArrayItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowArrayDataType) GetArrayItemTypeOk() (*WorkflowArrayItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArrayItemType.Get(), o.ArrayItemType.IsSet()
}

// HasArrayItemType returns a boolean if a field has been set.
func (o *WorkflowArrayDataType) HasArrayItemType() bool {
	if o != nil && o.ArrayItemType.IsSet() {
		return true
	}

	return false
}

// SetArrayItemType gets a reference to the given NullableWorkflowArrayItem and assigns it to the ArrayItemType field.
func (o *WorkflowArrayDataType) SetArrayItemType(v WorkflowArrayItem) {
	o.ArrayItemType.Set(&v)
}

// SetArrayItemTypeNil sets the value for ArrayItemType to be an explicit nil
func (o *WorkflowArrayDataType) SetArrayItemTypeNil() {
	o.ArrayItemType.Set(nil)
}

// UnsetArrayItemType ensures that no value is present for ArrayItemType, not even an explicit nil
func (o *WorkflowArrayDataType) UnsetArrayItemType() {
	o.ArrayItemType.Unset()
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *WorkflowArrayDataType) GetMax() int64 {
	if o == nil || o.Max == nil {
		var ret int64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowArrayDataType) GetMaxOk() (*int64, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *WorkflowArrayDataType) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given int64 and assigns it to the Max field.
func (o *WorkflowArrayDataType) SetMax(v int64) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *WorkflowArrayDataType) GetMin() int64 {
	if o == nil || o.Min == nil {
		var ret int64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowArrayDataType) GetMinOk() (*int64, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *WorkflowArrayDataType) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given int64 and assigns it to the Min field.
func (o *WorkflowArrayDataType) SetMin(v int64) {
	o.Min = &v
}

func (o WorkflowArrayDataType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowBaseDataType, errWorkflowBaseDataType := json.Marshal(o.WorkflowBaseDataType)
	if errWorkflowBaseDataType != nil {
		return []byte{}, errWorkflowBaseDataType
	}
	errWorkflowBaseDataType = json.Unmarshal([]byte(serializedWorkflowBaseDataType), &toSerialize)
	if errWorkflowBaseDataType != nil {
		return []byte{}, errWorkflowBaseDataType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ArrayItemType.IsSet() {
		toSerialize["ArrayItemType"] = o.ArrayItemType.Get()
	}
	if o.Max != nil {
		toSerialize["Max"] = o.Max
	}
	if o.Min != nil {
		toSerialize["Min"] = o.Min
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowArrayDataType) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowArrayDataTypeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType    string                    `json:"ObjectType"`
		ArrayItemType NullableWorkflowArrayItem `json:"ArrayItemType,omitempty"`
		// Specify the maximum value of the array.
		Max *int64 `json:"Max,omitempty"`
		// Specify the minimum value of the array.
		Min *int64 `json:"Min,omitempty"`
	}

	varWorkflowArrayDataTypeWithoutEmbeddedStruct := WorkflowArrayDataTypeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowArrayDataTypeWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowArrayDataType := _WorkflowArrayDataType{}
		varWorkflowArrayDataType.ClassId = varWorkflowArrayDataTypeWithoutEmbeddedStruct.ClassId
		varWorkflowArrayDataType.ObjectType = varWorkflowArrayDataTypeWithoutEmbeddedStruct.ObjectType
		varWorkflowArrayDataType.ArrayItemType = varWorkflowArrayDataTypeWithoutEmbeddedStruct.ArrayItemType
		varWorkflowArrayDataType.Max = varWorkflowArrayDataTypeWithoutEmbeddedStruct.Max
		varWorkflowArrayDataType.Min = varWorkflowArrayDataTypeWithoutEmbeddedStruct.Min
		*o = WorkflowArrayDataType(varWorkflowArrayDataType)
	} else {
		return err
	}

	varWorkflowArrayDataType := _WorkflowArrayDataType{}

	err = json.Unmarshal(bytes, &varWorkflowArrayDataType)
	if err == nil {
		o.WorkflowBaseDataType = varWorkflowArrayDataType.WorkflowBaseDataType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ArrayItemType")
		delete(additionalProperties, "Max")
		delete(additionalProperties, "Min")

		// remove fields from embedded structs
		reflectWorkflowBaseDataType := reflect.ValueOf(o.WorkflowBaseDataType)
		for i := 0; i < reflectWorkflowBaseDataType.Type().NumField(); i++ {
			t := reflectWorkflowBaseDataType.Type().Field(i)

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

type NullableWorkflowArrayDataType struct {
	value *WorkflowArrayDataType
	isSet bool
}

func (v NullableWorkflowArrayDataType) Get() *WorkflowArrayDataType {
	return v.value
}

func (v *NullableWorkflowArrayDataType) Set(val *WorkflowArrayDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowArrayDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowArrayDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowArrayDataType(val *WorkflowArrayDataType) *NullableWorkflowArrayDataType {
	return &NullableWorkflowArrayDataType{value: val, isSet: true}
}

func (v NullableWorkflowArrayDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowArrayDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
