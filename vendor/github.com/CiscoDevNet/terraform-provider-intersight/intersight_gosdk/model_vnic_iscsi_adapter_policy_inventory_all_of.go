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

// VnicIscsiAdapterPolicyInventoryAllOf Definition of the list of properties defined in 'vnic.IscsiAdapterPolicyInventory', excluding properties defined in parent classes.
type VnicIscsiAdapterPolicyInventoryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of seconds to wait until Cisco UCS assumes that the initial login has failed and the iSCSI adapter is unavailable.
	ConnectionTimeOut *int64 `json:"ConnectionTimeOut,omitempty"`
	// The number of seconds to wait before the initiator assumes that the DHCP server is unavailable.
	DhcpTimeout *int64 `json:"DhcpTimeout,omitempty"`
	// The number of times to retry the connection in case of a failure during iSCSI LUN discovery.
	LunBusyRetryCount    *int64                `json:"LunBusyRetryCount,omitempty"`
	TargetMo             *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicIscsiAdapterPolicyInventoryAllOf VnicIscsiAdapterPolicyInventoryAllOf

// NewVnicIscsiAdapterPolicyInventoryAllOf instantiates a new VnicIscsiAdapterPolicyInventoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicIscsiAdapterPolicyInventoryAllOf(classId string, objectType string) *VnicIscsiAdapterPolicyInventoryAllOf {
	this := VnicIscsiAdapterPolicyInventoryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVnicIscsiAdapterPolicyInventoryAllOfWithDefaults instantiates a new VnicIscsiAdapterPolicyInventoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicIscsiAdapterPolicyInventoryAllOfWithDefaults() *VnicIscsiAdapterPolicyInventoryAllOf {
	this := VnicIscsiAdapterPolicyInventoryAllOf{}
	var classId string = "vnic.IscsiAdapterPolicyInventory"
	this.ClassId = classId
	var objectType string = "vnic.IscsiAdapterPolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicIscsiAdapterPolicyInventoryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicIscsiAdapterPolicyInventoryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicIscsiAdapterPolicyInventoryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicIscsiAdapterPolicyInventoryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicIscsiAdapterPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicIscsiAdapterPolicyInventoryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConnectionTimeOut returns the ConnectionTimeOut field value if set, zero value otherwise.
func (o *VnicIscsiAdapterPolicyInventoryAllOf) GetConnectionTimeOut() int64 {
	if o == nil || o.ConnectionTimeOut == nil {
		var ret int64
		return ret
	}
	return *o.ConnectionTimeOut
}

// GetConnectionTimeOutOk returns a tuple with the ConnectionTimeOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiAdapterPolicyInventoryAllOf) GetConnectionTimeOutOk() (*int64, bool) {
	if o == nil || o.ConnectionTimeOut == nil {
		return nil, false
	}
	return o.ConnectionTimeOut, true
}

// HasConnectionTimeOut returns a boolean if a field has been set.
func (o *VnicIscsiAdapterPolicyInventoryAllOf) HasConnectionTimeOut() bool {
	if o != nil && o.ConnectionTimeOut != nil {
		return true
	}

	return false
}

// SetConnectionTimeOut gets a reference to the given int64 and assigns it to the ConnectionTimeOut field.
func (o *VnicIscsiAdapterPolicyInventoryAllOf) SetConnectionTimeOut(v int64) {
	o.ConnectionTimeOut = &v
}

// GetDhcpTimeout returns the DhcpTimeout field value if set, zero value otherwise.
func (o *VnicIscsiAdapterPolicyInventoryAllOf) GetDhcpTimeout() int64 {
	if o == nil || o.DhcpTimeout == nil {
		var ret int64
		return ret
	}
	return *o.DhcpTimeout
}

// GetDhcpTimeoutOk returns a tuple with the DhcpTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiAdapterPolicyInventoryAllOf) GetDhcpTimeoutOk() (*int64, bool) {
	if o == nil || o.DhcpTimeout == nil {
		return nil, false
	}
	return o.DhcpTimeout, true
}

