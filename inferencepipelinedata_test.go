// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/openlayer-go"
	"github.com/stainless-sdks/openlayer-go/internal/testutil"
	"github.com/stainless-sdks/openlayer-go/option"
)

func TestInferencePipelineDataStreamWithOptionalParams(t *testing.T) {
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
	_, err := client.InferencePipelines.Data.Stream(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.InferencePipelineDataStreamParams{
			Config: openlayer.F[openlayer.InferencePipelineDataStreamParamsConfigUnion](openlayer.InferencePipelineDataStreamParamsConfigLlmData{
				NumOfTokenColumnName:  openlayer.F("tokens"),
				ContextColumnName:     openlayer.F("context"),
				CostColumnName:        openlayer.F("cost"),
				GroundTruthColumnName: openlayer.F("ground_truth"),
				InferenceIDColumnName: openlayer.F("id"),
				InputVariableNames:    openlayer.F([]string{"user_query"}),
				LatencyColumnName:     openlayer.F("latency"),
				Metadata:              openlayer.F[any](map[string]interface{}{}),
				OutputColumnName:      openlayer.F("output"),
				Prompt: openlayer.F([]openlayer.InferencePipelineDataStreamParamsConfigLlmDataPrompt{{
					Role:    openlayer.F("user"),
					Content: openlayer.F("{{ user_query }}"),
				}}),
				QuestionColumnName:  openlayer.F("question"),
				TimestampColumnName: openlayer.F("timestamp"),
			}),
			Rows: openlayer.F([]map[string]interface{}{{
				"user_query": "bar",
				"output":     "bar",
				"tokens":     "bar",
				"cost":       "bar",
				"timestamp":  "bar",
			}}),
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
