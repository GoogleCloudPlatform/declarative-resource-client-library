// Copyright 2021 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package billingbudgets

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Budget struct {
	Name           *string                `json:"name"`
	DisplayName    *string                `json:"displayName"`
	BudgetFilter   *BudgetBudgetFilter    `json:"budgetFilter"`
	Amount         *BudgetAmount          `json:"amount"`
	ThresholdRules []BudgetThresholdRules `json:"thresholdRules"`
	Etag           *string                `json:"etag"`
	AllUpdatesRule *BudgetAllUpdatesRule  `json:"allUpdatesRule"`
	BillingAccount *string                `json:"billingAccount"`
}

func (r *Budget) String() string {
	return dcl.SprintResource(r)
}

// The enum BudgetBudgetFilterCreditTypesTreatmentEnum.
type BudgetBudgetFilterCreditTypesTreatmentEnum string

// BudgetBudgetFilterCreditTypesTreatmentEnumRef returns a *BudgetBudgetFilterCreditTypesTreatmentEnum with the value of string s
// If the empty string is provided, nil is returned.
func BudgetBudgetFilterCreditTypesTreatmentEnumRef(s string) *BudgetBudgetFilterCreditTypesTreatmentEnum {
	if s == "" {
		return nil
	}

	v := BudgetBudgetFilterCreditTypesTreatmentEnum(s)
	return &v
}

