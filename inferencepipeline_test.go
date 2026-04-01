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
	"github.com/openlayer-ai/openlayer-go/shared"
)

func TestInferencePipelineGetWithOptionalParams(t *testing.T) {
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
	_, err := client.InferencePipelines.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.InferencePipelineGetParams{
			Expand: openlayer.F([]openlayer.InferencePipelineGetParamsExpand{openlayer.InferencePipelineGetParamsExpandProject}),
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

func TestInferencePipelineUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.InferencePipelines.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.InferencePipelineUpdateParams{
			Description:         openlayer.F("This pipeline is used for production."),
			Name:                openlayer.F("production"),
			ReferenceDatasetUri: openlayer.F("referenceDatasetUri"),
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

func TestInferencePipelineDelete(t *testing.T) {
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
	err := client.InferencePipelines.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *openlayer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInferencePipelineGetSessionsWithOptionalParams(t *testing.T) {
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
	_, err := client.InferencePipelines.GetSessions(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.InferencePipelineGetSessionsParams{
			Asc:        openlayer.F(true),
			Page:       openlayer.F(int64(1)),
			PerPage:    openlayer.F(int64(1)),
			SortColumn: openlayer.F("sortColumn"),
			ColumnFilters: openlayer.F([]openlayer.InferencePipelineGetSessionsParamsColumnFilterUnion{openlayer.InferencePipelineGetSessionsParamsColumnFiltersSetColumnFilter{
				Measurement: openlayer.F("openlayer_token_set"),
				Operator:    openlayer.F(openlayer.InferencePipelineGetSessionsParamsColumnFiltersSetColumnFilterOperatorContainsNone),
				Value:       openlayer.F([]openlayer.InferencePipelineGetSessionsParamsColumnFiltersSetColumnFilterValueUnion{shared.UnionString("cat")}),
			}}),
			ExcludeRowIDList:  openlayer.F([]int64{int64(0)}),
			NotSearchQueryAnd: openlayer.F([]string{"string"}),
			NotSearchQueryOr:  openlayer.F([]string{"string"}),
			RowIDList:         openlayer.F([]int64{int64(0)}),
			SearchQueryAnd:    openlayer.F([]string{"string"}),
			SearchQueryOr:     openlayer.F([]string{"string"}),
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

func TestInferencePipelineGetUsersWithOptionalParams(t *testing.T) {
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
	_, err := client.InferencePipelines.GetUsers(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.InferencePipelineGetUsersParams{
			Asc:        openlayer.F(true),
			Page:       openlayer.F(int64(1)),
			PerPage:    openlayer.F(int64(1)),
			SortColumn: openlayer.F("sortColumn"),
			ColumnFilters: openlayer.F([]openlayer.InferencePipelineGetUsersParamsColumnFilterUnion{openlayer.InferencePipelineGetUsersParamsColumnFiltersSetColumnFilter{
				Measurement: openlayer.F("openlayer_token_set"),
				Operator:    openlayer.F(openlayer.InferencePipelineGetUsersParamsColumnFiltersSetColumnFilterOperatorContainsNone),
				Value:       openlayer.F([]openlayer.InferencePipelineGetUsersParamsColumnFiltersSetColumnFilterValueUnion{shared.UnionString("cat")}),
			}}),
			ExcludeRowIDList:  openlayer.F([]int64{int64(0)}),
			NotSearchQueryAnd: openlayer.F([]string{"string"}),
			NotSearchQueryOr:  openlayer.F([]string{"string"}),
			RowIDList:         openlayer.F([]int64{int64(0)}),
			SearchQueryAnd:    openlayer.F([]string{"string"}),
			SearchQueryOr:     openlayer.F([]string{"string"}),
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
