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

func TestProjectListWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.List(context.TODO(), openlayer.ProjectListParams{
		Name:     openlayer.F("string"),
		Page:     openlayer.F(int64(1)),
		PerPage:  openlayer.F(int64(1)),
		TaskType: openlayer.F(openlayer.ProjectListParamsTaskTypeLlmBase),
	})
	if err != nil {
		var apierr *openlayer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
