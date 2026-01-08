// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/openlayer-ai/openlayer-go/internal/apijson"
	"github.com/openlayer-ai/openlayer-go/internal/apiquery"
	"github.com/openlayer-ai/openlayer-go/internal/param"
	"github.com/openlayer-ai/openlayer-go/internal/requestconfig"
	"github.com/openlayer-ai/openlayer-go/option"
	"github.com/tidwall/gjson"
)

// ProjectInferencePipelineService contains methods and other services that help
// with interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectInferencePipelineService] method instead.
type ProjectInferencePipelineService struct {
	Options []option.RequestOption
}

// NewProjectInferencePipelineService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewProjectInferencePipelineService(opts ...option.RequestOption) (r *ProjectInferencePipelineService) {
	r = &ProjectInferencePipelineService{}
	r.Options = opts
	return
}

// Create an inference pipeline in a project.
func (r *ProjectInferencePipelineService) New(ctx context.Context, projectID string, body ProjectInferencePipelineNewParams, opts ...option.RequestOption) (res *ProjectInferencePipelineNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return
	}
	path := fmt.Sprintf("projects/%s/inference-pipelines", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List the inference pipelines in a project.
func (r *ProjectInferencePipelineService) List(ctx context.Context, projectID string, query ProjectInferencePipelineListParams, opts ...option.RequestOption) (res *ProjectInferencePipelineListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return
	}
	path := fmt.Sprintf("projects/%s/inference-pipelines", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ProjectInferencePipelineNewResponse struct {
	// The inference pipeline id.
	ID string `json:"id,required" format:"uuid"`
	// The creation date.
	DateCreated time.Time `json:"dateCreated,required" format:"date-time"`
	// The last test evaluation date.
	DateLastEvaluated time.Time `json:"dateLastEvaluated,required,nullable" format:"date-time"`
	// The last data sample received date.
	DateLastSampleReceived time.Time `json:"dateLastSampleReceived,required,nullable" format:"date-time"`
	// The next test evaluation date.
	DateOfNextEvaluation time.Time `json:"dateOfNextEvaluation,required,nullable" format:"date-time"`
	// The last updated date.
	DateUpdated time.Time `json:"dateUpdated,required" format:"date-time"`
	// The inference pipeline description.
	Description string `json:"description,required,nullable"`
	// The number of tests failing.
	FailingGoalCount int64                                    `json:"failingGoalCount,required"`
	Links            ProjectInferencePipelineNewResponseLinks `json:"links,required"`
	// The inference pipeline name.
	Name string `json:"name,required"`
	// The number of tests passing.
	PassingGoalCount int64 `json:"passingGoalCount,required"`
	// The project id.
	ProjectID string `json:"projectId,required" format:"uuid"`
	// The status of test evaluation for the inference pipeline.
	Status ProjectInferencePipelineNewResponseStatus `json:"status,required"`
	// The status message of test evaluation for the inference pipeline.
	StatusMessage string `json:"statusMessage,required,nullable"`
	// The total number of tests.
	TotalGoalCount int64                                          `json:"totalGoalCount,required"`
	DataBackend    ProjectInferencePipelineNewResponseDataBackend `json:"dataBackend,nullable"`
	// The last time the data was polled.
	DateLastPolled time.Time                                  `json:"dateLastPolled,nullable" format:"date-time"`
	Project        ProjectInferencePipelineNewResponseProject `json:"project,nullable"`
	// The total number of records in the data backend.
	TotalRecordsCount int64                                        `json:"totalRecordsCount,nullable"`
	Workspace         ProjectInferencePipelineNewResponseWorkspace `json:"workspace,nullable"`
	// The workspace id.
	WorkspaceID string                                  `json:"workspaceId" format:"uuid"`
	JSON        projectInferencePipelineNewResponseJSON `json:"-"`
}

// projectInferencePipelineNewResponseJSON contains the JSON metadata for the
// struct [ProjectInferencePipelineNewResponse]
type projectInferencePipelineNewResponseJSON struct {
	ID                     apijson.Field
	DateCreated            apijson.Field
	DateLastEvaluated      apijson.Field
	DateLastSampleReceived apijson.Field
	DateOfNextEvaluation   apijson.Field
	DateUpdated            apijson.Field
	Description            apijson.Field
	FailingGoalCount       apijson.Field
	Links                  apijson.Field
	Name                   apijson.Field
	PassingGoalCount       apijson.Field
	ProjectID              apijson.Field
	Status                 apijson.Field
	StatusMessage          apijson.Field
	TotalGoalCount         apijson.Field
	DataBackend            apijson.Field
	DateLastPolled         apijson.Field
	Project                apijson.Field
	TotalRecordsCount      apijson.Field
	Workspace              apijson.Field
	WorkspaceID            apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ProjectInferencePipelineNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineNewResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectInferencePipelineNewResponseLinks struct {
	App  string                                       `json:"app,required"`
	JSON projectInferencePipelineNewResponseLinksJSON `json:"-"`
}

// projectInferencePipelineNewResponseLinksJSON contains the JSON metadata for the
// struct [ProjectInferencePipelineNewResponseLinks]
type projectInferencePipelineNewResponseLinksJSON struct {
	App         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectInferencePipelineNewResponseLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineNewResponseLinksJSON) RawJSON() string {
	return r.raw
}

// The status of test evaluation for the inference pipeline.
type ProjectInferencePipelineNewResponseStatus string

const (
	ProjectInferencePipelineNewResponseStatusQueued    ProjectInferencePipelineNewResponseStatus = "queued"
	ProjectInferencePipelineNewResponseStatusRunning   ProjectInferencePipelineNewResponseStatus = "running"
	ProjectInferencePipelineNewResponseStatusPaused    ProjectInferencePipelineNewResponseStatus = "paused"
	ProjectInferencePipelineNewResponseStatusFailed    ProjectInferencePipelineNewResponseStatus = "failed"
	ProjectInferencePipelineNewResponseStatusCompleted ProjectInferencePipelineNewResponseStatus = "completed"
	ProjectInferencePipelineNewResponseStatusUnknown   ProjectInferencePipelineNewResponseStatus = "unknown"
)

func (r ProjectInferencePipelineNewResponseStatus) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewResponseStatusQueued, ProjectInferencePipelineNewResponseStatusRunning, ProjectInferencePipelineNewResponseStatusPaused, ProjectInferencePipelineNewResponseStatusFailed, ProjectInferencePipelineNewResponseStatusCompleted, ProjectInferencePipelineNewResponseStatusUnknown:
		return true
	}
	return false
}

type ProjectInferencePipelineNewResponseDataBackend struct {
	BackendType          ProjectInferencePipelineNewResponseDataBackendBackendType `json:"backendType,required"`
	BigqueryConnectionID string                                                    `json:"bigqueryConnectionId,nullable" format:"uuid"`
	// This field can have the runtime type of
	// [ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackendConfig],
	// [ProjectInferencePipelineNewResponseDataBackendSnowflakeDataBackendConfig],
	// [ProjectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackendConfig],
	// [ProjectInferencePipelineNewResponseDataBackendRedshiftDataBackendConfig],
	// [ProjectInferencePipelineNewResponseDataBackendPostgresDataBackendConfig].
	Config                    interface{}                                                 `json:"config"`
	Database                  string                                                      `json:"database"`
	DatabricksDtlConnectionID string                                                      `json:"databricksDtlConnectionId,nullable" format:"uuid"`
	DatasetID                 string                                                      `json:"datasetId"`
	PartitionType             ProjectInferencePipelineNewResponseDataBackendPartitionType `json:"partitionType,nullable"`
	PostgresConnectionID      string                                                      `json:"postgresConnectionId,nullable" format:"uuid"`
	ProjectID                 string                                                      `json:"projectId"`
	RedshiftConnectionID      string                                                      `json:"redshiftConnectionId,nullable" format:"uuid"`
	Schema                    string                                                      `json:"schema"`
	SchemaName                string                                                      `json:"schemaName"`
	SnowflakeConnectionID     string                                                      `json:"snowflakeConnectionId,nullable" format:"uuid"`
	Table                     string                                                      `json:"table,nullable"`
	TableID                   string                                                      `json:"tableId,nullable"`
	TableName                 string                                                      `json:"tableName"`
	JSON                      projectInferencePipelineNewResponseDataBackendJSON          `json:"-"`
	union                     ProjectInferencePipelineNewResponseDataBackendUnion
}

// projectInferencePipelineNewResponseDataBackendJSON contains the JSON metadata
// for the struct [ProjectInferencePipelineNewResponseDataBackend]
type projectInferencePipelineNewResponseDataBackendJSON struct {
	BackendType               apijson.Field
	BigqueryConnectionID      apijson.Field
	Config                    apijson.Field
	Database                  apijson.Field
	DatabricksDtlConnectionID apijson.Field
	DatasetID                 apijson.Field
	PartitionType             apijson.Field
	PostgresConnectionID      apijson.Field
	ProjectID                 apijson.Field
	RedshiftConnectionID      apijson.Field
	Schema                    apijson.Field
	SchemaName                apijson.Field
	SnowflakeConnectionID     apijson.Field
	Table                     apijson.Field
	TableID                   apijson.Field
	TableName                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r projectInferencePipelineNewResponseDataBackendJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectInferencePipelineNewResponseDataBackend) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectInferencePipelineNewResponseDataBackend{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectInferencePipelineNewResponseDataBackendUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackend],
// [ProjectInferencePipelineNewResponseDataBackendDefaultDataBackend],
// [ProjectInferencePipelineNewResponseDataBackendSnowflakeDataBackend],
// [ProjectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackend],
// [ProjectInferencePipelineNewResponseDataBackendRedshiftDataBackend],
// [ProjectInferencePipelineNewResponseDataBackendPostgresDataBackend].
func (r ProjectInferencePipelineNewResponseDataBackend) AsUnion() ProjectInferencePipelineNewResponseDataBackendUnion {
	return r.union
}

// Union satisfied by
// [ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackend],
// [ProjectInferencePipelineNewResponseDataBackendDefaultDataBackend],
// [ProjectInferencePipelineNewResponseDataBackendSnowflakeDataBackend],
// [ProjectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackend],
// [ProjectInferencePipelineNewResponseDataBackendRedshiftDataBackend] or
// [ProjectInferencePipelineNewResponseDataBackendPostgresDataBackend].
type ProjectInferencePipelineNewResponseDataBackendUnion interface {
	implementsProjectInferencePipelineNewResponseDataBackend()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectInferencePipelineNewResponseDataBackendUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackend{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineNewResponseDataBackendDefaultDataBackend{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineNewResponseDataBackendSnowflakeDataBackend{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackend{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineNewResponseDataBackendRedshiftDataBackend{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineNewResponseDataBackendPostgresDataBackend{}),
		},
	)
}

type ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackend struct {
	BackendType          ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackendBackendType   `json:"backendType,required"`
	BigqueryConnectionID string                                                                         `json:"bigqueryConnectionId,required,nullable" format:"uuid"`
	DatasetID            string                                                                         `json:"datasetId,required"`
	ProjectID            string                                                                         `json:"projectId,required"`
	TableID              string                                                                         `json:"tableId,required,nullable"`
	PartitionType        ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackendPartitionType `json:"partitionType,nullable"`
	JSON                 projectInferencePipelineNewResponseDataBackendBigQueryDataBackendJSON          `json:"-"`
}

// projectInferencePipelineNewResponseDataBackendBigQueryDataBackendJSON contains
// the JSON metadata for the struct
// [ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackend]
type projectInferencePipelineNewResponseDataBackendBigQueryDataBackendJSON struct {
	BackendType          apijson.Field
	BigqueryConnectionID apijson.Field
	DatasetID            apijson.Field
	ProjectID            apijson.Field
	TableID              apijson.Field
	PartitionType        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackend) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineNewResponseDataBackendBigQueryDataBackendJSON) RawJSON() string {
	return r.raw
}

func (r ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackend) implementsProjectInferencePipelineNewResponseDataBackend() {
}

type ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackendBackendType string

const (
	ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackendBackendTypeBigquery ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackendBackendType = "bigquery"
)

func (r ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackendBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackendBackendTypeBigquery:
		return true
	}
	return false
}

type ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackendConfig struct {
	// Name of the column with the ground truths.
	GroundTruthColumnName string `json:"groundTruthColumnName,nullable"`
	// Name of the column with human feedback.
	HumanFeedbackColumnName string `json:"humanFeedbackColumnName,nullable"`
	// Name of the column with the latencies.
	LatencyColumnName string `json:"latencyColumnName,nullable"`
	// Name of the column with the timestamps. Timestamps must be in UNIX sec format.
	// If not provided, the upload timestamp is used.
	TimestampColumnName string                                                                      `json:"timestampColumnName,nullable"`
	JSON                projectInferencePipelineNewResponseDataBackendBigQueryDataBackendConfigJSON `json:"-"`
}

// projectInferencePipelineNewResponseDataBackendBigQueryDataBackendConfigJSON
// contains the JSON metadata for the struct
// [ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackendConfig]
type projectInferencePipelineNewResponseDataBackendBigQueryDataBackendConfigJSON struct {
	GroundTruthColumnName   apijson.Field
	HumanFeedbackColumnName apijson.Field
	LatencyColumnName       apijson.Field
	TimestampColumnName     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackendConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineNewResponseDataBackendBigQueryDataBackendConfigJSON) RawJSON() string {
	return r.raw
}

type ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackendPartitionType string

const (
	ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackendPartitionTypeDay   ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackendPartitionType = "DAY"
	ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackendPartitionTypeMonth ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackendPartitionType = "MONTH"
	ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackendPartitionTypeYear  ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackendPartitionType = "YEAR"
)

func (r ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackendPartitionType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackendPartitionTypeDay, ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackendPartitionTypeMonth, ProjectInferencePipelineNewResponseDataBackendBigQueryDataBackendPartitionTypeYear:
		return true
	}
	return false
}

