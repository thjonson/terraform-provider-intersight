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
)

// HyperflexHealthCheckScriptInfoAllOf Definition of the list of properties defined in 'hyperflex.HealthCheckScriptInfo', excluding properties defined in parent classes.
type HyperflexHealthCheckScriptInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Health check aggregate script that runs in the HyperFlex Leader Node. | It aggregates the output of all HyperFlex nodes and provides the health check result.
	AggregateScriptName *string `json:"AggregateScriptName,omitempty"`
	// HyperFlex Data Platform version running on the target device.
	HyperflexVersion *string `json:"HyperflexVersion,omitempty"`
	// Location of the health check script's execution on the HyperFlex device.
	ScriptExecuteLocation *string `json:"ScriptExecuteLocation,omitempty"`
	// Input for the health check script execution.
	ScriptInput *string `json:"ScriptInput,omitempty"`
	// Name of the health check script to be executed.
	ScriptName *string `json:"ScriptName,omitempty"`
	// Version of the health check script associated with the health check definition.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHealthCheckScriptInfoAllOf HyperflexHealthCheckScriptInfoAllOf

// NewHyperflexHealthCheckScriptInfoAllOf instantiates a new HyperflexHealthCheckScriptInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHealthCheckScriptInfoAllOf(classId string, objectType string) *HyperflexHealthCheckScriptInfoAllOf {
	this := HyperflexHealthCheckScriptInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHealthCheckScriptInfoAllOfWithDefaults instantiates a new HyperflexHealthCheckScriptInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHealthCheckScriptInfoAllOfWithDefaults() *HyperflexHealthCheckScriptInfoAllOf {
	this := HyperflexHealthCheckScriptInfoAllOf{}
	var classId string = "hyperflex.HealthCheckScriptInfo"
	this.ClassId = classId
	var objectType string = "hyperflex.HealthCheckScriptInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHealthCheckScriptInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckScriptInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHealthCheckScriptInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHealthCheckScriptInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckScriptInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHealthCheckScriptInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAggregateScriptName returns the AggregateScriptName field value if set, zero value otherwise.
func (o *HyperflexHealthCheckScriptInfoAllOf) GetAggregateScriptName() string {
	if o == nil || o.AggregateScriptName == nil {
		var ret string
		return ret
	}
	return *o.AggregateScriptName
}

// GetAggregateScriptNameOk returns a tuple with the AggregateScriptName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckScriptInfoAllOf) GetAggregateScriptNameOk() (*string, bool) {
	if o == nil || o.AggregateScriptName == nil {
		return nil, false
	}
	return o.AggregateScriptName, true
}

// HasAggregateScriptName returns a boolean if a field has been set.
func (o *HyperflexHealthCheckScriptInfoAllOf) HasAggregateScriptName() bool {
	if o != nil && o.AggregateScriptName != nil {
		return true
	}

	return false
}

// SetAggregateScriptName gets a reference to the given string and assigns it to the AggregateScriptName field.
func (o *HyperflexHealthCheckScriptInfoAllOf) SetAggregateScriptName(v string) {
	o.AggregateScriptName = &v
}

// GetHyperflexVersion returns the HyperflexVersion field value if set, zero value otherwise.
func (o *HyperflexHealthCheckScriptInfoAllOf) GetHyperflexVersion() string {
	if o == nil || o.HyperflexVersion == nil {
		var ret string
		return ret
	}
	return *o.HyperflexVersion
}

// GetHyperflexVersionOk returns a tuple with the HyperflexVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckScriptInfoAllOf) GetHyperflexVersionOk() (*string, bool) {
	if o == nil || o.HyperflexVersion == nil {
		return nil, false
	}
	return o.HyperflexVersion, true
}

// HasHyperflexVersion returns a boolean if a field has been set.
func (o *HyperflexHealthCheckScriptInfoAllOf) HasHyperflexVersion() bool {
	if o != nil && o.HyperflexVersion != nil {
		return true
	}

	return false
}

// SetHyperflexVersion gets a reference to the given string and assigns it to the HyperflexVersion field.
func (o *HyperflexHealthCheckScriptInfoAllOf) SetHyperflexVersion(v string) {
	o.HyperflexVersion = &v
}

// GetScriptExecuteLocation returns the ScriptExecuteLocation field value if set, zero value otherwise.
func (o *HyperflexHealthCheckScriptInfoAllOf) GetScriptExecuteLocation() string {
	if o == nil || o.ScriptExecuteLocation == nil {
		var ret string
		return ret
	}
	return *o.ScriptExecuteLocation
}

// GetScriptExecuteLocationOk returns a tuple with the ScriptExecuteLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckScriptInfoAllOf) GetScriptExecuteLocationOk() (*string, bool) {
	if o == nil || o.ScriptExecuteLocation == nil {
		return nil, false
	}
	return o.ScriptExecuteLocation, true
}

