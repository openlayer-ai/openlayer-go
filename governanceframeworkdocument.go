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

// GovernanceFrameworkDocumentService contains methods and other services that help
// with interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGovernanceFrameworkDocumentService] method instead.
type GovernanceFrameworkDocumentService struct {
	Options []option.RequestOption
}

// NewGovernanceFrameworkDocumentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewGovernanceFrameworkDocumentService(opts ...option.RequestOption) (r *GovernanceFrameworkDocumentService) {
	r = &GovernanceFrameworkDocumentService{}
	r.Options = opts
	return
}

// Retrieve a framework document, including its sections, subsections, and the
// rules mapped to each.
//
// Each section and subsection carries a `ruleCount`, so you can tell which
// requirements have rules mapped to them before drilling in.
func (r *GovernanceFrameworkDocumentService) Get(ctx context.Context, frameworkID string, documentID string, opts ...option.RequestOption) (res *GovernanceFrameworkDocumentGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if frameworkID == "" {
		err = errors.New("missing required frameworkId parameter")
		return nil, err
	}
	if documentID == "" {
		err = errors.New("missing required documentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("frameworks/%s/documents/%s", frameworkID, documentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List the documents attached to a framework.
//
// A document holds the text of the standard the framework is based on, split into
// sections and subsections. Retrieve a single document to get that structure,
// along with the rules mapped to each part of it.
func (r *GovernanceFrameworkDocumentService) List(ctx context.Context, frameworkID string, query GovernanceFrameworkDocumentListParams, opts ...option.RequestOption) (res *GovernanceFrameworkDocumentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if frameworkID == "" {
		err = errors.New("missing required frameworkId parameter")
		return nil, err
	}
	path := fmt.Sprintf("frameworks/%s/documents", frameworkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type GovernanceFrameworkDocumentGetResponse struct {
	// The document id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The creation date.
	DateCreated time.Time `json:"dateCreated" api:"required" format:"date-time"`
	// The last update date.
	DateUpdated time.Time `json:"dateUpdated" api:"required" format:"date-time"`
	// The framework the document belongs to.
	FrameworkID string `json:"frameworkId" api:"required" format:"uuid"`
	// The document title.
	Title string `json:"title" api:"required"`
	// The document's sections, in display order. Only returned when retrieving a
	// single document.
	Sections []GovernanceFrameworkDocumentGetResponseSection `json:"sections"`
	JSON     governanceFrameworkDocumentGetResponseJSON      `json:"-"`
}

// governanceFrameworkDocumentGetResponseJSON contains the JSON metadata for the
// struct [GovernanceFrameworkDocumentGetResponse]
type governanceFrameworkDocumentGetResponseJSON struct {
	ID          apijson.Field
	DateCreated apijson.Field
	DateUpdated apijson.Field
	FrameworkID apijson.Field
	Title       apijson.Field
	Sections    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkDocumentGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkDocumentGetResponseJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkDocumentGetResponseSection struct {
	// The section id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The document the section belongs to.
	DocumentID string `json:"documentId" api:"required" format:"uuid"`
	// The section number as it appears in the source standard.
	Number string `json:"number" api:"required"`
	// The position of the section within the document.
	SortOrder int64 `json:"sortOrder" api:"required"`
	// The section title.
	Title string `json:"title" api:"required"`
	// How many rules are linked to this section, including its subsections. Use it to
	// decide whether to fetch the section's rules.
	RuleCount int64 `json:"ruleCount"`
	// The rules linked directly to this section.
	Rules []GovernanceFrameworkDocumentGetResponseSectionsRule `json:"rules"`
	// The section's subsections, in display order.
	Subsections []GovernanceFrameworkDocumentGetResponseSectionsSubsection `json:"subsections"`
	// The section text.
	Text string                                            `json:"text" api:"nullable"`
	JSON governanceFrameworkDocumentGetResponseSectionJSON `json:"-"`
}

// governanceFrameworkDocumentGetResponseSectionJSON contains the JSON metadata for
// the struct [GovernanceFrameworkDocumentGetResponseSection]
type governanceFrameworkDocumentGetResponseSectionJSON struct {
	ID          apijson.Field
	DocumentID  apijson.Field
	Number      apijson.Field
	SortOrder   apijson.Field
	Title       apijson.Field
	RuleCount   apijson.Field
	Rules       apijson.Field
	Subsections apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkDocumentGetResponseSection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkDocumentGetResponseSectionJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkDocumentGetResponseSectionsRule struct {
	// The rule id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The rule name.
	Name string `json:"name" api:"required"`
	// Whether the rule is evaluated once for the whole workspace, or once per project
	// the rule's frameworks apply to.
	Scope GovernanceFrameworkDocumentGetResponseSectionsRulesScope `json:"scope" api:"required"`
	// `platform` rules are evaluated automatically from the state of your Openlayer
	// workspace. `evidence` rules are satisfied by attaching evidence.
	Type GovernanceFrameworkDocumentGetResponseSectionsRulesType `json:"type" api:"required"`
	// Configuration for the platform check, when the automation takes parameters.
	AutomationParams map[string]interface{} `json:"automationParams" api:"nullable"`
	// Which workspace signal a platform rule checks, for example
	// `monitoring_mode_enabled`, `test_setup`, or `project_owner_set`. `null` for
	// evidence rules.
	AutomationType string `json:"automationType" api:"nullable"`
	// The creation date.
	DateCreated time.Time `json:"dateCreated" format:"date-time"`
	// The last update date.
	DateUpdated time.Time `json:"dateUpdated" format:"date-time"`
	// Whether the rule is excluded from compliance calculations.
	Deactivated bool `json:"deactivated"`
	// What the rule requires.
	Description string `json:"description" api:"nullable"`
	// The kind of evidence that satisfies the rule. `null` for platform rules.
	EvidenceType GovernanceFrameworkDocumentGetResponseSectionsRulesEvidenceType `json:"evidenceType" api:"nullable"`
	// Whether the rule is managed by Openlayer and cannot be edited.
	Immutable bool `json:"immutable"`
	// How often evidence must be renewed, in days. Once evidence is older than this,
	// the rule result becomes `due_soon` and then `failing`.
	RenewalCadenceDays int64                                                  `json:"renewalCadenceDays" api:"nullable"`
	JSON               governanceFrameworkDocumentGetResponseSectionsRuleJSON `json:"-"`
}

// governanceFrameworkDocumentGetResponseSectionsRuleJSON contains the JSON
// metadata for the struct [GovernanceFrameworkDocumentGetResponseSectionsRule]
type governanceFrameworkDocumentGetResponseSectionsRuleJSON struct {
	ID                 apijson.Field
	Name               apijson.Field
	Scope              apijson.Field
	Type               apijson.Field
	AutomationParams   apijson.Field
	AutomationType     apijson.Field
	DateCreated        apijson.Field
	DateUpdated        apijson.Field
	Deactivated        apijson.Field
	Description        apijson.Field
	EvidenceType       apijson.Field
	Immutable          apijson.Field
	RenewalCadenceDays apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *GovernanceFrameworkDocumentGetResponseSectionsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkDocumentGetResponseSectionsRuleJSON) RawJSON() string {
	return r.raw
}

// Whether the rule is evaluated once for the whole workspace, or once per project
// the rule's frameworks apply to.
type GovernanceFrameworkDocumentGetResponseSectionsRulesScope string

const (
	GovernanceFrameworkDocumentGetResponseSectionsRulesScopeProject   GovernanceFrameworkDocumentGetResponseSectionsRulesScope = "project"
	GovernanceFrameworkDocumentGetResponseSectionsRulesScopeWorkspace GovernanceFrameworkDocumentGetResponseSectionsRulesScope = "workspace"
)

func (r GovernanceFrameworkDocumentGetResponseSectionsRulesScope) IsKnown() bool {
	switch r {
	case GovernanceFrameworkDocumentGetResponseSectionsRulesScopeProject, GovernanceFrameworkDocumentGetResponseSectionsRulesScopeWorkspace:
		return true
	}
	return false
}

// `platform` rules are evaluated automatically from the state of your Openlayer
// workspace. `evidence` rules are satisfied by attaching evidence.
type GovernanceFrameworkDocumentGetResponseSectionsRulesType string

const (
	GovernanceFrameworkDocumentGetResponseSectionsRulesTypePlatform GovernanceFrameworkDocumentGetResponseSectionsRulesType = "platform"
	GovernanceFrameworkDocumentGetResponseSectionsRulesTypeEvidence GovernanceFrameworkDocumentGetResponseSectionsRulesType = "evidence"
)

func (r GovernanceFrameworkDocumentGetResponseSectionsRulesType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkDocumentGetResponseSectionsRulesTypePlatform, GovernanceFrameworkDocumentGetResponseSectionsRulesTypeEvidence:
		return true
	}
	return false
}

// The kind of evidence that satisfies the rule. `null` for platform rules.
type GovernanceFrameworkDocumentGetResponseSectionsRulesEvidenceType string

const (
	GovernanceFrameworkDocumentGetResponseSectionsRulesEvidenceTypeDocument      GovernanceFrameworkDocumentGetResponseSectionsRulesEvidenceType = "document"
	GovernanceFrameworkDocumentGetResponseSectionsRulesEvidenceTypeText          GovernanceFrameworkDocumentGetResponseSectionsRulesEvidenceType = "text"
	GovernanceFrameworkDocumentGetResponseSectionsRulesEvidenceTypeURL           GovernanceFrameworkDocumentGetResponseSectionsRulesEvidenceType = "url"
	GovernanceFrameworkDocumentGetResponseSectionsRulesEvidenceTypeCategoryValue GovernanceFrameworkDocumentGetResponseSectionsRulesEvidenceType = "categoryValue"
)

func (r GovernanceFrameworkDocumentGetResponseSectionsRulesEvidenceType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkDocumentGetResponseSectionsRulesEvidenceTypeDocument, GovernanceFrameworkDocumentGetResponseSectionsRulesEvidenceTypeText, GovernanceFrameworkDocumentGetResponseSectionsRulesEvidenceTypeURL, GovernanceFrameworkDocumentGetResponseSectionsRulesEvidenceTypeCategoryValue:
		return true
	}
	return false
}

type GovernanceFrameworkDocumentGetResponseSectionsSubsection struct {
	// The subsection id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The subsection number as it appears in the source standard.
	Number string `json:"number" api:"required"`
	// The section the subsection belongs to.
	SectionID string `json:"sectionId" api:"required" format:"uuid"`
	// The position of the subsection within its section.
	SortOrder int64 `json:"sortOrder" api:"required"`
	// The subsection title.
	Title string `json:"title" api:"required"`
	// How many rules are linked to this subsection.
	RuleCount int64 `json:"ruleCount"`
	// The rules linked to this subsection.
	Rules []GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRule `json:"rules"`
	// The subsection text. This is the requirement your rules are mapped against.
	Text string                                                       `json:"text" api:"nullable"`
	JSON governanceFrameworkDocumentGetResponseSectionsSubsectionJSON `json:"-"`
}

// governanceFrameworkDocumentGetResponseSectionsSubsectionJSON contains the JSON
// metadata for the struct
// [GovernanceFrameworkDocumentGetResponseSectionsSubsection]
type governanceFrameworkDocumentGetResponseSectionsSubsectionJSON struct {
	ID          apijson.Field
	Number      apijson.Field
	SectionID   apijson.Field
	SortOrder   apijson.Field
	Title       apijson.Field
	RuleCount   apijson.Field
	Rules       apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkDocumentGetResponseSectionsSubsection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkDocumentGetResponseSectionsSubsectionJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRule struct {
	// The rule id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The rule name.
	Name string `json:"name" api:"required"`
	// Whether the rule is evaluated once for the whole workspace, or once per project
	// the rule's frameworks apply to.
	Scope GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesScope `json:"scope" api:"required"`
	// `platform` rules are evaluated automatically from the state of your Openlayer
	// workspace. `evidence` rules are satisfied by attaching evidence.
	Type GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesType `json:"type" api:"required"`
	// Configuration for the platform check, when the automation takes parameters.
	AutomationParams map[string]interface{} `json:"automationParams" api:"nullable"`
	// Which workspace signal a platform rule checks, for example
	// `monitoring_mode_enabled`, `test_setup`, or `project_owner_set`. `null` for
	// evidence rules.
	AutomationType string `json:"automationType" api:"nullable"`
	// The creation date.
	DateCreated time.Time `json:"dateCreated" format:"date-time"`
	// The last update date.
	DateUpdated time.Time `json:"dateUpdated" format:"date-time"`
	// Whether the rule is excluded from compliance calculations.
	Deactivated bool `json:"deactivated"`
	// What the rule requires.
	Description string `json:"description" api:"nullable"`
	// The kind of evidence that satisfies the rule. `null` for platform rules.
	EvidenceType GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesEvidenceType `json:"evidenceType" api:"nullable"`
	// Whether the rule is managed by Openlayer and cannot be edited.
	Immutable bool `json:"immutable"`
	// How often evidence must be renewed, in days. Once evidence is older than this,
	// the rule result becomes `due_soon` and then `failing`.
	RenewalCadenceDays int64                                                             `json:"renewalCadenceDays" api:"nullable"`
	JSON               governanceFrameworkDocumentGetResponseSectionsSubsectionsRuleJSON `json:"-"`
}

// governanceFrameworkDocumentGetResponseSectionsSubsectionsRuleJSON contains the
// JSON metadata for the struct
// [GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRule]
type governanceFrameworkDocumentGetResponseSectionsSubsectionsRuleJSON struct {
	ID                 apijson.Field
	Name               apijson.Field
	Scope              apijson.Field
	Type               apijson.Field
	AutomationParams   apijson.Field
	AutomationType     apijson.Field
	DateCreated        apijson.Field
	DateUpdated        apijson.Field
	Deactivated        apijson.Field
	Description        apijson.Field
	EvidenceType       apijson.Field
	Immutable          apijson.Field
	RenewalCadenceDays apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkDocumentGetResponseSectionsSubsectionsRuleJSON) RawJSON() string {
	return r.raw
}

// Whether the rule is evaluated once for the whole workspace, or once per project
// the rule's frameworks apply to.
type GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesScope string

const (
	GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesScopeProject   GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesScope = "project"
	GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesScopeWorkspace GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesScope = "workspace"
)

func (r GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesScope) IsKnown() bool {
	switch r {
	case GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesScopeProject, GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesScopeWorkspace:
		return true
	}
	return false
}

// `platform` rules are evaluated automatically from the state of your Openlayer
// workspace. `evidence` rules are satisfied by attaching evidence.
type GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesType string

const (
	GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesTypePlatform GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesType = "platform"
	GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesTypeEvidence GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesType = "evidence"
)

func (r GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesTypePlatform, GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesTypeEvidence:
		return true
	}
	return false
}

// The kind of evidence that satisfies the rule. `null` for platform rules.
type GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesEvidenceType string

const (
	GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesEvidenceTypeDocument      GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesEvidenceType = "document"
	GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesEvidenceTypeText          GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesEvidenceType = "text"
	GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesEvidenceTypeURL           GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesEvidenceType = "url"
	GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesEvidenceTypeCategoryValue GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesEvidenceType = "categoryValue"
)

func (r GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesEvidenceType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesEvidenceTypeDocument, GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesEvidenceTypeText, GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesEvidenceTypeURL, GovernanceFrameworkDocumentGetResponseSectionsSubsectionsRulesEvidenceTypeCategoryValue:
		return true
	}
	return false
}

type GovernanceFrameworkDocumentListResponse struct {
	Items []GovernanceFrameworkDocumentListResponseItem `json:"items" api:"required"`
	JSON  governanceFrameworkDocumentListResponseJSON   `json:"-"`
}

// governanceFrameworkDocumentListResponseJSON contains the JSON metadata for the
// struct [GovernanceFrameworkDocumentListResponse]
type governanceFrameworkDocumentListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkDocumentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkDocumentListResponseJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkDocumentListResponseItem struct {
	// The document id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The creation date.
	DateCreated time.Time `json:"dateCreated" api:"required" format:"date-time"`
	// The last update date.
	DateUpdated time.Time `json:"dateUpdated" api:"required" format:"date-time"`
	// The framework the document belongs to.
	FrameworkID string `json:"frameworkId" api:"required" format:"uuid"`
	// The document title.
	Title string `json:"title" api:"required"`
	// The document's sections, in display order. Only returned when retrieving a
	// single document.
	Sections []GovernanceFrameworkDocumentListResponseItemsSection `json:"sections"`
	JSON     governanceFrameworkDocumentListResponseItemJSON       `json:"-"`
}

// governanceFrameworkDocumentListResponseItemJSON contains the JSON metadata for
// the struct [GovernanceFrameworkDocumentListResponseItem]
type governanceFrameworkDocumentListResponseItemJSON struct {
	ID          apijson.Field
	DateCreated apijson.Field
	DateUpdated apijson.Field
	FrameworkID apijson.Field
	Title       apijson.Field
	Sections    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkDocumentListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkDocumentListResponseItemJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkDocumentListResponseItemsSection struct {
	// The section id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The document the section belongs to.
	DocumentID string `json:"documentId" api:"required" format:"uuid"`
	// The section number as it appears in the source standard.
	Number string `json:"number" api:"required"`
	// The position of the section within the document.
	SortOrder int64 `json:"sortOrder" api:"required"`
	// The section title.
	Title string `json:"title" api:"required"`
	// How many rules are linked to this section, including its subsections. Use it to
	// decide whether to fetch the section's rules.
	RuleCount int64 `json:"ruleCount"`
	// The rules linked directly to this section.
	Rules []GovernanceFrameworkDocumentListResponseItemsSectionsRule `json:"rules"`
	// The section's subsections, in display order.
	Subsections []GovernanceFrameworkDocumentListResponseItemsSectionsSubsection `json:"subsections"`
	// The section text.
	Text string                                                  `json:"text" api:"nullable"`
	JSON governanceFrameworkDocumentListResponseItemsSectionJSON `json:"-"`
}

// governanceFrameworkDocumentListResponseItemsSectionJSON contains the JSON
// metadata for the struct [GovernanceFrameworkDocumentListResponseItemsSection]
type governanceFrameworkDocumentListResponseItemsSectionJSON struct {
	ID          apijson.Field
	DocumentID  apijson.Field
	Number      apijson.Field
	SortOrder   apijson.Field
	Title       apijson.Field
	RuleCount   apijson.Field
	Rules       apijson.Field
	Subsections apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkDocumentListResponseItemsSection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkDocumentListResponseItemsSectionJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkDocumentListResponseItemsSectionsRule struct {
	// The rule id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The rule name.
	Name string `json:"name" api:"required"`
	// Whether the rule is evaluated once for the whole workspace, or once per project
	// the rule's frameworks apply to.
	Scope GovernanceFrameworkDocumentListResponseItemsSectionsRulesScope `json:"scope" api:"required"`
	// `platform` rules are evaluated automatically from the state of your Openlayer
	// workspace. `evidence` rules are satisfied by attaching evidence.
	Type GovernanceFrameworkDocumentListResponseItemsSectionsRulesType `json:"type" api:"required"`
	// Configuration for the platform check, when the automation takes parameters.
	AutomationParams map[string]interface{} `json:"automationParams" api:"nullable"`
	// Which workspace signal a platform rule checks, for example
	// `monitoring_mode_enabled`, `test_setup`, or `project_owner_set`. `null` for
	// evidence rules.
	AutomationType string `json:"automationType" api:"nullable"`
	// The creation date.
	DateCreated time.Time `json:"dateCreated" format:"date-time"`
	// The last update date.
	DateUpdated time.Time `json:"dateUpdated" format:"date-time"`
	// Whether the rule is excluded from compliance calculations.
	Deactivated bool `json:"deactivated"`
	// What the rule requires.
	Description string `json:"description" api:"nullable"`
	// The kind of evidence that satisfies the rule. `null` for platform rules.
	EvidenceType GovernanceFrameworkDocumentListResponseItemsSectionsRulesEvidenceType `json:"evidenceType" api:"nullable"`
	// Whether the rule is managed by Openlayer and cannot be edited.
	Immutable bool `json:"immutable"`
	// How often evidence must be renewed, in days. Once evidence is older than this,
	// the rule result becomes `due_soon` and then `failing`.
	RenewalCadenceDays int64                                                        `json:"renewalCadenceDays" api:"nullable"`
	JSON               governanceFrameworkDocumentListResponseItemsSectionsRuleJSON `json:"-"`
}

// governanceFrameworkDocumentListResponseItemsSectionsRuleJSON contains the JSON
// metadata for the struct
// [GovernanceFrameworkDocumentListResponseItemsSectionsRule]
type governanceFrameworkDocumentListResponseItemsSectionsRuleJSON struct {
	ID                 apijson.Field
	Name               apijson.Field
	Scope              apijson.Field
	Type               apijson.Field
	AutomationParams   apijson.Field
	AutomationType     apijson.Field
	DateCreated        apijson.Field
	DateUpdated        apijson.Field
	Deactivated        apijson.Field
	Description        apijson.Field
	EvidenceType       apijson.Field
	Immutable          apijson.Field
	RenewalCadenceDays apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *GovernanceFrameworkDocumentListResponseItemsSectionsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkDocumentListResponseItemsSectionsRuleJSON) RawJSON() string {
	return r.raw
}

// Whether the rule is evaluated once for the whole workspace, or once per project
// the rule's frameworks apply to.
type GovernanceFrameworkDocumentListResponseItemsSectionsRulesScope string

const (
	GovernanceFrameworkDocumentListResponseItemsSectionsRulesScopeProject   GovernanceFrameworkDocumentListResponseItemsSectionsRulesScope = "project"
	GovernanceFrameworkDocumentListResponseItemsSectionsRulesScopeWorkspace GovernanceFrameworkDocumentListResponseItemsSectionsRulesScope = "workspace"
)

func (r GovernanceFrameworkDocumentListResponseItemsSectionsRulesScope) IsKnown() bool {
	switch r {
	case GovernanceFrameworkDocumentListResponseItemsSectionsRulesScopeProject, GovernanceFrameworkDocumentListResponseItemsSectionsRulesScopeWorkspace:
		return true
	}
	return false
}

// `platform` rules are evaluated automatically from the state of your Openlayer
// workspace. `evidence` rules are satisfied by attaching evidence.
type GovernanceFrameworkDocumentListResponseItemsSectionsRulesType string

const (
	GovernanceFrameworkDocumentListResponseItemsSectionsRulesTypePlatform GovernanceFrameworkDocumentListResponseItemsSectionsRulesType = "platform"
	GovernanceFrameworkDocumentListResponseItemsSectionsRulesTypeEvidence GovernanceFrameworkDocumentListResponseItemsSectionsRulesType = "evidence"
)

func (r GovernanceFrameworkDocumentListResponseItemsSectionsRulesType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkDocumentListResponseItemsSectionsRulesTypePlatform, GovernanceFrameworkDocumentListResponseItemsSectionsRulesTypeEvidence:
		return true
	}
	return false
}

// The kind of evidence that satisfies the rule. `null` for platform rules.
type GovernanceFrameworkDocumentListResponseItemsSectionsRulesEvidenceType string

const (
	GovernanceFrameworkDocumentListResponseItemsSectionsRulesEvidenceTypeDocument      GovernanceFrameworkDocumentListResponseItemsSectionsRulesEvidenceType = "document"
	GovernanceFrameworkDocumentListResponseItemsSectionsRulesEvidenceTypeText          GovernanceFrameworkDocumentListResponseItemsSectionsRulesEvidenceType = "text"
	GovernanceFrameworkDocumentListResponseItemsSectionsRulesEvidenceTypeURL           GovernanceFrameworkDocumentListResponseItemsSectionsRulesEvidenceType = "url"
	GovernanceFrameworkDocumentListResponseItemsSectionsRulesEvidenceTypeCategoryValue GovernanceFrameworkDocumentListResponseItemsSectionsRulesEvidenceType = "categoryValue"
)

func (r GovernanceFrameworkDocumentListResponseItemsSectionsRulesEvidenceType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkDocumentListResponseItemsSectionsRulesEvidenceTypeDocument, GovernanceFrameworkDocumentListResponseItemsSectionsRulesEvidenceTypeText, GovernanceFrameworkDocumentListResponseItemsSectionsRulesEvidenceTypeURL, GovernanceFrameworkDocumentListResponseItemsSectionsRulesEvidenceTypeCategoryValue:
		return true
	}
	return false
}

type GovernanceFrameworkDocumentListResponseItemsSectionsSubsection struct {
	// The subsection id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The subsection number as it appears in the source standard.
	Number string `json:"number" api:"required"`
	// The section the subsection belongs to.
	SectionID string `json:"sectionId" api:"required" format:"uuid"`
	// The position of the subsection within its section.
	SortOrder int64 `json:"sortOrder" api:"required"`
	// The subsection title.
	Title string `json:"title" api:"required"`
	// How many rules are linked to this subsection.
	RuleCount int64 `json:"ruleCount"`
	// The rules linked to this subsection.
	Rules []GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRule `json:"rules"`
	// The subsection text. This is the requirement your rules are mapped against.
	Text string                                                             `json:"text" api:"nullable"`
	JSON governanceFrameworkDocumentListResponseItemsSectionsSubsectionJSON `json:"-"`
}

// governanceFrameworkDocumentListResponseItemsSectionsSubsectionJSON contains the
// JSON metadata for the struct
// [GovernanceFrameworkDocumentListResponseItemsSectionsSubsection]
type governanceFrameworkDocumentListResponseItemsSectionsSubsectionJSON struct {
	ID          apijson.Field
	Number      apijson.Field
	SectionID   apijson.Field
	SortOrder   apijson.Field
	Title       apijson.Field
	RuleCount   apijson.Field
	Rules       apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceFrameworkDocumentListResponseItemsSectionsSubsection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkDocumentListResponseItemsSectionsSubsectionJSON) RawJSON() string {
	return r.raw
}

type GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRule struct {
	// The rule id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The rule name.
	Name string `json:"name" api:"required"`
	// Whether the rule is evaluated once for the whole workspace, or once per project
	// the rule's frameworks apply to.
	Scope GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesScope `json:"scope" api:"required"`
	// `platform` rules are evaluated automatically from the state of your Openlayer
	// workspace. `evidence` rules are satisfied by attaching evidence.
	Type GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesType `json:"type" api:"required"`
	// Configuration for the platform check, when the automation takes parameters.
	AutomationParams map[string]interface{} `json:"automationParams" api:"nullable"`
	// Which workspace signal a platform rule checks, for example
	// `monitoring_mode_enabled`, `test_setup`, or `project_owner_set`. `null` for
	// evidence rules.
	AutomationType string `json:"automationType" api:"nullable"`
	// The creation date.
	DateCreated time.Time `json:"dateCreated" format:"date-time"`
	// The last update date.
	DateUpdated time.Time `json:"dateUpdated" format:"date-time"`
	// Whether the rule is excluded from compliance calculations.
	Deactivated bool `json:"deactivated"`
	// What the rule requires.
	Description string `json:"description" api:"nullable"`
	// The kind of evidence that satisfies the rule. `null` for platform rules.
	EvidenceType GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesEvidenceType `json:"evidenceType" api:"nullable"`
	// Whether the rule is managed by Openlayer and cannot be edited.
	Immutable bool `json:"immutable"`
	// How often evidence must be renewed, in days. Once evidence is older than this,
	// the rule result becomes `due_soon` and then `failing`.
	RenewalCadenceDays int64                                                                   `json:"renewalCadenceDays" api:"nullable"`
	JSON               governanceFrameworkDocumentListResponseItemsSectionsSubsectionsRuleJSON `json:"-"`
}

// governanceFrameworkDocumentListResponseItemsSectionsSubsectionsRuleJSON contains
// the JSON metadata for the struct
// [GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRule]
type governanceFrameworkDocumentListResponseItemsSectionsSubsectionsRuleJSON struct {
	ID                 apijson.Field
	Name               apijson.Field
	Scope              apijson.Field
	Type               apijson.Field
	AutomationParams   apijson.Field
	AutomationType     apijson.Field
	DateCreated        apijson.Field
	DateUpdated        apijson.Field
	Deactivated        apijson.Field
	Description        apijson.Field
	EvidenceType       apijson.Field
	Immutable          apijson.Field
	RenewalCadenceDays apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceFrameworkDocumentListResponseItemsSectionsSubsectionsRuleJSON) RawJSON() string {
	return r.raw
}

// Whether the rule is evaluated once for the whole workspace, or once per project
// the rule's frameworks apply to.
type GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesScope string

const (
	GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesScopeProject   GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesScope = "project"
	GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesScopeWorkspace GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesScope = "workspace"
)

func (r GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesScope) IsKnown() bool {
	switch r {
	case GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesScopeProject, GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesScopeWorkspace:
		return true
	}
	return false
}

// `platform` rules are evaluated automatically from the state of your Openlayer
// workspace. `evidence` rules are satisfied by attaching evidence.
type GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesType string

const (
	GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesTypePlatform GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesType = "platform"
	GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesTypeEvidence GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesType = "evidence"
)

func (r GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesTypePlatform, GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesTypeEvidence:
		return true
	}
	return false
}

// The kind of evidence that satisfies the rule. `null` for platform rules.
type GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesEvidenceType string

const (
	GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesEvidenceTypeDocument      GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesEvidenceType = "document"
	GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesEvidenceTypeText          GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesEvidenceType = "text"
	GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesEvidenceTypeURL           GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesEvidenceType = "url"
	GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesEvidenceTypeCategoryValue GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesEvidenceType = "categoryValue"
)

func (r GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesEvidenceType) IsKnown() bool {
	switch r {
	case GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesEvidenceTypeDocument, GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesEvidenceTypeText, GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesEvidenceTypeURL, GovernanceFrameworkDocumentListResponseItemsSectionsSubsectionsRulesEvidenceTypeCategoryValue:
		return true
	}
	return false
}

type GovernanceFrameworkDocumentListParams struct {
	// The page to return in a paginated query.
	Page param.Field[int64] `query:"page"`
	// Maximum number of items to return per page.
	PerPage param.Field[int64] `query:"perPage"`
}

// URLQuery serializes [GovernanceFrameworkDocumentListParams]'s query parameters
// as `url.Values`.
func (r GovernanceFrameworkDocumentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
