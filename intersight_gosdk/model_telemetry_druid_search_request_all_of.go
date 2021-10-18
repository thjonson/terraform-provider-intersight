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

// TelemetryDruidSearchRequestAllOf struct for TelemetryDruidSearchRequestAllOf
type TelemetryDruidSearchRequestAllOf struct {
	DataSource TelemetryDruidDataSource `json:"dataSource"`
	// A JSON Object representing ISO-8601 Intervals. This defines the time ranges to run the query over.
	Intervals   []string                  `json:"intervals"`
	Granularity TelemetryDruidGranularity `json:"granularity"`
	Filter      *TelemetryDruidFilter     `json:"filter,omitempty"`
	// Aggregation functions are used to summarize data in buckets. Summarization functions include counting rows, calculating the min/max/sum of metrics and retrieving the first/last value of metrics for each bucket. Additional summarization functions are available with extensions. If no aggregator is provided, the results will be empty for each bucket.
	Aggregations *[]TelemetryDruidAggregator `json:"aggregations,omitempty"`
	// The list of dimensions to run the search over. Excluding this means the search is run over all dimensions.
	SearchDimensions *[]string                          `json:"searchDimensions,omitempty"`
	Query            *TelemetryDruidAggregateSearchSpec `json:"query,omitempty"`
	// An integer that limits the number of results. The default is unlimited.
	Limit                *int32                      `json:"limit,omitempty"`
	Context              *TelemetryDruidQueryContext `json:"context,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidSearchRequestAllOf TelemetryDruidSearchRequestAllOf

// NewTelemetryDruidSearchRequestAllOf instantiates a new TelemetryDruidSearchRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidSearchRequestAllOf(dataSource TelemetryDruidDataSource, intervals []string, granularity TelemetryDruidGranularity) *TelemetryDruidSearchRequestAllOf {
	this := TelemetryDruidSearchRequestAllOf{}
	this.DataSource = dataSource
	this.Intervals = intervals
	this.Granularity = granularity
	return &this
}

// NewTelemetryDruidSearchRequestAllOfWithDefaults instantiates a new TelemetryDruidSearchRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidSearchRequestAllOfWithDefaults() *TelemetryDruidSearchRequestAllOf {
	this := TelemetryDruidSearchRequestAllOf{}
	return &this
}

// GetDataSource returns the DataSource field value
func (o *TelemetryDruidSearchRequestAllOf) GetDataSource() TelemetryDruidDataSource {
	if o == nil {
		var ret TelemetryDruidDataSource
		return ret
	}

	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSearchRequestAllOf) GetDataSourceOk() (*TelemetryDruidDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value
func (o *TelemetryDruidSearchRequestAllOf) SetDataSource(v TelemetryDruidDataSource) {
	o.DataSource = v
}

// GetIntervals returns the Intervals field value
func (o *TelemetryDruidSearchRequestAllOf) GetIntervals() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Intervals
}

// GetIntervalsOk returns a tuple with the Intervals field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSearchRequestAllOf) GetIntervalsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Intervals, true
}

// SetIntervals sets field value
func (o *TelemetryDruidSearchRequestAllOf) SetIntervals(v []string) {
	o.Intervals = v
}

// GetGranularity returns the Granularity field value
func (o *TelemetryDruidSearchRequestAllOf) GetGranularity() TelemetryDruidGranularity {
	if o == nil {
		var ret TelemetryDruidGranularity
		return ret
	}

	return o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSearchRequestAllOf) GetGranularityOk() (*TelemetryDruidGranularity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Granularity, true
}

// SetGranularity sets field value
func (o *TelemetryDruidSearchRequestAllOf) SetGranularity(v TelemetryDruidGranularity) {
	o.Granularity = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *TelemetryDruidSearchRequestAllOf) GetFilter() TelemetryDruidFilter {
	if o == nil || o.Filter == nil {
		var ret TelemetryDruidFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSearchRequestAllOf) GetFilterOk() (*TelemetryDruidFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *TelemetryDruidSearchRequestAllOf) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given TelemetryDruidFilter and assigns it to the Filter field.
func (o *TelemetryDruidSearchRequestAllOf) SetFilter(v TelemetryDruidFilter) {
	o.Filter = &v
}

// GetAggregations returns the Aggregations field value if set, zero value otherwise.
func (o *TelemetryDruidSearchRequestAllOf) GetAggregations() []TelemetryDruidAggregator {
	if o == nil || o.Aggregations == nil {
		var ret []TelemetryDruidAggregator
		return ret
	}
	return *o.Aggregations
}

// GetAggregationsOk returns a tuple with the Aggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSearchRequestAllOf) GetAggregationsOk() (*[]TelemetryDruidAggregator, bool) {
	if o == nil || o.Aggregations == nil {
		return nil, false
	}
	return o.Aggregations, true
}

// HasAggregations returns a boolean if a field has been set.
func (o *TelemetryDruidSearchRequestAllOf) HasAggregations() bool {
	if o != nil && o.Aggregations != nil {
		return true
	}

	return false
}

// SetAggregations gets a reference to the given []TelemetryDruidAggregator and assigns it to the Aggregations field.
func (o *TelemetryDruidSearchRequestAllOf) SetAggregations(v []TelemetryDruidAggregator) {
	o.Aggregations = &v
}

// GetSearchDimensions returns the SearchDimensions field value if set, zero value otherwise.
func (o *TelemetryDruidSearchRequestAllOf) GetSearchDimensions() []string {
	if o == nil || o.SearchDimensions == nil {
		var ret []string
		return ret
	}
	return *o.SearchDimensions
}

// GetSearchDimensionsOk returns a tuple with the SearchDimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSearchRequestAllOf) GetSearchDimensionsOk() (*[]string, bool) {
	if o == nil || o.SearchDimensions == nil {
		return nil, false
	}
	return o.SearchDimensions, true
}

// HasSearchDimensions returns a boolean if a field has been set.
func (o *TelemetryDruidSearchRequestAllOf) HasSearchDimensions() bool {
	if o != nil && o.SearchDimensions != nil {
		return true
	}

	return false
}

// SetSearchDimensions gets a reference to the given []string and assigns it to the SearchDimensions field.
func (o *TelemetryDruidSearchRequestAllOf) SetSearchDimensions(v []string) {
	o.SearchDimensions = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *TelemetryDruidSearchRequestAllOf) GetQuery() TelemetryDruidAggregateSearchSpec {
	if o == nil || o.Query == nil {
		var ret TelemetryDruidAggregateSearchSpec
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSearchRequestAllOf) GetQueryOk() (*TelemetryDruidAggregateSearchSpec, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *TelemetryDruidSearchRequestAllOf) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given TelemetryDruidAggregateSearchSpec and assigns it to the Query field.
func (o *TelemetryDruidSearchRequestAllOf) SetQuery(v TelemetryDruidAggregateSearchSpec) {
	o.Query = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *TelemetryDruidSearchRequestAllOf) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSearchRequestAllOf) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *TelemetryDruidSearchRequestAllOf) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *TelemetryDruidSearchRequestAllOf) SetLimit(v int32) {
	o.Limit = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TelemetryDruidSearchRequestAllOf) GetContext() TelemetryDruidQueryContext {
	if o == nil || o.Context == nil {
		var ret TelemetryDruidQueryContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSearchRequestAllOf) GetContextOk() (*TelemetryDruidQueryContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TelemetryDruidSearchRequestAllOf) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given TelemetryDruidQueryContext and assigns it to the Context field.
func (o *TelemetryDruidSearchRequestAllOf) SetContext(v TelemetryDruidQueryContext) {
	o.Context = &v
}

func (o TelemetryDruidSearchRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dataSource"] = o.DataSource
	}
	if true {
		toSerialize["intervals"] = o.Intervals
	}
	if true {
		toSerialize["granularity"] = o.Granularity
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Aggregations != nil {
		toSerialize["aggregations"] = o.Aggregations
	}
	if o.SearchDimensions != nil {
		toSerialize["searchDimensions"] = o.SearchDimensions
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidSearchRequestAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidSearchRequestAllOf := _TelemetryDruidSearchRequestAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidSearchRequestAllOf); err == nil {
		*o = TelemetryDruidSearchRequestAllOf(varTelemetryDruidSearchRequestAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dataSource")
		delete(additionalProperties, "intervals")
		delete(additionalProperties, "granularity")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "aggregations")
		delete(additionalProperties, "searchDimensions")
		delete(additionalProperties, "query")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "context")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidSearchRequestAllOf struct {
	value *TelemetryDruidSearchRequestAllOf
	isSet bool
}

func (v NullableTelemetryDruidSearchRequestAllOf) Get() *TelemetryDruidSearchRequestAllOf {
	return v.value
}

func (v *NullableTelemetryDruidSearchRequestAllOf) Set(val *TelemetryDruidSearchRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidSearchRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidSearchRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidSearchRequestAllOf(val *TelemetryDruidSearchRequestAllOf) *NullableTelemetryDruidSearchRequestAllOf {
	return &NullableTelemetryDruidSearchRequestAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidSearchRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidSearchRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
