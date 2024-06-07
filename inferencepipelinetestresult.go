// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomopenlayeraiopenlayergo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/openlayer-ai/openlayer-go/internal/apijson"
	"github.com/openlayer-ai/openlayer-go/internal/apiquery"
	"github.com/openlayer-ai/openlayer-go/internal/param"
	"github.com/openlayer-ai/openlayer-go/internal/requestconfig"
	"github.com/openlayer-ai/openlayer-go/option"
	"github.com/openlayer-ai/openlayer-go/shared"
	"github.com/tidwall/gjson"
)

// InferencePipelineTestResultService contains methods and other services that help
// with interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInferencePipelineTestResultService] method instead.
type InferencePipelineTestResultService struct {
	Options []option.RequestOption
}

// NewInferencePipelineTestResultService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewInferencePipelineTestResultService(opts ...option.RequestOption) (r *InferencePipelineTestResultService) {
	r = &InferencePipelineTestResultService{}
	r.Options = opts
	return
}

// List the test results under an inference pipeline.
func (r *InferencePipelineTestResultService) List(ctx context.Context, id string, query InferencePipelineTestResultListParams, opts ...option.RequestOption) (res *InferencePipelineTestResultListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("inference-pipelines/%s/results", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type InferencePipelineTestResultListResponse struct {
	Meta  InferencePipelineTestResultListResponseMeta   `json:"_meta,required"`
	Items []InferencePipelineTestResultListResponseItem `json:"items,required"`
	JSON  inferencePipelineTestResultListResponseJSON   `json:"-"`
}

// inferencePipelineTestResultListResponseJSON contains the JSON metadata for the
// struct [InferencePipelineTestResultListResponse]
type inferencePipelineTestResultListResponseJSON struct {
	Meta        apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InferencePipelineTestResultListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineTestResultListResponseJSON) RawJSON() string {
	return r.raw
}

type InferencePipelineTestResultListResponseMeta struct {
	// The current page.
	Page int64 `json:"page,required"`
	// The number of items per page.
	PerPage int64 `json:"perPage,required"`
	// The total number of items.
	TotalItems int64 `json:"totalItems,required"`
	// The total number of pages.
	TotalPages int64                                           `json:"totalPages,required"`
	JSON       inferencePipelineTestResultListResponseMetaJSON `json:"-"`
}

// inferencePipelineTestResultListResponseMetaJSON contains the JSON metadata for
// the struct [InferencePipelineTestResultListResponseMeta]
type inferencePipelineTestResultListResponseMetaJSON struct {
	Page        apijson.Field
	PerPage     apijson.Field
	TotalItems  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InferencePipelineTestResultListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineTestResultListResponseMetaJSON) RawJSON() string {
	return r.raw
}

type InferencePipelineTestResultListResponseItem struct {
	// Project version (commit) id.
	ID string `json:"id,required" format:"uuid"`
	// The creation date.
	DateCreated time.Time `json:"dateCreated,required" format:"date-time"`
	// The data end date.
	DateDataEnds time.Time `json:"dateDataEnds,required,nullable" format:"date-time"`
	// The data start date.
	DateDataStarts time.Time `json:"dateDataStarts,required,nullable" format:"date-time"`
	// The last updated date.
	DateUpdated time.Time `json:"dateUpdated,required" format:"date-time"`
	// The inference pipeline id.
	InferencePipelineID string `json:"inferencePipelineId,required,nullable" format:"uuid"`
	// The project version (commit) id.
	ProjectVersionID string `json:"projectVersionId,required,nullable" format:"uuid"`
	// The status of the test.
	Status InferencePipelineTestResultListResponseItemsStatus `json:"status,required"`
	// The status message.
	StatusMessage string                                           `json:"statusMessage,required,nullable"`
	Goal          InferencePipelineTestResultListResponseItemsGoal `json:"goal"`
	// The test id.
	GoalID string                                          `json:"goalId,nullable" format:"uuid"`
	JSON   inferencePipelineTestResultListResponseItemJSON `json:"-"`
}

// inferencePipelineTestResultListResponseItemJSON contains the JSON metadata for
// the struct [InferencePipelineTestResultListResponseItem]
type inferencePipelineTestResultListResponseItemJSON struct {
	ID                  apijson.Field
	DateCreated         apijson.Field
	DateDataEnds        apijson.Field
	DateDataStarts      apijson.Field
	DateUpdated         apijson.Field
	InferencePipelineID apijson.Field
	ProjectVersionID    apijson.Field
	Status              apijson.Field
	StatusMessage       apijson.Field
	Goal                apijson.Field
	GoalID              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InferencePipelineTestResultListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineTestResultListResponseItemJSON) RawJSON() string {
	return r.raw
}

// The status of the test.
type InferencePipelineTestResultListResponseItemsStatus string

const (
	InferencePipelineTestResultListResponseItemsStatusRunning InferencePipelineTestResultListResponseItemsStatus = "running"
	InferencePipelineTestResultListResponseItemsStatusPassing InferencePipelineTestResultListResponseItemsStatus = "passing"
	InferencePipelineTestResultListResponseItemsStatusFailing InferencePipelineTestResultListResponseItemsStatus = "failing"
	InferencePipelineTestResultListResponseItemsStatusSkipped InferencePipelineTestResultListResponseItemsStatus = "skipped"
	InferencePipelineTestResultListResponseItemsStatusError   InferencePipelineTestResultListResponseItemsStatus = "error"
)

func (r InferencePipelineTestResultListResponseItemsStatus) IsKnown() bool {
	switch r {
	case InferencePipelineTestResultListResponseItemsStatusRunning, InferencePipelineTestResultListResponseItemsStatusPassing, InferencePipelineTestResultListResponseItemsStatusFailing, InferencePipelineTestResultListResponseItemsStatusSkipped, InferencePipelineTestResultListResponseItemsStatusError:
		return true
	}
	return false
}

type InferencePipelineTestResultListResponseItemsGoal struct {
	// The test id.
	ID string `json:"id,required" format:"uuid"`
	// The number of comments on the test.
	CommentCount int64 `json:"commentCount,required"`
	// The test creator id.
	CreatorID string `json:"creatorId,required,nullable" format:"uuid"`
	// The date the test was archived.
	DateArchived time.Time `json:"dateArchived,required,nullable" format:"date-time"`
	// The creation date.
	DateCreated time.Time `json:"dateCreated,required" format:"date-time"`
	// The last updated date.
	DateUpdated time.Time `json:"dateUpdated,required" format:"date-time"`
	// The test description.
	Description interface{} `json:"description,required,nullable"`
	// The test name.
	Name string `json:"name,required"`
	// The test number.
	Number int64 `json:"number,required"`
	// The project version (commit) id where the test was created.
	OriginProjectVersionID string `json:"originProjectVersionId,required,nullable" format:"uuid"`
	// The test subtype.
	Subtype string `json:"subtype,required"`
	// Whether the test is suggested or user-created.
	Suggested  bool                                                        `json:"suggested,required"`
	Thresholds []InferencePipelineTestResultListResponseItemsGoalThreshold `json:"thresholds,required"`
	// The test type.
	Type string `json:"type,required"`
	// Whether the test is archived.
	Archived bool `json:"archived"`
	// The delay window in seconds. Only applies to tests that use production data.
	DelayWindow float64 `json:"delayWindow,nullable"`
	// The evaluation window in seconds. Only applies to tests that use production
	// data.
	EvaluationWindow float64 `json:"evaluationWindow,nullable"`
	// Whether the test uses an ML model.
	UsesMlModel bool `json:"usesMlModel"`
	// Whether the test uses production data (monitoring mode only).
	UsesProductionData bool `json:"usesProductionData"`
	// Whether the test uses a reference dataset (monitoring mode only).
	UsesReferenceDataset bool `json:"usesReferenceDataset"`
	// Whether the test uses a training dataset.
	UsesTrainingDataset bool `json:"usesTrainingDataset"`
	// Whether the test uses a validation dataset.
	UsesValidationDataset bool                                                 `json:"usesValidationDataset"`
	JSON                  inferencePipelineTestResultListResponseItemsGoalJSON `json:"-"`
}

// inferencePipelineTestResultListResponseItemsGoalJSON contains the JSON metadata
// for the struct [InferencePipelineTestResultListResponseItemsGoal]
type inferencePipelineTestResultListResponseItemsGoalJSON struct {
	ID                     apijson.Field
	CommentCount           apijson.Field
	CreatorID              apijson.Field
	DateArchived           apijson.Field
	DateCreated            apijson.Field
	DateUpdated            apijson.Field
	Description            apijson.Field
	Name                   apijson.Field
	Number                 apijson.Field
	OriginProjectVersionID apijson.Field
	Subtype                apijson.Field
	Suggested              apijson.Field
	Thresholds             apijson.Field
	Type                   apijson.Field
	Archived               apijson.Field
	DelayWindow            apijson.Field
	EvaluationWindow       apijson.Field
	UsesMlModel            apijson.Field
	UsesProductionData     apijson.Field
	UsesReferenceDataset   apijson.Field
	UsesTrainingDataset    apijson.Field
	UsesValidationDataset  apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *InferencePipelineTestResultListResponseItemsGoal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineTestResultListResponseItemsGoalJSON) RawJSON() string {
	return r.raw
}

type InferencePipelineTestResultListResponseItemsGoalThreshold struct {
	// The insight name to be evaluated.
	InsightName       string        `json:"insightName"`
	InsightParameters []interface{} `json:"insightParameters"`
	// The measurement to be evaluated.
	Measurement string `json:"measurement"`
	// The operator to be used for the evaluation.
	Operator string `json:"operator"`
	// The value to be compared.
	Value InferencePipelineTestResultListResponseItemsGoalThresholdsValueUnion `json:"value"`
	JSON  inferencePipelineTestResultListResponseItemsGoalThresholdJSON        `json:"-"`
}

// inferencePipelineTestResultListResponseItemsGoalThresholdJSON contains the JSON
// metadata for the struct
// [InferencePipelineTestResultListResponseItemsGoalThreshold]
type inferencePipelineTestResultListResponseItemsGoalThresholdJSON struct {
	InsightName       apijson.Field
	InsightParameters apijson.Field
	Measurement       apijson.Field
	Operator          apijson.Field
	Value             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InferencePipelineTestResultListResponseItemsGoalThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineTestResultListResponseItemsGoalThresholdJSON) RawJSON() string {
	return r.raw
}

// The value to be compared.
//
// Union satisfied by [shared.UnionFloat], [shared.UnionBool], [shared.UnionString]
// or [InferencePipelineTestResultListResponseItemsGoalThresholdsValueArray].
type InferencePipelineTestResultListResponseItemsGoalThresholdsValueUnion interface {
	ImplementsInferencePipelineTestResultListResponseItemsGoalThresholdsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InferencePipelineTestResultListResponseItemsGoalThresholdsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InferencePipelineTestResultListResponseItemsGoalThresholdsValueArray{}),
		},
	)
}

