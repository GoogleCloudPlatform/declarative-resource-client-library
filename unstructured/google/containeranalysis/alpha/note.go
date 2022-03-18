// Copyright 2022 Google LLC. All Rights Reserved.
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
package containeranalysis

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeranalysis/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Note struct{}

func NoteToUnstructured(r *dclService.Note) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "containeranalysis",
			Version: "alpha",
			Type:    "Note",
		},
		Object: make(map[string]interface{}),
	}
	if r.Attestation != nil && r.Attestation != dclService.EmptyNoteAttestation {
		rAttestation := make(map[string]interface{})
		if r.Attestation.Hint != nil && r.Attestation.Hint != dclService.EmptyNoteAttestationHint {
			rAttestationHint := make(map[string]interface{})
			if r.Attestation.Hint.HumanReadableName != nil {
				rAttestationHint["humanReadableName"] = *r.Attestation.Hint.HumanReadableName
			}
			rAttestation["hint"] = rAttestationHint
		}
		u.Object["attestation"] = rAttestation
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Deployment != nil && r.Deployment != dclService.EmptyNoteDeployment {
		rDeployment := make(map[string]interface{})
		var rDeploymentResourceUri []interface{}
		for _, rDeploymentResourceUriVal := range r.Deployment.ResourceUri {
			rDeploymentResourceUri = append(rDeploymentResourceUri, rDeploymentResourceUriVal)
		}
		rDeployment["resourceUri"] = rDeploymentResourceUri
		u.Object["deployment"] = rDeployment
	}
	if r.Discovery != nil && r.Discovery != dclService.EmptyNoteDiscovery {
		rDiscovery := make(map[string]interface{})
		if r.Discovery.AnalysisKind != nil {
			rDiscovery["analysisKind"] = string(*r.Discovery.AnalysisKind)
		}
		u.Object["discovery"] = rDiscovery
	}
	if r.ExpirationTime != nil {
		u.Object["expirationTime"] = *r.ExpirationTime
	}
	if r.Image != nil && r.Image != dclService.EmptyNoteImage {
		rImage := make(map[string]interface{})
		if r.Image.Fingerprint != nil && r.Image.Fingerprint != dclService.EmptyNoteImageFingerprint {
			rImageFingerprint := make(map[string]interface{})
			if r.Image.Fingerprint.V1Name != nil {
				rImageFingerprint["v1Name"] = *r.Image.Fingerprint.V1Name
			}
			var rImageFingerprintV2Blob []interface{}
			for _, rImageFingerprintV2BlobVal := range r.Image.Fingerprint.V2Blob {
				rImageFingerprintV2Blob = append(rImageFingerprintV2Blob, rImageFingerprintV2BlobVal)
			}
			rImageFingerprint["v2Blob"] = rImageFingerprintV2Blob
			if r.Image.Fingerprint.V2Name != nil {
				rImageFingerprint["v2Name"] = *r.Image.Fingerprint.V2Name
			}
			rImage["fingerprint"] = rImageFingerprint
		}
		if r.Image.ResourceUrl != nil {
			rImage["resourceUrl"] = *r.Image.ResourceUrl
		}
		u.Object["image"] = rImage
	}
	if r.LongDescription != nil {
		u.Object["longDescription"] = *r.LongDescription
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Package != nil && r.Package != dclService.EmptyNotePackage {
		rPackage := make(map[string]interface{})
		var rPackageDistribution []interface{}
		for _, rPackageDistributionVal := range r.Package.Distribution {
			rPackageDistributionObject := make(map[string]interface{})
			if rPackageDistributionVal.Architecture != nil {
				rPackageDistributionObject["architecture"] = string(*rPackageDistributionVal.Architecture)
			}
			if rPackageDistributionVal.CpeUri != nil {
				rPackageDistributionObject["cpeUri"] = *rPackageDistributionVal.CpeUri
			}
			if rPackageDistributionVal.Description != nil {
				rPackageDistributionObject["description"] = *rPackageDistributionVal.Description
			}
			if rPackageDistributionVal.LatestVersion != nil && rPackageDistributionVal.LatestVersion != dclService.EmptyNotePackageDistributionLatestVersion {
				rPackageDistributionValLatestVersion := make(map[string]interface{})
				if rPackageDistributionVal.LatestVersion.Epoch != nil {
					rPackageDistributionValLatestVersion["epoch"] = *rPackageDistributionVal.LatestVersion.Epoch
				}
				if rPackageDistributionVal.LatestVersion.FullName != nil {
					rPackageDistributionValLatestVersion["fullName"] = *rPackageDistributionVal.LatestVersion.FullName
				}
				if rPackageDistributionVal.LatestVersion.Kind != nil {
					rPackageDistributionValLatestVersion["kind"] = string(*rPackageDistributionVal.LatestVersion.Kind)
				}
				if rPackageDistributionVal.LatestVersion.Name != nil {
					rPackageDistributionValLatestVersion["name"] = *rPackageDistributionVal.LatestVersion.Name
				}
				if rPackageDistributionVal.LatestVersion.Revision != nil {
					rPackageDistributionValLatestVersion["revision"] = *rPackageDistributionVal.LatestVersion.Revision
				}
				rPackageDistributionObject["latestVersion"] = rPackageDistributionValLatestVersion
			}
			if rPackageDistributionVal.Maintainer != nil {
				rPackageDistributionObject["maintainer"] = *rPackageDistributionVal.Maintainer
			}
			if rPackageDistributionVal.Url != nil {
				rPackageDistributionObject["url"] = *rPackageDistributionVal.Url
			}
			rPackageDistribution = append(rPackageDistribution, rPackageDistributionObject)
		}
		rPackage["distribution"] = rPackageDistribution
		if r.Package.Name != nil {
			rPackage["name"] = *r.Package.Name
		}
		u.Object["package"] = rPackage
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	var rRelatedUrl []interface{}
	for _, rRelatedUrlVal := range r.RelatedUrl {
		rRelatedUrlObject := make(map[string]interface{})
		if rRelatedUrlVal.Label != nil {
			rRelatedUrlObject["label"] = *rRelatedUrlVal.Label
		}
		if rRelatedUrlVal.Url != nil {
			rRelatedUrlObject["url"] = *rRelatedUrlVal.Url
		}
		rRelatedUrl = append(rRelatedUrl, rRelatedUrlObject)
	}
	u.Object["relatedUrl"] = rRelatedUrl
	if r.ShortDescription != nil {
		u.Object["shortDescription"] = *r.ShortDescription
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToNote(u *unstructured.Resource) (*dclService.Note, error) {
	r := &dclService.Note{}
	if _, ok := u.Object["attestation"]; ok {
		if rAttestation, ok := u.Object["attestation"].(map[string]interface{}); ok {
			r.Attestation = &dclService.NoteAttestation{}
			if _, ok := rAttestation["hint"]; ok {
				if rAttestationHint, ok := rAttestation["hint"].(map[string]interface{}); ok {
					r.Attestation.Hint = &dclService.NoteAttestationHint{}
					if _, ok := rAttestationHint["humanReadableName"]; ok {
						if s, ok := rAttestationHint["humanReadableName"].(string); ok {
							r.Attestation.Hint.HumanReadableName = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Attestation.Hint.HumanReadableName: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Attestation.Hint: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Attestation: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["deployment"]; ok {
		if rDeployment, ok := u.Object["deployment"].(map[string]interface{}); ok {
			r.Deployment = &dclService.NoteDeployment{}
			if _, ok := rDeployment["resourceUri"]; ok {
				if s, ok := rDeployment["resourceUri"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Deployment.ResourceUri = append(r.Deployment.ResourceUri, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Deployment.ResourceUri: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Deployment: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["discovery"]; ok {
		if rDiscovery, ok := u.Object["discovery"].(map[string]interface{}); ok {
			r.Discovery = &dclService.NoteDiscovery{}
			if _, ok := rDiscovery["analysisKind"]; ok {
				if s, ok := rDiscovery["analysisKind"].(string); ok {
					r.Discovery.AnalysisKind = dclService.NoteDiscoveryAnalysisKindEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.Discovery.AnalysisKind: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Discovery: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["expirationTime"]; ok {
		if s, ok := u.Object["expirationTime"].(string); ok {
			r.ExpirationTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ExpirationTime: expected string")
		}
	}
	if _, ok := u.Object["image"]; ok {
		if rImage, ok := u.Object["image"].(map[string]interface{}); ok {
			r.Image = &dclService.NoteImage{}
			if _, ok := rImage["fingerprint"]; ok {
				if rImageFingerprint, ok := rImage["fingerprint"].(map[string]interface{}); ok {
					r.Image.Fingerprint = &dclService.NoteImageFingerprint{}
					if _, ok := rImageFingerprint["v1Name"]; ok {
						if s, ok := rImageFingerprint["v1Name"].(string); ok {
							r.Image.Fingerprint.V1Name = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Image.Fingerprint.V1Name: expected string")
						}
					}
					if _, ok := rImageFingerprint["v2Blob"]; ok {
						if s, ok := rImageFingerprint["v2Blob"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.Image.Fingerprint.V2Blob = append(r.Image.Fingerprint.V2Blob, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Image.Fingerprint.V2Blob: expected []interface{}")
						}
					}
					if _, ok := rImageFingerprint["v2Name"]; ok {
						if s, ok := rImageFingerprint["v2Name"].(string); ok {
							r.Image.Fingerprint.V2Name = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Image.Fingerprint.V2Name: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Image.Fingerprint: expected map[string]interface{}")
				}
			}
			if _, ok := rImage["resourceUrl"]; ok {
				if s, ok := rImage["resourceUrl"].(string); ok {
					r.Image.ResourceUrl = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Image.ResourceUrl: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Image: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["longDescription"]; ok {
		if s, ok := u.Object["longDescription"].(string); ok {
			r.LongDescription = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.LongDescription: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["package"]; ok {
		if rPackage, ok := u.Object["package"].(map[string]interface{}); ok {
			r.Package = &dclService.NotePackage{}
			if _, ok := rPackage["distribution"]; ok {
				if s, ok := rPackage["distribution"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rPackageDistribution dclService.NotePackageDistribution
							if _, ok := objval["architecture"]; ok {
								if s, ok := objval["architecture"].(string); ok {
									rPackageDistribution.Architecture = dclService.NotePackageDistributionArchitectureEnumRef(s)
								} else {
									return nil, fmt.Errorf("rPackageDistribution.Architecture: expected string")
								}
							}
							if _, ok := objval["cpeUri"]; ok {
								if s, ok := objval["cpeUri"].(string); ok {
									rPackageDistribution.CpeUri = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rPackageDistribution.CpeUri: expected string")
								}
							}
							if _, ok := objval["description"]; ok {
								if s, ok := objval["description"].(string); ok {
									rPackageDistribution.Description = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rPackageDistribution.Description: expected string")
								}
							}
							if _, ok := objval["latestVersion"]; ok {
								if rPackageDistributionLatestVersion, ok := objval["latestVersion"].(map[string]interface{}); ok {
									rPackageDistribution.LatestVersion = &dclService.NotePackageDistributionLatestVersion{}
									if _, ok := rPackageDistributionLatestVersion["epoch"]; ok {
										if i, ok := rPackageDistributionLatestVersion["epoch"].(int64); ok {
											rPackageDistribution.LatestVersion.Epoch = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rPackageDistribution.LatestVersion.Epoch: expected int64")
										}
									}
									if _, ok := rPackageDistributionLatestVersion["fullName"]; ok {
										if s, ok := rPackageDistributionLatestVersion["fullName"].(string); ok {
											rPackageDistribution.LatestVersion.FullName = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rPackageDistribution.LatestVersion.FullName: expected string")
										}
									}
									if _, ok := rPackageDistributionLatestVersion["kind"]; ok {
										if s, ok := rPackageDistributionLatestVersion["kind"].(string); ok {
											rPackageDistribution.LatestVersion.Kind = dclService.NotePackageDistributionLatestVersionKindEnumRef(s)
										} else {
											return nil, fmt.Errorf("rPackageDistribution.LatestVersion.Kind: expected string")
										}
									}
									if _, ok := rPackageDistributionLatestVersion["name"]; ok {
										if s, ok := rPackageDistributionLatestVersion["name"].(string); ok {
											rPackageDistribution.LatestVersion.Name = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rPackageDistribution.LatestVersion.Name: expected string")
										}
									}
									if _, ok := rPackageDistributionLatestVersion["revision"]; ok {
										if s, ok := rPackageDistributionLatestVersion["revision"].(string); ok {
											rPackageDistribution.LatestVersion.Revision = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rPackageDistribution.LatestVersion.Revision: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rPackageDistribution.LatestVersion: expected map[string]interface{}")
								}
							}
							if _, ok := objval["maintainer"]; ok {
								if s, ok := objval["maintainer"].(string); ok {
									rPackageDistribution.Maintainer = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rPackageDistribution.Maintainer: expected string")
								}
							}
							if _, ok := objval["url"]; ok {
								if s, ok := objval["url"].(string); ok {
									rPackageDistribution.Url = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rPackageDistribution.Url: expected string")
								}
							}
							r.Package.Distribution = append(r.Package.Distribution, rPackageDistribution)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Package.Distribution: expected []interface{}")
				}
			}
			if _, ok := rPackage["name"]; ok {
				if s, ok := rPackage["name"].(string); ok {
					r.Package.Name = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Package.Name: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Package: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["relatedUrl"]; ok {
		if s, ok := u.Object["relatedUrl"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rRelatedUrl dclService.NoteRelatedUrl
					if _, ok := objval["label"]; ok {
						if s, ok := objval["label"].(string); ok {
							rRelatedUrl.Label = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rRelatedUrl.Label: expected string")
						}
					}
					if _, ok := objval["url"]; ok {
						if s, ok := objval["url"].(string); ok {
							rRelatedUrl.Url = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rRelatedUrl.Url: expected string")
						}
					}
					r.RelatedUrl = append(r.RelatedUrl, rRelatedUrl)
				}
			}
		} else {
			return nil, fmt.Errorf("r.RelatedUrl: expected []interface{}")
		}
	}
	if _, ok := u.Object["shortDescription"]; ok {
		if s, ok := u.Object["shortDescription"].(string); ok {
			r.ShortDescription = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ShortDescription: expected string")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	return r, nil
}

func GetNote(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNote(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetNote(ctx, r)
	if err != nil {
		return nil, err
	}
	return NoteToUnstructured(r), nil
}

func ListNote(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListNote(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, NoteToUnstructured(r))
		}
		if !l.HasNext() {
			break
		}
		if err := l.Next(ctx, c); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func ApplyNote(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNote(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToNote(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyNote(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return NoteToUnstructured(r), nil
}

func NoteHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNote(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToNote(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyNote(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteNote(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNote(u)
	if err != nil {
		return err
	}
	return c.DeleteNote(ctx, r)
}

func NoteID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToNote(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Note) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"containeranalysis",
		"Note",
		"alpha",
	}
}

func (r *Note) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Note) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Note) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Note) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Note) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Note) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Note) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetNote(ctx, config, resource)
}

func (r *Note) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyNote(ctx, config, resource, opts...)
}

func (r *Note) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return NoteHasDiff(ctx, config, resource, opts...)
}

func (r *Note) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteNote(ctx, config, resource)
}

func (r *Note) ID(resource *unstructured.Resource) (string, error) {
	return NoteID(resource)
}

func init() {
	unstructured.Register(&Note{})
}
