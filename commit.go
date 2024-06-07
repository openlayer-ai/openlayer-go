// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomopenlayeraiopenlayergo

import (
	"github.com/openlayer-ai/openlayer-go/option"
)

// CommitService contains methods and other services that help with interacting
// with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCommitService] method instead.
type CommitService struct {
	Options     []option.RequestOption
	TestResults *CommitTestResultService
}

// NewCommitService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCommitService(opts ...option.RequestOption) (r *CommitService) {
	r = &CommitService{}
	r.Options = opts
	r.TestResults = NewCommitTestResultService(opts...)
	return
}