type InferencePipelineTestResultListResponseItemsGoalThresholdsValueArray []string

func (r InferencePipelineTestResultListResponseItemsGoalThresholdsValueArray) ImplementsInferencePipelineTestResultListResponseItemsGoalThresholdsValueUnion() {
}

type InferencePipelineTestResultListParams struct {
	// Include archived goals.
	IncludeArchived param.Field[bool] `query:"includeArchived"`
	// The page to return in a paginated query.
	Page param.Field[int64] `query:"page"`
	// Maximum number of items to return per page.
	PerPage param.Field[int64] `query:"perPage"`
	// Filter list of test results by status. Available statuses are `running`,
	// `passing`, `failing`, `skipped`, and `error`.
	Status param.Field[InferencePipelineTestResultListParamsStatus] `query:"status"`
	// Filter objects by test type. Available types are `integrity`, `consistency`,
	// `performance`, `fairness`, and `robustness`.
	Type param.Field[InferencePipelineTestResultListParamsType] `query:"type"`
}

// URLQuery serializes [InferencePipelineTestResultListParams]'s query parameters
// as `url.Values`.
func (r InferencePipelineTestResultListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter list of test results by status. Available statuses are `running`,
// `passing`, `failing`, `skipped`, and `error`.
type InferencePipelineTestResultListParamsStatus string

const (
	InferencePipelineTestResultListParamsStatusRunning InferencePipelineTestResultListParamsStatus = "running"
	InferencePipelineTestResultListParamsStatusPassing InferencePipelineTestResultListParamsStatus = "passing"
	InferencePipelineTestResultListParamsStatusFailing InferencePipelineTestResultListParamsStatus = "failing"
	InferencePipelineTestResultListParamsStatusSkipped InferencePipelineTestResultListParamsStatus = "skipped"
	InferencePipelineTestResultListParamsStatusError   InferencePipelineTestResultListParamsStatus = "error"
)

func (r InferencePipelineTestResultListParamsStatus) IsKnown() bool {
	switch r {
	case InferencePipelineTestResultListParamsStatusRunning, InferencePipelineTestResultListParamsStatusPassing, InferencePipelineTestResultListParamsStatusFailing, InferencePipelineTestResultListParamsStatusSkipped, InferencePipelineTestResultListParamsStatusError:
		return true
	}
	return false
}

// Filter objects by test type. Available types are `integrity`, `consistency`,
// `performance`, `fairness`, and `robustness`.
type InferencePipelineTestResultListParamsType string

const (
	InferencePipelineTestResultListParamsTypeIntegrity   InferencePipelineTestResultListParamsType = "integrity"
	InferencePipelineTestResultListParamsTypeConsistency InferencePipelineTestResultListParamsType = "consistency"
	InferencePipelineTestResultListParamsTypePerformance InferencePipelineTestResultListParamsType = "performance"
	InferencePipelineTestResultListParamsTypeFairness    InferencePipelineTestResultListParamsType = "fairness"
	InferencePipelineTestResultListParamsTypeRobustness  InferencePipelineTestResultListParamsType = "robustness"
)

func (r InferencePipelineTestResultListParamsType) IsKnown() bool {
	switch r {
	case InferencePipelineTestResultListParamsTypeIntegrity, InferencePipelineTestResultListParamsTypeConsistency, InferencePipelineTestResultListParamsTypePerformance, InferencePipelineTestResultListParamsTypeFairness, InferencePipelineTestResultListParamsTypeRobustness:
		return true
	}
	return false
}
