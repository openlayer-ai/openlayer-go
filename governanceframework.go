// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/openlayer-ai/openlayer-go/internal/apijson"
	"github.com/openlayer-ai/openlayer-go/internal/apiquery"
	"github.com/openlayer-ai/openlayer-go/internal/param"
	"github.com/openlayer-ai/openlayer-go/internal/requestconfig"
	"github.com/openlayer-ai/openlayer-go/option"
)

// GovernanceFrameworkService contains methods and other services that help with
// interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGovernanceFrameworkService] method instead.
type GovernanceFrameworkService struct {
	Options     []option.RequestOption
	Documents   *GovernanceFrameworkDocumentService
	Sections    *GovernanceFrameworkSectionService
	Subsections *GovernanceFrameworkSubsectionService
}

// NewGovernanceFrameworkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGovernanceFrameworkService(opts ...option.RequestOption) (r *GovernanceFrameworkService) {
	r = &GovernanceFrameworkService{}
	r.Options = opts
	r.Documents = NewGovernanceFrameworkDocumentService(opts...)
	r.Sections = NewGovernanceFrameworkSectionService(opts...)
	r.Subsections = NewGovernanceFrameworkSubsectionService(opts...)
	return
}

// Create a custom governance framework in a workspace.
//
// Use this to track compliance against an internal policy, or against a standard
// Openlayer does not ship as a built-in framework. A new framework starts with no
// rules -- add them from the Openlayer app, or map an existing rule to it.
//
// A framework is created disabled unless you pass `enabled: true`. While it is
// disabled its rules are not evaluated and do not count towards compliance.
func (r *GovernanceFrameworkService) New(ctx context.Context, workspaceID string, body GovernanceFrameworkNewParams, opts ...option.RequestOption) (res *GovernanceFrameworkNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("workspaces/%s/frameworks", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a governance framework by its id.
func (r *GovernanceFrameworkService) Get(ctx context.Context, frameworkID string, opts ...option.RequestOption) (res *GovernanceFrameworkGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if frameworkID == "" {
		err = errors.New("missing required frameworkId parameter")
		return nil, err
	}
	path := fmt.Sprintf("frameworks/%s", frameworkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a governance framework.
//
// The most common use is activating or deactivating a framework for the workspace
// by setting `enabled`. Rules of a disabled framework are not evaluated and do not
// count towards compliance.
//
// Frameworks that ship with Openlayer report `immutable: true`. For those, only
// `enabled`, `tags`, and `projectSelector` can be changed -- their name and
// definition are managed by Openlayer.
//
// Only the fields you send are changed.
func (r *GovernanceFrameworkService) Update(ctx context.Context, frameworkID string, body GovernanceFrameworkUpdateParams, opts ...option.RequestOption) (res *GovernanceFrameworkUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if frameworkID == "" {
		err = errors.New("missing required frameworkId parameter")
		return nil, err
	}
	path := fmt.Sprintf("frameworks/%s", frameworkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// List the governance frameworks in a workspace.
//
// A framework is a set of rules -- drawn from a regulation, a standard, or your
// own internal policy -- that Openlayer tracks compliance against. Use this
// endpoint to find the framework you want to report on, then read its rules and
// rule results.
func (r *GovernanceFrameworkService) List(ctx context.Context, workspaceID string, query GovernanceFrameworkListParams, opts ...option.RequestOption) (res *GovernanceFrameworkListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("workspaces/%s/frameworks", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Export a framework's evidence and progress as an audit-ready zip archive.
//
// The archive holds every evidence file uploaded against the framework's
// evidence-based rules, a markdown report of the framework's progress and the
// status of all its rules (broken down by documentation section when the framework
// has documents), and CSV manifests of rules and evidence with SHA-256 checksums.
//
// Send `projectId` to export one project's compliance with the framework. Omit it
// for the workspace-wide view across every project in the framework, including
// workspace-scoped rules.
//
// The export runs as a background task, so this returns `202` immediately. To
// collect the archive:
//
//  1. Poll `GET /background-tasks/{taskId}` with the returned `taskResultId` until
//     `complete` is `true`.
//  2. Read `outputs.storageUri` off that task.
//  3. Exchange it for a download link at
//     `GET /storage/presigned-url?storageUri=<uri>`.
//
// Rate limited to 2 requests per minute per framework. Asking for an export while
// an identical one is still queued returns that task rather than starting a second
// one.
func (r *GovernanceFrameworkService) Export(ctx context.Context, frameworkID string, body GovernanceFrameworkExportParams, opts ...option.RequestOption) (res *GovernanceFrameworkExportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if frameworkID == "" {
		err = errors.New("missing required frameworkId parameter")
		return nil, err
	}
	path := fmt.Sprintf("frameworks/%s/export", frameworkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a compliance roll-up for a framework, one row per project it applies to.
//
// Each row counts the project's rule results by status, so you can report on where
// a framework is complete and where it is not without fetching every individual
// rule result.
func (r *GovernanceFrameworkService) ListProjectRuleStats(ctx context.Context, frameworkID string, query GovernanceFrameworkListProjectRuleStatsParams, opts ...option.RequestOption) (res *GovernanceFrameworkListProjectRuleStatsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if frameworkID == "" {
		err = errors.New("missing required frameworkId parameter")
		return nil, err
	}
	path := fmt.Sprintf("frameworks/%s/project-rule-stats", frameworkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List the projects a framework applies to.
//
// Which projects a framework covers is determined by its `projectSelector`. A
// framework with an empty selector applies to every project in the workspace.
func (r *GovernanceFrameworkService) ListProjects(ctx context.Context, frameworkID string, query GovernanceFrameworkListProjectsParams, opts ...option.RequestOption) (res *GovernanceFrameworkListProjectsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if frameworkID == "" {
		err = errors.New("missing required frameworkId parameter")
		return nil, err
	}
	path := fmt.Sprintf("frameworks/%s/projects", frameworkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List the rules that belong to a framework.
//
// To read the compliance status of these rules, use
// [List rule results](/api-reference/rest/governance/list-rule-results) with the
// `frameworkId` filter, or fetch the results of an individual rule.
func (r *GovernanceFrameworkService) ListRules(ctx context.Context, frameworkID string, query GovernanceFrameworkListRulesParams, opts ...option.RequestOption) (res *GovernanceFrameworkListRulesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if frameworkID == "" {
		err = errors.New("missing required frameworkId parameter")
		return nil, err
	}
	path := fmt.Sprintf("frameworks/%s/rules", frameworkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type GovernanceFrameworkNewResponse struct {
	// The framework id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The creation date.
	DateCreated time.Time `json:"dateCreated" api:"required" format:"date-time"`
	// The last update date.
	DateUpdated time.Time `json:"dateUpdated" api:"required" format:"date-time"`
	// Whether the framework is active. Rules of a disabled framework are not evaluated
	// and do not count towards compliance.
	Enabled bool `json:"enabled" api:"required"`
	// The framework name.
	Name string `json:"name" api:"required"`
	// Free-form labels on the framework.
	Tags []string `json:"tags" api:"required"`
	// The id of the workspace the framework belongs to.
	WorkspaceID string `json:"workspaceId" api:"required" format:"uuid"`
	// The icon shown for the framework.
	Avatar GovernanceFrameworkNewResponseAvatar `json:"avatar" api:"nullable"`
	// Identifies a framework that ships with Openlayer, for example `eu_ai_act`,
	// `iso_42001`, `nist_ai_rmf`, or `traiga`. `null` for frameworks you create
	// yourself.
	BuiltInSlug string `json:"builtInSlug" api:"nullable"`
	// The user who created the framework. `null` for built-in frameworks.
	CreatorID string `json:"creatorId" api:"nullable" format:"uuid"`
	// A short description of the framework.
	Description string `json:"description" api:"nullable"`
	// A longer, rich-text description, as a TipTap JSON document.
	ExtendedDescription map[string]interface{} `json:"extendedDescription" api:"nullable"`
	// A link to the external standard or regulation the framework is based on.
	Href string `json:"href" api:"nullable"`
	// Whether the framework definition is managed by Openlayer and cannot be edited.
	Immutable bool `json:"immutable"`
	// Determines which projects the framework applies to. An empty or `null` `match`
	// array applies the framework to every project in the workspace.
	ProjectSelector GovernanceFrameworkNewResponseProjectSelector `json:"projectSelector" api:"nullable"`
	// Compliance roll-up for the framework. Present only on
	// `GET /workspaces/{workspaceId}/frameworks` when the request sets
	// `includeRuleStats=true`.
	RuleStats GovernanceFrameworkNewResponseRuleStats `json:"ruleStats"`
	JSON      governanceFrameworkNewResponseJSON      `json:"-"`
}

// governanceFrameworkNewResponseJSON contains the JSON metadata for the struct
// [GovernanceFrameworkNewResponse]
type governanceFrameworkNewResponseJSON struct {
	ID                  apijson.Field
	DateCreated         apijson.Field
	DateUpdated         apijson.Field
	Enabled             apijson.Field
	Name                apijson.Field
	Tags                apijson.Field
	WorkspaceID         apijson.Field
	Avatar              apijson.Field
	BuiltInSlug         apijson.Field
	CreatorID           apijson.Field
	Description         apijson.Field
	ExtendedDescription apijson.Field
	Href                apijson.Field
	Immutable           apijson.Field
	ProjectSelector     apijson.Field
	RuleStats           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *GovernanceFrameworkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkNewResponseJSON) RawJSON() string {
	return r.raw
}

// The icon shown for the framework.
type GovernanceFrameworkNewResponseAvatar struct {
	Type  GovernanceFrameworkNewResponseAvatarType `json:"type" api:"required"`
	Value string                                   `json:"value" api:"required"`
	JSON  governanceFrameworkNewResponseAvatarJSON `json:"-"`
}

// governanceFrameworkNewResponseAvatarJSON contains the JSON metadata for the
// struct [GovernanceFrameworkNewResponseAvatar]
type governanceFrameworkNewResponseAvatarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkNewResponseAvatar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkNewResponseAvatarJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkNewResponseAvatarType string

const (
	GovernanceFrameworkNewResponseAvatarTypeEmoji        GovernanceFrameworkNewResponseAvatarType = "emoji"
	GovernanceFrameworkNewResponseAvatarTypeImageURL     GovernanceFrameworkNewResponseAvatarType = "imageUrl"
	GovernanceFrameworkNewResponseAvatarTypeBuiltinImage GovernanceFrameworkNewResponseAvatarType = "builtinImage"
)

func (r GovernanceFrameworkNewResponseAvatarType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkNewResponseAvatarTypeEmoji, GovernanceFrameworkNewResponseAvatarTypeImageURL, GovernanceFrameworkNewResponseAvatarTypeBuiltinImage:
		return true
	}
	return false
}

// Determines which projects the framework applies to. An empty or `null` `match`
// array applies the framework to every project in the workspace.
type GovernanceFrameworkNewResponseProjectSelector struct {
	// Match criteria, ANDed together.
	Match []GovernanceFrameworkNewResponseProjectSelectorMatch `json:"match" api:"nullable"`
	JSON  governanceFrameworkNewResponseProjectSelectorJSON    `json:"-"`
}

// governanceFrameworkNewResponseProjectSelectorJSON contains the JSON metadata for
// the struct [GovernanceFrameworkNewResponseProjectSelector]
type governanceFrameworkNewResponseProjectSelectorJSON struct {
	Match       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkNewResponseProjectSelector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkNewResponseProjectSelectorJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkNewResponseProjectSelectorMatch struct {
	// The project property to match against.
	Property GovernanceFrameworkNewResponseProjectSelectorMatchProperty `json:"property" api:"required"`
	// The value to match against. Pass an array to match any of several values, or
	// `null` to match projects where the property is unset. Omit it for `exists` and
	// `notExists`.
	Value interface{} `json:"value" api:"required"`
	// How to compare the project property with `value`. One of `equals`, `notEquals`,
	// `contains`, `notContains`, `startsWith`, `endsWith`, `in`, `notIn`,
	// `greaterThan`, `greaterThanOrEqual`, `lessThan`, `lessThanOrEqual`,
	// `equalsIgnoreCase`, `containsIgnoreCase`, `matches`, `exists`, or `notExists`.
	Operator string                                                 `json:"operator"`
	JSON     governanceFrameworkNewResponseProjectSelectorMatchJSON `json:"-"`
}

// governanceFrameworkNewResponseProjectSelectorMatchJSON contains the JSON
// metadata for the struct [GovernanceFrameworkNewResponseProjectSelectorMatch]
type governanceFrameworkNewResponseProjectSelectorMatchJSON struct {
	Property    apijson.Field
	Value       apijson.Field
	Operator    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkNewResponseProjectSelectorMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkNewResponseProjectSelectorMatchJSON) RawJSON() string {
	return r.raw
}

// The project property to match against.
type GovernanceFrameworkNewResponseProjectSelectorMatchProperty string

const (
	GovernanceFrameworkNewResponseProjectSelectorMatchPropertyTaskType       GovernanceFrameworkNewResponseProjectSelectorMatchProperty = "taskType"
	GovernanceFrameworkNewResponseProjectSelectorMatchPropertyRiskLevel      GovernanceFrameworkNewResponseProjectSelectorMatchProperty = "riskLevel"
	GovernanceFrameworkNewResponseProjectSelectorMatchPropertyRiskTotalScore GovernanceFrameworkNewResponseProjectSelectorMatchProperty = "riskTotalScore"
	GovernanceFrameworkNewResponseProjectSelectorMatchPropertyName           GovernanceFrameworkNewResponseProjectSelectorMatchProperty = "name"
	GovernanceFrameworkNewResponseProjectSelectorMatchPropertyOwnerID        GovernanceFrameworkNewResponseProjectSelectorMatchProperty = "ownerId"
	GovernanceFrameworkNewResponseProjectSelectorMatchPropertyModelTypes     GovernanceFrameworkNewResponseProjectSelectorMatchProperty = "modelTypes"
)

func (r GovernanceFrameworkNewResponseProjectSelectorMatchProperty) IsKnown() bool {
	switch r {
	case GovernanceFrameworkNewResponseProjectSelectorMatchPropertyTaskType, GovernanceFrameworkNewResponseProjectSelectorMatchPropertyRiskLevel, GovernanceFrameworkNewResponseProjectSelectorMatchPropertyRiskTotalScore, GovernanceFrameworkNewResponseProjectSelectorMatchPropertyName, GovernanceFrameworkNewResponseProjectSelectorMatchPropertyOwnerID, GovernanceFrameworkNewResponseProjectSelectorMatchPropertyModelTypes:
		return true
	}
	return false
}

// Compliance roll-up for the framework. Present only on
// `GET /workspaces/{workspaceId}/frameworks` when the request sets
// `includeRuleStats=true`.
type GovernanceFrameworkNewResponseRuleStats struct {
	// How many of the framework's projects fall into each completion band, where a
	// project's completion is the share of its rule results that are passing or
	// skipped. Projects with no evaluated results count as `low`.
	//
	// Zeroed when the request carries `projectId`: the bands compare a framework's
	// projects against each other, which says nothing about a single project.
	ProjectCompletion GovernanceFrameworkNewResponseRuleStatsProjectCompletion `json:"projectCompletion" api:"required"`
	// Rule result counts by status for this framework, matching what
	// `/workspaces/{workspaceId}/rule-stats?frameworkId=<id>` reports. Narrowed to a
	// single project when the request also carries `projectId`.
	RuleResults GovernanceFrameworkNewResponseRuleStatsRuleResults `json:"ruleResults" api:"required"`
	JSON        governanceFrameworkNewResponseRuleStatsJSON        `json:"-"`
}

// governanceFrameworkNewResponseRuleStatsJSON contains the JSON metadata for the
// struct [GovernanceFrameworkNewResponseRuleStats]
type governanceFrameworkNewResponseRuleStatsJSON struct {
	ProjectCompletion apijson.Field
	RuleResults       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *GovernanceFrameworkNewResponseRuleStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkNewResponseRuleStatsJSON) RawJSON() string {
	return r.raw
}

// How many of the framework's projects fall into each completion band, where a
// project's completion is the share of its rule results that are passing or
// skipped. Projects with no evaluated results count as `low`.
//
// Zeroed when the request carries `projectId`: the bands compare a framework's
// projects against each other, which says nothing about a single project.
type GovernanceFrameworkNewResponseRuleStatsProjectCompletion struct {
	// Projects at 80% completion or above.
	High int64 `json:"high" api:"required"`
	// Projects below 20% completion.
	Low int64 `json:"low" api:"required"`
	// Projects at or above 20% but below 80% completion.
	Mid  int64                                                        `json:"mid" api:"required"`
	JSON governanceFrameworkNewResponseRuleStatsProjectCompletionJSON `json:"-"`
}

// governanceFrameworkNewResponseRuleStatsProjectCompletionJSON contains the JSON
// metadata for the struct
// [GovernanceFrameworkNewResponseRuleStatsProjectCompletion]
type governanceFrameworkNewResponseRuleStatsProjectCompletionJSON struct {
	High        apijson.Field
	Low         apijson.Field
	Mid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkNewResponseRuleStatsProjectCompletion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkNewResponseRuleStatsProjectCompletionJSON) RawJSON() string {
	return r.raw
}

// Rule result counts by status for this framework, matching what
// `/workspaces/{workspaceId}/rule-stats?frameworkId=<id>` reports. Narrowed to a
// single project when the request also carries `projectId`.
type GovernanceFrameworkNewResponseRuleStatsRuleResults struct {
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
	TotalSkipped int64                                                  `json:"totalSkipped" api:"required"`
	JSON         governanceFrameworkNewResponseRuleStatsRuleResultsJSON `json:"-"`
}

// governanceFrameworkNewResponseRuleStatsRuleResultsJSON contains the JSON
// metadata for the struct [GovernanceFrameworkNewResponseRuleStatsRuleResults]
type governanceFrameworkNewResponseRuleStatsRuleResultsJSON struct {
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

func (r *GovernanceFrameworkNewResponseRuleStatsRuleResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkNewResponseRuleStatsRuleResultsJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkGetResponse struct {
	// The framework id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The creation date.
	DateCreated time.Time `json:"dateCreated" api:"required" format:"date-time"`
	// The last update date.
	DateUpdated time.Time `json:"dateUpdated" api:"required" format:"date-time"`
	// Whether the framework is active. Rules of a disabled framework are not evaluated
	// and do not count towards compliance.
	Enabled bool `json:"enabled" api:"required"`
	// The framework name.
	Name string `json:"name" api:"required"`
	// Free-form labels on the framework.
	Tags []string `json:"tags" api:"required"`
	// The id of the workspace the framework belongs to.
	WorkspaceID string `json:"workspaceId" api:"required" format:"uuid"`
	// The icon shown for the framework.
	Avatar GovernanceFrameworkGetResponseAvatar `json:"avatar" api:"nullable"`
	// Identifies a framework that ships with Openlayer, for example `eu_ai_act`,
	// `iso_42001`, `nist_ai_rmf`, or `traiga`. `null` for frameworks you create
	// yourself.
	BuiltInSlug string `json:"builtInSlug" api:"nullable"`
	// The user who created the framework. `null` for built-in frameworks.
	CreatorID string `json:"creatorId" api:"nullable" format:"uuid"`
	// A short description of the framework.
	Description string `json:"description" api:"nullable"`
	// A longer, rich-text description, as a TipTap JSON document.
	ExtendedDescription map[string]interface{} `json:"extendedDescription" api:"nullable"`
	// A link to the external standard or regulation the framework is based on.
	Href string `json:"href" api:"nullable"`
	// Whether the framework definition is managed by Openlayer and cannot be edited.
	Immutable bool `json:"immutable"`
	// Determines which projects the framework applies to. An empty or `null` `match`
	// array applies the framework to every project in the workspace.
	ProjectSelector GovernanceFrameworkGetResponseProjectSelector `json:"projectSelector" api:"nullable"`
	// Compliance roll-up for the framework. Present only on
	// `GET /workspaces/{workspaceId}/frameworks` when the request sets
	// `includeRuleStats=true`.
	RuleStats GovernanceFrameworkGetResponseRuleStats `json:"ruleStats"`
	JSON      governanceFrameworkGetResponseJSON      `json:"-"`
}

// governanceFrameworkGetResponseJSON contains the JSON metadata for the struct
// [GovernanceFrameworkGetResponse]
type governanceFrameworkGetResponseJSON struct {
	ID                  apijson.Field
	DateCreated         apijson.Field
	DateUpdated         apijson.Field
	Enabled             apijson.Field
	Name                apijson.Field
	Tags                apijson.Field
	WorkspaceID         apijson.Field
	Avatar              apijson.Field
	BuiltInSlug         apijson.Field
	CreatorID           apijson.Field
	Description         apijson.Field
	ExtendedDescription apijson.Field
	Href                apijson.Field
	Immutable           apijson.Field
	ProjectSelector     apijson.Field
	RuleStats           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *GovernanceFrameworkGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkGetResponseJSON) RawJSON() string {
	return r.raw
}

// The icon shown for the framework.
type GovernanceFrameworkGetResponseAvatar struct {
	Type  GovernanceFrameworkGetResponseAvatarType `json:"type" api:"required"`
	Value string                                   `json:"value" api:"required"`
	JSON  governanceFrameworkGetResponseAvatarJSON `json:"-"`
}

// governanceFrameworkGetResponseAvatarJSON contains the JSON metadata for the
// struct [GovernanceFrameworkGetResponseAvatar]
type governanceFrameworkGetResponseAvatarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkGetResponseAvatar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkGetResponseAvatarJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkGetResponseAvatarType string

const (
	GovernanceFrameworkGetResponseAvatarTypeEmoji        GovernanceFrameworkGetResponseAvatarType = "emoji"
	GovernanceFrameworkGetResponseAvatarTypeImageURL     GovernanceFrameworkGetResponseAvatarType = "imageUrl"
	GovernanceFrameworkGetResponseAvatarTypeBuiltinImage GovernanceFrameworkGetResponseAvatarType = "builtinImage"
)

func (r GovernanceFrameworkGetResponseAvatarType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkGetResponseAvatarTypeEmoji, GovernanceFrameworkGetResponseAvatarTypeImageURL, GovernanceFrameworkGetResponseAvatarTypeBuiltinImage:
		return true
	}
	return false
}

// Determines which projects the framework applies to. An empty or `null` `match`
// array applies the framework to every project in the workspace.
type GovernanceFrameworkGetResponseProjectSelector struct {
	// Match criteria, ANDed together.
	Match []GovernanceFrameworkGetResponseProjectSelectorMatch `json:"match" api:"nullable"`
	JSON  governanceFrameworkGetResponseProjectSelectorJSON    `json:"-"`
}

// governanceFrameworkGetResponseProjectSelectorJSON contains the JSON metadata for
// the struct [GovernanceFrameworkGetResponseProjectSelector]
type governanceFrameworkGetResponseProjectSelectorJSON struct {
	Match       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkGetResponseProjectSelector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkGetResponseProjectSelectorJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkGetResponseProjectSelectorMatch struct {
	// The project property to match against.
	Property GovernanceFrameworkGetResponseProjectSelectorMatchProperty `json:"property" api:"required"`
	// The value to match against. Pass an array to match any of several values, or
	// `null` to match projects where the property is unset. Omit it for `exists` and
	// `notExists`.
	Value interface{} `json:"value" api:"required"`
	// How to compare the project property with `value`. One of `equals`, `notEquals`,
	// `contains`, `notContains`, `startsWith`, `endsWith`, `in`, `notIn`,
	// `greaterThan`, `greaterThanOrEqual`, `lessThan`, `lessThanOrEqual`,
	// `equalsIgnoreCase`, `containsIgnoreCase`, `matches`, `exists`, or `notExists`.
	Operator string                                                 `json:"operator"`
	JSON     governanceFrameworkGetResponseProjectSelectorMatchJSON `json:"-"`
}

// governanceFrameworkGetResponseProjectSelectorMatchJSON contains the JSON
// metadata for the struct [GovernanceFrameworkGetResponseProjectSelectorMatch]
type governanceFrameworkGetResponseProjectSelectorMatchJSON struct {
	Property    apijson.Field
	Value       apijson.Field
	Operator    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkGetResponseProjectSelectorMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkGetResponseProjectSelectorMatchJSON) RawJSON() string {
	return r.raw
}

// The project property to match against.
type GovernanceFrameworkGetResponseProjectSelectorMatchProperty string

const (
	GovernanceFrameworkGetResponseProjectSelectorMatchPropertyTaskType       GovernanceFrameworkGetResponseProjectSelectorMatchProperty = "taskType"
	GovernanceFrameworkGetResponseProjectSelectorMatchPropertyRiskLevel      GovernanceFrameworkGetResponseProjectSelectorMatchProperty = "riskLevel"
	GovernanceFrameworkGetResponseProjectSelectorMatchPropertyRiskTotalScore GovernanceFrameworkGetResponseProjectSelectorMatchProperty = "riskTotalScore"
	GovernanceFrameworkGetResponseProjectSelectorMatchPropertyName           GovernanceFrameworkGetResponseProjectSelectorMatchProperty = "name"
	GovernanceFrameworkGetResponseProjectSelectorMatchPropertyOwnerID        GovernanceFrameworkGetResponseProjectSelectorMatchProperty = "ownerId"
	GovernanceFrameworkGetResponseProjectSelectorMatchPropertyModelTypes     GovernanceFrameworkGetResponseProjectSelectorMatchProperty = "modelTypes"
)

func (r GovernanceFrameworkGetResponseProjectSelectorMatchProperty) IsKnown() bool {
	switch r {
	case GovernanceFrameworkGetResponseProjectSelectorMatchPropertyTaskType, GovernanceFrameworkGetResponseProjectSelectorMatchPropertyRiskLevel, GovernanceFrameworkGetResponseProjectSelectorMatchPropertyRiskTotalScore, GovernanceFrameworkGetResponseProjectSelectorMatchPropertyName, GovernanceFrameworkGetResponseProjectSelectorMatchPropertyOwnerID, GovernanceFrameworkGetResponseProjectSelectorMatchPropertyModelTypes:
		return true
	}
	return false
}

// Compliance roll-up for the framework. Present only on
// `GET /workspaces/{workspaceId}/frameworks` when the request sets
// `includeRuleStats=true`.
type GovernanceFrameworkGetResponseRuleStats struct {
	// How many of the framework's projects fall into each completion band, where a
	// project's completion is the share of its rule results that are passing or
	// skipped. Projects with no evaluated results count as `low`.
	//
	// Zeroed when the request carries `projectId`: the bands compare a framework's
	// projects against each other, which says nothing about a single project.
	ProjectCompletion GovernanceFrameworkGetResponseRuleStatsProjectCompletion `json:"projectCompletion" api:"required"`
	// Rule result counts by status for this framework, matching what
	// `/workspaces/{workspaceId}/rule-stats?frameworkId=<id>` reports. Narrowed to a
	// single project when the request also carries `projectId`.
	RuleResults GovernanceFrameworkGetResponseRuleStatsRuleResults `json:"ruleResults" api:"required"`
	JSON        governanceFrameworkGetResponseRuleStatsJSON        `json:"-"`
}

// governanceFrameworkGetResponseRuleStatsJSON contains the JSON metadata for the
// struct [GovernanceFrameworkGetResponseRuleStats]
type governanceFrameworkGetResponseRuleStatsJSON struct {
	ProjectCompletion apijson.Field
	RuleResults       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *GovernanceFrameworkGetResponseRuleStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkGetResponseRuleStatsJSON) RawJSON() string {
	return r.raw
}

// How many of the framework's projects fall into each completion band, where a
// project's completion is the share of its rule results that are passing or
// skipped. Projects with no evaluated results count as `low`.
//
// Zeroed when the request carries `projectId`: the bands compare a framework's
// projects against each other, which says nothing about a single project.
type GovernanceFrameworkGetResponseRuleStatsProjectCompletion struct {
	// Projects at 80% completion or above.
	High int64 `json:"high" api:"required"`
	// Projects below 20% completion.
	Low int64 `json:"low" api:"required"`
	// Projects at or above 20% but below 80% completion.
	Mid  int64                                                        `json:"mid" api:"required"`
	JSON governanceFrameworkGetResponseRuleStatsProjectCompletionJSON `json:"-"`
}

// governanceFrameworkGetResponseRuleStatsProjectCompletionJSON contains the JSON
// metadata for the struct
// [GovernanceFrameworkGetResponseRuleStatsProjectCompletion]
type governanceFrameworkGetResponseRuleStatsProjectCompletionJSON struct {
	High        apijson.Field
	Low         apijson.Field
	Mid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkGetResponseRuleStatsProjectCompletion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkGetResponseRuleStatsProjectCompletionJSON) RawJSON() string {
	return r.raw
}

// Rule result counts by status for this framework, matching what
// `/workspaces/{workspaceId}/rule-stats?frameworkId=<id>` reports. Narrowed to a
// single project when the request also carries `projectId`.
type GovernanceFrameworkGetResponseRuleStatsRuleResults struct {
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
	TotalSkipped int64                                                  `json:"totalSkipped" api:"required"`
	JSON         governanceFrameworkGetResponseRuleStatsRuleResultsJSON `json:"-"`
}

// governanceFrameworkGetResponseRuleStatsRuleResultsJSON contains the JSON
// metadata for the struct [GovernanceFrameworkGetResponseRuleStatsRuleResults]
type governanceFrameworkGetResponseRuleStatsRuleResultsJSON struct {
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

func (r *GovernanceFrameworkGetResponseRuleStatsRuleResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkGetResponseRuleStatsRuleResultsJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkUpdateResponse struct {
	// The framework id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The creation date.
	DateCreated time.Time `json:"dateCreated" api:"required" format:"date-time"`
	// The last update date.
	DateUpdated time.Time `json:"dateUpdated" api:"required" format:"date-time"`
	// Whether the framework is active. Rules of a disabled framework are not evaluated
	// and do not count towards compliance.
	Enabled bool `json:"enabled" api:"required"`
	// The framework name.
	Name string `json:"name" api:"required"`
	// Free-form labels on the framework.
	Tags []string `json:"tags" api:"required"`
	// The id of the workspace the framework belongs to.
	WorkspaceID string `json:"workspaceId" api:"required" format:"uuid"`
	// The icon shown for the framework.
	Avatar GovernanceFrameworkUpdateResponseAvatar `json:"avatar" api:"nullable"`
	// Identifies a framework that ships with Openlayer, for example `eu_ai_act`,
	// `iso_42001`, `nist_ai_rmf`, or `traiga`. `null` for frameworks you create
	// yourself.
	BuiltInSlug string `json:"builtInSlug" api:"nullable"`
	// The user who created the framework. `null` for built-in frameworks.
	CreatorID string `json:"creatorId" api:"nullable" format:"uuid"`
	// A short description of the framework.
	Description string `json:"description" api:"nullable"`
	// A longer, rich-text description, as a TipTap JSON document.
	ExtendedDescription map[string]interface{} `json:"extendedDescription" api:"nullable"`
	// A link to the external standard or regulation the framework is based on.
	Href string `json:"href" api:"nullable"`
	// Whether the framework definition is managed by Openlayer and cannot be edited.
	Immutable bool `json:"immutable"`
	// Determines which projects the framework applies to. An empty or `null` `match`
	// array applies the framework to every project in the workspace.
	ProjectSelector GovernanceFrameworkUpdateResponseProjectSelector `json:"projectSelector" api:"nullable"`
	// Compliance roll-up for the framework. Present only on
	// `GET /workspaces/{workspaceId}/frameworks` when the request sets
	// `includeRuleStats=true`.
	RuleStats GovernanceFrameworkUpdateResponseRuleStats `json:"ruleStats"`
	JSON      governanceFrameworkUpdateResponseJSON      `json:"-"`
}

// governanceFrameworkUpdateResponseJSON contains the JSON metadata for the struct
// [GovernanceFrameworkUpdateResponse]
type governanceFrameworkUpdateResponseJSON struct {
	ID                  apijson.Field
	DateCreated         apijson.Field
	DateUpdated         apijson.Field
	Enabled             apijson.Field
	Name                apijson.Field
	Tags                apijson.Field
	WorkspaceID         apijson.Field
	Avatar              apijson.Field
	BuiltInSlug         apijson.Field
	CreatorID           apijson.Field
	Description         apijson.Field
	ExtendedDescription apijson.Field
	Href                apijson.Field
	Immutable           apijson.Field
	ProjectSelector     apijson.Field
	RuleStats           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *GovernanceFrameworkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The icon shown for the framework.
type GovernanceFrameworkUpdateResponseAvatar struct {
	Type  GovernanceFrameworkUpdateResponseAvatarType `json:"type" api:"required"`
	Value string                                      `json:"value" api:"required"`
	JSON  governanceFrameworkUpdateResponseAvatarJSON `json:"-"`
}

// governanceFrameworkUpdateResponseAvatarJSON contains the JSON metadata for the
// struct [GovernanceFrameworkUpdateResponseAvatar]
type governanceFrameworkUpdateResponseAvatarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkUpdateResponseAvatar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkUpdateResponseAvatarJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkUpdateResponseAvatarType string

const (
	GovernanceFrameworkUpdateResponseAvatarTypeEmoji        GovernanceFrameworkUpdateResponseAvatarType = "emoji"
	GovernanceFrameworkUpdateResponseAvatarTypeImageURL     GovernanceFrameworkUpdateResponseAvatarType = "imageUrl"
	GovernanceFrameworkUpdateResponseAvatarTypeBuiltinImage GovernanceFrameworkUpdateResponseAvatarType = "builtinImage"
)

func (r GovernanceFrameworkUpdateResponseAvatarType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkUpdateResponseAvatarTypeEmoji, GovernanceFrameworkUpdateResponseAvatarTypeImageURL, GovernanceFrameworkUpdateResponseAvatarTypeBuiltinImage:
		return true
	}
	return false
}

// Determines which projects the framework applies to. An empty or `null` `match`
// array applies the framework to every project in the workspace.
type GovernanceFrameworkUpdateResponseProjectSelector struct {
	// Match criteria, ANDed together.
	Match []GovernanceFrameworkUpdateResponseProjectSelectorMatch `json:"match" api:"nullable"`
	JSON  governanceFrameworkUpdateResponseProjectSelectorJSON    `json:"-"`
}

// governanceFrameworkUpdateResponseProjectSelectorJSON contains the JSON metadata
// for the struct [GovernanceFrameworkUpdateResponseProjectSelector]
type governanceFrameworkUpdateResponseProjectSelectorJSON struct {
	Match       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkUpdateResponseProjectSelector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkUpdateResponseProjectSelectorJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkUpdateResponseProjectSelectorMatch struct {
	// The project property to match against.
	Property GovernanceFrameworkUpdateResponseProjectSelectorMatchProperty `json:"property" api:"required"`
	// The value to match against. Pass an array to match any of several values, or
	// `null` to match projects where the property is unset. Omit it for `exists` and
	// `notExists`.
	Value interface{} `json:"value" api:"required"`
	// How to compare the project property with `value`. One of `equals`, `notEquals`,
	// `contains`, `notContains`, `startsWith`, `endsWith`, `in`, `notIn`,
	// `greaterThan`, `greaterThanOrEqual`, `lessThan`, `lessThanOrEqual`,
	// `equalsIgnoreCase`, `containsIgnoreCase`, `matches`, `exists`, or `notExists`.
	Operator string                                                    `json:"operator"`
	JSON     governanceFrameworkUpdateResponseProjectSelectorMatchJSON `json:"-"`
}

// governanceFrameworkUpdateResponseProjectSelectorMatchJSON contains the JSON
// metadata for the struct [GovernanceFrameworkUpdateResponseProjectSelectorMatch]
type governanceFrameworkUpdateResponseProjectSelectorMatchJSON struct {
	Property    apijson.Field
	Value       apijson.Field
	Operator    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkUpdateResponseProjectSelectorMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkUpdateResponseProjectSelectorMatchJSON) RawJSON() string {
	return r.raw
}

// The project property to match against.
type GovernanceFrameworkUpdateResponseProjectSelectorMatchProperty string

const (
	GovernanceFrameworkUpdateResponseProjectSelectorMatchPropertyTaskType       GovernanceFrameworkUpdateResponseProjectSelectorMatchProperty = "taskType"
	GovernanceFrameworkUpdateResponseProjectSelectorMatchPropertyRiskLevel      GovernanceFrameworkUpdateResponseProjectSelectorMatchProperty = "riskLevel"
	GovernanceFrameworkUpdateResponseProjectSelectorMatchPropertyRiskTotalScore GovernanceFrameworkUpdateResponseProjectSelectorMatchProperty = "riskTotalScore"
	GovernanceFrameworkUpdateResponseProjectSelectorMatchPropertyName           GovernanceFrameworkUpdateResponseProjectSelectorMatchProperty = "name"
	GovernanceFrameworkUpdateResponseProjectSelectorMatchPropertyOwnerID        GovernanceFrameworkUpdateResponseProjectSelectorMatchProperty = "ownerId"
	GovernanceFrameworkUpdateResponseProjectSelectorMatchPropertyModelTypes     GovernanceFrameworkUpdateResponseProjectSelectorMatchProperty = "modelTypes"
)

func (r GovernanceFrameworkUpdateResponseProjectSelectorMatchProperty) IsKnown() bool {
	switch r {
	case GovernanceFrameworkUpdateResponseProjectSelectorMatchPropertyTaskType, GovernanceFrameworkUpdateResponseProjectSelectorMatchPropertyRiskLevel, GovernanceFrameworkUpdateResponseProjectSelectorMatchPropertyRiskTotalScore, GovernanceFrameworkUpdateResponseProjectSelectorMatchPropertyName, GovernanceFrameworkUpdateResponseProjectSelectorMatchPropertyOwnerID, GovernanceFrameworkUpdateResponseProjectSelectorMatchPropertyModelTypes:
		return true
	}
	return false
}

// Compliance roll-up for the framework. Present only on
// `GET /workspaces/{workspaceId}/frameworks` when the request sets
// `includeRuleStats=true`.
type GovernanceFrameworkUpdateResponseRuleStats struct {
	// How many of the framework's projects fall into each completion band, where a
	// project's completion is the share of its rule results that are passing or
	// skipped. Projects with no evaluated results count as `low`.
	//
	// Zeroed when the request carries `projectId`: the bands compare a framework's
	// projects against each other, which says nothing about a single project.
	ProjectCompletion GovernanceFrameworkUpdateResponseRuleStatsProjectCompletion `json:"projectCompletion" api:"required"`
	// Rule result counts by status for this framework, matching what
	// `/workspaces/{workspaceId}/rule-stats?frameworkId=<id>` reports. Narrowed to a
	// single project when the request also carries `projectId`.
	RuleResults GovernanceFrameworkUpdateResponseRuleStatsRuleResults `json:"ruleResults" api:"required"`
	JSON        governanceFrameworkUpdateResponseRuleStatsJSON        `json:"-"`
}

// governanceFrameworkUpdateResponseRuleStatsJSON contains the JSON metadata for
// the struct [GovernanceFrameworkUpdateResponseRuleStats]
type governanceFrameworkUpdateResponseRuleStatsJSON struct {
	ProjectCompletion apijson.Field
	RuleResults       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *GovernanceFrameworkUpdateResponseRuleStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkUpdateResponseRuleStatsJSON) RawJSON() string {
	return r.raw
}

// How many of the framework's projects fall into each completion band, where a
// project's completion is the share of its rule results that are passing or
// skipped. Projects with no evaluated results count as `low`.
//
// Zeroed when the request carries `projectId`: the bands compare a framework's
// projects against each other, which says nothing about a single project.
type GovernanceFrameworkUpdateResponseRuleStatsProjectCompletion struct {
	// Projects at 80% completion or above.
	High int64 `json:"high" api:"required"`
	// Projects below 20% completion.
	Low int64 `json:"low" api:"required"`
	// Projects at or above 20% but below 80% completion.
	Mid  int64                                                           `json:"mid" api:"required"`
	JSON governanceFrameworkUpdateResponseRuleStatsProjectCompletionJSON `json:"-"`
}

// governanceFrameworkUpdateResponseRuleStatsProjectCompletionJSON contains the
// JSON metadata for the struct
// [GovernanceFrameworkUpdateResponseRuleStatsProjectCompletion]
type governanceFrameworkUpdateResponseRuleStatsProjectCompletionJSON struct {
	High        apijson.Field
	Low         apijson.Field
	Mid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkUpdateResponseRuleStatsProjectCompletion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkUpdateResponseRuleStatsProjectCompletionJSON) RawJSON() string {
	return r.raw
}

// Rule result counts by status for this framework, matching what
// `/workspaces/{workspaceId}/rule-stats?frameworkId=<id>` reports. Narrowed to a
// single project when the request also carries `projectId`.
type GovernanceFrameworkUpdateResponseRuleStatsRuleResults struct {
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
	TotalSkipped int64                                                     `json:"totalSkipped" api:"required"`
	JSON         governanceFrameworkUpdateResponseRuleStatsRuleResultsJSON `json:"-"`
}

// governanceFrameworkUpdateResponseRuleStatsRuleResultsJSON contains the JSON
// metadata for the struct [GovernanceFrameworkUpdateResponseRuleStatsRuleResults]
type governanceFrameworkUpdateResponseRuleStatsRuleResultsJSON struct {
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

func (r *GovernanceFrameworkUpdateResponseRuleStatsRuleResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkUpdateResponseRuleStatsRuleResultsJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkListResponse struct {
	Items []GovernanceFrameworkListResponseItem `json:"items" api:"required"`
	JSON  governanceFrameworkListResponseJSON   `json:"-"`
}

// governanceFrameworkListResponseJSON contains the JSON metadata for the struct
// [GovernanceFrameworkListResponse]
type governanceFrameworkListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListResponseJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkListResponseItem struct {
	// The framework id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The creation date.
	DateCreated time.Time `json:"dateCreated" api:"required" format:"date-time"`
	// The last update date.
	DateUpdated time.Time `json:"dateUpdated" api:"required" format:"date-time"`
	// Whether the framework is active. Rules of a disabled framework are not evaluated
	// and do not count towards compliance.
	Enabled bool `json:"enabled" api:"required"`
	// The framework name.
	Name string `json:"name" api:"required"`
	// Free-form labels on the framework.
	Tags []string `json:"tags" api:"required"`
	// The id of the workspace the framework belongs to.
	WorkspaceID string `json:"workspaceId" api:"required" format:"uuid"`
	// The icon shown for the framework.
	Avatar GovernanceFrameworkListResponseItemsAvatar `json:"avatar" api:"nullable"`
	// Identifies a framework that ships with Openlayer, for example `eu_ai_act`,
	// `iso_42001`, `nist_ai_rmf`, or `traiga`. `null` for frameworks you create
	// yourself.
	BuiltInSlug string `json:"builtInSlug" api:"nullable"`
	// The user who created the framework. `null` for built-in frameworks.
	CreatorID string `json:"creatorId" api:"nullable" format:"uuid"`
	// A short description of the framework.
	Description string `json:"description" api:"nullable"`
	// A longer, rich-text description, as a TipTap JSON document.
	ExtendedDescription map[string]interface{} `json:"extendedDescription" api:"nullable"`
	// A link to the external standard or regulation the framework is based on.
	Href string `json:"href" api:"nullable"`
	// Whether the framework definition is managed by Openlayer and cannot be edited.
	Immutable bool `json:"immutable"`
	// Determines which projects the framework applies to. An empty or `null` `match`
	// array applies the framework to every project in the workspace.
	ProjectSelector GovernanceFrameworkListResponseItemsProjectSelector `json:"projectSelector" api:"nullable"`
	// Compliance roll-up for the framework. Present only on
	// `GET /workspaces/{workspaceId}/frameworks` when the request sets
	// `includeRuleStats=true`.
	RuleStats GovernanceFrameworkListResponseItemsRuleStats `json:"ruleStats"`
	JSON      governanceFrameworkListResponseItemJSON       `json:"-"`
}

// governanceFrameworkListResponseItemJSON contains the JSON metadata for the
// struct [GovernanceFrameworkListResponseItem]
type governanceFrameworkListResponseItemJSON struct {
	ID                  apijson.Field
	DateCreated         apijson.Field
	DateUpdated         apijson.Field
	Enabled             apijson.Field
	Name                apijson.Field
	Tags                apijson.Field
	WorkspaceID         apijson.Field
	Avatar              apijson.Field
	BuiltInSlug         apijson.Field
	CreatorID           apijson.Field
	Description         apijson.Field
	ExtendedDescription apijson.Field
	Href                apijson.Field
	Immutable           apijson.Field
	ProjectSelector     apijson.Field
	RuleStats           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *GovernanceFrameworkListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListResponseItemJSON) RawJSON() string {
	return r.raw
}

// The icon shown for the framework.
type GovernanceFrameworkListResponseItemsAvatar struct {
	Type  GovernanceFrameworkListResponseItemsAvatarType `json:"type" api:"required"`
	Value string                                         `json:"value" api:"required"`
	JSON  governanceFrameworkListResponseItemsAvatarJSON `json:"-"`
}

// governanceFrameworkListResponseItemsAvatarJSON contains the JSON metadata for
// the struct [GovernanceFrameworkListResponseItemsAvatar]
type governanceFrameworkListResponseItemsAvatarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkListResponseItemsAvatar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListResponseItemsAvatarJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkListResponseItemsAvatarType string

const (
	GovernanceFrameworkListResponseItemsAvatarTypeEmoji        GovernanceFrameworkListResponseItemsAvatarType = "emoji"
	GovernanceFrameworkListResponseItemsAvatarTypeImageURL     GovernanceFrameworkListResponseItemsAvatarType = "imageUrl"
	GovernanceFrameworkListResponseItemsAvatarTypeBuiltinImage GovernanceFrameworkListResponseItemsAvatarType = "builtinImage"
)

func (r GovernanceFrameworkListResponseItemsAvatarType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkListResponseItemsAvatarTypeEmoji, GovernanceFrameworkListResponseItemsAvatarTypeImageURL, GovernanceFrameworkListResponseItemsAvatarTypeBuiltinImage:
		return true
	}
	return false
}

// Determines which projects the framework applies to. An empty or `null` `match`
// array applies the framework to every project in the workspace.
type GovernanceFrameworkListResponseItemsProjectSelector struct {
	// Match criteria, ANDed together.
	Match []GovernanceFrameworkListResponseItemsProjectSelectorMatch `json:"match" api:"nullable"`
	JSON  governanceFrameworkListResponseItemsProjectSelectorJSON    `json:"-"`
}

// governanceFrameworkListResponseItemsProjectSelectorJSON contains the JSON
// metadata for the struct [GovernanceFrameworkListResponseItemsProjectSelector]
type governanceFrameworkListResponseItemsProjectSelectorJSON struct {
	Match       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkListResponseItemsProjectSelector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListResponseItemsProjectSelectorJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkListResponseItemsProjectSelectorMatch struct {
	// The project property to match against.
	Property GovernanceFrameworkListResponseItemsProjectSelectorMatchProperty `json:"property" api:"required"`
	// The value to match against. Pass an array to match any of several values, or
	// `null` to match projects where the property is unset. Omit it for `exists` and
	// `notExists`.
	Value interface{} `json:"value" api:"required"`
	// How to compare the project property with `value`. One of `equals`, `notEquals`,
	// `contains`, `notContains`, `startsWith`, `endsWith`, `in`, `notIn`,
	// `greaterThan`, `greaterThanOrEqual`, `lessThan`, `lessThanOrEqual`,
	// `equalsIgnoreCase`, `containsIgnoreCase`, `matches`, `exists`, or `notExists`.
	Operator string                                                       `json:"operator"`
	JSON     governanceFrameworkListResponseItemsProjectSelectorMatchJSON `json:"-"`
}

// governanceFrameworkListResponseItemsProjectSelectorMatchJSON contains the JSON
// metadata for the struct
// [GovernanceFrameworkListResponseItemsProjectSelectorMatch]
type governanceFrameworkListResponseItemsProjectSelectorMatchJSON struct {
	Property    apijson.Field
	Value       apijson.Field
	Operator    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkListResponseItemsProjectSelectorMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListResponseItemsProjectSelectorMatchJSON) RawJSON() string {
	return r.raw
}

// The project property to match against.
type GovernanceFrameworkListResponseItemsProjectSelectorMatchProperty string

const (
	GovernanceFrameworkListResponseItemsProjectSelectorMatchPropertyTaskType       GovernanceFrameworkListResponseItemsProjectSelectorMatchProperty = "taskType"
	GovernanceFrameworkListResponseItemsProjectSelectorMatchPropertyRiskLevel      GovernanceFrameworkListResponseItemsProjectSelectorMatchProperty = "riskLevel"
	GovernanceFrameworkListResponseItemsProjectSelectorMatchPropertyRiskTotalScore GovernanceFrameworkListResponseItemsProjectSelectorMatchProperty = "riskTotalScore"
	GovernanceFrameworkListResponseItemsProjectSelectorMatchPropertyName           GovernanceFrameworkListResponseItemsProjectSelectorMatchProperty = "name"
	GovernanceFrameworkListResponseItemsProjectSelectorMatchPropertyOwnerID        GovernanceFrameworkListResponseItemsProjectSelectorMatchProperty = "ownerId"
	GovernanceFrameworkListResponseItemsProjectSelectorMatchPropertyModelTypes     GovernanceFrameworkListResponseItemsProjectSelectorMatchProperty = "modelTypes"
)

func (r GovernanceFrameworkListResponseItemsProjectSelectorMatchProperty) IsKnown() bool {
	switch r {
	case GovernanceFrameworkListResponseItemsProjectSelectorMatchPropertyTaskType, GovernanceFrameworkListResponseItemsProjectSelectorMatchPropertyRiskLevel, GovernanceFrameworkListResponseItemsProjectSelectorMatchPropertyRiskTotalScore, GovernanceFrameworkListResponseItemsProjectSelectorMatchPropertyName, GovernanceFrameworkListResponseItemsProjectSelectorMatchPropertyOwnerID, GovernanceFrameworkListResponseItemsProjectSelectorMatchPropertyModelTypes:
		return true
	}
	return false
}

// Compliance roll-up for the framework. Present only on
// `GET /workspaces/{workspaceId}/frameworks` when the request sets
// `includeRuleStats=true`.
type GovernanceFrameworkListResponseItemsRuleStats struct {
	// How many of the framework's projects fall into each completion band, where a
	// project's completion is the share of its rule results that are passing or
	// skipped. Projects with no evaluated results count as `low`.
	//
	// Zeroed when the request carries `projectId`: the bands compare a framework's
	// projects against each other, which says nothing about a single project.
	ProjectCompletion GovernanceFrameworkListResponseItemsRuleStatsProjectCompletion `json:"projectCompletion" api:"required"`
	// Rule result counts by status for this framework, matching what
	// `/workspaces/{workspaceId}/rule-stats?frameworkId=<id>` reports. Narrowed to a
	// single project when the request also carries `projectId`.
	RuleResults GovernanceFrameworkListResponseItemsRuleStatsRuleResults `json:"ruleResults" api:"required"`
	JSON        governanceFrameworkListResponseItemsRuleStatsJSON        `json:"-"`
}

// governanceFrameworkListResponseItemsRuleStatsJSON contains the JSON metadata for
// the struct [GovernanceFrameworkListResponseItemsRuleStats]
type governanceFrameworkListResponseItemsRuleStatsJSON struct {
	ProjectCompletion apijson.Field
	RuleResults       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *GovernanceFrameworkListResponseItemsRuleStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListResponseItemsRuleStatsJSON) RawJSON() string {
	return r.raw
}

// How many of the framework's projects fall into each completion band, where a
// project's completion is the share of its rule results that are passing or
// skipped. Projects with no evaluated results count as `low`.
//
// Zeroed when the request carries `projectId`: the bands compare a framework's
// projects against each other, which says nothing about a single project.
type GovernanceFrameworkListResponseItemsRuleStatsProjectCompletion struct {
	// Projects at 80% completion or above.
	High int64 `json:"high" api:"required"`
	// Projects below 20% completion.
	Low int64 `json:"low" api:"required"`
	// Projects at or above 20% but below 80% completion.
	Mid  int64                                                              `json:"mid" api:"required"`
	JSON governanceFrameworkListResponseItemsRuleStatsProjectCompletionJSON `json:"-"`
}

// governanceFrameworkListResponseItemsRuleStatsProjectCompletionJSON contains the
// JSON metadata for the struct
// [GovernanceFrameworkListResponseItemsRuleStatsProjectCompletion]
type governanceFrameworkListResponseItemsRuleStatsProjectCompletionJSON struct {
	High        apijson.Field
	Low         apijson.Field
	Mid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkListResponseItemsRuleStatsProjectCompletion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListResponseItemsRuleStatsProjectCompletionJSON) RawJSON() string {
	return r.raw
}

// Rule result counts by status for this framework, matching what
// `/workspaces/{workspaceId}/rule-stats?frameworkId=<id>` reports. Narrowed to a
// single project when the request also carries `projectId`.
type GovernanceFrameworkListResponseItemsRuleStatsRuleResults struct {
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
	TotalSkipped int64                                                        `json:"totalSkipped" api:"required"`
	JSON         governanceFrameworkListResponseItemsRuleStatsRuleResultsJSON `json:"-"`
}

// governanceFrameworkListResponseItemsRuleStatsRuleResultsJSON contains the JSON
// metadata for the struct
// [GovernanceFrameworkListResponseItemsRuleStatsRuleResults]
type governanceFrameworkListResponseItemsRuleStatsRuleResultsJSON struct {
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

func (r *GovernanceFrameworkListResponseItemsRuleStatsRuleResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListResponseItemsRuleStatsRuleResultsJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkExportResponse struct {
	// The background task id, for `GET /background-tasks/{taskId}`.
	TaskResultID string `json:"taskResultId" api:"required" format:"uuid"`
	// The path to poll for this export's status and result. Already `/v1`-prefixed, so
	// it is relative to the API host rather than to the `/v1` base url -- or just pass
	// `taskResultId` to `GET /background-tasks/{taskId}`.
	TaskResultURL string                                `json:"taskResultUrl" api:"required"`
	JSON          governanceFrameworkExportResponseJSON `json:"-"`
}

// governanceFrameworkExportResponseJSON contains the JSON metadata for the struct
// [GovernanceFrameworkExportResponse]
type governanceFrameworkExportResponseJSON struct {
	TaskResultID  apijson.Field
	TaskResultURL apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GovernanceFrameworkExportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkExportResponseJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkListProjectRuleStatsResponse struct {
	Items []GovernanceFrameworkListProjectRuleStatsResponseItem `json:"items" api:"required"`
	JSON  governanceFrameworkListProjectRuleStatsResponseJSON   `json:"-"`
}

// governanceFrameworkListProjectRuleStatsResponseJSON contains the JSON metadata
// for the struct [GovernanceFrameworkListProjectRuleStatsResponse]
type governanceFrameworkListProjectRuleStatsResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkListProjectRuleStatsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListProjectRuleStatsResponseJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkListProjectRuleStatsResponseItem struct {
	// The project id.
	ProjectID string `json:"projectId" api:"required" format:"uuid"`
	// The project name.
	ProjectName string `json:"projectName" api:"required"`
	// The project's task type.
	TaskType string `json:"taskType" api:"required"`
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
	TotalSkipped int64 `json:"totalSkipped" api:"required"`
	// The same counts, broken down by the type of the rule each result belongs to.
	ByRuleType GovernanceFrameworkListProjectRuleStatsResponseItemsByRuleType `json:"byRuleType"`
	JSON       governanceFrameworkListProjectRuleStatsResponseItemJSON        `json:"-"`
}

// governanceFrameworkListProjectRuleStatsResponseItemJSON contains the JSON
// metadata for the struct [GovernanceFrameworkListProjectRuleStatsResponseItem]
type governanceFrameworkListProjectRuleStatsResponseItemJSON struct {
	ProjectID    apijson.Field
	ProjectName  apijson.Field
	TaskType     apijson.Field
	Total        apijson.Field
	TotalDueSoon apijson.Field
	TotalError   apijson.Field
	TotalFailing apijson.Field
	TotalPassing apijson.Field
	TotalPending apijson.Field
	TotalRunning apijson.Field
	TotalSkipped apijson.Field
	ByRuleType   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GovernanceFrameworkListProjectRuleStatsResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListProjectRuleStatsResponseItemJSON) RawJSON() string {
	return r.raw
}

// The same counts, broken down by the type of the rule each result belongs to.
type GovernanceFrameworkListProjectRuleStatsResponseItemsByRuleType struct {
	Evidence GovernanceFrameworkListProjectRuleStatsResponseItemsByRuleTypeEvidence `json:"evidence"`
	Platform GovernanceFrameworkListProjectRuleStatsResponseItemsByRuleTypePlatform `json:"platform"`
	JSON     governanceFrameworkListProjectRuleStatsResponseItemsByRuleTypeJSON     `json:"-"`
}

// governanceFrameworkListProjectRuleStatsResponseItemsByRuleTypeJSON contains the
// JSON metadata for the struct
// [GovernanceFrameworkListProjectRuleStatsResponseItemsByRuleType]
type governanceFrameworkListProjectRuleStatsResponseItemsByRuleTypeJSON struct {
	Evidence    apijson.Field
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkListProjectRuleStatsResponseItemsByRuleType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListProjectRuleStatsResponseItemsByRuleTypeJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkListProjectRuleStatsResponseItemsByRuleTypeEvidence struct {
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
	TotalSkipped int64                                                                      `json:"totalSkipped" api:"required"`
	JSON         governanceFrameworkListProjectRuleStatsResponseItemsByRuleTypeEvidenceJSON `json:"-"`
}

// governanceFrameworkListProjectRuleStatsResponseItemsByRuleTypeEvidenceJSON
// contains the JSON metadata for the struct
// [GovernanceFrameworkListProjectRuleStatsResponseItemsByRuleTypeEvidence]
type governanceFrameworkListProjectRuleStatsResponseItemsByRuleTypeEvidenceJSON struct {
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

func (r *GovernanceFrameworkListProjectRuleStatsResponseItemsByRuleTypeEvidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListProjectRuleStatsResponseItemsByRuleTypeEvidenceJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkListProjectRuleStatsResponseItemsByRuleTypePlatform struct {
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
	TotalSkipped int64                                                                      `json:"totalSkipped" api:"required"`
	JSON         governanceFrameworkListProjectRuleStatsResponseItemsByRuleTypePlatformJSON `json:"-"`
}

// governanceFrameworkListProjectRuleStatsResponseItemsByRuleTypePlatformJSON
// contains the JSON metadata for the struct
// [GovernanceFrameworkListProjectRuleStatsResponseItemsByRuleTypePlatform]
type governanceFrameworkListProjectRuleStatsResponseItemsByRuleTypePlatformJSON struct {
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

func (r *GovernanceFrameworkListProjectRuleStatsResponseItemsByRuleTypePlatform) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListProjectRuleStatsResponseItemsByRuleTypePlatformJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkListProjectsResponse struct {
	Items []GovernanceFrameworkListProjectsResponseItem `json:"items" api:"required"`
	JSON  governanceFrameworkListProjectsResponseJSON   `json:"-"`
}

// governanceFrameworkListProjectsResponseJSON contains the JSON metadata for the
// struct [GovernanceFrameworkListProjectsResponse]
type governanceFrameworkListProjectsResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkListProjectsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListProjectsResponseJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkListProjectsResponseItem struct {
	// The project id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The project creator id.
	CreatorID string `json:"creatorId" api:"required,nullable" format:"uuid"`
	// The project creation date.
	DateCreated time.Time `json:"dateCreated" api:"required" format:"date-time"`
	// The project last updated date.
	DateUpdated time.Time `json:"dateUpdated" api:"required" format:"date-time"`
	// The number of tests in the development mode of the project.
	DevelopmentGoalCount int64 `json:"developmentGoalCount" api:"required"`
	// The total number of tests in the project.
	GoalCount int64 `json:"goalCount" api:"required"`
	// The number of inference pipelines in the project.
	InferencePipelineCount int64 `json:"inferencePipelineCount" api:"required"`
	// Links to the project.
	Links GovernanceFrameworkListProjectsResponseItemsLinks `json:"links" api:"required"`
	// The number of tests in the monitoring mode of the project.
	MonitoringGoalCount int64 `json:"monitoringGoalCount" api:"required"`
	// The project name.
	Name string `json:"name" api:"required"`
	// The source of the project.
	Source GovernanceFrameworkListProjectsResponseItemsSource `json:"source" api:"required,nullable"`
	// The task type of the project.
	TaskType GovernanceFrameworkListProjectsResponseItemsTaskType `json:"taskType" api:"required"`
	// The number of versions (commits) in the project.
	VersionCount int64 `json:"versionCount" api:"required"`
	// The workspace id.
	WorkspaceID string `json:"workspaceId" api:"required,nullable" format:"uuid"`
	// Number of days to retain monitoring data for this project. Null means data is
	// retained indefinitely.
	DataRetentionDays int64 `json:"dataRetentionDays" api:"nullable"`
	// The project description.
	Description string                                              `json:"description" api:"nullable"`
	GitRepo     GovernanceFrameworkListProjectsResponseItemsGitRepo `json:"gitRepo" api:"nullable"`
	// Who developed the model used in this project.
	ModelDeveloper string `json:"modelDeveloper" api:"nullable"`
	// The kinds of model used in this project.
	ModelTypes []string `json:"modelTypes" api:"nullable"`
	// What the system in this project is intended to do.
	Purpose string                                          `json:"purpose" api:"nullable"`
	JSON    governanceFrameworkListProjectsResponseItemJSON `json:"-"`
}

// governanceFrameworkListProjectsResponseItemJSON contains the JSON metadata for
// the struct [GovernanceFrameworkListProjectsResponseItem]
type governanceFrameworkListProjectsResponseItemJSON struct {
	ID                     apijson.Field
	CreatorID              apijson.Field
	DateCreated            apijson.Field
	DateUpdated            apijson.Field
	DevelopmentGoalCount   apijson.Field
	GoalCount              apijson.Field
	InferencePipelineCount apijson.Field
	Links                  apijson.Field
	MonitoringGoalCount    apijson.Field
	Name                   apijson.Field
	Source                 apijson.Field
	TaskType               apijson.Field
	VersionCount           apijson.Field
	WorkspaceID            apijson.Field
	DataRetentionDays      apijson.Field
	Description            apijson.Field
	GitRepo                apijson.Field
	ModelDeveloper         apijson.Field
	ModelTypes             apijson.Field
	Purpose                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *GovernanceFrameworkListProjectsResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListProjectsResponseItemJSON) RawJSON() string {
	return r.raw
}

// Links to the project.
type GovernanceFrameworkListProjectsResponseItemsLinks struct {
	App  string                                                `json:"app" api:"required"`
	JSON governanceFrameworkListProjectsResponseItemsLinksJSON `json:"-"`
}

// governanceFrameworkListProjectsResponseItemsLinksJSON contains the JSON metadata
// for the struct [GovernanceFrameworkListProjectsResponseItemsLinks]
type governanceFrameworkListProjectsResponseItemsLinksJSON struct {
	App         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkListProjectsResponseItemsLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListProjectsResponseItemsLinksJSON) RawJSON() string {
	return r.raw
}

// The source of the project.
type GovernanceFrameworkListProjectsResponseItemsSource string

const (
	GovernanceFrameworkListProjectsResponseItemsSourceWeb GovernanceFrameworkListProjectsResponseItemsSource = "web"
	GovernanceFrameworkListProjectsResponseItemsSourceAPI GovernanceFrameworkListProjectsResponseItemsSource = "api"
)

func (r GovernanceFrameworkListProjectsResponseItemsSource) IsKnown() bool {
	switch r {
	case GovernanceFrameworkListProjectsResponseItemsSourceWeb, GovernanceFrameworkListProjectsResponseItemsSourceAPI:
		return true
	}
	return false
}

// The task type of the project.
type GovernanceFrameworkListProjectsResponseItemsTaskType string

const (
	GovernanceFrameworkListProjectsResponseItemsTaskTypeLlmBase               GovernanceFrameworkListProjectsResponseItemsTaskType = "llm-base"
	GovernanceFrameworkListProjectsResponseItemsTaskTypeTabularClassification GovernanceFrameworkListProjectsResponseItemsTaskType = "tabular-classification"
	GovernanceFrameworkListProjectsResponseItemsTaskTypeTabularRegression     GovernanceFrameworkListProjectsResponseItemsTaskType = "tabular-regression"
	GovernanceFrameworkListProjectsResponseItemsTaskTypeTextClassification    GovernanceFrameworkListProjectsResponseItemsTaskType = "text-classification"
)

func (r GovernanceFrameworkListProjectsResponseItemsTaskType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkListProjectsResponseItemsTaskTypeLlmBase, GovernanceFrameworkListProjectsResponseItemsTaskTypeTabularClassification, GovernanceFrameworkListProjectsResponseItemsTaskTypeTabularRegression, GovernanceFrameworkListProjectsResponseItemsTaskTypeTextClassification:
		return true
	}
	return false
}

type GovernanceFrameworkListProjectsResponseItemsGitRepo struct {
	ID            string                                                  `json:"id" api:"required" format:"uuid"`
	DateConnected time.Time                                               `json:"dateConnected" api:"required" format:"date-time"`
	DateUpdated   time.Time                                               `json:"dateUpdated" api:"required" format:"date-time"`
	GitAccountID  string                                                  `json:"gitAccountId" api:"required" format:"uuid"`
	GitID         int64                                                   `json:"gitId" api:"required"`
	Name          string                                                  `json:"name" api:"required"`
	Private       bool                                                    `json:"private" api:"required"`
	ProjectID     string                                                  `json:"projectId" api:"required" format:"uuid"`
	Slug          string                                                  `json:"slug" api:"required"`
	URL           string                                                  `json:"url" api:"required" format:"url"`
	Branch        string                                                  `json:"branch"`
	RootDir       string                                                  `json:"rootDir"`
	JSON          governanceFrameworkListProjectsResponseItemsGitRepoJSON `json:"-"`
}

// governanceFrameworkListProjectsResponseItemsGitRepoJSON contains the JSON
// metadata for the struct [GovernanceFrameworkListProjectsResponseItemsGitRepo]
type governanceFrameworkListProjectsResponseItemsGitRepoJSON struct {
	ID            apijson.Field
	DateConnected apijson.Field
	DateUpdated   apijson.Field
	GitAccountID  apijson.Field
	GitID         apijson.Field
	Name          apijson.Field
	Private       apijson.Field
	ProjectID     apijson.Field
	Slug          apijson.Field
	URL           apijson.Field
	Branch        apijson.Field
	RootDir       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GovernanceFrameworkListProjectsResponseItemsGitRepo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListProjectsResponseItemsGitRepoJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkListRulesResponse struct {
	Items []GovernanceFrameworkListRulesResponseItem `json:"items" api:"required"`
	JSON  governanceFrameworkListRulesResponseJSON   `json:"-"`
}

// governanceFrameworkListRulesResponseJSON contains the JSON metadata for the
// struct [GovernanceFrameworkListRulesResponse]
type governanceFrameworkListRulesResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkListRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListRulesResponseJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkListRulesResponseItem struct {
	// The rule id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The creation date.
	DateCreated time.Time `json:"dateCreated" api:"required" format:"date-time"`
	// The last update date.
	DateUpdated time.Time `json:"dateUpdated" api:"required" format:"date-time"`
	// The rule name.
	Name string `json:"name" api:"required"`
	// Whether the rule is evaluated once for the whole workspace, or once per project
	// the rule's frameworks apply to.
	Scope GovernanceFrameworkListRulesResponseItemsScope `json:"scope" api:"required"`
	// `platform` rules are evaluated automatically from the state of your Openlayer
	// workspace. `evidence` rules are satisfied by attaching evidence.
	Type GovernanceFrameworkListRulesResponseItemsType `json:"type" api:"required"`
	// The id of the workspace the rule belongs to.
	WorkspaceID string `json:"workspaceId" api:"required" format:"uuid"`
	// The user responsible for satisfying the rule.
	AssigneeID string `json:"assigneeId" api:"nullable" format:"uuid"`
	// Configuration for the platform check, when the automation takes parameters.
	AutomationParams map[string]interface{} `json:"automationParams" api:"nullable"`
	// Which workspace signal a platform rule checks, for example
	// `monitoring_mode_enabled`, `test_setup`, or `project_owner_set`. `null` for
	// evidence rules.
	AutomationType string `json:"automationType" api:"nullable"`
	// Whether the rule is excluded from compliance calculations.
	Deactivated bool `json:"deactivated"`
	// What the rule requires.
	Description string `json:"description" api:"nullable"`
	// The kind of evidence that satisfies the rule. `null` for platform rules.
	EvidenceType GovernanceFrameworkListRulesResponseItemsEvidenceType `json:"evidenceType" api:"nullable"`
	// The frameworks that include this rule.
	Frameworks []GovernanceFrameworkListRulesResponseItemsFramework `json:"frameworks"`
	// Whether the rule is managed by Openlayer and cannot be edited.
	Immutable bool `json:"immutable"`
	// How often evidence must be renewed, in days. Once evidence is older than this,
	// the rule result becomes `due_soon` and then `failing`.
	RenewalCadenceDays int64 `json:"renewalCadenceDays" api:"nullable"`
	// The rule's results, one per entity the rule is evaluated against. Only returned
	// when `includeResults` is `true`.
	Results []GovernanceFrameworkListRulesResponseItemsResult `json:"results"`
	// Pass-rate counts across all of the rule's entities, independent of any status
	// filter applied to the request.
	ResultsSummary GovernanceFrameworkListRulesResponseItemsResultsSummary `json:"resultsSummary" api:"nullable"`
	// The rule tags associated with the rule.
	Tags []GovernanceFrameworkListRulesResponseItemsTag `json:"tags" api:"nullable"`
	JSON governanceFrameworkListRulesResponseItemJSON   `json:"-"`
}

// governanceFrameworkListRulesResponseItemJSON contains the JSON metadata for the
// struct [GovernanceFrameworkListRulesResponseItem]
type governanceFrameworkListRulesResponseItemJSON struct {
	ID                 apijson.Field
	DateCreated        apijson.Field
	DateUpdated        apijson.Field
	Name               apijson.Field
	Scope              apijson.Field
	Type               apijson.Field
	WorkspaceID        apijson.Field
	AssigneeID         apijson.Field
	AutomationParams   apijson.Field
	AutomationType     apijson.Field
	Deactivated        apijson.Field
	Description        apijson.Field
	EvidenceType       apijson.Field
	Frameworks         apijson.Field
	Immutable          apijson.Field
	RenewalCadenceDays apijson.Field
	Results            apijson.Field
	ResultsSummary     apijson.Field
	Tags               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *GovernanceFrameworkListRulesResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListRulesResponseItemJSON) RawJSON() string {
	return r.raw
}

// Whether the rule is evaluated once for the whole workspace, or once per project
// the rule's frameworks apply to.
type GovernanceFrameworkListRulesResponseItemsScope string

const (
	GovernanceFrameworkListRulesResponseItemsScopeProject   GovernanceFrameworkListRulesResponseItemsScope = "project"
	GovernanceFrameworkListRulesResponseItemsScopeWorkspace GovernanceFrameworkListRulesResponseItemsScope = "workspace"
)

func (r GovernanceFrameworkListRulesResponseItemsScope) IsKnown() bool {
	switch r {
	case GovernanceFrameworkListRulesResponseItemsScopeProject, GovernanceFrameworkListRulesResponseItemsScopeWorkspace:
		return true
	}
	return false
}

// `platform` rules are evaluated automatically from the state of your Openlayer
// workspace. `evidence` rules are satisfied by attaching evidence.
type GovernanceFrameworkListRulesResponseItemsType string

const (
	GovernanceFrameworkListRulesResponseItemsTypePlatform GovernanceFrameworkListRulesResponseItemsType = "platform"
	GovernanceFrameworkListRulesResponseItemsTypeEvidence GovernanceFrameworkListRulesResponseItemsType = "evidence"
)

func (r GovernanceFrameworkListRulesResponseItemsType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkListRulesResponseItemsTypePlatform, GovernanceFrameworkListRulesResponseItemsTypeEvidence:
		return true
	}
	return false
}

// The kind of evidence that satisfies the rule. `null` for platform rules.
type GovernanceFrameworkListRulesResponseItemsEvidenceType string

const (
	GovernanceFrameworkListRulesResponseItemsEvidenceTypeDocument      GovernanceFrameworkListRulesResponseItemsEvidenceType = "document"
	GovernanceFrameworkListRulesResponseItemsEvidenceTypeText          GovernanceFrameworkListRulesResponseItemsEvidenceType = "text"
	GovernanceFrameworkListRulesResponseItemsEvidenceTypeURL           GovernanceFrameworkListRulesResponseItemsEvidenceType = "url"
	GovernanceFrameworkListRulesResponseItemsEvidenceTypeCategoryValue GovernanceFrameworkListRulesResponseItemsEvidenceType = "categoryValue"
)

func (r GovernanceFrameworkListRulesResponseItemsEvidenceType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkListRulesResponseItemsEvidenceTypeDocument, GovernanceFrameworkListRulesResponseItemsEvidenceTypeText, GovernanceFrameworkListRulesResponseItemsEvidenceTypeURL, GovernanceFrameworkListRulesResponseItemsEvidenceTypeCategoryValue:
		return true
	}
	return false
}

type GovernanceFrameworkListRulesResponseItemsFramework struct {
	// The framework id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The icon shown for the framework.
	Avatar GovernanceFrameworkListRulesResponseItemsFrameworksAvatar `json:"avatar" api:"required,nullable"`
	// Identifies a framework that ships with Openlayer, for example `eu_ai_act`,
	// `iso_42001`, `nist_ai_rmf`, or `traiga`. `null` for frameworks you create
	// yourself.
	BuiltInSlug string `json:"builtInSlug" api:"required,nullable"`
	// Whether the framework is active. Rules of a disabled framework are not evaluated
	// and do not count towards compliance.
	Enabled bool `json:"enabled" api:"required"`
	// The framework name.
	Name string                                                 `json:"name" api:"required"`
	JSON governanceFrameworkListRulesResponseItemsFrameworkJSON `json:"-"`
}

// governanceFrameworkListRulesResponseItemsFrameworkJSON contains the JSON
// metadata for the struct [GovernanceFrameworkListRulesResponseItemsFramework]
type governanceFrameworkListRulesResponseItemsFrameworkJSON struct {
	ID          apijson.Field
	Avatar      apijson.Field
	BuiltInSlug apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkListRulesResponseItemsFramework) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListRulesResponseItemsFrameworkJSON) RawJSON() string {
	return r.raw
}

// The icon shown for the framework.
type GovernanceFrameworkListRulesResponseItemsFrameworksAvatar struct {
	Type  GovernanceFrameworkListRulesResponseItemsFrameworksAvatarType `json:"type" api:"required"`
	Value string                                                        `json:"value" api:"required"`
	JSON  governanceFrameworkListRulesResponseItemsFrameworksAvatarJSON `json:"-"`
}

// governanceFrameworkListRulesResponseItemsFrameworksAvatarJSON contains the JSON
// metadata for the struct
// [GovernanceFrameworkListRulesResponseItemsFrameworksAvatar]
type governanceFrameworkListRulesResponseItemsFrameworksAvatarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkListRulesResponseItemsFrameworksAvatar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListRulesResponseItemsFrameworksAvatarJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkListRulesResponseItemsFrameworksAvatarType string

const (
	GovernanceFrameworkListRulesResponseItemsFrameworksAvatarTypeEmoji        GovernanceFrameworkListRulesResponseItemsFrameworksAvatarType = "emoji"
	GovernanceFrameworkListRulesResponseItemsFrameworksAvatarTypeImageURL     GovernanceFrameworkListRulesResponseItemsFrameworksAvatarType = "imageUrl"
	GovernanceFrameworkListRulesResponseItemsFrameworksAvatarTypeBuiltinImage GovernanceFrameworkListRulesResponseItemsFrameworksAvatarType = "builtinImage"
)

func (r GovernanceFrameworkListRulesResponseItemsFrameworksAvatarType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkListRulesResponseItemsFrameworksAvatarTypeEmoji, GovernanceFrameworkListRulesResponseItemsFrameworksAvatarTypeImageURL, GovernanceFrameworkListRulesResponseItemsFrameworksAvatarTypeBuiltinImage:
		return true
	}
	return false
}

type GovernanceFrameworkListRulesResponseItemsResult struct {
	// The rule result id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The creation date.
	DateCreated time.Time `json:"dateCreated" api:"required" format:"date-time"`
	// The last update date.
	DateUpdated time.Time `json:"dateUpdated" api:"required" format:"date-time"`
	// Whether this result is excluded from compliance calculations.
	Deactivated bool `json:"deactivated" api:"required"`
	// The rule this result belongs to.
	RuleID string `json:"ruleId" api:"required" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status GovernanceFrameworkListRulesResponseItemsResultsStatus `json:"status" api:"required"`
	// The id of the workspace the rule result belongs to.
	WorkspaceID string `json:"workspaceId" api:"required" format:"uuid"`
	// The user responsible for this result.
	AssigneeID string `json:"assigneeId" api:"nullable" format:"uuid"`
	// Rule results that must pass before this one can be satisfied.
	BlockedBy []GovernanceFrameworkListRulesResponseItemsResultsBlockedBy `json:"blockedBy"`
	// Rule results that this one blocks.
	Blocking []GovernanceFrameworkListRulesResponseItemsResultsBlocking `json:"blocking"`
	// When the rule was last evaluated. Platform rules only.
	DateLastEvaluated time.Time `json:"dateLastEvaluated" api:"nullable" format:"date-time"`
	// When the most recent piece of evidence was attached. Evidence rules only.
	DateOfLatestEvidence time.Time `json:"dateOfLatestEvidence" api:"nullable" format:"date-time"`
	// When the rule will next be evaluated. Platform rules only.
	DateOfNextEvaluation time.Time `json:"dateOfNextEvaluation" api:"nullable" format:"date-time"`
	// When the evidence must be renewed. Evidence rules with a renewal cadence only.
	DateOfRenewal time.Time `json:"dateOfRenewal" api:"nullable" format:"date-time"`
	// Why the result was excluded.
	DeactivatedReason string `json:"deactivatedReason" api:"nullable"`
	// The project this result was evaluated for. `null` for workspace-scoped rules.
	ProjectID string `json:"projectId" api:"nullable" format:"uuid"`
	// A human-readable explanation of the status.
	StatusMessage string                                              `json:"statusMessage" api:"nullable"`
	JSON          governanceFrameworkListRulesResponseItemsResultJSON `json:"-"`
}

// governanceFrameworkListRulesResponseItemsResultJSON contains the JSON metadata
// for the struct [GovernanceFrameworkListRulesResponseItemsResult]
type governanceFrameworkListRulesResponseItemsResultJSON struct {
	ID                   apijson.Field
	DateCreated          apijson.Field
	DateUpdated          apijson.Field
	Deactivated          apijson.Field
	RuleID               apijson.Field
	Status               apijson.Field
	WorkspaceID          apijson.Field
	AssigneeID           apijson.Field
	BlockedBy            apijson.Field
	Blocking             apijson.Field
	DateLastEvaluated    apijson.Field
	DateOfLatestEvidence apijson.Field
	DateOfNextEvaluation apijson.Field
	DateOfRenewal        apijson.Field
	DeactivatedReason    apijson.Field
	ProjectID            apijson.Field
	StatusMessage        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *GovernanceFrameworkListRulesResponseItemsResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListRulesResponseItemsResultJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceFrameworkListRulesResponseItemsResultsStatus string

const (
	GovernanceFrameworkListRulesResponseItemsResultsStatusRunning GovernanceFrameworkListRulesResponseItemsResultsStatus = "running"
	GovernanceFrameworkListRulesResponseItemsResultsStatusPassing GovernanceFrameworkListRulesResponseItemsResultsStatus = "passing"
	GovernanceFrameworkListRulesResponseItemsResultsStatusFailing GovernanceFrameworkListRulesResponseItemsResultsStatus = "failing"
	GovernanceFrameworkListRulesResponseItemsResultsStatusSkipped GovernanceFrameworkListRulesResponseItemsResultsStatus = "skipped"
	GovernanceFrameworkListRulesResponseItemsResultsStatusError   GovernanceFrameworkListRulesResponseItemsResultsStatus = "error"
	GovernanceFrameworkListRulesResponseItemsResultsStatusPending GovernanceFrameworkListRulesResponseItemsResultsStatus = "pending"
	GovernanceFrameworkListRulesResponseItemsResultsStatusDueSoon GovernanceFrameworkListRulesResponseItemsResultsStatus = "due_soon"
)

func (r GovernanceFrameworkListRulesResponseItemsResultsStatus) IsKnown() bool {
	switch r {
	case GovernanceFrameworkListRulesResponseItemsResultsStatusRunning, GovernanceFrameworkListRulesResponseItemsResultsStatusPassing, GovernanceFrameworkListRulesResponseItemsResultsStatusFailing, GovernanceFrameworkListRulesResponseItemsResultsStatusSkipped, GovernanceFrameworkListRulesResponseItemsResultsStatusError, GovernanceFrameworkListRulesResponseItemsResultsStatusPending, GovernanceFrameworkListRulesResponseItemsResultsStatusDueSoon:
		return true
	}
	return false
}

type GovernanceFrameworkListRulesResponseItemsResultsBlockedBy struct {
	ID string `json:"id" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatus `json:"status"`
	JSON   governanceFrameworkListRulesResponseItemsResultsBlockedByJSON   `json:"-"`
}

// governanceFrameworkListRulesResponseItemsResultsBlockedByJSON contains the JSON
// metadata for the struct
// [GovernanceFrameworkListRulesResponseItemsResultsBlockedBy]
type governanceFrameworkListRulesResponseItemsResultsBlockedByJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkListRulesResponseItemsResultsBlockedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListRulesResponseItemsResultsBlockedByJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatus string

const (
	GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatusRunning GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatus = "running"
	GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatusPassing GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatus = "passing"
	GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatusFailing GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatus = "failing"
	GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatusSkipped GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatus = "skipped"
	GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatusError   GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatus = "error"
	GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatusPending GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatus = "pending"
	GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatusDueSoon GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatus = "due_soon"
)

func (r GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatus) IsKnown() bool {
	switch r {
	case GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatusRunning, GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatusPassing, GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatusFailing, GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatusSkipped, GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatusError, GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatusPending, GovernanceFrameworkListRulesResponseItemsResultsBlockedByStatusDueSoon:
		return true
	}
	return false
}

type GovernanceFrameworkListRulesResponseItemsResultsBlocking struct {
	ID string `json:"id" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status GovernanceFrameworkListRulesResponseItemsResultsBlockingStatus `json:"status"`
	JSON   governanceFrameworkListRulesResponseItemsResultsBlockingJSON   `json:"-"`
}

// governanceFrameworkListRulesResponseItemsResultsBlockingJSON contains the JSON
// metadata for the struct
// [GovernanceFrameworkListRulesResponseItemsResultsBlocking]
type governanceFrameworkListRulesResponseItemsResultsBlockingJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkListRulesResponseItemsResultsBlocking) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListRulesResponseItemsResultsBlockingJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceFrameworkListRulesResponseItemsResultsBlockingStatus string

const (
	GovernanceFrameworkListRulesResponseItemsResultsBlockingStatusRunning GovernanceFrameworkListRulesResponseItemsResultsBlockingStatus = "running"
	GovernanceFrameworkListRulesResponseItemsResultsBlockingStatusPassing GovernanceFrameworkListRulesResponseItemsResultsBlockingStatus = "passing"
	GovernanceFrameworkListRulesResponseItemsResultsBlockingStatusFailing GovernanceFrameworkListRulesResponseItemsResultsBlockingStatus = "failing"
	GovernanceFrameworkListRulesResponseItemsResultsBlockingStatusSkipped GovernanceFrameworkListRulesResponseItemsResultsBlockingStatus = "skipped"
	GovernanceFrameworkListRulesResponseItemsResultsBlockingStatusError   GovernanceFrameworkListRulesResponseItemsResultsBlockingStatus = "error"
	GovernanceFrameworkListRulesResponseItemsResultsBlockingStatusPending GovernanceFrameworkListRulesResponseItemsResultsBlockingStatus = "pending"
	GovernanceFrameworkListRulesResponseItemsResultsBlockingStatusDueSoon GovernanceFrameworkListRulesResponseItemsResultsBlockingStatus = "due_soon"
)

func (r GovernanceFrameworkListRulesResponseItemsResultsBlockingStatus) IsKnown() bool {
	switch r {
	case GovernanceFrameworkListRulesResponseItemsResultsBlockingStatusRunning, GovernanceFrameworkListRulesResponseItemsResultsBlockingStatusPassing, GovernanceFrameworkListRulesResponseItemsResultsBlockingStatusFailing, GovernanceFrameworkListRulesResponseItemsResultsBlockingStatusSkipped, GovernanceFrameworkListRulesResponseItemsResultsBlockingStatusError, GovernanceFrameworkListRulesResponseItemsResultsBlockingStatusPending, GovernanceFrameworkListRulesResponseItemsResultsBlockingStatusDueSoon:
		return true
	}
	return false
}

// Pass-rate counts across all of the rule's entities, independent of any status
// filter applied to the request.
type GovernanceFrameworkListRulesResponseItemsResultsSummary struct {
	Passing int64                                                       `json:"passing"`
	Total   int64                                                       `json:"total"`
	JSON    governanceFrameworkListRulesResponseItemsResultsSummaryJSON `json:"-"`
}

// governanceFrameworkListRulesResponseItemsResultsSummaryJSON contains the JSON
// metadata for the struct
// [GovernanceFrameworkListRulesResponseItemsResultsSummary]
type governanceFrameworkListRulesResponseItemsResultsSummaryJSON struct {
	Passing     apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkListRulesResponseItemsResultsSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListRulesResponseItemsResultsSummaryJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkListRulesResponseItemsTag struct {
	// The rule tag id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The user who created the tag. `null` for tags that ship with Openlayer.
	CreatorID string `json:"creatorId" api:"required,nullable" format:"uuid"`
	// The creation date.
	DateCreated time.Time `json:"dateCreated" api:"required" format:"date-time"`
	// The last update date.
	DateUpdated time.Time `json:"dateUpdated" api:"required" format:"date-time"`
	// Whether the tag is managed by Openlayer and cannot be edited or deleted.
	Immutable bool `json:"immutable" api:"required"`
	// The tag name.
	Name string `json:"name" api:"required"`
	// The id of the workspace the tag belongs to.
	WorkspaceID string `json:"workspaceId" api:"required" format:"uuid"`
	// The color the tag is displayed with.
	Color string                                           `json:"color" api:"nullable"`
	JSON  governanceFrameworkListRulesResponseItemsTagJSON `json:"-"`
}

// governanceFrameworkListRulesResponseItemsTagJSON contains the JSON metadata for
// the struct [GovernanceFrameworkListRulesResponseItemsTag]
type governanceFrameworkListRulesResponseItemsTagJSON struct {
	ID          apijson.Field
	CreatorID   apijson.Field
	DateCreated apijson.Field
	DateUpdated apijson.Field
	Immutable   apijson.Field
	Name        apijson.Field
	WorkspaceID apijson.Field
	Color       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkListRulesResponseItemsTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkListRulesResponseItemsTagJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkNewParams struct {
	// The framework name.
	Name param.Field[string] `json:"name" api:"required"`
	// A short description of the framework.
	Description param.Field[string] `json:"description"`
	// Whether the framework is active. Rules of a disabled framework are not evaluated
	// and do not count towards compliance.
	Enabled param.Field[bool] `json:"enabled"`
	// Determines which projects the framework applies to. An empty or `null` `match`
	// array applies the framework to every project in the workspace.
	ProjectSelector param.Field[GovernanceFrameworkNewParamsProjectSelector] `json:"projectSelector"`
	// Free-form labels on the framework.
	Tags param.Field[[]string] `json:"tags"`
}

func (r GovernanceFrameworkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines which projects the framework applies to. An empty or `null` `match`
// array applies the framework to every project in the workspace.
type GovernanceFrameworkNewParamsProjectSelector struct {
	// Match criteria, ANDed together.
	Match param.Field[[]GovernanceFrameworkNewParamsProjectSelectorMatch] `json:"match"`
}

func (r GovernanceFrameworkNewParamsProjectSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GovernanceFrameworkNewParamsProjectSelectorMatch struct {
	// The project property to match against.
	Property param.Field[GovernanceFrameworkNewParamsProjectSelectorMatchProperty] `json:"property" api:"required"`
	// The value to match against. Pass an array to match any of several values, or
	// `null` to match projects where the property is unset. Omit it for `exists` and
	// `notExists`.
	Value param.Field[interface{}] `json:"value" api:"required"`
	// How to compare the project property with `value`. One of `equals`, `notEquals`,
	// `contains`, `notContains`, `startsWith`, `endsWith`, `in`, `notIn`,
	// `greaterThan`, `greaterThanOrEqual`, `lessThan`, `lessThanOrEqual`,
	// `equalsIgnoreCase`, `containsIgnoreCase`, `matches`, `exists`, or `notExists`.
	Operator param.Field[string] `json:"operator"`
}

func (r GovernanceFrameworkNewParamsProjectSelectorMatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The project property to match against.
type GovernanceFrameworkNewParamsProjectSelectorMatchProperty string

const (
	GovernanceFrameworkNewParamsProjectSelectorMatchPropertyTaskType       GovernanceFrameworkNewParamsProjectSelectorMatchProperty = "taskType"
	GovernanceFrameworkNewParamsProjectSelectorMatchPropertyRiskLevel      GovernanceFrameworkNewParamsProjectSelectorMatchProperty = "riskLevel"
	GovernanceFrameworkNewParamsProjectSelectorMatchPropertyRiskTotalScore GovernanceFrameworkNewParamsProjectSelectorMatchProperty = "riskTotalScore"
	GovernanceFrameworkNewParamsProjectSelectorMatchPropertyName           GovernanceFrameworkNewParamsProjectSelectorMatchProperty = "name"
	GovernanceFrameworkNewParamsProjectSelectorMatchPropertyOwnerID        GovernanceFrameworkNewParamsProjectSelectorMatchProperty = "ownerId"
	GovernanceFrameworkNewParamsProjectSelectorMatchPropertyModelTypes     GovernanceFrameworkNewParamsProjectSelectorMatchProperty = "modelTypes"
)

func (r GovernanceFrameworkNewParamsProjectSelectorMatchProperty) IsKnown() bool {
	switch r {
	case GovernanceFrameworkNewParamsProjectSelectorMatchPropertyTaskType, GovernanceFrameworkNewParamsProjectSelectorMatchPropertyRiskLevel, GovernanceFrameworkNewParamsProjectSelectorMatchPropertyRiskTotalScore, GovernanceFrameworkNewParamsProjectSelectorMatchPropertyName, GovernanceFrameworkNewParamsProjectSelectorMatchPropertyOwnerID, GovernanceFrameworkNewParamsProjectSelectorMatchPropertyModelTypes:
		return true
	}
	return false
}

type GovernanceFrameworkUpdateParams struct {
	// The icon shown for the framework.
	Avatar param.Field[GovernanceFrameworkUpdateParamsAvatar] `json:"avatar"`
	// A short description of the framework.
	Description param.Field[string] `json:"description"`
	// Whether the framework is active. Rules of a disabled framework are not evaluated
	// and do not count towards compliance.
	Enabled param.Field[bool] `json:"enabled"`
	// A longer, rich-text description, as a TipTap JSON document.
	ExtendedDescription param.Field[map[string]interface{}] `json:"extendedDescription"`
	// A link to the external standard or regulation the framework is based on.
	Href param.Field[string] `json:"href"`
	// The framework name.
	Name param.Field[string] `json:"name"`
	// Determines which projects the framework applies to. An empty or `null` `match`
	// array applies the framework to every project in the workspace.
	ProjectSelector param.Field[GovernanceFrameworkUpdateParamsProjectSelector] `json:"projectSelector"`
	// Free-form labels on the framework.
	Tags param.Field[[]string] `json:"tags"`
}

func (r GovernanceFrameworkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The icon shown for the framework.
type GovernanceFrameworkUpdateParamsAvatar struct {
	Type  param.Field[GovernanceFrameworkUpdateParamsAvatarType] `json:"type" api:"required"`
	Value param.Field[string]                                    `json:"value" api:"required"`
}

func (r GovernanceFrameworkUpdateParamsAvatar) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GovernanceFrameworkUpdateParamsAvatarType string

const (
	GovernanceFrameworkUpdateParamsAvatarTypeEmoji        GovernanceFrameworkUpdateParamsAvatarType = "emoji"
	GovernanceFrameworkUpdateParamsAvatarTypeImageURL     GovernanceFrameworkUpdateParamsAvatarType = "imageUrl"
	GovernanceFrameworkUpdateParamsAvatarTypeBuiltinImage GovernanceFrameworkUpdateParamsAvatarType = "builtinImage"
)

func (r GovernanceFrameworkUpdateParamsAvatarType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkUpdateParamsAvatarTypeEmoji, GovernanceFrameworkUpdateParamsAvatarTypeImageURL, GovernanceFrameworkUpdateParamsAvatarTypeBuiltinImage:
		return true
	}
	return false
}

// Determines which projects the framework applies to. An empty or `null` `match`
// array applies the framework to every project in the workspace.
type GovernanceFrameworkUpdateParamsProjectSelector struct {
	// Match criteria, ANDed together.
	Match param.Field[[]GovernanceFrameworkUpdateParamsProjectSelectorMatch] `json:"match"`
}

func (r GovernanceFrameworkUpdateParamsProjectSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GovernanceFrameworkUpdateParamsProjectSelectorMatch struct {
	// The project property to match against.
	Property param.Field[GovernanceFrameworkUpdateParamsProjectSelectorMatchProperty] `json:"property" api:"required"`
	// The value to match against. Pass an array to match any of several values, or
	// `null` to match projects where the property is unset. Omit it for `exists` and
	// `notExists`.
	Value param.Field[interface{}] `json:"value" api:"required"`
	// How to compare the project property with `value`. One of `equals`, `notEquals`,
	// `contains`, `notContains`, `startsWith`, `endsWith`, `in`, `notIn`,
	// `greaterThan`, `greaterThanOrEqual`, `lessThan`, `lessThanOrEqual`,
	// `equalsIgnoreCase`, `containsIgnoreCase`, `matches`, `exists`, or `notExists`.
	Operator param.Field[string] `json:"operator"`
}

func (r GovernanceFrameworkUpdateParamsProjectSelectorMatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The project property to match against.
type GovernanceFrameworkUpdateParamsProjectSelectorMatchProperty string

const (
	GovernanceFrameworkUpdateParamsProjectSelectorMatchPropertyTaskType       GovernanceFrameworkUpdateParamsProjectSelectorMatchProperty = "taskType"
	GovernanceFrameworkUpdateParamsProjectSelectorMatchPropertyRiskLevel      GovernanceFrameworkUpdateParamsProjectSelectorMatchProperty = "riskLevel"
	GovernanceFrameworkUpdateParamsProjectSelectorMatchPropertyRiskTotalScore GovernanceFrameworkUpdateParamsProjectSelectorMatchProperty = "riskTotalScore"
	GovernanceFrameworkUpdateParamsProjectSelectorMatchPropertyName           GovernanceFrameworkUpdateParamsProjectSelectorMatchProperty = "name"
	GovernanceFrameworkUpdateParamsProjectSelectorMatchPropertyOwnerID        GovernanceFrameworkUpdateParamsProjectSelectorMatchProperty = "ownerId"
	GovernanceFrameworkUpdateParamsProjectSelectorMatchPropertyModelTypes     GovernanceFrameworkUpdateParamsProjectSelectorMatchProperty = "modelTypes"
)

func (r GovernanceFrameworkUpdateParamsProjectSelectorMatchProperty) IsKnown() bool {
	switch r {
	case GovernanceFrameworkUpdateParamsProjectSelectorMatchPropertyTaskType, GovernanceFrameworkUpdateParamsProjectSelectorMatchPropertyRiskLevel, GovernanceFrameworkUpdateParamsProjectSelectorMatchPropertyRiskTotalScore, GovernanceFrameworkUpdateParamsProjectSelectorMatchPropertyName, GovernanceFrameworkUpdateParamsProjectSelectorMatchPropertyOwnerID, GovernanceFrameworkUpdateParamsProjectSelectorMatchPropertyModelTypes:
		return true
	}
	return false
}

type GovernanceFrameworkListParams struct {
	// Whether to sort in ascending order.
	Asc param.Field[bool] `query:"asc"`
	// How to compare each framework's completion percentage with `completionValue`.
	// Must be sent together with `completionValue`.
	CompletionOperator param.Field[GovernanceFrameworkListParamsCompletionOperator] `query:"completionOperator"`
	// The completion percentage to compare against, from 0 to 100.
	CompletionValue param.Field[int64] `query:"completionValue"`
	// Only include frameworks that are enabled (or disabled).
	Enabled param.Field[bool] `query:"enabled"`
	// Whether to include a `ruleStats` object on each framework, with its rule result
	// status counts and its per-project completion buckets. Computed over the returned
	// page only.
	IncludeRuleStats param.Field[bool] `query:"includeRuleStats"`
	// The page to return in a paginated query.
	Page param.Field[int64] `query:"page"`
	// Maximum number of items to return per page.
	PerPage param.Field[int64] `query:"perPage"`
	// Only include items that apply to this project.
	ProjectID param.Field[string] `query:"projectId" format:"uuid"`
	// Filter by a free-text search over names and descriptions.
	SearchQuery param.Field[string] `query:"searchQuery"`
	// The column to sort on.
	SortColumn param.Field[GovernanceFrameworkListParamsSortColumn] `query:"sortColumn"`
	// Only include frameworks carrying all of these tags.
	Tags param.Field[[]string] `query:"tags"`
}

// URLQuery serializes [GovernanceFrameworkListParams]'s query parameters as
// `url.Values`.
func (r GovernanceFrameworkListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// How to compare each framework's completion percentage with `completionValue`.
// Must be sent together with `completionValue`.
type GovernanceFrameworkListParamsCompletionOperator string

const (
	GovernanceFrameworkListParamsCompletionOperatorIs              GovernanceFrameworkListParamsCompletionOperator = "is"
	GovernanceFrameworkListParamsCompletionOperatorGreater         GovernanceFrameworkListParamsCompletionOperator = ">"
	GovernanceFrameworkListParamsCompletionOperatorGreaterOrEquals GovernanceFrameworkListParamsCompletionOperator = ">="
	GovernanceFrameworkListParamsCompletionOperatorLess            GovernanceFrameworkListParamsCompletionOperator = "<"
	GovernanceFrameworkListParamsCompletionOperatorLessOrEquals    GovernanceFrameworkListParamsCompletionOperator = "<="
	GovernanceFrameworkListParamsCompletionOperatorNotEquals       GovernanceFrameworkListParamsCompletionOperator = "!="
)

func (r GovernanceFrameworkListParamsCompletionOperator) IsKnown() bool {
	switch r {
	case GovernanceFrameworkListParamsCompletionOperatorIs, GovernanceFrameworkListParamsCompletionOperatorGreater, GovernanceFrameworkListParamsCompletionOperatorGreaterOrEquals, GovernanceFrameworkListParamsCompletionOperatorLess, GovernanceFrameworkListParamsCompletionOperatorLessOrEquals, GovernanceFrameworkListParamsCompletionOperatorNotEquals:
		return true
	}
	return false
}

// The column to sort on.
type GovernanceFrameworkListParamsSortColumn string

const (
	GovernanceFrameworkListParamsSortColumnName                     GovernanceFrameworkListParamsSortColumn = "name"
	GovernanceFrameworkListParamsSortColumnEnabled                  GovernanceFrameworkListParamsSortColumn = "enabled"
	GovernanceFrameworkListParamsSortColumnDateCreated              GovernanceFrameworkListParamsSortColumn = "dateCreated"
	GovernanceFrameworkListParamsSortColumnDateUpdated              GovernanceFrameworkListParamsSortColumn = "dateUpdated"
	GovernanceFrameworkListParamsSortColumnOverallCompletion        GovernanceFrameworkListParamsSortColumn = "overallCompletion"
	GovernanceFrameworkListParamsSortColumnProjectCompletionBuckets GovernanceFrameworkListParamsSortColumn = "projectCompletionBuckets"
)

func (r GovernanceFrameworkListParamsSortColumn) IsKnown() bool {
	switch r {
	case GovernanceFrameworkListParamsSortColumnName, GovernanceFrameworkListParamsSortColumnEnabled, GovernanceFrameworkListParamsSortColumnDateCreated, GovernanceFrameworkListParamsSortColumnDateUpdated, GovernanceFrameworkListParamsSortColumnOverallCompletion, GovernanceFrameworkListParamsSortColumnProjectCompletionBuckets:
		return true
	}
	return false
}

type GovernanceFrameworkExportParams struct {
	// Scope the export to this project. It must belong to the framework.
	ProjectID param.Field[string] `json:"projectId" format:"uuid"`
}

func (r GovernanceFrameworkExportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GovernanceFrameworkListProjectRuleStatsParams struct {
	// Whether to sort in ascending order.
	Asc param.Field[bool] `query:"asc"`
	// The page to return in a paginated query.
	Page param.Field[int64] `query:"page"`
	// Maximum number of items to return per page.
	PerPage param.Field[int64] `query:"perPage"`
	// The column to sort on.
	SortColumn param.Field[GovernanceFrameworkListProjectRuleStatsParamsSortColumn] `query:"sortColumn"`
}

// URLQuery serializes [GovernanceFrameworkListProjectRuleStatsParams]'s query
// parameters as `url.Values`.
func (r GovernanceFrameworkListProjectRuleStatsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The column to sort on.
type GovernanceFrameworkListProjectRuleStatsParamsSortColumn string

const (
	GovernanceFrameworkListProjectRuleStatsParamsSortColumnProjectName       GovernanceFrameworkListProjectRuleStatsParamsSortColumn = "projectName"
	GovernanceFrameworkListProjectRuleStatsParamsSortColumnTotal             GovernanceFrameworkListProjectRuleStatsParamsSortColumn = "total"
	GovernanceFrameworkListProjectRuleStatsParamsSortColumnOverallCompletion GovernanceFrameworkListProjectRuleStatsParamsSortColumn = "overallCompletion"
	GovernanceFrameworkListProjectRuleStatsParamsSortColumnTotalPassing      GovernanceFrameworkListProjectRuleStatsParamsSortColumn = "totalPassing"
	GovernanceFrameworkListProjectRuleStatsParamsSortColumnTotalFailing      GovernanceFrameworkListProjectRuleStatsParamsSortColumn = "totalFailing"
	GovernanceFrameworkListProjectRuleStatsParamsSortColumnTotalSkipped      GovernanceFrameworkListProjectRuleStatsParamsSortColumn = "totalSkipped"
	GovernanceFrameworkListProjectRuleStatsParamsSortColumnTotalRunning      GovernanceFrameworkListProjectRuleStatsParamsSortColumn = "totalRunning"
	GovernanceFrameworkListProjectRuleStatsParamsSortColumnTotalError        GovernanceFrameworkListProjectRuleStatsParamsSortColumn = "totalError"
	GovernanceFrameworkListProjectRuleStatsParamsSortColumnTotalPending      GovernanceFrameworkListProjectRuleStatsParamsSortColumn = "totalPending"
	GovernanceFrameworkListProjectRuleStatsParamsSortColumnTotalDueSoon      GovernanceFrameworkListProjectRuleStatsParamsSortColumn = "totalDueSoon"
)

func (r GovernanceFrameworkListProjectRuleStatsParamsSortColumn) IsKnown() bool {
	switch r {
	case GovernanceFrameworkListProjectRuleStatsParamsSortColumnProjectName, GovernanceFrameworkListProjectRuleStatsParamsSortColumnTotal, GovernanceFrameworkListProjectRuleStatsParamsSortColumnOverallCompletion, GovernanceFrameworkListProjectRuleStatsParamsSortColumnTotalPassing, GovernanceFrameworkListProjectRuleStatsParamsSortColumnTotalFailing, GovernanceFrameworkListProjectRuleStatsParamsSortColumnTotalSkipped, GovernanceFrameworkListProjectRuleStatsParamsSortColumnTotalRunning, GovernanceFrameworkListProjectRuleStatsParamsSortColumnTotalError, GovernanceFrameworkListProjectRuleStatsParamsSortColumnTotalPending, GovernanceFrameworkListProjectRuleStatsParamsSortColumnTotalDueSoon:
		return true
	}
	return false
}

type GovernanceFrameworkListProjectsParams struct {
	// The page to return in a paginated query.
	Page param.Field[int64] `query:"page"`
	// Maximum number of items to return per page.
	PerPage param.Field[int64] `query:"perPage"`
}

// URLQuery serializes [GovernanceFrameworkListProjectsParams]'s query parameters
// as `url.Values`.
func (r GovernanceFrameworkListProjectsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GovernanceFrameworkListRulesParams struct {
	// The page to return in a paginated query.
	Page param.Field[int64] `query:"page"`
	// Maximum number of items to return per page.
	PerPage param.Field[int64] `query:"perPage"`
}

// URLQuery serializes [GovernanceFrameworkListRulesParams]'s query parameters as
// `url.Values`.
func (r GovernanceFrameworkListRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
