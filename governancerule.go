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

// GovernanceRuleService contains methods and other services that help with
// interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGovernanceRuleService] method instead.
type GovernanceRuleService struct {
	Options []option.RequestOption
}

// NewGovernanceRuleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGovernanceRuleService(opts ...option.RequestOption) (r *GovernanceRuleService) {
	r = &GovernanceRuleService{}
	r.Options = opts
	return
}

// Create a governance rule in a workspace.
//
// A rule is one requirement. Its `type` decides how it is satisfied, and the two
// types accept different fields:
//
//   - `platform` rules are evaluated automatically from the state of your workspace.
//     Set `automationType` to the signal to check. Their `scope` must be `project`,
//     and `evidenceType` and `renewalCadenceDays` must be omitted or `null`.
//   - `evidence` rules are satisfied by attaching evidence. Set `evidenceType` to
//     the kind of evidence that satisfies them. `automationType` and
//     `automationParams` must be omitted or `null`.
//
// A new rule belongs to no framework. Map it to one from the Openlayer app.
func (r *GovernanceRuleService) New(ctx context.Context, workspaceID string, body GovernanceRuleNewParams, opts ...option.RequestOption) (res *GovernanceRuleNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("workspaces/%s/rules", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a governance rule by its id, including the frameworks it belongs to and
// its tags.
func (r *GovernanceRuleService) Get(ctx context.Context, ruleID string, opts ...option.RequestOption) (res *GovernanceRuleGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if ruleID == "" {
		err = errors.New("missing required ruleId parameter")
		return nil, err
	}
	path := fmt.Sprintf("rules/%s", ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a governance rule. Only the fields you send are changed.
//
// Rules that ship with Openlayer report `immutable: true` and cannot be edited.
//
// A rule's `scope`, `type`, `evidenceType`, and automation are fixed once it
// exists -- create a new rule instead of converting one.
func (r *GovernanceRuleService) Update(ctx context.Context, ruleID string, body GovernanceRuleUpdateParams, opts ...option.RequestOption) (res *GovernanceRuleUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if ruleID == "" {
		err = errors.New("missing required ruleId parameter")
		return nil, err
	}
	path := fmt.Sprintf("rules/%s", ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// List the governance rules in a workspace.
//
// A rule is a single requirement Openlayer tracks. `platform` rules are evaluated
// automatically from the state of your workspace; `evidence` rules are satisfied
// by attaching evidence. A rule can belong to several frameworks at once, and
// rules that belong to none are returned too unless you pass
// `includeUnframed=false`.
//
// Pass `includeResults=true` to get each rule's compliance results inline instead
// of fetching them separately.
func (r *GovernanceRuleService) List(ctx context.Context, workspaceID string, query GovernanceRuleListParams, opts ...option.RequestOption) (res *GovernanceRuleListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("workspaces/%s/rules", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a governance rule and its rule results.
//
// Only rules you created can be deleted. Rules that ship with Openlayer report
// `immutable: true` and cannot be deleted -- exclude one from compliance by
// setting `deactivated` with `PUT /rules/{ruleId}` instead.
func (r *GovernanceRuleService) Delete(ctx context.Context, ruleID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if ruleID == "" {
		err = errors.New("missing required ruleId parameter")
		return err
	}
	path := fmt.Sprintf("rules/%s", ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type GovernanceRuleNewResponse struct {
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
	Scope GovernanceRuleNewResponseScope `json:"scope" api:"required"`
	// `platform` rules are evaluated automatically from the state of your Openlayer
	// workspace. `evidence` rules are satisfied by attaching evidence.
	Type GovernanceRuleNewResponseType `json:"type" api:"required"`
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
	EvidenceType GovernanceRuleNewResponseEvidenceType `json:"evidenceType" api:"nullable"`
	// The frameworks that include this rule.
	Frameworks []GovernanceRuleNewResponseFramework `json:"frameworks"`
	// Whether the rule is managed by Openlayer and cannot be edited.
	Immutable bool `json:"immutable"`
	// How often evidence must be renewed, in days. Once evidence is older than this,
	// the rule result becomes `due_soon` and then `failing`.
	RenewalCadenceDays int64 `json:"renewalCadenceDays" api:"nullable"`
	// The rule's results, one per entity the rule is evaluated against. Only returned
	// when `includeResults` is `true`.
	Results []GovernanceRuleNewResponseResult `json:"results"`
	// Pass-rate counts across all of the rule's entities, independent of any status
	// filter applied to the request.
	ResultsSummary GovernanceRuleNewResponseResultsSummary `json:"resultsSummary" api:"nullable"`
	// The rule tags associated with the rule.
	Tags []GovernanceRuleNewResponseTag `json:"tags" api:"nullable"`
	JSON governanceRuleNewResponseJSON  `json:"-"`
}

// governanceRuleNewResponseJSON contains the JSON metadata for the struct
// [GovernanceRuleNewResponse]
type governanceRuleNewResponseJSON struct {
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

func (r *GovernanceRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleNewResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the rule is evaluated once for the whole workspace, or once per project
// the rule's frameworks apply to.
type GovernanceRuleNewResponseScope string

const (
	GovernanceRuleNewResponseScopeProject   GovernanceRuleNewResponseScope = "project"
	GovernanceRuleNewResponseScopeWorkspace GovernanceRuleNewResponseScope = "workspace"
)

func (r GovernanceRuleNewResponseScope) IsKnown() bool {
	switch r {
	case GovernanceRuleNewResponseScopeProject, GovernanceRuleNewResponseScopeWorkspace:
		return true
	}
	return false
}

// `platform` rules are evaluated automatically from the state of your Openlayer
// workspace. `evidence` rules are satisfied by attaching evidence.
type GovernanceRuleNewResponseType string

const (
	GovernanceRuleNewResponseTypePlatform GovernanceRuleNewResponseType = "platform"
	GovernanceRuleNewResponseTypeEvidence GovernanceRuleNewResponseType = "evidence"
)

func (r GovernanceRuleNewResponseType) IsKnown() bool {
	switch r {
	case GovernanceRuleNewResponseTypePlatform, GovernanceRuleNewResponseTypeEvidence:
		return true
	}
	return false
}

// The kind of evidence that satisfies the rule. `null` for platform rules.
type GovernanceRuleNewResponseEvidenceType string

const (
	GovernanceRuleNewResponseEvidenceTypeDocument      GovernanceRuleNewResponseEvidenceType = "document"
	GovernanceRuleNewResponseEvidenceTypeText          GovernanceRuleNewResponseEvidenceType = "text"
	GovernanceRuleNewResponseEvidenceTypeURL           GovernanceRuleNewResponseEvidenceType = "url"
	GovernanceRuleNewResponseEvidenceTypeCategoryValue GovernanceRuleNewResponseEvidenceType = "categoryValue"
)

func (r GovernanceRuleNewResponseEvidenceType) IsKnown() bool {
	switch r {
	case GovernanceRuleNewResponseEvidenceTypeDocument, GovernanceRuleNewResponseEvidenceTypeText, GovernanceRuleNewResponseEvidenceTypeURL, GovernanceRuleNewResponseEvidenceTypeCategoryValue:
		return true
	}
	return false
}

type GovernanceRuleNewResponseFramework struct {
	// The framework id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The icon shown for the framework.
	Avatar GovernanceRuleNewResponseFrameworksAvatar `json:"avatar" api:"required,nullable"`
	// Identifies a framework that ships with Openlayer, for example `eu_ai_act`,
	// `iso_42001`, `nist_ai_rmf`, or `traiga`. `null` for frameworks you create
	// yourself.
	BuiltInSlug string `json:"builtInSlug" api:"required,nullable"`
	// Whether the framework is active. Rules of a disabled framework are not evaluated
	// and do not count towards compliance.
	Enabled bool `json:"enabled" api:"required"`
	// The framework name.
	Name string                                 `json:"name" api:"required"`
	JSON governanceRuleNewResponseFrameworkJSON `json:"-"`
}

// governanceRuleNewResponseFrameworkJSON contains the JSON metadata for the struct
// [GovernanceRuleNewResponseFramework]
type governanceRuleNewResponseFrameworkJSON struct {
	ID          apijson.Field
	Avatar      apijson.Field
	BuiltInSlug apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleNewResponseFramework) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleNewResponseFrameworkJSON) RawJSON() string {
	return r.raw
}

// The icon shown for the framework.
type GovernanceRuleNewResponseFrameworksAvatar struct {
	Type  GovernanceRuleNewResponseFrameworksAvatarType `json:"type" api:"required"`
	Value string                                        `json:"value" api:"required"`
	JSON  governanceRuleNewResponseFrameworksAvatarJSON `json:"-"`
}

// governanceRuleNewResponseFrameworksAvatarJSON contains the JSON metadata for the
// struct [GovernanceRuleNewResponseFrameworksAvatar]
type governanceRuleNewResponseFrameworksAvatarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleNewResponseFrameworksAvatar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleNewResponseFrameworksAvatarJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleNewResponseFrameworksAvatarType string

const (
	GovernanceRuleNewResponseFrameworksAvatarTypeEmoji        GovernanceRuleNewResponseFrameworksAvatarType = "emoji"
	GovernanceRuleNewResponseFrameworksAvatarTypeImageURL     GovernanceRuleNewResponseFrameworksAvatarType = "imageUrl"
	GovernanceRuleNewResponseFrameworksAvatarTypeBuiltinImage GovernanceRuleNewResponseFrameworksAvatarType = "builtinImage"
)

func (r GovernanceRuleNewResponseFrameworksAvatarType) IsKnown() bool {
	switch r {
	case GovernanceRuleNewResponseFrameworksAvatarTypeEmoji, GovernanceRuleNewResponseFrameworksAvatarTypeImageURL, GovernanceRuleNewResponseFrameworksAvatarTypeBuiltinImage:
		return true
	}
	return false
}

type GovernanceRuleNewResponseResult struct {
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
	Status GovernanceRuleNewResponseResultsStatus `json:"status" api:"required"`
	// The id of the workspace the rule result belongs to.
	WorkspaceID string `json:"workspaceId" api:"required" format:"uuid"`
	// The user responsible for this result.
	AssigneeID string `json:"assigneeId" api:"nullable" format:"uuid"`
	// Rule results that must pass before this one can be satisfied.
	BlockedBy []GovernanceRuleNewResponseResultsBlockedBy `json:"blockedBy"`
	// Rule results that this one blocks.
	Blocking []GovernanceRuleNewResponseResultsBlocking `json:"blocking"`
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
	StatusMessage string                              `json:"statusMessage" api:"nullable"`
	JSON          governanceRuleNewResponseResultJSON `json:"-"`
}

// governanceRuleNewResponseResultJSON contains the JSON metadata for the struct
// [GovernanceRuleNewResponseResult]
type governanceRuleNewResponseResultJSON struct {
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

func (r *GovernanceRuleNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleNewResponseResultJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceRuleNewResponseResultsStatus string

const (
	GovernanceRuleNewResponseResultsStatusRunning GovernanceRuleNewResponseResultsStatus = "running"
	GovernanceRuleNewResponseResultsStatusPassing GovernanceRuleNewResponseResultsStatus = "passing"
	GovernanceRuleNewResponseResultsStatusFailing GovernanceRuleNewResponseResultsStatus = "failing"
	GovernanceRuleNewResponseResultsStatusSkipped GovernanceRuleNewResponseResultsStatus = "skipped"
	GovernanceRuleNewResponseResultsStatusError   GovernanceRuleNewResponseResultsStatus = "error"
	GovernanceRuleNewResponseResultsStatusPending GovernanceRuleNewResponseResultsStatus = "pending"
	GovernanceRuleNewResponseResultsStatusDueSoon GovernanceRuleNewResponseResultsStatus = "due_soon"
)

func (r GovernanceRuleNewResponseResultsStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleNewResponseResultsStatusRunning, GovernanceRuleNewResponseResultsStatusPassing, GovernanceRuleNewResponseResultsStatusFailing, GovernanceRuleNewResponseResultsStatusSkipped, GovernanceRuleNewResponseResultsStatusError, GovernanceRuleNewResponseResultsStatusPending, GovernanceRuleNewResponseResultsStatusDueSoon:
		return true
	}
	return false
}

type GovernanceRuleNewResponseResultsBlockedBy struct {
	ID string `json:"id" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status GovernanceRuleNewResponseResultsBlockedByStatus `json:"status"`
	JSON   governanceRuleNewResponseResultsBlockedByJSON   `json:"-"`
}

// governanceRuleNewResponseResultsBlockedByJSON contains the JSON metadata for the
// struct [GovernanceRuleNewResponseResultsBlockedBy]
type governanceRuleNewResponseResultsBlockedByJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleNewResponseResultsBlockedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleNewResponseResultsBlockedByJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceRuleNewResponseResultsBlockedByStatus string

const (
	GovernanceRuleNewResponseResultsBlockedByStatusRunning GovernanceRuleNewResponseResultsBlockedByStatus = "running"
	GovernanceRuleNewResponseResultsBlockedByStatusPassing GovernanceRuleNewResponseResultsBlockedByStatus = "passing"
	GovernanceRuleNewResponseResultsBlockedByStatusFailing GovernanceRuleNewResponseResultsBlockedByStatus = "failing"
	GovernanceRuleNewResponseResultsBlockedByStatusSkipped GovernanceRuleNewResponseResultsBlockedByStatus = "skipped"
	GovernanceRuleNewResponseResultsBlockedByStatusError   GovernanceRuleNewResponseResultsBlockedByStatus = "error"
	GovernanceRuleNewResponseResultsBlockedByStatusPending GovernanceRuleNewResponseResultsBlockedByStatus = "pending"
	GovernanceRuleNewResponseResultsBlockedByStatusDueSoon GovernanceRuleNewResponseResultsBlockedByStatus = "due_soon"
)

func (r GovernanceRuleNewResponseResultsBlockedByStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleNewResponseResultsBlockedByStatusRunning, GovernanceRuleNewResponseResultsBlockedByStatusPassing, GovernanceRuleNewResponseResultsBlockedByStatusFailing, GovernanceRuleNewResponseResultsBlockedByStatusSkipped, GovernanceRuleNewResponseResultsBlockedByStatusError, GovernanceRuleNewResponseResultsBlockedByStatusPending, GovernanceRuleNewResponseResultsBlockedByStatusDueSoon:
		return true
	}
	return false
}

type GovernanceRuleNewResponseResultsBlocking struct {
	ID string `json:"id" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status GovernanceRuleNewResponseResultsBlockingStatus `json:"status"`
	JSON   governanceRuleNewResponseResultsBlockingJSON   `json:"-"`
}

// governanceRuleNewResponseResultsBlockingJSON contains the JSON metadata for the
// struct [GovernanceRuleNewResponseResultsBlocking]
type governanceRuleNewResponseResultsBlockingJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleNewResponseResultsBlocking) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleNewResponseResultsBlockingJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceRuleNewResponseResultsBlockingStatus string

const (
	GovernanceRuleNewResponseResultsBlockingStatusRunning GovernanceRuleNewResponseResultsBlockingStatus = "running"
	GovernanceRuleNewResponseResultsBlockingStatusPassing GovernanceRuleNewResponseResultsBlockingStatus = "passing"
	GovernanceRuleNewResponseResultsBlockingStatusFailing GovernanceRuleNewResponseResultsBlockingStatus = "failing"
	GovernanceRuleNewResponseResultsBlockingStatusSkipped GovernanceRuleNewResponseResultsBlockingStatus = "skipped"
	GovernanceRuleNewResponseResultsBlockingStatusError   GovernanceRuleNewResponseResultsBlockingStatus = "error"
	GovernanceRuleNewResponseResultsBlockingStatusPending GovernanceRuleNewResponseResultsBlockingStatus = "pending"
	GovernanceRuleNewResponseResultsBlockingStatusDueSoon GovernanceRuleNewResponseResultsBlockingStatus = "due_soon"
)

func (r GovernanceRuleNewResponseResultsBlockingStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleNewResponseResultsBlockingStatusRunning, GovernanceRuleNewResponseResultsBlockingStatusPassing, GovernanceRuleNewResponseResultsBlockingStatusFailing, GovernanceRuleNewResponseResultsBlockingStatusSkipped, GovernanceRuleNewResponseResultsBlockingStatusError, GovernanceRuleNewResponseResultsBlockingStatusPending, GovernanceRuleNewResponseResultsBlockingStatusDueSoon:
		return true
	}
	return false
}

// Pass-rate counts across all of the rule's entities, independent of any status
// filter applied to the request.
type GovernanceRuleNewResponseResultsSummary struct {
	Passing int64                                       `json:"passing"`
	Total   int64                                       `json:"total"`
	JSON    governanceRuleNewResponseResultsSummaryJSON `json:"-"`
}

// governanceRuleNewResponseResultsSummaryJSON contains the JSON metadata for the
// struct [GovernanceRuleNewResponseResultsSummary]
type governanceRuleNewResponseResultsSummaryJSON struct {
	Passing     apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleNewResponseResultsSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleNewResponseResultsSummaryJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleNewResponseTag struct {
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
	Color string                           `json:"color" api:"nullable"`
	JSON  governanceRuleNewResponseTagJSON `json:"-"`
}

// governanceRuleNewResponseTagJSON contains the JSON metadata for the struct
// [GovernanceRuleNewResponseTag]
type governanceRuleNewResponseTagJSON struct {
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

func (r *GovernanceRuleNewResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleNewResponseTagJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleGetResponse struct {
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
	Scope GovernanceRuleGetResponseScope `json:"scope" api:"required"`
	// `platform` rules are evaluated automatically from the state of your Openlayer
	// workspace. `evidence` rules are satisfied by attaching evidence.
	Type GovernanceRuleGetResponseType `json:"type" api:"required"`
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
	EvidenceType GovernanceRuleGetResponseEvidenceType `json:"evidenceType" api:"nullable"`
	// The frameworks that include this rule.
	Frameworks []GovernanceRuleGetResponseFramework `json:"frameworks"`
	// Whether the rule is managed by Openlayer and cannot be edited.
	Immutable bool `json:"immutable"`
	// How often evidence must be renewed, in days. Once evidence is older than this,
	// the rule result becomes `due_soon` and then `failing`.
	RenewalCadenceDays int64 `json:"renewalCadenceDays" api:"nullable"`
	// The rule's results, one per entity the rule is evaluated against. Only returned
	// when `includeResults` is `true`.
	Results []GovernanceRuleGetResponseResult `json:"results"`
	// Pass-rate counts across all of the rule's entities, independent of any status
	// filter applied to the request.
	ResultsSummary GovernanceRuleGetResponseResultsSummary `json:"resultsSummary" api:"nullable"`
	// The rule tags associated with the rule.
	Tags []GovernanceRuleGetResponseTag `json:"tags" api:"nullable"`
	JSON governanceRuleGetResponseJSON  `json:"-"`
}

// governanceRuleGetResponseJSON contains the JSON metadata for the struct
// [GovernanceRuleGetResponse]
type governanceRuleGetResponseJSON struct {
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

func (r *GovernanceRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the rule is evaluated once for the whole workspace, or once per project
// the rule's frameworks apply to.
type GovernanceRuleGetResponseScope string

const (
	GovernanceRuleGetResponseScopeProject   GovernanceRuleGetResponseScope = "project"
	GovernanceRuleGetResponseScopeWorkspace GovernanceRuleGetResponseScope = "workspace"
)

func (r GovernanceRuleGetResponseScope) IsKnown() bool {
	switch r {
	case GovernanceRuleGetResponseScopeProject, GovernanceRuleGetResponseScopeWorkspace:
		return true
	}
	return false
}

// `platform` rules are evaluated automatically from the state of your Openlayer
// workspace. `evidence` rules are satisfied by attaching evidence.
type GovernanceRuleGetResponseType string

const (
	GovernanceRuleGetResponseTypePlatform GovernanceRuleGetResponseType = "platform"
	GovernanceRuleGetResponseTypeEvidence GovernanceRuleGetResponseType = "evidence"
)

func (r GovernanceRuleGetResponseType) IsKnown() bool {
	switch r {
	case GovernanceRuleGetResponseTypePlatform, GovernanceRuleGetResponseTypeEvidence:
		return true
	}
	return false
}

// The kind of evidence that satisfies the rule. `null` for platform rules.
type GovernanceRuleGetResponseEvidenceType string

const (
	GovernanceRuleGetResponseEvidenceTypeDocument      GovernanceRuleGetResponseEvidenceType = "document"
	GovernanceRuleGetResponseEvidenceTypeText          GovernanceRuleGetResponseEvidenceType = "text"
	GovernanceRuleGetResponseEvidenceTypeURL           GovernanceRuleGetResponseEvidenceType = "url"
	GovernanceRuleGetResponseEvidenceTypeCategoryValue GovernanceRuleGetResponseEvidenceType = "categoryValue"
)

func (r GovernanceRuleGetResponseEvidenceType) IsKnown() bool {
	switch r {
	case GovernanceRuleGetResponseEvidenceTypeDocument, GovernanceRuleGetResponseEvidenceTypeText, GovernanceRuleGetResponseEvidenceTypeURL, GovernanceRuleGetResponseEvidenceTypeCategoryValue:
		return true
	}
	return false
}

type GovernanceRuleGetResponseFramework struct {
	// The framework id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The icon shown for the framework.
	Avatar GovernanceRuleGetResponseFrameworksAvatar `json:"avatar" api:"required,nullable"`
	// Identifies a framework that ships with Openlayer, for example `eu_ai_act`,
	// `iso_42001`, `nist_ai_rmf`, or `traiga`. `null` for frameworks you create
	// yourself.
	BuiltInSlug string `json:"builtInSlug" api:"required,nullable"`
	// Whether the framework is active. Rules of a disabled framework are not evaluated
	// and do not count towards compliance.
	Enabled bool `json:"enabled" api:"required"`
	// The framework name.
	Name string                                 `json:"name" api:"required"`
	JSON governanceRuleGetResponseFrameworkJSON `json:"-"`
}

// governanceRuleGetResponseFrameworkJSON contains the JSON metadata for the struct
// [GovernanceRuleGetResponseFramework]
type governanceRuleGetResponseFrameworkJSON struct {
	ID          apijson.Field
	Avatar      apijson.Field
	BuiltInSlug apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleGetResponseFramework) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleGetResponseFrameworkJSON) RawJSON() string {
	return r.raw
}

// The icon shown for the framework.
type GovernanceRuleGetResponseFrameworksAvatar struct {
	Type  GovernanceRuleGetResponseFrameworksAvatarType `json:"type" api:"required"`
	Value string                                        `json:"value" api:"required"`
	JSON  governanceRuleGetResponseFrameworksAvatarJSON `json:"-"`
}

// governanceRuleGetResponseFrameworksAvatarJSON contains the JSON metadata for the
// struct [GovernanceRuleGetResponseFrameworksAvatar]
type governanceRuleGetResponseFrameworksAvatarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleGetResponseFrameworksAvatar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleGetResponseFrameworksAvatarJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleGetResponseFrameworksAvatarType string

const (
	GovernanceRuleGetResponseFrameworksAvatarTypeEmoji        GovernanceRuleGetResponseFrameworksAvatarType = "emoji"
	GovernanceRuleGetResponseFrameworksAvatarTypeImageURL     GovernanceRuleGetResponseFrameworksAvatarType = "imageUrl"
	GovernanceRuleGetResponseFrameworksAvatarTypeBuiltinImage GovernanceRuleGetResponseFrameworksAvatarType = "builtinImage"
)

func (r GovernanceRuleGetResponseFrameworksAvatarType) IsKnown() bool {
	switch r {
	case GovernanceRuleGetResponseFrameworksAvatarTypeEmoji, GovernanceRuleGetResponseFrameworksAvatarTypeImageURL, GovernanceRuleGetResponseFrameworksAvatarTypeBuiltinImage:
		return true
	}
	return false
}

type GovernanceRuleGetResponseResult struct {
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
	Status GovernanceRuleGetResponseResultsStatus `json:"status" api:"required"`
	// The id of the workspace the rule result belongs to.
	WorkspaceID string `json:"workspaceId" api:"required" format:"uuid"`
	// The user responsible for this result.
	AssigneeID string `json:"assigneeId" api:"nullable" format:"uuid"`
	// Rule results that must pass before this one can be satisfied.
	BlockedBy []GovernanceRuleGetResponseResultsBlockedBy `json:"blockedBy"`
	// Rule results that this one blocks.
	Blocking []GovernanceRuleGetResponseResultsBlocking `json:"blocking"`
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
	StatusMessage string                              `json:"statusMessage" api:"nullable"`
	JSON          governanceRuleGetResponseResultJSON `json:"-"`
}

// governanceRuleGetResponseResultJSON contains the JSON metadata for the struct
// [GovernanceRuleGetResponseResult]
type governanceRuleGetResponseResultJSON struct {
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

func (r *GovernanceRuleGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceRuleGetResponseResultsStatus string

const (
	GovernanceRuleGetResponseResultsStatusRunning GovernanceRuleGetResponseResultsStatus = "running"
	GovernanceRuleGetResponseResultsStatusPassing GovernanceRuleGetResponseResultsStatus = "passing"
	GovernanceRuleGetResponseResultsStatusFailing GovernanceRuleGetResponseResultsStatus = "failing"
	GovernanceRuleGetResponseResultsStatusSkipped GovernanceRuleGetResponseResultsStatus = "skipped"
	GovernanceRuleGetResponseResultsStatusError   GovernanceRuleGetResponseResultsStatus = "error"
	GovernanceRuleGetResponseResultsStatusPending GovernanceRuleGetResponseResultsStatus = "pending"
	GovernanceRuleGetResponseResultsStatusDueSoon GovernanceRuleGetResponseResultsStatus = "due_soon"
)

func (r GovernanceRuleGetResponseResultsStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleGetResponseResultsStatusRunning, GovernanceRuleGetResponseResultsStatusPassing, GovernanceRuleGetResponseResultsStatusFailing, GovernanceRuleGetResponseResultsStatusSkipped, GovernanceRuleGetResponseResultsStatusError, GovernanceRuleGetResponseResultsStatusPending, GovernanceRuleGetResponseResultsStatusDueSoon:
		return true
	}
	return false
}

type GovernanceRuleGetResponseResultsBlockedBy struct {
	ID string `json:"id" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status GovernanceRuleGetResponseResultsBlockedByStatus `json:"status"`
	JSON   governanceRuleGetResponseResultsBlockedByJSON   `json:"-"`
}

// governanceRuleGetResponseResultsBlockedByJSON contains the JSON metadata for the
// struct [GovernanceRuleGetResponseResultsBlockedBy]
type governanceRuleGetResponseResultsBlockedByJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleGetResponseResultsBlockedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleGetResponseResultsBlockedByJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceRuleGetResponseResultsBlockedByStatus string

const (
	GovernanceRuleGetResponseResultsBlockedByStatusRunning GovernanceRuleGetResponseResultsBlockedByStatus = "running"
	GovernanceRuleGetResponseResultsBlockedByStatusPassing GovernanceRuleGetResponseResultsBlockedByStatus = "passing"
	GovernanceRuleGetResponseResultsBlockedByStatusFailing GovernanceRuleGetResponseResultsBlockedByStatus = "failing"
	GovernanceRuleGetResponseResultsBlockedByStatusSkipped GovernanceRuleGetResponseResultsBlockedByStatus = "skipped"
	GovernanceRuleGetResponseResultsBlockedByStatusError   GovernanceRuleGetResponseResultsBlockedByStatus = "error"
	GovernanceRuleGetResponseResultsBlockedByStatusPending GovernanceRuleGetResponseResultsBlockedByStatus = "pending"
	GovernanceRuleGetResponseResultsBlockedByStatusDueSoon GovernanceRuleGetResponseResultsBlockedByStatus = "due_soon"
)

func (r GovernanceRuleGetResponseResultsBlockedByStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleGetResponseResultsBlockedByStatusRunning, GovernanceRuleGetResponseResultsBlockedByStatusPassing, GovernanceRuleGetResponseResultsBlockedByStatusFailing, GovernanceRuleGetResponseResultsBlockedByStatusSkipped, GovernanceRuleGetResponseResultsBlockedByStatusError, GovernanceRuleGetResponseResultsBlockedByStatusPending, GovernanceRuleGetResponseResultsBlockedByStatusDueSoon:
		return true
	}
	return false
}

type GovernanceRuleGetResponseResultsBlocking struct {
	ID string `json:"id" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status GovernanceRuleGetResponseResultsBlockingStatus `json:"status"`
	JSON   governanceRuleGetResponseResultsBlockingJSON   `json:"-"`
}

// governanceRuleGetResponseResultsBlockingJSON contains the JSON metadata for the
// struct [GovernanceRuleGetResponseResultsBlocking]
type governanceRuleGetResponseResultsBlockingJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleGetResponseResultsBlocking) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleGetResponseResultsBlockingJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceRuleGetResponseResultsBlockingStatus string

const (
	GovernanceRuleGetResponseResultsBlockingStatusRunning GovernanceRuleGetResponseResultsBlockingStatus = "running"
	GovernanceRuleGetResponseResultsBlockingStatusPassing GovernanceRuleGetResponseResultsBlockingStatus = "passing"
	GovernanceRuleGetResponseResultsBlockingStatusFailing GovernanceRuleGetResponseResultsBlockingStatus = "failing"
	GovernanceRuleGetResponseResultsBlockingStatusSkipped GovernanceRuleGetResponseResultsBlockingStatus = "skipped"
	GovernanceRuleGetResponseResultsBlockingStatusError   GovernanceRuleGetResponseResultsBlockingStatus = "error"
	GovernanceRuleGetResponseResultsBlockingStatusPending GovernanceRuleGetResponseResultsBlockingStatus = "pending"
	GovernanceRuleGetResponseResultsBlockingStatusDueSoon GovernanceRuleGetResponseResultsBlockingStatus = "due_soon"
)

func (r GovernanceRuleGetResponseResultsBlockingStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleGetResponseResultsBlockingStatusRunning, GovernanceRuleGetResponseResultsBlockingStatusPassing, GovernanceRuleGetResponseResultsBlockingStatusFailing, GovernanceRuleGetResponseResultsBlockingStatusSkipped, GovernanceRuleGetResponseResultsBlockingStatusError, GovernanceRuleGetResponseResultsBlockingStatusPending, GovernanceRuleGetResponseResultsBlockingStatusDueSoon:
		return true
	}
	return false
}

// Pass-rate counts across all of the rule's entities, independent of any status
// filter applied to the request.
type GovernanceRuleGetResponseResultsSummary struct {
	Passing int64                                       `json:"passing"`
	Total   int64                                       `json:"total"`
	JSON    governanceRuleGetResponseResultsSummaryJSON `json:"-"`
}

// governanceRuleGetResponseResultsSummaryJSON contains the JSON metadata for the
// struct [GovernanceRuleGetResponseResultsSummary]
type governanceRuleGetResponseResultsSummaryJSON struct {
	Passing     apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleGetResponseResultsSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleGetResponseResultsSummaryJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleGetResponseTag struct {
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
	Color string                           `json:"color" api:"nullable"`
	JSON  governanceRuleGetResponseTagJSON `json:"-"`
}

// governanceRuleGetResponseTagJSON contains the JSON metadata for the struct
// [GovernanceRuleGetResponseTag]
type governanceRuleGetResponseTagJSON struct {
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

func (r *GovernanceRuleGetResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleGetResponseTagJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleUpdateResponse struct {
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
	Scope GovernanceRuleUpdateResponseScope `json:"scope" api:"required"`
	// `platform` rules are evaluated automatically from the state of your Openlayer
	// workspace. `evidence` rules are satisfied by attaching evidence.
	Type GovernanceRuleUpdateResponseType `json:"type" api:"required"`
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
	EvidenceType GovernanceRuleUpdateResponseEvidenceType `json:"evidenceType" api:"nullable"`
	// The frameworks that include this rule.
	Frameworks []GovernanceRuleUpdateResponseFramework `json:"frameworks"`
	// Whether the rule is managed by Openlayer and cannot be edited.
	Immutable bool `json:"immutable"`
	// How often evidence must be renewed, in days. Once evidence is older than this,
	// the rule result becomes `due_soon` and then `failing`.
	RenewalCadenceDays int64 `json:"renewalCadenceDays" api:"nullable"`
	// The rule's results, one per entity the rule is evaluated against. Only returned
	// when `includeResults` is `true`.
	Results []GovernanceRuleUpdateResponseResult `json:"results"`
	// Pass-rate counts across all of the rule's entities, independent of any status
	// filter applied to the request.
	ResultsSummary GovernanceRuleUpdateResponseResultsSummary `json:"resultsSummary" api:"nullable"`
	// The rule tags associated with the rule.
	Tags []GovernanceRuleUpdateResponseTag `json:"tags" api:"nullable"`
	JSON governanceRuleUpdateResponseJSON  `json:"-"`
}

// governanceRuleUpdateResponseJSON contains the JSON metadata for the struct
// [GovernanceRuleUpdateResponse]
type governanceRuleUpdateResponseJSON struct {
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

func (r *GovernanceRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the rule is evaluated once for the whole workspace, or once per project
// the rule's frameworks apply to.
type GovernanceRuleUpdateResponseScope string

const (
	GovernanceRuleUpdateResponseScopeProject   GovernanceRuleUpdateResponseScope = "project"
	GovernanceRuleUpdateResponseScopeWorkspace GovernanceRuleUpdateResponseScope = "workspace"
)

func (r GovernanceRuleUpdateResponseScope) IsKnown() bool {
	switch r {
	case GovernanceRuleUpdateResponseScopeProject, GovernanceRuleUpdateResponseScopeWorkspace:
		return true
	}
	return false
}

// `platform` rules are evaluated automatically from the state of your Openlayer
// workspace. `evidence` rules are satisfied by attaching evidence.
type GovernanceRuleUpdateResponseType string

const (
	GovernanceRuleUpdateResponseTypePlatform GovernanceRuleUpdateResponseType = "platform"
	GovernanceRuleUpdateResponseTypeEvidence GovernanceRuleUpdateResponseType = "evidence"
)

func (r GovernanceRuleUpdateResponseType) IsKnown() bool {
	switch r {
	case GovernanceRuleUpdateResponseTypePlatform, GovernanceRuleUpdateResponseTypeEvidence:
		return true
	}
	return false
}

// The kind of evidence that satisfies the rule. `null` for platform rules.
type GovernanceRuleUpdateResponseEvidenceType string

const (
	GovernanceRuleUpdateResponseEvidenceTypeDocument      GovernanceRuleUpdateResponseEvidenceType = "document"
	GovernanceRuleUpdateResponseEvidenceTypeText          GovernanceRuleUpdateResponseEvidenceType = "text"
	GovernanceRuleUpdateResponseEvidenceTypeURL           GovernanceRuleUpdateResponseEvidenceType = "url"
	GovernanceRuleUpdateResponseEvidenceTypeCategoryValue GovernanceRuleUpdateResponseEvidenceType = "categoryValue"
)

func (r GovernanceRuleUpdateResponseEvidenceType) IsKnown() bool {
	switch r {
	case GovernanceRuleUpdateResponseEvidenceTypeDocument, GovernanceRuleUpdateResponseEvidenceTypeText, GovernanceRuleUpdateResponseEvidenceTypeURL, GovernanceRuleUpdateResponseEvidenceTypeCategoryValue:
		return true
	}
	return false
}

type GovernanceRuleUpdateResponseFramework struct {
	// The framework id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The icon shown for the framework.
	Avatar GovernanceRuleUpdateResponseFrameworksAvatar `json:"avatar" api:"required,nullable"`
	// Identifies a framework that ships with Openlayer, for example `eu_ai_act`,
	// `iso_42001`, `nist_ai_rmf`, or `traiga`. `null` for frameworks you create
	// yourself.
	BuiltInSlug string `json:"builtInSlug" api:"required,nullable"`
	// Whether the framework is active. Rules of a disabled framework are not evaluated
	// and do not count towards compliance.
	Enabled bool `json:"enabled" api:"required"`
	// The framework name.
	Name string                                    `json:"name" api:"required"`
	JSON governanceRuleUpdateResponseFrameworkJSON `json:"-"`
}

// governanceRuleUpdateResponseFrameworkJSON contains the JSON metadata for the
// struct [GovernanceRuleUpdateResponseFramework]
type governanceRuleUpdateResponseFrameworkJSON struct {
	ID          apijson.Field
	Avatar      apijson.Field
	BuiltInSlug apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleUpdateResponseFramework) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleUpdateResponseFrameworkJSON) RawJSON() string {
	return r.raw
}

// The icon shown for the framework.
type GovernanceRuleUpdateResponseFrameworksAvatar struct {
	Type  GovernanceRuleUpdateResponseFrameworksAvatarType `json:"type" api:"required"`
	Value string                                           `json:"value" api:"required"`
	JSON  governanceRuleUpdateResponseFrameworksAvatarJSON `json:"-"`
}

// governanceRuleUpdateResponseFrameworksAvatarJSON contains the JSON metadata for
// the struct [GovernanceRuleUpdateResponseFrameworksAvatar]
type governanceRuleUpdateResponseFrameworksAvatarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleUpdateResponseFrameworksAvatar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleUpdateResponseFrameworksAvatarJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleUpdateResponseFrameworksAvatarType string

const (
	GovernanceRuleUpdateResponseFrameworksAvatarTypeEmoji        GovernanceRuleUpdateResponseFrameworksAvatarType = "emoji"
	GovernanceRuleUpdateResponseFrameworksAvatarTypeImageURL     GovernanceRuleUpdateResponseFrameworksAvatarType = "imageUrl"
	GovernanceRuleUpdateResponseFrameworksAvatarTypeBuiltinImage GovernanceRuleUpdateResponseFrameworksAvatarType = "builtinImage"
)

func (r GovernanceRuleUpdateResponseFrameworksAvatarType) IsKnown() bool {
	switch r {
	case GovernanceRuleUpdateResponseFrameworksAvatarTypeEmoji, GovernanceRuleUpdateResponseFrameworksAvatarTypeImageURL, GovernanceRuleUpdateResponseFrameworksAvatarTypeBuiltinImage:
		return true
	}
	return false
}

type GovernanceRuleUpdateResponseResult struct {
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
	Status GovernanceRuleUpdateResponseResultsStatus `json:"status" api:"required"`
	// The id of the workspace the rule result belongs to.
	WorkspaceID string `json:"workspaceId" api:"required" format:"uuid"`
	// The user responsible for this result.
	AssigneeID string `json:"assigneeId" api:"nullable" format:"uuid"`
	// Rule results that must pass before this one can be satisfied.
	BlockedBy []GovernanceRuleUpdateResponseResultsBlockedBy `json:"blockedBy"`
	// Rule results that this one blocks.
	Blocking []GovernanceRuleUpdateResponseResultsBlocking `json:"blocking"`
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
	StatusMessage string                                 `json:"statusMessage" api:"nullable"`
	JSON          governanceRuleUpdateResponseResultJSON `json:"-"`
}

// governanceRuleUpdateResponseResultJSON contains the JSON metadata for the struct
// [GovernanceRuleUpdateResponseResult]
type governanceRuleUpdateResponseResultJSON struct {
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

func (r *GovernanceRuleUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceRuleUpdateResponseResultsStatus string

const (
	GovernanceRuleUpdateResponseResultsStatusRunning GovernanceRuleUpdateResponseResultsStatus = "running"
	GovernanceRuleUpdateResponseResultsStatusPassing GovernanceRuleUpdateResponseResultsStatus = "passing"
	GovernanceRuleUpdateResponseResultsStatusFailing GovernanceRuleUpdateResponseResultsStatus = "failing"
	GovernanceRuleUpdateResponseResultsStatusSkipped GovernanceRuleUpdateResponseResultsStatus = "skipped"
	GovernanceRuleUpdateResponseResultsStatusError   GovernanceRuleUpdateResponseResultsStatus = "error"
	GovernanceRuleUpdateResponseResultsStatusPending GovernanceRuleUpdateResponseResultsStatus = "pending"
	GovernanceRuleUpdateResponseResultsStatusDueSoon GovernanceRuleUpdateResponseResultsStatus = "due_soon"
)

func (r GovernanceRuleUpdateResponseResultsStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleUpdateResponseResultsStatusRunning, GovernanceRuleUpdateResponseResultsStatusPassing, GovernanceRuleUpdateResponseResultsStatusFailing, GovernanceRuleUpdateResponseResultsStatusSkipped, GovernanceRuleUpdateResponseResultsStatusError, GovernanceRuleUpdateResponseResultsStatusPending, GovernanceRuleUpdateResponseResultsStatusDueSoon:
		return true
	}
	return false
}

type GovernanceRuleUpdateResponseResultsBlockedBy struct {
	ID string `json:"id" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status GovernanceRuleUpdateResponseResultsBlockedByStatus `json:"status"`
	JSON   governanceRuleUpdateResponseResultsBlockedByJSON   `json:"-"`
}

// governanceRuleUpdateResponseResultsBlockedByJSON contains the JSON metadata for
// the struct [GovernanceRuleUpdateResponseResultsBlockedBy]
type governanceRuleUpdateResponseResultsBlockedByJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleUpdateResponseResultsBlockedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleUpdateResponseResultsBlockedByJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceRuleUpdateResponseResultsBlockedByStatus string

const (
	GovernanceRuleUpdateResponseResultsBlockedByStatusRunning GovernanceRuleUpdateResponseResultsBlockedByStatus = "running"
	GovernanceRuleUpdateResponseResultsBlockedByStatusPassing GovernanceRuleUpdateResponseResultsBlockedByStatus = "passing"
	GovernanceRuleUpdateResponseResultsBlockedByStatusFailing GovernanceRuleUpdateResponseResultsBlockedByStatus = "failing"
	GovernanceRuleUpdateResponseResultsBlockedByStatusSkipped GovernanceRuleUpdateResponseResultsBlockedByStatus = "skipped"
	GovernanceRuleUpdateResponseResultsBlockedByStatusError   GovernanceRuleUpdateResponseResultsBlockedByStatus = "error"
	GovernanceRuleUpdateResponseResultsBlockedByStatusPending GovernanceRuleUpdateResponseResultsBlockedByStatus = "pending"
	GovernanceRuleUpdateResponseResultsBlockedByStatusDueSoon GovernanceRuleUpdateResponseResultsBlockedByStatus = "due_soon"
)

func (r GovernanceRuleUpdateResponseResultsBlockedByStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleUpdateResponseResultsBlockedByStatusRunning, GovernanceRuleUpdateResponseResultsBlockedByStatusPassing, GovernanceRuleUpdateResponseResultsBlockedByStatusFailing, GovernanceRuleUpdateResponseResultsBlockedByStatusSkipped, GovernanceRuleUpdateResponseResultsBlockedByStatusError, GovernanceRuleUpdateResponseResultsBlockedByStatusPending, GovernanceRuleUpdateResponseResultsBlockedByStatusDueSoon:
		return true
	}
	return false
}

type GovernanceRuleUpdateResponseResultsBlocking struct {
	ID string `json:"id" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status GovernanceRuleUpdateResponseResultsBlockingStatus `json:"status"`
	JSON   governanceRuleUpdateResponseResultsBlockingJSON   `json:"-"`
}

// governanceRuleUpdateResponseResultsBlockingJSON contains the JSON metadata for
// the struct [GovernanceRuleUpdateResponseResultsBlocking]
type governanceRuleUpdateResponseResultsBlockingJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleUpdateResponseResultsBlocking) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleUpdateResponseResultsBlockingJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceRuleUpdateResponseResultsBlockingStatus string

const (
	GovernanceRuleUpdateResponseResultsBlockingStatusRunning GovernanceRuleUpdateResponseResultsBlockingStatus = "running"
	GovernanceRuleUpdateResponseResultsBlockingStatusPassing GovernanceRuleUpdateResponseResultsBlockingStatus = "passing"
	GovernanceRuleUpdateResponseResultsBlockingStatusFailing GovernanceRuleUpdateResponseResultsBlockingStatus = "failing"
	GovernanceRuleUpdateResponseResultsBlockingStatusSkipped GovernanceRuleUpdateResponseResultsBlockingStatus = "skipped"
	GovernanceRuleUpdateResponseResultsBlockingStatusError   GovernanceRuleUpdateResponseResultsBlockingStatus = "error"
	GovernanceRuleUpdateResponseResultsBlockingStatusPending GovernanceRuleUpdateResponseResultsBlockingStatus = "pending"
	GovernanceRuleUpdateResponseResultsBlockingStatusDueSoon GovernanceRuleUpdateResponseResultsBlockingStatus = "due_soon"
)

func (r GovernanceRuleUpdateResponseResultsBlockingStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleUpdateResponseResultsBlockingStatusRunning, GovernanceRuleUpdateResponseResultsBlockingStatusPassing, GovernanceRuleUpdateResponseResultsBlockingStatusFailing, GovernanceRuleUpdateResponseResultsBlockingStatusSkipped, GovernanceRuleUpdateResponseResultsBlockingStatusError, GovernanceRuleUpdateResponseResultsBlockingStatusPending, GovernanceRuleUpdateResponseResultsBlockingStatusDueSoon:
		return true
	}
	return false
}

// Pass-rate counts across all of the rule's entities, independent of any status
// filter applied to the request.
type GovernanceRuleUpdateResponseResultsSummary struct {
	Passing int64                                          `json:"passing"`
	Total   int64                                          `json:"total"`
	JSON    governanceRuleUpdateResponseResultsSummaryJSON `json:"-"`
}

// governanceRuleUpdateResponseResultsSummaryJSON contains the JSON metadata for
// the struct [GovernanceRuleUpdateResponseResultsSummary]
type governanceRuleUpdateResponseResultsSummaryJSON struct {
	Passing     apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleUpdateResponseResultsSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleUpdateResponseResultsSummaryJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleUpdateResponseTag struct {
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
	Color string                              `json:"color" api:"nullable"`
	JSON  governanceRuleUpdateResponseTagJSON `json:"-"`
}

// governanceRuleUpdateResponseTagJSON contains the JSON metadata for the struct
// [GovernanceRuleUpdateResponseTag]
type governanceRuleUpdateResponseTagJSON struct {
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

func (r *GovernanceRuleUpdateResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleUpdateResponseTagJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleListResponse struct {
	Items []GovernanceRuleListResponseItem `json:"items" api:"required"`
	JSON  governanceRuleListResponseJSON   `json:"-"`
}

// governanceRuleListResponseJSON contains the JSON metadata for the struct
// [GovernanceRuleListResponse]
type governanceRuleListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleListResponseJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleListResponseItem struct {
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
	Scope GovernanceRuleListResponseItemsScope `json:"scope" api:"required"`
	// `platform` rules are evaluated automatically from the state of your Openlayer
	// workspace. `evidence` rules are satisfied by attaching evidence.
	Type GovernanceRuleListResponseItemsType `json:"type" api:"required"`
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
	EvidenceType GovernanceRuleListResponseItemsEvidenceType `json:"evidenceType" api:"nullable"`
	// The frameworks that include this rule.
	Frameworks []GovernanceRuleListResponseItemsFramework `json:"frameworks"`
	// Whether the rule is managed by Openlayer and cannot be edited.
	Immutable bool `json:"immutable"`
	// How often evidence must be renewed, in days. Once evidence is older than this,
	// the rule result becomes `due_soon` and then `failing`.
	RenewalCadenceDays int64 `json:"renewalCadenceDays" api:"nullable"`
	// The rule's results, one per entity the rule is evaluated against. Only returned
	// when `includeResults` is `true`.
	Results []GovernanceRuleListResponseItemsResult `json:"results"`
	// Pass-rate counts across all of the rule's entities, independent of any status
	// filter applied to the request.
	ResultsSummary GovernanceRuleListResponseItemsResultsSummary `json:"resultsSummary" api:"nullable"`
	// The rule tags associated with the rule.
	Tags []GovernanceRuleListResponseItemsTag `json:"tags" api:"nullable"`
	JSON governanceRuleListResponseItemJSON   `json:"-"`
}

// governanceRuleListResponseItemJSON contains the JSON metadata for the struct
// [GovernanceRuleListResponseItem]
type governanceRuleListResponseItemJSON struct {
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

func (r *GovernanceRuleListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleListResponseItemJSON) RawJSON() string {
	return r.raw
}

// Whether the rule is evaluated once for the whole workspace, or once per project
// the rule's frameworks apply to.
type GovernanceRuleListResponseItemsScope string

const (
	GovernanceRuleListResponseItemsScopeProject   GovernanceRuleListResponseItemsScope = "project"
	GovernanceRuleListResponseItemsScopeWorkspace GovernanceRuleListResponseItemsScope = "workspace"
)

func (r GovernanceRuleListResponseItemsScope) IsKnown() bool {
	switch r {
	case GovernanceRuleListResponseItemsScopeProject, GovernanceRuleListResponseItemsScopeWorkspace:
		return true
	}
	return false
}

// `platform` rules are evaluated automatically from the state of your Openlayer
// workspace. `evidence` rules are satisfied by attaching evidence.
type GovernanceRuleListResponseItemsType string

const (
	GovernanceRuleListResponseItemsTypePlatform GovernanceRuleListResponseItemsType = "platform"
	GovernanceRuleListResponseItemsTypeEvidence GovernanceRuleListResponseItemsType = "evidence"
)

func (r GovernanceRuleListResponseItemsType) IsKnown() bool {
	switch r {
	case GovernanceRuleListResponseItemsTypePlatform, GovernanceRuleListResponseItemsTypeEvidence:
		return true
	}
	return false
}

// The kind of evidence that satisfies the rule. `null` for platform rules.
type GovernanceRuleListResponseItemsEvidenceType string

const (
	GovernanceRuleListResponseItemsEvidenceTypeDocument      GovernanceRuleListResponseItemsEvidenceType = "document"
	GovernanceRuleListResponseItemsEvidenceTypeText          GovernanceRuleListResponseItemsEvidenceType = "text"
	GovernanceRuleListResponseItemsEvidenceTypeURL           GovernanceRuleListResponseItemsEvidenceType = "url"
	GovernanceRuleListResponseItemsEvidenceTypeCategoryValue GovernanceRuleListResponseItemsEvidenceType = "categoryValue"
)

func (r GovernanceRuleListResponseItemsEvidenceType) IsKnown() bool {
	switch r {
	case GovernanceRuleListResponseItemsEvidenceTypeDocument, GovernanceRuleListResponseItemsEvidenceTypeText, GovernanceRuleListResponseItemsEvidenceTypeURL, GovernanceRuleListResponseItemsEvidenceTypeCategoryValue:
		return true
	}
	return false
}

type GovernanceRuleListResponseItemsFramework struct {
	// The framework id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The icon shown for the framework.
	Avatar GovernanceRuleListResponseItemsFrameworksAvatar `json:"avatar" api:"required,nullable"`
	// Identifies a framework that ships with Openlayer, for example `eu_ai_act`,
	// `iso_42001`, `nist_ai_rmf`, or `traiga`. `null` for frameworks you create
	// yourself.
	BuiltInSlug string `json:"builtInSlug" api:"required,nullable"`
	// Whether the framework is active. Rules of a disabled framework are not evaluated
	// and do not count towards compliance.
	Enabled bool `json:"enabled" api:"required"`
	// The framework name.
	Name string                                       `json:"name" api:"required"`
	JSON governanceRuleListResponseItemsFrameworkJSON `json:"-"`
}

// governanceRuleListResponseItemsFrameworkJSON contains the JSON metadata for the
// struct [GovernanceRuleListResponseItemsFramework]
type governanceRuleListResponseItemsFrameworkJSON struct {
	ID          apijson.Field
	Avatar      apijson.Field
	BuiltInSlug apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleListResponseItemsFramework) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleListResponseItemsFrameworkJSON) RawJSON() string {
	return r.raw
}

// The icon shown for the framework.
type GovernanceRuleListResponseItemsFrameworksAvatar struct {
	Type  GovernanceRuleListResponseItemsFrameworksAvatarType `json:"type" api:"required"`
	Value string                                              `json:"value" api:"required"`
	JSON  governanceRuleListResponseItemsFrameworksAvatarJSON `json:"-"`
}

// governanceRuleListResponseItemsFrameworksAvatarJSON contains the JSON metadata
// for the struct [GovernanceRuleListResponseItemsFrameworksAvatar]
type governanceRuleListResponseItemsFrameworksAvatarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleListResponseItemsFrameworksAvatar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleListResponseItemsFrameworksAvatarJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleListResponseItemsFrameworksAvatarType string

const (
	GovernanceRuleListResponseItemsFrameworksAvatarTypeEmoji        GovernanceRuleListResponseItemsFrameworksAvatarType = "emoji"
	GovernanceRuleListResponseItemsFrameworksAvatarTypeImageURL     GovernanceRuleListResponseItemsFrameworksAvatarType = "imageUrl"
	GovernanceRuleListResponseItemsFrameworksAvatarTypeBuiltinImage GovernanceRuleListResponseItemsFrameworksAvatarType = "builtinImage"
)

func (r GovernanceRuleListResponseItemsFrameworksAvatarType) IsKnown() bool {
	switch r {
	case GovernanceRuleListResponseItemsFrameworksAvatarTypeEmoji, GovernanceRuleListResponseItemsFrameworksAvatarTypeImageURL, GovernanceRuleListResponseItemsFrameworksAvatarTypeBuiltinImage:
		return true
	}
	return false
}

type GovernanceRuleListResponseItemsResult struct {
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
	Status GovernanceRuleListResponseItemsResultsStatus `json:"status" api:"required"`
	// The id of the workspace the rule result belongs to.
	WorkspaceID string `json:"workspaceId" api:"required" format:"uuid"`
	// The user responsible for this result.
	AssigneeID string `json:"assigneeId" api:"nullable" format:"uuid"`
	// Rule results that must pass before this one can be satisfied.
	BlockedBy []GovernanceRuleListResponseItemsResultsBlockedBy `json:"blockedBy"`
	// Rule results that this one blocks.
	Blocking []GovernanceRuleListResponseItemsResultsBlocking `json:"blocking"`
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
	StatusMessage string                                    `json:"statusMessage" api:"nullable"`
	JSON          governanceRuleListResponseItemsResultJSON `json:"-"`
}

// governanceRuleListResponseItemsResultJSON contains the JSON metadata for the
// struct [GovernanceRuleListResponseItemsResult]
type governanceRuleListResponseItemsResultJSON struct {
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

func (r *GovernanceRuleListResponseItemsResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleListResponseItemsResultJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceRuleListResponseItemsResultsStatus string

const (
	GovernanceRuleListResponseItemsResultsStatusRunning GovernanceRuleListResponseItemsResultsStatus = "running"
	GovernanceRuleListResponseItemsResultsStatusPassing GovernanceRuleListResponseItemsResultsStatus = "passing"
	GovernanceRuleListResponseItemsResultsStatusFailing GovernanceRuleListResponseItemsResultsStatus = "failing"
	GovernanceRuleListResponseItemsResultsStatusSkipped GovernanceRuleListResponseItemsResultsStatus = "skipped"
	GovernanceRuleListResponseItemsResultsStatusError   GovernanceRuleListResponseItemsResultsStatus = "error"
	GovernanceRuleListResponseItemsResultsStatusPending GovernanceRuleListResponseItemsResultsStatus = "pending"
	GovernanceRuleListResponseItemsResultsStatusDueSoon GovernanceRuleListResponseItemsResultsStatus = "due_soon"
)

func (r GovernanceRuleListResponseItemsResultsStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleListResponseItemsResultsStatusRunning, GovernanceRuleListResponseItemsResultsStatusPassing, GovernanceRuleListResponseItemsResultsStatusFailing, GovernanceRuleListResponseItemsResultsStatusSkipped, GovernanceRuleListResponseItemsResultsStatusError, GovernanceRuleListResponseItemsResultsStatusPending, GovernanceRuleListResponseItemsResultsStatusDueSoon:
		return true
	}
	return false
}

type GovernanceRuleListResponseItemsResultsBlockedBy struct {
	ID string `json:"id" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status GovernanceRuleListResponseItemsResultsBlockedByStatus `json:"status"`
	JSON   governanceRuleListResponseItemsResultsBlockedByJSON   `json:"-"`
}

// governanceRuleListResponseItemsResultsBlockedByJSON contains the JSON metadata
// for the struct [GovernanceRuleListResponseItemsResultsBlockedBy]
type governanceRuleListResponseItemsResultsBlockedByJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleListResponseItemsResultsBlockedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleListResponseItemsResultsBlockedByJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceRuleListResponseItemsResultsBlockedByStatus string

const (
	GovernanceRuleListResponseItemsResultsBlockedByStatusRunning GovernanceRuleListResponseItemsResultsBlockedByStatus = "running"
	GovernanceRuleListResponseItemsResultsBlockedByStatusPassing GovernanceRuleListResponseItemsResultsBlockedByStatus = "passing"
	GovernanceRuleListResponseItemsResultsBlockedByStatusFailing GovernanceRuleListResponseItemsResultsBlockedByStatus = "failing"
	GovernanceRuleListResponseItemsResultsBlockedByStatusSkipped GovernanceRuleListResponseItemsResultsBlockedByStatus = "skipped"
	GovernanceRuleListResponseItemsResultsBlockedByStatusError   GovernanceRuleListResponseItemsResultsBlockedByStatus = "error"
	GovernanceRuleListResponseItemsResultsBlockedByStatusPending GovernanceRuleListResponseItemsResultsBlockedByStatus = "pending"
	GovernanceRuleListResponseItemsResultsBlockedByStatusDueSoon GovernanceRuleListResponseItemsResultsBlockedByStatus = "due_soon"
)

func (r GovernanceRuleListResponseItemsResultsBlockedByStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleListResponseItemsResultsBlockedByStatusRunning, GovernanceRuleListResponseItemsResultsBlockedByStatusPassing, GovernanceRuleListResponseItemsResultsBlockedByStatusFailing, GovernanceRuleListResponseItemsResultsBlockedByStatusSkipped, GovernanceRuleListResponseItemsResultsBlockedByStatusError, GovernanceRuleListResponseItemsResultsBlockedByStatusPending, GovernanceRuleListResponseItemsResultsBlockedByStatusDueSoon:
		return true
	}
	return false
}

type GovernanceRuleListResponseItemsResultsBlocking struct {
	ID string `json:"id" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status GovernanceRuleListResponseItemsResultsBlockingStatus `json:"status"`
	JSON   governanceRuleListResponseItemsResultsBlockingJSON   `json:"-"`
}

// governanceRuleListResponseItemsResultsBlockingJSON contains the JSON metadata
// for the struct [GovernanceRuleListResponseItemsResultsBlocking]
type governanceRuleListResponseItemsResultsBlockingJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleListResponseItemsResultsBlocking) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleListResponseItemsResultsBlockingJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceRuleListResponseItemsResultsBlockingStatus string

const (
	GovernanceRuleListResponseItemsResultsBlockingStatusRunning GovernanceRuleListResponseItemsResultsBlockingStatus = "running"
	GovernanceRuleListResponseItemsResultsBlockingStatusPassing GovernanceRuleListResponseItemsResultsBlockingStatus = "passing"
	GovernanceRuleListResponseItemsResultsBlockingStatusFailing GovernanceRuleListResponseItemsResultsBlockingStatus = "failing"
	GovernanceRuleListResponseItemsResultsBlockingStatusSkipped GovernanceRuleListResponseItemsResultsBlockingStatus = "skipped"
	GovernanceRuleListResponseItemsResultsBlockingStatusError   GovernanceRuleListResponseItemsResultsBlockingStatus = "error"
	GovernanceRuleListResponseItemsResultsBlockingStatusPending GovernanceRuleListResponseItemsResultsBlockingStatus = "pending"
	GovernanceRuleListResponseItemsResultsBlockingStatusDueSoon GovernanceRuleListResponseItemsResultsBlockingStatus = "due_soon"
)

func (r GovernanceRuleListResponseItemsResultsBlockingStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleListResponseItemsResultsBlockingStatusRunning, GovernanceRuleListResponseItemsResultsBlockingStatusPassing, GovernanceRuleListResponseItemsResultsBlockingStatusFailing, GovernanceRuleListResponseItemsResultsBlockingStatusSkipped, GovernanceRuleListResponseItemsResultsBlockingStatusError, GovernanceRuleListResponseItemsResultsBlockingStatusPending, GovernanceRuleListResponseItemsResultsBlockingStatusDueSoon:
		return true
	}
	return false
}

// Pass-rate counts across all of the rule's entities, independent of any status
// filter applied to the request.
type GovernanceRuleListResponseItemsResultsSummary struct {
	Passing int64                                             `json:"passing"`
	Total   int64                                             `json:"total"`
	JSON    governanceRuleListResponseItemsResultsSummaryJSON `json:"-"`
}

// governanceRuleListResponseItemsResultsSummaryJSON contains the JSON metadata for
// the struct [GovernanceRuleListResponseItemsResultsSummary]
type governanceRuleListResponseItemsResultsSummaryJSON struct {
	Passing     apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleListResponseItemsResultsSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleListResponseItemsResultsSummaryJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleListResponseItemsTag struct {
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
	Color string                                 `json:"color" api:"nullable"`
	JSON  governanceRuleListResponseItemsTagJSON `json:"-"`
}

// governanceRuleListResponseItemsTagJSON contains the JSON metadata for the struct
// [GovernanceRuleListResponseItemsTag]
type governanceRuleListResponseItemsTagJSON struct {
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

func (r *GovernanceRuleListResponseItemsTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleListResponseItemsTagJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleNewParams struct {
	// The rule name.
	Name param.Field[string] `json:"name" api:"required"`
	// Whether the rule is evaluated once for the whole workspace, or once per project
	// the rule's frameworks apply to.
	Scope param.Field[GovernanceRuleNewParamsScope] `json:"scope" api:"required"`
	// `platform` rules are evaluated automatically from the state of your Openlayer
	// workspace. `evidence` rules are satisfied by attaching evidence.
	Type param.Field[GovernanceRuleNewParamsType] `json:"type" api:"required"`
	// The user responsible for satisfying the rule.
	AssigneeID param.Field[string] `json:"assigneeId" format:"uuid"`
	// Configuration for the platform check, when the automation takes parameters.
	AutomationParams param.Field[map[string]interface{}] `json:"automationParams"`
	// Which workspace signal a platform rule checks, for example
	// `monitoring_mode_enabled`, `test_setup`, or `project_owner_set`. `null` for
	// evidence rules.
	AutomationType param.Field[string] `json:"automationType"`
	// Whether the rule is excluded from compliance calculations.
	Deactivated param.Field[bool] `json:"deactivated"`
	// What the rule requires.
	Description param.Field[string] `json:"description"`
	// The kind of evidence that satisfies the rule. `null` for platform rules.
	EvidenceType param.Field[GovernanceRuleNewParamsEvidenceType] `json:"evidenceType"`
	// How often evidence must be renewed, in days. Once evidence is older than this,
	// the rule result becomes `due_soon` and then `failing`.
	RenewalCadenceDays param.Field[int64] `json:"renewalCadenceDays"`
	// The ids of the rule tags to associate with the rule. Replaces the rule's tags.
	// Read them back from `tags`, and list the tags available in the workspace with
	// `GET /workspaces/{workspaceId}/rule-tags`.
	TagIDs param.Field[[]string] `json:"tagIds" format:"uuid"`
}

func (r GovernanceRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the rule is evaluated once for the whole workspace, or once per project
// the rule's frameworks apply to.
type GovernanceRuleNewParamsScope string

const (
	GovernanceRuleNewParamsScopeProject   GovernanceRuleNewParamsScope = "project"
	GovernanceRuleNewParamsScopeWorkspace GovernanceRuleNewParamsScope = "workspace"
)

func (r GovernanceRuleNewParamsScope) IsKnown() bool {
	switch r {
	case GovernanceRuleNewParamsScopeProject, GovernanceRuleNewParamsScopeWorkspace:
		return true
	}
	return false
}

// `platform` rules are evaluated automatically from the state of your Openlayer
// workspace. `evidence` rules are satisfied by attaching evidence.
type GovernanceRuleNewParamsType string

const (
	GovernanceRuleNewParamsTypePlatform GovernanceRuleNewParamsType = "platform"
	GovernanceRuleNewParamsTypeEvidence GovernanceRuleNewParamsType = "evidence"
)

func (r GovernanceRuleNewParamsType) IsKnown() bool {
	switch r {
	case GovernanceRuleNewParamsTypePlatform, GovernanceRuleNewParamsTypeEvidence:
		return true
	}
	return false
}

// The kind of evidence that satisfies the rule. `null` for platform rules.
type GovernanceRuleNewParamsEvidenceType string

const (
	GovernanceRuleNewParamsEvidenceTypeDocument      GovernanceRuleNewParamsEvidenceType = "document"
	GovernanceRuleNewParamsEvidenceTypeText          GovernanceRuleNewParamsEvidenceType = "text"
	GovernanceRuleNewParamsEvidenceTypeURL           GovernanceRuleNewParamsEvidenceType = "url"
	GovernanceRuleNewParamsEvidenceTypeCategoryValue GovernanceRuleNewParamsEvidenceType = "categoryValue"
)

func (r GovernanceRuleNewParamsEvidenceType) IsKnown() bool {
	switch r {
	case GovernanceRuleNewParamsEvidenceTypeDocument, GovernanceRuleNewParamsEvidenceTypeText, GovernanceRuleNewParamsEvidenceTypeURL, GovernanceRuleNewParamsEvidenceTypeCategoryValue:
		return true
	}
	return false
}

type GovernanceRuleUpdateParams struct {
	// The user responsible for satisfying the rule.
	AssigneeID param.Field[string] `json:"assigneeId" format:"uuid"`
	// Whether the rule is excluded from compliance calculations.
	Deactivated param.Field[bool] `json:"deactivated"`
	// What the rule requires.
	Description param.Field[string] `json:"description"`
	// The rule name.
	Name param.Field[string] `json:"name"`
	// How often evidence must be renewed, in days. Once evidence is older than this,
	// the rule result becomes `due_soon` and then `failing`.
	RenewalCadenceDays param.Field[int64] `json:"renewalCadenceDays"`
	// The ids of the rule tags to associate with the rule. Replaces the rule's tags.
	// Read them back from `tags`, and list the tags available in the workspace with
	// `GET /workspaces/{workspaceId}/rule-tags`.
	TagIDs param.Field[[]string] `json:"tagIds" format:"uuid"`
}

func (r GovernanceRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GovernanceRuleListParams struct {
	// Whether to sort in ascending order.
	Asc param.Field[bool] `query:"asc"`
	// Only include rules assigned to this user.
	AssigneeID param.Field[string] `query:"assigneeId" format:"uuid"`
	// Only include rules that are deactivated (or active).
	Deactivated param.Field[bool] `query:"deactivated"`
	// Only include items belonging to at least one enabled framework.
	EnabledFrameworkOnly param.Field[bool] `query:"enabledFrameworkOnly"`
	// Only include items belonging to this framework.
	FrameworkID param.Field[string] `query:"frameworkId" format:"uuid"`
	// Only include rules in one bucket of the compliance workflow. `open` covers rules
	// that still need attention, `done` covers rules that are fully satisfied, and
	// `excluded` covers rules that have been deactivated.
	Group param.Field[GovernanceRuleListParamsGroup] `query:"group"`
	// Whether to include each rule's results inline, in a `results` array.
	IncludeResults param.Field[bool] `query:"includeResults"`
	// Whether to include rules that are not part of any framework.
	IncludeUnframed param.Field[bool] `query:"includeUnframed"`
	// The page to return in a paginated query.
	Page param.Field[int64] `query:"page"`
	// Maximum number of items to return per page.
	PerPage param.Field[int64] `query:"perPage"`
	// Only include items that apply to this project.
	ProjectID param.Field[string] `query:"projectId" format:"uuid"`
	// Only include rules with this scope.
	Scope param.Field[GovernanceRuleListParamsScope] `query:"scope"`
	// Filter by a free-text search over names and descriptions.
	SearchQuery param.Field[string] `query:"searchQuery"`
	// The field to sort on.
	SortBy param.Field[GovernanceRuleListParamsSortBy] `query:"sortBy"`
	// Only include items whose rule result has this compliance status.
	Status param.Field[GovernanceRuleListParamsStatus] `query:"status"`
	// Only include rules carrying all of these rule tags. Pass tag ids, which you can
	// look up with [List rule tags](/api-reference/rest/governance/list-rule-tags).
	Tags param.Field[[]string] `query:"tags" format:"uuid"`
	// Only include rules of this type.
	Type param.Field[GovernanceRuleListParamsType] `query:"type"`
}

// URLQuery serializes [GovernanceRuleListParams]'s query parameters as
// `url.Values`.
func (r GovernanceRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only include rules in one bucket of the compliance workflow. `open` covers rules
// that still need attention, `done` covers rules that are fully satisfied, and
// `excluded` covers rules that have been deactivated.
type GovernanceRuleListParamsGroup string

const (
	GovernanceRuleListParamsGroupOpen     GovernanceRuleListParamsGroup = "open"
	GovernanceRuleListParamsGroupExcluded GovernanceRuleListParamsGroup = "excluded"
	GovernanceRuleListParamsGroupDone     GovernanceRuleListParamsGroup = "done"
)

func (r GovernanceRuleListParamsGroup) IsKnown() bool {
	switch r {
	case GovernanceRuleListParamsGroupOpen, GovernanceRuleListParamsGroupExcluded, GovernanceRuleListParamsGroupDone:
		return true
	}
	return false
}

// Only include rules with this scope.
type GovernanceRuleListParamsScope string

const (
	GovernanceRuleListParamsScopeProject   GovernanceRuleListParamsScope = "project"
	GovernanceRuleListParamsScopeWorkspace GovernanceRuleListParamsScope = "workspace"
)

func (r GovernanceRuleListParamsScope) IsKnown() bool {
	switch r {
	case GovernanceRuleListParamsScopeProject, GovernanceRuleListParamsScopeWorkspace:
		return true
	}
	return false
}

// The field to sort on.
type GovernanceRuleListParamsSortBy string

const (
	GovernanceRuleListParamsSortByName        GovernanceRuleListParamsSortBy = "name"
	GovernanceRuleListParamsSortByStatus      GovernanceRuleListParamsSortBy = "status"
	GovernanceRuleListParamsSortByFrameworks  GovernanceRuleListParamsSortBy = "frameworks"
	GovernanceRuleListParamsSortByScope       GovernanceRuleListParamsSortBy = "scope"
	GovernanceRuleListParamsSortByDateCreated GovernanceRuleListParamsSortBy = "dateCreated"
)

func (r GovernanceRuleListParamsSortBy) IsKnown() bool {
	switch r {
	case GovernanceRuleListParamsSortByName, GovernanceRuleListParamsSortByStatus, GovernanceRuleListParamsSortByFrameworks, GovernanceRuleListParamsSortByScope, GovernanceRuleListParamsSortByDateCreated:
		return true
	}
	return false
}

// Only include items whose rule result has this compliance status.
type GovernanceRuleListParamsStatus string

const (
	GovernanceRuleListParamsStatusRunning GovernanceRuleListParamsStatus = "running"
	GovernanceRuleListParamsStatusPassing GovernanceRuleListParamsStatus = "passing"
	GovernanceRuleListParamsStatusFailing GovernanceRuleListParamsStatus = "failing"
	GovernanceRuleListParamsStatusSkipped GovernanceRuleListParamsStatus = "skipped"
	GovernanceRuleListParamsStatusError   GovernanceRuleListParamsStatus = "error"
	GovernanceRuleListParamsStatusPending GovernanceRuleListParamsStatus = "pending"
	GovernanceRuleListParamsStatusDueSoon GovernanceRuleListParamsStatus = "due_soon"
)

func (r GovernanceRuleListParamsStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleListParamsStatusRunning, GovernanceRuleListParamsStatusPassing, GovernanceRuleListParamsStatusFailing, GovernanceRuleListParamsStatusSkipped, GovernanceRuleListParamsStatusError, GovernanceRuleListParamsStatusPending, GovernanceRuleListParamsStatusDueSoon:
		return true
	}
	return false
}

// Only include rules of this type.
type GovernanceRuleListParamsType string

const (
	GovernanceRuleListParamsTypePlatform GovernanceRuleListParamsType = "platform"
	GovernanceRuleListParamsTypeEvidence GovernanceRuleListParamsType = "evidence"
)

func (r GovernanceRuleListParamsType) IsKnown() bool {
	switch r {
	case GovernanceRuleListParamsTypePlatform, GovernanceRuleListParamsTypeEvidence:
		return true
	}
	return false
}
