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
)

func TestProjectInferencePipelineNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.InferencePipelines.New(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.ProjectInferencePipelineNewParams{
			Description: openlayer.F("This pipeline is used for production."),
			Name:        openlayer.F("production"),
			Project: openlayer.F(openlayer.ProjectInferencePipelineNewParamsProject{
				Name:        openlayer.F("My Project"),
				TaskType:    openlayer.F(openlayer.ProjectInferencePipelineNewParamsProjectTaskTypeLlmBase),
				Description: openlayer.F("My project description."),
			}),
			Workspace: openlayer.F(openlayer.ProjectInferencePipelineNewParamsWorkspace{
				Name:            openlayer.F("Openlayer"),
				Slug:            openlayer.F("openlayer"),
				SAMLOnlyAccess:  openlayer.F(true),
				WildcardDomains: openlayer.F([]string{"string"}),
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

func TestProjectInferencePipelineListWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.InferencePipelines.List(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.ProjectInferencePipelineListParams{
			Name:    openlayer.F("name"),
			Page:    openlayer.F(int64(1)),
			PerPage: openlayer.F(int64(1)),
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
