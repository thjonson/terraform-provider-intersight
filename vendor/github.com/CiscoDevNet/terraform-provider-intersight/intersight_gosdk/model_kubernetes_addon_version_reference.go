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

// KubernetesAddonVersionReference A reference to an Addon type and version that can be looked up in the Addon definition catalog.
type KubernetesAddonVersionReference struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the addon to lookup.
	Name *string `json:"Name,omitempty"`
	// Version number to filter the addon with.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesAddonVersionReference KubernetesAddonVersionReference

// NewKubernetesAddonVersionReference instantiates a new KubernetesAddonVersionReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAddonVersionReference(classId string, objectType string) *KubernetesAddonVersionReference {
	this := KubernetesAddonVersionReference{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesAddonVersionReferenceWithDefaults instantiates a new KubernetesAddonVersionReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAddonVersionReferenceWithDefaults() *KubernetesAddonVersionReference {
	this := KubernetesAddonVersionReference{}
	var classId string = "kubernetes.AddonVersionReference"
	this.ClassId = classId
	var objectType string = "kubernetes.AddonVersionReference"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesAddonVersionReference) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesAddonVersionReference) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesAddonVersionReference) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesAddonVersionReference) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesAddonVersionReference) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesAddonVersionReference) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesAddonVersionReference) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonVersionReference) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesAddonVersionReference) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesAddonVersionReference) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *KubernetesAddonVersionReference) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonVersionReference) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *KubernetesAddonVersionReference) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *KubernetesAddonVersionReference) SetVersion(v string) {
	o.Version = &v
}

func (o KubernetesAddonVersionReference) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesAddonVersionReference) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesAddonVersionReferenceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the addon to lookup.
		Name *string `json:"Name,omitempty"`
		// Version number to filter the addon with.
		Version *string `json:"Version,omitempty"`
	}

	varKubernetesAddonVersionReferenceWithoutEmbeddedStruct := KubernetesAddonVersionReferenceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesAddonVersionReferenceWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesAddonVersionReference := _KubernetesAddonVersionReference{}
		varKubernetesAddonVersionReference.ClassId = varKubernetesAddonVersionReferenceWithoutEmbeddedStruct.ClassId
		varKubernetesAddonVersionReference.ObjectType = varKubernetesAddonVersionReferenceWithoutEmbeddedStruct.ObjectType
		varKubernetesAddonVersionReference.Name = varKubernetesAddonVersionReferenceWithoutEmbeddedStruct.Name
		varKubernetesAddonVersionReference.Version = varKubernetesAddonVersionReferenceWithoutEmbeddedStruct.Version
		*o = KubernetesAddonVersionReference(varKubernetesAddonVersionReference)
	} else {
		return err
	}

	varKubernetesAddonVersionReference := _KubernetesAddonVersionReference{}

	err = json.Unmarshal(bytes, &varKubernetesAddonVersionReference)
	if err == nil {
		o.MoBaseComplexType = varKubernetesAddonVersionReference.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Version")

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

type NullableKubernetesAddonVersionReference struct {
	value *KubernetesAddonVersionReference
	isSet bool
}

func (v NullableKubernetesAddonVersionReference) Get() *KubernetesAddonVersionReference {
	return v.value
}

func (v *NullableKubernetesAddonVersionReference) Set(val *KubernetesAddonVersionReference) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAddonVersionReference) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAddonVersionReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAddonVersionReference(val *KubernetesAddonVersionReference) *NullableKubernetesAddonVersionReference {
	return &NullableKubernetesAddonVersionReference{value: val, isSet: true}
}

func (v NullableKubernetesAddonVersionReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAddonVersionReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
