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

func TestProjectTestNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.Tests.New(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.ProjectTestNewParams{
			Description: openlayer.F[any]("This test checks for duplicate rows in the dataset."),
			Name:        openlayer.F("No duplicate rows"),
			Subtype:     openlayer.F(openlayer.ProjectTestNewParamsSubtypeDuplicateRowCount),
			Thresholds: openlayer.F([]openlayer.ProjectTestNewParamsThreshold{{
				InsightName: openlayer.F("duplicateRowCount"),
				InsightParameters: openlayer.F([]openlayer.ProjectTestNewParamsThresholdsInsightParameter{{
					Name:  openlayer.F("column_name"),
					Value: openlayer.F[any]("Age"),
				}}),
				Measurement:   openlayer.F("duplicateRowCount"),
				Operator:      openlayer.F(openlayer.ProjectTestNewParamsThresholdsOperatorLessOrEquals),
				ThresholdMode: openlayer.F(openlayer.ProjectTestNewParamsThresholdsThresholdModeAutomatic),
				Value:         openlayer.F[openlayer.ProjectTestNewParamsThresholdsValueUnion](shared.UnionFloat(0.000000)),
			}}),
			Type:                  openlayer.F(openlayer.ProjectTestNewParamsTypeIntegrity),
			Archived:              openlayer.F(false),
			DelayWindow:           openlayer.F(0.000000),
			EvaluationWindow:      openlayer.F(3600.000000),
			UsesMlModel:           openlayer.F(false),
			UsesProductionData:    openlayer.F(false),
			UsesReferenceDataset:  openlayer.F(false),
			UsesTrainingDataset:   openlayer.F(false),
			UsesValidationDataset: openlayer.F(true),
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

func TestProjectTestListWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.Tests.List(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.ProjectTestListParams{
			IncludeArchived:    openlayer.F(true),
			OriginVersionID:    openlayer.F("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
			Page:               openlayer.F(int64(1)),
			PerPage:            openlayer.F(int64(1)),
			Suggested:          openlayer.F(true),
			Type:               openlayer.F(openlayer.ProjectTestListParamsTypeIntegrity),
			UsesProductionData: openlayer.F(true),
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