type ProjectInferencePipelineNewResponseDataBackendDefaultDataBackend struct {
	BackendType ProjectInferencePipelineNewResponseDataBackendDefaultDataBackendBackendType `json:"backendType,required"`
	JSON        projectInferencePipelineNewResponseDataBackendDefaultDataBackendJSON        `json:"-"`
}

// projectInferencePipelineNewResponseDataBackendDefaultDataBackendJSON contains
// the JSON metadata for the struct
// [ProjectInferencePipelineNewResponseDataBackendDefaultDataBackend]
type projectInferencePipelineNewResponseDataBackendDefaultDataBackendJSON struct {
	BackendType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectInferencePipelineNewResponseDataBackendDefaultDataBackend) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineNewResponseDataBackendDefaultDataBackendJSON) RawJSON() string {
	return r.raw
}

func (r ProjectInferencePipelineNewResponseDataBackendDefaultDataBackend) implementsProjectInferencePipelineNewResponseDataBackend() {
}

type ProjectInferencePipelineNewResponseDataBackendDefaultDataBackendBackendType string

const (
	ProjectInferencePipelineNewResponseDataBackendDefaultDataBackendBackendTypeDefault ProjectInferencePipelineNewResponseDataBackendDefaultDataBackendBackendType = "default"
)

func (r ProjectInferencePipelineNewResponseDataBackendDefaultDataBackendBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewResponseDataBackendDefaultDataBackendBackendTypeDefault:
		return true
	}
	return false
}

type ProjectInferencePipelineNewResponseDataBackendSnowflakeDataBackend struct {
	BackendType           ProjectInferencePipelineNewResponseDataBackendSnowflakeDataBackendBackendType `json:"backendType,required"`
	Database              string                                                                        `json:"database,required"`
	Schema                string                                                                        `json:"schema,required"`
	SnowflakeConnectionID string                                                                        `json:"snowflakeConnectionId,required,nullable" format:"uuid"`
	Table                 string                                                                        `json:"table,required,nullable"`
	JSON                  projectInferencePipelineNewResponseDataBackendSnowflakeDataBackendJSON        `json:"-"`
}

// projectInferencePipelineNewResponseDataBackendSnowflakeDataBackendJSON contains
// the JSON metadata for the struct
// [ProjectInferencePipelineNewResponseDataBackendSnowflakeDataBackend]
type projectInferencePipelineNewResponseDataBackendSnowflakeDataBackendJSON struct {
	BackendType           apijson.Field
	Database              apijson.Field
	Schema                apijson.Field
	SnowflakeConnectionID apijson.Field
	Table                 apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ProjectInferencePipelineNewResponseDataBackendSnowflakeDataBackend) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineNewResponseDataBackendSnowflakeDataBackendJSON) RawJSON() string {
	return r.raw
}

func (r ProjectInferencePipelineNewResponseDataBackendSnowflakeDataBackend) implementsProjectInferencePipelineNewResponseDataBackend() {
}

type ProjectInferencePipelineNewResponseDataBackendSnowflakeDataBackendBackendType string

const (
	ProjectInferencePipelineNewResponseDataBackendSnowflakeDataBackendBackendTypeSnowflake ProjectInferencePipelineNewResponseDataBackendSnowflakeDataBackendBackendType = "snowflake"
)

func (r ProjectInferencePipelineNewResponseDataBackendSnowflakeDataBackendBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewResponseDataBackendSnowflakeDataBackendBackendTypeSnowflake:
		return true
	}
	return false
}

type ProjectInferencePipelineNewResponseDataBackendSnowflakeDataBackendConfig struct {
	// Name of the column with the ground truths.
	GroundTruthColumnName string `json:"groundTruthColumnName,nullable"`
	// Name of the column with human feedback.
	HumanFeedbackColumnName string `json:"humanFeedbackColumnName,nullable"`
	// Name of the column with the latencies.
	LatencyColumnName string `json:"latencyColumnName,nullable"`
	// Name of the column with the timestamps. Timestamps must be in UNIX sec format.
	// If not provided, the upload timestamp is used.
	TimestampColumnName string                                                                       `json:"timestampColumnName,nullable"`
	JSON                projectInferencePipelineNewResponseDataBackendSnowflakeDataBackendConfigJSON `json:"-"`
}

// projectInferencePipelineNewResponseDataBackendSnowflakeDataBackendConfigJSON
// contains the JSON metadata for the struct
// [ProjectInferencePipelineNewResponseDataBackendSnowflakeDataBackendConfig]
type projectInferencePipelineNewResponseDataBackendSnowflakeDataBackendConfigJSON struct {
	GroundTruthColumnName   apijson.Field
	HumanFeedbackColumnName apijson.Field
	LatencyColumnName       apijson.Field
	TimestampColumnName     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ProjectInferencePipelineNewResponseDataBackendSnowflakeDataBackendConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineNewResponseDataBackendSnowflakeDataBackendConfigJSON) RawJSON() string {
	return r.raw
}

type ProjectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackend struct {
	BackendType               ProjectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackendBackendType `json:"backendType,required"`
	DatabricksDtlConnectionID string                                                                            `json:"databricksDtlConnectionId,required,nullable" format:"uuid"`
	TableID                   string                                                                            `json:"tableId,required,nullable"`
	JSON                      projectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackendJSON        `json:"-"`
}

// projectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackendJSON
// contains the JSON metadata for the struct
// [ProjectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackend]
type projectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackendJSON struct {
	BackendType               apijson.Field
	DatabricksDtlConnectionID apijson.Field
	TableID                   apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *ProjectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackend) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackendJSON) RawJSON() string {
	return r.raw
}

func (r ProjectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackend) implementsProjectInferencePipelineNewResponseDataBackend() {
}

type ProjectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackendBackendType string

const (
	ProjectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackendBackendTypeDatabricksDtl ProjectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackendBackendType = "databricks_dtl"
)

func (r ProjectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackendBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackendBackendTypeDatabricksDtl:
		return true
	}
	return false
}

type ProjectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackendConfig struct {
	// Name of the column with the ground truths.
	GroundTruthColumnName string `json:"groundTruthColumnName,nullable"`
	// Name of the column with human feedback.
	HumanFeedbackColumnName string `json:"humanFeedbackColumnName,nullable"`
	// Name of the column with the latencies.
	LatencyColumnName string `json:"latencyColumnName,nullable"`
	// Name of the column with the timestamps. Timestamps must be in UNIX sec format.
	// If not provided, the upload timestamp is used.
	TimestampColumnName string                                                                           `json:"timestampColumnName,nullable"`
	JSON                projectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackendConfigJSON `json:"-"`
}

// projectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackendConfigJSON
// contains the JSON metadata for the struct
// [ProjectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackendConfig]
type projectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackendConfigJSON struct {
	GroundTruthColumnName   apijson.Field
	HumanFeedbackColumnName apijson.Field
	LatencyColumnName       apijson.Field
	TimestampColumnName     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ProjectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackendConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineNewResponseDataBackendDatabricksDtlDataBackendConfigJSON) RawJSON() string {
	return r.raw
}

type ProjectInferencePipelineNewResponseDataBackendRedshiftDataBackend struct {
	BackendType          ProjectInferencePipelineNewResponseDataBackendRedshiftDataBackendBackendType `json:"backendType,required"`
	RedshiftConnectionID string                                                                       `json:"redshiftConnectionId,required,nullable" format:"uuid"`
	SchemaName           string                                                                       `json:"schemaName,required"`
	TableName            string                                                                       `json:"tableName,required"`
	JSON                 projectInferencePipelineNewResponseDataBackendRedshiftDataBackendJSON        `json:"-"`
}

// projectInferencePipelineNewResponseDataBackendRedshiftDataBackendJSON contains
// the JSON metadata for the struct
// [ProjectInferencePipelineNewResponseDataBackendRedshiftDataBackend]
type projectInferencePipelineNewResponseDataBackendRedshiftDataBackendJSON struct {
	BackendType          apijson.Field
	RedshiftConnectionID apijson.Field
	SchemaName           apijson.Field
	TableName            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectInferencePipelineNewResponseDataBackendRedshiftDataBackend) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineNewResponseDataBackendRedshiftDataBackendJSON) RawJSON() string {
	return r.raw
}

func (r ProjectInferencePipelineNewResponseDataBackendRedshiftDataBackend) implementsProjectInferencePipelineNewResponseDataBackend() {
}

type ProjectInferencePipelineNewResponseDataBackendRedshiftDataBackendBackendType string

const (
	ProjectInferencePipelineNewResponseDataBackendRedshiftDataBackendBackendTypeRedshift ProjectInferencePipelineNewResponseDataBackendRedshiftDataBackendBackendType = "redshift"
)

func (r ProjectInferencePipelineNewResponseDataBackendRedshiftDataBackendBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewResponseDataBackendRedshiftDataBackendBackendTypeRedshift:
		return true
	}
	return false
}

type ProjectInferencePipelineNewResponseDataBackendRedshiftDataBackendConfig struct {
	// Name of the column with the ground truths.
	GroundTruthColumnName string `json:"groundTruthColumnName,nullable"`
	// Name of the column with human feedback.
	HumanFeedbackColumnName string `json:"humanFeedbackColumnName,nullable"`
	// Name of the column with the latencies.
	LatencyColumnName string `json:"latencyColumnName,nullable"`
	// Name of the column with the timestamps. Timestamps must be in UNIX sec format.
	// If not provided, the upload timestamp is used.
	TimestampColumnName string                                                                      `json:"timestampColumnName,nullable"`
	JSON                projectInferencePipelineNewResponseDataBackendRedshiftDataBackendConfigJSON `json:"-"`
}

// projectInferencePipelineNewResponseDataBackendRedshiftDataBackendConfigJSON
// contains the JSON metadata for the struct
// [ProjectInferencePipelineNewResponseDataBackendRedshiftDataBackendConfig]
type projectInferencePipelineNewResponseDataBackendRedshiftDataBackendConfigJSON struct {
	GroundTruthColumnName   apijson.Field
	HumanFeedbackColumnName apijson.Field
	LatencyColumnName       apijson.Field
	TimestampColumnName     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ProjectInferencePipelineNewResponseDataBackendRedshiftDataBackendConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineNewResponseDataBackendRedshiftDataBackendConfigJSON) RawJSON() string {
	return r.raw
}

type ProjectInferencePipelineNewResponseDataBackendPostgresDataBackend struct {
	BackendType          ProjectInferencePipelineNewResponseDataBackendPostgresDataBackendBackendType `json:"backendType,required"`
	Database             string                                                                       `json:"database,required"`
	PostgresConnectionID string                                                                       `json:"postgresConnectionId,required,nullable" format:"uuid"`
	Schema               string                                                                       `json:"schema,required"`
	Table                string                                                                       `json:"table,required,nullable"`
	JSON                 projectInferencePipelineNewResponseDataBackendPostgresDataBackendJSON        `json:"-"`
}

// projectInferencePipelineNewResponseDataBackendPostgresDataBackendJSON contains
// the JSON metadata for the struct
// [ProjectInferencePipelineNewResponseDataBackendPostgresDataBackend]
type projectInferencePipelineNewResponseDataBackendPostgresDataBackendJSON struct {
	BackendType          apijson.Field
	Database             apijson.Field
	PostgresConnectionID apijson.Field
	Schema               apijson.Field
	Table                apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectInferencePipelineNewResponseDataBackendPostgresDataBackend) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineNewResponseDataBackendPostgresDataBackendJSON) RawJSON() string {
	return r.raw
}

func (r ProjectInferencePipelineNewResponseDataBackendPostgresDataBackend) implementsProjectInferencePipelineNewResponseDataBackend() {
}

type ProjectInferencePipelineNewResponseDataBackendPostgresDataBackendBackendType string

const (
	ProjectInferencePipelineNewResponseDataBackendPostgresDataBackendBackendTypePostgres ProjectInferencePipelineNewResponseDataBackendPostgresDataBackendBackendType = "postgres"
)

func (r ProjectInferencePipelineNewResponseDataBackendPostgresDataBackendBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewResponseDataBackendPostgresDataBackendBackendTypePostgres:
		return true
	}
	return false
}

type ProjectInferencePipelineNewResponseDataBackendPostgresDataBackendConfig struct {
	// Name of the column with the ground truths.
	GroundTruthColumnName string `json:"groundTruthColumnName,nullable"`
	// Name of the column with human feedback.
	HumanFeedbackColumnName string `json:"humanFeedbackColumnName,nullable"`
	// Name of the column with the latencies.
	LatencyColumnName string `json:"latencyColumnName,nullable"`
	// Name of the column with the timestamps. Timestamps must be in UNIX sec format.
	// If not provided, the upload timestamp is used.
	TimestampColumnName string                                                                      `json:"timestampColumnName,nullable"`
	JSON                projectInferencePipelineNewResponseDataBackendPostgresDataBackendConfigJSON `json:"-"`
}

