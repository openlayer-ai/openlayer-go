// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/openlayer-ai/openlayer-go/internal/apijson"
	"github.com/openlayer-ai/openlayer-go/internal/param"
	"github.com/openlayer-ai/openlayer-go/internal/requestconfig"
	"github.com/openlayer-ai/openlayer-go/option"
	"github.com/openlayer-ai/openlayer-go/shared"
	"github.com/tidwall/gjson"
)

// ProjectTestService contains methods and other services that help with
// interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectTestService] method instead.
type ProjectTestService struct {
	Options []option.RequestOption
}

// NewProjectTestService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProjectTestService(opts ...option.RequestOption) (r *ProjectTestService) {
	r = &ProjectTestService{}
	r.Options = opts
	return
}

// Create a test.
func (r *ProjectTestService) New(ctx context.Context, projectID string, body ProjectTestNewParams, opts ...option.RequestOption) (res *ProjectTestNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return
	}
	path := fmt.Sprintf("projects/%s/tests", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ProjectTestNewResponse struct {
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
	Suggested  bool                              `json:"suggested,required"`
	Thresholds []ProjectTestNewResponseThreshold `json:"thresholds,required"`
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
	UsesValidationDataset bool                       `json:"usesValidationDataset"`
	JSON                  projectTestNewResponseJSON `json:"-"`
}

// projectTestNewResponseJSON contains the JSON metadata for the struct
// [ProjectTestNewResponse]
type projectTestNewResponseJSON struct {
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

func (r *ProjectTestNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectTestNewResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectTestNewResponseThreshold struct {
	// The insight name to be evaluated.
	InsightName string `json:"insightName"`
	// The insight parameters. Required only for some test subtypes.
	InsightParameters []ProjectTestNewResponseThresholdsInsightParameter `json:"insightParameters,nullable"`
	// The measurement to be evaluated.
	Measurement string `json:"measurement"`
	// The operator to be used for the evaluation.
	Operator ProjectTestNewResponseThresholdsOperator `json:"operator"`
	// Whether to use automatic anomaly detection or manual thresholds
	ThresholdMode ProjectTestNewResponseThresholdsThresholdMode `json:"thresholdMode"`
	// The value to be compared.
	Value ProjectTestNewResponseThresholdsValueUnion `json:"value"`
	JSON  projectTestNewResponseThresholdJSON        `json:"-"`
}

// projectTestNewResponseThresholdJSON contains the JSON metadata for the struct
// [ProjectTestNewResponseThreshold]
type projectTestNewResponseThresholdJSON struct {
	InsightName       apijson.Field
	InsightParameters apijson.Field
	Measurement       apijson.Field
	Operator          apijson.Field
	ThresholdMode     apijson.Field
	Value             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectTestNewResponseThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectTestNewResponseThresholdJSON) RawJSON() string {
	return r.raw
}

type ProjectTestNewResponseThresholdsInsightParameter struct {
	// The name of the insight filter.
	Name  string                                               `json:"name,required"`
	Value interface{}                                          `json:"value,required"`
	JSON  projectTestNewResponseThresholdsInsightParameterJSON `json:"-"`
}

// projectTestNewResponseThresholdsInsightParameterJSON contains the JSON metadata
// for the struct [ProjectTestNewResponseThresholdsInsightParameter]
type projectTestNewResponseThresholdsInsightParameterJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectTestNewResponseThresholdsInsightParameter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectTestNewResponseThresholdsInsightParameterJSON) RawJSON() string {
	return r.raw
}

// The operator to be used for the evaluation.
type ProjectTestNewResponseThresholdsOperator string

const (
	ProjectTestNewResponseThresholdsOperatorIs              ProjectTestNewResponseThresholdsOperator = "is"
	ProjectTestNewResponseThresholdsOperatorGreater         ProjectTestNewResponseThresholdsOperator = ">"
	ProjectTestNewResponseThresholdsOperatorGreaterOrEquals ProjectTestNewResponseThresholdsOperator = ">="
	ProjectTestNewResponseThresholdsOperatorLess            ProjectTestNewResponseThresholdsOperator = "<"
	ProjectTestNewResponseThresholdsOperatorLessOrEquals    ProjectTestNewResponseThresholdsOperator = "<="
	ProjectTestNewResponseThresholdsOperatorNotEquals       ProjectTestNewResponseThresholdsOperator = "!="
)

func (r ProjectTestNewResponseThresholdsOperator) IsKnown() bool {
	switch r {
	case ProjectTestNewResponseThresholdsOperatorIs, ProjectTestNewResponseThresholdsOperatorGreater, ProjectTestNewResponseThresholdsOperatorGreaterOrEquals, ProjectTestNewResponseThresholdsOperatorLess, ProjectTestNewResponseThresholdsOperatorLessOrEquals, ProjectTestNewResponseThresholdsOperatorNotEquals:
		return true
	}
	return false
}

// Whether to use automatic anomaly detection or manual thresholds
type ProjectTestNewResponseThresholdsThresholdMode string

const (
	ProjectTestNewResponseThresholdsThresholdModeAutomatic ProjectTestNewResponseThresholdsThresholdMode = "automatic"
	ProjectTestNewResponseThresholdsThresholdModeManual    ProjectTestNewResponseThresholdsThresholdMode = "manual"
)

func (r ProjectTestNewResponseThresholdsThresholdMode) IsKnown() bool {
	switch r {
	case ProjectTestNewResponseThresholdsThresholdModeAutomatic, ProjectTestNewResponseThresholdsThresholdModeManual:
		return true
	}
	return false
}

// The value to be compared.
//
// Union satisfied by [shared.UnionFloat], [shared.UnionBool], [shared.UnionString]
// or [ProjectTestNewResponseThresholdsValueArray].
type ProjectTestNewResponseThresholdsValueUnion interface {
	ImplementsProjectTestNewResponseThresholdsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectTestNewResponseThresholdsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(ProjectTestNewResponseThresholdsValueArray{}),
		},
	)
}

