// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer

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

// CommitTestResultService contains methods and other services that help with
// interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCommitTestResultService] method instead.
type CommitTestResultService struct {
	Options []option.RequestOption
}

// NewCommitTestResultService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCommitTestResultService(opts ...option.RequestOption) (r *CommitTestResultService) {
	r = &CommitTestResultService{}
	r.Options = opts
	return
}

// List the test results for a project commit (project version).
func (r *CommitTestResultService) List(ctx context.Context, projectVersionID string, query CommitTestResultListParams, opts ...option.RequestOption) (res *CommitTestResultListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if projectVersionID == "" {
		err = errors.New("missing required projectVersionId parameter")
		return
	}
	path := fmt.Sprintf("versions/%s/results", projectVersionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CommitTestResultListResponse struct {
	Items []CommitTestResultListResponseItem `json:"items,required"`
	JSON  commitTestResultListResponseJSON   `json:"-"`
}

// commitTestResultListResponseJSON contains the JSON metadata for the struct
// [CommitTestResultListResponse]
type commitTestResultListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitTestResultListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitTestResultListResponseJSON) RawJSON() string {
	return r.raw
}

type CommitTestResultListResponseItem struct {
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
	Status CommitTestResultListResponseItemsStatus `json:"status,required"`
	// The status message.
	StatusMessage string                                `json:"statusMessage,required,nullable"`
	Goal          CommitTestResultListResponseItemsGoal `json:"goal"`
	// The test id.
	GoalID string                               `json:"goalId,nullable" format:"uuid"`
	JSON   commitTestResultListResponseItemJSON `json:"-"`
}

// commitTestResultListResponseItemJSON contains the JSON metadata for the struct
// [CommitTestResultListResponseItem]
type commitTestResultListResponseItemJSON struct {
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

func (r *CommitTestResultListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitTestResultListResponseItemJSON) RawJSON() string {
	return r.raw
}

// The status of the test.
type CommitTestResultListResponseItemsStatus string

const (
	CommitTestResultListResponseItemsStatusRunning CommitTestResultListResponseItemsStatus = "running"
	CommitTestResultListResponseItemsStatusPassing CommitTestResultListResponseItemsStatus = "passing"
	CommitTestResultListResponseItemsStatusFailing CommitTestResultListResponseItemsStatus = "failing"
	CommitTestResultListResponseItemsStatusSkipped CommitTestResultListResponseItemsStatus = "skipped"
	CommitTestResultListResponseItemsStatusError   CommitTestResultListResponseItemsStatus = "error"
)

func (r CommitTestResultListResponseItemsStatus) IsKnown() bool {
	switch r {
	case CommitTestResultListResponseItemsStatusRunning, CommitTestResultListResponseItemsStatusPassing, CommitTestResultListResponseItemsStatusFailing, CommitTestResultListResponseItemsStatusSkipped, CommitTestResultListResponseItemsStatusError:
		return true
	}
	return false
}

type CommitTestResultListResponseItemsGoal struct {
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
	Suggested  bool                                             `json:"suggested,required"`
	Thresholds []CommitTestResultListResponseItemsGoalThreshold `json:"thresholds,required"`
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
	UsesValidationDataset bool                                      `json:"usesValidationDataset"`
	JSON                  commitTestResultListResponseItemsGoalJSON `json:"-"`
}

// commitTestResultListResponseItemsGoalJSON contains the JSON metadata for the
// struct [CommitTestResultListResponseItemsGoal]
type commitTestResultListResponseItemsGoalJSON struct {
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

func (r *CommitTestResultListResponseItemsGoal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitTestResultListResponseItemsGoalJSON) RawJSON() string {
	return r.raw
}

type CommitTestResultListResponseItemsGoalThreshold struct {
	// The insight name to be evaluated.
	InsightName       string        `json:"insightName"`
	InsightParameters []interface{} `json:"insightParameters"`
	// The measurement to be evaluated.
	Measurement string `json:"measurement"`
	// The operator to be used for the evaluation.
	Operator string `json:"operator"`
	// The value to be compared.
	Value CommitTestResultListResponseItemsGoalThresholdsValueUnion `json:"value"`
	JSON  commitTestResultListResponseItemsGoalThresholdJSON        `json:"-"`
}

// commitTestResultListResponseItemsGoalThresholdJSON contains the JSON metadata
// for the struct [CommitTestResultListResponseItemsGoalThreshold]
type commitTestResultListResponseItemsGoalThresholdJSON struct {
	InsightName       apijson.Field
	InsightParameters apijson.Field
	Measurement       apijson.Field
	Operator          apijson.Field
	Value             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CommitTestResultListResponseItemsGoalThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitTestResultListResponseItemsGoalThresholdJSON) RawJSON() string {
	return r.raw
}

// The value to be compared.
//
// Union satisfied by [shared.UnionFloat], [shared.UnionBool], [shared.UnionString]
// or [CommitTestResultListResponseItemsGoalThresholdsValueArray].
type CommitTestResultListResponseItemsGoalThresholdsValueUnion interface {
	ImplementsCommitTestResultListResponseItemsGoalThresholdsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CommitTestResultListResponseItemsGoalThresholdsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(CommitTestResultListResponseItemsGoalThresholdsValueArray{}),
		},
	)
}

type CommitTestResultListResponseItemsGoalThresholdsValueArray []string

func (r CommitTestResultListResponseItemsGoalThresholdsValueArray) ImplementsCommitTestResultListResponseItemsGoalThresholdsValueUnion() {
}

type CommitTestResultListParams struct {
	// Include archived goals.
	IncludeArchived param.Field[bool] `query:"includeArchived"`
	// The page to return in a paginated query.
	Page param.Field[int64] `query:"page"`
	// Maximum number of items to return per page.
	PerPage param.Field[int64] `query:"perPage"`
	// Filter list of test results by status. Available statuses are `running`,
	// `passing`, `failing`, `skipped`, and `error`.
	Status param.Field[CommitTestResultListParamsStatus] `query:"status"`
	// Filter objects by test type. Available types are `integrity`, `consistency`,
	// `performance`, `fairness`, and `robustness`.
	Type param.Field[CommitTestResultListParamsType] `query:"type"`
}

// URLQuery serializes [CommitTestResultListParams]'s query parameters as
// `url.Values`.
func (r CommitTestResultListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter list of test results by status. Available statuses are `running`,
// `passing`, `failing`, `skipped`, and `error`.
type CommitTestResultListParamsStatus string

const (
	CommitTestResultListParamsStatusRunning CommitTestResultListParamsStatus = "running"
	CommitTestResultListParamsStatusPassing CommitTestResultListParamsStatus = "passing"
	CommitTestResultListParamsStatusFailing CommitTestResultListParamsStatus = "failing"
	CommitTestResultListParamsStatusSkipped CommitTestResultListParamsStatus = "skipped"
	CommitTestResultListParamsStatusError   CommitTestResultListParamsStatus = "error"
)

func (r CommitTestResultListParamsStatus) IsKnown() bool {
	switch r {
	case CommitTestResultListParamsStatusRunning, CommitTestResultListParamsStatusPassing, CommitTestResultListParamsStatusFailing, CommitTestResultListParamsStatusSkipped, CommitTestResultListParamsStatusError:
		return true
	}
	return false
}

// Filter objects by test type. Available types are `integrity`, `consistency`,
// `performance`, `fairness`, and `robustness`.
type CommitTestResultListParamsType string

const (
	CommitTestResultListParamsTypeIntegrity   CommitTestResultListParamsType = "integrity"
	CommitTestResultListParamsTypeConsistency CommitTestResultListParamsType = "consistency"
	CommitTestResultListParamsTypePerformance CommitTestResultListParamsType = "performance"
	CommitTestResultListParamsTypeFairness    CommitTestResultListParamsType = "fairness"
	CommitTestResultListParamsTypeRobustness  CommitTestResultListParamsType = "robustness"
)

func (r CommitTestResultListParamsType) IsKnown() bool {
	switch r {
	case CommitTestResultListParamsTypeIntegrity, CommitTestResultListParamsTypeConsistency, CommitTestResultListParamsTypePerformance, CommitTestResultListParamsTypeFairness, CommitTestResultListParamsTypeRobustness:
		return true
	}
	return false
}
