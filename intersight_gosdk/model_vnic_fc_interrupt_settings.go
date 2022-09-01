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

// VnicFcInterruptSettings Interrupt Settings for the virtual fibre channel interface.
type VnicFcInterruptSettings struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The preferred driver interrupt mode. This can be one of the following:- MSIx - Message Signaled Interrupts (MSI) with the optional extension. MSI  - MSI only. INTx - PCI INTx interrupts. MSIx is the recommended option. * `MSIx` - Message Signaled Interrupt (MSI) mechanism with the optional extension (MSIx). MSIx is the recommended and default option. * `MSI` - Message Signaled Interrupt (MSI) mechanism that treats messages as interrupts. * `INTx` - Line-based interrupt (INTx) mechanism similar to the one used in Legacy systems.
	Mode                 *string `json:"Mode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicFcInterruptSettings VnicFcInterruptSettings

// NewVnicFcInterruptSettings instantiates a new VnicFcInterruptSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicFcInterruptSettings(classId string, objectType string) *VnicFcInterruptSettings {
	this := VnicFcInterruptSettings{}
	this.ClassId = classId
	this.ObjectType = objectType
	var mode string = "MSIx"
	this.Mode = &mode
	return &this
}

// NewVnicFcInterruptSettingsWithDefaults instantiates a new VnicFcInterruptSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicFcInterruptSettingsWithDefaults() *VnicFcInterruptSettings {
	this := VnicFcInterruptSettings{}
	var classId string = "vnic.FcInterruptSettings"
	this.ClassId = classId
	var objectType string = "vnic.FcInterruptSettings"
	this.ObjectType = objectType
	var mode string = "MSIx"
	this.Mode = &mode
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicFcInterruptSettings) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicFcInterruptSettings) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicFcInterruptSettings) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicFcInterruptSettings) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicFcInterruptSettings) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicFcInterruptSettings) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *VnicFcInterruptSettings) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcInterruptSettings) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *VnicFcInterruptSettings) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *VnicFcInterruptSettings) SetMode(v string) {
	o.Mode = &v
}

func (o VnicFcInterruptSettings) MarshalJSON() ([]byte, error) {
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
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicFcInterruptSettings) UnmarshalJSON(bytes []byte) (err error) {
	type VnicFcInterruptSettingsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The preferred driver interrupt mode. This can be one of the following:- MSIx - Message Signaled Interrupts (MSI) with the optional extension. MSI  - MSI only. INTx - PCI INTx interrupts. MSIx is the recommended option. * `MSIx` - Message Signaled Interrupt (MSI) mechanism with the optional extension (MSIx). MSIx is the recommended and default option. * `MSI` - Message Signaled Interrupt (MSI) mechanism that treats messages as interrupts. * `INTx` - Line-based interrupt (INTx) mechanism similar to the one used in Legacy systems.
		Mode *string `json:"Mode,omitempty"`
	}

	varVnicFcInterruptSettingsWithoutEmbeddedStruct := VnicFcInterruptSettingsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicFcInterruptSettingsWithoutEmbeddedStruct)
	if err == nil {
		varVnicFcInterruptSettings := _VnicFcInterruptSettings{}
		varVnicFcInterruptSettings.ClassId = varVnicFcInterruptSettingsWithoutEmbeddedStruct.ClassId
		varVnicFcInterruptSettings.ObjectType = varVnicFcInterruptSettingsWithoutEmbeddedStruct.ObjectType
		varVnicFcInterruptSettings.Mode = varVnicFcInterruptSettingsWithoutEmbeddedStruct.Mode
		*o = VnicFcInterruptSettings(varVnicFcInterruptSettings)
	} else {
		return err
	}

	varVnicFcInterruptSettings := _VnicFcInterruptSettings{}

	err = json.Unmarshal(bytes, &varVnicFcInterruptSettings)
	if err == nil {
		o.MoBaseComplexType = varVnicFcInterruptSettings.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Mode")

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

type NullableVnicFcInterruptSettings struct {
	value *VnicFcInterruptSettings
	isSet bool
}

func (v NullableVnicFcInterruptSettings) Get() *VnicFcInterruptSettings {
	return v.value
}

func (v *NullableVnicFcInterruptSettings) Set(val *VnicFcInterruptSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicFcInterruptSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicFcInterruptSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicFcInterruptSettings(val *VnicFcInterruptSettings) *NullableVnicFcInterruptSettings {
	return &NullableVnicFcInterruptSettings{value: val, isSet: true}
}

func (v NullableVnicFcInterruptSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicFcInterruptSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
