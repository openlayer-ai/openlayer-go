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

// GovernanceRuleTagService contains methods and other services that help with
// interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGovernanceRuleTagService] method instead.
type GovernanceRuleTagService struct {
	Options []option.RequestOption
}

// NewGovernanceRuleTagService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGovernanceRuleTagService(opts ...option.RequestOption) (r *GovernanceRuleTagService) {
	r = &GovernanceRuleTagService{}
	r.Options = opts
	return
}

// List the rule tags in a workspace.
//
// Tags group rules across frameworks, for example by team or by control family.
// Use the ids returned here with the `tags` filter on
// [List rules](/api-reference/rest/governance/list-rules).
func (r *GovernanceRuleTagService) List(ctx context.Context, workspaceID string, query GovernanceRuleTagListParams, opts ...option.RequestOption) (res *GovernanceRuleTagListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("workspaces/%s/rule-tags", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type GovernanceRuleTagListResponse struct {
	Items []GovernanceRuleTagListResponseItem `json:"items" api:"required"`
	JSON  governanceRuleTagListResponseJSON   `json:"-"`
}

// governanceRuleTagListResponseJSON contains the JSON metadata for the struct
// [GovernanceRuleTagListResponse]
type governanceRuleTagListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GovernanceRuleTagListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleTagListResponseJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleTagListResponseItem struct {
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
	Color string                                `json:"color" api:"nullable"`
	JSON  governanceRuleTagListResponseItemJSON `json:"-"`
}

// governanceRuleTagListResponseItemJSON contains the JSON metadata for the struct
// [GovernanceRuleTagListResponseItem]
type governanceRuleTagListResponseItemJSON struct {
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

func (r *GovernanceRuleTagListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r governanceRuleTagListResponseItemJSON) RawJSON() string {
	return r.raw
}

type GovernanceRuleTagListParams struct {
	// The page to return in a paginated query.
	Page param.Field[int64] `query:"page"`
	// Maximum number of items to return per page.
	PerPage param.Field[int64] `query:"perPage"`
}

// URLQuery serializes [GovernanceRuleTagListParams]'s query parameters as
// `url.Values`.
func (r GovernanceRuleTagListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
