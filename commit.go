// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/openlayer-ai/openlayer-go/internal/apijson"
	"github.com/openlayer-ai/openlayer-go/internal/param"
	"github.com/openlayer-ai/openlayer-go/internal/requestconfig"
	"github.com/openlayer-ai/openlayer-go/option"
)

// CommitService contains methods and other services that help with interacting
// with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCommitService] method instead.
type CommitService struct {
	Options     []option.RequestOption
	TestResults *CommitTestResultService
}

// NewCommitService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCommitService(opts ...option.RequestOption) (r *CommitService) {
	r = &CommitService{}
	r.Options = opts
	r.TestResults = NewCommitTestResultService(opts...)
	return
}

// Create a new commit (project version) in a project.
func (r *CommitService) New(ctx context.Context, projectID string, body CommitNewParams, opts ...option.RequestOption) (res *CommitNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return
	}
	path := fmt.Sprintf("projects/%s/versions", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CommitNewResponse struct {
	// The project version (commit) id.
	ID string `json:"id,required" format:"uuid"`
	// The details of a commit (project version).
	Commit CommitNewResponseCommit `json:"commit,required"`
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
	Status CommitNewResponseStatus `json:"status,required"`
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
	DeploymentStatus string                 `json:"deploymentStatus"`
	Links            CommitNewResponseLinks `json:"links"`
	JSON             commitNewResponseJSON  `json:"-"`
}

// commitNewResponseJSON contains the JSON metadata for the struct
// [CommitNewResponse]
type commitNewResponseJSON struct {
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

func (r *CommitNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitNewResponseJSON) RawJSON() string {
	return r.raw
}

// The details of a commit (project version).
type CommitNewResponseCommit struct {
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
	GitCommitURL string                      `json:"gitCommitUrl"`
	JSON         commitNewResponseCommitJSON `json:"-"`
}

// commitNewResponseCommitJSON contains the JSON metadata for the struct
// [CommitNewResponseCommit]
type commitNewResponseCommitJSON struct {
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

func (r *CommitNewResponseCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitNewResponseCommitJSON) RawJSON() string {
	return r.raw
}

// The commit status. Initially, the commit is `queued`, then, it switches to
// `running`. Finally, it can be `paused`, `failed`, or `completed`.
type CommitNewResponseStatus string

const (
	CommitNewResponseStatusQueued    CommitNewResponseStatus = "queued"
	CommitNewResponseStatusRunning   CommitNewResponseStatus = "running"
	CommitNewResponseStatusPaused    CommitNewResponseStatus = "paused"
	CommitNewResponseStatusFailed    CommitNewResponseStatus = "failed"
	CommitNewResponseStatusCompleted CommitNewResponseStatus = "completed"
	CommitNewResponseStatusUnknown   CommitNewResponseStatus = "unknown"
)

func (r CommitNewResponseStatus) IsKnown() bool {
	switch r {
	case CommitNewResponseStatusQueued, CommitNewResponseStatusRunning, CommitNewResponseStatusPaused, CommitNewResponseStatusFailed, CommitNewResponseStatusCompleted, CommitNewResponseStatusUnknown:
		return true
	}
	return false
}

type CommitNewResponseLinks struct {
	App  string                     `json:"app,required"`
	JSON commitNewResponseLinksJSON `json:"-"`
}

// commitNewResponseLinksJSON contains the JSON metadata for the struct
// [CommitNewResponseLinks]
type commitNewResponseLinksJSON struct {
	App         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitNewResponseLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitNewResponseLinksJSON) RawJSON() string {
	return r.raw
}

type CommitNewParams struct {
	// The details of a commit (project version).
	Commit param.Field[CommitNewParamsCommit] `json:"commit,required"`
	// The storage URI where the commit bundle is stored.
	StorageUri param.Field[string] `json:"storageUri,required"`
	// Whether the commit is archived.
	Archived param.Field[bool] `json:"archived"`
	// The deployment status associated with the commit's model.
	DeploymentStatus param.Field[string] `json:"deploymentStatus"`
}

func (r CommitNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The details of a commit (project version).
type CommitNewParamsCommit struct {
	// The commit message.
	Message param.Field[string] `json:"message,required"`
}

func (r CommitNewParamsCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
