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
	"fmt"
)

// KubernetesVersionPolicyRelationship - A relationship to the 'kubernetes.VersionPolicy' resource, or the expanded 'kubernetes.VersionPolicy' resource, or the 'null' value.
type KubernetesVersionPolicyRelationship struct {
	KubernetesVersionPolicy *KubernetesVersionPolicy
	MoMoRef                 *MoMoRef
}

// KubernetesVersionPolicyAsKubernetesVersionPolicyRelationship is a convenience function that returns KubernetesVersionPolicy wrapped in KubernetesVersionPolicyRelationship
func KubernetesVersionPolicyAsKubernetesVersionPolicyRelationship(v *KubernetesVersionPolicy) KubernetesVersionPolicyRelationship {
	return KubernetesVersionPolicyRelationship{
		KubernetesVersionPolicy: v,
	}
}

// MoMoRefAsKubernetesVersionPolicyRelationship is a convenience function that returns MoMoRef wrapped in KubernetesVersionPolicyRelationship
func MoMoRefAsKubernetesVersionPolicyRelationship(v *MoMoRef) KubernetesVersionPolicyRelationship {
	return KubernetesVersionPolicyRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *KubernetesVersionPolicyRelationship) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'kubernetes.VersionPolicy'
	if jsonDict["ClassId"] == "kubernetes.VersionPolicy" {
		// try to unmarshal JSON data into KubernetesVersionPolicy
		err = json.Unmarshal(data, &dst.KubernetesVersionPolicy)
		if err == nil {
			return nil // data stored in dst.KubernetesVersionPolicy, return on the first match
		} else {
			dst.KubernetesVersionPolicy = nil
			return fmt.Errorf("Failed to unmarshal KubernetesVersionPolicyRelationship as KubernetesVersionPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.MoRef'
	if jsonDict["ClassId"] == "mo.MoRef" {
		// try to unmarshal JSON data into MoMoRef
		err = json.Unmarshal(data, &dst.MoMoRef)
		if err == nil {
			return nil // data stored in dst.MoMoRef, return on the first match
		} else {
			dst.MoMoRef = nil
			return fmt.Errorf("Failed to unmarshal KubernetesVersionPolicyRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src KubernetesVersionPolicyRelationship) MarshalJSON() ([]byte, error) {
	if src.KubernetesVersionPolicy != nil {
		return json.Marshal(&src.KubernetesVersionPolicy)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *KubernetesVersionPolicyRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.KubernetesVersionPolicy != nil {
		return obj.KubernetesVersionPolicy
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableKubernetesVersionPolicyRelationship struct {
	value *KubernetesVersionPolicyRelationship
	isSet bool
}

func (v NullableKubernetesVersionPolicyRelationship) Get() *KubernetesVersionPolicyRelationship {
	return v.value
}

func (v *NullableKubernetesVersionPolicyRelationship) Set(val *KubernetesVersionPolicyRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesVersionPolicyRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesVersionPolicyRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesVersionPolicyRelationship(val *KubernetesVersionPolicyRelationship) *NullableKubernetesVersionPolicyRelationship {
	return &NullableKubernetesVersionPolicyRelationship{value: val, isSet: true}
}

func (v NullableKubernetesVersionPolicyRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesVersionPolicyRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
