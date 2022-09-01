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

// WorkflowTaskLoopInfoAllOf Definition of the list of properties defined in 'workflow.TaskLoopInfo', excluding properties defined in parent classes.
type WorkflowTaskLoopInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This specifies the count of iteration for the specific task executed inside the loop.
	Iteration *int64 `json:"Iteration,omitempty"`
	// Label of the loop task inside which this task is executed.
	LoopTaskLabel *string `json:"LoopTaskLabel,omitempty"`
	// Name of the loop task inside which this task is executed.
	LoopTaskName *string `json:"LoopTaskName,omitempty"`
	// This specifies the type of loop, Serial or Parallel. * `None` - The enum specifies the option as None which implies this is not a Loop type and this is the default value for loop type. * `Parallel` - The enum specifies the option as Parallel where the loop task type is parallel loop. * `Serial` - The enum specifies the option as Serial where the loop task type is serial loop.
	LoopType             *string `json:"LoopType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTaskLoopInfoAllOf WorkflowTaskLoopInfoAllOf

// NewWorkflowTaskLoopInfoAllOf instantiates a new WorkflowTaskLoopInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskLoopInfoAllOf(classId string, objectType string) *WorkflowTaskLoopInfoAllOf {
	this := WorkflowTaskLoopInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var loopType string = "None"
	this.LoopType = &loopType
	return &this
}

// NewWorkflowTaskLoopInfoAllOfWithDefaults instantiates a new WorkflowTaskLoopInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskLoopInfoAllOfWithDefaults() *WorkflowTaskLoopInfoAllOf {
	this := WorkflowTaskLoopInfoAllOf{}
	var classId string = "workflow.TaskLoopInfo"
	this.ClassId = classId
	var objectType string = "workflow.TaskLoopInfo"
	this.ObjectType = objectType
	var loopType string = "None"
	this.LoopType = &loopType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowTaskLoopInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskLoopInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowTaskLoopInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowTaskLoopInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskLoopInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowTaskLoopInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIteration returns the Iteration field value if set, zero value otherwise.
func (o *WorkflowTaskLoopInfoAllOf) GetIteration() int64 {
	if o == nil || o.Iteration == nil {
		var ret int64
		return ret
	}
	return *o.Iteration
}

// GetIterationOk returns a tuple with the Iteration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskLoopInfoAllOf) GetIterationOk() (*int64, bool) {
	if o == nil || o.Iteration == nil {
		return nil, false
	}
	return o.Iteration, true
}

// HasIteration returns a boolean if a field has been set.
func (o *WorkflowTaskLoopInfoAllOf) HasIteration() bool {
	if o != nil && o.Iteration != nil {
		return true
	}

	return false
}

// SetIteration gets a reference to the given int64 and assigns it to the Iteration field.
func (o *WorkflowTaskLoopInfoAllOf) SetIteration(v int64) {
	o.Iteration = &v
}

// GetLoopTaskLabel returns the LoopTaskLabel field value if set, zero value otherwise.
func (o *WorkflowTaskLoopInfoAllOf) GetLoopTaskLabel() string {
	if o == nil || o.LoopTaskLabel == nil {
		var ret string
		return ret
	}
	return *o.LoopTaskLabel
}

// GetLoopTaskLabelOk returns a tuple with the LoopTaskLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskLoopInfoAllOf) GetLoopTaskLabelOk() (*string, bool) {
	if o == nil || o.LoopTaskLabel == nil {
		return nil, false
	}
	return o.LoopTaskLabel, true
}

// HasLoopTaskLabel returns a boolean if a field has been set.
func (o *WorkflowTaskLoopInfoAllOf) HasLoopTaskLabel() bool {
	if o != nil && o.LoopTaskLabel != nil {
		return true
	}

	return false
}

// SetLoopTaskLabel gets a reference to the given string and assigns it to the LoopTaskLabel field.
func (o *WorkflowTaskLoopInfoAllOf) SetLoopTaskLabel(v string) {
	o.LoopTaskLabel = &v
}

// GetLoopTaskName returns the LoopTaskName field value if set, zero value otherwise.
func (o *WorkflowTaskLoopInfoAllOf) GetLoopTaskName() string {
	if o == nil || o.LoopTaskName == nil {
		var ret string
		return ret
	}
	return *o.LoopTaskName
}

// GetLoopTaskNameOk returns a tuple with the LoopTaskName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskLoopInfoAllOf) GetLoopTaskNameOk() (*string, bool) {
	if o == nil || o.LoopTaskName == nil {
		return nil, false
	}
	return o.LoopTaskName, true
}

// HasLoopTaskName returns a boolean if a field has been set.
func (o *WorkflowTaskLoopInfoAllOf) HasLoopTaskName() bool {
	if o != nil && o.LoopTaskName != nil {
		return true
	}

	return false
}

// SetLoopTaskName gets a reference to the given string and assigns it to the LoopTaskName field.
func (o *WorkflowTaskLoopInfoAllOf) SetLoopTaskName(v string) {
	o.LoopTaskName = &v
}

// GetLoopType returns the LoopType field value if set, zero value otherwise.
func (o *WorkflowTaskLoopInfoAllOf) GetLoopType() string {
	if o == nil || o.LoopType == nil {
		var ret string
		return ret
	}
	return *o.LoopType
}

// GetLoopTypeOk returns a tuple with the LoopType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskLoopInfoAllOf) GetLoopTypeOk() (*string, bool) {
	if o == nil || o.LoopType == nil {
		return nil, false
	}
	return o.LoopType, true
}

// HasLoopType returns a boolean if a field has been set.
func (o *WorkflowTaskLoopInfoAllOf) HasLoopType() bool {
	if o != nil && o.LoopType != nil {
		return true
	}

	return false
}

// SetLoopType gets a reference to the given string and assigns it to the LoopType field.
func (o *WorkflowTaskLoopInfoAllOf) SetLoopType(v string) {
	o.LoopType = &v
}

func (o WorkflowTaskLoopInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Iteration != nil {
		toSerialize["Iteration"] = o.Iteration
	}
	if o.LoopTaskLabel != nil {
		toSerialize["LoopTaskLabel"] = o.LoopTaskLabel
	}
	if o.LoopTaskName != nil {
		toSerialize["LoopTaskName"] = o.LoopTaskName
	}
	if o.LoopType != nil {
		toSerialize["LoopType"] = o.LoopType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowTaskLoopInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowTaskLoopInfoAllOf := _WorkflowTaskLoopInfoAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowTaskLoopInfoAllOf); err == nil {
		*o = WorkflowTaskLoopInfoAllOf(varWorkflowTaskLoopInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Iteration")
		delete(additionalProperties, "LoopTaskLabel")
		delete(additionalProperties, "LoopTaskName")
		delete(additionalProperties, "LoopType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowTaskLoopInfoAllOf struct {
	value *WorkflowTaskLoopInfoAllOf
	isSet bool
}

func (v NullableWorkflowTaskLoopInfoAllOf) Get() *WorkflowTaskLoopInfoAllOf {
	return v.value
}

func (v *NullableWorkflowTaskLoopInfoAllOf) Set(val *WorkflowTaskLoopInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskLoopInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskLoopInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskLoopInfoAllOf(val *WorkflowTaskLoopInfoAllOf) *NullableWorkflowTaskLoopInfoAllOf {
	return &NullableWorkflowTaskLoopInfoAllOf{value: val, isSet: true}
}

func (v NullableWorkflowTaskLoopInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskLoopInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
