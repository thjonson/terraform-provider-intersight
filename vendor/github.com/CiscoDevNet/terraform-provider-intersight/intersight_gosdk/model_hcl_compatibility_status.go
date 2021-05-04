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

// HclCompatibilityStatus Check the compatibility status for the given hardware and software configurations.
type HclCompatibilityStatus struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType  string                            `json:"ObjectType"`
	ProfileList []HclHardwareCompatibilityProfile `json:"ProfileList,omitempty"`
	// Type of the request to be served. * `FillSupportedVersions` - Responds with the supported firmware and driver versions. The API doesn't expect firmware and driver versions to be passed in the request and ignores if passed. * `CheckCompatibility` - Checks the compatibility for the given firmware and driver versions. This request type expects the firmware and driver versions to filled and the service validates the values and responds back with the error codes. * `GetRecommendedDrivers` - Responds with the supported drivers. The API expects firmware version to be filled. The API populates driver ISO url for the given server model. Today the link is same for all servers managed by UCSM whereas it depends on the model for Standalone servers.
	RequestType          *string `json:"RequestType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HclCompatibilityStatus HclCompatibilityStatus

// NewHclCompatibilityStatus instantiates a new HclCompatibilityStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclCompatibilityStatus(classId string, objectType string) *HclCompatibilityStatus {
	this := HclCompatibilityStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	var requestType string = "FillSupportedVersions"
	this.RequestType = &requestType
	return &this
}

// NewHclCompatibilityStatusWithDefaults instantiates a new HclCompatibilityStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclCompatibilityStatusWithDefaults() *HclCompatibilityStatus {
	this := HclCompatibilityStatus{}
	var classId string = "hcl.CompatibilityStatus"
	this.ClassId = classId
	var objectType string = "hcl.CompatibilityStatus"
	this.ObjectType = objectType
	var requestType string = "FillSupportedVersions"
	this.RequestType = &requestType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HclCompatibilityStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HclCompatibilityStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HclCompatibilityStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HclCompatibilityStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HclCompatibilityStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HclCompatibilityStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetProfileList returns the ProfileList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HclCompatibilityStatus) GetProfileList() []HclHardwareCompatibilityProfile {
	if o == nil {
		var ret []HclHardwareCompatibilityProfile
		return ret
	}
	return o.ProfileList
}

// GetProfileListOk returns a tuple with the ProfileList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HclCompatibilityStatus) GetProfileListOk() (*[]HclHardwareCompatibilityProfile, bool) {
	if o == nil || o.ProfileList == nil {
		return nil, false
	}
	return &o.ProfileList, true
}

// HasProfileList returns a boolean if a field has been set.
func (o *HclCompatibilityStatus) HasProfileList() bool {
	if o != nil && o.ProfileList != nil {
		return true
	}

	return false
}

// SetProfileList gets a reference to the given []HclHardwareCompatibilityProfile and assigns it to the ProfileList field.
func (o *HclCompatibilityStatus) SetProfileList(v []HclHardwareCompatibilityProfile) {
	o.ProfileList = v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *HclCompatibilityStatus) GetRequestType() string {
	if o == nil || o.RequestType == nil {
		var ret string
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclCompatibilityStatus) GetRequestTypeOk() (*string, bool) {
	if o == nil || o.RequestType == nil {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *HclCompatibilityStatus) HasRequestType() bool {
	if o != nil && o.RequestType != nil {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given string and assigns it to the RequestType field.
func (o *HclCompatibilityStatus) SetRequestType(v string) {
	o.RequestType = &v
}

func (o HclCompatibilityStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ProfileList != nil {
		toSerialize["ProfileList"] = o.ProfileList
	}
	if o.RequestType != nil {
		toSerialize["RequestType"] = o.RequestType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HclCompatibilityStatus) UnmarshalJSON(bytes []byte) (err error) {
	type HclCompatibilityStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType  string                            `json:"ObjectType"`
		ProfileList []HclHardwareCompatibilityProfile `json:"ProfileList,omitempty"`
		// Type of the request to be served. * `FillSupportedVersions` - Responds with the supported firmware and driver versions. The API doesn't expect firmware and driver versions to be passed in the request and ignores if passed. * `CheckCompatibility` - Checks the compatibility for the given firmware and driver versions. This request type expects the firmware and driver versions to filled and the service validates the values and responds back with the error codes. * `GetRecommendedDrivers` - Responds with the supported drivers. The API expects firmware version to be filled. The API populates driver ISO url for the given server model. Today the link is same for all servers managed by UCSM whereas it depends on the model for Standalone servers.
		RequestType *string `json:"RequestType,omitempty"`
	}

	varHclCompatibilityStatusWithoutEmbeddedStruct := HclCompatibilityStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHclCompatibilityStatusWithoutEmbeddedStruct)
	if err == nil {
		varHclCompatibilityStatus := _HclCompatibilityStatus{}
		varHclCompatibilityStatus.ClassId = varHclCompatibilityStatusWithoutEmbeddedStruct.ClassId
		varHclCompatibilityStatus.ObjectType = varHclCompatibilityStatusWithoutEmbeddedStruct.ObjectType
		varHclCompatibilityStatus.ProfileList = varHclCompatibilityStatusWithoutEmbeddedStruct.ProfileList
		varHclCompatibilityStatus.RequestType = varHclCompatibilityStatusWithoutEmbeddedStruct.RequestType
		*o = HclCompatibilityStatus(varHclCompatibilityStatus)
	} else {
		return err
	}

	varHclCompatibilityStatus := _HclCompatibilityStatus{}

	err = json.Unmarshal(bytes, &varHclCompatibilityStatus)
	if err == nil {
		o.MoBaseMo = varHclCompatibilityStatus.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ProfileList")
		delete(additionalProperties, "RequestType")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableHclCompatibilityStatus struct {
	value *HclCompatibilityStatus
	isSet bool
}

func (v NullableHclCompatibilityStatus) Get() *HclCompatibilityStatus {
	return v.value
}

func (v *NullableHclCompatibilityStatus) Set(val *HclCompatibilityStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableHclCompatibilityStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableHclCompatibilityStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclCompatibilityStatus(val *HclCompatibilityStatus) *NullableHclCompatibilityStatus {
	return &NullableHclCompatibilityStatus{value: val, isSet: true}
}

func (v NullableHclCompatibilityStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclCompatibilityStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
