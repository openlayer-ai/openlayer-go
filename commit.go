// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/openlayer-ai/openlayer-go/internal/apijson"
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

// Retrieve a project version (commit) by its id.
func (r *CommitService) Get(ctx context.Context, projectVersionID string, opts ...option.RequestOption) (res *CommitGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if projectVersionID == "" {
		err = errors.New("missing required projectVersionId parameter")
		return
	}
	path := fmt.Sprintf("versions/%s", projectVersionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CommitGetResponse struct {
	// The project version (commit) id.
	ID string `json:"id,required" format:"uuid"`
	// The details of a commit (project version).
	Commit CommitGetResponseCommit `json:"commit,required"`
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
	Status CommitGetResponseStatus `json:"status,required"`
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
	Links            CommitGetResponseLinks `json:"links"`
	JSON             commitGetResponseJSON  `json:"-"`
}

// commitGetResponseJSON contains the JSON metadata for the struct
// [CommitGetResponse]
type commitGetResponseJSON struct {
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

func (r *CommitGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitGetResponseJSON) RawJSON() string {
	return r.raw
}

// The details of a commit (project version).
type CommitGetResponseCommit struct {
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
	JSON         commitGetResponseCommitJSON `json:"-"`
}

// commitGetResponseCommitJSON contains the JSON metadata for the struct
// [CommitGetResponseCommit]
type commitGetResponseCommitJSON struct {
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

func (r *CommitGetResponseCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitGetResponseCommitJSON) RawJSON() string {
	return r.raw
}

// The commit status. Initially, the commit is `queued`, then, it switches to
// `running`. Finally, it can be `paused`, `failed`, or `completed`.
type CommitGetResponseStatus string

const (
	CommitGetResponseStatusQueued    CommitGetResponseStatus = "queued"
	CommitGetResponseStatusRunning   CommitGetResponseStatus = "running"
	CommitGetResponseStatusPaused    CommitGetResponseStatus = "paused"
	CommitGetResponseStatusFailed    CommitGetResponseStatus = "failed"
	CommitGetResponseStatusCompleted CommitGetResponseStatus = "completed"
	CommitGetResponseStatusUnknown   CommitGetResponseStatus = "unknown"
)

func (r CommitGetResponseStatus) IsKnown() bool {
	switch r {
	case CommitGetResponseStatusQueued, CommitGetResponseStatusRunning, CommitGetResponseStatusPaused, CommitGetResponseStatusFailed, CommitGetResponseStatusCompleted, CommitGetResponseStatusUnknown:
		return true
	}
	return false
}

type CommitGetResponseLinks struct {
	App  string                     `json:"app,required"`
	JSON commitGetResponseLinksJSON `json:"-"`
}

// commitGetResponseLinksJSON contains the JSON metadata for the struct
// [CommitGetResponseLinks]
type commitGetResponseLinksJSON struct {
	App         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitGetResponseLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitGetResponseLinksJSON) RawJSON() string {
	return r.raw
}
