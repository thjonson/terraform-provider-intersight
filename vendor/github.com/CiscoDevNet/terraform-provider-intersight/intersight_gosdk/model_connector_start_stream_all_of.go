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

// ConnectorStartStreamAllOf Definition of the list of properties defined in 'connector.StartStream', excluding properties defined in parent classes.
type ConnectorStartStreamAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of outputs from a plugin to collect into a single message. Applicable only to streams that involve polling plugins and plugins which support emitting batchable data. Default value of zero indicates no batching.
	BatchSize *int64 `json:"BatchSize,omitempty"`
	// Flag to force a rebuild of an existing stream. To be used if a stream is unable to recover itself in response to dropped messages.
	ForceRebuild *bool `json:"ForceRebuild,omitempty"`
	// Input to the plugin to start the start the stream or collect stream messages.
	Input *string `json:"Input,omitempty"`
	// Interval at which device should emit a keepalive message for this stream. Device will also expect a keepalive response from the cloud within the interval. If zero, no keepalive is required and stream should not timeout.
	KeepAliveInterval *int64 `json:"KeepAliveInterval,omitempty"`
	// The plugin to run the stream on.
	PluginName *string `json:"PluginName,omitempty"`
	// The desired interval to emit messages from this stream. The stream plugin will poll plugins at this interval to create a stream event.
	PollInterval *int64 `json:"PollInterval,omitempty"`
	// The priority level to apply to messages emitted by this stream.
	Priority *int64 `json:"Priority,omitempty"`
	// The version of the device connector stream protocol. Used to change behavior of the device connector stream plugin based on the version of the Intersight service. Allows for multiple versions of Intersight services to interact with the stream plugin of devices.
	ProtocolVersion *int64 `json:"ProtocolVersion,omitempty"`
	// The topic for the device connector to publish messages to.
	ResponseTopic        *string `json:"ResponseTopic,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorStartStreamAllOf ConnectorStartStreamAllOf

// NewConnectorStartStreamAllOf instantiates a new ConnectorStartStreamAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorStartStreamAllOf(classId string, objectType string) *ConnectorStartStreamAllOf {
	this := ConnectorStartStreamAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorStartStreamAllOfWithDefaults instantiates a new ConnectorStartStreamAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorStartStreamAllOfWithDefaults() *ConnectorStartStreamAllOf {
	this := ConnectorStartStreamAllOf{}
	var classId string = "connector.StartStream"
	this.ClassId = classId
	var objectType string = "connector.StartStream"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorStartStreamAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorStartStreamAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorStartStreamAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorStartStreamAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorStartStreamAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorStartStreamAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBatchSize returns the BatchSize field value if set, zero value otherwise.
func (o *ConnectorStartStreamAllOf) GetBatchSize() int64 {
	if o == nil || o.BatchSize == nil {
		var ret int64
		return ret
	}
	return *o.BatchSize
}

// GetBatchSizeOk returns a tuple with the BatchSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorStartStreamAllOf) GetBatchSizeOk() (*int64, bool) {
	if o == nil || o.BatchSize == nil {
		return nil, false
	}
	return o.BatchSize, true
}

// HasBatchSize returns a boolean if a field has been set.
func (o *ConnectorStartStreamAllOf) HasBatchSize() bool {
	if o != nil && o.BatchSize != nil {
		return true
	}

	return false
}

// SetBatchSize gets a reference to the given int64 and assigns it to the BatchSize field.
func (o *ConnectorStartStreamAllOf) SetBatchSize(v int64) {
	o.BatchSize = &v
}

// GetForceRebuild returns the ForceRebuild field value if set, zero value otherwise.
func (o *ConnectorStartStreamAllOf) GetForceRebuild() bool {
	if o == nil || o.ForceRebuild == nil {
		var ret bool
		return ret
	}
	return *o.ForceRebuild
}

// GetForceRebuildOk returns a tuple with the ForceRebuild field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorStartStreamAllOf) GetForceRebuildOk() (*bool, bool) {
	if o == nil || o.ForceRebuild == nil {
		return nil, false
	}
	return o.ForceRebuild, true
}

// HasForceRebuild returns a boolean if a field has been set.
func (o *ConnectorStartStreamAllOf) HasForceRebuild() bool {
	if o != nil && o.ForceRebuild != nil {
		return true
	}

	return false
}

// SetForceRebuild gets a reference to the given bool and assigns it to the ForceRebuild field.
func (o *ConnectorStartStreamAllOf) SetForceRebuild(v bool) {
	o.ForceRebuild = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *ConnectorStartStreamAllOf) GetInput() string {
	if o == nil || o.Input == nil {
		var ret string
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorStartStreamAllOf) GetInputOk() (*string, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *ConnectorStartStreamAllOf) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given string and assigns it to the Input field.
func (o *ConnectorStartStreamAllOf) SetInput(v string) {
	o.Input = &v
}

// GetKeepAliveInterval returns the KeepAliveInterval field value if set, zero value otherwise.
func (o *ConnectorStartStreamAllOf) GetKeepAliveInterval() int64 {
	if o == nil || o.KeepAliveInterval == nil {
		var ret int64
		return ret
	}
	return *o.KeepAliveInterval
}

// GetKeepAliveIntervalOk returns a tuple with the KeepAliveInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorStartStreamAllOf) GetKeepAliveIntervalOk() (*int64, bool) {
	if o == nil || o.KeepAliveInterval == nil {
		return nil, false
	}
	return o.KeepAliveInterval, true
}

// HasKeepAliveInterval returns a boolean if a field has been set.
func (o *ConnectorStartStreamAllOf) HasKeepAliveInterval() bool {
	if o != nil && o.KeepAliveInterval != nil {
		return true
	}

	return false
}

// SetKeepAliveInterval gets a reference to the given int64 and assigns it to the KeepAliveInterval field.
func (o *ConnectorStartStreamAllOf) SetKeepAliveInterval(v int64) {
	o.KeepAliveInterval = &v
}

// GetPluginName returns the PluginName field value if set, zero value otherwise.
func (o *ConnectorStartStreamAllOf) GetPluginName() string {
	if o == nil || o.PluginName == nil {
		var ret string
		return ret
	}
	return *o.PluginName
}

// GetPluginNameOk returns a tuple with the PluginName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorStartStreamAllOf) GetPluginNameOk() (*string, bool) {
	if o == nil || o.PluginName == nil {
		return nil, false
	}
	return o.PluginName, true
}

// HasPluginName returns a boolean if a field has been set.
func (o *ConnectorStartStreamAllOf) HasPluginName() bool {
	if o != nil && o.PluginName != nil {
		return true
	}

	return false
}

// SetPluginName gets a reference to the given string and assigns it to the PluginName field.
func (o *ConnectorStartStreamAllOf) SetPluginName(v string) {
	o.PluginName = &v
}

// GetPollInterval returns the PollInterval field value if set, zero value otherwise.
func (o *ConnectorStartStreamAllOf) GetPollInterval() int64 {
	if o == nil || o.PollInterval == nil {
		var ret int64
		return ret
	}
	return *o.PollInterval
}

// GetPollIntervalOk returns a tuple with the PollInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorStartStreamAllOf) GetPollIntervalOk() (*int64, bool) {
	if o == nil || o.PollInterval == nil {
		return nil, false
	}
	return o.PollInterval, true
}

// HasPollInterval returns a boolean if a field has been set.
func (o *ConnectorStartStreamAllOf) HasPollInterval() bool {
	if o != nil && o.PollInterval != nil {
		return true
	}

	return false
}

// SetPollInterval gets a reference to the given int64 and assigns it to the PollInterval field.
func (o *ConnectorStartStreamAllOf) SetPollInterval(v int64) {
	o.PollInterval = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ConnectorStartStreamAllOf) GetPriority() int64 {
	if o == nil || o.Priority == nil {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorStartStreamAllOf) GetPriorityOk() (*int64, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ConnectorStartStreamAllOf) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *ConnectorStartStreamAllOf) SetPriority(v int64) {
	o.Priority = &v
}

// GetProtocolVersion returns the ProtocolVersion field value if set, zero value otherwise.
func (o *ConnectorStartStreamAllOf) GetProtocolVersion() int64 {
	if o == nil || o.ProtocolVersion == nil {
		var ret int64
		return ret
	}
	return *o.ProtocolVersion
}

// GetProtocolVersionOk returns a tuple with the ProtocolVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorStartStreamAllOf) GetProtocolVersionOk() (*int64, bool) {
	if o == nil || o.ProtocolVersion == nil {
		return nil, false
	}
	return o.ProtocolVersion, true
}

// HasProtocolVersion returns a boolean if a field has been set.
func (o *ConnectorStartStreamAllOf) HasProtocolVersion() bool {
	if o != nil && o.ProtocolVersion != nil {
		return true
	}

	return false
}

// SetProtocolVersion gets a reference to the given int64 and assigns it to the ProtocolVersion field.
func (o *ConnectorStartStreamAllOf) SetProtocolVersion(v int64) {
	o.ProtocolVersion = &v
}

// GetResponseTopic returns the ResponseTopic field value if set, zero value otherwise.
func (o *ConnectorStartStreamAllOf) GetResponseTopic() string {
	if o == nil || o.ResponseTopic == nil {
		var ret string
		return ret
	}
	return *o.ResponseTopic
}

// GetResponseTopicOk returns a tuple with the ResponseTopic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorStartStreamAllOf) GetResponseTopicOk() (*string, bool) {
	if o == nil || o.ResponseTopic == nil {
		return nil, false
	}
	return o.ResponseTopic, true
}

// HasResponseTopic returns a boolean if a field has been set.
func (o *ConnectorStartStreamAllOf) HasResponseTopic() bool {
	if o != nil && o.ResponseTopic != nil {
		return true
	}

	return false
}

// SetResponseTopic gets a reference to the given string and assigns it to the ResponseTopic field.
func (o *ConnectorStartStreamAllOf) SetResponseTopic(v string) {
	o.ResponseTopic = &v
}

func (o ConnectorStartStreamAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BatchSize != nil {
		toSerialize["BatchSize"] = o.BatchSize
	}
	if o.ForceRebuild != nil {
		toSerialize["ForceRebuild"] = o.ForceRebuild
	}
	if o.Input != nil {
		toSerialize["Input"] = o.Input
	}
	if o.KeepAliveInterval != nil {
		toSerialize["KeepAliveInterval"] = o.KeepAliveInterval
	}
	if o.PluginName != nil {
		toSerialize["PluginName"] = o.PluginName
	}
	if o.PollInterval != nil {
		toSerialize["PollInterval"] = o.PollInterval
	}
	if o.Priority != nil {
		toSerialize["Priority"] = o.Priority
	}
	if o.ProtocolVersion != nil {
		toSerialize["ProtocolVersion"] = o.ProtocolVersion
	}
	if o.ResponseTopic != nil {
		toSerialize["ResponseTopic"] = o.ResponseTopic
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorStartStreamAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConnectorStartStreamAllOf := _ConnectorStartStreamAllOf{}

	if err = json.Unmarshal(bytes, &varConnectorStartStreamAllOf); err == nil {
		*o = ConnectorStartStreamAllOf(varConnectorStartStreamAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BatchSize")
		delete(additionalProperties, "ForceRebuild")
		delete(additionalProperties, "Input")
		delete(additionalProperties, "KeepAliveInterval")
		delete(additionalProperties, "PluginName")
		delete(additionalProperties, "PollInterval")
		delete(additionalProperties, "Priority")
		delete(additionalProperties, "ProtocolVersion")
		delete(additionalProperties, "ResponseTopic")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorStartStreamAllOf struct {
	value *ConnectorStartStreamAllOf
	isSet bool
}

func (v NullableConnectorStartStreamAllOf) Get() *ConnectorStartStreamAllOf {
	return v.value
}

func (v *NullableConnectorStartStreamAllOf) Set(val *ConnectorStartStreamAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorStartStreamAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorStartStreamAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorStartStreamAllOf(val *ConnectorStartStreamAllOf) *NullableConnectorStartStreamAllOf {
	return &NullableConnectorStartStreamAllOf{value: val, isSet: true}
}

func (v NullableConnectorStartStreamAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorStartStreamAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
