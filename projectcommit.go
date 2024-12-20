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

// ProjectCommitService contains methods and other services that help with
// interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectCommitService] method instead.
type ProjectCommitService struct {
	Options []option.RequestOption
}

// NewProjectCommitService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProjectCommitService(opts ...option.RequestOption) (r *ProjectCommitService) {
	r = &ProjectCommitService{}
	r.Options = opts
	return
}

// Create a new commit (project version) in a project.
func (r *ProjectCommitService) New(ctx context.Context, projectID string, body ProjectCommitNewParams, opts ...option.RequestOption) (res *ProjectCommitNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return
	}
	path := fmt.Sprintf("projects/%s/versions", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List the commits (project versions) in a project.
func (r *ProjectCommitService) List(ctx context.Context, projectID string, query ProjectCommitListParams, opts ...option.RequestOption) (res *ProjectCommitListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return
	}
	path := fmt.Sprintf("projects/%s/versions", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ProjectCommitNewResponse struct {
	// The project version (commit) id.
	ID string `json:"id,required" format:"uuid"`
	// The details of a commit (project version).
	Commit ProjectCommitNewResponseCommit `json:"commit,required"`
	// The commit archive date.
	DateArchived time.Time `json:"dateArchived,required,nullable" format:"date-time"`
	// The project version (commit) creation date.
	DateCreated time.Time `json:"dateCreated,required" format:"date-time"`
	// The number of tests that are failing for the commit.
	FailingGoalCount int64 `json:"failingGoalCount,required"`
	// The model id.
	MlModelID string `json:"mlModelId,required,nullable" format:"uuid"`
	// The number of tests that are passing for the commit.
	PassingGoalCount int64 `json:"passingGoalCount,required"`
	// The project id.
	ProjectID string `json:"projectId,required" format:"uuid"`
	// The commit status. Initially, the commit is `queued`, then, it switches to
	// `running`. Finally, it can be `paused`, `failed`, or `completed`.
	Status ProjectCommitNewResponseStatus `json:"status,required"`
	// The commit status message.
	StatusMessage string `json:"statusMessage,required,nullable"`
	// The total number of tests for the commit.
	TotalGoalCount int64 `json:"totalGoalCount,required"`
	// The training dataset id.
	TrainingDatasetID string `json:"trainingDatasetId,required,nullable" format:"uuid"`
	// The validation dataset id.
	ValidationDatasetID string `json:"validationDatasetId,required,nullable" format:"uuid"`
	// Whether the commit is archived.
	Archived bool `json:"archived,nullable"`
	// The deployment status associated with the commit's model.
	DeploymentStatus string                        `json:"deploymentStatus"`
	Links            ProjectCommitNewResponseLinks `json:"links"`
	JSON             projectCommitNewResponseJSON  `json:"-"`
}

// projectCommitNewResponseJSON contains the JSON metadata for the struct
// [ProjectCommitNewResponse]
type projectCommitNewResponseJSON struct {
	ID                  apijson.Field
	Commit              apijson.Field
	DateArchived        apijson.Field
	DateCreated         apijson.Field
	FailingGoalCount    apijson.Field
	MlModelID           apijson.Field
	PassingGoalCount    apijson.Field
	ProjectID           apijson.Field
	Status              apijson.Field
	StatusMessage       apijson.Field
	TotalGoalCount      apijson.Field
	TrainingDatasetID   apijson.Field
	ValidationDatasetID apijson.Field
	Archived            apijson.Field
	DeploymentStatus    apijson.Field
	Links               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProjectCommitNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectCommitNewResponseJSON) RawJSON() string {
	return r.raw
}

// The details of a commit (project version).
type ProjectCommitNewResponseCommit struct {
	// The commit id.
	ID string `json:"id,required" format:"uuid"`
	// The author id of the commit.
	AuthorID string `json:"authorId,required" format:"uuid"`
	// The size of the commit bundle in bytes.
	FileSize int64 `json:"fileSize,required,nullable"`
	// The commit message.
	Message string `json:"message,required"`
	// The model id.
	MlModelID string `json:"mlModelId,required,nullable" format:"uuid"`
	// The storage URI where the commit bundle is stored.
	StorageUri string `json:"storageUri,required"`
	// The training dataset id.
	TrainingDatasetID string `json:"trainingDatasetId,required,nullable" format:"uuid"`
	// The validation dataset id.
	ValidationDatasetID string `json:"validationDatasetId,required,nullable" format:"uuid"`
	// The commit creation date.
	DateCreated time.Time `json:"dateCreated" format:"date-time"`
	// The ref of the corresponding git commit.
	GitCommitRef string `json:"gitCommitRef"`
	// The SHA of the corresponding git commit.
	GitCommitSha int64 `json:"gitCommitSha"`
	// The URL of the corresponding git commit.
	GitCommitURL string                             `json:"gitCommitUrl"`
	JSON         projectCommitNewResponseCommitJSON `json:"-"`
}

// projectCommitNewResponseCommitJSON contains the JSON metadata for the struct
// [ProjectCommitNewResponseCommit]
type projectCommitNewResponseCommitJSON struct {
	ID                  apijson.Field
	AuthorID            apijson.Field
	FileSize            apijson.Field
	Message             apijson.Field
	MlModelID           apijson.Field
	StorageUri          apijson.Field
	TrainingDatasetID   apijson.Field
	ValidationDatasetID apijson.Field
	DateCreated         apijson.Field
	GitCommitRef        apijson.Field
	GitCommitSha        apijson.Field
	GitCommitURL        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProjectCommitNewResponseCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectCommitNewResponseCommitJSON) RawJSON() string {
	return r.raw
}

// The commit status. Initially, the commit is `queued`, then, it switches to
// `running`. Finally, it can be `paused`, `failed`, or `completed`.
type ProjectCommitNewResponseStatus string

const (
	ProjectCommitNewResponseStatusQueued    ProjectCommitNewResponseStatus = "queued"
	ProjectCommitNewResponseStatusRunning   ProjectCommitNewResponseStatus = "running"
	ProjectCommitNewResponseStatusPaused    ProjectCommitNewResponseStatus = "paused"
	ProjectCommitNewResponseStatusFailed    ProjectCommitNewResponseStatus = "failed"
	ProjectCommitNewResponseStatusCompleted ProjectCommitNewResponseStatus = "completed"
	ProjectCommitNewResponseStatusUnknown   ProjectCommitNewResponseStatus = "unknown"
)

func (r ProjectCommitNewResponseStatus) IsKnown() bool {
	switch r {
	case ProjectCommitNewResponseStatusQueued, ProjectCommitNewResponseStatusRunning, ProjectCommitNewResponseStatusPaused, ProjectCommitNewResponseStatusFailed, ProjectCommitNewResponseStatusCompleted, ProjectCommitNewResponseStatusUnknown:
		return true
	}
	return false
}

type ProjectCommitNewResponseLinks struct {
	App  string                            `json:"app,required"`
	JSON projectCommitNewResponseLinksJSON `json:"-"`
}

// projectCommitNewResponseLinksJSON contains the JSON metadata for the struct
// [ProjectCommitNewResponseLinks]
type projectCommitNewResponseLinksJSON struct {
	App         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectCommitNewResponseLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectCommitNewResponseLinksJSON) RawJSON() string {
	return r.raw
}

type ProjectCommitListResponse struct {
	Items []ProjectCommitListResponseItem `json:"items,required"`
	JSON  projectCommitListResponseJSON   `json:"-"`
}

// projectCommitListResponseJSON contains the JSON metadata for the struct
// [ProjectCommitListResponse]
type projectCommitListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectCommitListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectCommitListResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectCommitListResponseItem struct {
	// The project version (commit) id.
	ID string `json:"id,required" format:"uuid"`
	// The details of a commit (project version).
	Commit ProjectCommitListResponseItemsCommit `json:"commit,required"`
	// The commit archive date.
	DateArchived time.Time `json:"dateArchived,required,nullable" format:"date-time"`
	// The project version (commit) creation date.
	DateCreated time.Time `json:"dateCreated,required" format:"date-time"`
	// The number of tests that are failing for the commit.
	FailingGoalCount int64 `json:"failingGoalCount,required"`
	// The model id.
	MlModelID string `json:"mlModelId,required,nullable" format:"uuid"`
	// The number of tests that are passing for the commit.
	PassingGoalCount int64 `json:"passingGoalCount,required"`
	// The project id.
	ProjectID string `json:"projectId,required" format:"uuid"`
	// The commit status. Initially, the commit is `queued`, then, it switches to
	// `running`. Finally, it can be `paused`, `failed`, or `completed`.
	Status ProjectCommitListResponseItemsStatus `json:"status,required"`
	// The commit status message.
	StatusMessage string `json:"statusMessage,required,nullable"`
	// The total number of tests for the commit.
	TotalGoalCount int64 `json:"totalGoalCount,required"`
	// The training dataset id.
	TrainingDatasetID string `json:"trainingDatasetId,required,nullable" format:"uuid"`
	// The validation dataset id.
	ValidationDatasetID string `json:"validationDatasetId,required,nullable" format:"uuid"`
	// Whether the commit is archived.
	Archived bool `json:"archived,nullable"`
	// The deployment status associated with the commit's model.
	DeploymentStatus string                              `json:"deploymentStatus"`
	Links            ProjectCommitListResponseItemsLinks `json:"links"`
	JSON             projectCommitListResponseItemJSON   `json:"-"`
}

// projectCommitListResponseItemJSON contains the JSON metadata for the struct
// [ProjectCommitListResponseItem]
type projectCommitListResponseItemJSON struct {
	ID                  apijson.Field
	Commit              apijson.Field
	DateArchived        apijson.Field
	DateCreated         apijson.Field
	FailingGoalCount    apijson.Field
	MlModelID           apijson.Field
	PassingGoalCount    apijson.Field
	ProjectID           apijson.Field
	Status              apijson.Field
	StatusMessage       apijson.Field
	TotalGoalCount      apijson.Field
	TrainingDatasetID   apijson.Field
	ValidationDatasetID apijson.Field
	Archived            apijson.Field
	DeploymentStatus    apijson.Field
	Links               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProjectCommitListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectCommitListResponseItemJSON) RawJSON() string {
	return r.raw
}

// The details of a commit (project version).
type ProjectCommitListResponseItemsCommit struct {
	// The commit id.
	ID string `json:"id,required" format:"uuid"`
	// The author id of the commit.
	AuthorID string `json:"authorId,required" format:"uuid"`
	// The size of the commit bundle in bytes.
	FileSize int64 `json:"fileSize,required,nullable"`
	// The commit message.
	Message string `json:"message,required"`
	// The model id.
	MlModelID string `json:"mlModelId,required,nullable" format:"uuid"`
	// The storage URI where the commit bundle is stored.
	StorageUri string `json:"storageUri,required"`
	// The training dataset id.
	TrainingDatasetID string `json:"trainingDatasetId,required,nullable" format:"uuid"`
	// The validation dataset id.
	ValidationDatasetID string `json:"validationDatasetId,required,nullable" format:"uuid"`
	// The commit creation date.
	DateCreated time.Time `json:"dateCreated" format:"date-time"`
	// The ref of the corresponding git commit.
	GitCommitRef string `json:"gitCommitRef"`
	// The SHA of the corresponding git commit.
	GitCommitSha int64 `json:"gitCommitSha"`
	// The URL of the corresponding git commit.
	GitCommitURL string                                   `json:"gitCommitUrl"`
	JSON         projectCommitListResponseItemsCommitJSON `json:"-"`
}

// projectCommitListResponseItemsCommitJSON contains the JSON metadata for the
// struct [ProjectCommitListResponseItemsCommit]
type projectCommitListResponseItemsCommitJSON struct {
	ID                  apijson.Field
	AuthorID            apijson.Field
	FileSize            apijson.Field
	Message             apijson.Field
	MlModelID           apijson.Field
	StorageUri          apijson.Field
	TrainingDatasetID   apijson.Field
	ValidationDatasetID apijson.Field
	DateCreated         apijson.Field
	GitCommitRef        apijson.Field
	GitCommitSha        apijson.Field
	GitCommitURL        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProjectCommitListResponseItemsCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectCommitListResponseItemsCommitJSON) RawJSON() string {
	return r.raw
}

// The commit status. Initially, the commit is `queued`, then, it switches to
// `running`. Finally, it can be `paused`, `failed`, or `completed`.
type ProjectCommitListResponseItemsStatus string

const (
	ProjectCommitListResponseItemsStatusQueued    ProjectCommitListResponseItemsStatus = "queued"
	ProjectCommitListResponseItemsStatusRunning   ProjectCommitListResponseItemsStatus = "running"
	ProjectCommitListResponseItemsStatusPaused    ProjectCommitListResponseItemsStatus = "paused"
	ProjectCommitListResponseItemsStatusFailed    ProjectCommitListResponseItemsStatus = "failed"
	ProjectCommitListResponseItemsStatusCompleted ProjectCommitListResponseItemsStatus = "completed"
	ProjectCommitListResponseItemsStatusUnknown   ProjectCommitListResponseItemsStatus = "unknown"
)

func (r ProjectCommitListResponseItemsStatus) IsKnown() bool {
	switch r {
	case ProjectCommitListResponseItemsStatusQueued, ProjectCommitListResponseItemsStatusRunning, ProjectCommitListResponseItemsStatusPaused, ProjectCommitListResponseItemsStatusFailed, ProjectCommitListResponseItemsStatusCompleted, ProjectCommitListResponseItemsStatusUnknown:
		return true
	}
	return false
}

type ProjectCommitListResponseItemsLinks struct {
	App  string                                  `json:"app,required"`
	JSON projectCommitListResponseItemsLinksJSON `json:"-"`
}

// projectCommitListResponseItemsLinksJSON contains the JSON metadata for the
// struct [ProjectCommitListResponseItemsLinks]
type projectCommitListResponseItemsLinksJSON struct {
	App         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectCommitListResponseItemsLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectCommitListResponseItemsLinksJSON) RawJSON() string {
	return r.raw
}

type ProjectCommitNewParams struct {
	// The details of a commit (project version).
	Commit param.Field[ProjectCommitNewParamsCommit] `json:"commit,required"`
	// The storage URI where the commit bundle is stored.
	StorageUri param.Field[string] `json:"storageUri,required"`
	// Whether the commit is archived.
	Archived param.Field[bool] `json:"archived"`
	// The deployment status associated with the commit's model.
	DeploymentStatus param.Field[string] `json:"deploymentStatus"`
}

func (r ProjectCommitNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The details of a commit (project version).
type ProjectCommitNewParamsCommit struct {
	// The commit message.
	Message param.Field[string] `json:"message,required"`
}

func (r ProjectCommitNewParamsCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectCommitListParams struct {
	// The page to return in a paginated query.
	Page param.Field[int64] `query:"page"`
	// Maximum number of items to return per page.
	PerPage param.Field[int64] `query:"perPage"`
}

// URLQuery serializes [ProjectCommitListParams]'s query parameters as
// `url.Values`.
func (r ProjectCommitListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
