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

// PoolAbstractPoolAllOf Definition of the list of properties defined in 'pool.AbstractPool', excluding properties defined in parent classes.
type PoolAbstractPoolAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Number of IDs that are currently assigned.
	Assigned *int64 `json:"Assigned,omitempty"`
	// Assignment order decides the order in which the next identifier is allocated. * `sequential` - Identifiers are assigned in a sequential order. * `default` - Assignment order is decided by the system.
	AssignmentOrder *string `json:"AssignmentOrder,omitempty"`
	// Total number of identifiers in this pool.
	Size                 *int64 `json:"Size,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PoolAbstractPoolAllOf PoolAbstractPoolAllOf

// NewPoolAbstractPoolAllOf instantiates a new PoolAbstractPoolAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolAbstractPoolAllOf(classId string, objectType string) *PoolAbstractPoolAllOf {
	this := PoolAbstractPoolAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var assignmentOrder string = "sequential"
	this.AssignmentOrder = &assignmentOrder
	return &this
}

// NewPoolAbstractPoolAllOfWithDefaults instantiates a new PoolAbstractPoolAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolAbstractPoolAllOfWithDefaults() *PoolAbstractPoolAllOf {
	this := PoolAbstractPoolAllOf{}
	var assignmentOrder string = "sequential"
	this.AssignmentOrder = &assignmentOrder
	return &this
}

// GetClassId returns the ClassId field value
func (o *PoolAbstractPoolAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PoolAbstractPoolAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PoolAbstractPoolAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PoolAbstractPoolAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PoolAbstractPoolAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PoolAbstractPoolAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *PoolAbstractPoolAllOf) GetAssigned() int64 {
	if o == nil || o.Assigned == nil {
		var ret int64
		return ret
	}
	return *o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractPoolAllOf) GetAssignedOk() (*int64, bool) {
	if o == nil || o.Assigned == nil {
		return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *PoolAbstractPoolAllOf) HasAssigned() bool {
	if o != nil && o.Assigned != nil {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given int64 and assigns it to the Assigned field.
func (o *PoolAbstractPoolAllOf) SetAssigned(v int64) {
	o.Assigned = &v
}

// GetAssignmentOrder returns the AssignmentOrder field value if set, zero value otherwise.
func (o *PoolAbstractPoolAllOf) GetAssignmentOrder() string {
	if o == nil || o.AssignmentOrder == nil {
		var ret string
		return ret
	}
	return *o.AssignmentOrder
}

// GetAssignmentOrderOk returns a tuple with the AssignmentOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractPoolAllOf) GetAssignmentOrderOk() (*string, bool) {
	if o == nil || o.AssignmentOrder == nil {
		return nil, false
	}
	return o.AssignmentOrder, true
}

// HasAssignmentOrder returns a boolean if a field has been set.
func (o *PoolAbstractPoolAllOf) HasAssignmentOrder() bool {
	if o != nil && o.AssignmentOrder != nil {
		return true
	}

	return false
}

// SetAssignmentOrder gets a reference to the given string and assigns it to the AssignmentOrder field.
func (o *PoolAbstractPoolAllOf) SetAssignmentOrder(v string) {
	o.AssignmentOrder = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *PoolAbstractPoolAllOf) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractPoolAllOf) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *PoolAbstractPoolAllOf) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *PoolAbstractPoolAllOf) SetSize(v int64) {
	o.Size = &v
}

func (o PoolAbstractPoolAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Assigned != nil {
		toSerialize["Assigned"] = o.Assigned
	}
	if o.AssignmentOrder != nil {
		toSerialize["AssignmentOrder"] = o.AssignmentOrder
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PoolAbstractPoolAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPoolAbstractPoolAllOf := _PoolAbstractPoolAllOf{}

	if err = json.Unmarshal(bytes, &varPoolAbstractPoolAllOf); err == nil {
		*o = PoolAbstractPoolAllOf(varPoolAbstractPoolAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Assigned")
		delete(additionalProperties, "AssignmentOrder")
		delete(additionalProperties, "Size")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePoolAbstractPoolAllOf struct {
	value *PoolAbstractPoolAllOf
	isSet bool
}

func (v NullablePoolAbstractPoolAllOf) Get() *PoolAbstractPoolAllOf {
	return v.value
}

func (v *NullablePoolAbstractPoolAllOf) Set(val *PoolAbstractPoolAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolAbstractPoolAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolAbstractPoolAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolAbstractPoolAllOf(val *PoolAbstractPoolAllOf) *NullablePoolAbstractPoolAllOf {
	return &NullablePoolAbstractPoolAllOf{value: val, isSet: true}
}

func (v NullablePoolAbstractPoolAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolAbstractPoolAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
