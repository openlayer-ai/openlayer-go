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
	projectNewResponse, err := client.Projects.New(context.TODO(), openlayer.ProjectNewParams{
		Name:     openlayer.F("My Project"),
		TaskType: openlayer.F(openlayer.ProjectNewParamsTaskTypeLlmBase),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", projectNewResponse.ID)
}
