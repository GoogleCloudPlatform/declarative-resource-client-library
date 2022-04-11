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
package beta

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *Table) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "dataset"); err != nil {
		return err
	}
	if err := dcl.Required(r, "project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Model) {
		if err := r.Model.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Schema) {
		if err := r.Schema.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.TimePartitioning) {
		if err := r.TimePartitioning.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RangePartitioning) {
		if err := r.RangePartitioning.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Clustering) {
		if err := r.Clustering.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.View) {
		if err := r.View.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MaterializedView) {
		if err := r.MaterializedView.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExternalDataConfiguration) {
		if err := r.ExternalDataConfiguration.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.StreamingBuffer) {
		if err := r.StreamingBuffer.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.EncryptionConfiguration) {
		if err := r.EncryptionConfiguration.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SnapshotDefinition) {
		if err := r.SnapshotDefinition.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *TableModel) validate() error {
	if !dcl.IsEmptyValueIndirect(r.ModelOptions) {
		if err := r.ModelOptions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *TableModelModelOptions) validate() error {
	return nil
}
func (r *TableModelTrainingRuns) validate() error {
	if !dcl.IsEmptyValueIndirect(r.TrainingOptions) {
		if err := r.TrainingOptions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *TableModelTrainingRunsTrainingOptions) validate() error {
	return nil
}
func (r *TableModelTrainingRunsIterationResults) validate() error {
	return nil
}
func (r *TableSchema) validate() error {
	return nil
}
func (r *TableGooglecloudbigqueryv2Tablefieldschema) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "type"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Categories) {
		if err := r.Categories.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PolicyTags) {
		if err := r.PolicyTags.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *TableGooglecloudbigqueryv2TablefieldschemaCategories) validate() error {
	return nil
}
func (r *TableGooglecloudbigqueryv2TablefieldschemaPolicyTags) validate() error {
	return nil
}
func (r *TableTimePartitioning) validate() error {
	if err := dcl.Required(r, "type"); err != nil {
		return err
	}
	return nil
}
func (r *TableRangePartitioning) validate() error {
	if err := dcl.Required(r, "field"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Range) {
		if err := r.Range.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *TableRangePartitioningRange) validate() error {
	if err := dcl.Required(r, "start"); err != nil {
		return err
	}
	if err := dcl.Required(r, "end"); err != nil {
		return err
	}
	if err := dcl.Required(r, "interval"); err != nil {
		return err
	}
	return nil
}
func (r *TableClustering) validate() error {
	return nil
}
func (r *TableView) validate() error {
	if err := dcl.Required(r, "query"); err != nil {
		return err
	}
	return nil
}
func (r *TableViewUserDefinedFunctionResources) validate() error {
	return nil
}
func (r *TableMaterializedView) validate() error {
	if err := dcl.Required(r, "query"); err != nil {
		return err
	}
	return nil
}
func (r *TableExternalDataConfiguration) validate() error {
	if err := dcl.Required(r, "sourceUris"); err != nil {
		return err
	}
	if err := dcl.Required(r, "sourceFormat"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Schema) {
		if err := r.Schema.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CsvOptions) {
		if err := r.CsvOptions.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.BigtableOptions) {
		if err := r.BigtableOptions.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.GoogleSheetsOptions) {
		if err := r.GoogleSheetsOptions.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HivePartitioningOptions) {
		if err := r.HivePartitioningOptions.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ValueConversionModes) {
		if err := r.ValueConversionModes.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.AvroOptions) {
		if err := r.AvroOptions.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ParquetOptions) {
		if err := r.ParquetOptions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *TableExternalDataConfigurationSchema) validate() error {
	return nil
}
func (r *TableExternalDataConfigurationCsvOptions) validate() error {
	return nil
}
func (r *TableExternalDataConfigurationBigtableOptions) validate() error {
	return nil
}
func (r *TableExternalDataConfigurationBigtableOptionsColumnFamilies) validate() error {
	return nil
}
func (r *TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns) validate() error {
	if err := dcl.Required(r, "qualifierEncoded"); err != nil {
		return err
	}
	return nil
}
func (r *TableExternalDataConfigurationGoogleSheetsOptions) validate() error {
	return nil
}
func (r *TableExternalDataConfigurationHivePartitioningOptions) validate() error {
	return nil
}
func (r *TableExternalDataConfigurationValueConversionModes) validate() error {
	return nil
}
func (r *TableExternalDataConfigurationAvroOptions) validate() error {
	return nil
}
func (r *TableExternalDataConfigurationParquetOptions) validate() error {
	return nil
}
func (r *TableStreamingBuffer) validate() error {
	return nil
}
func (r *TableEncryptionConfiguration) validate() error {
	return nil
}
func (r *TableSnapshotDefinition) validate() error {
	if err := dcl.Required(r, "table"); err != nil {
		return err
	}
	if err := dcl.Required(r, "dataset"); err != nil {
		return err
	}
	if err := dcl.Required(r, "project"); err != nil {
		return err
	}
	if err := dcl.Required(r, "snapshotTime"); err != nil {
		return err
	}
	return nil
}
func (r *Table) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://bigquery.googleapis.com/bigquery/v2/", params)
}

func (r *Table) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"dataset": dcl.ValueOrEmptyString(nr.Dataset),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/datasets/{{dataset}}/tables/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Table) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"dataset": dcl.ValueOrEmptyString(nr.Dataset),
	}
	return dcl.URL("projects/{{project}}/datasets/{{dataset}}/tables", nr.basePath(), userBasePath, params), nil

}

func (r *Table) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"dataset": dcl.ValueOrEmptyString(nr.Dataset),
	}
	return dcl.URL("projects/{{project}}/datasets/{{dataset}}/tables", nr.basePath(), userBasePath, params), nil

}

func (r *Table) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"dataset": dcl.ValueOrEmptyString(nr.Dataset),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/datasets/{{dataset}}/tables/{{name}}", nr.basePath(), userBasePath, params), nil
}

// tableApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type tableApiOperation interface {
	do(context.Context, *Table, *Client) error
}

