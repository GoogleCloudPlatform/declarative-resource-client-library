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
package dcl

import (
	"fmt"
	"reflect"
)

// Info is a struct that contains all information about the diff that's about to occur.
type Info struct {
	// Ignore + OutputOnly cause the diff checker to always return no-diff.
	Ignore          bool
	OutputOnly      bool
	IgnoredPrefixes []string
	Type            string

	FieldName string
}

// AddIndex adds an index to a FieldName and returns the same item.
// Info is always pass-by-value, so the original field name still exists.
func (i Info) AddIndex(index int) Info {
	newInfo := i
	newInfo.FieldName = newInfo.FieldName + fmt.Sprintf("[%v]", index)
	return newInfo
}

// FieldDiff contains all information about a diff that exists in the resource.
type FieldDiff struct {
	FieldName string
	Message   string
	Desired   interface{}
	Actual    interface{}

	ToAdd    []interface{}
	ToRemove []interface{}
}

func (d *FieldDiff) String() string {
	if d.Message != "" {
		return fmt.Sprintf("Field %s diff: %s", d.FieldName, d.Message)
	} else if len(d.ToAdd) != 0 || len(d.ToRemove) != 0 {
		return fmt.Sprintf("Field %s: add %v, remove %v", d.FieldName, d.ToAdd, d.ToRemove)
	}
	return fmt.Sprintf("Field %s: got %s, want %s", d.FieldName, stringValue(d.Actual), stringValue(d.Desired))
}

func stringValue(i interface{}) string {
	if reflect.ValueOf(i).Kind() == reflect.Ptr {
		ptrV := reflect.Indirect(reflect.ValueOf(i))
		return ptrV.String()
	}
	return fmt.Sprintf("%v", i)
}

// Diff takes in two interfaces and diffs them according to Info.
func Diff(desired, actual interface{}, info Info) ([]*FieldDiff, error) {
	var diffs []*FieldDiff
	// All Output-only fields should not be diffed.
	if info.OutputOnly || info.Ignore {
		return nil, nil
	}

	// If desired is set to nil, we do not care about the field.
	if desired == nil {
		return nil, nil
	}

	desiredType := ValueType(desired)

	if desiredType == "invalid" {
		return nil, nil
	}

	if info.Type == "ReferenceType" {
		dStr, aStr, err := strs(desired, actual)
		if err != nil {
			return nil, err
		}
		if !StringEqualsWithSelfLink(dStr, aStr) {
			diffs = append(diffs, &FieldDiff{FieldName: info.FieldName, Desired: dStr, Actual: aStr})
			return diffs, nil
		}
		return nil, nil
	} else if info.Type == "EnumType" {
		if !reflect.DeepEqual(desired, actual) {
			diffs = append(diffs, &FieldDiff{FieldName: info.FieldName, Desired: desired, Actual: actual})
			return diffs, nil
		}
		return nil, nil
	}

	switch desiredType {
	case "string":
		dStr, aStr, err := strs(desired, actual)
		if err != nil {
			return nil, err
		}
		if !StringCanonicalize(dStr, aStr) {
			diffs = append(diffs, &FieldDiff{FieldName: info.FieldName, Desired: dStr, Actual: aStr})
		}

	case "map":
		if !MapEquals(desired, actual, info.IgnoredPrefixes) {
			diffs = append(diffs, &FieldDiff{FieldName: info.FieldName, Desired: desired, Actual: actual})
		}

	case "int", "float64", "int64":
		if !reflect.DeepEqual(desired, actual) {
			diffs = append(diffs, &FieldDiff{FieldName: info.FieldName, Desired: desired, Actual: actual})
		}

	case "bool":
		dBool, aBool, err := bools(desired, actual)
		if err != nil {
			return nil, err
		}
		if !BoolCanonicalize(dBool, aBool) {
			diffs = append(diffs, &FieldDiff{FieldName: info.FieldName, Desired: dBool, Actual: aBool})
		}

	case "slice":
		dSlice, iSlice, err := slices(desired, actual)
		if err != nil {
			return nil, err
		}

		var arrDiffs []*FieldDiff
		if info.Type == "Set" {
			arrDiffs, err = setDiff(dSlice, iSlice, info)
		} else {
			arrDiffs, err = arrayDiff(dSlice, iSlice, info)
		}

		if err != nil {
			return nil, err
		}
		diffs = append(diffs, arrDiffs...)
	default:
		return nil, fmt.Errorf("no diffing logic exists for type: %q", desiredType)
	}

	return diffs, nil
}

