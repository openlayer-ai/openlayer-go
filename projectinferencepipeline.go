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
	// [ProjectInferencePipelineNewResponseDataBackendObjectConfig].
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
// [ProjectInferencePipelineNewResponseDataBackendObject],
// [ProjectInferencePipelineNewResponseDataBackendBackendType],
// [ProjectInferencePipelineNewResponseDataBackendObject],
// [ProjectInferencePipelineNewResponseDataBackendObject],
// [ProjectInferencePipelineNewResponseDataBackendObject],
// [ProjectInferencePipelineNewResponseDataBackendObject].
func (r ProjectInferencePipelineNewResponseDataBackend) AsUnion() ProjectInferencePipelineNewResponseDataBackendUnion {
	return r.union
}

// Union satisfied by [ProjectInferencePipelineNewResponseDataBackendObject],
// [ProjectInferencePipelineNewResponseDataBackendBackendType],
// [ProjectInferencePipelineNewResponseDataBackendObject],
// [ProjectInferencePipelineNewResponseDataBackendObject],
// [ProjectInferencePipelineNewResponseDataBackendObject] or
// [ProjectInferencePipelineNewResponseDataBackendObject].
type ProjectInferencePipelineNewResponseDataBackendUnion interface {
	implementsProjectInferencePipelineNewResponseDataBackend()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectInferencePipelineNewResponseDataBackendUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineNewResponseDataBackendObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineNewResponseDataBackendBackendType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineNewResponseDataBackendObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineNewResponseDataBackendObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineNewResponseDataBackendObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineNewResponseDataBackendObject{}),
		},
	)
}

type ProjectInferencePipelineNewResponseDataBackendObject struct {
	BackendType          ProjectInferencePipelineNewResponseDataBackendObjectBackendType   `json:"backendType,required"`
	BigqueryConnectionID string                                                            `json:"bigqueryConnectionId,required,nullable" format:"uuid"`
	DatasetID            string                                                            `json:"datasetId,required"`
	ProjectID            string                                                            `json:"projectId,required"`
	TableID              string                                                            `json:"tableId,required,nullable"`
	PartitionType        ProjectInferencePipelineNewResponseDataBackendObjectPartitionType `json:"partitionType,nullable"`
	JSON                 projectInferencePipelineNewResponseDataBackendObjectJSON          `json:"-"`
}

// projectInferencePipelineNewResponseDataBackendObjectJSON contains the JSON
// metadata for the struct [ProjectInferencePipelineNewResponseDataBackendObject]
type projectInferencePipelineNewResponseDataBackendObjectJSON struct {
	BackendType          apijson.Field
	BigqueryConnectionID apijson.Field
	DatasetID            apijson.Field
	ProjectID            apijson.Field
	TableID              apijson.Field
	PartitionType        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectInferencePipelineNewResponseDataBackendObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineNewResponseDataBackendObjectJSON) RawJSON() string {
	return r.raw
}

func (r ProjectInferencePipelineNewResponseDataBackendObject) implementsProjectInferencePipelineNewResponseDataBackend() {
}

type ProjectInferencePipelineNewResponseDataBackendObjectBackendType string

const (
	ProjectInferencePipelineNewResponseDataBackendObjectBackendTypeBigquery ProjectInferencePipelineNewResponseDataBackendObjectBackendType = "bigquery"
)

func (r ProjectInferencePipelineNewResponseDataBackendObjectBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewResponseDataBackendObjectBackendTypeBigquery:
		return true
	}
	return false
}

type ProjectInferencePipelineNewResponseDataBackendObjectConfig struct {
	// Name of the column with the ground truths.
	GroundTruthColumnName string `json:"groundTruthColumnName,nullable"`
	// Name of the column with human feedback.
	HumanFeedbackColumnName string `json:"humanFeedbackColumnName,nullable"`
	// Name of the column with the latencies.
	LatencyColumnName string `json:"latencyColumnName,nullable"`
	// Name of the column with the timestamps. Timestamps must be in UNIX sec format.
	// If not provided, the upload timestamp is used.
	TimestampColumnName string                                                         `json:"timestampColumnName,nullable"`
	JSON                projectInferencePipelineNewResponseDataBackendObjectConfigJSON `json:"-"`
}

// projectInferencePipelineNewResponseDataBackendObjectConfigJSON contains the JSON
// metadata for the struct
// [ProjectInferencePipelineNewResponseDataBackendObjectConfig]
type projectInferencePipelineNewResponseDataBackendObjectConfigJSON struct {
	GroundTruthColumnName   apijson.Field
	HumanFeedbackColumnName apijson.Field
	LatencyColumnName       apijson.Field
	TimestampColumnName     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ProjectInferencePipelineNewResponseDataBackendObjectConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineNewResponseDataBackendObjectConfigJSON) RawJSON() string {
	return r.raw
}

