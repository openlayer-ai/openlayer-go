// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer

import (
	"github.com/openlayer-ai/openlayer-go/option"
)

// GovernanceService contains methods and other services that help with interacting
// with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGovernanceService] method instead.
type GovernanceService struct {
	Options     []option.RequestOption
	Frameworks  *GovernanceFrameworkService
	Rules       *GovernanceRuleService
	RuleResults *GovernanceRuleResultService
	RuleStats   *GovernanceRuleStatService
	RuleTags    *GovernanceRuleTagService
}

// NewGovernanceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGovernanceService(opts ...option.RequestOption) (r *GovernanceService) {
	r = &GovernanceService{}
	r.Options = opts
	r.Frameworks = NewGovernanceFrameworkService(opts...)
	r.Rules = NewGovernanceRuleService(opts...)
	r.RuleResults = NewGovernanceRuleResultService(opts...)
	r.RuleStats = NewGovernanceRuleStatService(opts...)
	r.RuleTags = NewGovernanceRuleTagService(opts...)
	return
}
