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

// GovernanceFrameworkSubsectionService contains methods and other services that
// help with interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGovernanceFrameworkSubsectionService] method instead.
type GovernanceFrameworkSubsectionService struct {
	Options []option.RequestOption
}

// NewGovernanceFrameworkSubsectionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewGovernanceFrameworkSubsectionService(opts ...option.RequestOption) (r *GovernanceFrameworkSubsectionService) {
	r = &GovernanceFrameworkSubsectionService{}
	r.Options = opts
	return
}

// List the rules mapped to a subsection of a framework document.
//
// A subsection is usually the level at which a standard states an individual
// requirement, so this is the endpoint to use when you want to show which rules
// cover a specific clause.
func (r *GovernanceFrameworkSubsectionService) ListRules(ctx context.Context, frameworkID string, subsectionID string, query GovernanceFrameworkSubsectionListRulesParams, opts ...option.RequestOption) (res *GovernanceFrameworkSubsectionListRulesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if frameworkID == "" {
		err = errors.New("missing required frameworkId parameter")
		return nil, err
	}
	if subsectionID == "" {
		err = errors.New("missing required subsectionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("frameworks/%s/subsections/%s/rules", frameworkID, subsectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type GovernanceFrameworkSubsectionListRulesResponse struct {
	Items []GovernanceFrameworkSubsectionListRulesResponseItem `json:"items" api:"required"`
	JSON  governanceFrameworkSubsectionListRulesResponseJSON   `json:"-"`
}

// governanceFrameworkSubsectionListRulesResponseJSON contains the JSON metadata
// for the struct [GovernanceFrameworkSubsectionListRulesResponse]
type governanceFrameworkSubsectionListRulesResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkSubsectionListRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkSubsectionListRulesResponseJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkSubsectionListRulesResponseItem struct {
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
	Scope GovernanceFrameworkSubsectionListRulesResponseItemsScope `json:"scope" api:"required"`
	// `platform` rules are evaluated automatically from the state of your Openlayer
	// workspace. `evidence` rules are satisfied by attaching evidence.
	Type GovernanceFrameworkSubsectionListRulesResponseItemsType `json:"type" api:"required"`
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
	EvidenceType GovernanceFrameworkSubsectionListRulesResponseItemsEvidenceType `json:"evidenceType" api:"nullable"`
	// The frameworks that include this rule.
	Frameworks []GovernanceFrameworkSubsectionListRulesResponseItemsFramework `json:"frameworks"`
	// Whether the rule is managed by Openlayer and cannot be edited.
	Immutable bool `json:"immutable"`
	// How often evidence must be renewed, in days. Once evidence is older than this,
	// the rule result becomes `due_soon` and then `failing`.
	RenewalCadenceDays int64 `json:"renewalCadenceDays" api:"nullable"`
	// The rule's results, one per entity the rule is evaluated against. Only returned
	// when `includeResults` is `true`.
	Results []GovernanceFrameworkSubsectionListRulesResponseItemsResult `json:"results"`
	// Pass-rate counts across all of the rule's entities, independent of any status
	// filter applied to the request.
	ResultsSummary GovernanceFrameworkSubsectionListRulesResponseItemsResultsSummary `json:"resultsSummary" api:"nullable"`
	// The rule tags associated with the rule.
	Tags []GovernanceFrameworkSubsectionListRulesResponseItemsTag `json:"tags" api:"nullable"`
	JSON governanceFrameworkSubsectionListRulesResponseItemJSON   `json:"-"`
}

// governanceFrameworkSubsectionListRulesResponseItemJSON contains the JSON
// metadata for the struct [GovernanceFrameworkSubsectionListRulesResponseItem]
type governanceFrameworkSubsectionListRulesResponseItemJSON struct {
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

func (r *GovernanceFrameworkSubsectionListRulesResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkSubsectionListRulesResponseItemJSON) RawJSON() string {
	return r.raw
}

// Whether the rule is evaluated once for the whole workspace, or once per project
// the rule's frameworks apply to.
type GovernanceFrameworkSubsectionListRulesResponseItemsScope string

const (
	GovernanceFrameworkSubsectionListRulesResponseItemsScopeProject   GovernanceFrameworkSubsectionListRulesResponseItemsScope = "project"
	GovernanceFrameworkSubsectionListRulesResponseItemsScopeWorkspace GovernanceFrameworkSubsectionListRulesResponseItemsScope = "workspace"
)

func (r GovernanceFrameworkSubsectionListRulesResponseItemsScope) IsKnown() bool {
	switch r {
	case GovernanceFrameworkSubsectionListRulesResponseItemsScopeProject, GovernanceFrameworkSubsectionListRulesResponseItemsScopeWorkspace:
		return true
	}
	return false
}

// `platform` rules are evaluated automatically from the state of your Openlayer
// workspace. `evidence` rules are satisfied by attaching evidence.
type GovernanceFrameworkSubsectionListRulesResponseItemsType string

const (
	GovernanceFrameworkSubsectionListRulesResponseItemsTypePlatform GovernanceFrameworkSubsectionListRulesResponseItemsType = "platform"
	GovernanceFrameworkSubsectionListRulesResponseItemsTypeEvidence GovernanceFrameworkSubsectionListRulesResponseItemsType = "evidence"
)

func (r GovernanceFrameworkSubsectionListRulesResponseItemsType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkSubsectionListRulesResponseItemsTypePlatform, GovernanceFrameworkSubsectionListRulesResponseItemsTypeEvidence:
		return true
	}
	return false
}

// The kind of evidence that satisfies the rule. `null` for platform rules.
type GovernanceFrameworkSubsectionListRulesResponseItemsEvidenceType string

const (
	GovernanceFrameworkSubsectionListRulesResponseItemsEvidenceTypeDocument      GovernanceFrameworkSubsectionListRulesResponseItemsEvidenceType = "document"
	GovernanceFrameworkSubsectionListRulesResponseItemsEvidenceTypeText          GovernanceFrameworkSubsectionListRulesResponseItemsEvidenceType = "text"
	GovernanceFrameworkSubsectionListRulesResponseItemsEvidenceTypeURL           GovernanceFrameworkSubsectionListRulesResponseItemsEvidenceType = "url"
	GovernanceFrameworkSubsectionListRulesResponseItemsEvidenceTypeCategoryValue GovernanceFrameworkSubsectionListRulesResponseItemsEvidenceType = "categoryValue"
)

func (r GovernanceFrameworkSubsectionListRulesResponseItemsEvidenceType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkSubsectionListRulesResponseItemsEvidenceTypeDocument, GovernanceFrameworkSubsectionListRulesResponseItemsEvidenceTypeText, GovernanceFrameworkSubsectionListRulesResponseItemsEvidenceTypeURL, GovernanceFrameworkSubsectionListRulesResponseItemsEvidenceTypeCategoryValue:
		return true
	}
	return false
}

type GovernanceFrameworkSubsectionListRulesResponseItemsFramework struct {
	// The framework id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The icon shown for the framework.
	Avatar GovernanceFrameworkSubsectionListRulesResponseItemsFrameworksAvatar `json:"avatar" api:"required,nullable"`
	// Identifies a framework that ships with Openlayer, for example `eu_ai_act`,
	// `iso_42001`, `nist_ai_rmf`, or `traiga`. `null` for frameworks you create
	// yourself.
	BuiltInSlug string `json:"builtInSlug" api:"required,nullable"`
	// Whether the framework is active. Rules of a disabled framework are not evaluated
	// and do not count towards compliance.
	Enabled bool `json:"enabled" api:"required"`
	// The framework name.
	Name string                                                           `json:"name" api:"required"`
	JSON governanceFrameworkSubsectionListRulesResponseItemsFrameworkJSON `json:"-"`
}

// governanceFrameworkSubsectionListRulesResponseItemsFrameworkJSON contains the
// JSON metadata for the struct
// [GovernanceFrameworkSubsectionListRulesResponseItemsFramework]
type governanceFrameworkSubsectionListRulesResponseItemsFrameworkJSON struct {
	ID          apijson.Field
	Avatar      apijson.Field
	BuiltInSlug apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkSubsectionListRulesResponseItemsFramework) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkSubsectionListRulesResponseItemsFrameworkJSON) RawJSON() string {
	return r.raw
}

