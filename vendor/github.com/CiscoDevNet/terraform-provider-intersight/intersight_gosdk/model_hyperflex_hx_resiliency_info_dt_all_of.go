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

// HyperflexHxResiliencyInfoDtAllOf Definition of the list of properties defined in 'hyperflex.HxResiliencyInfoDt', excluding properties defined in parent classes.
type HyperflexHxResiliencyInfoDtAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of copies of data replicated by the cluster. * `ONE_COPY` - The HyperFlex cluster does not replicate data. * `TWO_COPIES` - The HyperFlex cluster keeps 2 copies of data. * `THREE_COPIES` - The HyperFlex cluster keeps 3 copies of data. * `FOUR_COPIES` - The HyperFlex cluster keeps 4 copies of data. * `SIX_COPIES` - The HyperFlex cluster keeps 6 copies of data.
	DataReplicationFactor *string `json:"DataReplicationFactor,omitempty"`
	// The number of persistent device disruptions the HyperFlex storage cluster can handle at this point in time.
	HddFailuresTolerable *int64   `json:"HddFailuresTolerable,omitempty"`
	Messages             []string `json:"Messages,omitempty"`
	// The number of node disruptions the HyperFlex storage cluster can handle at this point in time.
	NodeFailuresTolerable *int64 `json:"NodeFailuresTolerable,omitempty"`
	// The state of the storage cluster's compliance with the cluster access policy. * `UNKNOWN` - The HyperFlex cluster's compliance with the data replication policy could not be determined. * `COMPLIANT` - The HyperFlex cluster is compliant with the data replication policy and data is replicated to the configured replication factor. * `NON_COMPLIANT` - The HyperFlex cluster is not compliant with the data replication policy because some data is not replicated.
	PolicyCompliance *string `json:"PolicyCompliance,omitempty"`
	// The resiliency state of the storage platform. The resiliency state is the storage cluster's current ability to maintain data. * `UNKNOWN` - The resiliency status of the HyperFlex cluster cannot be determined, or the cluster is transitioning into ONLINE state. * `HEALTHY` - The HyperFlex cluster is healthy. The cluster is able to perform IO operations and data is available. * `WARNING` - The HyperFlex cluster or data availability is adversely affected. This can happen if there are node or storage device failures beyond the tolerable failure threshold. * `OFFLINE` - The HyperFlex cluster is offline and not performing IO operations. * `CRITICAL` - The HyperFlex cluster has severe faults that affect cluster and data availability.
	ResiliencyState *string `json:"ResiliencyState,omitempty"`
	// The number of cache device disruptions the HyperFlex storage cluster can handle at this point in time.
	SsdFailuresTolerable *int64 `json:"SsdFailuresTolerable,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxResiliencyInfoDtAllOf HyperflexHxResiliencyInfoDtAllOf

// NewHyperflexHxResiliencyInfoDtAllOf instantiates a new HyperflexHxResiliencyInfoDtAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxResiliencyInfoDtAllOf(classId string, objectType string) *HyperflexHxResiliencyInfoDtAllOf {
	this := HyperflexHxResiliencyInfoDtAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHxResiliencyInfoDtAllOfWithDefaults instantiates a new HyperflexHxResiliencyInfoDtAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxResiliencyInfoDtAllOfWithDefaults() *HyperflexHxResiliencyInfoDtAllOf {
	this := HyperflexHxResiliencyInfoDtAllOf{}
	var classId string = "hyperflex.HxResiliencyInfoDt"
	this.ClassId = classId
	var objectType string = "hyperflex.HxResiliencyInfoDt"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxResiliencyInfoDtAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxResiliencyInfoDtAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxResiliencyInfoDtAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxResiliencyInfoDtAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxResiliencyInfoDtAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxResiliencyInfoDtAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDataReplicationFactor returns the DataReplicationFactor field value if set, zero value otherwise.
func (o *HyperflexHxResiliencyInfoDtAllOf) GetDataReplicationFactor() string {
	if o == nil || o.DataReplicationFactor == nil {
		var ret string
		return ret
	}
	return *o.DataReplicationFactor
}

// GetDataReplicationFactorOk returns a tuple with the DataReplicationFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxResiliencyInfoDtAllOf) GetDataReplicationFactorOk() (*string, bool) {
	if o == nil || o.DataReplicationFactor == nil {
		return nil, false
	}
	return o.DataReplicationFactor, true
}

// HasDataReplicationFactor returns a boolean if a field has been set.
func (o *HyperflexHxResiliencyInfoDtAllOf) HasDataReplicationFactor() bool {
	if o != nil && o.DataReplicationFactor != nil {
		return true
	}

	return false
}

// SetDataReplicationFactor gets a reference to the given string and assigns it to the DataReplicationFactor field.
func (o *HyperflexHxResiliencyInfoDtAllOf) SetDataReplicationFactor(v string) {
	o.DataReplicationFactor = &v
}

// GetHddFailuresTolerable returns the HddFailuresTolerable field value if set, zero value otherwise.
func (o *HyperflexHxResiliencyInfoDtAllOf) GetHddFailuresTolerable() int64 {
	if o == nil || o.HddFailuresTolerable == nil {
		var ret int64
		return ret
	}
	return *o.HddFailuresTolerable
}

// GetHddFailuresTolerableOk returns a tuple with the HddFailuresTolerable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxResiliencyInfoDtAllOf) GetHddFailuresTolerableOk() (*int64, bool) {
	if o == nil || o.HddFailuresTolerable == nil {
		return nil, false
	}
	return o.HddFailuresTolerable, true
}

// HasHddFailuresTolerable returns a boolean if a field has been set.
func (o *HyperflexHxResiliencyInfoDtAllOf) HasHddFailuresTolerable() bool {
	if o != nil && o.HddFailuresTolerable != nil {
		return true
	}

	return false
}

// SetHddFailuresTolerable gets a reference to the given int64 and assigns it to the HddFailuresTolerable field.
func (o *HyperflexHxResiliencyInfoDtAllOf) SetHddFailuresTolerable(v int64) {
	o.HddFailuresTolerable = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHxResiliencyInfoDtAllOf) GetMessages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHxResiliencyInfoDtAllOf) GetMessagesOk() ([]string, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *HyperflexHxResiliencyInfoDtAllOf) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *HyperflexHxResiliencyInfoDtAllOf) SetMessages(v []string) {
	o.Messages = v
}

// GetNodeFailuresTolerable returns the NodeFailuresTolerable field value if set, zero value otherwise.
func (o *HyperflexHxResiliencyInfoDtAllOf) GetNodeFailuresTolerable() int64 {
	if o == nil || o.NodeFailuresTolerable == nil {
		var ret int64
		return ret
	}
	return *o.NodeFailuresTolerable
}

// GetNodeFailuresTolerableOk returns a tuple with the NodeFailuresTolerable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxResiliencyInfoDtAllOf) GetNodeFailuresTolerableOk() (*int64, bool) {
	if o == nil || o.NodeFailuresTolerable == nil {
		return nil, false
	}
	return o.NodeFailuresTolerable, true
}

// HasNodeFailuresTolerable returns a boolean if a field has been set.
func (o *HyperflexHxResiliencyInfoDtAllOf) HasNodeFailuresTolerable() bool {
	if o != nil && o.NodeFailuresTolerable != nil {
		return true
	}

	return false
}

// SetNodeFailuresTolerable gets a reference to the given int64 and assigns it to the NodeFailuresTolerable field.
func (o *HyperflexHxResiliencyInfoDtAllOf) SetNodeFailuresTolerable(v int64) {
	o.NodeFailuresTolerable = &v
}

// GetPolicyCompliance returns the PolicyCompliance field value if set, zero value otherwise.
func (o *HyperflexHxResiliencyInfoDtAllOf) GetPolicyCompliance() string {
	if o == nil || o.PolicyCompliance == nil {
		var ret string
		return ret
	}
	return *o.PolicyCompliance
}

// GetPolicyComplianceOk returns a tuple with the PolicyCompliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxResiliencyInfoDtAllOf) GetPolicyComplianceOk() (*string, bool) {
	if o == nil || o.PolicyCompliance == nil {
		return nil, false
	}
	return o.PolicyCompliance, true
}

// HasPolicyCompliance returns a boolean if a field has been set.
func (o *HyperflexHxResiliencyInfoDtAllOf) HasPolicyCompliance() bool {
	if o != nil && o.PolicyCompliance != nil {
		return true
	}

	return false
}

// SetPolicyCompliance gets a reference to the given string and assigns it to the PolicyCompliance field.
func (o *HyperflexHxResiliencyInfoDtAllOf) SetPolicyCompliance(v string) {
	o.PolicyCompliance = &v
}

// GetResiliencyState returns the ResiliencyState field value if set, zero value otherwise.
func (o *HyperflexHxResiliencyInfoDtAllOf) GetResiliencyState() string {
	if o == nil || o.ResiliencyState == nil {
		var ret string
		return ret
	}
	return *o.ResiliencyState
}

// GetResiliencyStateOk returns a tuple with the ResiliencyState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxResiliencyInfoDtAllOf) GetResiliencyStateOk() (*string, bool) {
	if o == nil || o.ResiliencyState == nil {
		return nil, false
	}
	return o.ResiliencyState, true
}

// HasResiliencyState returns a boolean if a field has been set.
func (o *HyperflexHxResiliencyInfoDtAllOf) HasResiliencyState() bool {
	if o != nil && o.ResiliencyState != nil {
		return true
	}

	return false
}

// SetResiliencyState gets a reference to the given string and assigns it to the ResiliencyState field.
func (o *HyperflexHxResiliencyInfoDtAllOf) SetResiliencyState(v string) {
	o.ResiliencyState = &v
}

// GetSsdFailuresTolerable returns the SsdFailuresTolerable field value if set, zero value otherwise.
func (o *HyperflexHxResiliencyInfoDtAllOf) GetSsdFailuresTolerable() int64 {
	if o == nil || o.SsdFailuresTolerable == nil {
		var ret int64
		return ret
	}
	return *o.SsdFailuresTolerable
}

// GetSsdFailuresTolerableOk returns a tuple with the SsdFailuresTolerable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxResiliencyInfoDtAllOf) GetSsdFailuresTolerableOk() (*int64, bool) {
	if o == nil || o.SsdFailuresTolerable == nil {
		return nil, false
	}
	return o.SsdFailuresTolerable, true
}

// HasSsdFailuresTolerable returns a boolean if a field has been set.
func (o *HyperflexHxResiliencyInfoDtAllOf) HasSsdFailuresTolerable() bool {
	if o != nil && o.SsdFailuresTolerable != nil {
		return true
	}

	return false
}

// SetSsdFailuresTolerable gets a reference to the given int64 and assigns it to the SsdFailuresTolerable field.
func (o *HyperflexHxResiliencyInfoDtAllOf) SetSsdFailuresTolerable(v int64) {
	o.SsdFailuresTolerable = &v
}

func (o HyperflexHxResiliencyInfoDtAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DataReplicationFactor != nil {
		toSerialize["DataReplicationFactor"] = o.DataReplicationFactor
	}
	if o.HddFailuresTolerable != nil {
		toSerialize["HddFailuresTolerable"] = o.HddFailuresTolerable
	}
	if o.Messages != nil {
		toSerialize["Messages"] = o.Messages
	}
	if o.NodeFailuresTolerable != nil {
		toSerialize["NodeFailuresTolerable"] = o.NodeFailuresTolerable
	}
	if o.PolicyCompliance != nil {
		toSerialize["PolicyCompliance"] = o.PolicyCompliance
	}
	if o.ResiliencyState != nil {
		toSerialize["ResiliencyState"] = o.ResiliencyState
	}
	if o.SsdFailuresTolerable != nil {
		toSerialize["SsdFailuresTolerable"] = o.SsdFailuresTolerable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxResiliencyInfoDtAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHxResiliencyInfoDtAllOf := _HyperflexHxResiliencyInfoDtAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexHxResiliencyInfoDtAllOf); err == nil {
		*o = HyperflexHxResiliencyInfoDtAllOf(varHyperflexHxResiliencyInfoDtAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DataReplicationFactor")
		delete(additionalProperties, "HddFailuresTolerable")
		delete(additionalProperties, "Messages")
		delete(additionalProperties, "NodeFailuresTolerable")
		delete(additionalProperties, "PolicyCompliance")
		delete(additionalProperties, "ResiliencyState")
		delete(additionalProperties, "SsdFailuresTolerable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHxResiliencyInfoDtAllOf struct {
	value *HyperflexHxResiliencyInfoDtAllOf
	isSet bool
}

func (v NullableHyperflexHxResiliencyInfoDtAllOf) Get() *HyperflexHxResiliencyInfoDtAllOf {
	return v.value
}

func (v *NullableHyperflexHxResiliencyInfoDtAllOf) Set(val *HyperflexHxResiliencyInfoDtAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxResiliencyInfoDtAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxResiliencyInfoDtAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxResiliencyInfoDtAllOf(val *HyperflexHxResiliencyInfoDtAllOf) *NullableHyperflexHxResiliencyInfoDtAllOf {
	return &NullableHyperflexHxResiliencyInfoDtAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHxResiliencyInfoDtAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxResiliencyInfoDtAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
