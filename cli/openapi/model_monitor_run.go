/*
TraceTest

OpenAPI definition for TraceTest endpoint and resources

API version: 0.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the MonitorRun type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitorRun{}

// MonitorRun struct for MonitorRun
type MonitorRun struct {
	Id                 *int32             `json:"id,omitempty"`
	MonitorId          *string            `json:"monitorId,omitempty"`
	MonitorVersion     *int32             `json:"monitorVersion,omitempty"`
	RunGroupId         *string            `json:"runGroupId,omitempty"`
	CreatedAt          *time.Time         `json:"createdAt,omitempty"`
	CompletedAt        *time.Time         `json:"completedAt,omitempty"`
	ExecutionType      *string            `json:"executionType,omitempty"`
	State              *string            `json:"state,omitempty"`
	VariableSet        *VariableSet       `json:"variableSet,omitempty"`
	Metadata           *map[string]string `json:"metadata,omitempty"`
	TestRunsCount      *int32             `json:"testRunsCount,omitempty"`
	TestSuiteRunsCount *int32             `json:"testSuiteRunsCount,omitempty"`
	// list of test runs of the Monitor Run
	TestRuns []TestRun `json:"testRuns,omitempty"`
	// list of test suite runs of the Monitor Run
	TestSuiteRuns []TestSuiteRun `json:"testSuiteRuns,omitempty"`
	Alerts        []AlertResult  `json:"alerts,omitempty"`
}

// NewMonitorRun instantiates a new MonitorRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorRun() *MonitorRun {
	this := MonitorRun{}
	var executionType string = "SCHEDULED"
	this.ExecutionType = &executionType
	return &this
}

// NewMonitorRunWithDefaults instantiates a new MonitorRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorRunWithDefaults() *MonitorRun {
	this := MonitorRun{}
	var executionType string = "SCHEDULED"
	this.ExecutionType = &executionType
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MonitorRun) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorRun) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MonitorRun) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MonitorRun) SetId(v int32) {
	o.Id = &v
}

// GetMonitorId returns the MonitorId field value if set, zero value otherwise.
func (o *MonitorRun) GetMonitorId() string {
	if o == nil || isNil(o.MonitorId) {
		var ret string
		return ret
	}
	return *o.MonitorId
}

// GetMonitorIdOk returns a tuple with the MonitorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorRun) GetMonitorIdOk() (*string, bool) {
	if o == nil || isNil(o.MonitorId) {
		return nil, false
	}
	return o.MonitorId, true
}

// HasMonitorId returns a boolean if a field has been set.
func (o *MonitorRun) HasMonitorId() bool {
	if o != nil && !isNil(o.MonitorId) {
		return true
	}

	return false
}

// SetMonitorId gets a reference to the given string and assigns it to the MonitorId field.
func (o *MonitorRun) SetMonitorId(v string) {
	o.MonitorId = &v
}

// GetMonitorVersion returns the MonitorVersion field value if set, zero value otherwise.
func (o *MonitorRun) GetMonitorVersion() int32 {
	if o == nil || isNil(o.MonitorVersion) {
		var ret int32
		return ret
	}
	return *o.MonitorVersion
}

// GetMonitorVersionOk returns a tuple with the MonitorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorRun) GetMonitorVersionOk() (*int32, bool) {
	if o == nil || isNil(o.MonitorVersion) {
		return nil, false
	}
	return o.MonitorVersion, true
}

// HasMonitorVersion returns a boolean if a field has been set.
func (o *MonitorRun) HasMonitorVersion() bool {
	if o != nil && !isNil(o.MonitorVersion) {
		return true
	}

	return false
}

// SetMonitorVersion gets a reference to the given int32 and assigns it to the MonitorVersion field.
func (o *MonitorRun) SetMonitorVersion(v int32) {
	o.MonitorVersion = &v
}

// GetRunGroupId returns the RunGroupId field value if set, zero value otherwise.
func (o *MonitorRun) GetRunGroupId() string {
	if o == nil || isNil(o.RunGroupId) {
		var ret string
		return ret
	}
	return *o.RunGroupId
}

// GetRunGroupIdOk returns a tuple with the RunGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorRun) GetRunGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.RunGroupId) {
		return nil, false
	}
	return o.RunGroupId, true
}

// HasRunGroupId returns a boolean if a field has been set.
func (o *MonitorRun) HasRunGroupId() bool {
	if o != nil && !isNil(o.RunGroupId) {
		return true
	}

	return false
}

// SetRunGroupId gets a reference to the given string and assigns it to the RunGroupId field.
func (o *MonitorRun) SetRunGroupId(v string) {
	o.RunGroupId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MonitorRun) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorRun) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MonitorRun) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *MonitorRun) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *MonitorRun) GetCompletedAt() time.Time {
	if o == nil || isNil(o.CompletedAt) {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorRun) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CompletedAt) {
		return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *MonitorRun) HasCompletedAt() bool {
	if o != nil && !isNil(o.CompletedAt) {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given time.Time and assigns it to the CompletedAt field.
func (o *MonitorRun) SetCompletedAt(v time.Time) {
	o.CompletedAt = &v
}

// GetExecutionType returns the ExecutionType field value if set, zero value otherwise.
func (o *MonitorRun) GetExecutionType() string {
	if o == nil || isNil(o.ExecutionType) {
		var ret string
		return ret
	}
	return *o.ExecutionType
}

// GetExecutionTypeOk returns a tuple with the ExecutionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorRun) GetExecutionTypeOk() (*string, bool) {
	if o == nil || isNil(o.ExecutionType) {
		return nil, false
	}
	return o.ExecutionType, true
}

// HasExecutionType returns a boolean if a field has been set.
func (o *MonitorRun) HasExecutionType() bool {
	if o != nil && !isNil(o.ExecutionType) {
		return true
	}

	return false
}

// SetExecutionType gets a reference to the given string and assigns it to the ExecutionType field.
func (o *MonitorRun) SetExecutionType(v string) {
	o.ExecutionType = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *MonitorRun) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorRun) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MonitorRun) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *MonitorRun) SetState(v string) {
	o.State = &v
}

// GetVariableSet returns the VariableSet field value if set, zero value otherwise.
func (o *MonitorRun) GetVariableSet() VariableSet {
	if o == nil || isNil(o.VariableSet) {
		var ret VariableSet
		return ret
	}
	return *o.VariableSet
}

// GetVariableSetOk returns a tuple with the VariableSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorRun) GetVariableSetOk() (*VariableSet, bool) {
	if o == nil || isNil(o.VariableSet) {
		return nil, false
	}
	return o.VariableSet, true
}

// HasVariableSet returns a boolean if a field has been set.
func (o *MonitorRun) HasVariableSet() bool {
	if o != nil && !isNil(o.VariableSet) {
		return true
	}

	return false
}

// SetVariableSet gets a reference to the given VariableSet and assigns it to the VariableSet field.
func (o *MonitorRun) SetVariableSet(v VariableSet) {
	o.VariableSet = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *MonitorRun) GetMetadata() map[string]string {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorRun) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *MonitorRun) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *MonitorRun) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetTestRunsCount returns the TestRunsCount field value if set, zero value otherwise.
func (o *MonitorRun) GetTestRunsCount() int32 {
	if o == nil || isNil(o.TestRunsCount) {
		var ret int32
		return ret
	}
	return *o.TestRunsCount
}

// GetTestRunsCountOk returns a tuple with the TestRunsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorRun) GetTestRunsCountOk() (*int32, bool) {
	if o == nil || isNil(o.TestRunsCount) {
		return nil, false
	}
	return o.TestRunsCount, true
}

// HasTestRunsCount returns a boolean if a field has been set.
func (o *MonitorRun) HasTestRunsCount() bool {
	if o != nil && !isNil(o.TestRunsCount) {
		return true
	}

	return false
}

// SetTestRunsCount gets a reference to the given int32 and assigns it to the TestRunsCount field.
func (o *MonitorRun) SetTestRunsCount(v int32) {
	o.TestRunsCount = &v
}

// GetTestSuiteRunsCount returns the TestSuiteRunsCount field value if set, zero value otherwise.
func (o *MonitorRun) GetTestSuiteRunsCount() int32 {
	if o == nil || isNil(o.TestSuiteRunsCount) {
		var ret int32
		return ret
	}
	return *o.TestSuiteRunsCount
}

// GetTestSuiteRunsCountOk returns a tuple with the TestSuiteRunsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorRun) GetTestSuiteRunsCountOk() (*int32, bool) {
	if o == nil || isNil(o.TestSuiteRunsCount) {
		return nil, false
	}
	return o.TestSuiteRunsCount, true
}

// HasTestSuiteRunsCount returns a boolean if a field has been set.
func (o *MonitorRun) HasTestSuiteRunsCount() bool {
	if o != nil && !isNil(o.TestSuiteRunsCount) {
		return true
	}

	return false
}

// SetTestSuiteRunsCount gets a reference to the given int32 and assigns it to the TestSuiteRunsCount field.
func (o *MonitorRun) SetTestSuiteRunsCount(v int32) {
	o.TestSuiteRunsCount = &v
}

// GetTestRuns returns the TestRuns field value if set, zero value otherwise.
func (o *MonitorRun) GetTestRuns() []TestRun {
	if o == nil || isNil(o.TestRuns) {
		var ret []TestRun
		return ret
	}
	return o.TestRuns
}

// GetTestRunsOk returns a tuple with the TestRuns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorRun) GetTestRunsOk() ([]TestRun, bool) {
	if o == nil || isNil(o.TestRuns) {
		return nil, false
	}
	return o.TestRuns, true
}

// HasTestRuns returns a boolean if a field has been set.
func (o *MonitorRun) HasTestRuns() bool {
	if o != nil && !isNil(o.TestRuns) {
		return true
	}

	return false
}

// SetTestRuns gets a reference to the given []TestRun and assigns it to the TestRuns field.
func (o *MonitorRun) SetTestRuns(v []TestRun) {
	o.TestRuns = v
}

// GetTestSuiteRuns returns the TestSuiteRuns field value if set, zero value otherwise.
func (o *MonitorRun) GetTestSuiteRuns() []TestSuiteRun {
	if o == nil || isNil(o.TestSuiteRuns) {
		var ret []TestSuiteRun
		return ret
	}
	return o.TestSuiteRuns
}

// GetTestSuiteRunsOk returns a tuple with the TestSuiteRuns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorRun) GetTestSuiteRunsOk() ([]TestSuiteRun, bool) {
	if o == nil || isNil(o.TestSuiteRuns) {
		return nil, false
	}
	return o.TestSuiteRuns, true
}

// HasTestSuiteRuns returns a boolean if a field has been set.
func (o *MonitorRun) HasTestSuiteRuns() bool {
	if o != nil && !isNil(o.TestSuiteRuns) {
		return true
	}

	return false
}

// SetTestSuiteRuns gets a reference to the given []TestSuiteRun and assigns it to the TestSuiteRuns field.
func (o *MonitorRun) SetTestSuiteRuns(v []TestSuiteRun) {
	o.TestSuiteRuns = v
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *MonitorRun) GetAlerts() []AlertResult {
	if o == nil || isNil(o.Alerts) {
		var ret []AlertResult
		return ret
	}
	return o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorRun) GetAlertsOk() ([]AlertResult, bool) {
	if o == nil || isNil(o.Alerts) {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *MonitorRun) HasAlerts() bool {
	if o != nil && !isNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given []AlertResult and assigns it to the Alerts field.
func (o *MonitorRun) SetAlerts(v []AlertResult) {
	o.Alerts = v
}

func (o MonitorRun) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitorRun) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	if !isNil(o.MonitorId) {
		toSerialize["monitorId"] = o.MonitorId
	}
	if !isNil(o.MonitorVersion) {
		toSerialize["monitorVersion"] = o.MonitorVersion
	}
	if !isNil(o.RunGroupId) {
		toSerialize["runGroupId"] = o.RunGroupId
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.CompletedAt) {
		toSerialize["completedAt"] = o.CompletedAt
	}
	if !isNil(o.ExecutionType) {
		toSerialize["executionType"] = o.ExecutionType
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.VariableSet) {
		toSerialize["variableSet"] = o.VariableSet
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !isNil(o.TestRunsCount) {
		toSerialize["testRunsCount"] = o.TestRunsCount
	}
	if !isNil(o.TestSuiteRunsCount) {
		toSerialize["testSuiteRunsCount"] = o.TestSuiteRunsCount
	}
	if !isNil(o.TestRuns) {
		toSerialize["testRuns"] = o.TestRuns
	}
	if !isNil(o.TestSuiteRuns) {
		toSerialize["testSuiteRuns"] = o.TestSuiteRuns
	}
	if !isNil(o.Alerts) {
		toSerialize["alerts"] = o.Alerts
	}
	return toSerialize, nil
}

type NullableMonitorRun struct {
	value *MonitorRun
	isSet bool
}

func (v NullableMonitorRun) Get() *MonitorRun {
	return v.value
}

func (v *NullableMonitorRun) Set(val *MonitorRun) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorRun) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorRun(val *MonitorRun) *NullableMonitorRun {
	return &NullableMonitorRun{value: val, isSet: true}
}

func (v NullableMonitorRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
