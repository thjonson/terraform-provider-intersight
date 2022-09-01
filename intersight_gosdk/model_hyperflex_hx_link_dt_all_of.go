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

// HyperflexHxLinkDtAllOf Definition of the list of properties defined in 'hyperflex.HxLinkDt', excluding properties defined in parent classes.
type HyperflexHxLinkDtAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Comment for this HyperFlex resource.
	Comments *string `json:"Comments,omitempty"`
	// URI of resource. Target URL for making REST call.
	Href *string `json:"Href,omitempty"`
	// HTTP verb that this HyperFlex link DT is referring to. * `POST` - HTTP verb POST for this task definition. * `GET` - HTTP verb GET for this task definition. * `PUT` - HTTP verb PUT for this task definition. * `DELETE` - HTTP verb DELETE for this task definition.
	Method *string `json:"Method,omitempty"`
	// Relationship of  link to this resource.
	Rel                  *string `json:"Rel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxLinkDtAllOf HyperflexHxLinkDtAllOf

// NewHyperflexHxLinkDtAllOf instantiates a new HyperflexHxLinkDtAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxLinkDtAllOf(classId string, objectType string) *HyperflexHxLinkDtAllOf {
	this := HyperflexHxLinkDtAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHxLinkDtAllOfWithDefaults instantiates a new HyperflexHxLinkDtAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxLinkDtAllOfWithDefaults() *HyperflexHxLinkDtAllOf {
	this := HyperflexHxLinkDtAllOf{}
	var classId string = "hyperflex.HxLinkDt"
	this.ClassId = classId
	var objectType string = "hyperflex.HxLinkDt"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxLinkDtAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxLinkDtAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxLinkDtAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxLinkDtAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxLinkDtAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxLinkDtAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *HyperflexHxLinkDtAllOf) GetComments() string {
	if o == nil || o.Comments == nil {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxLinkDtAllOf) GetCommentsOk() (*string, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *HyperflexHxLinkDtAllOf) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *HyperflexHxLinkDtAllOf) SetComments(v string) {
	o.Comments = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *HyperflexHxLinkDtAllOf) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxLinkDtAllOf) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *HyperflexHxLinkDtAllOf) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *HyperflexHxLinkDtAllOf) SetHref(v string) {
	o.Href = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *HyperflexHxLinkDtAllOf) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxLinkDtAllOf) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *HyperflexHxLinkDtAllOf) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *HyperflexHxLinkDtAllOf) SetMethod(v string) {
	o.Method = &v
}

// GetRel returns the Rel field value if set, zero value otherwise.
func (o *HyperflexHxLinkDtAllOf) GetRel() string {
	if o == nil || o.Rel == nil {
		var ret string
		return ret
	}
	return *o.Rel
}

// GetRelOk returns a tuple with the Rel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxLinkDtAllOf) GetRelOk() (*string, bool) {
	if o == nil || o.Rel == nil {
		return nil, false
	}
	return o.Rel, true
}

// HasRel returns a boolean if a field has been set.
func (o *HyperflexHxLinkDtAllOf) HasRel() bool {
	if o != nil && o.Rel != nil {
		return true
	}

	return false
}

// SetRel gets a reference to the given string and assigns it to the Rel field.
func (o *HyperflexHxLinkDtAllOf) SetRel(v string) {
	o.Rel = &v
}

func (o HyperflexHxLinkDtAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Comments != nil {
		toSerialize["Comments"] = o.Comments
	}
	if o.Href != nil {
		toSerialize["Href"] = o.Href
	}
	if o.Method != nil {
		toSerialize["Method"] = o.Method
	}
	if o.Rel != nil {
		toSerialize["Rel"] = o.Rel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxLinkDtAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHxLinkDtAllOf := _HyperflexHxLinkDtAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexHxLinkDtAllOf); err == nil {
		*o = HyperflexHxLinkDtAllOf(varHyperflexHxLinkDtAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Comments")
		delete(additionalProperties, "Href")
		delete(additionalProperties, "Method")
		delete(additionalProperties, "Rel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHxLinkDtAllOf struct {
	value *HyperflexHxLinkDtAllOf
	isSet bool
}

func (v NullableHyperflexHxLinkDtAllOf) Get() *HyperflexHxLinkDtAllOf {
	return v.value
}

func (v *NullableHyperflexHxLinkDtAllOf) Set(val *HyperflexHxLinkDtAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxLinkDtAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxLinkDtAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxLinkDtAllOf(val *HyperflexHxLinkDtAllOf) *NullableHyperflexHxLinkDtAllOf {
	return &NullableHyperflexHxLinkDtAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHxLinkDtAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxLinkDtAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