func (v BudgetBudgetFilterCreditTypesTreatmentEnum) Validate() error {
	for _, s := range []string{"INCLUDE_ALL_CREDITS", "EXCLUDE_ALL_CREDITS", "INCLUDE_SPECIFIED_CREDITS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "BudgetBudgetFilterCreditTypesTreatmentEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum BudgetBudgetFilterCalendarPeriodEnum.
type BudgetBudgetFilterCalendarPeriodEnum string

// BudgetBudgetFilterCalendarPeriodEnumRef returns a *BudgetBudgetFilterCalendarPeriodEnum with the value of string s
// If the empty string is provided, nil is returned.
func BudgetBudgetFilterCalendarPeriodEnumRef(s string) *BudgetBudgetFilterCalendarPeriodEnum {
	if s == "" {
		return nil
	}

	v := BudgetBudgetFilterCalendarPeriodEnum(s)
	return &v
}

func (v BudgetBudgetFilterCalendarPeriodEnum) Validate() error {
	for _, s := range []string{"CALENDAR_PERIOD_UNSPECIFIED", "MONTH", "QUARTER", "YEAR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "BudgetBudgetFilterCalendarPeriodEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum BudgetThresholdRulesSpendBasisEnum.
type BudgetThresholdRulesSpendBasisEnum string

// BudgetThresholdRulesSpendBasisEnumRef returns a *BudgetThresholdRulesSpendBasisEnum with the value of string s
// If the empty string is provided, nil is returned.
func BudgetThresholdRulesSpendBasisEnumRef(s string) *BudgetThresholdRulesSpendBasisEnum {
	if s == "" {
		return nil
	}

	v := BudgetThresholdRulesSpendBasisEnum(s)
	return &v
}

func (v BudgetThresholdRulesSpendBasisEnum) Validate() error {
	for _, s := range []string{"BASIS_UNSPECIFIED", "CURRENT_SPEND", "FORECASTED_SPEND"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "BudgetThresholdRulesSpendBasisEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type BudgetBudgetFilter struct {
	empty                bool                                        `json:"-"`
	Projects             []string                                    `json:"projects"`
	CreditTypes          []string                                    `json:"creditTypes"`
	CreditTypesTreatment *BudgetBudgetFilterCreditTypesTreatmentEnum `json:"creditTypesTreatment"`
	Services             []string                                    `json:"services"`
	Subaccounts          []string                                    `json:"subaccounts"`
	Labels               map[string]BudgetBudgetFilterLabels         `json:"labels"`
	CalendarPeriod       *BudgetBudgetFilterCalendarPeriodEnum       `json:"calendarPeriod"`
	CustomPeriod         *BudgetBudgetFilterCustomPeriod             `json:"customPeriod"`
}

type jsonBudgetBudgetFilter BudgetBudgetFilter

func (r *BudgetBudgetFilter) UnmarshalJSON(data []byte) error {
	var res jsonBudgetBudgetFilter
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBudgetBudgetFilter
	} else {

		r.Projects = res.Projects

		r.CreditTypes = res.CreditTypes

		r.CreditTypesTreatment = res.CreditTypesTreatment

		r.Services = res.Services

		r.Subaccounts = res.Subaccounts

		r.Labels = res.Labels

		r.CalendarPeriod = res.CalendarPeriod

		r.CustomPeriod = res.CustomPeriod

	}
	return nil
}

// This object is used to assert a desired state where this BudgetBudgetFilter is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBudgetBudgetFilter *BudgetBudgetFilter = &BudgetBudgetFilter{empty: true}

func (r *BudgetBudgetFilter) Empty() bool {
	return r.empty
}

func (r *BudgetBudgetFilter) String() string {
	return dcl.SprintResource(r)
}

func (r *BudgetBudgetFilter) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BudgetBudgetFilterLabels struct {
	empty  bool     `json:"-"`
	Values []string `json:"values"`
}

type jsonBudgetBudgetFilterLabels BudgetBudgetFilterLabels

func (r *BudgetBudgetFilterLabels) UnmarshalJSON(data []byte) error {
	var res jsonBudgetBudgetFilterLabels
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBudgetBudgetFilterLabels
	} else {

		r.Values = res.Values

	}
	return nil
}

// This object is used to assert a desired state where this BudgetBudgetFilterLabels is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBudgetBudgetFilterLabels *BudgetBudgetFilterLabels = &BudgetBudgetFilterLabels{empty: true}

func (r *BudgetBudgetFilterLabels) Empty() bool {
	return r.empty
}

func (r *BudgetBudgetFilterLabels) String() string {
	return dcl.SprintResource(r)
}

func (r *BudgetBudgetFilterLabels) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BudgetBudgetFilterCustomPeriod struct {
	empty     bool                                     `json:"-"`
	StartDate *BudgetBudgetFilterCustomPeriodStartDate `json:"startDate"`
	EndDate   *BudgetBudgetFilterCustomPeriodEndDate   `json:"endDate"`
}

type jsonBudgetBudgetFilterCustomPeriod BudgetBudgetFilterCustomPeriod

func (r *BudgetBudgetFilterCustomPeriod) UnmarshalJSON(data []byte) error {
	var res jsonBudgetBudgetFilterCustomPeriod
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBudgetBudgetFilterCustomPeriod
	} else {

		r.StartDate = res.StartDate

		r.EndDate = res.EndDate

	}
	return nil
}

// This object is used to assert a desired state where this BudgetBudgetFilterCustomPeriod is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBudgetBudgetFilterCustomPeriod *BudgetBudgetFilterCustomPeriod = &BudgetBudgetFilterCustomPeriod{empty: true}

func (r *BudgetBudgetFilterCustomPeriod) Empty() bool {
	return r.empty
}

func (r *BudgetBudgetFilterCustomPeriod) String() string {
	return dcl.SprintResource(r)
}

func (r *BudgetBudgetFilterCustomPeriod) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BudgetBudgetFilterCustomPeriodStartDate struct {
	empty bool   `json:"-"`
	Year  *int64 `json:"year"`
	Month *int64 `json:"month"`
	Day   *int64 `json:"day"`
}

type jsonBudgetBudgetFilterCustomPeriodStartDate BudgetBudgetFilterCustomPeriodStartDate

func (r *BudgetBudgetFilterCustomPeriodStartDate) UnmarshalJSON(data []byte) error {
	var res jsonBudgetBudgetFilterCustomPeriodStartDate
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBudgetBudgetFilterCustomPeriodStartDate
	} else {

		r.Year = res.Year

		r.Month = res.Month

		r.Day = res.Day

	}
	return nil
}

// This object is used to assert a desired state where this BudgetBudgetFilterCustomPeriodStartDate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBudgetBudgetFilterCustomPeriodStartDate *BudgetBudgetFilterCustomPeriodStartDate = &BudgetBudgetFilterCustomPeriodStartDate{empty: true}

func (r *BudgetBudgetFilterCustomPeriodStartDate) Empty() bool {
	return r.empty
}

func (r *BudgetBudgetFilterCustomPeriodStartDate) String() string {
	return dcl.SprintResource(r)
}

func (r *BudgetBudgetFilterCustomPeriodStartDate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BudgetBudgetFilterCustomPeriodEndDate struct {
	empty bool   `json:"-"`
	Year  *int64 `json:"year"`
	Month *int64 `json:"month"`
	Day   *int64 `json:"day"`
}

type jsonBudgetBudgetFilterCustomPeriodEndDate BudgetBudgetFilterCustomPeriodEndDate

func (r *BudgetBudgetFilterCustomPeriodEndDate) UnmarshalJSON(data []byte) error {
	var res jsonBudgetBudgetFilterCustomPeriodEndDate
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBudgetBudgetFilterCustomPeriodEndDate
	} else {

		r.Year = res.Year

		r.Month = res.Month

		r.Day = res.Day

	}
	return nil
}

// This object is used to assert a desired state where this BudgetBudgetFilterCustomPeriodEndDate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBudgetBudgetFilterCustomPeriodEndDate *BudgetBudgetFilterCustomPeriodEndDate = &BudgetBudgetFilterCustomPeriodEndDate{empty: true}

func (r *BudgetBudgetFilterCustomPeriodEndDate) Empty() bool {
	return r.empty
}

func (r *BudgetBudgetFilterCustomPeriodEndDate) String() string {
	return dcl.SprintResource(r)
}

func (r *BudgetBudgetFilterCustomPeriodEndDate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BudgetAmount struct {
	empty            bool                          `json:"-"`
	SpecifiedAmount  *BudgetAmountSpecifiedAmount  `json:"specifiedAmount"`
	LastPeriodAmount *BudgetAmountLastPeriodAmount `json:"lastPeriodAmount"`
}

type jsonBudgetAmount BudgetAmount

func (r *BudgetAmount) UnmarshalJSON(data []byte) error {
	var res jsonBudgetAmount
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBudgetAmount
	} else {

		r.SpecifiedAmount = res.SpecifiedAmount

		r.LastPeriodAmount = res.LastPeriodAmount

	}
	return nil
}

// This object is used to assert a desired state where this BudgetAmount is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBudgetAmount *BudgetAmount = &BudgetAmount{empty: true}

func (r *BudgetAmount) Empty() bool {
	return r.empty
}

func (r *BudgetAmount) String() string {
	return dcl.SprintResource(r)
}

func (r *BudgetAmount) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BudgetAmountSpecifiedAmount struct {
	empty        bool    `json:"-"`
	CurrencyCode *string `json:"currencyCode"`
	Units        *int64  `json:"units"`
	Nanos        *int64  `json:"nanos"`
}

type jsonBudgetAmountSpecifiedAmount BudgetAmountSpecifiedAmount

func (r *BudgetAmountSpecifiedAmount) UnmarshalJSON(data []byte) error {
	var res jsonBudgetAmountSpecifiedAmount
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBudgetAmountSpecifiedAmount
	} else {

		r.CurrencyCode = res.CurrencyCode

		r.Units = res.Units

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this BudgetAmountSpecifiedAmount is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBudgetAmountSpecifiedAmount *BudgetAmountSpecifiedAmount = &BudgetAmountSpecifiedAmount{empty: true}

func (r *BudgetAmountSpecifiedAmount) Empty() bool {
	return r.empty
}

func (r *BudgetAmountSpecifiedAmount) String() string {
	return dcl.SprintResource(r)
}

func (r *BudgetAmountSpecifiedAmount) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BudgetAmountLastPeriodAmount struct {
	empty bool `json:"-"`
}

type jsonBudgetAmountLastPeriodAmount BudgetAmountLastPeriodAmount

func (r *BudgetAmountLastPeriodAmount) UnmarshalJSON(data []byte) error {
	var res jsonBudgetAmountLastPeriodAmount
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBudgetAmountLastPeriodAmount
	} else {

	}
	return nil
}

// This object is used to assert a desired state where this BudgetAmountLastPeriodAmount is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBudgetAmountLastPeriodAmount *BudgetAmountLastPeriodAmount = &BudgetAmountLastPeriodAmount{empty: true}

func (r *BudgetAmountLastPeriodAmount) Empty() bool {
	return r.empty
}

func (r *BudgetAmountLastPeriodAmount) String() string {
	return dcl.SprintResource(r)
}

func (r *BudgetAmountLastPeriodAmount) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BudgetThresholdRules struct {
	empty            bool                                `json:"-"`
	ThresholdPercent *float64                            `json:"thresholdPercent"`
	SpendBasis       *BudgetThresholdRulesSpendBasisEnum `json:"spendBasis"`
}

type jsonBudgetThresholdRules BudgetThresholdRules

func (r *BudgetThresholdRules) UnmarshalJSON(data []byte) error {
	var res jsonBudgetThresholdRules
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBudgetThresholdRules
	} else {

		r.ThresholdPercent = res.ThresholdPercent

		r.SpendBasis = res.SpendBasis

	}
	return nil
}

