/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-07-21T16:37:30Z.
 *
 * API version: 1.0.9-4403
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// NotificationAccountSubscriptionAllOf Definition of the list of properties defined in 'notification.AccountSubscription', excluding properties defined in parent classes.
type NotificationAccountSubscriptionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The subscription name.
	Name                 *string                 `json:"Name,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotificationAccountSubscriptionAllOf NotificationAccountSubscriptionAllOf

// NewNotificationAccountSubscriptionAllOf instantiates a new NotificationAccountSubscriptionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationAccountSubscriptionAllOf(classId string, objectType string) *NotificationAccountSubscriptionAllOf {
	this := NotificationAccountSubscriptionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNotificationAccountSubscriptionAllOfWithDefaults instantiates a new NotificationAccountSubscriptionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationAccountSubscriptionAllOfWithDefaults() *NotificationAccountSubscriptionAllOf {
	this := NotificationAccountSubscriptionAllOf{}
	var classId string = "notification.AccountSubscription"
	this.ClassId = classId
	var objectType string = "notification.AccountSubscription"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NotificationAccountSubscriptionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NotificationAccountSubscriptionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NotificationAccountSubscriptionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NotificationAccountSubscriptionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NotificationAccountSubscriptionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NotificationAccountSubscriptionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NotificationAccountSubscriptionAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationAccountSubscriptionAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NotificationAccountSubscriptionAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NotificationAccountSubscriptionAllOf) SetName(v string) {
	o.Name = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *NotificationAccountSubscriptionAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationAccountSubscriptionAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *NotificationAccountSubscriptionAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *NotificationAccountSubscriptionAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o NotificationAccountSubscriptionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NotificationAccountSubscriptionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNotificationAccountSubscriptionAllOf := _NotificationAccountSubscriptionAllOf{}

	if err = json.Unmarshal(bytes, &varNotificationAccountSubscriptionAllOf); err == nil {
		*o = NotificationAccountSubscriptionAllOf(varNotificationAccountSubscriptionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotificationAccountSubscriptionAllOf struct {
	value *NotificationAccountSubscriptionAllOf
	isSet bool
}

func (v NullableNotificationAccountSubscriptionAllOf) Get() *NotificationAccountSubscriptionAllOf {
	return v.value
}

func (v *NullableNotificationAccountSubscriptionAllOf) Set(val *NotificationAccountSubscriptionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationAccountSubscriptionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationAccountSubscriptionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationAccountSubscriptionAllOf(val *NotificationAccountSubscriptionAllOf) *NullableNotificationAccountSubscriptionAllOf {
	return &NullableNotificationAccountSubscriptionAllOf{value: val, isSet: true}
}

func (v NullableNotificationAccountSubscriptionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationAccountSubscriptionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
