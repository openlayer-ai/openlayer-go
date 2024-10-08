// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/openlayer-ai/openlayer-go/internal/apijson"
	"github.com/openlayer-ai/openlayer-go/internal/apiquery"
	"github.com/openlayer-ai/openlayer-go/internal/param"
	"github.com/openlayer-ai/openlayer-go/internal/requestconfig"
	"github.com/openlayer-ai/openlayer-go/option"
)

// InferencePipelineRowService contains methods and other services that help with
// interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInferencePipelineRowService] method instead.
type InferencePipelineRowService struct {
	Options []option.RequestOption
}

// NewInferencePipelineRowService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInferencePipelineRowService(opts ...option.RequestOption) (r *InferencePipelineRowService) {
	r = &InferencePipelineRowService{}
	r.Options = opts
	return
}

// Update an inference data point in an inference pipeline.
func (r *InferencePipelineRowService) Update(ctx context.Context, inferencePipelineID string, params InferencePipelineRowUpdateParams, opts ...option.RequestOption) (res *InferencePipelineRowUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if inferencePipelineID == "" {
		err = errors.New("missing required inferencePipelineId parameter")
		return
	}
	path := fmt.Sprintf("inference-pipelines/%s/rows", inferencePipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type InferencePipelineRowUpdateResponse struct {
	Success InferencePipelineRowUpdateResponseSuccess `json:"success,required"`
	JSON    inferencePipelineRowUpdateResponseJSON    `json:"-"`
}

// inferencePipelineRowUpdateResponseJSON contains the JSON metadata for the struct
// [InferencePipelineRowUpdateResponse]
type inferencePipelineRowUpdateResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InferencePipelineRowUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineRowUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type InferencePipelineRowUpdateResponseSuccess bool

const (
	InferencePipelineRowUpdateResponseSuccessTrue InferencePipelineRowUpdateResponseSuccess = true
)

func (r InferencePipelineRowUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case InferencePipelineRowUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type InferencePipelineRowUpdateParams struct {
	// Specify the inference id as a query param.
	InferenceID param.Field[string]                                 `query:"inferenceId,required"`
	Row         param.Field[interface{}]                            `json:"row,required"`
	Config      param.Field[InferencePipelineRowUpdateParamsConfig] `json:"config"`
}

func (r InferencePipelineRowUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [InferencePipelineRowUpdateParams]'s query parameters as
// `url.Values`.
func (r InferencePipelineRowUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InferencePipelineRowUpdateParamsConfig struct {
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

func (r InferencePipelineRowUpdateParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
