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

// CertificatemanagementPolicyInventory Certificate Management policy models a reusable certificate and private key configuration that can be applied to multiple servers via profile association.
type CertificatemanagementPolicyInventory struct {
	PolicyAbstractPolicyInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                 `json:"ObjectType"`
	Certificates         []CertificatemanagementCertificateBase `json:"Certificates,omitempty"`
	TargetMo             *MoBaseMoRelationship                  `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificatemanagementPolicyInventory CertificatemanagementPolicyInventory

// NewCertificatemanagementPolicyInventory instantiates a new CertificatemanagementPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificatemanagementPolicyInventory(classId string, objectType string) *CertificatemanagementPolicyInventory {
	this := CertificatemanagementPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCertificatemanagementPolicyInventoryWithDefaults instantiates a new CertificatemanagementPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificatemanagementPolicyInventoryWithDefaults() *CertificatemanagementPolicyInventory {
	this := CertificatemanagementPolicyInventory{}
	var classId string = "certificatemanagement.PolicyInventory"
	this.ClassId = classId
	var objectType string = "certificatemanagement.PolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CertificatemanagementPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CertificatemanagementPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CertificatemanagementPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CertificatemanagementPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CertificatemanagementPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CertificatemanagementPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCertificates returns the Certificates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificatemanagementPolicyInventory) GetCertificates() []CertificatemanagementCertificateBase {
	if o == nil {
		var ret []CertificatemanagementCertificateBase
		return ret
	}
	return o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificatemanagementPolicyInventory) GetCertificatesOk() ([]CertificatemanagementCertificateBase, bool) {
	if o == nil || o.Certificates == nil {
		return nil, false
	}
	return o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *CertificatemanagementPolicyInventory) HasCertificates() bool {
	if o != nil && o.Certificates != nil {
		return true
	}

	return false
}

// SetCertificates gets a reference to the given []CertificatemanagementCertificateBase and assigns it to the Certificates field.
func (o *CertificatemanagementPolicyInventory) SetCertificates(v []CertificatemanagementCertificateBase) {
	o.Certificates = v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *CertificatemanagementPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificatemanagementPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *CertificatemanagementPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *CertificatemanagementPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o CertificatemanagementPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicyInventory, errPolicyAbstractPolicyInventory := json.Marshal(o.PolicyAbstractPolicyInventory)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
	errPolicyAbstractPolicyInventory = json.Unmarshal([]byte(serializedPolicyAbstractPolicyInventory), &toSerialize)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Certificates != nil {
		toSerialize["Certificates"] = o.Certificates
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CertificatemanagementPolicyInventory) UnmarshalJSON(bytes []byte) (err error) {
	type CertificatemanagementPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string                                 `json:"ObjectType"`
		Certificates []CertificatemanagementCertificateBase `json:"Certificates,omitempty"`
		TargetMo     *MoBaseMoRelationship                  `json:"TargetMo,omitempty"`
	}

	varCertificatemanagementPolicyInventoryWithoutEmbeddedStruct := CertificatemanagementPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCertificatemanagementPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varCertificatemanagementPolicyInventory := _CertificatemanagementPolicyInventory{}
		varCertificatemanagementPolicyInventory.ClassId = varCertificatemanagementPolicyInventoryWithoutEmbeddedStruct.ClassId
		varCertificatemanagementPolicyInventory.ObjectType = varCertificatemanagementPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varCertificatemanagementPolicyInventory.Certificates = varCertificatemanagementPolicyInventoryWithoutEmbeddedStruct.Certificates
		varCertificatemanagementPolicyInventory.TargetMo = varCertificatemanagementPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = CertificatemanagementPolicyInventory(varCertificatemanagementPolicyInventory)
	} else {
		return err
	}

	varCertificatemanagementPolicyInventory := _CertificatemanagementPolicyInventory{}

	err = json.Unmarshal(bytes, &varCertificatemanagementPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varCertificatemanagementPolicyInventory.PolicyAbstractPolicyInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Certificates")
		delete(additionalProperties, "TargetMo")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicyInventory := reflect.ValueOf(o.PolicyAbstractPolicyInventory)
		for i := 0; i < reflectPolicyAbstractPolicyInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicyInventory.Type().Field(i)

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

type NullableCertificatemanagementPolicyInventory struct {
	value *CertificatemanagementPolicyInventory
	isSet bool
}

func (v NullableCertificatemanagementPolicyInventory) Get() *CertificatemanagementPolicyInventory {
	return v.value
}

func (v *NullableCertificatemanagementPolicyInventory) Set(val *CertificatemanagementPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificatemanagementPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificatemanagementPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificatemanagementPolicyInventory(val *CertificatemanagementPolicyInventory) *NullableCertificatemanagementPolicyInventory {
	return &NullableCertificatemanagementPolicyInventory{value: val, isSet: true}
}

func (v NullableCertificatemanagementPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificatemanagementPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