// The icon shown for the framework.
type GovernanceFrameworkSubsectionListRulesResponseItemsFrameworksAvatar struct {
	Type  GovernanceFrameworkSubsectionListRulesResponseItemsFrameworksAvatarType `json:"type" api:"required"`
	Value string                                                                  `json:"value" api:"required"`
	JSON  governanceFrameworkSubsectionListRulesResponseItemsFrameworksAvatarJSON `json:"-"`
}

// governanceFrameworkSubsectionListRulesResponseItemsFrameworksAvatarJSON contains
// the JSON metadata for the struct
// [GovernanceFrameworkSubsectionListRulesResponseItemsFrameworksAvatar]
type governanceFrameworkSubsectionListRulesResponseItemsFrameworksAvatarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkSubsectionListRulesResponseItemsFrameworksAvatar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkSubsectionListRulesResponseItemsFrameworksAvatarJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkSubsectionListRulesResponseItemsFrameworksAvatarType string

const (
	GovernanceFrameworkSubsectionListRulesResponseItemsFrameworksAvatarTypeEmoji        GovernanceFrameworkSubsectionListRulesResponseItemsFrameworksAvatarType = "emoji"
	GovernanceFrameworkSubsectionListRulesResponseItemsFrameworksAvatarTypeImageURL     GovernanceFrameworkSubsectionListRulesResponseItemsFrameworksAvatarType = "imageUrl"
	GovernanceFrameworkSubsectionListRulesResponseItemsFrameworksAvatarTypeBuiltinImage GovernanceFrameworkSubsectionListRulesResponseItemsFrameworksAvatarType = "builtinImage"
)

