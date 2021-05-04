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
	"fmt"
)

// HclHyperflexSoftwareCompatibilityInfoRelationship - A relationship to the 'hcl.HyperflexSoftwareCompatibilityInfo' resource, or the expanded 'hcl.HyperflexSoftwareCompatibilityInfo' resource, or the 'null' value.
type HclHyperflexSoftwareCompatibilityInfoRelationship struct {
	HclHyperflexSoftwareCompatibilityInfo *HclHyperflexSoftwareCompatibilityInfo
	MoMoRef                               *MoMoRef
}

// HclHyperflexSoftwareCompatibilityInfoAsHclHyperflexSoftwareCompatibilityInfoRelationship is a convenience function that returns HclHyperflexSoftwareCompatibilityInfo wrapped in HclHyperflexSoftwareCompatibilityInfoRelationship
func HclHyperflexSoftwareCompatibilityInfoAsHclHyperflexSoftwareCompatibilityInfoRelationship(v *HclHyperflexSoftwareCompatibilityInfo) HclHyperflexSoftwareCompatibilityInfoRelationship {
	return HclHyperflexSoftwareCompatibilityInfoRelationship{HclHyperflexSoftwareCompatibilityInfo: v}
}

// MoMoRefAsHclHyperflexSoftwareCompatibilityInfoRelationship is a convenience function that returns MoMoRef wrapped in HclHyperflexSoftwareCompatibilityInfoRelationship
func MoMoRefAsHclHyperflexSoftwareCompatibilityInfoRelationship(v *MoMoRef) HclHyperflexSoftwareCompatibilityInfoRelationship {
	return HclHyperflexSoftwareCompatibilityInfoRelationship{MoMoRef: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *HclHyperflexSoftwareCompatibilityInfoRelationship) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'hcl.HyperflexSoftwareCompatibilityInfo'
	if jsonDict["ClassId"] == "hcl.HyperflexSoftwareCompatibilityInfo" {
		// try to unmarshal JSON data into HclHyperflexSoftwareCompatibilityInfo
		err = json.Unmarshal(data, &dst.HclHyperflexSoftwareCompatibilityInfo)
		if err == nil {
			return nil // data stored in dst.HclHyperflexSoftwareCompatibilityInfo, return on the first match
		} else {
			dst.HclHyperflexSoftwareCompatibilityInfo = nil
			return fmt.Errorf("Failed to unmarshal HclHyperflexSoftwareCompatibilityInfoRelationship as HclHyperflexSoftwareCompatibilityInfo: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal HclHyperflexSoftwareCompatibilityInfoRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HclHyperflexSoftwareCompatibilityInfoRelationship) MarshalJSON() ([]byte, error) {
	if src.HclHyperflexSoftwareCompatibilityInfo != nil {
		return json.Marshal(&src.HclHyperflexSoftwareCompatibilityInfo)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HclHyperflexSoftwareCompatibilityInfoRelationship) GetActualInstance() interface{} {
	if obj.HclHyperflexSoftwareCompatibilityInfo != nil {
		return obj.HclHyperflexSoftwareCompatibilityInfo
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableHclHyperflexSoftwareCompatibilityInfoRelationship struct {
	value *HclHyperflexSoftwareCompatibilityInfoRelationship
	isSet bool
}

func (v NullableHclHyperflexSoftwareCompatibilityInfoRelationship) Get() *HclHyperflexSoftwareCompatibilityInfoRelationship {
	return v.value
}

func (v *NullableHclHyperflexSoftwareCompatibilityInfoRelationship) Set(val *HclHyperflexSoftwareCompatibilityInfoRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableHclHyperflexSoftwareCompatibilityInfoRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableHclHyperflexSoftwareCompatibilityInfoRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclHyperflexSoftwareCompatibilityInfoRelationship(val *HclHyperflexSoftwareCompatibilityInfoRelationship) *NullableHclHyperflexSoftwareCompatibilityInfoRelationship {
	return &NullableHclHyperflexSoftwareCompatibilityInfoRelationship{value: val, isSet: true}
}

func (v NullableHclHyperflexSoftwareCompatibilityInfoRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclHyperflexSoftwareCompatibilityInfoRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