// projectInferencePipelineNewResponseDataBackendPostgresDataBackendConfigJSON
// contains the JSON metadata for the struct
// [ProjectInferencePipelineNewResponseDataBackendPostgresDataBackendConfig]
type projectInferencePipelineNewResponseDataBackendPostgresDataBackendConfigJSON struct {
	GroundTruthColumnName   apijson.Field
	HumanFeedbackColumnName apijson.Field
	LatencyColumnName       apijson.Field
	TimestampColumnName     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ProjectInferencePipelineNewResponseDataBackendPostgresDataBackendConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineNewResponseDataBackendPostgresDataBackendConfigJSON) RawJSON() string {
	return r.raw
}

type ProjectInferencePipelineNewResponseDataBackendBackendType string

const (
	ProjectInferencePipelineNewResponseDataBackendBackendTypeBigquery      ProjectInferencePipelineNewResponseDataBackendBackendType = "bigquery"
	ProjectInferencePipelineNewResponseDataBackendBackendTypeDefault       ProjectInferencePipelineNewResponseDataBackendBackendType = "default"
	ProjectInferencePipelineNewResponseDataBackendBackendTypeSnowflake     ProjectInferencePipelineNewResponseDataBackendBackendType = "snowflake"
	ProjectInferencePipelineNewResponseDataBackendBackendTypeDatabricksDtl ProjectInferencePipelineNewResponseDataBackendBackendType = "databricks_dtl"
	ProjectInferencePipelineNewResponseDataBackendBackendTypeRedshift      ProjectInferencePipelineNewResponseDataBackendBackendType = "redshift"
	ProjectInferencePipelineNewResponseDataBackendBackendTypePostgres      ProjectInferencePipelineNewResponseDataBackendBackendType = "postgres"
)

func (r ProjectInferencePipelineNewResponseDataBackendBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewResponseDataBackendBackendTypeBigquery, ProjectInferencePipelineNewResponseDataBackendBackendTypeDefault, ProjectInferencePipelineNewResponseDataBackendBackendTypeSnowflake, ProjectInferencePipelineNewResponseDataBackendBackendTypeDatabricksDtl, ProjectInferencePipelineNewResponseDataBackendBackendTypeRedshift, ProjectInferencePipelineNewResponseDataBackendBackendTypePostgres:
		return true
	}
	return false
}

type ProjectInferencePipelineNewResponseDataBackendPartitionType string

const (
	ProjectInferencePipelineNewResponseDataBackendPartitionTypeDay   ProjectInferencePipelineNewResponseDataBackendPartitionType = "DAY"
	ProjectInferencePipelineNewResponseDataBackendPartitionTypeMonth ProjectInferencePipelineNewResponseDataBackendPartitionType = "MONTH"
	ProjectInferencePipelineNewResponseDataBackendPartitionTypeYear  ProjectInferencePipelineNewResponseDataBackendPartitionType = "YEAR"
)

func (r ProjectInferencePipelineNewResponseDataBackendPartitionType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewResponseDataBackendPartitionTypeDay, ProjectInferencePipelineNewResponseDataBackendPartitionTypeMonth, ProjectInferencePipelineNewResponseDataBackendPartitionTypeYear:
		return true
	}
	return false
}

type ProjectInferencePipelineNewResponseProject struct {
	// The project id.
	ID string `json:"id,required" format:"uuid"`
	// The project creator id.
	CreatorID string `json:"creatorId,required,nullable" format:"uuid"`
	// The project creation date.
	DateCreated time.Time `json:"dateCreated,required" format:"date-time"`
	// The project last updated date.
	DateUpdated time.Time `json:"dateUpdated,required" format:"date-time"`
	// The number of tests in the development mode of the project.
	DevelopmentGoalCount int64 `json:"developmentGoalCount,required"`
	// The total number of tests in the project.
	GoalCount int64 `json:"goalCount,required"`
	// The number of inference pipelines in the project.
	InferencePipelineCount int64 `json:"inferencePipelineCount,required"`
	// Links to the project.
	Links ProjectInferencePipelineNewResponseProjectLinks `json:"links,required"`
	// The number of tests in the monitoring mode of the project.
	MonitoringGoalCount int64 `json:"monitoringGoalCount,required"`
	// The project name.
	Name string `json:"name,required"`
	// The source of the project.
	Source ProjectInferencePipelineNewResponseProjectSource `json:"source,required,nullable"`
	// The task type of the project.
	TaskType ProjectInferencePipelineNewResponseProjectTaskType `json:"taskType,required"`
	// The number of versions (commits) in the project.
	VersionCount int64 `json:"versionCount,required"`
	// The workspace id.
	WorkspaceID string `json:"workspaceId,required,nullable" format:"uuid"`
	// The project description.
	Description string                                            `json:"description,nullable"`
	GitRepo     ProjectInferencePipelineNewResponseProjectGitRepo `json:"gitRepo,nullable"`
	JSON        projectInferencePipelineNewResponseProjectJSON    `json:"-"`
}

// projectInferencePipelineNewResponseProjectJSON contains the JSON metadata for
// the struct [ProjectInferencePipelineNewResponseProject]
type projectInferencePipelineNewResponseProjectJSON struct {
	ID                     apijson.Field
	CreatorID              apijson.Field
	DateCreated            apijson.Field
	DateUpdated            apijson.Field
	DevelopmentGoalCount   apijson.Field
	GoalCount              apijson.Field
	InferencePipelineCount apijson.Field
	Links                  apijson.Field
	MonitoringGoalCount    apijson.Field
	Name                   apijson.Field
	Source                 apijson.Field
	TaskType               apijson.Field
	VersionCount           apijson.Field
	WorkspaceID            apijson.Field
	Description            apijson.Field
	GitRepo                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ProjectInferencePipelineNewResponseProject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineNewResponseProjectJSON) RawJSON() string {
	return r.raw
}

// Links to the project.
type ProjectInferencePipelineNewResponseProjectLinks struct {
	App  string                                              `json:"app,required"`
	JSON projectInferencePipelineNewResponseProjectLinksJSON `json:"-"`
}

// projectInferencePipelineNewResponseProjectLinksJSON contains the JSON metadata
// for the struct [ProjectInferencePipelineNewResponseProjectLinks]
type projectInferencePipelineNewResponseProjectLinksJSON struct {
	App         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectInferencePipelineNewResponseProjectLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineNewResponseProjectLinksJSON) RawJSON() string {
	return r.raw
}

// The source of the project.
type ProjectInferencePipelineNewResponseProjectSource string

const (
	ProjectInferencePipelineNewResponseProjectSourceWeb  ProjectInferencePipelineNewResponseProjectSource = "web"
	ProjectInferencePipelineNewResponseProjectSourceAPI  ProjectInferencePipelineNewResponseProjectSource = "api"
	ProjectInferencePipelineNewResponseProjectSourceNull ProjectInferencePipelineNewResponseProjectSource = "null"
)

func (r ProjectInferencePipelineNewResponseProjectSource) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewResponseProjectSourceWeb, ProjectInferencePipelineNewResponseProjectSourceAPI, ProjectInferencePipelineNewResponseProjectSourceNull:
		return true
	}
	return false
}

// The task type of the project.
type ProjectInferencePipelineNewResponseProjectTaskType string

const (
	ProjectInferencePipelineNewResponseProjectTaskTypeLlmBase               ProjectInferencePipelineNewResponseProjectTaskType = "llm-base"
	ProjectInferencePipelineNewResponseProjectTaskTypeTabularClassification ProjectInferencePipelineNewResponseProjectTaskType = "tabular-classification"
	ProjectInferencePipelineNewResponseProjectTaskTypeTabularRegression     ProjectInferencePipelineNewResponseProjectTaskType = "tabular-regression"
	ProjectInferencePipelineNewResponseProjectTaskTypeTextClassification    ProjectInferencePipelineNewResponseProjectTaskType = "text-classification"
)

func (r ProjectInferencePipelineNewResponseProjectTaskType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewResponseProjectTaskTypeLlmBase, ProjectInferencePipelineNewResponseProjectTaskTypeTabularClassification, ProjectInferencePipelineNewResponseProjectTaskTypeTabularRegression, ProjectInferencePipelineNewResponseProjectTaskTypeTextClassification:
		return true
	}
	return false
}

type ProjectInferencePipelineNewResponseProjectGitRepo struct {
	ID            string                                                `json:"id,required" format:"uuid"`
	DateConnected time.Time                                             `json:"dateConnected,required" format:"date-time"`
	DateUpdated   time.Time                                             `json:"dateUpdated,required" format:"date-time"`
	GitAccountID  string                                                `json:"gitAccountId,required" format:"uuid"`
	GitID         int64                                                 `json:"gitId,required"`
	Name          string                                                `json:"name,required"`
	Private       bool                                                  `json:"private,required"`
	ProjectID     string                                                `json:"projectId,required" format:"uuid"`
	Slug          string                                                `json:"slug,required"`
	URL           string                                                `json:"url,required" format:"url"`
	Branch        string                                                `json:"branch"`
	RootDir       string                                                `json:"rootDir"`
	JSON          projectInferencePipelineNewResponseProjectGitRepoJSON `json:"-"`
}

// projectInferencePipelineNewResponseProjectGitRepoJSON contains the JSON metadata
// for the struct [ProjectInferencePipelineNewResponseProjectGitRepo]
type projectInferencePipelineNewResponseProjectGitRepoJSON struct {
	ID            apijson.Field
	DateConnected apijson.Field
	DateUpdated   apijson.Field
	GitAccountID  apijson.Field
	GitID         apijson.Field
	Name          apijson.Field
	Private       apijson.Field
	ProjectID     apijson.Field
	Slug          apijson.Field
	URL           apijson.Field
	Branch        apijson.Field
	RootDir       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectInferencePipelineNewResponseProjectGitRepo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineNewResponseProjectGitRepoJSON) RawJSON() string {
	return r.raw
}

type ProjectInferencePipelineNewResponseWorkspace struct {
	// The workspace id.
	ID string `json:"id,required" format:"uuid"`
	// The workspace creator id.
	CreatorID string `json:"creatorId,required,nullable" format:"uuid"`
	// The workspace creation date.
	DateCreated time.Time `json:"dateCreated,required" format:"date-time"`
	// The workspace last updated date.
	DateUpdated time.Time `json:"dateUpdated,required" format:"date-time"`
	// The number of invites in the workspace.
	InviteCount int64 `json:"inviteCount,required"`
	// The number of members in the workspace.
	MemberCount int64 `json:"memberCount,required"`
	// The workspace name.
	Name string `json:"name,required"`
	// The end date of the current billing period.
	PeriodEndDate time.Time `json:"periodEndDate,required,nullable" format:"date-time"`
	// The start date of the current billing period.
	PeriodStartDate time.Time `json:"periodStartDate,required,nullable" format:"date-time"`
	// The number of projects in the workspace.
	ProjectCount int64 `json:"projectCount,required"`
	// The workspace slug.
	Slug         string                                                     `json:"slug,required"`
	Status       ProjectInferencePipelineNewResponseWorkspaceStatus         `json:"status,required"`
	MonthlyUsage []ProjectInferencePipelineNewResponseWorkspaceMonthlyUsage `json:"monthlyUsage"`
	// Whether the workspace only allows SAML authentication.
	SAMLOnlyAccess  bool                                             `json:"samlOnlyAccess"`
	WildcardDomains []string                                         `json:"wildcardDomains"`
	JSON            projectInferencePipelineNewResponseWorkspaceJSON `json:"-"`
}

// projectInferencePipelineNewResponseWorkspaceJSON contains the JSON metadata for
// the struct [ProjectInferencePipelineNewResponseWorkspace]
type projectInferencePipelineNewResponseWorkspaceJSON struct {
	ID              apijson.Field
	CreatorID       apijson.Field
	DateCreated     apijson.Field
	DateUpdated     apijson.Field
	InviteCount     apijson.Field
	MemberCount     apijson.Field
	Name            apijson.Field
	PeriodEndDate   apijson.Field
	PeriodStartDate apijson.Field
	ProjectCount    apijson.Field
	Slug            apijson.Field
	Status          apijson.Field
	MonthlyUsage    apijson.Field
	SAMLOnlyAccess  apijson.Field
	WildcardDomains apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProjectInferencePipelineNewResponseWorkspace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineNewResponseWorkspaceJSON) RawJSON() string {
	return r.raw
}

type ProjectInferencePipelineNewResponseWorkspaceStatus string

const (
	ProjectInferencePipelineNewResponseWorkspaceStatusActive            ProjectInferencePipelineNewResponseWorkspaceStatus = "active"
	ProjectInferencePipelineNewResponseWorkspaceStatusPastDue           ProjectInferencePipelineNewResponseWorkspaceStatus = "past_due"
	ProjectInferencePipelineNewResponseWorkspaceStatusUnpaid            ProjectInferencePipelineNewResponseWorkspaceStatus = "unpaid"
	ProjectInferencePipelineNewResponseWorkspaceStatusCanceled          ProjectInferencePipelineNewResponseWorkspaceStatus = "canceled"
	ProjectInferencePipelineNewResponseWorkspaceStatusIncomplete        ProjectInferencePipelineNewResponseWorkspaceStatus = "incomplete"
	ProjectInferencePipelineNewResponseWorkspaceStatusIncompleteExpired ProjectInferencePipelineNewResponseWorkspaceStatus = "incomplete_expired"
	ProjectInferencePipelineNewResponseWorkspaceStatusTrialing          ProjectInferencePipelineNewResponseWorkspaceStatus = "trialing"
	ProjectInferencePipelineNewResponseWorkspaceStatusPaused            ProjectInferencePipelineNewResponseWorkspaceStatus = "paused"
)

