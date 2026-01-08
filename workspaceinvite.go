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

// WorkspaceInviteService contains methods and other services that help with
// interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceInviteService] method instead.
type WorkspaceInviteService struct {
	Options []option.RequestOption
}

// NewWorkspaceInviteService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkspaceInviteService(opts ...option.RequestOption) (r *WorkspaceInviteService) {
	r = &WorkspaceInviteService{}
	r.Options = opts
	return
}

// Invite users to a workspace.
func (r *WorkspaceInviteService) New(ctx context.Context, workspaceID string, body WorkspaceInviteNewParams, opts ...option.RequestOption) (res *WorkspaceInviteNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return
	}
	path := fmt.Sprintf("workspaces/%s/invites", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a list of invites in a workspace.
func (r *WorkspaceInviteService) List(ctx context.Context, workspaceID string, query WorkspaceInviteListParams, opts ...option.RequestOption) (res *WorkspaceInviteListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return
	}
	path := fmt.Sprintf("workspaces/%s/invites", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type WorkspaceInviteNewResponse struct {
	Items []WorkspaceInviteNewResponseItem `json:"items,required"`
	JSON  workspaceInviteNewResponseJSON   `json:"-"`
}

// workspaceInviteNewResponseJSON contains the JSON metadata for the struct
// [WorkspaceInviteNewResponse]
type workspaceInviteNewResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceInviteNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceInviteNewResponseJSON) RawJSON() string {
	return r.raw
}

type WorkspaceInviteNewResponseItem struct {
	// The invite id.
	ID      string                                 `json:"id,required" format:"uuid"`
	Creator WorkspaceInviteNewResponseItemsCreator `json:"creator,required"`
	// The invite creation date.
	DateCreated time.Time `json:"dateCreated,required" format:"date-time"`
	// The invite email.
	Email string `json:"email,required" format:"email"`
	// The invite role.
	Role WorkspaceInviteNewResponseItemsRole `json:"role,required"`
	// The invite status.
	Status    WorkspaceInviteNewResponseItemsStatus    `json:"status,required"`
	Workspace WorkspaceInviteNewResponseItemsWorkspace `json:"workspace,required"`
	JSON      workspaceInviteNewResponseItemJSON       `json:"-"`
}

// workspaceInviteNewResponseItemJSON contains the JSON metadata for the struct
// [WorkspaceInviteNewResponseItem]
type workspaceInviteNewResponseItemJSON struct {
	ID          apijson.Field
	Creator     apijson.Field
	DateCreated apijson.Field
	Email       apijson.Field
	Role        apijson.Field
	Status      apijson.Field
	Workspace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceInviteNewResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceInviteNewResponseItemJSON) RawJSON() string {
	return r.raw
}

type WorkspaceInviteNewResponseItemsCreator struct {
	// The invite creator id.
	ID string `json:"id" format:"uuid"`
	// The invite creator name.
	Name string `json:"name,nullable"`
	// The invite creator username.
	Username string                                     `json:"username,nullable"`
	JSON     workspaceInviteNewResponseItemsCreatorJSON `json:"-"`
}

// workspaceInviteNewResponseItemsCreatorJSON contains the JSON metadata for the
// struct [WorkspaceInviteNewResponseItemsCreator]
type workspaceInviteNewResponseItemsCreatorJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceInviteNewResponseItemsCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceInviteNewResponseItemsCreatorJSON) RawJSON() string {
	return r.raw
}

// The invite role.
type WorkspaceInviteNewResponseItemsRole string

const (
	WorkspaceInviteNewResponseItemsRoleAdmin  WorkspaceInviteNewResponseItemsRole = "ADMIN"
	WorkspaceInviteNewResponseItemsRoleMember WorkspaceInviteNewResponseItemsRole = "MEMBER"
	WorkspaceInviteNewResponseItemsRoleViewer WorkspaceInviteNewResponseItemsRole = "VIEWER"
)

func (r WorkspaceInviteNewResponseItemsRole) IsKnown() bool {
	switch r {
	case WorkspaceInviteNewResponseItemsRoleAdmin, WorkspaceInviteNewResponseItemsRoleMember, WorkspaceInviteNewResponseItemsRoleViewer:
		return true
	}
	return false
}

