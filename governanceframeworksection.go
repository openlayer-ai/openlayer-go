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

// GovernanceFrameworkSectionService contains methods and other services that help
// with interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGovernanceFrameworkSectionService] method instead.
type GovernanceFrameworkSectionService struct {
	Options []option.RequestOption
}

// NewGovernanceFrameworkSectionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewGovernanceFrameworkSectionService(opts ...option.RequestOption) (r *GovernanceFrameworkSectionService) {
	r = &GovernanceFrameworkSectionService{}
	r.Options = opts
	return
}

// List the rules mapped to a section of a framework document.
//
// Pass `includeSubsectionRules=true` to also return the rules mapped to the
// section's subsections, which is how you get every rule covering a requirement
// and everything under it.
func (r *GovernanceFrameworkSectionService) ListRules(ctx context.Context, frameworkID string, sectionID string, query GovernanceFrameworkSectionListRulesParams, opts ...option.RequestOption) (res *GovernanceFrameworkSectionListRulesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if frameworkID == "" {
		err = errors.New("missing required frameworkId parameter")
		return nil, err
	}
	if sectionID == "" {
		err = errors.New("missing required sectionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("frameworks/%s/sections/%s/rules", frameworkID, sectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type GovernanceFrameworkSectionListRulesResponse struct {
	Items []GovernanceFrameworkSectionListRulesResponseItem `json:"items" api:"required"`
	JSON  governanceFrameworkSectionListRulesResponseJSON   `json:"-"`
}

// governanceFrameworkSectionListRulesResponseJSON contains the JSON metadata for
// the struct [GovernanceFrameworkSectionListRulesResponse]
type governanceFrameworkSectionListRulesResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkSectionListRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkSectionListRulesResponseJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkSectionListRulesResponseItem struct {
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
	Scope GovernanceFrameworkSectionListRulesResponseItemsScope `json:"scope" api:"required"`
	// `platform` rules are evaluated automatically from the state of your Openlayer
	// workspace. `evidence` rules are satisfied by attaching evidence.
	Type GovernanceFrameworkSectionListRulesResponseItemsType `json:"type" api:"required"`
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
	EvidenceType GovernanceFrameworkSectionListRulesResponseItemsEvidenceType `json:"evidenceType" api:"nullable"`
	// The frameworks that include this rule.
	Frameworks []GovernanceFrameworkSectionListRulesResponseItemsFramework `json:"frameworks"`
	// Whether the rule is managed by Openlayer and cannot be edited.
	Immutable bool `json:"immutable"`
	// How often evidence must be renewed, in days. Once evidence is older than this,
	// the rule result becomes `due_soon` and then `failing`.
	RenewalCadenceDays int64 `json:"renewalCadenceDays" api:"nullable"`
	// The rule's results, one per entity the rule is evaluated against. Only returned
	// when `includeResults` is `true`.
	Results []GovernanceFrameworkSectionListRulesResponseItemsResult `json:"results"`
	// Pass-rate counts across all of the rule's entities, independent of any status
	// filter applied to the request.
	ResultsSummary GovernanceFrameworkSectionListRulesResponseItemsResultsSummary `json:"resultsSummary" api:"nullable"`
	// The rule tags associated with the rule.
	Tags []GovernanceFrameworkSectionListRulesResponseItemsTag `json:"tags" api:"nullable"`
	JSON governanceFrameworkSectionListRulesResponseItemJSON   `json:"-"`
}

// governanceFrameworkSectionListRulesResponseItemJSON contains the JSON metadata
// for the struct [GovernanceFrameworkSectionListRulesResponseItem]
type governanceFrameworkSectionListRulesResponseItemJSON struct {
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

func (r *GovernanceFrameworkSectionListRulesResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkSectionListRulesResponseItemJSON) RawJSON() string {
	return r.raw
}

// Whether the rule is evaluated once for the whole workspace, or once per project
// the rule's frameworks apply to.
type GovernanceFrameworkSectionListRulesResponseItemsScope string

const (
	GovernanceFrameworkSectionListRulesResponseItemsScopeProject   GovernanceFrameworkSectionListRulesResponseItemsScope = "project"
	GovernanceFrameworkSectionListRulesResponseItemsScopeWorkspace GovernanceFrameworkSectionListRulesResponseItemsScope = "workspace"
)

func (r GovernanceFrameworkSectionListRulesResponseItemsScope) IsKnown() bool {
	switch r {
	case GovernanceFrameworkSectionListRulesResponseItemsScopeProject, GovernanceFrameworkSectionListRulesResponseItemsScopeWorkspace:
		return true
	}
	return false
}

// `platform` rules are evaluated automatically from the state of your Openlayer
// workspace. `evidence` rules are satisfied by attaching evidence.
type GovernanceFrameworkSectionListRulesResponseItemsType string

const (
	GovernanceFrameworkSectionListRulesResponseItemsTypePlatform GovernanceFrameworkSectionListRulesResponseItemsType = "platform"
	GovernanceFrameworkSectionListRulesResponseItemsTypeEvidence GovernanceFrameworkSectionListRulesResponseItemsType = "evidence"
)

func (r GovernanceFrameworkSectionListRulesResponseItemsType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkSectionListRulesResponseItemsTypePlatform, GovernanceFrameworkSectionListRulesResponseItemsTypeEvidence:
		return true
	}
	return false
}

// The kind of evidence that satisfies the rule. `null` for platform rules.
type GovernanceFrameworkSectionListRulesResponseItemsEvidenceType string

const (
	GovernanceFrameworkSectionListRulesResponseItemsEvidenceTypeDocument      GovernanceFrameworkSectionListRulesResponseItemsEvidenceType = "document"
	GovernanceFrameworkSectionListRulesResponseItemsEvidenceTypeText          GovernanceFrameworkSectionListRulesResponseItemsEvidenceType = "text"
	GovernanceFrameworkSectionListRulesResponseItemsEvidenceTypeURL           GovernanceFrameworkSectionListRulesResponseItemsEvidenceType = "url"
	GovernanceFrameworkSectionListRulesResponseItemsEvidenceTypeCategoryValue GovernanceFrameworkSectionListRulesResponseItemsEvidenceType = "categoryValue"
)

func (r GovernanceFrameworkSectionListRulesResponseItemsEvidenceType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkSectionListRulesResponseItemsEvidenceTypeDocument, GovernanceFrameworkSectionListRulesResponseItemsEvidenceTypeText, GovernanceFrameworkSectionListRulesResponseItemsEvidenceTypeURL, GovernanceFrameworkSectionListRulesResponseItemsEvidenceTypeCategoryValue:
		return true
	}
	return false
}

type GovernanceFrameworkSectionListRulesResponseItemsFramework struct {
	// The framework id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The icon shown for the framework.
	Avatar GovernanceFrameworkSectionListRulesResponseItemsFrameworksAvatar `json:"avatar" api:"required,nullable"`
	// Identifies a framework that ships with Openlayer, for example `eu_ai_act`,
	// `iso_42001`, `nist_ai_rmf`, or `traiga`. `null` for frameworks you create
	// yourself.
	BuiltInSlug string `json:"builtInSlug" api:"required,nullable"`
	// Whether the framework is active. Rules of a disabled framework are not evaluated
	// and do not count towards compliance.
	Enabled bool `json:"enabled" api:"required"`
	// The framework name.
	Name string                                                        `json:"name" api:"required"`
	JSON governanceFrameworkSectionListRulesResponseItemsFrameworkJSON `json:"-"`
}

// governanceFrameworkSectionListRulesResponseItemsFrameworkJSON contains the JSON
// metadata for the struct
// [GovernanceFrameworkSectionListRulesResponseItemsFramework]
type governanceFrameworkSectionListRulesResponseItemsFrameworkJSON struct {
	ID          apijson.Field
	Avatar      apijson.Field
	BuiltInSlug apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkSectionListRulesResponseItemsFramework) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkSectionListRulesResponseItemsFrameworkJSON) RawJSON() string {
	return r.raw
}

