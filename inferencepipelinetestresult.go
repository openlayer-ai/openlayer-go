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

// List the latest test results for an inference pipeline.
func (r *InferencePipelineTestResultService) List(ctx context.Context, inferencePipelineID string, query InferencePipelineTestResultListParams, opts ...option.RequestOption) (res *InferencePipelineTestResultListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if inferencePipelineID == "" {
		err = errors.New("missing required inferencePipelineId parameter")
		return
	}
	path := fmt.Sprintf("inference-pipelines/%s/results", inferencePipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type InferencePipelineTestResultListResponse struct {
	Items []InferencePipelineTestResultListResponseItem `json:"items,required"`
	JSON  inferencePipelineTestResultListResponseJSON   `json:"-"`
}

// inferencePipelineTestResultListResponseJSON contains the JSON metadata for the
// struct [InferencePipelineTestResultListResponse]
type inferencePipelineTestResultListResponseJSON struct {
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
	Subtype InferencePipelineTestResultListResponseItemsGoalSubtype `json:"subtype,required"`
	// Whether the test is suggested or user-created.
	Suggested  bool                                                        `json:"suggested,required"`
	Thresholds []InferencePipelineTestResultListResponseItemsGoalThreshold `json:"thresholds,required"`
	// The test type.
	Type InferencePipelineTestResultListResponseItemsGoalType `json:"type,required"`
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

// The test subtype.
type InferencePipelineTestResultListResponseItemsGoalSubtype string

const (
	InferencePipelineTestResultListResponseItemsGoalSubtypeAnomalousColumnCount       InferencePipelineTestResultListResponseItemsGoalSubtype = "anomalousColumnCount"
	InferencePipelineTestResultListResponseItemsGoalSubtypeCharacterLength            InferencePipelineTestResultListResponseItemsGoalSubtype = "characterLength"
	InferencePipelineTestResultListResponseItemsGoalSubtypeClassImbalanceRatio        InferencePipelineTestResultListResponseItemsGoalSubtype = "classImbalanceRatio"
	InferencePipelineTestResultListResponseItemsGoalSubtypeExpectColumnAToBeInColumnB InferencePipelineTestResultListResponseItemsGoalSubtype = "expectColumnAToBeInColumnB"
	InferencePipelineTestResultListResponseItemsGoalSubtypeColumnAverage              InferencePipelineTestResultListResponseItemsGoalSubtype = "columnAverage"
	InferencePipelineTestResultListResponseItemsGoalSubtypeColumnDrift                InferencePipelineTestResultListResponseItemsGoalSubtype = "columnDrift"
	InferencePipelineTestResultListResponseItemsGoalSubtypeColumnStatistic            InferencePipelineTestResultListResponseItemsGoalSubtype = "columnStatistic"
	InferencePipelineTestResultListResponseItemsGoalSubtypeColumnValuesMatch          InferencePipelineTestResultListResponseItemsGoalSubtype = "columnValuesMatch"
	InferencePipelineTestResultListResponseItemsGoalSubtypeConflictingLabelRowCount   InferencePipelineTestResultListResponseItemsGoalSubtype = "conflictingLabelRowCount"
	InferencePipelineTestResultListResponseItemsGoalSubtypeContainsPii                InferencePipelineTestResultListResponseItemsGoalSubtype = "containsPii"
	InferencePipelineTestResultListResponseItemsGoalSubtypeContainsValidURL           InferencePipelineTestResultListResponseItemsGoalSubtype = "containsValidUrl"
	InferencePipelineTestResultListResponseItemsGoalSubtypeCorrelatedFeatureCount     InferencePipelineTestResultListResponseItemsGoalSubtype = "correlatedFeatureCount"
	InferencePipelineTestResultListResponseItemsGoalSubtypeCustomMetricThreshold      InferencePipelineTestResultListResponseItemsGoalSubtype = "customMetricThreshold"
	InferencePipelineTestResultListResponseItemsGoalSubtypeDuplicateRowCount          InferencePipelineTestResultListResponseItemsGoalSubtype = "duplicateRowCount"
	InferencePipelineTestResultListResponseItemsGoalSubtypeEmptyFeature               InferencePipelineTestResultListResponseItemsGoalSubtype = "emptyFeature"
	InferencePipelineTestResultListResponseItemsGoalSubtypeEmptyFeatureCount          InferencePipelineTestResultListResponseItemsGoalSubtype = "emptyFeatureCount"
	InferencePipelineTestResultListResponseItemsGoalSubtypeDriftedFeatureCount        InferencePipelineTestResultListResponseItemsGoalSubtype = "driftedFeatureCount"
	InferencePipelineTestResultListResponseItemsGoalSubtypeFeatureMissingValues       InferencePipelineTestResultListResponseItemsGoalSubtype = "featureMissingValues"
	InferencePipelineTestResultListResponseItemsGoalSubtypeFeatureValueValidation     InferencePipelineTestResultListResponseItemsGoalSubtype = "featureValueValidation"
	InferencePipelineTestResultListResponseItemsGoalSubtypeGreatExpectations          InferencePipelineTestResultListResponseItemsGoalSubtype = "greatExpectations"
	InferencePipelineTestResultListResponseItemsGoalSubtypeGroupByColumnStatsCheck    InferencePipelineTestResultListResponseItemsGoalSubtype = "groupByColumnStatsCheck"
	InferencePipelineTestResultListResponseItemsGoalSubtypeIllFormedRowCount          InferencePipelineTestResultListResponseItemsGoalSubtype = "illFormedRowCount"
	InferencePipelineTestResultListResponseItemsGoalSubtypeIsCode                     InferencePipelineTestResultListResponseItemsGoalSubtype = "isCode"
	InferencePipelineTestResultListResponseItemsGoalSubtypeIsJson                     InferencePipelineTestResultListResponseItemsGoalSubtype = "isJson"
	InferencePipelineTestResultListResponseItemsGoalSubtypeLlmRubricThresholdV2       InferencePipelineTestResultListResponseItemsGoalSubtype = "llmRubricThresholdV2"
	InferencePipelineTestResultListResponseItemsGoalSubtypeLabelDrift                 InferencePipelineTestResultListResponseItemsGoalSubtype = "labelDrift"
	InferencePipelineTestResultListResponseItemsGoalSubtypeMetricThreshold            InferencePipelineTestResultListResponseItemsGoalSubtype = "metricThreshold"
	InferencePipelineTestResultListResponseItemsGoalSubtypeNewCategoryCount           InferencePipelineTestResultListResponseItemsGoalSubtype = "newCategoryCount"
	InferencePipelineTestResultListResponseItemsGoalSubtypeNewLabelCount              InferencePipelineTestResultListResponseItemsGoalSubtype = "newLabelCount"
	InferencePipelineTestResultListResponseItemsGoalSubtypeNullRowCount               InferencePipelineTestResultListResponseItemsGoalSubtype = "nullRowCount"
	InferencePipelineTestResultListResponseItemsGoalSubtypeRowCount                   InferencePipelineTestResultListResponseItemsGoalSubtype = "rowCount"
	InferencePipelineTestResultListResponseItemsGoalSubtypePpScoreValueValidation     InferencePipelineTestResultListResponseItemsGoalSubtype = "ppScoreValueValidation"
	InferencePipelineTestResultListResponseItemsGoalSubtypeQuasiConstantFeature       InferencePipelineTestResultListResponseItemsGoalSubtype = "quasiConstantFeature"
	InferencePipelineTestResultListResponseItemsGoalSubtypeQuasiConstantFeatureCount  InferencePipelineTestResultListResponseItemsGoalSubtype = "quasiConstantFeatureCount"
	InferencePipelineTestResultListResponseItemsGoalSubtypeSqlQuery                   InferencePipelineTestResultListResponseItemsGoalSubtype = "sqlQuery"
	InferencePipelineTestResultListResponseItemsGoalSubtypeDtypeValidation            InferencePipelineTestResultListResponseItemsGoalSubtype = "dtypeValidation"
	InferencePipelineTestResultListResponseItemsGoalSubtypeSentenceLength             InferencePipelineTestResultListResponseItemsGoalSubtype = "sentenceLength"
	InferencePipelineTestResultListResponseItemsGoalSubtypeSizeRatio                  InferencePipelineTestResultListResponseItemsGoalSubtype = "sizeRatio"
	InferencePipelineTestResultListResponseItemsGoalSubtypeSpecialCharactersRatio     InferencePipelineTestResultListResponseItemsGoalSubtype = "specialCharactersRatio"
	InferencePipelineTestResultListResponseItemsGoalSubtypeStringValidation           InferencePipelineTestResultListResponseItemsGoalSubtype = "stringValidation"
	InferencePipelineTestResultListResponseItemsGoalSubtypeTrainValLeakageRowCount    InferencePipelineTestResultListResponseItemsGoalSubtype = "trainValLeakageRowCount"
)

func (r InferencePipelineTestResultListResponseItemsGoalSubtype) IsKnown() bool {
	switch r {
	case InferencePipelineTestResultListResponseItemsGoalSubtypeAnomalousColumnCount, InferencePipelineTestResultListResponseItemsGoalSubtypeCharacterLength, InferencePipelineTestResultListResponseItemsGoalSubtypeClassImbalanceRatio, InferencePipelineTestResultListResponseItemsGoalSubtypeExpectColumnAToBeInColumnB, InferencePipelineTestResultListResponseItemsGoalSubtypeColumnAverage, InferencePipelineTestResultListResponseItemsGoalSubtypeColumnDrift, InferencePipelineTestResultListResponseItemsGoalSubtypeColumnStatistic, InferencePipelineTestResultListResponseItemsGoalSubtypeColumnValuesMatch, InferencePipelineTestResultListResponseItemsGoalSubtypeConflictingLabelRowCount, InferencePipelineTestResultListResponseItemsGoalSubtypeContainsPii, InferencePipelineTestResultListResponseItemsGoalSubtypeContainsValidURL, InferencePipelineTestResultListResponseItemsGoalSubtypeCorrelatedFeatureCount, InferencePipelineTestResultListResponseItemsGoalSubtypeCustomMetricThreshold, InferencePipelineTestResultListResponseItemsGoalSubtypeDuplicateRowCount, InferencePipelineTestResultListResponseItemsGoalSubtypeEmptyFeature, InferencePipelineTestResultListResponseItemsGoalSubtypeEmptyFeatureCount, InferencePipelineTestResultListResponseItemsGoalSubtypeDriftedFeatureCount, InferencePipelineTestResultListResponseItemsGoalSubtypeFeatureMissingValues, InferencePipelineTestResultListResponseItemsGoalSubtypeFeatureValueValidation, InferencePipelineTestResultListResponseItemsGoalSubtypeGreatExpectations, InferencePipelineTestResultListResponseItemsGoalSubtypeGroupByColumnStatsCheck, InferencePipelineTestResultListResponseItemsGoalSubtypeIllFormedRowCount, InferencePipelineTestResultListResponseItemsGoalSubtypeIsCode, InferencePipelineTestResultListResponseItemsGoalSubtypeIsJson, InferencePipelineTestResultListResponseItemsGoalSubtypeLlmRubricThresholdV2, InferencePipelineTestResultListResponseItemsGoalSubtypeLabelDrift, InferencePipelineTestResultListResponseItemsGoalSubtypeMetricThreshold, InferencePipelineTestResultListResponseItemsGoalSubtypeNewCategoryCount, InferencePipelineTestResultListResponseItemsGoalSubtypeNewLabelCount, InferencePipelineTestResultListResponseItemsGoalSubtypeNullRowCount, InferencePipelineTestResultListResponseItemsGoalSubtypeRowCount, InferencePipelineTestResultListResponseItemsGoalSubtypePpScoreValueValidation, InferencePipelineTestResultListResponseItemsGoalSubtypeQuasiConstantFeature, InferencePipelineTestResultListResponseItemsGoalSubtypeQuasiConstantFeatureCount, InferencePipelineTestResultListResponseItemsGoalSubtypeSqlQuery, InferencePipelineTestResultListResponseItemsGoalSubtypeDtypeValidation, InferencePipelineTestResultListResponseItemsGoalSubtypeSentenceLength, InferencePipelineTestResultListResponseItemsGoalSubtypeSizeRatio, InferencePipelineTestResultListResponseItemsGoalSubtypeSpecialCharactersRatio, InferencePipelineTestResultListResponseItemsGoalSubtypeStringValidation, InferencePipelineTestResultListResponseItemsGoalSubtypeTrainValLeakageRowCount:
		return true
	}
	return false
}

type InferencePipelineTestResultListResponseItemsGoalThreshold struct {
	// The insight name to be evaluated.
	InsightName InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName `json:"insightName"`
	// The insight parameters. Required only for some test subtypes. For example, for
	// tests that require a column name, the insight parameters will be [{'name':
	// 'column_name', 'value': 'Age'}]
	InsightParameters []InferencePipelineTestResultListResponseItemsGoalThresholdsInsightParameter `json:"insightParameters,nullable"`
	// The measurement to be evaluated.
	Measurement string `json:"measurement"`
	// The operator to be used for the evaluation.
	Operator InferencePipelineTestResultListResponseItemsGoalThresholdsOperator `json:"operator"`
	// Whether to use automatic anomaly detection or manual thresholds
	ThresholdMode InferencePipelineTestResultListResponseItemsGoalThresholdsThresholdMode `json:"thresholdMode"`
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
	ThresholdMode     apijson.Field
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

// The insight name to be evaluated.
type InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName string

const (
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameCharacterLength            InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "characterLength"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameClassImbalance             InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "classImbalance"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameExpectColumnAToBeInColumnB InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "expectColumnAToBeInColumnB"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameColumnAverage              InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "columnAverage"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameColumnDrift                InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "columnDrift"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameColumnValuesMatch          InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "columnValuesMatch"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameConfidenceDistribution     InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "confidenceDistribution"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameConflictingLabelRowCount   InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "conflictingLabelRowCount"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameContainsPii                InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "containsPii"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameContainsValidURL           InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "containsValidUrl"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameCorrelatedFeatures         InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "correlatedFeatures"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameCustomMetric               InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "customMetric"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameDuplicateRowCount          InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "duplicateRowCount"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameEmptyFeatures              InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "emptyFeatures"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameFeatureDrift               InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "featureDrift"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameFeatureProfile             InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "featureProfile"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameGreatExpectations          InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "greatExpectations"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameGroupByColumnStatsCheck    InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "groupByColumnStatsCheck"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameIllFormedRowCount          InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "illFormedRowCount"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameIsCode                     InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "isCode"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameIsJson                     InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "isJson"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameLlmRubricV2                InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "llmRubricV2"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameLabelDrift                 InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "labelDrift"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameMetrics                    InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "metrics"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameNewCategories              InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "newCategories"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameNewLabels                  InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "newLabels"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameNullRowCount               InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "nullRowCount"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNamePpScore                    InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "ppScore"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameQuasiConstantFeatures      InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "quasiConstantFeatures"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameSentenceLength             InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "sentenceLength"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameSizeRatio                  InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "sizeRatio"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameSpecialCharacters          InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "specialCharacters"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameStringValidation           InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "stringValidation"
	InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameTrainValLeakageRowCount    InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName = "trainValLeakageRowCount"
)

func (r InferencePipelineTestResultListResponseItemsGoalThresholdsInsightName) IsKnown() bool {
	switch r {
	case InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameCharacterLength, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameClassImbalance, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameExpectColumnAToBeInColumnB, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameColumnAverage, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameColumnDrift, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameColumnValuesMatch, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameConfidenceDistribution, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameConflictingLabelRowCount, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameContainsPii, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameContainsValidURL, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameCorrelatedFeatures, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameCustomMetric, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameDuplicateRowCount, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameEmptyFeatures, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameFeatureDrift, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameFeatureProfile, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameGreatExpectations, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameGroupByColumnStatsCheck, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameIllFormedRowCount, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameIsCode, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameIsJson, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameLlmRubricV2, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameLabelDrift, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameMetrics, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameNewCategories, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameNewLabels, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameNullRowCount, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNamePpScore, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameQuasiConstantFeatures, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameSentenceLength, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameSizeRatio, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameSpecialCharacters, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameStringValidation, InferencePipelineTestResultListResponseItemsGoalThresholdsInsightNameTrainValLeakageRowCount:
		return true
	}
	return false
}

type InferencePipelineTestResultListResponseItemsGoalThresholdsInsightParameter struct {
	// The name of the insight filter.
	Name  string                                                                         `json:"name,required"`
	Value interface{}                                                                    `json:"value,required"`
	JSON  inferencePipelineTestResultListResponseItemsGoalThresholdsInsightParameterJSON `json:"-"`
}

// inferencePipelineTestResultListResponseItemsGoalThresholdsInsightParameterJSON
// contains the JSON metadata for the struct
// [InferencePipelineTestResultListResponseItemsGoalThresholdsInsightParameter]
type inferencePipelineTestResultListResponseItemsGoalThresholdsInsightParameterJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InferencePipelineTestResultListResponseItemsGoalThresholdsInsightParameter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineTestResultListResponseItemsGoalThresholdsInsightParameterJSON) RawJSON() string {
	return r.raw
}

// The operator to be used for the evaluation.
type InferencePipelineTestResultListResponseItemsGoalThresholdsOperator string

const (
	InferencePipelineTestResultListResponseItemsGoalThresholdsOperatorIs              InferencePipelineTestResultListResponseItemsGoalThresholdsOperator = "is"
	InferencePipelineTestResultListResponseItemsGoalThresholdsOperatorGreater         InferencePipelineTestResultListResponseItemsGoalThresholdsOperator = ">"
	InferencePipelineTestResultListResponseItemsGoalThresholdsOperatorGreaterOrEquals InferencePipelineTestResultListResponseItemsGoalThresholdsOperator = ">="
	InferencePipelineTestResultListResponseItemsGoalThresholdsOperatorLess            InferencePipelineTestResultListResponseItemsGoalThresholdsOperator = "<"
	InferencePipelineTestResultListResponseItemsGoalThresholdsOperatorLessOrEquals    InferencePipelineTestResultListResponseItemsGoalThresholdsOperator = "<="
	InferencePipelineTestResultListResponseItemsGoalThresholdsOperatorNotEquals       InferencePipelineTestResultListResponseItemsGoalThresholdsOperator = "!="
)

func (r InferencePipelineTestResultListResponseItemsGoalThresholdsOperator) IsKnown() bool {
	switch r {
	case InferencePipelineTestResultListResponseItemsGoalThresholdsOperatorIs, InferencePipelineTestResultListResponseItemsGoalThresholdsOperatorGreater, InferencePipelineTestResultListResponseItemsGoalThresholdsOperatorGreaterOrEquals, InferencePipelineTestResultListResponseItemsGoalThresholdsOperatorLess, InferencePipelineTestResultListResponseItemsGoalThresholdsOperatorLessOrEquals, InferencePipelineTestResultListResponseItemsGoalThresholdsOperatorNotEquals:
		return true
	}
	return false
}

// Whether to use automatic anomaly detection or manual thresholds
type InferencePipelineTestResultListResponseItemsGoalThresholdsThresholdMode string

const (
	InferencePipelineTestResultListResponseItemsGoalThresholdsThresholdModeAutomatic InferencePipelineTestResultListResponseItemsGoalThresholdsThresholdMode = "automatic"
	InferencePipelineTestResultListResponseItemsGoalThresholdsThresholdModeManual    InferencePipelineTestResultListResponseItemsGoalThresholdsThresholdMode = "manual"
)

func (r InferencePipelineTestResultListResponseItemsGoalThresholdsThresholdMode) IsKnown() bool {
	switch r {
	case InferencePipelineTestResultListResponseItemsGoalThresholdsThresholdModeAutomatic, InferencePipelineTestResultListResponseItemsGoalThresholdsThresholdModeManual:
		return true
	}
	return false
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

// The test type.
type InferencePipelineTestResultListResponseItemsGoalType string

const (
	InferencePipelineTestResultListResponseItemsGoalTypeIntegrity   InferencePipelineTestResultListResponseItemsGoalType = "integrity"
	InferencePipelineTestResultListResponseItemsGoalTypeConsistency InferencePipelineTestResultListResponseItemsGoalType = "consistency"
	InferencePipelineTestResultListResponseItemsGoalTypePerformance InferencePipelineTestResultListResponseItemsGoalType = "performance"
)

func (r InferencePipelineTestResultListResponseItemsGoalType) IsKnown() bool {
	switch r {
	case InferencePipelineTestResultListResponseItemsGoalTypeIntegrity, InferencePipelineTestResultListResponseItemsGoalTypeConsistency, InferencePipelineTestResultListResponseItemsGoalTypePerformance:
		return true
	}
	return false
}

type InferencePipelineTestResultListParams struct {
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