type ProjectInferencePipelineNewResponseDataBackendObjectPartitionType string

const (
	ProjectInferencePipelineNewResponseDataBackendObjectPartitionTypeDay   ProjectInferencePipelineNewResponseDataBackendObjectPartitionType = "DAY"
	ProjectInferencePipelineNewResponseDataBackendObjectPartitionTypeMonth ProjectInferencePipelineNewResponseDataBackendObjectPartitionType = "MONTH"
	ProjectInferencePipelineNewResponseDataBackendObjectPartitionTypeYear  ProjectInferencePipelineNewResponseDataBackendObjectPartitionType = "YEAR"
)

func (r ProjectInferencePipelineNewResponseDataBackendObjectPartitionType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewResponseDataBackendObjectPartitionTypeDay, ProjectInferencePipelineNewResponseDataBackendObjectPartitionTypeMonth, ProjectInferencePipelineNewResponseDataBackendObjectPartitionTypeYear:
		return true
	}
	return false
}

type ProjectInferencePipelineNewResponseDataBackendBackendType struct {
	BackendType ProjectInferencePipelineNewResponseDataBackendBackendTypeBackendType `json:"backendType,required"`
	JSON        projectInferencePipelineNewResponseDataBackendBackendTypeJSON        `json:"-"`
}

// projectInferencePipelineNewResponseDataBackendBackendTypeJSON contains the JSON
// metadata for the struct
// [ProjectInferencePipelineNewResponseDataBackendBackendType]
type projectInferencePipelineNewResponseDataBackendBackendTypeJSON struct {
	BackendType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectInferencePipelineNewResponseDataBackendBackendType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineNewResponseDataBackendBackendTypeJSON) RawJSON() string {
	return r.raw
}

func (r ProjectInferencePipelineNewResponseDataBackendBackendType) implementsProjectInferencePipelineNewResponseDataBackend() {
}

type ProjectInferencePipelineNewResponseDataBackendBackendTypeBackendType string

const (
	ProjectInferencePipelineNewResponseDataBackendBackendTypeBackendTypeDefault ProjectInferencePipelineNewResponseDataBackendBackendTypeBackendType = "default"
)

func (r ProjectInferencePipelineNewResponseDataBackendBackendTypeBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewResponseDataBackendBackendTypeBackendTypeDefault:
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
	// [ProjectInferencePipelineListResponseItemsDataBackendObjectConfig].
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
// [ProjectInferencePipelineListResponseItemsDataBackendObject],
// [ProjectInferencePipelineListResponseItemsDataBackendBackendType],
// [ProjectInferencePipelineListResponseItemsDataBackendObject],
// [ProjectInferencePipelineListResponseItemsDataBackendObject],
// [ProjectInferencePipelineListResponseItemsDataBackendObject],
// [ProjectInferencePipelineListResponseItemsDataBackendObject].
func (r ProjectInferencePipelineListResponseItemsDataBackend) AsUnion() ProjectInferencePipelineListResponseItemsDataBackendUnion {
	return r.union
}

// Union satisfied by [ProjectInferencePipelineListResponseItemsDataBackendObject],
// [ProjectInferencePipelineListResponseItemsDataBackendBackendType],
// [ProjectInferencePipelineListResponseItemsDataBackendObject],
// [ProjectInferencePipelineListResponseItemsDataBackendObject],
// [ProjectInferencePipelineListResponseItemsDataBackendObject] or
// [ProjectInferencePipelineListResponseItemsDataBackendObject].
type ProjectInferencePipelineListResponseItemsDataBackendUnion interface {
	implementsProjectInferencePipelineListResponseItemsDataBackend()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectInferencePipelineListResponseItemsDataBackendUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineListResponseItemsDataBackendObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineListResponseItemsDataBackendBackendType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineListResponseItemsDataBackendObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineListResponseItemsDataBackendObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineListResponseItemsDataBackendObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectInferencePipelineListResponseItemsDataBackendObject{}),
		},
	)
}

type ProjectInferencePipelineListResponseItemsDataBackendObject struct {
	BackendType          ProjectInferencePipelineListResponseItemsDataBackendObjectBackendType   `json:"backendType,required"`
	BigqueryConnectionID string                                                                  `json:"bigqueryConnectionId,required,nullable" format:"uuid"`
	DatasetID            string                                                                  `json:"datasetId,required"`
	ProjectID            string                                                                  `json:"projectId,required"`
	TableID              string                                                                  `json:"tableId,required,nullable"`
	PartitionType        ProjectInferencePipelineListResponseItemsDataBackendObjectPartitionType `json:"partitionType,nullable"`
	JSON                 projectInferencePipelineListResponseItemsDataBackendObjectJSON          `json:"-"`
}