func (r ProjectInferencePipelineNewResponseWorkspaceStatus) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewResponseWorkspaceStatusActive, ProjectInferencePipelineNewResponseWorkspaceStatusPastDue, ProjectInferencePipelineNewResponseWorkspaceStatusUnpaid, ProjectInferencePipelineNewResponseWorkspaceStatusCanceled, ProjectInferencePipelineNewResponseWorkspaceStatusIncomplete, ProjectInferencePipelineNewResponseWorkspaceStatusIncompleteExpired, ProjectInferencePipelineNewResponseWorkspaceStatusTrialing, ProjectInferencePipelineNewResponseWorkspaceStatusPaused:
		return true
	}
	return false
}

type ProjectInferencePipelineNewResponseWorkspaceMonthlyUsage struct {
	ExecutionTimeMs int64                                                        `json:"executionTimeMs,nullable"`
	MonthYear       time.Time                                                    `json:"monthYear" format:"date"`
	PredictionCount int64                                                        `json:"predictionCount"`
	JSON            projectInferencePipelineNewResponseWorkspaceMonthlyUsageJSON `json:"-"`
}

// projectInferencePipelineNewResponseWorkspaceMonthlyUsageJSON contains the JSON
// metadata for the struct
// [ProjectInferencePipelineNewResponseWorkspaceMonthlyUsage]
type projectInferencePipelineNewResponseWorkspaceMonthlyUsageJSON struct {
	ExecutionTimeMs apijson.Field
	MonthYear       apijson.Field
	PredictionCount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProjectInferencePipelineNewResponseWorkspaceMonthlyUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineNewResponseWorkspaceMonthlyUsageJSON) RawJSON() string {
	return r.raw
}

type ProjectInferencePipelineListResponse struct {
	Items []ProjectInferencePipelineListResponseItem `json:"items,required"`
	JSON  projectInferencePipelineListResponseJSON   `json:"-"`
}

// projectInferencePipelineListResponseJSON contains the JSON metadata for the
// struct [ProjectInferencePipelineListResponse]
type projectInferencePipelineListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectInferencePipelineListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineListResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectInferencePipelineListResponseItem struct {
	// The inference pipeline id.
	ID string `json:"id,required" format:"uuid"`
	// The creation date.
	DateCreated time.Time `json:"dateCreated,required" format:"date-time"`
	// The last test evaluation date.
	DateLastEvaluated time.Time `json:"dateLastEvaluated,required,nullable" format:"date-time"`
	// The last data sample received date.
	DateLastSampleReceived time.Time `json:"dateLastSampleReceived,required,nullable" format:"date-time"`
	// The next test evaluation date.
	DateOfNextEvaluation time.Time `json:"dateOfNextEvaluation,required,nullable" format:"date-time"`
	// The last updated date.
	DateUpdated time.Time `json:"dateUpdated,required" format:"date-time"`
	// The inference pipeline description.
	Description string `json:"description,required,nullable"`
	// The number of tests failing.
	FailingGoalCount int64                                          `json:"failingGoalCount,required"`
	Links            ProjectInferencePipelineListResponseItemsLinks `json:"links,required"`
	// The inference pipeline name.
	Name string `json:"name,required"`
	// The number of tests passing.
	PassingGoalCount int64 `json:"passingGoalCount,required"`
	// The project id.
	ProjectID string `json:"projectId,required" format:"uuid"`
	// The status of test evaluation for the inference pipeline.
	Status ProjectInferencePipelineListResponseItemsStatus `json:"status,required"`
	// The status message of test evaluation for the inference pipeline.
	StatusMessage string `json:"statusMessage,required,nullable"`
	// The total number of tests.
	TotalGoalCount int64                                                `json:"totalGoalCount,required"`
	DataBackend    ProjectInferencePipelineListResponseItemsDataBackend `json:"dataBackend,nullable"`
	// The last time the data was polled.
	DateLastPolled time.Time                                        `json:"dateLastPolled,nullable" format:"date-time"`
	Project        ProjectInferencePipelineListResponseItemsProject `json:"project,nullable"`
	// The total number of records in the data backend.
	TotalRecordsCount int64                                              `json:"totalRecordsCount,nullable"`
	Workspace         ProjectInferencePipelineListResponseItemsWorkspace `json:"workspace,nullable"`
	// The workspace id.
	WorkspaceID string                                       `json:"workspaceId" format:"uuid"`
	JSON        projectInferencePipelineListResponseItemJSON `json:"-"`
}

// projectInferencePipelineListResponseItemJSON contains the JSON metadata for the
// struct [ProjectInferencePipelineListResponseItem]
type projectInferencePipelineListResponseItemJSON struct {
	ID                     apijson.Field
	DateCreated            apijson.Field
	DateLastEvaluated      apijson.Field
	DateLastSampleReceived apijson.Field
	DateOfNextEvaluation   apijson.Field
	DateUpdated            apijson.Field
	Description            apijson.Field
	FailingGoalCount       apijson.Field
	Links                  apijson.Field
	Name                   apijson.Field
	PassingGoalCount       apijson.Field
	ProjectID              apijson.Field
	Status                 apijson.Field
	StatusMessage          apijson.Field
	TotalGoalCount         apijson.Field
	DataBackend            apijson.Field
	DateLastPolled         apijson.Field
	Project                apijson.Field
	TotalRecordsCount      apijson.Field
	Workspace              apijson.Field
	WorkspaceID            apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ProjectInferencePipelineListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineListResponseItemJSON) RawJSON() string {
	return r.raw
}

type ProjectInferencePipelineListResponseItemsLinks struct {
	App  string                                             `json:"app,required"`
	JSON projectInferencePipelineListResponseItemsLinksJSON `json:"-"`
}

// projectInferencePipelineListResponseItemsLinksJSON contains the JSON metadata
// for the struct [ProjectInferencePipelineListResponseItemsLinks]
type projectInferencePipelineListResponseItemsLinksJSON struct {
	App         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectInferencePipelineListResponseItemsLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineListResponseItemsLinksJSON) RawJSON() string {
	return r.raw
}

// The status of test evaluation for the inference pipeline.
type ProjectInferencePipelineListResponseItemsStatus string

const (
	ProjectInferencePipelineListResponseItemsStatusQueued    ProjectInferencePipelineListResponseItemsStatus = "queued"
	ProjectInferencePipelineListResponseItemsStatusRunning   ProjectInferencePipelineListResponseItemsStatus = "running"
	ProjectInferencePipelineListResponseItemsStatusPaused    ProjectInferencePipelineListResponseItemsStatus = "paused"
	ProjectInferencePipelineListResponseItemsStatusFailed    ProjectInferencePipelineListResponseItemsStatus = "failed"
	ProjectInferencePipelineListResponseItemsStatusCompleted ProjectInferencePipelineListResponseItemsStatus = "completed"
	ProjectInferencePipelineListResponseItemsStatusUnknown   ProjectInferencePipelineListResponseItemsStatus = "unknown"
)

func (r ProjectInferencePipelineListResponseItemsStatus) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineListResponseItemsStatusQueued, ProjectInferencePipelineListResponseItemsStatusRunning, ProjectInferencePipelineListResponseItemsStatusPaused, ProjectInferencePipelineListResponseItemsStatusFailed, ProjectInferencePipelineListResponseItemsStatusCompleted, ProjectInferencePipelineListResponseItemsStatusUnknown:
		return true
	}
	return false
}

type ProjectInferencePipelineListResponseItemsDataBackend struct {
	BackendType          ProjectInferencePipelineListResponseItemsDataBackendBackendType `json:"backendType,required"`
	BigqueryConnectionID string                                                          `json:"bigqueryConnectionId,nullable" format:"uuid"`
	// This field can have the runtime type of
	// [ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendConfig],
	// [ProjectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackendConfig],
	// [ProjectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackendConfig],
	// [ProjectInferencePipelineListResponseItemsDataBackendRedshiftDataBackendConfig],
	// [ProjectInferencePipelineListResponseItemsDataBackendPostgresDataBackendConfig].
	Config                    interface{}                                                       `json:"config"`
	Database                  string                                                            `json:"database"`
	DatabricksDtlConnectionID string                                                            `json:"databricksDtlConnectionId,nullable" format:"uuid"`
	DatasetID                 string                                                            `json:"datasetId"`
	PartitionType             ProjectInferencePipelineListResponseItemsDataBackendPartitionType `json:"partitionType,nullable"`
	PostgresConnectionID      string                                                            `json:"postgresConnectionId,nullable" format:"uuid"`
	ProjectID                 string                                                            `json:"projectId"`
	RedshiftConnectionID      string                                                            `json:"redshiftConnectionId,nullable" format:"uuid"`
	Schema                    string                                                            `json:"schema"`
	SchemaName                string                                                            `json:"schemaName"`
	SnowflakeConnectionID     string                                                            `json:"snowflakeConnectionId,nullable" format:"uuid"`
	Table                     string                                                            `json:"table,nullable"`
	TableID                   string                                                            `json:"tableId,nullable"`
	TableName                 string                                                            `json:"tableName"`
	JSON                      projectInferencePipelineListResponseItemsDataBackendJSON          `json:"-"`
	union                     ProjectInferencePipelineListResponseItemsDataBackendUnion
}

// projectInferencePipelineListResponseItemsDataBackendJSON contains the JSON
// metadata for the struct [ProjectInferencePipelineListResponseItemsDataBackend]
type projectInferencePipelineListResponseItemsDataBackendJSON struct {
	BackendType               apijson.Field
	BigqueryConnectionID      apijson.Field
	Config                    apijson.Field
	Database                  apijson.Field
	DatabricksDtlConnectionID apijson.Field
	DatasetID                 apijson.Field
	PartitionType             apijson.Field
	PostgresConnectionID      apijson.Field
	ProjectID                 apijson.Field
	RedshiftConnectionID      apijson.Field
	Schema                    apijson.Field
	SchemaName                apijson.Field
	SnowflakeConnectionID     apijson.Field
	Table                     apijson.Field
	TableID                   apijson.Field
	TableName                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r projectInferencePipelineListResponseItemsDataBackendJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectInferencePipelineListResponseItemsDataBackend) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectInferencePipelineListResponseItemsDataBackend{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectInferencePipelineListResponseItemsDataBackendUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackend],
// [ProjectInferencePipelineListResponseItemsDataBackendDefaultDataBackend],
// [ProjectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackend],
// [ProjectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackend],
// [ProjectInferencePipelineListResponseItemsDataBackendRedshiftDataBackend],
// [ProjectInferencePipelineListResponseItemsDataBackendPostgresDataBackend].
func (r ProjectInferencePipelineListResponseItemsDataBackend) AsUnion() ProjectInferencePipelineListResponseItemsDataBackendUnion {
	return r.union
}

// Union satisfied by
// [ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackend],
// [ProjectInferencePipelineListResponseItemsDataBackendDefaultDataBackend],
// [ProjectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackend],
// [ProjectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackend],
// [ProjectInferencePipelineListResponseItemsDataBackendRedshiftDataBackend] or
// [ProjectInferencePipelineListResponseItemsDataBackendPostgresDataBackend].
type ProjectInferencePipelineListResponseItemsDataBackendUnion interface {
	implementsProjectInferencePipelineListResponseItemsDataBackend()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectInferencePipelineListResponseItemsDataBackendUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackend{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineListResponseItemsDataBackendDefaultDataBackend{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackend{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackend{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineListResponseItemsDataBackendRedshiftDataBackend{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineListResponseItemsDataBackendPostgresDataBackend{}),
		},
	)
}

type ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackend struct {
	BackendType          ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendBackendType   `json:"backendType,required"`
	BigqueryConnectionID string                                                                               `json:"bigqueryConnectionId,required,nullable" format:"uuid"`
	DatasetID            string                                                                               `json:"datasetId,required"`
	ProjectID            string                                                                               `json:"projectId,required"`
	TableID              string                                                                               `json:"tableId,required,nullable"`
	PartitionType        ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendPartitionType `json:"partitionType,nullable"`
	JSON                 projectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendJSON          `json:"-"`
}

// projectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendJSON
// contains the JSON metadata for the struct
// [ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackend]
type projectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendJSON struct {
	BackendType          apijson.Field
	BigqueryConnectionID apijson.Field
	DatasetID            apijson.Field
	ProjectID            apijson.Field
	TableID              apijson.Field
	PartitionType        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackend) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendJSON) RawJSON() string {
	return r.raw
}

func (r ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackend) implementsProjectInferencePipelineListResponseItemsDataBackend() {
}

type ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendBackendType string

const (
	ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendBackendTypeBigquery ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendBackendType = "bigquery"
)

func (r ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendBackendTypeBigquery:
		return true
	}
	return false
}

type ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendConfig struct {
	// Name of the column with the ground truths.
	GroundTruthColumnName string `json:"groundTruthColumnName,nullable"`
	// Name of the column with human feedback.
	HumanFeedbackColumnName string `json:"humanFeedbackColumnName,nullable"`
	// Name of the column with the latencies.
	LatencyColumnName string `json:"latencyColumnName,nullable"`
	// Name of the column with the timestamps. Timestamps must be in UNIX sec format.
	// If not provided, the upload timestamp is used.
	TimestampColumnName string                                                                            `json:"timestampColumnName,nullable"`
	JSON                projectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendConfigJSON `json:"-"`
}

// projectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendConfigJSON
// contains the JSON metadata for the struct
// [ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendConfig]
type projectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendConfigJSON struct {
	GroundTruthColumnName   apijson.Field
	HumanFeedbackColumnName apijson.Field
	LatencyColumnName       apijson.Field
	TimestampColumnName     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendConfigJSON) RawJSON() string {
	return r.raw
}

type ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendPartitionType string

const (
	ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendPartitionTypeDay   ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendPartitionType = "DAY"
	ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendPartitionTypeMonth ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendPartitionType = "MONTH"
	ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendPartitionTypeYear  ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendPartitionType = "YEAR"
)

func (r ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendPartitionType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendPartitionTypeDay, ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendPartitionTypeMonth, ProjectInferencePipelineListResponseItemsDataBackendBigQueryDataBackendPartitionTypeYear:
		return true
	}
	return false
}

type ProjectInferencePipelineListResponseItemsDataBackendDefaultDataBackend struct {
	BackendType ProjectInferencePipelineListResponseItemsDataBackendDefaultDataBackendBackendType `json:"backendType,required"`
	JSON        projectInferencePipelineListResponseItemsDataBackendDefaultDataBackendJSON        `json:"-"`
}

// projectInferencePipelineListResponseItemsDataBackendDefaultDataBackendJSON
// contains the JSON metadata for the struct
// [ProjectInferencePipelineListResponseItemsDataBackendDefaultDataBackend]
type projectInferencePipelineListResponseItemsDataBackendDefaultDataBackendJSON struct {
	BackendType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectInferencePipelineListResponseItemsDataBackendDefaultDataBackend) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineListResponseItemsDataBackendDefaultDataBackendJSON) RawJSON() string {
	return r.raw
}

func (r ProjectInferencePipelineListResponseItemsDataBackendDefaultDataBackend) implementsProjectInferencePipelineListResponseItemsDataBackend() {
}

type ProjectInferencePipelineListResponseItemsDataBackendDefaultDataBackendBackendType string

const (
	ProjectInferencePipelineListResponseItemsDataBackendDefaultDataBackendBackendTypeDefault ProjectInferencePipelineListResponseItemsDataBackendDefaultDataBackendBackendType = "default"
)

func (r ProjectInferencePipelineListResponseItemsDataBackendDefaultDataBackendBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineListResponseItemsDataBackendDefaultDataBackendBackendTypeDefault:
		return true
	}
	return false
}

type ProjectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackend struct {
	BackendType           ProjectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackendBackendType `json:"backendType,required"`
	Database              string                                                                              `json:"database,required"`
	Schema                string                                                                              `json:"schema,required"`
	SnowflakeConnectionID string                                                                              `json:"snowflakeConnectionId,required,nullable" format:"uuid"`
	Table                 string                                                                              `json:"table,required,nullable"`
	JSON                  projectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackendJSON        `json:"-"`
}

// projectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackendJSON
// contains the JSON metadata for the struct
// [ProjectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackend]
type projectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackendJSON struct {
	BackendType           apijson.Field
	Database              apijson.Field
	Schema                apijson.Field
	SnowflakeConnectionID apijson.Field
	Table                 apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ProjectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackend) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackendJSON) RawJSON() string {
	return r.raw
}

func (r ProjectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackend) implementsProjectInferencePipelineListResponseItemsDataBackend() {
}

type ProjectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackendBackendType string

const (
	ProjectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackendBackendTypeSnowflake ProjectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackendBackendType = "snowflake"
)

func (r ProjectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackendBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackendBackendTypeSnowflake:
		return true
	}
	return false
}

type ProjectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackendConfig struct {
	// Name of the column with the ground truths.
	GroundTruthColumnName string `json:"groundTruthColumnName,nullable"`
	// Name of the column with human feedback.
	HumanFeedbackColumnName string `json:"humanFeedbackColumnName,nullable"`
	// Name of the column with the latencies.
	LatencyColumnName string `json:"latencyColumnName,nullable"`
	// Name of the column with the timestamps. Timestamps must be in UNIX sec format.
	// If not provided, the upload timestamp is used.
	TimestampColumnName string                                                                             `json:"timestampColumnName,nullable"`
	JSON                projectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackendConfigJSON `json:"-"`
}

// projectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackendConfigJSON
// contains the JSON metadata for the struct
// [ProjectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackendConfig]
type projectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackendConfigJSON struct {
	GroundTruthColumnName   apijson.Field
	HumanFeedbackColumnName apijson.Field
	LatencyColumnName       apijson.Field
	TimestampColumnName     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ProjectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackendConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineListResponseItemsDataBackendSnowflakeDataBackendConfigJSON) RawJSON() string {
	return r.raw
}

type ProjectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackend struct {
	BackendType               ProjectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackendBackendType `json:"backendType,required"`
	DatabricksDtlConnectionID string                                                                                  `json:"databricksDtlConnectionId,required,nullable" format:"uuid"`
	TableID                   string                                                                                  `json:"tableId,required,nullable"`
	JSON                      projectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackendJSON        `json:"-"`
}

// projectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackendJSON
// contains the JSON metadata for the struct
// [ProjectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackend]
type projectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackendJSON struct {
	BackendType               apijson.Field
	DatabricksDtlConnectionID apijson.Field
	TableID                   apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *ProjectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackend) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackendJSON) RawJSON() string {
	return r.raw
}

func (r ProjectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackend) implementsProjectInferencePipelineListResponseItemsDataBackend() {
}

type ProjectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackendBackendType string

const (
	ProjectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackendBackendTypeDatabricksDtl ProjectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackendBackendType = "databricks_dtl"
)

func (r ProjectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackendBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackendBackendTypeDatabricksDtl:
		return true
	}
	return false
}

type ProjectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackendConfig struct {
	// Name of the column with the ground truths.
	GroundTruthColumnName string `json:"groundTruthColumnName,nullable"`
	// Name of the column with human feedback.
	HumanFeedbackColumnName string `json:"humanFeedbackColumnName,nullable"`
	// Name of the column with the latencies.
	LatencyColumnName string `json:"latencyColumnName,nullable"`
	// Name of the column with the timestamps. Timestamps must be in UNIX sec format.
	// If not provided, the upload timestamp is used.
	TimestampColumnName string                                                                                 `json:"timestampColumnName,nullable"`
	JSON                projectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackendConfigJSON `json:"-"`
}

// projectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackendConfigJSON
// contains the JSON metadata for the struct
// [ProjectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackendConfig]
type projectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackendConfigJSON struct {
	GroundTruthColumnName   apijson.Field
	HumanFeedbackColumnName apijson.Field
	LatencyColumnName       apijson.Field
	TimestampColumnName     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ProjectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackendConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineListResponseItemsDataBackendDatabricksDtlDataBackendConfigJSON) RawJSON() string {
	return r.raw
}

type ProjectInferencePipelineListResponseItemsDataBackendRedshiftDataBackend struct {
	BackendType          ProjectInferencePipelineListResponseItemsDataBackendRedshiftDataBackendBackendType `json:"backendType,required"`
	RedshiftConnectionID string                                                                             `json:"redshiftConnectionId,required,nullable" format:"uuid"`
	SchemaName           string                                                                             `json:"schemaName,required"`
	TableName            string                                                                             `json:"tableName,required"`
	JSON                 projectInferencePipelineListResponseItemsDataBackendRedshiftDataBackendJSON        `json:"-"`
}

// projectInferencePipelineListResponseItemsDataBackendRedshiftDataBackendJSON
// contains the JSON metadata for the struct
// [ProjectInferencePipelineListResponseItemsDataBackendRedshiftDataBackend]
type projectInferencePipelineListResponseItemsDataBackendRedshiftDataBackendJSON struct {
	BackendType          apijson.Field
	RedshiftConnectionID apijson.Field
	SchemaName           apijson.Field
	TableName            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectInferencePipelineListResponseItemsDataBackendRedshiftDataBackend) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineListResponseItemsDataBackendRedshiftDataBackendJSON) RawJSON() string {
	return r.raw
}

func (r ProjectInferencePipelineListResponseItemsDataBackendRedshiftDataBackend) implementsProjectInferencePipelineListResponseItemsDataBackend() {
}

type ProjectInferencePipelineListResponseItemsDataBackendRedshiftDataBackendBackendType string

const (
	ProjectInferencePipelineListResponseItemsDataBackendRedshiftDataBackendBackendTypeRedshift ProjectInferencePipelineListResponseItemsDataBackendRedshiftDataBackendBackendType = "redshift"
)

func (r ProjectInferencePipelineListResponseItemsDataBackendRedshiftDataBackendBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineListResponseItemsDataBackendRedshiftDataBackendBackendTypeRedshift:
		return true
	}
	return false
}

type ProjectInferencePipelineListResponseItemsDataBackendRedshiftDataBackendConfig struct {
	// Name of the column with the ground truths.
	GroundTruthColumnName string `json:"groundTruthColumnName,nullable"`
	// Name of the column with human feedback.
	HumanFeedbackColumnName string `json:"humanFeedbackColumnName,nullable"`
	// Name of the column with the latencies.
	LatencyColumnName string `json:"latencyColumnName,nullable"`
	// Name of the column with the timestamps. Timestamps must be in UNIX sec format.
	// If not provided, the upload timestamp is used.
	TimestampColumnName string                                                                            `json:"timestampColumnName,nullable"`
	JSON                projectInferencePipelineListResponseItemsDataBackendRedshiftDataBackendConfigJSON `json:"-"`
}

// projectInferencePipelineListResponseItemsDataBackendRedshiftDataBackendConfigJSON
// contains the JSON metadata for the struct
// [ProjectInferencePipelineListResponseItemsDataBackendRedshiftDataBackendConfig]
type projectInferencePipelineListResponseItemsDataBackendRedshiftDataBackendConfigJSON struct {
	GroundTruthColumnName   apijson.Field
	HumanFeedbackColumnName apijson.Field
	LatencyColumnName       apijson.Field
	TimestampColumnName     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ProjectInferencePipelineListResponseItemsDataBackendRedshiftDataBackendConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineListResponseItemsDataBackendRedshiftDataBackendConfigJSON) RawJSON() string {
	return r.raw
}

type ProjectInferencePipelineListResponseItemsDataBackendPostgresDataBackend struct {
	BackendType          ProjectInferencePipelineListResponseItemsDataBackendPostgresDataBackendBackendType `json:"backendType,required"`
	Database             string                                                                             `json:"database,required"`
	PostgresConnectionID string                                                                             `json:"postgresConnectionId,required,nullable" format:"uuid"`
	Schema               string                                                                             `json:"schema,required"`
	Table                string                                                                             `json:"table,required,nullable"`
	JSON                 projectInferencePipelineListResponseItemsDataBackendPostgresDataBackendJSON        `json:"-"`
}

// projectInferencePipelineListResponseItemsDataBackendPostgresDataBackendJSON
// contains the JSON metadata for the struct
// [ProjectInferencePipelineListResponseItemsDataBackendPostgresDataBackend]
type projectInferencePipelineListResponseItemsDataBackendPostgresDataBackendJSON struct {
	BackendType          apijson.Field
	Database             apijson.Field
	PostgresConnectionID apijson.Field
	Schema               apijson.Field
	Table                apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectInferencePipelineListResponseItemsDataBackendPostgresDataBackend) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineListResponseItemsDataBackendPostgresDataBackendJSON) RawJSON() string {
	return r.raw
}

func (r ProjectInferencePipelineListResponseItemsDataBackendPostgresDataBackend) implementsProjectInferencePipelineListResponseItemsDataBackend() {
}

type ProjectInferencePipelineListResponseItemsDataBackendPostgresDataBackendBackendType string

const (
	ProjectInferencePipelineListResponseItemsDataBackendPostgresDataBackendBackendTypePostgres ProjectInferencePipelineListResponseItemsDataBackendPostgresDataBackendBackendType = "postgres"
)

func (r ProjectInferencePipelineListResponseItemsDataBackendPostgresDataBackendBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineListResponseItemsDataBackendPostgresDataBackendBackendTypePostgres:
		return true
	}
	return false
}

type ProjectInferencePipelineListResponseItemsDataBackendPostgresDataBackendConfig struct {
	// Name of the column with the ground truths.
	GroundTruthColumnName string `json:"groundTruthColumnName,nullable"`
	// Name of the column with human feedback.
	HumanFeedbackColumnName string `json:"humanFeedbackColumnName,nullable"`
	// Name of the column with the latencies.
	LatencyColumnName string `json:"latencyColumnName,nullable"`
	// Name of the column with the timestamps. Timestamps must be in UNIX sec format.
	// If not provided, the upload timestamp is used.
	TimestampColumnName string                                                                            `json:"timestampColumnName,nullable"`
	JSON                projectInferencePipelineListResponseItemsDataBackendPostgresDataBackendConfigJSON `json:"-"`
}

