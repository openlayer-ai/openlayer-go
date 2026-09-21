// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/openlayer-ai/openlayer-go/internal/apijson"
	"github.com/openlayer-ai/openlayer-go/internal/apiquery"
	"github.com/openlayer-ai/openlayer-go/internal/param"
	"github.com/openlayer-ai/openlayer-go/internal/requestconfig"
	"github.com/openlayer-ai/openlayer-go/option"
)

// GovernanceRuleStatService contains methods and other services that help with
// interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGovernanceRuleStatService] method instead.
type GovernanceRuleStatService struct {
	Options []option.RequestOption
}

// NewGovernanceRuleStatService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGovernanceRuleStatService(opts ...option.RequestOption) (r *GovernanceRuleStatService) {
	r = &GovernanceRuleStatService{}
	r.Options = opts
	return
}

// Get a compliance roll-up for a workspace: how many rules exist, and how many of
// their results are passing, failing, pending, or due for renewal.
//
// Counts respect the filters you pass, so `frameworkId` gives you a single
// framework's overall compliance and `projectId` gives you a single project's.
func (r *GovernanceRuleStatService) Get(ctx context.Context, workspaceID string, query GovernanceRuleStatGetParams, opts ...option.RequestOption) (res *GovernanceRuleStatGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("workspaces/%s/rule-stats", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type GovernanceRuleStatGetResponse struct {
	// Counts of rule results, after any filters in the request, with breakdowns by the
	// type and scope of the rule each result belongs to.
	RuleResults GovernanceRuleStatGetResponseRuleResults `json:"ruleResults" api:"required"`
	// Counts of the rules themselves, after any filters in the request.
	Rules GovernanceRuleStatGetResponseRules `json:"rules" api:"required"`
	JSON  governanceRuleStatGetResponseJSON  `json:"-"`
}

// governanceRuleStatGetResponseJSON contains the JSON metadata for the struct
// [GovernanceRuleStatGetResponse]
type governanceRuleStatGetResponseJSON struct {
	RuleResults apijson.Field
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleStatGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleStatGetResponseJSON) RawJSON() string {
	return r.raw
}

// Counts of rule results, after any filters in the request, with breakdowns by the
// type and scope of the rule each result belongs to.
type GovernanceRuleStatGetResponseRuleResults struct {
	// The total number of rule results.
	Total int64 `json:"total" api:"required"`
	// The number of rule results whose evidence is about to expire.
	TotalDueSoon int64 `json:"totalDueSoon" api:"required"`
	// The number of rule results that errored during evaluation.
	TotalError int64 `json:"totalError" api:"required"`
	// The number of failing rule results.
	TotalFailing int64 `json:"totalFailing" api:"required"`
	// The number of passing rule results.
	TotalPassing int64 `json:"totalPassing" api:"required"`
	// The number of rule results that have not been satisfied yet.
	TotalPending int64 `json:"totalPending" api:"required"`
	// The number of rule results currently being evaluated.
	TotalRunning int64 `json:"totalRunning" api:"required"`
	// The number of skipped rule results.
	TotalSkipped int64                                               `json:"totalSkipped" api:"required"`
	ByRuleScope  GovernanceRuleStatGetResponseRuleResultsByRuleScope `json:"byRuleScope"`
	ByRuleType   GovernanceRuleStatGetResponseRuleResultsByRuleType  `json:"byRuleType"`
	JSON         governanceRuleStatGetResponseRuleResultsJSON        `json:"-"`
}

// governanceRuleStatGetResponseRuleResultsJSON contains the JSON metadata for the
// struct [GovernanceRuleStatGetResponseRuleResults]
type governanceRuleStatGetResponseRuleResultsJSON struct {
	Total        apijson.Field
	TotalDueSoon apijson.Field
	TotalError   apijson.Field
	TotalFailing apijson.Field
	TotalPassing apijson.Field
	TotalPending apijson.Field
	TotalRunning apijson.Field
	TotalSkipped apijson.Field
	ByRuleScope  apijson.Field
	ByRuleType   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GovernanceRuleStatGetResponseRuleResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleStatGetResponseRuleResultsJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleStatGetResponseRuleResultsByRuleScope struct {
	Project   GovernanceRuleStatGetResponseRuleResultsByRuleScopeProject   `json:"project"`
	Workspace GovernanceRuleStatGetResponseRuleResultsByRuleScopeWorkspace `json:"workspace"`
	JSON      governanceRuleStatGetResponseRuleResultsByRuleScopeJSON      `json:"-"`
}

// governanceRuleStatGetResponseRuleResultsByRuleScopeJSON contains the JSON
// metadata for the struct [GovernanceRuleStatGetResponseRuleResultsByRuleScope]
type governanceRuleStatGetResponseRuleResultsByRuleScopeJSON struct {
	Project     apijson.Field
	Workspace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleStatGetResponseRuleResultsByRuleScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleStatGetResponseRuleResultsByRuleScopeJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleStatGetResponseRuleResultsByRuleScopeProject struct {
	// The total number of rule results.
	Total int64 `json:"total" api:"required"`
	// The number of rule results whose evidence is about to expire.
	TotalDueSoon int64 `json:"totalDueSoon" api:"required"`
	// The number of rule results that errored during evaluation.
	TotalError int64 `json:"totalError" api:"required"`
	// The number of failing rule results.
	TotalFailing int64 `json:"totalFailing" api:"required"`
	// The number of passing rule results.
	TotalPassing int64 `json:"totalPassing" api:"required"`
	// The number of rule results that have not been satisfied yet.
	TotalPending int64 `json:"totalPending" api:"required"`
	// The number of rule results currently being evaluated.
	TotalRunning int64 `json:"totalRunning" api:"required"`
	// The number of skipped rule results.
	TotalSkipped int64                                                          `json:"totalSkipped" api:"required"`
	JSON         governanceRuleStatGetResponseRuleResultsByRuleScopeProjectJSON `json:"-"`
}

// governanceRuleStatGetResponseRuleResultsByRuleScopeProjectJSON contains the JSON
// metadata for the struct
// [GovernanceRuleStatGetResponseRuleResultsByRuleScopeProject]
type governanceRuleStatGetResponseRuleResultsByRuleScopeProjectJSON struct {
	Total        apijson.Field
	TotalDueSoon apijson.Field
	TotalError   apijson.Field
	TotalFailing apijson.Field
	TotalPassing apijson.Field
	TotalPending apijson.Field
	TotalRunning apijson.Field
	TotalSkipped apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GovernanceRuleStatGetResponseRuleResultsByRuleScopeProject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleStatGetResponseRuleResultsByRuleScopeProjectJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleStatGetResponseRuleResultsByRuleScopeWorkspace struct {
	// The total number of rule results.
	Total int64 `json:"total" api:"required"`
	// The number of rule results whose evidence is about to expire.
	TotalDueSoon int64 `json:"totalDueSoon" api:"required"`
	// The number of rule results that errored during evaluation.
	TotalError int64 `json:"totalError" api:"required"`
	// The number of failing rule results.
	TotalFailing int64 `json:"totalFailing" api:"required"`
	// The number of passing rule results.
	TotalPassing int64 `json:"totalPassing" api:"required"`
	// The number of rule results that have not been satisfied yet.
	TotalPending int64 `json:"totalPending" api:"required"`
	// The number of rule results currently being evaluated.
	TotalRunning int64 `json:"totalRunning" api:"required"`
	// The number of skipped rule results.
	TotalSkipped int64                                                            `json:"totalSkipped" api:"required"`
	JSON         governanceRuleStatGetResponseRuleResultsByRuleScopeWorkspaceJSON `json:"-"`
}

// governanceRuleStatGetResponseRuleResultsByRuleScopeWorkspaceJSON contains the
// JSON metadata for the struct
// [GovernanceRuleStatGetResponseRuleResultsByRuleScopeWorkspace]
type governanceRuleStatGetResponseRuleResultsByRuleScopeWorkspaceJSON struct {
	Total        apijson.Field
	TotalDueSoon apijson.Field
	TotalError   apijson.Field
	TotalFailing apijson.Field
	TotalPassing apijson.Field
	TotalPending apijson.Field
	TotalRunning apijson.Field
	TotalSkipped apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GovernanceRuleStatGetResponseRuleResultsByRuleScopeWorkspace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleStatGetResponseRuleResultsByRuleScopeWorkspaceJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleStatGetResponseRuleResultsByRuleType struct {
	Evidence GovernanceRuleStatGetResponseRuleResultsByRuleTypeEvidence `json:"evidence"`
	Platform GovernanceRuleStatGetResponseRuleResultsByRuleTypePlatform `json:"platform"`
	JSON     governanceRuleStatGetResponseRuleResultsByRuleTypeJSON     `json:"-"`
}

// governanceRuleStatGetResponseRuleResultsByRuleTypeJSON contains the JSON
// metadata for the struct [GovernanceRuleStatGetResponseRuleResultsByRuleType]
type governanceRuleStatGetResponseRuleResultsByRuleTypeJSON struct {
	Evidence    apijson.Field
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleStatGetResponseRuleResultsByRuleType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleStatGetResponseRuleResultsByRuleTypeJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleStatGetResponseRuleResultsByRuleTypeEvidence struct {
	// The total number of rule results.
	Total int64 `json:"total" api:"required"`
	// The number of rule results whose evidence is about to expire.
	TotalDueSoon int64 `json:"totalDueSoon" api:"required"`
	// The number of rule results that errored during evaluation.
	TotalError int64 `json:"totalError" api:"required"`
	// The number of failing rule results.
	TotalFailing int64 `json:"totalFailing" api:"required"`
	// The number of passing rule results.
	TotalPassing int64 `json:"totalPassing" api:"required"`
	// The number of rule results that have not been satisfied yet.
	TotalPending int64 `json:"totalPending" api:"required"`
	// The number of rule results currently being evaluated.
	TotalRunning int64 `json:"totalRunning" api:"required"`
	// The number of skipped rule results.
	TotalSkipped int64                                                          `json:"totalSkipped" api:"required"`
	JSON         governanceRuleStatGetResponseRuleResultsByRuleTypeEvidenceJSON `json:"-"`
}

// governanceRuleStatGetResponseRuleResultsByRuleTypeEvidenceJSON contains the JSON
// metadata for the struct
// [GovernanceRuleStatGetResponseRuleResultsByRuleTypeEvidence]
type governanceRuleStatGetResponseRuleResultsByRuleTypeEvidenceJSON struct {
	Total        apijson.Field
	TotalDueSoon apijson.Field
	TotalError   apijson.Field
	TotalFailing apijson.Field
	TotalPassing apijson.Field
	TotalPending apijson.Field
	TotalRunning apijson.Field
	TotalSkipped apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GovernanceRuleStatGetResponseRuleResultsByRuleTypeEvidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleStatGetResponseRuleResultsByRuleTypeEvidenceJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleStatGetResponseRuleResultsByRuleTypePlatform struct {
	// The total number of rule results.
	Total int64 `json:"total" api:"required"`
	// The number of rule results whose evidence is about to expire.
	TotalDueSoon int64 `json:"totalDueSoon" api:"required"`
	// The number of rule results that errored during evaluation.
	TotalError int64 `json:"totalError" api:"required"`
	// The number of failing rule results.
	TotalFailing int64 `json:"totalFailing" api:"required"`
	// The number of passing rule results.
	TotalPassing int64 `json:"totalPassing" api:"required"`
	// The number of rule results that have not been satisfied yet.
	TotalPending int64 `json:"totalPending" api:"required"`
	// The number of rule results currently being evaluated.
	TotalRunning int64 `json:"totalRunning" api:"required"`
	// The number of skipped rule results.
	TotalSkipped int64                                                          `json:"totalSkipped" api:"required"`
	JSON         governanceRuleStatGetResponseRuleResultsByRuleTypePlatformJSON `json:"-"`
}

// governanceRuleStatGetResponseRuleResultsByRuleTypePlatformJSON contains the JSON
// metadata for the struct
// [GovernanceRuleStatGetResponseRuleResultsByRuleTypePlatform]
type governanceRuleStatGetResponseRuleResultsByRuleTypePlatformJSON struct {
	Total        apijson.Field
	TotalDueSoon apijson.Field
	TotalError   apijson.Field
	TotalFailing apijson.Field
	TotalPassing apijson.Field
	TotalPending apijson.Field
	TotalRunning apijson.Field
	TotalSkipped apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GovernanceRuleStatGetResponseRuleResultsByRuleTypePlatform) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleStatGetResponseRuleResultsByRuleTypePlatformJSON) RawJSON() string {
	return r.raw
}

// Counts of the rules themselves, after any filters in the request.
type GovernanceRuleStatGetResponseRules struct {
	// Rule counts by scope.
	ByScope GovernanceRuleStatGetResponseRulesByScope `json:"byScope"`
	// Rule counts by type.
	ByType GovernanceRuleStatGetResponseRulesByType `json:"byType"`
	// The total number of rules.
	Total int64                                  `json:"total"`
	JSON  governanceRuleStatGetResponseRulesJSON `json:"-"`
}

// governanceRuleStatGetResponseRulesJSON contains the JSON metadata for the struct
// [GovernanceRuleStatGetResponseRules]
type governanceRuleStatGetResponseRulesJSON struct {
	ByScope     apijson.Field
	ByType      apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleStatGetResponseRules) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleStatGetResponseRulesJSON) RawJSON() string {
	return r.raw
}

// Rule counts by scope.
type GovernanceRuleStatGetResponseRulesByScope struct {
	Project   int64                                         `json:"project"`
	Workspace int64                                         `json:"workspace"`
	JSON      governanceRuleStatGetResponseRulesByScopeJSON `json:"-"`
}

// governanceRuleStatGetResponseRulesByScopeJSON contains the JSON metadata for the
// struct [GovernanceRuleStatGetResponseRulesByScope]
type governanceRuleStatGetResponseRulesByScopeJSON struct {
	Project     apijson.Field
	Workspace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleStatGetResponseRulesByScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleStatGetResponseRulesByScopeJSON) RawJSON() string {
	return r.raw
}

// Rule counts by type.
type GovernanceRuleStatGetResponseRulesByType struct {
	Evidence int64                                        `json:"evidence"`
	Platform int64                                        `json:"platform"`
	JSON     governanceRuleStatGetResponseRulesByTypeJSON `json:"-"`
}

// governanceRuleStatGetResponseRulesByTypeJSON contains the JSON metadata for the
// struct [GovernanceRuleStatGetResponseRulesByType]
type governanceRuleStatGetResponseRulesByTypeJSON struct {
	Evidence    apijson.Field
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleStatGetResponseRulesByType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleStatGetResponseRulesByTypeJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleStatGetParams struct {
	// Only include items belonging to this framework.
	FrameworkID param.Field[string] `query:"frameworkId" format:"uuid"`
	// Only include items that apply to this project.
	ProjectID param.Field[string] `query:"projectId" format:"uuid"`
}

// URLQuery serializes [GovernanceRuleStatGetParams]'s query parameters as
// `url.Values`.
func (r GovernanceRuleStatGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
