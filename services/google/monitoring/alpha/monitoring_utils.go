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
// Package monitoring provides methods and types for managing monitoring GCP resources.
package alpha

func equalsMetricDescriptorValueType(m, n *MetricDescriptorValueTypeEnum) bool {
	if m == nil && n == nil {
		return true
	}
	v := *MetricDescriptorValueTypeEnumRef("STRING")
	if m == nil {
		return *n == v
	} else if n == nil {
		return *m == v
	}
	return *m == *n
}

func equalsMetricDescriptorLabelsValueType(m, n *MetricDescriptorLabelsValueTypeEnum) bool {
	if m == nil && n == nil {
		return true
	}
	v := *MetricDescriptorLabelsValueTypeEnumRef("STRING")
	if m == nil {
		return *n == v
	} else if n == nil {
		return *m == v
	}
	return *m == *n
}

func canonicalizeMetricDescriptorValueType(m, n interface{}) bool {
	if m == nil && n == nil {
		return true
	}

	mVal, _ := m.(*MetricDescriptorValueTypeEnum)
	nVal, _ := n.(*MetricDescriptorValueTypeEnum)
	return equalsMetricDescriptorValueType(mVal, nVal)
}

func canonicalizeMetricDescriptorLabelsValueType(m, n interface{}) bool {
	if m == nil && n == nil {
		return true
	}

	mVal, _ := m.(*MetricDescriptorLabelsValueTypeEnum)
	nVal, _ := n.(*MetricDescriptorLabelsValueTypeEnum)
	return equalsMetricDescriptorLabelsValueType(mVal, nVal)
}
