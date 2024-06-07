// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomopenlayeraiopenlayergo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/openlayer-go"
	"github.com/stainless-sdks/openlayer-go/internal/testutil"
	"github.com/stainless-sdks/openlayer-go/option"
)

func TestCommitTestResultListWithOptionalParams(t *testing.T) {
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
	_, err := client.Commits.TestResults.List(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		githubcomopenlayeraiopenlayergo.CommitTestResultListParams{
			IncludeArchived: githubcomopenlayeraiopenlayergo.F(true),
			Page:            githubcomopenlayeraiopenlayergo.F(int64(1)),
			PerPage:         githubcomopenlayeraiopenlayergo.F(int64(1)),
			Status:          githubcomopenlayeraiopenlayergo.F(githubcomopenlayeraiopenlayergo.CommitTestResultListParamsStatusPassing),
			Type:            githubcomopenlayeraiopenlayergo.F(githubcomopenlayeraiopenlayergo.CommitTestResultListParamsTypeIntegrity),
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
