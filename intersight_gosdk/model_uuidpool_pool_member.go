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

// UuidpoolPoolMember PoolMember represents a single UUID that is part of a pool.
type UuidpoolPoolMember struct {
	PoolAbstractPoolMember
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// UUID Prefix+Suffix of this PoolMember.
	Uuid                 *string                        `json:"Uuid,omitempty"`
	AssignedToEntity     *MoBaseMoRelationship          `json:"AssignedToEntity,omitempty"`
	BlockHead            *UuidpoolBlockRelationship     `json:"BlockHead,omitempty"`
	Peer                 *UuidpoolUuidLeaseRelationship `json:"Peer,omitempty"`
	Pool                 *UuidpoolPoolRelationship      `json:"Pool,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UuidpoolPoolMember UuidpoolPoolMember

// NewUuidpoolPoolMember instantiates a new UuidpoolPoolMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUuidpoolPoolMember(classId string, objectType string) *UuidpoolPoolMember {
	this := UuidpoolPoolMember{}
	this.ClassId = classId
	this.ObjectType = objectType
	var assigned bool = false
	this.Assigned = &assigned
	return &this
}

// NewUuidpoolPoolMemberWithDefaults instantiates a new UuidpoolPoolMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUuidpoolPoolMemberWithDefaults() *UuidpoolPoolMember {
	this := UuidpoolPoolMember{}
	var classId string = "uuidpool.PoolMember"
	this.ClassId = classId
	var objectType string = "uuidpool.PoolMember"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *UuidpoolPoolMember) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *UuidpoolPoolMember) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *UuidpoolPoolMember) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *UuidpoolPoolMember) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *UuidpoolPoolMember) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *UuidpoolPoolMember) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *UuidpoolPoolMember) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolPoolMember) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *UuidpoolPoolMember) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *UuidpoolPoolMember) SetUuid(v string) {
	o.Uuid = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise.
func (o *UuidpoolPoolMember) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || o.AssignedToEntity == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolPoolMember) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AssignedToEntity == nil {
		return nil, false
	}
	return o.AssignedToEntity, true
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *UuidpoolPoolMember) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity != nil {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given MoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *UuidpoolPoolMember) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity = &v
}

// GetBlockHead returns the BlockHead field value if set, zero value otherwise.
func (o *UuidpoolPoolMember) GetBlockHead() UuidpoolBlockRelationship {
	if o == nil || o.BlockHead == nil {
		var ret UuidpoolBlockRelationship
		return ret
	}
	return *o.BlockHead
}

// GetBlockHeadOk returns a tuple with the BlockHead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolPoolMember) GetBlockHeadOk() (*UuidpoolBlockRelationship, bool) {
	if o == nil || o.BlockHead == nil {
		return nil, false
	}
	return o.BlockHead, true
}

// HasBlockHead returns a boolean if a field has been set.
func (o *UuidpoolPoolMember) HasBlockHead() bool {
	if o != nil && o.BlockHead != nil {
		return true
	}

	return false
}

// SetBlockHead gets a reference to the given UuidpoolBlockRelationship and assigns it to the BlockHead field.
func (o *UuidpoolPoolMember) SetBlockHead(v UuidpoolBlockRelationship) {
	o.BlockHead = &v
}

// GetPeer returns the Peer field value if set, zero value otherwise.
func (o *UuidpoolPoolMember) GetPeer() UuidpoolUuidLeaseRelationship {
	if o == nil || o.Peer == nil {
		var ret UuidpoolUuidLeaseRelationship
		return ret
	}
	return *o.Peer
}

// GetPeerOk returns a tuple with the Peer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolPoolMember) GetPeerOk() (*UuidpoolUuidLeaseRelationship, bool) {
	if o == nil || o.Peer == nil {
		return nil, false
	}
	return o.Peer, true
}

// HasPeer returns a boolean if a field has been set.
func (o *UuidpoolPoolMember) HasPeer() bool {
	if o != nil && o.Peer != nil {
		return true
	}

	return false
}

// SetPeer gets a reference to the given UuidpoolUuidLeaseRelationship and assigns it to the Peer field.
func (o *UuidpoolPoolMember) SetPeer(v UuidpoolUuidLeaseRelationship) {
	o.Peer = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *UuidpoolPoolMember) GetPool() UuidpoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret UuidpoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolPoolMember) GetPoolOk() (*UuidpoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *UuidpoolPoolMember) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given UuidpoolPoolRelationship and assigns it to the Pool field.
func (o *UuidpoolPoolMember) SetPool(v UuidpoolPoolRelationship) {
	o.Pool = &v
}

func (o UuidpoolPoolMember) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractPoolMember, errPoolAbstractPoolMember := json.Marshal(o.PoolAbstractPoolMember)
	if errPoolAbstractPoolMember != nil {
		return []byte{}, errPoolAbstractPoolMember
	}
	errPoolAbstractPoolMember = json.Unmarshal([]byte(serializedPoolAbstractPoolMember), &toSerialize)
	if errPoolAbstractPoolMember != nil {
		return []byte{}, errPoolAbstractPoolMember
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.AssignedToEntity != nil {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity
	}
	if o.BlockHead != nil {
		toSerialize["BlockHead"] = o.BlockHead
	}
	if o.Peer != nil {
		toSerialize["Peer"] = o.Peer
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UuidpoolPoolMember) UnmarshalJSON(bytes []byte) (err error) {
	type UuidpoolPoolMemberWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// UUID Prefix+Suffix of this PoolMember.
		Uuid             *string                        `json:"Uuid,omitempty"`
		AssignedToEntity *MoBaseMoRelationship          `json:"AssignedToEntity,omitempty"`
		BlockHead        *UuidpoolBlockRelationship     `json:"BlockHead,omitempty"`
		Peer             *UuidpoolUuidLeaseRelationship `json:"Peer,omitempty"`
		Pool             *UuidpoolPoolRelationship      `json:"Pool,omitempty"`
	}

	varUuidpoolPoolMemberWithoutEmbeddedStruct := UuidpoolPoolMemberWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varUuidpoolPoolMemberWithoutEmbeddedStruct)
	if err == nil {
		varUuidpoolPoolMember := _UuidpoolPoolMember{}
		varUuidpoolPoolMember.ClassId = varUuidpoolPoolMemberWithoutEmbeddedStruct.ClassId
		varUuidpoolPoolMember.ObjectType = varUuidpoolPoolMemberWithoutEmbeddedStruct.ObjectType
		varUuidpoolPoolMember.Uuid = varUuidpoolPoolMemberWithoutEmbeddedStruct.Uuid
		varUuidpoolPoolMember.AssignedToEntity = varUuidpoolPoolMemberWithoutEmbeddedStruct.AssignedToEntity
		varUuidpoolPoolMember.BlockHead = varUuidpoolPoolMemberWithoutEmbeddedStruct.BlockHead
		varUuidpoolPoolMember.Peer = varUuidpoolPoolMemberWithoutEmbeddedStruct.Peer
		varUuidpoolPoolMember.Pool = varUuidpoolPoolMemberWithoutEmbeddedStruct.Pool
		*o = UuidpoolPoolMember(varUuidpoolPoolMember)
	} else {
		return err
	}

	varUuidpoolPoolMember := _UuidpoolPoolMember{}

	err = json.Unmarshal(bytes, &varUuidpoolPoolMember)
	if err == nil {
		o.PoolAbstractPoolMember = varUuidpoolPoolMember.PoolAbstractPoolMember
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "AssignedToEntity")
		delete(additionalProperties, "BlockHead")
		delete(additionalProperties, "Peer")
		delete(additionalProperties, "Pool")

		// remove fields from embedded structs
		reflectPoolAbstractPoolMember := reflect.ValueOf(o.PoolAbstractPoolMember)
		for i := 0; i < reflectPoolAbstractPoolMember.Type().NumField(); i++ {
			t := reflectPoolAbstractPoolMember.Type().Field(i)

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

type NullableUuidpoolPoolMember struct {
	value *UuidpoolPoolMember
	isSet bool
}

func (v NullableUuidpoolPoolMember) Get() *UuidpoolPoolMember {
	return v.value
}

func (v *NullableUuidpoolPoolMember) Set(val *UuidpoolPoolMember) {
	v.value = val
	v.isSet = true
}

func (v NullableUuidpoolPoolMember) IsSet() bool {
	return v.isSet
}

func (v *NullableUuidpoolPoolMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUuidpoolPoolMember(val *UuidpoolPoolMember) *NullableUuidpoolPoolMember {
	return &NullableUuidpoolPoolMember{value: val, isSet: true}
}

func (v NullableUuidpoolPoolMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUuidpoolPoolMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
