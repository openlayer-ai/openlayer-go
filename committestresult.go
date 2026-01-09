// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
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
	opts = slices.Concat(r.Options, opts)
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
	StatusMessage  string                                           `json:"statusMessage,required,nullable"`
	ExpectedValues []CommitTestResultListResponseItemsExpectedValue `json:"expectedValues"`
	Goal           CommitTestResultListResponseItemsGoal            `json:"goal"`
	// The test id.
	GoalID string `json:"goalId,nullable" format:"uuid"`
	// The URL to the rows of the test result.
	Rows string `json:"rows"`
	// The body of the rows request.
	RowsBody CommitTestResultListResponseItemsRowsBody `json:"rowsBody,nullable"`
	JSON     commitTestResultListResponseItemJSON      `json:"-"`
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
	ExpectedValues      apijson.Field
	Goal                apijson.Field
	GoalID              apijson.Field
	Rows                apijson.Field
	RowsBody            apijson.Field
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

type CommitTestResultListResponseItemsExpectedValue struct {
	// the lower threshold for the expected value
	LowerThreshold float64 `json:"lowerThreshold,nullable"`
	// One of the `measurement` values in the test's thresholds
	Measurement string `json:"measurement"`
	// The upper threshold for the expected value
	UpperThreshold float64                                            `json:"upperThreshold,nullable"`
	JSON           commitTestResultListResponseItemsExpectedValueJSON `json:"-"`
}

