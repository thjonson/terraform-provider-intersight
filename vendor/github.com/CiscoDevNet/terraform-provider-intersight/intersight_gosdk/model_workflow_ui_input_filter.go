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

// WorkflowUiInputFilter A cascade filter enables the definition or extension of an optional filter for Intersight APIs, to narrow down the search results. This can be defined for workflow inputs only.
type WorkflowUiInputFilter struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string   `json:"ObjectType"`
	Filters    []string `json:"Filters,omitempty"`
	// Name for the input definition to which this filter applies. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character. When defining the cascade filter for a sub property, use a period (.) to seperate each section of the name like \"StorageConfig.Volume\" where 'StorageConfig' is an input name and 'Volume' is a sub property defined through custom data type definition.
	Name *string `json:"Name,omitempty"`
	// Help message shown to the user about which prior input needs to be selected to enable the input mapped to this filter.
	UserHelpMessage      *string `json:"UserHelpMessage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowUiInputFilter WorkflowUiInputFilter

// NewWorkflowUiInputFilter instantiates a new WorkflowUiInputFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowUiInputFilter(classId string, objectType string) *WorkflowUiInputFilter {
	this := WorkflowUiInputFilter{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowUiInputFilterWithDefaults instantiates a new WorkflowUiInputFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowUiInputFilterWithDefaults() *WorkflowUiInputFilter {
	this := WorkflowUiInputFilter{}
	var classId string = "workflow.UiInputFilter"
	this.ClassId = classId
	var objectType string = "workflow.UiInputFilter"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowUiInputFilter) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowUiInputFilter) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowUiInputFilter) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowUiInputFilter) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowUiInputFilter) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowUiInputFilter) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFilters returns the Filters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowUiInputFilter) GetFilters() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowUiInputFilter) GetFiltersOk() (*[]string, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return &o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *WorkflowUiInputFilter) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []string and assigns it to the Filters field.
func (o *WorkflowUiInputFilter) SetFilters(v []string) {
	o.Filters = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowUiInputFilter) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowUiInputFilter) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowUiInputFilter) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowUiInputFilter) SetName(v string) {
	o.Name = &v
}

// GetUserHelpMessage returns the UserHelpMessage field value if set, zero value otherwise.
func (o *WorkflowUiInputFilter) GetUserHelpMessage() string {
	if o == nil || o.UserHelpMessage == nil {
		var ret string
		return ret
	}
	return *o.UserHelpMessage
}

// GetUserHelpMessageOk returns a tuple with the UserHelpMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowUiInputFilter) GetUserHelpMessageOk() (*string, bool) {
	if o == nil || o.UserHelpMessage == nil {
		return nil, false
	}
	return o.UserHelpMessage, true
}

// HasUserHelpMessage returns a boolean if a field has been set.
func (o *WorkflowUiInputFilter) HasUserHelpMessage() bool {
	if o != nil && o.UserHelpMessage != nil {
		return true
	}

	return false
}

// SetUserHelpMessage gets a reference to the given string and assigns it to the UserHelpMessage field.
func (o *WorkflowUiInputFilter) SetUserHelpMessage(v string) {
	o.UserHelpMessage = &v
}

func (o WorkflowUiInputFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Filters != nil {
		toSerialize["Filters"] = o.Filters
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.UserHelpMessage != nil {
		toSerialize["UserHelpMessage"] = o.UserHelpMessage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowUiInputFilter) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowUiInputFilterWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string   `json:"ObjectType"`
		Filters    []string `json:"Filters,omitempty"`
		// Name for the input definition to which this filter applies. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character. When defining the cascade filter for a sub property, use a period (.) to seperate each section of the name like \"StorageConfig.Volume\" where 'StorageConfig' is an input name and 'Volume' is a sub property defined through custom data type definition.
		Name *string `json:"Name,omitempty"`
		// Help message shown to the user about which prior input needs to be selected to enable the input mapped to this filter.
		UserHelpMessage *string `json:"UserHelpMessage,omitempty"`
	}

	varWorkflowUiInputFilterWithoutEmbeddedStruct := WorkflowUiInputFilterWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowUiInputFilterWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowUiInputFilter := _WorkflowUiInputFilter{}
		varWorkflowUiInputFilter.ClassId = varWorkflowUiInputFilterWithoutEmbeddedStruct.ClassId
		varWorkflowUiInputFilter.ObjectType = varWorkflowUiInputFilterWithoutEmbeddedStruct.ObjectType
		varWorkflowUiInputFilter.Filters = varWorkflowUiInputFilterWithoutEmbeddedStruct.Filters
		varWorkflowUiInputFilter.Name = varWorkflowUiInputFilterWithoutEmbeddedStruct.Name
		varWorkflowUiInputFilter.UserHelpMessage = varWorkflowUiInputFilterWithoutEmbeddedStruct.UserHelpMessage
		*o = WorkflowUiInputFilter(varWorkflowUiInputFilter)
	} else {
		return err
	}

	varWorkflowUiInputFilter := _WorkflowUiInputFilter{}

	err = json.Unmarshal(bytes, &varWorkflowUiInputFilter)
	if err == nil {
		o.MoBaseComplexType = varWorkflowUiInputFilter.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Filters")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "UserHelpMessage")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableWorkflowUiInputFilter struct {
	value *WorkflowUiInputFilter
	isSet bool
}

func (v NullableWorkflowUiInputFilter) Get() *WorkflowUiInputFilter {
	return v.value
}

func (v *NullableWorkflowUiInputFilter) Set(val *WorkflowUiInputFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowUiInputFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowUiInputFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowUiInputFilter(val *WorkflowUiInputFilter) *NullableWorkflowUiInputFilter {
	return &NullableWorkflowUiInputFilter{value: val, isSet: true}
}

func (v NullableWorkflowUiInputFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowUiInputFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
