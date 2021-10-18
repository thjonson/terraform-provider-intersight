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

// IamPrivateKeySpec Parameters used to generate a private key.
type IamPrivateKeySpec struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                             `json:"ObjectType"`
	Algorithm            NullablePkixKeyGenerationSpec      `json:"Algorithm,omitempty"`
	CertificateRequest   *IamCertificateRequestRelationship `json:"CertificateRequest,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamPrivateKeySpec IamPrivateKeySpec

// NewIamPrivateKeySpec instantiates a new IamPrivateKeySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamPrivateKeySpec(classId string, objectType string) *IamPrivateKeySpec {
	this := IamPrivateKeySpec{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamPrivateKeySpecWithDefaults instantiates a new IamPrivateKeySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamPrivateKeySpecWithDefaults() *IamPrivateKeySpec {
	this := IamPrivateKeySpec{}
	var classId string = "iam.PrivateKeySpec"
	this.ClassId = classId
	var objectType string = "iam.PrivateKeySpec"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamPrivateKeySpec) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamPrivateKeySpec) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamPrivateKeySpec) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamPrivateKeySpec) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamPrivateKeySpec) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamPrivateKeySpec) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPrivateKeySpec) GetAlgorithm() PkixKeyGenerationSpec {
	if o == nil || o.Algorithm.Get() == nil {
		var ret PkixKeyGenerationSpec
		return ret
	}
	return *o.Algorithm.Get()
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPrivateKeySpec) GetAlgorithmOk() (*PkixKeyGenerationSpec, bool) {
	if o == nil {
		return nil, false
	}
	return o.Algorithm.Get(), o.Algorithm.IsSet()
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *IamPrivateKeySpec) HasAlgorithm() bool {
	if o != nil && o.Algorithm.IsSet() {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given NullablePkixKeyGenerationSpec and assigns it to the Algorithm field.
func (o *IamPrivateKeySpec) SetAlgorithm(v PkixKeyGenerationSpec) {
	o.Algorithm.Set(&v)
}

// SetAlgorithmNil sets the value for Algorithm to be an explicit nil
func (o *IamPrivateKeySpec) SetAlgorithmNil() {
	o.Algorithm.Set(nil)
}

// UnsetAlgorithm ensures that no value is present for Algorithm, not even an explicit nil
func (o *IamPrivateKeySpec) UnsetAlgorithm() {
	o.Algorithm.Unset()
}

// GetCertificateRequest returns the CertificateRequest field value if set, zero value otherwise.
func (o *IamPrivateKeySpec) GetCertificateRequest() IamCertificateRequestRelationship {
	if o == nil || o.CertificateRequest == nil {
		var ret IamCertificateRequestRelationship
		return ret
	}
	return *o.CertificateRequest
}

// GetCertificateRequestOk returns a tuple with the CertificateRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPrivateKeySpec) GetCertificateRequestOk() (*IamCertificateRequestRelationship, bool) {
	if o == nil || o.CertificateRequest == nil {
		return nil, false
	}
	return o.CertificateRequest, true
}

// HasCertificateRequest returns a boolean if a field has been set.
func (o *IamPrivateKeySpec) HasCertificateRequest() bool {
	if o != nil && o.CertificateRequest != nil {
		return true
	}

	return false
}

// SetCertificateRequest gets a reference to the given IamCertificateRequestRelationship and assigns it to the CertificateRequest field.
func (o *IamPrivateKeySpec) SetCertificateRequest(v IamCertificateRequestRelationship) {
	o.CertificateRequest = &v
}

func (o IamPrivateKeySpec) MarshalJSON() ([]byte, error) {
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
	if o.Algorithm.IsSet() {
		toSerialize["Algorithm"] = o.Algorithm.Get()
	}
	if o.CertificateRequest != nil {
		toSerialize["CertificateRequest"] = o.CertificateRequest
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamPrivateKeySpec) UnmarshalJSON(bytes []byte) (err error) {
	type IamPrivateKeySpecWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType         string                             `json:"ObjectType"`
		Algorithm          NullablePkixKeyGenerationSpec      `json:"Algorithm,omitempty"`
		CertificateRequest *IamCertificateRequestRelationship `json:"CertificateRequest,omitempty"`
	}

	varIamPrivateKeySpecWithoutEmbeddedStruct := IamPrivateKeySpecWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamPrivateKeySpecWithoutEmbeddedStruct)
	if err == nil {
		varIamPrivateKeySpec := _IamPrivateKeySpec{}
		varIamPrivateKeySpec.ClassId = varIamPrivateKeySpecWithoutEmbeddedStruct.ClassId
		varIamPrivateKeySpec.ObjectType = varIamPrivateKeySpecWithoutEmbeddedStruct.ObjectType
		varIamPrivateKeySpec.Algorithm = varIamPrivateKeySpecWithoutEmbeddedStruct.Algorithm
		varIamPrivateKeySpec.CertificateRequest = varIamPrivateKeySpecWithoutEmbeddedStruct.CertificateRequest
		*o = IamPrivateKeySpec(varIamPrivateKeySpec)
	} else {
		return err
	}

	varIamPrivateKeySpec := _IamPrivateKeySpec{}

	err = json.Unmarshal(bytes, &varIamPrivateKeySpec)
	if err == nil {
		o.MoBaseMo = varIamPrivateKeySpec.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Algorithm")
		delete(additionalProperties, "CertificateRequest")

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

type NullableIamPrivateKeySpec struct {
	value *IamPrivateKeySpec
	isSet bool
}

func (v NullableIamPrivateKeySpec) Get() *IamPrivateKeySpec {
	return v.value
}

func (v *NullableIamPrivateKeySpec) Set(val *IamPrivateKeySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableIamPrivateKeySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableIamPrivateKeySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamPrivateKeySpec(val *IamPrivateKeySpec) *NullableIamPrivateKeySpec {
	return &NullableIamPrivateKeySpec{value: val, isSet: true}
}

func (v NullableIamPrivateKeySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamPrivateKeySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