// This object is used to assert a desired state where this BudgetThresholdRules is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBudgetThresholdRules *BudgetThresholdRules = &BudgetThresholdRules{empty: true}

func (r *BudgetThresholdRules) Empty() bool {
	return r.empty
}

func (r *BudgetThresholdRules) String() string {
	return dcl.SprintResource(r)
}

func (r *BudgetThresholdRules) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BudgetAllUpdatesRule struct {
	empty                          bool     `json:"-"`
	PubsubTopic                    *string  `json:"pubsubTopic"`
	SchemaVersion                  *string  `json:"schemaVersion"`
	MonitoringNotificationChannels []string `json:"monitoringNotificationChannels"`
	DisableDefaultIamRecipients    *bool    `json:"disableDefaultIamRecipients"`
}

type jsonBudgetAllUpdatesRule BudgetAllUpdatesRule

func (r *BudgetAllUpdatesRule) UnmarshalJSON(data []byte) error {
	var res jsonBudgetAllUpdatesRule
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBudgetAllUpdatesRule
	} else {

		r.PubsubTopic = res.PubsubTopic

		r.SchemaVersion = res.SchemaVersion

		r.MonitoringNotificationChannels = res.MonitoringNotificationChannels

		r.DisableDefaultIamRecipients = res.DisableDefaultIamRecipients

	}
	return nil
}

