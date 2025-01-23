// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/openlayer-ai/openlayer-go/internal/apijson"
	"github.com/openlayer-ai/openlayer-go/internal/apiquery"
	"github.com/openlayer-ai/openlayer-go/internal/param"
	"github.com/openlayer-ai/openlayer-go/internal/requestconfig"
	"github.com/openlayer-ai/openlayer-go/option"
)

// InferencePipelineService contains methods and other services that help with
// interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInferencePipelineService] method instead.
type InferencePipelineService struct {
	Options     []option.RequestOption
	Data        *InferencePipelineDataService
	Rows        *InferencePipelineRowService
	TestResults *InferencePipelineTestResultService
}

// NewInferencePipelineService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInferencePipelineService(opts ...option.RequestOption) (r *InferencePipelineService) {
	r = &InferencePipelineService{}
	r.Options = opts
	r.Data = NewInferencePipelineDataService(opts...)
	r.Rows = NewInferencePipelineRowService(opts...)
	r.TestResults = NewInferencePipelineTestResultService(opts...)
	return
}

// Retrieve inference pipeline.
func (r *InferencePipelineService) Get(ctx context.Context, inferencePipelineID string, query InferencePipelineGetParams, opts ...option.RequestOption) (res *InferencePipelineGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if inferencePipelineID == "" {
		err = errors.New("missing required inferencePipelineId parameter")
		return
	}
	path := fmt.Sprintf("inference-pipelines/%s", inferencePipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update inference pipeline.
func (r *InferencePipelineService) Update(ctx context.Context, inferencePipelineID string, body InferencePipelineUpdateParams, opts ...option.RequestOption) (res *InferencePipelineUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if inferencePipelineID == "" {
		err = errors.New("missing required inferencePipelineId parameter")
		return
	}
	path := fmt.Sprintf("inference-pipelines/%s", inferencePipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete inference pipeline.
func (r *InferencePipelineService) Delete(ctx context.Context, inferencePipelineID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if inferencePipelineID == "" {
		err = errors.New("missing required inferencePipelineId parameter")
		return
	}
	path := fmt.Sprintf("inference-pipelines/%s", inferencePipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type InferencePipelineGetResponse struct {
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
	FailingGoalCount int64                             `json:"failingGoalCount,required"`
	Links            InferencePipelineGetResponseLinks `json:"links,required"`
	// The inference pipeline name.
	Name string `json:"name,required"`
	// The number of tests passing.
	PassingGoalCount int64 `json:"passingGoalCount,required"`
	// The project id.
	ProjectID string `json:"projectId,required" format:"uuid"`
	// The status of test evaluation for the inference pipeline.
	Status InferencePipelineGetResponseStatus `json:"status,required"`
	// The status message of test evaluation for the inference pipeline.
	StatusMessage string `json:"statusMessage,required,nullable"`
	// The total number of tests.
	TotalGoalCount int64                                 `json:"totalGoalCount,required"`
	Project        InferencePipelineGetResponseProject   `json:"project,nullable"`
	Workspace      InferencePipelineGetResponseWorkspace `json:"workspace,nullable"`
	// The workspace id.
	WorkspaceID string                           `json:"workspaceId" format:"uuid"`
	JSON        inferencePipelineGetResponseJSON `json:"-"`
}

// inferencePipelineGetResponseJSON contains the JSON metadata for the struct
// [InferencePipelineGetResponse]
type inferencePipelineGetResponseJSON struct {
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
	Project                apijson.Field
	Workspace              apijson.Field
	WorkspaceID            apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *InferencePipelineGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineGetResponseJSON) RawJSON() string {
	return r.raw
}

type InferencePipelineGetResponseLinks struct {
	App  string                                `json:"app,required"`
	JSON inferencePipelineGetResponseLinksJSON `json:"-"`
}

// inferencePipelineGetResponseLinksJSON contains the JSON metadata for the struct
// [InferencePipelineGetResponseLinks]
type inferencePipelineGetResponseLinksJSON struct {
	App         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InferencePipelineGetResponseLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineGetResponseLinksJSON) RawJSON() string {
	return r.raw
}

// The status of test evaluation for the inference pipeline.
type InferencePipelineGetResponseStatus string

const (
	InferencePipelineGetResponseStatusQueued    InferencePipelineGetResponseStatus = "queued"
	InferencePipelineGetResponseStatusRunning   InferencePipelineGetResponseStatus = "running"
	InferencePipelineGetResponseStatusPaused    InferencePipelineGetResponseStatus = "paused"
	InferencePipelineGetResponseStatusFailed    InferencePipelineGetResponseStatus = "failed"
	InferencePipelineGetResponseStatusCompleted InferencePipelineGetResponseStatus = "completed"
	InferencePipelineGetResponseStatusUnknown   InferencePipelineGetResponseStatus = "unknown"
)

func (r InferencePipelineGetResponseStatus) IsKnown() bool {
	switch r {
	case InferencePipelineGetResponseStatusQueued, InferencePipelineGetResponseStatusRunning, InferencePipelineGetResponseStatusPaused, InferencePipelineGetResponseStatusFailed, InferencePipelineGetResponseStatusCompleted, InferencePipelineGetResponseStatusUnknown:
		return true
	}
	return false
}

type InferencePipelineGetResponseProject struct {
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
	Links InferencePipelineGetResponseProjectLinks `json:"links,required"`
	// The number of tests in the monitoring mode of the project.
	MonitoringGoalCount int64 `json:"monitoringGoalCount,required"`
	// The project name.
	Name string `json:"name,required"`
	// The source of the project.
	Source InferencePipelineGetResponseProjectSource `json:"source,required,nullable"`
	// The task type of the project.
	TaskType InferencePipelineGetResponseProjectTaskType `json:"taskType,required"`
	// The number of versions (commits) in the project.
	VersionCount int64 `json:"versionCount,required"`
	// The workspace id.
	WorkspaceID string `json:"workspaceId,required,nullable" format:"uuid"`
	// The project description.
	Description string                                     `json:"description,nullable"`
	GitRepo     InferencePipelineGetResponseProjectGitRepo `json:"gitRepo,nullable"`
	JSON        inferencePipelineGetResponseProjectJSON    `json:"-"`
}

// inferencePipelineGetResponseProjectJSON contains the JSON metadata for the
// struct [InferencePipelineGetResponseProject]
type inferencePipelineGetResponseProjectJSON struct {
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

func (r *InferencePipelineGetResponseProject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineGetResponseProjectJSON) RawJSON() string {
	return r.raw
}

// Links to the project.
type InferencePipelineGetResponseProjectLinks struct {
	App  string                                       `json:"app,required"`
	JSON inferencePipelineGetResponseProjectLinksJSON `json:"-"`
}

// inferencePipelineGetResponseProjectLinksJSON contains the JSON metadata for the
// struct [InferencePipelineGetResponseProjectLinks]
type inferencePipelineGetResponseProjectLinksJSON struct {
	App         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InferencePipelineGetResponseProjectLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineGetResponseProjectLinksJSON) RawJSON() string {
	return r.raw
}

// The source of the project.
type InferencePipelineGetResponseProjectSource string

const (
	InferencePipelineGetResponseProjectSourceWeb  InferencePipelineGetResponseProjectSource = "web"
	InferencePipelineGetResponseProjectSourceAPI  InferencePipelineGetResponseProjectSource = "api"
	InferencePipelineGetResponseProjectSourceNull InferencePipelineGetResponseProjectSource = "null"
)

func (r InferencePipelineGetResponseProjectSource) IsKnown() bool {
	switch r {
	case InferencePipelineGetResponseProjectSourceWeb, InferencePipelineGetResponseProjectSourceAPI, InferencePipelineGetResponseProjectSourceNull:
		return true
	}
	return false
}

// The task type of the project.
type InferencePipelineGetResponseProjectTaskType string

const (
	InferencePipelineGetResponseProjectTaskTypeLlmBase               InferencePipelineGetResponseProjectTaskType = "llm-base"
	InferencePipelineGetResponseProjectTaskTypeTabularClassification InferencePipelineGetResponseProjectTaskType = "tabular-classification"
	InferencePipelineGetResponseProjectTaskTypeTabularRegression     InferencePipelineGetResponseProjectTaskType = "tabular-regression"
	InferencePipelineGetResponseProjectTaskTypeTextClassification    InferencePipelineGetResponseProjectTaskType = "text-classification"
)

func (r InferencePipelineGetResponseProjectTaskType) IsKnown() bool {
	switch r {
	case InferencePipelineGetResponseProjectTaskTypeLlmBase, InferencePipelineGetResponseProjectTaskTypeTabularClassification, InferencePipelineGetResponseProjectTaskTypeTabularRegression, InferencePipelineGetResponseProjectTaskTypeTextClassification:
		return true
	}
	return false
}

type InferencePipelineGetResponseProjectGitRepo struct {
	ID            string                                         `json:"id,required" format:"uuid"`
	DateConnected time.Time                                      `json:"dateConnected,required" format:"date-time"`
	DateUpdated   time.Time                                      `json:"dateUpdated,required" format:"date-time"`
	GitAccountID  string                                         `json:"gitAccountId,required" format:"uuid"`
	GitID         int64                                          `json:"gitId,required"`
	Name          string                                         `json:"name,required"`
	Private       bool                                           `json:"private,required"`
	ProjectID     string                                         `json:"projectId,required" format:"uuid"`
	Slug          string                                         `json:"slug,required"`
	URL           string                                         `json:"url,required" format:"url"`
	Branch        string                                         `json:"branch"`
	RootDir       string                                         `json:"rootDir"`
	JSON          inferencePipelineGetResponseProjectGitRepoJSON `json:"-"`
}

// inferencePipelineGetResponseProjectGitRepoJSON contains the JSON metadata for
// the struct [InferencePipelineGetResponseProjectGitRepo]
type inferencePipelineGetResponseProjectGitRepoJSON struct {
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

func (r *InferencePipelineGetResponseProjectGitRepo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineGetResponseProjectGitRepoJSON) RawJSON() string {
	return r.raw
}

type InferencePipelineGetResponseWorkspace struct {
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
	Slug         string                                              `json:"slug,required"`
	Status       InferencePipelineGetResponseWorkspaceStatus         `json:"status,required"`
	MonthlyUsage []InferencePipelineGetResponseWorkspaceMonthlyUsage `json:"monthlyUsage"`
	// Whether the workspace only allows SAML authentication.
	SAMLOnlyAccess  bool                                      `json:"samlOnlyAccess"`
	WildcardDomains []string                                  `json:"wildcardDomains"`
	JSON            inferencePipelineGetResponseWorkspaceJSON `json:"-"`
}

// inferencePipelineGetResponseWorkspaceJSON contains the JSON metadata for the
// struct [InferencePipelineGetResponseWorkspace]
type inferencePipelineGetResponseWorkspaceJSON struct {
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

func (r *InferencePipelineGetResponseWorkspace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineGetResponseWorkspaceJSON) RawJSON() string {
	return r.raw
}

type InferencePipelineGetResponseWorkspaceStatus string

const (
	InferencePipelineGetResponseWorkspaceStatusActive            InferencePipelineGetResponseWorkspaceStatus = "active"
	InferencePipelineGetResponseWorkspaceStatusPastDue           InferencePipelineGetResponseWorkspaceStatus = "past_due"
	InferencePipelineGetResponseWorkspaceStatusUnpaid            InferencePipelineGetResponseWorkspaceStatus = "unpaid"
	InferencePipelineGetResponseWorkspaceStatusCanceled          InferencePipelineGetResponseWorkspaceStatus = "canceled"
	InferencePipelineGetResponseWorkspaceStatusIncomplete        InferencePipelineGetResponseWorkspaceStatus = "incomplete"
	InferencePipelineGetResponseWorkspaceStatusIncompleteExpired InferencePipelineGetResponseWorkspaceStatus = "incomplete_expired"
	InferencePipelineGetResponseWorkspaceStatusTrialing          InferencePipelineGetResponseWorkspaceStatus = "trialing"
	InferencePipelineGetResponseWorkspaceStatusPaused            InferencePipelineGetResponseWorkspaceStatus = "paused"
)

func (r InferencePipelineGetResponseWorkspaceStatus) IsKnown() bool {
	switch r {
	case InferencePipelineGetResponseWorkspaceStatusActive, InferencePipelineGetResponseWorkspaceStatusPastDue, InferencePipelineGetResponseWorkspaceStatusUnpaid, InferencePipelineGetResponseWorkspaceStatusCanceled, InferencePipelineGetResponseWorkspaceStatusIncomplete, InferencePipelineGetResponseWorkspaceStatusIncompleteExpired, InferencePipelineGetResponseWorkspaceStatusTrialing, InferencePipelineGetResponseWorkspaceStatusPaused:
		return true
	}
	return false
}

type InferencePipelineGetResponseWorkspaceMonthlyUsage struct {
	ExecutionTimeMs int64                                                 `json:"executionTimeMs,nullable"`
	MonthYear       time.Time                                             `json:"monthYear" format:"date"`
	PredictionCount int64                                                 `json:"predictionCount"`
	JSON            inferencePipelineGetResponseWorkspaceMonthlyUsageJSON `json:"-"`
}

// inferencePipelineGetResponseWorkspaceMonthlyUsageJSON contains the JSON metadata
// for the struct [InferencePipelineGetResponseWorkspaceMonthlyUsage]
type inferencePipelineGetResponseWorkspaceMonthlyUsageJSON struct {
	ExecutionTimeMs apijson.Field
	MonthYear       apijson.Field
	PredictionCount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InferencePipelineGetResponseWorkspaceMonthlyUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineGetResponseWorkspaceMonthlyUsageJSON) RawJSON() string {
	return r.raw
}

type InferencePipelineUpdateResponse struct {
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
	FailingGoalCount int64                                `json:"failingGoalCount,required"`
	Links            InferencePipelineUpdateResponseLinks `json:"links,required"`
	// The inference pipeline name.
	Name string `json:"name,required"`
	// The number of tests passing.
	PassingGoalCount int64 `json:"passingGoalCount,required"`
	// The project id.
	ProjectID string `json:"projectId,required" format:"uuid"`
	// The status of test evaluation for the inference pipeline.
	Status InferencePipelineUpdateResponseStatus `json:"status,required"`
	// The status message of test evaluation for the inference pipeline.
	StatusMessage string `json:"statusMessage,required,nullable"`
	// The total number of tests.
	TotalGoalCount int64                                    `json:"totalGoalCount,required"`
	Project        InferencePipelineUpdateResponseProject   `json:"project,nullable"`
	Workspace      InferencePipelineUpdateResponseWorkspace `json:"workspace,nullable"`
	// The workspace id.
	WorkspaceID string                              `json:"workspaceId" format:"uuid"`
	JSON        inferencePipelineUpdateResponseJSON `json:"-"`
}

// inferencePipelineUpdateResponseJSON contains the JSON metadata for the struct
// [InferencePipelineUpdateResponse]
type inferencePipelineUpdateResponseJSON struct {
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
	Project                apijson.Field
	Workspace              apijson.Field
	WorkspaceID            apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *InferencePipelineUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type InferencePipelineUpdateResponseLinks struct {
	App  string                                   `json:"app,required"`
	JSON inferencePipelineUpdateResponseLinksJSON `json:"-"`
}

// inferencePipelineUpdateResponseLinksJSON contains the JSON metadata for the
// struct [InferencePipelineUpdateResponseLinks]
type inferencePipelineUpdateResponseLinksJSON struct {
	App         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InferencePipelineUpdateResponseLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineUpdateResponseLinksJSON) RawJSON() string {
	return r.raw
}

// The status of test evaluation for the inference pipeline.
type InferencePipelineUpdateResponseStatus string

const (
	InferencePipelineUpdateResponseStatusQueued    InferencePipelineUpdateResponseStatus = "queued"
	InferencePipelineUpdateResponseStatusRunning   InferencePipelineUpdateResponseStatus = "running"
	InferencePipelineUpdateResponseStatusPaused    InferencePipelineUpdateResponseStatus = "paused"
	InferencePipelineUpdateResponseStatusFailed    InferencePipelineUpdateResponseStatus = "failed"
	InferencePipelineUpdateResponseStatusCompleted InferencePipelineUpdateResponseStatus = "completed"
	InferencePipelineUpdateResponseStatusUnknown   InferencePipelineUpdateResponseStatus = "unknown"
)

func (r InferencePipelineUpdateResponseStatus) IsKnown() bool {
	switch r {
	case InferencePipelineUpdateResponseStatusQueued, InferencePipelineUpdateResponseStatusRunning, InferencePipelineUpdateResponseStatusPaused, InferencePipelineUpdateResponseStatusFailed, InferencePipelineUpdateResponseStatusCompleted, InferencePipelineUpdateResponseStatusUnknown:
		return true
	}
	return false
}

type InferencePipelineUpdateResponseProject struct {
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
	Links InferencePipelineUpdateResponseProjectLinks `json:"links,required"`
	// The number of tests in the monitoring mode of the project.
	MonitoringGoalCount int64 `json:"monitoringGoalCount,required"`
	// The project name.
	Name string `json:"name,required"`
	// The source of the project.
	Source InferencePipelineUpdateResponseProjectSource `json:"source,required,nullable"`
	// The task type of the project.
	TaskType InferencePipelineUpdateResponseProjectTaskType `json:"taskType,required"`
	// The number of versions (commits) in the project.
	VersionCount int64 `json:"versionCount,required"`
	// The workspace id.
	WorkspaceID string `json:"workspaceId,required,nullable" format:"uuid"`
	// The project description.
	Description string                                        `json:"description,nullable"`
	GitRepo     InferencePipelineUpdateResponseProjectGitRepo `json:"gitRepo,nullable"`
	JSON        inferencePipelineUpdateResponseProjectJSON    `json:"-"`
}

// inferencePipelineUpdateResponseProjectJSON contains the JSON metadata for the
// struct [InferencePipelineUpdateResponseProject]
type inferencePipelineUpdateResponseProjectJSON struct {
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

func (r *InferencePipelineUpdateResponseProject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineUpdateResponseProjectJSON) RawJSON() string {
	return r.raw
}

// Links to the project.
type InferencePipelineUpdateResponseProjectLinks struct {
	App  string                                          `json:"app,required"`
	JSON inferencePipelineUpdateResponseProjectLinksJSON `json:"-"`
}

// inferencePipelineUpdateResponseProjectLinksJSON contains the JSON metadata for
// the struct [InferencePipelineUpdateResponseProjectLinks]
type inferencePipelineUpdateResponseProjectLinksJSON struct {
	App         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InferencePipelineUpdateResponseProjectLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineUpdateResponseProjectLinksJSON) RawJSON() string {
	return r.raw
}

// The source of the project.
type InferencePipelineUpdateResponseProjectSource string

const (
	InferencePipelineUpdateResponseProjectSourceWeb  InferencePipelineUpdateResponseProjectSource = "web"
	InferencePipelineUpdateResponseProjectSourceAPI  InferencePipelineUpdateResponseProjectSource = "api"
	InferencePipelineUpdateResponseProjectSourceNull InferencePipelineUpdateResponseProjectSource = "null"
)

func (r InferencePipelineUpdateResponseProjectSource) IsKnown() bool {
	switch r {
	case InferencePipelineUpdateResponseProjectSourceWeb, InferencePipelineUpdateResponseProjectSourceAPI, InferencePipelineUpdateResponseProjectSourceNull:
		return true
	}
	return false
}

// The task type of the project.
type InferencePipelineUpdateResponseProjectTaskType string

const (
	InferencePipelineUpdateResponseProjectTaskTypeLlmBase               InferencePipelineUpdateResponseProjectTaskType = "llm-base"
	InferencePipelineUpdateResponseProjectTaskTypeTabularClassification InferencePipelineUpdateResponseProjectTaskType = "tabular-classification"
	InferencePipelineUpdateResponseProjectTaskTypeTabularRegression     InferencePipelineUpdateResponseProjectTaskType = "tabular-regression"
	InferencePipelineUpdateResponseProjectTaskTypeTextClassification    InferencePipelineUpdateResponseProjectTaskType = "text-classification"
)

func (r InferencePipelineUpdateResponseProjectTaskType) IsKnown() bool {
	switch r {
	case InferencePipelineUpdateResponseProjectTaskTypeLlmBase, InferencePipelineUpdateResponseProjectTaskTypeTabularClassification, InferencePipelineUpdateResponseProjectTaskTypeTabularRegression, InferencePipelineUpdateResponseProjectTaskTypeTextClassification:
		return true
	}
	return false
}

type InferencePipelineUpdateResponseProjectGitRepo struct {
	ID            string                                            `json:"id,required" format:"uuid"`
	DateConnected time.Time                                         `json:"dateConnected,required" format:"date-time"`
	DateUpdated   time.Time                                         `json:"dateUpdated,required" format:"date-time"`
	GitAccountID  string                                            `json:"gitAccountId,required" format:"uuid"`
	GitID         int64                                             `json:"gitId,required"`
	Name          string                                            `json:"name,required"`
	Private       bool                                              `json:"private,required"`
	ProjectID     string                                            `json:"projectId,required" format:"uuid"`
	Slug          string                                            `json:"slug,required"`
	URL           string                                            `json:"url,required" format:"url"`
	Branch        string                                            `json:"branch"`
	RootDir       string                                            `json:"rootDir"`
	JSON          inferencePipelineUpdateResponseProjectGitRepoJSON `json:"-"`
}

// inferencePipelineUpdateResponseProjectGitRepoJSON contains the JSON metadata for
// the struct [InferencePipelineUpdateResponseProjectGitRepo]
type inferencePipelineUpdateResponseProjectGitRepoJSON struct {
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

func (r *InferencePipelineUpdateResponseProjectGitRepo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineUpdateResponseProjectGitRepoJSON) RawJSON() string {
	return r.raw
}

type InferencePipelineUpdateResponseWorkspace struct {
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
	Slug         string                                                 `json:"slug,required"`
	Status       InferencePipelineUpdateResponseWorkspaceStatus         `json:"status,required"`
	MonthlyUsage []InferencePipelineUpdateResponseWorkspaceMonthlyUsage `json:"monthlyUsage"`
	// Whether the workspace only allows SAML authentication.
	SAMLOnlyAccess  bool                                         `json:"samlOnlyAccess"`
	WildcardDomains []string                                     `json:"wildcardDomains"`
	JSON            inferencePipelineUpdateResponseWorkspaceJSON `json:"-"`
}

// inferencePipelineUpdateResponseWorkspaceJSON contains the JSON metadata for the
// struct [InferencePipelineUpdateResponseWorkspace]
type inferencePipelineUpdateResponseWorkspaceJSON struct {
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

func (r *InferencePipelineUpdateResponseWorkspace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineUpdateResponseWorkspaceJSON) RawJSON() string {
	return r.raw
}

type InferencePipelineUpdateResponseWorkspaceStatus string

const (
	InferencePipelineUpdateResponseWorkspaceStatusActive            InferencePipelineUpdateResponseWorkspaceStatus = "active"
	InferencePipelineUpdateResponseWorkspaceStatusPastDue           InferencePipelineUpdateResponseWorkspaceStatus = "past_due"
	InferencePipelineUpdateResponseWorkspaceStatusUnpaid            InferencePipelineUpdateResponseWorkspaceStatus = "unpaid"
	InferencePipelineUpdateResponseWorkspaceStatusCanceled          InferencePipelineUpdateResponseWorkspaceStatus = "canceled"
	InferencePipelineUpdateResponseWorkspaceStatusIncomplete        InferencePipelineUpdateResponseWorkspaceStatus = "incomplete"
	InferencePipelineUpdateResponseWorkspaceStatusIncompleteExpired InferencePipelineUpdateResponseWorkspaceStatus = "incomplete_expired"
	InferencePipelineUpdateResponseWorkspaceStatusTrialing          InferencePipelineUpdateResponseWorkspaceStatus = "trialing"
	InferencePipelineUpdateResponseWorkspaceStatusPaused            InferencePipelineUpdateResponseWorkspaceStatus = "paused"
)

func (r InferencePipelineUpdateResponseWorkspaceStatus) IsKnown() bool {
	switch r {
	case InferencePipelineUpdateResponseWorkspaceStatusActive, InferencePipelineUpdateResponseWorkspaceStatusPastDue, InferencePipelineUpdateResponseWorkspaceStatusUnpaid, InferencePipelineUpdateResponseWorkspaceStatusCanceled, InferencePipelineUpdateResponseWorkspaceStatusIncomplete, InferencePipelineUpdateResponseWorkspaceStatusIncompleteExpired, InferencePipelineUpdateResponseWorkspaceStatusTrialing, InferencePipelineUpdateResponseWorkspaceStatusPaused:
		return true
	}
	return false
}

type InferencePipelineUpdateResponseWorkspaceMonthlyUsage struct {
	ExecutionTimeMs int64                                                    `json:"executionTimeMs,nullable"`
	MonthYear       time.Time                                                `json:"monthYear" format:"date"`
	PredictionCount int64                                                    `json:"predictionCount"`
	JSON            inferencePipelineUpdateResponseWorkspaceMonthlyUsageJSON `json:"-"`
}

// inferencePipelineUpdateResponseWorkspaceMonthlyUsageJSON contains the JSON
// metadata for the struct [InferencePipelineUpdateResponseWorkspaceMonthlyUsage]
type inferencePipelineUpdateResponseWorkspaceMonthlyUsageJSON struct {
	ExecutionTimeMs apijson.Field
	MonthYear       apijson.Field
	PredictionCount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InferencePipelineUpdateResponseWorkspaceMonthlyUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineUpdateResponseWorkspaceMonthlyUsageJSON) RawJSON() string {
	return r.raw
}

type InferencePipelineGetParams struct {
	// Expand specific nested objects.
	Expand param.Field[[]InferencePipelineGetParamsExpand] `query:"expand"`
}

// URLQuery serializes [InferencePipelineGetParams]'s query parameters as
// `url.Values`.
func (r InferencePipelineGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InferencePipelineGetParamsExpand string

const (
	InferencePipelineGetParamsExpandProject   InferencePipelineGetParamsExpand = "project"
	InferencePipelineGetParamsExpandWorkspace InferencePipelineGetParamsExpand = "workspace"
)

func (r InferencePipelineGetParamsExpand) IsKnown() bool {
	switch r {
	case InferencePipelineGetParamsExpandProject, InferencePipelineGetParamsExpandWorkspace:
		return true
	}
	return false
}

type InferencePipelineUpdateParams struct {
	// The inference pipeline description.
	Description param.Field[string] `json:"description"`
	// The inference pipeline name.
	Name param.Field[string] `json:"name"`
	// The storage uri of your reference dataset. We recommend using the Python SDK or
	// the UI to handle your reference dataset updates.
	ReferenceDatasetUri param.Field[string] `json:"referenceDatasetUri"`
}

func (r InferencePipelineUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
