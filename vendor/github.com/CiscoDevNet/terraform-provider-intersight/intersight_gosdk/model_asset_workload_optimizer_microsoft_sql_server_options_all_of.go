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
)

// AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf Definition of the list of properties defined in 'asset.WorkloadOptimizerMicrosoftSqlServerOptions', excluding properties defined in parent classes.
type AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Port that Microsoft SQL Server Browser listens for incoming requests for SQL Server resources and provides information about SQL Server instances that are installed on the computer. When this port is specified, Database will be communicated through the Browser Service with this port instead of default SQLServer port.
	BrowserServicePort *int64 `json:"BrowserServicePort,omitempty"`
	// Active Directory domain, if required for this account.
	FullDomainName       *string `json:"FullDomainName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf

// NewAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf instantiates a new AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf(classId string, objectType string) *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf {
	this := AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOfWithDefaults instantiates a new AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOfWithDefaults() *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf {
	this := AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf{}
	var classId string = "asset.WorkloadOptimizerMicrosoftSqlServerOptions"
	this.ClassId = classId
	var objectType string = "asset.WorkloadOptimizerMicrosoftSqlServerOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBrowserServicePort returns the BrowserServicePort field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) GetBrowserServicePort() int64 {
	if o == nil || o.BrowserServicePort == nil {
		var ret int64
		return ret
	}
	return *o.BrowserServicePort
}

// GetBrowserServicePortOk returns a tuple with the BrowserServicePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) GetBrowserServicePortOk() (*int64, bool) {
	if o == nil || o.BrowserServicePort == nil {
		return nil, false
	}
	return o.BrowserServicePort, true
}

// HasBrowserServicePort returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) HasBrowserServicePort() bool {
	if o != nil && o.BrowserServicePort != nil {
		return true
	}

	return false
}

// SetBrowserServicePort gets a reference to the given int64 and assigns it to the BrowserServicePort field.
func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) SetBrowserServicePort(v int64) {
	o.BrowserServicePort = &v
}

// GetFullDomainName returns the FullDomainName field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) GetFullDomainName() string {
	if o == nil || o.FullDomainName == nil {
		var ret string
		return ret
	}
	return *o.FullDomainName
}

// GetFullDomainNameOk returns a tuple with the FullDomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) GetFullDomainNameOk() (*string, bool) {
	if o == nil || o.FullDomainName == nil {
		return nil, false
	}
	return o.FullDomainName, true
}

// HasFullDomainName returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) HasFullDomainName() bool {
	if o != nil && o.FullDomainName != nil {
		return true
	}

	return false
}

// SetFullDomainName gets a reference to the given string and assigns it to the FullDomainName field.
func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) SetFullDomainName(v string) {
	o.FullDomainName = &v
}

func (o AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BrowserServicePort != nil {
		toSerialize["BrowserServicePort"] = o.BrowserServicePort
	}
	if o.FullDomainName != nil {
		toSerialize["FullDomainName"] = o.FullDomainName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf := _AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf{}

	if err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf); err == nil {
		*o = AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf(varAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BrowserServicePort")
		delete(additionalProperties, "FullDomainName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf struct {
	value *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf
	isSet bool
}

func (v NullableAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) Get() *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf {
	return v.value
}

func (v *NullableAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) Set(val *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf(val *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) *NullableAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf {
	return &NullableAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf{value: val, isSet: true}
}

func (v NullableAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
