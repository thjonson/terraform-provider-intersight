# FabricUplinkPcRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.UplinkPcRole"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.UplinkPcRole"]
**AdminSpeed** | Pointer to **string** | Admin configured speed for the port. * &#x60;Auto&#x60; - Admin configurable speed AUTO ( default ). * &#x60;1Gbps&#x60; - Admin configurable speed 1Gbps. * &#x60;10Gbps&#x60; - Admin configurable speed 10Gbps. * &#x60;25Gbps&#x60; - Admin configurable speed 25Gbps. * &#x60;40Gbps&#x60; - Admin configurable speed 40Gbps. * &#x60;100Gbps&#x60; - Admin configurable speed 100Gbps. | [optional] [default to "Auto"]
**FlowControlPolicy** | Pointer to [**FabricFlowControlPolicyRelationship**](FabricFlowControlPolicyRelationship.md) |  | [optional] 
**LinkAggregationPolicy** | Pointer to [**FabricLinkAggregationPolicyRelationship**](FabricLinkAggregationPolicyRelationship.md) |  | [optional] 
**LinkControlPolicy** | Pointer to [**FabricLinkControlPolicyRelationship**](FabricLinkControlPolicyRelationship.md) |  | [optional] 

## Methods

### NewFabricUplinkPcRole

`func NewFabricUplinkPcRole(classId string, objectType string, ) *FabricUplinkPcRole`

NewFabricUplinkPcRole instantiates a new FabricUplinkPcRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricUplinkPcRoleWithDefaults

`func NewFabricUplinkPcRoleWithDefaults() *FabricUplinkPcRole`

NewFabricUplinkPcRoleWithDefaults instantiates a new FabricUplinkPcRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricUplinkPcRole) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricUplinkPcRole) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricUplinkPcRole) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricUplinkPcRole) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricUplinkPcRole) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricUplinkPcRole) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminSpeed

`func (o *FabricUplinkPcRole) GetAdminSpeed() string`

GetAdminSpeed returns the AdminSpeed field if non-nil, zero value otherwise.

### GetAdminSpeedOk

`func (o *FabricUplinkPcRole) GetAdminSpeedOk() (*string, bool)`

GetAdminSpeedOk returns a tuple with the AdminSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSpeed

`func (o *FabricUplinkPcRole) SetAdminSpeed(v string)`

SetAdminSpeed sets AdminSpeed field to given value.

### HasAdminSpeed

`func (o *FabricUplinkPcRole) HasAdminSpeed() bool`

HasAdminSpeed returns a boolean if a field has been set.

### GetFlowControlPolicy

`func (o *FabricUplinkPcRole) GetFlowControlPolicy() FabricFlowControlPolicyRelationship`

GetFlowControlPolicy returns the FlowControlPolicy field if non-nil, zero value otherwise.

### GetFlowControlPolicyOk

`func (o *FabricUplinkPcRole) GetFlowControlPolicyOk() (*FabricFlowControlPolicyRelationship, bool)`

GetFlowControlPolicyOk returns a tuple with the FlowControlPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControlPolicy

`func (o *FabricUplinkPcRole) SetFlowControlPolicy(v FabricFlowControlPolicyRelationship)`

SetFlowControlPolicy sets FlowControlPolicy field to given value.

### HasFlowControlPolicy

`func (o *FabricUplinkPcRole) HasFlowControlPolicy() bool`

HasFlowControlPolicy returns a boolean if a field has been set.

### GetLinkAggregationPolicy

`func (o *FabricUplinkPcRole) GetLinkAggregationPolicy() FabricLinkAggregationPolicyRelationship`

GetLinkAggregationPolicy returns the LinkAggregationPolicy field if non-nil, zero value otherwise.

### GetLinkAggregationPolicyOk

`func (o *FabricUplinkPcRole) GetLinkAggregationPolicyOk() (*FabricLinkAggregationPolicyRelationship, bool)`

GetLinkAggregationPolicyOk returns a tuple with the LinkAggregationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAggregationPolicy

`func (o *FabricUplinkPcRole) SetLinkAggregationPolicy(v FabricLinkAggregationPolicyRelationship)`

SetLinkAggregationPolicy sets LinkAggregationPolicy field to given value.

### HasLinkAggregationPolicy

`func (o *FabricUplinkPcRole) HasLinkAggregationPolicy() bool`

HasLinkAggregationPolicy returns a boolean if a field has been set.

### GetLinkControlPolicy

`func (o *FabricUplinkPcRole) GetLinkControlPolicy() FabricLinkControlPolicyRelationship`

GetLinkControlPolicy returns the LinkControlPolicy field if non-nil, zero value otherwise.

### GetLinkControlPolicyOk

`func (o *FabricUplinkPcRole) GetLinkControlPolicyOk() (*FabricLinkControlPolicyRelationship, bool)`

GetLinkControlPolicyOk returns a tuple with the LinkControlPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkControlPolicy

`func (o *FabricUplinkPcRole) SetLinkControlPolicy(v FabricLinkControlPolicyRelationship)`

SetLinkControlPolicy sets LinkControlPolicy field to given value.

### HasLinkControlPolicy

`func (o *FabricUplinkPcRole) HasLinkControlPolicy() bool`

HasLinkControlPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


