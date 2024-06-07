// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomopenlayeraiopenlayergo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/openlayer-ai/openlayer-go"
	"github.com/openlayer-ai/openlayer-go/internal/testutil"
	"github.com/openlayer-ai/openlayer-go/option"
)

func TestInferencePipelineDataStreamWithOptionalParams(t *testing.T) {
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
	_, err := client.InferencePipelines.Data.Stream(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		githubcomopenlayeraiopenlayergo.InferencePipelineDataStreamParams{
			Config: githubcomopenlayeraiopenlayergo.F[githubcomopenlayeraiopenlayergo.InferencePipelineDataStreamParamsConfigUnion](githubcomopenlayeraiopenlayergo.InferencePipelineDataStreamParamsConfigLlmData{
				NumOfTokenColumnName:  githubcomopenlayeraiopenlayergo.F("tokens"),
				ContextColumnName:     githubcomopenlayeraiopenlayergo.F("context"),
				CostColumnName:        githubcomopenlayeraiopenlayergo.F("cost"),
				GroundTruthColumnName: githubcomopenlayeraiopenlayergo.F("ground_truth"),
				InferenceIDColumnName: githubcomopenlayeraiopenlayergo.F("id"),
				InputVariableNames:    githubcomopenlayeraiopenlayergo.F([]string{"user_query"}),
				LatencyColumnName:     githubcomopenlayeraiopenlayergo.F("latency"),
				Metadata:              githubcomopenlayeraiopenlayergo.F[any](map[string]interface{}{}),
				OutputColumnName:      githubcomopenlayeraiopenlayergo.F("output"),
				Prompt: githubcomopenlayeraiopenlayergo.F([]githubcomopenlayeraiopenlayergo.InferencePipelineDataStreamParamsConfigLlmDataPrompt{{
					Role:    githubcomopenlayeraiopenlayergo.F("user"),
					Content: githubcomopenlayeraiopenlayergo.F("{{ user_query }}"),
				}}),
				QuestionColumnName:  githubcomopenlayeraiopenlayergo.F("question"),
				TimestampColumnName: githubcomopenlayeraiopenlayergo.F("timestamp"),
			}),
			Rows: githubcomopenlayeraiopenlayergo.F([]map[string]interface{}{{
				"user_query": "bar",
				"output":     "bar",
				"tokens":     "bar",
				"cost":       "bar",
				"timestamp":  "bar",
			}}),
		},
	)
	if err != nil {
		var apierr *githubcomopenlayeraiopenlayergo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