func (r GovernanceFrameworkSubsectionListRulesResponseItemsFrameworksAvatarType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkSubsectionListRulesResponseItemsFrameworksAvatarTypeEmoji, GovernanceFrameworkSubsectionListRulesResponseItemsFrameworksAvatarTypeImageURL, GovernanceFrameworkSubsectionListRulesResponseItemsFrameworksAvatarTypeBuiltinImage:
		return true
	}
	return false
}

type GovernanceFrameworkSubsectionListRulesResponseItemsResult struct {
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
	Status GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatus `json:"status" api:"required"`
	// The id of the workspace the rule result belongs to.
	WorkspaceID string `json:"workspaceId" api:"required" format:"uuid"`
	// The user responsible for this result.
	AssigneeID string `json:"assigneeId" api:"nullable" format:"uuid"`
	// Rule results that must pass before this one can be satisfied.
	BlockedBy []GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedBy `json:"blockedBy"`
	// Rule results that this one blocks.
	Blocking []GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlocking `json:"blocking"`
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
	StatusMessage string                                                        `json:"statusMessage" api:"nullable"`
	JSON          governanceFrameworkSubsectionListRulesResponseItemsResultJSON `json:"-"`
}

// governanceFrameworkSubsectionListRulesResponseItemsResultJSON contains the JSON
// metadata for the struct
// [GovernanceFrameworkSubsectionListRulesResponseItemsResult]
type governanceFrameworkSubsectionListRulesResponseItemsResultJSON struct {
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

func (r *GovernanceFrameworkSubsectionListRulesResponseItemsResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkSubsectionListRulesResponseItemsResultJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatus string

const (
	GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatusRunning GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatus = "running"
	GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatusPassing GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatus = "passing"
	GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatusFailing GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatus = "failing"
	GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatusSkipped GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatus = "skipped"
	GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatusError   GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatus = "error"
	GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatusPending GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatus = "pending"
	GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatusDueSoon GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatus = "due_soon"
)

func (r GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatus) IsKnown() bool {
	switch r {
	case GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatusRunning, GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatusPassing, GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatusFailing, GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatusSkipped, GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatusError, GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatusPending, GovernanceFrameworkSubsectionListRulesResponseItemsResultsStatusDueSoon:
		return true
	}
	return false
}

type GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedBy struct {
	ID string `json:"id" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatus `json:"status"`
	JSON   governanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByJSON   `json:"-"`
}

// governanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByJSON contains
// the JSON metadata for the struct
// [GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedBy]
type governanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatus string

const (
	GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatusRunning GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatus = "running"
	GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatusPassing GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatus = "passing"
	GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatusFailing GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatus = "failing"
	GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatusSkipped GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatus = "skipped"
	GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatusError   GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatus = "error"
	GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatusPending GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatus = "pending"
	GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatusDueSoon GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatus = "due_soon"
)

func (r GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatus) IsKnown() bool {
	switch r {
	case GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatusRunning, GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatusPassing, GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatusFailing, GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatusSkipped, GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatusError, GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatusPending, GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockedByStatusDueSoon:
		return true
	}
	return false
}

type GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlocking struct {
	ID string `json:"id" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatus `json:"status"`
	JSON   governanceFrameworkSubsectionListRulesResponseItemsResultsBlockingJSON   `json:"-"`
}

// governanceFrameworkSubsectionListRulesResponseItemsResultsBlockingJSON contains
// the JSON metadata for the struct
// [GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlocking]
type governanceFrameworkSubsectionListRulesResponseItemsResultsBlockingJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlocking) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkSubsectionListRulesResponseItemsResultsBlockingJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatus string

const (
	GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatusRunning GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatus = "running"
	GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatusPassing GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatus = "passing"
	GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatusFailing GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatus = "failing"
	GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatusSkipped GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatus = "skipped"
	GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatusError   GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatus = "error"
	GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatusPending GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatus = "pending"
	GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatusDueSoon GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatus = "due_soon"
)

func (r GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatus) IsKnown() bool {
	switch r {
	case GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatusRunning, GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatusPassing, GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatusFailing, GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatusSkipped, GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatusError, GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatusPending, GovernanceFrameworkSubsectionListRulesResponseItemsResultsBlockingStatusDueSoon:
		return true
	}
	return false
}

// Pass-rate counts across all of the rule's entities, independent of any status
// filter applied to the request.
type GovernanceFrameworkSubsectionListRulesResponseItemsResultsSummary struct {
	Passing int64                                                                 `json:"passing"`
	Total   int64                                                                 `json:"total"`
	JSON    governanceFrameworkSubsectionListRulesResponseItemsResultsSummaryJSON `json:"-"`
}

// governanceFrameworkSubsectionListRulesResponseItemsResultsSummaryJSON contains
// the JSON metadata for the struct
// [GovernanceFrameworkSubsectionListRulesResponseItemsResultsSummary]
type governanceFrameworkSubsectionListRulesResponseItemsResultsSummaryJSON struct {
	Passing     apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkSubsectionListRulesResponseItemsResultsSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkSubsectionListRulesResponseItemsResultsSummaryJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkSubsectionListRulesResponseItemsTag struct {
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
	Color string                                                     `json:"color" api:"nullable"`
	JSON  governanceFrameworkSubsectionListRulesResponseItemsTagJSON `json:"-"`
}

// governanceFrameworkSubsectionListRulesResponseItemsTagJSON contains the JSON
// metadata for the struct [GovernanceFrameworkSubsectionListRulesResponseItemsTag]
type governanceFrameworkSubsectionListRulesResponseItemsTagJSON struct {
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

func (r *GovernanceFrameworkSubsectionListRulesResponseItemsTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkSubsectionListRulesResponseItemsTagJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkSubsectionListRulesParams struct {
	// Whether to include each rule's results inline, in a `results` array.
	IncludeResults param.Field[bool] `query:"includeResults"`
	// The page to return in a paginated query.
	Page param.Field[int64] `query:"page"`
	// Maximum number of items to return per page.
	PerPage param.Field[int64] `query:"perPage"`
	// Only include items that apply to this project.
	ProjectID param.Field[string] `query:"projectId" format:"uuid"`
	// Only include items whose rule result has this compliance status.
	Status param.Field[GovernanceFrameworkSubsectionListRulesParamsStatus] `query:"status"`
}

// URLQuery serializes [GovernanceFrameworkSubsectionListRulesParams]'s query
// parameters as `url.Values`.
func (r GovernanceFrameworkSubsectionListRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only include items whose rule result has this compliance status.
type GovernanceFrameworkSubsectionListRulesParamsStatus string

const (
	GovernanceFrameworkSubsectionListRulesParamsStatusRunning GovernanceFrameworkSubsectionListRulesParamsStatus = "running"
	GovernanceFrameworkSubsectionListRulesParamsStatusPassing GovernanceFrameworkSubsectionListRulesParamsStatus = "passing"
	GovernanceFrameworkSubsectionListRulesParamsStatusFailing GovernanceFrameworkSubsectionListRulesParamsStatus = "failing"
	GovernanceFrameworkSubsectionListRulesParamsStatusSkipped GovernanceFrameworkSubsectionListRulesParamsStatus = "skipped"
	GovernanceFrameworkSubsectionListRulesParamsStatusError   GovernanceFrameworkSubsectionListRulesParamsStatus = "error"
	GovernanceFrameworkSubsectionListRulesParamsStatusPending GovernanceFrameworkSubsectionListRulesParamsStatus = "pending"
	GovernanceFrameworkSubsectionListRulesParamsStatusDueSoon GovernanceFrameworkSubsectionListRulesParamsStatus = "due_soon"
)

func (r GovernanceFrameworkSubsectionListRulesParamsStatus) IsKnown() bool {
	switch r {
	case GovernanceFrameworkSubsectionListRulesParamsStatusRunning, GovernanceFrameworkSubsectionListRulesParamsStatusPassing, GovernanceFrameworkSubsectionListRulesParamsStatusFailing, GovernanceFrameworkSubsectionListRulesParamsStatusSkipped, GovernanceFrameworkSubsectionListRulesParamsStatusError, GovernanceFrameworkSubsectionListRulesParamsStatusPending, GovernanceFrameworkSubsectionListRulesParamsStatusDueSoon:
		return true
	}
	return false
}
