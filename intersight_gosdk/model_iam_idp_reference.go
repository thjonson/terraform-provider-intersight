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

// IamIdpReference Default Cisco IdP for authentication.
type IamIdpReference struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The email domain name for this IdP of the user. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication.
	DomainName *string `json:"DomainName,omitempty"`
	// Entity ID of the IdP. In SAML, the entity ID uniquely identifies the IdP/Service Provider.
	IdpEntityId *string `json:"IdpEntityId,omitempty"`
	// The flag represents if the second factor of authentication is required for Cisco IdP users.
	MultiFactorAuthentication *bool `json:"MultiFactorAuthentication,omitempty"`
	// Cisco IdP reference in an account.
	Name    *string                 `json:"Name,omitempty"`
	Account *IamAccountRelationship `json:"Account,omitempty"`
	Idp     *IamIdpRelationship     `json:"Idp,omitempty"`
	// An array of relationships to iamUserPreference resources.
	UserPreferences []IamUserPreferenceRelationship `json:"UserPreferences,omitempty"`
	// An array of relationships to iamUserGroup resources.
	Usergroups []IamUserGroupRelationship `json:"Usergroups,omitempty"`
	// An array of relationships to iamUser resources.
	Users                []IamUserRelationship `json:"Users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamIdpReference IamIdpReference

// NewIamIdpReference instantiates a new IamIdpReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamIdpReference(classId string, objectType string) *IamIdpReference {
	this := IamIdpReference{}
	this.ClassId = classId
	this.ObjectType = objectType
	var multiFactorAuthentication bool = false
	this.MultiFactorAuthentication = &multiFactorAuthentication
	return &this
}

// NewIamIdpReferenceWithDefaults instantiates a new IamIdpReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamIdpReferenceWithDefaults() *IamIdpReference {
	this := IamIdpReference{}
	var classId string = "iam.IdpReference"
	this.ClassId = classId
	var objectType string = "iam.IdpReference"
	this.ObjectType = objectType
	var multiFactorAuthentication bool = false
	this.MultiFactorAuthentication = &multiFactorAuthentication
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamIdpReference) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamIdpReference) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamIdpReference) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamIdpReference) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamIdpReference) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamIdpReference) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *IamIdpReference) GetDomainName() string {
	if o == nil || o.DomainName == nil {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdpReference) GetDomainNameOk() (*string, bool) {
	if o == nil || o.DomainName == nil {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *IamIdpReference) HasDomainName() bool {
	if o != nil && o.DomainName != nil {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *IamIdpReference) SetDomainName(v string) {
	o.DomainName = &v
}

// GetIdpEntityId returns the IdpEntityId field value if set, zero value otherwise.
func (o *IamIdpReference) GetIdpEntityId() string {
	if o == nil || o.IdpEntityId == nil {
		var ret string
		return ret
	}
	return *o.IdpEntityId
}

// GetIdpEntityIdOk returns a tuple with the IdpEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdpReference) GetIdpEntityIdOk() (*string, bool) {
	if o == nil || o.IdpEntityId == nil {
		return nil, false
	}
	return o.IdpEntityId, true
}

// HasIdpEntityId returns a boolean if a field has been set.
func (o *IamIdpReference) HasIdpEntityId() bool {
	if o != nil && o.IdpEntityId != nil {
		return true
	}

	return false
}

// SetIdpEntityId gets a reference to the given string and assigns it to the IdpEntityId field.
func (o *IamIdpReference) SetIdpEntityId(v string) {
	o.IdpEntityId = &v
}

// GetMultiFactorAuthentication returns the MultiFactorAuthentication field value if set, zero value otherwise.
func (o *IamIdpReference) GetMultiFactorAuthentication() bool {
	if o == nil || o.MultiFactorAuthentication == nil {
		var ret bool
		return ret
	}
	return *o.MultiFactorAuthentication
}

// GetMultiFactorAuthenticationOk returns a tuple with the MultiFactorAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdpReference) GetMultiFactorAuthenticationOk() (*bool, bool) {
	if o == nil || o.MultiFactorAuthentication == nil {
		return nil, false
	}
	return o.MultiFactorAuthentication, true
}

// HasMultiFactorAuthentication returns a boolean if a field has been set.
func (o *IamIdpReference) HasMultiFactorAuthentication() bool {
	if o != nil && o.MultiFactorAuthentication != nil {
		return true
	}

	return false
}

// SetMultiFactorAuthentication gets a reference to the given bool and assigns it to the MultiFactorAuthentication field.
func (o *IamIdpReference) SetMultiFactorAuthentication(v bool) {
	o.MultiFactorAuthentication = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamIdpReference) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdpReference) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamIdpReference) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamIdpReference) SetName(v string) {
	o.Name = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamIdpReference) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdpReference) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamIdpReference) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamIdpReference) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *IamIdpReference) GetIdp() IamIdpRelationship {
	if o == nil || o.Idp == nil {
		var ret IamIdpRelationship
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIdpReference) GetIdpOk() (*IamIdpRelationship, bool) {
	if o == nil || o.Idp == nil {
		return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *IamIdpReference) HasIdp() bool {
	if o != nil && o.Idp != nil {
		return true
	}

	return false
}

// SetIdp gets a reference to the given IamIdpRelationship and assigns it to the Idp field.
func (o *IamIdpReference) SetIdp(v IamIdpRelationship) {
	o.Idp = &v
}

// GetUserPreferences returns the UserPreferences field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamIdpReference) GetUserPreferences() []IamUserPreferenceRelationship {
	if o == nil {
		var ret []IamUserPreferenceRelationship
		return ret
	}
	return o.UserPreferences
}

// GetUserPreferencesOk returns a tuple with the UserPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamIdpReference) GetUserPreferencesOk() (*[]IamUserPreferenceRelationship, bool) {
	if o == nil || o.UserPreferences == nil {
		return nil, false
	}
	return &o.UserPreferences, true
}

// HasUserPreferences returns a boolean if a field has been set.
func (o *IamIdpReference) HasUserPreferences() bool {
	if o != nil && o.UserPreferences != nil {
		return true
	}

	return false
}

// SetUserPreferences gets a reference to the given []IamUserPreferenceRelationship and assigns it to the UserPreferences field.
func (o *IamIdpReference) SetUserPreferences(v []IamUserPreferenceRelationship) {
	o.UserPreferences = v
}

// GetUsergroups returns the Usergroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamIdpReference) GetUsergroups() []IamUserGroupRelationship {
	if o == nil {
		var ret []IamUserGroupRelationship
		return ret
	}
	return o.Usergroups
}

// GetUsergroupsOk returns a tuple with the Usergroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamIdpReference) GetUsergroupsOk() (*[]IamUserGroupRelationship, bool) {
	if o == nil || o.Usergroups == nil {
		return nil, false
	}
	return &o.Usergroups, true
}

// HasUsergroups returns a boolean if a field has been set.
func (o *IamIdpReference) HasUsergroups() bool {
	if o != nil && o.Usergroups != nil {
		return true
	}

	return false
}

// SetUsergroups gets a reference to the given []IamUserGroupRelationship and assigns it to the Usergroups field.
func (o *IamIdpReference) SetUsergroups(v []IamUserGroupRelationship) {
	o.Usergroups = v
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamIdpReference) GetUsers() []IamUserRelationship {
	if o == nil {
		var ret []IamUserRelationship
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamIdpReference) GetUsersOk() (*[]IamUserRelationship, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return &o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *IamIdpReference) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []IamUserRelationship and assigns it to the Users field.
func (o *IamIdpReference) SetUsers(v []IamUserRelationship) {
	o.Users = v
}

func (o IamIdpReference) MarshalJSON() ([]byte, error) {
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
	if o.DomainName != nil {
		toSerialize["DomainName"] = o.DomainName
	}
	if o.IdpEntityId != nil {
		toSerialize["IdpEntityId"] = o.IdpEntityId
	}
	if o.MultiFactorAuthentication != nil {
		toSerialize["MultiFactorAuthentication"] = o.MultiFactorAuthentication
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.Idp != nil {
		toSerialize["Idp"] = o.Idp
	}
	if o.UserPreferences != nil {
		toSerialize["UserPreferences"] = o.UserPreferences
	}
	if o.Usergroups != nil {
		toSerialize["Usergroups"] = o.Usergroups
	}
	if o.Users != nil {
		toSerialize["Users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamIdpReference) UnmarshalJSON(bytes []byte) (err error) {
	type IamIdpReferenceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The email domain name for this IdP of the user. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication.
		DomainName *string `json:"DomainName,omitempty"`
		// Entity ID of the IdP. In SAML, the entity ID uniquely identifies the IdP/Service Provider.
		IdpEntityId *string `json:"IdpEntityId,omitempty"`
		// The flag represents if the second factor of authentication is required for Cisco IdP users.
		MultiFactorAuthentication *bool `json:"MultiFactorAuthentication,omitempty"`
		// Cisco IdP reference in an account.
		Name    *string                 `json:"Name,omitempty"`
		Account *IamAccountRelationship `json:"Account,omitempty"`
		Idp     *IamIdpRelationship     `json:"Idp,omitempty"`
		// An array of relationships to iamUserPreference resources.
		UserPreferences []IamUserPreferenceRelationship `json:"UserPreferences,omitempty"`
		// An array of relationships to iamUserGroup resources.
		Usergroups []IamUserGroupRelationship `json:"Usergroups,omitempty"`
		// An array of relationships to iamUser resources.
		Users []IamUserRelationship `json:"Users,omitempty"`
	}

	varIamIdpReferenceWithoutEmbeddedStruct := IamIdpReferenceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamIdpReferenceWithoutEmbeddedStruct)
	if err == nil {
		varIamIdpReference := _IamIdpReference{}
		varIamIdpReference.ClassId = varIamIdpReferenceWithoutEmbeddedStruct.ClassId
		varIamIdpReference.ObjectType = varIamIdpReferenceWithoutEmbeddedStruct.ObjectType
		varIamIdpReference.DomainName = varIamIdpReferenceWithoutEmbeddedStruct.DomainName
		varIamIdpReference.IdpEntityId = varIamIdpReferenceWithoutEmbeddedStruct.IdpEntityId
		varIamIdpReference.MultiFactorAuthentication = varIamIdpReferenceWithoutEmbeddedStruct.MultiFactorAuthentication
		varIamIdpReference.Name = varIamIdpReferenceWithoutEmbeddedStruct.Name
		varIamIdpReference.Account = varIamIdpReferenceWithoutEmbeddedStruct.Account
		varIamIdpReference.Idp = varIamIdpReferenceWithoutEmbeddedStruct.Idp
		varIamIdpReference.UserPreferences = varIamIdpReferenceWithoutEmbeddedStruct.UserPreferences
		varIamIdpReference.Usergroups = varIamIdpReferenceWithoutEmbeddedStruct.Usergroups
		varIamIdpReference.Users = varIamIdpReferenceWithoutEmbeddedStruct.Users
		*o = IamIdpReference(varIamIdpReference)
	} else {
		return err
	}

	varIamIdpReference := _IamIdpReference{}

	err = json.Unmarshal(bytes, &varIamIdpReference)
	if err == nil {
		o.MoBaseMo = varIamIdpReference.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DomainName")
		delete(additionalProperties, "IdpEntityId")
		delete(additionalProperties, "MultiFactorAuthentication")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "Idp")
		delete(additionalProperties, "UserPreferences")
		delete(additionalProperties, "Usergroups")
		delete(additionalProperties, "Users")

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

type NullableIamIdpReference struct {
	value *IamIdpReference
	isSet bool
}

func (v NullableIamIdpReference) Get() *IamIdpReference {
	return v.value
}

func (v *NullableIamIdpReference) Set(val *IamIdpReference) {
	v.value = val
	v.isSet = true
}

func (v NullableIamIdpReference) IsSet() bool {
	return v.isSet
}

func (v *NullableIamIdpReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamIdpReference(val *IamIdpReference) *NullableIamIdpReference {
	return &NullableIamIdpReference{value: val, isSet: true}
}

func (v NullableIamIdpReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamIdpReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
