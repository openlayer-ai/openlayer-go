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

func TestGovernanceRuleResultGet(t *testing.T) {
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
	_, err := client.Governance.RuleResults.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *openlayer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGovernanceRuleResultUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Governance.RuleResults.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.GovernanceRuleResultUpdateParams{
			AssigneeID: openlayer.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			BlockedBy: openlayer.F([]openlayer.GovernanceRuleResultUpdateParamsBlockedBy{{
				ID:     openlayer.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Status: openlayer.F(openlayer.GovernanceRuleResultUpdateParamsBlockedByStatusPassing),
			}}),
			Blocking: openlayer.F([]openlayer.GovernanceRuleResultUpdateParamsBlocking{{
				ID:     openlayer.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Status: openlayer.F(openlayer.GovernanceRuleResultUpdateParamsBlockingStatusPassing),
			}}),
			Deactivated:       openlayer.F(true),
			DeactivatedReason: openlayer.F("deactivatedReason"),
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

func TestGovernanceRuleResultListWithOptionalParams(t *testing.T) {
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
	_, err := client.Governance.RuleResults.List(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.GovernanceRuleResultListParams{
			EnabledFrameworkOnly: openlayer.F(true),
			FrameworkID:          openlayer.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			IncludeUnframed:      openlayer.F(true),
			Page:                 openlayer.F(int64(1)),
			PerPage:              openlayer.F(int64(1)),
			ProjectID:            openlayer.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			RuleID:               openlayer.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Scope:                openlayer.F(openlayer.GovernanceRuleResultListParamsScopeProject),
			SearchQuery:          openlayer.F("searchQuery"),
			Status:               openlayer.F(openlayer.GovernanceRuleResultListParamsStatusPassing),
			Type:                 openlayer.F(openlayer.GovernanceRuleResultListParamsTypePlatform),
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

func TestGovernanceRuleResultNewEvidenceWithOptionalParams(t *testing.T) {
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
	_, err := client.Governance.RuleResults.NewEvidence(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.GovernanceRuleResultNewEvidenceParams{
			Description: openlayer.F("description"),
			Name:        openlayer.F("Model risk assessment 2026"),
			StorageUri:  openlayer.F("s3://openlayer-evidence/evidence.pdf"),
			Text:        openlayer.F("text"),
			URL:         openlayer.F("https://openlayer.com/evidence"),
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

func TestGovernanceRuleResultListEvidenceWithOptionalParams(t *testing.T) {
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
	_, err := client.Governance.RuleResults.ListEvidence(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.GovernanceRuleResultListEvidenceParams{
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
