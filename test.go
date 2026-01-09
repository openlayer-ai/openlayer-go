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

// List the test results for a test.
func (r *TestService) ListResults(ctx context.Context, testID string, query TestListResultsParams, opts ...option.RequestOption) (res *TestListResultsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if testID == "" {
		err = errors.New("missing required testId parameter")
		return
	}
	path := fmt.Sprintf("tests/%s/results", testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
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

type TestListResultsResponse struct {
	Items               []TestListResultsResponseItem              `json:"items,required"`
	LastUnskippedResult TestListResultsResponseLastUnskippedResult `json:"lastUnskippedResult,nullable"`
	JSON                testListResultsResponseJSON                `json:"-"`
}

// testListResultsResponseJSON contains the JSON metadata for the struct
// [TestListResultsResponse]
type testListResultsResponseJSON struct {
	Items               apijson.Field
	LastUnskippedResult apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *TestListResultsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResultsResponseJSON) RawJSON() string {
	return r.raw
}

type TestListResultsResponseItem struct {
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
	Status TestListResultsResponseItemsStatus `json:"status,required"`
	// The status message.
	StatusMessage  string                                      `json:"statusMessage,required,nullable"`
	ExpectedValues []TestListResultsResponseItemsExpectedValue `json:"expectedValues"`
	Goal           TestListResultsResponseItemsGoal            `json:"goal"`
	// The test id.
	GoalID string `json:"goalId,nullable" format:"uuid"`
	// The URL to the rows of the test result.
	Rows string `json:"rows"`
	// The body of the rows request.
	RowsBody TestListResultsResponseItemsRowsBody `json:"rowsBody,nullable"`
	JSON     testListResultsResponseItemJSON      `json:"-"`
}

// testListResultsResponseItemJSON contains the JSON metadata for the struct
// [TestListResultsResponseItem]
type testListResultsResponseItemJSON struct {
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

func (r *TestListResultsResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResultsResponseItemJSON) RawJSON() string {
	return r.raw
}

// The status of the test.
type TestListResultsResponseItemsStatus string

const (
	TestListResultsResponseItemsStatusRunning TestListResultsResponseItemsStatus = "running"
	TestListResultsResponseItemsStatusPassing TestListResultsResponseItemsStatus = "passing"
	TestListResultsResponseItemsStatusFailing TestListResultsResponseItemsStatus = "failing"
	TestListResultsResponseItemsStatusSkipped TestListResultsResponseItemsStatus = "skipped"
	TestListResultsResponseItemsStatusError   TestListResultsResponseItemsStatus = "error"
)

func (r TestListResultsResponseItemsStatus) IsKnown() bool {
	switch r {
	case TestListResultsResponseItemsStatusRunning, TestListResultsResponseItemsStatusPassing, TestListResultsResponseItemsStatusFailing, TestListResultsResponseItemsStatusSkipped, TestListResultsResponseItemsStatusError:
		return true
	}
	return false
}

type TestListResultsResponseItemsExpectedValue struct {
	// the lower threshold for the expected value
	LowerThreshold float64 `json:"lowerThreshold,nullable"`
	// One of the `measurement` values in the test's thresholds
	Measurement string `json:"measurement"`
	// The upper threshold for the expected value
	UpperThreshold float64                                       `json:"upperThreshold,nullable"`
	JSON           testListResultsResponseItemsExpectedValueJSON `json:"-"`
}

// testListResultsResponseItemsExpectedValueJSON contains the JSON metadata for the
// struct [TestListResultsResponseItemsExpectedValue]
type testListResultsResponseItemsExpectedValueJSON struct {
	LowerThreshold apijson.Field
	Measurement    apijson.Field
	UpperThreshold apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TestListResultsResponseItemsExpectedValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResultsResponseItemsExpectedValueJSON) RawJSON() string {
	return r.raw
}

type TestListResultsResponseItemsGoal struct {
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
	Subtype TestListResultsResponseItemsGoalSubtype `json:"subtype,required"`
	// Whether the test is suggested or user-created.
	Suggested  bool                                        `json:"suggested,required"`
	Thresholds []TestListResultsResponseItemsGoalThreshold `json:"thresholds,required"`
	// The test type.
	Type TestListResultsResponseItemsGoalType `json:"type,required"`
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
	UsesValidationDataset bool                                 `json:"usesValidationDataset"`
	JSON                  testListResultsResponseItemsGoalJSON `json:"-"`
}

// testListResultsResponseItemsGoalJSON contains the JSON metadata for the struct
// [TestListResultsResponseItemsGoal]
type testListResultsResponseItemsGoalJSON struct {
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

func (r *TestListResultsResponseItemsGoal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResultsResponseItemsGoalJSON) RawJSON() string {
	return r.raw
}

// The test subtype.
type TestListResultsResponseItemsGoalSubtype string

const (
	TestListResultsResponseItemsGoalSubtypeAnomalousColumnCount       TestListResultsResponseItemsGoalSubtype = "anomalousColumnCount"
	TestListResultsResponseItemsGoalSubtypeCharacterLength            TestListResultsResponseItemsGoalSubtype = "characterLength"
	TestListResultsResponseItemsGoalSubtypeClassImbalanceRatio        TestListResultsResponseItemsGoalSubtype = "classImbalanceRatio"
	TestListResultsResponseItemsGoalSubtypeExpectColumnAToBeInColumnB TestListResultsResponseItemsGoalSubtype = "expectColumnAToBeInColumnB"
	TestListResultsResponseItemsGoalSubtypeColumnAverage              TestListResultsResponseItemsGoalSubtype = "columnAverage"
	TestListResultsResponseItemsGoalSubtypeColumnDrift                TestListResultsResponseItemsGoalSubtype = "columnDrift"
	TestListResultsResponseItemsGoalSubtypeColumnStatistic            TestListResultsResponseItemsGoalSubtype = "columnStatistic"
	TestListResultsResponseItemsGoalSubtypeColumnValuesMatch          TestListResultsResponseItemsGoalSubtype = "columnValuesMatch"
	TestListResultsResponseItemsGoalSubtypeConflictingLabelRowCount   TestListResultsResponseItemsGoalSubtype = "conflictingLabelRowCount"
	TestListResultsResponseItemsGoalSubtypeContainsPii                TestListResultsResponseItemsGoalSubtype = "containsPii"
	TestListResultsResponseItemsGoalSubtypeContainsValidURL           TestListResultsResponseItemsGoalSubtype = "containsValidUrl"
	TestListResultsResponseItemsGoalSubtypeCorrelatedFeatureCount     TestListResultsResponseItemsGoalSubtype = "correlatedFeatureCount"
	TestListResultsResponseItemsGoalSubtypeCustomMetricThreshold      TestListResultsResponseItemsGoalSubtype = "customMetricThreshold"
	TestListResultsResponseItemsGoalSubtypeDuplicateRowCount          TestListResultsResponseItemsGoalSubtype = "duplicateRowCount"
	TestListResultsResponseItemsGoalSubtypeEmptyFeature               TestListResultsResponseItemsGoalSubtype = "emptyFeature"
	TestListResultsResponseItemsGoalSubtypeEmptyFeatureCount          TestListResultsResponseItemsGoalSubtype = "emptyFeatureCount"
	TestListResultsResponseItemsGoalSubtypeDriftedFeatureCount        TestListResultsResponseItemsGoalSubtype = "driftedFeatureCount"
	TestListResultsResponseItemsGoalSubtypeFeatureMissingValues       TestListResultsResponseItemsGoalSubtype = "featureMissingValues"
	TestListResultsResponseItemsGoalSubtypeFeatureValueValidation     TestListResultsResponseItemsGoalSubtype = "featureValueValidation"
	TestListResultsResponseItemsGoalSubtypeGreatExpectations          TestListResultsResponseItemsGoalSubtype = "greatExpectations"
	TestListResultsResponseItemsGoalSubtypeGroupByColumnStatsCheck    TestListResultsResponseItemsGoalSubtype = "groupByColumnStatsCheck"
	TestListResultsResponseItemsGoalSubtypeIllFormedRowCount          TestListResultsResponseItemsGoalSubtype = "illFormedRowCount"
	TestListResultsResponseItemsGoalSubtypeIsCode                     TestListResultsResponseItemsGoalSubtype = "isCode"
	TestListResultsResponseItemsGoalSubtypeIsJson                     TestListResultsResponseItemsGoalSubtype = "isJson"
	TestListResultsResponseItemsGoalSubtypeLlmRubricThresholdV2       TestListResultsResponseItemsGoalSubtype = "llmRubricThresholdV2"
	TestListResultsResponseItemsGoalSubtypeLabelDrift                 TestListResultsResponseItemsGoalSubtype = "labelDrift"
	TestListResultsResponseItemsGoalSubtypeMetricThreshold            TestListResultsResponseItemsGoalSubtype = "metricThreshold"
	TestListResultsResponseItemsGoalSubtypeNewCategoryCount           TestListResultsResponseItemsGoalSubtype = "newCategoryCount"
	TestListResultsResponseItemsGoalSubtypeNewLabelCount              TestListResultsResponseItemsGoalSubtype = "newLabelCount"
	TestListResultsResponseItemsGoalSubtypeNullRowCount               TestListResultsResponseItemsGoalSubtype = "nullRowCount"
	TestListResultsResponseItemsGoalSubtypeRowCount                   TestListResultsResponseItemsGoalSubtype = "rowCount"
	TestListResultsResponseItemsGoalSubtypePpScoreValueValidation     TestListResultsResponseItemsGoalSubtype = "ppScoreValueValidation"
	TestListResultsResponseItemsGoalSubtypeQuasiConstantFeature       TestListResultsResponseItemsGoalSubtype = "quasiConstantFeature"
	TestListResultsResponseItemsGoalSubtypeQuasiConstantFeatureCount  TestListResultsResponseItemsGoalSubtype = "quasiConstantFeatureCount"
	TestListResultsResponseItemsGoalSubtypeSqlQuery                   TestListResultsResponseItemsGoalSubtype = "sqlQuery"
	TestListResultsResponseItemsGoalSubtypeDtypeValidation            TestListResultsResponseItemsGoalSubtype = "dtypeValidation"
	TestListResultsResponseItemsGoalSubtypeSentenceLength             TestListResultsResponseItemsGoalSubtype = "sentenceLength"
	TestListResultsResponseItemsGoalSubtypeSizeRatio                  TestListResultsResponseItemsGoalSubtype = "sizeRatio"
	TestListResultsResponseItemsGoalSubtypeSpecialCharactersRatio     TestListResultsResponseItemsGoalSubtype = "specialCharactersRatio"
	TestListResultsResponseItemsGoalSubtypeStringValidation           TestListResultsResponseItemsGoalSubtype = "stringValidation"
	TestListResultsResponseItemsGoalSubtypeTrainValLeakageRowCount    TestListResultsResponseItemsGoalSubtype = "trainValLeakageRowCount"
)

func (r TestListResultsResponseItemsGoalSubtype) IsKnown() bool {
	switch r {
	case TestListResultsResponseItemsGoalSubtypeAnomalousColumnCount, TestListResultsResponseItemsGoalSubtypeCharacterLength, TestListResultsResponseItemsGoalSubtypeClassImbalanceRatio, TestListResultsResponseItemsGoalSubtypeExpectColumnAToBeInColumnB, TestListResultsResponseItemsGoalSubtypeColumnAverage, TestListResultsResponseItemsGoalSubtypeColumnDrift, TestListResultsResponseItemsGoalSubtypeColumnStatistic, TestListResultsResponseItemsGoalSubtypeColumnValuesMatch, TestListResultsResponseItemsGoalSubtypeConflictingLabelRowCount, TestListResultsResponseItemsGoalSubtypeContainsPii, TestListResultsResponseItemsGoalSubtypeContainsValidURL, TestListResultsResponseItemsGoalSubtypeCorrelatedFeatureCount, TestListResultsResponseItemsGoalSubtypeCustomMetricThreshold, TestListResultsResponseItemsGoalSubtypeDuplicateRowCount, TestListResultsResponseItemsGoalSubtypeEmptyFeature, TestListResultsResponseItemsGoalSubtypeEmptyFeatureCount, TestListResultsResponseItemsGoalSubtypeDriftedFeatureCount, TestListResultsResponseItemsGoalSubtypeFeatureMissingValues, TestListResultsResponseItemsGoalSubtypeFeatureValueValidation, TestListResultsResponseItemsGoalSubtypeGreatExpectations, TestListResultsResponseItemsGoalSubtypeGroupByColumnStatsCheck, TestListResultsResponseItemsGoalSubtypeIllFormedRowCount, TestListResultsResponseItemsGoalSubtypeIsCode, TestListResultsResponseItemsGoalSubtypeIsJson, TestListResultsResponseItemsGoalSubtypeLlmRubricThresholdV2, TestListResultsResponseItemsGoalSubtypeLabelDrift, TestListResultsResponseItemsGoalSubtypeMetricThreshold, TestListResultsResponseItemsGoalSubtypeNewCategoryCount, TestListResultsResponseItemsGoalSubtypeNewLabelCount, TestListResultsResponseItemsGoalSubtypeNullRowCount, TestListResultsResponseItemsGoalSubtypeRowCount, TestListResultsResponseItemsGoalSubtypePpScoreValueValidation, TestListResultsResponseItemsGoalSubtypeQuasiConstantFeature, TestListResultsResponseItemsGoalSubtypeQuasiConstantFeatureCount, TestListResultsResponseItemsGoalSubtypeSqlQuery, TestListResultsResponseItemsGoalSubtypeDtypeValidation, TestListResultsResponseItemsGoalSubtypeSentenceLength, TestListResultsResponseItemsGoalSubtypeSizeRatio, TestListResultsResponseItemsGoalSubtypeSpecialCharactersRatio, TestListResultsResponseItemsGoalSubtypeStringValidation, TestListResultsResponseItemsGoalSubtypeTrainValLeakageRowCount:
		return true
	}
	return false
}

type TestListResultsResponseItemsGoalThreshold struct {
	// The insight name to be evaluated.
	InsightName TestListResultsResponseItemsGoalThresholdsInsightName `json:"insightName"`
	// The insight parameters. Required only for some test subtypes. For example, for
	// tests that require a column name, the insight parameters will be [{'name':
	// 'column_name', 'value': 'Age'}]
	InsightParameters []TestListResultsResponseItemsGoalThresholdsInsightParameter `json:"insightParameters,nullable"`
	// The measurement to be evaluated.
	Measurement string `json:"measurement"`
	// The operator to be used for the evaluation.
	Operator TestListResultsResponseItemsGoalThresholdsOperator `json:"operator"`
	// Whether to use automatic anomaly detection or manual thresholds
	ThresholdMode TestListResultsResponseItemsGoalThresholdsThresholdMode `json:"thresholdMode"`
	// The value to be compared.
	Value TestListResultsResponseItemsGoalThresholdsValueUnion `json:"value"`
	JSON  testListResultsResponseItemsGoalThresholdJSON        `json:"-"`
}

// testListResultsResponseItemsGoalThresholdJSON contains the JSON metadata for the
// struct [TestListResultsResponseItemsGoalThreshold]
type testListResultsResponseItemsGoalThresholdJSON struct {
	InsightName       apijson.Field
	InsightParameters apijson.Field
	Measurement       apijson.Field
	Operator          apijson.Field
	ThresholdMode     apijson.Field
	Value             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TestListResultsResponseItemsGoalThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResultsResponseItemsGoalThresholdJSON) RawJSON() string {
	return r.raw
}

// The insight name to be evaluated.
type TestListResultsResponseItemsGoalThresholdsInsightName string

const (
	TestListResultsResponseItemsGoalThresholdsInsightNameCharacterLength            TestListResultsResponseItemsGoalThresholdsInsightName = "characterLength"
	TestListResultsResponseItemsGoalThresholdsInsightNameClassImbalance             TestListResultsResponseItemsGoalThresholdsInsightName = "classImbalance"
	TestListResultsResponseItemsGoalThresholdsInsightNameExpectColumnAToBeInColumnB TestListResultsResponseItemsGoalThresholdsInsightName = "expectColumnAToBeInColumnB"
	TestListResultsResponseItemsGoalThresholdsInsightNameColumnAverage              TestListResultsResponseItemsGoalThresholdsInsightName = "columnAverage"
	TestListResultsResponseItemsGoalThresholdsInsightNameColumnDrift                TestListResultsResponseItemsGoalThresholdsInsightName = "columnDrift"
	TestListResultsResponseItemsGoalThresholdsInsightNameColumnValuesMatch          TestListResultsResponseItemsGoalThresholdsInsightName = "columnValuesMatch"
	TestListResultsResponseItemsGoalThresholdsInsightNameConfidenceDistribution     TestListResultsResponseItemsGoalThresholdsInsightName = "confidenceDistribution"
	TestListResultsResponseItemsGoalThresholdsInsightNameConflictingLabelRowCount   TestListResultsResponseItemsGoalThresholdsInsightName = "conflictingLabelRowCount"
	TestListResultsResponseItemsGoalThresholdsInsightNameContainsPii                TestListResultsResponseItemsGoalThresholdsInsightName = "containsPii"
	TestListResultsResponseItemsGoalThresholdsInsightNameContainsValidURL           TestListResultsResponseItemsGoalThresholdsInsightName = "containsValidUrl"
	TestListResultsResponseItemsGoalThresholdsInsightNameCorrelatedFeatures         TestListResultsResponseItemsGoalThresholdsInsightName = "correlatedFeatures"
	TestListResultsResponseItemsGoalThresholdsInsightNameCustomMetric               TestListResultsResponseItemsGoalThresholdsInsightName = "customMetric"
	TestListResultsResponseItemsGoalThresholdsInsightNameDuplicateRowCount          TestListResultsResponseItemsGoalThresholdsInsightName = "duplicateRowCount"
	TestListResultsResponseItemsGoalThresholdsInsightNameEmptyFeatures              TestListResultsResponseItemsGoalThresholdsInsightName = "emptyFeatures"
	TestListResultsResponseItemsGoalThresholdsInsightNameFeatureDrift               TestListResultsResponseItemsGoalThresholdsInsightName = "featureDrift"
	TestListResultsResponseItemsGoalThresholdsInsightNameFeatureProfile             TestListResultsResponseItemsGoalThresholdsInsightName = "featureProfile"
	TestListResultsResponseItemsGoalThresholdsInsightNameGreatExpectations          TestListResultsResponseItemsGoalThresholdsInsightName = "greatExpectations"
	TestListResultsResponseItemsGoalThresholdsInsightNameGroupByColumnStatsCheck    TestListResultsResponseItemsGoalThresholdsInsightName = "groupByColumnStatsCheck"
	TestListResultsResponseItemsGoalThresholdsInsightNameIllFormedRowCount          TestListResultsResponseItemsGoalThresholdsInsightName = "illFormedRowCount"
	TestListResultsResponseItemsGoalThresholdsInsightNameIsCode                     TestListResultsResponseItemsGoalThresholdsInsightName = "isCode"
	TestListResultsResponseItemsGoalThresholdsInsightNameIsJson                     TestListResultsResponseItemsGoalThresholdsInsightName = "isJson"
	TestListResultsResponseItemsGoalThresholdsInsightNameLlmRubricV2                TestListResultsResponseItemsGoalThresholdsInsightName = "llmRubricV2"
	TestListResultsResponseItemsGoalThresholdsInsightNameLabelDrift                 TestListResultsResponseItemsGoalThresholdsInsightName = "labelDrift"
	TestListResultsResponseItemsGoalThresholdsInsightNameMetrics                    TestListResultsResponseItemsGoalThresholdsInsightName = "metrics"
	TestListResultsResponseItemsGoalThresholdsInsightNameNewCategories              TestListResultsResponseItemsGoalThresholdsInsightName = "newCategories"
	TestListResultsResponseItemsGoalThresholdsInsightNameNewLabels                  TestListResultsResponseItemsGoalThresholdsInsightName = "newLabels"
	TestListResultsResponseItemsGoalThresholdsInsightNameNullRowCount               TestListResultsResponseItemsGoalThresholdsInsightName = "nullRowCount"
	TestListResultsResponseItemsGoalThresholdsInsightNamePpScore                    TestListResultsResponseItemsGoalThresholdsInsightName = "ppScore"
	TestListResultsResponseItemsGoalThresholdsInsightNameQuasiConstantFeatures      TestListResultsResponseItemsGoalThresholdsInsightName = "quasiConstantFeatures"
	TestListResultsResponseItemsGoalThresholdsInsightNameSentenceLength             TestListResultsResponseItemsGoalThresholdsInsightName = "sentenceLength"
	TestListResultsResponseItemsGoalThresholdsInsightNameSizeRatio                  TestListResultsResponseItemsGoalThresholdsInsightName = "sizeRatio"
	TestListResultsResponseItemsGoalThresholdsInsightNameSpecialCharacters          TestListResultsResponseItemsGoalThresholdsInsightName = "specialCharacters"
	TestListResultsResponseItemsGoalThresholdsInsightNameStringValidation           TestListResultsResponseItemsGoalThresholdsInsightName = "stringValidation"
	TestListResultsResponseItemsGoalThresholdsInsightNameTrainValLeakageRowCount    TestListResultsResponseItemsGoalThresholdsInsightName = "trainValLeakageRowCount"
)

func (r TestListResultsResponseItemsGoalThresholdsInsightName) IsKnown() bool {
	switch r {
	case TestListResultsResponseItemsGoalThresholdsInsightNameCharacterLength, TestListResultsResponseItemsGoalThresholdsInsightNameClassImbalance, TestListResultsResponseItemsGoalThresholdsInsightNameExpectColumnAToBeInColumnB, TestListResultsResponseItemsGoalThresholdsInsightNameColumnAverage, TestListResultsResponseItemsGoalThresholdsInsightNameColumnDrift, TestListResultsResponseItemsGoalThresholdsInsightNameColumnValuesMatch, TestListResultsResponseItemsGoalThresholdsInsightNameConfidenceDistribution, TestListResultsResponseItemsGoalThresholdsInsightNameConflictingLabelRowCount, TestListResultsResponseItemsGoalThresholdsInsightNameContainsPii, TestListResultsResponseItemsGoalThresholdsInsightNameContainsValidURL, TestListResultsResponseItemsGoalThresholdsInsightNameCorrelatedFeatures, TestListResultsResponseItemsGoalThresholdsInsightNameCustomMetric, TestListResultsResponseItemsGoalThresholdsInsightNameDuplicateRowCount, TestListResultsResponseItemsGoalThresholdsInsightNameEmptyFeatures, TestListResultsResponseItemsGoalThresholdsInsightNameFeatureDrift, TestListResultsResponseItemsGoalThresholdsInsightNameFeatureProfile, TestListResultsResponseItemsGoalThresholdsInsightNameGreatExpectations, TestListResultsResponseItemsGoalThresholdsInsightNameGroupByColumnStatsCheck, TestListResultsResponseItemsGoalThresholdsInsightNameIllFormedRowCount, TestListResultsResponseItemsGoalThresholdsInsightNameIsCode, TestListResultsResponseItemsGoalThresholdsInsightNameIsJson, TestListResultsResponseItemsGoalThresholdsInsightNameLlmRubricV2, TestListResultsResponseItemsGoalThresholdsInsightNameLabelDrift, TestListResultsResponseItemsGoalThresholdsInsightNameMetrics, TestListResultsResponseItemsGoalThresholdsInsightNameNewCategories, TestListResultsResponseItemsGoalThresholdsInsightNameNewLabels, TestListResultsResponseItemsGoalThresholdsInsightNameNullRowCount, TestListResultsResponseItemsGoalThresholdsInsightNamePpScore, TestListResultsResponseItemsGoalThresholdsInsightNameQuasiConstantFeatures, TestListResultsResponseItemsGoalThresholdsInsightNameSentenceLength, TestListResultsResponseItemsGoalThresholdsInsightNameSizeRatio, TestListResultsResponseItemsGoalThresholdsInsightNameSpecialCharacters, TestListResultsResponseItemsGoalThresholdsInsightNameStringValidation, TestListResultsResponseItemsGoalThresholdsInsightNameTrainValLeakageRowCount:
		return true
	}
	return false
}

type TestListResultsResponseItemsGoalThresholdsInsightParameter struct {
	// The name of the insight filter.
	Name  string                                                         `json:"name,required"`
	Value interface{}                                                    `json:"value,required"`
	JSON  testListResultsResponseItemsGoalThresholdsInsightParameterJSON `json:"-"`
}

// testListResultsResponseItemsGoalThresholdsInsightParameterJSON contains the JSON
// metadata for the struct
// [TestListResultsResponseItemsGoalThresholdsInsightParameter]
type testListResultsResponseItemsGoalThresholdsInsightParameterJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestListResultsResponseItemsGoalThresholdsInsightParameter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResultsResponseItemsGoalThresholdsInsightParameterJSON) RawJSON() string {
	return r.raw
}

// The operator to be used for the evaluation.
type TestListResultsResponseItemsGoalThresholdsOperator string

const (
	TestListResultsResponseItemsGoalThresholdsOperatorIs              TestListResultsResponseItemsGoalThresholdsOperator = "is"
	TestListResultsResponseItemsGoalThresholdsOperatorGreater         TestListResultsResponseItemsGoalThresholdsOperator = ">"
	TestListResultsResponseItemsGoalThresholdsOperatorGreaterOrEquals TestListResultsResponseItemsGoalThresholdsOperator = ">="
	TestListResultsResponseItemsGoalThresholdsOperatorLess            TestListResultsResponseItemsGoalThresholdsOperator = "<"
	TestListResultsResponseItemsGoalThresholdsOperatorLessOrEquals    TestListResultsResponseItemsGoalThresholdsOperator = "<="
	TestListResultsResponseItemsGoalThresholdsOperatorNotEquals       TestListResultsResponseItemsGoalThresholdsOperator = "!="
)

func (r TestListResultsResponseItemsGoalThresholdsOperator) IsKnown() bool {
	switch r {
	case TestListResultsResponseItemsGoalThresholdsOperatorIs, TestListResultsResponseItemsGoalThresholdsOperatorGreater, TestListResultsResponseItemsGoalThresholdsOperatorGreaterOrEquals, TestListResultsResponseItemsGoalThresholdsOperatorLess, TestListResultsResponseItemsGoalThresholdsOperatorLessOrEquals, TestListResultsResponseItemsGoalThresholdsOperatorNotEquals:
		return true
	}
	return false
}

// Whether to use automatic anomaly detection or manual thresholds
type TestListResultsResponseItemsGoalThresholdsThresholdMode string

const (
	TestListResultsResponseItemsGoalThresholdsThresholdModeAutomatic TestListResultsResponseItemsGoalThresholdsThresholdMode = "automatic"
	TestListResultsResponseItemsGoalThresholdsThresholdModeManual    TestListResultsResponseItemsGoalThresholdsThresholdMode = "manual"
)

func (r TestListResultsResponseItemsGoalThresholdsThresholdMode) IsKnown() bool {
	switch r {
	case TestListResultsResponseItemsGoalThresholdsThresholdModeAutomatic, TestListResultsResponseItemsGoalThresholdsThresholdModeManual:
		return true
	}
	return false
}

// The value to be compared.
//
// Union satisfied by [shared.UnionFloat], [shared.UnionBool], [shared.UnionString]
// or [TestListResultsResponseItemsGoalThresholdsValueArray].
type TestListResultsResponseItemsGoalThresholdsValueUnion interface {
	ImplementsTestListResultsResponseItemsGoalThresholdsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TestListResultsResponseItemsGoalThresholdsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(TestListResultsResponseItemsGoalThresholdsValueArray{}),
		},
	)
}

type TestListResultsResponseItemsGoalThresholdsValueArray []string

func (r TestListResultsResponseItemsGoalThresholdsValueArray) ImplementsTestListResultsResponseItemsGoalThresholdsValueUnion() {
}

// The test type.
type TestListResultsResponseItemsGoalType string

const (
	TestListResultsResponseItemsGoalTypeIntegrity   TestListResultsResponseItemsGoalType = "integrity"
	TestListResultsResponseItemsGoalTypeConsistency TestListResultsResponseItemsGoalType = "consistency"
	TestListResultsResponseItemsGoalTypePerformance TestListResultsResponseItemsGoalType = "performance"
)

func (r TestListResultsResponseItemsGoalType) IsKnown() bool {
	switch r {
	case TestListResultsResponseItemsGoalTypeIntegrity, TestListResultsResponseItemsGoalTypeConsistency, TestListResultsResponseItemsGoalTypePerformance:
		return true
	}
	return false
}

// The body of the rows request.
type TestListResultsResponseItemsRowsBody struct {
	ColumnFilters     []TestListResultsResponseItemsRowsBodyColumnFilter `json:"columnFilters,nullable"`
	ExcludeRowIDList  []int64                                            `json:"excludeRowIdList,nullable"`
	NotSearchQueryAnd []string                                           `json:"notSearchQueryAnd,nullable"`
	NotSearchQueryOr  []string                                           `json:"notSearchQueryOr,nullable"`
	RowIDList         []int64                                            `json:"rowIdList,nullable"`
	SearchQueryAnd    []string                                           `json:"searchQueryAnd,nullable"`
	SearchQueryOr     []string                                           `json:"searchQueryOr,nullable"`
	JSON              testListResultsResponseItemsRowsBodyJSON           `json:"-"`
}

// testListResultsResponseItemsRowsBodyJSON contains the JSON metadata for the
// struct [TestListResultsResponseItemsRowsBody]
type testListResultsResponseItemsRowsBodyJSON struct {
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

func (r *TestListResultsResponseItemsRowsBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResultsResponseItemsRowsBodyJSON) RawJSON() string {
	return r.raw
}

type TestListResultsResponseItemsRowsBodyColumnFilter struct {
	// The name of the column.
	Measurement string                                                    `json:"measurement,required"`
	Operator    TestListResultsResponseItemsRowsBodyColumnFiltersOperator `json:"operator,required"`
	// This field can have the runtime type of
	// [[]TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterValueUnion],
	// [float64],
	// [TestListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilterValueUnion].
	Value interface{}                                          `json:"value,required"`
	JSON  testListResultsResponseItemsRowsBodyColumnFilterJSON `json:"-"`
	union TestListResultsResponseItemsRowsBodyColumnFiltersUnion
}

// testListResultsResponseItemsRowsBodyColumnFilterJSON contains the JSON metadata
// for the struct [TestListResultsResponseItemsRowsBodyColumnFilter]
type testListResultsResponseItemsRowsBodyColumnFilterJSON struct {
	Measurement apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r testListResultsResponseItemsRowsBodyColumnFilterJSON) RawJSON() string {
	return r.raw
}

func (r *TestListResultsResponseItemsRowsBodyColumnFilter) UnmarshalJSON(data []byte) (err error) {
	*r = TestListResultsResponseItemsRowsBodyColumnFilter{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TestListResultsResponseItemsRowsBodyColumnFiltersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilter],
// [TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilter],
// [TestListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilter].
func (r TestListResultsResponseItemsRowsBodyColumnFilter) AsUnion() TestListResultsResponseItemsRowsBodyColumnFiltersUnion {
	return r.union
}

// Union satisfied by
// [TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilter],
// [TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilter] or
// [TestListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilter].
type TestListResultsResponseItemsRowsBodyColumnFiltersUnion interface {
	implementsTestListResultsResponseItemsRowsBodyColumnFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TestListResultsResponseItemsRowsBodyColumnFiltersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TestListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilter{}),
		},
	)
}

type TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilter struct {
	// The name of the column.
	Measurement string                                                                       `json:"measurement,required"`
	Operator    TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterOperator     `json:"operator,required"`
	Value       []TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterValueUnion `json:"value,required"`
	JSON        testListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterJSON         `json:"-"`
}

// testListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterJSON contains
// the JSON metadata for the struct
// [TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilter]
type testListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterJSON struct {
	Measurement apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterJSON) RawJSON() string {
	return r.raw
}

func (r TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilter) implementsTestListResultsResponseItemsRowsBodyColumnFilter() {
}

type TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterOperator string

const (
	TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterOperatorContainsNone TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterOperator = "contains_none"
	TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterOperatorContainsAny  TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterOperator = "contains_any"
	TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterOperatorContainsAll  TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterOperator = "contains_all"
	TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterOperatorOneOf        TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterOperator = "one_of"
	TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterOperatorNoneOf       TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterOperator = "none_of"
)

func (r TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterOperator) IsKnown() bool {
	switch r {
	case TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterOperatorContainsNone, TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterOperatorContainsAny, TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterOperatorContainsAll, TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterOperatorOneOf, TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterOperatorNoneOf:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterValueUnion interface {
	ImplementsTestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TestListResultsResponseItemsRowsBodyColumnFiltersSetColumnFilterValueUnion)(nil)).Elem(),
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

type TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilter struct {
	// The name of the column.
	Measurement string                                                                       `json:"measurement,required"`
	Operator    TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperator `json:"operator,required"`
	Value       float64                                                                      `json:"value,required,nullable"`
	JSON        testListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterJSON     `json:"-"`
}

// testListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterJSON
// contains the JSON metadata for the struct
// [TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilter]
type testListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterJSON struct {
	Measurement apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterJSON) RawJSON() string {
	return r.raw
}

func (r TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilter) implementsTestListResultsResponseItemsRowsBodyColumnFilter() {
}

type TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperator string

const (
	TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorGreater         TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperator = ">"
	TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorGreaterOrEquals TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperator = ">="
	TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorIs              TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperator = "is"
	TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorLess            TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperator = "<"
	TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorLessOrEquals    TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperator = "<="
	TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorNotEquals       TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperator = "!="
)

func (r TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperator) IsKnown() bool {
	switch r {
	case TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorGreater, TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorGreaterOrEquals, TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorIs, TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorLess, TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorLessOrEquals, TestListResultsResponseItemsRowsBodyColumnFiltersNumericColumnFilterOperatorNotEquals:
		return true
	}
	return false
}

type TestListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilter struct {
	// The name of the column.
	Measurement string                                                                        `json:"measurement,required"`
	Operator    TestListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilterOperator   `json:"operator,required"`
	Value       TestListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilterValueUnion `json:"value,required"`
	JSON        testListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilterJSON       `json:"-"`
}

// testListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilterJSON contains
// the JSON metadata for the struct
// [TestListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilter]
type testListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilterJSON struct {
	Measurement apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilterJSON) RawJSON() string {
	return r.raw
}

func (r TestListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilter) implementsTestListResultsResponseItemsRowsBodyColumnFilter() {
}

type TestListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilterOperator string

const (
	TestListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilterOperatorIs        TestListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilterOperator = "is"
	TestListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilterOperatorNotEquals TestListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilterOperator = "!="
)

func (r TestListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilterOperator) IsKnown() bool {
	switch r {
	case TestListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilterOperatorIs, TestListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilterOperatorNotEquals:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type TestListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilterValueUnion interface {
	ImplementsTestListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilterValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TestListResultsResponseItemsRowsBodyColumnFiltersStringColumnFilterValueUnion)(nil)).Elem(),
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

type TestListResultsResponseItemsRowsBodyColumnFiltersOperator string

const (
	TestListResultsResponseItemsRowsBodyColumnFiltersOperatorContainsNone    TestListResultsResponseItemsRowsBodyColumnFiltersOperator = "contains_none"
	TestListResultsResponseItemsRowsBodyColumnFiltersOperatorContainsAny     TestListResultsResponseItemsRowsBodyColumnFiltersOperator = "contains_any"
	TestListResultsResponseItemsRowsBodyColumnFiltersOperatorContainsAll     TestListResultsResponseItemsRowsBodyColumnFiltersOperator = "contains_all"
	TestListResultsResponseItemsRowsBodyColumnFiltersOperatorOneOf           TestListResultsResponseItemsRowsBodyColumnFiltersOperator = "one_of"
	TestListResultsResponseItemsRowsBodyColumnFiltersOperatorNoneOf          TestListResultsResponseItemsRowsBodyColumnFiltersOperator = "none_of"
	TestListResultsResponseItemsRowsBodyColumnFiltersOperatorGreater         TestListResultsResponseItemsRowsBodyColumnFiltersOperator = ">"
	TestListResultsResponseItemsRowsBodyColumnFiltersOperatorGreaterOrEquals TestListResultsResponseItemsRowsBodyColumnFiltersOperator = ">="
	TestListResultsResponseItemsRowsBodyColumnFiltersOperatorIs              TestListResultsResponseItemsRowsBodyColumnFiltersOperator = "is"
	TestListResultsResponseItemsRowsBodyColumnFiltersOperatorLess            TestListResultsResponseItemsRowsBodyColumnFiltersOperator = "<"
	TestListResultsResponseItemsRowsBodyColumnFiltersOperatorLessOrEquals    TestListResultsResponseItemsRowsBodyColumnFiltersOperator = "<="
	TestListResultsResponseItemsRowsBodyColumnFiltersOperatorNotEquals       TestListResultsResponseItemsRowsBodyColumnFiltersOperator = "!="
)

func (r TestListResultsResponseItemsRowsBodyColumnFiltersOperator) IsKnown() bool {
	switch r {
	case TestListResultsResponseItemsRowsBodyColumnFiltersOperatorContainsNone, TestListResultsResponseItemsRowsBodyColumnFiltersOperatorContainsAny, TestListResultsResponseItemsRowsBodyColumnFiltersOperatorContainsAll, TestListResultsResponseItemsRowsBodyColumnFiltersOperatorOneOf, TestListResultsResponseItemsRowsBodyColumnFiltersOperatorNoneOf, TestListResultsResponseItemsRowsBodyColumnFiltersOperatorGreater, TestListResultsResponseItemsRowsBodyColumnFiltersOperatorGreaterOrEquals, TestListResultsResponseItemsRowsBodyColumnFiltersOperatorIs, TestListResultsResponseItemsRowsBodyColumnFiltersOperatorLess, TestListResultsResponseItemsRowsBodyColumnFiltersOperatorLessOrEquals, TestListResultsResponseItemsRowsBodyColumnFiltersOperatorNotEquals:
		return true
	}
	return false
}

type TestListResultsResponseLastUnskippedResult struct {
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
	Status TestListResultsResponseLastUnskippedResultStatus `json:"status,required"`
	// The status message.
	StatusMessage  string                                                    `json:"statusMessage,required,nullable"`
	ExpectedValues []TestListResultsResponseLastUnskippedResultExpectedValue `json:"expectedValues"`
	Goal           TestListResultsResponseLastUnskippedResultGoal            `json:"goal"`
	// The test id.
	GoalID string `json:"goalId,nullable" format:"uuid"`
	// The URL to the rows of the test result.
	Rows string `json:"rows"`
	// The body of the rows request.
	RowsBody TestListResultsResponseLastUnskippedResultRowsBody `json:"rowsBody,nullable"`
	JSON     testListResultsResponseLastUnskippedResultJSON     `json:"-"`
}

// testListResultsResponseLastUnskippedResultJSON contains the JSON metadata for
// the struct [TestListResultsResponseLastUnskippedResult]
type testListResultsResponseLastUnskippedResultJSON struct {
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

func (r *TestListResultsResponseLastUnskippedResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResultsResponseLastUnskippedResultJSON) RawJSON() string {
	return r.raw
}

// The status of the test.
type TestListResultsResponseLastUnskippedResultStatus string

const (
	TestListResultsResponseLastUnskippedResultStatusRunning TestListResultsResponseLastUnskippedResultStatus = "running"
	TestListResultsResponseLastUnskippedResultStatusPassing TestListResultsResponseLastUnskippedResultStatus = "passing"
	TestListResultsResponseLastUnskippedResultStatusFailing TestListResultsResponseLastUnskippedResultStatus = "failing"
	TestListResultsResponseLastUnskippedResultStatusSkipped TestListResultsResponseLastUnskippedResultStatus = "skipped"
	TestListResultsResponseLastUnskippedResultStatusError   TestListResultsResponseLastUnskippedResultStatus = "error"
)

func (r TestListResultsResponseLastUnskippedResultStatus) IsKnown() bool {
	switch r {
	case TestListResultsResponseLastUnskippedResultStatusRunning, TestListResultsResponseLastUnskippedResultStatusPassing, TestListResultsResponseLastUnskippedResultStatusFailing, TestListResultsResponseLastUnskippedResultStatusSkipped, TestListResultsResponseLastUnskippedResultStatusError:
		return true
	}
	return false
}

type TestListResultsResponseLastUnskippedResultExpectedValue struct {
	// the lower threshold for the expected value
	LowerThreshold float64 `json:"lowerThreshold,nullable"`
	// One of the `measurement` values in the test's thresholds
	Measurement string `json:"measurement"`
	// The upper threshold for the expected value
	UpperThreshold float64                                                     `json:"upperThreshold,nullable"`
	JSON           testListResultsResponseLastUnskippedResultExpectedValueJSON `json:"-"`
}

// testListResultsResponseLastUnskippedResultExpectedValueJSON contains the JSON
// metadata for the struct
// [TestListResultsResponseLastUnskippedResultExpectedValue]
type testListResultsResponseLastUnskippedResultExpectedValueJSON struct {
	LowerThreshold apijson.Field
	Measurement    apijson.Field
	UpperThreshold apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TestListResultsResponseLastUnskippedResultExpectedValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResultsResponseLastUnskippedResultExpectedValueJSON) RawJSON() string {
	return r.raw
}

type TestListResultsResponseLastUnskippedResultGoal struct {
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
	Subtype TestListResultsResponseLastUnskippedResultGoalSubtype `json:"subtype,required"`
	// Whether the test is suggested or user-created.
	Suggested  bool                                                      `json:"suggested,required"`
	Thresholds []TestListResultsResponseLastUnskippedResultGoalThreshold `json:"thresholds,required"`
	// The test type.
	Type TestListResultsResponseLastUnskippedResultGoalType `json:"type,required"`
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
	UsesValidationDataset bool                                               `json:"usesValidationDataset"`
	JSON                  testListResultsResponseLastUnskippedResultGoalJSON `json:"-"`
}

// testListResultsResponseLastUnskippedResultGoalJSON contains the JSON metadata
// for the struct [TestListResultsResponseLastUnskippedResultGoal]
type testListResultsResponseLastUnskippedResultGoalJSON struct {
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

func (r *TestListResultsResponseLastUnskippedResultGoal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResultsResponseLastUnskippedResultGoalJSON) RawJSON() string {
	return r.raw
}

// The test subtype.
type TestListResultsResponseLastUnskippedResultGoalSubtype string

const (
	TestListResultsResponseLastUnskippedResultGoalSubtypeAnomalousColumnCount       TestListResultsResponseLastUnskippedResultGoalSubtype = "anomalousColumnCount"
	TestListResultsResponseLastUnskippedResultGoalSubtypeCharacterLength            TestListResultsResponseLastUnskippedResultGoalSubtype = "characterLength"
	TestListResultsResponseLastUnskippedResultGoalSubtypeClassImbalanceRatio        TestListResultsResponseLastUnskippedResultGoalSubtype = "classImbalanceRatio"
	TestListResultsResponseLastUnskippedResultGoalSubtypeExpectColumnAToBeInColumnB TestListResultsResponseLastUnskippedResultGoalSubtype = "expectColumnAToBeInColumnB"
	TestListResultsResponseLastUnskippedResultGoalSubtypeColumnAverage              TestListResultsResponseLastUnskippedResultGoalSubtype = "columnAverage"
	TestListResultsResponseLastUnskippedResultGoalSubtypeColumnDrift                TestListResultsResponseLastUnskippedResultGoalSubtype = "columnDrift"
	TestListResultsResponseLastUnskippedResultGoalSubtypeColumnStatistic            TestListResultsResponseLastUnskippedResultGoalSubtype = "columnStatistic"
	TestListResultsResponseLastUnskippedResultGoalSubtypeColumnValuesMatch          TestListResultsResponseLastUnskippedResultGoalSubtype = "columnValuesMatch"
	TestListResultsResponseLastUnskippedResultGoalSubtypeConflictingLabelRowCount   TestListResultsResponseLastUnskippedResultGoalSubtype = "conflictingLabelRowCount"
	TestListResultsResponseLastUnskippedResultGoalSubtypeContainsPii                TestListResultsResponseLastUnskippedResultGoalSubtype = "containsPii"
	TestListResultsResponseLastUnskippedResultGoalSubtypeContainsValidURL           TestListResultsResponseLastUnskippedResultGoalSubtype = "containsValidUrl"
	TestListResultsResponseLastUnskippedResultGoalSubtypeCorrelatedFeatureCount     TestListResultsResponseLastUnskippedResultGoalSubtype = "correlatedFeatureCount"
	TestListResultsResponseLastUnskippedResultGoalSubtypeCustomMetricThreshold      TestListResultsResponseLastUnskippedResultGoalSubtype = "customMetricThreshold"
	TestListResultsResponseLastUnskippedResultGoalSubtypeDuplicateRowCount          TestListResultsResponseLastUnskippedResultGoalSubtype = "duplicateRowCount"
	TestListResultsResponseLastUnskippedResultGoalSubtypeEmptyFeature               TestListResultsResponseLastUnskippedResultGoalSubtype = "emptyFeature"
	TestListResultsResponseLastUnskippedResultGoalSubtypeEmptyFeatureCount          TestListResultsResponseLastUnskippedResultGoalSubtype = "emptyFeatureCount"
	TestListResultsResponseLastUnskippedResultGoalSubtypeDriftedFeatureCount        TestListResultsResponseLastUnskippedResultGoalSubtype = "driftedFeatureCount"
	TestListResultsResponseLastUnskippedResultGoalSubtypeFeatureMissingValues       TestListResultsResponseLastUnskippedResultGoalSubtype = "featureMissingValues"
	TestListResultsResponseLastUnskippedResultGoalSubtypeFeatureValueValidation     TestListResultsResponseLastUnskippedResultGoalSubtype = "featureValueValidation"
	TestListResultsResponseLastUnskippedResultGoalSubtypeGreatExpectations          TestListResultsResponseLastUnskippedResultGoalSubtype = "greatExpectations"
	TestListResultsResponseLastUnskippedResultGoalSubtypeGroupByColumnStatsCheck    TestListResultsResponseLastUnskippedResultGoalSubtype = "groupByColumnStatsCheck"
	TestListResultsResponseLastUnskippedResultGoalSubtypeIllFormedRowCount          TestListResultsResponseLastUnskippedResultGoalSubtype = "illFormedRowCount"
	TestListResultsResponseLastUnskippedResultGoalSubtypeIsCode                     TestListResultsResponseLastUnskippedResultGoalSubtype = "isCode"
	TestListResultsResponseLastUnskippedResultGoalSubtypeIsJson                     TestListResultsResponseLastUnskippedResultGoalSubtype = "isJson"
	TestListResultsResponseLastUnskippedResultGoalSubtypeLlmRubricThresholdV2       TestListResultsResponseLastUnskippedResultGoalSubtype = "llmRubricThresholdV2"
	TestListResultsResponseLastUnskippedResultGoalSubtypeLabelDrift                 TestListResultsResponseLastUnskippedResultGoalSubtype = "labelDrift"
	TestListResultsResponseLastUnskippedResultGoalSubtypeMetricThreshold            TestListResultsResponseLastUnskippedResultGoalSubtype = "metricThreshold"
	TestListResultsResponseLastUnskippedResultGoalSubtypeNewCategoryCount           TestListResultsResponseLastUnskippedResultGoalSubtype = "newCategoryCount"
	TestListResultsResponseLastUnskippedResultGoalSubtypeNewLabelCount              TestListResultsResponseLastUnskippedResultGoalSubtype = "newLabelCount"
	TestListResultsResponseLastUnskippedResultGoalSubtypeNullRowCount               TestListResultsResponseLastUnskippedResultGoalSubtype = "nullRowCount"
	TestListResultsResponseLastUnskippedResultGoalSubtypeRowCount                   TestListResultsResponseLastUnskippedResultGoalSubtype = "rowCount"
	TestListResultsResponseLastUnskippedResultGoalSubtypePpScoreValueValidation     TestListResultsResponseLastUnskippedResultGoalSubtype = "ppScoreValueValidation"
	TestListResultsResponseLastUnskippedResultGoalSubtypeQuasiConstantFeature       TestListResultsResponseLastUnskippedResultGoalSubtype = "quasiConstantFeature"
	TestListResultsResponseLastUnskippedResultGoalSubtypeQuasiConstantFeatureCount  TestListResultsResponseLastUnskippedResultGoalSubtype = "quasiConstantFeatureCount"
	TestListResultsResponseLastUnskippedResultGoalSubtypeSqlQuery                   TestListResultsResponseLastUnskippedResultGoalSubtype = "sqlQuery"
	TestListResultsResponseLastUnskippedResultGoalSubtypeDtypeValidation            TestListResultsResponseLastUnskippedResultGoalSubtype = "dtypeValidation"
	TestListResultsResponseLastUnskippedResultGoalSubtypeSentenceLength             TestListResultsResponseLastUnskippedResultGoalSubtype = "sentenceLength"
	TestListResultsResponseLastUnskippedResultGoalSubtypeSizeRatio                  TestListResultsResponseLastUnskippedResultGoalSubtype = "sizeRatio"
	TestListResultsResponseLastUnskippedResultGoalSubtypeSpecialCharactersRatio     TestListResultsResponseLastUnskippedResultGoalSubtype = "specialCharactersRatio"
	TestListResultsResponseLastUnskippedResultGoalSubtypeStringValidation           TestListResultsResponseLastUnskippedResultGoalSubtype = "stringValidation"
	TestListResultsResponseLastUnskippedResultGoalSubtypeTrainValLeakageRowCount    TestListResultsResponseLastUnskippedResultGoalSubtype = "trainValLeakageRowCount"
)

func (r TestListResultsResponseLastUnskippedResultGoalSubtype) IsKnown() bool {
	switch r {
	case TestListResultsResponseLastUnskippedResultGoalSubtypeAnomalousColumnCount, TestListResultsResponseLastUnskippedResultGoalSubtypeCharacterLength, TestListResultsResponseLastUnskippedResultGoalSubtypeClassImbalanceRatio, TestListResultsResponseLastUnskippedResultGoalSubtypeExpectColumnAToBeInColumnB, TestListResultsResponseLastUnskippedResultGoalSubtypeColumnAverage, TestListResultsResponseLastUnskippedResultGoalSubtypeColumnDrift, TestListResultsResponseLastUnskippedResultGoalSubtypeColumnStatistic, TestListResultsResponseLastUnskippedResultGoalSubtypeColumnValuesMatch, TestListResultsResponseLastUnskippedResultGoalSubtypeConflictingLabelRowCount, TestListResultsResponseLastUnskippedResultGoalSubtypeContainsPii, TestListResultsResponseLastUnskippedResultGoalSubtypeContainsValidURL, TestListResultsResponseLastUnskippedResultGoalSubtypeCorrelatedFeatureCount, TestListResultsResponseLastUnskippedResultGoalSubtypeCustomMetricThreshold, TestListResultsResponseLastUnskippedResultGoalSubtypeDuplicateRowCount, TestListResultsResponseLastUnskippedResultGoalSubtypeEmptyFeature, TestListResultsResponseLastUnskippedResultGoalSubtypeEmptyFeatureCount, TestListResultsResponseLastUnskippedResultGoalSubtypeDriftedFeatureCount, TestListResultsResponseLastUnskippedResultGoalSubtypeFeatureMissingValues, TestListResultsResponseLastUnskippedResultGoalSubtypeFeatureValueValidation, TestListResultsResponseLastUnskippedResultGoalSubtypeGreatExpectations, TestListResultsResponseLastUnskippedResultGoalSubtypeGroupByColumnStatsCheck, TestListResultsResponseLastUnskippedResultGoalSubtypeIllFormedRowCount, TestListResultsResponseLastUnskippedResultGoalSubtypeIsCode, TestListResultsResponseLastUnskippedResultGoalSubtypeIsJson, TestListResultsResponseLastUnskippedResultGoalSubtypeLlmRubricThresholdV2, TestListResultsResponseLastUnskippedResultGoalSubtypeLabelDrift, TestListResultsResponseLastUnskippedResultGoalSubtypeMetricThreshold, TestListResultsResponseLastUnskippedResultGoalSubtypeNewCategoryCount, TestListResultsResponseLastUnskippedResultGoalSubtypeNewLabelCount, TestListResultsResponseLastUnskippedResultGoalSubtypeNullRowCount, TestListResultsResponseLastUnskippedResultGoalSubtypeRowCount, TestListResultsResponseLastUnskippedResultGoalSubtypePpScoreValueValidation, TestListResultsResponseLastUnskippedResultGoalSubtypeQuasiConstantFeature, TestListResultsResponseLastUnskippedResultGoalSubtypeQuasiConstantFeatureCount, TestListResultsResponseLastUnskippedResultGoalSubtypeSqlQuery, TestListResultsResponseLastUnskippedResultGoalSubtypeDtypeValidation, TestListResultsResponseLastUnskippedResultGoalSubtypeSentenceLength, TestListResultsResponseLastUnskippedResultGoalSubtypeSizeRatio, TestListResultsResponseLastUnskippedResultGoalSubtypeSpecialCharactersRatio, TestListResultsResponseLastUnskippedResultGoalSubtypeStringValidation, TestListResultsResponseLastUnskippedResultGoalSubtypeTrainValLeakageRowCount:
		return true
	}
	return false
}

type TestListResultsResponseLastUnskippedResultGoalThreshold struct {
	// The insight name to be evaluated.
	InsightName TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName `json:"insightName"`
	// The insight parameters. Required only for some test subtypes. For example, for
	// tests that require a column name, the insight parameters will be [{'name':
	// 'column_name', 'value': 'Age'}]
	InsightParameters []TestListResultsResponseLastUnskippedResultGoalThresholdsInsightParameter `json:"insightParameters,nullable"`
	// The measurement to be evaluated.
	Measurement string `json:"measurement"`
	// The operator to be used for the evaluation.
	Operator TestListResultsResponseLastUnskippedResultGoalThresholdsOperator `json:"operator"`
	// Whether to use automatic anomaly detection or manual thresholds
	ThresholdMode TestListResultsResponseLastUnskippedResultGoalThresholdsThresholdMode `json:"thresholdMode"`
	// The value to be compared.
	Value TestListResultsResponseLastUnskippedResultGoalThresholdsValueUnion `json:"value"`
	JSON  testListResultsResponseLastUnskippedResultGoalThresholdJSON        `json:"-"`
}

// testListResultsResponseLastUnskippedResultGoalThresholdJSON contains the JSON
// metadata for the struct
// [TestListResultsResponseLastUnskippedResultGoalThreshold]
type testListResultsResponseLastUnskippedResultGoalThresholdJSON struct {
	InsightName       apijson.Field
	InsightParameters apijson.Field
	Measurement       apijson.Field
	Operator          apijson.Field
	ThresholdMode     apijson.Field
	Value             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TestListResultsResponseLastUnskippedResultGoalThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResultsResponseLastUnskippedResultGoalThresholdJSON) RawJSON() string {
	return r.raw
}

// The insight name to be evaluated.
type TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName string

const (
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameCharacterLength            TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "characterLength"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameClassImbalance             TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "classImbalance"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameExpectColumnAToBeInColumnB TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "expectColumnAToBeInColumnB"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameColumnAverage              TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "columnAverage"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameColumnDrift                TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "columnDrift"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameColumnValuesMatch          TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "columnValuesMatch"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameConfidenceDistribution     TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "confidenceDistribution"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameConflictingLabelRowCount   TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "conflictingLabelRowCount"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameContainsPii                TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "containsPii"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameContainsValidURL           TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "containsValidUrl"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameCorrelatedFeatures         TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "correlatedFeatures"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameCustomMetric               TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "customMetric"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameDuplicateRowCount          TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "duplicateRowCount"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameEmptyFeatures              TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "emptyFeatures"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameFeatureDrift               TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "featureDrift"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameFeatureProfile             TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "featureProfile"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameGreatExpectations          TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "greatExpectations"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameGroupByColumnStatsCheck    TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "groupByColumnStatsCheck"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameIllFormedRowCount          TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "illFormedRowCount"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameIsCode                     TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "isCode"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameIsJson                     TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "isJson"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameLlmRubricV2                TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "llmRubricV2"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameLabelDrift                 TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "labelDrift"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameMetrics                    TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "metrics"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameNewCategories              TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "newCategories"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameNewLabels                  TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "newLabels"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameNullRowCount               TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "nullRowCount"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNamePpScore                    TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "ppScore"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameQuasiConstantFeatures      TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "quasiConstantFeatures"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameSentenceLength             TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "sentenceLength"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameSizeRatio                  TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "sizeRatio"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameSpecialCharacters          TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "specialCharacters"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameStringValidation           TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "stringValidation"
	TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameTrainValLeakageRowCount    TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName = "trainValLeakageRowCount"
)

func (r TestListResultsResponseLastUnskippedResultGoalThresholdsInsightName) IsKnown() bool {
	switch r {
	case TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameCharacterLength, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameClassImbalance, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameExpectColumnAToBeInColumnB, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameColumnAverage, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameColumnDrift, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameColumnValuesMatch, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameConfidenceDistribution, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameConflictingLabelRowCount, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameContainsPii, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameContainsValidURL, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameCorrelatedFeatures, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameCustomMetric, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameDuplicateRowCount, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameEmptyFeatures, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameFeatureDrift, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameFeatureProfile, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameGreatExpectations, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameGroupByColumnStatsCheck, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameIllFormedRowCount, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameIsCode, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameIsJson, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameLlmRubricV2, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameLabelDrift, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameMetrics, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameNewCategories, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameNewLabels, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameNullRowCount, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNamePpScore, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameQuasiConstantFeatures, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameSentenceLength, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameSizeRatio, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameSpecialCharacters, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameStringValidation, TestListResultsResponseLastUnskippedResultGoalThresholdsInsightNameTrainValLeakageRowCount:
		return true
	}
	return false
}

type TestListResultsResponseLastUnskippedResultGoalThresholdsInsightParameter struct {
	// The name of the insight filter.
	Name  string                                                                       `json:"name,required"`
	Value interface{}                                                                  `json:"value,required"`
	JSON  testListResultsResponseLastUnskippedResultGoalThresholdsInsightParameterJSON `json:"-"`
}

// testListResultsResponseLastUnskippedResultGoalThresholdsInsightParameterJSON
// contains the JSON metadata for the struct
// [TestListResultsResponseLastUnskippedResultGoalThresholdsInsightParameter]
type testListResultsResponseLastUnskippedResultGoalThresholdsInsightParameterJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestListResultsResponseLastUnskippedResultGoalThresholdsInsightParameter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResultsResponseLastUnskippedResultGoalThresholdsInsightParameterJSON) RawJSON() string {
	return r.raw
}

// The operator to be used for the evaluation.
type TestListResultsResponseLastUnskippedResultGoalThresholdsOperator string

const (
	TestListResultsResponseLastUnskippedResultGoalThresholdsOperatorIs              TestListResultsResponseLastUnskippedResultGoalThresholdsOperator = "is"
	TestListResultsResponseLastUnskippedResultGoalThresholdsOperatorGreater         TestListResultsResponseLastUnskippedResultGoalThresholdsOperator = ">"
	TestListResultsResponseLastUnskippedResultGoalThresholdsOperatorGreaterOrEquals TestListResultsResponseLastUnskippedResultGoalThresholdsOperator = ">="
	TestListResultsResponseLastUnskippedResultGoalThresholdsOperatorLess            TestListResultsResponseLastUnskippedResultGoalThresholdsOperator = "<"
	TestListResultsResponseLastUnskippedResultGoalThresholdsOperatorLessOrEquals    TestListResultsResponseLastUnskippedResultGoalThresholdsOperator = "<="
	TestListResultsResponseLastUnskippedResultGoalThresholdsOperatorNotEquals       TestListResultsResponseLastUnskippedResultGoalThresholdsOperator = "!="
)

func (r TestListResultsResponseLastUnskippedResultGoalThresholdsOperator) IsKnown() bool {
	switch r {
	case TestListResultsResponseLastUnskippedResultGoalThresholdsOperatorIs, TestListResultsResponseLastUnskippedResultGoalThresholdsOperatorGreater, TestListResultsResponseLastUnskippedResultGoalThresholdsOperatorGreaterOrEquals, TestListResultsResponseLastUnskippedResultGoalThresholdsOperatorLess, TestListResultsResponseLastUnskippedResultGoalThresholdsOperatorLessOrEquals, TestListResultsResponseLastUnskippedResultGoalThresholdsOperatorNotEquals:
		return true
	}
	return false
}

// Whether to use automatic anomaly detection or manual thresholds
type TestListResultsResponseLastUnskippedResultGoalThresholdsThresholdMode string

const (
	TestListResultsResponseLastUnskippedResultGoalThresholdsThresholdModeAutomatic TestListResultsResponseLastUnskippedResultGoalThresholdsThresholdMode = "automatic"
	TestListResultsResponseLastUnskippedResultGoalThresholdsThresholdModeManual    TestListResultsResponseLastUnskippedResultGoalThresholdsThresholdMode = "manual"
)

func (r TestListResultsResponseLastUnskippedResultGoalThresholdsThresholdMode) IsKnown() bool {
	switch r {
	case TestListResultsResponseLastUnskippedResultGoalThresholdsThresholdModeAutomatic, TestListResultsResponseLastUnskippedResultGoalThresholdsThresholdModeManual:
		return true
	}
	return false
}

// The value to be compared.
//
// Union satisfied by [shared.UnionFloat], [shared.UnionBool], [shared.UnionString]
// or [TestListResultsResponseLastUnskippedResultGoalThresholdsValueArray].
type TestListResultsResponseLastUnskippedResultGoalThresholdsValueUnion interface {
	ImplementsTestListResultsResponseLastUnskippedResultGoalThresholdsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TestListResultsResponseLastUnskippedResultGoalThresholdsValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(TestListResultsResponseLastUnskippedResultGoalThresholdsValueArray{}),
		},
	)
}

type TestListResultsResponseLastUnskippedResultGoalThresholdsValueArray []string

func (r TestListResultsResponseLastUnskippedResultGoalThresholdsValueArray) ImplementsTestListResultsResponseLastUnskippedResultGoalThresholdsValueUnion() {
}

// The test type.
type TestListResultsResponseLastUnskippedResultGoalType string

const (
	TestListResultsResponseLastUnskippedResultGoalTypeIntegrity   TestListResultsResponseLastUnskippedResultGoalType = "integrity"
	TestListResultsResponseLastUnskippedResultGoalTypeConsistency TestListResultsResponseLastUnskippedResultGoalType = "consistency"
	TestListResultsResponseLastUnskippedResultGoalTypePerformance TestListResultsResponseLastUnskippedResultGoalType = "performance"
)

func (r TestListResultsResponseLastUnskippedResultGoalType) IsKnown() bool {
	switch r {
	case TestListResultsResponseLastUnskippedResultGoalTypeIntegrity, TestListResultsResponseLastUnskippedResultGoalTypeConsistency, TestListResultsResponseLastUnskippedResultGoalTypePerformance:
		return true
	}
	return false
}

// The body of the rows request.
type TestListResultsResponseLastUnskippedResultRowsBody struct {
	ColumnFilters     []TestListResultsResponseLastUnskippedResultRowsBodyColumnFilter `json:"columnFilters,nullable"`
	ExcludeRowIDList  []int64                                                          `json:"excludeRowIdList,nullable"`
	NotSearchQueryAnd []string                                                         `json:"notSearchQueryAnd,nullable"`
	NotSearchQueryOr  []string                                                         `json:"notSearchQueryOr,nullable"`
	RowIDList         []int64                                                          `json:"rowIdList,nullable"`
	SearchQueryAnd    []string                                                         `json:"searchQueryAnd,nullable"`
	SearchQueryOr     []string                                                         `json:"searchQueryOr,nullable"`
	JSON              testListResultsResponseLastUnskippedResultRowsBodyJSON           `json:"-"`
}

// testListResultsResponseLastUnskippedResultRowsBodyJSON contains the JSON
// metadata for the struct [TestListResultsResponseLastUnskippedResultRowsBody]
type testListResultsResponseLastUnskippedResultRowsBodyJSON struct {
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

func (r *TestListResultsResponseLastUnskippedResultRowsBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResultsResponseLastUnskippedResultRowsBodyJSON) RawJSON() string {
	return r.raw
}

type TestListResultsResponseLastUnskippedResultRowsBodyColumnFilter struct {
	// The name of the column.
	Measurement string                                                                  `json:"measurement,required"`
	Operator    TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperator `json:"operator,required"`
	// This field can have the runtime type of
	// [[]TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterValueUnion],
	// [float64],
	// [TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilterValueUnion].
	Value interface{}                                                        `json:"value,required"`
	JSON  testListResultsResponseLastUnskippedResultRowsBodyColumnFilterJSON `json:"-"`
	union TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersUnion
}

// testListResultsResponseLastUnskippedResultRowsBodyColumnFilterJSON contains the
// JSON metadata for the struct
// [TestListResultsResponseLastUnskippedResultRowsBodyColumnFilter]
type testListResultsResponseLastUnskippedResultRowsBodyColumnFilterJSON struct {
	Measurement apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r testListResultsResponseLastUnskippedResultRowsBodyColumnFilterJSON) RawJSON() string {
	return r.raw
}

func (r *TestListResultsResponseLastUnskippedResultRowsBodyColumnFilter) UnmarshalJSON(data []byte) (err error) {
	*r = TestListResultsResponseLastUnskippedResultRowsBodyColumnFilter{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilter],
// [TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilter],
// [TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilter].
func (r TestListResultsResponseLastUnskippedResultRowsBodyColumnFilter) AsUnion() TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersUnion {
	return r.union
}

// Union satisfied by
// [TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilter],
// [TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilter]
// or
// [TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilter].
type TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersUnion interface {
	implementsTestListResultsResponseLastUnskippedResultRowsBodyColumnFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilter{}),
		},
	)
}

type TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilter struct {
	// The name of the column.
	Measurement string                                                                                     `json:"measurement,required"`
	Operator    TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterOperator     `json:"operator,required"`
	Value       []TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterValueUnion `json:"value,required"`
	JSON        testListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterJSON         `json:"-"`
}

// testListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterJSON
// contains the JSON metadata for the struct
// [TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilter]
type testListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterJSON struct {
	Measurement apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterJSON) RawJSON() string {
	return r.raw
}

func (r TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilter) implementsTestListResultsResponseLastUnskippedResultRowsBodyColumnFilter() {
}

type TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterOperator string

const (
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterOperatorContainsNone TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterOperator = "contains_none"
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterOperatorContainsAny  TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterOperator = "contains_any"
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterOperatorContainsAll  TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterOperator = "contains_all"
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterOperatorOneOf        TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterOperator = "one_of"
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterOperatorNoneOf       TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterOperator = "none_of"
)

func (r TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterOperator) IsKnown() bool {
	switch r {
	case TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterOperatorContainsNone, TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterOperatorContainsAny, TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterOperatorContainsAll, TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterOperatorOneOf, TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterOperatorNoneOf:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterValueUnion interface {
	ImplementsTestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersSetColumnFilterValueUnion)(nil)).Elem(),
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

type TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilter struct {
	// The name of the column.
	Measurement string                                                                                     `json:"measurement,required"`
	Operator    TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterOperator `json:"operator,required"`
	Value       float64                                                                                    `json:"value,required,nullable"`
	JSON        testListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterJSON     `json:"-"`
}

// testListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterJSON
// contains the JSON metadata for the struct
// [TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilter]
type testListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterJSON struct {
	Measurement apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterJSON) RawJSON() string {
	return r.raw
}

func (r TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilter) implementsTestListResultsResponseLastUnskippedResultRowsBodyColumnFilter() {
}

type TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterOperator string

const (
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterOperatorGreater         TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterOperator = ">"
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterOperatorGreaterOrEquals TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterOperator = ">="
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterOperatorIs              TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterOperator = "is"
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterOperatorLess            TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterOperator = "<"
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterOperatorLessOrEquals    TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterOperator = "<="
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterOperatorNotEquals       TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterOperator = "!="
)

func (r TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterOperator) IsKnown() bool {
	switch r {
	case TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterOperatorGreater, TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterOperatorGreaterOrEquals, TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterOperatorIs, TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterOperatorLess, TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterOperatorLessOrEquals, TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersNumericColumnFilterOperatorNotEquals:
		return true
	}
	return false
}

type TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilter struct {
	// The name of the column.
	Measurement string                                                                                      `json:"measurement,required"`
	Operator    TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilterOperator   `json:"operator,required"`
	Value       TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilterValueUnion `json:"value,required"`
	JSON        testListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilterJSON       `json:"-"`
}

// testListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilterJSON
// contains the JSON metadata for the struct
// [TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilter]
type testListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilterJSON struct {
	Measurement apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilterJSON) RawJSON() string {
	return r.raw
}

func (r TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilter) implementsTestListResultsResponseLastUnskippedResultRowsBodyColumnFilter() {
}

type TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilterOperator string

const (
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilterOperatorIs        TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilterOperator = "is"
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilterOperatorNotEquals TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilterOperator = "!="
)

func (r TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilterOperator) IsKnown() bool {
	switch r {
	case TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilterOperatorIs, TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilterOperatorNotEquals:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilterValueUnion interface {
	ImplementsTestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilterValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersStringColumnFilterValueUnion)(nil)).Elem(),
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

type TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperator string

const (
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperatorContainsNone    TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperator = "contains_none"
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperatorContainsAny     TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperator = "contains_any"
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperatorContainsAll     TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperator = "contains_all"
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperatorOneOf           TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperator = "one_of"
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperatorNoneOf          TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperator = "none_of"
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperatorGreater         TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperator = ">"
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperatorGreaterOrEquals TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperator = ">="
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperatorIs              TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperator = "is"
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperatorLess            TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperator = "<"
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperatorLessOrEquals    TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperator = "<="
	TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperatorNotEquals       TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperator = "!="
)

func (r TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperator) IsKnown() bool {
	switch r {
	case TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperatorContainsNone, TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperatorContainsAny, TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperatorContainsAll, TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperatorOneOf, TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperatorNoneOf, TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperatorGreater, TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperatorGreaterOrEquals, TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperatorIs, TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperatorLess, TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperatorLessOrEquals, TestListResultsResponseLastUnskippedResultRowsBodyColumnFiltersOperatorNotEquals:
		return true
	}
	return false
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

type TestListResultsParams struct {
	// Filter for results that use data starting before the end timestamp.
	EndTimestamp param.Field[float64] `query:"endTimestamp"`
	// Include the insights linked to each test result
	IncludeInsights param.Field[bool] `query:"includeInsights"`
	// Retrive test results for a specific inference pipeline.
	InferencePipelineID param.Field[string] `query:"inferencePipelineId" format:"uuid"`
	// The page to return in a paginated query.
	Page param.Field[int64] `query:"page"`
	// Maximum number of items to return per page.
	PerPage param.Field[int64] `query:"perPage"`
	// Retrive test results for a specific project version.
	ProjectVersionID param.Field[string] `query:"projectVersionId" format:"uuid"`
	// Filter for results that use data ending after the start timestamp.
	StartTimestamp param.Field[float64] `query:"startTimestamp"`
	// Filter by status(es).
	Status param.Field[[]string] `query:"status"`
}

// URLQuery serializes [TestListResultsParams]'s query parameters as `url.Values`.
func (r TestListResultsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
