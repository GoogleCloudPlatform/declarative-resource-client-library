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
package server

import (
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	billingbudgetspb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/billingbudgets/billingbudgets_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/billingbudgets"
)

// Server implements the gRPC interface for Budget.
type BudgetServer struct{}

// ProtoToBudgetBudgetFilterCreditTypesTreatmentEnum converts a BudgetBudgetFilterCreditTypesTreatmentEnum enum from its proto representation.
func ProtoToBillingbudgetsBudgetBudgetFilterCreditTypesTreatmentEnum(e billingbudgetspb.BillingbudgetsBudgetBudgetFilterCreditTypesTreatmentEnum) *billingbudgets.BudgetBudgetFilterCreditTypesTreatmentEnum {
	if e == 0 {
		return nil
	}
	if n, ok := billingbudgetspb.BillingbudgetsBudgetBudgetFilterCreditTypesTreatmentEnum_name[int32(e)]; ok {
		e := billingbudgets.BudgetBudgetFilterCreditTypesTreatmentEnum(n[len("BillingbudgetsBudgetBudgetFilterCreditTypesTreatmentEnum"):])
		return &e
	}
	return nil
}

// ProtoToBudgetBudgetFilterCalendarPeriodEnum converts a BudgetBudgetFilterCalendarPeriodEnum enum from its proto representation.
func ProtoToBillingbudgetsBudgetBudgetFilterCalendarPeriodEnum(e billingbudgetspb.BillingbudgetsBudgetBudgetFilterCalendarPeriodEnum) *billingbudgets.BudgetBudgetFilterCalendarPeriodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := billingbudgetspb.BillingbudgetsBudgetBudgetFilterCalendarPeriodEnum_name[int32(e)]; ok {
		e := billingbudgets.BudgetBudgetFilterCalendarPeriodEnum(n[len("BillingbudgetsBudgetBudgetFilterCalendarPeriodEnum"):])
		return &e
	}
	return nil
}

// ProtoToBudgetThresholdRulesSpendBasisEnum converts a BudgetThresholdRulesSpendBasisEnum enum from its proto representation.
func ProtoToBillingbudgetsBudgetThresholdRulesSpendBasisEnum(e billingbudgetspb.BillingbudgetsBudgetThresholdRulesSpendBasisEnum) *billingbudgets.BudgetThresholdRulesSpendBasisEnum {
	if e == 0 {
		return nil
	}
	if n, ok := billingbudgetspb.BillingbudgetsBudgetThresholdRulesSpendBasisEnum_name[int32(e)]; ok {
		e := billingbudgets.BudgetThresholdRulesSpendBasisEnum(n[len("BillingbudgetsBudgetThresholdRulesSpendBasisEnum"):])
		return &e
	}
	return nil
}

// ProtoToBudgetBudgetFilter converts a BudgetBudgetFilter resource from its proto representation.
func ProtoToBillingbudgetsBudgetBudgetFilter(p *billingbudgetspb.BillingbudgetsBudgetBudgetFilter) *billingbudgets.BudgetBudgetFilter {
	if p == nil {
		return nil
	}
	obj := &billingbudgets.BudgetBudgetFilter{
		CreditTypesTreatment: ProtoToBillingbudgetsBudgetBudgetFilterCreditTypesTreatmentEnum(p.GetCreditTypesTreatment()),
		CalendarPeriod:       ProtoToBillingbudgetsBudgetBudgetFilterCalendarPeriodEnum(p.GetCalendarPeriod()),
		CustomPeriod:         ProtoToBillingbudgetsBudgetBudgetFilterCustomPeriod(p.GetCustomPeriod()),
	}
	for _, r := range p.GetProjects() {
		obj.Projects = append(obj.Projects, r)
	}
	for _, r := range p.GetCreditTypes() {
		obj.CreditTypes = append(obj.CreditTypes, r)
	}
	for _, r := range p.GetServices() {
		obj.Services = append(obj.Services, r)
	}
	for _, r := range p.GetSubaccounts() {
		obj.Subaccounts = append(obj.Subaccounts, r)
	}
	return obj
}