type ProjectTestNewResponseThresholdsValueArray []string

func (r ProjectTestNewResponseThresholdsValueArray) ImplementsProjectTestNewResponseThresholdsValueUnion() {
}

type ProjectTestNewParams struct {
	// The test description.
	Description param.Field[interface{}] `json:"description,required"`
	// The test name.
	Name param.Field[string] `json:"name,required"`
	// The test subtype.
	Subtype    param.Field[string]                          `json:"subtype,required"`
	Thresholds param.Field[[]ProjectTestNewParamsThreshold] `json:"thresholds,required"`
	// The test type.
	Type param.Field[string] `json:"type,required"`
	// Whether the test is archived.
	Archived param.Field[bool] `json:"archived"`
	// The delay window in seconds. Only applies to tests that use production data.
	DelayWindow param.Field[float64] `json:"delayWindow"`
	// The evaluation window in seconds. Only applies to tests that use production
	// data.
	EvaluationWindow param.Field[float64] `json:"evaluationWindow"`
	// Whether the test uses an ML model.
	UsesMlModel param.Field[bool] `json:"usesMlModel"`
	// Whether the test uses production data (monitoring mode only).
	UsesProductionData param.Field[bool] `json:"usesProductionData"`
	// Whether the test uses a reference dataset (monitoring mode only).
	UsesReferenceDataset param.Field[bool] `json:"usesReferenceDataset"`
	// Whether the test uses a training dataset.
	UsesTrainingDataset param.Field[bool] `json:"usesTrainingDataset"`
	// Whether the test uses a validation dataset.
	UsesValidationDataset param.Field[bool] `json:"usesValidationDataset"`
}

func (r ProjectTestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectTestNewParamsThreshold struct {
	// The insight name to be evaluated.
	InsightName param.Field[string] `json:"insightName"`
	// The insight parameters. Required only for some test subtypes.
	InsightParameters param.Field[[]ProjectTestNewParamsThresholdsInsightParameter] `json:"insightParameters"`
	// The measurement to be evaluated.
	Measurement param.Field[string] `json:"measurement"`
	// The operator to be used for the evaluation.
	Operator param.Field[ProjectTestNewParamsThresholdsOperator] `json:"operator"`
	// Whether to use automatic anomaly detection or manual thresholds
	ThresholdMode param.Field[ProjectTestNewParamsThresholdsThresholdMode] `json:"thresholdMode"`
	// The value to be compared.
	Value param.Field[ProjectTestNewParamsThresholdsValueUnion] `json:"value"`
}

func (r ProjectTestNewParamsThreshold) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectTestNewParamsThresholdsInsightParameter struct {
	// The name of the insight filter.
	Name  param.Field[string]      `json:"name,required"`
	Value param.Field[interface{}] `json:"value,required"`
}

func (r ProjectTestNewParamsThresholdsInsightParameter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The operator to be used for the evaluation.
type ProjectTestNewParamsThresholdsOperator string

const (
	ProjectTestNewParamsThresholdsOperatorIs              ProjectTestNewParamsThresholdsOperator = "is"
	ProjectTestNewParamsThresholdsOperatorGreater         ProjectTestNewParamsThresholdsOperator = ">"
	ProjectTestNewParamsThresholdsOperatorGreaterOrEquals ProjectTestNewParamsThresholdsOperator = ">="
	ProjectTestNewParamsThresholdsOperatorLess            ProjectTestNewParamsThresholdsOperator = "<"
	ProjectTestNewParamsThresholdsOperatorLessOrEquals    ProjectTestNewParamsThresholdsOperator = "<="
	ProjectTestNewParamsThresholdsOperatorNotEquals       ProjectTestNewParamsThresholdsOperator = "!="
)

func (r ProjectTestNewParamsThresholdsOperator) IsKnown() bool {
	switch r {
	case ProjectTestNewParamsThresholdsOperatorIs, ProjectTestNewParamsThresholdsOperatorGreater, ProjectTestNewParamsThresholdsOperatorGreaterOrEquals, ProjectTestNewParamsThresholdsOperatorLess, ProjectTestNewParamsThresholdsOperatorLessOrEquals, ProjectTestNewParamsThresholdsOperatorNotEquals:
		return true
	}
	return false
}

// Whether to use automatic anomaly detection or manual thresholds
type ProjectTestNewParamsThresholdsThresholdMode string

const (
	ProjectTestNewParamsThresholdsThresholdModeAutomatic ProjectTestNewParamsThresholdsThresholdMode = "automatic"
	ProjectTestNewParamsThresholdsThresholdModeManual    ProjectTestNewParamsThresholdsThresholdMode = "manual"
)

func (r ProjectTestNewParamsThresholdsThresholdMode) IsKnown() bool {
	switch r {
	case ProjectTestNewParamsThresholdsThresholdModeAutomatic, ProjectTestNewParamsThresholdsThresholdModeManual:
		return true
	}
	return false
}

// The value to be compared.
//
// Satisfied by [shared.UnionFloat], [shared.UnionBool], [shared.UnionString],
// [ProjectTestNewParamsThresholdsValueArray].
type ProjectTestNewParamsThresholdsValueUnion interface {
	ImplementsProjectTestNewParamsThresholdsValueUnion()
}

type ProjectTestNewParamsThresholdsValueArray []string

func (r ProjectTestNewParamsThresholdsValueArray) ImplementsProjectTestNewParamsThresholdsValueUnion() {
}