// The invite status.
type WorkspaceInviteNewResponseItemsStatus string

const (
	WorkspaceInviteNewResponseItemsStatusAccepted WorkspaceInviteNewResponseItemsStatus = "accepted"
	WorkspaceInviteNewResponseItemsStatusPending  WorkspaceInviteNewResponseItemsStatus = "pending"
)

func (r WorkspaceInviteNewResponseItemsStatus) IsKnown() bool {
	switch r {
	case WorkspaceInviteNewResponseItemsStatusAccepted, WorkspaceInviteNewResponseItemsStatusPending:
		return true
	}
	return false
}

type WorkspaceInviteNewResponseItemsWorkspace struct {
	ID          string                                       `json:"id,required" format:"uuid"`
	DateCreated time.Time                                    `json:"dateCreated,required" format:"date-time"`
	MemberCount int64                                        `json:"memberCount,required"`
	Name        string                                       `json:"name,required"`
	Slug        string                                       `json:"slug,required"`
	JSON        workspaceInviteNewResponseItemsWorkspaceJSON `json:"-"`
}

// workspaceInviteNewResponseItemsWorkspaceJSON contains the JSON metadata for the
// struct [WorkspaceInviteNewResponseItemsWorkspace]
type workspaceInviteNewResponseItemsWorkspaceJSON struct {
	ID          apijson.Field
	DateCreated apijson.Field
	MemberCount apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceInviteNewResponseItemsWorkspace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceInviteNewResponseItemsWorkspaceJSON) RawJSON() string {
	return r.raw
}

type WorkspaceInviteListResponse struct {
	Meta  WorkspaceInviteListResponse_Meta  `json:"_meta,required"`
	Items []WorkspaceInviteListResponseItem `json:"items,required"`
	JSON  workspaceInviteListResponseJSON   `json:"-"`
}

// workspaceInviteListResponseJSON contains the JSON metadata for the struct
// [WorkspaceInviteListResponse]
type workspaceInviteListResponseJSON struct {
	Meta        apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceInviteListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceInviteListResponseJSON) RawJSON() string {
	return r.raw
}

type WorkspaceInviteListResponse_Meta struct {
	// The current page.
	Page int64 `json:"page,required"`
	// The number of items per page.
	PerPage int64 `json:"perPage,required"`
	// The total number of items.
	TotalItems int64 `json:"totalItems,required"`
	// The total number of pages.
	TotalPages int64                               `json:"totalPages,required"`
	JSON       workspaceInviteListResponseMetaJSON `json:"-"`
}

// workspaceInviteListResponseMetaJSON contains the JSON metadata for the struct
// [WorkspaceInviteListResponse_Meta]
type workspaceInviteListResponseMetaJSON struct {
	Page        apijson.Field
	PerPage     apijson.Field
	TotalItems  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceInviteListResponse_Meta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceInviteListResponseMetaJSON) RawJSON() string {
	return r.raw
}

type WorkspaceInviteListResponseItem struct {
	// The invite id.
	ID      string                                  `json:"id,required" format:"uuid"`
	Creator WorkspaceInviteListResponseItemsCreator `json:"creator,required"`
	// The invite creation date.
	DateCreated time.Time `json:"dateCreated,required" format:"date-time"`
	// The invite email.
	Email string `json:"email,required" format:"email"`
	// The invite role.
	Role WorkspaceInviteListResponseItemsRole `json:"role,required"`
	// The invite status.
	Status    WorkspaceInviteListResponseItemsStatus    `json:"status,required"`
	Workspace WorkspaceInviteListResponseItemsWorkspace `json:"workspace,required"`
	JSON      workspaceInviteListResponseItemJSON       `json:"-"`
}

// workspaceInviteListResponseItemJSON contains the JSON metadata for the struct
// [WorkspaceInviteListResponseItem]
type workspaceInviteListResponseItemJSON struct {
	ID          apijson.Field
	Creator     apijson.Field
	DateCreated apijson.Field
	Email       apijson.Field
	Role        apijson.Field
	Status      apijson.Field
	Workspace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceInviteListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceInviteListResponseItemJSON) RawJSON() string {
	return r.raw
}

