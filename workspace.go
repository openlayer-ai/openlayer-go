// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/openlayer-ai/openlayer-go/internal/apijson"
	"github.com/openlayer-ai/openlayer-go/internal/param"
	"github.com/openlayer-ai/openlayer-go/internal/requestconfig"
	"github.com/openlayer-ai/openlayer-go/option"
)

// WorkspaceService contains methods and other services that help with interacting
// with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceService] method instead.
type WorkspaceService struct {
	Options []option.RequestOption
	Invites *WorkspaceInviteService
	APIKeys *WorkspaceAPIKeyService
}

// NewWorkspaceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkspaceService(opts ...option.RequestOption) (r *WorkspaceService) {
	r = &WorkspaceService{}
	r.Options = opts
	r.Invites = NewWorkspaceInviteService(opts...)
	r.APIKeys = NewWorkspaceAPIKeyService(opts...)
	return
}

// Retrieve a workspace by its ID.
func (r *WorkspaceService) Get(ctx context.Context, workspaceID string, opts ...option.RequestOption) (res *WorkspaceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return
	}
	path := fmt.Sprintf("workspaces/%s", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a workspace.
func (r *WorkspaceService) Update(ctx context.Context, workspaceID string, body WorkspaceUpdateParams, opts ...option.RequestOption) (res *WorkspaceUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return
	}
	path := fmt.Sprintf("workspaces/%s", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type WorkspaceGetResponse struct {
	// The workspace id.
	ID string `json:"id,required" format:"uuid"`
	// The workspace creator id.
	CreatorID string `json:"creatorId,required,nullable" format:"uuid"`
	// The workspace creation date.
	DateCreated time.Time `json:"dateCreated,required" format:"date-time"`
	// The workspace last updated date.
	DateUpdated time.Time `json:"dateUpdated,required" format:"date-time"`
	// The number of invites in the workspace.
	InviteCount int64 `json:"inviteCount,required"`
	// The number of members in the workspace.
	MemberCount int64 `json:"memberCount,required"`
	// The workspace name.
	Name string `json:"name,required"`
	// The end date of the current billing period.
	PeriodEndDate time.Time `json:"periodEndDate,required,nullable" format:"date-time"`
	// The start date of the current billing period.
	PeriodStartDate time.Time `json:"periodStartDate,required,nullable" format:"date-time"`
	// The number of projects in the workspace.
	ProjectCount int64 `json:"projectCount,required"`
	// The workspace slug.
	Slug         string                             `json:"slug,required"`
	Status       WorkspaceGetResponseStatus         `json:"status,required"`
	MonthlyUsage []WorkspaceGetResponseMonthlyUsage `json:"monthlyUsage"`
	// Whether the workspace only allows SAML authentication.
	SAMLOnlyAccess  bool                     `json:"samlOnlyAccess"`
	WildcardDomains []string                 `json:"wildcardDomains"`
	JSON            workspaceGetResponseJSON `json:"-"`
}

// workspaceGetResponseJSON contains the JSON metadata for the struct
// [WorkspaceGetResponse]
type workspaceGetResponseJSON struct {
	ID              apijson.Field
	CreatorID       apijson.Field
	DateCreated     apijson.Field
	DateUpdated     apijson.Field
	InviteCount     apijson.Field
	MemberCount     apijson.Field
	Name            apijson.Field
	PeriodEndDate   apijson.Field
	PeriodStartDate apijson.Field
	ProjectCount    apijson.Field
	Slug            apijson.Field
	Status          apijson.Field
	MonthlyUsage    apijson.Field
	SAMLOnlyAccess  apijson.Field
	WildcardDomains apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WorkspaceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceGetResponseJSON) RawJSON() string {
	return r.raw
}

type WorkspaceGetResponseStatus string

const (
	WorkspaceGetResponseStatusActive            WorkspaceGetResponseStatus = "active"
	WorkspaceGetResponseStatusPastDue           WorkspaceGetResponseStatus = "past_due"
	WorkspaceGetResponseStatusUnpaid            WorkspaceGetResponseStatus = "unpaid"
	WorkspaceGetResponseStatusCanceled          WorkspaceGetResponseStatus = "canceled"
	WorkspaceGetResponseStatusIncomplete        WorkspaceGetResponseStatus = "incomplete"
	WorkspaceGetResponseStatusIncompleteExpired WorkspaceGetResponseStatus = "incomplete_expired"
	WorkspaceGetResponseStatusTrialing          WorkspaceGetResponseStatus = "trialing"
	WorkspaceGetResponseStatusPaused            WorkspaceGetResponseStatus = "paused"
)

func (r WorkspaceGetResponseStatus) IsKnown() bool {
	switch r {
	case WorkspaceGetResponseStatusActive, WorkspaceGetResponseStatusPastDue, WorkspaceGetResponseStatusUnpaid, WorkspaceGetResponseStatusCanceled, WorkspaceGetResponseStatusIncomplete, WorkspaceGetResponseStatusIncompleteExpired, WorkspaceGetResponseStatusTrialing, WorkspaceGetResponseStatusPaused:
		return true
	}
	return false
}

type WorkspaceGetResponseMonthlyUsage struct {
	ExecutionTimeMs int64                                `json:"executionTimeMs,nullable"`
	MonthYear       time.Time                            `json:"monthYear" format:"date"`
	PredictionCount int64                                `json:"predictionCount"`
	JSON            workspaceGetResponseMonthlyUsageJSON `json:"-"`
}

// workspaceGetResponseMonthlyUsageJSON contains the JSON metadata for the struct
// [WorkspaceGetResponseMonthlyUsage]
type workspaceGetResponseMonthlyUsageJSON struct {
	ExecutionTimeMs apijson.Field
	MonthYear       apijson.Field
	PredictionCount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WorkspaceGetResponseMonthlyUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceGetResponseMonthlyUsageJSON) RawJSON() string {
	return r.raw
}

type WorkspaceUpdateResponse struct {
	// The workspace id.
	ID string `json:"id,required" format:"uuid"`
	// The workspace creator id.
	CreatorID string `json:"creatorId,required,nullable" format:"uuid"`
	// The workspace creation date.
	DateCreated time.Time `json:"dateCreated,required" format:"date-time"`
	// The workspace last updated date.
	DateUpdated time.Time `json:"dateUpdated,required" format:"date-time"`
	// The number of invites in the workspace.
	InviteCount int64 `json:"inviteCount,required"`
	// The number of members in the workspace.
	MemberCount int64 `json:"memberCount,required"`
	// The workspace name.
	Name string `json:"name,required"`
	// The end date of the current billing period.
	PeriodEndDate time.Time `json:"periodEndDate,required,nullable" format:"date-time"`
	// The start date of the current billing period.
	PeriodStartDate time.Time `json:"periodStartDate,required,nullable" format:"date-time"`
	// The number of projects in the workspace.
	ProjectCount int64 `json:"projectCount,required"`
	// The workspace slug.
	Slug         string                                `json:"slug,required"`
	Status       WorkspaceUpdateResponseStatus         `json:"status,required"`
	MonthlyUsage []WorkspaceUpdateResponseMonthlyUsage `json:"monthlyUsage"`
	// Whether the workspace only allows SAML authentication.
	SAMLOnlyAccess  bool                        `json:"samlOnlyAccess"`
	WildcardDomains []string                    `json:"wildcardDomains"`
	JSON            workspaceUpdateResponseJSON `json:"-"`
}

// workspaceUpdateResponseJSON contains the JSON metadata for the struct
// [WorkspaceUpdateResponse]
type workspaceUpdateResponseJSON struct {
	ID              apijson.Field
	CreatorID       apijson.Field
	DateCreated     apijson.Field
	DateUpdated     apijson.Field
	InviteCount     apijson.Field
	MemberCount     apijson.Field
	Name            apijson.Field
	PeriodEndDate   apijson.Field
	PeriodStartDate apijson.Field
	ProjectCount    apijson.Field
	Slug            apijson.Field
	Status          apijson.Field
	MonthlyUsage    apijson.Field
	SAMLOnlyAccess  apijson.Field
	WildcardDomains apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WorkspaceUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type WorkspaceUpdateResponseStatus string

const (
	WorkspaceUpdateResponseStatusActive            WorkspaceUpdateResponseStatus = "active"
	WorkspaceUpdateResponseStatusPastDue           WorkspaceUpdateResponseStatus = "past_due"
	WorkspaceUpdateResponseStatusUnpaid            WorkspaceUpdateResponseStatus = "unpaid"
	WorkspaceUpdateResponseStatusCanceled          WorkspaceUpdateResponseStatus = "canceled"
	WorkspaceUpdateResponseStatusIncomplete        WorkspaceUpdateResponseStatus = "incomplete"
	WorkspaceUpdateResponseStatusIncompleteExpired WorkspaceUpdateResponseStatus = "incomplete_expired"
	WorkspaceUpdateResponseStatusTrialing          WorkspaceUpdateResponseStatus = "trialing"
	WorkspaceUpdateResponseStatusPaused            WorkspaceUpdateResponseStatus = "paused"
)

func (r WorkspaceUpdateResponseStatus) IsKnown() bool {
	switch r {
	case WorkspaceUpdateResponseStatusActive, WorkspaceUpdateResponseStatusPastDue, WorkspaceUpdateResponseStatusUnpaid, WorkspaceUpdateResponseStatusCanceled, WorkspaceUpdateResponseStatusIncomplete, WorkspaceUpdateResponseStatusIncompleteExpired, WorkspaceUpdateResponseStatusTrialing, WorkspaceUpdateResponseStatusPaused:
		return true
	}
	return false
}

type WorkspaceUpdateResponseMonthlyUsage struct {
	ExecutionTimeMs int64                                   `json:"executionTimeMs,nullable"`
	MonthYear       time.Time                               `json:"monthYear" format:"date"`
	PredictionCount int64                                   `json:"predictionCount"`
	JSON            workspaceUpdateResponseMonthlyUsageJSON `json:"-"`
}

// workspaceUpdateResponseMonthlyUsageJSON contains the JSON metadata for the
// struct [WorkspaceUpdateResponseMonthlyUsage]
type workspaceUpdateResponseMonthlyUsageJSON struct {
	ExecutionTimeMs apijson.Field
	MonthYear       apijson.Field
	PredictionCount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WorkspaceUpdateResponseMonthlyUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceUpdateResponseMonthlyUsageJSON) RawJSON() string {
	return r.raw
}

type WorkspaceUpdateParams struct {
	// The workspace invite code.
	InviteCode param.Field[string] `json:"inviteCode"`
	// The workspace name.
	Name param.Field[string] `json:"name"`
	// The workspace slug.
	Slug param.Field[string] `json:"slug"`
}

func (r WorkspaceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