// This object is used to assert a desired state where this BudgetAllUpdatesRule is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyBudgetAllUpdatesRule *BudgetAllUpdatesRule = &BudgetAllUpdatesRule{empty: true}

func (r *BudgetAllUpdatesRule) Empty() bool {
	return r.empty
}

func (r *BudgetAllUpdatesRule) String() string {
	return dcl.SprintResource(r)
}

func (r *BudgetAllUpdatesRule) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Budget) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "billingbudgets",
		Type:    "Budget",
		Version: "billingbudgets",
	}
}

const BudgetMaxPage = -1

type BudgetList struct {
	Items []*Budget

	nextToken string

	pageSize int32

	billingAccount string
}

func (l *BudgetList) HasNext() bool {
	return l.nextToken != ""
}

func (l *BudgetList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listBudget(ctx, l.billingAccount, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListBudget(ctx context.Context, billingAccount string) (*BudgetList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListBudgetWithMaxResults(ctx, billingAccount, BudgetMaxPage)

}

func (c *Client) ListBudgetWithMaxResults(ctx context.Context, billingAccount string, pageSize int32) (*BudgetList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listBudget(ctx, billingAccount, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &BudgetList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		billingAccount: billingAccount,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Budget) URLNormalized() *Budget {
	normalized := dcl.Copy(*r).(Budget)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.BillingAccount = dcl.SelfLinkToName(r.BillingAccount)
	return &normalized
}

func (c *Client) GetBudget(ctx context.Context, r *Budget) (*Budget, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getBudgetRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalBudget(b, c)
	if err != nil {
		return nil, err
	}
	result.BillingAccount = r.BillingAccount
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeBudgetNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteBudget(ctx context.Context, r *Budget) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Budget resource is nil")
	}
	c.Config.Logger.Info("Deleting Budget...")
	deleteOp := deleteBudgetOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllBudget deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllBudget(ctx context.Context, billingAccount string, filter func(*Budget) bool) error {
	listObj, err := c.ListBudget(ctx, billingAccount)
	if err != nil {
		return err
	}

	err = c.deleteAllBudget(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllBudget(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyBudget(ctx context.Context, rawDesired *Budget, opts ...dcl.ApplyOption) (*Budget, error) {
	var resultNewState *Budget
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyBudgetHelper(c, ctx, rawDesired, opts...)
		resultNewState = newState
		if err != nil {
			// If the error is 409, there is conflict in resource update.
			// Here we want to apply changes based on latest state.
			if dcl.IsConflictError(err) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	return resultNewState, err
}

func applyBudgetHelper(c *Client, ctx context.Context, rawDesired *Budget, opts ...dcl.ApplyOption) (*Budget, error) {
	c.Config.Logger.Info("Beginning ApplyBudget...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.budgetDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToBudgetDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	var recreate bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				if dcl.HasLifecycleParam(lp, dcl.BlockDestruction) || dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
					return nil, dcl.ApplyInfeasibleError{
						Message: fmt.Sprintf("Infeasible update: (%v) would require recreation.", d),
					}
				}
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []budgetApiOperation
	if create {
		ops = append(ops, &createBudgetOperation{})
	} else if recreate {
		ops = append(ops, &deleteBudgetOperation{})
		ops = append(ops, &createBudgetOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeBudgetDesiredState(rawDesired, nil)
		if err != nil {
			return nil, err
		}
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.Infof("Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.Infof("Performing operation %#v", op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.Infof("Failed operation %#v: %v", op, err)
			return nil, err
		}
		c.Config.Logger.Infof("Finished operation %#v", op)
	}

	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.Info("Retrieving raw new state...")
	rawNew, err := c.GetBudget(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createBudgetOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapBudget(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeBudgetNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeBudgetNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeBudgetDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffBudget(c, newDesired, newState)
	if err != nil {
		return nil, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.Info("No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.Infof("Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.Info("Done Apply.")
	return newState, nil
}
