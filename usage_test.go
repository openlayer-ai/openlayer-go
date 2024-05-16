// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer_test

import (
  "context"
  "os"
  "testing"

  "github.com/stainless-sdks/openlayer-go"
  "github.com/stainless-sdks/openlayer-go/internal/testutil"
  "github.com/stainless-sdks/openlayer-go/option"
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
      Config: openlayer.F[openlayer.InferencePipelineDataStreamParamsConfigUnion](openlayer.InferencePipelineDataStreamParamsConfigLlmDataConfig{
        Number0: "R",
        Number1: "E",
        Number2: "P",
        Number3: "L",
        Number4: "A",
        Number5: "C",
        Number6: "E",
        Number7: "_",
        Number8: "M",
        Number9: "E",
      }),
      Rows: openlayer.F([]),
    },
  )
  if err != nil {
    t.Error(err)
  }
  t.Logf("%+v\n", inferencePipelineDataStreamResponse.Success)
}