// HasScriptExecuteLocation returns a boolean if a field has been set.
func (o *HyperflexHealthCheckScriptInfoAllOf) HasScriptExecuteLocation() bool {
	if o != nil && o.ScriptExecuteLocation != nil {
		return true
	}

	return false
}

// SetScriptExecuteLocation gets a reference to the given string and assigns it to the ScriptExecuteLocation field.
func (o *HyperflexHealthCheckScriptInfoAllOf) SetScriptExecuteLocation(v string) {
	o.ScriptExecuteLocation = &v
}

// GetScriptInput returns the ScriptInput field value if set, zero value otherwise.
func (o *HyperflexHealthCheckScriptInfoAllOf) GetScriptInput() string {
	if o == nil || o.ScriptInput == nil {
		var ret string
		return ret
	}
	return *o.ScriptInput
}

// GetScriptInputOk returns a tuple with the ScriptInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckScriptInfoAllOf) GetScriptInputOk() (*string, bool) {
	if o == nil || o.ScriptInput == nil {
		return nil, false
	}
	return o.ScriptInput, true
}

// HasScriptInput returns a boolean if a field has been set.
func (o *HyperflexHealthCheckScriptInfoAllOf) HasScriptInput() bool {
	if o != nil && o.ScriptInput != nil {
		return true
	}

	return false
}

// SetScriptInput gets a reference to the given string and assigns it to the ScriptInput field.
func (o *HyperflexHealthCheckScriptInfoAllOf) SetScriptInput(v string) {
	o.ScriptInput = &v
}

// GetScriptName returns the ScriptName field value if set, zero value otherwise.
func (o *HyperflexHealthCheckScriptInfoAllOf) GetScriptName() string {
	if o == nil || o.ScriptName == nil {
		var ret string
		return ret
	}
	return *o.ScriptName
}

// GetScriptNameOk returns a tuple with the ScriptName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckScriptInfoAllOf) GetScriptNameOk() (*string, bool) {
	if o == nil || o.ScriptName == nil {
		return nil, false
	}
	return o.ScriptName, true
}

// HasScriptName returns a boolean if a field has been set.
func (o *HyperflexHealthCheckScriptInfoAllOf) HasScriptName() bool {
	if o != nil && o.ScriptName != nil {
		return true
	}

	return false
}

// SetScriptName gets a reference to the given string and assigns it to the ScriptName field.
func (o *HyperflexHealthCheckScriptInfoAllOf) SetScriptName(v string) {
	o.ScriptName = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexHealthCheckScriptInfoAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckScriptInfoAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexHealthCheckScriptInfoAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexHealthCheckScriptInfoAllOf) SetVersion(v string) {
	o.Version = &v
}

func (o HyperflexHealthCheckScriptInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AggregateScriptName != nil {
		toSerialize["AggregateScriptName"] = o.AggregateScriptName
	}
	if o.HyperflexVersion != nil {
		toSerialize["HyperflexVersion"] = o.HyperflexVersion
	}
	if o.ScriptExecuteLocation != nil {
		toSerialize["ScriptExecuteLocation"] = o.ScriptExecuteLocation
	}
	if o.ScriptInput != nil {
		toSerialize["ScriptInput"] = o.ScriptInput
	}
	if o.ScriptName != nil {
		toSerialize["ScriptName"] = o.ScriptName
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHealthCheckScriptInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHealthCheckScriptInfoAllOf := _HyperflexHealthCheckScriptInfoAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexHealthCheckScriptInfoAllOf); err == nil {
		*o = HyperflexHealthCheckScriptInfoAllOf(varHyperflexHealthCheckScriptInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AggregateScriptName")
		delete(additionalProperties, "HyperflexVersion")
		delete(additionalProperties, "ScriptExecuteLocation")
		delete(additionalProperties, "ScriptInput")
		delete(additionalProperties, "ScriptName")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHealthCheckScriptInfoAllOf struct {
	value *HyperflexHealthCheckScriptInfoAllOf
	isSet bool
}

func (v NullableHyperflexHealthCheckScriptInfoAllOf) Get() *HyperflexHealthCheckScriptInfoAllOf {
	return v.value
}

func (v *NullableHyperflexHealthCheckScriptInfoAllOf) Set(val *HyperflexHealthCheckScriptInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHealthCheckScriptInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHealthCheckScriptInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHealthCheckScriptInfoAllOf(val *HyperflexHealthCheckScriptInfoAllOf) *NullableHyperflexHealthCheckScriptInfoAllOf {
	return &NullableHyperflexHealthCheckScriptInfoAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHealthCheckScriptInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHealthCheckScriptInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
