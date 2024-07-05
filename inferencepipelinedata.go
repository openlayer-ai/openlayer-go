// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/openlayer-ai/openlayer-go/internal/apijson"
	"github.com/openlayer-ai/openlayer-go/internal/param"
	"github.com/openlayer-ai/openlayer-go/internal/requestconfig"
	"github.com/openlayer-ai/openlayer-go/option"
)

// InferencePipelineDataService contains methods and other services that help with
// interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInferencePipelineDataService] method instead.
type InferencePipelineDataService struct {
	Options []option.RequestOption
}

// NewInferencePipelineDataService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInferencePipelineDataService(opts ...option.RequestOption) (r *InferencePipelineDataService) {
	r = &InferencePipelineDataService{}
	r.Options = opts
	return
}

// Stream production data to an inference pipeline in Openlayer.
func (r *InferencePipelineDataService) Stream(ctx context.Context, id string, body InferencePipelineDataStreamParams, opts ...option.RequestOption) (res *InferencePipelineDataStreamResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("inference-pipelines/%s/data-stream", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type InferencePipelineDataStreamResponse struct {
	Success InferencePipelineDataStreamResponseSuccess `json:"success,required"`
	JSON    inferencePipelineDataStreamResponseJSON    `json:"-"`
}

// inferencePipelineDataStreamResponseJSON contains the JSON metadata for the
// struct [InferencePipelineDataStreamResponse]
type inferencePipelineDataStreamResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InferencePipelineDataStreamResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferencePipelineDataStreamResponseJSON) RawJSON() string {
	return r.raw
}

type InferencePipelineDataStreamResponseSuccess bool

const (
	InferencePipelineDataStreamResponseSuccessTrue InferencePipelineDataStreamResponseSuccess = true
)

func (r InferencePipelineDataStreamResponseSuccess) IsKnown() bool {
	switch r {
	case InferencePipelineDataStreamResponseSuccessTrue:
		return true
	}
	return false
}

type InferencePipelineDataStreamParams struct {
	// Configuration for the data stream. Depends on your **Openlayer project task
	// type**.
	Config param.Field[InferencePipelineDataStreamParamsConfigUnion] `json:"config,required"`
	// A list of entries that represent rows of a csv file
	Rows param.Field[[]map[string]interface{}] `json:"rows,required"`
}

func (r InferencePipelineDataStreamParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for the data stream. Depends on your **Openlayer project task
// type**.
type InferencePipelineDataStreamParamsConfig struct {
	// Name of the column with the total number of tokens.
	NumOfTokenColumnName param.Field[string] `json:"numOfTokenColumnName"`
	// Name of the column with the context retrieved. Applies to RAG use cases.
	// Providing the context enables RAG-specific metrics.
	ContextColumnName param.Field[string] `json:"contextColumnName"`
	// Name of the column with the cost associated with each row.
	CostColumnName param.Field[string] `json:"costColumnName"`
	// Name of the column with the ground truths.
	GroundTruthColumnName param.Field[string] `json:"groundTruthColumnName"`
	// Name of the column with the inference ids. This is useful if you want to update
	// rows at a later point in time. If not provided, a unique id is generated by
	// Openlayer.
	InferenceIDColumnName param.Field[string]      `json:"inferenceIdColumnName"`
	InputVariableNames    param.Field[interface{}] `json:"inputVariableNames,required"`
	// Name of the column with the latencies.
	LatencyColumnName param.Field[string]      `json:"latencyColumnName"`
	Metadata          param.Field[interface{}] `json:"metadata,required"`
	// Name of the column with the model outputs.
	OutputColumnName param.Field[string]      `json:"outputColumnName"`
	Prompt           param.Field[interface{}] `json:"prompt,required"`
	// Name of the column with the questions. Applies to RAG use cases. Providing the
	// question enables RAG-specific metrics.
	QuestionColumnName param.Field[string] `json:"questionColumnName"`
	// Name of the column with the timestamps. Timestamps must be in UNIX sec format.
	// If not provided, the upload timestamp is used.
	TimestampColumnName     param.Field[string]      `json:"timestampColumnName"`
	CategoricalFeatureNames param.Field[interface{}] `json:"categoricalFeatureNames,required"`
	ClassNames              param.Field[interface{}] `json:"classNames,required"`
	FeatureNames            param.Field[interface{}] `json:"featureNames,required"`
	// Name of the column with the labels. The data in this column must be
	// **zero-indexed integers**, matching the list provided in `classNames`.
	LabelColumnName param.Field[string] `json:"labelColumnName"`
	// Name of the column with the model's predictions as **zero-indexed integers**.
	PredictionsColumnName param.Field[string] `json:"predictionsColumnName"`
	// Name of the column with the model's predictions as **lists of class
	// probabilities**.
	PredictionScoresColumnName param.Field[string] `json:"predictionScoresColumnName"`
	// Name of the column with the targets (ground truth values).
	TargetColumnName param.Field[string] `json:"targetColumnName"`
	// Name of the column with the text data.
	TextColumnName param.Field[string] `json:"textColumnName"`
}

func (r InferencePipelineDataStreamParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InferencePipelineDataStreamParamsConfig) implementsInferencePipelineDataStreamParamsConfigUnion() {
}

// Configuration for the data stream. Depends on your **Openlayer project task
// type**.
//
// Satisfied by [InferencePipelineDataStreamParamsConfigLlmData],
// [InferencePipelineDataStreamParamsConfigTabularClassificationData],
// [InferencePipelineDataStreamParamsConfigTabularRegressionData],
// [InferencePipelineDataStreamParamsConfigTextClassificationData],
// [InferencePipelineDataStreamParamsConfig].
type InferencePipelineDataStreamParamsConfigUnion interface {
	implementsInferencePipelineDataStreamParamsConfigUnion()
}

type InferencePipelineDataStreamParamsConfigLlmData struct {
	// Name of the column with the model outputs.
	OutputColumnName param.Field[string] `json:"outputColumnName,required"`
	// Name of the column with the context retrieved. Applies to RAG use cases.
	// Providing the context enables RAG-specific metrics.
	ContextColumnName param.Field[string] `json:"contextColumnName"`
	// Name of the column with the cost associated with each row.
	CostColumnName param.Field[string] `json:"costColumnName"`
	// Name of the column with the ground truths.
	GroundTruthColumnName param.Field[string] `json:"groundTruthColumnName"`
	// Name of the column with the inference ids. This is useful if you want to update
	// rows at a later point in time. If not provided, a unique id is generated by
	// Openlayer.
	InferenceIDColumnName param.Field[string] `json:"inferenceIdColumnName"`
	// Array of input variable names. Each input variable should be a dataset column.
	InputVariableNames param.Field[[]string] `json:"inputVariableNames"`
	// Name of the column with the latencies.
	LatencyColumnName param.Field[string] `json:"latencyColumnName"`
	// Object with metadata.
	Metadata param.Field[interface{}] `json:"metadata"`
	// Name of the column with the total number of tokens.
	NumOfTokenColumnName param.Field[string] `json:"numOfTokenColumnName"`
	// Prompt for the LLM.
	Prompt param.Field[[]InferencePipelineDataStreamParamsConfigLlmDataPrompt] `json:"prompt"`
	// Name of the column with the questions. Applies to RAG use cases. Providing the
	// question enables RAG-specific metrics.
	QuestionColumnName param.Field[string] `json:"questionColumnName"`
	// Name of the column with the timestamps. Timestamps must be in UNIX sec format.
	// If not provided, the upload timestamp is used.
	TimestampColumnName param.Field[string] `json:"timestampColumnName"`
}

func (r InferencePipelineDataStreamParamsConfigLlmData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InferencePipelineDataStreamParamsConfigLlmData) implementsInferencePipelineDataStreamParamsConfigUnion() {
}

type InferencePipelineDataStreamParamsConfigLlmDataPrompt struct {
	// Content of the prompt.
	Content param.Field[string] `json:"content"`
	// Role of the prompt.
	Role param.Field[string] `json:"role"`
}

func (r InferencePipelineDataStreamParamsConfigLlmDataPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InferencePipelineDataStreamParamsConfigTabularClassificationData struct {
	// List of class names indexed by label integer in the dataset. E.g. ["Retained",
	// "Exited"] when 0, 1 are in your label column.
	ClassNames param.Field[[]string] `json:"classNames,required"`
	// Array with the names of all categorical features in the dataset. E.g. ["Gender",
	// "Geography"].
	CategoricalFeatureNames param.Field[[]string] `json:"categoricalFeatureNames"`
	// Array with all input feature names.
	FeatureNames param.Field[[]string] `json:"featureNames"`
	// Name of the column with the inference ids. This is useful if you want to update
	// rows at a later point in time. If not provided, a unique id is generated by
	// Openlayer.
	InferenceIDColumnName param.Field[string] `json:"inferenceIdColumnName"`
	// Name of the column with the labels. The data in this column must be
	// **zero-indexed integers**, matching the list provided in `classNames`.
	LabelColumnName param.Field[string] `json:"labelColumnName"`
	// Name of the column with the latencies.
	LatencyColumnName param.Field[string] `json:"latencyColumnName"`
	// Object with metadata.
	Metadata param.Field[interface{}] `json:"metadata"`
	// Name of the column with the model's predictions as **zero-indexed integers**.
	PredictionsColumnName param.Field[string] `json:"predictionsColumnName"`
	// Name of the column with the model's predictions as **lists of class
	// probabilities**.
	PredictionScoresColumnName param.Field[string] `json:"predictionScoresColumnName"`
	// Name of the column with the timestamps. Timestamps must be in UNIX sec format.
	// If not provided, the upload timestamp is used.
	TimestampColumnName param.Field[string] `json:"timestampColumnName"`
}

func (r InferencePipelineDataStreamParamsConfigTabularClassificationData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InferencePipelineDataStreamParamsConfigTabularClassificationData) implementsInferencePipelineDataStreamParamsConfigUnion() {
}

type InferencePipelineDataStreamParamsConfigTabularRegressionData struct {
	// Array with the names of all categorical features in the dataset. E.g. ["Gender",
	// "Geography"].
	CategoricalFeatureNames param.Field[[]string] `json:"categoricalFeatureNames"`
	// Array with all input feature names.
	FeatureNames param.Field[[]string] `json:"featureNames"`
	// Name of the column with the inference ids. This is useful if you want to update
	// rows at a later point in time. If not provided, a unique id is generated by
	// Openlayer.
	InferenceIDColumnName param.Field[string] `json:"inferenceIdColumnName"`
	// Name of the column with the latencies.
	LatencyColumnName param.Field[string] `json:"latencyColumnName"`
	// Object with metadata.
	Metadata param.Field[interface{}] `json:"metadata"`
	// Name of the column with the model's predictions.
	PredictionsColumnName param.Field[string] `json:"predictionsColumnName"`
	// Name of the column with the targets (ground truth values).
	TargetColumnName param.Field[string] `json:"targetColumnName"`
	// Name of the column with the timestamps. Timestamps must be in UNIX sec format.
	// If not provided, the upload timestamp is used.
	TimestampColumnName param.Field[string] `json:"timestampColumnName"`
}

func (r InferencePipelineDataStreamParamsConfigTabularRegressionData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InferencePipelineDataStreamParamsConfigTabularRegressionData) implementsInferencePipelineDataStreamParamsConfigUnion() {
}

type InferencePipelineDataStreamParamsConfigTextClassificationData struct {
	// List of class names indexed by label integer in the dataset. E.g. ["Retained",
	// "Exited"] when 0, 1 are in your label column.
	ClassNames param.Field[[]string] `json:"classNames,required"`
	// Name of the column with the inference ids. This is useful if you want to update
	// rows at a later point in time. If not provided, a unique id is generated by
	// Openlayer.
	InferenceIDColumnName param.Field[string] `json:"inferenceIdColumnName"`
	// Name of the column with the labels. The data in this column must be
	// **zero-indexed integers**, matching the list provided in `classNames`.
	LabelColumnName param.Field[string] `json:"labelColumnName"`
	// Name of the column with the latencies.
	LatencyColumnName param.Field[string] `json:"latencyColumnName"`
	// Object with metadata.
	Metadata param.Field[interface{}] `json:"metadata"`
	// Name of the column with the model's predictions as **zero-indexed integers**.
	PredictionsColumnName param.Field[string] `json:"predictionsColumnName"`
	// Name of the column with the model's predictions as **lists of class
	// probabilities**.
	PredictionScoresColumnName param.Field[string] `json:"predictionScoresColumnName"`
	// Name of the column with the text data.
	TextColumnName param.Field[string] `json:"textColumnName"`
	// Name of the column with the timestamps. Timestamps must be in UNIX sec format.
	// If not provided, the upload timestamp is used.
	TimestampColumnName param.Field[string] `json:"timestampColumnName"`
}

func (r InferencePipelineDataStreamParamsConfigTextClassificationData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InferencePipelineDataStreamParamsConfigTextClassificationData) implementsInferencePipelineDataStreamParamsConfigUnion() {
}