// The icon shown for the framework.
type GovernanceFrameworkSectionListRulesResponseItemsFrameworksAvatar struct {
	Type  GovernanceFrameworkSectionListRulesResponseItemsFrameworksAvatarType `json:"type" api:"required"`
	Value string                                                               `json:"value" api:"required"`
	JSON  governanceFrameworkSectionListRulesResponseItemsFrameworksAvatarJSON `json:"-"`
}

// governanceFrameworkSectionListRulesResponseItemsFrameworksAvatarJSON contains
// the JSON metadata for the struct
// [GovernanceFrameworkSectionListRulesResponseItemsFrameworksAvatar]
type governanceFrameworkSectionListRulesResponseItemsFrameworksAvatarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkSectionListRulesResponseItemsFrameworksAvatar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkSectionListRulesResponseItemsFrameworksAvatarJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkSectionListRulesResponseItemsFrameworksAvatarType string

const (
	GovernanceFrameworkSectionListRulesResponseItemsFrameworksAvatarTypeEmoji        GovernanceFrameworkSectionListRulesResponseItemsFrameworksAvatarType = "emoji"
	GovernanceFrameworkSectionListRulesResponseItemsFrameworksAvatarTypeImageURL     GovernanceFrameworkSectionListRulesResponseItemsFrameworksAvatarType = "imageUrl"
	GovernanceFrameworkSectionListRulesResponseItemsFrameworksAvatarTypeBuiltinImage GovernanceFrameworkSectionListRulesResponseItemsFrameworksAvatarType = "builtinImage"
)

