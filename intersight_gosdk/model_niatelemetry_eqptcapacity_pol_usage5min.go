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

// NiatelemetryEqptcapacityPolUsage5min Aci node tcam utilization information.
type NiatelemetryEqptcapacityPolUsage5min struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// PolUsageBase information for aci nodes.
	PolUsageBase *string `json:"PolUsageBase,omitempty"`
	// PolUsageCapCum information for aci nodes.
	PolUsageCapCum *string `json:"PolUsageCapCum,omitempty"`
	// PolUsageCum information for aci nodes.
	PolUsageCum          *string `json:"PolUsageCum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryEqptcapacityPolUsage5min NiatelemetryEqptcapacityPolUsage5min

// NewNiatelemetryEqptcapacityPolUsage5min instantiates a new NiatelemetryEqptcapacityPolUsage5min object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryEqptcapacityPolUsage5min(classId string, objectType string) *NiatelemetryEqptcapacityPolUsage5min {
	this := NiatelemetryEqptcapacityPolUsage5min{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryEqptcapacityPolUsage5minWithDefaults instantiates a new NiatelemetryEqptcapacityPolUsage5min object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryEqptcapacityPolUsage5minWithDefaults() *NiatelemetryEqptcapacityPolUsage5min {
	this := NiatelemetryEqptcapacityPolUsage5min{}
	var classId string = "niatelemetry.EqptcapacityPolUsage5min"
	this.ClassId = classId
	var objectType string = "niatelemetry.EqptcapacityPolUsage5min"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryEqptcapacityPolUsage5min) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryEqptcapacityPolUsage5min) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryEqptcapacityPolUsage5min) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryEqptcapacityPolUsage5min) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryEqptcapacityPolUsage5min) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryEqptcapacityPolUsage5min) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPolUsageBase returns the PolUsageBase field value if set, zero value otherwise.
func (o *NiatelemetryEqptcapacityPolUsage5min) GetPolUsageBase() string {
	if o == nil || o.PolUsageBase == nil {
		var ret string
		return ret
	}
	return *o.PolUsageBase
}

// GetPolUsageBaseOk returns a tuple with the PolUsageBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEqptcapacityPolUsage5min) GetPolUsageBaseOk() (*string, bool) {
	if o == nil || o.PolUsageBase == nil {
		return nil, false
	}
	return o.PolUsageBase, true
}

// HasPolUsageBase returns a boolean if a field has been set.
func (o *NiatelemetryEqptcapacityPolUsage5min) HasPolUsageBase() bool {
	if o != nil && o.PolUsageBase != nil {
		return true
	}

	return false
}

// SetPolUsageBase gets a reference to the given string and assigns it to the PolUsageBase field.
func (o *NiatelemetryEqptcapacityPolUsage5min) SetPolUsageBase(v string) {
	o.PolUsageBase = &v
}

// GetPolUsageCapCum returns the PolUsageCapCum field value if set, zero value otherwise.
func (o *NiatelemetryEqptcapacityPolUsage5min) GetPolUsageCapCum() string {
	if o == nil || o.PolUsageCapCum == nil {
		var ret string
		return ret
	}
	return *o.PolUsageCapCum
}

// GetPolUsageCapCumOk returns a tuple with the PolUsageCapCum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEqptcapacityPolUsage5min) GetPolUsageCapCumOk() (*string, bool) {
	if o == nil || o.PolUsageCapCum == nil {
		return nil, false
	}
	return o.PolUsageCapCum, true
}

// HasPolUsageCapCum returns a boolean if a field has been set.
func (o *NiatelemetryEqptcapacityPolUsage5min) HasPolUsageCapCum() bool {
	if o != nil && o.PolUsageCapCum != nil {
		return true
	}

	return false
}

// SetPolUsageCapCum gets a reference to the given string and assigns it to the PolUsageCapCum field.
func (o *NiatelemetryEqptcapacityPolUsage5min) SetPolUsageCapCum(v string) {
	o.PolUsageCapCum = &v
}

// GetPolUsageCum returns the PolUsageCum field value if set, zero value otherwise.
func (o *NiatelemetryEqptcapacityPolUsage5min) GetPolUsageCum() string {
	if o == nil || o.PolUsageCum == nil {
		var ret string
		return ret
	}
	return *o.PolUsageCum
}

// GetPolUsageCumOk returns a tuple with the PolUsageCum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEqptcapacityPolUsage5min) GetPolUsageCumOk() (*string, bool) {
	if o == nil || o.PolUsageCum == nil {
		return nil, false
	}
	return o.PolUsageCum, true
}

// HasPolUsageCum returns a boolean if a field has been set.
func (o *NiatelemetryEqptcapacityPolUsage5min) HasPolUsageCum() bool {
	if o != nil && o.PolUsageCum != nil {
		return true
	}

	return false
}

// SetPolUsageCum gets a reference to the given string and assigns it to the PolUsageCum field.
func (o *NiatelemetryEqptcapacityPolUsage5min) SetPolUsageCum(v string) {
	o.PolUsageCum = &v
}

func (o NiatelemetryEqptcapacityPolUsage5min) MarshalJSON() ([]byte, error) {
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
	if o.PolUsageBase != nil {
		toSerialize["PolUsageBase"] = o.PolUsageBase
	}
	if o.PolUsageCapCum != nil {
		toSerialize["PolUsageCapCum"] = o.PolUsageCapCum
	}
	if o.PolUsageCum != nil {
		toSerialize["PolUsageCum"] = o.PolUsageCum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryEqptcapacityPolUsage5min) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryEqptcapacityPolUsage5minWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// PolUsageBase information for aci nodes.
		PolUsageBase *string `json:"PolUsageBase,omitempty"`
		// PolUsageCapCum information for aci nodes.
		PolUsageCapCum *string `json:"PolUsageCapCum,omitempty"`
		// PolUsageCum information for aci nodes.
		PolUsageCum *string `json:"PolUsageCum,omitempty"`
	}

	varNiatelemetryEqptcapacityPolUsage5minWithoutEmbeddedStruct := NiatelemetryEqptcapacityPolUsage5minWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryEqptcapacityPolUsage5minWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryEqptcapacityPolUsage5min := _NiatelemetryEqptcapacityPolUsage5min{}
		varNiatelemetryEqptcapacityPolUsage5min.ClassId = varNiatelemetryEqptcapacityPolUsage5minWithoutEmbeddedStruct.ClassId
		varNiatelemetryEqptcapacityPolUsage5min.ObjectType = varNiatelemetryEqptcapacityPolUsage5minWithoutEmbeddedStruct.ObjectType
		varNiatelemetryEqptcapacityPolUsage5min.PolUsageBase = varNiatelemetryEqptcapacityPolUsage5minWithoutEmbeddedStruct.PolUsageBase
		varNiatelemetryEqptcapacityPolUsage5min.PolUsageCapCum = varNiatelemetryEqptcapacityPolUsage5minWithoutEmbeddedStruct.PolUsageCapCum
		varNiatelemetryEqptcapacityPolUsage5min.PolUsageCum = varNiatelemetryEqptcapacityPolUsage5minWithoutEmbeddedStruct.PolUsageCum
		*o = NiatelemetryEqptcapacityPolUsage5min(varNiatelemetryEqptcapacityPolUsage5min)
	} else {
		return err
	}

	varNiatelemetryEqptcapacityPolUsage5min := _NiatelemetryEqptcapacityPolUsage5min{}

	err = json.Unmarshal(bytes, &varNiatelemetryEqptcapacityPolUsage5min)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetryEqptcapacityPolUsage5min.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PolUsageBase")
		delete(additionalProperties, "PolUsageCapCum")
		delete(additionalProperties, "PolUsageCum")

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

type NullableNiatelemetryEqptcapacityPolUsage5min struct {
	value *NiatelemetryEqptcapacityPolUsage5min
	isSet bool
}

func (v NullableNiatelemetryEqptcapacityPolUsage5min) Get() *NiatelemetryEqptcapacityPolUsage5min {
	return v.value
}

func (v *NullableNiatelemetryEqptcapacityPolUsage5min) Set(val *NiatelemetryEqptcapacityPolUsage5min) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryEqptcapacityPolUsage5min) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryEqptcapacityPolUsage5min) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryEqptcapacityPolUsage5min(val *NiatelemetryEqptcapacityPolUsage5min) *NullableNiatelemetryEqptcapacityPolUsage5min {
	return &NullableNiatelemetryEqptcapacityPolUsage5min{value: val, isSet: true}
}

func (v NullableNiatelemetryEqptcapacityPolUsage5min) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryEqptcapacityPolUsage5min) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
