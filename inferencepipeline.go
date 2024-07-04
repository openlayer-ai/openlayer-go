// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer

import (
	"github.com/stainless-sdks/openlayer-go/option"
)

// InferencePipelineService contains methods and other services that help with
// interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInferencePipelineService] method instead.
type InferencePipelineService struct {
	Options     []option.RequestOption
	Data        *InferencePipelineDataService
	TestResults *InferencePipelineTestResultService
}

// NewInferencePipelineService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInferencePipelineService(opts ...option.RequestOption) (r *InferencePipelineService) {
	r = &InferencePipelineService{}
	r.Options = opts
	r.Data = NewInferencePipelineDataService(opts...)
	r.TestResults = NewInferencePipelineTestResultService(opts...)
	return
}