// projectInferencePipelineListResponseItemsDataBackendPostgresDataBackendConfigJSON
// contains the JSON metadata for the struct
// [ProjectInferencePipelineListResponseItemsDataBackendPostgresDataBackendConfig]
type projectInferencePipelineListResponseItemsDataBackendPostgresDataBackendConfigJSON struct {
	GroundTruthColumnName   apijson.Field
	HumanFeedbackColumnName apijson.Field
	LatencyColumnName       apijson.Field
	TimestampColumnName     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ProjectInferencePipelineListResponseItemsDataBackendPostgresDataBackendConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineListResponseItemsDataBackendPostgresDataBackendConfigJSON) RawJSON() string {
	return r.raw
}

type ProjectInferencePipelineListResponseItemsDataBackendBackendType string

const (
	ProjectInferencePipelineListResponseItemsDataBackendBackendTypeBigquery      ProjectInferencePipelineListResponseItemsDataBackendBackendType = "bigquery"
	ProjectInferencePipelineListResponseItemsDataBackendBackendTypeDefault       ProjectInferencePipelineListResponseItemsDataBackendBackendType = "default"
	ProjectInferencePipelineListResponseItemsDataBackendBackendTypeSnowflake     ProjectInferencePipelineListResponseItemsDataBackendBackendType = "snowflake"
	ProjectInferencePipelineListResponseItemsDataBackendBackendTypeDatabricksDtl ProjectInferencePipelineListResponseItemsDataBackendBackendType = "databricks_dtl"
	ProjectInferencePipelineListResponseItemsDataBackendBackendTypeRedshift      ProjectInferencePipelineListResponseItemsDataBackendBackendType = "redshift"
	ProjectInferencePipelineListResponseItemsDataBackendBackendTypePostgres      ProjectInferencePipelineListResponseItemsDataBackendBackendType = "postgres"
)

func (r ProjectInferencePipelineListResponseItemsDataBackendBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineListResponseItemsDataBackendBackendTypeBigquery, ProjectInferencePipelineListResponseItemsDataBackendBackendTypeDefault, ProjectInferencePipelineListResponseItemsDataBackendBackendTypeSnowflake, ProjectInferencePipelineListResponseItemsDataBackendBackendTypeDatabricksDtl, ProjectInferencePipelineListResponseItemsDataBackendBackendTypeRedshift, ProjectInferencePipelineListResponseItemsDataBackendBackendTypePostgres:
		return true
	}
	return false
}

type ProjectInferencePipelineListResponseItemsDataBackendPartitionType string

const (
	ProjectInferencePipelineListResponseItemsDataBackendPartitionTypeDay   ProjectInferencePipelineListResponseItemsDataBackendPartitionType = "DAY"
	ProjectInferencePipelineListResponseItemsDataBackendPartitionTypeMonth ProjectInferencePipelineListResponseItemsDataBackendPartitionType = "MONTH"
	ProjectInferencePipelineListResponseItemsDataBackendPartitionTypeYear  ProjectInferencePipelineListResponseItemsDataBackendPartitionType = "YEAR"
)

func (r ProjectInferencePipelineListResponseItemsDataBackendPartitionType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineListResponseItemsDataBackendPartitionTypeDay, ProjectInferencePipelineListResponseItemsDataBackendPartitionTypeMonth, ProjectInferencePipelineListResponseItemsDataBackendPartitionTypeYear:
		return true
	}
	return false
}

type ProjectInferencePipelineListResponseItemsProject struct {
	// The project id.
	ID string `json:"id,required" format:"uuid"`
	// The project creator id.
	CreatorID string `json:"creatorId,required,nullable" format:"uuid"`
	// The project creation date.
	DateCreated time.Time `json:"dateCreated,required" format:"date-time"`
	// The project last updated date.
	DateUpdated time.Time `json:"dateUpdated,required" format:"date-time"`
	// The number of tests in the development mode of the project.
	DevelopmentGoalCount int64 `json:"developmentGoalCount,required"`
	// The total number of tests in the project.
	GoalCount int64 `json:"goalCount,required"`
	// The number of inference pipelines in the project.
	InferencePipelineCount int64 `json:"inferencePipelineCount,required"`
	// Links to the project.
	Links ProjectInferencePipelineListResponseItemsProjectLinks `json:"links,required"`
	// The number of tests in the monitoring mode of the project.
	MonitoringGoalCount int64 `json:"monitoringGoalCount,required"`
	// The project name.
	Name string `json:"name,required"`
	// The source of the project.
	Source ProjectInferencePipelineListResponseItemsProjectSource `json:"source,required,nullable"`
	// The task type of the project.
	TaskType ProjectInferencePipelineListResponseItemsProjectTaskType `json:"taskType,required"`
	// The number of versions (commits) in the project.
	VersionCount int64 `json:"versionCount,required"`
	// The workspace id.
	WorkspaceID string `json:"workspaceId,required,nullable" format:"uuid"`
	// The project description.
	Description string                                                  `json:"description,nullable"`
	GitRepo     ProjectInferencePipelineListResponseItemsProjectGitRepo `json:"gitRepo,nullable"`
	JSON        projectInferencePipelineListResponseItemsProjectJSON    `json:"-"`
}

// projectInferencePipelineListResponseItemsProjectJSON contains the JSON metadata
// for the struct [ProjectInferencePipelineListResponseItemsProject]
type projectInferencePipelineListResponseItemsProjectJSON struct {
	ID                     apijson.Field
	CreatorID              apijson.Field
	DateCreated            apijson.Field
	DateUpdated            apijson.Field
	DevelopmentGoalCount   apijson.Field
	GoalCount              apijson.Field
	InferencePipelineCount apijson.Field
	Links                  apijson.Field
	MonitoringGoalCount    apijson.Field
	Name                   apijson.Field
	Source                 apijson.Field
	TaskType               apijson.Field
	VersionCount           apijson.Field
	WorkspaceID            apijson.Field
	Description            apijson.Field
	GitRepo                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ProjectInferencePipelineListResponseItemsProject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineListResponseItemsProjectJSON) RawJSON() string {
	return r.raw
}

// Links to the project.
type ProjectInferencePipelineListResponseItemsProjectLinks struct {
	App  string                                                    `json:"app,required"`
	JSON projectInferencePipelineListResponseItemsProjectLinksJSON `json:"-"`
}

// projectInferencePipelineListResponseItemsProjectLinksJSON contains the JSON
// metadata for the struct [ProjectInferencePipelineListResponseItemsProjectLinks]
type projectInferencePipelineListResponseItemsProjectLinksJSON struct {
	App         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectInferencePipelineListResponseItemsProjectLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineListResponseItemsProjectLinksJSON) RawJSON() string {
	return r.raw
}

// The source of the project.
type ProjectInferencePipelineListResponseItemsProjectSource string

const (
	ProjectInferencePipelineListResponseItemsProjectSourceWeb  ProjectInferencePipelineListResponseItemsProjectSource = "web"
	ProjectInferencePipelineListResponseItemsProjectSourceAPI  ProjectInferencePipelineListResponseItemsProjectSource = "api"
	ProjectInferencePipelineListResponseItemsProjectSourceNull ProjectInferencePipelineListResponseItemsProjectSource = "null"
)

func (r ProjectInferencePipelineListResponseItemsProjectSource) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineListResponseItemsProjectSourceWeb, ProjectInferencePipelineListResponseItemsProjectSourceAPI, ProjectInferencePipelineListResponseItemsProjectSourceNull:
		return true
	}
	return false
}

// The task type of the project.
type ProjectInferencePipelineListResponseItemsProjectTaskType string

const (
	ProjectInferencePipelineListResponseItemsProjectTaskTypeLlmBase               ProjectInferencePipelineListResponseItemsProjectTaskType = "llm-base"
	ProjectInferencePipelineListResponseItemsProjectTaskTypeTabularClassification ProjectInferencePipelineListResponseItemsProjectTaskType = "tabular-classification"
	ProjectInferencePipelineListResponseItemsProjectTaskTypeTabularRegression     ProjectInferencePipelineListResponseItemsProjectTaskType = "tabular-regression"
	ProjectInferencePipelineListResponseItemsProjectTaskTypeTextClassification    ProjectInferencePipelineListResponseItemsProjectTaskType = "text-classification"
)

func (r ProjectInferencePipelineListResponseItemsProjectTaskType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineListResponseItemsProjectTaskTypeLlmBase, ProjectInferencePipelineListResponseItemsProjectTaskTypeTabularClassification, ProjectInferencePipelineListResponseItemsProjectTaskTypeTabularRegression, ProjectInferencePipelineListResponseItemsProjectTaskTypeTextClassification:
		return true
	}
	return false
}

type ProjectInferencePipelineListResponseItemsProjectGitRepo struct {
	ID            string                                                      `json:"id,required" format:"uuid"`
	DateConnected time.Time                                                   `json:"dateConnected,required" format:"date-time"`
	DateUpdated   time.Time                                                   `json:"dateUpdated,required" format:"date-time"`
	GitAccountID  string                                                      `json:"gitAccountId,required" format:"uuid"`
	GitID         int64                                                       `json:"gitId,required"`
	Name          string                                                      `json:"name,required"`
	Private       bool                                                        `json:"private,required"`
	ProjectID     string                                                      `json:"projectId,required" format:"uuid"`
	Slug          string                                                      `json:"slug,required"`
	URL           string                                                      `json:"url,required" format:"url"`
	Branch        string                                                      `json:"branch"`
	RootDir       string                                                      `json:"rootDir"`
	JSON          projectInferencePipelineListResponseItemsProjectGitRepoJSON `json:"-"`
}

// projectInferencePipelineListResponseItemsProjectGitRepoJSON contains the JSON
// metadata for the struct
// [ProjectInferencePipelineListResponseItemsProjectGitRepo]
type projectInferencePipelineListResponseItemsProjectGitRepoJSON struct {
	ID            apijson.Field
	DateConnected apijson.Field
	DateUpdated   apijson.Field
	GitAccountID  apijson.Field
	GitID         apijson.Field
	Name          apijson.Field
	Private       apijson.Field
	ProjectID     apijson.Field
	Slug          apijson.Field
	URL           apijson.Field
	Branch        apijson.Field
	RootDir       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectInferencePipelineListResponseItemsProjectGitRepo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineListResponseItemsProjectGitRepoJSON) RawJSON() string {
	return r.raw
}

type ProjectInferencePipelineListResponseItemsWorkspace struct {
	// The workspace id.
	ID string `json:"id,required" format:"uuid"`
	// The workspace creator id.
	CreatorID string `json:"creatorId,required,nullable" format:"uuid"`
	// The workspace creation date.
	DateCreated time.Time `json:"dateCreated,required" format:"date-time"`
	// The workspace last updated date.
	DateUpdated time.Time `json:"dateUpdated,required" format:"date-time"`
	// The number of invites in the workspace.
	InviteCount int64 `json:"inviteCount,required"`
	// The number of members in the workspace.
	MemberCount int64 `json:"memberCount,required"`
	// The workspace name.
	Name string `json:"name,required"`
	// The end date of the current billing period.
	PeriodEndDate time.Time `json:"periodEndDate,required,nullable" format:"date-time"`
	// The start date of the current billing period.
	PeriodStartDate time.Time `json:"periodStartDate,required,nullable" format:"date-time"`
	// The number of projects in the workspace.
	ProjectCount int64 `json:"projectCount,required"`
	// The workspace slug.
	Slug         string                                                           `json:"slug,required"`
	Status       ProjectInferencePipelineListResponseItemsWorkspaceStatus         `json:"status,required"`
	MonthlyUsage []ProjectInferencePipelineListResponseItemsWorkspaceMonthlyUsage `json:"monthlyUsage"`
	// Whether the workspace only allows SAML authentication.
	SAMLOnlyAccess  bool                                                   `json:"samlOnlyAccess"`
	WildcardDomains []string                                               `json:"wildcardDomains"`
	JSON            projectInferencePipelineListResponseItemsWorkspaceJSON `json:"-"`
}

// projectInferencePipelineListResponseItemsWorkspaceJSON contains the JSON
// metadata for the struct [ProjectInferencePipelineListResponseItemsWorkspace]
type projectInferencePipelineListResponseItemsWorkspaceJSON struct {
	ID              apijson.Field
	CreatorID       apijson.Field
	DateCreated     apijson.Field
	DateUpdated     apijson.Field
	InviteCount     apijson.Field
	MemberCount     apijson.Field
	Name            apijson.Field
	PeriodEndDate   apijson.Field
	PeriodStartDate apijson.Field
	ProjectCount    apijson.Field
	Slug            apijson.Field
	Status          apijson.Field
	MonthlyUsage    apijson.Field
	SAMLOnlyAccess  apijson.Field
	WildcardDomains apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProjectInferencePipelineListResponseItemsWorkspace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineListResponseItemsWorkspaceJSON) RawJSON() string {
	return r.raw
}

type ProjectInferencePipelineListResponseItemsWorkspaceStatus string

const (
	ProjectInferencePipelineListResponseItemsWorkspaceStatusActive            ProjectInferencePipelineListResponseItemsWorkspaceStatus = "active"
	ProjectInferencePipelineListResponseItemsWorkspaceStatusPastDue           ProjectInferencePipelineListResponseItemsWorkspaceStatus = "past_due"
	ProjectInferencePipelineListResponseItemsWorkspaceStatusUnpaid            ProjectInferencePipelineListResponseItemsWorkspaceStatus = "unpaid"
	ProjectInferencePipelineListResponseItemsWorkspaceStatusCanceled          ProjectInferencePipelineListResponseItemsWorkspaceStatus = "canceled"
	ProjectInferencePipelineListResponseItemsWorkspaceStatusIncomplete        ProjectInferencePipelineListResponseItemsWorkspaceStatus = "incomplete"
	ProjectInferencePipelineListResponseItemsWorkspaceStatusIncompleteExpired ProjectInferencePipelineListResponseItemsWorkspaceStatus = "incomplete_expired"
	ProjectInferencePipelineListResponseItemsWorkspaceStatusTrialing          ProjectInferencePipelineListResponseItemsWorkspaceStatus = "trialing"
	ProjectInferencePipelineListResponseItemsWorkspaceStatusPaused            ProjectInferencePipelineListResponseItemsWorkspaceStatus = "paused"
)