// newUpdateTablePatchTableRequest creates a request for an
// Table resource's PatchTable update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateTablePatchTableRequest(ctx context.Context, f *Table, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.FriendlyName; !dcl.IsEmptyValueIndirect(v) {
		req["friendlyName"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v, err := expandTableSchema(c, f.Schema, res); err != nil {
		return nil, fmt.Errorf("error expanding Schema into schema: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["schema"] = v
	}
	if v, err := expandTableTimePartitioning(c, f.TimePartitioning, res); err != nil {
		return nil, fmt.Errorf("error expanding TimePartitioning into timePartitioning: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["timePartitioning"] = v
	}
	if v, err := expandTableRangePartitioning(c, f.RangePartitioning, res); err != nil {
		return nil, fmt.Errorf("error expanding RangePartitioning into rangePartitioning: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["rangePartitioning"] = v
	}
	if v, err := expandTableClustering(c, f.Clustering, res); err != nil {
		return nil, fmt.Errorf("error expanding Clustering into clustering: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["clustering"] = v
	}
	if v := f.RequirePartitionFilter; !dcl.IsEmptyValueIndirect(v) {
		req["requirePartitionFilter"] = v
	}
	if v := f.ExpirationTime; !dcl.IsEmptyValueIndirect(v) {
		req["expirationTime"] = v
	}
	if v, err := expandTableView(c, f.View, res); err != nil {
		return nil, fmt.Errorf("error expanding View into view: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["view"] = v
	}
	if v, err := expandTableMaterializedView(c, f.MaterializedView, res); err != nil {
		return nil, fmt.Errorf("error expanding MaterializedView into materializedView: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["materializedView"] = v
	}
	if v, err := expandTableExternalDataConfiguration(c, f.ExternalDataConfiguration, res); err != nil {
		return nil, fmt.Errorf("error expanding ExternalDataConfiguration into externalDataConfiguration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["externalDataConfiguration"] = v
	}
	if v, err := expandTableEncryptionConfiguration(c, f.EncryptionConfiguration, res); err != nil {
		return nil, fmt.Errorf("error expanding EncryptionConfiguration into encryptionConfiguration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["encryptionConfiguration"] = v
	}
	if v, err := expandTableSnapshotDefinition(c, f.SnapshotDefinition, res); err != nil {
		return nil, fmt.Errorf("error expanding SnapshotDefinition into snapshotDefinition: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["snapshotDefinition"] = v
	}
	if v := f.DefaultCollation; !dcl.IsEmptyValueIndirect(v) {
		req["defaultCollation"] = v
	}
	b, err := c.getTableRaw(ctx, f)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	rawEtag, err := dcl.GetMapEntry(
		m,
		[]string{"etag"},
	)
	if err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "Failed to fetch from JSON Path: %v", err)
	} else {
		req["etag"] = rawEtag.(string)
	}
	return req, nil
}

// marshalUpdateTablePatchTableRequest converts the update into
// the final JSON request body.
func marshalUpdateTablePatchTableRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	dcl.MoveMapEntry(
		m,
		[]string{"name"},
		[]string{"tableReference", "tableId"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"dataset"},
		[]string{"tableReference", "datasetId"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"project"},
		[]string{"tableReference", "projectId"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"snapshotDefinition", "table"},
		[]string{"snapshotDefinition", "tableReference", "tableId"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"snapshotDefinition", "dataset"},
		[]string{"snapshotDefinition", "tableReference", "datasetId"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"snapshotDefinition", "project"},
		[]string{"snapshotDefinition", "tableReference", "projectId"},
	)
	return json.Marshal(m)
}

type updateTablePatchTableOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateTablePatchTableOperation) do(ctx context.Context, r *Table, c *Client) error {
	_, err := c.GetTable(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "PatchTable")
	if err != nil {
		return err
	}

	req, err := newUpdateTablePatchTableRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateTablePatchTableRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listTableRaw(ctx context.Context, r *Table, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != TableMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listTableOperation struct {
	Tables []map[string]interface{} `json:"tables"`
	Token  string                   `json:"nextPageToken"`
}

func (c *Client) listTable(ctx context.Context, r *Table, pageToken string, pageSize int32) ([]*Table, string, error) {
	b, err := c.listTableRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listTableOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Table
	for _, v := range m.Tables {
		res, err := unmarshalMapTable(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Dataset = r.Dataset
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllTable(ctx context.Context, f func(*Table) bool, resources []*Table) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteTable(ctx, res)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", strings.Join(errors, "\n"))
	} else {
		return nil
	}
}

type deleteTableOperation struct{}

func (op *deleteTableOperation) do(ctx context.Context, r *Table, c *Client) error {
	r, err := c.GetTable(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Table not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetTable checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Table: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetTable(ctx, r)
		if !dcl.IsNotFound(err) {
			if i == maxRetry {
				return dcl.NotDeletedError{ExistingResource: r}
			}
			time.Sleep(1000 * time.Millisecond)
		} else {
			break
		}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createTableOperation struct {
	response map[string]interface{}
}

func (op *createTableOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createTableOperation) do(ctx context.Context, r *Table, c *Client) error {
	c.Config.Logger.InfoWithContextf(ctx, "Attempting to create %v", r)
	u, err := r.createURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o

	if _, err := c.GetTable(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getTableRaw(ctx context.Context, r *Table) ([]byte, error) {

	u, err := r.getURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) tableDiffsForRawDesired(ctx context.Context, rawDesired *Table, opts ...dcl.ApplyOption) (initial, desired *Table, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Table
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Table); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Table, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetTable(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Table resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Table resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Table resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeTableDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Table: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Table: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractTableFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeTableInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Table: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeTableDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Table: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffTable(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeTableInitialState(rawInitial, rawDesired *Table) (*Table, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeTableDesiredState(rawDesired, rawInitial *Table, opts ...dcl.ApplyOption) (*Table, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Model = canonicalizeTableModel(rawDesired.Model, nil, opts...)
		rawDesired.Schema = canonicalizeTableSchema(rawDesired.Schema, nil, opts...)
		rawDesired.TimePartitioning = canonicalizeTableTimePartitioning(rawDesired.TimePartitioning, nil, opts...)
		rawDesired.RangePartitioning = canonicalizeTableRangePartitioning(rawDesired.RangePartitioning, nil, opts...)
		rawDesired.Clustering = canonicalizeTableClustering(rawDesired.Clustering, nil, opts...)
		rawDesired.View = canonicalizeTableView(rawDesired.View, nil, opts...)
		rawDesired.MaterializedView = canonicalizeTableMaterializedView(rawDesired.MaterializedView, nil, opts...)
		rawDesired.ExternalDataConfiguration = canonicalizeTableExternalDataConfiguration(rawDesired.ExternalDataConfiguration, nil, opts...)
		rawDesired.StreamingBuffer = canonicalizeTableStreamingBuffer(rawDesired.StreamingBuffer, nil, opts...)
		rawDesired.EncryptionConfiguration = canonicalizeTableEncryptionConfiguration(rawDesired.EncryptionConfiguration, nil, opts...)
		rawDesired.SnapshotDefinition = canonicalizeTableSnapshotDefinition(rawDesired.SnapshotDefinition, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Table{}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.IsZeroValue(rawDesired.Dataset) || (dcl.IsEmptyValueIndirect(rawDesired.Dataset) && dcl.IsEmptyValueIndirect(rawInitial.Dataset)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Dataset = rawInitial.Dataset
	} else {
		canonicalDesired.Dataset = rawDesired.Dataset
	}
	if dcl.IsZeroValue(rawDesired.Project) || (dcl.IsEmptyValueIndirect(rawDesired.Project) && dcl.IsEmptyValueIndirect(rawInitial.Project)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.StringCanonicalize(rawDesired.FriendlyName, rawInitial.FriendlyName) {
		canonicalDesired.FriendlyName = rawInitial.FriendlyName
	} else {
		canonicalDesired.FriendlyName = rawDesired.FriendlyName
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.IsZeroValue(rawDesired.Labels) || (dcl.IsEmptyValueIndirect(rawDesired.Labels) && dcl.IsEmptyValueIndirect(rawInitial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	canonicalDesired.Schema = canonicalizeTableSchema(rawDesired.Schema, rawInitial.Schema, opts...)
	canonicalDesired.TimePartitioning = canonicalizeTableTimePartitioning(rawDesired.TimePartitioning, rawInitial.TimePartitioning, opts...)
	canonicalDesired.RangePartitioning = canonicalizeTableRangePartitioning(rawDesired.RangePartitioning, rawInitial.RangePartitioning, opts...)
	canonicalDesired.Clustering = canonicalizeTableClustering(rawDesired.Clustering, rawInitial.Clustering, opts...)
	if dcl.BoolCanonicalize(rawDesired.RequirePartitionFilter, rawInitial.RequirePartitionFilter) {
		canonicalDesired.RequirePartitionFilter = rawInitial.RequirePartitionFilter
	} else {
		canonicalDesired.RequirePartitionFilter = rawDesired.RequirePartitionFilter
	}
	if dcl.IsZeroValue(rawDesired.ExpirationTime) || (dcl.IsEmptyValueIndirect(rawDesired.ExpirationTime) && dcl.IsEmptyValueIndirect(rawInitial.ExpirationTime)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.ExpirationTime = rawInitial.ExpirationTime
	} else {
		canonicalDesired.ExpirationTime = rawDesired.ExpirationTime
	}
	canonicalDesired.View = canonicalizeTableView(rawDesired.View, rawInitial.View, opts...)
	canonicalDesired.MaterializedView = canonicalizeTableMaterializedView(rawDesired.MaterializedView, rawInitial.MaterializedView, opts...)
	canonicalDesired.ExternalDataConfiguration = canonicalizeTableExternalDataConfiguration(rawDesired.ExternalDataConfiguration, rawInitial.ExternalDataConfiguration, opts...)
	canonicalDesired.EncryptionConfiguration = canonicalizeTableEncryptionConfiguration(rawDesired.EncryptionConfiguration, rawInitial.EncryptionConfiguration, opts...)
	if dcl.StringCanonicalize(rawDesired.DefaultCollation, rawInitial.DefaultCollation) {
		canonicalDesired.DefaultCollation = rawInitial.DefaultCollation
	} else {
		canonicalDesired.DefaultCollation = rawDesired.DefaultCollation
	}

	return canonicalDesired, nil
}

func canonicalizeTableNewState(c *Client, rawNew, rawDesired *Table) (*Table, error) {

	if dcl.IsNotReturnedByServer(rawNew.Etag) && dcl.IsNotReturnedByServer(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Id) && dcl.IsNotReturnedByServer(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
		if dcl.StringCanonicalize(rawDesired.Id, rawNew.Id) {
			rawNew.Id = rawDesired.Id
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.SelfLink) && dcl.IsNotReturnedByServer(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Name) && dcl.IsNotReturnedByServer(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Dataset) && dcl.IsNotReturnedByServer(rawDesired.Dataset) {
		rawNew.Dataset = rawDesired.Dataset
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.Project) && dcl.IsNotReturnedByServer(rawDesired.Project) {
		rawNew.Project = rawDesired.Project
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.FriendlyName) && dcl.IsNotReturnedByServer(rawDesired.FriendlyName) {
		rawNew.FriendlyName = rawDesired.FriendlyName
	} else {
		if dcl.StringCanonicalize(rawDesired.FriendlyName, rawNew.FriendlyName) {
			rawNew.FriendlyName = rawDesired.FriendlyName
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Description) && dcl.IsNotReturnedByServer(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Labels) && dcl.IsNotReturnedByServer(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.Model) && dcl.IsNotReturnedByServer(rawDesired.Model) {
		rawNew.Model = rawDesired.Model
	} else {
		rawNew.Model = canonicalizeNewTableModel(c, rawDesired.Model, rawNew.Model)
	}

	if dcl.IsNotReturnedByServer(rawNew.Schema) && dcl.IsNotReturnedByServer(rawDesired.Schema) {
		rawNew.Schema = rawDesired.Schema
	} else {
		rawNew.Schema = canonicalizeNewTableSchema(c, rawDesired.Schema, rawNew.Schema)
	}

	if dcl.IsNotReturnedByServer(rawNew.TimePartitioning) && dcl.IsNotReturnedByServer(rawDesired.TimePartitioning) {
		rawNew.TimePartitioning = rawDesired.TimePartitioning
	} else {
		rawNew.TimePartitioning = canonicalizeNewTableTimePartitioning(c, rawDesired.TimePartitioning, rawNew.TimePartitioning)
	}

	if dcl.IsNotReturnedByServer(rawNew.RangePartitioning) && dcl.IsNotReturnedByServer(rawDesired.RangePartitioning) {
		rawNew.RangePartitioning = rawDesired.RangePartitioning
	} else {
		rawNew.RangePartitioning = canonicalizeNewTableRangePartitioning(c, rawDesired.RangePartitioning, rawNew.RangePartitioning)
	}

	if dcl.IsNotReturnedByServer(rawNew.Clustering) && dcl.IsNotReturnedByServer(rawDesired.Clustering) {
		rawNew.Clustering = rawDesired.Clustering
	} else {
		rawNew.Clustering = canonicalizeNewTableClustering(c, rawDesired.Clustering, rawNew.Clustering)
	}

	if dcl.IsNotReturnedByServer(rawNew.RequirePartitionFilter) && dcl.IsNotReturnedByServer(rawDesired.RequirePartitionFilter) {
		rawNew.RequirePartitionFilter = rawDesired.RequirePartitionFilter
	} else {
		if dcl.BoolCanonicalize(rawDesired.RequirePartitionFilter, rawNew.RequirePartitionFilter) {
			rawNew.RequirePartitionFilter = rawDesired.RequirePartitionFilter
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.NumBytes) && dcl.IsNotReturnedByServer(rawDesired.NumBytes) {
		rawNew.NumBytes = rawDesired.NumBytes
	} else {
		if dcl.StringCanonicalize(rawDesired.NumBytes, rawNew.NumBytes) {
			rawNew.NumBytes = rawDesired.NumBytes
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.NumPhysicalBytes) && dcl.IsNotReturnedByServer(rawDesired.NumPhysicalBytes) {
		rawNew.NumPhysicalBytes = rawDesired.NumPhysicalBytes
	} else {
		if dcl.StringCanonicalize(rawDesired.NumPhysicalBytes, rawNew.NumPhysicalBytes) {
			rawNew.NumPhysicalBytes = rawDesired.NumPhysicalBytes
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.NumLongTermBytes) && dcl.IsNotReturnedByServer(rawDesired.NumLongTermBytes) {
		rawNew.NumLongTermBytes = rawDesired.NumLongTermBytes
	} else {
		if dcl.StringCanonicalize(rawDesired.NumLongTermBytes, rawNew.NumLongTermBytes) {
			rawNew.NumLongTermBytes = rawDesired.NumLongTermBytes
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.NumRows) && dcl.IsNotReturnedByServer(rawDesired.NumRows) {
		rawNew.NumRows = rawDesired.NumRows
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.CreationTime) && dcl.IsNotReturnedByServer(rawDesired.CreationTime) {
		rawNew.CreationTime = rawDesired.CreationTime
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.ExpirationTime) && dcl.IsNotReturnedByServer(rawDesired.ExpirationTime) {
		rawNew.ExpirationTime = rawDesired.ExpirationTime
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.LastModifiedTime) && dcl.IsNotReturnedByServer(rawDesired.LastModifiedTime) {
		rawNew.LastModifiedTime = rawDesired.LastModifiedTime
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.Type) && dcl.IsNotReturnedByServer(rawDesired.Type) {
		rawNew.Type = rawDesired.Type
	} else {
		if dcl.StringCanonicalize(rawDesired.Type, rawNew.Type) {
			rawNew.Type = rawDesired.Type
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.View) && dcl.IsNotReturnedByServer(rawDesired.View) {
		rawNew.View = rawDesired.View
	} else {
		rawNew.View = canonicalizeNewTableView(c, rawDesired.View, rawNew.View)
	}

	if dcl.IsNotReturnedByServer(rawNew.MaterializedView) && dcl.IsNotReturnedByServer(rawDesired.MaterializedView) {
		rawNew.MaterializedView = rawDesired.MaterializedView
	} else {
		rawNew.MaterializedView = canonicalizeNewTableMaterializedView(c, rawDesired.MaterializedView, rawNew.MaterializedView)
	}

	if dcl.IsNotReturnedByServer(rawNew.ExternalDataConfiguration) && dcl.IsNotReturnedByServer(rawDesired.ExternalDataConfiguration) {
		rawNew.ExternalDataConfiguration = rawDesired.ExternalDataConfiguration
	} else {
		rawNew.ExternalDataConfiguration = canonicalizeNewTableExternalDataConfiguration(c, rawDesired.ExternalDataConfiguration, rawNew.ExternalDataConfiguration)
	}

	if dcl.IsNotReturnedByServer(rawNew.Location) && dcl.IsNotReturnedByServer(rawDesired.Location) {
		rawNew.Location = rawDesired.Location
	} else {
		if dcl.StringCanonicalize(rawDesired.Location, rawNew.Location) {
			rawNew.Location = rawDesired.Location
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.StreamingBuffer) && dcl.IsNotReturnedByServer(rawDesired.StreamingBuffer) {
		rawNew.StreamingBuffer = rawDesired.StreamingBuffer
	} else {
		rawNew.StreamingBuffer = canonicalizeNewTableStreamingBuffer(c, rawDesired.StreamingBuffer, rawNew.StreamingBuffer)
	}

	if dcl.IsNotReturnedByServer(rawNew.EncryptionConfiguration) && dcl.IsNotReturnedByServer(rawDesired.EncryptionConfiguration) {
		rawNew.EncryptionConfiguration = rawDesired.EncryptionConfiguration
	} else {
		rawNew.EncryptionConfiguration = canonicalizeNewTableEncryptionConfiguration(c, rawDesired.EncryptionConfiguration, rawNew.EncryptionConfiguration)
	}

	if dcl.IsNotReturnedByServer(rawNew.SnapshotDefinition) && dcl.IsNotReturnedByServer(rawDesired.SnapshotDefinition) {
		rawNew.SnapshotDefinition = rawDesired.SnapshotDefinition
	} else {
		rawNew.SnapshotDefinition = canonicalizeNewTableSnapshotDefinition(c, rawDesired.SnapshotDefinition, rawNew.SnapshotDefinition)
	}

	if dcl.IsNotReturnedByServer(rawNew.DefaultCollation) && dcl.IsNotReturnedByServer(rawDesired.DefaultCollation) {
		rawNew.DefaultCollation = rawDesired.DefaultCollation
	} else {
		if dcl.StringCanonicalize(rawDesired.DefaultCollation, rawNew.DefaultCollation) {
			rawNew.DefaultCollation = rawDesired.DefaultCollation
		}
	}

	return rawNew, nil
}

func canonicalizeTableModel(des, initial *TableModel, opts ...dcl.ApplyOption) *TableModel {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableModel{}

	return cDes
}

func canonicalizeTableModelSlice(des, initial []TableModel, opts ...dcl.ApplyOption) []TableModel {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableModel, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableModel(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableModel, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableModel(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableModel(c *Client, des, nw *TableModel) *TableModel {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableModel while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.ModelOptions = canonicalizeNewTableModelModelOptions(c, des.ModelOptions, nw.ModelOptions)
	nw.TrainingRuns = canonicalizeNewTableModelTrainingRunsSlice(c, des.TrainingRuns, nw.TrainingRuns)

	return nw
}

func canonicalizeNewTableModelSet(c *Client, des, nw []TableModel) []TableModel {
	if des == nil {
		return nw
	}
	var reorderedNew []TableModel
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableModelNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableModelSlice(c *Client, des, nw []TableModel) []TableModel {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableModel
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableModel(c, &d, &n))
	}

	return items
}

func canonicalizeTableModelModelOptions(des, initial *TableModelModelOptions, opts ...dcl.ApplyOption) *TableModelModelOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableModelModelOptions{}

	return cDes
}

func canonicalizeTableModelModelOptionsSlice(des, initial []TableModelModelOptions, opts ...dcl.ApplyOption) []TableModelModelOptions {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableModelModelOptions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableModelModelOptions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableModelModelOptions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableModelModelOptions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableModelModelOptions(c *Client, des, nw *TableModelModelOptions) *TableModelModelOptions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableModelModelOptions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.ModelType, nw.ModelType) {
		nw.ModelType = des.ModelType
	}
	if dcl.StringArrayCanonicalize(des.Labels, nw.Labels) {
		nw.Labels = des.Labels
	}
	if dcl.StringCanonicalize(des.LossType, nw.LossType) {
		nw.LossType = des.LossType
	}

	return nw
}

func canonicalizeNewTableModelModelOptionsSet(c *Client, des, nw []TableModelModelOptions) []TableModelModelOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []TableModelModelOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableModelModelOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableModelModelOptionsSlice(c *Client, des, nw []TableModelModelOptions) []TableModelModelOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableModelModelOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableModelModelOptions(c, &d, &n))
	}

	return items
}

func canonicalizeTableModelTrainingRuns(des, initial *TableModelTrainingRuns, opts ...dcl.ApplyOption) *TableModelTrainingRuns {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableModelTrainingRuns{}

	return cDes
}

func canonicalizeTableModelTrainingRunsSlice(des, initial []TableModelTrainingRuns, opts ...dcl.ApplyOption) []TableModelTrainingRuns {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableModelTrainingRuns, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableModelTrainingRuns(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableModelTrainingRuns, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableModelTrainingRuns(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableModelTrainingRuns(c *Client, des, nw *TableModelTrainingRuns) *TableModelTrainingRuns {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableModelTrainingRuns while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.State, nw.State) {
		nw.State = des.State
	}
	nw.TrainingOptions = canonicalizeNewTableModelTrainingRunsTrainingOptions(c, des.TrainingOptions, nw.TrainingOptions)
	nw.IterationResults = canonicalizeNewTableModelTrainingRunsIterationResultsSlice(c, des.IterationResults, nw.IterationResults)

	return nw
}

func canonicalizeNewTableModelTrainingRunsSet(c *Client, des, nw []TableModelTrainingRuns) []TableModelTrainingRuns {
	if des == nil {
		return nw
	}
	var reorderedNew []TableModelTrainingRuns
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableModelTrainingRunsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableModelTrainingRunsSlice(c *Client, des, nw []TableModelTrainingRuns) []TableModelTrainingRuns {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableModelTrainingRuns
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableModelTrainingRuns(c, &d, &n))
	}

	return items
}

func canonicalizeTableModelTrainingRunsTrainingOptions(des, initial *TableModelTrainingRunsTrainingOptions, opts ...dcl.ApplyOption) *TableModelTrainingRunsTrainingOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableModelTrainingRunsTrainingOptions{}

	return cDes
}

func canonicalizeTableModelTrainingRunsTrainingOptionsSlice(des, initial []TableModelTrainingRunsTrainingOptions, opts ...dcl.ApplyOption) []TableModelTrainingRunsTrainingOptions {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableModelTrainingRunsTrainingOptions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableModelTrainingRunsTrainingOptions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableModelTrainingRunsTrainingOptions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableModelTrainingRunsTrainingOptions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableModelTrainingRunsTrainingOptions(c *Client, des, nw *TableModelTrainingRunsTrainingOptions) *TableModelTrainingRunsTrainingOptions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableModelTrainingRunsTrainingOptions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.WarmStart, nw.WarmStart) {
		nw.WarmStart = des.WarmStart
	}
	if dcl.BoolCanonicalize(des.EarlyStop, nw.EarlyStop) {
		nw.EarlyStop = des.EarlyStop
	}
	if dcl.StringCanonicalize(des.LearnRateStrategy, nw.LearnRateStrategy) {
		nw.LearnRateStrategy = des.LearnRateStrategy
	}

	return nw
}

func canonicalizeNewTableModelTrainingRunsTrainingOptionsSet(c *Client, des, nw []TableModelTrainingRunsTrainingOptions) []TableModelTrainingRunsTrainingOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []TableModelTrainingRunsTrainingOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableModelTrainingRunsTrainingOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableModelTrainingRunsTrainingOptionsSlice(c *Client, des, nw []TableModelTrainingRunsTrainingOptions) []TableModelTrainingRunsTrainingOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableModelTrainingRunsTrainingOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableModelTrainingRunsTrainingOptions(c, &d, &n))
	}

	return items
}

func canonicalizeTableModelTrainingRunsIterationResults(des, initial *TableModelTrainingRunsIterationResults, opts ...dcl.ApplyOption) *TableModelTrainingRunsIterationResults {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableModelTrainingRunsIterationResults{}

	return cDes
}

func canonicalizeTableModelTrainingRunsIterationResultsSlice(des, initial []TableModelTrainingRunsIterationResults, opts ...dcl.ApplyOption) []TableModelTrainingRunsIterationResults {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableModelTrainingRunsIterationResults, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableModelTrainingRunsIterationResults(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableModelTrainingRunsIterationResults, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableModelTrainingRunsIterationResults(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableModelTrainingRunsIterationResults(c *Client, des, nw *TableModelTrainingRunsIterationResults) *TableModelTrainingRunsIterationResults {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableModelTrainingRunsIterationResults while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.DurationMs, nw.DurationMs) {
		nw.DurationMs = des.DurationMs
	}

	return nw
}

func canonicalizeNewTableModelTrainingRunsIterationResultsSet(c *Client, des, nw []TableModelTrainingRunsIterationResults) []TableModelTrainingRunsIterationResults {
	if des == nil {
		return nw
	}
	var reorderedNew []TableModelTrainingRunsIterationResults
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableModelTrainingRunsIterationResultsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableModelTrainingRunsIterationResultsSlice(c *Client, des, nw []TableModelTrainingRunsIterationResults) []TableModelTrainingRunsIterationResults {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableModelTrainingRunsIterationResults
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableModelTrainingRunsIterationResults(c, &d, &n))
	}

	return items
}

func canonicalizeTableSchema(des, initial *TableSchema, opts ...dcl.ApplyOption) *TableSchema {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableSchema{}

	cDes.Fields = canonicalizeTableGooglecloudbigqueryv2TablefieldschemaSlice(des.Fields, initial.Fields, opts...)

	return cDes
}

func canonicalizeTableSchemaSlice(des, initial []TableSchema, opts ...dcl.ApplyOption) []TableSchema {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableSchema, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableSchema(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableSchema, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableSchema(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableSchema(c *Client, des, nw *TableSchema) *TableSchema {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableSchema while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Fields = canonicalizeNewTableGooglecloudbigqueryv2TablefieldschemaSlice(c, des.Fields, nw.Fields)

	return nw
}

func canonicalizeNewTableSchemaSet(c *Client, des, nw []TableSchema) []TableSchema {
	if des == nil {
		return nw
	}
	var reorderedNew []TableSchema
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableSchemaNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableSchemaSlice(c *Client, des, nw []TableSchema) []TableSchema {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableSchema
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableSchema(c, &d, &n))
	}

	return items
}

func canonicalizeTableGooglecloudbigqueryv2Tablefieldschema(des, initial *TableGooglecloudbigqueryv2Tablefieldschema, opts ...dcl.ApplyOption) *TableGooglecloudbigqueryv2Tablefieldschema {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableGooglecloudbigqueryv2Tablefieldschema{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Type, initial.Type) || dcl.IsZeroValue(des.Type) {
		cDes.Type = initial.Type
	} else {
		cDes.Type = des.Type
	}
	if dcl.StringCanonicalize(des.Mode, initial.Mode) || dcl.IsZeroValue(des.Mode) {
		cDes.Mode = initial.Mode
	} else {
		cDes.Mode = des.Mode
	}
	if dcl.IsZeroValue(des.Fields) || (dcl.IsEmptyValueIndirect(des.Fields) && dcl.IsEmptyValueIndirect(initial.Fields)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Fields = initial.Fields
	} else {
		cDes.Fields = des.Fields
	}
	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		cDes.Description = initial.Description
	} else {
		cDes.Description = des.Description
	}
	cDes.Categories = canonicalizeTableGooglecloudbigqueryv2TablefieldschemaCategories(des.Categories, initial.Categories, opts...)
	cDes.PolicyTags = canonicalizeTableGooglecloudbigqueryv2TablefieldschemaPolicyTags(des.PolicyTags, initial.PolicyTags, opts...)
	if dcl.StringArrayCanonicalize(des.NameAlternative, initial.NameAlternative) {
		cDes.NameAlternative = initial.NameAlternative
	} else {
		cDes.NameAlternative = des.NameAlternative
	}
	if dcl.IsZeroValue(des.MaxLength) || (dcl.IsEmptyValueIndirect(des.MaxLength) && dcl.IsEmptyValueIndirect(initial.MaxLength)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MaxLength = initial.MaxLength
	} else {
		cDes.MaxLength = des.MaxLength
	}
	if dcl.IsZeroValue(des.Precision) || (dcl.IsEmptyValueIndirect(des.Precision) && dcl.IsEmptyValueIndirect(initial.Precision)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Precision = initial.Precision
	} else {
		cDes.Precision = des.Precision
	}
	if dcl.IsZeroValue(des.Scale) || (dcl.IsEmptyValueIndirect(des.Scale) && dcl.IsEmptyValueIndirect(initial.Scale)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Scale = initial.Scale
	} else {
		cDes.Scale = des.Scale
	}
	if dcl.StringCanonicalize(des.Collation, initial.Collation) || dcl.IsZeroValue(des.Collation) {
		cDes.Collation = initial.Collation
	} else {
		cDes.Collation = des.Collation
	}
	if dcl.StringCanonicalize(des.DefaultValueExpression, initial.DefaultValueExpression) || dcl.IsZeroValue(des.DefaultValueExpression) {
		cDes.DefaultValueExpression = initial.DefaultValueExpression
	} else {
		cDes.DefaultValueExpression = des.DefaultValueExpression
	}

	return cDes
}

func canonicalizeTableGooglecloudbigqueryv2TablefieldschemaSlice(des, initial []TableGooglecloudbigqueryv2Tablefieldschema, opts ...dcl.ApplyOption) []TableGooglecloudbigqueryv2Tablefieldschema {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableGooglecloudbigqueryv2Tablefieldschema, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableGooglecloudbigqueryv2Tablefieldschema(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableGooglecloudbigqueryv2Tablefieldschema, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableGooglecloudbigqueryv2Tablefieldschema(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableGooglecloudbigqueryv2Tablefieldschema(c *Client, des, nw *TableGooglecloudbigqueryv2Tablefieldschema) *TableGooglecloudbigqueryv2Tablefieldschema {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableGooglecloudbigqueryv2Tablefieldschema while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Type, nw.Type) {
		nw.Type = des.Type
	}
	if dcl.StringCanonicalize(des.Mode, nw.Mode) {
		nw.Mode = des.Mode
	}
	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}
	nw.Categories = canonicalizeNewTableGooglecloudbigqueryv2TablefieldschemaCategories(c, des.Categories, nw.Categories)
	nw.PolicyTags = canonicalizeNewTableGooglecloudbigqueryv2TablefieldschemaPolicyTags(c, des.PolicyTags, nw.PolicyTags)
	if dcl.StringArrayCanonicalize(des.NameAlternative, nw.NameAlternative) {
		nw.NameAlternative = des.NameAlternative
	}
	if dcl.StringCanonicalize(des.Collation, nw.Collation) {
		nw.Collation = des.Collation
	}
	if dcl.StringCanonicalize(des.DefaultValueExpression, nw.DefaultValueExpression) {
		nw.DefaultValueExpression = des.DefaultValueExpression
	}

	return nw
}

func canonicalizeNewTableGooglecloudbigqueryv2TablefieldschemaSet(c *Client, des, nw []TableGooglecloudbigqueryv2Tablefieldschema) []TableGooglecloudbigqueryv2Tablefieldschema {
	if des == nil {
		return nw
	}
	var reorderedNew []TableGooglecloudbigqueryv2Tablefieldschema
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableGooglecloudbigqueryv2TablefieldschemaNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableGooglecloudbigqueryv2TablefieldschemaSlice(c *Client, des, nw []TableGooglecloudbigqueryv2Tablefieldschema) []TableGooglecloudbigqueryv2Tablefieldschema {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableGooglecloudbigqueryv2Tablefieldschema
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableGooglecloudbigqueryv2Tablefieldschema(c, &d, &n))
	}

	return items
}

func canonicalizeTableGooglecloudbigqueryv2TablefieldschemaCategories(des, initial *TableGooglecloudbigqueryv2TablefieldschemaCategories, opts ...dcl.ApplyOption) *TableGooglecloudbigqueryv2TablefieldschemaCategories {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableGooglecloudbigqueryv2TablefieldschemaCategories{}

	if dcl.StringArrayCanonicalize(des.Names, initial.Names) {
		cDes.Names = initial.Names
	} else {
		cDes.Names = des.Names
	}

	return cDes
}

func canonicalizeTableGooglecloudbigqueryv2TablefieldschemaCategoriesSlice(des, initial []TableGooglecloudbigqueryv2TablefieldschemaCategories, opts ...dcl.ApplyOption) []TableGooglecloudbigqueryv2TablefieldschemaCategories {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableGooglecloudbigqueryv2TablefieldschemaCategories, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableGooglecloudbigqueryv2TablefieldschemaCategories(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableGooglecloudbigqueryv2TablefieldschemaCategories, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableGooglecloudbigqueryv2TablefieldschemaCategories(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableGooglecloudbigqueryv2TablefieldschemaCategories(c *Client, des, nw *TableGooglecloudbigqueryv2TablefieldschemaCategories) *TableGooglecloudbigqueryv2TablefieldschemaCategories {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableGooglecloudbigqueryv2TablefieldschemaCategories while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Names, nw.Names) {
		nw.Names = des.Names
	}

	return nw
}

func canonicalizeNewTableGooglecloudbigqueryv2TablefieldschemaCategoriesSet(c *Client, des, nw []TableGooglecloudbigqueryv2TablefieldschemaCategories) []TableGooglecloudbigqueryv2TablefieldschemaCategories {
	if des == nil {
		return nw
	}
	var reorderedNew []TableGooglecloudbigqueryv2TablefieldschemaCategories
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableGooglecloudbigqueryv2TablefieldschemaCategoriesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableGooglecloudbigqueryv2TablefieldschemaCategoriesSlice(c *Client, des, nw []TableGooglecloudbigqueryv2TablefieldschemaCategories) []TableGooglecloudbigqueryv2TablefieldschemaCategories {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableGooglecloudbigqueryv2TablefieldschemaCategories
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableGooglecloudbigqueryv2TablefieldschemaCategories(c, &d, &n))
	}

	return items
}

func canonicalizeTableGooglecloudbigqueryv2TablefieldschemaPolicyTags(des, initial *TableGooglecloudbigqueryv2TablefieldschemaPolicyTags, opts ...dcl.ApplyOption) *TableGooglecloudbigqueryv2TablefieldschemaPolicyTags {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableGooglecloudbigqueryv2TablefieldschemaPolicyTags{}

	if dcl.StringArrayCanonicalize(des.Names, initial.Names) {
		cDes.Names = initial.Names
	} else {
		cDes.Names = des.Names
	}

	return cDes
}

func canonicalizeTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsSlice(des, initial []TableGooglecloudbigqueryv2TablefieldschemaPolicyTags, opts ...dcl.ApplyOption) []TableGooglecloudbigqueryv2TablefieldschemaPolicyTags {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableGooglecloudbigqueryv2TablefieldschemaPolicyTags, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableGooglecloudbigqueryv2TablefieldschemaPolicyTags(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableGooglecloudbigqueryv2TablefieldschemaPolicyTags, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableGooglecloudbigqueryv2TablefieldschemaPolicyTags(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableGooglecloudbigqueryv2TablefieldschemaPolicyTags(c *Client, des, nw *TableGooglecloudbigqueryv2TablefieldschemaPolicyTags) *TableGooglecloudbigqueryv2TablefieldschemaPolicyTags {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableGooglecloudbigqueryv2TablefieldschemaPolicyTags while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Names, nw.Names) {
		nw.Names = des.Names
	}

	return nw
}

func canonicalizeNewTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsSet(c *Client, des, nw []TableGooglecloudbigqueryv2TablefieldschemaPolicyTags) []TableGooglecloudbigqueryv2TablefieldschemaPolicyTags {
	if des == nil {
		return nw
	}
	var reorderedNew []TableGooglecloudbigqueryv2TablefieldschemaPolicyTags
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsSlice(c *Client, des, nw []TableGooglecloudbigqueryv2TablefieldschemaPolicyTags) []TableGooglecloudbigqueryv2TablefieldschemaPolicyTags {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableGooglecloudbigqueryv2TablefieldschemaPolicyTags
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableGooglecloudbigqueryv2TablefieldschemaPolicyTags(c, &d, &n))
	}

	return items
}

func canonicalizeTableTimePartitioning(des, initial *TableTimePartitioning, opts ...dcl.ApplyOption) *TableTimePartitioning {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableTimePartitioning{}

	if dcl.StringCanonicalize(des.Type, initial.Type) || dcl.IsZeroValue(des.Type) {
		cDes.Type = initial.Type
	} else {
		cDes.Type = des.Type
	}
	if dcl.StringCanonicalize(des.ExpirationMs, initial.ExpirationMs) || dcl.IsZeroValue(des.ExpirationMs) {
		cDes.ExpirationMs = initial.ExpirationMs
	} else {
		cDes.ExpirationMs = des.ExpirationMs
	}
	if dcl.StringCanonicalize(des.Field, initial.Field) || dcl.IsZeroValue(des.Field) {
		cDes.Field = initial.Field
	} else {
		cDes.Field = des.Field
	}

	return cDes
}

func canonicalizeTableTimePartitioningSlice(des, initial []TableTimePartitioning, opts ...dcl.ApplyOption) []TableTimePartitioning {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableTimePartitioning, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableTimePartitioning(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableTimePartitioning, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableTimePartitioning(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableTimePartitioning(c *Client, des, nw *TableTimePartitioning) *TableTimePartitioning {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableTimePartitioning while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Type, nw.Type) {
		nw.Type = des.Type
	}
	if dcl.StringCanonicalize(des.ExpirationMs, nw.ExpirationMs) {
		nw.ExpirationMs = des.ExpirationMs
	}
	if dcl.StringCanonicalize(des.Field, nw.Field) {
		nw.Field = des.Field
	}

	return nw
}

func canonicalizeNewTableTimePartitioningSet(c *Client, des, nw []TableTimePartitioning) []TableTimePartitioning {
	if des == nil {
		return nw
	}
	var reorderedNew []TableTimePartitioning
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableTimePartitioningNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableTimePartitioningSlice(c *Client, des, nw []TableTimePartitioning) []TableTimePartitioning {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableTimePartitioning
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableTimePartitioning(c, &d, &n))
	}

	return items
}

func canonicalizeTableRangePartitioning(des, initial *TableRangePartitioning, opts ...dcl.ApplyOption) *TableRangePartitioning {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableRangePartitioning{}

	if dcl.StringCanonicalize(des.Field, initial.Field) || dcl.IsZeroValue(des.Field) {
		cDes.Field = initial.Field
	} else {
		cDes.Field = des.Field
	}
	cDes.Range = canonicalizeTableRangePartitioningRange(des.Range, initial.Range, opts...)

	return cDes
}

func canonicalizeTableRangePartitioningSlice(des, initial []TableRangePartitioning, opts ...dcl.ApplyOption) []TableRangePartitioning {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableRangePartitioning, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableRangePartitioning(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableRangePartitioning, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableRangePartitioning(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableRangePartitioning(c *Client, des, nw *TableRangePartitioning) *TableRangePartitioning {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableRangePartitioning while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Field, nw.Field) {
		nw.Field = des.Field
	}
	nw.Range = canonicalizeNewTableRangePartitioningRange(c, des.Range, nw.Range)

	return nw
}

func canonicalizeNewTableRangePartitioningSet(c *Client, des, nw []TableRangePartitioning) []TableRangePartitioning {
	if des == nil {
		return nw
	}
	var reorderedNew []TableRangePartitioning
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableRangePartitioningNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableRangePartitioningSlice(c *Client, des, nw []TableRangePartitioning) []TableRangePartitioning {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableRangePartitioning
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableRangePartitioning(c, &d, &n))
	}

	return items
}

func canonicalizeTableRangePartitioningRange(des, initial *TableRangePartitioningRange, opts ...dcl.ApplyOption) *TableRangePartitioningRange {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableRangePartitioningRange{}

	if dcl.StringCanonicalize(des.Start, initial.Start) || dcl.IsZeroValue(des.Start) {
		cDes.Start = initial.Start
	} else {
		cDes.Start = des.Start
	}
	if dcl.StringCanonicalize(des.End, initial.End) || dcl.IsZeroValue(des.End) {
		cDes.End = initial.End
	} else {
		cDes.End = des.End
	}
	if dcl.StringCanonicalize(des.Interval, initial.Interval) || dcl.IsZeroValue(des.Interval) {
		cDes.Interval = initial.Interval
	} else {
		cDes.Interval = des.Interval
	}

	return cDes
}

func canonicalizeTableRangePartitioningRangeSlice(des, initial []TableRangePartitioningRange, opts ...dcl.ApplyOption) []TableRangePartitioningRange {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableRangePartitioningRange, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableRangePartitioningRange(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableRangePartitioningRange, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableRangePartitioningRange(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableRangePartitioningRange(c *Client, des, nw *TableRangePartitioningRange) *TableRangePartitioningRange {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableRangePartitioningRange while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Start, nw.Start) {
		nw.Start = des.Start
	}
	if dcl.StringCanonicalize(des.End, nw.End) {
		nw.End = des.End
	}
	if dcl.StringCanonicalize(des.Interval, nw.Interval) {
		nw.Interval = des.Interval
	}

	return nw
}

func canonicalizeNewTableRangePartitioningRangeSet(c *Client, des, nw []TableRangePartitioningRange) []TableRangePartitioningRange {
	if des == nil {
		return nw
	}
	var reorderedNew []TableRangePartitioningRange
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableRangePartitioningRangeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableRangePartitioningRangeSlice(c *Client, des, nw []TableRangePartitioningRange) []TableRangePartitioningRange {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableRangePartitioningRange
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableRangePartitioningRange(c, &d, &n))
	}

	return items
}

func canonicalizeTableClustering(des, initial *TableClustering, opts ...dcl.ApplyOption) *TableClustering {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableClustering{}

	if dcl.StringArrayCanonicalize(des.Fields, initial.Fields) {
		cDes.Fields = initial.Fields
	} else {
		cDes.Fields = des.Fields
	}

	return cDes
}

func canonicalizeTableClusteringSlice(des, initial []TableClustering, opts ...dcl.ApplyOption) []TableClustering {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableClustering, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableClustering(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableClustering, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableClustering(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableClustering(c *Client, des, nw *TableClustering) *TableClustering {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableClustering while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Fields, nw.Fields) {
		nw.Fields = des.Fields
	}

	return nw
}

func canonicalizeNewTableClusteringSet(c *Client, des, nw []TableClustering) []TableClustering {
	if des == nil {
		return nw
	}
	var reorderedNew []TableClustering
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableClusteringNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableClusteringSlice(c *Client, des, nw []TableClustering) []TableClustering {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableClustering
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableClustering(c, &d, &n))
	}

	return items
}

func canonicalizeTableView(des, initial *TableView, opts ...dcl.ApplyOption) *TableView {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableView{}

	if dcl.StringCanonicalize(des.Query, initial.Query) || dcl.IsZeroValue(des.Query) {
		cDes.Query = initial.Query
	} else {
		cDes.Query = des.Query
	}
	cDes.UserDefinedFunctionResources = canonicalizeTableViewUserDefinedFunctionResourcesSlice(des.UserDefinedFunctionResources, initial.UserDefinedFunctionResources, opts...)
	if dcl.BoolCanonicalize(des.UseLegacySql, initial.UseLegacySql) || dcl.IsZeroValue(des.UseLegacySql) {
		cDes.UseLegacySql = initial.UseLegacySql
	} else {
		cDes.UseLegacySql = des.UseLegacySql
	}
	if dcl.BoolCanonicalize(des.UseExplicitColumnNames, initial.UseExplicitColumnNames) || dcl.IsZeroValue(des.UseExplicitColumnNames) {
		cDes.UseExplicitColumnNames = initial.UseExplicitColumnNames
	} else {
		cDes.UseExplicitColumnNames = des.UseExplicitColumnNames
	}

	return cDes
}

func canonicalizeTableViewSlice(des, initial []TableView, opts ...dcl.ApplyOption) []TableView {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableView, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableView(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableView, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableView(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableView(c *Client, des, nw *TableView) *TableView {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableView while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Query, nw.Query) {
		nw.Query = des.Query
	}
	nw.UserDefinedFunctionResources = canonicalizeNewTableViewUserDefinedFunctionResourcesSlice(c, des.UserDefinedFunctionResources, nw.UserDefinedFunctionResources)
	if dcl.BoolCanonicalize(des.UseLegacySql, nw.UseLegacySql) {
		nw.UseLegacySql = des.UseLegacySql
	}
	if dcl.BoolCanonicalize(des.UseExplicitColumnNames, nw.UseExplicitColumnNames) {
		nw.UseExplicitColumnNames = des.UseExplicitColumnNames
	}

	return nw
}

func canonicalizeNewTableViewSet(c *Client, des, nw []TableView) []TableView {
	if des == nil {
		return nw
	}
	var reorderedNew []TableView
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableViewNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableViewSlice(c *Client, des, nw []TableView) []TableView {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableView
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableView(c, &d, &n))
	}

	return items
}

func canonicalizeTableViewUserDefinedFunctionResources(des, initial *TableViewUserDefinedFunctionResources, opts ...dcl.ApplyOption) *TableViewUserDefinedFunctionResources {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableViewUserDefinedFunctionResources{}

	if dcl.StringCanonicalize(des.ResourceUri, initial.ResourceUri) || dcl.IsZeroValue(des.ResourceUri) {
		cDes.ResourceUri = initial.ResourceUri
	} else {
		cDes.ResourceUri = des.ResourceUri
	}
	if dcl.StringCanonicalize(des.InlineCode, initial.InlineCode) || dcl.IsZeroValue(des.InlineCode) {
		cDes.InlineCode = initial.InlineCode
	} else {
		cDes.InlineCode = des.InlineCode
	}
	if dcl.StringArrayCanonicalize(des.InlineCodeAlternative, initial.InlineCodeAlternative) {
		cDes.InlineCodeAlternative = initial.InlineCodeAlternative
	} else {
		cDes.InlineCodeAlternative = des.InlineCodeAlternative
	}

	return cDes
}

func canonicalizeTableViewUserDefinedFunctionResourcesSlice(des, initial []TableViewUserDefinedFunctionResources, opts ...dcl.ApplyOption) []TableViewUserDefinedFunctionResources {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableViewUserDefinedFunctionResources, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableViewUserDefinedFunctionResources(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableViewUserDefinedFunctionResources, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableViewUserDefinedFunctionResources(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableViewUserDefinedFunctionResources(c *Client, des, nw *TableViewUserDefinedFunctionResources) *TableViewUserDefinedFunctionResources {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableViewUserDefinedFunctionResources while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.ResourceUri, nw.ResourceUri) {
		nw.ResourceUri = des.ResourceUri
	}
	if dcl.StringCanonicalize(des.InlineCode, nw.InlineCode) {
		nw.InlineCode = des.InlineCode
	}
	if dcl.StringArrayCanonicalize(des.InlineCodeAlternative, nw.InlineCodeAlternative) {
		nw.InlineCodeAlternative = des.InlineCodeAlternative
	}

	return nw
}

func canonicalizeNewTableViewUserDefinedFunctionResourcesSet(c *Client, des, nw []TableViewUserDefinedFunctionResources) []TableViewUserDefinedFunctionResources {
	if des == nil {
		return nw
	}
	var reorderedNew []TableViewUserDefinedFunctionResources
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableViewUserDefinedFunctionResourcesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableViewUserDefinedFunctionResourcesSlice(c *Client, des, nw []TableViewUserDefinedFunctionResources) []TableViewUserDefinedFunctionResources {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableViewUserDefinedFunctionResources
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableViewUserDefinedFunctionResources(c, &d, &n))
	}

	return items
}

func canonicalizeTableMaterializedView(des, initial *TableMaterializedView, opts ...dcl.ApplyOption) *TableMaterializedView {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableMaterializedView{}

	if dcl.StringCanonicalize(des.Query, initial.Query) || dcl.IsZeroValue(des.Query) {
		cDes.Query = initial.Query
	} else {
		cDes.Query = des.Query
	}
	if dcl.BoolCanonicalize(des.EnableRefresh, initial.EnableRefresh) || dcl.IsZeroValue(des.EnableRefresh) {
		cDes.EnableRefresh = initial.EnableRefresh
	} else {
		cDes.EnableRefresh = des.EnableRefresh
	}
	if dcl.IsZeroValue(des.RefreshIntervalMs) || (dcl.IsEmptyValueIndirect(des.RefreshIntervalMs) && dcl.IsEmptyValueIndirect(initial.RefreshIntervalMs)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.RefreshIntervalMs = initial.RefreshIntervalMs
	} else {
		cDes.RefreshIntervalMs = des.RefreshIntervalMs
	}

	return cDes
}

func canonicalizeTableMaterializedViewSlice(des, initial []TableMaterializedView, opts ...dcl.ApplyOption) []TableMaterializedView {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableMaterializedView, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableMaterializedView(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableMaterializedView, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableMaterializedView(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableMaterializedView(c *Client, des, nw *TableMaterializedView) *TableMaterializedView {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableMaterializedView while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Query, nw.Query) {
		nw.Query = des.Query
	}
	if dcl.BoolCanonicalize(des.EnableRefresh, nw.EnableRefresh) {
		nw.EnableRefresh = des.EnableRefresh
	}

	return nw
}

func canonicalizeNewTableMaterializedViewSet(c *Client, des, nw []TableMaterializedView) []TableMaterializedView {
	if des == nil {
		return nw
	}
	var reorderedNew []TableMaterializedView
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableMaterializedViewNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableMaterializedViewSlice(c *Client, des, nw []TableMaterializedView) []TableMaterializedView {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableMaterializedView
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableMaterializedView(c, &d, &n))
	}

	return items
}

func canonicalizeTableExternalDataConfiguration(des, initial *TableExternalDataConfiguration, opts ...dcl.ApplyOption) *TableExternalDataConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableExternalDataConfiguration{}

	if dcl.StringArrayCanonicalize(des.SourceUris, initial.SourceUris) {
		cDes.SourceUris = initial.SourceUris
	} else {
		cDes.SourceUris = des.SourceUris
	}
	cDes.Schema = canonicalizeTableExternalDataConfigurationSchema(des.Schema, initial.Schema, opts...)
	if dcl.StringCanonicalize(des.SourceFormat, initial.SourceFormat) || dcl.IsZeroValue(des.SourceFormat) {
		cDes.SourceFormat = initial.SourceFormat
	} else {
		cDes.SourceFormat = des.SourceFormat
	}
	if dcl.IsZeroValue(des.MaxBadRecords) || (dcl.IsEmptyValueIndirect(des.MaxBadRecords) && dcl.IsEmptyValueIndirect(initial.MaxBadRecords)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MaxBadRecords = initial.MaxBadRecords
	} else {
		cDes.MaxBadRecords = des.MaxBadRecords
	}
	if dcl.BoolCanonicalize(des.Autodetect, initial.Autodetect) || dcl.IsZeroValue(des.Autodetect) {
		cDes.Autodetect = initial.Autodetect
	} else {
		cDes.Autodetect = des.Autodetect
	}
	if dcl.BoolCanonicalize(des.IgnoreUnknownValues, initial.IgnoreUnknownValues) || dcl.IsZeroValue(des.IgnoreUnknownValues) {
		cDes.IgnoreUnknownValues = initial.IgnoreUnknownValues
	} else {
		cDes.IgnoreUnknownValues = des.IgnoreUnknownValues
	}
	if dcl.StringCanonicalize(des.Compression, initial.Compression) || dcl.IsZeroValue(des.Compression) {
		cDes.Compression = initial.Compression
	} else {
		cDes.Compression = des.Compression
	}
	cDes.CsvOptions = canonicalizeTableExternalDataConfigurationCsvOptions(des.CsvOptions, initial.CsvOptions, opts...)
	cDes.BigtableOptions = canonicalizeTableExternalDataConfigurationBigtableOptions(des.BigtableOptions, initial.BigtableOptions, opts...)
	cDes.GoogleSheetsOptions = canonicalizeTableExternalDataConfigurationGoogleSheetsOptions(des.GoogleSheetsOptions, initial.GoogleSheetsOptions, opts...)
	if dcl.IsZeroValue(des.MaxBadRecordsAlternative) || (dcl.IsEmptyValueIndirect(des.MaxBadRecordsAlternative) && dcl.IsEmptyValueIndirect(initial.MaxBadRecordsAlternative)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MaxBadRecordsAlternative = initial.MaxBadRecordsAlternative
	} else {
		cDes.MaxBadRecordsAlternative = des.MaxBadRecordsAlternative
	}
	cDes.HivePartitioningOptions = canonicalizeTableExternalDataConfigurationHivePartitioningOptions(des.HivePartitioningOptions, initial.HivePartitioningOptions, opts...)
	if dcl.StringCanonicalize(des.ConnectionId, initial.ConnectionId) || dcl.IsZeroValue(des.ConnectionId) {
		cDes.ConnectionId = initial.ConnectionId
	} else {
		cDes.ConnectionId = des.ConnectionId
	}
	cDes.ValueConversionModes = canonicalizeTableExternalDataConfigurationValueConversionModes(des.ValueConversionModes, initial.ValueConversionModes, opts...)
	if dcl.IsZeroValue(des.DecimalTargetTypes) || (dcl.IsEmptyValueIndirect(des.DecimalTargetTypes) && dcl.IsEmptyValueIndirect(initial.DecimalTargetTypes)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.DecimalTargetTypes = initial.DecimalTargetTypes
	} else {
		cDes.DecimalTargetTypes = des.DecimalTargetTypes
	}
	cDes.AvroOptions = canonicalizeTableExternalDataConfigurationAvroOptions(des.AvroOptions, initial.AvroOptions, opts...)
	if dcl.IsZeroValue(des.JsonExtension) || (dcl.IsEmptyValueIndirect(des.JsonExtension) && dcl.IsEmptyValueIndirect(initial.JsonExtension)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.JsonExtension = initial.JsonExtension
	} else {
		cDes.JsonExtension = des.JsonExtension
	}
	cDes.ParquetOptions = canonicalizeTableExternalDataConfigurationParquetOptions(des.ParquetOptions, initial.ParquetOptions, opts...)

	return cDes
}

func canonicalizeTableExternalDataConfigurationSlice(des, initial []TableExternalDataConfiguration, opts ...dcl.ApplyOption) []TableExternalDataConfiguration {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableExternalDataConfiguration, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableExternalDataConfiguration(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableExternalDataConfiguration, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableExternalDataConfiguration(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableExternalDataConfiguration(c *Client, des, nw *TableExternalDataConfiguration) *TableExternalDataConfiguration {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableExternalDataConfiguration while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.SourceUris, nw.SourceUris) {
		nw.SourceUris = des.SourceUris
	}
	nw.Schema = canonicalizeNewTableExternalDataConfigurationSchema(c, des.Schema, nw.Schema)
	if dcl.StringCanonicalize(des.SourceFormat, nw.SourceFormat) {
		nw.SourceFormat = des.SourceFormat
	}
	if dcl.BoolCanonicalize(des.Autodetect, nw.Autodetect) {
		nw.Autodetect = des.Autodetect
	}
	if dcl.BoolCanonicalize(des.IgnoreUnknownValues, nw.IgnoreUnknownValues) {
		nw.IgnoreUnknownValues = des.IgnoreUnknownValues
	}
	if dcl.StringCanonicalize(des.Compression, nw.Compression) {
		nw.Compression = des.Compression
	}
	nw.CsvOptions = canonicalizeNewTableExternalDataConfigurationCsvOptions(c, des.CsvOptions, nw.CsvOptions)
	nw.BigtableOptions = canonicalizeNewTableExternalDataConfigurationBigtableOptions(c, des.BigtableOptions, nw.BigtableOptions)
	nw.GoogleSheetsOptions = canonicalizeNewTableExternalDataConfigurationGoogleSheetsOptions(c, des.GoogleSheetsOptions, nw.GoogleSheetsOptions)
	nw.HivePartitioningOptions = canonicalizeNewTableExternalDataConfigurationHivePartitioningOptions(c, des.HivePartitioningOptions, nw.HivePartitioningOptions)
	if dcl.StringCanonicalize(des.ConnectionId, nw.ConnectionId) {
		nw.ConnectionId = des.ConnectionId
	}
	nw.ValueConversionModes = canonicalizeNewTableExternalDataConfigurationValueConversionModes(c, des.ValueConversionModes, nw.ValueConversionModes)
	nw.AvroOptions = canonicalizeNewTableExternalDataConfigurationAvroOptions(c, des.AvroOptions, nw.AvroOptions)
	nw.ParquetOptions = canonicalizeNewTableExternalDataConfigurationParquetOptions(c, des.ParquetOptions, nw.ParquetOptions)

	return nw
}

func canonicalizeNewTableExternalDataConfigurationSet(c *Client, des, nw []TableExternalDataConfiguration) []TableExternalDataConfiguration {
	if des == nil {
		return nw
	}
	var reorderedNew []TableExternalDataConfiguration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableExternalDataConfigurationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableExternalDataConfigurationSlice(c *Client, des, nw []TableExternalDataConfiguration) []TableExternalDataConfiguration {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableExternalDataConfiguration
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableExternalDataConfiguration(c, &d, &n))
	}

	return items
}

func canonicalizeTableExternalDataConfigurationSchema(des, initial *TableExternalDataConfigurationSchema, opts ...dcl.ApplyOption) *TableExternalDataConfigurationSchema {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableExternalDataConfigurationSchema{}

	if dcl.IsZeroValue(des.Fields) || (dcl.IsEmptyValueIndirect(des.Fields) && dcl.IsEmptyValueIndirect(initial.Fields)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Fields = initial.Fields
	} else {
		cDes.Fields = des.Fields
	}

	return cDes
}

func canonicalizeTableExternalDataConfigurationSchemaSlice(des, initial []TableExternalDataConfigurationSchema, opts ...dcl.ApplyOption) []TableExternalDataConfigurationSchema {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableExternalDataConfigurationSchema, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableExternalDataConfigurationSchema(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableExternalDataConfigurationSchema, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableExternalDataConfigurationSchema(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableExternalDataConfigurationSchema(c *Client, des, nw *TableExternalDataConfigurationSchema) *TableExternalDataConfigurationSchema {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableExternalDataConfigurationSchema while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewTableExternalDataConfigurationSchemaSet(c *Client, des, nw []TableExternalDataConfigurationSchema) []TableExternalDataConfigurationSchema {
	if des == nil {
		return nw
	}
	var reorderedNew []TableExternalDataConfigurationSchema
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableExternalDataConfigurationSchemaNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableExternalDataConfigurationSchemaSlice(c *Client, des, nw []TableExternalDataConfigurationSchema) []TableExternalDataConfigurationSchema {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableExternalDataConfigurationSchema
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableExternalDataConfigurationSchema(c, &d, &n))
	}

	return items
}

func canonicalizeTableExternalDataConfigurationCsvOptions(des, initial *TableExternalDataConfigurationCsvOptions, opts ...dcl.ApplyOption) *TableExternalDataConfigurationCsvOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableExternalDataConfigurationCsvOptions{}

	if dcl.StringCanonicalize(des.FieldDelimiter, initial.FieldDelimiter) || dcl.IsZeroValue(des.FieldDelimiter) {
		cDes.FieldDelimiter = initial.FieldDelimiter
	} else {
		cDes.FieldDelimiter = des.FieldDelimiter
	}
	if dcl.StringCanonicalize(des.SkipLeadingRows, initial.SkipLeadingRows) || dcl.IsZeroValue(des.SkipLeadingRows) {
		cDes.SkipLeadingRows = initial.SkipLeadingRows
	} else {
		cDes.SkipLeadingRows = des.SkipLeadingRows
	}
	if dcl.StringCanonicalize(des.Quote, initial.Quote) || dcl.IsZeroValue(des.Quote) {
		cDes.Quote = initial.Quote
	} else {
		cDes.Quote = des.Quote
	}
	if dcl.BoolCanonicalize(des.AllowQuotedNewlines, initial.AllowQuotedNewlines) || dcl.IsZeroValue(des.AllowQuotedNewlines) {
		cDes.AllowQuotedNewlines = initial.AllowQuotedNewlines
	} else {
		cDes.AllowQuotedNewlines = des.AllowQuotedNewlines
	}
	if dcl.BoolCanonicalize(des.AllowJaggedRows, initial.AllowJaggedRows) || dcl.IsZeroValue(des.AllowJaggedRows) {
		cDes.AllowJaggedRows = initial.AllowJaggedRows
	} else {
		cDes.AllowJaggedRows = des.AllowJaggedRows
	}
	if dcl.StringCanonicalize(des.Encoding, initial.Encoding) || dcl.IsZeroValue(des.Encoding) {
		cDes.Encoding = initial.Encoding
	} else {
		cDes.Encoding = des.Encoding
	}

	return cDes
}

func canonicalizeTableExternalDataConfigurationCsvOptionsSlice(des, initial []TableExternalDataConfigurationCsvOptions, opts ...dcl.ApplyOption) []TableExternalDataConfigurationCsvOptions {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableExternalDataConfigurationCsvOptions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableExternalDataConfigurationCsvOptions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableExternalDataConfigurationCsvOptions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableExternalDataConfigurationCsvOptions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableExternalDataConfigurationCsvOptions(c *Client, des, nw *TableExternalDataConfigurationCsvOptions) *TableExternalDataConfigurationCsvOptions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableExternalDataConfigurationCsvOptions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.FieldDelimiter, nw.FieldDelimiter) {
		nw.FieldDelimiter = des.FieldDelimiter
	}
	if dcl.StringCanonicalize(des.SkipLeadingRows, nw.SkipLeadingRows) {
		nw.SkipLeadingRows = des.SkipLeadingRows
	}
	if dcl.StringCanonicalize(des.Quote, nw.Quote) {
		nw.Quote = des.Quote
	}
	if dcl.BoolCanonicalize(des.AllowQuotedNewlines, nw.AllowQuotedNewlines) {
		nw.AllowQuotedNewlines = des.AllowQuotedNewlines
	}
	if dcl.BoolCanonicalize(des.AllowJaggedRows, nw.AllowJaggedRows) {
		nw.AllowJaggedRows = des.AllowJaggedRows
	}
	if dcl.StringCanonicalize(des.Encoding, nw.Encoding) {
		nw.Encoding = des.Encoding
	}

	return nw
}

func canonicalizeNewTableExternalDataConfigurationCsvOptionsSet(c *Client, des, nw []TableExternalDataConfigurationCsvOptions) []TableExternalDataConfigurationCsvOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []TableExternalDataConfigurationCsvOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableExternalDataConfigurationCsvOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableExternalDataConfigurationCsvOptionsSlice(c *Client, des, nw []TableExternalDataConfigurationCsvOptions) []TableExternalDataConfigurationCsvOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableExternalDataConfigurationCsvOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableExternalDataConfigurationCsvOptions(c, &d, &n))
	}

	return items
}

func canonicalizeTableExternalDataConfigurationBigtableOptions(des, initial *TableExternalDataConfigurationBigtableOptions, opts ...dcl.ApplyOption) *TableExternalDataConfigurationBigtableOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableExternalDataConfigurationBigtableOptions{}

	cDes.ColumnFamilies = canonicalizeTableExternalDataConfigurationBigtableOptionsColumnFamiliesSlice(des.ColumnFamilies, initial.ColumnFamilies, opts...)
	if dcl.BoolCanonicalize(des.IgnoreUnspecifiedColumnFamilies, initial.IgnoreUnspecifiedColumnFamilies) || dcl.IsZeroValue(des.IgnoreUnspecifiedColumnFamilies) {
		cDes.IgnoreUnspecifiedColumnFamilies = initial.IgnoreUnspecifiedColumnFamilies
	} else {
		cDes.IgnoreUnspecifiedColumnFamilies = des.IgnoreUnspecifiedColumnFamilies
	}
	if dcl.BoolCanonicalize(des.ReadRowkeyAsString, initial.ReadRowkeyAsString) || dcl.IsZeroValue(des.ReadRowkeyAsString) {
		cDes.ReadRowkeyAsString = initial.ReadRowkeyAsString
	} else {
		cDes.ReadRowkeyAsString = des.ReadRowkeyAsString
	}

	return cDes
}

func canonicalizeTableExternalDataConfigurationBigtableOptionsSlice(des, initial []TableExternalDataConfigurationBigtableOptions, opts ...dcl.ApplyOption) []TableExternalDataConfigurationBigtableOptions {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableExternalDataConfigurationBigtableOptions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableExternalDataConfigurationBigtableOptions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableExternalDataConfigurationBigtableOptions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableExternalDataConfigurationBigtableOptions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableExternalDataConfigurationBigtableOptions(c *Client, des, nw *TableExternalDataConfigurationBigtableOptions) *TableExternalDataConfigurationBigtableOptions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableExternalDataConfigurationBigtableOptions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.ColumnFamilies = canonicalizeNewTableExternalDataConfigurationBigtableOptionsColumnFamiliesSlice(c, des.ColumnFamilies, nw.ColumnFamilies)
	if dcl.BoolCanonicalize(des.IgnoreUnspecifiedColumnFamilies, nw.IgnoreUnspecifiedColumnFamilies) {
		nw.IgnoreUnspecifiedColumnFamilies = des.IgnoreUnspecifiedColumnFamilies
	}
	if dcl.BoolCanonicalize(des.ReadRowkeyAsString, nw.ReadRowkeyAsString) {
		nw.ReadRowkeyAsString = des.ReadRowkeyAsString
	}

	return nw
}

func canonicalizeNewTableExternalDataConfigurationBigtableOptionsSet(c *Client, des, nw []TableExternalDataConfigurationBigtableOptions) []TableExternalDataConfigurationBigtableOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []TableExternalDataConfigurationBigtableOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableExternalDataConfigurationBigtableOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableExternalDataConfigurationBigtableOptionsSlice(c *Client, des, nw []TableExternalDataConfigurationBigtableOptions) []TableExternalDataConfigurationBigtableOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableExternalDataConfigurationBigtableOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableExternalDataConfigurationBigtableOptions(c, &d, &n))
	}

	return items
}

func canonicalizeTableExternalDataConfigurationBigtableOptionsColumnFamilies(des, initial *TableExternalDataConfigurationBigtableOptionsColumnFamilies, opts ...dcl.ApplyOption) *TableExternalDataConfigurationBigtableOptionsColumnFamilies {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableExternalDataConfigurationBigtableOptionsColumnFamilies{}

	if dcl.StringCanonicalize(des.FamilyId, initial.FamilyId) || dcl.IsZeroValue(des.FamilyId) {
		cDes.FamilyId = initial.FamilyId
	} else {
		cDes.FamilyId = des.FamilyId
	}
	if dcl.StringCanonicalize(des.Type, initial.Type) || dcl.IsZeroValue(des.Type) {
		cDes.Type = initial.Type
	} else {
		cDes.Type = des.Type
	}
	if dcl.StringCanonicalize(des.Encoding, initial.Encoding) || dcl.IsZeroValue(des.Encoding) {
		cDes.Encoding = initial.Encoding
	} else {
		cDes.Encoding = des.Encoding
	}
	cDes.Columns = canonicalizeTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsSlice(des.Columns, initial.Columns, opts...)
	if dcl.BoolCanonicalize(des.OnlyReadLatest, initial.OnlyReadLatest) || dcl.IsZeroValue(des.OnlyReadLatest) {
		cDes.OnlyReadLatest = initial.OnlyReadLatest
	} else {
		cDes.OnlyReadLatest = des.OnlyReadLatest
	}

	return cDes
}

func canonicalizeTableExternalDataConfigurationBigtableOptionsColumnFamiliesSlice(des, initial []TableExternalDataConfigurationBigtableOptionsColumnFamilies, opts ...dcl.ApplyOption) []TableExternalDataConfigurationBigtableOptionsColumnFamilies {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableExternalDataConfigurationBigtableOptionsColumnFamilies, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableExternalDataConfigurationBigtableOptionsColumnFamilies(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableExternalDataConfigurationBigtableOptionsColumnFamilies, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableExternalDataConfigurationBigtableOptionsColumnFamilies(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableExternalDataConfigurationBigtableOptionsColumnFamilies(c *Client, des, nw *TableExternalDataConfigurationBigtableOptionsColumnFamilies) *TableExternalDataConfigurationBigtableOptionsColumnFamilies {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableExternalDataConfigurationBigtableOptionsColumnFamilies while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.FamilyId, nw.FamilyId) {
		nw.FamilyId = des.FamilyId
	}
	if dcl.StringCanonicalize(des.Type, nw.Type) {
		nw.Type = des.Type
	}
	if dcl.StringCanonicalize(des.Encoding, nw.Encoding) {
		nw.Encoding = des.Encoding
	}
	nw.Columns = canonicalizeNewTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsSlice(c, des.Columns, nw.Columns)
	if dcl.BoolCanonicalize(des.OnlyReadLatest, nw.OnlyReadLatest) {
		nw.OnlyReadLatest = des.OnlyReadLatest
	}

	return nw
}

func canonicalizeNewTableExternalDataConfigurationBigtableOptionsColumnFamiliesSet(c *Client, des, nw []TableExternalDataConfigurationBigtableOptionsColumnFamilies) []TableExternalDataConfigurationBigtableOptionsColumnFamilies {
	if des == nil {
		return nw
	}
	var reorderedNew []TableExternalDataConfigurationBigtableOptionsColumnFamilies
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableExternalDataConfigurationBigtableOptionsColumnFamiliesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableExternalDataConfigurationBigtableOptionsColumnFamiliesSlice(c *Client, des, nw []TableExternalDataConfigurationBigtableOptionsColumnFamilies) []TableExternalDataConfigurationBigtableOptionsColumnFamilies {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableExternalDataConfigurationBigtableOptionsColumnFamilies
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableExternalDataConfigurationBigtableOptionsColumnFamilies(c, &d, &n))
	}

	return items
}

func canonicalizeTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns(des, initial *TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns, opts ...dcl.ApplyOption) *TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns{}

	if dcl.StringCanonicalize(des.QualifierEncoded, initial.QualifierEncoded) || dcl.IsZeroValue(des.QualifierEncoded) {
		cDes.QualifierEncoded = initial.QualifierEncoded
	} else {
		cDes.QualifierEncoded = des.QualifierEncoded
	}
	if dcl.StringCanonicalize(des.QualifierString, initial.QualifierString) || dcl.IsZeroValue(des.QualifierString) {
		cDes.QualifierString = initial.QualifierString
	} else {
		cDes.QualifierString = des.QualifierString
	}
	if dcl.StringCanonicalize(des.FieldName, initial.FieldName) || dcl.IsZeroValue(des.FieldName) {
		cDes.FieldName = initial.FieldName
	} else {
		cDes.FieldName = des.FieldName
	}
	if dcl.StringCanonicalize(des.Type, initial.Type) || dcl.IsZeroValue(des.Type) {
		cDes.Type = initial.Type
	} else {
		cDes.Type = des.Type
	}
	if dcl.StringCanonicalize(des.Encoding, initial.Encoding) || dcl.IsZeroValue(des.Encoding) {
		cDes.Encoding = initial.Encoding
	} else {
		cDes.Encoding = des.Encoding
	}
	if dcl.BoolCanonicalize(des.OnlyReadLatest, initial.OnlyReadLatest) || dcl.IsZeroValue(des.OnlyReadLatest) {
		cDes.OnlyReadLatest = initial.OnlyReadLatest
	} else {
		cDes.OnlyReadLatest = des.OnlyReadLatest
	}

	return cDes
}

func canonicalizeTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsSlice(des, initial []TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns, opts ...dcl.ApplyOption) []TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns(c *Client, des, nw *TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns) *TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.QualifierEncoded, nw.QualifierEncoded) {
		nw.QualifierEncoded = des.QualifierEncoded
	}
	if dcl.StringCanonicalize(des.QualifierString, nw.QualifierString) {
		nw.QualifierString = des.QualifierString
	}
	if dcl.StringCanonicalize(des.FieldName, nw.FieldName) {
		nw.FieldName = des.FieldName
	}
	if dcl.StringCanonicalize(des.Type, nw.Type) {
		nw.Type = des.Type
	}
	if dcl.StringCanonicalize(des.Encoding, nw.Encoding) {
		nw.Encoding = des.Encoding
	}
	if dcl.BoolCanonicalize(des.OnlyReadLatest, nw.OnlyReadLatest) {
		nw.OnlyReadLatest = des.OnlyReadLatest
	}

	return nw
}

func canonicalizeNewTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsSet(c *Client, des, nw []TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns) []TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns {
	if des == nil {
		return nw
	}
	var reorderedNew []TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsSlice(c *Client, des, nw []TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns) []TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns(c, &d, &n))
	}

	return items
}

func canonicalizeTableExternalDataConfigurationGoogleSheetsOptions(des, initial *TableExternalDataConfigurationGoogleSheetsOptions, opts ...dcl.ApplyOption) *TableExternalDataConfigurationGoogleSheetsOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableExternalDataConfigurationGoogleSheetsOptions{}

	if dcl.StringCanonicalize(des.SkipLeadingRows, initial.SkipLeadingRows) || dcl.IsZeroValue(des.SkipLeadingRows) {
		cDes.SkipLeadingRows = initial.SkipLeadingRows
	} else {
		cDes.SkipLeadingRows = des.SkipLeadingRows
	}
	if dcl.StringCanonicalize(des.Range, initial.Range) || dcl.IsZeroValue(des.Range) {
		cDes.Range = initial.Range
	} else {
		cDes.Range = des.Range
	}

	return cDes
}

func canonicalizeTableExternalDataConfigurationGoogleSheetsOptionsSlice(des, initial []TableExternalDataConfigurationGoogleSheetsOptions, opts ...dcl.ApplyOption) []TableExternalDataConfigurationGoogleSheetsOptions {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableExternalDataConfigurationGoogleSheetsOptions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableExternalDataConfigurationGoogleSheetsOptions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableExternalDataConfigurationGoogleSheetsOptions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableExternalDataConfigurationGoogleSheetsOptions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableExternalDataConfigurationGoogleSheetsOptions(c *Client, des, nw *TableExternalDataConfigurationGoogleSheetsOptions) *TableExternalDataConfigurationGoogleSheetsOptions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableExternalDataConfigurationGoogleSheetsOptions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.SkipLeadingRows, nw.SkipLeadingRows) {
		nw.SkipLeadingRows = des.SkipLeadingRows
	}
	if dcl.StringCanonicalize(des.Range, nw.Range) {
		nw.Range = des.Range
	}

	return nw
}

func canonicalizeNewTableExternalDataConfigurationGoogleSheetsOptionsSet(c *Client, des, nw []TableExternalDataConfigurationGoogleSheetsOptions) []TableExternalDataConfigurationGoogleSheetsOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []TableExternalDataConfigurationGoogleSheetsOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableExternalDataConfigurationGoogleSheetsOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableExternalDataConfigurationGoogleSheetsOptionsSlice(c *Client, des, nw []TableExternalDataConfigurationGoogleSheetsOptions) []TableExternalDataConfigurationGoogleSheetsOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableExternalDataConfigurationGoogleSheetsOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableExternalDataConfigurationGoogleSheetsOptions(c, &d, &n))
	}

	return items
}

func canonicalizeTableExternalDataConfigurationHivePartitioningOptions(des, initial *TableExternalDataConfigurationHivePartitioningOptions, opts ...dcl.ApplyOption) *TableExternalDataConfigurationHivePartitioningOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableExternalDataConfigurationHivePartitioningOptions{}

	if dcl.StringCanonicalize(des.Mode, initial.Mode) || dcl.IsZeroValue(des.Mode) {
		cDes.Mode = initial.Mode
	} else {
		cDes.Mode = des.Mode
	}
	if dcl.StringCanonicalize(des.SourceUriPrefix, initial.SourceUriPrefix) || dcl.IsZeroValue(des.SourceUriPrefix) {
		cDes.SourceUriPrefix = initial.SourceUriPrefix
	} else {
		cDes.SourceUriPrefix = des.SourceUriPrefix
	}
	if dcl.BoolCanonicalize(des.RequirePartitionFilter, initial.RequirePartitionFilter) || dcl.IsZeroValue(des.RequirePartitionFilter) {
		cDes.RequirePartitionFilter = initial.RequirePartitionFilter
	} else {
		cDes.RequirePartitionFilter = des.RequirePartitionFilter
	}

	return cDes
}

func canonicalizeTableExternalDataConfigurationHivePartitioningOptionsSlice(des, initial []TableExternalDataConfigurationHivePartitioningOptions, opts ...dcl.ApplyOption) []TableExternalDataConfigurationHivePartitioningOptions {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableExternalDataConfigurationHivePartitioningOptions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableExternalDataConfigurationHivePartitioningOptions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableExternalDataConfigurationHivePartitioningOptions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableExternalDataConfigurationHivePartitioningOptions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableExternalDataConfigurationHivePartitioningOptions(c *Client, des, nw *TableExternalDataConfigurationHivePartitioningOptions) *TableExternalDataConfigurationHivePartitioningOptions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableExternalDataConfigurationHivePartitioningOptions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Mode, nw.Mode) {
		nw.Mode = des.Mode
	}
	if dcl.StringCanonicalize(des.SourceUriPrefix, nw.SourceUriPrefix) {
		nw.SourceUriPrefix = des.SourceUriPrefix
	}
	if dcl.BoolCanonicalize(des.RequirePartitionFilter, nw.RequirePartitionFilter) {
		nw.RequirePartitionFilter = des.RequirePartitionFilter
	}
	if dcl.StringArrayCanonicalize(des.Fields, nw.Fields) {
		nw.Fields = des.Fields
	}

	return nw
}

func canonicalizeNewTableExternalDataConfigurationHivePartitioningOptionsSet(c *Client, des, nw []TableExternalDataConfigurationHivePartitioningOptions) []TableExternalDataConfigurationHivePartitioningOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []TableExternalDataConfigurationHivePartitioningOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableExternalDataConfigurationHivePartitioningOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableExternalDataConfigurationHivePartitioningOptionsSlice(c *Client, des, nw []TableExternalDataConfigurationHivePartitioningOptions) []TableExternalDataConfigurationHivePartitioningOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableExternalDataConfigurationHivePartitioningOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableExternalDataConfigurationHivePartitioningOptions(c, &d, &n))
	}

	return items
}

func canonicalizeTableExternalDataConfigurationValueConversionModes(des, initial *TableExternalDataConfigurationValueConversionModes, opts ...dcl.ApplyOption) *TableExternalDataConfigurationValueConversionModes {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableExternalDataConfigurationValueConversionModes{}

	if dcl.IsZeroValue(des.TemporalTypesOutOfRangeConversionMode) || (dcl.IsEmptyValueIndirect(des.TemporalTypesOutOfRangeConversionMode) && dcl.IsEmptyValueIndirect(initial.TemporalTypesOutOfRangeConversionMode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.TemporalTypesOutOfRangeConversionMode = initial.TemporalTypesOutOfRangeConversionMode
	} else {
		cDes.TemporalTypesOutOfRangeConversionMode = des.TemporalTypesOutOfRangeConversionMode
	}
	if dcl.IsZeroValue(des.NumericTypeOutOfRangeConversionMode) || (dcl.IsEmptyValueIndirect(des.NumericTypeOutOfRangeConversionMode) && dcl.IsEmptyValueIndirect(initial.NumericTypeOutOfRangeConversionMode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.NumericTypeOutOfRangeConversionMode = initial.NumericTypeOutOfRangeConversionMode
	} else {
		cDes.NumericTypeOutOfRangeConversionMode = des.NumericTypeOutOfRangeConversionMode
	}

	return cDes
}

func canonicalizeTableExternalDataConfigurationValueConversionModesSlice(des, initial []TableExternalDataConfigurationValueConversionModes, opts ...dcl.ApplyOption) []TableExternalDataConfigurationValueConversionModes {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableExternalDataConfigurationValueConversionModes, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableExternalDataConfigurationValueConversionModes(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableExternalDataConfigurationValueConversionModes, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableExternalDataConfigurationValueConversionModes(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableExternalDataConfigurationValueConversionModes(c *Client, des, nw *TableExternalDataConfigurationValueConversionModes) *TableExternalDataConfigurationValueConversionModes {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableExternalDataConfigurationValueConversionModes while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewTableExternalDataConfigurationValueConversionModesSet(c *Client, des, nw []TableExternalDataConfigurationValueConversionModes) []TableExternalDataConfigurationValueConversionModes {
	if des == nil {
		return nw
	}
	var reorderedNew []TableExternalDataConfigurationValueConversionModes
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableExternalDataConfigurationValueConversionModesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableExternalDataConfigurationValueConversionModesSlice(c *Client, des, nw []TableExternalDataConfigurationValueConversionModes) []TableExternalDataConfigurationValueConversionModes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableExternalDataConfigurationValueConversionModes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableExternalDataConfigurationValueConversionModes(c, &d, &n))
	}

	return items
}

func canonicalizeTableExternalDataConfigurationAvroOptions(des, initial *TableExternalDataConfigurationAvroOptions, opts ...dcl.ApplyOption) *TableExternalDataConfigurationAvroOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableExternalDataConfigurationAvroOptions{}

	if dcl.BoolCanonicalize(des.UseAvroLogicalTypes, initial.UseAvroLogicalTypes) || dcl.IsZeroValue(des.UseAvroLogicalTypes) {
		cDes.UseAvroLogicalTypes = initial.UseAvroLogicalTypes
	} else {
		cDes.UseAvroLogicalTypes = des.UseAvroLogicalTypes
	}

	return cDes
}

func canonicalizeTableExternalDataConfigurationAvroOptionsSlice(des, initial []TableExternalDataConfigurationAvroOptions, opts ...dcl.ApplyOption) []TableExternalDataConfigurationAvroOptions {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableExternalDataConfigurationAvroOptions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableExternalDataConfigurationAvroOptions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableExternalDataConfigurationAvroOptions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableExternalDataConfigurationAvroOptions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableExternalDataConfigurationAvroOptions(c *Client, des, nw *TableExternalDataConfigurationAvroOptions) *TableExternalDataConfigurationAvroOptions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableExternalDataConfigurationAvroOptions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.UseAvroLogicalTypes, nw.UseAvroLogicalTypes) {
		nw.UseAvroLogicalTypes = des.UseAvroLogicalTypes
	}

	return nw
}

func canonicalizeNewTableExternalDataConfigurationAvroOptionsSet(c *Client, des, nw []TableExternalDataConfigurationAvroOptions) []TableExternalDataConfigurationAvroOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []TableExternalDataConfigurationAvroOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableExternalDataConfigurationAvroOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableExternalDataConfigurationAvroOptionsSlice(c *Client, des, nw []TableExternalDataConfigurationAvroOptions) []TableExternalDataConfigurationAvroOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableExternalDataConfigurationAvroOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableExternalDataConfigurationAvroOptions(c, &d, &n))
	}

	return items
}

func canonicalizeTableExternalDataConfigurationParquetOptions(des, initial *TableExternalDataConfigurationParquetOptions, opts ...dcl.ApplyOption) *TableExternalDataConfigurationParquetOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableExternalDataConfigurationParquetOptions{}

	if dcl.BoolCanonicalize(des.EnumAsString, initial.EnumAsString) || dcl.IsZeroValue(des.EnumAsString) {
		cDes.EnumAsString = initial.EnumAsString
	} else {
		cDes.EnumAsString = des.EnumAsString
	}
	if dcl.BoolCanonicalize(des.EnableListInference, initial.EnableListInference) || dcl.IsZeroValue(des.EnableListInference) {
		cDes.EnableListInference = initial.EnableListInference
	} else {
		cDes.EnableListInference = des.EnableListInference
	}

	return cDes
}

func canonicalizeTableExternalDataConfigurationParquetOptionsSlice(des, initial []TableExternalDataConfigurationParquetOptions, opts ...dcl.ApplyOption) []TableExternalDataConfigurationParquetOptions {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableExternalDataConfigurationParquetOptions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableExternalDataConfigurationParquetOptions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableExternalDataConfigurationParquetOptions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableExternalDataConfigurationParquetOptions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableExternalDataConfigurationParquetOptions(c *Client, des, nw *TableExternalDataConfigurationParquetOptions) *TableExternalDataConfigurationParquetOptions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableExternalDataConfigurationParquetOptions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.EnumAsString, nw.EnumAsString) {
		nw.EnumAsString = des.EnumAsString
	}
	if dcl.BoolCanonicalize(des.EnableListInference, nw.EnableListInference) {
		nw.EnableListInference = des.EnableListInference
	}

	return nw
}

func canonicalizeNewTableExternalDataConfigurationParquetOptionsSet(c *Client, des, nw []TableExternalDataConfigurationParquetOptions) []TableExternalDataConfigurationParquetOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []TableExternalDataConfigurationParquetOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableExternalDataConfigurationParquetOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableExternalDataConfigurationParquetOptionsSlice(c *Client, des, nw []TableExternalDataConfigurationParquetOptions) []TableExternalDataConfigurationParquetOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableExternalDataConfigurationParquetOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableExternalDataConfigurationParquetOptions(c, &d, &n))
	}

	return items
}

func canonicalizeTableStreamingBuffer(des, initial *TableStreamingBuffer, opts ...dcl.ApplyOption) *TableStreamingBuffer {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableStreamingBuffer{}

	return cDes
}

func canonicalizeTableStreamingBufferSlice(des, initial []TableStreamingBuffer, opts ...dcl.ApplyOption) []TableStreamingBuffer {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableStreamingBuffer, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableStreamingBuffer(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableStreamingBuffer, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableStreamingBuffer(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableStreamingBuffer(c *Client, des, nw *TableStreamingBuffer) *TableStreamingBuffer {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableStreamingBuffer while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewTableStreamingBufferSet(c *Client, des, nw []TableStreamingBuffer) []TableStreamingBuffer {
	if des == nil {
		return nw
	}
	var reorderedNew []TableStreamingBuffer
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableStreamingBufferNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableStreamingBufferSlice(c *Client, des, nw []TableStreamingBuffer) []TableStreamingBuffer {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableStreamingBuffer
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableStreamingBuffer(c, &d, &n))
	}

	return items
}

func canonicalizeTableEncryptionConfiguration(des, initial *TableEncryptionConfiguration, opts ...dcl.ApplyOption) *TableEncryptionConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableEncryptionConfiguration{}

	if dcl.StringCanonicalize(des.KmsKeyName, initial.KmsKeyName) || dcl.IsZeroValue(des.KmsKeyName) {
		cDes.KmsKeyName = initial.KmsKeyName
	} else {
		cDes.KmsKeyName = des.KmsKeyName
	}

	return cDes
}

func canonicalizeTableEncryptionConfigurationSlice(des, initial []TableEncryptionConfiguration, opts ...dcl.ApplyOption) []TableEncryptionConfiguration {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableEncryptionConfiguration, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableEncryptionConfiguration(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableEncryptionConfiguration, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableEncryptionConfiguration(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableEncryptionConfiguration(c *Client, des, nw *TableEncryptionConfiguration) *TableEncryptionConfiguration {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableEncryptionConfiguration while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.KmsKeyName, nw.KmsKeyName) {
		nw.KmsKeyName = des.KmsKeyName
	}

	return nw
}

func canonicalizeNewTableEncryptionConfigurationSet(c *Client, des, nw []TableEncryptionConfiguration) []TableEncryptionConfiguration {
	if des == nil {
		return nw
	}
	var reorderedNew []TableEncryptionConfiguration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableEncryptionConfigurationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableEncryptionConfigurationSlice(c *Client, des, nw []TableEncryptionConfiguration) []TableEncryptionConfiguration {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableEncryptionConfiguration
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableEncryptionConfiguration(c, &d, &n))
	}

	return items
}

func canonicalizeTableSnapshotDefinition(des, initial *TableSnapshotDefinition, opts ...dcl.ApplyOption) *TableSnapshotDefinition {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TableSnapshotDefinition{}

	if dcl.IsZeroValue(des.Table) || (dcl.IsEmptyValueIndirect(des.Table) && dcl.IsEmptyValueIndirect(initial.Table)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Table = initial.Table
	} else {
		cDes.Table = des.Table
	}
	if dcl.IsZeroValue(des.Dataset) || (dcl.IsEmptyValueIndirect(des.Dataset) && dcl.IsEmptyValueIndirect(initial.Dataset)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Dataset = initial.Dataset
	} else {
		cDes.Dataset = des.Dataset
	}
	if dcl.IsZeroValue(des.Project) || (dcl.IsEmptyValueIndirect(des.Project) && dcl.IsEmptyValueIndirect(initial.Project)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Project = initial.Project
	} else {
		cDes.Project = des.Project
	}
	if dcl.IsZeroValue(des.SnapshotTime) || (dcl.IsEmptyValueIndirect(des.SnapshotTime) && dcl.IsEmptyValueIndirect(initial.SnapshotTime)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.SnapshotTime = initial.SnapshotTime
	} else {
		cDes.SnapshotTime = des.SnapshotTime
	}

	return cDes
}

func canonicalizeTableSnapshotDefinitionSlice(des, initial []TableSnapshotDefinition, opts ...dcl.ApplyOption) []TableSnapshotDefinition {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TableSnapshotDefinition, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTableSnapshotDefinition(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TableSnapshotDefinition, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTableSnapshotDefinition(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTableSnapshotDefinition(c *Client, des, nw *TableSnapshotDefinition) *TableSnapshotDefinition {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for TableSnapshotDefinition while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewTableSnapshotDefinitionSet(c *Client, des, nw []TableSnapshotDefinition) []TableSnapshotDefinition {
	if des == nil {
		return nw
	}
	var reorderedNew []TableSnapshotDefinition
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareTableSnapshotDefinitionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewTableSnapshotDefinitionSlice(c *Client, des, nw []TableSnapshotDefinition) []TableSnapshotDefinition {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TableSnapshotDefinition
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTableSnapshotDefinition(c, &d, &n))
	}

	return items
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffTable(c *Client, desired, actual *Table, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SelfLink")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Dataset, actual.Dataset, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Dataset")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FriendlyName, actual.FriendlyName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("FriendlyName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Model, actual.Model, dcl.Info{OutputOnly: true, ObjectFunction: compareTableModelNewStyle, EmptyObject: EmptyTableModel, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Model")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Schema, actual.Schema, dcl.Info{ObjectFunction: compareTableSchemaNewStyle, EmptyObject: EmptyTableSchema, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Schema")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimePartitioning, actual.TimePartitioning, dcl.Info{ObjectFunction: compareTableTimePartitioningNewStyle, EmptyObject: EmptyTableTimePartitioning, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("TimePartitioning")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RangePartitioning, actual.RangePartitioning, dcl.Info{ObjectFunction: compareTableRangePartitioningNewStyle, EmptyObject: EmptyTableRangePartitioning, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("RangePartitioning")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Clustering, actual.Clustering, dcl.Info{ObjectFunction: compareTableClusteringNewStyle, EmptyObject: EmptyTableClustering, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Clustering")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequirePartitionFilter, actual.RequirePartitionFilter, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("RequirePartitionFilter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NumBytes, actual.NumBytes, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumBytes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NumPhysicalBytes, actual.NumPhysicalBytes, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumPhysicalBytes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NumLongTermBytes, actual.NumLongTermBytes, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumLongTermBytes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NumRows, actual.NumRows, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NumRows")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreationTime, actual.CreationTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreationTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExpirationTime, actual.ExpirationTime, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("ExpirationTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastModifiedTime, actual.LastModifiedTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LastModifiedTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.View, actual.View, dcl.Info{ObjectFunction: compareTableViewNewStyle, EmptyObject: EmptyTableView, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("View")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaterializedView, actual.MaterializedView, dcl.Info{ObjectFunction: compareTableMaterializedViewNewStyle, EmptyObject: EmptyTableMaterializedView, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("MaterializedView")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExternalDataConfiguration, actual.ExternalDataConfiguration, dcl.Info{ObjectFunction: compareTableExternalDataConfigurationNewStyle, EmptyObject: EmptyTableExternalDataConfiguration, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("ExternalDataConfiguration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StreamingBuffer, actual.StreamingBuffer, dcl.Info{OutputOnly: true, ObjectFunction: compareTableStreamingBufferNewStyle, EmptyObject: EmptyTableStreamingBuffer, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StreamingBuffer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EncryptionConfiguration, actual.EncryptionConfiguration, dcl.Info{ObjectFunction: compareTableEncryptionConfigurationNewStyle, EmptyObject: EmptyTableEncryptionConfiguration, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("EncryptionConfiguration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SnapshotDefinition, actual.SnapshotDefinition, dcl.Info{OutputOnly: true, ObjectFunction: compareTableSnapshotDefinitionNewStyle, EmptyObject: EmptyTableSnapshotDefinition, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SnapshotDefinition")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultCollation, actual.DefaultCollation, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("DefaultCollation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareTableModelNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableModel)
	if !ok {
		desiredNotPointer, ok := d.(TableModel)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableModel or *TableModel", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableModel)
	if !ok {
		actualNotPointer, ok := a.(TableModel)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableModel", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ModelOptions, actual.ModelOptions, dcl.Info{OutputOnly: true, ObjectFunction: compareTableModelModelOptionsNewStyle, EmptyObject: EmptyTableModelModelOptions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ModelOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TrainingRuns, actual.TrainingRuns, dcl.Info{OutputOnly: true, ObjectFunction: compareTableModelTrainingRunsNewStyle, EmptyObject: EmptyTableModelTrainingRuns, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TrainingRuns")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableModelModelOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableModelModelOptions)
	if !ok {
		desiredNotPointer, ok := d.(TableModelModelOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableModelModelOptions or *TableModelModelOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableModelModelOptions)
	if !ok {
		actualNotPointer, ok := a.(TableModelModelOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableModelModelOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ModelType, actual.ModelType, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ModelType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LossType, actual.LossType, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LossType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableModelTrainingRunsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableModelTrainingRuns)
	if !ok {
		desiredNotPointer, ok := d.(TableModelTrainingRuns)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableModelTrainingRuns or *TableModelTrainingRuns", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableModelTrainingRuns)
	if !ok {
		actualNotPointer, ok := a.(TableModelTrainingRuns)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableModelTrainingRuns", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StartTime, actual.StartTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StartTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TrainingOptions, actual.TrainingOptions, dcl.Info{OutputOnly: true, ObjectFunction: compareTableModelTrainingRunsTrainingOptionsNewStyle, EmptyObject: EmptyTableModelTrainingRunsTrainingOptions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TrainingOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IterationResults, actual.IterationResults, dcl.Info{OutputOnly: true, ObjectFunction: compareTableModelTrainingRunsIterationResultsNewStyle, EmptyObject: EmptyTableModelTrainingRunsIterationResults, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IterationResults")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableModelTrainingRunsTrainingOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableModelTrainingRunsTrainingOptions)
	if !ok {
		desiredNotPointer, ok := d.(TableModelTrainingRunsTrainingOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableModelTrainingRunsTrainingOptions or *TableModelTrainingRunsTrainingOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableModelTrainingRunsTrainingOptions)
	if !ok {
		actualNotPointer, ok := a.(TableModelTrainingRunsTrainingOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableModelTrainingRunsTrainingOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MaxIteration, actual.MaxIteration, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxIteration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LearnRate, actual.LearnRate, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LearnRate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.L1Reg, actual.L1Reg, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("L1Reg")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.L2Reg, actual.L2Reg, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("L2Reg")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinRelProgress, actual.MinRelProgress, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MinRelProgress")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WarmStart, actual.WarmStart, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("WarmStart")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EarlyStop, actual.EarlyStop, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EarlyStop")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LearnRateStrategy, actual.LearnRateStrategy, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LearnRateStrategy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LineSearchInitLearnRate, actual.LineSearchInitLearnRate, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LineSearchInitLearnRate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableModelTrainingRunsIterationResultsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableModelTrainingRunsIterationResults)
	if !ok {
		desiredNotPointer, ok := d.(TableModelTrainingRunsIterationResults)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableModelTrainingRunsIterationResults or *TableModelTrainingRunsIterationResults", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableModelTrainingRunsIterationResults)
	if !ok {
		actualNotPointer, ok := a.(TableModelTrainingRunsIterationResults)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableModelTrainingRunsIterationResults", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Index, actual.Index, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Index")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LearnRate, actual.LearnRate, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LearnRate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TrainingLoss, actual.TrainingLoss, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TrainingLoss")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EvalLoss, actual.EvalLoss, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EvalLoss")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DurationMs, actual.DurationMs, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DurationMs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableSchemaNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableSchema)
	if !ok {
		desiredNotPointer, ok := d.(TableSchema)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableSchema or *TableSchema", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableSchema)
	if !ok {
		actualNotPointer, ok := a.(TableSchema)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableSchema", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Fields, actual.Fields, dcl.Info{ObjectFunction: compareTableGooglecloudbigqueryv2TablefieldschemaNewStyle, EmptyObject: EmptyTableGooglecloudbigqueryv2Tablefieldschema, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Fields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableGooglecloudbigqueryv2TablefieldschemaNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableGooglecloudbigqueryv2Tablefieldschema)
	if !ok {
		desiredNotPointer, ok := d.(TableGooglecloudbigqueryv2Tablefieldschema)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableGooglecloudbigqueryv2Tablefieldschema or *TableGooglecloudbigqueryv2Tablefieldschema", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableGooglecloudbigqueryv2Tablefieldschema)
	if !ok {
		actualNotPointer, ok := a.(TableGooglecloudbigqueryv2Tablefieldschema)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableGooglecloudbigqueryv2Tablefieldschema", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Mode, actual.Mode, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Mode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Fields, actual.Fields, dcl.Info{ObjectFunction: compareTableGooglecloudbigqueryv2TablefieldschemaNewStyle, EmptyObject: EmptyTableGooglecloudbigqueryv2Tablefieldschema, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Fields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Categories, actual.Categories, dcl.Info{ObjectFunction: compareTableGooglecloudbigqueryv2TablefieldschemaCategoriesNewStyle, EmptyObject: EmptyTableGooglecloudbigqueryv2TablefieldschemaCategories, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Categories")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PolicyTags, actual.PolicyTags, dcl.Info{ObjectFunction: compareTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsNewStyle, EmptyObject: EmptyTableGooglecloudbigqueryv2TablefieldschemaPolicyTags, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("PolicyTags")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NameAlternative, actual.NameAlternative, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("NameAlternative")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxLength, actual.MaxLength, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("MaxLength")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Precision, actual.Precision, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Precision")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Scale, actual.Scale, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Scale")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Collation, actual.Collation, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Collation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultValueExpression, actual.DefaultValueExpression, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("DefaultValueExpression")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableGooglecloudbigqueryv2TablefieldschemaCategoriesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableGooglecloudbigqueryv2TablefieldschemaCategories)
	if !ok {
		desiredNotPointer, ok := d.(TableGooglecloudbigqueryv2TablefieldschemaCategories)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableGooglecloudbigqueryv2TablefieldschemaCategories or *TableGooglecloudbigqueryv2TablefieldschemaCategories", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableGooglecloudbigqueryv2TablefieldschemaCategories)
	if !ok {
		actualNotPointer, ok := a.(TableGooglecloudbigqueryv2TablefieldschemaCategories)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableGooglecloudbigqueryv2TablefieldschemaCategories", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Names, actual.Names, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Names")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableGooglecloudbigqueryv2TablefieldschemaPolicyTags)
	if !ok {
		desiredNotPointer, ok := d.(TableGooglecloudbigqueryv2TablefieldschemaPolicyTags)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableGooglecloudbigqueryv2TablefieldschemaPolicyTags or *TableGooglecloudbigqueryv2TablefieldschemaPolicyTags", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableGooglecloudbigqueryv2TablefieldschemaPolicyTags)
	if !ok {
		actualNotPointer, ok := a.(TableGooglecloudbigqueryv2TablefieldschemaPolicyTags)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableGooglecloudbigqueryv2TablefieldschemaPolicyTags", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Names, actual.Names, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Names")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableTimePartitioningNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableTimePartitioning)
	if !ok {
		desiredNotPointer, ok := d.(TableTimePartitioning)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableTimePartitioning or *TableTimePartitioning", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableTimePartitioning)
	if !ok {
		actualNotPointer, ok := a.(TableTimePartitioning)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableTimePartitioning", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExpirationMs, actual.ExpirationMs, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("ExpirationMs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Field, actual.Field, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Field")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableRangePartitioningNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableRangePartitioning)
	if !ok {
		desiredNotPointer, ok := d.(TableRangePartitioning)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableRangePartitioning or *TableRangePartitioning", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableRangePartitioning)
	if !ok {
		actualNotPointer, ok := a.(TableRangePartitioning)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableRangePartitioning", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Field, actual.Field, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Field")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Range, actual.Range, dcl.Info{ObjectFunction: compareTableRangePartitioningRangeNewStyle, EmptyObject: EmptyTableRangePartitioningRange, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Range")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableRangePartitioningRangeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableRangePartitioningRange)
	if !ok {
		desiredNotPointer, ok := d.(TableRangePartitioningRange)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableRangePartitioningRange or *TableRangePartitioningRange", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableRangePartitioningRange)
	if !ok {
		actualNotPointer, ok := a.(TableRangePartitioningRange)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableRangePartitioningRange", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Start, actual.Start, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Start")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.End, actual.End, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("End")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Interval, actual.Interval, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Interval")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableClusteringNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableClustering)
	if !ok {
		desiredNotPointer, ok := d.(TableClustering)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableClustering or *TableClustering", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableClustering)
	if !ok {
		actualNotPointer, ok := a.(TableClustering)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableClustering", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Fields, actual.Fields, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Fields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableViewNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableView)
	if !ok {
		desiredNotPointer, ok := d.(TableView)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableView or *TableView", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableView)
	if !ok {
		actualNotPointer, ok := a.(TableView)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableView", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Query, actual.Query, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Query")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UserDefinedFunctionResources, actual.UserDefinedFunctionResources, dcl.Info{ObjectFunction: compareTableViewUserDefinedFunctionResourcesNewStyle, EmptyObject: EmptyTableViewUserDefinedFunctionResources, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("UserDefinedFunctionResources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UseLegacySql, actual.UseLegacySql, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("UseLegacySql")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UseExplicitColumnNames, actual.UseExplicitColumnNames, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("UseExplicitColumnNames")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableViewUserDefinedFunctionResourcesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableViewUserDefinedFunctionResources)
	if !ok {
		desiredNotPointer, ok := d.(TableViewUserDefinedFunctionResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableViewUserDefinedFunctionResources or *TableViewUserDefinedFunctionResources", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableViewUserDefinedFunctionResources)
	if !ok {
		actualNotPointer, ok := a.(TableViewUserDefinedFunctionResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableViewUserDefinedFunctionResources", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ResourceUri, actual.ResourceUri, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("ResourceUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InlineCode, actual.InlineCode, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("InlineCode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InlineCodeAlternative, actual.InlineCodeAlternative, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("InlineCodeAlternative")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableMaterializedViewNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableMaterializedView)
	if !ok {
		desiredNotPointer, ok := d.(TableMaterializedView)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableMaterializedView or *TableMaterializedView", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableMaterializedView)
	if !ok {
		actualNotPointer, ok := a.(TableMaterializedView)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableMaterializedView", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Query, actual.Query, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Query")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastRefreshTime, actual.LastRefreshTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LastRefreshTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableRefresh, actual.EnableRefresh, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("EnableRefresh")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RefreshIntervalMs, actual.RefreshIntervalMs, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("RefreshIntervalMs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableExternalDataConfigurationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableExternalDataConfiguration)
	if !ok {
		desiredNotPointer, ok := d.(TableExternalDataConfiguration)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableExternalDataConfiguration or *TableExternalDataConfiguration", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableExternalDataConfiguration)
	if !ok {
		actualNotPointer, ok := a.(TableExternalDataConfiguration)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableExternalDataConfiguration", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SourceUris, actual.SourceUris, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("SourceUris")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Schema, actual.Schema, dcl.Info{ObjectFunction: compareTableExternalDataConfigurationSchemaNewStyle, EmptyObject: EmptyTableExternalDataConfigurationSchema, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Schema")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceFormat, actual.SourceFormat, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("SourceFormat")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxBadRecords, actual.MaxBadRecords, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("MaxBadRecords")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Autodetect, actual.Autodetect, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Autodetect")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IgnoreUnknownValues, actual.IgnoreUnknownValues, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("IgnoreUnknownValues")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Compression, actual.Compression, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Compression")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CsvOptions, actual.CsvOptions, dcl.Info{ObjectFunction: compareTableExternalDataConfigurationCsvOptionsNewStyle, EmptyObject: EmptyTableExternalDataConfigurationCsvOptions, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("CsvOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BigtableOptions, actual.BigtableOptions, dcl.Info{ObjectFunction: compareTableExternalDataConfigurationBigtableOptionsNewStyle, EmptyObject: EmptyTableExternalDataConfigurationBigtableOptions, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("BigtableOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GoogleSheetsOptions, actual.GoogleSheetsOptions, dcl.Info{ObjectFunction: compareTableExternalDataConfigurationGoogleSheetsOptionsNewStyle, EmptyObject: EmptyTableExternalDataConfigurationGoogleSheetsOptions, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("GoogleSheetsOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxBadRecordsAlternative, actual.MaxBadRecordsAlternative, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("MaxBadRecordsAlternative")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HivePartitioningOptions, actual.HivePartitioningOptions, dcl.Info{ObjectFunction: compareTableExternalDataConfigurationHivePartitioningOptionsNewStyle, EmptyObject: EmptyTableExternalDataConfigurationHivePartitioningOptions, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("HivePartitioningOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConnectionId, actual.ConnectionId, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("ConnectionId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ValueConversionModes, actual.ValueConversionModes, dcl.Info{ObjectFunction: compareTableExternalDataConfigurationValueConversionModesNewStyle, EmptyObject: EmptyTableExternalDataConfigurationValueConversionModes, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("ValueConversionModes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DecimalTargetTypes, actual.DecimalTargetTypes, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("DecimalTargetTypes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AvroOptions, actual.AvroOptions, dcl.Info{ObjectFunction: compareTableExternalDataConfigurationAvroOptionsNewStyle, EmptyObject: EmptyTableExternalDataConfigurationAvroOptions, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("AvroOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.JsonExtension, actual.JsonExtension, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("JsonExtension")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ParquetOptions, actual.ParquetOptions, dcl.Info{ObjectFunction: compareTableExternalDataConfigurationParquetOptionsNewStyle, EmptyObject: EmptyTableExternalDataConfigurationParquetOptions, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("ParquetOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableExternalDataConfigurationSchemaNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableExternalDataConfigurationSchema)
	if !ok {
		desiredNotPointer, ok := d.(TableExternalDataConfigurationSchema)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableExternalDataConfigurationSchema or *TableExternalDataConfigurationSchema", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableExternalDataConfigurationSchema)
	if !ok {
		actualNotPointer, ok := a.(TableExternalDataConfigurationSchema)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableExternalDataConfigurationSchema", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Fields, actual.Fields, dcl.Info{ObjectFunction: compareTableGooglecloudbigqueryv2TablefieldschemaNewStyle, EmptyObject: EmptyTableGooglecloudbigqueryv2Tablefieldschema, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Fields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableExternalDataConfigurationCsvOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableExternalDataConfigurationCsvOptions)
	if !ok {
		desiredNotPointer, ok := d.(TableExternalDataConfigurationCsvOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableExternalDataConfigurationCsvOptions or *TableExternalDataConfigurationCsvOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableExternalDataConfigurationCsvOptions)
	if !ok {
		actualNotPointer, ok := a.(TableExternalDataConfigurationCsvOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableExternalDataConfigurationCsvOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.FieldDelimiter, actual.FieldDelimiter, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("FieldDelimiter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SkipLeadingRows, actual.SkipLeadingRows, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("SkipLeadingRows")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Quote, actual.Quote, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Quote")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowQuotedNewlines, actual.AllowQuotedNewlines, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("AllowQuotedNewlines")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowJaggedRows, actual.AllowJaggedRows, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("AllowJaggedRows")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Encoding, actual.Encoding, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Encoding")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableExternalDataConfigurationBigtableOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableExternalDataConfigurationBigtableOptions)
	if !ok {
		desiredNotPointer, ok := d.(TableExternalDataConfigurationBigtableOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableExternalDataConfigurationBigtableOptions or *TableExternalDataConfigurationBigtableOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableExternalDataConfigurationBigtableOptions)
	if !ok {
		actualNotPointer, ok := a.(TableExternalDataConfigurationBigtableOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableExternalDataConfigurationBigtableOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ColumnFamilies, actual.ColumnFamilies, dcl.Info{ObjectFunction: compareTableExternalDataConfigurationBigtableOptionsColumnFamiliesNewStyle, EmptyObject: EmptyTableExternalDataConfigurationBigtableOptionsColumnFamilies, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("ColumnFamilies")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IgnoreUnspecifiedColumnFamilies, actual.IgnoreUnspecifiedColumnFamilies, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("IgnoreUnspecifiedColumnFamilies")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReadRowkeyAsString, actual.ReadRowkeyAsString, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("ReadRowkeyAsString")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableExternalDataConfigurationBigtableOptionsColumnFamiliesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableExternalDataConfigurationBigtableOptionsColumnFamilies)
	if !ok {
		desiredNotPointer, ok := d.(TableExternalDataConfigurationBigtableOptionsColumnFamilies)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableExternalDataConfigurationBigtableOptionsColumnFamilies or *TableExternalDataConfigurationBigtableOptionsColumnFamilies", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableExternalDataConfigurationBigtableOptionsColumnFamilies)
	if !ok {
		actualNotPointer, ok := a.(TableExternalDataConfigurationBigtableOptionsColumnFamilies)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableExternalDataConfigurationBigtableOptionsColumnFamilies", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.FamilyId, actual.FamilyId, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("FamilyId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Encoding, actual.Encoding, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Encoding")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Columns, actual.Columns, dcl.Info{ObjectFunction: compareTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsNewStyle, EmptyObject: EmptyTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns, OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Columns")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OnlyReadLatest, actual.OnlyReadLatest, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("OnlyReadLatest")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns)
	if !ok {
		desiredNotPointer, ok := d.(TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns or *TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns)
	if !ok {
		actualNotPointer, ok := a.(TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.QualifierEncoded, actual.QualifierEncoded, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("QualifierEncoded")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.QualifierString, actual.QualifierString, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("QualifierString")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FieldName, actual.FieldName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("FieldName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Encoding, actual.Encoding, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Encoding")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OnlyReadLatest, actual.OnlyReadLatest, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("OnlyReadLatest")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableExternalDataConfigurationGoogleSheetsOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableExternalDataConfigurationGoogleSheetsOptions)
	if !ok {
		desiredNotPointer, ok := d.(TableExternalDataConfigurationGoogleSheetsOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableExternalDataConfigurationGoogleSheetsOptions or *TableExternalDataConfigurationGoogleSheetsOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableExternalDataConfigurationGoogleSheetsOptions)
	if !ok {
		actualNotPointer, ok := a.(TableExternalDataConfigurationGoogleSheetsOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableExternalDataConfigurationGoogleSheetsOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SkipLeadingRows, actual.SkipLeadingRows, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("SkipLeadingRows")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Range, actual.Range, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Range")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableExternalDataConfigurationHivePartitioningOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableExternalDataConfigurationHivePartitioningOptions)
	if !ok {
		desiredNotPointer, ok := d.(TableExternalDataConfigurationHivePartitioningOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableExternalDataConfigurationHivePartitioningOptions or *TableExternalDataConfigurationHivePartitioningOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableExternalDataConfigurationHivePartitioningOptions)
	if !ok {
		actualNotPointer, ok := a.(TableExternalDataConfigurationHivePartitioningOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableExternalDataConfigurationHivePartitioningOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Mode, actual.Mode, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("Mode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceUriPrefix, actual.SourceUriPrefix, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("SourceUriPrefix")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequirePartitionFilter, actual.RequirePartitionFilter, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("RequirePartitionFilter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Fields, actual.Fields, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Fields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableExternalDataConfigurationValueConversionModesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableExternalDataConfigurationValueConversionModes)
	if !ok {
		desiredNotPointer, ok := d.(TableExternalDataConfigurationValueConversionModes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableExternalDataConfigurationValueConversionModes or *TableExternalDataConfigurationValueConversionModes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableExternalDataConfigurationValueConversionModes)
	if !ok {
		actualNotPointer, ok := a.(TableExternalDataConfigurationValueConversionModes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableExternalDataConfigurationValueConversionModes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TemporalTypesOutOfRangeConversionMode, actual.TemporalTypesOutOfRangeConversionMode, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("TemporalTypesOutOfRangeConversionMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NumericTypeOutOfRangeConversionMode, actual.NumericTypeOutOfRangeConversionMode, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("NumericTypeOutOfRangeConversionMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableExternalDataConfigurationAvroOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableExternalDataConfigurationAvroOptions)
	if !ok {
		desiredNotPointer, ok := d.(TableExternalDataConfigurationAvroOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableExternalDataConfigurationAvroOptions or *TableExternalDataConfigurationAvroOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableExternalDataConfigurationAvroOptions)
	if !ok {
		actualNotPointer, ok := a.(TableExternalDataConfigurationAvroOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableExternalDataConfigurationAvroOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.UseAvroLogicalTypes, actual.UseAvroLogicalTypes, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("UseAvroLogicalTypes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableExternalDataConfigurationParquetOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableExternalDataConfigurationParquetOptions)
	if !ok {
		desiredNotPointer, ok := d.(TableExternalDataConfigurationParquetOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableExternalDataConfigurationParquetOptions or *TableExternalDataConfigurationParquetOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableExternalDataConfigurationParquetOptions)
	if !ok {
		actualNotPointer, ok := a.(TableExternalDataConfigurationParquetOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableExternalDataConfigurationParquetOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EnumAsString, actual.EnumAsString, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("EnumAsString")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableListInference, actual.EnableListInference, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("EnableListInference")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableStreamingBufferNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableStreamingBuffer)
	if !ok {
		desiredNotPointer, ok := d.(TableStreamingBuffer)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableStreamingBuffer or *TableStreamingBuffer", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableStreamingBuffer)
	if !ok {
		actualNotPointer, ok := a.(TableStreamingBuffer)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableStreamingBuffer", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EstimatedBytes, actual.EstimatedBytes, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EstimatedBytes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EstimatedRows, actual.EstimatedRows, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EstimatedRows")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OldestEntryTime, actual.OldestEntryTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("OldestEntryTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableEncryptionConfigurationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableEncryptionConfiguration)
	if !ok {
		desiredNotPointer, ok := d.(TableEncryptionConfiguration)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableEncryptionConfiguration or *TableEncryptionConfiguration", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableEncryptionConfiguration)
	if !ok {
		actualNotPointer, ok := a.(TableEncryptionConfiguration)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableEncryptionConfiguration", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.KmsKeyName, actual.KmsKeyName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("KmsKeyName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTableSnapshotDefinitionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TableSnapshotDefinition)
	if !ok {
		desiredNotPointer, ok := d.(TableSnapshotDefinition)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableSnapshotDefinition or *TableSnapshotDefinition", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TableSnapshotDefinition)
	if !ok {
		actualNotPointer, ok := a.(TableSnapshotDefinition)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TableSnapshotDefinition", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Table, actual.Table, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Table")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Dataset, actual.Dataset, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Dataset")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SnapshotTime, actual.SnapshotTime, dcl.Info{OperationSelector: dcl.TriggersOperation("updateTablePatchTableOperation")}, fn.AddNest("SnapshotTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Table) urlNormalized() *Table {
	normalized := dcl.Copy(*r).(Table)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Id = dcl.SelfLinkToName(r.Id)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Dataset = dcl.SelfLinkToName(r.Dataset)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.FriendlyName = dcl.SelfLinkToName(r.FriendlyName)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.NumBytes = dcl.SelfLinkToName(r.NumBytes)
	normalized.NumPhysicalBytes = dcl.SelfLinkToName(r.NumPhysicalBytes)
	normalized.NumLongTermBytes = dcl.SelfLinkToName(r.NumLongTermBytes)
	normalized.Type = dcl.SelfLinkToName(r.Type)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.DefaultCollation = dcl.SelfLinkToName(r.DefaultCollation)
	return &normalized
}

func (r *Table) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "PatchTable" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(nr.Project),
			"dataset": dcl.ValueOrEmptyString(nr.Dataset),
			"name":    dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/datasets/{{dataset}}/tables/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Table resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Table) marshal(c *Client) ([]byte, error) {
	m, err := expandTable(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Table: %w", err)
	}
	dcl.MoveMapEntry(
		m,
		[]string{"name"},
		[]string{"tableReference", "tableId"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"dataset"},
		[]string{"tableReference", "datasetId"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"project"},
		[]string{"tableReference", "projectId"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"snapshotDefinition", "table"},
		[]string{"snapshotDefinition", "tableReference", "tableId"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"snapshotDefinition", "dataset"},
		[]string{"snapshotDefinition", "tableReference", "datasetId"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"snapshotDefinition", "project"},
		[]string{"snapshotDefinition", "tableReference", "projectId"},
	)

	return json.Marshal(m)
}

// unmarshalTable decodes JSON responses into the Table resource schema.
func unmarshalTable(b []byte, c *Client, res *Table) (*Table, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapTable(m, c, res)
}

func unmarshalMapTable(m map[string]interface{}, c *Client, res *Table) (*Table, error) {
	dcl.MoveMapEntry(
		m,
		[]string{"tableReference", "tableId"},
		[]string{"name"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"tableReference", "datasetId"},
		[]string{"dataset"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"tableReference", "projectId"},
		[]string{"project"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"snapshotDefinition", "tableReference", "tableId"},
		[]string{"snapshotDefinition", "table"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"snapshotDefinition", "tableReference", "datasetId"},
		[]string{"snapshotDefinition", "dataset"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"snapshotDefinition", "tableReference", "projectId"},
		[]string{"snapshotDefinition", "project"},
	)

	flattened := flattenTable(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandTable expands Table into a JSON request object.
func expandTable(c *Client, f *Table) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v := f.Name; dcl.ValueShouldBeSent(v) {
		m["name"] = v
	}
	if v := f.Dataset; dcl.ValueShouldBeSent(v) {
		m["dataset"] = v
	}
	if v := f.Project; dcl.ValueShouldBeSent(v) {
		m["project"] = v
	}
	if v := f.FriendlyName; dcl.ValueShouldBeSent(v) {
		m["friendlyName"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
	}
	if v, err := expandTableSchema(c, f.Schema, res); err != nil {
		return nil, fmt.Errorf("error expanding Schema into schema: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["schema"] = v
	}
	if v, err := expandTableTimePartitioning(c, f.TimePartitioning, res); err != nil {
		return nil, fmt.Errorf("error expanding TimePartitioning into timePartitioning: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["timePartitioning"] = v
	}
	if v, err := expandTableRangePartitioning(c, f.RangePartitioning, res); err != nil {
		return nil, fmt.Errorf("error expanding RangePartitioning into rangePartitioning: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["rangePartitioning"] = v
	}
	if v, err := expandTableClustering(c, f.Clustering, res); err != nil {
		return nil, fmt.Errorf("error expanding Clustering into clustering: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["clustering"] = v
	}
	if v := f.RequirePartitionFilter; dcl.ValueShouldBeSent(v) {
		m["requirePartitionFilter"] = v
	}
	if v := f.ExpirationTime; dcl.ValueShouldBeSent(v) {
		m["expirationTime"] = v
	}
	if v, err := expandTableView(c, f.View, res); err != nil {
		return nil, fmt.Errorf("error expanding View into view: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["view"] = v
	}
	if v, err := expandTableMaterializedView(c, f.MaterializedView, res); err != nil {
		return nil, fmt.Errorf("error expanding MaterializedView into materializedView: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["materializedView"] = v
	}
	if v, err := expandTableExternalDataConfiguration(c, f.ExternalDataConfiguration, res); err != nil {
		return nil, fmt.Errorf("error expanding ExternalDataConfiguration into externalDataConfiguration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["externalDataConfiguration"] = v
	}
	if v, err := expandTableEncryptionConfiguration(c, f.EncryptionConfiguration, res); err != nil {
		return nil, fmt.Errorf("error expanding EncryptionConfiguration into encryptionConfiguration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["encryptionConfiguration"] = v
	}
	if v := f.DefaultCollation; dcl.ValueShouldBeSent(v) {
		m["defaultCollation"] = v
	}

	return m, nil
}

// flattenTable flattens Table from a JSON request object into the
// Table type.
func flattenTable(c *Client, i interface{}, res *Table) *Table {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Table{}
	resultRes.Etag = dcl.FlattenString(m["etag"])
	resultRes.Id = dcl.FlattenString(m["id"])
	resultRes.SelfLink = dcl.FlattenString(m["selfLink"])
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Dataset = dcl.FlattenString(m["dataset"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.FriendlyName = dcl.FlattenString(m["friendlyName"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.Model = flattenTableModel(c, m["model"], res)
	resultRes.Schema = flattenTableSchema(c, m["schema"], res)
	resultRes.TimePartitioning = flattenTableTimePartitioning(c, m["timePartitioning"], res)
	resultRes.RangePartitioning = flattenTableRangePartitioning(c, m["rangePartitioning"], res)
	resultRes.Clustering = flattenTableClustering(c, m["clustering"], res)
	resultRes.RequirePartitionFilter = dcl.FlattenBool(m["requirePartitionFilter"])
	resultRes.NumBytes = dcl.FlattenString(m["numBytes"])
	resultRes.NumPhysicalBytes = dcl.FlattenString(m["numPhysicalBytes"])
	resultRes.NumLongTermBytes = dcl.FlattenString(m["numLongTermBytes"])
	resultRes.NumRows = dcl.FlattenInteger(m["numRows"])
	resultRes.CreationTime = dcl.FlattenInteger(m["creationTime"])
	resultRes.ExpirationTime = dcl.FlattenInteger(m["expirationTime"])
	resultRes.LastModifiedTime = dcl.FlattenInteger(m["lastModifiedTime"])
	resultRes.Type = dcl.FlattenString(m["type"])
	resultRes.View = flattenTableView(c, m["view"], res)
	resultRes.MaterializedView = flattenTableMaterializedView(c, m["materializedView"], res)
	resultRes.ExternalDataConfiguration = flattenTableExternalDataConfiguration(c, m["externalDataConfiguration"], res)
	resultRes.Location = dcl.FlattenString(m["location"])
	resultRes.StreamingBuffer = flattenTableStreamingBuffer(c, m["streamingBuffer"], res)
	resultRes.EncryptionConfiguration = flattenTableEncryptionConfiguration(c, m["encryptionConfiguration"], res)
	resultRes.SnapshotDefinition = flattenTableSnapshotDefinition(c, m["snapshotDefinition"], res)
	resultRes.DefaultCollation = dcl.FlattenString(m["defaultCollation"])

	return resultRes
}

// expandTableModelMap expands the contents of TableModel into a JSON
// request object.
func expandTableModelMap(c *Client, f map[string]TableModel, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableModel(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableModelSlice expands the contents of TableModel into a JSON
// request object.
func expandTableModelSlice(c *Client, f []TableModel, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableModel(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableModelMap flattens the contents of TableModel from a JSON
// response object.
func flattenTableModelMap(c *Client, i interface{}, res *Table) map[string]TableModel {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableModel{}
	}

	if len(a) == 0 {
		return map[string]TableModel{}
	}

	items := make(map[string]TableModel)
	for k, item := range a {
		items[k] = *flattenTableModel(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableModelSlice flattens the contents of TableModel from a JSON
// response object.
func flattenTableModelSlice(c *Client, i interface{}, res *Table) []TableModel {
	a, ok := i.([]interface{})
	if !ok {
		return []TableModel{}
	}

	if len(a) == 0 {
		return []TableModel{}
	}

	items := make([]TableModel, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableModel(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableModel expands an instance of TableModel into a JSON
// request object.
func expandTableModel(c *Client, f *TableModel, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenTableModel flattens an instance of TableModel from a JSON
// response object.
func flattenTableModel(c *Client, i interface{}, res *Table) *TableModel {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableModel{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableModel
	}
	r.ModelOptions = flattenTableModelModelOptions(c, m["modelOptions"], res)
	r.TrainingRuns = flattenTableModelTrainingRunsSlice(c, m["trainingRuns"], res)

	return r
}

// expandTableModelModelOptionsMap expands the contents of TableModelModelOptions into a JSON
// request object.
func expandTableModelModelOptionsMap(c *Client, f map[string]TableModelModelOptions, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableModelModelOptions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableModelModelOptionsSlice expands the contents of TableModelModelOptions into a JSON
// request object.
func expandTableModelModelOptionsSlice(c *Client, f []TableModelModelOptions, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableModelModelOptions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableModelModelOptionsMap flattens the contents of TableModelModelOptions from a JSON
// response object.
func flattenTableModelModelOptionsMap(c *Client, i interface{}, res *Table) map[string]TableModelModelOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableModelModelOptions{}
	}

	if len(a) == 0 {
		return map[string]TableModelModelOptions{}
	}

	items := make(map[string]TableModelModelOptions)
	for k, item := range a {
		items[k] = *flattenTableModelModelOptions(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableModelModelOptionsSlice flattens the contents of TableModelModelOptions from a JSON
// response object.
func flattenTableModelModelOptionsSlice(c *Client, i interface{}, res *Table) []TableModelModelOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []TableModelModelOptions{}
	}

	if len(a) == 0 {
		return []TableModelModelOptions{}
	}

	items := make([]TableModelModelOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableModelModelOptions(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableModelModelOptions expands an instance of TableModelModelOptions into a JSON
// request object.
func expandTableModelModelOptions(c *Client, f *TableModelModelOptions, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenTableModelModelOptions flattens an instance of TableModelModelOptions from a JSON
// response object.
func flattenTableModelModelOptions(c *Client, i interface{}, res *Table) *TableModelModelOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableModelModelOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableModelModelOptions
	}
	r.ModelType = dcl.FlattenString(m["modelType"])
	r.Labels = dcl.FlattenStringSlice(m["labels"])
	r.LossType = dcl.FlattenString(m["lossType"])

	return r
}

// expandTableModelTrainingRunsMap expands the contents of TableModelTrainingRuns into a JSON
// request object.
func expandTableModelTrainingRunsMap(c *Client, f map[string]TableModelTrainingRuns, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableModelTrainingRuns(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableModelTrainingRunsSlice expands the contents of TableModelTrainingRuns into a JSON
// request object.
func expandTableModelTrainingRunsSlice(c *Client, f []TableModelTrainingRuns, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableModelTrainingRuns(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableModelTrainingRunsMap flattens the contents of TableModelTrainingRuns from a JSON
// response object.
func flattenTableModelTrainingRunsMap(c *Client, i interface{}, res *Table) map[string]TableModelTrainingRuns {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableModelTrainingRuns{}
	}

	if len(a) == 0 {
		return map[string]TableModelTrainingRuns{}
	}

	items := make(map[string]TableModelTrainingRuns)
	for k, item := range a {
		items[k] = *flattenTableModelTrainingRuns(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableModelTrainingRunsSlice flattens the contents of TableModelTrainingRuns from a JSON
// response object.
func flattenTableModelTrainingRunsSlice(c *Client, i interface{}, res *Table) []TableModelTrainingRuns {
	a, ok := i.([]interface{})
	if !ok {
		return []TableModelTrainingRuns{}
	}

	if len(a) == 0 {
		return []TableModelTrainingRuns{}
	}

	items := make([]TableModelTrainingRuns, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableModelTrainingRuns(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableModelTrainingRuns expands an instance of TableModelTrainingRuns into a JSON
// request object.
func expandTableModelTrainingRuns(c *Client, f *TableModelTrainingRuns, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenTableModelTrainingRuns flattens an instance of TableModelTrainingRuns from a JSON
// response object.
func flattenTableModelTrainingRuns(c *Client, i interface{}, res *Table) *TableModelTrainingRuns {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableModelTrainingRuns{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableModelTrainingRuns
	}
	r.State = dcl.FlattenString(m["state"])
	r.StartTime = dcl.FlattenString(m["startTime"])
	r.TrainingOptions = flattenTableModelTrainingRunsTrainingOptions(c, m["trainingOptions"], res)
	r.IterationResults = flattenTableModelTrainingRunsIterationResultsSlice(c, m["iterationResults"], res)

	return r
}

// expandTableModelTrainingRunsTrainingOptionsMap expands the contents of TableModelTrainingRunsTrainingOptions into a JSON
// request object.
func expandTableModelTrainingRunsTrainingOptionsMap(c *Client, f map[string]TableModelTrainingRunsTrainingOptions, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableModelTrainingRunsTrainingOptions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableModelTrainingRunsTrainingOptionsSlice expands the contents of TableModelTrainingRunsTrainingOptions into a JSON
// request object.
func expandTableModelTrainingRunsTrainingOptionsSlice(c *Client, f []TableModelTrainingRunsTrainingOptions, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableModelTrainingRunsTrainingOptions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableModelTrainingRunsTrainingOptionsMap flattens the contents of TableModelTrainingRunsTrainingOptions from a JSON
// response object.
func flattenTableModelTrainingRunsTrainingOptionsMap(c *Client, i interface{}, res *Table) map[string]TableModelTrainingRunsTrainingOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableModelTrainingRunsTrainingOptions{}
	}

	if len(a) == 0 {
		return map[string]TableModelTrainingRunsTrainingOptions{}
	}

	items := make(map[string]TableModelTrainingRunsTrainingOptions)
	for k, item := range a {
		items[k] = *flattenTableModelTrainingRunsTrainingOptions(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableModelTrainingRunsTrainingOptionsSlice flattens the contents of TableModelTrainingRunsTrainingOptions from a JSON
// response object.
func flattenTableModelTrainingRunsTrainingOptionsSlice(c *Client, i interface{}, res *Table) []TableModelTrainingRunsTrainingOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []TableModelTrainingRunsTrainingOptions{}
	}

	if len(a) == 0 {
		return []TableModelTrainingRunsTrainingOptions{}
	}

	items := make([]TableModelTrainingRunsTrainingOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableModelTrainingRunsTrainingOptions(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableModelTrainingRunsTrainingOptions expands an instance of TableModelTrainingRunsTrainingOptions into a JSON
// request object.
func expandTableModelTrainingRunsTrainingOptions(c *Client, f *TableModelTrainingRunsTrainingOptions, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenTableModelTrainingRunsTrainingOptions flattens an instance of TableModelTrainingRunsTrainingOptions from a JSON
// response object.
func flattenTableModelTrainingRunsTrainingOptions(c *Client, i interface{}, res *Table) *TableModelTrainingRunsTrainingOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableModelTrainingRunsTrainingOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableModelTrainingRunsTrainingOptions
	}
	r.MaxIteration = dcl.FlattenInteger(m["maxIteration"])
	r.LearnRate = dcl.FlattenDouble(m["learnRate"])
	r.L1Reg = dcl.FlattenDouble(m["l1Reg"])
	r.L2Reg = dcl.FlattenDouble(m["l2Reg"])
	r.MinRelProgress = dcl.FlattenDouble(m["minRelProgress"])
	r.WarmStart = dcl.FlattenBool(m["warmStart"])
	r.EarlyStop = dcl.FlattenBool(m["earlyStop"])
	r.LearnRateStrategy = dcl.FlattenString(m["learnRateStrategy"])
	r.LineSearchInitLearnRate = dcl.FlattenDouble(m["lineSearchInitLearnRate"])

	return r
}

// expandTableModelTrainingRunsIterationResultsMap expands the contents of TableModelTrainingRunsIterationResults into a JSON
// request object.
func expandTableModelTrainingRunsIterationResultsMap(c *Client, f map[string]TableModelTrainingRunsIterationResults, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableModelTrainingRunsIterationResults(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableModelTrainingRunsIterationResultsSlice expands the contents of TableModelTrainingRunsIterationResults into a JSON
// request object.
func expandTableModelTrainingRunsIterationResultsSlice(c *Client, f []TableModelTrainingRunsIterationResults, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableModelTrainingRunsIterationResults(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableModelTrainingRunsIterationResultsMap flattens the contents of TableModelTrainingRunsIterationResults from a JSON
// response object.
func flattenTableModelTrainingRunsIterationResultsMap(c *Client, i interface{}, res *Table) map[string]TableModelTrainingRunsIterationResults {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableModelTrainingRunsIterationResults{}
	}

	if len(a) == 0 {
		return map[string]TableModelTrainingRunsIterationResults{}
	}

	items := make(map[string]TableModelTrainingRunsIterationResults)
	for k, item := range a {
		items[k] = *flattenTableModelTrainingRunsIterationResults(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableModelTrainingRunsIterationResultsSlice flattens the contents of TableModelTrainingRunsIterationResults from a JSON
// response object.
func flattenTableModelTrainingRunsIterationResultsSlice(c *Client, i interface{}, res *Table) []TableModelTrainingRunsIterationResults {
	a, ok := i.([]interface{})
	if !ok {
		return []TableModelTrainingRunsIterationResults{}
	}

	if len(a) == 0 {
		return []TableModelTrainingRunsIterationResults{}
	}

	items := make([]TableModelTrainingRunsIterationResults, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableModelTrainingRunsIterationResults(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableModelTrainingRunsIterationResults expands an instance of TableModelTrainingRunsIterationResults into a JSON
// request object.
func expandTableModelTrainingRunsIterationResults(c *Client, f *TableModelTrainingRunsIterationResults, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenTableModelTrainingRunsIterationResults flattens an instance of TableModelTrainingRunsIterationResults from a JSON
// response object.
func flattenTableModelTrainingRunsIterationResults(c *Client, i interface{}, res *Table) *TableModelTrainingRunsIterationResults {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableModelTrainingRunsIterationResults{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableModelTrainingRunsIterationResults
	}
	r.Index = dcl.FlattenInteger(m["index"])
	r.LearnRate = dcl.FlattenDouble(m["learnRate"])
	r.TrainingLoss = dcl.FlattenDouble(m["trainingLoss"])
	r.EvalLoss = dcl.FlattenDouble(m["evalLoss"])
	r.DurationMs = dcl.FlattenString(m["durationMs"])

	return r
}

// expandTableSchemaMap expands the contents of TableSchema into a JSON
// request object.
func expandTableSchemaMap(c *Client, f map[string]TableSchema, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableSchema(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableSchemaSlice expands the contents of TableSchema into a JSON
// request object.
func expandTableSchemaSlice(c *Client, f []TableSchema, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableSchema(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableSchemaMap flattens the contents of TableSchema from a JSON
// response object.
func flattenTableSchemaMap(c *Client, i interface{}, res *Table) map[string]TableSchema {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableSchema{}
	}

	if len(a) == 0 {
		return map[string]TableSchema{}
	}

	items := make(map[string]TableSchema)
	for k, item := range a {
		items[k] = *flattenTableSchema(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableSchemaSlice flattens the contents of TableSchema from a JSON
// response object.
func flattenTableSchemaSlice(c *Client, i interface{}, res *Table) []TableSchema {
	a, ok := i.([]interface{})
	if !ok {
		return []TableSchema{}
	}

	if len(a) == 0 {
		return []TableSchema{}
	}

	items := make([]TableSchema, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableSchema(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableSchema expands an instance of TableSchema into a JSON
// request object.
func expandTableSchema(c *Client, f *TableSchema, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandTableGooglecloudbigqueryv2TablefieldschemaSlice(c, f.Fields, res); err != nil {
		return nil, fmt.Errorf("error expanding Fields into fields: %w", err)
	} else if v != nil {
		m["fields"] = v
	}

	return m, nil
}

// flattenTableSchema flattens an instance of TableSchema from a JSON
// response object.
func flattenTableSchema(c *Client, i interface{}, res *Table) *TableSchema {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableSchema{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableSchema
	}
	r.Fields = flattenTableGooglecloudbigqueryv2TablefieldschemaSlice(c, m["fields"], res)

	return r
}

// expandTableGooglecloudbigqueryv2TablefieldschemaMap expands the contents of TableGooglecloudbigqueryv2Tablefieldschema into a JSON
// request object.
func expandTableGooglecloudbigqueryv2TablefieldschemaMap(c *Client, f map[string]TableGooglecloudbigqueryv2Tablefieldschema, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableGooglecloudbigqueryv2Tablefieldschema(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableGooglecloudbigqueryv2TablefieldschemaSlice expands the contents of TableGooglecloudbigqueryv2Tablefieldschema into a JSON
// request object.
func expandTableGooglecloudbigqueryv2TablefieldschemaSlice(c *Client, f []TableGooglecloudbigqueryv2Tablefieldschema, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableGooglecloudbigqueryv2Tablefieldschema(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableGooglecloudbigqueryv2TablefieldschemaMap flattens the contents of TableGooglecloudbigqueryv2Tablefieldschema from a JSON
// response object.
func flattenTableGooglecloudbigqueryv2TablefieldschemaMap(c *Client, i interface{}, res *Table) map[string]TableGooglecloudbigqueryv2Tablefieldschema {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableGooglecloudbigqueryv2Tablefieldschema{}
	}

	if len(a) == 0 {
		return map[string]TableGooglecloudbigqueryv2Tablefieldschema{}
	}

	items := make(map[string]TableGooglecloudbigqueryv2Tablefieldschema)
	for k, item := range a {
		items[k] = *flattenTableGooglecloudbigqueryv2Tablefieldschema(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableGooglecloudbigqueryv2TablefieldschemaSlice flattens the contents of TableGooglecloudbigqueryv2Tablefieldschema from a JSON
// response object.
func flattenTableGooglecloudbigqueryv2TablefieldschemaSlice(c *Client, i interface{}, res *Table) []TableGooglecloudbigqueryv2Tablefieldschema {
	a, ok := i.([]interface{})
	if !ok {
		return []TableGooglecloudbigqueryv2Tablefieldschema{}
	}

	if len(a) == 0 {
		return []TableGooglecloudbigqueryv2Tablefieldschema{}
	}

	items := make([]TableGooglecloudbigqueryv2Tablefieldschema, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableGooglecloudbigqueryv2Tablefieldschema(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableGooglecloudbigqueryv2Tablefieldschema expands an instance of TableGooglecloudbigqueryv2Tablefieldschema into a JSON
// request object.
func expandTableGooglecloudbigqueryv2Tablefieldschema(c *Client, f *TableGooglecloudbigqueryv2Tablefieldschema, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.Mode; !dcl.IsEmptyValueIndirect(v) {
		m["mode"] = v
	}
	if v := f.Fields; v != nil {
		m["fields"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v, err := expandTableGooglecloudbigqueryv2TablefieldschemaCategories(c, f.Categories, res); err != nil {
		return nil, fmt.Errorf("error expanding Categories into categories: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["categories"] = v
	}
	if v, err := expandTableGooglecloudbigqueryv2TablefieldschemaPolicyTags(c, f.PolicyTags, res); err != nil {
		return nil, fmt.Errorf("error expanding PolicyTags into policyTags: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["policyTags"] = v
	}
	if v := f.NameAlternative; v != nil {
		m["nameAlternative"] = v
	}
	if v := f.MaxLength; !dcl.IsEmptyValueIndirect(v) {
		m["maxLength"] = v
	}
	if v := f.Precision; !dcl.IsEmptyValueIndirect(v) {
		m["precision"] = v
	}
	if v := f.Scale; !dcl.IsEmptyValueIndirect(v) {
		m["scale"] = v
	}
	if v := f.Collation; !dcl.IsEmptyValueIndirect(v) {
		m["collation"] = v
	}
	if v := f.DefaultValueExpression; !dcl.IsEmptyValueIndirect(v) {
		m["defaultValueExpression"] = v
	}

	return m, nil
}

// flattenTableGooglecloudbigqueryv2Tablefieldschema flattens an instance of TableGooglecloudbigqueryv2Tablefieldschema from a JSON
// response object.
func flattenTableGooglecloudbigqueryv2Tablefieldschema(c *Client, i interface{}, res *Table) *TableGooglecloudbigqueryv2Tablefieldschema {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableGooglecloudbigqueryv2Tablefieldschema{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableGooglecloudbigqueryv2Tablefieldschema
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Type = dcl.FlattenString(m["type"])
	r.Mode = dcl.FlattenString(m["mode"])
	r.Fields = flattenTableGooglecloudbigqueryv2TablefieldschemaSlice(c, m["fields"], res)
	r.Description = dcl.FlattenString(m["description"])
	r.Categories = flattenTableGooglecloudbigqueryv2TablefieldschemaCategories(c, m["categories"], res)
	r.PolicyTags = flattenTableGooglecloudbigqueryv2TablefieldschemaPolicyTags(c, m["policyTags"], res)
	r.NameAlternative = dcl.FlattenStringSlice(m["nameAlternative"])
	r.MaxLength = dcl.FlattenInteger(m["maxLength"])
	r.Precision = dcl.FlattenInteger(m["precision"])
	r.Scale = dcl.FlattenInteger(m["scale"])
	r.Collation = dcl.FlattenString(m["collation"])
	r.DefaultValueExpression = dcl.FlattenString(m["defaultValueExpression"])

	return r
}

// expandTableGooglecloudbigqueryv2TablefieldschemaCategoriesMap expands the contents of TableGooglecloudbigqueryv2TablefieldschemaCategories into a JSON
// request object.
func expandTableGooglecloudbigqueryv2TablefieldschemaCategoriesMap(c *Client, f map[string]TableGooglecloudbigqueryv2TablefieldschemaCategories, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableGooglecloudbigqueryv2TablefieldschemaCategories(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableGooglecloudbigqueryv2TablefieldschemaCategoriesSlice expands the contents of TableGooglecloudbigqueryv2TablefieldschemaCategories into a JSON
// request object.
func expandTableGooglecloudbigqueryv2TablefieldschemaCategoriesSlice(c *Client, f []TableGooglecloudbigqueryv2TablefieldschemaCategories, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableGooglecloudbigqueryv2TablefieldschemaCategories(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableGooglecloudbigqueryv2TablefieldschemaCategoriesMap flattens the contents of TableGooglecloudbigqueryv2TablefieldschemaCategories from a JSON
// response object.
func flattenTableGooglecloudbigqueryv2TablefieldschemaCategoriesMap(c *Client, i interface{}, res *Table) map[string]TableGooglecloudbigqueryv2TablefieldschemaCategories {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableGooglecloudbigqueryv2TablefieldschemaCategories{}
	}

	if len(a) == 0 {
		return map[string]TableGooglecloudbigqueryv2TablefieldschemaCategories{}
	}

	items := make(map[string]TableGooglecloudbigqueryv2TablefieldschemaCategories)
	for k, item := range a {
		items[k] = *flattenTableGooglecloudbigqueryv2TablefieldschemaCategories(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableGooglecloudbigqueryv2TablefieldschemaCategoriesSlice flattens the contents of TableGooglecloudbigqueryv2TablefieldschemaCategories from a JSON
// response object.
func flattenTableGooglecloudbigqueryv2TablefieldschemaCategoriesSlice(c *Client, i interface{}, res *Table) []TableGooglecloudbigqueryv2TablefieldschemaCategories {
	a, ok := i.([]interface{})
	if !ok {
		return []TableGooglecloudbigqueryv2TablefieldschemaCategories{}
	}

	if len(a) == 0 {
		return []TableGooglecloudbigqueryv2TablefieldschemaCategories{}
	}

	items := make([]TableGooglecloudbigqueryv2TablefieldschemaCategories, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableGooglecloudbigqueryv2TablefieldschemaCategories(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableGooglecloudbigqueryv2TablefieldschemaCategories expands an instance of TableGooglecloudbigqueryv2TablefieldschemaCategories into a JSON
// request object.
func expandTableGooglecloudbigqueryv2TablefieldschemaCategories(c *Client, f *TableGooglecloudbigqueryv2TablefieldschemaCategories, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Names; v != nil {
		m["names"] = v
	}

	return m, nil
}

// flattenTableGooglecloudbigqueryv2TablefieldschemaCategories flattens an instance of TableGooglecloudbigqueryv2TablefieldschemaCategories from a JSON
// response object.
func flattenTableGooglecloudbigqueryv2TablefieldschemaCategories(c *Client, i interface{}, res *Table) *TableGooglecloudbigqueryv2TablefieldschemaCategories {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableGooglecloudbigqueryv2TablefieldschemaCategories{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableGooglecloudbigqueryv2TablefieldschemaCategories
	}
	r.Names = dcl.FlattenStringSlice(m["names"])

	return r
}

// expandTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsMap expands the contents of TableGooglecloudbigqueryv2TablefieldschemaPolicyTags into a JSON
// request object.
func expandTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsMap(c *Client, f map[string]TableGooglecloudbigqueryv2TablefieldschemaPolicyTags, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableGooglecloudbigqueryv2TablefieldschemaPolicyTags(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsSlice expands the contents of TableGooglecloudbigqueryv2TablefieldschemaPolicyTags into a JSON
// request object.
func expandTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsSlice(c *Client, f []TableGooglecloudbigqueryv2TablefieldschemaPolicyTags, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableGooglecloudbigqueryv2TablefieldschemaPolicyTags(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsMap flattens the contents of TableGooglecloudbigqueryv2TablefieldschemaPolicyTags from a JSON
// response object.
func flattenTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsMap(c *Client, i interface{}, res *Table) map[string]TableGooglecloudbigqueryv2TablefieldschemaPolicyTags {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableGooglecloudbigqueryv2TablefieldschemaPolicyTags{}
	}

	if len(a) == 0 {
		return map[string]TableGooglecloudbigqueryv2TablefieldschemaPolicyTags{}
	}

	items := make(map[string]TableGooglecloudbigqueryv2TablefieldschemaPolicyTags)
	for k, item := range a {
		items[k] = *flattenTableGooglecloudbigqueryv2TablefieldschemaPolicyTags(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsSlice flattens the contents of TableGooglecloudbigqueryv2TablefieldschemaPolicyTags from a JSON
// response object.
func flattenTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsSlice(c *Client, i interface{}, res *Table) []TableGooglecloudbigqueryv2TablefieldschemaPolicyTags {
	a, ok := i.([]interface{})
	if !ok {
		return []TableGooglecloudbigqueryv2TablefieldschemaPolicyTags{}
	}

	if len(a) == 0 {
		return []TableGooglecloudbigqueryv2TablefieldschemaPolicyTags{}
	}

	items := make([]TableGooglecloudbigqueryv2TablefieldschemaPolicyTags, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableGooglecloudbigqueryv2TablefieldschemaPolicyTags(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableGooglecloudbigqueryv2TablefieldschemaPolicyTags expands an instance of TableGooglecloudbigqueryv2TablefieldschemaPolicyTags into a JSON
// request object.
func expandTableGooglecloudbigqueryv2TablefieldschemaPolicyTags(c *Client, f *TableGooglecloudbigqueryv2TablefieldschemaPolicyTags, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Names; v != nil {
		m["names"] = v
	}

	return m, nil
}

// flattenTableGooglecloudbigqueryv2TablefieldschemaPolicyTags flattens an instance of TableGooglecloudbigqueryv2TablefieldschemaPolicyTags from a JSON
// response object.
func flattenTableGooglecloudbigqueryv2TablefieldschemaPolicyTags(c *Client, i interface{}, res *Table) *TableGooglecloudbigqueryv2TablefieldschemaPolicyTags {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableGooglecloudbigqueryv2TablefieldschemaPolicyTags{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableGooglecloudbigqueryv2TablefieldschemaPolicyTags
	}
	r.Names = dcl.FlattenStringSlice(m["names"])

	return r
}

// expandTableTimePartitioningMap expands the contents of TableTimePartitioning into a JSON
// request object.
func expandTableTimePartitioningMap(c *Client, f map[string]TableTimePartitioning, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableTimePartitioning(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableTimePartitioningSlice expands the contents of TableTimePartitioning into a JSON
// request object.
func expandTableTimePartitioningSlice(c *Client, f []TableTimePartitioning, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableTimePartitioning(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableTimePartitioningMap flattens the contents of TableTimePartitioning from a JSON
// response object.
func flattenTableTimePartitioningMap(c *Client, i interface{}, res *Table) map[string]TableTimePartitioning {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableTimePartitioning{}
	}

	if len(a) == 0 {
		return map[string]TableTimePartitioning{}
	}

	items := make(map[string]TableTimePartitioning)
	for k, item := range a {
		items[k] = *flattenTableTimePartitioning(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableTimePartitioningSlice flattens the contents of TableTimePartitioning from a JSON
// response object.
func flattenTableTimePartitioningSlice(c *Client, i interface{}, res *Table) []TableTimePartitioning {
	a, ok := i.([]interface{})
	if !ok {
		return []TableTimePartitioning{}
	}

	if len(a) == 0 {
		return []TableTimePartitioning{}
	}

	items := make([]TableTimePartitioning, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableTimePartitioning(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableTimePartitioning expands an instance of TableTimePartitioning into a JSON
// request object.
func expandTableTimePartitioning(c *Client, f *TableTimePartitioning, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.ExpirationMs; !dcl.IsEmptyValueIndirect(v) {
		m["expirationMs"] = v
	}
	if v := f.Field; !dcl.IsEmptyValueIndirect(v) {
		m["field"] = v
	}

	return m, nil
}

// flattenTableTimePartitioning flattens an instance of TableTimePartitioning from a JSON
// response object.
func flattenTableTimePartitioning(c *Client, i interface{}, res *Table) *TableTimePartitioning {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableTimePartitioning{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableTimePartitioning
	}
	r.Type = dcl.FlattenString(m["type"])
	r.ExpirationMs = dcl.FlattenString(m["expirationMs"])
	r.Field = dcl.FlattenString(m["field"])

	return r
}

// expandTableRangePartitioningMap expands the contents of TableRangePartitioning into a JSON
// request object.
func expandTableRangePartitioningMap(c *Client, f map[string]TableRangePartitioning, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableRangePartitioning(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableRangePartitioningSlice expands the contents of TableRangePartitioning into a JSON
// request object.
func expandTableRangePartitioningSlice(c *Client, f []TableRangePartitioning, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableRangePartitioning(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableRangePartitioningMap flattens the contents of TableRangePartitioning from a JSON
// response object.
func flattenTableRangePartitioningMap(c *Client, i interface{}, res *Table) map[string]TableRangePartitioning {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableRangePartitioning{}
	}

	if len(a) == 0 {
		return map[string]TableRangePartitioning{}
	}

	items := make(map[string]TableRangePartitioning)
	for k, item := range a {
		items[k] = *flattenTableRangePartitioning(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableRangePartitioningSlice flattens the contents of TableRangePartitioning from a JSON
// response object.
func flattenTableRangePartitioningSlice(c *Client, i interface{}, res *Table) []TableRangePartitioning {
	a, ok := i.([]interface{})
	if !ok {
		return []TableRangePartitioning{}
	}

	if len(a) == 0 {
		return []TableRangePartitioning{}
	}

	items := make([]TableRangePartitioning, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableRangePartitioning(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableRangePartitioning expands an instance of TableRangePartitioning into a JSON
// request object.
func expandTableRangePartitioning(c *Client, f *TableRangePartitioning, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Field; !dcl.IsEmptyValueIndirect(v) {
		m["field"] = v
	}
	if v, err := expandTableRangePartitioningRange(c, f.Range, res); err != nil {
		return nil, fmt.Errorf("error expanding Range into range: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["range"] = v
	}

	return m, nil
}

// flattenTableRangePartitioning flattens an instance of TableRangePartitioning from a JSON
// response object.
func flattenTableRangePartitioning(c *Client, i interface{}, res *Table) *TableRangePartitioning {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableRangePartitioning{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableRangePartitioning
	}
	r.Field = dcl.FlattenString(m["field"])
	r.Range = flattenTableRangePartitioningRange(c, m["range"], res)

	return r
}

// expandTableRangePartitioningRangeMap expands the contents of TableRangePartitioningRange into a JSON
// request object.
func expandTableRangePartitioningRangeMap(c *Client, f map[string]TableRangePartitioningRange, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableRangePartitioningRange(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableRangePartitioningRangeSlice expands the contents of TableRangePartitioningRange into a JSON
// request object.
func expandTableRangePartitioningRangeSlice(c *Client, f []TableRangePartitioningRange, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableRangePartitioningRange(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableRangePartitioningRangeMap flattens the contents of TableRangePartitioningRange from a JSON
// response object.
func flattenTableRangePartitioningRangeMap(c *Client, i interface{}, res *Table) map[string]TableRangePartitioningRange {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableRangePartitioningRange{}
	}

	if len(a) == 0 {
		return map[string]TableRangePartitioningRange{}
	}

	items := make(map[string]TableRangePartitioningRange)
	for k, item := range a {
		items[k] = *flattenTableRangePartitioningRange(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableRangePartitioningRangeSlice flattens the contents of TableRangePartitioningRange from a JSON
// response object.
func flattenTableRangePartitioningRangeSlice(c *Client, i interface{}, res *Table) []TableRangePartitioningRange {
	a, ok := i.([]interface{})
	if !ok {
		return []TableRangePartitioningRange{}
	}

	if len(a) == 0 {
		return []TableRangePartitioningRange{}
	}

	items := make([]TableRangePartitioningRange, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableRangePartitioningRange(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableRangePartitioningRange expands an instance of TableRangePartitioningRange into a JSON
// request object.
func expandTableRangePartitioningRange(c *Client, f *TableRangePartitioningRange, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Start; !dcl.IsEmptyValueIndirect(v) {
		m["start"] = v
	}
	if v := f.End; !dcl.IsEmptyValueIndirect(v) {
		m["end"] = v
	}
	if v := f.Interval; !dcl.IsEmptyValueIndirect(v) {
		m["interval"] = v
	}

	return m, nil
}

// flattenTableRangePartitioningRange flattens an instance of TableRangePartitioningRange from a JSON
// response object.
func flattenTableRangePartitioningRange(c *Client, i interface{}, res *Table) *TableRangePartitioningRange {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableRangePartitioningRange{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableRangePartitioningRange
	}
	r.Start = dcl.FlattenString(m["start"])
	r.End = dcl.FlattenString(m["end"])
	r.Interval = dcl.FlattenString(m["interval"])

	return r
}

// expandTableClusteringMap expands the contents of TableClustering into a JSON
// request object.
func expandTableClusteringMap(c *Client, f map[string]TableClustering, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableClustering(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableClusteringSlice expands the contents of TableClustering into a JSON
// request object.
func expandTableClusteringSlice(c *Client, f []TableClustering, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableClustering(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableClusteringMap flattens the contents of TableClustering from a JSON
// response object.
func flattenTableClusteringMap(c *Client, i interface{}, res *Table) map[string]TableClustering {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableClustering{}
	}

	if len(a) == 0 {
		return map[string]TableClustering{}
	}

	items := make(map[string]TableClustering)
	for k, item := range a {
		items[k] = *flattenTableClustering(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableClusteringSlice flattens the contents of TableClustering from a JSON
// response object.
func flattenTableClusteringSlice(c *Client, i interface{}, res *Table) []TableClustering {
	a, ok := i.([]interface{})
	if !ok {
		return []TableClustering{}
	}

	if len(a) == 0 {
		return []TableClustering{}
	}

	items := make([]TableClustering, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableClustering(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableClustering expands an instance of TableClustering into a JSON
// request object.
func expandTableClustering(c *Client, f *TableClustering, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Fields; v != nil {
		m["fields"] = v
	}

	return m, nil
}

// flattenTableClustering flattens an instance of TableClustering from a JSON
// response object.
func flattenTableClustering(c *Client, i interface{}, res *Table) *TableClustering {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableClustering{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableClustering
	}
	r.Fields = dcl.FlattenStringSlice(m["fields"])

	return r
}

// expandTableViewMap expands the contents of TableView into a JSON
// request object.
func expandTableViewMap(c *Client, f map[string]TableView, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableView(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableViewSlice expands the contents of TableView into a JSON
// request object.
func expandTableViewSlice(c *Client, f []TableView, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableView(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableViewMap flattens the contents of TableView from a JSON
// response object.
func flattenTableViewMap(c *Client, i interface{}, res *Table) map[string]TableView {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableView{}
	}

	if len(a) == 0 {
		return map[string]TableView{}
	}

	items := make(map[string]TableView)
	for k, item := range a {
		items[k] = *flattenTableView(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableViewSlice flattens the contents of TableView from a JSON
// response object.
func flattenTableViewSlice(c *Client, i interface{}, res *Table) []TableView {
	a, ok := i.([]interface{})
	if !ok {
		return []TableView{}
	}

	if len(a) == 0 {
		return []TableView{}
	}

	items := make([]TableView, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableView(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableView expands an instance of TableView into a JSON
// request object.
func expandTableView(c *Client, f *TableView, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Query; !dcl.IsEmptyValueIndirect(v) {
		m["query"] = v
	}
	if v, err := expandTableViewUserDefinedFunctionResourcesSlice(c, f.UserDefinedFunctionResources, res); err != nil {
		return nil, fmt.Errorf("error expanding UserDefinedFunctionResources into userDefinedFunctionResources: %w", err)
	} else if v != nil {
		m["userDefinedFunctionResources"] = v
	}
	if v := f.UseLegacySql; !dcl.IsEmptyValueIndirect(v) {
		m["useLegacySql"] = v
	}
	if v := f.UseExplicitColumnNames; !dcl.IsEmptyValueIndirect(v) {
		m["useExplicitColumnNames"] = v
	}

	return m, nil
}

// flattenTableView flattens an instance of TableView from a JSON
// response object.
func flattenTableView(c *Client, i interface{}, res *Table) *TableView {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableView{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableView
	}
	r.Query = dcl.FlattenString(m["query"])
	r.UserDefinedFunctionResources = flattenTableViewUserDefinedFunctionResourcesSlice(c, m["userDefinedFunctionResources"], res)
	r.UseLegacySql = dcl.FlattenBool(m["useLegacySql"])
	r.UseExplicitColumnNames = dcl.FlattenBool(m["useExplicitColumnNames"])

	return r
}

// expandTableViewUserDefinedFunctionResourcesMap expands the contents of TableViewUserDefinedFunctionResources into a JSON
// request object.
func expandTableViewUserDefinedFunctionResourcesMap(c *Client, f map[string]TableViewUserDefinedFunctionResources, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableViewUserDefinedFunctionResources(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableViewUserDefinedFunctionResourcesSlice expands the contents of TableViewUserDefinedFunctionResources into a JSON
// request object.
func expandTableViewUserDefinedFunctionResourcesSlice(c *Client, f []TableViewUserDefinedFunctionResources, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableViewUserDefinedFunctionResources(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableViewUserDefinedFunctionResourcesMap flattens the contents of TableViewUserDefinedFunctionResources from a JSON
// response object.
func flattenTableViewUserDefinedFunctionResourcesMap(c *Client, i interface{}, res *Table) map[string]TableViewUserDefinedFunctionResources {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableViewUserDefinedFunctionResources{}
	}

	if len(a) == 0 {
		return map[string]TableViewUserDefinedFunctionResources{}
	}

	items := make(map[string]TableViewUserDefinedFunctionResources)
	for k, item := range a {
		items[k] = *flattenTableViewUserDefinedFunctionResources(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableViewUserDefinedFunctionResourcesSlice flattens the contents of TableViewUserDefinedFunctionResources from a JSON
// response object.
func flattenTableViewUserDefinedFunctionResourcesSlice(c *Client, i interface{}, res *Table) []TableViewUserDefinedFunctionResources {
	a, ok := i.([]interface{})
	if !ok {
		return []TableViewUserDefinedFunctionResources{}
	}

	if len(a) == 0 {
		return []TableViewUserDefinedFunctionResources{}
	}

	items := make([]TableViewUserDefinedFunctionResources, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableViewUserDefinedFunctionResources(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableViewUserDefinedFunctionResources expands an instance of TableViewUserDefinedFunctionResources into a JSON
// request object.
func expandTableViewUserDefinedFunctionResources(c *Client, f *TableViewUserDefinedFunctionResources, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ResourceUri; !dcl.IsEmptyValueIndirect(v) {
		m["resourceUri"] = v
	}
	if v := f.InlineCode; !dcl.IsEmptyValueIndirect(v) {
		m["inlineCode"] = v
	}
	if v := f.InlineCodeAlternative; v != nil {
		m["inlineCodeAlternative"] = v
	}

	return m, nil
}

// flattenTableViewUserDefinedFunctionResources flattens an instance of TableViewUserDefinedFunctionResources from a JSON
// response object.
func flattenTableViewUserDefinedFunctionResources(c *Client, i interface{}, res *Table) *TableViewUserDefinedFunctionResources {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableViewUserDefinedFunctionResources{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableViewUserDefinedFunctionResources
	}
	r.ResourceUri = dcl.FlattenString(m["resourceUri"])
	r.InlineCode = dcl.FlattenString(m["inlineCode"])
	r.InlineCodeAlternative = dcl.FlattenStringSlice(m["inlineCodeAlternative"])

	return r
}

// expandTableMaterializedViewMap expands the contents of TableMaterializedView into a JSON
// request object.
func expandTableMaterializedViewMap(c *Client, f map[string]TableMaterializedView, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableMaterializedView(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableMaterializedViewSlice expands the contents of TableMaterializedView into a JSON
// request object.
func expandTableMaterializedViewSlice(c *Client, f []TableMaterializedView, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableMaterializedView(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableMaterializedViewMap flattens the contents of TableMaterializedView from a JSON
// response object.
func flattenTableMaterializedViewMap(c *Client, i interface{}, res *Table) map[string]TableMaterializedView {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableMaterializedView{}
	}

	if len(a) == 0 {
		return map[string]TableMaterializedView{}
	}

	items := make(map[string]TableMaterializedView)
	for k, item := range a {
		items[k] = *flattenTableMaterializedView(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableMaterializedViewSlice flattens the contents of TableMaterializedView from a JSON
// response object.
func flattenTableMaterializedViewSlice(c *Client, i interface{}, res *Table) []TableMaterializedView {
	a, ok := i.([]interface{})
	if !ok {
		return []TableMaterializedView{}
	}

	if len(a) == 0 {
		return []TableMaterializedView{}
	}

	items := make([]TableMaterializedView, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableMaterializedView(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableMaterializedView expands an instance of TableMaterializedView into a JSON
// request object.
func expandTableMaterializedView(c *Client, f *TableMaterializedView, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Query; !dcl.IsEmptyValueIndirect(v) {
		m["query"] = v
	}
	if v := f.EnableRefresh; !dcl.IsEmptyValueIndirect(v) {
		m["enableRefresh"] = v
	}
	if v := f.RefreshIntervalMs; !dcl.IsEmptyValueIndirect(v) {
		m["refreshIntervalMs"] = v
	}

	return m, nil
}

// flattenTableMaterializedView flattens an instance of TableMaterializedView from a JSON
// response object.
func flattenTableMaterializedView(c *Client, i interface{}, res *Table) *TableMaterializedView {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableMaterializedView{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableMaterializedView
	}
	r.Query = dcl.FlattenString(m["query"])
	r.LastRefreshTime = dcl.FlattenInteger(m["lastRefreshTime"])
	r.EnableRefresh = dcl.FlattenBool(m["enableRefresh"])
	r.RefreshIntervalMs = dcl.FlattenInteger(m["refreshIntervalMs"])

	return r
}

// expandTableExternalDataConfigurationMap expands the contents of TableExternalDataConfiguration into a JSON
// request object.
func expandTableExternalDataConfigurationMap(c *Client, f map[string]TableExternalDataConfiguration, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableExternalDataConfiguration(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableExternalDataConfigurationSlice expands the contents of TableExternalDataConfiguration into a JSON
// request object.
func expandTableExternalDataConfigurationSlice(c *Client, f []TableExternalDataConfiguration, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableExternalDataConfiguration(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableExternalDataConfigurationMap flattens the contents of TableExternalDataConfiguration from a JSON
// response object.
func flattenTableExternalDataConfigurationMap(c *Client, i interface{}, res *Table) map[string]TableExternalDataConfiguration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableExternalDataConfiguration{}
	}

	if len(a) == 0 {
		return map[string]TableExternalDataConfiguration{}
	}

	items := make(map[string]TableExternalDataConfiguration)
	for k, item := range a {
		items[k] = *flattenTableExternalDataConfiguration(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableExternalDataConfigurationSlice flattens the contents of TableExternalDataConfiguration from a JSON
// response object.
func flattenTableExternalDataConfigurationSlice(c *Client, i interface{}, res *Table) []TableExternalDataConfiguration {
	a, ok := i.([]interface{})
	if !ok {
		return []TableExternalDataConfiguration{}
	}

	if len(a) == 0 {
		return []TableExternalDataConfiguration{}
	}

	items := make([]TableExternalDataConfiguration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableExternalDataConfiguration(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableExternalDataConfiguration expands an instance of TableExternalDataConfiguration into a JSON
// request object.
func expandTableExternalDataConfiguration(c *Client, f *TableExternalDataConfiguration, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SourceUris; v != nil {
		m["sourceUris"] = v
	}
	if v, err := expandTableExternalDataConfigurationSchema(c, f.Schema, res); err != nil {
		return nil, fmt.Errorf("error expanding Schema into schema: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["schema"] = v
	}
	if v := f.SourceFormat; !dcl.IsEmptyValueIndirect(v) {
		m["sourceFormat"] = v
	}
	if v := f.MaxBadRecords; !dcl.IsEmptyValueIndirect(v) {
		m["maxBadRecords"] = v
	}
	if v := f.Autodetect; !dcl.IsEmptyValueIndirect(v) {
		m["autodetect"] = v
	}
	if v := f.IgnoreUnknownValues; !dcl.IsEmptyValueIndirect(v) {
		m["ignoreUnknownValues"] = v
	}
	if v := f.Compression; !dcl.IsEmptyValueIndirect(v) {
		m["compression"] = v
	}
	if v, err := expandTableExternalDataConfigurationCsvOptions(c, f.CsvOptions, res); err != nil {
		return nil, fmt.Errorf("error expanding CsvOptions into csvOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["csvOptions"] = v
	}
	if v, err := expandTableExternalDataConfigurationBigtableOptions(c, f.BigtableOptions, res); err != nil {
		return nil, fmt.Errorf("error expanding BigtableOptions into bigtableOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["bigtableOptions"] = v
	}
	if v, err := expandTableExternalDataConfigurationGoogleSheetsOptions(c, f.GoogleSheetsOptions, res); err != nil {
		return nil, fmt.Errorf("error expanding GoogleSheetsOptions into googleSheetsOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["googleSheetsOptions"] = v
	}
	if v := f.MaxBadRecordsAlternative; v != nil {
		m["maxBadRecordsAlternative"] = v
	}
	if v, err := expandTableExternalDataConfigurationHivePartitioningOptions(c, f.HivePartitioningOptions, res); err != nil {
		return nil, fmt.Errorf("error expanding HivePartitioningOptions into hivePartitioningOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["hivePartitioningOptions"] = v
	}
	if v := f.ConnectionId; !dcl.IsEmptyValueIndirect(v) {
		m["connectionId"] = v
	}
	if v, err := expandTableExternalDataConfigurationValueConversionModes(c, f.ValueConversionModes, res); err != nil {
		return nil, fmt.Errorf("error expanding ValueConversionModes into valueConversionModes: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["valueConversionModes"] = v
	}
	if v := f.DecimalTargetTypes; v != nil {
		m["decimalTargetTypes"] = v
	}
	if v, err := expandTableExternalDataConfigurationAvroOptions(c, f.AvroOptions, res); err != nil {
		return nil, fmt.Errorf("error expanding AvroOptions into avroOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["avroOptions"] = v
	}
	if v := f.JsonExtension; !dcl.IsEmptyValueIndirect(v) {
		m["jsonExtension"] = v
	}
	if v, err := expandTableExternalDataConfigurationParquetOptions(c, f.ParquetOptions, res); err != nil {
		return nil, fmt.Errorf("error expanding ParquetOptions into parquetOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["parquetOptions"] = v
	}

	return m, nil
}

// flattenTableExternalDataConfiguration flattens an instance of TableExternalDataConfiguration from a JSON
// response object.
func flattenTableExternalDataConfiguration(c *Client, i interface{}, res *Table) *TableExternalDataConfiguration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableExternalDataConfiguration{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableExternalDataConfiguration
	}
	r.SourceUris = dcl.FlattenStringSlice(m["sourceUris"])
	r.Schema = flattenTableExternalDataConfigurationSchema(c, m["schema"], res)
	r.SourceFormat = dcl.FlattenString(m["sourceFormat"])
	r.MaxBadRecords = dcl.FlattenInteger(m["maxBadRecords"])
	r.Autodetect = dcl.FlattenBool(m["autodetect"])
	r.IgnoreUnknownValues = dcl.FlattenBool(m["ignoreUnknownValues"])
	r.Compression = dcl.FlattenString(m["compression"])
	r.CsvOptions = flattenTableExternalDataConfigurationCsvOptions(c, m["csvOptions"], res)
	r.BigtableOptions = flattenTableExternalDataConfigurationBigtableOptions(c, m["bigtableOptions"], res)
	r.GoogleSheetsOptions = flattenTableExternalDataConfigurationGoogleSheetsOptions(c, m["googleSheetsOptions"], res)
	r.MaxBadRecordsAlternative = dcl.FlattenIntSlice(m["maxBadRecordsAlternative"])
	r.HivePartitioningOptions = flattenTableExternalDataConfigurationHivePartitioningOptions(c, m["hivePartitioningOptions"], res)
	r.ConnectionId = dcl.FlattenString(m["connectionId"])
	r.ValueConversionModes = flattenTableExternalDataConfigurationValueConversionModes(c, m["valueConversionModes"], res)
	r.DecimalTargetTypes = flattenTableExternalDataConfigurationDecimalTargetTypesEnumSlice(c, m["decimalTargetTypes"], res)
	r.AvroOptions = flattenTableExternalDataConfigurationAvroOptions(c, m["avroOptions"], res)
	r.JsonExtension = flattenTableExternalDataConfigurationJsonExtensionEnum(m["jsonExtension"])
	r.ParquetOptions = flattenTableExternalDataConfigurationParquetOptions(c, m["parquetOptions"], res)

	return r
}

// expandTableExternalDataConfigurationSchemaMap expands the contents of TableExternalDataConfigurationSchema into a JSON
// request object.
func expandTableExternalDataConfigurationSchemaMap(c *Client, f map[string]TableExternalDataConfigurationSchema, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableExternalDataConfigurationSchema(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableExternalDataConfigurationSchemaSlice expands the contents of TableExternalDataConfigurationSchema into a JSON
// request object.
func expandTableExternalDataConfigurationSchemaSlice(c *Client, f []TableExternalDataConfigurationSchema, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableExternalDataConfigurationSchema(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableExternalDataConfigurationSchemaMap flattens the contents of TableExternalDataConfigurationSchema from a JSON
// response object.
func flattenTableExternalDataConfigurationSchemaMap(c *Client, i interface{}, res *Table) map[string]TableExternalDataConfigurationSchema {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableExternalDataConfigurationSchema{}
	}

	if len(a) == 0 {
		return map[string]TableExternalDataConfigurationSchema{}
	}

	items := make(map[string]TableExternalDataConfigurationSchema)
	for k, item := range a {
		items[k] = *flattenTableExternalDataConfigurationSchema(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableExternalDataConfigurationSchemaSlice flattens the contents of TableExternalDataConfigurationSchema from a JSON
// response object.
func flattenTableExternalDataConfigurationSchemaSlice(c *Client, i interface{}, res *Table) []TableExternalDataConfigurationSchema {
	a, ok := i.([]interface{})
	if !ok {
		return []TableExternalDataConfigurationSchema{}
	}

	if len(a) == 0 {
		return []TableExternalDataConfigurationSchema{}
	}

	items := make([]TableExternalDataConfigurationSchema, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableExternalDataConfigurationSchema(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableExternalDataConfigurationSchema expands an instance of TableExternalDataConfigurationSchema into a JSON
// request object.
func expandTableExternalDataConfigurationSchema(c *Client, f *TableExternalDataConfigurationSchema, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Fields; v != nil {
		m["fields"] = v
	}

	return m, nil
}

// flattenTableExternalDataConfigurationSchema flattens an instance of TableExternalDataConfigurationSchema from a JSON
// response object.
func flattenTableExternalDataConfigurationSchema(c *Client, i interface{}, res *Table) *TableExternalDataConfigurationSchema {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableExternalDataConfigurationSchema{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableExternalDataConfigurationSchema
	}
	r.Fields = flattenTableGooglecloudbigqueryv2TablefieldschemaSlice(c, m["fields"], res)

	return r
}

// expandTableExternalDataConfigurationCsvOptionsMap expands the contents of TableExternalDataConfigurationCsvOptions into a JSON
// request object.
func expandTableExternalDataConfigurationCsvOptionsMap(c *Client, f map[string]TableExternalDataConfigurationCsvOptions, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableExternalDataConfigurationCsvOptions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableExternalDataConfigurationCsvOptionsSlice expands the contents of TableExternalDataConfigurationCsvOptions into a JSON
// request object.
func expandTableExternalDataConfigurationCsvOptionsSlice(c *Client, f []TableExternalDataConfigurationCsvOptions, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableExternalDataConfigurationCsvOptions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableExternalDataConfigurationCsvOptionsMap flattens the contents of TableExternalDataConfigurationCsvOptions from a JSON
// response object.
func flattenTableExternalDataConfigurationCsvOptionsMap(c *Client, i interface{}, res *Table) map[string]TableExternalDataConfigurationCsvOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableExternalDataConfigurationCsvOptions{}
	}

	if len(a) == 0 {
		return map[string]TableExternalDataConfigurationCsvOptions{}
	}

	items := make(map[string]TableExternalDataConfigurationCsvOptions)
	for k, item := range a {
		items[k] = *flattenTableExternalDataConfigurationCsvOptions(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableExternalDataConfigurationCsvOptionsSlice flattens the contents of TableExternalDataConfigurationCsvOptions from a JSON
// response object.
func flattenTableExternalDataConfigurationCsvOptionsSlice(c *Client, i interface{}, res *Table) []TableExternalDataConfigurationCsvOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []TableExternalDataConfigurationCsvOptions{}
	}

	if len(a) == 0 {
		return []TableExternalDataConfigurationCsvOptions{}
	}

	items := make([]TableExternalDataConfigurationCsvOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableExternalDataConfigurationCsvOptions(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableExternalDataConfigurationCsvOptions expands an instance of TableExternalDataConfigurationCsvOptions into a JSON
// request object.
func expandTableExternalDataConfigurationCsvOptions(c *Client, f *TableExternalDataConfigurationCsvOptions, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.FieldDelimiter; !dcl.IsEmptyValueIndirect(v) {
		m["fieldDelimiter"] = v
	}
	if v := f.SkipLeadingRows; !dcl.IsEmptyValueIndirect(v) {
		m["skipLeadingRows"] = v
	}
	if v := f.Quote; !dcl.IsEmptyValueIndirect(v) {
		m["quote"] = v
	}
	if v := f.AllowQuotedNewlines; !dcl.IsEmptyValueIndirect(v) {
		m["allowQuotedNewlines"] = v
	}
	if v := f.AllowJaggedRows; !dcl.IsEmptyValueIndirect(v) {
		m["allowJaggedRows"] = v
	}
	if v := f.Encoding; !dcl.IsEmptyValueIndirect(v) {
		m["encoding"] = v
	}

	return m, nil
}

// flattenTableExternalDataConfigurationCsvOptions flattens an instance of TableExternalDataConfigurationCsvOptions from a JSON
// response object.
func flattenTableExternalDataConfigurationCsvOptions(c *Client, i interface{}, res *Table) *TableExternalDataConfigurationCsvOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableExternalDataConfigurationCsvOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableExternalDataConfigurationCsvOptions
	}
	r.FieldDelimiter = dcl.FlattenString(m["fieldDelimiter"])
	r.SkipLeadingRows = dcl.FlattenString(m["skipLeadingRows"])
	r.Quote = dcl.FlattenString(m["quote"])
	r.AllowQuotedNewlines = dcl.FlattenBool(m["allowQuotedNewlines"])
	r.AllowJaggedRows = dcl.FlattenBool(m["allowJaggedRows"])
	r.Encoding = dcl.FlattenString(m["encoding"])

	return r
}

// expandTableExternalDataConfigurationBigtableOptionsMap expands the contents of TableExternalDataConfigurationBigtableOptions into a JSON
// request object.
func expandTableExternalDataConfigurationBigtableOptionsMap(c *Client, f map[string]TableExternalDataConfigurationBigtableOptions, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableExternalDataConfigurationBigtableOptions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableExternalDataConfigurationBigtableOptionsSlice expands the contents of TableExternalDataConfigurationBigtableOptions into a JSON
// request object.
func expandTableExternalDataConfigurationBigtableOptionsSlice(c *Client, f []TableExternalDataConfigurationBigtableOptions, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableExternalDataConfigurationBigtableOptions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableExternalDataConfigurationBigtableOptionsMap flattens the contents of TableExternalDataConfigurationBigtableOptions from a JSON
// response object.
func flattenTableExternalDataConfigurationBigtableOptionsMap(c *Client, i interface{}, res *Table) map[string]TableExternalDataConfigurationBigtableOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableExternalDataConfigurationBigtableOptions{}
	}

	if len(a) == 0 {
		return map[string]TableExternalDataConfigurationBigtableOptions{}
	}

	items := make(map[string]TableExternalDataConfigurationBigtableOptions)
	for k, item := range a {
		items[k] = *flattenTableExternalDataConfigurationBigtableOptions(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableExternalDataConfigurationBigtableOptionsSlice flattens the contents of TableExternalDataConfigurationBigtableOptions from a JSON
// response object.
func flattenTableExternalDataConfigurationBigtableOptionsSlice(c *Client, i interface{}, res *Table) []TableExternalDataConfigurationBigtableOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []TableExternalDataConfigurationBigtableOptions{}
	}

	if len(a) == 0 {
		return []TableExternalDataConfigurationBigtableOptions{}
	}

	items := make([]TableExternalDataConfigurationBigtableOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableExternalDataConfigurationBigtableOptions(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableExternalDataConfigurationBigtableOptions expands an instance of TableExternalDataConfigurationBigtableOptions into a JSON
// request object.
func expandTableExternalDataConfigurationBigtableOptions(c *Client, f *TableExternalDataConfigurationBigtableOptions, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandTableExternalDataConfigurationBigtableOptionsColumnFamiliesSlice(c, f.ColumnFamilies, res); err != nil {
		return nil, fmt.Errorf("error expanding ColumnFamilies into columnFamilies: %w", err)
	} else if v != nil {
		m["columnFamilies"] = v
	}
	if v := f.IgnoreUnspecifiedColumnFamilies; !dcl.IsEmptyValueIndirect(v) {
		m["ignoreUnspecifiedColumnFamilies"] = v
	}
	if v := f.ReadRowkeyAsString; !dcl.IsEmptyValueIndirect(v) {
		m["readRowkeyAsString"] = v
	}

	return m, nil
}

// flattenTableExternalDataConfigurationBigtableOptions flattens an instance of TableExternalDataConfigurationBigtableOptions from a JSON
// response object.
func flattenTableExternalDataConfigurationBigtableOptions(c *Client, i interface{}, res *Table) *TableExternalDataConfigurationBigtableOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableExternalDataConfigurationBigtableOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableExternalDataConfigurationBigtableOptions
	}
	r.ColumnFamilies = flattenTableExternalDataConfigurationBigtableOptionsColumnFamiliesSlice(c, m["columnFamilies"], res)
	r.IgnoreUnspecifiedColumnFamilies = dcl.FlattenBool(m["ignoreUnspecifiedColumnFamilies"])
	r.ReadRowkeyAsString = dcl.FlattenBool(m["readRowkeyAsString"])

	return r
}

// expandTableExternalDataConfigurationBigtableOptionsColumnFamiliesMap expands the contents of TableExternalDataConfigurationBigtableOptionsColumnFamilies into a JSON
// request object.
func expandTableExternalDataConfigurationBigtableOptionsColumnFamiliesMap(c *Client, f map[string]TableExternalDataConfigurationBigtableOptionsColumnFamilies, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableExternalDataConfigurationBigtableOptionsColumnFamilies(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableExternalDataConfigurationBigtableOptionsColumnFamiliesSlice expands the contents of TableExternalDataConfigurationBigtableOptionsColumnFamilies into a JSON
// request object.
func expandTableExternalDataConfigurationBigtableOptionsColumnFamiliesSlice(c *Client, f []TableExternalDataConfigurationBigtableOptionsColumnFamilies, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableExternalDataConfigurationBigtableOptionsColumnFamilies(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableExternalDataConfigurationBigtableOptionsColumnFamiliesMap flattens the contents of TableExternalDataConfigurationBigtableOptionsColumnFamilies from a JSON
// response object.
func flattenTableExternalDataConfigurationBigtableOptionsColumnFamiliesMap(c *Client, i interface{}, res *Table) map[string]TableExternalDataConfigurationBigtableOptionsColumnFamilies {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableExternalDataConfigurationBigtableOptionsColumnFamilies{}
	}

	if len(a) == 0 {
		return map[string]TableExternalDataConfigurationBigtableOptionsColumnFamilies{}
	}

	items := make(map[string]TableExternalDataConfigurationBigtableOptionsColumnFamilies)
	for k, item := range a {
		items[k] = *flattenTableExternalDataConfigurationBigtableOptionsColumnFamilies(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableExternalDataConfigurationBigtableOptionsColumnFamiliesSlice flattens the contents of TableExternalDataConfigurationBigtableOptionsColumnFamilies from a JSON
// response object.
func flattenTableExternalDataConfigurationBigtableOptionsColumnFamiliesSlice(c *Client, i interface{}, res *Table) []TableExternalDataConfigurationBigtableOptionsColumnFamilies {
	a, ok := i.([]interface{})
	if !ok {
		return []TableExternalDataConfigurationBigtableOptionsColumnFamilies{}
	}

	if len(a) == 0 {
		return []TableExternalDataConfigurationBigtableOptionsColumnFamilies{}
	}

	items := make([]TableExternalDataConfigurationBigtableOptionsColumnFamilies, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableExternalDataConfigurationBigtableOptionsColumnFamilies(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableExternalDataConfigurationBigtableOptionsColumnFamilies expands an instance of TableExternalDataConfigurationBigtableOptionsColumnFamilies into a JSON
// request object.
func expandTableExternalDataConfigurationBigtableOptionsColumnFamilies(c *Client, f *TableExternalDataConfigurationBigtableOptionsColumnFamilies, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.FamilyId; !dcl.IsEmptyValueIndirect(v) {
		m["familyId"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.Encoding; !dcl.IsEmptyValueIndirect(v) {
		m["encoding"] = v
	}
	if v, err := expandTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsSlice(c, f.Columns, res); err != nil {
		return nil, fmt.Errorf("error expanding Columns into columns: %w", err)
	} else if v != nil {
		m["columns"] = v
	}
	if v := f.OnlyReadLatest; !dcl.IsEmptyValueIndirect(v) {
		m["onlyReadLatest"] = v
	}

	return m, nil
}

// flattenTableExternalDataConfigurationBigtableOptionsColumnFamilies flattens an instance of TableExternalDataConfigurationBigtableOptionsColumnFamilies from a JSON
// response object.
func flattenTableExternalDataConfigurationBigtableOptionsColumnFamilies(c *Client, i interface{}, res *Table) *TableExternalDataConfigurationBigtableOptionsColumnFamilies {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableExternalDataConfigurationBigtableOptionsColumnFamilies{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableExternalDataConfigurationBigtableOptionsColumnFamilies
	}
	r.FamilyId = dcl.FlattenString(m["familyId"])
	r.Type = dcl.FlattenString(m["type"])
	r.Encoding = dcl.FlattenString(m["encoding"])
	r.Columns = flattenTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsSlice(c, m["columns"], res)
	r.OnlyReadLatest = dcl.FlattenBool(m["onlyReadLatest"])

	return r
}

// expandTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsMap expands the contents of TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns into a JSON
// request object.
func expandTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsMap(c *Client, f map[string]TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsSlice expands the contents of TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns into a JSON
// request object.
func expandTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsSlice(c *Client, f []TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsMap flattens the contents of TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns from a JSON
// response object.
func flattenTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsMap(c *Client, i interface{}, res *Table) map[string]TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns{}
	}

	if len(a) == 0 {
		return map[string]TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns{}
	}

	items := make(map[string]TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns)
	for k, item := range a {
		items[k] = *flattenTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsSlice flattens the contents of TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns from a JSON
// response object.
func flattenTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsSlice(c *Client, i interface{}, res *Table) []TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns {
	a, ok := i.([]interface{})
	if !ok {
		return []TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns{}
	}

	if len(a) == 0 {
		return []TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns{}
	}

	items := make([]TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns expands an instance of TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns into a JSON
// request object.
func expandTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns(c *Client, f *TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.QualifierEncoded; !dcl.IsEmptyValueIndirect(v) {
		m["qualifierEncoded"] = v
	}
	if v := f.QualifierString; !dcl.IsEmptyValueIndirect(v) {
		m["qualifierString"] = v
	}
	if v := f.FieldName; !dcl.IsEmptyValueIndirect(v) {
		m["fieldName"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.Encoding; !dcl.IsEmptyValueIndirect(v) {
		m["encoding"] = v
	}
	if v := f.OnlyReadLatest; !dcl.IsEmptyValueIndirect(v) {
		m["onlyReadLatest"] = v
	}

	return m, nil
}

// flattenTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns flattens an instance of TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns from a JSON
// response object.
func flattenTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns(c *Client, i interface{}, res *Table) *TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns
	}
	r.QualifierEncoded = dcl.FlattenString(m["qualifierEncoded"])
	r.QualifierString = dcl.FlattenString(m["qualifierString"])
	r.FieldName = dcl.FlattenString(m["fieldName"])
	r.Type = dcl.FlattenString(m["type"])
	r.Encoding = dcl.FlattenString(m["encoding"])
	r.OnlyReadLatest = dcl.FlattenBool(m["onlyReadLatest"])

	return r
}

// expandTableExternalDataConfigurationGoogleSheetsOptionsMap expands the contents of TableExternalDataConfigurationGoogleSheetsOptions into a JSON
// request object.
func expandTableExternalDataConfigurationGoogleSheetsOptionsMap(c *Client, f map[string]TableExternalDataConfigurationGoogleSheetsOptions, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableExternalDataConfigurationGoogleSheetsOptions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableExternalDataConfigurationGoogleSheetsOptionsSlice expands the contents of TableExternalDataConfigurationGoogleSheetsOptions into a JSON
// request object.
func expandTableExternalDataConfigurationGoogleSheetsOptionsSlice(c *Client, f []TableExternalDataConfigurationGoogleSheetsOptions, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableExternalDataConfigurationGoogleSheetsOptions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableExternalDataConfigurationGoogleSheetsOptionsMap flattens the contents of TableExternalDataConfigurationGoogleSheetsOptions from a JSON
// response object.
func flattenTableExternalDataConfigurationGoogleSheetsOptionsMap(c *Client, i interface{}, res *Table) map[string]TableExternalDataConfigurationGoogleSheetsOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableExternalDataConfigurationGoogleSheetsOptions{}
	}

	if len(a) == 0 {
		return map[string]TableExternalDataConfigurationGoogleSheetsOptions{}
	}

	items := make(map[string]TableExternalDataConfigurationGoogleSheetsOptions)
	for k, item := range a {
		items[k] = *flattenTableExternalDataConfigurationGoogleSheetsOptions(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableExternalDataConfigurationGoogleSheetsOptionsSlice flattens the contents of TableExternalDataConfigurationGoogleSheetsOptions from a JSON
// response object.
func flattenTableExternalDataConfigurationGoogleSheetsOptionsSlice(c *Client, i interface{}, res *Table) []TableExternalDataConfigurationGoogleSheetsOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []TableExternalDataConfigurationGoogleSheetsOptions{}
	}

	if len(a) == 0 {
		return []TableExternalDataConfigurationGoogleSheetsOptions{}
	}

	items := make([]TableExternalDataConfigurationGoogleSheetsOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableExternalDataConfigurationGoogleSheetsOptions(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableExternalDataConfigurationGoogleSheetsOptions expands an instance of TableExternalDataConfigurationGoogleSheetsOptions into a JSON
// request object.
func expandTableExternalDataConfigurationGoogleSheetsOptions(c *Client, f *TableExternalDataConfigurationGoogleSheetsOptions, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SkipLeadingRows; !dcl.IsEmptyValueIndirect(v) {
		m["skipLeadingRows"] = v
	}
	if v := f.Range; !dcl.IsEmptyValueIndirect(v) {
		m["range"] = v
	}

	return m, nil
}

// flattenTableExternalDataConfigurationGoogleSheetsOptions flattens an instance of TableExternalDataConfigurationGoogleSheetsOptions from a JSON
// response object.
func flattenTableExternalDataConfigurationGoogleSheetsOptions(c *Client, i interface{}, res *Table) *TableExternalDataConfigurationGoogleSheetsOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableExternalDataConfigurationGoogleSheetsOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableExternalDataConfigurationGoogleSheetsOptions
	}
	r.SkipLeadingRows = dcl.FlattenString(m["skipLeadingRows"])
	r.Range = dcl.FlattenString(m["range"])

	return r
}

// expandTableExternalDataConfigurationHivePartitioningOptionsMap expands the contents of TableExternalDataConfigurationHivePartitioningOptions into a JSON
// request object.
func expandTableExternalDataConfigurationHivePartitioningOptionsMap(c *Client, f map[string]TableExternalDataConfigurationHivePartitioningOptions, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableExternalDataConfigurationHivePartitioningOptions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableExternalDataConfigurationHivePartitioningOptionsSlice expands the contents of TableExternalDataConfigurationHivePartitioningOptions into a JSON
// request object.
func expandTableExternalDataConfigurationHivePartitioningOptionsSlice(c *Client, f []TableExternalDataConfigurationHivePartitioningOptions, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableExternalDataConfigurationHivePartitioningOptions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableExternalDataConfigurationHivePartitioningOptionsMap flattens the contents of TableExternalDataConfigurationHivePartitioningOptions from a JSON
// response object.
func flattenTableExternalDataConfigurationHivePartitioningOptionsMap(c *Client, i interface{}, res *Table) map[string]TableExternalDataConfigurationHivePartitioningOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableExternalDataConfigurationHivePartitioningOptions{}
	}

	if len(a) == 0 {
		return map[string]TableExternalDataConfigurationHivePartitioningOptions{}
	}

	items := make(map[string]TableExternalDataConfigurationHivePartitioningOptions)
	for k, item := range a {
		items[k] = *flattenTableExternalDataConfigurationHivePartitioningOptions(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableExternalDataConfigurationHivePartitioningOptionsSlice flattens the contents of TableExternalDataConfigurationHivePartitioningOptions from a JSON
// response object.
func flattenTableExternalDataConfigurationHivePartitioningOptionsSlice(c *Client, i interface{}, res *Table) []TableExternalDataConfigurationHivePartitioningOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []TableExternalDataConfigurationHivePartitioningOptions{}
	}

	if len(a) == 0 {
		return []TableExternalDataConfigurationHivePartitioningOptions{}
	}

	items := make([]TableExternalDataConfigurationHivePartitioningOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableExternalDataConfigurationHivePartitioningOptions(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableExternalDataConfigurationHivePartitioningOptions expands an instance of TableExternalDataConfigurationHivePartitioningOptions into a JSON
// request object.
func expandTableExternalDataConfigurationHivePartitioningOptions(c *Client, f *TableExternalDataConfigurationHivePartitioningOptions, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Mode; !dcl.IsEmptyValueIndirect(v) {
		m["mode"] = v
	}
	if v := f.SourceUriPrefix; !dcl.IsEmptyValueIndirect(v) {
		m["sourceUriPrefix"] = v
	}
	if v := f.RequirePartitionFilter; !dcl.IsEmptyValueIndirect(v) {
		m["requirePartitionFilter"] = v
	}

	return m, nil
}

// flattenTableExternalDataConfigurationHivePartitioningOptions flattens an instance of TableExternalDataConfigurationHivePartitioningOptions from a JSON
// response object.
func flattenTableExternalDataConfigurationHivePartitioningOptions(c *Client, i interface{}, res *Table) *TableExternalDataConfigurationHivePartitioningOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableExternalDataConfigurationHivePartitioningOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableExternalDataConfigurationHivePartitioningOptions
	}
	r.Mode = dcl.FlattenString(m["mode"])
	r.SourceUriPrefix = dcl.FlattenString(m["sourceUriPrefix"])
	r.RequirePartitionFilter = dcl.FlattenBool(m["requirePartitionFilter"])
	r.Fields = dcl.FlattenStringSlice(m["fields"])

	return r
}

// expandTableExternalDataConfigurationValueConversionModesMap expands the contents of TableExternalDataConfigurationValueConversionModes into a JSON
// request object.
func expandTableExternalDataConfigurationValueConversionModesMap(c *Client, f map[string]TableExternalDataConfigurationValueConversionModes, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableExternalDataConfigurationValueConversionModes(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableExternalDataConfigurationValueConversionModesSlice expands the contents of TableExternalDataConfigurationValueConversionModes into a JSON
// request object.
func expandTableExternalDataConfigurationValueConversionModesSlice(c *Client, f []TableExternalDataConfigurationValueConversionModes, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableExternalDataConfigurationValueConversionModes(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableExternalDataConfigurationValueConversionModesMap flattens the contents of TableExternalDataConfigurationValueConversionModes from a JSON
// response object.
func flattenTableExternalDataConfigurationValueConversionModesMap(c *Client, i interface{}, res *Table) map[string]TableExternalDataConfigurationValueConversionModes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableExternalDataConfigurationValueConversionModes{}
	}

	if len(a) == 0 {
		return map[string]TableExternalDataConfigurationValueConversionModes{}
	}

	items := make(map[string]TableExternalDataConfigurationValueConversionModes)
	for k, item := range a {
		items[k] = *flattenTableExternalDataConfigurationValueConversionModes(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableExternalDataConfigurationValueConversionModesSlice flattens the contents of TableExternalDataConfigurationValueConversionModes from a JSON
// response object.
func flattenTableExternalDataConfigurationValueConversionModesSlice(c *Client, i interface{}, res *Table) []TableExternalDataConfigurationValueConversionModes {
	a, ok := i.([]interface{})
	if !ok {
		return []TableExternalDataConfigurationValueConversionModes{}
	}

	if len(a) == 0 {
		return []TableExternalDataConfigurationValueConversionModes{}
	}

	items := make([]TableExternalDataConfigurationValueConversionModes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableExternalDataConfigurationValueConversionModes(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableExternalDataConfigurationValueConversionModes expands an instance of TableExternalDataConfigurationValueConversionModes into a JSON
// request object.
func expandTableExternalDataConfigurationValueConversionModes(c *Client, f *TableExternalDataConfigurationValueConversionModes, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.TemporalTypesOutOfRangeConversionMode; !dcl.IsEmptyValueIndirect(v) {
		m["temporalTypesOutOfRangeConversionMode"] = v
	}
	if v := f.NumericTypeOutOfRangeConversionMode; !dcl.IsEmptyValueIndirect(v) {
		m["numericTypeOutOfRangeConversionMode"] = v
	}

	return m, nil
}

// flattenTableExternalDataConfigurationValueConversionModes flattens an instance of TableExternalDataConfigurationValueConversionModes from a JSON
// response object.
func flattenTableExternalDataConfigurationValueConversionModes(c *Client, i interface{}, res *Table) *TableExternalDataConfigurationValueConversionModes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableExternalDataConfigurationValueConversionModes{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableExternalDataConfigurationValueConversionModes
	}
	r.TemporalTypesOutOfRangeConversionMode = flattenTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(m["temporalTypesOutOfRangeConversionMode"])
	r.NumericTypeOutOfRangeConversionMode = flattenTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(m["numericTypeOutOfRangeConversionMode"])

	return r
}

// expandTableExternalDataConfigurationAvroOptionsMap expands the contents of TableExternalDataConfigurationAvroOptions into a JSON
// request object.
func expandTableExternalDataConfigurationAvroOptionsMap(c *Client, f map[string]TableExternalDataConfigurationAvroOptions, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableExternalDataConfigurationAvroOptions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableExternalDataConfigurationAvroOptionsSlice expands the contents of TableExternalDataConfigurationAvroOptions into a JSON
// request object.
func expandTableExternalDataConfigurationAvroOptionsSlice(c *Client, f []TableExternalDataConfigurationAvroOptions, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableExternalDataConfigurationAvroOptions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableExternalDataConfigurationAvroOptionsMap flattens the contents of TableExternalDataConfigurationAvroOptions from a JSON
// response object.
func flattenTableExternalDataConfigurationAvroOptionsMap(c *Client, i interface{}, res *Table) map[string]TableExternalDataConfigurationAvroOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableExternalDataConfigurationAvroOptions{}
	}

	if len(a) == 0 {
		return map[string]TableExternalDataConfigurationAvroOptions{}
	}

	items := make(map[string]TableExternalDataConfigurationAvroOptions)
	for k, item := range a {
		items[k] = *flattenTableExternalDataConfigurationAvroOptions(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableExternalDataConfigurationAvroOptionsSlice flattens the contents of TableExternalDataConfigurationAvroOptions from a JSON
// response object.
func flattenTableExternalDataConfigurationAvroOptionsSlice(c *Client, i interface{}, res *Table) []TableExternalDataConfigurationAvroOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []TableExternalDataConfigurationAvroOptions{}
	}

	if len(a) == 0 {
		return []TableExternalDataConfigurationAvroOptions{}
	}

	items := make([]TableExternalDataConfigurationAvroOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableExternalDataConfigurationAvroOptions(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableExternalDataConfigurationAvroOptions expands an instance of TableExternalDataConfigurationAvroOptions into a JSON
// request object.
func expandTableExternalDataConfigurationAvroOptions(c *Client, f *TableExternalDataConfigurationAvroOptions, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.UseAvroLogicalTypes; !dcl.IsEmptyValueIndirect(v) {
		m["useAvroLogicalTypes"] = v
	}

	return m, nil
}

// flattenTableExternalDataConfigurationAvroOptions flattens an instance of TableExternalDataConfigurationAvroOptions from a JSON
// response object.
func flattenTableExternalDataConfigurationAvroOptions(c *Client, i interface{}, res *Table) *TableExternalDataConfigurationAvroOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableExternalDataConfigurationAvroOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableExternalDataConfigurationAvroOptions
	}
	r.UseAvroLogicalTypes = dcl.FlattenBool(m["useAvroLogicalTypes"])

	return r
}

// expandTableExternalDataConfigurationParquetOptionsMap expands the contents of TableExternalDataConfigurationParquetOptions into a JSON
// request object.
func expandTableExternalDataConfigurationParquetOptionsMap(c *Client, f map[string]TableExternalDataConfigurationParquetOptions, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableExternalDataConfigurationParquetOptions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableExternalDataConfigurationParquetOptionsSlice expands the contents of TableExternalDataConfigurationParquetOptions into a JSON
// request object.
func expandTableExternalDataConfigurationParquetOptionsSlice(c *Client, f []TableExternalDataConfigurationParquetOptions, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableExternalDataConfigurationParquetOptions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableExternalDataConfigurationParquetOptionsMap flattens the contents of TableExternalDataConfigurationParquetOptions from a JSON
// response object.
func flattenTableExternalDataConfigurationParquetOptionsMap(c *Client, i interface{}, res *Table) map[string]TableExternalDataConfigurationParquetOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableExternalDataConfigurationParquetOptions{}
	}

	if len(a) == 0 {
		return map[string]TableExternalDataConfigurationParquetOptions{}
	}

	items := make(map[string]TableExternalDataConfigurationParquetOptions)
	for k, item := range a {
		items[k] = *flattenTableExternalDataConfigurationParquetOptions(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableExternalDataConfigurationParquetOptionsSlice flattens the contents of TableExternalDataConfigurationParquetOptions from a JSON
// response object.
func flattenTableExternalDataConfigurationParquetOptionsSlice(c *Client, i interface{}, res *Table) []TableExternalDataConfigurationParquetOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []TableExternalDataConfigurationParquetOptions{}
	}

	if len(a) == 0 {
		return []TableExternalDataConfigurationParquetOptions{}
	}

	items := make([]TableExternalDataConfigurationParquetOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableExternalDataConfigurationParquetOptions(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableExternalDataConfigurationParquetOptions expands an instance of TableExternalDataConfigurationParquetOptions into a JSON
// request object.
func expandTableExternalDataConfigurationParquetOptions(c *Client, f *TableExternalDataConfigurationParquetOptions, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EnumAsString; !dcl.IsEmptyValueIndirect(v) {
		m["enumAsString"] = v
	}
	if v := f.EnableListInference; !dcl.IsEmptyValueIndirect(v) {
		m["enableListInference"] = v
	}

	return m, nil
}

// flattenTableExternalDataConfigurationParquetOptions flattens an instance of TableExternalDataConfigurationParquetOptions from a JSON
// response object.
func flattenTableExternalDataConfigurationParquetOptions(c *Client, i interface{}, res *Table) *TableExternalDataConfigurationParquetOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableExternalDataConfigurationParquetOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableExternalDataConfigurationParquetOptions
	}
	r.EnumAsString = dcl.FlattenBool(m["enumAsString"])
	r.EnableListInference = dcl.FlattenBool(m["enableListInference"])

	return r
}

// expandTableStreamingBufferMap expands the contents of TableStreamingBuffer into a JSON
// request object.
func expandTableStreamingBufferMap(c *Client, f map[string]TableStreamingBuffer, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableStreamingBuffer(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableStreamingBufferSlice expands the contents of TableStreamingBuffer into a JSON
// request object.
func expandTableStreamingBufferSlice(c *Client, f []TableStreamingBuffer, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableStreamingBuffer(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableStreamingBufferMap flattens the contents of TableStreamingBuffer from a JSON
// response object.
func flattenTableStreamingBufferMap(c *Client, i interface{}, res *Table) map[string]TableStreamingBuffer {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableStreamingBuffer{}
	}

	if len(a) == 0 {
		return map[string]TableStreamingBuffer{}
	}

	items := make(map[string]TableStreamingBuffer)
	for k, item := range a {
		items[k] = *flattenTableStreamingBuffer(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableStreamingBufferSlice flattens the contents of TableStreamingBuffer from a JSON
// response object.
func flattenTableStreamingBufferSlice(c *Client, i interface{}, res *Table) []TableStreamingBuffer {
	a, ok := i.([]interface{})
	if !ok {
		return []TableStreamingBuffer{}
	}

	if len(a) == 0 {
		return []TableStreamingBuffer{}
	}

	items := make([]TableStreamingBuffer, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableStreamingBuffer(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableStreamingBuffer expands an instance of TableStreamingBuffer into a JSON
// request object.
func expandTableStreamingBuffer(c *Client, f *TableStreamingBuffer, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenTableStreamingBuffer flattens an instance of TableStreamingBuffer from a JSON
// response object.
func flattenTableStreamingBuffer(c *Client, i interface{}, res *Table) *TableStreamingBuffer {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableStreamingBuffer{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableStreamingBuffer
	}
	r.EstimatedBytes = dcl.FlattenInteger(m["estimatedBytes"])
	r.EstimatedRows = dcl.FlattenInteger(m["estimatedRows"])
	r.OldestEntryTime = dcl.FlattenInteger(m["oldestEntryTime"])

	return r
}

// expandTableEncryptionConfigurationMap expands the contents of TableEncryptionConfiguration into a JSON
// request object.
func expandTableEncryptionConfigurationMap(c *Client, f map[string]TableEncryptionConfiguration, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableEncryptionConfiguration(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableEncryptionConfigurationSlice expands the contents of TableEncryptionConfiguration into a JSON
// request object.
func expandTableEncryptionConfigurationSlice(c *Client, f []TableEncryptionConfiguration, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableEncryptionConfiguration(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableEncryptionConfigurationMap flattens the contents of TableEncryptionConfiguration from a JSON
// response object.
func flattenTableEncryptionConfigurationMap(c *Client, i interface{}, res *Table) map[string]TableEncryptionConfiguration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableEncryptionConfiguration{}
	}

	if len(a) == 0 {
		return map[string]TableEncryptionConfiguration{}
	}

	items := make(map[string]TableEncryptionConfiguration)
	for k, item := range a {
		items[k] = *flattenTableEncryptionConfiguration(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableEncryptionConfigurationSlice flattens the contents of TableEncryptionConfiguration from a JSON
// response object.
func flattenTableEncryptionConfigurationSlice(c *Client, i interface{}, res *Table) []TableEncryptionConfiguration {
	a, ok := i.([]interface{})
	if !ok {
		return []TableEncryptionConfiguration{}
	}

	if len(a) == 0 {
		return []TableEncryptionConfiguration{}
	}

	items := make([]TableEncryptionConfiguration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableEncryptionConfiguration(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableEncryptionConfiguration expands an instance of TableEncryptionConfiguration into a JSON
// request object.
func expandTableEncryptionConfiguration(c *Client, f *TableEncryptionConfiguration, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.KmsKeyName; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyName"] = v
	}

	return m, nil
}

// flattenTableEncryptionConfiguration flattens an instance of TableEncryptionConfiguration from a JSON
// response object.
func flattenTableEncryptionConfiguration(c *Client, i interface{}, res *Table) *TableEncryptionConfiguration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableEncryptionConfiguration{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableEncryptionConfiguration
	}
	r.KmsKeyName = dcl.FlattenString(m["kmsKeyName"])

	return r
}

// expandTableSnapshotDefinitionMap expands the contents of TableSnapshotDefinition into a JSON
// request object.
func expandTableSnapshotDefinitionMap(c *Client, f map[string]TableSnapshotDefinition, res *Table) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTableSnapshotDefinition(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTableSnapshotDefinitionSlice expands the contents of TableSnapshotDefinition into a JSON
// request object.
func expandTableSnapshotDefinitionSlice(c *Client, f []TableSnapshotDefinition, res *Table) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTableSnapshotDefinition(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTableSnapshotDefinitionMap flattens the contents of TableSnapshotDefinition from a JSON
// response object.
func flattenTableSnapshotDefinitionMap(c *Client, i interface{}, res *Table) map[string]TableSnapshotDefinition {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableSnapshotDefinition{}
	}

	if len(a) == 0 {
		return map[string]TableSnapshotDefinition{}
	}

	items := make(map[string]TableSnapshotDefinition)
	for k, item := range a {
		items[k] = *flattenTableSnapshotDefinition(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTableSnapshotDefinitionSlice flattens the contents of TableSnapshotDefinition from a JSON
// response object.
func flattenTableSnapshotDefinitionSlice(c *Client, i interface{}, res *Table) []TableSnapshotDefinition {
	a, ok := i.([]interface{})
	if !ok {
		return []TableSnapshotDefinition{}
	}

	if len(a) == 0 {
		return []TableSnapshotDefinition{}
	}

	items := make([]TableSnapshotDefinition, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableSnapshotDefinition(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTableSnapshotDefinition expands an instance of TableSnapshotDefinition into a JSON
// request object.
func expandTableSnapshotDefinition(c *Client, f *TableSnapshotDefinition, res *Table) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Table; !dcl.IsEmptyValueIndirect(v) {
		m["table"] = v
	}
	if v := f.Dataset; !dcl.IsEmptyValueIndirect(v) {
		m["dataset"] = v
	}
	if v := f.Project; !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v := f.SnapshotTime; !dcl.IsEmptyValueIndirect(v) {
		m["snapshotTime"] = v
	}

	return m, nil
}

// flattenTableSnapshotDefinition flattens an instance of TableSnapshotDefinition from a JSON
// response object.
func flattenTableSnapshotDefinition(c *Client, i interface{}, res *Table) *TableSnapshotDefinition {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TableSnapshotDefinition{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTableSnapshotDefinition
	}
	r.Table = dcl.FlattenString(m["table"])
	r.Dataset = dcl.FlattenString(m["dataset"])
	r.Project = dcl.FlattenString(m["project"])
	r.SnapshotTime = dcl.FlattenString(m["snapshotTime"])

	return r
}

// flattenTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnumMap flattens the contents of TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum from a JSON
// response object.
func flattenTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnumMap(c *Client, i interface{}, res *Table) map[string]TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum{}
	}

	if len(a) == 0 {
		return map[string]TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum{}
	}

	items := make(map[string]TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum)
	for k, item := range a {
		items[k] = *flattenTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(item.(interface{}))
	}

	return items
}

// flattenTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnumSlice flattens the contents of TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum from a JSON
// response object.
func flattenTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnumSlice(c *Client, i interface{}, res *Table) []TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum{}
	}

	if len(a) == 0 {
		return []TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum{}
	}

	items := make([]TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(item.(interface{})))
	}

	return items
}

// flattenTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum asserts that an interface is a string, and returns a
// pointer to a *TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum with the same value as that string.
func flattenTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(i interface{}) *TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnumRef(s)
}

// flattenTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnumMap flattens the contents of TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum from a JSON
// response object.
func flattenTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnumMap(c *Client, i interface{}, res *Table) map[string]TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum{}
	}

	if len(a) == 0 {
		return map[string]TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum{}
	}

	items := make(map[string]TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum)
	for k, item := range a {
		items[k] = *flattenTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(item.(interface{}))
	}

	return items
}

// flattenTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnumSlice flattens the contents of TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum from a JSON
// response object.
func flattenTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnumSlice(c *Client, i interface{}, res *Table) []TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum{}
	}

	if len(a) == 0 {
		return []TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum{}
	}

	items := make([]TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(item.(interface{})))
	}

	return items
}

// flattenTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum asserts that an interface is a string, and returns a
// pointer to a *TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum with the same value as that string.
func flattenTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(i interface{}) *TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnumRef(s)
}

// flattenTableExternalDataConfigurationDecimalTargetTypesEnumMap flattens the contents of TableExternalDataConfigurationDecimalTargetTypesEnum from a JSON
// response object.
func flattenTableExternalDataConfigurationDecimalTargetTypesEnumMap(c *Client, i interface{}, res *Table) map[string]TableExternalDataConfigurationDecimalTargetTypesEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableExternalDataConfigurationDecimalTargetTypesEnum{}
	}

	if len(a) == 0 {
		return map[string]TableExternalDataConfigurationDecimalTargetTypesEnum{}
	}

	items := make(map[string]TableExternalDataConfigurationDecimalTargetTypesEnum)
	for k, item := range a {
		items[k] = *flattenTableExternalDataConfigurationDecimalTargetTypesEnum(item.(interface{}))
	}

	return items
}

// flattenTableExternalDataConfigurationDecimalTargetTypesEnumSlice flattens the contents of TableExternalDataConfigurationDecimalTargetTypesEnum from a JSON
// response object.
func flattenTableExternalDataConfigurationDecimalTargetTypesEnumSlice(c *Client, i interface{}, res *Table) []TableExternalDataConfigurationDecimalTargetTypesEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []TableExternalDataConfigurationDecimalTargetTypesEnum{}
	}

	if len(a) == 0 {
		return []TableExternalDataConfigurationDecimalTargetTypesEnum{}
	}

	items := make([]TableExternalDataConfigurationDecimalTargetTypesEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableExternalDataConfigurationDecimalTargetTypesEnum(item.(interface{})))
	}

	return items
}

// flattenTableExternalDataConfigurationDecimalTargetTypesEnum asserts that an interface is a string, and returns a
// pointer to a *TableExternalDataConfigurationDecimalTargetTypesEnum with the same value as that string.
func flattenTableExternalDataConfigurationDecimalTargetTypesEnum(i interface{}) *TableExternalDataConfigurationDecimalTargetTypesEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return TableExternalDataConfigurationDecimalTargetTypesEnumRef(s)
}

// flattenTableExternalDataConfigurationJsonExtensionEnumMap flattens the contents of TableExternalDataConfigurationJsonExtensionEnum from a JSON
// response object.
func flattenTableExternalDataConfigurationJsonExtensionEnumMap(c *Client, i interface{}, res *Table) map[string]TableExternalDataConfigurationJsonExtensionEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TableExternalDataConfigurationJsonExtensionEnum{}
	}

	if len(a) == 0 {
		return map[string]TableExternalDataConfigurationJsonExtensionEnum{}
	}

	items := make(map[string]TableExternalDataConfigurationJsonExtensionEnum)
	for k, item := range a {
		items[k] = *flattenTableExternalDataConfigurationJsonExtensionEnum(item.(interface{}))
	}

	return items
}

// flattenTableExternalDataConfigurationJsonExtensionEnumSlice flattens the contents of TableExternalDataConfigurationJsonExtensionEnum from a JSON
// response object.
func flattenTableExternalDataConfigurationJsonExtensionEnumSlice(c *Client, i interface{}, res *Table) []TableExternalDataConfigurationJsonExtensionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []TableExternalDataConfigurationJsonExtensionEnum{}
	}

	if len(a) == 0 {
		return []TableExternalDataConfigurationJsonExtensionEnum{}
	}

	items := make([]TableExternalDataConfigurationJsonExtensionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTableExternalDataConfigurationJsonExtensionEnum(item.(interface{})))
	}

	return items
}

// flattenTableExternalDataConfigurationJsonExtensionEnum asserts that an interface is a string, and returns a
// pointer to a *TableExternalDataConfigurationJsonExtensionEnum with the same value as that string.
func flattenTableExternalDataConfigurationJsonExtensionEnum(i interface{}) *TableExternalDataConfigurationJsonExtensionEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return TableExternalDataConfigurationJsonExtensionEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Table) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalTable(b, c, r)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
			return false
		}
		if nr.Dataset == nil && ncr.Dataset == nil {
			c.Config.Logger.Info("Both Dataset fields null - considering equal.")
		} else if nr.Dataset == nil || ncr.Dataset == nil {
			c.Config.Logger.Info("Only one Dataset field is null - considering unequal.")
			return false
		} else if *nr.Dataset != *ncr.Dataset {
			return false
		}
		if nr.Name == nil && ncr.Name == nil {
			c.Config.Logger.Info("Both Name fields null - considering equal.")
		} else if nr.Name == nil || ncr.Name == nil {
			c.Config.Logger.Info("Only one Name field is null - considering unequal.")
			return false
		} else if *nr.Name != *ncr.Name {
			return false
		}
		return true
	}
}

type tableDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         tableApiOperation
}

func convertFieldDiffsToTableDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]tableDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff: %v", ro, fd)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []tableDiff
	// For each operation name, create a tableDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := tableDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToTableApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToTableApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (tableApiOperation, error) {
	switch opName {

	case "updateTablePatchTableOperation":
		return &updateTablePatchTableOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractTableFields(r *Table) error {
	vModel := r.Model
	if vModel == nil {
		// note: explicitly not the empty object.
		vModel = &TableModel{}
	}
	if err := extractTableModelFields(r, vModel); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vModel) {
		r.Model = vModel
	}
	vSchema := r.Schema
	if vSchema == nil {
		// note: explicitly not the empty object.
		vSchema = &TableSchema{}
	}
	if err := extractTableSchemaFields(r, vSchema); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSchema) {
		r.Schema = vSchema
	}
	vTimePartitioning := r.TimePartitioning
	if vTimePartitioning == nil {
		// note: explicitly not the empty object.
		vTimePartitioning = &TableTimePartitioning{}
	}
	if err := extractTableTimePartitioningFields(r, vTimePartitioning); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vTimePartitioning) {
		r.TimePartitioning = vTimePartitioning
	}
	vRangePartitioning := r.RangePartitioning
	if vRangePartitioning == nil {
		// note: explicitly not the empty object.
		vRangePartitioning = &TableRangePartitioning{}
	}
	if err := extractTableRangePartitioningFields(r, vRangePartitioning); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vRangePartitioning) {
		r.RangePartitioning = vRangePartitioning
	}
	vClustering := r.Clustering
	if vClustering == nil {
		// note: explicitly not the empty object.
		vClustering = &TableClustering{}
	}
	if err := extractTableClusteringFields(r, vClustering); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vClustering) {
		r.Clustering = vClustering
	}
	vView := r.View
	if vView == nil {
		// note: explicitly not the empty object.
		vView = &TableView{}
	}
	if err := extractTableViewFields(r, vView); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vView) {
		r.View = vView
	}
	vMaterializedView := r.MaterializedView
	if vMaterializedView == nil {
		// note: explicitly not the empty object.
		vMaterializedView = &TableMaterializedView{}
	}
	if err := extractTableMaterializedViewFields(r, vMaterializedView); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vMaterializedView) {
		r.MaterializedView = vMaterializedView
	}
	vExternalDataConfiguration := r.ExternalDataConfiguration
	if vExternalDataConfiguration == nil {
		// note: explicitly not the empty object.
		vExternalDataConfiguration = &TableExternalDataConfiguration{}
	}
	if err := extractTableExternalDataConfigurationFields(r, vExternalDataConfiguration); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vExternalDataConfiguration) {
		r.ExternalDataConfiguration = vExternalDataConfiguration
	}
	vStreamingBuffer := r.StreamingBuffer
	if vStreamingBuffer == nil {
		// note: explicitly not the empty object.
		vStreamingBuffer = &TableStreamingBuffer{}
	}
	if err := extractTableStreamingBufferFields(r, vStreamingBuffer); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vStreamingBuffer) {
		r.StreamingBuffer = vStreamingBuffer
	}
	vEncryptionConfiguration := r.EncryptionConfiguration
	if vEncryptionConfiguration == nil {
		// note: explicitly not the empty object.
		vEncryptionConfiguration = &TableEncryptionConfiguration{}
	}
	if err := extractTableEncryptionConfigurationFields(r, vEncryptionConfiguration); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vEncryptionConfiguration) {
		r.EncryptionConfiguration = vEncryptionConfiguration
	}
	vSnapshotDefinition := r.SnapshotDefinition
	if vSnapshotDefinition == nil {
		// note: explicitly not the empty object.
		vSnapshotDefinition = &TableSnapshotDefinition{}
	}
	if err := extractTableSnapshotDefinitionFields(r, vSnapshotDefinition); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSnapshotDefinition) {
		r.SnapshotDefinition = vSnapshotDefinition
	}
	return nil
}
func extractTableModelFields(r *Table, o *TableModel) error {
	vModelOptions := o.ModelOptions
	if vModelOptions == nil {
		// note: explicitly not the empty object.
		vModelOptions = &TableModelModelOptions{}
	}
	if err := extractTableModelModelOptionsFields(r, vModelOptions); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vModelOptions) {
		o.ModelOptions = vModelOptions
	}
	return nil
}
func extractTableModelModelOptionsFields(r *Table, o *TableModelModelOptions) error {
	return nil
}
func extractTableModelTrainingRunsFields(r *Table, o *TableModelTrainingRuns) error {
	vTrainingOptions := o.TrainingOptions
	if vTrainingOptions == nil {
		// note: explicitly not the empty object.
		vTrainingOptions = &TableModelTrainingRunsTrainingOptions{}
	}
	if err := extractTableModelTrainingRunsTrainingOptionsFields(r, vTrainingOptions); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vTrainingOptions) {
		o.TrainingOptions = vTrainingOptions
	}
	return nil
}
func extractTableModelTrainingRunsTrainingOptionsFields(r *Table, o *TableModelTrainingRunsTrainingOptions) error {
	return nil
}
func extractTableModelTrainingRunsIterationResultsFields(r *Table, o *TableModelTrainingRunsIterationResults) error {
	return nil
}
func extractTableSchemaFields(r *Table, o *TableSchema) error {
	return nil
}
func extractTableGooglecloudbigqueryv2TablefieldschemaFields(r *Table, o *TableGooglecloudbigqueryv2Tablefieldschema) error {
	vCategories := o.Categories
	if vCategories == nil {
		// note: explicitly not the empty object.
		vCategories = &TableGooglecloudbigqueryv2TablefieldschemaCategories{}
	}
	if err := extractTableGooglecloudbigqueryv2TablefieldschemaCategoriesFields(r, vCategories); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vCategories) {
		o.Categories = vCategories
	}
	vPolicyTags := o.PolicyTags
	if vPolicyTags == nil {
		// note: explicitly not the empty object.
		vPolicyTags = &TableGooglecloudbigqueryv2TablefieldschemaPolicyTags{}
	}
	if err := extractTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsFields(r, vPolicyTags); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vPolicyTags) {
		o.PolicyTags = vPolicyTags
	}
	return nil
}
func extractTableGooglecloudbigqueryv2TablefieldschemaCategoriesFields(r *Table, o *TableGooglecloudbigqueryv2TablefieldschemaCategories) error {
	return nil
}
func extractTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsFields(r *Table, o *TableGooglecloudbigqueryv2TablefieldschemaPolicyTags) error {
	return nil
}
func extractTableTimePartitioningFields(r *Table, o *TableTimePartitioning) error {
	return nil
}
func extractTableRangePartitioningFields(r *Table, o *TableRangePartitioning) error {
	vRange := o.Range
	if vRange == nil {
		// note: explicitly not the empty object.
		vRange = &TableRangePartitioningRange{}
	}
	if err := extractTableRangePartitioningRangeFields(r, vRange); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vRange) {
		o.Range = vRange
	}
	return nil
}
func extractTableRangePartitioningRangeFields(r *Table, o *TableRangePartitioningRange) error {
	return nil
}
func extractTableClusteringFields(r *Table, o *TableClustering) error {
	return nil
}
func extractTableViewFields(r *Table, o *TableView) error {
	return nil
}
func extractTableViewUserDefinedFunctionResourcesFields(r *Table, o *TableViewUserDefinedFunctionResources) error {
	return nil
}
func extractTableMaterializedViewFields(r *Table, o *TableMaterializedView) error {
	return nil
}
func extractTableExternalDataConfigurationFields(r *Table, o *TableExternalDataConfiguration) error {
	vSchema := o.Schema
	if vSchema == nil {
		// note: explicitly not the empty object.
		vSchema = &TableExternalDataConfigurationSchema{}
	}
	if err := extractTableExternalDataConfigurationSchemaFields(r, vSchema); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSchema) {
		o.Schema = vSchema
	}
	vCsvOptions := o.CsvOptions
	if vCsvOptions == nil {
		// note: explicitly not the empty object.
		vCsvOptions = &TableExternalDataConfigurationCsvOptions{}
	}
	if err := extractTableExternalDataConfigurationCsvOptionsFields(r, vCsvOptions); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vCsvOptions) {
		o.CsvOptions = vCsvOptions
	}
	vBigtableOptions := o.BigtableOptions
	if vBigtableOptions == nil {
		// note: explicitly not the empty object.
		vBigtableOptions = &TableExternalDataConfigurationBigtableOptions{}
	}
	if err := extractTableExternalDataConfigurationBigtableOptionsFields(r, vBigtableOptions); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vBigtableOptions) {
		o.BigtableOptions = vBigtableOptions
	}
	vGoogleSheetsOptions := o.GoogleSheetsOptions
	if vGoogleSheetsOptions == nil {
		// note: explicitly not the empty object.
		vGoogleSheetsOptions = &TableExternalDataConfigurationGoogleSheetsOptions{}
	}
	if err := extractTableExternalDataConfigurationGoogleSheetsOptionsFields(r, vGoogleSheetsOptions); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vGoogleSheetsOptions) {
		o.GoogleSheetsOptions = vGoogleSheetsOptions
	}
	vHivePartitioningOptions := o.HivePartitioningOptions
	if vHivePartitioningOptions == nil {
		// note: explicitly not the empty object.
		vHivePartitioningOptions = &TableExternalDataConfigurationHivePartitioningOptions{}
	}
	if err := extractTableExternalDataConfigurationHivePartitioningOptionsFields(r, vHivePartitioningOptions); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vHivePartitioningOptions) {
		o.HivePartitioningOptions = vHivePartitioningOptions
	}
	vValueConversionModes := o.ValueConversionModes
	if vValueConversionModes == nil {
		// note: explicitly not the empty object.
		vValueConversionModes = &TableExternalDataConfigurationValueConversionModes{}
	}
	if err := extractTableExternalDataConfigurationValueConversionModesFields(r, vValueConversionModes); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vValueConversionModes) {
		o.ValueConversionModes = vValueConversionModes
	}
	vAvroOptions := o.AvroOptions
	if vAvroOptions == nil {
		// note: explicitly not the empty object.
		vAvroOptions = &TableExternalDataConfigurationAvroOptions{}
	}
	if err := extractTableExternalDataConfigurationAvroOptionsFields(r, vAvroOptions); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vAvroOptions) {
		o.AvroOptions = vAvroOptions
	}
	vParquetOptions := o.ParquetOptions
	if vParquetOptions == nil {
		// note: explicitly not the empty object.
		vParquetOptions = &TableExternalDataConfigurationParquetOptions{}
	}
	if err := extractTableExternalDataConfigurationParquetOptionsFields(r, vParquetOptions); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vParquetOptions) {
		o.ParquetOptions = vParquetOptions
	}
	return nil
}
func extractTableExternalDataConfigurationSchemaFields(r *Table, o *TableExternalDataConfigurationSchema) error {
	return nil
}
func extractTableExternalDataConfigurationCsvOptionsFields(r *Table, o *TableExternalDataConfigurationCsvOptions) error {
	return nil
}
func extractTableExternalDataConfigurationBigtableOptionsFields(r *Table, o *TableExternalDataConfigurationBigtableOptions) error {
	return nil
}
func extractTableExternalDataConfigurationBigtableOptionsColumnFamiliesFields(r *Table, o *TableExternalDataConfigurationBigtableOptionsColumnFamilies) error {
	return nil
}
func extractTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsFields(r *Table, o *TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns) error {
	return nil
}
func extractTableExternalDataConfigurationGoogleSheetsOptionsFields(r *Table, o *TableExternalDataConfigurationGoogleSheetsOptions) error {
	return nil
}
func extractTableExternalDataConfigurationHivePartitioningOptionsFields(r *Table, o *TableExternalDataConfigurationHivePartitioningOptions) error {
	return nil
}
func extractTableExternalDataConfigurationValueConversionModesFields(r *Table, o *TableExternalDataConfigurationValueConversionModes) error {
	return nil
}
func extractTableExternalDataConfigurationAvroOptionsFields(r *Table, o *TableExternalDataConfigurationAvroOptions) error {
	return nil
}
func extractTableExternalDataConfigurationParquetOptionsFields(r *Table, o *TableExternalDataConfigurationParquetOptions) error {
	return nil
}
func extractTableStreamingBufferFields(r *Table, o *TableStreamingBuffer) error {
	return nil
}
func extractTableEncryptionConfigurationFields(r *Table, o *TableEncryptionConfiguration) error {
	return nil
}
func extractTableSnapshotDefinitionFields(r *Table, o *TableSnapshotDefinition) error {
	return nil
}

func postReadExtractTableFields(r *Table) error {
	vModel := r.Model
	if vModel == nil {
		// note: explicitly not the empty object.
		vModel = &TableModel{}
	}
	if err := postReadExtractTableModelFields(r, vModel); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vModel) {
		r.Model = vModel
	}
	vSchema := r.Schema
	if vSchema == nil {
		// note: explicitly not the empty object.
		vSchema = &TableSchema{}
	}
	if err := postReadExtractTableSchemaFields(r, vSchema); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSchema) {
		r.Schema = vSchema
	}
	vTimePartitioning := r.TimePartitioning
	if vTimePartitioning == nil {
		// note: explicitly not the empty object.
		vTimePartitioning = &TableTimePartitioning{}
	}
	if err := postReadExtractTableTimePartitioningFields(r, vTimePartitioning); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vTimePartitioning) {
		r.TimePartitioning = vTimePartitioning
	}
	vRangePartitioning := r.RangePartitioning
	if vRangePartitioning == nil {
		// note: explicitly not the empty object.
		vRangePartitioning = &TableRangePartitioning{}
	}
	if err := postReadExtractTableRangePartitioningFields(r, vRangePartitioning); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vRangePartitioning) {
		r.RangePartitioning = vRangePartitioning
	}
	vClustering := r.Clustering
	if vClustering == nil {
		// note: explicitly not the empty object.
		vClustering = &TableClustering{}
	}
	if err := postReadExtractTableClusteringFields(r, vClustering); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vClustering) {
		r.Clustering = vClustering
	}
	vView := r.View
	if vView == nil {
		// note: explicitly not the empty object.
		vView = &TableView{}
	}
	if err := postReadExtractTableViewFields(r, vView); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vView) {
		r.View = vView
	}
	vMaterializedView := r.MaterializedView
	if vMaterializedView == nil {
		// note: explicitly not the empty object.
		vMaterializedView = &TableMaterializedView{}
	}
	if err := postReadExtractTableMaterializedViewFields(r, vMaterializedView); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vMaterializedView) {
		r.MaterializedView = vMaterializedView
	}
	vExternalDataConfiguration := r.ExternalDataConfiguration
	if vExternalDataConfiguration == nil {
		// note: explicitly not the empty object.
		vExternalDataConfiguration = &TableExternalDataConfiguration{}
	}
	if err := postReadExtractTableExternalDataConfigurationFields(r, vExternalDataConfiguration); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vExternalDataConfiguration) {
		r.ExternalDataConfiguration = vExternalDataConfiguration
	}
	vStreamingBuffer := r.StreamingBuffer
	if vStreamingBuffer == nil {
		// note: explicitly not the empty object.
		vStreamingBuffer = &TableStreamingBuffer{}
	}
	if err := postReadExtractTableStreamingBufferFields(r, vStreamingBuffer); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vStreamingBuffer) {
		r.StreamingBuffer = vStreamingBuffer
	}
	vEncryptionConfiguration := r.EncryptionConfiguration
	if vEncryptionConfiguration == nil {
		// note: explicitly not the empty object.
		vEncryptionConfiguration = &TableEncryptionConfiguration{}
	}
	if err := postReadExtractTableEncryptionConfigurationFields(r, vEncryptionConfiguration); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vEncryptionConfiguration) {
		r.EncryptionConfiguration = vEncryptionConfiguration
	}
	vSnapshotDefinition := r.SnapshotDefinition
	if vSnapshotDefinition == nil {
		// note: explicitly not the empty object.
		vSnapshotDefinition = &TableSnapshotDefinition{}
	}
	if err := postReadExtractTableSnapshotDefinitionFields(r, vSnapshotDefinition); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSnapshotDefinition) {
		r.SnapshotDefinition = vSnapshotDefinition
	}
	return nil
}
func postReadExtractTableModelFields(r *Table, o *TableModel) error {
	vModelOptions := o.ModelOptions
	if vModelOptions == nil {
		// note: explicitly not the empty object.
		vModelOptions = &TableModelModelOptions{}
	}
	if err := extractTableModelModelOptionsFields(r, vModelOptions); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vModelOptions) {
		o.ModelOptions = vModelOptions
	}
	return nil
}
func postReadExtractTableModelModelOptionsFields(r *Table, o *TableModelModelOptions) error {
	return nil
}
func postReadExtractTableModelTrainingRunsFields(r *Table, o *TableModelTrainingRuns) error {
	vTrainingOptions := o.TrainingOptions
	if vTrainingOptions == nil {
		// note: explicitly not the empty object.
		vTrainingOptions = &TableModelTrainingRunsTrainingOptions{}
	}
	if err := extractTableModelTrainingRunsTrainingOptionsFields(r, vTrainingOptions); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vTrainingOptions) {
		o.TrainingOptions = vTrainingOptions
	}
	return nil
}
func postReadExtractTableModelTrainingRunsTrainingOptionsFields(r *Table, o *TableModelTrainingRunsTrainingOptions) error {
	return nil
}
func postReadExtractTableModelTrainingRunsIterationResultsFields(r *Table, o *TableModelTrainingRunsIterationResults) error {
	return nil
}
func postReadExtractTableSchemaFields(r *Table, o *TableSchema) error {
	return nil
}
func postReadExtractTableGooglecloudbigqueryv2TablefieldschemaFields(r *Table, o *TableGooglecloudbigqueryv2Tablefieldschema) error {
	vCategories := o.Categories
	if vCategories == nil {
		// note: explicitly not the empty object.
		vCategories = &TableGooglecloudbigqueryv2TablefieldschemaCategories{}
	}
	if err := extractTableGooglecloudbigqueryv2TablefieldschemaCategoriesFields(r, vCategories); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vCategories) {
		o.Categories = vCategories
	}
	vPolicyTags := o.PolicyTags
	if vPolicyTags == nil {
		// note: explicitly not the empty object.
		vPolicyTags = &TableGooglecloudbigqueryv2TablefieldschemaPolicyTags{}
	}
	if err := extractTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsFields(r, vPolicyTags); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vPolicyTags) {
		o.PolicyTags = vPolicyTags
	}
	return nil
}
func postReadExtractTableGooglecloudbigqueryv2TablefieldschemaCategoriesFields(r *Table, o *TableGooglecloudbigqueryv2TablefieldschemaCategories) error {
	return nil
}
func postReadExtractTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsFields(r *Table, o *TableGooglecloudbigqueryv2TablefieldschemaPolicyTags) error {
	return nil
}
func postReadExtractTableTimePartitioningFields(r *Table, o *TableTimePartitioning) error {
	return nil
}
func postReadExtractTableRangePartitioningFields(r *Table, o *TableRangePartitioning) error {
	vRange := o.Range
	if vRange == nil {
		// note: explicitly not the empty object.
		vRange = &TableRangePartitioningRange{}
	}
	if err := extractTableRangePartitioningRangeFields(r, vRange); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vRange) {
		o.Range = vRange
	}
	return nil
}
func postReadExtractTableRangePartitioningRangeFields(r *Table, o *TableRangePartitioningRange) error {
	return nil
}
func postReadExtractTableClusteringFields(r *Table, o *TableClustering) error {
	return nil
}
func postReadExtractTableViewFields(r *Table, o *TableView) error {
	return nil
}
func postReadExtractTableViewUserDefinedFunctionResourcesFields(r *Table, o *TableViewUserDefinedFunctionResources) error {
	return nil
}
func postReadExtractTableMaterializedViewFields(r *Table, o *TableMaterializedView) error {
	return nil
}
func postReadExtractTableExternalDataConfigurationFields(r *Table, o *TableExternalDataConfiguration) error {
	vSchema := o.Schema
	if vSchema == nil {
		// note: explicitly not the empty object.
		vSchema = &TableExternalDataConfigurationSchema{}
	}
	if err := extractTableExternalDataConfigurationSchemaFields(r, vSchema); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vSchema) {
		o.Schema = vSchema
	}
	vCsvOptions := o.CsvOptions
	if vCsvOptions == nil {
		// note: explicitly not the empty object.
		vCsvOptions = &TableExternalDataConfigurationCsvOptions{}
	}
	if err := extractTableExternalDataConfigurationCsvOptionsFields(r, vCsvOptions); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vCsvOptions) {
		o.CsvOptions = vCsvOptions
	}
	vBigtableOptions := o.BigtableOptions
	if vBigtableOptions == nil {
		// note: explicitly not the empty object.
		vBigtableOptions = &TableExternalDataConfigurationBigtableOptions{}
	}
	if err := extractTableExternalDataConfigurationBigtableOptionsFields(r, vBigtableOptions); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vBigtableOptions) {
		o.BigtableOptions = vBigtableOptions
	}
	vGoogleSheetsOptions := o.GoogleSheetsOptions
	if vGoogleSheetsOptions == nil {
		// note: explicitly not the empty object.
		vGoogleSheetsOptions = &TableExternalDataConfigurationGoogleSheetsOptions{}
	}
	if err := extractTableExternalDataConfigurationGoogleSheetsOptionsFields(r, vGoogleSheetsOptions); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vGoogleSheetsOptions) {
		o.GoogleSheetsOptions = vGoogleSheetsOptions
	}
	vHivePartitioningOptions := o.HivePartitioningOptions
	if vHivePartitioningOptions == nil {
		// note: explicitly not the empty object.
		vHivePartitioningOptions = &TableExternalDataConfigurationHivePartitioningOptions{}
	}
	if err := extractTableExternalDataConfigurationHivePartitioningOptionsFields(r, vHivePartitioningOptions); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vHivePartitioningOptions) {
		o.HivePartitioningOptions = vHivePartitioningOptions
	}
	vValueConversionModes := o.ValueConversionModes
	if vValueConversionModes == nil {
		// note: explicitly not the empty object.
		vValueConversionModes = &TableExternalDataConfigurationValueConversionModes{}
	}
	if err := extractTableExternalDataConfigurationValueConversionModesFields(r, vValueConversionModes); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vValueConversionModes) {
		o.ValueConversionModes = vValueConversionModes
	}
	vAvroOptions := o.AvroOptions
	if vAvroOptions == nil {
		// note: explicitly not the empty object.
		vAvroOptions = &TableExternalDataConfigurationAvroOptions{}
	}
	if err := extractTableExternalDataConfigurationAvroOptionsFields(r, vAvroOptions); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vAvroOptions) {
		o.AvroOptions = vAvroOptions
	}
	vParquetOptions := o.ParquetOptions
	if vParquetOptions == nil {
		// note: explicitly not the empty object.
		vParquetOptions = &TableExternalDataConfigurationParquetOptions{}
	}
	if err := extractTableExternalDataConfigurationParquetOptionsFields(r, vParquetOptions); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vParquetOptions) {
		o.ParquetOptions = vParquetOptions
	}
	return nil
}
func postReadExtractTableExternalDataConfigurationSchemaFields(r *Table, o *TableExternalDataConfigurationSchema) error {
	return nil
}
func postReadExtractTableExternalDataConfigurationCsvOptionsFields(r *Table, o *TableExternalDataConfigurationCsvOptions) error {
	return nil
}
func postReadExtractTableExternalDataConfigurationBigtableOptionsFields(r *Table, o *TableExternalDataConfigurationBigtableOptions) error {
	return nil
}
func postReadExtractTableExternalDataConfigurationBigtableOptionsColumnFamiliesFields(r *Table, o *TableExternalDataConfigurationBigtableOptionsColumnFamilies) error {
	return nil
}
func postReadExtractTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsFields(r *Table, o *TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns) error {
	return nil
}
func postReadExtractTableExternalDataConfigurationGoogleSheetsOptionsFields(r *Table, o *TableExternalDataConfigurationGoogleSheetsOptions) error {
	return nil
}
func postReadExtractTableExternalDataConfigurationHivePartitioningOptionsFields(r *Table, o *TableExternalDataConfigurationHivePartitioningOptions) error {
	return nil
}
func postReadExtractTableExternalDataConfigurationValueConversionModesFields(r *Table, o *TableExternalDataConfigurationValueConversionModes) error {
	return nil
}
func postReadExtractTableExternalDataConfigurationAvroOptionsFields(r *Table, o *TableExternalDataConfigurationAvroOptions) error {
	return nil
}
func postReadExtractTableExternalDataConfigurationParquetOptionsFields(r *Table, o *TableExternalDataConfigurationParquetOptions) error {
	return nil
}
func postReadExtractTableStreamingBufferFields(r *Table, o *TableStreamingBuffer) error {
	return nil
}
func postReadExtractTableEncryptionConfigurationFields(r *Table, o *TableEncryptionConfiguration) error {
	return nil
}
func postReadExtractTableSnapshotDefinitionFields(r *Table, o *TableSnapshotDefinition) error {
	return nil
}
