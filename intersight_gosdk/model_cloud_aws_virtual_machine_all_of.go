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

// CloudAwsVirtualMachineAllOf Definition of the list of properties defined in 'cloud.AwsVirtualMachine', excluding properties defined in parent classes.
type CloudAwsVirtualMachineAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType     string                           `json:"ObjectType"`
	AwsBillingUnit *CloudAwsBillingUnitRelationship `json:"AwsBillingUnit,omitempty"`
	KeyPair        *CloudAwsKeyPairRelationship     `json:"KeyPair,omitempty"`
	Location       *CloudAwsVpcRelationship         `json:"Location,omitempty"`
	// An array of relationships to cloudAwsSecurityGroup resources.
	SecurityGroups       []CloudAwsSecurityGroupRelationship `json:"SecurityGroups,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudAwsVirtualMachineAllOf CloudAwsVirtualMachineAllOf

// NewCloudAwsVirtualMachineAllOf instantiates a new CloudAwsVirtualMachineAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudAwsVirtualMachineAllOf(classId string, objectType string) *CloudAwsVirtualMachineAllOf {
	this := CloudAwsVirtualMachineAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudAwsVirtualMachineAllOfWithDefaults instantiates a new CloudAwsVirtualMachineAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudAwsVirtualMachineAllOfWithDefaults() *CloudAwsVirtualMachineAllOf {
	this := CloudAwsVirtualMachineAllOf{}
	var classId string = "cloud.AwsVirtualMachine"
	this.ClassId = classId
	var objectType string = "cloud.AwsVirtualMachine"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudAwsVirtualMachineAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudAwsVirtualMachineAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudAwsVirtualMachineAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudAwsVirtualMachineAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudAwsVirtualMachineAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudAwsVirtualMachineAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAwsBillingUnit returns the AwsBillingUnit field value if set, zero value otherwise.
func (o *CloudAwsVirtualMachineAllOf) GetAwsBillingUnit() CloudAwsBillingUnitRelationship {
	if o == nil || o.AwsBillingUnit == nil {
		var ret CloudAwsBillingUnitRelationship
		return ret
	}
	return *o.AwsBillingUnit
}

// GetAwsBillingUnitOk returns a tuple with the AwsBillingUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsVirtualMachineAllOf) GetAwsBillingUnitOk() (*CloudAwsBillingUnitRelationship, bool) {
	if o == nil || o.AwsBillingUnit == nil {
		return nil, false
	}
	return o.AwsBillingUnit, true
}

// HasAwsBillingUnit returns a boolean if a field has been set.
func (o *CloudAwsVirtualMachineAllOf) HasAwsBillingUnit() bool {
	if o != nil && o.AwsBillingUnit != nil {
		return true
	}

	return false
}

// SetAwsBillingUnit gets a reference to the given CloudAwsBillingUnitRelationship and assigns it to the AwsBillingUnit field.
func (o *CloudAwsVirtualMachineAllOf) SetAwsBillingUnit(v CloudAwsBillingUnitRelationship) {
	o.AwsBillingUnit = &v
}

// GetKeyPair returns the KeyPair field value if set, zero value otherwise.
func (o *CloudAwsVirtualMachineAllOf) GetKeyPair() CloudAwsKeyPairRelationship {
	if o == nil || o.KeyPair == nil {
		var ret CloudAwsKeyPairRelationship
		return ret
	}
	return *o.KeyPair
}

// GetKeyPairOk returns a tuple with the KeyPair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsVirtualMachineAllOf) GetKeyPairOk() (*CloudAwsKeyPairRelationship, bool) {
	if o == nil || o.KeyPair == nil {
		return nil, false
	}
	return o.KeyPair, true
}

// HasKeyPair returns a boolean if a field has been set.
func (o *CloudAwsVirtualMachineAllOf) HasKeyPair() bool {
	if o != nil && o.KeyPair != nil {
		return true
	}

	return false
}

// SetKeyPair gets a reference to the given CloudAwsKeyPairRelationship and assigns it to the KeyPair field.
func (o *CloudAwsVirtualMachineAllOf) SetKeyPair(v CloudAwsKeyPairRelationship) {
	o.KeyPair = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *CloudAwsVirtualMachineAllOf) GetLocation() CloudAwsVpcRelationship {
	if o == nil || o.Location == nil {
		var ret CloudAwsVpcRelationship
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsVirtualMachineAllOf) GetLocationOk() (*CloudAwsVpcRelationship, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *CloudAwsVirtualMachineAllOf) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given CloudAwsVpcRelationship and assigns it to the Location field.
func (o *CloudAwsVirtualMachineAllOf) SetLocation(v CloudAwsVpcRelationship) {
	o.Location = &v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudAwsVirtualMachineAllOf) GetSecurityGroups() []CloudAwsSecurityGroupRelationship {
	if o == nil {
		var ret []CloudAwsSecurityGroupRelationship
		return ret
	}
	return o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudAwsVirtualMachineAllOf) GetSecurityGroupsOk() ([]CloudAwsSecurityGroupRelationship, bool) {
	if o == nil || o.SecurityGroups == nil {
		return nil, false
	}
	return o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *CloudAwsVirtualMachineAllOf) HasSecurityGroups() bool {
	if o != nil && o.SecurityGroups != nil {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []CloudAwsSecurityGroupRelationship and assigns it to the SecurityGroups field.
func (o *CloudAwsVirtualMachineAllOf) SetSecurityGroups(v []CloudAwsSecurityGroupRelationship) {
	o.SecurityGroups = v
}

func (o CloudAwsVirtualMachineAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AwsBillingUnit != nil {
		toSerialize["AwsBillingUnit"] = o.AwsBillingUnit
	}
	if o.KeyPair != nil {
		toSerialize["KeyPair"] = o.KeyPair
	}
	if o.Location != nil {
		toSerialize["Location"] = o.Location
	}
	if o.SecurityGroups != nil {
		toSerialize["SecurityGroups"] = o.SecurityGroups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudAwsVirtualMachineAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCloudAwsVirtualMachineAllOf := _CloudAwsVirtualMachineAllOf{}

	if err = json.Unmarshal(bytes, &varCloudAwsVirtualMachineAllOf); err == nil {
		*o = CloudAwsVirtualMachineAllOf(varCloudAwsVirtualMachineAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AwsBillingUnit")
		delete(additionalProperties, "KeyPair")
		delete(additionalProperties, "Location")
		delete(additionalProperties, "SecurityGroups")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudAwsVirtualMachineAllOf struct {
	value *CloudAwsVirtualMachineAllOf
	isSet bool
}

func (v NullableCloudAwsVirtualMachineAllOf) Get() *CloudAwsVirtualMachineAllOf {
	return v.value
}

func (v *NullableCloudAwsVirtualMachineAllOf) Set(val *CloudAwsVirtualMachineAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudAwsVirtualMachineAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudAwsVirtualMachineAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudAwsVirtualMachineAllOf(val *CloudAwsVirtualMachineAllOf) *NullableCloudAwsVirtualMachineAllOf {
	return &NullableCloudAwsVirtualMachineAllOf{value: val, isSet: true}
}

func (v NullableCloudAwsVirtualMachineAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudAwsVirtualMachineAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
