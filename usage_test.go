// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer_test

import (
	"context"
	"os"
	"testing"

	"github.com/openlayer-ai/openlayer-go"
	"github.com/openlayer-ai/openlayer-go/internal/testutil"
	"github.com/openlayer-ai/openlayer-go/option"
)

func TestUsage(t *testing.T) {
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
	inferencePipelineDataStreamResponse, err := client.InferencePipelines.Data.Stream(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.InferencePipelineDataStreamParams{
			Config: openlayer.F[openlayer.InferencePipelineDataStreamParamsConfigUnion](openlayer.InferencePipelineDataStreamParamsConfigLlmData{
				InputVariableNames:   openlayer.F([]string{"user_query"}),
				OutputColumnName:     openlayer.F("output"),
				NumOfTokenColumnName: openlayer.F("tokens"),
				CostColumnName:       openlayer.F("cost"),
				TimestampColumnName:  openlayer.F("timestamp"),
			}),
			Rows: openlayer.F([]map[string]interface{}{{
				"user_query": "what is the meaning of life?",
				"output":     "42",
				"tokens":     map[string]interface{}{},
				"cost":       map[string]interface{}{},
				"timestamp":  map[string]interface{}{},
			}}),
		},
	)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", inferencePipelineDataStreamResponse.Success)
}