func (r GovernanceFrameworkSectionListRulesResponseItemsFrameworksAvatarType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkSectionListRulesResponseItemsFrameworksAvatarTypeEmoji, GovernanceFrameworkSectionListRulesResponseItemsFrameworksAvatarTypeImageURL, GovernanceFrameworkSectionListRulesResponseItemsFrameworksAvatarTypeBuiltinImage:
		return true
	}
	return false
}

type GovernanceFrameworkSectionListRulesResponseItemsResult struct {
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
	Status GovernanceFrameworkSectionListRulesResponseItemsResultsStatus `json:"status" api:"required"`
	// The id of the workspace the rule result belongs to.
	WorkspaceID string `json:"workspaceId" api:"required" format:"uuid"`
	// The user responsible for this result.
	AssigneeID string `json:"assigneeId" api:"nullable" format:"uuid"`
	// Rule results that must pass before this one can be satisfied.
	BlockedBy []GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedBy `json:"blockedBy"`
	// Rule results that this one blocks.
	Blocking []GovernanceFrameworkSectionListRulesResponseItemsResultsBlocking `json:"blocking"`
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
	StatusMessage string                                                     `json:"statusMessage" api:"nullable"`
	JSON          governanceFrameworkSectionListRulesResponseItemsResultJSON `json:"-"`
}

