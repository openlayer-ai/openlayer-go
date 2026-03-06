// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/openlayer-ai/openlayer-go"
	"github.com/openlayer-ai/openlayer-go/internal/testutil"
	"github.com/openlayer-ai/openlayer-go/option"
	"github.com/openlayer-ai/openlayer-go/shared"
)

func TestInferencePipelineRowNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := openlayer.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.InferencePipelines.Rows.New(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.InferencePipelineRowNewParams{
			Asc:        openlayer.F(true),
			Page:       openlayer.F(int64(1)),
			PerPage:    openlayer.F(int64(1)),
			SortColumn: openlayer.F("sortColumn"),
			ColumnFilters: openlayer.F([]openlayer.InferencePipelineRowNewParamsColumnFilterUnion{openlayer.InferencePipelineRowNewParamsColumnFiltersSetColumnFilter{
				Measurement: openlayer.F("openlayer_token_set"),
				Operator:    openlayer.F(openlayer.InferencePipelineRowNewParamsColumnFiltersSetColumnFilterOperatorContainsNone),
				Value:       openlayer.F([]openlayer.InferencePipelineRowNewParamsColumnFiltersSetColumnFilterValueUnion{shared.UnionString("cat")}),
			}}),
			ExcludeRowIDList:  openlayer.F([]int64{int64(0)}),
			NotSearchQueryAnd: openlayer.F([]string{"string"}),
			NotSearchQueryOr:  openlayer.F([]string{"string"}),
			RowIDList:         openlayer.F([]int64{int64(0)}),
			SearchQueryAnd:    openlayer.F([]string{"string"}),
			SearchQueryOr:     openlayer.F([]string{"string"}),
		},
	)
	if err != nil {
		var apierr *openlayer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInferencePipelineRowUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := openlayer.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.InferencePipelines.Rows.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.InferencePipelineRowUpdateParams{
			InferenceID: openlayer.F("inferenceId"),
			Row:         openlayer.F[any](map[string]interface{}{}),
			Config: openlayer.F(openlayer.InferencePipelineRowUpdateParamsConfig{
				GroundTruthColumnName:   openlayer.F("ground_truth"),
				HumanFeedbackColumnName: openlayer.F("human_feedback"),
				InferenceIDColumnName:   openlayer.F("id"),
				LatencyColumnName:       openlayer.F("latency"),
				TimestampColumnName:     openlayer.F("timestamp"),
			}),
		},
	)
	if err != nil {
		var apierr *openlayer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