func (r ProjectInferencePipelineListResponseItemsWorkspaceStatus) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineListResponseItemsWorkspaceStatusActive, ProjectInferencePipelineListResponseItemsWorkspaceStatusPastDue, ProjectInferencePipelineListResponseItemsWorkspaceStatusUnpaid, ProjectInferencePipelineListResponseItemsWorkspaceStatusCanceled, ProjectInferencePipelineListResponseItemsWorkspaceStatusIncomplete, ProjectInferencePipelineListResponseItemsWorkspaceStatusIncompleteExpired, ProjectInferencePipelineListResponseItemsWorkspaceStatusTrialing, ProjectInferencePipelineListResponseItemsWorkspaceStatusPaused:
		return true
	}
	return false
}

type ProjectInferencePipelineListResponseItemsWorkspaceMonthlyUsage struct {
	ExecutionTimeMs int64                                                              `json:"executionTimeMs,nullable"`
	MonthYear       time.Time                                                          `json:"monthYear" format:"date"`
	PredictionCount int64                                                              `json:"predictionCount"`
	JSON            projectInferencePipelineListResponseItemsWorkspaceMonthlyUsageJSON `json:"-"`
}

// projectInferencePipelineListResponseItemsWorkspaceMonthlyUsageJSON contains the
// JSON metadata for the struct
// [ProjectInferencePipelineListResponseItemsWorkspaceMonthlyUsage]
type projectInferencePipelineListResponseItemsWorkspaceMonthlyUsageJSON struct {
	ExecutionTimeMs apijson.Field
	MonthYear       apijson.Field
	PredictionCount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProjectInferencePipelineListResponseItemsWorkspaceMonthlyUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineListResponseItemsWorkspaceMonthlyUsageJSON) RawJSON() string {
	return r.raw
}

type ProjectInferencePipelineNewParams struct {
	// The inference pipeline description.
	Description param.Field[string] `json:"description,required"`
	// The inference pipeline name.
	Name        param.Field[string]                                            `json:"name,required"`
	DataBackend param.Field[ProjectInferencePipelineNewParamsDataBackendUnion] `json:"dataBackend"`
	Project     param.Field[ProjectInferencePipelineNewParamsProject]          `json:"project"`
	Workspace   param.Field[ProjectInferencePipelineNewParamsWorkspace]        `json:"workspace"`
}

func (r ProjectInferencePipelineNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectInferencePipelineNewParamsDataBackend struct {
	BackendType               param.Field[ProjectInferencePipelineNewParamsDataBackendBackendType]   `json:"backendType,required"`
	BigqueryConnectionID      param.Field[string]                                                    `json:"bigqueryConnectionId" format:"uuid"`
	Config                    param.Field[interface{}]                                               `json:"config"`
	Database                  param.Field[string]                                                    `json:"database"`
	DatabricksDtlConnectionID param.Field[string]                                                    `json:"databricksDtlConnectionId" format:"uuid"`
	DatasetID                 param.Field[string]                                                    `json:"datasetId"`
	PartitionType             param.Field[ProjectInferencePipelineNewParamsDataBackendPartitionType] `json:"partitionType"`
	PostgresConnectionID      param.Field[string]                                                    `json:"postgresConnectionId" format:"uuid"`
	ProjectID                 param.Field[string]                                                    `json:"projectId"`
	RedshiftConnectionID      param.Field[string]                                                    `json:"redshiftConnectionId" format:"uuid"`
	Schema                    param.Field[string]                                                    `json:"schema"`
	SchemaName                param.Field[string]                                                    `json:"schemaName"`
	SnowflakeConnectionID     param.Field[string]                                                    `json:"snowflakeConnectionId" format:"uuid"`
	Table                     param.Field[string]                                                    `json:"table"`
	TableID                   param.Field[string]                                                    `json:"tableId"`
	TableName                 param.Field[string]                                                    `json:"tableName"`
}

func (r ProjectInferencePipelineNewParamsDataBackend) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectInferencePipelineNewParamsDataBackend) implementsProjectInferencePipelineNewParamsDataBackendUnion() {
}

// Satisfied by [ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackend],
// [ProjectInferencePipelineNewParamsDataBackendDefaultDataBackend],
// [ProjectInferencePipelineNewParamsDataBackendSnowflakeDataBackend],
// [ProjectInferencePipelineNewParamsDataBackendDatabricksDtlDataBackend],
// [ProjectInferencePipelineNewParamsDataBackendRedshiftDataBackend],
// [ProjectInferencePipelineNewParamsDataBackendPostgresDataBackend],
// [ProjectInferencePipelineNewParamsDataBackend].
type ProjectInferencePipelineNewParamsDataBackendUnion interface {
	implementsProjectInferencePipelineNewParamsDataBackendUnion()
}

type ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackend struct {
	BackendType          param.Field[ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackendBackendType]   `json:"backendType,required"`
	BigqueryConnectionID param.Field[string]                                                                       `json:"bigqueryConnectionId,required" format:"uuid"`
	Config               param.Field[ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackendConfig]        `json:"config,required"`
	DatasetID            param.Field[string]                                                                       `json:"datasetId,required"`
	ProjectID            param.Field[string]                                                                       `json:"projectId,required"`
	TableID              param.Field[string]                                                                       `json:"tableId,required"`
	PartitionType        param.Field[ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackendPartitionType] `json:"partitionType"`
}

func (r ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackend) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackend) implementsProjectInferencePipelineNewParamsDataBackendUnion() {
}

type ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackendBackendType string

const (
	ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackendBackendTypeBigquery ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackendBackendType = "bigquery"
)

func (r ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackendBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackendBackendTypeBigquery:
		return true
	}
	return false
}

type ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackendConfig struct {
	// Name of the column with the ground truths.
	GroundTruthColumnName param.Field[string] `json:"groundTruthColumnName"`
	// Name of the column with human feedback.
	HumanFeedbackColumnName param.Field[string] `json:"humanFeedbackColumnName"`
	// Name of the column with the inference ids. This is useful if you want to update
	// rows at a later point in time. If not provided, a unique id is generated by
	// Openlayer.
	InferenceIDColumnName param.Field[string] `json:"inferenceIdColumnName"`
	// Name of the column with the latencies.
	LatencyColumnName param.Field[string] `json:"latencyColumnName"`
	// Name of the column with the timestamps. Timestamps must be in UNIX sec format.
	// If not provided, the upload timestamp is used.
	TimestampColumnName param.Field[string] `json:"timestampColumnName"`
}

func (r ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackendConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackendPartitionType string

const (
	ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackendPartitionTypeDay   ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackendPartitionType = "DAY"
	ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackendPartitionTypeMonth ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackendPartitionType = "MONTH"
	ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackendPartitionTypeYear  ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackendPartitionType = "YEAR"
)

func (r ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackendPartitionType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackendPartitionTypeDay, ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackendPartitionTypeMonth, ProjectInferencePipelineNewParamsDataBackendBigQueryDataBackendPartitionTypeYear:
		return true
	}
	return false
}

type ProjectInferencePipelineNewParamsDataBackendDefaultDataBackend struct {
	BackendType param.Field[ProjectInferencePipelineNewParamsDataBackendDefaultDataBackendBackendType] `json:"backendType,required"`
}

func (r ProjectInferencePipelineNewParamsDataBackendDefaultDataBackend) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectInferencePipelineNewParamsDataBackendDefaultDataBackend) implementsProjectInferencePipelineNewParamsDataBackendUnion() {
}

type ProjectInferencePipelineNewParamsDataBackendDefaultDataBackendBackendType string

const (
	ProjectInferencePipelineNewParamsDataBackendDefaultDataBackendBackendTypeDefault ProjectInferencePipelineNewParamsDataBackendDefaultDataBackendBackendType = "default"
)

func (r ProjectInferencePipelineNewParamsDataBackendDefaultDataBackendBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewParamsDataBackendDefaultDataBackendBackendTypeDefault:
		return true
	}
	return false
}

type ProjectInferencePipelineNewParamsDataBackendSnowflakeDataBackend struct {
	BackendType           param.Field[ProjectInferencePipelineNewParamsDataBackendSnowflakeDataBackendBackendType] `json:"backendType,required"`
	Config                param.Field[ProjectInferencePipelineNewParamsDataBackendSnowflakeDataBackendConfig]      `json:"config,required"`
	Database              param.Field[string]                                                                      `json:"database,required"`
	Schema                param.Field[string]                                                                      `json:"schema,required"`
	SnowflakeConnectionID param.Field[string]                                                                      `json:"snowflakeConnectionId,required" format:"uuid"`
	Table                 param.Field[string]                                                                      `json:"table,required"`
}

func (r ProjectInferencePipelineNewParamsDataBackendSnowflakeDataBackend) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectInferencePipelineNewParamsDataBackendSnowflakeDataBackend) implementsProjectInferencePipelineNewParamsDataBackendUnion() {
}

type ProjectInferencePipelineNewParamsDataBackendSnowflakeDataBackendBackendType string

const (
	ProjectInferencePipelineNewParamsDataBackendSnowflakeDataBackendBackendTypeSnowflake ProjectInferencePipelineNewParamsDataBackendSnowflakeDataBackendBackendType = "snowflake"
)

func (r ProjectInferencePipelineNewParamsDataBackendSnowflakeDataBackendBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewParamsDataBackendSnowflakeDataBackendBackendTypeSnowflake:
		return true
	}
	return false
}

type ProjectInferencePipelineNewParamsDataBackendSnowflakeDataBackendConfig struct {
	// Name of the column with the ground truths.
	GroundTruthColumnName param.Field[string] `json:"groundTruthColumnName"`
	// Name of the column with human feedback.
	HumanFeedbackColumnName param.Field[string] `json:"humanFeedbackColumnName"`
	// Name of the column with the inference ids. This is useful if you want to update
	// rows at a later point in time. If not provided, a unique id is generated by
	// Openlayer.
	InferenceIDColumnName param.Field[string] `json:"inferenceIdColumnName"`
	// Name of the column with the latencies.
	LatencyColumnName param.Field[string] `json:"latencyColumnName"`
	// Name of the column with the timestamps. Timestamps must be in UNIX sec format.
	// If not provided, the upload timestamp is used.
	TimestampColumnName param.Field[string] `json:"timestampColumnName"`
}

func (r ProjectInferencePipelineNewParamsDataBackendSnowflakeDataBackendConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectInferencePipelineNewParamsDataBackendDatabricksDtlDataBackend struct {
	BackendType               param.Field[ProjectInferencePipelineNewParamsDataBackendDatabricksDtlDataBackendBackendType] `json:"backendType,required"`
	Config                    param.Field[ProjectInferencePipelineNewParamsDataBackendDatabricksDtlDataBackendConfig]      `json:"config,required"`
	DatabricksDtlConnectionID param.Field[string]                                                                          `json:"databricksDtlConnectionId,required" format:"uuid"`
	TableID                   param.Field[string]                                                                          `json:"tableId,required"`
}

func (r ProjectInferencePipelineNewParamsDataBackendDatabricksDtlDataBackend) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectInferencePipelineNewParamsDataBackendDatabricksDtlDataBackend) implementsProjectInferencePipelineNewParamsDataBackendUnion() {
}

type ProjectInferencePipelineNewParamsDataBackendDatabricksDtlDataBackendBackendType string

const (
	ProjectInferencePipelineNewParamsDataBackendDatabricksDtlDataBackendBackendTypeDatabricksDtl ProjectInferencePipelineNewParamsDataBackendDatabricksDtlDataBackendBackendType = "databricks_dtl"
)

func (r ProjectInferencePipelineNewParamsDataBackendDatabricksDtlDataBackendBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewParamsDataBackendDatabricksDtlDataBackendBackendTypeDatabricksDtl:
		return true
	}
	return false
}

type ProjectInferencePipelineNewParamsDataBackendDatabricksDtlDataBackendConfig struct {
	// Name of the column with the ground truths.
	GroundTruthColumnName param.Field[string] `json:"groundTruthColumnName"`
	// Name of the column with human feedback.
	HumanFeedbackColumnName param.Field[string] `json:"humanFeedbackColumnName"`
	// Name of the column with the inference ids. This is useful if you want to update
	// rows at a later point in time. If not provided, a unique id is generated by
	// Openlayer.
	InferenceIDColumnName param.Field[string] `json:"inferenceIdColumnName"`
	// Name of the column with the latencies.
	LatencyColumnName param.Field[string] `json:"latencyColumnName"`
	// Name of the column with the timestamps. Timestamps must be in UNIX sec format.
	// If not provided, the upload timestamp is used.
	TimestampColumnName param.Field[string] `json:"timestampColumnName"`
}