// governanceFrameworkSectionListRulesResponseItemsResultJSON contains the JSON
// metadata for the struct [GovernanceFrameworkSectionListRulesResponseItemsResult]
type governanceFrameworkSectionListRulesResponseItemsResultJSON struct {
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

func (r *GovernanceFrameworkSectionListRulesResponseItemsResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkSectionListRulesResponseItemsResultJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceFrameworkSectionListRulesResponseItemsResultsStatus string

const (
	GovernanceFrameworkSectionListRulesResponseItemsResultsStatusRunning GovernanceFrameworkSectionListRulesResponseItemsResultsStatus = "running"
	GovernanceFrameworkSectionListRulesResponseItemsResultsStatusPassing GovernanceFrameworkSectionListRulesResponseItemsResultsStatus = "passing"
	GovernanceFrameworkSectionListRulesResponseItemsResultsStatusFailing GovernanceFrameworkSectionListRulesResponseItemsResultsStatus = "failing"
	GovernanceFrameworkSectionListRulesResponseItemsResultsStatusSkipped GovernanceFrameworkSectionListRulesResponseItemsResultsStatus = "skipped"
	GovernanceFrameworkSectionListRulesResponseItemsResultsStatusError   GovernanceFrameworkSectionListRulesResponseItemsResultsStatus = "error"
	GovernanceFrameworkSectionListRulesResponseItemsResultsStatusPending GovernanceFrameworkSectionListRulesResponseItemsResultsStatus = "pending"
	GovernanceFrameworkSectionListRulesResponseItemsResultsStatusDueSoon GovernanceFrameworkSectionListRulesResponseItemsResultsStatus = "due_soon"
)

func (r GovernanceFrameworkSectionListRulesResponseItemsResultsStatus) IsKnown() bool {
	switch r {
	case GovernanceFrameworkSectionListRulesResponseItemsResultsStatusRunning, GovernanceFrameworkSectionListRulesResponseItemsResultsStatusPassing, GovernanceFrameworkSectionListRulesResponseItemsResultsStatusFailing, GovernanceFrameworkSectionListRulesResponseItemsResultsStatusSkipped, GovernanceFrameworkSectionListRulesResponseItemsResultsStatusError, GovernanceFrameworkSectionListRulesResponseItemsResultsStatusPending, GovernanceFrameworkSectionListRulesResponseItemsResultsStatusDueSoon:
		return true
	}
	return false
}

type GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedBy struct {
	ID string `json:"id" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatus `json:"status"`
	JSON   governanceFrameworkSectionListRulesResponseItemsResultsBlockedByJSON   `json:"-"`
}

// governanceFrameworkSectionListRulesResponseItemsResultsBlockedByJSON contains
// the JSON metadata for the struct
// [GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedBy]
type governanceFrameworkSectionListRulesResponseItemsResultsBlockedByJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkSectionListRulesResponseItemsResultsBlockedByJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatus string

const (
	GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatusRunning GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatus = "running"
	GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatusPassing GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatus = "passing"
	GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatusFailing GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatus = "failing"
	GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatusSkipped GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatus = "skipped"
	GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatusError   GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatus = "error"
	GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatusPending GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatus = "pending"
	GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatusDueSoon GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatus = "due_soon"
)

func (r GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatus) IsKnown() bool {
	switch r {
	case GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatusRunning, GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatusPassing, GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatusFailing, GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatusSkipped, GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatusError, GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatusPending, GovernanceFrameworkSectionListRulesResponseItemsResultsBlockedByStatusDueSoon:
		return true
	}
	return false
}

type GovernanceFrameworkSectionListRulesResponseItemsResultsBlocking struct {
	ID string `json:"id" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatus `json:"status"`
	JSON   governanceFrameworkSectionListRulesResponseItemsResultsBlockingJSON   `json:"-"`
}

// governanceFrameworkSectionListRulesResponseItemsResultsBlockingJSON contains the
// JSON metadata for the struct
// [GovernanceFrameworkSectionListRulesResponseItemsResultsBlocking]
type governanceFrameworkSectionListRulesResponseItemsResultsBlockingJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkSectionListRulesResponseItemsResultsBlocking) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkSectionListRulesResponseItemsResultsBlockingJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatus string

const (
	GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatusRunning GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatus = "running"
	GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatusPassing GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatus = "passing"
	GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatusFailing GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatus = "failing"
	GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatusSkipped GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatus = "skipped"
	GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatusError   GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatus = "error"
	GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatusPending GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatus = "pending"
	GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatusDueSoon GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatus = "due_soon"
)

func (r GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatus) IsKnown() bool {
	switch r {
	case GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatusRunning, GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatusPassing, GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatusFailing, GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatusSkipped, GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatusError, GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatusPending, GovernanceFrameworkSectionListRulesResponseItemsResultsBlockingStatusDueSoon:
		return true
	}
	return false
}

// Pass-rate counts across all of the rule's entities, independent of any status
// filter applied to the request.
type GovernanceFrameworkSectionListRulesResponseItemsResultsSummary struct {
	Passing int64                                                              `json:"passing"`
	Total   int64                                                              `json:"total"`
	JSON    governanceFrameworkSectionListRulesResponseItemsResultsSummaryJSON `json:"-"`
}

// governanceFrameworkSectionListRulesResponseItemsResultsSummaryJSON contains the
// JSON metadata for the struct
// [GovernanceFrameworkSectionListRulesResponseItemsResultsSummary]
type governanceFrameworkSectionListRulesResponseItemsResultsSummaryJSON struct {
	Passing     apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkSectionListRulesResponseItemsResultsSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkSectionListRulesResponseItemsResultsSummaryJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkSectionListRulesResponseItemsTag struct {
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
	Color string                                                  `json:"color" api:"nullable"`
	JSON  governanceFrameworkSectionListRulesResponseItemsTagJSON `json:"-"`
}

// governanceFrameworkSectionListRulesResponseItemsTagJSON contains the JSON
// metadata for the struct [GovernanceFrameworkSectionListRulesResponseItemsTag]
type governanceFrameworkSectionListRulesResponseItemsTagJSON struct {
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

func (r *GovernanceFrameworkSectionListRulesResponseItemsTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkSectionListRulesResponseItemsTagJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkSectionListRulesParams struct {
	// Whether to include each rule's results inline, in a `results` array.
	IncludeResults param.Field[bool] `query:"includeResults"`
	// Whether to also include the rules mapped to the section's subsections.
	IncludeSubsectionRules param.Field[bool] `query:"includeSubsectionRules"`
	// The page to return in a paginated query.
	Page param.Field[int64] `query:"page"`
	// Maximum number of items to return per page.
	PerPage param.Field[int64] `query:"perPage"`
	// Only include items that apply to this project.
	ProjectID param.Field[string] `query:"projectId" format:"uuid"`
	// Only include items whose rule result has this compliance status.
	Status param.Field[GovernanceFrameworkSectionListRulesParamsStatus] `query:"status"`
}

// URLQuery serializes [GovernanceFrameworkSectionListRulesParams]'s query
// parameters as `url.Values`.
func (r GovernanceFrameworkSectionListRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only include items whose rule result has this compliance status.
type GovernanceFrameworkSectionListRulesParamsStatus string

const (
	GovernanceFrameworkSectionListRulesParamsStatusRunning GovernanceFrameworkSectionListRulesParamsStatus = "running"
	GovernanceFrameworkSectionListRulesParamsStatusPassing GovernanceFrameworkSectionListRulesParamsStatus = "passing"
	GovernanceFrameworkSectionListRulesParamsStatusFailing GovernanceFrameworkSectionListRulesParamsStatus = "failing"
	GovernanceFrameworkSectionListRulesParamsStatusSkipped GovernanceFrameworkSectionListRulesParamsStatus = "skipped"
	GovernanceFrameworkSectionListRulesParamsStatusError   GovernanceFrameworkSectionListRulesParamsStatus = "error"
	GovernanceFrameworkSectionListRulesParamsStatusPending GovernanceFrameworkSectionListRulesParamsStatus = "pending"
	GovernanceFrameworkSectionListRulesParamsStatusDueSoon GovernanceFrameworkSectionListRulesParamsStatus = "due_soon"
)

func (r GovernanceFrameworkSectionListRulesParamsStatus) IsKnown() bool {
	switch r {
	case GovernanceFrameworkSectionListRulesParamsStatusRunning, GovernanceFrameworkSectionListRulesParamsStatusPassing, GovernanceFrameworkSectionListRulesParamsStatusFailing, GovernanceFrameworkSectionListRulesParamsStatusSkipped, GovernanceFrameworkSectionListRulesParamsStatusError, GovernanceFrameworkSectionListRulesParamsStatusPending, GovernanceFrameworkSectionListRulesParamsStatusDueSoon:
		return true
	}
	return false
}
