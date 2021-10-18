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

// KubernetesEthernet Configuration to apply to a Physical Network Interface Card.
type KubernetesEthernet struct {
	KubernetesNetworkInterface
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                            `json:"ObjectType"`
	Matcher              NullableKubernetesEthernetMatcher `json:"Matcher,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesEthernet KubernetesEthernet

// NewKubernetesEthernet instantiates a new KubernetesEthernet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesEthernet(classId string, objectType string) *KubernetesEthernet {
	this := KubernetesEthernet{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesEthernetWithDefaults instantiates a new KubernetesEthernet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesEthernetWithDefaults() *KubernetesEthernet {
	this := KubernetesEthernet{}
	var classId string = "kubernetes.Ethernet"
	this.ClassId = classId
	var objectType string = "kubernetes.Ethernet"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesEthernet) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesEthernet) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesEthernet) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesEthernet) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesEthernet) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesEthernet) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMatcher returns the Matcher field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesEthernet) GetMatcher() KubernetesEthernetMatcher {
	if o == nil || o.Matcher.Get() == nil {
		var ret KubernetesEthernetMatcher
		return ret
	}
	return *o.Matcher.Get()
}

// GetMatcherOk returns a tuple with the Matcher field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesEthernet) GetMatcherOk() (*KubernetesEthernetMatcher, bool) {
	if o == nil {
		return nil, false
	}
	return o.Matcher.Get(), o.Matcher.IsSet()
}

// HasMatcher returns a boolean if a field has been set.
func (o *KubernetesEthernet) HasMatcher() bool {
	if o != nil && o.Matcher.IsSet() {
		return true
	}

	return false
}

// SetMatcher gets a reference to the given NullableKubernetesEthernetMatcher and assigns it to the Matcher field.
func (o *KubernetesEthernet) SetMatcher(v KubernetesEthernetMatcher) {
	o.Matcher.Set(&v)
}

// SetMatcherNil sets the value for Matcher to be an explicit nil
func (o *KubernetesEthernet) SetMatcherNil() {
	o.Matcher.Set(nil)
}

// UnsetMatcher ensures that no value is present for Matcher, not even an explicit nil
func (o *KubernetesEthernet) UnsetMatcher() {
	o.Matcher.Unset()
}

func (o KubernetesEthernet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKubernetesNetworkInterface, errKubernetesNetworkInterface := json.Marshal(o.KubernetesNetworkInterface)
	if errKubernetesNetworkInterface != nil {
		return []byte{}, errKubernetesNetworkInterface
	}
	errKubernetesNetworkInterface = json.Unmarshal([]byte(serializedKubernetesNetworkInterface), &toSerialize)
	if errKubernetesNetworkInterface != nil {
		return []byte{}, errKubernetesNetworkInterface
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Matcher.IsSet() {
		toSerialize["Matcher"] = o.Matcher.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesEthernet) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesEthernetWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                            `json:"ObjectType"`
		Matcher    NullableKubernetesEthernetMatcher `json:"Matcher,omitempty"`
	}

	varKubernetesEthernetWithoutEmbeddedStruct := KubernetesEthernetWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesEthernetWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesEthernet := _KubernetesEthernet{}
		varKubernetesEthernet.ClassId = varKubernetesEthernetWithoutEmbeddedStruct.ClassId
		varKubernetesEthernet.ObjectType = varKubernetesEthernetWithoutEmbeddedStruct.ObjectType
		varKubernetesEthernet.Matcher = varKubernetesEthernetWithoutEmbeddedStruct.Matcher
		*o = KubernetesEthernet(varKubernetesEthernet)
	} else {
		return err
	}

	varKubernetesEthernet := _KubernetesEthernet{}

	err = json.Unmarshal(bytes, &varKubernetesEthernet)
	if err == nil {
		o.KubernetesNetworkInterface = varKubernetesEthernet.KubernetesNetworkInterface
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Matcher")

		// remove fields from embedded structs
		reflectKubernetesNetworkInterface := reflect.ValueOf(o.KubernetesNetworkInterface)
		for i := 0; i < reflectKubernetesNetworkInterface.Type().NumField(); i++ {
			t := reflectKubernetesNetworkInterface.Type().Field(i)

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

type NullableKubernetesEthernet struct {
	value *KubernetesEthernet
	isSet bool
}

func (v NullableKubernetesEthernet) Get() *KubernetesEthernet {
	return v.value
}

func (v *NullableKubernetesEthernet) Set(val *KubernetesEthernet) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesEthernet) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesEthernet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesEthernet(val *KubernetesEthernet) *NullableKubernetesEthernet {
	return &NullableKubernetesEthernet{value: val, isSet: true}
}

func (v NullableKubernetesEthernet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesEthernet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
