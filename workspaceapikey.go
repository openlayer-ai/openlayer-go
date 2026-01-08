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

// WorkspaceAPIKeyService contains methods and other services that help with
// interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceAPIKeyService] method instead.
type WorkspaceAPIKeyService struct {
	Options []option.RequestOption
}

// NewWorkspaceAPIKeyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkspaceAPIKeyService(opts ...option.RequestOption) (r *WorkspaceAPIKeyService) {
	r = &WorkspaceAPIKeyService{}
	r.Options = opts
	return
}

// Create a new API key in a workspace.
func (r *WorkspaceAPIKeyService) New(ctx context.Context, workspaceID string, body WorkspaceAPIKeyNewParams, opts ...option.RequestOption) (res *WorkspaceAPIKeyNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return
	}
	path := fmt.Sprintf("workspaces/%s/api-keys", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type WorkspaceAPIKeyNewResponse struct {
	// The API key id.
	ID string `json:"id,required" format:"uuid"`
	// The API key creation date.
	DateCreated time.Time `json:"dateCreated,required" format:"date-time"`
	// The API key last use date.
	DateLastUsed time.Time `json:"dateLastUsed,required,nullable" format:"date-time"`
	// The API key last update date.
	DateUpdated time.Time `json:"dateUpdated,required" format:"date-time"`
	// The API key value.
	SecureKey string `json:"secureKey,required"`
	// The API key name.
	Name string                         `json:"name,nullable"`
	JSON workspaceAPIKeyNewResponseJSON `json:"-"`
}

// workspaceAPIKeyNewResponseJSON contains the JSON metadata for the struct
// [WorkspaceAPIKeyNewResponse]
type workspaceAPIKeyNewResponseJSON struct {
	ID           apijson.Field
	DateCreated  apijson.Field
	DateLastUsed apijson.Field
	DateUpdated  apijson.Field
	SecureKey    apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WorkspaceAPIKeyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceAPIKeyNewResponseJSON) RawJSON() string {
	return r.raw
}

type WorkspaceAPIKeyNewParams struct {
	// The API key name.
	Name param.Field[string] `json:"name"`
}

func (r WorkspaceAPIKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
