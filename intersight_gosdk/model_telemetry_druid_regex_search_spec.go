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
)

// TelemetryDruidRegexSearchSpec The specification for a Druid 'insensitive_contains' search
type TelemetryDruidRegexSearchSpec struct {
	// null
	Type string `json:"type"`
	// The regular expression to match.  If any part of a dimension value contains the pattern specified in this search query a \"match\" occurs.
	Regex                string `json:"regex"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidRegexSearchSpec TelemetryDruidRegexSearchSpec

// NewTelemetryDruidRegexSearchSpec instantiates a new TelemetryDruidRegexSearchSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidRegexSearchSpec(type_ string, regex string) *TelemetryDruidRegexSearchSpec {
	this := TelemetryDruidRegexSearchSpec{}
	this.Type = type_
	this.Regex = regex
	return &this
}

// NewTelemetryDruidRegexSearchSpecWithDefaults instantiates a new TelemetryDruidRegexSearchSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidRegexSearchSpecWithDefaults() *TelemetryDruidRegexSearchSpec {
	this := TelemetryDruidRegexSearchSpec{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidRegexSearchSpec) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidRegexSearchSpec) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidRegexSearchSpec) SetType(v string) {
	o.Type = v
}

// GetRegex returns the Regex field value
func (o *TelemetryDruidRegexSearchSpec) GetRegex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Regex
}

// GetRegexOk returns a tuple with the Regex field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidRegexSearchSpec) GetRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Regex, true
}

// SetRegex sets field value
func (o *TelemetryDruidRegexSearchSpec) SetRegex(v string) {
	o.Regex = v
}

func (o TelemetryDruidRegexSearchSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["regex"] = o.Regex
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidRegexSearchSpec) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidRegexSearchSpec := _TelemetryDruidRegexSearchSpec{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidRegexSearchSpec); err == nil {
		*o = TelemetryDruidRegexSearchSpec(varTelemetryDruidRegexSearchSpec)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "regex")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidRegexSearchSpec struct {
	value *TelemetryDruidRegexSearchSpec
	isSet bool
}

func (v NullableTelemetryDruidRegexSearchSpec) Get() *TelemetryDruidRegexSearchSpec {
	return v.value
}

func (v *NullableTelemetryDruidRegexSearchSpec) Set(val *TelemetryDruidRegexSearchSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidRegexSearchSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidRegexSearchSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidRegexSearchSpec(val *TelemetryDruidRegexSearchSpec) *NullableTelemetryDruidRegexSearchSpec {
	return &NullableTelemetryDruidRegexSearchSpec{value: val, isSet: true}
}

func (v NullableTelemetryDruidRegexSearchSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidRegexSearchSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
