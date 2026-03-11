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

// Update an inference data point in an inference pipeline.
func (r *InferencePipelineRowService) Update(ctx context.Context, inferencePipelineID string, params InferencePipelineRowUpdateParams, opts ...option.RequestOption) (res *InferencePipelineRowUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if inferencePipelineID == "" {
		err = errors.New("missing required inferencePipelineId parameter")
		return nil, err
	}
	path := fmt.Sprintf("inference-pipelines/%s/rows", inferencePipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// A list of rows for an inference pipeline.
func (r *InferencePipelineRowService) List(ctx context.Context, inferencePipelineID string, params InferencePipelineRowListParams, opts ...option.RequestOption) (res *InferencePipelineRowListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if inferencePipelineID == "" {
		err = errors.New("missing required inferencePipelineId parameter")
		return nil, err
	}
	path := fmt.Sprintf("inference-pipelines/%s/rows", inferencePipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
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

type InferencePipelineRowListResponse struct {
	Items []InferencePipelineRowListResponseItem `json:"items" api:"required"`
	JSON  inferencePipelineRowListResponseJSON   `json:"-"`
}

// inferencePipelineRowListResponseJSON contains the JSON metadata for the struct
// [InferencePipelineRowListResponse]
type inferencePipelineRowListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InferencePipelineRowListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineRowListResponseJSON) RawJSON() string {
	return r.raw
}

type InferencePipelineRowListResponseItem struct {
	OpenlayerRowID int64                                    `json:"openlayer_row_id" api:"required"`
	JSON           inferencePipelineRowListResponseItemJSON `json:"-"`
}

// inferencePipelineRowListResponseItemJSON contains the JSON metadata for the
// struct [InferencePipelineRowListResponseItem]
type inferencePipelineRowListResponseItemJSON struct {
	OpenlayerRowID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InferencePipelineRowListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineRowListResponseItemJSON) RawJSON() string {
	return r.raw
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

type InferencePipelineRowListParams struct {
	// Whether or not to sort on the sortColumn in ascending order.
	Asc param.Field[bool] `query:"asc"`
	// The page to return in a paginated query.
	Page param.Field[int64] `query:"page"`
	// Maximum number of items to return per page.
	PerPage param.Field[int64] `query:"perPage"`
	// Name of the column to sort on
	SortColumn        param.Field[string]                                            `query:"sortColumn"`
	ColumnFilters     param.Field[[]InferencePipelineRowListParamsColumnFilterUnion] `json:"columnFilters"`
	ExcludeRowIDList  param.Field[[]int64]                                           `json:"excludeRowIdList"`
	NotSearchQueryAnd param.Field[[]string]                                          `json:"notSearchQueryAnd"`
	NotSearchQueryOr  param.Field[[]string]                                          `json:"notSearchQueryOr"`
	RowIDList         param.Field[[]int64]                                           `json:"rowIdList"`
	SearchQueryAnd    param.Field[[]string]                                          `json:"searchQueryAnd"`
	SearchQueryOr     param.Field[[]string]                                          `json:"searchQueryOr"`
}

func (r InferencePipelineRowListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [InferencePipelineRowListParams]'s query parameters as
// `url.Values`.
func (r InferencePipelineRowListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InferencePipelineRowListParamsColumnFilter struct {
	// The name of the column.
	Measurement param.Field[string]                                              `json:"measurement" api:"required"`
	Operator    param.Field[InferencePipelineRowListParamsColumnFiltersOperator] `json:"operator" api:"required"`
	Value       param.Field[interface{}]                                         `json:"value" api:"required"`
}

func (r InferencePipelineRowListParamsColumnFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InferencePipelineRowListParamsColumnFilter) implementsInferencePipelineRowListParamsColumnFilterUnion() {
}

// Satisfied by [InferencePipelineRowListParamsColumnFiltersSetColumnFilter],
// [InferencePipelineRowListParamsColumnFiltersNumericColumnFilter],
// [InferencePipelineRowListParamsColumnFiltersStringColumnFilter],
// [InferencePipelineRowListParamsColumnFilter].
type InferencePipelineRowListParamsColumnFilterUnion interface {
	implementsInferencePipelineRowListParamsColumnFilterUnion()
}

type InferencePipelineRowListParamsColumnFiltersSetColumnFilter struct {
	// The name of the column.
	Measurement param.Field[string]                                                                 `json:"measurement" api:"required"`
	Operator    param.Field[InferencePipelineRowListParamsColumnFiltersSetColumnFilterOperator]     `json:"operator" api:"required"`
	Value       param.Field[[]InferencePipelineRowListParamsColumnFiltersSetColumnFilterValueUnion] `json:"value" api:"required"`
}

func (r InferencePipelineRowListParamsColumnFiltersSetColumnFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InferencePipelineRowListParamsColumnFiltersSetColumnFilter) implementsInferencePipelineRowListParamsColumnFilterUnion() {
}

type InferencePipelineRowListParamsColumnFiltersSetColumnFilterOperator string

const (
	InferencePipelineRowListParamsColumnFiltersSetColumnFilterOperatorContainsNone InferencePipelineRowListParamsColumnFiltersSetColumnFilterOperator = "contains_none"
	InferencePipelineRowListParamsColumnFiltersSetColumnFilterOperatorContainsAny  InferencePipelineRowListParamsColumnFiltersSetColumnFilterOperator = "contains_any"
	InferencePipelineRowListParamsColumnFiltersSetColumnFilterOperatorContainsAll  InferencePipelineRowListParamsColumnFiltersSetColumnFilterOperator = "contains_all"
	InferencePipelineRowListParamsColumnFiltersSetColumnFilterOperatorOneOf        InferencePipelineRowListParamsColumnFiltersSetColumnFilterOperator = "one_of"
	InferencePipelineRowListParamsColumnFiltersSetColumnFilterOperatorNoneOf       InferencePipelineRowListParamsColumnFiltersSetColumnFilterOperator = "none_of"
)

func (r InferencePipelineRowListParamsColumnFiltersSetColumnFilterOperator) IsKnown() bool {
	switch r {
	case InferencePipelineRowListParamsColumnFiltersSetColumnFilterOperatorContainsNone, InferencePipelineRowListParamsColumnFiltersSetColumnFilterOperatorContainsAny, InferencePipelineRowListParamsColumnFiltersSetColumnFilterOperatorContainsAll, InferencePipelineRowListParamsColumnFiltersSetColumnFilterOperatorOneOf, InferencePipelineRowListParamsColumnFiltersSetColumnFilterOperatorNoneOf:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type InferencePipelineRowListParamsColumnFiltersSetColumnFilterValueUnion interface {
	ImplementsInferencePipelineRowListParamsColumnFiltersSetColumnFilterValueUnion()
}

type InferencePipelineRowListParamsColumnFiltersNumericColumnFilter struct {
	// The name of the column.
	Measurement param.Field[string]                                                                 `json:"measurement" api:"required"`
	Operator    param.Field[InferencePipelineRowListParamsColumnFiltersNumericColumnFilterOperator] `json:"operator" api:"required"`
	Value       param.Field[float64]                                                                `json:"value" api:"required"`
}

func (r InferencePipelineRowListParamsColumnFiltersNumericColumnFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InferencePipelineRowListParamsColumnFiltersNumericColumnFilter) implementsInferencePipelineRowListParamsColumnFilterUnion() {
}

type InferencePipelineRowListParamsColumnFiltersNumericColumnFilterOperator string

const (
	InferencePipelineRowListParamsColumnFiltersNumericColumnFilterOperatorGreater         InferencePipelineRowListParamsColumnFiltersNumericColumnFilterOperator = ">"
	InferencePipelineRowListParamsColumnFiltersNumericColumnFilterOperatorGreaterOrEquals InferencePipelineRowListParamsColumnFiltersNumericColumnFilterOperator = ">="
	InferencePipelineRowListParamsColumnFiltersNumericColumnFilterOperatorIs              InferencePipelineRowListParamsColumnFiltersNumericColumnFilterOperator = "is"
	InferencePipelineRowListParamsColumnFiltersNumericColumnFilterOperatorLess            InferencePipelineRowListParamsColumnFiltersNumericColumnFilterOperator = "<"
	InferencePipelineRowListParamsColumnFiltersNumericColumnFilterOperatorLessOrEquals    InferencePipelineRowListParamsColumnFiltersNumericColumnFilterOperator = "<="
	InferencePipelineRowListParamsColumnFiltersNumericColumnFilterOperatorNotEquals       InferencePipelineRowListParamsColumnFiltersNumericColumnFilterOperator = "!="
)

func (r InferencePipelineRowListParamsColumnFiltersNumericColumnFilterOperator) IsKnown() bool {
	switch r {
	case InferencePipelineRowListParamsColumnFiltersNumericColumnFilterOperatorGreater, InferencePipelineRowListParamsColumnFiltersNumericColumnFilterOperatorGreaterOrEquals, InferencePipelineRowListParamsColumnFiltersNumericColumnFilterOperatorIs, InferencePipelineRowListParamsColumnFiltersNumericColumnFilterOperatorLess, InferencePipelineRowListParamsColumnFiltersNumericColumnFilterOperatorLessOrEquals, InferencePipelineRowListParamsColumnFiltersNumericColumnFilterOperatorNotEquals:
		return true
	}
	return false
}

type InferencePipelineRowListParamsColumnFiltersStringColumnFilter struct {
	// The name of the column.
	Measurement param.Field[string]                                                                  `json:"measurement" api:"required"`
	Operator    param.Field[InferencePipelineRowListParamsColumnFiltersStringColumnFilterOperator]   `json:"operator" api:"required"`
	Value       param.Field[InferencePipelineRowListParamsColumnFiltersStringColumnFilterValueUnion] `json:"value" api:"required"`
}

func (r InferencePipelineRowListParamsColumnFiltersStringColumnFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InferencePipelineRowListParamsColumnFiltersStringColumnFilter) implementsInferencePipelineRowListParamsColumnFilterUnion() {
}

type InferencePipelineRowListParamsColumnFiltersStringColumnFilterOperator string

const (
	InferencePipelineRowListParamsColumnFiltersStringColumnFilterOperatorIs        InferencePipelineRowListParamsColumnFiltersStringColumnFilterOperator = "is"
	InferencePipelineRowListParamsColumnFiltersStringColumnFilterOperatorNotEquals InferencePipelineRowListParamsColumnFiltersStringColumnFilterOperator = "!="
)

func (r InferencePipelineRowListParamsColumnFiltersStringColumnFilterOperator) IsKnown() bool {
	switch r {
	case InferencePipelineRowListParamsColumnFiltersStringColumnFilterOperatorIs, InferencePipelineRowListParamsColumnFiltersStringColumnFilterOperatorNotEquals:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type InferencePipelineRowListParamsColumnFiltersStringColumnFilterValueUnion interface {
	ImplementsInferencePipelineRowListParamsColumnFiltersStringColumnFilterValueUnion()
}

type InferencePipelineRowListParamsColumnFiltersOperator string

const (
	InferencePipelineRowListParamsColumnFiltersOperatorContainsNone    InferencePipelineRowListParamsColumnFiltersOperator = "contains_none"
	InferencePipelineRowListParamsColumnFiltersOperatorContainsAny     InferencePipelineRowListParamsColumnFiltersOperator = "contains_any"
	InferencePipelineRowListParamsColumnFiltersOperatorContainsAll     InferencePipelineRowListParamsColumnFiltersOperator = "contains_all"
	InferencePipelineRowListParamsColumnFiltersOperatorOneOf           InferencePipelineRowListParamsColumnFiltersOperator = "one_of"
	InferencePipelineRowListParamsColumnFiltersOperatorNoneOf          InferencePipelineRowListParamsColumnFiltersOperator = "none_of"
	InferencePipelineRowListParamsColumnFiltersOperatorGreater         InferencePipelineRowListParamsColumnFiltersOperator = ">"
	InferencePipelineRowListParamsColumnFiltersOperatorGreaterOrEquals InferencePipelineRowListParamsColumnFiltersOperator = ">="
	InferencePipelineRowListParamsColumnFiltersOperatorIs              InferencePipelineRowListParamsColumnFiltersOperator = "is"
	InferencePipelineRowListParamsColumnFiltersOperatorLess            InferencePipelineRowListParamsColumnFiltersOperator = "<"
	InferencePipelineRowListParamsColumnFiltersOperatorLessOrEquals    InferencePipelineRowListParamsColumnFiltersOperator = "<="
	InferencePipelineRowListParamsColumnFiltersOperatorNotEquals       InferencePipelineRowListParamsColumnFiltersOperator = "!="
)

func (r InferencePipelineRowListParamsColumnFiltersOperator) IsKnown() bool {
	switch r {
	case InferencePipelineRowListParamsColumnFiltersOperatorContainsNone, InferencePipelineRowListParamsColumnFiltersOperatorContainsAny, InferencePipelineRowListParamsColumnFiltersOperatorContainsAll, InferencePipelineRowListParamsColumnFiltersOperatorOneOf, InferencePipelineRowListParamsColumnFiltersOperatorNoneOf, InferencePipelineRowListParamsColumnFiltersOperatorGreater, InferencePipelineRowListParamsColumnFiltersOperatorGreaterOrEquals, InferencePipelineRowListParamsColumnFiltersOperatorIs, InferencePipelineRowListParamsColumnFiltersOperatorLess, InferencePipelineRowListParamsColumnFiltersOperatorLessOrEquals, InferencePipelineRowListParamsColumnFiltersOperatorNotEquals:
		return true
	}
	return false
}
