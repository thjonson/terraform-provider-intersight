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
)

// HyperflexHxSiteDtAllOf Definition of the list of properties defined in 'hyperflex.HxSiteDt', excluding properties defined in parent classes.
type HyperflexHxSiteDtAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the site for this HyperFlex cluster.
	Name                 *string                       `json:"Name,omitempty"`
	Zone                 NullableHyperflexHxZoneInfoDt `json:"Zone,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxSiteDtAllOf HyperflexHxSiteDtAllOf

// NewHyperflexHxSiteDtAllOf instantiates a new HyperflexHxSiteDtAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxSiteDtAllOf(classId string, objectType string) *HyperflexHxSiteDtAllOf {
	this := HyperflexHxSiteDtAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHxSiteDtAllOfWithDefaults instantiates a new HyperflexHxSiteDtAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxSiteDtAllOfWithDefaults() *HyperflexHxSiteDtAllOf {
	this := HyperflexHxSiteDtAllOf{}
	var classId string = "hyperflex.HxSiteDt"
	this.ClassId = classId
	var objectType string = "hyperflex.HxSiteDt"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxSiteDtAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxSiteDtAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxSiteDtAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxSiteDtAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxSiteDtAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxSiteDtAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexHxSiteDtAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxSiteDtAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexHxSiteDtAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexHxSiteDtAllOf) SetName(v string) {
	o.Name = &v
}

// GetZone returns the Zone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHxSiteDtAllOf) GetZone() HyperflexHxZoneInfoDt {
	if o == nil || o.Zone.Get() == nil {
		var ret HyperflexHxZoneInfoDt
		return ret
	}
	return *o.Zone.Get()
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHxSiteDtAllOf) GetZoneOk() (*HyperflexHxZoneInfoDt, bool) {
	if o == nil {
		return nil, false
	}
	return o.Zone.Get(), o.Zone.IsSet()
}

// HasZone returns a boolean if a field has been set.
func (o *HyperflexHxSiteDtAllOf) HasZone() bool {
	if o != nil && o.Zone.IsSet() {
		return true
	}

	return false
}

// SetZone gets a reference to the given NullableHyperflexHxZoneInfoDt and assigns it to the Zone field.
func (o *HyperflexHxSiteDtAllOf) SetZone(v HyperflexHxZoneInfoDt) {
	o.Zone.Set(&v)
}

// SetZoneNil sets the value for Zone to be an explicit nil
func (o *HyperflexHxSiteDtAllOf) SetZoneNil() {
	o.Zone.Set(nil)
}

// UnsetZone ensures that no value is present for Zone, not even an explicit nil
func (o *HyperflexHxSiteDtAllOf) UnsetZone() {
	o.Zone.Unset()
}

func (o HyperflexHxSiteDtAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Zone.IsSet() {
		toSerialize["Zone"] = o.Zone.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxSiteDtAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHxSiteDtAllOf := _HyperflexHxSiteDtAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexHxSiteDtAllOf); err == nil {
		*o = HyperflexHxSiteDtAllOf(varHyperflexHxSiteDtAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Zone")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHxSiteDtAllOf struct {
	value *HyperflexHxSiteDtAllOf
	isSet bool
}

func (v NullableHyperflexHxSiteDtAllOf) Get() *HyperflexHxSiteDtAllOf {
	return v.value
}

func (v *NullableHyperflexHxSiteDtAllOf) Set(val *HyperflexHxSiteDtAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxSiteDtAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxSiteDtAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxSiteDtAllOf(val *HyperflexHxSiteDtAllOf) *NullableHyperflexHxSiteDtAllOf {
	return &NullableHyperflexHxSiteDtAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHxSiteDtAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxSiteDtAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
