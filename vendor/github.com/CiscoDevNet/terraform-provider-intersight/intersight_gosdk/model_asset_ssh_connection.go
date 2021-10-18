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

// AssetSshConnection SshConnection provides the necessary details for Intersight to connect and authenticate with  a managed target over an SSH connection.
type AssetSshConnection struct {
	AssetConnection
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The DNS hostname or IP Address, either IPv4 or IPv6, to be used to connect to the managed target.
	ManagementAddress *string `json:"ManagementAddress,omitempty"`
	// The port number to be used to connect to the managed target. Values 1-65535 indicate a port number to be used. A value of 0 is not a valid port number and instead indicates that the default management port, as defined by the documentation of the managed target, should be used to establish a connection.
	Port                 *int64 `json:"Port,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetSshConnection AssetSshConnection

// NewAssetSshConnection instantiates a new AssetSshConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetSshConnection(classId string, objectType string) *AssetSshConnection {
	this := AssetSshConnection{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetSshConnectionWithDefaults instantiates a new AssetSshConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetSshConnectionWithDefaults() *AssetSshConnection {
	this := AssetSshConnection{}
	var classId string = "asset.SshConnection"
	this.ClassId = classId
	var objectType string = "asset.SshConnection"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetSshConnection) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetSshConnection) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetSshConnection) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetSshConnection) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetSshConnection) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetSshConnection) SetObjectType(v string) {
	o.ObjectType = v
}

// GetManagementAddress returns the ManagementAddress field value if set, zero value otherwise.
func (o *AssetSshConnection) GetManagementAddress() string {
	if o == nil || o.ManagementAddress == nil {
		var ret string
		return ret
	}
	return *o.ManagementAddress
}

// GetManagementAddressOk returns a tuple with the ManagementAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSshConnection) GetManagementAddressOk() (*string, bool) {
	if o == nil || o.ManagementAddress == nil {
		return nil, false
	}
	return o.ManagementAddress, true
}

// HasManagementAddress returns a boolean if a field has been set.
func (o *AssetSshConnection) HasManagementAddress() bool {
	if o != nil && o.ManagementAddress != nil {
		return true
	}

	return false
}

// SetManagementAddress gets a reference to the given string and assigns it to the ManagementAddress field.
func (o *AssetSshConnection) SetManagementAddress(v string) {
	o.ManagementAddress = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *AssetSshConnection) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSshConnection) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *AssetSshConnection) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *AssetSshConnection) SetPort(v int64) {
	o.Port = &v
}

func (o AssetSshConnection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetConnection, errAssetConnection := json.Marshal(o.AssetConnection)
	if errAssetConnection != nil {
		return []byte{}, errAssetConnection
	}
	errAssetConnection = json.Unmarshal([]byte(serializedAssetConnection), &toSerialize)
	if errAssetConnection != nil {
		return []byte{}, errAssetConnection
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ManagementAddress != nil {
		toSerialize["ManagementAddress"] = o.ManagementAddress
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetSshConnection) UnmarshalJSON(bytes []byte) (err error) {
	type AssetSshConnectionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The DNS hostname or IP Address, either IPv4 or IPv6, to be used to connect to the managed target.
		ManagementAddress *string `json:"ManagementAddress,omitempty"`
		// The port number to be used to connect to the managed target. Values 1-65535 indicate a port number to be used. A value of 0 is not a valid port number and instead indicates that the default management port, as defined by the documentation of the managed target, should be used to establish a connection.
		Port *int64 `json:"Port,omitempty"`
	}

	varAssetSshConnectionWithoutEmbeddedStruct := AssetSshConnectionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetSshConnectionWithoutEmbeddedStruct)
	if err == nil {
		varAssetSshConnection := _AssetSshConnection{}
		varAssetSshConnection.ClassId = varAssetSshConnectionWithoutEmbeddedStruct.ClassId
		varAssetSshConnection.ObjectType = varAssetSshConnectionWithoutEmbeddedStruct.ObjectType
		varAssetSshConnection.ManagementAddress = varAssetSshConnectionWithoutEmbeddedStruct.ManagementAddress
		varAssetSshConnection.Port = varAssetSshConnectionWithoutEmbeddedStruct.Port
		*o = AssetSshConnection(varAssetSshConnection)
	} else {
		return err
	}

	varAssetSshConnection := _AssetSshConnection{}

	err = json.Unmarshal(bytes, &varAssetSshConnection)
	if err == nil {
		o.AssetConnection = varAssetSshConnection.AssetConnection
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ManagementAddress")
		delete(additionalProperties, "Port")

		// remove fields from embedded structs
		reflectAssetConnection := reflect.ValueOf(o.AssetConnection)
		for i := 0; i < reflectAssetConnection.Type().NumField(); i++ {
			t := reflectAssetConnection.Type().Field(i)

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

type NullableAssetSshConnection struct {
	value *AssetSshConnection
	isSet bool
}

func (v NullableAssetSshConnection) Get() *AssetSshConnection {
	return v.value
}

func (v *NullableAssetSshConnection) Set(val *AssetSshConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetSshConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetSshConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetSshConnection(val *AssetSshConnection) *NullableAssetSshConnection {
	return &NullableAssetSshConnection{value: val, isSet: true}
}

func (v NullableAssetSshConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetSshConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