// projectInferencePipelineListResponseItemsDataBackendObjectJSON contains the JSON
// metadata for the struct
// [ProjectInferencePipelineListResponseItemsDataBackendObject]
type projectInferencePipelineListResponseItemsDataBackendObjectJSON struct {
	BackendType          apijson.Field
	BigqueryConnectionID apijson.Field
	DatasetID            apijson.Field
	ProjectID            apijson.Field
	TableID              apijson.Field
	PartitionType        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectInferencePipelineListResponseItemsDataBackendObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineListResponseItemsDataBackendObjectJSON) RawJSON() string {
	return r.raw
}

func (r ProjectInferencePipelineListResponseItemsDataBackendObject) implementsProjectInferencePipelineListResponseItemsDataBackend() {
}

type ProjectInferencePipelineListResponseItemsDataBackendObjectBackendType string

const (
	ProjectInferencePipelineListResponseItemsDataBackendObjectBackendTypeBigquery ProjectInferencePipelineListResponseItemsDataBackendObjectBackendType = "bigquery"
)

func (r ProjectInferencePipelineListResponseItemsDataBackendObjectBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineListResponseItemsDataBackendObjectBackendTypeBigquery:
		return true
	}
	return false
}

type ProjectInferencePipelineListResponseItemsDataBackendObjectConfig struct {
	// Name of the column with the ground truths.
	GroundTruthColumnName string `json:"groundTruthColumnName,nullable"`
	// Name of the column with human feedback.
	HumanFeedbackColumnName string `json:"humanFeedbackColumnName,nullable"`
	// Name of the column with the latencies.
	LatencyColumnName string `json:"latencyColumnName,nullable"`
	// Name of the column with the timestamps. Timestamps must be in UNIX sec format.
	// If not provided, the upload timestamp is used.
	TimestampColumnName string                                                               `json:"timestampColumnName,nullable"`
	JSON                projectInferencePipelineListResponseItemsDataBackendObjectConfigJSON `json:"-"`
}

// projectInferencePipelineListResponseItemsDataBackendObjectConfigJSON contains
// the JSON metadata for the struct
// [ProjectInferencePipelineListResponseItemsDataBackendObjectConfig]
type projectInferencePipelineListResponseItemsDataBackendObjectConfigJSON struct {
	GroundTruthColumnName   apijson.Field
	HumanFeedbackColumnName apijson.Field
	LatencyColumnName       apijson.Field
	TimestampColumnName     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ProjectInferencePipelineListResponseItemsDataBackendObjectConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineListResponseItemsDataBackendObjectConfigJSON) RawJSON() string {
	return r.raw
}

type ProjectInferencePipelineListResponseItemsDataBackendObjectPartitionType string

const (
	ProjectInferencePipelineListResponseItemsDataBackendObjectPartitionTypeDay   ProjectInferencePipelineListResponseItemsDataBackendObjectPartitionType = "DAY"
	ProjectInferencePipelineListResponseItemsDataBackendObjectPartitionTypeMonth ProjectInferencePipelineListResponseItemsDataBackendObjectPartitionType = "MONTH"
	ProjectInferencePipelineListResponseItemsDataBackendObjectPartitionTypeYear  ProjectInferencePipelineListResponseItemsDataBackendObjectPartitionType = "YEAR"
)

func (r ProjectInferencePipelineListResponseItemsDataBackendObjectPartitionType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineListResponseItemsDataBackendObjectPartitionTypeDay, ProjectInferencePipelineListResponseItemsDataBackendObjectPartitionTypeMonth, ProjectInferencePipelineListResponseItemsDataBackendObjectPartitionTypeYear:
		return true
	}
	return false
}

type ProjectInferencePipelineListResponseItemsDataBackendBackendType struct {
	BackendType ProjectInferencePipelineListResponseItemsDataBackendBackendTypeBackendType `json:"backendType,required"`
	JSON        projectInferencePipelineListResponseItemsDataBackendBackendTypeJSON        `json:"-"`
}

// projectInferencePipelineListResponseItemsDataBackendBackendTypeJSON contains the
// JSON metadata for the struct
// [ProjectInferencePipelineListResponseItemsDataBackendBackendType]
type projectInferencePipelineListResponseItemsDataBackendBackendTypeJSON struct {
	BackendType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectInferencePipelineListResponseItemsDataBackendBackendType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectInferencePipelineListResponseItemsDataBackendBackendTypeJSON) RawJSON() string {
	return r.raw
}