func arrayDiff(desired, actual []interface{}, info Info) ([]*FieldDiff, error) {
	var diffs []*FieldDiff

	// Nothing to diff against.
	if actual == nil {
		return diffs, nil
	}

	if len(desired) != len(actual) {
		diffs = append(diffs, &FieldDiff{FieldName: info.FieldName, Message: fmt.Sprintf("different lengths: desired %d, actual %d", len(desired), len(actual))})
		return diffs, nil
	}

	for i, dItem := range desired {
		aItem := actual[i]
		diff, err := Diff(dItem, aItem, info.AddIndex(i))
		if err != nil {
			return nil, err
		}
		if diff != nil {
			diffs = append(diffs, diff...)
		}
	}
	return diffs, nil
}

func setDiff(desired, actual []interface{}, info Info) ([]*FieldDiff, error) {
	var diffs []*FieldDiff

	// Everything should be added.
	if actual == nil {
		diffs = append(diffs, &FieldDiff{FieldName: info.FieldName, ToAdd: desired})
		return diffs, nil
	}

	var toAdd, toRemove []interface{}

	for i, aItem := range actual {
		found := false
		for _, desItem := range desired {
			if ds, _ := Diff(desItem, aItem, info.AddIndex(i)); len(ds) == 0 {
				found = true
				break
			}
		}
		if !found {
			toRemove = append(toRemove, aItem)
		}
	}

	for i, dItem := range desired {
		found := false
		for _, actItem := range actual {
			if ds, _ := Diff(dItem, actItem, info.AddIndex(i)); len(ds) == 0 {
				found = true
				break
			}
		}
		if !found {
			toAdd = append(toAdd, dItem)
		}
	}

	if len(toAdd) > 0 || len(toRemove) > 0 {
		return []*FieldDiff{&FieldDiff{FieldName: info.FieldName, ToAdd: toAdd, ToRemove: toRemove}}, nil
	}
	return nil, nil
}

// ValueType returns the reflect-style kind of an interface or the underlying type of a pointer.
func ValueType(i interface{}) string {
	if reflect.ValueOf(i).Kind() == reflect.Ptr {
		return reflect.Indirect(reflect.ValueOf(i)).Kind().String()
	}
	return reflect.ValueOf(i).Kind().String()
}

func strs(d, i interface{}) (*string, *string, error) {
	dStr, err := str(d)
	if err != nil {
		return nil, nil, err
	}

	iStr, err := str(i)
	if err != nil {
		return nil, nil, err
	}

	return dStr, iStr, nil
}

func str(d interface{}) (*string, error) {
	dPtr, dOk := d.(*string)
	if !dOk {
		dStr, dOk2 := d.(string)
		if !dOk2 {
			return nil, fmt.Errorf("was given non string %v", d)
		}
		dPtr = String(dStr)
	}
	return dPtr, nil
}

func bools(d, i interface{}) (*bool, *bool, error) {
	dBool, err := boolean(d)
	if err != nil {
		return nil, nil, err
	}

	iBool, err := boolean(i)
	if err != nil {
		return nil, nil, err
	}

	return dBool, iBool, nil
}

func boolean(d interface{}) (*bool, error) {
	dPtr, dOk := d.(*bool)
	if !dOk {
		return nil, nil
	}
	return dPtr, nil
}

func mapVal(d interface{}) (map[string]string, error) {
	dMap, dOk := d.(map[string]string)
	if !dOk {
		return nil, fmt.Errorf("value is not a map[string]string")
	}
	return dMap, nil
}

func slices(d, i interface{}) ([]interface{}, []interface{}, error) {
	dSlice, err := slice(d)
	if err != nil {
		return nil, nil, err
	}

	iSlice, err := slice(i)
	if err != nil {
		return nil, nil, err
	}

	return dSlice, iSlice, nil
}

func slice(slice interface{}) ([]interface{}, error) {
	// Keep the distinction between nil and empty slice input
	// This isn't going to be an error though.
	if slice == nil {
		return nil, nil
	}

	s := reflect.ValueOf(slice)

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret, nil
}