// commitTestResultListResponseItemsExpectedValueJSON contains the JSON metadata
// for the struct [CommitTestResultListResponseItemsExpectedValue]
type commitTestResultListResponseItemsExpectedValueJSON struct {
	LowerThreshold apijson.Field
	Measurement    apijson.Field
	UpperThreshold apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CommitTestResultListResponseItemsExpectedValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitTestResultListResponseItemsExpectedValueJSON) RawJSON() string {
	return r.raw
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
	Subtype CommitTestResultListResponseItemsGoalSubtype `json:"subtype,required"`
	// Whether the test is suggested or user-created.
	Suggested  bool                                             `json:"suggested,required"`
	Thresholds []CommitTestResultListResponseItemsGoalThreshold `json:"thresholds,required"`
	// The test type.
	Type CommitTestResultListResponseItemsGoalType `json:"type,required"`
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

// The test subtype.
type CommitTestResultListResponseItemsGoalSubtype string

const (
	CommitTestResultListResponseItemsGoalSubtypeAnomalousColumnCount       CommitTestResultListResponseItemsGoalSubtype = "anomalousColumnCount"
	CommitTestResultListResponseItemsGoalSubtypeCharacterLength            CommitTestResultListResponseItemsGoalSubtype = "characterLength"
	CommitTestResultListResponseItemsGoalSubtypeClassImbalanceRatio        CommitTestResultListResponseItemsGoalSubtype = "classImbalanceRatio"
	CommitTestResultListResponseItemsGoalSubtypeExpectColumnAToBeInColumnB CommitTestResultListResponseItemsGoalSubtype = "expectColumnAToBeInColumnB"
	CommitTestResultListResponseItemsGoalSubtypeColumnAverage              CommitTestResultListResponseItemsGoalSubtype = "columnAverage"
	CommitTestResultListResponseItemsGoalSubtypeColumnDrift                CommitTestResultListResponseItemsGoalSubtype = "columnDrift"
	CommitTestResultListResponseItemsGoalSubtypeColumnStatistic            CommitTestResultListResponseItemsGoalSubtype = "columnStatistic"
	CommitTestResultListResponseItemsGoalSubtypeColumnValuesMatch          CommitTestResultListResponseItemsGoalSubtype = "columnValuesMatch"
	CommitTestResultListResponseItemsGoalSubtypeConflictingLabelRowCount   CommitTestResultListResponseItemsGoalSubtype = "conflictingLabelRowCount"
	CommitTestResultListResponseItemsGoalSubtypeContainsPii                CommitTestResultListResponseItemsGoalSubtype = "containsPii"
	CommitTestResultListResponseItemsGoalSubtypeContainsValidURL           CommitTestResultListResponseItemsGoalSubtype = "containsValidUrl"
	CommitTestResultListResponseItemsGoalSubtypeCorrelatedFeatureCount     CommitTestResultListResponseItemsGoalSubtype = "correlatedFeatureCount"
	CommitTestResultListResponseItemsGoalSubtypeCustomMetricThreshold      CommitTestResultListResponseItemsGoalSubtype = "customMetricThreshold"
	CommitTestResultListResponseItemsGoalSubtypeDuplicateRowCount          CommitTestResultListResponseItemsGoalSubtype = "duplicateRowCount"
	CommitTestResultListResponseItemsGoalSubtypeEmptyFeature               CommitTestResultListResponseItemsGoalSubtype = "emptyFeature"
	CommitTestResultListResponseItemsGoalSubtypeEmptyFeatureCount          CommitTestResultListResponseItemsGoalSubtype = "emptyFeatureCount"
	CommitTestResultListResponseItemsGoalSubtypeDriftedFeatureCount        CommitTestResultListResponseItemsGoalSubtype = "driftedFeatureCount"
	CommitTestResultListResponseItemsGoalSubtypeFeatureMissingValues       CommitTestResultListResponseItemsGoalSubtype = "featureMissingValues"
	CommitTestResultListResponseItemsGoalSubtypeFeatureValueValidation     CommitTestResultListResponseItemsGoalSubtype = "featureValueValidation"
	CommitTestResultListResponseItemsGoalSubtypeGreatExpectations          CommitTestResultListResponseItemsGoalSubtype = "greatExpectations"
	CommitTestResultListResponseItemsGoalSubtypeGroupByColumnStatsCheck    CommitTestResultListResponseItemsGoalSubtype = "groupByColumnStatsCheck"
	CommitTestResultListResponseItemsGoalSubtypeIllFormedRowCount          CommitTestResultListResponseItemsGoalSubtype = "illFormedRowCount"
	CommitTestResultListResponseItemsGoalSubtypeIsCode                     CommitTestResultListResponseItemsGoalSubtype = "isCode"
	CommitTestResultListResponseItemsGoalSubtypeIsJson                     CommitTestResultListResponseItemsGoalSubtype = "isJson"
	CommitTestResultListResponseItemsGoalSubtypeLlmRubricThresholdV2       CommitTestResultListResponseItemsGoalSubtype = "llmRubricThresholdV2"
	CommitTestResultListResponseItemsGoalSubtypeLabelDrift                 CommitTestResultListResponseItemsGoalSubtype = "labelDrift"
	CommitTestResultListResponseItemsGoalSubtypeMetricThreshold            CommitTestResultListResponseItemsGoalSubtype = "metricThreshold"
	CommitTestResultListResponseItemsGoalSubtypeNewCategoryCount           CommitTestResultListResponseItemsGoalSubtype = "newCategoryCount"
	CommitTestResultListResponseItemsGoalSubtypeNewLabelCount              CommitTestResultListResponseItemsGoalSubtype = "newLabelCount"
	CommitTestResultListResponseItemsGoalSubtypeNullRowCount               CommitTestResultListResponseItemsGoalSubtype = "nullRowCount"
	CommitTestResultListResponseItemsGoalSubtypeRowCount                   CommitTestResultListResponseItemsGoalSubtype = "rowCount"
	CommitTestResultListResponseItemsGoalSubtypePpScoreValueValidation     CommitTestResultListResponseItemsGoalSubtype = "ppScoreValueValidation"
	CommitTestResultListResponseItemsGoalSubtypeQuasiConstantFeature       CommitTestResultListResponseItemsGoalSubtype = "quasiConstantFeature"
	CommitTestResultListResponseItemsGoalSubtypeQuasiConstantFeatureCount  CommitTestResultListResponseItemsGoalSubtype = "quasiConstantFeatureCount"
	CommitTestResultListResponseItemsGoalSubtypeSqlQuery                   CommitTestResultListResponseItemsGoalSubtype = "sqlQuery"
	CommitTestResultListResponseItemsGoalSubtypeDtypeValidation            CommitTestResultListResponseItemsGoalSubtype = "dtypeValidation"
	CommitTestResultListResponseItemsGoalSubtypeSentenceLength             CommitTestResultListResponseItemsGoalSubtype = "sentenceLength"
	CommitTestResultListResponseItemsGoalSubtypeSizeRatio                  CommitTestResultListResponseItemsGoalSubtype = "sizeRatio"
	CommitTestResultListResponseItemsGoalSubtypeSpecialCharactersRatio     CommitTestResultListResponseItemsGoalSubtype = "specialCharactersRatio"
	CommitTestResultListResponseItemsGoalSubtypeStringValidation           CommitTestResultListResponseItemsGoalSubtype = "stringValidation"
	CommitTestResultListResponseItemsGoalSubtypeTrainValLeakageRowCount    CommitTestResultListResponseItemsGoalSubtype = "trainValLeakageRowCount"
)

func (r CommitTestResultListResponseItemsGoalSubtype) IsKnown() bool {
	switch r {
	case CommitTestResultListResponseItemsGoalSubtypeAnomalousColumnCount, CommitTestResultListResponseItemsGoalSubtypeCharacterLength, CommitTestResultListResponseItemsGoalSubtypeClassImbalanceRatio, CommitTestResultListResponseItemsGoalSubtypeExpectColumnAToBeInColumnB, CommitTestResultListResponseItemsGoalSubtypeColumnAverage, CommitTestResultListResponseItemsGoalSubtypeColumnDrift, CommitTestResultListResponseItemsGoalSubtypeColumnStatistic, CommitTestResultListResponseItemsGoalSubtypeColumnValuesMatch, CommitTestResultListResponseItemsGoalSubtypeConflictingLabelRowCount, CommitTestResultListResponseItemsGoalSubtypeContainsPii, CommitTestResultListResponseItemsGoalSubtypeContainsValidURL, CommitTestResultListResponseItemsGoalSubtypeCorrelatedFeatureCount, CommitTestResultListResponseItemsGoalSubtypeCustomMetricThreshold, CommitTestResultListResponseItemsGoalSubtypeDuplicateRowCount, CommitTestResultListResponseItemsGoalSubtypeEmptyFeature, CommitTestResultListResponseItemsGoalSubtypeEmptyFeatureCount, CommitTestResultListResponseItemsGoalSubtypeDriftedFeatureCount, CommitTestResultListResponseItemsGoalSubtypeFeatureMissingValues, CommitTestResultListResponseItemsGoalSubtypeFeatureValueValidation, CommitTestResultListResponseItemsGoalSubtypeGreatExpectations, CommitTestResultListResponseItemsGoalSubtypeGroupByColumnStatsCheck, CommitTestResultListResponseItemsGoalSubtypeIllFormedRowCount, CommitTestResultListResponseItemsGoalSubtypeIsCode, CommitTestResultListResponseItemsGoalSubtypeIsJson, CommitTestResultListResponseItemsGoalSubtypeLlmRubricThresholdV2, CommitTestResultListResponseItemsGoalSubtypeLabelDrift, CommitTestResultListResponseItemsGoalSubtypeMetricThreshold, CommitTestResultListResponseItemsGoalSubtypeNewCategoryCount, CommitTestResultListResponseItemsGoalSubtypeNewLabelCount, CommitTestResultListResponseItemsGoalSubtypeNullRowCount, CommitTestResultListResponseItemsGoalSubtypeRowCount, CommitTestResultListResponseItemsGoalSubtypePpScoreValueValidation, CommitTestResultListResponseItemsGoalSubtypeQuasiConstantFeature, CommitTestResultListResponseItemsGoalSubtypeQuasiConstantFeatureCount, CommitTestResultListResponseItemsGoalSubtypeSqlQuery, CommitTestResultListResponseItemsGoalSubtypeDtypeValidation, CommitTestResultListResponseItemsGoalSubtypeSentenceLength, CommitTestResultListResponseItemsGoalSubtypeSizeRatio, CommitTestResultListResponseItemsGoalSubtypeSpecialCharactersRatio, CommitTestResultListResponseItemsGoalSubtypeStringValidation, CommitTestResultListResponseItemsGoalSubtypeTrainValLeakageRowCount:
		return true
	}
	return false
}

type CommitTestResultListResponseItemsGoalThreshold struct {
	// The insight name to be evaluated.
	InsightName CommitTestResultListResponseItemsGoalThresholdsInsightName `json:"insightName"`
	// The insight parameters. Required only for some test subtypes. For example, for
	// tests that require a column name, the insight parameters will be [{'name':
	// 'column_name', 'value': 'Age'}]
	InsightParameters []CommitTestResultListResponseItemsGoalThresholdsInsightParameter `json:"insightParameters,nullable"`
	// The measurement to be evaluated.
	Measurement string `json:"measurement"`
	// The operator to be used for the evaluation.
	Operator CommitTestResultListResponseItemsGoalThresholdsOperator `json:"operator"`
	// Whether to use automatic anomaly detection or manual thresholds
	ThresholdMode CommitTestResultListResponseItemsGoalThresholdsThresholdMode `json:"thresholdMode"`
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
	ThresholdMode     apijson.Field
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

// The insight name to be evaluated.
type CommitTestResultListResponseItemsGoalThresholdsInsightName string

const (
	CommitTestResultListResponseItemsGoalThresholdsInsightNameCharacterLength            CommitTestResultListResponseItemsGoalThresholdsInsightName = "characterLength"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameClassImbalance             CommitTestResultListResponseItemsGoalThresholdsInsightName = "classImbalance"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameExpectColumnAToBeInColumnB CommitTestResultListResponseItemsGoalThresholdsInsightName = "expectColumnAToBeInColumnB"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameColumnAverage              CommitTestResultListResponseItemsGoalThresholdsInsightName = "columnAverage"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameColumnDrift                CommitTestResultListResponseItemsGoalThresholdsInsightName = "columnDrift"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameColumnValuesMatch          CommitTestResultListResponseItemsGoalThresholdsInsightName = "columnValuesMatch"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameConfidenceDistribution     CommitTestResultListResponseItemsGoalThresholdsInsightName = "confidenceDistribution"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameConflictingLabelRowCount   CommitTestResultListResponseItemsGoalThresholdsInsightName = "conflictingLabelRowCount"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameContainsPii                CommitTestResultListResponseItemsGoalThresholdsInsightName = "containsPii"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameContainsValidURL           CommitTestResultListResponseItemsGoalThresholdsInsightName = "containsValidUrl"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameCorrelatedFeatures         CommitTestResultListResponseItemsGoalThresholdsInsightName = "correlatedFeatures"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameCustomMetric               CommitTestResultListResponseItemsGoalThresholdsInsightName = "customMetric"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameDuplicateRowCount          CommitTestResultListResponseItemsGoalThresholdsInsightName = "duplicateRowCount"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameEmptyFeatures              CommitTestResultListResponseItemsGoalThresholdsInsightName = "emptyFeatures"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameFeatureDrift               CommitTestResultListResponseItemsGoalThresholdsInsightName = "featureDrift"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameFeatureProfile             CommitTestResultListResponseItemsGoalThresholdsInsightName = "featureProfile"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameGreatExpectations          CommitTestResultListResponseItemsGoalThresholdsInsightName = "greatExpectations"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameGroupByColumnStatsCheck    CommitTestResultListResponseItemsGoalThresholdsInsightName = "groupByColumnStatsCheck"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameIllFormedRowCount          CommitTestResultListResponseItemsGoalThresholdsInsightName = "illFormedRowCount"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameIsCode                     CommitTestResultListResponseItemsGoalThresholdsInsightName = "isCode"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameIsJson                     CommitTestResultListResponseItemsGoalThresholdsInsightName = "isJson"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameLlmRubricV2                CommitTestResultListResponseItemsGoalThresholdsInsightName = "llmRubricV2"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameLabelDrift                 CommitTestResultListResponseItemsGoalThresholdsInsightName = "labelDrift"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameMetrics                    CommitTestResultListResponseItemsGoalThresholdsInsightName = "metrics"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameNewCategories              CommitTestResultListResponseItemsGoalThresholdsInsightName = "newCategories"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameNewLabels                  CommitTestResultListResponseItemsGoalThresholdsInsightName = "newLabels"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameNullRowCount               CommitTestResultListResponseItemsGoalThresholdsInsightName = "nullRowCount"
	CommitTestResultListResponseItemsGoalThresholdsInsightNamePpScore                    CommitTestResultListResponseItemsGoalThresholdsInsightName = "ppScore"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameQuasiConstantFeatures      CommitTestResultListResponseItemsGoalThresholdsInsightName = "quasiConstantFeatures"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameSentenceLength             CommitTestResultListResponseItemsGoalThresholdsInsightName = "sentenceLength"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameSizeRatio                  CommitTestResultListResponseItemsGoalThresholdsInsightName = "sizeRatio"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameSpecialCharacters          CommitTestResultListResponseItemsGoalThresholdsInsightName = "specialCharacters"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameStringValidation           CommitTestResultListResponseItemsGoalThresholdsInsightName = "stringValidation"
	CommitTestResultListResponseItemsGoalThresholdsInsightNameTrainValLeakageRowCount    CommitTestResultListResponseItemsGoalThresholdsInsightName = "trainValLeakageRowCount"
)

func (r CommitTestResultListResponseItemsGoalThresholdsInsightName) IsKnown() bool {
	switch r {
	case CommitTestResultListResponseItemsGoalThresholdsInsightNameCharacterLength, CommitTestResultListResponseItemsGoalThresholdsInsightNameClassImbalance, CommitTestResultListResponseItemsGoalThresholdsInsightNameExpectColumnAToBeInColumnB, CommitTestResultListResponseItemsGoalThresholdsInsightNameColumnAverage, CommitTestResultListResponseItemsGoalThresholdsInsightNameColumnDrift, CommitTestResultListResponseItemsGoalThresholdsInsightNameColumnValuesMatch, CommitTestResultListResponseItemsGoalThresholdsInsightNameConfidenceDistribution, CommitTestResultListResponseItemsGoalThresholdsInsightNameConflictingLabelRowCount, CommitTestResultListResponseItemsGoalThresholdsInsightNameContainsPii, CommitTestResultListResponseItemsGoalThresholdsInsightNameContainsValidURL, CommitTestResultListResponseItemsGoalThresholdsInsightNameCorrelatedFeatures, CommitTestResultListResponseItemsGoalThresholdsInsightNameCustomMetric, CommitTestResultListResponseItemsGoalThresholdsInsightNameDuplicateRowCount, CommitTestResultListResponseItemsGoalThresholdsInsightNameEmptyFeatures, CommitTestResultListResponseItemsGoalThresholdsInsightNameFeatureDrift, CommitTestResultListResponseItemsGoalThresholdsInsightNameFeatureProfile, CommitTestResultListResponseItemsGoalThresholdsInsightNameGreatExpectations, CommitTestResultListResponseItemsGoalThresholdsInsightNameGroupByColumnStatsCheck, CommitTestResultListResponseItemsGoalThresholdsInsightNameIllFormedRowCount, CommitTestResultListResponseItemsGoalThresholdsInsightNameIsCode, CommitTestResultListResponseItemsGoalThresholdsInsightNameIsJson, CommitTestResultListResponseItemsGoalThresholdsInsightNameLlmRubricV2, CommitTestResultListResponseItemsGoalThresholdsInsightNameLabelDrift, CommitTestResultListResponseItemsGoalThresholdsInsightNameMetrics, CommitTestResultListResponseItemsGoalThresholdsInsightNameNewCategories, CommitTestResultListResponseItemsGoalThresholdsInsightNameNewLabels, CommitTestResultListResponseItemsGoalThresholdsInsightNameNullRowCount, CommitTestResultListResponseItemsGoalThresholdsInsightNamePpScore, CommitTestResultListResponseItemsGoalThresholdsInsightNameQuasiConstantFeatures, CommitTestResultListResponseItemsGoalThresholdsInsightNameSentenceLength, CommitTestResultListResponseItemsGoalThresholdsInsightNameSizeRatio, CommitTestResultListResponseItemsGoalThresholdsInsightNameSpecialCharacters, CommitTestResultListResponseItemsGoalThresholdsInsightNameStringValidation, CommitTestResultListResponseItemsGoalThresholdsInsightNameTrainValLeakageRowCount:
		return true
	}
	return false
}

type CommitTestResultListResponseItemsGoalThresholdsInsightParameter struct {
	// The name of the insight filter.
	Name  string                                                              `json:"name,required"`
	Value interface{}                                                         `json:"value,required"`
	JSON  commitTestResultListResponseItemsGoalThresholdsInsightParameterJSON `json:"-"`
}

// commitTestResultListResponseItemsGoalThresholdsInsightParameterJSON contains the
// JSON metadata for the struct
// [CommitTestResultListResponseItemsGoalThresholdsInsightParameter]
type commitTestResultListResponseItemsGoalThresholdsInsightParameterJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitTestResultListResponseItemsGoalThresholdsInsightParameter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitTestResultListResponseItemsGoalThresholdsInsightParameterJSON) RawJSON() string {
	return r.raw
}

// The operator to be used for the evaluation.
type CommitTestResultListResponseItemsGoalThresholdsOperator string

const (
	CommitTestResultListResponseItemsGoalThresholdsOperatorIs              CommitTestResultListResponseItemsGoalThresholdsOperator = "is"
	CommitTestResultListResponseItemsGoalThresholdsOperatorGreater         CommitTestResultListResponseItemsGoalThresholdsOperator = ">"
	CommitTestResultListResponseItemsGoalThresholdsOperatorGreaterOrEquals CommitTestResultListResponseItemsGoalThresholdsOperator = ">="
	CommitTestResultListResponseItemsGoalThresholdsOperatorLess            CommitTestResultListResponseItemsGoalThresholdsOperator = "<"
	CommitTestResultListResponseItemsGoalThresholdsOperatorLessOrEquals    CommitTestResultListResponseItemsGoalThresholdsOperator = "<="
	CommitTestResultListResponseItemsGoalThresholdsOperatorNotEquals       CommitTestResultListResponseItemsGoalThresholdsOperator = "!="
)

func (r CommitTestResultListResponseItemsGoalThresholdsOperator) IsKnown() bool {
	switch r {
	case CommitTestResultListResponseItemsGoalThresholdsOperatorIs, CommitTestResultListResponseItemsGoalThresholdsOperatorGreater, CommitTestResultListResponseItemsGoalThresholdsOperatorGreaterOrEquals, CommitTestResultListResponseItemsGoalThresholdsOperatorLess, CommitTestResultListResponseItemsGoalThresholdsOperatorLessOrEquals, CommitTestResultListResponseItemsGoalThresholdsOperatorNotEquals:
		return true
	}
	return false
}

// Whether to use automatic anomaly detection or manual thresholds
type CommitTestResultListResponseItemsGoalThresholdsThresholdMode string

const (
	CommitTestResultListResponseItemsGoalThresholdsThresholdModeAutomatic CommitTestResultListResponseItemsGoalThresholdsThresholdMode = "automatic"
	CommitTestResultListResponseItemsGoalThresholdsThresholdModeManual    CommitTestResultListResponseItemsGoalThresholdsThresholdMode = "manual"
)

func (r CommitTestResultListResponseItemsGoalThresholdsThresholdMode) IsKnown() bool {
	switch r {
	case CommitTestResultListResponseItemsGoalThresholdsThresholdModeAutomatic, CommitTestResultListResponseItemsGoalThresholdsThresholdModeManual:
		return true
	}
	return false
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

// The test type.
type CommitTestResultListResponseItemsGoalType string

const (
	CommitTestResultListResponseItemsGoalTypeIntegrity   CommitTestResultListResponseItemsGoalType = "integrity"
	CommitTestResultListResponseItemsGoalTypeConsistency CommitTestResultListResponseItemsGoalType = "consistency"
	CommitTestResultListResponseItemsGoalTypePerformance CommitTestResultListResponseItemsGoalType = "performance"
)

func (r CommitTestResultListResponseItemsGoalType) IsKnown() bool {
	switch r {
	case CommitTestResultListResponseItemsGoalTypeIntegrity, CommitTestResultListResponseItemsGoalTypeConsistency, CommitTestResultListResponseItemsGoalTypePerformance:
		return true
	}
	return false
}

// The body of the rows request.
type CommitTestResultListResponseItemsRowsBody struct {
	ColumnFilters     []CommitTestResultListResponseItemsRowsBodyColumnFilter `json:"columnFilters,nullable"`
	ExcludeRowIDList  []int64                                                 `json:"excludeRowIdList,nullable"`
	NotSearchQueryAnd []string                                                `json:"notSearchQueryAnd,nullable"`
	NotSearchQueryOr  []string                                                `json:"notSearchQueryOr,nullable"`
	RowIDList         []int64                                                 `json:"rowIdList,nullable"`
	SearchQueryAnd    []string                                                `json:"searchQueryAnd,nullable"`
	SearchQueryOr     []string                                                `json:"searchQueryOr,nullable"`
	JSON              commitTestResultListResponseItemsRowsBodyJSON           `json:"-"`
}

// commitTestResultListResponseItemsRowsBodyJSON contains the JSON metadata for the
// struct [CommitTestResultListResponseItemsRowsBody]
type commitTestResultListResponseItemsRowsBodyJSON struct {
	ColumnFilters     apijson.Field
	ExcludeRowIDList  apijson.Field
	NotSearchQueryAnd apijson.Field
	NotSearchQueryOr  apijson.Field
	RowIDList         apijson.Field
	SearchQueryAnd    apijson.Field
	SearchQueryOr     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CommitTestResultListResponseItemsRowsBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitTestResultListResponseItemsRowsBodyJSON) RawJSON() string {
	return r.raw
}

type CommitTestResultListResponseItemsRowsBodyColumnFilter struct {
	// The name of the column.
	Measurement string                                                         `json:"measurement,required"`
	Operator    CommitTestResultListResponseItemsRowsBodyColumnFiltersOperator `json:"operator,required"`
	// This field can have the runtime type of
	// [[]CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterValueUnion],
	// [float64],
	// [CommitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilterValueUnion].
	Value interface{}                                               `json:"value,required"`
	JSON  commitTestResultListResponseItemsRowsBodyColumnFilterJSON `json:"-"`
	union CommitTestResultListResponseItemsRowsBodyColumnFiltersUnion
}

// commitTestResultListResponseItemsRowsBodyColumnFilterJSON contains the JSON
// metadata for the struct [CommitTestResultListResponseItemsRowsBodyColumnFilter]
type commitTestResultListResponseItemsRowsBodyColumnFilterJSON struct {
	Measurement apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r commitTestResultListResponseItemsRowsBodyColumnFilterJSON) RawJSON() string {
	return r.raw
}

func (r *CommitTestResultListResponseItemsRowsBodyColumnFilter) UnmarshalJSON(data []byte) (err error) {
	*r = CommitTestResultListResponseItemsRowsBodyColumnFilter{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CommitTestResultListResponseItemsRowsBodyColumnFiltersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilter],
// [CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilter],
// [CommitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilter].
func (r CommitTestResultListResponseItemsRowsBodyColumnFilter) AsUnion() CommitTestResultListResponseItemsRowsBodyColumnFiltersUnion {
	return r.union
}

// Union satisfied by
// [CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilter],
// [CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilter] or
// [CommitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilter].
type CommitTestResultListResponseItemsRowsBodyColumnFiltersUnion interface {
	implementsCommitTestResultListResponseItemsRowsBodyColumnFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CommitTestResultListResponseItemsRowsBodyColumnFiltersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilter{}),
		},
	)
}

type CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilter struct {
	// The name of the column.
	Measurement string                                                                            `json:"measurement,required"`
	Operator    CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterOperator     `json:"operator,required"`
	Value       []CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterValueUnion `json:"value,required"`
	JSON        commitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterJSON         `json:"-"`
}

// commitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterJSON
// contains the JSON metadata for the struct
// [CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilter]
type commitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterJSON struct {
	Measurement apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterJSON) RawJSON() string {
	return r.raw
}

func (r CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilter) implementsCommitTestResultListResponseItemsRowsBodyColumnFilter() {
}

type CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterOperator string

const (
	CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterOperatorContainsNone CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterOperator = "contains_none"
	CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterOperatorContainsAny  CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterOperator = "contains_any"
	CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterOperatorContainsAll  CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterOperator = "contains_all"
	CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterOperatorOneOf        CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterOperator = "one_of"
	CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterOperatorNoneOf       CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterOperator = "none_of"
)

func (r CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterOperator) IsKnown() bool {
	switch r {
	case CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterOperatorContainsNone, CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterOperatorContainsAny, CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterOperatorContainsAll, CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterOperatorOneOf, CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterOperatorNoneOf:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterValueUnion interface {
	ImplementsCommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CommitTestResultListResponseItemsRowsBodyColumnFiltersSetColumnFilterValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilter struct {
	// The name of the column.
	Measurement string                                                                            `json:"measurement,required"`
	Operator    CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperator `json:"operator,required"`
	Value       float64                                                                           `json:"value,required,nullable"`
	JSON        commitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterJSON     `json:"-"`
}

// commitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterJSON
// contains the JSON metadata for the struct
// [CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilter]
type commitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterJSON struct {
	Measurement apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterJSON) RawJSON() string {
	return r.raw
}

func (r CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilter) implementsCommitTestResultListResponseItemsRowsBodyColumnFilter() {
}

type CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperator string

const (
	CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorGreater         CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperator = ">"
	CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorGreaterOrEquals CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperator = ">="
	CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorIs              CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperator = "is"
	CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorLess            CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperator = "<"
	CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorLessOrEquals    CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperator = "<="
	CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorNotEquals       CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperator = "!="
)

func (r CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperator) IsKnown() bool {
	switch r {
	case CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorGreater, CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorGreaterOrEquals, CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorIs, CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorLess, CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorLessOrEquals, CommitTestResultListResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorNotEquals:
		return true
	}
	return false
}

type CommitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilter struct {
	// The name of the column.
	Measurement string                                                                             `json:"measurement,required"`
	Operator    CommitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilterOperator   `json:"operator,required"`
	Value       CommitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilterValueUnion `json:"value,required"`
	JSON        commitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilterJSON       `json:"-"`
}

// commitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilterJSON
// contains the JSON metadata for the struct
// [CommitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilter]
type commitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilterJSON struct {
	Measurement apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilterJSON) RawJSON() string {
	return r.raw
}

func (r CommitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilter) implementsCommitTestResultListResponseItemsRowsBodyColumnFilter() {
}

type CommitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilterOperator string

const (
	CommitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilterOperatorIs        CommitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilterOperator = "is"
	CommitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilterOperatorNotEquals CommitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilterOperator = "!="
)

func (r CommitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilterOperator) IsKnown() bool {
	switch r {
	case CommitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilterOperatorIs, CommitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilterOperatorNotEquals:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type CommitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilterValueUnion interface {
	ImplementsCommitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilterValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CommitTestResultListResponseItemsRowsBodyColumnFiltersStringColumnFilterValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type CommitTestResultListResponseItemsRowsBodyColumnFiltersOperator string

const (
	CommitTestResultListResponseItemsRowsBodyColumnFiltersOperatorContainsNone    CommitTestResultListResponseItemsRowsBodyColumnFiltersOperator = "contains_none"
	CommitTestResultListResponseItemsRowsBodyColumnFiltersOperatorContainsAny     CommitTestResultListResponseItemsRowsBodyColumnFiltersOperator = "contains_any"
	CommitTestResultListResponseItemsRowsBodyColumnFiltersOperatorContainsAll     CommitTestResultListResponseItemsRowsBodyColumnFiltersOperator = "contains_all"
	CommitTestResultListResponseItemsRowsBodyColumnFiltersOperatorOneOf           CommitTestResultListResponseItemsRowsBodyColumnFiltersOperator = "one_of"
	CommitTestResultListResponseItemsRowsBodyColumnFiltersOperatorNoneOf          CommitTestResultListResponseItemsRowsBodyColumnFiltersOperator = "none_of"
	CommitTestResultListResponseItemsRowsBodyColumnFiltersOperatorGreater         CommitTestResultListResponseItemsRowsBodyColumnFiltersOperator = ">"
	CommitTestResultListResponseItemsRowsBodyColumnFiltersOperatorGreaterOrEquals CommitTestResultListResponseItemsRowsBodyColumnFiltersOperator = ">="
	CommitTestResultListResponseItemsRowsBodyColumnFiltersOperatorIs              CommitTestResultListResponseItemsRowsBodyColumnFiltersOperator = "is"
	CommitTestResultListResponseItemsRowsBodyColumnFiltersOperatorLess            CommitTestResultListResponseItemsRowsBodyColumnFiltersOperator = "<"
	CommitTestResultListResponseItemsRowsBodyColumnFiltersOperatorLessOrEquals    CommitTestResultListResponseItemsRowsBodyColumnFiltersOperator = "<="
	CommitTestResultListResponseItemsRowsBodyColumnFiltersOperatorNotEquals       CommitTestResultListResponseItemsRowsBodyColumnFiltersOperator = "!="
)

func (r CommitTestResultListResponseItemsRowsBodyColumnFiltersOperator) IsKnown() bool {
	switch r {
	case CommitTestResultListResponseItemsRowsBodyColumnFiltersOperatorContainsNone, CommitTestResultListResponseItemsRowsBodyColumnFiltersOperatorContainsAny, CommitTestResultListResponseItemsRowsBodyColumnFiltersOperatorContainsAll, CommitTestResultListResponseItemsRowsBodyColumnFiltersOperatorOneOf, CommitTestResultListResponseItemsRowsBodyColumnFiltersOperatorNoneOf, CommitTestResultListResponseItemsRowsBodyColumnFiltersOperatorGreater, CommitTestResultListResponseItemsRowsBodyColumnFiltersOperatorGreaterOrEquals, CommitTestResultListResponseItemsRowsBodyColumnFiltersOperatorIs, CommitTestResultListResponseItemsRowsBodyColumnFiltersOperatorLess, CommitTestResultListResponseItemsRowsBodyColumnFiltersOperatorLessOrEquals, CommitTestResultListResponseItemsRowsBodyColumnFiltersOperatorNotEquals:
		return true
	}
	return false
}

type CommitTestResultListParams struct {
	// Filter for archived tests.
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
