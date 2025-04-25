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

// List tests under a project.
func (r *ProjectTestService) List(ctx context.Context, projectID string, query ProjectTestListParams, opts ...option.RequestOption) (res *ProjectTestListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return
	}
	path := fmt.Sprintf("projects/%s/tests", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
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
	Subtype ProjectTestNewResponseSubtype `json:"subtype,required"`
	// Whether the test is suggested or user-created.
	Suggested  bool                              `json:"suggested,required"`
	Thresholds []ProjectTestNewResponseThreshold `json:"thresholds,required"`
	// The test type.
	Type ProjectTestNewResponseType `json:"type,required"`
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

// The test subtype.
type ProjectTestNewResponseSubtype string

const (
	ProjectTestNewResponseSubtypeAnomalousColumnCount       ProjectTestNewResponseSubtype = "anomalousColumnCount"
	ProjectTestNewResponseSubtypeCharacterLength            ProjectTestNewResponseSubtype = "characterLength"
	ProjectTestNewResponseSubtypeClassImbalanceRatio        ProjectTestNewResponseSubtype = "classImbalanceRatio"
	ProjectTestNewResponseSubtypeExpectColumnAToBeInColumnB ProjectTestNewResponseSubtype = "expectColumnAToBeInColumnB"
	ProjectTestNewResponseSubtypeColumnAverage              ProjectTestNewResponseSubtype = "columnAverage"
	ProjectTestNewResponseSubtypeColumnDrift                ProjectTestNewResponseSubtype = "columnDrift"
	ProjectTestNewResponseSubtypeColumnStatistic            ProjectTestNewResponseSubtype = "columnStatistic"
	ProjectTestNewResponseSubtypeColumnValuesMatch          ProjectTestNewResponseSubtype = "columnValuesMatch"
	ProjectTestNewResponseSubtypeConflictingLabelRowCount   ProjectTestNewResponseSubtype = "conflictingLabelRowCount"
	ProjectTestNewResponseSubtypeContainsPii                ProjectTestNewResponseSubtype = "containsPii"
	ProjectTestNewResponseSubtypeContainsValidURL           ProjectTestNewResponseSubtype = "containsValidUrl"
	ProjectTestNewResponseSubtypeCorrelatedFeatureCount     ProjectTestNewResponseSubtype = "correlatedFeatureCount"
	ProjectTestNewResponseSubtypeCustomMetricThreshold      ProjectTestNewResponseSubtype = "customMetricThreshold"
	ProjectTestNewResponseSubtypeDuplicateRowCount          ProjectTestNewResponseSubtype = "duplicateRowCount"
	ProjectTestNewResponseSubtypeEmptyFeature               ProjectTestNewResponseSubtype = "emptyFeature"
	ProjectTestNewResponseSubtypeEmptyFeatureCount          ProjectTestNewResponseSubtype = "emptyFeatureCount"
	ProjectTestNewResponseSubtypeDriftedFeatureCount        ProjectTestNewResponseSubtype = "driftedFeatureCount"
	ProjectTestNewResponseSubtypeFeatureMissingValues       ProjectTestNewResponseSubtype = "featureMissingValues"
	ProjectTestNewResponseSubtypeFeatureValueValidation     ProjectTestNewResponseSubtype = "featureValueValidation"
	ProjectTestNewResponseSubtypeGreatExpectations          ProjectTestNewResponseSubtype = "greatExpectations"
	ProjectTestNewResponseSubtypeGroupByColumnStatsCheck    ProjectTestNewResponseSubtype = "groupByColumnStatsCheck"
	ProjectTestNewResponseSubtypeIllFormedRowCount          ProjectTestNewResponseSubtype = "illFormedRowCount"
	ProjectTestNewResponseSubtypeIsCode                     ProjectTestNewResponseSubtype = "isCode"
	ProjectTestNewResponseSubtypeIsJson                     ProjectTestNewResponseSubtype = "isJson"
	ProjectTestNewResponseSubtypeLlmRubricThresholdV2       ProjectTestNewResponseSubtype = "llmRubricThresholdV2"
	ProjectTestNewResponseSubtypeLabelDrift                 ProjectTestNewResponseSubtype = "labelDrift"
	ProjectTestNewResponseSubtypeMetricThreshold            ProjectTestNewResponseSubtype = "metricThreshold"
	ProjectTestNewResponseSubtypeNewCategoryCount           ProjectTestNewResponseSubtype = "newCategoryCount"
	ProjectTestNewResponseSubtypeNewLabelCount              ProjectTestNewResponseSubtype = "newLabelCount"
	ProjectTestNewResponseSubtypeNullRowCount               ProjectTestNewResponseSubtype = "nullRowCount"
	ProjectTestNewResponseSubtypeRowCount                   ProjectTestNewResponseSubtype = "rowCount"
	ProjectTestNewResponseSubtypePpScoreValueValidation     ProjectTestNewResponseSubtype = "ppScoreValueValidation"
	ProjectTestNewResponseSubtypeQuasiConstantFeature       ProjectTestNewResponseSubtype = "quasiConstantFeature"
	ProjectTestNewResponseSubtypeQuasiConstantFeatureCount  ProjectTestNewResponseSubtype = "quasiConstantFeatureCount"
	ProjectTestNewResponseSubtypeSqlQuery                   ProjectTestNewResponseSubtype = "sqlQuery"
	ProjectTestNewResponseSubtypeDtypeValidation            ProjectTestNewResponseSubtype = "dtypeValidation"
	ProjectTestNewResponseSubtypeSentenceLength             ProjectTestNewResponseSubtype = "sentenceLength"
	ProjectTestNewResponseSubtypeSizeRatio                  ProjectTestNewResponseSubtype = "sizeRatio"
	ProjectTestNewResponseSubtypeSpecialCharactersRatio     ProjectTestNewResponseSubtype = "specialCharactersRatio"
	ProjectTestNewResponseSubtypeStringValidation           ProjectTestNewResponseSubtype = "stringValidation"
	ProjectTestNewResponseSubtypeTrainValLeakageRowCount    ProjectTestNewResponseSubtype = "trainValLeakageRowCount"
)

func (r ProjectTestNewResponseSubtype) IsKnown() bool {
	switch r {
	case ProjectTestNewResponseSubtypeAnomalousColumnCount, ProjectTestNewResponseSubtypeCharacterLength, ProjectTestNewResponseSubtypeClassImbalanceRatio, ProjectTestNewResponseSubtypeExpectColumnAToBeInColumnB, ProjectTestNewResponseSubtypeColumnAverage, ProjectTestNewResponseSubtypeColumnDrift, ProjectTestNewResponseSubtypeColumnStatistic, ProjectTestNewResponseSubtypeColumnValuesMatch, ProjectTestNewResponseSubtypeConflictingLabelRowCount, ProjectTestNewResponseSubtypeContainsPii, ProjectTestNewResponseSubtypeContainsValidURL, ProjectTestNewResponseSubtypeCorrelatedFeatureCount, ProjectTestNewResponseSubtypeCustomMetricThreshold, ProjectTestNewResponseSubtypeDuplicateRowCount, ProjectTestNewResponseSubtypeEmptyFeature, ProjectTestNewResponseSubtypeEmptyFeatureCount, ProjectTestNewResponseSubtypeDriftedFeatureCount, ProjectTestNewResponseSubtypeFeatureMissingValues, ProjectTestNewResponseSubtypeFeatureValueValidation, ProjectTestNewResponseSubtypeGreatExpectations, ProjectTestNewResponseSubtypeGroupByColumnStatsCheck, ProjectTestNewResponseSubtypeIllFormedRowCount, ProjectTestNewResponseSubtypeIsCode, ProjectTestNewResponseSubtypeIsJson, ProjectTestNewResponseSubtypeLlmRubricThresholdV2, ProjectTestNewResponseSubtypeLabelDrift, ProjectTestNewResponseSubtypeMetricThreshold, ProjectTestNewResponseSubtypeNewCategoryCount, ProjectTestNewResponseSubtypeNewLabelCount, ProjectTestNewResponseSubtypeNullRowCount, ProjectTestNewResponseSubtypeRowCount, ProjectTestNewResponseSubtypePpScoreValueValidation, ProjectTestNewResponseSubtypeQuasiConstantFeature, ProjectTestNewResponseSubtypeQuasiConstantFeatureCount, ProjectTestNewResponseSubtypeSqlQuery, ProjectTestNewResponseSubtypeDtypeValidation, ProjectTestNewResponseSubtypeSentenceLength, ProjectTestNewResponseSubtypeSizeRatio, ProjectTestNewResponseSubtypeSpecialCharactersRatio, ProjectTestNewResponseSubtypeStringValidation, ProjectTestNewResponseSubtypeTrainValLeakageRowCount:
		return true
	}
	return false
}

type ProjectTestNewResponseThreshold struct {
	// The insight name to be evaluated.
	InsightName ProjectTestNewResponseThresholdsInsightName `json:"insightName"`
	// The insight parameters. Required only for some test subtypes. For example, for
	// tests that require a column name, the insight parameters will be [{'name':
	// 'column_name', 'value': 'Age'}]
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

// The insight name to be evaluated.
type ProjectTestNewResponseThresholdsInsightName string

const (
	ProjectTestNewResponseThresholdsInsightNameCharacterLength            ProjectTestNewResponseThresholdsInsightName = "characterLength"
	ProjectTestNewResponseThresholdsInsightNameClassImbalance             ProjectTestNewResponseThresholdsInsightName = "classImbalance"
	ProjectTestNewResponseThresholdsInsightNameExpectColumnAToBeInColumnB ProjectTestNewResponseThresholdsInsightName = "expectColumnAToBeInColumnB"
	ProjectTestNewResponseThresholdsInsightNameColumnAverage              ProjectTestNewResponseThresholdsInsightName = "columnAverage"
	ProjectTestNewResponseThresholdsInsightNameColumnDrift                ProjectTestNewResponseThresholdsInsightName = "columnDrift"
	ProjectTestNewResponseThresholdsInsightNameColumnValuesMatch          ProjectTestNewResponseThresholdsInsightName = "columnValuesMatch"
	ProjectTestNewResponseThresholdsInsightNameConfidenceDistribution     ProjectTestNewResponseThresholdsInsightName = "confidenceDistribution"
	ProjectTestNewResponseThresholdsInsightNameConflictingLabelRowCount   ProjectTestNewResponseThresholdsInsightName = "conflictingLabelRowCount"
	ProjectTestNewResponseThresholdsInsightNameContainsPii                ProjectTestNewResponseThresholdsInsightName = "containsPii"
	ProjectTestNewResponseThresholdsInsightNameContainsValidURL           ProjectTestNewResponseThresholdsInsightName = "containsValidUrl"
	ProjectTestNewResponseThresholdsInsightNameCorrelatedFeatures         ProjectTestNewResponseThresholdsInsightName = "correlatedFeatures"
	ProjectTestNewResponseThresholdsInsightNameCustomMetric               ProjectTestNewResponseThresholdsInsightName = "customMetric"
	ProjectTestNewResponseThresholdsInsightNameDuplicateRowCount          ProjectTestNewResponseThresholdsInsightName = "duplicateRowCount"
	ProjectTestNewResponseThresholdsInsightNameEmptyFeatures              ProjectTestNewResponseThresholdsInsightName = "emptyFeatures"
	ProjectTestNewResponseThresholdsInsightNameFeatureDrift               ProjectTestNewResponseThresholdsInsightName = "featureDrift"
	ProjectTestNewResponseThresholdsInsightNameFeatureProfile             ProjectTestNewResponseThresholdsInsightName = "featureProfile"
	ProjectTestNewResponseThresholdsInsightNameGreatExpectations          ProjectTestNewResponseThresholdsInsightName = "greatExpectations"
	ProjectTestNewResponseThresholdsInsightNameGroupByColumnStatsCheck    ProjectTestNewResponseThresholdsInsightName = "groupByColumnStatsCheck"
	ProjectTestNewResponseThresholdsInsightNameIllFormedRowCount          ProjectTestNewResponseThresholdsInsightName = "illFormedRowCount"
	ProjectTestNewResponseThresholdsInsightNameIsCode                     ProjectTestNewResponseThresholdsInsightName = "isCode"
	ProjectTestNewResponseThresholdsInsightNameIsJson                     ProjectTestNewResponseThresholdsInsightName = "isJson"
	ProjectTestNewResponseThresholdsInsightNameLlmRubricV2                ProjectTestNewResponseThresholdsInsightName = "llmRubricV2"
	ProjectTestNewResponseThresholdsInsightNameLabelDrift                 ProjectTestNewResponseThresholdsInsightName = "labelDrift"
	ProjectTestNewResponseThresholdsInsightNameMetrics                    ProjectTestNewResponseThresholdsInsightName = "metrics"
	ProjectTestNewResponseThresholdsInsightNameNewCategories              ProjectTestNewResponseThresholdsInsightName = "newCategories"
	ProjectTestNewResponseThresholdsInsightNameNewLabels                  ProjectTestNewResponseThresholdsInsightName = "newLabels"
	ProjectTestNewResponseThresholdsInsightNameNullRowCount               ProjectTestNewResponseThresholdsInsightName = "nullRowCount"
	ProjectTestNewResponseThresholdsInsightNamePpScore                    ProjectTestNewResponseThresholdsInsightName = "ppScore"
	ProjectTestNewResponseThresholdsInsightNameQuasiConstantFeatures      ProjectTestNewResponseThresholdsInsightName = "quasiConstantFeatures"
	ProjectTestNewResponseThresholdsInsightNameSentenceLength             ProjectTestNewResponseThresholdsInsightName = "sentenceLength"
	ProjectTestNewResponseThresholdsInsightNameSizeRatio                  ProjectTestNewResponseThresholdsInsightName = "sizeRatio"
	ProjectTestNewResponseThresholdsInsightNameSpecialCharacters          ProjectTestNewResponseThresholdsInsightName = "specialCharacters"
	ProjectTestNewResponseThresholdsInsightNameStringValidation           ProjectTestNewResponseThresholdsInsightName = "stringValidation"
	ProjectTestNewResponseThresholdsInsightNameTrainValLeakageRowCount    ProjectTestNewResponseThresholdsInsightName = "trainValLeakageRowCount"
)

func (r ProjectTestNewResponseThresholdsInsightName) IsKnown() bool {
	switch r {
	case ProjectTestNewResponseThresholdsInsightNameCharacterLength, ProjectTestNewResponseThresholdsInsightNameClassImbalance, ProjectTestNewResponseThresholdsInsightNameExpectColumnAToBeInColumnB, ProjectTestNewResponseThresholdsInsightNameColumnAverage, ProjectTestNewResponseThresholdsInsightNameColumnDrift, ProjectTestNewResponseThresholdsInsightNameColumnValuesMatch, ProjectTestNewResponseThresholdsInsightNameConfidenceDistribution, ProjectTestNewResponseThresholdsInsightNameConflictingLabelRowCount, ProjectTestNewResponseThresholdsInsightNameContainsPii, ProjectTestNewResponseThresholdsInsightNameContainsValidURL, ProjectTestNewResponseThresholdsInsightNameCorrelatedFeatures, ProjectTestNewResponseThresholdsInsightNameCustomMetric, ProjectTestNewResponseThresholdsInsightNameDuplicateRowCount, ProjectTestNewResponseThresholdsInsightNameEmptyFeatures, ProjectTestNewResponseThresholdsInsightNameFeatureDrift, ProjectTestNewResponseThresholdsInsightNameFeatureProfile, ProjectTestNewResponseThresholdsInsightNameGreatExpectations, ProjectTestNewResponseThresholdsInsightNameGroupByColumnStatsCheck, ProjectTestNewResponseThresholdsInsightNameIllFormedRowCount, ProjectTestNewResponseThresholdsInsightNameIsCode, ProjectTestNewResponseThresholdsInsightNameIsJson, ProjectTestNewResponseThresholdsInsightNameLlmRubricV2, ProjectTestNewResponseThresholdsInsightNameLabelDrift, ProjectTestNewResponseThresholdsInsightNameMetrics, ProjectTestNewResponseThresholdsInsightNameNewCategories, ProjectTestNewResponseThresholdsInsightNameNewLabels, ProjectTestNewResponseThresholdsInsightNameNullRowCount, ProjectTestNewResponseThresholdsInsightNamePpScore, ProjectTestNewResponseThresholdsInsightNameQuasiConstantFeatures, ProjectTestNewResponseThresholdsInsightNameSentenceLength, ProjectTestNewResponseThresholdsInsightNameSizeRatio, ProjectTestNewResponseThresholdsInsightNameSpecialCharacters, ProjectTestNewResponseThresholdsInsightNameStringValidation, ProjectTestNewResponseThresholdsInsightNameTrainValLeakageRowCount:
		return true
	}
	return false
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

// The test type.
type ProjectTestNewResponseType string

const (
	ProjectTestNewResponseTypeIntegrity   ProjectTestNewResponseType = "integrity"
	ProjectTestNewResponseTypeConsistency ProjectTestNewResponseType = "consistency"
	ProjectTestNewResponseTypePerformance ProjectTestNewResponseType = "performance"
)

func (r ProjectTestNewResponseType) IsKnown() bool {
	switch r {
	case ProjectTestNewResponseTypeIntegrity, ProjectTestNewResponseTypeConsistency, ProjectTestNewResponseTypePerformance:
		return true
	}
	return false
}

type ProjectTestListResponse struct {
	Items []ProjectTestListResponseItem `json:"items,required"`
	JSON  projectTestListResponseJSON   `json:"-"`
}

// projectTestListResponseJSON contains the JSON metadata for the struct
// [ProjectTestListResponse]
type projectTestListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectTestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectTestListResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectTestListResponseItem struct {
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
	Subtype ProjectTestListResponseItemsSubtype `json:"subtype,required"`
	// Whether the test is suggested or user-created.
	Suggested  bool                                    `json:"suggested,required"`
	Thresholds []ProjectTestListResponseItemsThreshold `json:"thresholds,required"`
	// The test type.
	Type ProjectTestListResponseItemsType `json:"type,required"`
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
	UsesValidationDataset bool                            `json:"usesValidationDataset"`
	JSON                  projectTestListResponseItemJSON `json:"-"`
}

// projectTestListResponseItemJSON contains the JSON metadata for the struct
// [ProjectTestListResponseItem]
type projectTestListResponseItemJSON struct {
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

func (r *ProjectTestListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectTestListResponseItemJSON) RawJSON() string {
	return r.raw
}

// The test subtype.
type ProjectTestListResponseItemsSubtype string

const (
	ProjectTestListResponseItemsSubtypeAnomalousColumnCount       ProjectTestListResponseItemsSubtype = "anomalousColumnCount"
	ProjectTestListResponseItemsSubtypeCharacterLength            ProjectTestListResponseItemsSubtype = "characterLength"
	ProjectTestListResponseItemsSubtypeClassImbalanceRatio        ProjectTestListResponseItemsSubtype = "classImbalanceRatio"
	ProjectTestListResponseItemsSubtypeExpectColumnAToBeInColumnB ProjectTestListResponseItemsSubtype = "expectColumnAToBeInColumnB"
	ProjectTestListResponseItemsSubtypeColumnAverage              ProjectTestListResponseItemsSubtype = "columnAverage"
	ProjectTestListResponseItemsSubtypeColumnDrift                ProjectTestListResponseItemsSubtype = "columnDrift"
	ProjectTestListResponseItemsSubtypeColumnStatistic            ProjectTestListResponseItemsSubtype = "columnStatistic"
	ProjectTestListResponseItemsSubtypeColumnValuesMatch          ProjectTestListResponseItemsSubtype = "columnValuesMatch"
	ProjectTestListResponseItemsSubtypeConflictingLabelRowCount   ProjectTestListResponseItemsSubtype = "conflictingLabelRowCount"
	ProjectTestListResponseItemsSubtypeContainsPii                ProjectTestListResponseItemsSubtype = "containsPii"
	ProjectTestListResponseItemsSubtypeContainsValidURL           ProjectTestListResponseItemsSubtype = "containsValidUrl"
	ProjectTestListResponseItemsSubtypeCorrelatedFeatureCount     ProjectTestListResponseItemsSubtype = "correlatedFeatureCount"
	ProjectTestListResponseItemsSubtypeCustomMetricThreshold      ProjectTestListResponseItemsSubtype = "customMetricThreshold"
	ProjectTestListResponseItemsSubtypeDuplicateRowCount          ProjectTestListResponseItemsSubtype = "duplicateRowCount"
	ProjectTestListResponseItemsSubtypeEmptyFeature               ProjectTestListResponseItemsSubtype = "emptyFeature"
	ProjectTestListResponseItemsSubtypeEmptyFeatureCount          ProjectTestListResponseItemsSubtype = "emptyFeatureCount"
	ProjectTestListResponseItemsSubtypeDriftedFeatureCount        ProjectTestListResponseItemsSubtype = "driftedFeatureCount"
	ProjectTestListResponseItemsSubtypeFeatureMissingValues       ProjectTestListResponseItemsSubtype = "featureMissingValues"
	ProjectTestListResponseItemsSubtypeFeatureValueValidation     ProjectTestListResponseItemsSubtype = "featureValueValidation"
	ProjectTestListResponseItemsSubtypeGreatExpectations          ProjectTestListResponseItemsSubtype = "greatExpectations"
	ProjectTestListResponseItemsSubtypeGroupByColumnStatsCheck    ProjectTestListResponseItemsSubtype = "groupByColumnStatsCheck"
	ProjectTestListResponseItemsSubtypeIllFormedRowCount          ProjectTestListResponseItemsSubtype = "illFormedRowCount"
	ProjectTestListResponseItemsSubtypeIsCode                     ProjectTestListResponseItemsSubtype = "isCode"
	ProjectTestListResponseItemsSubtypeIsJson                     ProjectTestListResponseItemsSubtype = "isJson"
	ProjectTestListResponseItemsSubtypeLlmRubricThresholdV2       ProjectTestListResponseItemsSubtype = "llmRubricThresholdV2"
	ProjectTestListResponseItemsSubtypeLabelDrift                 ProjectTestListResponseItemsSubtype = "labelDrift"
	ProjectTestListResponseItemsSubtypeMetricThreshold            ProjectTestListResponseItemsSubtype = "metricThreshold"
	ProjectTestListResponseItemsSubtypeNewCategoryCount           ProjectTestListResponseItemsSubtype = "newCategoryCount"
	ProjectTestListResponseItemsSubtypeNewLabelCount              ProjectTestListResponseItemsSubtype = "newLabelCount"
	ProjectTestListResponseItemsSubtypeNullRowCount               ProjectTestListResponseItemsSubtype = "nullRowCount"
	ProjectTestListResponseItemsSubtypeRowCount                   ProjectTestListResponseItemsSubtype = "rowCount"
	ProjectTestListResponseItemsSubtypePpScoreValueValidation     ProjectTestListResponseItemsSubtype = "ppScoreValueValidation"
	ProjectTestListResponseItemsSubtypeQuasiConstantFeature       ProjectTestListResponseItemsSubtype = "quasiConstantFeature"
	ProjectTestListResponseItemsSubtypeQuasiConstantFeatureCount  ProjectTestListResponseItemsSubtype = "quasiConstantFeatureCount"
	ProjectTestListResponseItemsSubtypeSqlQuery                   ProjectTestListResponseItemsSubtype = "sqlQuery"
	ProjectTestListResponseItemsSubtypeDtypeValidation            ProjectTestListResponseItemsSubtype = "dtypeValidation"
	ProjectTestListResponseItemsSubtypeSentenceLength             ProjectTestListResponseItemsSubtype = "sentenceLength"
	ProjectTestListResponseItemsSubtypeSizeRatio                  ProjectTestListResponseItemsSubtype = "sizeRatio"
	ProjectTestListResponseItemsSubtypeSpecialCharactersRatio     ProjectTestListResponseItemsSubtype = "specialCharactersRatio"
	ProjectTestListResponseItemsSubtypeStringValidation           ProjectTestListResponseItemsSubtype = "stringValidation"
	ProjectTestListResponseItemsSubtypeTrainValLeakageRowCount    ProjectTestListResponseItemsSubtype = "trainValLeakageRowCount"
)

func (r ProjectTestListResponseItemsSubtype) IsKnown() bool {
	switch r {
	case ProjectTestListResponseItemsSubtypeAnomalousColumnCount, ProjectTestListResponseItemsSubtypeCharacterLength, ProjectTestListResponseItemsSubtypeClassImbalanceRatio, ProjectTestListResponseItemsSubtypeExpectColumnAToBeInColumnB, ProjectTestListResponseItemsSubtypeColumnAverage, ProjectTestListResponseItemsSubtypeColumnDrift, ProjectTestListResponseItemsSubtypeColumnStatistic, ProjectTestListResponseItemsSubtypeColumnValuesMatch, ProjectTestListResponseItemsSubtypeConflictingLabelRowCount, ProjectTestListResponseItemsSubtypeContainsPii, ProjectTestListResponseItemsSubtypeContainsValidURL, ProjectTestListResponseItemsSubtypeCorrelatedFeatureCount, ProjectTestListResponseItemsSubtypeCustomMetricThreshold, ProjectTestListResponseItemsSubtypeDuplicateRowCount, ProjectTestListResponseItemsSubtypeEmptyFeature, ProjectTestListResponseItemsSubtypeEmptyFeatureCount, ProjectTestListResponseItemsSubtypeDriftedFeatureCount, ProjectTestListResponseItemsSubtypeFeatureMissingValues, ProjectTestListResponseItemsSubtypeFeatureValueValidation, ProjectTestListResponseItemsSubtypeGreatExpectations, ProjectTestListResponseItemsSubtypeGroupByColumnStatsCheck, ProjectTestListResponseItemsSubtypeIllFormedRowCount, ProjectTestListResponseItemsSubtypeIsCode, ProjectTestListResponseItemsSubtypeIsJson, ProjectTestListResponseItemsSubtypeLlmRubricThresholdV2, ProjectTestListResponseItemsSubtypeLabelDrift, ProjectTestListResponseItemsSubtypeMetricThreshold, ProjectTestListResponseItemsSubtypeNewCategoryCount, ProjectTestListResponseItemsSubtypeNewLabelCount, ProjectTestListResponseItemsSubtypeNullRowCount, ProjectTestListResponseItemsSubtypeRowCount, ProjectTestListResponseItemsSubtypePpScoreValueValidation, ProjectTestListResponseItemsSubtypeQuasiConstantFeature, ProjectTestListResponseItemsSubtypeQuasiConstantFeatureCount, ProjectTestListResponseItemsSubtypeSqlQuery, ProjectTestListResponseItemsSubtypeDtypeValidation, ProjectTestListResponseItemsSubtypeSentenceLength, ProjectTestListResponseItemsSubtypeSizeRatio, ProjectTestListResponseItemsSubtypeSpecialCharactersRatio, ProjectTestListResponseItemsSubtypeStringValidation, ProjectTestListResponseItemsSubtypeTrainValLeakageRowCount:
		return true
	}
	return false
}

type ProjectTestListResponseItemsThreshold struct {
	// The insight name to be evaluated.
	InsightName ProjectTestListResponseItemsThresholdsInsightName `json:"insightName"`
	// The insight parameters. Required only for some test subtypes. For example, for
	// tests that require a column name, the insight parameters will be [{'name':
	// 'column_name', 'value': 'Age'}]
	InsightParameters []ProjectTestListResponseItemsThresholdsInsightParameter `json:"insightParameters,nullable"`
	// The measurement to be evaluated.
	Measurement string `json:"measurement"`
	// The operator to be used for the evaluation.
	Operator ProjectTestListResponseItemsThresholdsOperator `json:"operator"`
	// Whether to use automatic anomaly detection or manual thresholds
	ThresholdMode ProjectTestListResponseItemsThresholdsThresholdMode `json:"thresholdMode"`
	// The value to be compared.
	Value ProjectTestListResponseItemsThresholdsValueUnion `json:"value"`
	JSON  projectTestListResponseItemsThresholdJSON        `json:"-"`
}

// projectTestListResponseItemsThresholdJSON contains the JSON metadata for the
// struct [ProjectTestListResponseItemsThreshold]
type projectTestListResponseItemsThresholdJSON struct {
	InsightName       apijson.Field
	InsightParameters apijson.Field
	Measurement       apijson.Field
	Operator          apijson.Field
	ThresholdMode     apijson.Field
	Value             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectTestListResponseItemsThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectTestListResponseItemsThresholdJSON) RawJSON() string {
	return r.raw
}

// The insight name to be evaluated.
type ProjectTestListResponseItemsThresholdsInsightName string

const (
	ProjectTestListResponseItemsThresholdsInsightNameCharacterLength            ProjectTestListResponseItemsThresholdsInsightName = "characterLength"
	ProjectTestListResponseItemsThresholdsInsightNameClassImbalance             ProjectTestListResponseItemsThresholdsInsightName = "classImbalance"
	ProjectTestListResponseItemsThresholdsInsightNameExpectColumnAToBeInColumnB ProjectTestListResponseItemsThresholdsInsightName = "expectColumnAToBeInColumnB"
	ProjectTestListResponseItemsThresholdsInsightNameColumnAverage              ProjectTestListResponseItemsThresholdsInsightName = "columnAverage"
	ProjectTestListResponseItemsThresholdsInsightNameColumnDrift                ProjectTestListResponseItemsThresholdsInsightName = "columnDrift"
	ProjectTestListResponseItemsThresholdsInsightNameColumnValuesMatch          ProjectTestListResponseItemsThresholdsInsightName = "columnValuesMatch"
	ProjectTestListResponseItemsThresholdsInsightNameConfidenceDistribution     ProjectTestListResponseItemsThresholdsInsightName = "confidenceDistribution"
	ProjectTestListResponseItemsThresholdsInsightNameConflictingLabelRowCount   ProjectTestListResponseItemsThresholdsInsightName = "conflictingLabelRowCount"
	ProjectTestListResponseItemsThresholdsInsightNameContainsPii                ProjectTestListResponseItemsThresholdsInsightName = "containsPii"
	ProjectTestListResponseItemsThresholdsInsightNameContainsValidURL           ProjectTestListResponseItemsThresholdsInsightName = "containsValidUrl"
	ProjectTestListResponseItemsThresholdsInsightNameCorrelatedFeatures         ProjectTestListResponseItemsThresholdsInsightName = "correlatedFeatures"
	ProjectTestListResponseItemsThresholdsInsightNameCustomMetric               ProjectTestListResponseItemsThresholdsInsightName = "customMetric"
	ProjectTestListResponseItemsThresholdsInsightNameDuplicateRowCount          ProjectTestListResponseItemsThresholdsInsightName = "duplicateRowCount"
	ProjectTestListResponseItemsThresholdsInsightNameEmptyFeatures              ProjectTestListResponseItemsThresholdsInsightName = "emptyFeatures"
	ProjectTestListResponseItemsThresholdsInsightNameFeatureDrift               ProjectTestListResponseItemsThresholdsInsightName = "featureDrift"
	ProjectTestListResponseItemsThresholdsInsightNameFeatureProfile             ProjectTestListResponseItemsThresholdsInsightName = "featureProfile"
	ProjectTestListResponseItemsThresholdsInsightNameGreatExpectations          ProjectTestListResponseItemsThresholdsInsightName = "greatExpectations"
	ProjectTestListResponseItemsThresholdsInsightNameGroupByColumnStatsCheck    ProjectTestListResponseItemsThresholdsInsightName = "groupByColumnStatsCheck"
	ProjectTestListResponseItemsThresholdsInsightNameIllFormedRowCount          ProjectTestListResponseItemsThresholdsInsightName = "illFormedRowCount"
	ProjectTestListResponseItemsThresholdsInsightNameIsCode                     ProjectTestListResponseItemsThresholdsInsightName = "isCode"
	ProjectTestListResponseItemsThresholdsInsightNameIsJson                     ProjectTestListResponseItemsThresholdsInsightName = "isJson"
	ProjectTestListResponseItemsThresholdsInsightNameLlmRubricV2                ProjectTestListResponseItemsThresholdsInsightName = "llmRubricV2"
	ProjectTestListResponseItemsThresholdsInsightNameLabelDrift                 ProjectTestListResponseItemsThresholdsInsightName = "labelDrift"
	ProjectTestListResponseItemsThresholdsInsightNameMetrics                    ProjectTestListResponseItemsThresholdsInsightName = "metrics"
	ProjectTestListResponseItemsThresholdsInsightNameNewCategories              ProjectTestListResponseItemsThresholdsInsightName = "newCategories"
	ProjectTestListResponseItemsThresholdsInsightNameNewLabels                  ProjectTestListResponseItemsThresholdsInsightName = "newLabels"
	ProjectTestListResponseItemsThresholdsInsightNameNullRowCount               ProjectTestListResponseItemsThresholdsInsightName = "nullRowCount"
	ProjectTestListResponseItemsThresholdsInsightNamePpScore                    ProjectTestListResponseItemsThresholdsInsightName = "ppScore"
	ProjectTestListResponseItemsThresholdsInsightNameQuasiConstantFeatures      ProjectTestListResponseItemsThresholdsInsightName = "quasiConstantFeatures"
	ProjectTestListResponseItemsThresholdsInsightNameSentenceLength             ProjectTestListResponseItemsThresholdsInsightName = "sentenceLength"
	ProjectTestListResponseItemsThresholdsInsightNameSizeRatio                  ProjectTestListResponseItemsThresholdsInsightName = "sizeRatio"
	ProjectTestListResponseItemsThresholdsInsightNameSpecialCharacters          ProjectTestListResponseItemsThresholdsInsightName = "specialCharacters"
	ProjectTestListResponseItemsThresholdsInsightNameStringValidation           ProjectTestListResponseItemsThresholdsInsightName = "stringValidation"
	ProjectTestListResponseItemsThresholdsInsightNameTrainValLeakageRowCount    ProjectTestListResponseItemsThresholdsInsightName = "trainValLeakageRowCount"
)

func (r ProjectTestListResponseItemsThresholdsInsightName) IsKnown() bool {
	switch r {
	case ProjectTestListResponseItemsThresholdsInsightNameCharacterLength, ProjectTestListResponseItemsThresholdsInsightNameClassImbalance, ProjectTestListResponseItemsThresholdsInsightNameExpectColumnAToBeInColumnB, ProjectTestListResponseItemsThresholdsInsightNameColumnAverage, ProjectTestListResponseItemsThresholdsInsightNameColumnDrift, ProjectTestListResponseItemsThresholdsInsightNameColumnValuesMatch, ProjectTestListResponseItemsThresholdsInsightNameConfidenceDistribution, ProjectTestListResponseItemsThresholdsInsightNameConflictingLabelRowCount, ProjectTestListResponseItemsThresholdsInsightNameContainsPii, ProjectTestListResponseItemsThresholdsInsightNameContainsValidURL, ProjectTestListResponseItemsThresholdsInsightNameCorrelatedFeatures, ProjectTestListResponseItemsThresholdsInsightNameCustomMetric, ProjectTestListResponseItemsThresholdsInsightNameDuplicateRowCount, ProjectTestListResponseItemsThresholdsInsightNameEmptyFeatures, ProjectTestListResponseItemsThresholdsInsightNameFeatureDrift, ProjectTestListResponseItemsThresholdsInsightNameFeatureProfile, ProjectTestListResponseItemsThresholdsInsightNameGreatExpectations, ProjectTestListResponseItemsThresholdsInsightNameGroupByColumnStatsCheck, ProjectTestListResponseItemsThresholdsInsightNameIllFormedRowCount, ProjectTestListResponseItemsThresholdsInsightNameIsCode, ProjectTestListResponseItemsThresholdsInsightNameIsJson, ProjectTestListResponseItemsThresholdsInsightNameLlmRubricV2, ProjectTestListResponseItemsThresholdsInsightNameLabelDrift, ProjectTestListResponseItemsThresholdsInsightNameMetrics, ProjectTestListResponseItemsThresholdsInsightNameNewCategories, ProjectTestListResponseItemsThresholdsInsightNameNewLabels, ProjectTestListResponseItemsThresholdsInsightNameNullRowCount, ProjectTestListResponseItemsThresholdsInsightNamePpScore, ProjectTestListResponseItemsThresholdsInsightNameQuasiConstantFeatures, ProjectTestListResponseItemsThresholdsInsightNameSentenceLength, ProjectTestListResponseItemsThresholdsInsightNameSizeRatio, ProjectTestListResponseItemsThresholdsInsightNameSpecialCharacters, ProjectTestListResponseItemsThresholdsInsightNameStringValidation, ProjectTestListResponseItemsThresholdsInsightNameTrainValLeakageRowCount:
		return true
	}
	return false
}

type ProjectTestListResponseItemsThresholdsInsightParameter struct {
	// The name of the insight filter.
	Name  string                                                     `json:"name,required"`
	Value interface{}                                                `json:"value,required"`
	JSON  projectTestListResponseItemsThresholdsInsightParameterJSON `json:"-"`
}

// projectTestListResponseItemsThresholdsInsightParameterJSON contains the JSON
// metadata for the struct [ProjectTestListResponseItemsThresholdsInsightParameter]
type projectTestListResponseItemsThresholdsInsightParameterJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectTestListResponseItemsThresholdsInsightParameter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectTestListResponseItemsThresholdsInsightParameterJSON) RawJSON() string {
	return r.raw
}

// The operator to be used for the evaluation.
type ProjectTestListResponseItemsThresholdsOperator string

const (
	ProjectTestListResponseItemsThresholdsOperatorIs              ProjectTestListResponseItemsThresholdsOperator = "is"
	ProjectTestListResponseItemsThresholdsOperatorGreater         ProjectTestListResponseItemsThresholdsOperator = ">"
	ProjectTestListResponseItemsThresholdsOperatorGreaterOrEquals ProjectTestListResponseItemsThresholdsOperator = ">="
	ProjectTestListResponseItemsThresholdsOperatorLess            ProjectTestListResponseItemsThresholdsOperator = "<"
	ProjectTestListResponseItemsThresholdsOperatorLessOrEquals    ProjectTestListResponseItemsThresholdsOperator = "<="
	ProjectTestListResponseItemsThresholdsOperatorNotEquals       ProjectTestListResponseItemsThresholdsOperator = "!="
)

func (r ProjectTestListResponseItemsThresholdsOperator) IsKnown() bool {
	switch r {
	case ProjectTestListResponseItemsThresholdsOperatorIs, ProjectTestListResponseItemsThresholdsOperatorGreater, ProjectTestListResponseItemsThresholdsOperatorGreaterOrEquals, ProjectTestListResponseItemsThresholdsOperatorLess, ProjectTestListResponseItemsThresholdsOperatorLessOrEquals, ProjectTestListResponseItemsThresholdsOperatorNotEquals:
		return true
	}
	return false
}

// Whether to use automatic anomaly detection or manual thresholds
type ProjectTestListResponseItemsThresholdsThresholdMode string

const (
	ProjectTestListResponseItemsThresholdsThresholdModeAutomatic ProjectTestListResponseItemsThresholdsThresholdMode = "automatic"
	ProjectTestListResponseItemsThresholdsThresholdModeManual    ProjectTestListResponseItemsThresholdsThresholdMode = "manual"
)

func (r ProjectTestListResponseItemsThresholdsThresholdMode) IsKnown() bool {
	switch r {
	case ProjectTestListResponseItemsThresholdsThresholdModeAutomatic, ProjectTestListResponseItemsThresholdsThresholdModeManual:
		return true
	}
	return false
}

// The value to be compared.
//
// Union satisfied by [shared.UnionFloat], [shared.UnionBool], [shared.UnionString]
// or [ProjectTestListResponseItemsThresholdsValueArray].
type ProjectTestListResponseItemsThresholdsValueUnion interface {
	ImplementsProjectTestListResponseItemsThresholdsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectTestListResponseItemsThresholdsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(ProjectTestListResponseItemsThresholdsValueArray{}),
		},
	)
}

type ProjectTestListResponseItemsThresholdsValueArray []string

func (r ProjectTestListResponseItemsThresholdsValueArray) ImplementsProjectTestListResponseItemsThresholdsValueUnion() {
}

// The test type.
type ProjectTestListResponseItemsType string

const (
	ProjectTestListResponseItemsTypeIntegrity   ProjectTestListResponseItemsType = "integrity"
	ProjectTestListResponseItemsTypeConsistency ProjectTestListResponseItemsType = "consistency"
	ProjectTestListResponseItemsTypePerformance ProjectTestListResponseItemsType = "performance"
)

func (r ProjectTestListResponseItemsType) IsKnown() bool {
	switch r {
	case ProjectTestListResponseItemsTypeIntegrity, ProjectTestListResponseItemsTypeConsistency, ProjectTestListResponseItemsTypePerformance:
		return true
	}
	return false
}

type ProjectTestNewParams struct {
	// The test description.
	Description param.Field[interface{}] `json:"description,required"`
	// The test name.
	Name param.Field[string] `json:"name,required"`
	// The test subtype.
	Subtype    param.Field[ProjectTestNewParamsSubtype]     `json:"subtype,required"`
	Thresholds param.Field[[]ProjectTestNewParamsThreshold] `json:"thresholds,required"`
	// The test type.
	Type param.Field[ProjectTestNewParamsType] `json:"type,required"`
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

// The test subtype.
type ProjectTestNewParamsSubtype string

const (
	ProjectTestNewParamsSubtypeAnomalousColumnCount       ProjectTestNewParamsSubtype = "anomalousColumnCount"
	ProjectTestNewParamsSubtypeCharacterLength            ProjectTestNewParamsSubtype = "characterLength"
	ProjectTestNewParamsSubtypeClassImbalanceRatio        ProjectTestNewParamsSubtype = "classImbalanceRatio"
	ProjectTestNewParamsSubtypeExpectColumnAToBeInColumnB ProjectTestNewParamsSubtype = "expectColumnAToBeInColumnB"
	ProjectTestNewParamsSubtypeColumnAverage              ProjectTestNewParamsSubtype = "columnAverage"
	ProjectTestNewParamsSubtypeColumnDrift                ProjectTestNewParamsSubtype = "columnDrift"
	ProjectTestNewParamsSubtypeColumnStatistic            ProjectTestNewParamsSubtype = "columnStatistic"
	ProjectTestNewParamsSubtypeColumnValuesMatch          ProjectTestNewParamsSubtype = "columnValuesMatch"
	ProjectTestNewParamsSubtypeConflictingLabelRowCount   ProjectTestNewParamsSubtype = "conflictingLabelRowCount"
	ProjectTestNewParamsSubtypeContainsPii                ProjectTestNewParamsSubtype = "containsPii"
	ProjectTestNewParamsSubtypeContainsValidURL           ProjectTestNewParamsSubtype = "containsValidUrl"
	ProjectTestNewParamsSubtypeCorrelatedFeatureCount     ProjectTestNewParamsSubtype = "correlatedFeatureCount"
	ProjectTestNewParamsSubtypeCustomMetricThreshold      ProjectTestNewParamsSubtype = "customMetricThreshold"
	ProjectTestNewParamsSubtypeDuplicateRowCount          ProjectTestNewParamsSubtype = "duplicateRowCount"
	ProjectTestNewParamsSubtypeEmptyFeature               ProjectTestNewParamsSubtype = "emptyFeature"
	ProjectTestNewParamsSubtypeEmptyFeatureCount          ProjectTestNewParamsSubtype = "emptyFeatureCount"
	ProjectTestNewParamsSubtypeDriftedFeatureCount        ProjectTestNewParamsSubtype = "driftedFeatureCount"
	ProjectTestNewParamsSubtypeFeatureMissingValues       ProjectTestNewParamsSubtype = "featureMissingValues"
	ProjectTestNewParamsSubtypeFeatureValueValidation     ProjectTestNewParamsSubtype = "featureValueValidation"
	ProjectTestNewParamsSubtypeGreatExpectations          ProjectTestNewParamsSubtype = "greatExpectations"
	ProjectTestNewParamsSubtypeGroupByColumnStatsCheck    ProjectTestNewParamsSubtype = "groupByColumnStatsCheck"
	ProjectTestNewParamsSubtypeIllFormedRowCount          ProjectTestNewParamsSubtype = "illFormedRowCount"
	ProjectTestNewParamsSubtypeIsCode                     ProjectTestNewParamsSubtype = "isCode"
	ProjectTestNewParamsSubtypeIsJson                     ProjectTestNewParamsSubtype = "isJson"
	ProjectTestNewParamsSubtypeLlmRubricThresholdV2       ProjectTestNewParamsSubtype = "llmRubricThresholdV2"
	ProjectTestNewParamsSubtypeLabelDrift                 ProjectTestNewParamsSubtype = "labelDrift"
	ProjectTestNewParamsSubtypeMetricThreshold            ProjectTestNewParamsSubtype = "metricThreshold"
	ProjectTestNewParamsSubtypeNewCategoryCount           ProjectTestNewParamsSubtype = "newCategoryCount"
	ProjectTestNewParamsSubtypeNewLabelCount              ProjectTestNewParamsSubtype = "newLabelCount"
	ProjectTestNewParamsSubtypeNullRowCount               ProjectTestNewParamsSubtype = "nullRowCount"
	ProjectTestNewParamsSubtypeRowCount                   ProjectTestNewParamsSubtype = "rowCount"
	ProjectTestNewParamsSubtypePpScoreValueValidation     ProjectTestNewParamsSubtype = "ppScoreValueValidation"
	ProjectTestNewParamsSubtypeQuasiConstantFeature       ProjectTestNewParamsSubtype = "quasiConstantFeature"
	ProjectTestNewParamsSubtypeQuasiConstantFeatureCount  ProjectTestNewParamsSubtype = "quasiConstantFeatureCount"
	ProjectTestNewParamsSubtypeSqlQuery                   ProjectTestNewParamsSubtype = "sqlQuery"
	ProjectTestNewParamsSubtypeDtypeValidation            ProjectTestNewParamsSubtype = "dtypeValidation"
	ProjectTestNewParamsSubtypeSentenceLength             ProjectTestNewParamsSubtype = "sentenceLength"
	ProjectTestNewParamsSubtypeSizeRatio                  ProjectTestNewParamsSubtype = "sizeRatio"
	ProjectTestNewParamsSubtypeSpecialCharactersRatio     ProjectTestNewParamsSubtype = "specialCharactersRatio"
	ProjectTestNewParamsSubtypeStringValidation           ProjectTestNewParamsSubtype = "stringValidation"
	ProjectTestNewParamsSubtypeTrainValLeakageRowCount    ProjectTestNewParamsSubtype = "trainValLeakageRowCount"
)

func (r ProjectTestNewParamsSubtype) IsKnown() bool {
	switch r {
	case ProjectTestNewParamsSubtypeAnomalousColumnCount, ProjectTestNewParamsSubtypeCharacterLength, ProjectTestNewParamsSubtypeClassImbalanceRatio, ProjectTestNewParamsSubtypeExpectColumnAToBeInColumnB, ProjectTestNewParamsSubtypeColumnAverage, ProjectTestNewParamsSubtypeColumnDrift, ProjectTestNewParamsSubtypeColumnStatistic, ProjectTestNewParamsSubtypeColumnValuesMatch, ProjectTestNewParamsSubtypeConflictingLabelRowCount, ProjectTestNewParamsSubtypeContainsPii, ProjectTestNewParamsSubtypeContainsValidURL, ProjectTestNewParamsSubtypeCorrelatedFeatureCount, ProjectTestNewParamsSubtypeCustomMetricThreshold, ProjectTestNewParamsSubtypeDuplicateRowCount, ProjectTestNewParamsSubtypeEmptyFeature, ProjectTestNewParamsSubtypeEmptyFeatureCount, ProjectTestNewParamsSubtypeDriftedFeatureCount, ProjectTestNewParamsSubtypeFeatureMissingValues, ProjectTestNewParamsSubtypeFeatureValueValidation, ProjectTestNewParamsSubtypeGreatExpectations, ProjectTestNewParamsSubtypeGroupByColumnStatsCheck, ProjectTestNewParamsSubtypeIllFormedRowCount, ProjectTestNewParamsSubtypeIsCode, ProjectTestNewParamsSubtypeIsJson, ProjectTestNewParamsSubtypeLlmRubricThresholdV2, ProjectTestNewParamsSubtypeLabelDrift, ProjectTestNewParamsSubtypeMetricThreshold, ProjectTestNewParamsSubtypeNewCategoryCount, ProjectTestNewParamsSubtypeNewLabelCount, ProjectTestNewParamsSubtypeNullRowCount, ProjectTestNewParamsSubtypeRowCount, ProjectTestNewParamsSubtypePpScoreValueValidation, ProjectTestNewParamsSubtypeQuasiConstantFeature, ProjectTestNewParamsSubtypeQuasiConstantFeatureCount, ProjectTestNewParamsSubtypeSqlQuery, ProjectTestNewParamsSubtypeDtypeValidation, ProjectTestNewParamsSubtypeSentenceLength, ProjectTestNewParamsSubtypeSizeRatio, ProjectTestNewParamsSubtypeSpecialCharactersRatio, ProjectTestNewParamsSubtypeStringValidation, ProjectTestNewParamsSubtypeTrainValLeakageRowCount:
		return true
	}
	return false
}

type ProjectTestNewParamsThreshold struct {
	// The insight name to be evaluated.
	InsightName param.Field[ProjectTestNewParamsThresholdsInsightName] `json:"insightName"`
	// The insight parameters. Required only for some test subtypes. For example, for
	// tests that require a column name, the insight parameters will be [{'name':
	// 'column_name', 'value': 'Age'}]
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

// The insight name to be evaluated.
type ProjectTestNewParamsThresholdsInsightName string

const (
	ProjectTestNewParamsThresholdsInsightNameCharacterLength            ProjectTestNewParamsThresholdsInsightName = "characterLength"
	ProjectTestNewParamsThresholdsInsightNameClassImbalance             ProjectTestNewParamsThresholdsInsightName = "classImbalance"
	ProjectTestNewParamsThresholdsInsightNameExpectColumnAToBeInColumnB ProjectTestNewParamsThresholdsInsightName = "expectColumnAToBeInColumnB"
	ProjectTestNewParamsThresholdsInsightNameColumnAverage              ProjectTestNewParamsThresholdsInsightName = "columnAverage"
	ProjectTestNewParamsThresholdsInsightNameColumnDrift                ProjectTestNewParamsThresholdsInsightName = "columnDrift"
	ProjectTestNewParamsThresholdsInsightNameColumnValuesMatch          ProjectTestNewParamsThresholdsInsightName = "columnValuesMatch"
	ProjectTestNewParamsThresholdsInsightNameConfidenceDistribution     ProjectTestNewParamsThresholdsInsightName = "confidenceDistribution"
	ProjectTestNewParamsThresholdsInsightNameConflictingLabelRowCount   ProjectTestNewParamsThresholdsInsightName = "conflictingLabelRowCount"
	ProjectTestNewParamsThresholdsInsightNameContainsPii                ProjectTestNewParamsThresholdsInsightName = "containsPii"
	ProjectTestNewParamsThresholdsInsightNameContainsValidURL           ProjectTestNewParamsThresholdsInsightName = "containsValidUrl"
	ProjectTestNewParamsThresholdsInsightNameCorrelatedFeatures         ProjectTestNewParamsThresholdsInsightName = "correlatedFeatures"
	ProjectTestNewParamsThresholdsInsightNameCustomMetric               ProjectTestNewParamsThresholdsInsightName = "customMetric"
	ProjectTestNewParamsThresholdsInsightNameDuplicateRowCount          ProjectTestNewParamsThresholdsInsightName = "duplicateRowCount"
	ProjectTestNewParamsThresholdsInsightNameEmptyFeatures              ProjectTestNewParamsThresholdsInsightName = "emptyFeatures"
	ProjectTestNewParamsThresholdsInsightNameFeatureDrift               ProjectTestNewParamsThresholdsInsightName = "featureDrift"
	ProjectTestNewParamsThresholdsInsightNameFeatureProfile             ProjectTestNewParamsThresholdsInsightName = "featureProfile"
	ProjectTestNewParamsThresholdsInsightNameGreatExpectations          ProjectTestNewParamsThresholdsInsightName = "greatExpectations"
	ProjectTestNewParamsThresholdsInsightNameGroupByColumnStatsCheck    ProjectTestNewParamsThresholdsInsightName = "groupByColumnStatsCheck"
	ProjectTestNewParamsThresholdsInsightNameIllFormedRowCount          ProjectTestNewParamsThresholdsInsightName = "illFormedRowCount"
	ProjectTestNewParamsThresholdsInsightNameIsCode                     ProjectTestNewParamsThresholdsInsightName = "isCode"
	ProjectTestNewParamsThresholdsInsightNameIsJson                     ProjectTestNewParamsThresholdsInsightName = "isJson"
	ProjectTestNewParamsThresholdsInsightNameLlmRubricV2                ProjectTestNewParamsThresholdsInsightName = "llmRubricV2"
	ProjectTestNewParamsThresholdsInsightNameLabelDrift                 ProjectTestNewParamsThresholdsInsightName = "labelDrift"
	ProjectTestNewParamsThresholdsInsightNameMetrics                    ProjectTestNewParamsThresholdsInsightName = "metrics"
	ProjectTestNewParamsThresholdsInsightNameNewCategories              ProjectTestNewParamsThresholdsInsightName = "newCategories"
	ProjectTestNewParamsThresholdsInsightNameNewLabels                  ProjectTestNewParamsThresholdsInsightName = "newLabels"
	ProjectTestNewParamsThresholdsInsightNameNullRowCount               ProjectTestNewParamsThresholdsInsightName = "nullRowCount"
	ProjectTestNewParamsThresholdsInsightNamePpScore                    ProjectTestNewParamsThresholdsInsightName = "ppScore"
	ProjectTestNewParamsThresholdsInsightNameQuasiConstantFeatures      ProjectTestNewParamsThresholdsInsightName = "quasiConstantFeatures"
	ProjectTestNewParamsThresholdsInsightNameSentenceLength             ProjectTestNewParamsThresholdsInsightName = "sentenceLength"
	ProjectTestNewParamsThresholdsInsightNameSizeRatio                  ProjectTestNewParamsThresholdsInsightName = "sizeRatio"
	ProjectTestNewParamsThresholdsInsightNameSpecialCharacters          ProjectTestNewParamsThresholdsInsightName = "specialCharacters"
	ProjectTestNewParamsThresholdsInsightNameStringValidation           ProjectTestNewParamsThresholdsInsightName = "stringValidation"
	ProjectTestNewParamsThresholdsInsightNameTrainValLeakageRowCount    ProjectTestNewParamsThresholdsInsightName = "trainValLeakageRowCount"
)

func (r ProjectTestNewParamsThresholdsInsightName) IsKnown() bool {
	switch r {
	case ProjectTestNewParamsThresholdsInsightNameCharacterLength, ProjectTestNewParamsThresholdsInsightNameClassImbalance, ProjectTestNewParamsThresholdsInsightNameExpectColumnAToBeInColumnB, ProjectTestNewParamsThresholdsInsightNameColumnAverage, ProjectTestNewParamsThresholdsInsightNameColumnDrift, ProjectTestNewParamsThresholdsInsightNameColumnValuesMatch, ProjectTestNewParamsThresholdsInsightNameConfidenceDistribution, ProjectTestNewParamsThresholdsInsightNameConflictingLabelRowCount, ProjectTestNewParamsThresholdsInsightNameContainsPii, ProjectTestNewParamsThresholdsInsightNameContainsValidURL, ProjectTestNewParamsThresholdsInsightNameCorrelatedFeatures, ProjectTestNewParamsThresholdsInsightNameCustomMetric, ProjectTestNewParamsThresholdsInsightNameDuplicateRowCount, ProjectTestNewParamsThresholdsInsightNameEmptyFeatures, ProjectTestNewParamsThresholdsInsightNameFeatureDrift, ProjectTestNewParamsThresholdsInsightNameFeatureProfile, ProjectTestNewParamsThresholdsInsightNameGreatExpectations, ProjectTestNewParamsThresholdsInsightNameGroupByColumnStatsCheck, ProjectTestNewParamsThresholdsInsightNameIllFormedRowCount, ProjectTestNewParamsThresholdsInsightNameIsCode, ProjectTestNewParamsThresholdsInsightNameIsJson, ProjectTestNewParamsThresholdsInsightNameLlmRubricV2, ProjectTestNewParamsThresholdsInsightNameLabelDrift, ProjectTestNewParamsThresholdsInsightNameMetrics, ProjectTestNewParamsThresholdsInsightNameNewCategories, ProjectTestNewParamsThresholdsInsightNameNewLabels, ProjectTestNewParamsThresholdsInsightNameNullRowCount, ProjectTestNewParamsThresholdsInsightNamePpScore, ProjectTestNewParamsThresholdsInsightNameQuasiConstantFeatures, ProjectTestNewParamsThresholdsInsightNameSentenceLength, ProjectTestNewParamsThresholdsInsightNameSizeRatio, ProjectTestNewParamsThresholdsInsightNameSpecialCharacters, ProjectTestNewParamsThresholdsInsightNameStringValidation, ProjectTestNewParamsThresholdsInsightNameTrainValLeakageRowCount:
		return true
	}
	return false
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

// The test type.
type ProjectTestNewParamsType string

const (
	ProjectTestNewParamsTypeIntegrity   ProjectTestNewParamsType = "integrity"
	ProjectTestNewParamsTypeConsistency ProjectTestNewParamsType = "consistency"
	ProjectTestNewParamsTypePerformance ProjectTestNewParamsType = "performance"
)

func (r ProjectTestNewParamsType) IsKnown() bool {
	switch r {
	case ProjectTestNewParamsTypeIntegrity, ProjectTestNewParamsTypeConsistency, ProjectTestNewParamsTypePerformance:
		return true
	}
	return false
}

type ProjectTestListParams struct {
	// Filter for archived tests.
	IncludeArchived param.Field[bool] `query:"includeArchived"`
	// Retrive tests created by a specific project version.
	OriginVersionID param.Field[string] `query:"originVersionId" format:"uuid"`
	// The page to return in a paginated query.
	Page param.Field[int64] `query:"page"`
	// Maximum number of items to return per page.
	PerPage param.Field[int64] `query:"perPage"`
	// Filter for suggested tests.
	Suggested param.Field[bool] `query:"suggested"`
	// Filter objects by test type. Available types are `integrity`, `consistency`,
	// `performance`, `fairness`, and `robustness`.
	Type param.Field[ProjectTestListParamsType] `query:"type"`
	// Retrive tests with usesProductionData (monitoring).
	UsesProductionData param.Field[bool] `query:"usesProductionData"`
}

// URLQuery serializes [ProjectTestListParams]'s query parameters as `url.Values`.
func (r ProjectTestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter objects by test type. Available types are `integrity`, `consistency`,
// `performance`, `fairness`, and `robustness`.
type ProjectTestListParamsType string

const (
	ProjectTestListParamsTypeIntegrity   ProjectTestListParamsType = "integrity"
	ProjectTestListParamsTypeConsistency ProjectTestListParamsType = "consistency"
	ProjectTestListParamsTypePerformance ProjectTestListParamsType = "performance"
	ProjectTestListParamsTypeFairness    ProjectTestListParamsType = "fairness"
	ProjectTestListParamsTypeRobustness  ProjectTestListParamsType = "robustness"
)

func (r ProjectTestListParamsType) IsKnown() bool {
	switch r {
	case ProjectTestListParamsTypeIntegrity, ProjectTestListParamsTypeConsistency, ProjectTestListParamsTypePerformance, ProjectTestListParamsTypeFairness, ProjectTestListParamsTypeRobustness:
		return true
	}
	return false
}
