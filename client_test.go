// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stainless-sdks/openlayer-go"
	"github.com/stainless-sdks/openlayer-go/internal"
	"github.com/stainless-sdks/openlayer-go/option"
)

type closureTransport struct {
	fn func(req *http.Request) (*http.Response, error)
}

func (t *closureTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.fn(req)
}

func TestUserAgentHeader(t *testing.T) {
	var userAgent string
	client := openlayer.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					userAgent = req.Header.Get("User-Agent")
					return &http.Response{
						StatusCode: http.StatusOK,
					}, nil
				},
			},
		}),
	)
	client.InferencePipelines.Data.Stream(
		context.Background(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.InferencePipelineDataStreamParams{
			Config: openlayer.F[openlayer.InferencePipelineDataStreamParamsConfigUnion](openlayer.InferencePipelineDataStreamParamsConfigLlmDataConfig{
				InputVariableNames:   openlayer.F([]string{"user_query"}),
				OutputColumnName:     openlayer.F("output"),
				NumOfTokenColumnName: "tokens",
				CostColumnName:       openlayer.F("cost"),
				TimestampColumnName:  openlayer.F("timestamp"),
			}),
			Rows: openlayer.F([]map[string]interface{}{{
				"user_query": "what's the meaning of life?",
				"output":     "42",
				"tokens":     map[string]interface{}{},
				"cost":       map[string]interface{}{},
				"timestamp":  map[string]interface{}{},
			}}),
		},
	)
	if userAgent != fmt.Sprintf("Openlayer/Go %s", internal.PackageVersion) {
		t.Errorf("Expected User-Agent to be correct, but got: %#v", userAgent)
	}
}

func TestRetryAfter(t *testing.T) {
	attempts := 0
	client := openlayer.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					attempts++
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After"): []string{"0.1"},
						},
					}, nil
				},
			},
		}),
	)
	res, err := client.InferencePipelines.Data.Stream(
		context.Background(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.InferencePipelineDataStreamParams{
			Config: openlayer.F[openlayer.InferencePipelineDataStreamParamsConfigUnion](openlayer.InferencePipelineDataStreamParamsConfigLlmDataConfig{
				InputVariableNames:   openlayer.F([]string{"user_query"}),
				OutputColumnName:     openlayer.F("output"),
				NumOfTokenColumnName: "tokens",
				CostColumnName:       openlayer.F("cost"),
				TimestampColumnName:  openlayer.F("timestamp"),
			}),
			Rows: openlayer.F([]map[string]interface{}{{
				"user_query": "what's the meaning of life?",
				"output":     "42",
				"tokens":     map[string]interface{}{},
				"cost":       map[string]interface{}{},
				"timestamp":  map[string]interface{}{},
			}}),
		},
	)
	if err == nil || res != nil {
		t.Error("Expected there to be a cancel error and for the response to be nil")
	}
	if want := 3; attempts != want {
		t.Errorf("Expected %d attempts, got %d", want, attempts)
	}
}

func TestRetryAfterMs(t *testing.T) {
	attempts := 0
	client := openlayer.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					attempts++
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After-Ms"): []string{"100"},
						},
					}, nil
				},
			},
		}),
	)
	res, err := client.InferencePipelines.Data.Stream(
		context.Background(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.InferencePipelineDataStreamParams{
			Config: openlayer.F[openlayer.InferencePipelineDataStreamParamsConfigUnion](openlayer.InferencePipelineDataStreamParamsConfigLlmDataConfig{
				InputVariableNames:   openlayer.F([]string{"user_query"}),
				OutputColumnName:     openlayer.F("output"),
				NumOfTokenColumnName: "tokens",
				CostColumnName:       openlayer.F("cost"),
				TimestampColumnName:  openlayer.F("timestamp"),
			}),
			Rows: openlayer.F([]map[string]interface{}{{
				"user_query": "what's the meaning of life?",
				"output":     "42",
				"tokens":     map[string]interface{}{},
				"cost":       map[string]interface{}{},
				"timestamp":  map[string]interface{}{},
			}}),
		},
	)
	if err == nil || res != nil {
		t.Error("Expected there to be a cancel error and for the response to be nil")
	}
	if want := 3; attempts != want {
		t.Errorf("Expected %d attempts, got %d", want, attempts)
	}
}