// ProtoToBudgetBudgetFilterLabels converts a BudgetBudgetFilterLabels resource from its proto representation.
func ProtoToBillingbudgetsBudgetBudgetFilterLabels(p *billingbudgetspb.BillingbudgetsBudgetBudgetFilterLabels) *billingbudgets.BudgetBudgetFilterLabels {
	if p == nil {
		return nil
	}
	obj := &billingbudgets.BudgetBudgetFilterLabels{}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToBudgetBudgetFilterCustomPeriod converts a BudgetBudgetFilterCustomPeriod resource from its proto representation.
func ProtoToBillingbudgetsBudgetBudgetFilterCustomPeriod(p *billingbudgetspb.BillingbudgetsBudgetBudgetFilterCustomPeriod) *billingbudgets.BudgetBudgetFilterCustomPeriod {
	if p == nil {
		return nil
	}
	obj := &billingbudgets.BudgetBudgetFilterCustomPeriod{
		StartDate: ProtoToBillingbudgetsBudgetBudgetFilterCustomPeriodStartDate(p.GetStartDate()),
		EndDate:   ProtoToBillingbudgetsBudgetBudgetFilterCustomPeriodEndDate(p.GetEndDate()),
	}
	return obj
}

// ProtoToBudgetBudgetFilterCustomPeriodStartDate converts a BudgetBudgetFilterCustomPeriodStartDate resource from its proto representation.
func ProtoToBillingbudgetsBudgetBudgetFilterCustomPeriodStartDate(p *billingbudgetspb.BillingbudgetsBudgetBudgetFilterCustomPeriodStartDate) *billingbudgets.BudgetBudgetFilterCustomPeriodStartDate {
	if p == nil {
		return nil
	}
	obj := &billingbudgets.BudgetBudgetFilterCustomPeriodStartDate{
		Year:  dcl.Int64OrNil(p.Year),
		Month: dcl.Int64OrNil(p.Month),
		Day:   dcl.Int64OrNil(p.Day),
	}
	return obj
}

// ProtoToBudgetBudgetFilterCustomPeriodEndDate converts a BudgetBudgetFilterCustomPeriodEndDate resource from its proto representation.
func ProtoToBillingbudgetsBudgetBudgetFilterCustomPeriodEndDate(p *billingbudgetspb.BillingbudgetsBudgetBudgetFilterCustomPeriodEndDate) *billingbudgets.BudgetBudgetFilterCustomPeriodEndDate {
	if p == nil {
		return nil
	}
	obj := &billingbudgets.BudgetBudgetFilterCustomPeriodEndDate{
		Year:  dcl.Int64OrNil(p.Year),
		Month: dcl.Int64OrNil(p.Month),
		Day:   dcl.Int64OrNil(p.Day),
	}
	return obj
}

// ProtoToBudgetAmount converts a BudgetAmount resource from its proto representation.
func ProtoToBillingbudgetsBudgetAmount(p *billingbudgetspb.BillingbudgetsBudgetAmount) *billingbudgets.BudgetAmount {
	if p == nil {
		return nil
	}
	obj := &billingbudgets.BudgetAmount{
		SpecifiedAmount:  ProtoToBillingbudgetsBudgetAmountSpecifiedAmount(p.GetSpecifiedAmount()),
		LastPeriodAmount: ProtoToBillingbudgetsBudgetAmountLastPeriodAmount(p.GetLastPeriodAmount()),
	}
	return obj
}

// ProtoToBudgetAmountSpecifiedAmount converts a BudgetAmountSpecifiedAmount resource from its proto representation.
func ProtoToBillingbudgetsBudgetAmountSpecifiedAmount(p *billingbudgetspb.BillingbudgetsBudgetAmountSpecifiedAmount) *billingbudgets.BudgetAmountSpecifiedAmount {
	if p == nil {
		return nil
	}
	obj := &billingbudgets.BudgetAmountSpecifiedAmount{
		CurrencyCode: dcl.StringOrNil(p.CurrencyCode),
		Units:        dcl.Int64OrNil(p.Units),
		Nanos:        dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToBudgetAmountLastPeriodAmount converts a BudgetAmountLastPeriodAmount resource from its proto representation.
func ProtoToBillingbudgetsBudgetAmountLastPeriodAmount(p *billingbudgetspb.BillingbudgetsBudgetAmountLastPeriodAmount) *billingbudgets.BudgetAmountLastPeriodAmount {
	if p == nil {
		return nil
	}
	obj := &billingbudgets.BudgetAmountLastPeriodAmount{}
	return obj
}

// ProtoToBudgetThresholdRules converts a BudgetThresholdRules resource from its proto representation.
func ProtoToBillingbudgetsBudgetThresholdRules(p *billingbudgetspb.BillingbudgetsBudgetThresholdRules) *billingbudgets.BudgetThresholdRules {
	if p == nil {
		return nil
	}
	obj := &billingbudgets.BudgetThresholdRules{
		ThresholdPercent: dcl.Float64OrNil(p.ThresholdPercent),
		SpendBasis:       ProtoToBillingbudgetsBudgetThresholdRulesSpendBasisEnum(p.GetSpendBasis()),
	}
	return obj
}

// ProtoToBudgetAllUpdatesRule converts a BudgetAllUpdatesRule resource from its proto representation.
func ProtoToBillingbudgetsBudgetAllUpdatesRule(p *billingbudgetspb.BillingbudgetsBudgetAllUpdatesRule) *billingbudgets.BudgetAllUpdatesRule {
	if p == nil {
		return nil
	}
	obj := &billingbudgets.BudgetAllUpdatesRule{
		PubsubTopic:                 dcl.StringOrNil(p.PubsubTopic),
		SchemaVersion:               dcl.StringOrNil(p.SchemaVersion),
		DisableDefaultIamRecipients: dcl.Bool(p.DisableDefaultIamRecipients),
	}
	for _, r := range p.GetMonitoringNotificationChannels() {
		obj.MonitoringNotificationChannels = append(obj.MonitoringNotificationChannels, r)
	}
	return obj
}

// ProtoToBudget converts a Budget resource from its proto representation.
func ProtoToBudget(p *billingbudgetspb.BillingbudgetsBudget) *billingbudgets.Budget {
	obj := &billingbudgets.Budget{
		Name:           dcl.StringOrNil(p.Name),
		DisplayName:    dcl.StringOrNil(p.DisplayName),
		BudgetFilter:   ProtoToBillingbudgetsBudgetBudgetFilter(p.GetBudgetFilter()),
		Amount:         ProtoToBillingbudgetsBudgetAmount(p.GetAmount()),
		Etag:           dcl.StringOrNil(p.Etag),
		AllUpdatesRule: ProtoToBillingbudgetsBudgetAllUpdatesRule(p.GetAllUpdatesRule()),
		BillingAccount: dcl.StringOrNil(p.BillingAccount),
	}
	for _, r := range p.GetThresholdRules() {
		obj.ThresholdRules = append(obj.ThresholdRules, *ProtoToBillingbudgetsBudgetThresholdRules(r))
	}
	return obj
}

// BudgetBudgetFilterCreditTypesTreatmentEnumToProto converts a BudgetBudgetFilterCreditTypesTreatmentEnum enum to its proto representation.
func BillingbudgetsBudgetBudgetFilterCreditTypesTreatmentEnumToProto(e *billingbudgets.BudgetBudgetFilterCreditTypesTreatmentEnum) billingbudgetspb.BillingbudgetsBudgetBudgetFilterCreditTypesTreatmentEnum {
	if e == nil {
		return billingbudgetspb.BillingbudgetsBudgetBudgetFilterCreditTypesTreatmentEnum(0)
	}
	if v, ok := billingbudgetspb.BillingbudgetsBudgetBudgetFilterCreditTypesTreatmentEnum_value["BudgetBudgetFilterCreditTypesTreatmentEnum"+string(*e)]; ok {
		return billingbudgetspb.BillingbudgetsBudgetBudgetFilterCreditTypesTreatmentEnum(v)
	}
	return billingbudgetspb.BillingbudgetsBudgetBudgetFilterCreditTypesTreatmentEnum(0)
}

// BudgetBudgetFilterCalendarPeriodEnumToProto converts a BudgetBudgetFilterCalendarPeriodEnum enum to its proto representation.
func BillingbudgetsBudgetBudgetFilterCalendarPeriodEnumToProto(e *billingbudgets.BudgetBudgetFilterCalendarPeriodEnum) billingbudgetspb.BillingbudgetsBudgetBudgetFilterCalendarPeriodEnum {
	if e == nil {
		return billingbudgetspb.BillingbudgetsBudgetBudgetFilterCalendarPeriodEnum(0)
	}
	if v, ok := billingbudgetspb.BillingbudgetsBudgetBudgetFilterCalendarPeriodEnum_value["BudgetBudgetFilterCalendarPeriodEnum"+string(*e)]; ok {
		return billingbudgetspb.BillingbudgetsBudgetBudgetFilterCalendarPeriodEnum(v)
	}
	return billingbudgetspb.BillingbudgetsBudgetBudgetFilterCalendarPeriodEnum(0)
}

// BudgetThresholdRulesSpendBasisEnumToProto converts a BudgetThresholdRulesSpendBasisEnum enum to its proto representation.
func BillingbudgetsBudgetThresholdRulesSpendBasisEnumToProto(e *billingbudgets.BudgetThresholdRulesSpendBasisEnum) billingbudgetspb.BillingbudgetsBudgetThresholdRulesSpendBasisEnum {
	if e == nil {
		return billingbudgetspb.BillingbudgetsBudgetThresholdRulesSpendBasisEnum(0)
	}
	if v, ok := billingbudgetspb.BillingbudgetsBudgetThresholdRulesSpendBasisEnum_value["BudgetThresholdRulesSpendBasisEnum"+string(*e)]; ok {
		return billingbudgetspb.BillingbudgetsBudgetThresholdRulesSpendBasisEnum(v)
	}
	return billingbudgetspb.BillingbudgetsBudgetThresholdRulesSpendBasisEnum(0)
}

// BudgetBudgetFilterToProto converts a BudgetBudgetFilter resource to its proto representation.
func BillingbudgetsBudgetBudgetFilterToProto(o *billingbudgets.BudgetBudgetFilter) *billingbudgetspb.BillingbudgetsBudgetBudgetFilter {
	if o == nil {
		return nil
	}
	p := &billingbudgetspb.BillingbudgetsBudgetBudgetFilter{
		CreditTypesTreatment: BillingbudgetsBudgetBudgetFilterCreditTypesTreatmentEnumToProto(o.CreditTypesTreatment),
		CalendarPeriod:       BillingbudgetsBudgetBudgetFilterCalendarPeriodEnumToProto(o.CalendarPeriod),
		CustomPeriod:         BillingbudgetsBudgetBudgetFilterCustomPeriodToProto(o.CustomPeriod),
	}
	for _, r := range o.Projects {
		p.Projects = append(p.Projects, r)
	}
	for _, r := range o.CreditTypes {
		p.CreditTypes = append(p.CreditTypes, r)
	}
	for _, r := range o.Services {
		p.Services = append(p.Services, r)
	}
	for _, r := range o.Subaccounts {
		p.Subaccounts = append(p.Subaccounts, r)
	}
	p.Labels = make(map[string]*billingbudgetspb.BillingbudgetsBudgetBudgetFilterLabels)
	for k, r := range o.Labels {
		p.Labels[k] = BillingbudgetsBudgetBudgetFilterLabelsToProto(&r)
	}
	return p
}

// BudgetBudgetFilterLabelsToProto converts a BudgetBudgetFilterLabels resource to its proto representation.
func BillingbudgetsBudgetBudgetFilterLabelsToProto(o *billingbudgets.BudgetBudgetFilterLabels) *billingbudgetspb.BillingbudgetsBudgetBudgetFilterLabels {
	if o == nil {
		return nil
	}
	p := &billingbudgetspb.BillingbudgetsBudgetBudgetFilterLabels{}
	for _, r := range o.Values {
		p.Values = append(p.Values, r)
	}
	return p
}

// BudgetBudgetFilterCustomPeriodToProto converts a BudgetBudgetFilterCustomPeriod resource to its proto representation.
func BillingbudgetsBudgetBudgetFilterCustomPeriodToProto(o *billingbudgets.BudgetBudgetFilterCustomPeriod) *billingbudgetspb.BillingbudgetsBudgetBudgetFilterCustomPeriod {
	if o == nil {
		return nil
	}
	p := &billingbudgetspb.BillingbudgetsBudgetBudgetFilterCustomPeriod{
		StartDate: BillingbudgetsBudgetBudgetFilterCustomPeriodStartDateToProto(o.StartDate),
		EndDate:   BillingbudgetsBudgetBudgetFilterCustomPeriodEndDateToProto(o.EndDate),
	}
	return p
}

// BudgetBudgetFilterCustomPeriodStartDateToProto converts a BudgetBudgetFilterCustomPeriodStartDate resource to its proto representation.
func BillingbudgetsBudgetBudgetFilterCustomPeriodStartDateToProto(o *billingbudgets.BudgetBudgetFilterCustomPeriodStartDate) *billingbudgetspb.BillingbudgetsBudgetBudgetFilterCustomPeriodStartDate {
	if o == nil {
		return nil
	}
	p := &billingbudgetspb.BillingbudgetsBudgetBudgetFilterCustomPeriodStartDate{
		Year:  dcl.ValueOrEmptyInt64(o.Year),
		Month: dcl.ValueOrEmptyInt64(o.Month),
		Day:   dcl.ValueOrEmptyInt64(o.Day),
	}
	return p
}

// BudgetBudgetFilterCustomPeriodEndDateToProto converts a BudgetBudgetFilterCustomPeriodEndDate resource to its proto representation.
func BillingbudgetsBudgetBudgetFilterCustomPeriodEndDateToProto(o *billingbudgets.BudgetBudgetFilterCustomPeriodEndDate) *billingbudgetspb.BillingbudgetsBudgetBudgetFilterCustomPeriodEndDate {
	if o == nil {
		return nil
	}
	p := &billingbudgetspb.BillingbudgetsBudgetBudgetFilterCustomPeriodEndDate{
		Year:  dcl.ValueOrEmptyInt64(o.Year),
		Month: dcl.ValueOrEmptyInt64(o.Month),
		Day:   dcl.ValueOrEmptyInt64(o.Day),
	}
	return p
}

// BudgetAmountToProto converts a BudgetAmount resource to its proto representation.
func BillingbudgetsBudgetAmountToProto(o *billingbudgets.BudgetAmount) *billingbudgetspb.BillingbudgetsBudgetAmount {
	if o == nil {
		return nil
	}
	p := &billingbudgetspb.BillingbudgetsBudgetAmount{
		SpecifiedAmount:  BillingbudgetsBudgetAmountSpecifiedAmountToProto(o.SpecifiedAmount),
		LastPeriodAmount: BillingbudgetsBudgetAmountLastPeriodAmountToProto(o.LastPeriodAmount),
	}
	return p
}

// BudgetAmountSpecifiedAmountToProto converts a BudgetAmountSpecifiedAmount resource to its proto representation.
func BillingbudgetsBudgetAmountSpecifiedAmountToProto(o *billingbudgets.BudgetAmountSpecifiedAmount) *billingbudgetspb.BillingbudgetsBudgetAmountSpecifiedAmount {
	if o == nil {
		return nil
	}
	p := &billingbudgetspb.BillingbudgetsBudgetAmountSpecifiedAmount{
		CurrencyCode: dcl.ValueOrEmptyString(o.CurrencyCode),
		Units:        dcl.ValueOrEmptyInt64(o.Units),
		Nanos:        dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// BudgetAmountLastPeriodAmountToProto converts a BudgetAmountLastPeriodAmount resource to its proto representation.
func BillingbudgetsBudgetAmountLastPeriodAmountToProto(o *billingbudgets.BudgetAmountLastPeriodAmount) *billingbudgetspb.BillingbudgetsBudgetAmountLastPeriodAmount {
	if o == nil {
		return nil
	}
	p := &billingbudgetspb.BillingbudgetsBudgetAmountLastPeriodAmount{}
	return p
}

// BudgetThresholdRulesToProto converts a BudgetThresholdRules resource to its proto representation.
func BillingbudgetsBudgetThresholdRulesToProto(o *billingbudgets.BudgetThresholdRules) *billingbudgetspb.BillingbudgetsBudgetThresholdRules {
	if o == nil {
		return nil
	}
	p := &billingbudgetspb.BillingbudgetsBudgetThresholdRules{
		ThresholdPercent: dcl.ValueOrEmptyDouble(o.ThresholdPercent),
		SpendBasis:       BillingbudgetsBudgetThresholdRulesSpendBasisEnumToProto(o.SpendBasis),
	}
	return p
}

// BudgetAllUpdatesRuleToProto converts a BudgetAllUpdatesRule resource to its proto representation.
func BillingbudgetsBudgetAllUpdatesRuleToProto(o *billingbudgets.BudgetAllUpdatesRule) *billingbudgetspb.BillingbudgetsBudgetAllUpdatesRule {
	if o == nil {
		return nil
	}
	p := &billingbudgetspb.BillingbudgetsBudgetAllUpdatesRule{
		PubsubTopic:                 dcl.ValueOrEmptyString(o.PubsubTopic),
		SchemaVersion:               dcl.ValueOrEmptyString(o.SchemaVersion),
		DisableDefaultIamRecipients: dcl.ValueOrEmptyBool(o.DisableDefaultIamRecipients),
	}
	for _, r := range o.MonitoringNotificationChannels {
		p.MonitoringNotificationChannels = append(p.MonitoringNotificationChannels, r)
	}
	return p
}

// BudgetToProto converts a Budget resource to its proto representation.
func BudgetToProto(resource *billingbudgets.Budget) *billingbudgetspb.BillingbudgetsBudget {
	p := &billingbudgetspb.BillingbudgetsBudget{
		Name:           dcl.ValueOrEmptyString(resource.Name),
		DisplayName:    dcl.ValueOrEmptyString(resource.DisplayName),
		BudgetFilter:   BillingbudgetsBudgetBudgetFilterToProto(resource.BudgetFilter),
		Amount:         BillingbudgetsBudgetAmountToProto(resource.Amount),
		Etag:           dcl.ValueOrEmptyString(resource.Etag),
		AllUpdatesRule: BillingbudgetsBudgetAllUpdatesRuleToProto(resource.AllUpdatesRule),
		BillingAccount: dcl.ValueOrEmptyString(resource.BillingAccount),
	}
	for _, r := range resource.ThresholdRules {
		p.ThresholdRules = append(p.ThresholdRules, BillingbudgetsBudgetThresholdRulesToProto(&r))
	}

	return p
}

// ApplyBudget handles the gRPC request by passing it to the underlying Budget Apply() method.
func (s *BudgetServer) applyBudget(ctx context.Context, c *billingbudgets.Client, request *billingbudgetspb.ApplyBillingbudgetsBudgetRequest) (*billingbudgetspb.BillingbudgetsBudget, error) {
	p := ProtoToBudget(request.GetResource())
	res, err := c.ApplyBudget(ctx, p)
	if err != nil {
		return nil, err
	}
	r := BudgetToProto(res)
	return r, nil
}

// ApplyBudget handles the gRPC request by passing it to the underlying Budget Apply() method.
func (s *BudgetServer) ApplyBillingbudgetsBudget(ctx context.Context, request *billingbudgetspb.ApplyBillingbudgetsBudgetRequest) (*billingbudgetspb.BillingbudgetsBudget, error) {
	cl, err := createConfigBudget(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyBudget(ctx, cl, request)
}

// DeleteBudget handles the gRPC request by passing it to the underlying Budget Delete() method.
func (s *BudgetServer) DeleteBillingbudgetsBudget(ctx context.Context, request *billingbudgetspb.DeleteBillingbudgetsBudgetRequest) (*emptypb.Empty, error) {

	cl, err := createConfigBudget(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteBudget(ctx, ProtoToBudget(request.GetResource()))

}

// ListBillingbudgetsBudget handles the gRPC request by passing it to the underlying BudgetList() method.
func (s *BudgetServer) ListBillingbudgetsBudget(ctx context.Context, request *billingbudgetspb.ListBillingbudgetsBudgetRequest) (*billingbudgetspb.ListBillingbudgetsBudgetResponse, error) {
	cl, err := createConfigBudget(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListBudget(ctx, request.BillingAccount)
	if err != nil {
		return nil, err
	}
	var protos []*billingbudgetspb.BillingbudgetsBudget
	for _, r := range resources.Items {
		rp := BudgetToProto(r)
		protos = append(protos, rp)
	}
	return &billingbudgetspb.ListBillingbudgetsBudgetResponse{Items: protos}, nil
}

func createConfigBudget(ctx context.Context, service_account_file string) (*billingbudgets.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return billingbudgets.NewClient(conf), nil
}
