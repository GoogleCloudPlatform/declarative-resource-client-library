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
package serialization

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/hcl/hcl/printer"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// InstanceStateToHCL converts a terraform.InstanceState into the HCL that represents the maximally hydrated form.  The generated
// HCL will not include any references to other resources, since that information cannot be reconstructed without a full view
// of all resources present..
func InstanceStateToHCL(state *terraform.InstanceState, info *terraform.InstanceInfo, provider *tfschema.Provider) (string, error) {
	providerSchema := provider.ResourcesMap[info.Type].Schema
	str, err := resourceOrSubresourceHCL(providerSchema, state.Attributes)
	if err != nil {
		return "", err
	}
	str = fmt.Sprintf("resource %q %q {\n%s\n}\n", info.Type, info.Id, str)
	hBytes, err := printer.Format([]byte(str))
	if err != nil {
		return "", fmt.Errorf("could not pretty print hcl: %w", err)
	}
	return string(hBytes), nil
}

func resourceOrSubresourceHCL(schema map[string]*tfschema.Schema, attributes map[string]string) (string, error) {
	var hcl strings.Builder
	keys := make([]string, 0, len(schema))
	for k := range schema {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, aname := range keys {
		aschema := schema[aname]
		if !aschema.Optional && !aschema.Required {
			continue
		}
		switch aschema.Type {
		case tfschema.TypeFloat:
			fallthrough
		case tfschema.TypeInt:
			fallthrough
		case tfschema.TypeBool:
			if attributes[aname] == "" {
				continue
			}
			hcl.WriteString(fmt.Sprintf("%s = %s\n", aname, attributes[aname]))
		case tfschema.TypeString:
			if attributes[aname] == "" {
				continue
			}
			hcl.WriteString(fmt.Sprintf("%s = %q\n", aname, attributes[aname]))
		// for each non-primitive, find if there are any values and set them.
		case tfschema.TypeMap:
			if attributes[aname+".%"] == "0" {
				continue
			}
			hcl.WriteString(fmt.Sprintf("%s {\n", aname))
			m, err := mapFromPrefix(attributes, aname)
			if err != nil {
				return "", err
			}
			for k, v := range m {
				hcl.WriteString(fmt.Sprintf("%s = %q\n", k, v))
			}
			hcl.WriteString("}\n")
		case tfschema.TypeSet:
			fallthrough
		case tfschema.TypeList:
			if attributes[aname+".#"] == "0" {
				continue
			}
			if subtype, ok := aschema.Elem.(*tfschema.Schema); ok {
				hcl.WriteString(fmt.Sprintf("%s = [", aname))
				var l []string
				var err error
				if aschema.Type == tfschema.TypeList {
					l, err = listFromPrefix(attributes, aname)
					if err != nil {
						return "", err
					}
				} else if aschema.Type == tfschema.TypeSet {
					l, err = setFromPrefix(attributes, aname)
					if err != nil {
						return "", err
					}
				}
				for _, v := range l {
					switch subtype.Type {
					case tfschema.TypeFloat:
						fallthrough
					case tfschema.TypeInt:
						fallthrough
					case tfschema.TypeBool:
						hcl.WriteString(fmt.Sprintf("%s, ", v))
					case tfschema.TypeString:
						hcl.WriteString(fmt.Sprintf("%q, ", v))
					}
				}
				hcl.WriteString("]\n")
			} else if subtype, ok := aschema.Elem.(*tfschema.Resource); ok {
				cnt, err := strconv.Atoi(attributes[aname+".#"])
				if err != nil {
					return "", fmt.Errorf("could not parse count of %s, %w", aname, err)
				}
				for i := 0; i < cnt; i++ {
					subAttrs, err := mapsFromPrefix(attributes, fmt.Sprintf("%s.%d", aname, i))
					if err != nil {
						return "", fmt.Errorf("could not get subresource attributes for %s: %w", aname, err)
					}
					subresource, err := resourceOrSubresourceHCL(subtype.Schema, subAttrs)
					if err != nil {
						return "", fmt.Errorf("could not create subresource %s: %w", aname, err)
					}
					hcl.WriteString(fmt.Sprintf("%s {\n%s\n}\n", aname, subresource))
				}
			}
		}
	}
	return hcl.String(), nil
}

func mapsFromPrefix(attributes map[string]string, prefix string) (map[string]string, error) {
	a := make(map[string]string)
	for k, v := range attributes {
		if strings.HasPrefix(k, prefix+".") {
			a[strings.TrimPrefix(k, prefix+".")] = v
		}
	}
	return a, nil
}

func setFromPrefix(attributes map[string]string, prefix string) ([]string, error) {
	size, err := strconv.Atoi(attributes[prefix+".#"])
	if err != nil {
		return nil, fmt.Errorf("could not parse size of list %s: %w", prefix, err)
	}
	out := make([]string, 0, size)
	for k, v := range attributes {
		if strings.HasPrefix(k, prefix+".") && k != prefix+".#" {
			out = append(out, v)
		}
	}
	// sort the list so that the ordering is deterministic for tests
	sort.Strings(out)
	return out, nil
}

func listFromPrefix(attributes map[string]string, prefix string) ([]string, error) {
	size, err := strconv.Atoi(attributes[prefix+".#"])
	if err != nil {
		return nil, fmt.Errorf("could not parse size of list %s: %w", prefix, err)
	}
	out := make([]string, size)
	for k, v := range attributes {
		if strings.HasPrefix(k, prefix+".") && k != prefix+".#" {
			kparts := strings.Split(k, ".")
			c, err := strconv.Atoi(kparts[len(kparts)-1])
			if err != nil {
				return nil, fmt.Errorf("could not parse index of %s: %w", k, err)
			}
			out[c] = v
		}
	}
	return out, nil
}

func mapFromPrefix(attributes map[string]string, prefix string) (map[string]string, error) {
	size, err := strconv.Atoi(attributes[prefix+".%"])
	if err != nil {
		return nil, fmt.Errorf("could not parse size of map %s: %w", prefix, err)
	}
	out := make(map[string]string, size)
	for k, v := range attributes {
		if strings.HasPrefix(k, prefix+".") && k != prefix+".%" {
			kparts := strings.Split(k, ".")
			out[kparts[len(kparts)-1]] = v
		}
	}
	return out, nil
}
