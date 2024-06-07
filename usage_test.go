// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomopenlayeraiopenlayergo_test

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
	client := githubcomopenlayeraiopenlayergo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	inferencePipelineDataStreamResponse, err := client.InferencePipelines.Data.Stream(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		githubcomopenlayeraiopenlayergo.InferencePipelineDataStreamParams{
			Config: githubcomopenlayeraiopenlayergo.F[githubcomopenlayeraiopenlayergo.InferencePipelineDataStreamParamsConfigUnion](githubcomopenlayeraiopenlayergo.InferencePipelineDataStreamParamsConfigLlmData{
				InputVariableNames:   githubcomopenlayeraiopenlayergo.F([]string{"user_query"}),
				OutputColumnName:     githubcomopenlayeraiopenlayergo.F("output"),
				NumOfTokenColumnName: githubcomopenlayeraiopenlayergo.F("tokens"),
				CostColumnName:       githubcomopenlayeraiopenlayergo.F("cost"),
				TimestampColumnName:  githubcomopenlayeraiopenlayergo.F("timestamp"),
			}),
			Rows: githubcomopenlayeraiopenlayergo.F([]map[string]interface{}{{
				"user_query": "what's the meaning of life?",
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
