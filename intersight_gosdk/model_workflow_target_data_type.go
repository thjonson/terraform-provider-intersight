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

// WorkflowTargetDataType Data type to capture a target endpoint or device.
type WorkflowTargetDataType struct {
	WorkflowBaseDataType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType               string                             `json:"ObjectType"`
	CustomDataTypeProperties NullableWorkflowCustomDataProperty `json:"CustomDataTypeProperties,omitempty"`
	// When this property is true then an array of targets can be passed as input.
	IsArray *bool `json:"IsArray,omitempty"`
	// Specify the maximum value of the array.
	Max *int64 `json:"Max,omitempty"`
	// Specify the minimum value of the array.
	Min                  *int64                   `json:"Min,omitempty"`
	Properties           []WorkflowTargetProperty `json:"Properties,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTargetDataType WorkflowTargetDataType

// NewWorkflowTargetDataType instantiates a new WorkflowTargetDataType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTargetDataType(classId string, objectType string) *WorkflowTargetDataType {
	this := WorkflowTargetDataType{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isArray bool = false
	this.IsArray = &isArray
	return &this
}

// NewWorkflowTargetDataTypeWithDefaults instantiates a new WorkflowTargetDataType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTargetDataTypeWithDefaults() *WorkflowTargetDataType {
	this := WorkflowTargetDataType{}
	var classId string = "workflow.TargetDataType"
	this.ClassId = classId
	var objectType string = "workflow.TargetDataType"
	this.ObjectType = objectType
	var isArray bool = false
	this.IsArray = &isArray
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowTargetDataType) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowTargetDataType) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowTargetDataType) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowTargetDataType) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowTargetDataType) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowTargetDataType) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCustomDataTypeProperties returns the CustomDataTypeProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTargetDataType) GetCustomDataTypeProperties() WorkflowCustomDataProperty {
	if o == nil || o.CustomDataTypeProperties.Get() == nil {
		var ret WorkflowCustomDataProperty
		return ret
	}
	return *o.CustomDataTypeProperties.Get()
}

// GetCustomDataTypePropertiesOk returns a tuple with the CustomDataTypeProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTargetDataType) GetCustomDataTypePropertiesOk() (*WorkflowCustomDataProperty, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomDataTypeProperties.Get(), o.CustomDataTypeProperties.IsSet()
}

// HasCustomDataTypeProperties returns a boolean if a field has been set.
func (o *WorkflowTargetDataType) HasCustomDataTypeProperties() bool {
	if o != nil && o.CustomDataTypeProperties.IsSet() {
		return true
	}

	return false
}

// SetCustomDataTypeProperties gets a reference to the given NullableWorkflowCustomDataProperty and assigns it to the CustomDataTypeProperties field.
func (o *WorkflowTargetDataType) SetCustomDataTypeProperties(v WorkflowCustomDataProperty) {
	o.CustomDataTypeProperties.Set(&v)
}

// SetCustomDataTypePropertiesNil sets the value for CustomDataTypeProperties to be an explicit nil
func (o *WorkflowTargetDataType) SetCustomDataTypePropertiesNil() {
	o.CustomDataTypeProperties.Set(nil)
}

// UnsetCustomDataTypeProperties ensures that no value is present for CustomDataTypeProperties, not even an explicit nil
func (o *WorkflowTargetDataType) UnsetCustomDataTypeProperties() {
	o.CustomDataTypeProperties.Unset()
}

// GetIsArray returns the IsArray field value if set, zero value otherwise.
func (o *WorkflowTargetDataType) GetIsArray() bool {
	if o == nil || o.IsArray == nil {
		var ret bool
		return ret
	}
	return *o.IsArray
}

// GetIsArrayOk returns a tuple with the IsArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetDataType) GetIsArrayOk() (*bool, bool) {
	if o == nil || o.IsArray == nil {
		return nil, false
	}
	return o.IsArray, true
}

// HasIsArray returns a boolean if a field has been set.
func (o *WorkflowTargetDataType) HasIsArray() bool {
	if o != nil && o.IsArray != nil {
		return true
	}

	return false
}

// SetIsArray gets a reference to the given bool and assigns it to the IsArray field.
func (o *WorkflowTargetDataType) SetIsArray(v bool) {
	o.IsArray = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *WorkflowTargetDataType) GetMax() int64 {
	if o == nil || o.Max == nil {
		var ret int64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetDataType) GetMaxOk() (*int64, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *WorkflowTargetDataType) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given int64 and assigns it to the Max field.
func (o *WorkflowTargetDataType) SetMax(v int64) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *WorkflowTargetDataType) GetMin() int64 {
	if o == nil || o.Min == nil {
		var ret int64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetDataType) GetMinOk() (*int64, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *WorkflowTargetDataType) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given int64 and assigns it to the Min field.
func (o *WorkflowTargetDataType) SetMin(v int64) {
	o.Min = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTargetDataType) GetProperties() []WorkflowTargetProperty {
	if o == nil {
		var ret []WorkflowTargetProperty
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTargetDataType) GetPropertiesOk() (*[]WorkflowTargetProperty, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *WorkflowTargetDataType) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []WorkflowTargetProperty and assigns it to the Properties field.
func (o *WorkflowTargetDataType) SetProperties(v []WorkflowTargetProperty) {
	o.Properties = v
}

func (o WorkflowTargetDataType) MarshalJSON() ([]byte, error) {
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
	if o.CustomDataTypeProperties.IsSet() {
		toSerialize["CustomDataTypeProperties"] = o.CustomDataTypeProperties.Get()
	}
	if o.IsArray != nil {
		toSerialize["IsArray"] = o.IsArray
	}
	if o.Max != nil {
		toSerialize["Max"] = o.Max
	}
	if o.Min != nil {
		toSerialize["Min"] = o.Min
	}
	if o.Properties != nil {
		toSerialize["Properties"] = o.Properties
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowTargetDataType) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowTargetDataTypeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType               string                             `json:"ObjectType"`
		CustomDataTypeProperties NullableWorkflowCustomDataProperty `json:"CustomDataTypeProperties,omitempty"`
		// When this property is true then an array of targets can be passed as input.
		IsArray *bool `json:"IsArray,omitempty"`
		// Specify the maximum value of the array.
		Max *int64 `json:"Max,omitempty"`
		// Specify the minimum value of the array.
		Min        *int64                   `json:"Min,omitempty"`
		Properties []WorkflowTargetProperty `json:"Properties,omitempty"`
	}

	varWorkflowTargetDataTypeWithoutEmbeddedStruct := WorkflowTargetDataTypeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowTargetDataTypeWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowTargetDataType := _WorkflowTargetDataType{}
		varWorkflowTargetDataType.ClassId = varWorkflowTargetDataTypeWithoutEmbeddedStruct.ClassId
		varWorkflowTargetDataType.ObjectType = varWorkflowTargetDataTypeWithoutEmbeddedStruct.ObjectType
		varWorkflowTargetDataType.CustomDataTypeProperties = varWorkflowTargetDataTypeWithoutEmbeddedStruct.CustomDataTypeProperties
		varWorkflowTargetDataType.IsArray = varWorkflowTargetDataTypeWithoutEmbeddedStruct.IsArray
		varWorkflowTargetDataType.Max = varWorkflowTargetDataTypeWithoutEmbeddedStruct.Max
		varWorkflowTargetDataType.Min = varWorkflowTargetDataTypeWithoutEmbeddedStruct.Min
		varWorkflowTargetDataType.Properties = varWorkflowTargetDataTypeWithoutEmbeddedStruct.Properties
		*o = WorkflowTargetDataType(varWorkflowTargetDataType)
	} else {
		return err
	}

	varWorkflowTargetDataType := _WorkflowTargetDataType{}

	err = json.Unmarshal(bytes, &varWorkflowTargetDataType)
	if err == nil {
		o.WorkflowBaseDataType = varWorkflowTargetDataType.WorkflowBaseDataType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CustomDataTypeProperties")
		delete(additionalProperties, "IsArray")
		delete(additionalProperties, "Max")
		delete(additionalProperties, "Min")
		delete(additionalProperties, "Properties")

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

type NullableWorkflowTargetDataType struct {
	value *WorkflowTargetDataType
	isSet bool
}

func (v NullableWorkflowTargetDataType) Get() *WorkflowTargetDataType {
	return v.value
}

func (v *NullableWorkflowTargetDataType) Set(val *WorkflowTargetDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTargetDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTargetDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTargetDataType(val *WorkflowTargetDataType) *NullableWorkflowTargetDataType {
	return &NullableWorkflowTargetDataType{value: val, isSet: true}
}

func (v NullableWorkflowTargetDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTargetDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