func (r ProjectInferencePipelineNewParamsDataBackendDatabricksDtlDataBackendConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectInferencePipelineNewParamsDataBackendRedshiftDataBackend struct {
	BackendType          param.Field[ProjectInferencePipelineNewParamsDataBackendRedshiftDataBackendBackendType] `json:"backendType,required"`
	Config               param.Field[ProjectInferencePipelineNewParamsDataBackendRedshiftDataBackendConfig]      `json:"config,required"`
	RedshiftConnectionID param.Field[string]                                                                     `json:"redshiftConnectionId,required" format:"uuid"`
	SchemaName           param.Field[string]                                                                     `json:"schemaName,required"`
	TableName            param.Field[string]                                                                     `json:"tableName,required"`
}

func (r ProjectInferencePipelineNewParamsDataBackendRedshiftDataBackend) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectInferencePipelineNewParamsDataBackendRedshiftDataBackend) implementsProjectInferencePipelineNewParamsDataBackendUnion() {
}

type ProjectInferencePipelineNewParamsDataBackendRedshiftDataBackendBackendType string

const (
	ProjectInferencePipelineNewParamsDataBackendRedshiftDataBackendBackendTypeRedshift ProjectInferencePipelineNewParamsDataBackendRedshiftDataBackendBackendType = "redshift"
)

func (r ProjectInferencePipelineNewParamsDataBackendRedshiftDataBackendBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewParamsDataBackendRedshiftDataBackendBackendTypeRedshift:
		return true
	}
	return false
}

type ProjectInferencePipelineNewParamsDataBackendRedshiftDataBackendConfig struct {
	// Name of the column with the ground truths.
	GroundTruthColumnName param.Field[string] `json:"groundTruthColumnName"`
	// Name of the column with human feedback.
	HumanFeedbackColumnName param.Field[string] `json:"humanFeedbackColumnName"`
	// Name of the column with the inference ids. This is useful if you want to update
	// rows at a later point in time. If not provided, a unique id is generated by
	// Openlayer.
	InferenceIDColumnName param.Field[string] `json:"inferenceIdColumnName"`
	// Name of the column with the latencies.
	LatencyColumnName param.Field[string] `json:"latencyColumnName"`
	// Name of the column with the timestamps. Timestamps must be in UNIX sec format.
	// If not provided, the upload timestamp is used.
	TimestampColumnName param.Field[string] `json:"timestampColumnName"`
}

func (r ProjectInferencePipelineNewParamsDataBackendRedshiftDataBackendConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectInferencePipelineNewParamsDataBackendPostgresDataBackend struct {
	BackendType          param.Field[ProjectInferencePipelineNewParamsDataBackendPostgresDataBackendBackendType] `json:"backendType,required"`
	Config               param.Field[ProjectInferencePipelineNewParamsDataBackendPostgresDataBackendConfig]      `json:"config,required"`
	Database             param.Field[string]                                                                     `json:"database,required"`
	PostgresConnectionID param.Field[string]                                                                     `json:"postgresConnectionId,required" format:"uuid"`
	Schema               param.Field[string]                                                                     `json:"schema,required"`
	Table                param.Field[string]                                                                     `json:"table,required"`
}

func (r ProjectInferencePipelineNewParamsDataBackendPostgresDataBackend) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectInferencePipelineNewParamsDataBackendPostgresDataBackend) implementsProjectInferencePipelineNewParamsDataBackendUnion() {
}

type ProjectInferencePipelineNewParamsDataBackendPostgresDataBackendBackendType string

const (
	ProjectInferencePipelineNewParamsDataBackendPostgresDataBackendBackendTypePostgres ProjectInferencePipelineNewParamsDataBackendPostgresDataBackendBackendType = "postgres"
)

func (r ProjectInferencePipelineNewParamsDataBackendPostgresDataBackendBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewParamsDataBackendPostgresDataBackendBackendTypePostgres:
		return true
	}
	return false
}

type ProjectInferencePipelineNewParamsDataBackendPostgresDataBackendConfig struct {
	// Name of the column with the ground truths.
	GroundTruthColumnName param.Field[string] `json:"groundTruthColumnName"`
	// Name of the column with human feedback.
	HumanFeedbackColumnName param.Field[string] `json:"humanFeedbackColumnName"`
	// Name of the column with the inference ids. This is useful if you want to update
	// rows at a later point in time. If not provided, a unique id is generated by
	// Openlayer.
	InferenceIDColumnName param.Field[string] `json:"inferenceIdColumnName"`
	// Name of the column with the latencies.
	LatencyColumnName param.Field[string] `json:"latencyColumnName"`
	// Name of the column with the timestamps. Timestamps must be in UNIX sec format.
	// If not provided, the upload timestamp is used.
	TimestampColumnName param.Field[string] `json:"timestampColumnName"`
}

func (r ProjectInferencePipelineNewParamsDataBackendPostgresDataBackendConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectInferencePipelineNewParamsDataBackendBackendType string

const (
	ProjectInferencePipelineNewParamsDataBackendBackendTypeBigquery      ProjectInferencePipelineNewParamsDataBackendBackendType = "bigquery"
	ProjectInferencePipelineNewParamsDataBackendBackendTypeDefault       ProjectInferencePipelineNewParamsDataBackendBackendType = "default"
	ProjectInferencePipelineNewParamsDataBackendBackendTypeSnowflake     ProjectInferencePipelineNewParamsDataBackendBackendType = "snowflake"
	ProjectInferencePipelineNewParamsDataBackendBackendTypeDatabricksDtl ProjectInferencePipelineNewParamsDataBackendBackendType = "databricks_dtl"
	ProjectInferencePipelineNewParamsDataBackendBackendTypeRedshift      ProjectInferencePipelineNewParamsDataBackendBackendType = "redshift"
	ProjectInferencePipelineNewParamsDataBackendBackendTypePostgres      ProjectInferencePipelineNewParamsDataBackendBackendType = "postgres"
)

func (r ProjectInferencePipelineNewParamsDataBackendBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewParamsDataBackendBackendTypeBigquery, ProjectInferencePipelineNewParamsDataBackendBackendTypeDefault, ProjectInferencePipelineNewParamsDataBackendBackendTypeSnowflake, ProjectInferencePipelineNewParamsDataBackendBackendTypeDatabricksDtl, ProjectInferencePipelineNewParamsDataBackendBackendTypeRedshift, ProjectInferencePipelineNewParamsDataBackendBackendTypePostgres:
		return true
	}
	return false
}

type ProjectInferencePipelineNewParamsDataBackendPartitionType string

const (
	ProjectInferencePipelineNewParamsDataBackendPartitionTypeDay   ProjectInferencePipelineNewParamsDataBackendPartitionType = "DAY"
	ProjectInferencePipelineNewParamsDataBackendPartitionTypeMonth ProjectInferencePipelineNewParamsDataBackendPartitionType = "MONTH"
	ProjectInferencePipelineNewParamsDataBackendPartitionTypeYear  ProjectInferencePipelineNewParamsDataBackendPartitionType = "YEAR"
)

func (r ProjectInferencePipelineNewParamsDataBackendPartitionType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewParamsDataBackendPartitionTypeDay, ProjectInferencePipelineNewParamsDataBackendPartitionTypeMonth, ProjectInferencePipelineNewParamsDataBackendPartitionTypeYear:
		return true
	}
	return false
}

type ProjectInferencePipelineNewParamsProject struct {
	// The project name.
	Name param.Field[string] `json:"name,required"`
	// The task type of the project.
	TaskType param.Field[ProjectInferencePipelineNewParamsProjectTaskType] `json:"taskType,required"`
	// The project description.
	Description param.Field[string] `json:"description"`
}

func (r ProjectInferencePipelineNewParamsProject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Links to the project.
type ProjectInferencePipelineNewParamsProjectLinks struct {
	App param.Field[string] `json:"app,required"`
}

func (r ProjectInferencePipelineNewParamsProjectLinks) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The source of the project.
type ProjectInferencePipelineNewParamsProjectSource string

const (
	ProjectInferencePipelineNewParamsProjectSourceWeb  ProjectInferencePipelineNewParamsProjectSource = "web"
	ProjectInferencePipelineNewParamsProjectSourceAPI  ProjectInferencePipelineNewParamsProjectSource = "api"
	ProjectInferencePipelineNewParamsProjectSourceNull ProjectInferencePipelineNewParamsProjectSource = "null"
)

func (r ProjectInferencePipelineNewParamsProjectSource) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewParamsProjectSourceWeb, ProjectInferencePipelineNewParamsProjectSourceAPI, ProjectInferencePipelineNewParamsProjectSourceNull:
		return true
	}
	return false
}

// The task type of the project.
type ProjectInferencePipelineNewParamsProjectTaskType string

const (
	ProjectInferencePipelineNewParamsProjectTaskTypeLlmBase               ProjectInferencePipelineNewParamsProjectTaskType = "llm-base"
	ProjectInferencePipelineNewParamsProjectTaskTypeTabularClassification ProjectInferencePipelineNewParamsProjectTaskType = "tabular-classification"
	ProjectInferencePipelineNewParamsProjectTaskTypeTabularRegression     ProjectInferencePipelineNewParamsProjectTaskType = "tabular-regression"
	ProjectInferencePipelineNewParamsProjectTaskTypeTextClassification    ProjectInferencePipelineNewParamsProjectTaskType = "text-classification"
)

func (r ProjectInferencePipelineNewParamsProjectTaskType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewParamsProjectTaskTypeLlmBase, ProjectInferencePipelineNewParamsProjectTaskTypeTabularClassification, ProjectInferencePipelineNewParamsProjectTaskTypeTabularRegression, ProjectInferencePipelineNewParamsProjectTaskTypeTextClassification:
		return true
	}
	return false
}

type ProjectInferencePipelineNewParamsProjectGitRepo struct {
	GitAccountID param.Field[string] `json:"gitAccountId,required" format:"uuid"`
	GitID        param.Field[int64]  `json:"gitId,required"`
	Branch       param.Field[string] `json:"branch"`
	RootDir      param.Field[string] `json:"rootDir"`
}

func (r ProjectInferencePipelineNewParamsProjectGitRepo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectInferencePipelineNewParamsWorkspace struct {
	// The workspace name.
	Name param.Field[string] `json:"name,required"`
	// The workspace slug.
	Slug param.Field[string] `json:"slug,required"`
	// The workspace invite code.
	InviteCode param.Field[string] `json:"inviteCode"`
	// Whether the workspace only allows SAML authentication.
	SAMLOnlyAccess  param.Field[bool]     `json:"samlOnlyAccess"`
	WildcardDomains param.Field[[]string] `json:"wildcardDomains"`
}

func (r ProjectInferencePipelineNewParamsWorkspace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectInferencePipelineNewParamsWorkspaceStatus string

const (
	ProjectInferencePipelineNewParamsWorkspaceStatusActive            ProjectInferencePipelineNewParamsWorkspaceStatus = "active"
	ProjectInferencePipelineNewParamsWorkspaceStatusPastDue           ProjectInferencePipelineNewParamsWorkspaceStatus = "past_due"
	ProjectInferencePipelineNewParamsWorkspaceStatusUnpaid            ProjectInferencePipelineNewParamsWorkspaceStatus = "unpaid"
	ProjectInferencePipelineNewParamsWorkspaceStatusCanceled          ProjectInferencePipelineNewParamsWorkspaceStatus = "canceled"
	ProjectInferencePipelineNewParamsWorkspaceStatusIncomplete        ProjectInferencePipelineNewParamsWorkspaceStatus = "incomplete"
	ProjectInferencePipelineNewParamsWorkspaceStatusIncompleteExpired ProjectInferencePipelineNewParamsWorkspaceStatus = "incomplete_expired"
	ProjectInferencePipelineNewParamsWorkspaceStatusTrialing          ProjectInferencePipelineNewParamsWorkspaceStatus = "trialing"
	ProjectInferencePipelineNewParamsWorkspaceStatusPaused            ProjectInferencePipelineNewParamsWorkspaceStatus = "paused"
)

func (r ProjectInferencePipelineNewParamsWorkspaceStatus) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewParamsWorkspaceStatusActive, ProjectInferencePipelineNewParamsWorkspaceStatusPastDue, ProjectInferencePipelineNewParamsWorkspaceStatusUnpaid, ProjectInferencePipelineNewParamsWorkspaceStatusCanceled, ProjectInferencePipelineNewParamsWorkspaceStatusIncomplete, ProjectInferencePipelineNewParamsWorkspaceStatusIncompleteExpired, ProjectInferencePipelineNewParamsWorkspaceStatusTrialing, ProjectInferencePipelineNewParamsWorkspaceStatusPaused:
		return true
	}
	return false
}

type ProjectInferencePipelineNewParamsWorkspaceMonthlyUsage struct {
	ExecutionTimeMs param.Field[int64]     `json:"executionTimeMs"`
	MonthYear       param.Field[time.Time] `json:"monthYear" format:"date"`
	PredictionCount param.Field[int64]     `json:"predictionCount"`
}

func (r ProjectInferencePipelineNewParamsWorkspaceMonthlyUsage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectInferencePipelineListParams struct {
	// Filter list of items by name.
	Name param.Field[string] `query:"name"`
	// The page to return in a paginated query.
	Page param.Field[int64] `query:"page"`
	// Maximum number of items to return per page.
	PerPage param.Field[int64] `query:"perPage"`
}

// URLQuery serializes [ProjectInferencePipelineListParams]'s query parameters as
// `url.Values`.
func (r ProjectInferencePipelineListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
