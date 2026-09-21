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

func TestGovernanceFrameworkNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Governance.Frameworks.New(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.GovernanceFrameworkNewParams{
			Name:        openlayer.F("EU AI Act"),
			Description: openlayer.F("Requirements for high-risk AI systems under the EU AI Act."),
			Enabled:     openlayer.F(true),
			ProjectSelector: openlayer.F(openlayer.GovernanceFrameworkNewParamsProjectSelector{
				Match: openlayer.F([]openlayer.GovernanceFrameworkNewParamsProjectSelectorMatch{{
					Property: openlayer.F(openlayer.GovernanceFrameworkNewParamsProjectSelectorMatchPropertyRiskLevel),
					Value: openlayer.F[any]([]string{
						"high",
						"critical",
					}),
					Operator: openlayer.F("operator"),
				}}),
			}),
			Tags: openlayer.F([]string{"regulation", "eu"}),
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

func TestGovernanceFrameworkGet(t *testing.T) {
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
	_, err := client.Governance.Frameworks.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *openlayer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGovernanceFrameworkUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Governance.Frameworks.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.GovernanceFrameworkUpdateParams{
			Avatar: openlayer.F(openlayer.GovernanceFrameworkUpdateParamsAvatar{
				Type:  openlayer.F(openlayer.GovernanceFrameworkUpdateParamsAvatarTypeEmoji),
				Value: openlayer.F("🧭"),
			}),
			Description: openlayer.F("Requirements for high-risk AI systems under the EU AI Act."),
			Enabled:     openlayer.F(true),
			ExtendedDescription: openlayer.F(map[string]interface{}{
				"foo": "bar",
			}),
			Href: openlayer.F("href"),
			Name: openlayer.F("EU AI Act"),
			ProjectSelector: openlayer.F(openlayer.GovernanceFrameworkUpdateParamsProjectSelector{
				Match: openlayer.F([]openlayer.GovernanceFrameworkUpdateParamsProjectSelectorMatch{{
					Property: openlayer.F(openlayer.GovernanceFrameworkUpdateParamsProjectSelectorMatchPropertyRiskLevel),
					Value: openlayer.F[any]([]string{
						"high",
						"critical",
					}),
					Operator: openlayer.F("operator"),
				}}),
			}),
			Tags: openlayer.F([]string{"regulation", "eu"}),
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

func TestGovernanceFrameworkListWithOptionalParams(t *testing.T) {
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
	_, err := client.Governance.Frameworks.List(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.GovernanceFrameworkListParams{
			Asc:                openlayer.F(true),
			CompletionOperator: openlayer.F(openlayer.GovernanceFrameworkListParamsCompletionOperatorIs),
			CompletionValue:    openlayer.F(int64(0)),
			Enabled:            openlayer.F(true),
			IncludeRuleStats:   openlayer.F(true),
			Page:               openlayer.F(int64(1)),
			PerPage:            openlayer.F(int64(1)),
			ProjectID:          openlayer.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			SearchQuery:        openlayer.F("searchQuery"),
			SortColumn:         openlayer.F(openlayer.GovernanceFrameworkListParamsSortColumnName),
			Tags:               openlayer.F([]string{"string"}),
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

func TestGovernanceFrameworkExportWithOptionalParams(t *testing.T) {
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
	_, err := client.Governance.Frameworks.Export(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.GovernanceFrameworkExportParams{
			ProjectID: openlayer.F("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
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

func TestGovernanceFrameworkListProjectRuleStatsWithOptionalParams(t *testing.T) {
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
	_, err := client.Governance.Frameworks.ListProjectRuleStats(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.GovernanceFrameworkListProjectRuleStatsParams{
			Asc:        openlayer.F(true),
			Page:       openlayer.F(int64(1)),
			PerPage:    openlayer.F(int64(1)),
			SortColumn: openlayer.F(openlayer.GovernanceFrameworkListProjectRuleStatsParamsSortColumnProjectName),
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

func TestGovernanceFrameworkListProjectsWithOptionalParams(t *testing.T) {
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
	_, err := client.Governance.Frameworks.ListProjects(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.GovernanceFrameworkListProjectsParams{
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

func TestGovernanceFrameworkListRulesWithOptionalParams(t *testing.T) {
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
	_, err := client.Governance.Frameworks.ListRules(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.GovernanceFrameworkListRulesParams{
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
