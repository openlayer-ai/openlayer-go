// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

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

// A list of rows for an inference pipeline.
func (r *InferencePipelineRowService) New(ctx context.Context, inferencePipelineID string, params InferencePipelineRowNewParams, opts ...option.RequestOption) (res *InferencePipelineRowNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if inferencePipelineID == "" {
		err = errors.New("missing required inferencePipelineId parameter")
		return
	}
	path := fmt.Sprintf("inference-pipelines/%s/rows", inferencePipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update an inference data point in an inference pipeline.
func (r *InferencePipelineRowService) Update(ctx context.Context, inferencePipelineID string, params InferencePipelineRowUpdateParams, opts ...option.RequestOption) (res *InferencePipelineRowUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if inferencePipelineID == "" {
		err = errors.New("missing required inferencePipelineId parameter")
		return
	}
	path := fmt.Sprintf("inference-pipelines/%s/rows", inferencePipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type InferencePipelineRowNewResponse struct {
	Items []InferencePipelineRowNewResponseItem `json:"items" api:"required"`
	JSON  inferencePipelineRowNewResponseJSON   `json:"-"`
}

// inferencePipelineRowNewResponseJSON contains the JSON metadata for the struct
// [InferencePipelineRowNewResponse]
type inferencePipelineRowNewResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InferencePipelineRowNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineRowNewResponseJSON) RawJSON() string {
	return r.raw
}

type InferencePipelineRowNewResponseItem struct {
	OpenlayerRowID int64                                   `json:"openlayer_row_id" api:"required"`
	JSON           inferencePipelineRowNewResponseItemJSON `json:"-"`
}

// inferencePipelineRowNewResponseItemJSON contains the JSON metadata for the
// struct [InferencePipelineRowNewResponseItem]
type inferencePipelineRowNewResponseItemJSON struct {
	OpenlayerRowID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InferencePipelineRowNewResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineRowNewResponseItemJSON) RawJSON() string {
	return r.raw
}

type InferencePipelineRowUpdateResponse struct {
	Success InferencePipelineRowUpdateResponseSuccess `json:"success" api:"required"`
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

type InferencePipelineRowNewParams struct {
	// Whether or not to sort on the sortColumn in ascending order.
	Asc param.Field[bool] `query:"asc"`
	// The page to return in a paginated query.
	Page param.Field[int64] `query:"page"`
	// Maximum number of items to return per page.
	PerPage param.Field[int64] `query:"perPage"`
	// Name of the column to sort on
	SortColumn        param.Field[string]                                           `query:"sortColumn"`
	ColumnFilters     param.Field[[]InferencePipelineRowNewParamsColumnFilterUnion] `json:"columnFilters"`
	ExcludeRowIDList  param.Field[[]int64]                                          `json:"excludeRowIdList"`
	NotSearchQueryAnd param.Field[[]string]                                         `json:"notSearchQueryAnd"`
	NotSearchQueryOr  param.Field[[]string]                                         `json:"notSearchQueryOr"`
	RowIDList         param.Field[[]int64]                                          `json:"rowIdList"`
	SearchQueryAnd    param.Field[[]string]                                         `json:"searchQueryAnd"`
	SearchQueryOr     param.Field[[]string]                                         `json:"searchQueryOr"`
}

func (r InferencePipelineRowNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [InferencePipelineRowNewParams]'s query parameters as
// `url.Values`.
func (r InferencePipelineRowNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InferencePipelineRowNewParamsColumnFilter struct {
	// The name of the column.
	Measurement param.Field[string]                                             `json:"measurement" api:"required"`
	Operator    param.Field[InferencePipelineRowNewParamsColumnFiltersOperator] `json:"operator" api:"required"`
	Value       param.Field[interface{}]                                        `json:"value" api:"required"`
}

func (r InferencePipelineRowNewParamsColumnFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InferencePipelineRowNewParamsColumnFilter) implementsInferencePipelineRowNewParamsColumnFilterUnion() {
}

// Satisfied by [InferencePipelineRowNewParamsColumnFiltersSetColumnFilter],
// [InferencePipelineRowNewParamsColumnFiltersNumericColumnFilter],
// [InferencePipelineRowNewParamsColumnFiltersStringColumnFilter],
// [InferencePipelineRowNewParamsColumnFilter].
type InferencePipelineRowNewParamsColumnFilterUnion interface {
	implementsInferencePipelineRowNewParamsColumnFilterUnion()
}

type InferencePipelineRowNewParamsColumnFiltersSetColumnFilter struct {
	// The name of the column.
	Measurement param.Field[string]                                                                `json:"measurement" api:"required"`
	Operator    param.Field[InferencePipelineRowNewParamsColumnFiltersSetColumnFilterOperator]     `json:"operator" api:"required"`
	Value       param.Field[[]InferencePipelineRowNewParamsColumnFiltersSetColumnFilterValueUnion] `json:"value" api:"required"`
}

func (r InferencePipelineRowNewParamsColumnFiltersSetColumnFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InferencePipelineRowNewParamsColumnFiltersSetColumnFilter) implementsInferencePipelineRowNewParamsColumnFilterUnion() {
}

type InferencePipelineRowNewParamsColumnFiltersSetColumnFilterOperator string

const (
	InferencePipelineRowNewParamsColumnFiltersSetColumnFilterOperatorContainsNone InferencePipelineRowNewParamsColumnFiltersSetColumnFilterOperator = "contains_none"
	InferencePipelineRowNewParamsColumnFiltersSetColumnFilterOperatorContainsAny  InferencePipelineRowNewParamsColumnFiltersSetColumnFilterOperator = "contains_any"
	InferencePipelineRowNewParamsColumnFiltersSetColumnFilterOperatorContainsAll  InferencePipelineRowNewParamsColumnFiltersSetColumnFilterOperator = "contains_all"
	InferencePipelineRowNewParamsColumnFiltersSetColumnFilterOperatorOneOf        InferencePipelineRowNewParamsColumnFiltersSetColumnFilterOperator = "one_of"
	InferencePipelineRowNewParamsColumnFiltersSetColumnFilterOperatorNoneOf       InferencePipelineRowNewParamsColumnFiltersSetColumnFilterOperator = "none_of"
)

func (r InferencePipelineRowNewParamsColumnFiltersSetColumnFilterOperator) IsKnown() bool {
	switch r {
	case InferencePipelineRowNewParamsColumnFiltersSetColumnFilterOperatorContainsNone, InferencePipelineRowNewParamsColumnFiltersSetColumnFilterOperatorContainsAny, InferencePipelineRowNewParamsColumnFiltersSetColumnFilterOperatorContainsAll, InferencePipelineRowNewParamsColumnFiltersSetColumnFilterOperatorOneOf, InferencePipelineRowNewParamsColumnFiltersSetColumnFilterOperatorNoneOf:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type InferencePipelineRowNewParamsColumnFiltersSetColumnFilterValueUnion interface {
	ImplementsInferencePipelineRowNewParamsColumnFiltersSetColumnFilterValueUnion()
}

type InferencePipelineRowNewParamsColumnFiltersNumericColumnFilter struct {
	// The name of the column.
	Measurement param.Field[string]                                                                `json:"measurement" api:"required"`
	Operator    param.Field[InferencePipelineRowNewParamsColumnFiltersNumericColumnFilterOperator] `json:"operator" api:"required"`
	Value       param.Field[float64]                                                               `json:"value" api:"required"`
}

func (r InferencePipelineRowNewParamsColumnFiltersNumericColumnFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InferencePipelineRowNewParamsColumnFiltersNumericColumnFilter) implementsInferencePipelineRowNewParamsColumnFilterUnion() {
}

type InferencePipelineRowNewParamsColumnFiltersNumericColumnFilterOperator string

const (
	InferencePipelineRowNewParamsColumnFiltersNumericColumnFilterOperatorGreater         InferencePipelineRowNewParamsColumnFiltersNumericColumnFilterOperator = ">"
	InferencePipelineRowNewParamsColumnFiltersNumericColumnFilterOperatorGreaterOrEquals InferencePipelineRowNewParamsColumnFiltersNumericColumnFilterOperator = ">="
	InferencePipelineRowNewParamsColumnFiltersNumericColumnFilterOperatorIs              InferencePipelineRowNewParamsColumnFiltersNumericColumnFilterOperator = "is"
	InferencePipelineRowNewParamsColumnFiltersNumericColumnFilterOperatorLess            InferencePipelineRowNewParamsColumnFiltersNumericColumnFilterOperator = "<"
	InferencePipelineRowNewParamsColumnFiltersNumericColumnFilterOperatorLessOrEquals    InferencePipelineRowNewParamsColumnFiltersNumericColumnFilterOperator = "<="
	InferencePipelineRowNewParamsColumnFiltersNumericColumnFilterOperatorNotEquals       InferencePipelineRowNewParamsColumnFiltersNumericColumnFilterOperator = "!="
)

func (r InferencePipelineRowNewParamsColumnFiltersNumericColumnFilterOperator) IsKnown() bool {
	switch r {
	case InferencePipelineRowNewParamsColumnFiltersNumericColumnFilterOperatorGreater, InferencePipelineRowNewParamsColumnFiltersNumericColumnFilterOperatorGreaterOrEquals, InferencePipelineRowNewParamsColumnFiltersNumericColumnFilterOperatorIs, InferencePipelineRowNewParamsColumnFiltersNumericColumnFilterOperatorLess, InferencePipelineRowNewParamsColumnFiltersNumericColumnFilterOperatorLessOrEquals, InferencePipelineRowNewParamsColumnFiltersNumericColumnFilterOperatorNotEquals:
		return true
	}
	return false
}

type InferencePipelineRowNewParamsColumnFiltersStringColumnFilter struct {
	// The name of the column.
	Measurement param.Field[string]                                                                 `json:"measurement" api:"required"`
	Operator    param.Field[InferencePipelineRowNewParamsColumnFiltersStringColumnFilterOperator]   `json:"operator" api:"required"`
	Value       param.Field[InferencePipelineRowNewParamsColumnFiltersStringColumnFilterValueUnion] `json:"value" api:"required"`
}

func (r InferencePipelineRowNewParamsColumnFiltersStringColumnFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InferencePipelineRowNewParamsColumnFiltersStringColumnFilter) implementsInferencePipelineRowNewParamsColumnFilterUnion() {
}

type InferencePipelineRowNewParamsColumnFiltersStringColumnFilterOperator string

const (
	InferencePipelineRowNewParamsColumnFiltersStringColumnFilterOperatorIs        InferencePipelineRowNewParamsColumnFiltersStringColumnFilterOperator = "is"
	InferencePipelineRowNewParamsColumnFiltersStringColumnFilterOperatorNotEquals InferencePipelineRowNewParamsColumnFiltersStringColumnFilterOperator = "!="
)

func (r InferencePipelineRowNewParamsColumnFiltersStringColumnFilterOperator) IsKnown() bool {
	switch r {
	case InferencePipelineRowNewParamsColumnFiltersStringColumnFilterOperatorIs, InferencePipelineRowNewParamsColumnFiltersStringColumnFilterOperatorNotEquals:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type InferencePipelineRowNewParamsColumnFiltersStringColumnFilterValueUnion interface {
	ImplementsInferencePipelineRowNewParamsColumnFiltersStringColumnFilterValueUnion()
}

type InferencePipelineRowNewParamsColumnFiltersOperator string

const (
	InferencePipelineRowNewParamsColumnFiltersOperatorContainsNone    InferencePipelineRowNewParamsColumnFiltersOperator = "contains_none"
	InferencePipelineRowNewParamsColumnFiltersOperatorContainsAny     InferencePipelineRowNewParamsColumnFiltersOperator = "contains_any"
	InferencePipelineRowNewParamsColumnFiltersOperatorContainsAll     InferencePipelineRowNewParamsColumnFiltersOperator = "contains_all"
	InferencePipelineRowNewParamsColumnFiltersOperatorOneOf           InferencePipelineRowNewParamsColumnFiltersOperator = "one_of"
	InferencePipelineRowNewParamsColumnFiltersOperatorNoneOf          InferencePipelineRowNewParamsColumnFiltersOperator = "none_of"
	InferencePipelineRowNewParamsColumnFiltersOperatorGreater         InferencePipelineRowNewParamsColumnFiltersOperator = ">"
	InferencePipelineRowNewParamsColumnFiltersOperatorGreaterOrEquals InferencePipelineRowNewParamsColumnFiltersOperator = ">="
	InferencePipelineRowNewParamsColumnFiltersOperatorIs              InferencePipelineRowNewParamsColumnFiltersOperator = "is"
	InferencePipelineRowNewParamsColumnFiltersOperatorLess            InferencePipelineRowNewParamsColumnFiltersOperator = "<"
	InferencePipelineRowNewParamsColumnFiltersOperatorLessOrEquals    InferencePipelineRowNewParamsColumnFiltersOperator = "<="
	InferencePipelineRowNewParamsColumnFiltersOperatorNotEquals       InferencePipelineRowNewParamsColumnFiltersOperator = "!="
)

func (r InferencePipelineRowNewParamsColumnFiltersOperator) IsKnown() bool {
	switch r {
	case InferencePipelineRowNewParamsColumnFiltersOperatorContainsNone, InferencePipelineRowNewParamsColumnFiltersOperatorContainsAny, InferencePipelineRowNewParamsColumnFiltersOperatorContainsAll, InferencePipelineRowNewParamsColumnFiltersOperatorOneOf, InferencePipelineRowNewParamsColumnFiltersOperatorNoneOf, InferencePipelineRowNewParamsColumnFiltersOperatorGreater, InferencePipelineRowNewParamsColumnFiltersOperatorGreaterOrEquals, InferencePipelineRowNewParamsColumnFiltersOperatorIs, InferencePipelineRowNewParamsColumnFiltersOperatorLess, InferencePipelineRowNewParamsColumnFiltersOperatorLessOrEquals, InferencePipelineRowNewParamsColumnFiltersOperatorNotEquals:
		return true
	}
	return false
}

type InferencePipelineRowUpdateParams struct {
	// Specify the inference id as a query param.
	InferenceID param.Field[string]                                 `query:"inferenceId" api:"required"`
	Row         param.Field[interface{}]                            `json:"row" api:"required"`
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
