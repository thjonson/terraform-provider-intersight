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

// SoftwareUcsdDistributableAllOf Definition of the list of properties defined in 'software.UcsdDistributable', excluding properties defined in parent classes.
type SoftwareUcsdDistributableAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                 `json:"ObjectType"`
	Catalog              *SoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwareUcsdDistributableAllOf SoftwareUcsdDistributableAllOf

// NewSoftwareUcsdDistributableAllOf instantiates a new SoftwareUcsdDistributableAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareUcsdDistributableAllOf(classId string, objectType string) *SoftwareUcsdDistributableAllOf {
	this := SoftwareUcsdDistributableAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSoftwareUcsdDistributableAllOfWithDefaults instantiates a new SoftwareUcsdDistributableAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareUcsdDistributableAllOfWithDefaults() *SoftwareUcsdDistributableAllOf {
	this := SoftwareUcsdDistributableAllOf{}
	var classId string = "software.UcsdDistributable"
	this.ClassId = classId
	var objectType string = "software.UcsdDistributable"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwareUcsdDistributableAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwareUcsdDistributableAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwareUcsdDistributableAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwareUcsdDistributableAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwareUcsdDistributableAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwareUcsdDistributableAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *SoftwareUcsdDistributableAllOf) GetCatalog() SoftwarerepositoryCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret SoftwarerepositoryCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareUcsdDistributableAllOf) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *SoftwareUcsdDistributableAllOf) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given SoftwarerepositoryCatalogRelationship and assigns it to the Catalog field.
func (o *SoftwareUcsdDistributableAllOf) SetCatalog(v SoftwarerepositoryCatalogRelationship) {
	o.Catalog = &v
}

func (o SoftwareUcsdDistributableAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwareUcsdDistributableAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSoftwareUcsdDistributableAllOf := _SoftwareUcsdDistributableAllOf{}

	if err = json.Unmarshal(bytes, &varSoftwareUcsdDistributableAllOf); err == nil {
		*o = SoftwareUcsdDistributableAllOf(varSoftwareUcsdDistributableAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Catalog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSoftwareUcsdDistributableAllOf struct {
	value *SoftwareUcsdDistributableAllOf
	isSet bool
}

func (v NullableSoftwareUcsdDistributableAllOf) Get() *SoftwareUcsdDistributableAllOf {
	return v.value
}

func (v *NullableSoftwareUcsdDistributableAllOf) Set(val *SoftwareUcsdDistributableAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareUcsdDistributableAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareUcsdDistributableAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareUcsdDistributableAllOf(val *SoftwareUcsdDistributableAllOf) *NullableSoftwareUcsdDistributableAllOf {
	return &NullableSoftwareUcsdDistributableAllOf{value: val, isSet: true}
}

func (v NullableSoftwareUcsdDistributableAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareUcsdDistributableAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
