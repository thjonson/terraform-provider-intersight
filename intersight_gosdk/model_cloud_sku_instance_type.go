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
	"reflect"
	"strings"
)

// CloudSkuInstanceType Details for an instance type.
type CloudSkuInstanceType struct {
	CloudBaseSku
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates if the instance type supports 32 or 64 bit or both. * `64Bit` - Enum to indicate that the instance type suppports only 64 bit architecture. * `32Bit` - Enum to indicate that the instance type supports only 32 bit architecture. * `both` - Enum to indicate that the instance type supports both 32 and 64 bit architecture.
	ArchitectureType *string `json:"ArchitectureType,omitempty"`
	// The cpu unit for this instance type. * `VIRTUAL_CPU` - The CPU unit used for virtual machines. * `MILLI_CPU` - The CPU unit used by containers.
	CpuUnit *string `json:"CpuUnit,omitempty"`
	// Does the instanceType support CUDA architecture.
	CudaSupport *bool `json:"CudaSupport,omitempty"`
	// Total size of local storage for this instance type.
	LocalStorageSize *float32 `json:"LocalStorageSize,omitempty"`
	// The units for this storage. For e.g. MB or GB or TB. * `MB` - Enum to represent mega byte storage unit. * `GB` - Enum to represent Giga byte storage unit. * `TB` - Enum to represent Tera byte storage unit.
	LocalStorageSizeUnit *string `json:"LocalStorageSizeUnit,omitempty"`
	// The RAM size for this instance type.
	MemorySize *float32 `json:"MemorySize,omitempty"`
	// Units to represent memory, for e.g. MB or GB. * `MB` - Enum to represent mega byte storage unit. * `GB` - Enum to represent Giga byte storage unit. * `TB` - Enum to represent Tera byte storage unit.
	MemorySizeUnit *string `json:"MemorySizeUnit,omitempty"`
	// Network performance of this instance type.
	NetworkPerformance *string `json:"NetworkPerformance,omitempty"`
	// The number of CPUs in this instance type.
	NumOfCpus *int64 `json:"NumOfCpus,omitempty"`
	// Maximum number of NICs supported by this instance type.
	NumOfNics            *int64 `json:"NumOfNics,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudSkuInstanceType CloudSkuInstanceType

// NewCloudSkuInstanceType instantiates a new CloudSkuInstanceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudSkuInstanceType(classId string, objectType string) *CloudSkuInstanceType {
	this := CloudSkuInstanceType{}
	this.ClassId = classId
	this.ObjectType = objectType
	var architectureType string = "64Bit"
	this.ArchitectureType = &architectureType
	var cpuUnit string = "VIRTUAL_CPU"
	this.CpuUnit = &cpuUnit
	var localStorageSizeUnit string = "MB"
	this.LocalStorageSizeUnit = &localStorageSizeUnit
	var memorySizeUnit string = "MB"
	this.MemorySizeUnit = &memorySizeUnit
	return &this
}

// NewCloudSkuInstanceTypeWithDefaults instantiates a new CloudSkuInstanceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudSkuInstanceTypeWithDefaults() *CloudSkuInstanceType {
	this := CloudSkuInstanceType{}
	var classId string = "cloud.SkuInstanceType"
	this.ClassId = classId
	var objectType string = "cloud.SkuInstanceType"
	this.ObjectType = objectType
	var architectureType string = "64Bit"
	this.ArchitectureType = &architectureType
	var cpuUnit string = "VIRTUAL_CPU"
	this.CpuUnit = &cpuUnit
	var localStorageSizeUnit string = "MB"
	this.LocalStorageSizeUnit = &localStorageSizeUnit
	var memorySizeUnit string = "MB"
	this.MemorySizeUnit = &memorySizeUnit
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudSkuInstanceType) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceType) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudSkuInstanceType) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudSkuInstanceType) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceType) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudSkuInstanceType) SetObjectType(v string) {
	o.ObjectType = v
}

// GetArchitectureType returns the ArchitectureType field value if set, zero value otherwise.
func (o *CloudSkuInstanceType) GetArchitectureType() string {
	if o == nil || o.ArchitectureType == nil {
		var ret string
		return ret
	}
	return *o.ArchitectureType
}

// GetArchitectureTypeOk returns a tuple with the ArchitectureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceType) GetArchitectureTypeOk() (*string, bool) {
	if o == nil || o.ArchitectureType == nil {
		return nil, false
	}
	return o.ArchitectureType, true
}

// HasArchitectureType returns a boolean if a field has been set.
func (o *CloudSkuInstanceType) HasArchitectureType() bool {
	if o != nil && o.ArchitectureType != nil {
		return true
	}

	return false
}

// SetArchitectureType gets a reference to the given string and assigns it to the ArchitectureType field.
func (o *CloudSkuInstanceType) SetArchitectureType(v string) {
	o.ArchitectureType = &v
}

// GetCpuUnit returns the CpuUnit field value if set, zero value otherwise.
func (o *CloudSkuInstanceType) GetCpuUnit() string {
	if o == nil || o.CpuUnit == nil {
		var ret string
		return ret
	}
	return *o.CpuUnit
}

// GetCpuUnitOk returns a tuple with the CpuUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceType) GetCpuUnitOk() (*string, bool) {
	if o == nil || o.CpuUnit == nil {
		return nil, false
	}
	return o.CpuUnit, true
}

// HasCpuUnit returns a boolean if a field has been set.
func (o *CloudSkuInstanceType) HasCpuUnit() bool {
	if o != nil && o.CpuUnit != nil {
		return true
	}

	return false
}

// SetCpuUnit gets a reference to the given string and assigns it to the CpuUnit field.
func (o *CloudSkuInstanceType) SetCpuUnit(v string) {
	o.CpuUnit = &v
}

// GetCudaSupport returns the CudaSupport field value if set, zero value otherwise.
func (o *CloudSkuInstanceType) GetCudaSupport() bool {
	if o == nil || o.CudaSupport == nil {
		var ret bool
		return ret
	}
	return *o.CudaSupport
}

// GetCudaSupportOk returns a tuple with the CudaSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceType) GetCudaSupportOk() (*bool, bool) {
	if o == nil || o.CudaSupport == nil {
		return nil, false
	}
	return o.CudaSupport, true
}

// HasCudaSupport returns a boolean if a field has been set.
func (o *CloudSkuInstanceType) HasCudaSupport() bool {
	if o != nil && o.CudaSupport != nil {
		return true
	}

	return false
}

// SetCudaSupport gets a reference to the given bool and assigns it to the CudaSupport field.
func (o *CloudSkuInstanceType) SetCudaSupport(v bool) {
	o.CudaSupport = &v
}

// GetLocalStorageSize returns the LocalStorageSize field value if set, zero value otherwise.
func (o *CloudSkuInstanceType) GetLocalStorageSize() float32 {
	if o == nil || o.LocalStorageSize == nil {
		var ret float32
		return ret
	}
	return *o.LocalStorageSize
}

// GetLocalStorageSizeOk returns a tuple with the LocalStorageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceType) GetLocalStorageSizeOk() (*float32, bool) {
	if o == nil || o.LocalStorageSize == nil {
		return nil, false
	}
	return o.LocalStorageSize, true
}

// HasLocalStorageSize returns a boolean if a field has been set.
func (o *CloudSkuInstanceType) HasLocalStorageSize() bool {
	if o != nil && o.LocalStorageSize != nil {
		return true
	}

	return false
}

// SetLocalStorageSize gets a reference to the given float32 and assigns it to the LocalStorageSize field.
func (o *CloudSkuInstanceType) SetLocalStorageSize(v float32) {
	o.LocalStorageSize = &v
}

// GetLocalStorageSizeUnit returns the LocalStorageSizeUnit field value if set, zero value otherwise.
func (o *CloudSkuInstanceType) GetLocalStorageSizeUnit() string {
	if o == nil || o.LocalStorageSizeUnit == nil {
		var ret string
		return ret
	}
	return *o.LocalStorageSizeUnit
}

// GetLocalStorageSizeUnitOk returns a tuple with the LocalStorageSizeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceType) GetLocalStorageSizeUnitOk() (*string, bool) {
	if o == nil || o.LocalStorageSizeUnit == nil {
		return nil, false
	}
	return o.LocalStorageSizeUnit, true
}

// HasLocalStorageSizeUnit returns a boolean if a field has been set.
func (o *CloudSkuInstanceType) HasLocalStorageSizeUnit() bool {
	if o != nil && o.LocalStorageSizeUnit != nil {
		return true
	}

	return false
}

// SetLocalStorageSizeUnit gets a reference to the given string and assigns it to the LocalStorageSizeUnit field.
func (o *CloudSkuInstanceType) SetLocalStorageSizeUnit(v string) {
	o.LocalStorageSizeUnit = &v
}

// GetMemorySize returns the MemorySize field value if set, zero value otherwise.
func (o *CloudSkuInstanceType) GetMemorySize() float32 {
	if o == nil || o.MemorySize == nil {
		var ret float32
		return ret
	}
	return *o.MemorySize
}

// GetMemorySizeOk returns a tuple with the MemorySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceType) GetMemorySizeOk() (*float32, bool) {
	if o == nil || o.MemorySize == nil {
		return nil, false
	}
	return o.MemorySize, true
}

// HasMemorySize returns a boolean if a field has been set.
func (o *CloudSkuInstanceType) HasMemorySize() bool {
	if o != nil && o.MemorySize != nil {
		return true
	}

	return false
}

// SetMemorySize gets a reference to the given float32 and assigns it to the MemorySize field.
func (o *CloudSkuInstanceType) SetMemorySize(v float32) {
	o.MemorySize = &v
}

// GetMemorySizeUnit returns the MemorySizeUnit field value if set, zero value otherwise.
func (o *CloudSkuInstanceType) GetMemorySizeUnit() string {
	if o == nil || o.MemorySizeUnit == nil {
		var ret string
		return ret
	}
	return *o.MemorySizeUnit
}

// GetMemorySizeUnitOk returns a tuple with the MemorySizeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceType) GetMemorySizeUnitOk() (*string, bool) {
	if o == nil || o.MemorySizeUnit == nil {
		return nil, false
	}
	return o.MemorySizeUnit, true
}

// HasMemorySizeUnit returns a boolean if a field has been set.
func (o *CloudSkuInstanceType) HasMemorySizeUnit() bool {
	if o != nil && o.MemorySizeUnit != nil {
		return true
	}

	return false
}

// SetMemorySizeUnit gets a reference to the given string and assigns it to the MemorySizeUnit field.
func (o *CloudSkuInstanceType) SetMemorySizeUnit(v string) {
	o.MemorySizeUnit = &v
}

// GetNetworkPerformance returns the NetworkPerformance field value if set, zero value otherwise.
func (o *CloudSkuInstanceType) GetNetworkPerformance() string {
	if o == nil || o.NetworkPerformance == nil {
		var ret string
		return ret
	}
	return *o.NetworkPerformance
}

// GetNetworkPerformanceOk returns a tuple with the NetworkPerformance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceType) GetNetworkPerformanceOk() (*string, bool) {
	if o == nil || o.NetworkPerformance == nil {
		return nil, false
	}
	return o.NetworkPerformance, true
}

// HasNetworkPerformance returns a boolean if a field has been set.
func (o *CloudSkuInstanceType) HasNetworkPerformance() bool {
	if o != nil && o.NetworkPerformance != nil {
		return true
	}

	return false
}

// SetNetworkPerformance gets a reference to the given string and assigns it to the NetworkPerformance field.
func (o *CloudSkuInstanceType) SetNetworkPerformance(v string) {
	o.NetworkPerformance = &v
}

// GetNumOfCpus returns the NumOfCpus field value if set, zero value otherwise.
func (o *CloudSkuInstanceType) GetNumOfCpus() int64 {
	if o == nil || o.NumOfCpus == nil {
		var ret int64
		return ret
	}
	return *o.NumOfCpus
}

// GetNumOfCpusOk returns a tuple with the NumOfCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceType) GetNumOfCpusOk() (*int64, bool) {
	if o == nil || o.NumOfCpus == nil {
		return nil, false
	}
	return o.NumOfCpus, true
}

// HasNumOfCpus returns a boolean if a field has been set.
func (o *CloudSkuInstanceType) HasNumOfCpus() bool {
	if o != nil && o.NumOfCpus != nil {
		return true
	}

	return false
}

// SetNumOfCpus gets a reference to the given int64 and assigns it to the NumOfCpus field.
func (o *CloudSkuInstanceType) SetNumOfCpus(v int64) {
	o.NumOfCpus = &v
}

// GetNumOfNics returns the NumOfNics field value if set, zero value otherwise.
func (o *CloudSkuInstanceType) GetNumOfNics() int64 {
	if o == nil || o.NumOfNics == nil {
		var ret int64
		return ret
	}
	return *o.NumOfNics
}

// GetNumOfNicsOk returns a tuple with the NumOfNics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceType) GetNumOfNicsOk() (*int64, bool) {
	if o == nil || o.NumOfNics == nil {
		return nil, false
	}
	return o.NumOfNics, true
}

// HasNumOfNics returns a boolean if a field has been set.
func (o *CloudSkuInstanceType) HasNumOfNics() bool {
	if o != nil && o.NumOfNics != nil {
		return true
	}

	return false
}

// SetNumOfNics gets a reference to the given int64 and assigns it to the NumOfNics field.
func (o *CloudSkuInstanceType) SetNumOfNics(v int64) {
	o.NumOfNics = &v
}

func (o CloudSkuInstanceType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseSku, errCloudBaseSku := json.Marshal(o.CloudBaseSku)
	if errCloudBaseSku != nil {
		return []byte{}, errCloudBaseSku
	}
	errCloudBaseSku = json.Unmarshal([]byte(serializedCloudBaseSku), &toSerialize)
	if errCloudBaseSku != nil {
		return []byte{}, errCloudBaseSku
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ArchitectureType != nil {
		toSerialize["ArchitectureType"] = o.ArchitectureType
	}
	if o.CpuUnit != nil {
		toSerialize["CpuUnit"] = o.CpuUnit
	}
	if o.CudaSupport != nil {
		toSerialize["CudaSupport"] = o.CudaSupport
	}
	if o.LocalStorageSize != nil {
		toSerialize["LocalStorageSize"] = o.LocalStorageSize
	}
	if o.LocalStorageSizeUnit != nil {
		toSerialize["LocalStorageSizeUnit"] = o.LocalStorageSizeUnit
	}
	if o.MemorySize != nil {
		toSerialize["MemorySize"] = o.MemorySize
	}
	if o.MemorySizeUnit != nil {
		toSerialize["MemorySizeUnit"] = o.MemorySizeUnit
	}
	if o.NetworkPerformance != nil {
		toSerialize["NetworkPerformance"] = o.NetworkPerformance
	}
	if o.NumOfCpus != nil {
		toSerialize["NumOfCpus"] = o.NumOfCpus
	}
	if o.NumOfNics != nil {
		toSerialize["NumOfNics"] = o.NumOfNics
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudSkuInstanceType) UnmarshalJSON(bytes []byte) (err error) {
	type CloudSkuInstanceTypeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Indicates if the instance type supports 32 or 64 bit or both. * `64Bit` - Enum to indicate that the instance type suppports only 64 bit architecture. * `32Bit` - Enum to indicate that the instance type supports only 32 bit architecture. * `both` - Enum to indicate that the instance type supports both 32 and 64 bit architecture.
		ArchitectureType *string `json:"ArchitectureType,omitempty"`
		// The cpu unit for this instance type. * `VIRTUAL_CPU` - The CPU unit used for virtual machines. * `MILLI_CPU` - The CPU unit used by containers.
		CpuUnit *string `json:"CpuUnit,omitempty"`
		// Does the instanceType support CUDA architecture.
		CudaSupport *bool `json:"CudaSupport,omitempty"`
		// Total size of local storage for this instance type.
		LocalStorageSize *float32 `json:"LocalStorageSize,omitempty"`
		// The units for this storage. For e.g. MB or GB or TB. * `MB` - Enum to represent mega byte storage unit. * `GB` - Enum to represent Giga byte storage unit. * `TB` - Enum to represent Tera byte storage unit.
		LocalStorageSizeUnit *string `json:"LocalStorageSizeUnit,omitempty"`
		// The RAM size for this instance type.
		MemorySize *float32 `json:"MemorySize,omitempty"`
		// Units to represent memory, for e.g. MB or GB. * `MB` - Enum to represent mega byte storage unit. * `GB` - Enum to represent Giga byte storage unit. * `TB` - Enum to represent Tera byte storage unit.
		MemorySizeUnit *string `json:"MemorySizeUnit,omitempty"`
		// Network performance of this instance type.
		NetworkPerformance *string `json:"NetworkPerformance,omitempty"`
		// The number of CPUs in this instance type.
		NumOfCpus *int64 `json:"NumOfCpus,omitempty"`
		// Maximum number of NICs supported by this instance type.
		NumOfNics *int64 `json:"NumOfNics,omitempty"`
	}

	varCloudSkuInstanceTypeWithoutEmbeddedStruct := CloudSkuInstanceTypeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudSkuInstanceTypeWithoutEmbeddedStruct)
	if err == nil {
		varCloudSkuInstanceType := _CloudSkuInstanceType{}
		varCloudSkuInstanceType.ClassId = varCloudSkuInstanceTypeWithoutEmbeddedStruct.ClassId
		varCloudSkuInstanceType.ObjectType = varCloudSkuInstanceTypeWithoutEmbeddedStruct.ObjectType
		varCloudSkuInstanceType.ArchitectureType = varCloudSkuInstanceTypeWithoutEmbeddedStruct.ArchitectureType
		varCloudSkuInstanceType.CpuUnit = varCloudSkuInstanceTypeWithoutEmbeddedStruct.CpuUnit
		varCloudSkuInstanceType.CudaSupport = varCloudSkuInstanceTypeWithoutEmbeddedStruct.CudaSupport
		varCloudSkuInstanceType.LocalStorageSize = varCloudSkuInstanceTypeWithoutEmbeddedStruct.LocalStorageSize
		varCloudSkuInstanceType.LocalStorageSizeUnit = varCloudSkuInstanceTypeWithoutEmbeddedStruct.LocalStorageSizeUnit
		varCloudSkuInstanceType.MemorySize = varCloudSkuInstanceTypeWithoutEmbeddedStruct.MemorySize
		varCloudSkuInstanceType.MemorySizeUnit = varCloudSkuInstanceTypeWithoutEmbeddedStruct.MemorySizeUnit
		varCloudSkuInstanceType.NetworkPerformance = varCloudSkuInstanceTypeWithoutEmbeddedStruct.NetworkPerformance
		varCloudSkuInstanceType.NumOfCpus = varCloudSkuInstanceTypeWithoutEmbeddedStruct.NumOfCpus
		varCloudSkuInstanceType.NumOfNics = varCloudSkuInstanceTypeWithoutEmbeddedStruct.NumOfNics
		*o = CloudSkuInstanceType(varCloudSkuInstanceType)
	} else {
		return err
	}

	varCloudSkuInstanceType := _CloudSkuInstanceType{}

	err = json.Unmarshal(bytes, &varCloudSkuInstanceType)
	if err == nil {
		o.CloudBaseSku = varCloudSkuInstanceType.CloudBaseSku
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ArchitectureType")
		delete(additionalProperties, "CpuUnit")
		delete(additionalProperties, "CudaSupport")
		delete(additionalProperties, "LocalStorageSize")
		delete(additionalProperties, "LocalStorageSizeUnit")
		delete(additionalProperties, "MemorySize")
		delete(additionalProperties, "MemorySizeUnit")
		delete(additionalProperties, "NetworkPerformance")
		delete(additionalProperties, "NumOfCpus")
		delete(additionalProperties, "NumOfNics")

		// remove fields from embedded structs
		reflectCloudBaseSku := reflect.ValueOf(o.CloudBaseSku)
		for i := 0; i < reflectCloudBaseSku.Type().NumField(); i++ {
			t := reflectCloudBaseSku.Type().Field(i)

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

type NullableCloudSkuInstanceType struct {
	value *CloudSkuInstanceType
	isSet bool
}

func (v NullableCloudSkuInstanceType) Get() *CloudSkuInstanceType {
	return v.value
}

func (v *NullableCloudSkuInstanceType) Set(val *CloudSkuInstanceType) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudSkuInstanceType) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudSkuInstanceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudSkuInstanceType(val *CloudSkuInstanceType) *NullableCloudSkuInstanceType {
	return &NullableCloudSkuInstanceType{value: val, isSet: true}
}

func (v NullableCloudSkuInstanceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudSkuInstanceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