// HasDhcpTimeout returns a boolean if a field has been set.
func (o *VnicIscsiAdapterPolicyInventoryAllOf) HasDhcpTimeout() bool {
	if o != nil && o.DhcpTimeout != nil {
		return true
	}

	return false
}

// SetDhcpTimeout gets a reference to the given int64 and assigns it to the DhcpTimeout field.
func (o *VnicIscsiAdapterPolicyInventoryAllOf) SetDhcpTimeout(v int64) {
	o.DhcpTimeout = &v
}

// GetLunBusyRetryCount returns the LunBusyRetryCount field value if set, zero value otherwise.
func (o *VnicIscsiAdapterPolicyInventoryAllOf) GetLunBusyRetryCount() int64 {
	if o == nil || o.LunBusyRetryCount == nil {
		var ret int64
		return ret
	}
	return *o.LunBusyRetryCount
}

// GetLunBusyRetryCountOk returns a tuple with the LunBusyRetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiAdapterPolicyInventoryAllOf) GetLunBusyRetryCountOk() (*int64, bool) {
	if o == nil || o.LunBusyRetryCount == nil {
		return nil, false
	}
	return o.LunBusyRetryCount, true
}

// HasLunBusyRetryCount returns a boolean if a field has been set.
func (o *VnicIscsiAdapterPolicyInventoryAllOf) HasLunBusyRetryCount() bool {
	if o != nil && o.LunBusyRetryCount != nil {
		return true
	}

	return false
}

// SetLunBusyRetryCount gets a reference to the given int64 and assigns it to the LunBusyRetryCount field.
func (o *VnicIscsiAdapterPolicyInventoryAllOf) SetLunBusyRetryCount(v int64) {
	o.LunBusyRetryCount = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *VnicIscsiAdapterPolicyInventoryAllOf) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiAdapterPolicyInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *VnicIscsiAdapterPolicyInventoryAllOf) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *VnicIscsiAdapterPolicyInventoryAllOf) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o VnicIscsiAdapterPolicyInventoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConnectionTimeOut != nil {
		toSerialize["ConnectionTimeOut"] = o.ConnectionTimeOut
	}
	if o.DhcpTimeout != nil {
		toSerialize["DhcpTimeout"] = o.DhcpTimeout
	}
	if o.LunBusyRetryCount != nil {
		toSerialize["LunBusyRetryCount"] = o.LunBusyRetryCount
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicIscsiAdapterPolicyInventoryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicIscsiAdapterPolicyInventoryAllOf := _VnicIscsiAdapterPolicyInventoryAllOf{}

	if err = json.Unmarshal(bytes, &varVnicIscsiAdapterPolicyInventoryAllOf); err == nil {
		*o = VnicIscsiAdapterPolicyInventoryAllOf(varVnicIscsiAdapterPolicyInventoryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConnectionTimeOut")
		delete(additionalProperties, "DhcpTimeout")
		delete(additionalProperties, "LunBusyRetryCount")
		delete(additionalProperties, "TargetMo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicIscsiAdapterPolicyInventoryAllOf struct {
	value *VnicIscsiAdapterPolicyInventoryAllOf
	isSet bool
}

func (v NullableVnicIscsiAdapterPolicyInventoryAllOf) Get() *VnicIscsiAdapterPolicyInventoryAllOf {
	return v.value
}

func (v *NullableVnicIscsiAdapterPolicyInventoryAllOf) Set(val *VnicIscsiAdapterPolicyInventoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicIscsiAdapterPolicyInventoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicIscsiAdapterPolicyInventoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicIscsiAdapterPolicyInventoryAllOf(val *VnicIscsiAdapterPolicyInventoryAllOf) *NullableVnicIscsiAdapterPolicyInventoryAllOf {
	return &NullableVnicIscsiAdapterPolicyInventoryAllOf{value: val, isSet: true}
}

func (v NullableVnicIscsiAdapterPolicyInventoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicIscsiAdapterPolicyInventoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
