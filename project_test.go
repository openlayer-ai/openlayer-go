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

func TestProjectNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.New(context.TODO(), openlayer.ProjectNewParams{
		Name:        openlayer.F("My Project"),
		TaskType:    openlayer.F(openlayer.ProjectNewParamsTaskTypeLlmBase),
		Description: openlayer.F("My project description."),
		GitRepo: openlayer.F(openlayer.ProjectNewParamsGitRepo{
			GitID:        openlayer.F(int64(0)),
			Branch:       openlayer.F("string"),
			RootDir:      openlayer.F("string"),
			GitAccountID: openlayer.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}),
		SlackChannelID:                   openlayer.F("C01B2PZQX1Z"),
		SlackChannelName:                 openlayer.F("#my-project"),
		SlackChannelNotificationsEnabled: openlayer.F(true),
	})
	if err != nil {
		var apierr *openlayer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

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