func TestContextCancel(t *testing.T) {
	client := openlayer.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					<-req.Context().Done()
					return nil, req.Context().Err()
				},
			},
		}),
	)
	cancelCtx, cancel := context.WithCancel(context.Background())
	cancel()
	res, err := client.InferencePipelines.Data.Stream(
		cancelCtx,
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.InferencePipelineDataStreamParams{
			Config: openlayer.F[openlayer.InferencePipelineDataStreamParamsConfigUnion](openlayer.InferencePipelineDataStreamParamsConfigLlmDataConfig{
				InputVariableNames:   openlayer.F([]string{"user_query"}),
				OutputColumnName:     openlayer.F("output"),
				NumOfTokenColumnName: "tokens",
				CostColumnName:       openlayer.F("cost"),
				TimestampColumnName:  openlayer.F("timestamp"),
			}),
			Rows: openlayer.F([]map[string]interface{}{{
				"user_query": "what's the meaning of life?",
				"output":     "42",
				"tokens":     map[string]interface{}{},
				"cost":       map[string]interface{}{},
				"timestamp":  map[string]interface{}{},
			}}),
		},
	)
	if err == nil || res != nil {
		t.Error("Expected there to be a cancel error and for the response to be nil")
	}
}

func TestContextCancelDelay(t *testing.T) {
	client := openlayer.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					<-req.Context().Done()
					return nil, req.Context().Err()
				},
			},
		}),
	)
	cancelCtx, cancel := context.WithTimeout(context.Background(), 2*time.Millisecond)
	defer cancel()
	res, err := client.InferencePipelines.Data.Stream(
		cancelCtx,
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		openlayer.InferencePipelineDataStreamParams{
			Config: openlayer.F[openlayer.InferencePipelineDataStreamParamsConfigUnion](openlayer.InferencePipelineDataStreamParamsConfigLlmDataConfig{
				InputVariableNames:   openlayer.F([]string{"user_query"}),
				OutputColumnName:     openlayer.F("output"),
				NumOfTokenColumnName: "tokens",
				CostColumnName:       openlayer.F("cost"),
				TimestampColumnName:  openlayer.F("timestamp"),
			}),
			Rows: openlayer.F([]map[string]interface{}{{
				"user_query": "what's the meaning of life?",
				"output":     "42",
				"tokens":     map[string]interface{}{},
				"cost":       map[string]interface{}{},
				"timestamp":  map[string]interface{}{},
			}}),
		},
	)
	if err == nil || res != nil {
		t.Error("expected there to be a cancel error and for the response to be nil")
	}
}

func TestContextDeadline(t *testing.T) {
	testTimeout := time.After(3 * time.Second)
	testDone := make(chan struct{})

	deadline := time.Now().Add(100 * time.Millisecond)
	deadlineCtx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go func() {
		client := openlayer.NewClient(
			option.WithHTTPClient(&http.Client{
				Transport: &closureTransport{
					fn: func(req *http.Request) (*http.Response, error) {
						<-req.Context().Done()
						return nil, req.Context().Err()
					},
				},
			}),
		)
		res, err := client.InferencePipelines.Data.Stream(
			deadlineCtx,
			"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			openlayer.InferencePipelineDataStreamParams{
				Config: openlayer.F[openlayer.InferencePipelineDataStreamParamsConfigUnion](openlayer.InferencePipelineDataStreamParamsConfigLlmDataConfig{
					InputVariableNames:   openlayer.F([]string{"user_query"}),
					OutputColumnName:     openlayer.F("output"),
					NumOfTokenColumnName: "tokens",
					CostColumnName:       openlayer.F("cost"),
					TimestampColumnName:  openlayer.F("timestamp"),
				}),
				Rows: openlayer.F([]map[string]interface{}{{
					"user_query": "what's the meaning of life?",
					"output":     "42",
					"tokens":     map[string]interface{}{},
					"cost":       map[string]interface{}{},
					"timestamp":  map[string]interface{}{},
				}}),
			},
		)
		if err == nil || res != nil {
			t.Error("expected there to be a deadline error and for the response to be nil")
		}
		close(testDone)
	}()

	select {
	case <-testTimeout:
		t.Fatal("client didn't finish in time")
	case <-testDone:
		if diff := time.Since(deadline); diff < -30*time.Millisecond || 30*time.Millisecond < diff {
			t.Fatalf("client did not return within 30ms of context deadline, got %s", diff)
		}
	}
}
