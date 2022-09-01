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

// FirmwareComponentDescriptorAllOf Definition of the list of properties defined in 'firmware.ComponentDescriptor', excluding properties defined in parent classes.
type FirmwareComponentDescriptorAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The brand string of the endpoint for which this capability information is applicable.
	BrandString *string `json:"BrandString,omitempty"`
	// The label type for the component.
	Label *string `json:"Label,omitempty"`
	// The revision for the component.
	Revision             *string `json:"Revision,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareComponentDescriptorAllOf FirmwareComponentDescriptorAllOf

// NewFirmwareComponentDescriptorAllOf instantiates a new FirmwareComponentDescriptorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareComponentDescriptorAllOf(classId string, objectType string) *FirmwareComponentDescriptorAllOf {
	this := FirmwareComponentDescriptorAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwareComponentDescriptorAllOfWithDefaults instantiates a new FirmwareComponentDescriptorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareComponentDescriptorAllOfWithDefaults() *FirmwareComponentDescriptorAllOf {
	this := FirmwareComponentDescriptorAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareComponentDescriptorAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareComponentDescriptorAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareComponentDescriptorAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareComponentDescriptorAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareComponentDescriptorAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareComponentDescriptorAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBrandString returns the BrandString field value if set, zero value otherwise.
func (o *FirmwareComponentDescriptorAllOf) GetBrandString() string {
	if o == nil || o.BrandString == nil {
		var ret string
		return ret
	}
	return *o.BrandString
}

// GetBrandStringOk returns a tuple with the BrandString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentDescriptorAllOf) GetBrandStringOk() (*string, bool) {
	if o == nil || o.BrandString == nil {
		return nil, false
	}
	return o.BrandString, true
}

// HasBrandString returns a boolean if a field has been set.
func (o *FirmwareComponentDescriptorAllOf) HasBrandString() bool {
	if o != nil && o.BrandString != nil {
		return true
	}

	return false
}

// SetBrandString gets a reference to the given string and assigns it to the BrandString field.
func (o *FirmwareComponentDescriptorAllOf) SetBrandString(v string) {
	o.BrandString = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FirmwareComponentDescriptorAllOf) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentDescriptorAllOf) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FirmwareComponentDescriptorAllOf) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *FirmwareComponentDescriptorAllOf) SetLabel(v string) {
	o.Label = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *FirmwareComponentDescriptorAllOf) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentDescriptorAllOf) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *FirmwareComponentDescriptorAllOf) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *FirmwareComponentDescriptorAllOf) SetRevision(v string) {
	o.Revision = &v
}

func (o FirmwareComponentDescriptorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BrandString != nil {
		toSerialize["BrandString"] = o.BrandString
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Revision != nil {
		toSerialize["Revision"] = o.Revision
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareComponentDescriptorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFirmwareComponentDescriptorAllOf := _FirmwareComponentDescriptorAllOf{}

	if err = json.Unmarshal(bytes, &varFirmwareComponentDescriptorAllOf); err == nil {
		*o = FirmwareComponentDescriptorAllOf(varFirmwareComponentDescriptorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BrandString")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Revision")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareComponentDescriptorAllOf struct {
	value *FirmwareComponentDescriptorAllOf
	isSet bool
}

func (v NullableFirmwareComponentDescriptorAllOf) Get() *FirmwareComponentDescriptorAllOf {
	return v.value
}

func (v *NullableFirmwareComponentDescriptorAllOf) Set(val *FirmwareComponentDescriptorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareComponentDescriptorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareComponentDescriptorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareComponentDescriptorAllOf(val *FirmwareComponentDescriptorAllOf) *NullableFirmwareComponentDescriptorAllOf {
	return &NullableFirmwareComponentDescriptorAllOf{value: val, isSet: true}
}

func (v NullableFirmwareComponentDescriptorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareComponentDescriptorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
