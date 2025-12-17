// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/openlayer-ai/openlayer-go/internal/apijson"
	"github.com/openlayer-ai/openlayer-go/internal/param"
	"github.com/openlayer-ai/openlayer-go/internal/requestconfig"
	"github.com/openlayer-ai/openlayer-go/option"
)

// TestService contains methods and other services that help with interacting with
// the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTestService] method instead.
type TestService struct {
	Options []option.RequestOption
}

// NewTestService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTestService(opts ...option.RequestOption) (r *TestService) {
	r = &TestService{}
	r.Options = opts
	return
}

// Triggers one-off evaluation of a specific monitoring test for a custom timestamp
// range. This allows evaluating tests for historical data or custom time periods
// outside the regular evaluation window schedule. It also allows overwriting the
// existing test results.
func (r *TestService) Evaluate(ctx context.Context, testID string, body TestEvaluateParams, opts ...option.RequestOption) (res *TestEvaluateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if testID == "" {
		err = errors.New("missing required testId parameter")
		return
	}
	path := fmt.Sprintf("tests/%s/evaluate", testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TestEvaluateResponse struct {
	Message string `json:"message,required"`
	// Number of inference pipelines the test was queued for evaluation on
	PipelineCount int64 `json:"pipelineCount,required"`
	// The end timestamp you requested (in seconds)
	RequestedEndTimestamp int64 `json:"requestedEndTimestamp,required"`
	// The start timestamp you requested (in seconds)
	RequestedStartTimestamp int64 `json:"requestedStartTimestamp,required"`
	// Array of background task information for each pipeline evaluation
	Tasks []TestEvaluateResponseTask `json:"tasks,required"`
	JSON  testEvaluateResponseJSON   `json:"-"`
}

// testEvaluateResponseJSON contains the JSON metadata for the struct
// [TestEvaluateResponse]
type testEvaluateResponseJSON struct {
	Message                 apijson.Field
	PipelineCount           apijson.Field
	RequestedEndTimestamp   apijson.Field
	RequestedStartTimestamp apijson.Field
	Tasks                   apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *TestEvaluateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testEvaluateResponseJSON) RawJSON() string {
	return r.raw
}

type TestEvaluateResponseTask struct {
	// ID of the inference pipeline this task is for
	PipelineID string `json:"pipelineId,required" format:"uuid"`
	// ID of the background task
	TaskResultID string `json:"taskResultId,required" format:"uuid"`
	// URL to check the status of this background task
	TaskResultURL string                       `json:"taskResultUrl,required"`
	JSON          testEvaluateResponseTaskJSON `json:"-"`
}

// testEvaluateResponseTaskJSON contains the JSON metadata for the struct
// [TestEvaluateResponseTask]
type testEvaluateResponseTaskJSON struct {
	PipelineID    apijson.Field
	TaskResultID  apijson.Field
	TaskResultURL apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TestEvaluateResponseTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testEvaluateResponseTaskJSON) RawJSON() string {
	return r.raw
}

type TestEvaluateParams struct {
	// End timestamp in seconds (Unix epoch)
	EndTimestamp param.Field[int64] `json:"endTimestamp,required"`
	// Start timestamp in seconds (Unix epoch)
	StartTimestamp param.Field[int64] `json:"startTimestamp,required"`
	// ID of the inference pipeline to evaluate. If not provided, all inference
	// pipelines the test applies to will be evaluated.
	InferencePipelineID param.Field[string] `json:"inferencePipelineId" format:"uuid"`
	// Whether to overwrite existing test results
	OverwriteResults param.Field[bool] `json:"overwriteResults"`
}

func (r TestEvaluateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
