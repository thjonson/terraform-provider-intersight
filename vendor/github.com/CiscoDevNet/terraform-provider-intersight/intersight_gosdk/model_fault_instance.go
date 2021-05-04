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

// FaultInstance An endpoint anomaly is represented by this object.
type FaultInstance struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The user acknowledgement state of the fault.
	Acknowledged *string `json:"Acknowledged,omitempty"`
	// The Distinguished Name of the Managed object which was affected.
	AffectedDn *string `json:"AffectedDn,omitempty"`
	// Managed object Id which was affected.
	AffectedMoId *string `json:"AffectedMoId,omitempty"`
	// Managed object type which was affected.
	AffectedMoType *string `json:"AffectedMoType,omitempty"`
	// Object Id of the parent of the Managed object which was affected.
	AncestorMoId *string `json:"AncestorMoId,omitempty"`
	// Object type of the parent of the Managed object which was affected.
	AncestorMoType *string `json:"AncestorMoType,omitempty"`
	// Numerical fault code of the fault found.
	Code *string `json:"Code,omitempty"`
	// The time of creation of the fault instance.
	CreationTime *string `json:"CreationTime,omitempty"`
	// Detailed message of the fault.
	Description *string `json:"Description,omitempty"`
	// Last transition time of the fault.
	LastTransitionTime *string `json:"LastTransitionTime,omitempty"`
	// The number of times this fault has occured.
	NumOccurrences *int64 `json:"NumOccurrences,omitempty"`
	// Current Severity of the fault found.
	OriginalSeverity *string `json:"OriginalSeverity,omitempty"`
	// The Severity of the fault prior to user update.
	PreviousSeverity *string `json:"PreviousSeverity,omitempty"`
	// The rule that is responsible for generation of the fault.
	Rule *string `json:"Rule,omitempty"`
	// Severity of the fault found.
	Severity             *string                              `json:"Severity,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FaultInstance FaultInstance

// NewFaultInstance instantiates a new FaultInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFaultInstance(classId string, objectType string) *FaultInstance {
	this := FaultInstance{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFaultInstanceWithDefaults instantiates a new FaultInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFaultInstanceWithDefaults() *FaultInstance {
	this := FaultInstance{}
	var classId string = "fault.Instance"
	this.ClassId = classId
	var objectType string = "fault.Instance"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FaultInstance) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FaultInstance) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FaultInstance) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FaultInstance) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FaultInstance) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FaultInstance) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAcknowledged returns the Acknowledged field value if set, zero value otherwise.
func (o *FaultInstance) GetAcknowledged() string {
	if o == nil || o.Acknowledged == nil {
		var ret string
		return ret
	}
	return *o.Acknowledged
}

// GetAcknowledgedOk returns a tuple with the Acknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaultInstance) GetAcknowledgedOk() (*string, bool) {
	if o == nil || o.Acknowledged == nil {
		return nil, false
	}
	return o.Acknowledged, true
}

// HasAcknowledged returns a boolean if a field has been set.
func (o *FaultInstance) HasAcknowledged() bool {
	if o != nil && o.Acknowledged != nil {
		return true
	}

	return false
}

// SetAcknowledged gets a reference to the given string and assigns it to the Acknowledged field.
func (o *FaultInstance) SetAcknowledged(v string) {
	o.Acknowledged = &v
}

// GetAffectedDn returns the AffectedDn field value if set, zero value otherwise.
func (o *FaultInstance) GetAffectedDn() string {
	if o == nil || o.AffectedDn == nil {
		var ret string
		return ret
	}
	return *o.AffectedDn
}

// GetAffectedDnOk returns a tuple with the AffectedDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaultInstance) GetAffectedDnOk() (*string, bool) {
	if o == nil || o.AffectedDn == nil {
		return nil, false
	}
	return o.AffectedDn, true
}

// HasAffectedDn returns a boolean if a field has been set.
func (o *FaultInstance) HasAffectedDn() bool {
	if o != nil && o.AffectedDn != nil {
		return true
	}

	return false
}

// SetAffectedDn gets a reference to the given string and assigns it to the AffectedDn field.
func (o *FaultInstance) SetAffectedDn(v string) {
	o.AffectedDn = &v
}

// GetAffectedMoId returns the AffectedMoId field value if set, zero value otherwise.
func (o *FaultInstance) GetAffectedMoId() string {
	if o == nil || o.AffectedMoId == nil {
		var ret string
		return ret
	}
	return *o.AffectedMoId
}

// GetAffectedMoIdOk returns a tuple with the AffectedMoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaultInstance) GetAffectedMoIdOk() (*string, bool) {
	if o == nil || o.AffectedMoId == nil {
		return nil, false
	}
	return o.AffectedMoId, true
}

// HasAffectedMoId returns a boolean if a field has been set.
func (o *FaultInstance) HasAffectedMoId() bool {
	if o != nil && o.AffectedMoId != nil {
		return true
	}

	return false
}

// SetAffectedMoId gets a reference to the given string and assigns it to the AffectedMoId field.
func (o *FaultInstance) SetAffectedMoId(v string) {
	o.AffectedMoId = &v
}

// GetAffectedMoType returns the AffectedMoType field value if set, zero value otherwise.
func (o *FaultInstance) GetAffectedMoType() string {
	if o == nil || o.AffectedMoType == nil {
		var ret string
		return ret
	}
	return *o.AffectedMoType
}

// GetAffectedMoTypeOk returns a tuple with the AffectedMoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaultInstance) GetAffectedMoTypeOk() (*string, bool) {
	if o == nil || o.AffectedMoType == nil {
		return nil, false
	}
	return o.AffectedMoType, true
}

// HasAffectedMoType returns a boolean if a field has been set.
func (o *FaultInstance) HasAffectedMoType() bool {
	if o != nil && o.AffectedMoType != nil {
		return true
	}

	return false
}

// SetAffectedMoType gets a reference to the given string and assigns it to the AffectedMoType field.
func (o *FaultInstance) SetAffectedMoType(v string) {
	o.AffectedMoType = &v
}

// GetAncestorMoId returns the AncestorMoId field value if set, zero value otherwise.
func (o *FaultInstance) GetAncestorMoId() string {
	if o == nil || o.AncestorMoId == nil {
		var ret string
		return ret
	}
	return *o.AncestorMoId
}

// GetAncestorMoIdOk returns a tuple with the AncestorMoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaultInstance) GetAncestorMoIdOk() (*string, bool) {
	if o == nil || o.AncestorMoId == nil {
		return nil, false
	}
	return o.AncestorMoId, true
}

// HasAncestorMoId returns a boolean if a field has been set.
func (o *FaultInstance) HasAncestorMoId() bool {
	if o != nil && o.AncestorMoId != nil {
		return true
	}

	return false
}

// SetAncestorMoId gets a reference to the given string and assigns it to the AncestorMoId field.
func (o *FaultInstance) SetAncestorMoId(v string) {
	o.AncestorMoId = &v
}

// GetAncestorMoType returns the AncestorMoType field value if set, zero value otherwise.
func (o *FaultInstance) GetAncestorMoType() string {
	if o == nil || o.AncestorMoType == nil {
		var ret string
		return ret
	}
	return *o.AncestorMoType
}

// GetAncestorMoTypeOk returns a tuple with the AncestorMoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaultInstance) GetAncestorMoTypeOk() (*string, bool) {
	if o == nil || o.AncestorMoType == nil {
		return nil, false
	}
	return o.AncestorMoType, true
}

// HasAncestorMoType returns a boolean if a field has been set.
func (o *FaultInstance) HasAncestorMoType() bool {
	if o != nil && o.AncestorMoType != nil {
		return true
	}

	return false
}

// SetAncestorMoType gets a reference to the given string and assigns it to the AncestorMoType field.
func (o *FaultInstance) SetAncestorMoType(v string) {
	o.AncestorMoType = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *FaultInstance) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaultInstance) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *FaultInstance) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *FaultInstance) SetCode(v string) {
	o.Code = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *FaultInstance) GetCreationTime() string {
	if o == nil || o.CreationTime == nil {
		var ret string
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaultInstance) GetCreationTimeOk() (*string, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *FaultInstance) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given string and assigns it to the CreationTime field.
func (o *FaultInstance) SetCreationTime(v string) {
	o.CreationTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FaultInstance) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaultInstance) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FaultInstance) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FaultInstance) SetDescription(v string) {
	o.Description = &v
}

// GetLastTransitionTime returns the LastTransitionTime field value if set, zero value otherwise.
func (o *FaultInstance) GetLastTransitionTime() string {
	if o == nil || o.LastTransitionTime == nil {
		var ret string
		return ret
	}
	return *o.LastTransitionTime
}

// GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaultInstance) GetLastTransitionTimeOk() (*string, bool) {
	if o == nil || o.LastTransitionTime == nil {
		return nil, false
	}
	return o.LastTransitionTime, true
}

// HasLastTransitionTime returns a boolean if a field has been set.
func (o *FaultInstance) HasLastTransitionTime() bool {
	if o != nil && o.LastTransitionTime != nil {
		return true
	}

	return false
}

// SetLastTransitionTime gets a reference to the given string and assigns it to the LastTransitionTime field.
func (o *FaultInstance) SetLastTransitionTime(v string) {
	o.LastTransitionTime = &v
}

// GetNumOccurrences returns the NumOccurrences field value if set, zero value otherwise.
func (o *FaultInstance) GetNumOccurrences() int64 {
	if o == nil || o.NumOccurrences == nil {
		var ret int64
		return ret
	}
	return *o.NumOccurrences
}

// GetNumOccurrencesOk returns a tuple with the NumOccurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaultInstance) GetNumOccurrencesOk() (*int64, bool) {
	if o == nil || o.NumOccurrences == nil {
		return nil, false
	}
	return o.NumOccurrences, true
}

// HasNumOccurrences returns a boolean if a field has been set.
func (o *FaultInstance) HasNumOccurrences() bool {
	if o != nil && o.NumOccurrences != nil {
		return true
	}

	return false
}

// SetNumOccurrences gets a reference to the given int64 and assigns it to the NumOccurrences field.
func (o *FaultInstance) SetNumOccurrences(v int64) {
	o.NumOccurrences = &v
}

// GetOriginalSeverity returns the OriginalSeverity field value if set, zero value otherwise.
func (o *FaultInstance) GetOriginalSeverity() string {
	if o == nil || o.OriginalSeverity == nil {
		var ret string
		return ret
	}
	return *o.OriginalSeverity
}

// GetOriginalSeverityOk returns a tuple with the OriginalSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaultInstance) GetOriginalSeverityOk() (*string, bool) {
	if o == nil || o.OriginalSeverity == nil {
		return nil, false
	}
	return o.OriginalSeverity, true
}

// HasOriginalSeverity returns a boolean if a field has been set.
func (o *FaultInstance) HasOriginalSeverity() bool {
	if o != nil && o.OriginalSeverity != nil {
		return true
	}

	return false
}

// SetOriginalSeverity gets a reference to the given string and assigns it to the OriginalSeverity field.
func (o *FaultInstance) SetOriginalSeverity(v string) {
	o.OriginalSeverity = &v
}

// GetPreviousSeverity returns the PreviousSeverity field value if set, zero value otherwise.
func (o *FaultInstance) GetPreviousSeverity() string {
	if o == nil || o.PreviousSeverity == nil {
		var ret string
		return ret
	}
	return *o.PreviousSeverity
}

// GetPreviousSeverityOk returns a tuple with the PreviousSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaultInstance) GetPreviousSeverityOk() (*string, bool) {
	if o == nil || o.PreviousSeverity == nil {
		return nil, false
	}
	return o.PreviousSeverity, true
}

// HasPreviousSeverity returns a boolean if a field has been set.
func (o *FaultInstance) HasPreviousSeverity() bool {
	if o != nil && o.PreviousSeverity != nil {
		return true
	}

	return false
}

// SetPreviousSeverity gets a reference to the given string and assigns it to the PreviousSeverity field.
func (o *FaultInstance) SetPreviousSeverity(v string) {
	o.PreviousSeverity = &v
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *FaultInstance) GetRule() string {
	if o == nil || o.Rule == nil {
		var ret string
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaultInstance) GetRuleOk() (*string, bool) {
	if o == nil || o.Rule == nil {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *FaultInstance) HasRule() bool {
	if o != nil && o.Rule != nil {
		return true
	}

	return false
}

// SetRule gets a reference to the given string and assigns it to the Rule field.
func (o *FaultInstance) SetRule(v string) {
	o.Rule = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *FaultInstance) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaultInstance) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *FaultInstance) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *FaultInstance) SetSeverity(v string) {
	o.Severity = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *FaultInstance) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaultInstance) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *FaultInstance) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *FaultInstance) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *FaultInstance) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaultInstance) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *FaultInstance) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *FaultInstance) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o FaultInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Acknowledged != nil {
		toSerialize["Acknowledged"] = o.Acknowledged
	}
	if o.AffectedDn != nil {
		toSerialize["AffectedDn"] = o.AffectedDn
	}
	if o.AffectedMoId != nil {
		toSerialize["AffectedMoId"] = o.AffectedMoId
	}
	if o.AffectedMoType != nil {
		toSerialize["AffectedMoType"] = o.AffectedMoType
	}
	if o.AncestorMoId != nil {
		toSerialize["AncestorMoId"] = o.AncestorMoId
	}
	if o.AncestorMoType != nil {
		toSerialize["AncestorMoType"] = o.AncestorMoType
	}
	if o.Code != nil {
		toSerialize["Code"] = o.Code
	}
	if o.CreationTime != nil {
		toSerialize["CreationTime"] = o.CreationTime
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.LastTransitionTime != nil {
		toSerialize["LastTransitionTime"] = o.LastTransitionTime
	}
	if o.NumOccurrences != nil {
		toSerialize["NumOccurrences"] = o.NumOccurrences
	}
	if o.OriginalSeverity != nil {
		toSerialize["OriginalSeverity"] = o.OriginalSeverity
	}
	if o.PreviousSeverity != nil {
		toSerialize["PreviousSeverity"] = o.PreviousSeverity
	}
	if o.Rule != nil {
		toSerialize["Rule"] = o.Rule
	}
	if o.Severity != nil {
		toSerialize["Severity"] = o.Severity
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FaultInstance) UnmarshalJSON(bytes []byte) (err error) {
	type FaultInstanceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The user acknowledgement state of the fault.
		Acknowledged *string `json:"Acknowledged,omitempty"`
		// The Distinguished Name of the Managed object which was affected.
		AffectedDn *string `json:"AffectedDn,omitempty"`
		// Managed object Id which was affected.
		AffectedMoId *string `json:"AffectedMoId,omitempty"`
		// Managed object type which was affected.
		AffectedMoType *string `json:"AffectedMoType,omitempty"`
		// Object Id of the parent of the Managed object which was affected.
		AncestorMoId *string `json:"AncestorMoId,omitempty"`
		// Object type of the parent of the Managed object which was affected.
		AncestorMoType *string `json:"AncestorMoType,omitempty"`
		// Numerical fault code of the fault found.
		Code *string `json:"Code,omitempty"`
		// The time of creation of the fault instance.
		CreationTime *string `json:"CreationTime,omitempty"`
		// Detailed message of the fault.
		Description *string `json:"Description,omitempty"`
		// Last transition time of the fault.
		LastTransitionTime *string `json:"LastTransitionTime,omitempty"`
		// The number of times this fault has occured.
		NumOccurrences *int64 `json:"NumOccurrences,omitempty"`
		// Current Severity of the fault found.
		OriginalSeverity *string `json:"OriginalSeverity,omitempty"`
		// The Severity of the fault prior to user update.
		PreviousSeverity *string `json:"PreviousSeverity,omitempty"`
		// The rule that is responsible for generation of the fault.
		Rule *string `json:"Rule,omitempty"`
		// Severity of the fault found.
		Severity            *string                              `json:"Severity,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varFaultInstanceWithoutEmbeddedStruct := FaultInstanceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFaultInstanceWithoutEmbeddedStruct)
	if err == nil {
		varFaultInstance := _FaultInstance{}
		varFaultInstance.ClassId = varFaultInstanceWithoutEmbeddedStruct.ClassId
		varFaultInstance.ObjectType = varFaultInstanceWithoutEmbeddedStruct.ObjectType
		varFaultInstance.Acknowledged = varFaultInstanceWithoutEmbeddedStruct.Acknowledged
		varFaultInstance.AffectedDn = varFaultInstanceWithoutEmbeddedStruct.AffectedDn
		varFaultInstance.AffectedMoId = varFaultInstanceWithoutEmbeddedStruct.AffectedMoId
		varFaultInstance.AffectedMoType = varFaultInstanceWithoutEmbeddedStruct.AffectedMoType
		varFaultInstance.AncestorMoId = varFaultInstanceWithoutEmbeddedStruct.AncestorMoId
		varFaultInstance.AncestorMoType = varFaultInstanceWithoutEmbeddedStruct.AncestorMoType
		varFaultInstance.Code = varFaultInstanceWithoutEmbeddedStruct.Code
		varFaultInstance.CreationTime = varFaultInstanceWithoutEmbeddedStruct.CreationTime
		varFaultInstance.Description = varFaultInstanceWithoutEmbeddedStruct.Description
		varFaultInstance.LastTransitionTime = varFaultInstanceWithoutEmbeddedStruct.LastTransitionTime
		varFaultInstance.NumOccurrences = varFaultInstanceWithoutEmbeddedStruct.NumOccurrences
		varFaultInstance.OriginalSeverity = varFaultInstanceWithoutEmbeddedStruct.OriginalSeverity
		varFaultInstance.PreviousSeverity = varFaultInstanceWithoutEmbeddedStruct.PreviousSeverity
		varFaultInstance.Rule = varFaultInstanceWithoutEmbeddedStruct.Rule
		varFaultInstance.Severity = varFaultInstanceWithoutEmbeddedStruct.Severity
		varFaultInstance.InventoryDeviceInfo = varFaultInstanceWithoutEmbeddedStruct.InventoryDeviceInfo
		varFaultInstance.RegisteredDevice = varFaultInstanceWithoutEmbeddedStruct.RegisteredDevice
		*o = FaultInstance(varFaultInstance)
	} else {
		return err
	}

	varFaultInstance := _FaultInstance{}

	err = json.Unmarshal(bytes, &varFaultInstance)
	if err == nil {
		o.InventoryBase = varFaultInstance.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Acknowledged")
		delete(additionalProperties, "AffectedDn")
		delete(additionalProperties, "AffectedMoId")
		delete(additionalProperties, "AffectedMoType")
		delete(additionalProperties, "AncestorMoId")
		delete(additionalProperties, "AncestorMoType")
		delete(additionalProperties, "Code")
		delete(additionalProperties, "CreationTime")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "LastTransitionTime")
		delete(additionalProperties, "NumOccurrences")
		delete(additionalProperties, "OriginalSeverity")
		delete(additionalProperties, "PreviousSeverity")
		delete(additionalProperties, "Rule")
		delete(additionalProperties, "Severity")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableFaultInstance struct {
	value *FaultInstance
	isSet bool
}

func (v NullableFaultInstance) Get() *FaultInstance {
	return v.value
}

func (v *NullableFaultInstance) Set(val *FaultInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableFaultInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableFaultInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFaultInstance(val *FaultInstance) *NullableFaultInstance {
	return &NullableFaultInstance{value: val, isSet: true}
}

func (v NullableFaultInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFaultInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
