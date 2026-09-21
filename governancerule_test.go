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

func TestGovernanceRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Governance.Rules.New(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.GovernanceRuleNewParams{
			Name:       openlayer.F("Monitoring enabled"),
			Scope:      openlayer.F(openlayer.GovernanceRuleNewParamsScopeProject),
			Type:       openlayer.F(openlayer.GovernanceRuleNewParamsTypePlatform),
			AssigneeID: openlayer.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AutomationParams: openlayer.F(map[string]interface{}{
				"foo": "bar",
			}),
			AutomationType:     openlayer.F("monitoring_mode_enabled"),
			Deactivated:        openlayer.F(true),
			Description:        openlayer.F("Each project must have Openlayer monitoring mode enabled."),
			EvidenceType:       openlayer.F(openlayer.GovernanceRuleNewParamsEvidenceTypeDocument),
			RenewalCadenceDays: openlayer.F(int64(90)),
			TagIDs:             openlayer.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
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

func TestGovernanceRuleGet(t *testing.T) {
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
	_, err := client.Governance.Rules.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *openlayer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGovernanceRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Governance.Rules.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.GovernanceRuleUpdateParams{
			AssigneeID:         openlayer.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Deactivated:        openlayer.F(true),
			Description:        openlayer.F("Each project must have Openlayer monitoring mode enabled."),
			Name:               openlayer.F("Monitoring enabled"),
			RenewalCadenceDays: openlayer.F(int64(90)),
			TagIDs:             openlayer.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
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

func TestGovernanceRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Governance.Rules.List(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.GovernanceRuleListParams{
			Asc:                  openlayer.F(true),
			AssigneeID:           openlayer.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Deactivated:          openlayer.F(true),
			EnabledFrameworkOnly: openlayer.F(true),
			FrameworkID:          openlayer.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Group:                openlayer.F(openlayer.GovernanceRuleListParamsGroupOpen),
			IncludeResults:       openlayer.F(true),
			IncludeUnframed:      openlayer.F(true),
			Page:                 openlayer.F(int64(1)),
			PerPage:              openlayer.F(int64(1)),
			ProjectID:            openlayer.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Scope:                openlayer.F(openlayer.GovernanceRuleListParamsScopeProject),
			SearchQuery:          openlayer.F("searchQuery"),
			SortBy:               openlayer.F(openlayer.GovernanceRuleListParamsSortByName),
			Status:               openlayer.F(openlayer.GovernanceRuleListParamsStatusPassing),
			Tags:                 openlayer.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			Type:                 openlayer.F(openlayer.GovernanceRuleListParamsTypePlatform),
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

func TestGovernanceRuleDelete(t *testing.T) {
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
	err := client.Governance.Rules.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *openlayer.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
