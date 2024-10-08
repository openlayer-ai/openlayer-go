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
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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
	TotalGoalCount int64                                   `json:"totalGoalCount,required"`
	JSON           projectInferencePipelineNewResponseJSON `json:"-"`
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
	TotalGoalCount int64                                        `json:"totalGoalCount,required"`
	JSON           projectInferencePipelineListResponseItemJSON `json:"-"`
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

type ProjectInferencePipelineNewParams struct {
	// The inference pipeline description.
	Description param.Field[string] `json:"description,required"`
	// The inference pipeline name.
	Name param.Field[string] `json:"name,required"`
}

func (r ProjectInferencePipelineNewParams) MarshalJSON() (data []byte, err error) {
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
