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
	"reflect"
	"strings"
)

// WorkflowActionWorkflowDefinition Definition to capture the workflow definition which will be used in the action of a solution definition.
type WorkflowActionWorkflowDefinition struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Specify the catalog moid that this workflow belongs. When catalog moid is not specified then the catalog of the solution is used first and if no workflow can be found in that catalog, then the shared catalog from Intersight is used.
	CatalogMoid *string `json:"CatalogMoid,omitempty"`
	// The description of this workflow instance.
	Description *string `json:"Description,omitempty"`
	// Capture the mapping of ActionDefinition inputDefinition to workflow definition.
	InputParameters interface{} `json:"InputParameters,omitempty"`
	// A user defined label identifier of the workflow used for UI display.
	Label *string `json:"Label,omitempty"`
	// The name of the workflow, this name must be unique across all the workflow definition used within the action definitions.
	Name *string `json:"Name,omitempty"`
	// The workflow definition version to use as subworkflow. When no version is specified then the default version of the workflow at the time of creating or updating this workflow is used.
	Version *int64 `json:"Version,omitempty"`
	// The qualified name of workflow that should be executed.
	WorkflowDefinitionName *string `json:"WorkflowDefinitionName,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _WorkflowActionWorkflowDefinition WorkflowActionWorkflowDefinition

// NewWorkflowActionWorkflowDefinition instantiates a new WorkflowActionWorkflowDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowActionWorkflowDefinition(classId string, objectType string) *WorkflowActionWorkflowDefinition {
	this := WorkflowActionWorkflowDefinition{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowActionWorkflowDefinitionWithDefaults instantiates a new WorkflowActionWorkflowDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowActionWorkflowDefinitionWithDefaults() *WorkflowActionWorkflowDefinition {
	this := WorkflowActionWorkflowDefinition{}
	var classId string = "workflow.ActionWorkflowDefinition"
	this.ClassId = classId
	var objectType string = "workflow.ActionWorkflowDefinition"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowActionWorkflowDefinition) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowActionWorkflowDefinition) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowActionWorkflowDefinition) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowActionWorkflowDefinition) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowActionWorkflowDefinition) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowActionWorkflowDefinition) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCatalogMoid returns the CatalogMoid field value if set, zero value otherwise.
func (o *WorkflowActionWorkflowDefinition) GetCatalogMoid() string {
	if o == nil || o.CatalogMoid == nil {
		var ret string
		return ret
	}
	return *o.CatalogMoid
}

// GetCatalogMoidOk returns a tuple with the CatalogMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActionWorkflowDefinition) GetCatalogMoidOk() (*string, bool) {
	if o == nil || o.CatalogMoid == nil {
		return nil, false
	}
	return o.CatalogMoid, true
}

// HasCatalogMoid returns a boolean if a field has been set.
func (o *WorkflowActionWorkflowDefinition) HasCatalogMoid() bool {
	if o != nil && o.CatalogMoid != nil {
		return true
	}

	return false
}

// SetCatalogMoid gets a reference to the given string and assigns it to the CatalogMoid field.
func (o *WorkflowActionWorkflowDefinition) SetCatalogMoid(v string) {
	o.CatalogMoid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowActionWorkflowDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActionWorkflowDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowActionWorkflowDefinition) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowActionWorkflowDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetInputParameters returns the InputParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowActionWorkflowDefinition) GetInputParameters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.InputParameters
}

// GetInputParametersOk returns a tuple with the InputParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowActionWorkflowDefinition) GetInputParametersOk() (*interface{}, bool) {
	if o == nil || o.InputParameters == nil {
		return nil, false
	}
	return &o.InputParameters, true
}

// HasInputParameters returns a boolean if a field has been set.
func (o *WorkflowActionWorkflowDefinition) HasInputParameters() bool {
	if o != nil && o.InputParameters != nil {
		return true
	}

	return false
}

// SetInputParameters gets a reference to the given interface{} and assigns it to the InputParameters field.
func (o *WorkflowActionWorkflowDefinition) SetInputParameters(v interface{}) {
	o.InputParameters = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowActionWorkflowDefinition) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActionWorkflowDefinition) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowActionWorkflowDefinition) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowActionWorkflowDefinition) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowActionWorkflowDefinition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActionWorkflowDefinition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowActionWorkflowDefinition) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowActionWorkflowDefinition) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WorkflowActionWorkflowDefinition) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActionWorkflowDefinition) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WorkflowActionWorkflowDefinition) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *WorkflowActionWorkflowDefinition) SetVersion(v int64) {
	o.Version = &v
}

// GetWorkflowDefinitionName returns the WorkflowDefinitionName field value if set, zero value otherwise.
func (o *WorkflowActionWorkflowDefinition) GetWorkflowDefinitionName() string {
	if o == nil || o.WorkflowDefinitionName == nil {
		var ret string
		return ret
	}
	return *o.WorkflowDefinitionName
}

// GetWorkflowDefinitionNameOk returns a tuple with the WorkflowDefinitionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActionWorkflowDefinition) GetWorkflowDefinitionNameOk() (*string, bool) {
	if o == nil || o.WorkflowDefinitionName == nil {
		return nil, false
	}
	return o.WorkflowDefinitionName, true
}

// HasWorkflowDefinitionName returns a boolean if a field has been set.
func (o *WorkflowActionWorkflowDefinition) HasWorkflowDefinitionName() bool {
	if o != nil && o.WorkflowDefinitionName != nil {
		return true
	}

	return false
}

// SetWorkflowDefinitionName gets a reference to the given string and assigns it to the WorkflowDefinitionName field.
func (o *WorkflowActionWorkflowDefinition) SetWorkflowDefinitionName(v string) {
	o.WorkflowDefinitionName = &v
}

func (o WorkflowActionWorkflowDefinition) MarshalJSON() ([]byte, error) {
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
	if o.CatalogMoid != nil {
		toSerialize["CatalogMoid"] = o.CatalogMoid
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.InputParameters != nil {
		toSerialize["InputParameters"] = o.InputParameters
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.WorkflowDefinitionName != nil {
		toSerialize["WorkflowDefinitionName"] = o.WorkflowDefinitionName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowActionWorkflowDefinition) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowActionWorkflowDefinitionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Specify the catalog moid that this workflow belongs. When catalog moid is not specified then the catalog of the solution is used first and if no workflow can be found in that catalog, then the shared catalog from Intersight is used.
		CatalogMoid *string `json:"CatalogMoid,omitempty"`
		// The description of this workflow instance.
		Description *string `json:"Description,omitempty"`
		// Capture the mapping of ActionDefinition inputDefinition to workflow definition.
		InputParameters interface{} `json:"InputParameters,omitempty"`
		// A user defined label identifier of the workflow used for UI display.
		Label *string `json:"Label,omitempty"`
		// The name of the workflow, this name must be unique across all the workflow definition used within the action definitions.
		Name *string `json:"Name,omitempty"`
		// The workflow definition version to use as subworkflow. When no version is specified then the default version of the workflow at the time of creating or updating this workflow is used.
		Version *int64 `json:"Version,omitempty"`
		// The qualified name of workflow that should be executed.
		WorkflowDefinitionName *string `json:"WorkflowDefinitionName,omitempty"`
	}

	varWorkflowActionWorkflowDefinitionWithoutEmbeddedStruct := WorkflowActionWorkflowDefinitionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowActionWorkflowDefinitionWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowActionWorkflowDefinition := _WorkflowActionWorkflowDefinition{}
		varWorkflowActionWorkflowDefinition.ClassId = varWorkflowActionWorkflowDefinitionWithoutEmbeddedStruct.ClassId
		varWorkflowActionWorkflowDefinition.ObjectType = varWorkflowActionWorkflowDefinitionWithoutEmbeddedStruct.ObjectType
		varWorkflowActionWorkflowDefinition.CatalogMoid = varWorkflowActionWorkflowDefinitionWithoutEmbeddedStruct.CatalogMoid
		varWorkflowActionWorkflowDefinition.Description = varWorkflowActionWorkflowDefinitionWithoutEmbeddedStruct.Description
		varWorkflowActionWorkflowDefinition.InputParameters = varWorkflowActionWorkflowDefinitionWithoutEmbeddedStruct.InputParameters
		varWorkflowActionWorkflowDefinition.Label = varWorkflowActionWorkflowDefinitionWithoutEmbeddedStruct.Label
		varWorkflowActionWorkflowDefinition.Name = varWorkflowActionWorkflowDefinitionWithoutEmbeddedStruct.Name
		varWorkflowActionWorkflowDefinition.Version = varWorkflowActionWorkflowDefinitionWithoutEmbeddedStruct.Version
		varWorkflowActionWorkflowDefinition.WorkflowDefinitionName = varWorkflowActionWorkflowDefinitionWithoutEmbeddedStruct.WorkflowDefinitionName
		*o = WorkflowActionWorkflowDefinition(varWorkflowActionWorkflowDefinition)
	} else {
		return err
	}

	varWorkflowActionWorkflowDefinition := _WorkflowActionWorkflowDefinition{}

	err = json.Unmarshal(bytes, &varWorkflowActionWorkflowDefinition)
	if err == nil {
		o.MoBaseComplexType = varWorkflowActionWorkflowDefinition.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CatalogMoid")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "InputParameters")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "WorkflowDefinitionName")

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

type NullableWorkflowActionWorkflowDefinition struct {
	value *WorkflowActionWorkflowDefinition
	isSet bool
}

func (v NullableWorkflowActionWorkflowDefinition) Get() *WorkflowActionWorkflowDefinition {
	return v.value
}

func (v *NullableWorkflowActionWorkflowDefinition) Set(val *WorkflowActionWorkflowDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowActionWorkflowDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowActionWorkflowDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowActionWorkflowDefinition(val *WorkflowActionWorkflowDefinition) *NullableWorkflowActionWorkflowDefinition {
	return &NullableWorkflowActionWorkflowDefinition{value: val, isSet: true}
}

func (v NullableWorkflowActionWorkflowDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowActionWorkflowDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
