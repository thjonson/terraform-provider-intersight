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

// NiatelemetrySiteInventory Details of sites onboarded onto an ND.
type NiatelemetrySiteInventory struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string   `json:"ObjectType"`
	Apps       []string `json:"Apps,omitempty"`
	// Returns the total number of connectivity Analysis run for EPs in NDFC Fabrics.
	ConnectivityAnalysisCount *int64 `json:"ConnectivityAnalysisCount,omitempty"`
	// Version of the specified site.
	FirmwareVersion *string `json:"FirmwareVersion,omitempty"`
	// Fine-grained type DCNM either SAN or LAN.
	InstallType *string  `json:"InstallType,omitempty"`
	IpAddress   []string `json:"IpAddress,omitempty"`
	// Name of the APIC / DCNM site onboarded.
	Name *string `json:"Name,omitempty"`
	// Name of ND on which site has been onboarded.
	NexusDashboard *string `json:"NexusDashboard,omitempty"`
	// Number of nodes the site contains.
	Nodes *int64 `json:"Nodes,omitempty"`
	// Specifies whether Site object is DCNM or APIC or ND.
	RecordType *string `json:"RecordType,omitempty"`
	// Type of site onboarded either APIC or DCNM.
	Type                 *string                              `json:"Type,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetrySiteInventory NiatelemetrySiteInventory

// NewNiatelemetrySiteInventory instantiates a new NiatelemetrySiteInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetrySiteInventory(classId string, objectType string) *NiatelemetrySiteInventory {
	this := NiatelemetrySiteInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetrySiteInventoryWithDefaults instantiates a new NiatelemetrySiteInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetrySiteInventoryWithDefaults() *NiatelemetrySiteInventory {
	this := NiatelemetrySiteInventory{}
	var classId string = "niatelemetry.SiteInventory"
	this.ClassId = classId
	var objectType string = "niatelemetry.SiteInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetrySiteInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySiteInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetrySiteInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetrySiteInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySiteInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetrySiteInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetApps returns the Apps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetrySiteInventory) GetApps() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetrySiteInventory) GetAppsOk() ([]string, bool) {
	if o == nil || o.Apps == nil {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *NiatelemetrySiteInventory) HasApps() bool {
	if o != nil && o.Apps != nil {
		return true
	}

	return false
}

// SetApps gets a reference to the given []string and assigns it to the Apps field.
func (o *NiatelemetrySiteInventory) SetApps(v []string) {
	o.Apps = v
}

// GetConnectivityAnalysisCount returns the ConnectivityAnalysisCount field value if set, zero value otherwise.
func (o *NiatelemetrySiteInventory) GetConnectivityAnalysisCount() int64 {
	if o == nil || o.ConnectivityAnalysisCount == nil {
		var ret int64
		return ret
	}
	return *o.ConnectivityAnalysisCount
}

// GetConnectivityAnalysisCountOk returns a tuple with the ConnectivityAnalysisCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySiteInventory) GetConnectivityAnalysisCountOk() (*int64, bool) {
	if o == nil || o.ConnectivityAnalysisCount == nil {
		return nil, false
	}
	return o.ConnectivityAnalysisCount, true
}

// HasConnectivityAnalysisCount returns a boolean if a field has been set.
func (o *NiatelemetrySiteInventory) HasConnectivityAnalysisCount() bool {
	if o != nil && o.ConnectivityAnalysisCount != nil {
		return true
	}

	return false
}

// SetConnectivityAnalysisCount gets a reference to the given int64 and assigns it to the ConnectivityAnalysisCount field.
func (o *NiatelemetrySiteInventory) SetConnectivityAnalysisCount(v int64) {
	o.ConnectivityAnalysisCount = &v
}

// GetFirmwareVersion returns the FirmwareVersion field value if set, zero value otherwise.
func (o *NiatelemetrySiteInventory) GetFirmwareVersion() string {
	if o == nil || o.FirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.FirmwareVersion
}

// GetFirmwareVersionOk returns a tuple with the FirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySiteInventory) GetFirmwareVersionOk() (*string, bool) {
	if o == nil || o.FirmwareVersion == nil {
		return nil, false
	}
	return o.FirmwareVersion, true
}

// HasFirmwareVersion returns a boolean if a field has been set.
func (o *NiatelemetrySiteInventory) HasFirmwareVersion() bool {
	if o != nil && o.FirmwareVersion != nil {
		return true
	}

	return false
}

// SetFirmwareVersion gets a reference to the given string and assigns it to the FirmwareVersion field.
func (o *NiatelemetrySiteInventory) SetFirmwareVersion(v string) {
	o.FirmwareVersion = &v
}

// GetInstallType returns the InstallType field value if set, zero value otherwise.
func (o *NiatelemetrySiteInventory) GetInstallType() string {
	if o == nil || o.InstallType == nil {
		var ret string
		return ret
	}
	return *o.InstallType
}

// GetInstallTypeOk returns a tuple with the InstallType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySiteInventory) GetInstallTypeOk() (*string, bool) {
	if o == nil || o.InstallType == nil {
		return nil, false
	}
	return o.InstallType, true
}

// HasInstallType returns a boolean if a field has been set.
func (o *NiatelemetrySiteInventory) HasInstallType() bool {
	if o != nil && o.InstallType != nil {
		return true
	}

	return false
}

// SetInstallType gets a reference to the given string and assigns it to the InstallType field.
func (o *NiatelemetrySiteInventory) SetInstallType(v string) {
	o.InstallType = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetrySiteInventory) GetIpAddress() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetrySiteInventory) GetIpAddressOk() ([]string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *NiatelemetrySiteInventory) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given []string and assigns it to the IpAddress field.
func (o *NiatelemetrySiteInventory) SetIpAddress(v []string) {
	o.IpAddress = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NiatelemetrySiteInventory) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySiteInventory) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NiatelemetrySiteInventory) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NiatelemetrySiteInventory) SetName(v string) {
	o.Name = &v
}

// GetNexusDashboard returns the NexusDashboard field value if set, zero value otherwise.
func (o *NiatelemetrySiteInventory) GetNexusDashboard() string {
	if o == nil || o.NexusDashboard == nil {
		var ret string
		return ret
	}
	return *o.NexusDashboard
}

// GetNexusDashboardOk returns a tuple with the NexusDashboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySiteInventory) GetNexusDashboardOk() (*string, bool) {
	if o == nil || o.NexusDashboard == nil {
		return nil, false
	}
	return o.NexusDashboard, true
}

// HasNexusDashboard returns a boolean if a field has been set.
func (o *NiatelemetrySiteInventory) HasNexusDashboard() bool {
	if o != nil && o.NexusDashboard != nil {
		return true
	}

	return false
}

// SetNexusDashboard gets a reference to the given string and assigns it to the NexusDashboard field.
func (o *NiatelemetrySiteInventory) SetNexusDashboard(v string) {
	o.NexusDashboard = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *NiatelemetrySiteInventory) GetNodes() int64 {
	if o == nil || o.Nodes == nil {
		var ret int64
		return ret
	}
	return *o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySiteInventory) GetNodesOk() (*int64, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *NiatelemetrySiteInventory) HasNodes() bool {
	if o != nil && o.Nodes != nil {
		return true
	}

	return false
}

// SetNodes gets a reference to the given int64 and assigns it to the Nodes field.
func (o *NiatelemetrySiteInventory) SetNodes(v int64) {
	o.Nodes = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetrySiteInventory) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySiteInventory) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetrySiteInventory) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetrySiteInventory) SetRecordType(v string) {
	o.RecordType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NiatelemetrySiteInventory) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySiteInventory) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NiatelemetrySiteInventory) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NiatelemetrySiteInventory) SetType(v string) {
	o.Type = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetrySiteInventory) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySiteInventory) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetrySiteInventory) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetrySiteInventory) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetrySiteInventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Apps != nil {
		toSerialize["Apps"] = o.Apps
	}
	if o.ConnectivityAnalysisCount != nil {
		toSerialize["ConnectivityAnalysisCount"] = o.ConnectivityAnalysisCount
	}
	if o.FirmwareVersion != nil {
		toSerialize["FirmwareVersion"] = o.FirmwareVersion
	}
	if o.InstallType != nil {
		toSerialize["InstallType"] = o.InstallType
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.NexusDashboard != nil {
		toSerialize["NexusDashboard"] = o.NexusDashboard
	}
	if o.Nodes != nil {
		toSerialize["Nodes"] = o.Nodes
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetrySiteInventory) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetrySiteInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string   `json:"ObjectType"`
		Apps       []string `json:"Apps,omitempty"`
		// Returns the total number of connectivity Analysis run for EPs in NDFC Fabrics.
		ConnectivityAnalysisCount *int64 `json:"ConnectivityAnalysisCount,omitempty"`
		// Version of the specified site.
		FirmwareVersion *string `json:"FirmwareVersion,omitempty"`
		// Fine-grained type DCNM either SAN or LAN.
		InstallType *string  `json:"InstallType,omitempty"`
		IpAddress   []string `json:"IpAddress,omitempty"`
		// Name of the APIC / DCNM site onboarded.
		Name *string `json:"Name,omitempty"`
		// Name of ND on which site has been onboarded.
		NexusDashboard *string `json:"NexusDashboard,omitempty"`
		// Number of nodes the site contains.
		Nodes *int64 `json:"Nodes,omitempty"`
		// Specifies whether Site object is DCNM or APIC or ND.
		RecordType *string `json:"RecordType,omitempty"`
		// Type of site onboarded either APIC or DCNM.
		Type             *string                              `json:"Type,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetrySiteInventoryWithoutEmbeddedStruct := NiatelemetrySiteInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetrySiteInventoryWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetrySiteInventory := _NiatelemetrySiteInventory{}
		varNiatelemetrySiteInventory.ClassId = varNiatelemetrySiteInventoryWithoutEmbeddedStruct.ClassId
		varNiatelemetrySiteInventory.ObjectType = varNiatelemetrySiteInventoryWithoutEmbeddedStruct.ObjectType
		varNiatelemetrySiteInventory.Apps = varNiatelemetrySiteInventoryWithoutEmbeddedStruct.Apps
		varNiatelemetrySiteInventory.ConnectivityAnalysisCount = varNiatelemetrySiteInventoryWithoutEmbeddedStruct.ConnectivityAnalysisCount
		varNiatelemetrySiteInventory.FirmwareVersion = varNiatelemetrySiteInventoryWithoutEmbeddedStruct.FirmwareVersion
		varNiatelemetrySiteInventory.InstallType = varNiatelemetrySiteInventoryWithoutEmbeddedStruct.InstallType
		varNiatelemetrySiteInventory.IpAddress = varNiatelemetrySiteInventoryWithoutEmbeddedStruct.IpAddress
		varNiatelemetrySiteInventory.Name = varNiatelemetrySiteInventoryWithoutEmbeddedStruct.Name
		varNiatelemetrySiteInventory.NexusDashboard = varNiatelemetrySiteInventoryWithoutEmbeddedStruct.NexusDashboard
		varNiatelemetrySiteInventory.Nodes = varNiatelemetrySiteInventoryWithoutEmbeddedStruct.Nodes
		varNiatelemetrySiteInventory.RecordType = varNiatelemetrySiteInventoryWithoutEmbeddedStruct.RecordType
		varNiatelemetrySiteInventory.Type = varNiatelemetrySiteInventoryWithoutEmbeddedStruct.Type
		varNiatelemetrySiteInventory.RegisteredDevice = varNiatelemetrySiteInventoryWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetrySiteInventory(varNiatelemetrySiteInventory)
	} else {
		return err
	}

	varNiatelemetrySiteInventory := _NiatelemetrySiteInventory{}

	err = json.Unmarshal(bytes, &varNiatelemetrySiteInventory)
	if err == nil {
		o.MoBaseMo = varNiatelemetrySiteInventory.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Apps")
		delete(additionalProperties, "ConnectivityAnalysisCount")
		delete(additionalProperties, "FirmwareVersion")
		delete(additionalProperties, "InstallType")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "NexusDashboard")
		delete(additionalProperties, "Nodes")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableNiatelemetrySiteInventory struct {
	value *NiatelemetrySiteInventory
	isSet bool
}

func (v NullableNiatelemetrySiteInventory) Get() *NiatelemetrySiteInventory {
	return v.value
}

func (v *NullableNiatelemetrySiteInventory) Set(val *NiatelemetrySiteInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetrySiteInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetrySiteInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetrySiteInventory(val *NiatelemetrySiteInventory) *NullableNiatelemetrySiteInventory {
	return &NullableNiatelemetrySiteInventory{value: val, isSet: true}
}

func (v NullableNiatelemetrySiteInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetrySiteInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