func (r ProjectInferencePipelineListResponseItemsDataBackendBackendType) implementsProjectInferencePipelineListResponseItemsDataBackend() {
}

type ProjectInferencePipelineListResponseItemsDataBackendBackendTypeBackendType string

const (
	ProjectInferencePipelineListResponseItemsDataBackendBackendTypeBackendTypeDefault ProjectInferencePipelineListResponseItemsDataBackendBackendTypeBackendType = "default"
)

func (r ProjectInferencePipelineListResponseItemsDataBackendBackendTypeBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineListResponseItemsDataBackendBackendTypeBackendTypeDefault:
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

// Satisfied by [ProjectInferencePipelineNewParamsDataBackendObject],
// [ProjectInferencePipelineNewParamsDataBackendBackendType],
// [ProjectInferencePipelineNewParamsDataBackendObject],
// [ProjectInferencePipelineNewParamsDataBackendObject],
// [ProjectInferencePipelineNewParamsDataBackendObject],
// [ProjectInferencePipelineNewParamsDataBackendObject],
// [ProjectInferencePipelineNewParamsDataBackend].
type ProjectInferencePipelineNewParamsDataBackendUnion interface {
	implementsProjectInferencePipelineNewParamsDataBackendUnion()
}

type ProjectInferencePipelineNewParamsDataBackendObject struct {
	BackendType          param.Field[ProjectInferencePipelineNewParamsDataBackendObjectBackendType]   `json:"backendType,required"`
	BigqueryConnectionID param.Field[string]                                                          `json:"bigqueryConnectionId,required" format:"uuid"`
	Config               param.Field[ProjectInferencePipelineNewParamsDataBackendObjectConfig]        `json:"config,required"`
	DatasetID            param.Field[string]                                                          `json:"datasetId,required"`
	ProjectID            param.Field[string]                                                          `json:"projectId,required"`
	TableID              param.Field[string]                                                          `json:"tableId,required"`
	PartitionType        param.Field[ProjectInferencePipelineNewParamsDataBackendObjectPartitionType] `json:"partitionType"`
}

func (r ProjectInferencePipelineNewParamsDataBackendObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectInferencePipelineNewParamsDataBackendObject) implementsProjectInferencePipelineNewParamsDataBackendUnion() {
}

type ProjectInferencePipelineNewParamsDataBackendObjectBackendType string

const (
	ProjectInferencePipelineNewParamsDataBackendObjectBackendTypeBigquery ProjectInferencePipelineNewParamsDataBackendObjectBackendType = "bigquery"
)

func (r ProjectInferencePipelineNewParamsDataBackendObjectBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewParamsDataBackendObjectBackendTypeBigquery:
		return true
	}
	return false
}

type ProjectInferencePipelineNewParamsDataBackendObjectConfig struct {
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

func (r ProjectInferencePipelineNewParamsDataBackendObjectConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectInferencePipelineNewParamsDataBackendObjectPartitionType string

const (
	ProjectInferencePipelineNewParamsDataBackendObjectPartitionTypeDay   ProjectInferencePipelineNewParamsDataBackendObjectPartitionType = "DAY"
	ProjectInferencePipelineNewParamsDataBackendObjectPartitionTypeMonth ProjectInferencePipelineNewParamsDataBackendObjectPartitionType = "MONTH"
	ProjectInferencePipelineNewParamsDataBackendObjectPartitionTypeYear  ProjectInferencePipelineNewParamsDataBackendObjectPartitionType = "YEAR"
)

func (r ProjectInferencePipelineNewParamsDataBackendObjectPartitionType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewParamsDataBackendObjectPartitionTypeDay, ProjectInferencePipelineNewParamsDataBackendObjectPartitionTypeMonth, ProjectInferencePipelineNewParamsDataBackendObjectPartitionTypeYear:
		return true
	}
	return false
}

type ProjectInferencePipelineNewParamsDataBackendBackendType struct {
	BackendType param.Field[ProjectInferencePipelineNewParamsDataBackendBackendTypeBackendType] `json:"backendType,required"`
}

func (r ProjectInferencePipelineNewParamsDataBackendBackendType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectInferencePipelineNewParamsDataBackendBackendType) implementsProjectInferencePipelineNewParamsDataBackendUnion() {
}

type ProjectInferencePipelineNewParamsDataBackendBackendTypeBackendType string

const (
	ProjectInferencePipelineNewParamsDataBackendBackendTypeBackendTypeDefault ProjectInferencePipelineNewParamsDataBackendBackendTypeBackendType = "default"
)

func (r ProjectInferencePipelineNewParamsDataBackendBackendTypeBackendType) IsKnown() bool {
	switch r {
	case ProjectInferencePipelineNewParamsDataBackendBackendTypeBackendTypeDefault:
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
