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
func (r *InferencePipelineService) Get(ctx context.Context, inferencePipelineID string, opts ...option.RequestOption) (res *InferencePipelineGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if inferencePipelineID == "" {
		err = errors.New("missing required inferencePipelineId parameter")
		return
	}
	path := fmt.Sprintf("inference-pipelines/%s", inferencePipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
	TotalGoalCount int64                            `json:"totalGoalCount,required"`
	JSON           inferencePipelineGetResponseJSON `json:"-"`
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
	TotalGoalCount int64                               `json:"totalGoalCount,required"`
	JSON           inferencePipelineUpdateResponseJSON `json:"-"`
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
