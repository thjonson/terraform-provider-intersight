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

// TelemetryDruidCountAggregator Computes the count of Druid rows that match the filters. The count aggregator counts the number of Druid rows, which does not always reflect the number of raw events ingested. This is because Druid can be configured to roll up data at ingestion time To count the number of ingested rows of data, include a count aggregator at ingestion time, and a longSum aggregator at query time.
type TelemetryDruidCountAggregator struct {
	// The aggregator type.
	Type string `json:"type"`
	// the output name
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidCountAggregator TelemetryDruidCountAggregator

// NewTelemetryDruidCountAggregator instantiates a new TelemetryDruidCountAggregator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidCountAggregator(type_ string) *TelemetryDruidCountAggregator {
	this := TelemetryDruidCountAggregator{}
	this.Type = type_
	return &this
}

// NewTelemetryDruidCountAggregatorWithDefaults instantiates a new TelemetryDruidCountAggregator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidCountAggregatorWithDefaults() *TelemetryDruidCountAggregator {
	this := TelemetryDruidCountAggregator{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidCountAggregator) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidCountAggregator) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidCountAggregator) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TelemetryDruidCountAggregator) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidCountAggregator) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TelemetryDruidCountAggregator) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TelemetryDruidCountAggregator) SetName(v string) {
	o.Name = &v
}

func (o TelemetryDruidCountAggregator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidCountAggregator) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidCountAggregator := _TelemetryDruidCountAggregator{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidCountAggregator); err == nil {
		*o = TelemetryDruidCountAggregator(varTelemetryDruidCountAggregator)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidCountAggregator struct {
	value *TelemetryDruidCountAggregator
	isSet bool
}

func (v NullableTelemetryDruidCountAggregator) Get() *TelemetryDruidCountAggregator {
	return v.value
}

func (v *NullableTelemetryDruidCountAggregator) Set(val *TelemetryDruidCountAggregator) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidCountAggregator) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidCountAggregator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidCountAggregator(val *TelemetryDruidCountAggregator) *NullableTelemetryDruidCountAggregator {
	return &NullableTelemetryDruidCountAggregator{value: val, isSet: true}
}

func (v NullableTelemetryDruidCountAggregator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidCountAggregator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
