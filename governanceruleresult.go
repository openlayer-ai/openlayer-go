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

// GovernanceRuleResultService contains methods and other services that help with
// interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGovernanceRuleResultService] method instead.
type GovernanceRuleResultService struct {
	Options []option.RequestOption
}

// NewGovernanceRuleResultService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGovernanceRuleResultService(opts ...option.RequestOption) (r *GovernanceRuleResultService) {
	r = &GovernanceRuleResultService{}
	r.Options = opts
	return
}

// Retrieve a rule result by its id.
//
// Alongside the status, the response carries the evaluation and renewal dates that
// explain it: `dateLastEvaluated` and `dateOfNextEvaluation` for platform rules,
// `dateOfLatestEvidence` and `dateOfRenewal` for evidence rules.
func (r *GovernanceRuleResultService) Get(ctx context.Context, ruleResultID string, opts ...option.RequestOption) (res *GovernanceRuleResultGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if ruleResultID == "" {
		err = errors.New("missing required ruleResultId parameter")
		return nil, err
	}
	path := fmt.Sprintf("rule-results/%s", ruleResultID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a rule result. Only the fields you send are changed.
//
// Use this to assign an owner, or to exclude a single result from compliance
// without deactivating the rule everywhere. `deactivatedReason` is required when
// setting `deactivated` to `true`.
//
// A result's `status` is computed by Openlayer and cannot be set directly.
func (r *GovernanceRuleResultService) Update(ctx context.Context, ruleResultID string, body GovernanceRuleResultUpdateParams, opts ...option.RequestOption) (res *GovernanceRuleResultUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if ruleResultID == "" {
		err = errors.New("missing required ruleResultId parameter")
		return nil, err
	}
	path := fmt.Sprintf("rule-results/%s", ruleResultID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List rule results across a workspace.
//
// A rule result is the compliance status of one rule for one entity: a project for
// project-scoped rules, or the workspace itself for workspace-scoped rules. This
// is the endpoint to poll or export when you want your current compliance state,
// filtered to a framework, a project, or a status.
func (r *GovernanceRuleResultService) List(ctx context.Context, workspaceID string, query GovernanceRuleResultListParams, opts ...option.RequestOption) (res *GovernanceRuleResultListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("workspaces/%s/rule-results", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Attach evidence to a rule result, satisfying an evidence rule.
//
// Send the field that matches the rule's `evidenceType`: `storageUri` for an
// uploaded document, `text` for a written statement, or `url` for a link.
//
// For a document, upload the file first with `POST /storage/presigned-url` and
// send the resulting storage URI as `storageUri`.
//
// Attaching evidence re-evaluates the rule result. If the rule sets
// `renewalCadenceDays`, the renewal window restarts from this evidence.
func (r *GovernanceRuleResultService) NewEvidence(ctx context.Context, ruleResultID string, body GovernanceRuleResultNewEvidenceParams, opts ...option.RequestOption) (res *GovernanceRuleResultNewEvidenceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if ruleResultID == "" {
		err = errors.New("missing required ruleResultId parameter")
		return nil, err
	}
	path := fmt.Sprintf("rule-results/%s/evidence", ruleResultID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List the evidence attached to a rule result.
//
// Which field carries the evidence depends on the rule's `evidenceType`:
// `storageUri` for uploaded documents, `text` for written statements, and `url`
// for links.
func (r *GovernanceRuleResultService) ListEvidence(ctx context.Context, ruleResultID string, query GovernanceRuleResultListEvidenceParams, opts ...option.RequestOption) (res *GovernanceRuleResultListEvidenceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if ruleResultID == "" {
		err = errors.New("missing required ruleResultId parameter")
		return nil, err
	}
	path := fmt.Sprintf("rule-results/%s/evidence", ruleResultID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type GovernanceRuleResultGetResponse struct {
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
	Status GovernanceRuleResultGetResponseStatus `json:"status" api:"required"`
	// The id of the workspace the rule result belongs to.
	WorkspaceID string `json:"workspaceId" api:"required" format:"uuid"`
	// The user responsible for this result.
	AssigneeID string `json:"assigneeId" api:"nullable" format:"uuid"`
	// Rule results that must pass before this one can be satisfied.
	BlockedBy []GovernanceRuleResultGetResponseBlockedBy `json:"blockedBy"`
	// Rule results that this one blocks.
	Blocking []GovernanceRuleResultGetResponseBlocking `json:"blocking"`
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
	JSON          governanceRuleResultGetResponseJSON `json:"-"`
}

// governanceRuleResultGetResponseJSON contains the JSON metadata for the struct
// [GovernanceRuleResultGetResponse]
type governanceRuleResultGetResponseJSON struct {
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

func (r *GovernanceRuleResultGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleResultGetResponseJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceRuleResultGetResponseStatus string

const (
	GovernanceRuleResultGetResponseStatusRunning GovernanceRuleResultGetResponseStatus = "running"
	GovernanceRuleResultGetResponseStatusPassing GovernanceRuleResultGetResponseStatus = "passing"
	GovernanceRuleResultGetResponseStatusFailing GovernanceRuleResultGetResponseStatus = "failing"
	GovernanceRuleResultGetResponseStatusSkipped GovernanceRuleResultGetResponseStatus = "skipped"
	GovernanceRuleResultGetResponseStatusError   GovernanceRuleResultGetResponseStatus = "error"
	GovernanceRuleResultGetResponseStatusPending GovernanceRuleResultGetResponseStatus = "pending"
	GovernanceRuleResultGetResponseStatusDueSoon GovernanceRuleResultGetResponseStatus = "due_soon"
)

func (r GovernanceRuleResultGetResponseStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleResultGetResponseStatusRunning, GovernanceRuleResultGetResponseStatusPassing, GovernanceRuleResultGetResponseStatusFailing, GovernanceRuleResultGetResponseStatusSkipped, GovernanceRuleResultGetResponseStatusError, GovernanceRuleResultGetResponseStatusPending, GovernanceRuleResultGetResponseStatusDueSoon:
		return true
	}
	return false
}

type GovernanceRuleResultGetResponseBlockedBy struct {
	ID string `json:"id" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status GovernanceRuleResultGetResponseBlockedByStatus `json:"status"`
	JSON   governanceRuleResultGetResponseBlockedByJSON   `json:"-"`
}

// governanceRuleResultGetResponseBlockedByJSON contains the JSON metadata for the
// struct [GovernanceRuleResultGetResponseBlockedBy]
type governanceRuleResultGetResponseBlockedByJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleResultGetResponseBlockedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleResultGetResponseBlockedByJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceRuleResultGetResponseBlockedByStatus string

const (
	GovernanceRuleResultGetResponseBlockedByStatusRunning GovernanceRuleResultGetResponseBlockedByStatus = "running"
	GovernanceRuleResultGetResponseBlockedByStatusPassing GovernanceRuleResultGetResponseBlockedByStatus = "passing"
	GovernanceRuleResultGetResponseBlockedByStatusFailing GovernanceRuleResultGetResponseBlockedByStatus = "failing"
	GovernanceRuleResultGetResponseBlockedByStatusSkipped GovernanceRuleResultGetResponseBlockedByStatus = "skipped"
	GovernanceRuleResultGetResponseBlockedByStatusError   GovernanceRuleResultGetResponseBlockedByStatus = "error"
	GovernanceRuleResultGetResponseBlockedByStatusPending GovernanceRuleResultGetResponseBlockedByStatus = "pending"
	GovernanceRuleResultGetResponseBlockedByStatusDueSoon GovernanceRuleResultGetResponseBlockedByStatus = "due_soon"
)

func (r GovernanceRuleResultGetResponseBlockedByStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleResultGetResponseBlockedByStatusRunning, GovernanceRuleResultGetResponseBlockedByStatusPassing, GovernanceRuleResultGetResponseBlockedByStatusFailing, GovernanceRuleResultGetResponseBlockedByStatusSkipped, GovernanceRuleResultGetResponseBlockedByStatusError, GovernanceRuleResultGetResponseBlockedByStatusPending, GovernanceRuleResultGetResponseBlockedByStatusDueSoon:
		return true
	}
	return false
}

type GovernanceRuleResultGetResponseBlocking struct {
	ID string `json:"id" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status GovernanceRuleResultGetResponseBlockingStatus `json:"status"`
	JSON   governanceRuleResultGetResponseBlockingJSON   `json:"-"`
}

// governanceRuleResultGetResponseBlockingJSON contains the JSON metadata for the
// struct [GovernanceRuleResultGetResponseBlocking]
type governanceRuleResultGetResponseBlockingJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleResultGetResponseBlocking) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleResultGetResponseBlockingJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceRuleResultGetResponseBlockingStatus string

const (
	GovernanceRuleResultGetResponseBlockingStatusRunning GovernanceRuleResultGetResponseBlockingStatus = "running"
	GovernanceRuleResultGetResponseBlockingStatusPassing GovernanceRuleResultGetResponseBlockingStatus = "passing"
	GovernanceRuleResultGetResponseBlockingStatusFailing GovernanceRuleResultGetResponseBlockingStatus = "failing"
	GovernanceRuleResultGetResponseBlockingStatusSkipped GovernanceRuleResultGetResponseBlockingStatus = "skipped"
	GovernanceRuleResultGetResponseBlockingStatusError   GovernanceRuleResultGetResponseBlockingStatus = "error"
	GovernanceRuleResultGetResponseBlockingStatusPending GovernanceRuleResultGetResponseBlockingStatus = "pending"
	GovernanceRuleResultGetResponseBlockingStatusDueSoon GovernanceRuleResultGetResponseBlockingStatus = "due_soon"
)

func (r GovernanceRuleResultGetResponseBlockingStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleResultGetResponseBlockingStatusRunning, GovernanceRuleResultGetResponseBlockingStatusPassing, GovernanceRuleResultGetResponseBlockingStatusFailing, GovernanceRuleResultGetResponseBlockingStatusSkipped, GovernanceRuleResultGetResponseBlockingStatusError, GovernanceRuleResultGetResponseBlockingStatusPending, GovernanceRuleResultGetResponseBlockingStatusDueSoon:
		return true
	}
	return false
}

type GovernanceRuleResultUpdateResponse struct {
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
	Status GovernanceRuleResultUpdateResponseStatus `json:"status" api:"required"`
	// The id of the workspace the rule result belongs to.
	WorkspaceID string `json:"workspaceId" api:"required" format:"uuid"`
	// The user responsible for this result.
	AssigneeID string `json:"assigneeId" api:"nullable" format:"uuid"`
	// Rule results that must pass before this one can be satisfied.
	BlockedBy []GovernanceRuleResultUpdateResponseBlockedBy `json:"blockedBy"`
	// Rule results that this one blocks.
	Blocking []GovernanceRuleResultUpdateResponseBlocking `json:"blocking"`
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
	JSON          governanceRuleResultUpdateResponseJSON `json:"-"`
}

// governanceRuleResultUpdateResponseJSON contains the JSON metadata for the struct
// [GovernanceRuleResultUpdateResponse]
type governanceRuleResultUpdateResponseJSON struct {
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

func (r *GovernanceRuleResultUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleResultUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceRuleResultUpdateResponseStatus string

const (
	GovernanceRuleResultUpdateResponseStatusRunning GovernanceRuleResultUpdateResponseStatus = "running"
	GovernanceRuleResultUpdateResponseStatusPassing GovernanceRuleResultUpdateResponseStatus = "passing"
	GovernanceRuleResultUpdateResponseStatusFailing GovernanceRuleResultUpdateResponseStatus = "failing"
	GovernanceRuleResultUpdateResponseStatusSkipped GovernanceRuleResultUpdateResponseStatus = "skipped"
	GovernanceRuleResultUpdateResponseStatusError   GovernanceRuleResultUpdateResponseStatus = "error"
	GovernanceRuleResultUpdateResponseStatusPending GovernanceRuleResultUpdateResponseStatus = "pending"
	GovernanceRuleResultUpdateResponseStatusDueSoon GovernanceRuleResultUpdateResponseStatus = "due_soon"
)

func (r GovernanceRuleResultUpdateResponseStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleResultUpdateResponseStatusRunning, GovernanceRuleResultUpdateResponseStatusPassing, GovernanceRuleResultUpdateResponseStatusFailing, GovernanceRuleResultUpdateResponseStatusSkipped, GovernanceRuleResultUpdateResponseStatusError, GovernanceRuleResultUpdateResponseStatusPending, GovernanceRuleResultUpdateResponseStatusDueSoon:
		return true
	}
	return false
}

type GovernanceRuleResultUpdateResponseBlockedBy struct {
	ID string `json:"id" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status GovernanceRuleResultUpdateResponseBlockedByStatus `json:"status"`
	JSON   governanceRuleResultUpdateResponseBlockedByJSON   `json:"-"`
}

// governanceRuleResultUpdateResponseBlockedByJSON contains the JSON metadata for
// the struct [GovernanceRuleResultUpdateResponseBlockedBy]
type governanceRuleResultUpdateResponseBlockedByJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleResultUpdateResponseBlockedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleResultUpdateResponseBlockedByJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceRuleResultUpdateResponseBlockedByStatus string

const (
	GovernanceRuleResultUpdateResponseBlockedByStatusRunning GovernanceRuleResultUpdateResponseBlockedByStatus = "running"
	GovernanceRuleResultUpdateResponseBlockedByStatusPassing GovernanceRuleResultUpdateResponseBlockedByStatus = "passing"
	GovernanceRuleResultUpdateResponseBlockedByStatusFailing GovernanceRuleResultUpdateResponseBlockedByStatus = "failing"
	GovernanceRuleResultUpdateResponseBlockedByStatusSkipped GovernanceRuleResultUpdateResponseBlockedByStatus = "skipped"
	GovernanceRuleResultUpdateResponseBlockedByStatusError   GovernanceRuleResultUpdateResponseBlockedByStatus = "error"
	GovernanceRuleResultUpdateResponseBlockedByStatusPending GovernanceRuleResultUpdateResponseBlockedByStatus = "pending"
	GovernanceRuleResultUpdateResponseBlockedByStatusDueSoon GovernanceRuleResultUpdateResponseBlockedByStatus = "due_soon"
)

func (r GovernanceRuleResultUpdateResponseBlockedByStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleResultUpdateResponseBlockedByStatusRunning, GovernanceRuleResultUpdateResponseBlockedByStatusPassing, GovernanceRuleResultUpdateResponseBlockedByStatusFailing, GovernanceRuleResultUpdateResponseBlockedByStatusSkipped, GovernanceRuleResultUpdateResponseBlockedByStatusError, GovernanceRuleResultUpdateResponseBlockedByStatusPending, GovernanceRuleResultUpdateResponseBlockedByStatusDueSoon:
		return true
	}
	return false
}

type GovernanceRuleResultUpdateResponseBlocking struct {
	ID string `json:"id" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status GovernanceRuleResultUpdateResponseBlockingStatus `json:"status"`
	JSON   governanceRuleResultUpdateResponseBlockingJSON   `json:"-"`
}

// governanceRuleResultUpdateResponseBlockingJSON contains the JSON metadata for
// the struct [GovernanceRuleResultUpdateResponseBlocking]
type governanceRuleResultUpdateResponseBlockingJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleResultUpdateResponseBlocking) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleResultUpdateResponseBlockingJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceRuleResultUpdateResponseBlockingStatus string

const (
	GovernanceRuleResultUpdateResponseBlockingStatusRunning GovernanceRuleResultUpdateResponseBlockingStatus = "running"
	GovernanceRuleResultUpdateResponseBlockingStatusPassing GovernanceRuleResultUpdateResponseBlockingStatus = "passing"
	GovernanceRuleResultUpdateResponseBlockingStatusFailing GovernanceRuleResultUpdateResponseBlockingStatus = "failing"
	GovernanceRuleResultUpdateResponseBlockingStatusSkipped GovernanceRuleResultUpdateResponseBlockingStatus = "skipped"
	GovernanceRuleResultUpdateResponseBlockingStatusError   GovernanceRuleResultUpdateResponseBlockingStatus = "error"
	GovernanceRuleResultUpdateResponseBlockingStatusPending GovernanceRuleResultUpdateResponseBlockingStatus = "pending"
	GovernanceRuleResultUpdateResponseBlockingStatusDueSoon GovernanceRuleResultUpdateResponseBlockingStatus = "due_soon"
)

func (r GovernanceRuleResultUpdateResponseBlockingStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleResultUpdateResponseBlockingStatusRunning, GovernanceRuleResultUpdateResponseBlockingStatusPassing, GovernanceRuleResultUpdateResponseBlockingStatusFailing, GovernanceRuleResultUpdateResponseBlockingStatusSkipped, GovernanceRuleResultUpdateResponseBlockingStatusError, GovernanceRuleResultUpdateResponseBlockingStatusPending, GovernanceRuleResultUpdateResponseBlockingStatusDueSoon:
		return true
	}
	return false
}

type GovernanceRuleResultListResponse struct {
	Items []GovernanceRuleResultListResponseItem `json:"items" api:"required"`
	JSON  governanceRuleResultListResponseJSON   `json:"-"`
}

// governanceRuleResultListResponseJSON contains the JSON metadata for the struct
// [GovernanceRuleResultListResponse]
type governanceRuleResultListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleResultListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleResultListResponseJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleResultListResponseItem struct {
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
	Status GovernanceRuleResultListResponseItemsStatus `json:"status" api:"required"`
	// The id of the workspace the rule result belongs to.
	WorkspaceID string `json:"workspaceId" api:"required" format:"uuid"`
	// The user responsible for this result.
	AssigneeID string `json:"assigneeId" api:"nullable" format:"uuid"`
	// Rule results that must pass before this one can be satisfied.
	BlockedBy []GovernanceRuleResultListResponseItemsBlockedBy `json:"blockedBy"`
	// Rule results that this one blocks.
	Blocking []GovernanceRuleResultListResponseItemsBlocking `json:"blocking"`
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
	StatusMessage string                                   `json:"statusMessage" api:"nullable"`
	JSON          governanceRuleResultListResponseItemJSON `json:"-"`
}

// governanceRuleResultListResponseItemJSON contains the JSON metadata for the
// struct [GovernanceRuleResultListResponseItem]
type governanceRuleResultListResponseItemJSON struct {
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

func (r *GovernanceRuleResultListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleResultListResponseItemJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceRuleResultListResponseItemsStatus string

const (
	GovernanceRuleResultListResponseItemsStatusRunning GovernanceRuleResultListResponseItemsStatus = "running"
	GovernanceRuleResultListResponseItemsStatusPassing GovernanceRuleResultListResponseItemsStatus = "passing"
	GovernanceRuleResultListResponseItemsStatusFailing GovernanceRuleResultListResponseItemsStatus = "failing"
	GovernanceRuleResultListResponseItemsStatusSkipped GovernanceRuleResultListResponseItemsStatus = "skipped"
	GovernanceRuleResultListResponseItemsStatusError   GovernanceRuleResultListResponseItemsStatus = "error"
	GovernanceRuleResultListResponseItemsStatusPending GovernanceRuleResultListResponseItemsStatus = "pending"
	GovernanceRuleResultListResponseItemsStatusDueSoon GovernanceRuleResultListResponseItemsStatus = "due_soon"
)

func (r GovernanceRuleResultListResponseItemsStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleResultListResponseItemsStatusRunning, GovernanceRuleResultListResponseItemsStatusPassing, GovernanceRuleResultListResponseItemsStatusFailing, GovernanceRuleResultListResponseItemsStatusSkipped, GovernanceRuleResultListResponseItemsStatusError, GovernanceRuleResultListResponseItemsStatusPending, GovernanceRuleResultListResponseItemsStatusDueSoon:
		return true
	}
	return false
}

type GovernanceRuleResultListResponseItemsBlockedBy struct {
	ID string `json:"id" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status GovernanceRuleResultListResponseItemsBlockedByStatus `json:"status"`
	JSON   governanceRuleResultListResponseItemsBlockedByJSON   `json:"-"`
}

// governanceRuleResultListResponseItemsBlockedByJSON contains the JSON metadata
// for the struct [GovernanceRuleResultListResponseItemsBlockedBy]
type governanceRuleResultListResponseItemsBlockedByJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleResultListResponseItemsBlockedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleResultListResponseItemsBlockedByJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceRuleResultListResponseItemsBlockedByStatus string

const (
	GovernanceRuleResultListResponseItemsBlockedByStatusRunning GovernanceRuleResultListResponseItemsBlockedByStatus = "running"
	GovernanceRuleResultListResponseItemsBlockedByStatusPassing GovernanceRuleResultListResponseItemsBlockedByStatus = "passing"
	GovernanceRuleResultListResponseItemsBlockedByStatusFailing GovernanceRuleResultListResponseItemsBlockedByStatus = "failing"
	GovernanceRuleResultListResponseItemsBlockedByStatusSkipped GovernanceRuleResultListResponseItemsBlockedByStatus = "skipped"
	GovernanceRuleResultListResponseItemsBlockedByStatusError   GovernanceRuleResultListResponseItemsBlockedByStatus = "error"
	GovernanceRuleResultListResponseItemsBlockedByStatusPending GovernanceRuleResultListResponseItemsBlockedByStatus = "pending"
	GovernanceRuleResultListResponseItemsBlockedByStatusDueSoon GovernanceRuleResultListResponseItemsBlockedByStatus = "due_soon"
)

func (r GovernanceRuleResultListResponseItemsBlockedByStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleResultListResponseItemsBlockedByStatusRunning, GovernanceRuleResultListResponseItemsBlockedByStatusPassing, GovernanceRuleResultListResponseItemsBlockedByStatusFailing, GovernanceRuleResultListResponseItemsBlockedByStatusSkipped, GovernanceRuleResultListResponseItemsBlockedByStatusError, GovernanceRuleResultListResponseItemsBlockedByStatusPending, GovernanceRuleResultListResponseItemsBlockedByStatusDueSoon:
		return true
	}
	return false
}

type GovernanceRuleResultListResponseItemsBlocking struct {
	ID string `json:"id" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status GovernanceRuleResultListResponseItemsBlockingStatus `json:"status"`
	JSON   governanceRuleResultListResponseItemsBlockingJSON   `json:"-"`
}

// governanceRuleResultListResponseItemsBlockingJSON contains the JSON metadata for
// the struct [GovernanceRuleResultListResponseItemsBlocking]
type governanceRuleResultListResponseItemsBlockingJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleResultListResponseItemsBlocking) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleResultListResponseItemsBlockingJSON) RawJSON() string {
	return r.raw
}

// The compliance status of the rule for this entity.
type GovernanceRuleResultListResponseItemsBlockingStatus string

const (
	GovernanceRuleResultListResponseItemsBlockingStatusRunning GovernanceRuleResultListResponseItemsBlockingStatus = "running"
	GovernanceRuleResultListResponseItemsBlockingStatusPassing GovernanceRuleResultListResponseItemsBlockingStatus = "passing"
	GovernanceRuleResultListResponseItemsBlockingStatusFailing GovernanceRuleResultListResponseItemsBlockingStatus = "failing"
	GovernanceRuleResultListResponseItemsBlockingStatusSkipped GovernanceRuleResultListResponseItemsBlockingStatus = "skipped"
	GovernanceRuleResultListResponseItemsBlockingStatusError   GovernanceRuleResultListResponseItemsBlockingStatus = "error"
	GovernanceRuleResultListResponseItemsBlockingStatusPending GovernanceRuleResultListResponseItemsBlockingStatus = "pending"
	GovernanceRuleResultListResponseItemsBlockingStatusDueSoon GovernanceRuleResultListResponseItemsBlockingStatus = "due_soon"
)

func (r GovernanceRuleResultListResponseItemsBlockingStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleResultListResponseItemsBlockingStatusRunning, GovernanceRuleResultListResponseItemsBlockingStatusPassing, GovernanceRuleResultListResponseItemsBlockingStatusFailing, GovernanceRuleResultListResponseItemsBlockingStatusSkipped, GovernanceRuleResultListResponseItemsBlockingStatusError, GovernanceRuleResultListResponseItemsBlockingStatusPending, GovernanceRuleResultListResponseItemsBlockingStatusDueSoon:
		return true
	}
	return false
}

type GovernanceRuleResultNewEvidenceResponse struct {
	// The evidence id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The user who attached the evidence.
	CreatorID string `json:"creatorId" api:"required,nullable" format:"uuid"`
	// The creation date.
	DateCreated time.Time `json:"dateCreated" api:"required" format:"date-time"`
	// The last update date.
	DateUpdated time.Time `json:"dateUpdated" api:"required" format:"date-time"`
	// A description of what the evidence shows.
	Description string `json:"description" api:"nullable"`
	// The evidence name.
	Name string `json:"name" api:"nullable"`
	// Where the uploaded file is stored. Set when the rule's `evidenceType` is
	// `document`.
	StorageUri string `json:"storageUri" api:"nullable"`
	// The evidence text. Set when the rule's `evidenceType` is `text`.
	Text string `json:"text" api:"nullable"`
	// A link to the evidence. Set when the rule's `evidenceType` is `url`.
	URL  string                                      `json:"url" api:"nullable" format:"uri"`
	JSON governanceRuleResultNewEvidenceResponseJSON `json:"-"`
}

// governanceRuleResultNewEvidenceResponseJSON contains the JSON metadata for the
// struct [GovernanceRuleResultNewEvidenceResponse]
type governanceRuleResultNewEvidenceResponseJSON struct {
	ID          apijson.Field
	CreatorID   apijson.Field
	DateCreated apijson.Field
	DateUpdated apijson.Field
	Description apijson.Field
	Name        apijson.Field
	StorageUri  apijson.Field
	Text        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleResultNewEvidenceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleResultNewEvidenceResponseJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleResultListEvidenceResponse struct {
	Items []GovernanceRuleResultListEvidenceResponseItem `json:"items" api:"required"`
	JSON  governanceRuleResultListEvidenceResponseJSON   `json:"-"`
}

// governanceRuleResultListEvidenceResponseJSON contains the JSON metadata for the
// struct [GovernanceRuleResultListEvidenceResponse]
type governanceRuleResultListEvidenceResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleResultListEvidenceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleResultListEvidenceResponseJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleResultListEvidenceResponseItem struct {
	// The evidence id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The user who attached the evidence.
	CreatorID string `json:"creatorId" api:"required,nullable" format:"uuid"`
	// The creation date.
	DateCreated time.Time `json:"dateCreated" api:"required" format:"date-time"`
	// The last update date.
	DateUpdated time.Time `json:"dateUpdated" api:"required" format:"date-time"`
	// A description of what the evidence shows.
	Description string `json:"description" api:"nullable"`
	// The evidence name.
	Name string `json:"name" api:"nullable"`
	// Where the uploaded file is stored. Set when the rule's `evidenceType` is
	// `document`.
	StorageUri string `json:"storageUri" api:"nullable"`
	// The evidence text. Set when the rule's `evidenceType` is `text`.
	Text string `json:"text" api:"nullable"`
	// A link to the evidence. Set when the rule's `evidenceType` is `url`.
	URL  string                                           `json:"url" api:"nullable" format:"uri"`
	JSON governanceRuleResultListEvidenceResponseItemJSON `json:"-"`
}

// governanceRuleResultListEvidenceResponseItemJSON contains the JSON metadata for
// the struct [GovernanceRuleResultListEvidenceResponseItem]
type governanceRuleResultListEvidenceResponseItemJSON struct {
	ID          apijson.Field
	CreatorID   apijson.Field
	DateCreated apijson.Field
	DateUpdated apijson.Field
	Description apijson.Field
	Name        apijson.Field
	StorageUri  apijson.Field
	Text        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleResultListEvidenceResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleResultListEvidenceResponseItemJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleResultUpdateParams struct {
	// The user responsible for this result.
	AssigneeID param.Field[string] `json:"assigneeId" format:"uuid"`
	// Rule results that must pass before this one can be satisfied.
	BlockedBy param.Field[[]GovernanceRuleResultUpdateParamsBlockedBy] `json:"blockedBy"`
	// Rule results that this one blocks.
	Blocking param.Field[[]GovernanceRuleResultUpdateParamsBlocking] `json:"blocking"`
	// Whether this result is excluded from compliance calculations.
	Deactivated param.Field[bool] `json:"deactivated"`
	// Why the result was excluded.
	DeactivatedReason param.Field[string] `json:"deactivatedReason"`
}

func (r GovernanceRuleResultUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GovernanceRuleResultUpdateParamsBlockedBy struct {
	ID param.Field[string] `json:"id" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status param.Field[GovernanceRuleResultUpdateParamsBlockedByStatus] `json:"status"`
}

func (r GovernanceRuleResultUpdateParamsBlockedBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The compliance status of the rule for this entity.
type GovernanceRuleResultUpdateParamsBlockedByStatus string

const (
	GovernanceRuleResultUpdateParamsBlockedByStatusRunning GovernanceRuleResultUpdateParamsBlockedByStatus = "running"
	GovernanceRuleResultUpdateParamsBlockedByStatusPassing GovernanceRuleResultUpdateParamsBlockedByStatus = "passing"
	GovernanceRuleResultUpdateParamsBlockedByStatusFailing GovernanceRuleResultUpdateParamsBlockedByStatus = "failing"
	GovernanceRuleResultUpdateParamsBlockedByStatusSkipped GovernanceRuleResultUpdateParamsBlockedByStatus = "skipped"
	GovernanceRuleResultUpdateParamsBlockedByStatusError   GovernanceRuleResultUpdateParamsBlockedByStatus = "error"
	GovernanceRuleResultUpdateParamsBlockedByStatusPending GovernanceRuleResultUpdateParamsBlockedByStatus = "pending"
	GovernanceRuleResultUpdateParamsBlockedByStatusDueSoon GovernanceRuleResultUpdateParamsBlockedByStatus = "due_soon"
)

func (r GovernanceRuleResultUpdateParamsBlockedByStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleResultUpdateParamsBlockedByStatusRunning, GovernanceRuleResultUpdateParamsBlockedByStatusPassing, GovernanceRuleResultUpdateParamsBlockedByStatusFailing, GovernanceRuleResultUpdateParamsBlockedByStatusSkipped, GovernanceRuleResultUpdateParamsBlockedByStatusError, GovernanceRuleResultUpdateParamsBlockedByStatusPending, GovernanceRuleResultUpdateParamsBlockedByStatusDueSoon:
		return true
	}
	return false
}

type GovernanceRuleResultUpdateParamsBlocking struct {
	ID param.Field[string] `json:"id" format:"uuid"`
	// The compliance status of the rule for this entity.
	Status param.Field[GovernanceRuleResultUpdateParamsBlockingStatus] `json:"status"`
}

func (r GovernanceRuleResultUpdateParamsBlocking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The compliance status of the rule for this entity.
type GovernanceRuleResultUpdateParamsBlockingStatus string

const (
	GovernanceRuleResultUpdateParamsBlockingStatusRunning GovernanceRuleResultUpdateParamsBlockingStatus = "running"
	GovernanceRuleResultUpdateParamsBlockingStatusPassing GovernanceRuleResultUpdateParamsBlockingStatus = "passing"
	GovernanceRuleResultUpdateParamsBlockingStatusFailing GovernanceRuleResultUpdateParamsBlockingStatus = "failing"
	GovernanceRuleResultUpdateParamsBlockingStatusSkipped GovernanceRuleResultUpdateParamsBlockingStatus = "skipped"
	GovernanceRuleResultUpdateParamsBlockingStatusError   GovernanceRuleResultUpdateParamsBlockingStatus = "error"
	GovernanceRuleResultUpdateParamsBlockingStatusPending GovernanceRuleResultUpdateParamsBlockingStatus = "pending"
	GovernanceRuleResultUpdateParamsBlockingStatusDueSoon GovernanceRuleResultUpdateParamsBlockingStatus = "due_soon"
)

func (r GovernanceRuleResultUpdateParamsBlockingStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleResultUpdateParamsBlockingStatusRunning, GovernanceRuleResultUpdateParamsBlockingStatusPassing, GovernanceRuleResultUpdateParamsBlockingStatusFailing, GovernanceRuleResultUpdateParamsBlockingStatusSkipped, GovernanceRuleResultUpdateParamsBlockingStatusError, GovernanceRuleResultUpdateParamsBlockingStatusPending, GovernanceRuleResultUpdateParamsBlockingStatusDueSoon:
		return true
	}
	return false
}

type GovernanceRuleResultListParams struct {
	// Only include items belonging to at least one enabled framework.
	EnabledFrameworkOnly param.Field[bool] `query:"enabledFrameworkOnly"`
	// Only include items belonging to this framework.
	FrameworkID param.Field[string] `query:"frameworkId" format:"uuid"`
	// Whether to include rules that are not part of any framework.
	IncludeUnframed param.Field[bool] `query:"includeUnframed"`
	// The page to return in a paginated query.
	Page param.Field[int64] `query:"page"`
	// Maximum number of items to return per page.
	PerPage param.Field[int64] `query:"perPage"`
	// Only include items that apply to this project.
	ProjectID param.Field[string] `query:"projectId" format:"uuid"`
	// Only include results of this rule.
	RuleID param.Field[string] `query:"ruleId" format:"uuid"`
	// Only include rules with this scope.
	Scope param.Field[GovernanceRuleResultListParamsScope] `query:"scope"`
	// Filter by a free-text search over names and descriptions.
	SearchQuery param.Field[string] `query:"searchQuery"`
	// Only include items whose rule result has this compliance status.
	Status param.Field[GovernanceRuleResultListParamsStatus] `query:"status"`
	// Only include rules of this type.
	Type param.Field[GovernanceRuleResultListParamsType] `query:"type"`
}

// URLQuery serializes [GovernanceRuleResultListParams]'s query parameters as
// `url.Values`.
func (r GovernanceRuleResultListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only include rules with this scope.
type GovernanceRuleResultListParamsScope string

const (
	GovernanceRuleResultListParamsScopeProject   GovernanceRuleResultListParamsScope = "project"
	GovernanceRuleResultListParamsScopeWorkspace GovernanceRuleResultListParamsScope = "workspace"
)

func (r GovernanceRuleResultListParamsScope) IsKnown() bool {
	switch r {
	case GovernanceRuleResultListParamsScopeProject, GovernanceRuleResultListParamsScopeWorkspace:
		return true
	}
	return false
}

// Only include items whose rule result has this compliance status.
type GovernanceRuleResultListParamsStatus string

const (
	GovernanceRuleResultListParamsStatusRunning GovernanceRuleResultListParamsStatus = "running"
	GovernanceRuleResultListParamsStatusPassing GovernanceRuleResultListParamsStatus = "passing"
	GovernanceRuleResultListParamsStatusFailing GovernanceRuleResultListParamsStatus = "failing"
	GovernanceRuleResultListParamsStatusSkipped GovernanceRuleResultListParamsStatus = "skipped"
	GovernanceRuleResultListParamsStatusError   GovernanceRuleResultListParamsStatus = "error"
	GovernanceRuleResultListParamsStatusPending GovernanceRuleResultListParamsStatus = "pending"
	GovernanceRuleResultListParamsStatusDueSoon GovernanceRuleResultListParamsStatus = "due_soon"
)

func (r GovernanceRuleResultListParamsStatus) IsKnown() bool {
	switch r {
	case GovernanceRuleResultListParamsStatusRunning, GovernanceRuleResultListParamsStatusPassing, GovernanceRuleResultListParamsStatusFailing, GovernanceRuleResultListParamsStatusSkipped, GovernanceRuleResultListParamsStatusError, GovernanceRuleResultListParamsStatusPending, GovernanceRuleResultListParamsStatusDueSoon:
		return true
	}
	return false
}

// Only include rules of this type.
type GovernanceRuleResultListParamsType string

const (
	GovernanceRuleResultListParamsTypePlatform GovernanceRuleResultListParamsType = "platform"
	GovernanceRuleResultListParamsTypeEvidence GovernanceRuleResultListParamsType = "evidence"
)

func (r GovernanceRuleResultListParamsType) IsKnown() bool {
	switch r {
	case GovernanceRuleResultListParamsTypePlatform, GovernanceRuleResultListParamsTypeEvidence:
		return true
	}
	return false
}

type GovernanceRuleResultNewEvidenceParams struct {
	// A description of what the evidence shows.
	Description param.Field[string] `json:"description"`
	// The evidence name.
	Name param.Field[string] `json:"name"`
	// Where the uploaded file is stored. Set when the rule's `evidenceType` is
	// `document`.
	StorageUri param.Field[string] `json:"storageUri"`
	// The evidence text. Set when the rule's `evidenceType` is `text`.
	Text param.Field[string] `json:"text"`
	// A link to the evidence. Set when the rule's `evidenceType` is `url`.
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r GovernanceRuleResultNewEvidenceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GovernanceRuleResultListEvidenceParams struct {
	// The page to return in a paginated query.
	Page param.Field[int64] `query:"page"`
	// Maximum number of items to return per page.
	PerPage param.Field[int64] `query:"perPage"`
}

// URLQuery serializes [GovernanceRuleResultListEvidenceParams]'s query parameters
// as `url.Values`.
func (r GovernanceRuleResultListEvidenceParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
