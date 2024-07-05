// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer

import (
	"github.com/openlayer-ai/openlayer-go/option"
)

// InferencePipelineDataService contains methods and other services that help with
// interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInferencePipelineDataService] method instead.
type InferencePipelineDataService struct {
	Options []option.RequestOption
}

// NewInferencePipelineDataService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInferencePipelineDataService(opts ...option.RequestOption) (r *InferencePipelineDataService) {
	r = &InferencePipelineDataService{}
	r.Options = opts
	return
}