type WorkspaceInviteListResponseItemsCreator struct {
	// The invite creator id.
	ID string `json:"id" format:"uuid"`
	// The invite creator name.
	Name string `json:"name,nullable"`
	// The invite creator username.
	Username string                                      `json:"username,nullable"`
	JSON     workspaceInviteListResponseItemsCreatorJSON `json:"-"`
}

// workspaceInviteListResponseItemsCreatorJSON contains the JSON metadata for the
// struct [WorkspaceInviteListResponseItemsCreator]
type workspaceInviteListResponseItemsCreatorJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceInviteListResponseItemsCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceInviteListResponseItemsCreatorJSON) RawJSON() string {
	return r.raw
}

// The invite role.
type WorkspaceInviteListResponseItemsRole string

const (
	WorkspaceInviteListResponseItemsRoleAdmin  WorkspaceInviteListResponseItemsRole = "ADMIN"
	WorkspaceInviteListResponseItemsRoleMember WorkspaceInviteListResponseItemsRole = "MEMBER"
	WorkspaceInviteListResponseItemsRoleViewer WorkspaceInviteListResponseItemsRole = "VIEWER"
)

func (r WorkspaceInviteListResponseItemsRole) IsKnown() bool {
	switch r {
	case WorkspaceInviteListResponseItemsRoleAdmin, WorkspaceInviteListResponseItemsRoleMember, WorkspaceInviteListResponseItemsRoleViewer:
		return true
	}
	return false
}

// The invite status.
type WorkspaceInviteListResponseItemsStatus string

const (
	WorkspaceInviteListResponseItemsStatusAccepted WorkspaceInviteListResponseItemsStatus = "accepted"
	WorkspaceInviteListResponseItemsStatusPending  WorkspaceInviteListResponseItemsStatus = "pending"
)

func (r WorkspaceInviteListResponseItemsStatus) IsKnown() bool {
	switch r {
	case WorkspaceInviteListResponseItemsStatusAccepted, WorkspaceInviteListResponseItemsStatusPending:
		return true
	}
	return false
}

type WorkspaceInviteListResponseItemsWorkspace struct {
	ID          string                                        `json:"id,required" format:"uuid"`
	DateCreated time.Time                                     `json:"dateCreated,required" format:"date-time"`
	MemberCount int64                                         `json:"memberCount,required"`
	Name        string                                        `json:"name,required"`
	Slug        string                                        `json:"slug,required"`
	JSON        workspaceInviteListResponseItemsWorkspaceJSON `json:"-"`
}

// workspaceInviteListResponseItemsWorkspaceJSON contains the JSON metadata for the
// struct [WorkspaceInviteListResponseItemsWorkspace]
type workspaceInviteListResponseItemsWorkspaceJSON struct {
	ID          apijson.Field
	DateCreated apijson.Field
	MemberCount apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceInviteListResponseItemsWorkspace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceInviteListResponseItemsWorkspaceJSON) RawJSON() string {
	return r.raw
}

type WorkspaceInviteNewParams struct {
	Emails param.Field[[]string] `json:"emails" format:"email"`
	// The member role.
	Role param.Field[WorkspaceInviteNewParamsRole] `json:"role"`
}

func (r WorkspaceInviteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The member role.
type WorkspaceInviteNewParamsRole string

const (
	WorkspaceInviteNewParamsRoleAdmin  WorkspaceInviteNewParamsRole = "ADMIN"
	WorkspaceInviteNewParamsRoleMember WorkspaceInviteNewParamsRole = "MEMBER"
	WorkspaceInviteNewParamsRoleViewer WorkspaceInviteNewParamsRole = "VIEWER"
)

func (r WorkspaceInviteNewParamsRole) IsKnown() bool {
	switch r {
	case WorkspaceInviteNewParamsRoleAdmin, WorkspaceInviteNewParamsRoleMember, WorkspaceInviteNewParamsRoleViewer:
		return true
	}
	return false
}

type WorkspaceInviteListParams struct {
	// The page to return in a paginated query.
	Page param.Field[int64] `query:"page"`
	// Maximum number of items to return per page.
	PerPage param.Field[int64] `query:"perPage"`
}

// URLQuery serializes [WorkspaceInviteListParams]'s query parameters as
// `url.Values`.
func (r WorkspaceInviteListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
