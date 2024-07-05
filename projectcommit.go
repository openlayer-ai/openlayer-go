// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer

import (
	"github.com/openlayer-ai/openlayer-go/option"
)

// ProjectCommitService contains methods and other services that help with
// interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectCommitService] method instead.
type ProjectCommitService struct {
	Options []option.RequestOption
}

// NewProjectCommitService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProjectCommitService(opts ...option.RequestOption) (r *ProjectCommitService) {
	r = &ProjectCommitService{}
	r.Options = opts
	return
}
