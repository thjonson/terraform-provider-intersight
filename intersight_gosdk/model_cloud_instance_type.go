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

// CloudInstanceType The cpu, memory, storage and network capacity of this instance.
type CloudInstanceType struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Cpu architecture, for example, x86. * `x86_64` - The virtual machine cpu architecture type x86_64. * `x86` - The virtual machine cpu architecture type x86.
	Architecture *string `json:"Architecture,omitempty"`
	// The speed of the processor, in GHz.
	CpuSpeed *int64 `json:"CpuSpeed,omitempty"`
	// The number of CPUs for the instance type.
	Cpus *int64 `json:"Cpus,omitempty"`
	// The ID of the instance type.
	InstanceTypeId *string `json:"InstanceTypeId,omitempty"`
	// The size of the memory, in bytes.
	Memory *int64 `json:"Memory,omitempty"`
	// Name to identity the instance type.
	Name *string `json:"Name,omitempty"`
	// Operation System, for example, Linux.
	Platform             *string `json:"Platform,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudInstanceType CloudInstanceType

// NewCloudInstanceType instantiates a new CloudInstanceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudInstanceType(classId string, objectType string) *CloudInstanceType {
	this := CloudInstanceType{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudInstanceTypeWithDefaults instantiates a new CloudInstanceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudInstanceTypeWithDefaults() *CloudInstanceType {
	this := CloudInstanceType{}
	var classId string = "cloud.InstanceType"
	this.ClassId = classId
	var objectType string = "cloud.InstanceType"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudInstanceType) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudInstanceType) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudInstanceType) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudInstanceType) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudInstanceType) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudInstanceType) SetObjectType(v string) {
	o.ObjectType = v
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise.
func (o *CloudInstanceType) GetArchitecture() string {
	if o == nil || o.Architecture == nil {
		var ret string
		return ret
	}
	return *o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInstanceType) GetArchitectureOk() (*string, bool) {
	if o == nil || o.Architecture == nil {
		return nil, false
	}
	return o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *CloudInstanceType) HasArchitecture() bool {
	if o != nil && o.Architecture != nil {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given string and assigns it to the Architecture field.
func (o *CloudInstanceType) SetArchitecture(v string) {
	o.Architecture = &v
}

// GetCpuSpeed returns the CpuSpeed field value if set, zero value otherwise.
func (o *CloudInstanceType) GetCpuSpeed() int64 {
	if o == nil || o.CpuSpeed == nil {
		var ret int64
		return ret
	}
	return *o.CpuSpeed
}

// GetCpuSpeedOk returns a tuple with the CpuSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInstanceType) GetCpuSpeedOk() (*int64, bool) {
	if o == nil || o.CpuSpeed == nil {
		return nil, false
	}
	return o.CpuSpeed, true
}

// HasCpuSpeed returns a boolean if a field has been set.
func (o *CloudInstanceType) HasCpuSpeed() bool {
	if o != nil && o.CpuSpeed != nil {
		return true
	}

	return false
}

// SetCpuSpeed gets a reference to the given int64 and assigns it to the CpuSpeed field.
func (o *CloudInstanceType) SetCpuSpeed(v int64) {
	o.CpuSpeed = &v
}

// GetCpus returns the Cpus field value if set, zero value otherwise.
func (o *CloudInstanceType) GetCpus() int64 {
	if o == nil || o.Cpus == nil {
		var ret int64
		return ret
	}
	return *o.Cpus
}

// GetCpusOk returns a tuple with the Cpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInstanceType) GetCpusOk() (*int64, bool) {
	if o == nil || o.Cpus == nil {
		return nil, false
	}
	return o.Cpus, true
}

// HasCpus returns a boolean if a field has been set.
func (o *CloudInstanceType) HasCpus() bool {
	if o != nil && o.Cpus != nil {
		return true
	}

	return false
}

// SetCpus gets a reference to the given int64 and assigns it to the Cpus field.
func (o *CloudInstanceType) SetCpus(v int64) {
	o.Cpus = &v
}

// GetInstanceTypeId returns the InstanceTypeId field value if set, zero value otherwise.
func (o *CloudInstanceType) GetInstanceTypeId() string {
	if o == nil || o.InstanceTypeId == nil {
		var ret string
		return ret
	}
	return *o.InstanceTypeId
}

// GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInstanceType) GetInstanceTypeIdOk() (*string, bool) {
	if o == nil || o.InstanceTypeId == nil {
		return nil, false
	}
	return o.InstanceTypeId, true
}

// HasInstanceTypeId returns a boolean if a field has been set.
func (o *CloudInstanceType) HasInstanceTypeId() bool {
	if o != nil && o.InstanceTypeId != nil {
		return true
	}

	return false
}

// SetInstanceTypeId gets a reference to the given string and assigns it to the InstanceTypeId field.
func (o *CloudInstanceType) SetInstanceTypeId(v string) {
	o.InstanceTypeId = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *CloudInstanceType) GetMemory() int64 {
	if o == nil || o.Memory == nil {
		var ret int64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInstanceType) GetMemoryOk() (*int64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *CloudInstanceType) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int64 and assigns it to the Memory field.
func (o *CloudInstanceType) SetMemory(v int64) {
	o.Memory = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudInstanceType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInstanceType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudInstanceType) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudInstanceType) SetName(v string) {
	o.Name = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *CloudInstanceType) GetPlatform() string {
	if o == nil || o.Platform == nil {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInstanceType) GetPlatformOk() (*string, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *CloudInstanceType) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *CloudInstanceType) SetPlatform(v string) {
	o.Platform = &v
}

func (o CloudInstanceType) MarshalJSON() ([]byte, error) {
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
	if o.Architecture != nil {
		toSerialize["Architecture"] = o.Architecture
	}
	if o.CpuSpeed != nil {
		toSerialize["CpuSpeed"] = o.CpuSpeed
	}
	if o.Cpus != nil {
		toSerialize["Cpus"] = o.Cpus
	}
	if o.InstanceTypeId != nil {
		toSerialize["InstanceTypeId"] = o.InstanceTypeId
	}
	if o.Memory != nil {
		toSerialize["Memory"] = o.Memory
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Platform != nil {
		toSerialize["Platform"] = o.Platform
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudInstanceType) UnmarshalJSON(bytes []byte) (err error) {
	type CloudInstanceTypeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Cpu architecture, for example, x86. * `x86_64` - The virtual machine cpu architecture type x86_64. * `x86` - The virtual machine cpu architecture type x86.
		Architecture *string `json:"Architecture,omitempty"`
		// The speed of the processor, in GHz.
		CpuSpeed *int64 `json:"CpuSpeed,omitempty"`
		// The number of CPUs for the instance type.
		Cpus *int64 `json:"Cpus,omitempty"`
		// The ID of the instance type.
		InstanceTypeId *string `json:"InstanceTypeId,omitempty"`
		// The size of the memory, in bytes.
		Memory *int64 `json:"Memory,omitempty"`
		// Name to identity the instance type.
		Name *string `json:"Name,omitempty"`
		// Operation System, for example, Linux.
		Platform *string `json:"Platform,omitempty"`
	}

	varCloudInstanceTypeWithoutEmbeddedStruct := CloudInstanceTypeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudInstanceTypeWithoutEmbeddedStruct)
	if err == nil {
		varCloudInstanceType := _CloudInstanceType{}
		varCloudInstanceType.ClassId = varCloudInstanceTypeWithoutEmbeddedStruct.ClassId
		varCloudInstanceType.ObjectType = varCloudInstanceTypeWithoutEmbeddedStruct.ObjectType
		varCloudInstanceType.Architecture = varCloudInstanceTypeWithoutEmbeddedStruct.Architecture
		varCloudInstanceType.CpuSpeed = varCloudInstanceTypeWithoutEmbeddedStruct.CpuSpeed
		varCloudInstanceType.Cpus = varCloudInstanceTypeWithoutEmbeddedStruct.Cpus
		varCloudInstanceType.InstanceTypeId = varCloudInstanceTypeWithoutEmbeddedStruct.InstanceTypeId
		varCloudInstanceType.Memory = varCloudInstanceTypeWithoutEmbeddedStruct.Memory
		varCloudInstanceType.Name = varCloudInstanceTypeWithoutEmbeddedStruct.Name
		varCloudInstanceType.Platform = varCloudInstanceTypeWithoutEmbeddedStruct.Platform
		*o = CloudInstanceType(varCloudInstanceType)
	} else {
		return err
	}

	varCloudInstanceType := _CloudInstanceType{}

	err = json.Unmarshal(bytes, &varCloudInstanceType)
	if err == nil {
		o.MoBaseComplexType = varCloudInstanceType.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Architecture")
		delete(additionalProperties, "CpuSpeed")
		delete(additionalProperties, "Cpus")
		delete(additionalProperties, "InstanceTypeId")
		delete(additionalProperties, "Memory")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Platform")

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

type NullableCloudInstanceType struct {
	value *CloudInstanceType
	isSet bool
}

func (v NullableCloudInstanceType) Get() *CloudInstanceType {
	return v.value
}

func (v *NullableCloudInstanceType) Set(val *CloudInstanceType) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudInstanceType) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudInstanceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudInstanceType(val *CloudInstanceType) *NullableCloudInstanceType {
	return &NullableCloudInstanceType{value: val, isSet: true}
}

func (v NullableCloudInstanceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudInstanceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
