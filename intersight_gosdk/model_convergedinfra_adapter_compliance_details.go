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

// ConvergedinfraAdapterComplianceDetails The compliance details for an adapter present in a server which is part of a converged infrastructure pod.
type ConvergedinfraAdapterComplianceDetails struct {
	ConvergedinfraBaseComplianceDetails
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The driver name (e.g. nenic, nfnic) of the adapter.
	DriverName *string `json:"DriverName,omitempty"`
	// The driver version of the adapter.
	DriverVersion *string `json:"DriverVersion,omitempty"`
	// The firmware version of the adapter.
	Firmware *string `json:"Firmware,omitempty"`
	// The HCL compatibility status for the adapter. * `NotEvaluated` - The interoperability compliance for the component has not be checked. * `Approved` - The component is valid as per the interoperability compliance check. * `NotApproved` - The component is not valid as per the interoperability compliance check. * `Incomplete` - The interoperability compliance check could not be completed for the component due to incomplete data.
	HclStatus *string `json:"HclStatus,omitempty"`
	// The reason for the HCL status when it is not Approved. * `Missing-Os-Driver-Info` - The validation failed becaue the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information. * `Incompatible-Server-With-Component` - The validation failed for this component because he server model and component model combination was not found in the HCL. * `Incompatible-Processor` - The validation failed because the given processor was not found for the given server PID. * `Incompatible-Os-Info` - The validation failed because the given OS vendor and version was not found in HCL for the server PID and processor combination. * `Incompatible-Component-Model` - The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and OS vendor and version. * `Incompatible-Firmware` - The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, OS vendor and version and component model. * `Incompatible-Driver` - The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, OS vendor and version, server firmware and component firmware. * `Incompatible-Firmware-Driver` - The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, OS vendor and version and server firmware. * `Service-Unavailable` - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data. * `Service-Error` - The validation has failed because the HCL data service has return a service error or unrecognized result. * `Unrecognized-Protocol` - The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values. * `Not-Evaluated` - The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server's hardware profile fails to validate with HCL, then the server's software status will not be evaluated. * `Compatible` - The validation has passed for this server PID, processor, OS vendor and version, component model, component firmware and driver version.
	HclStatusReason *string `json:"HclStatusReason,omitempty"`
	// The model information of the adapter.
	Model                   *string                                            `json:"Model,omitempty"`
	Adapter                 *AdapterUnitRelationship                           `json:"Adapter,omitempty"`
	ServerComplianceDetails *ConvergedinfraServerComplianceDetailsRelationship `json:"ServerComplianceDetails,omitempty"`
	// An array of relationships to convergedinfraStorageComplianceDetails resources.
	StorageCompliances   []ConvergedinfraStorageComplianceDetailsRelationship `json:"StorageCompliances,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConvergedinfraAdapterComplianceDetails ConvergedinfraAdapterComplianceDetails

// NewConvergedinfraAdapterComplianceDetails instantiates a new ConvergedinfraAdapterComplianceDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvergedinfraAdapterComplianceDetails(classId string, objectType string) *ConvergedinfraAdapterComplianceDetails {
	this := ConvergedinfraAdapterComplianceDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConvergedinfraAdapterComplianceDetailsWithDefaults instantiates a new ConvergedinfraAdapterComplianceDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvergedinfraAdapterComplianceDetailsWithDefaults() *ConvergedinfraAdapterComplianceDetails {
	this := ConvergedinfraAdapterComplianceDetails{}
	var classId string = "convergedinfra.AdapterComplianceDetails"
	this.ClassId = classId
	var objectType string = "convergedinfra.AdapterComplianceDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConvergedinfraAdapterComplianceDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraAdapterComplianceDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConvergedinfraAdapterComplianceDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConvergedinfraAdapterComplianceDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraAdapterComplianceDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConvergedinfraAdapterComplianceDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDriverName returns the DriverName field value if set, zero value otherwise.
func (o *ConvergedinfraAdapterComplianceDetails) GetDriverName() string {
	if o == nil || o.DriverName == nil {
		var ret string
		return ret
	}
	return *o.DriverName
}

// GetDriverNameOk returns a tuple with the DriverName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraAdapterComplianceDetails) GetDriverNameOk() (*string, bool) {
	if o == nil || o.DriverName == nil {
		return nil, false
	}
	return o.DriverName, true
}

// HasDriverName returns a boolean if a field has been set.
func (o *ConvergedinfraAdapterComplianceDetails) HasDriverName() bool {
	if o != nil && o.DriverName != nil {
		return true
	}

	return false
}

// SetDriverName gets a reference to the given string and assigns it to the DriverName field.
func (o *ConvergedinfraAdapterComplianceDetails) SetDriverName(v string) {
	o.DriverName = &v
}

// GetDriverVersion returns the DriverVersion field value if set, zero value otherwise.
func (o *ConvergedinfraAdapterComplianceDetails) GetDriverVersion() string {
	if o == nil || o.DriverVersion == nil {
		var ret string
		return ret
	}
	return *o.DriverVersion
}

// GetDriverVersionOk returns a tuple with the DriverVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraAdapterComplianceDetails) GetDriverVersionOk() (*string, bool) {
	if o == nil || o.DriverVersion == nil {
		return nil, false
	}
	return o.DriverVersion, true
}

// HasDriverVersion returns a boolean if a field has been set.
func (o *ConvergedinfraAdapterComplianceDetails) HasDriverVersion() bool {
	if o != nil && o.DriverVersion != nil {
		return true
	}

	return false
}

// SetDriverVersion gets a reference to the given string and assigns it to the DriverVersion field.
func (o *ConvergedinfraAdapterComplianceDetails) SetDriverVersion(v string) {
	o.DriverVersion = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *ConvergedinfraAdapterComplianceDetails) GetFirmware() string {
	if o == nil || o.Firmware == nil {
		var ret string
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraAdapterComplianceDetails) GetFirmwareOk() (*string, bool) {
	if o == nil || o.Firmware == nil {
		return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *ConvergedinfraAdapterComplianceDetails) HasFirmware() bool {
	if o != nil && o.Firmware != nil {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given string and assigns it to the Firmware field.
func (o *ConvergedinfraAdapterComplianceDetails) SetFirmware(v string) {
	o.Firmware = &v
}

// GetHclStatus returns the HclStatus field value if set, zero value otherwise.
func (o *ConvergedinfraAdapterComplianceDetails) GetHclStatus() string {
	if o == nil || o.HclStatus == nil {
		var ret string
		return ret
	}
	return *o.HclStatus
}

// GetHclStatusOk returns a tuple with the HclStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraAdapterComplianceDetails) GetHclStatusOk() (*string, bool) {
	if o == nil || o.HclStatus == nil {
		return nil, false
	}
	return o.HclStatus, true
}

// HasHclStatus returns a boolean if a field has been set.
func (o *ConvergedinfraAdapterComplianceDetails) HasHclStatus() bool {
	if o != nil && o.HclStatus != nil {
		return true
	}

	return false
}

// SetHclStatus gets a reference to the given string and assigns it to the HclStatus field.
func (o *ConvergedinfraAdapterComplianceDetails) SetHclStatus(v string) {
	o.HclStatus = &v
}

// GetHclStatusReason returns the HclStatusReason field value if set, zero value otherwise.
func (o *ConvergedinfraAdapterComplianceDetails) GetHclStatusReason() string {
	if o == nil || o.HclStatusReason == nil {
		var ret string
		return ret
	}
	return *o.HclStatusReason
}

// GetHclStatusReasonOk returns a tuple with the HclStatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraAdapterComplianceDetails) GetHclStatusReasonOk() (*string, bool) {
	if o == nil || o.HclStatusReason == nil {
		return nil, false
	}
	return o.HclStatusReason, true
}

// HasHclStatusReason returns a boolean if a field has been set.
func (o *ConvergedinfraAdapterComplianceDetails) HasHclStatusReason() bool {
	if o != nil && o.HclStatusReason != nil {
		return true
	}

	return false
}

// SetHclStatusReason gets a reference to the given string and assigns it to the HclStatusReason field.
func (o *ConvergedinfraAdapterComplianceDetails) SetHclStatusReason(v string) {
	o.HclStatusReason = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ConvergedinfraAdapterComplianceDetails) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraAdapterComplianceDetails) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ConvergedinfraAdapterComplianceDetails) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *ConvergedinfraAdapterComplianceDetails) SetModel(v string) {
	o.Model = &v
}

// GetAdapter returns the Adapter field value if set, zero value otherwise.
func (o *ConvergedinfraAdapterComplianceDetails) GetAdapter() AdapterUnitRelationship {
	if o == nil || o.Adapter == nil {
		var ret AdapterUnitRelationship
		return ret
	}
	return *o.Adapter
}

// GetAdapterOk returns a tuple with the Adapter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraAdapterComplianceDetails) GetAdapterOk() (*AdapterUnitRelationship, bool) {
	if o == nil || o.Adapter == nil {
		return nil, false
	}
	return o.Adapter, true
}

// HasAdapter returns a boolean if a field has been set.
func (o *ConvergedinfraAdapterComplianceDetails) HasAdapter() bool {
	if o != nil && o.Adapter != nil {
		return true
	}

	return false
}

// SetAdapter gets a reference to the given AdapterUnitRelationship and assigns it to the Adapter field.
func (o *ConvergedinfraAdapterComplianceDetails) SetAdapter(v AdapterUnitRelationship) {
	o.Adapter = &v
}

// GetServerComplianceDetails returns the ServerComplianceDetails field value if set, zero value otherwise.
func (o *ConvergedinfraAdapterComplianceDetails) GetServerComplianceDetails() ConvergedinfraServerComplianceDetailsRelationship {
	if o == nil || o.ServerComplianceDetails == nil {
		var ret ConvergedinfraServerComplianceDetailsRelationship
		return ret
	}
	return *o.ServerComplianceDetails
}

// GetServerComplianceDetailsOk returns a tuple with the ServerComplianceDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraAdapterComplianceDetails) GetServerComplianceDetailsOk() (*ConvergedinfraServerComplianceDetailsRelationship, bool) {
	if o == nil || o.ServerComplianceDetails == nil {
		return nil, false
	}
	return o.ServerComplianceDetails, true
}

// HasServerComplianceDetails returns a boolean if a field has been set.
func (o *ConvergedinfraAdapterComplianceDetails) HasServerComplianceDetails() bool {
	if o != nil && o.ServerComplianceDetails != nil {
		return true
	}

	return false
}

// SetServerComplianceDetails gets a reference to the given ConvergedinfraServerComplianceDetailsRelationship and assigns it to the ServerComplianceDetails field.
func (o *ConvergedinfraAdapterComplianceDetails) SetServerComplianceDetails(v ConvergedinfraServerComplianceDetailsRelationship) {
	o.ServerComplianceDetails = &v
}

// GetStorageCompliances returns the StorageCompliances field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvergedinfraAdapterComplianceDetails) GetStorageCompliances() []ConvergedinfraStorageComplianceDetailsRelationship {
	if o == nil {
		var ret []ConvergedinfraStorageComplianceDetailsRelationship
		return ret
	}
	return o.StorageCompliances
}

// GetStorageCompliancesOk returns a tuple with the StorageCompliances field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvergedinfraAdapterComplianceDetails) GetStorageCompliancesOk() ([]ConvergedinfraStorageComplianceDetailsRelationship, bool) {
	if o == nil || o.StorageCompliances == nil {
		return nil, false
	}
	return o.StorageCompliances, true
}

// HasStorageCompliances returns a boolean if a field has been set.
func (o *ConvergedinfraAdapterComplianceDetails) HasStorageCompliances() bool {
	if o != nil && o.StorageCompliances != nil {
		return true
	}

	return false
}

// SetStorageCompliances gets a reference to the given []ConvergedinfraStorageComplianceDetailsRelationship and assigns it to the StorageCompliances field.
func (o *ConvergedinfraAdapterComplianceDetails) SetStorageCompliances(v []ConvergedinfraStorageComplianceDetailsRelationship) {
	o.StorageCompliances = v
}

func (o ConvergedinfraAdapterComplianceDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConvergedinfraBaseComplianceDetails, errConvergedinfraBaseComplianceDetails := json.Marshal(o.ConvergedinfraBaseComplianceDetails)
	if errConvergedinfraBaseComplianceDetails != nil {
		return []byte{}, errConvergedinfraBaseComplianceDetails
	}
	errConvergedinfraBaseComplianceDetails = json.Unmarshal([]byte(serializedConvergedinfraBaseComplianceDetails), &toSerialize)
	if errConvergedinfraBaseComplianceDetails != nil {
		return []byte{}, errConvergedinfraBaseComplianceDetails
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DriverName != nil {
		toSerialize["DriverName"] = o.DriverName
	}
	if o.DriverVersion != nil {
		toSerialize["DriverVersion"] = o.DriverVersion
	}
	if o.Firmware != nil {
		toSerialize["Firmware"] = o.Firmware
	}
	if o.HclStatus != nil {
		toSerialize["HclStatus"] = o.HclStatus
	}
	if o.HclStatusReason != nil {
		toSerialize["HclStatusReason"] = o.HclStatusReason
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.Adapter != nil {
		toSerialize["Adapter"] = o.Adapter
	}
	if o.ServerComplianceDetails != nil {
		toSerialize["ServerComplianceDetails"] = o.ServerComplianceDetails
	}
	if o.StorageCompliances != nil {
		toSerialize["StorageCompliances"] = o.StorageCompliances
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConvergedinfraAdapterComplianceDetails) UnmarshalJSON(bytes []byte) (err error) {
	type ConvergedinfraAdapterComplianceDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The driver name (e.g. nenic, nfnic) of the adapter.
		DriverName *string `json:"DriverName,omitempty"`
		// The driver version of the adapter.
		DriverVersion *string `json:"DriverVersion,omitempty"`
		// The firmware version of the adapter.
		Firmware *string `json:"Firmware,omitempty"`
		// The HCL compatibility status for the adapter. * `NotEvaluated` - The interoperability compliance for the component has not be checked. * `Approved` - The component is valid as per the interoperability compliance check. * `NotApproved` - The component is not valid as per the interoperability compliance check. * `Incomplete` - The interoperability compliance check could not be completed for the component due to incomplete data.
		HclStatus *string `json:"HclStatus,omitempty"`
		// The reason for the HCL status when it is not Approved. * `Missing-Os-Driver-Info` - The validation failed becaue the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information. * `Incompatible-Server-With-Component` - The validation failed for this component because he server model and component model combination was not found in the HCL. * `Incompatible-Processor` - The validation failed because the given processor was not found for the given server PID. * `Incompatible-Os-Info` - The validation failed because the given OS vendor and version was not found in HCL for the server PID and processor combination. * `Incompatible-Component-Model` - The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and OS vendor and version. * `Incompatible-Firmware` - The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, OS vendor and version and component model. * `Incompatible-Driver` - The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, OS vendor and version, server firmware and component firmware. * `Incompatible-Firmware-Driver` - The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, OS vendor and version and server firmware. * `Service-Unavailable` - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data. * `Service-Error` - The validation has failed because the HCL data service has return a service error or unrecognized result. * `Unrecognized-Protocol` - The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values. * `Not-Evaluated` - The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server's hardware profile fails to validate with HCL, then the server's software status will not be evaluated. * `Compatible` - The validation has passed for this server PID, processor, OS vendor and version, component model, component firmware and driver version.
		HclStatusReason *string `json:"HclStatusReason,omitempty"`
		// The model information of the adapter.
		Model                   *string                                            `json:"Model,omitempty"`
		Adapter                 *AdapterUnitRelationship                           `json:"Adapter,omitempty"`
		ServerComplianceDetails *ConvergedinfraServerComplianceDetailsRelationship `json:"ServerComplianceDetails,omitempty"`
		// An array of relationships to convergedinfraStorageComplianceDetails resources.
		StorageCompliances []ConvergedinfraStorageComplianceDetailsRelationship `json:"StorageCompliances,omitempty"`
	}

	varConvergedinfraAdapterComplianceDetailsWithoutEmbeddedStruct := ConvergedinfraAdapterComplianceDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConvergedinfraAdapterComplianceDetailsWithoutEmbeddedStruct)
	if err == nil {
		varConvergedinfraAdapterComplianceDetails := _ConvergedinfraAdapterComplianceDetails{}
		varConvergedinfraAdapterComplianceDetails.ClassId = varConvergedinfraAdapterComplianceDetailsWithoutEmbeddedStruct.ClassId
		varConvergedinfraAdapterComplianceDetails.ObjectType = varConvergedinfraAdapterComplianceDetailsWithoutEmbeddedStruct.ObjectType
		varConvergedinfraAdapterComplianceDetails.DriverName = varConvergedinfraAdapterComplianceDetailsWithoutEmbeddedStruct.DriverName
		varConvergedinfraAdapterComplianceDetails.DriverVersion = varConvergedinfraAdapterComplianceDetailsWithoutEmbeddedStruct.DriverVersion
		varConvergedinfraAdapterComplianceDetails.Firmware = varConvergedinfraAdapterComplianceDetailsWithoutEmbeddedStruct.Firmware
		varConvergedinfraAdapterComplianceDetails.HclStatus = varConvergedinfraAdapterComplianceDetailsWithoutEmbeddedStruct.HclStatus
		varConvergedinfraAdapterComplianceDetails.HclStatusReason = varConvergedinfraAdapterComplianceDetailsWithoutEmbeddedStruct.HclStatusReason
		varConvergedinfraAdapterComplianceDetails.Model = varConvergedinfraAdapterComplianceDetailsWithoutEmbeddedStruct.Model
		varConvergedinfraAdapterComplianceDetails.Adapter = varConvergedinfraAdapterComplianceDetailsWithoutEmbeddedStruct.Adapter
		varConvergedinfraAdapterComplianceDetails.ServerComplianceDetails = varConvergedinfraAdapterComplianceDetailsWithoutEmbeddedStruct.ServerComplianceDetails
		varConvergedinfraAdapterComplianceDetails.StorageCompliances = varConvergedinfraAdapterComplianceDetailsWithoutEmbeddedStruct.StorageCompliances
		*o = ConvergedinfraAdapterComplianceDetails(varConvergedinfraAdapterComplianceDetails)
	} else {
		return err
	}

	varConvergedinfraAdapterComplianceDetails := _ConvergedinfraAdapterComplianceDetails{}

	err = json.Unmarshal(bytes, &varConvergedinfraAdapterComplianceDetails)
	if err == nil {
		o.ConvergedinfraBaseComplianceDetails = varConvergedinfraAdapterComplianceDetails.ConvergedinfraBaseComplianceDetails
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DriverName")
		delete(additionalProperties, "DriverVersion")
		delete(additionalProperties, "Firmware")
		delete(additionalProperties, "HclStatus")
		delete(additionalProperties, "HclStatusReason")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "Adapter")
		delete(additionalProperties, "ServerComplianceDetails")
		delete(additionalProperties, "StorageCompliances")

		// remove fields from embedded structs
		reflectConvergedinfraBaseComplianceDetails := reflect.ValueOf(o.ConvergedinfraBaseComplianceDetails)
		for i := 0; i < reflectConvergedinfraBaseComplianceDetails.Type().NumField(); i++ {
			t := reflectConvergedinfraBaseComplianceDetails.Type().Field(i)

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

type NullableConvergedinfraAdapterComplianceDetails struct {
	value *ConvergedinfraAdapterComplianceDetails
	isSet bool
}

func (v NullableConvergedinfraAdapterComplianceDetails) Get() *ConvergedinfraAdapterComplianceDetails {
	return v.value
}

func (v *NullableConvergedinfraAdapterComplianceDetails) Set(val *ConvergedinfraAdapterComplianceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableConvergedinfraAdapterComplianceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableConvergedinfraAdapterComplianceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvergedinfraAdapterComplianceDetails(val *ConvergedinfraAdapterComplianceDetails) *NullableConvergedinfraAdapterComplianceDetails {
	return &NullableConvergedinfraAdapterComplianceDetails{value: val, isSet: true}
}

func (v NullableConvergedinfraAdapterComplianceDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvergedinfraAdapterComplianceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
