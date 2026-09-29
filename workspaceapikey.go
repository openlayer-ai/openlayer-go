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

// Create a new API key in a workspace. The full secret is returned in `secret`,
// only in this response. Optionally set `expiresAt`. When you authenticate with an
// API key that expires, the new key can't outlive it: omit `expiresAt` to inherit
// that expiry, and a later expiry (or `null`) is rejected with 400.
func (r *WorkspaceAPIKeyService) New(ctx context.Context, workspaceID string, body WorkspaceAPIKeyNewParams, opts ...option.RequestOption) (res *WorkspaceAPIKeyNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("workspaces/%s/api-keys", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve one of your API keys, with its lifecycle status. The secret is never
// returned; `secureKey` is an obfuscated hint.
func (r *WorkspaceAPIKeyService) Get(ctx context.Context, workspaceID string, apiKeyID string, opts ...option.RequestOption) (res *WorkspaceAPIKeyGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if apiKeyID == "" {
		err = errors.New("missing required apiKeyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("workspaces/%s/api-keys/%s", workspaceID, apiKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Rename one of your API keys. A key's expiry can't be updated; rotate the key
// with a new `expiresAt` instead, so extending a key's life always issues a new
// secret.
func (r *WorkspaceAPIKeyService) Update(ctx context.Context, workspaceID string, apiKeyID string, body WorkspaceAPIKeyUpdateParams, opts ...option.RequestOption) (res *WorkspaceAPIKeyUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if apiKeyID == "" {
		err = errors.New("missing required apiKeyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("workspaces/%s/api-keys/%s", workspaceID, apiKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// List the API keys you own in a workspace, with their lifecycle status. Secrets
// are never returned; `secureKey` is an obfuscated hint.
func (r *WorkspaceAPIKeyService) List(ctx context.Context, workspaceID string, opts ...option.RequestOption) (res *[]WorkspaceAPIKeyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("workspaces/%s/api-keys", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete one of your API keys. Every secret for the key stops working immediately,
// including a previous secret still in its rotation grace period.
func (r *WorkspaceAPIKeyService) Delete(ctx context.Context, workspaceID string, apiKeyID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return err
	}
	if apiKeyID == "" {
		err = errors.New("missing required apiKeyId parameter")
		return err
	}
	path := fmt.Sprintf("workspaces/%s/api-keys/%s", workspaceID, apiKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Replace an API key's secret now. The new secret is returned in `secret`, only in
// this response. Send `expiresAt` to change the key's expiry (`null` for never);
// omit it to keep the current one. The previous secret keeps authenticating for
// `gracePeriodHours` (default 0, so it stops working immediately), and never past
// `expiresAt`. The key keeps its id and name. Expired keys cannot be rotated. Only
// one previous secret is kept, so rotating again during a grace period retires the
// older one immediately.
func (r *WorkspaceAPIKeyService) Rotate(ctx context.Context, workspaceID string, apiKeyID string, body WorkspaceAPIKeyRotateParams, opts ...option.RequestOption) (res *WorkspaceAPIKeyRotateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if apiKeyID == "" {
		err = errors.New("missing required apiKeyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("workspaces/%s/api-keys/%s/rotate", workspaceID, apiKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type WorkspaceAPIKeyNewResponse struct {
	// The API key id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The API key creation date.
	DateCreated time.Time `json:"dateCreated" api:"required" format:"date-time"`
	// The API key last use date.
	DateLastUsed time.Time `json:"dateLastUsed" api:"required,nullable" format:"date-time"`
	// The API key last update date.
	DateUpdated time.Time `json:"dateUpdated" api:"required" format:"date-time"`
	// An obfuscated hint of the API key value. When a key is created or rotated this
	// also holds the full secret, for backward compatibility; prefer `secret`.
	SecureKey string `json:"secureKey" api:"required"`
	// The key's lifecycle state. `active`: the current secret authenticates.
	// `rotating`: the key was rotated and the previous secret still authenticates
	// until `previousKeyExpiresAt`. `expired`: `expiresAt` has passed and no secret
	// authenticates.
	Status WorkspaceAPIKeyNewResponseStatus `json:"status" api:"required"`
	// When the key stops authenticating. `null` means the key never expires. Set when
	// the key is created or rotated, and must be in the future. When the request is
	// authenticated with an API key that expires, the result can't be later than that
	// key's expiry.
	ExpiresAt time.Time `json:"expiresAt" api:"nullable" format:"date-time"`
	// When the key was last rotated.
	LastRotatedAt time.Time `json:"lastRotatedAt" api:"nullable" format:"date-time"`
	// The API key name.
	Name string `json:"name" api:"nullable"`
	// While `status` is `rotating`, when the previous secret stops authenticating.
	PreviousKeyExpiresAt time.Time `json:"previousKeyExpiresAt" api:"nullable" format:"date-time"`
	// The full API key. Only present in the response that creates or rotates the key,
	// and never shown again.
	Secret string                         `json:"secret"`
	JSON   workspaceAPIKeyNewResponseJSON `json:"-"`
}

// workspaceAPIKeyNewResponseJSON contains the JSON metadata for the struct
// [WorkspaceAPIKeyNewResponse]
type workspaceAPIKeyNewResponseJSON struct {
	ID                   apijson.Field
	DateCreated          apijson.Field
	DateLastUsed         apijson.Field
	DateUpdated          apijson.Field
	SecureKey            apijson.Field
	Status               apijson.Field
	ExpiresAt            apijson.Field
	LastRotatedAt        apijson.Field
	Name                 apijson.Field
	PreviousKeyExpiresAt apijson.Field
	Secret               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *WorkspaceAPIKeyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceAPIKeyNewResponseJSON) RawJSON() string {
	return r.raw
}

// The key's lifecycle state. `active`: the current secret authenticates.
// `rotating`: the key was rotated and the previous secret still authenticates
// until `previousKeyExpiresAt`. `expired`: `expiresAt` has passed and no secret
// authenticates.
type WorkspaceAPIKeyNewResponseStatus string

const (
	WorkspaceAPIKeyNewResponseStatusActive   WorkspaceAPIKeyNewResponseStatus = "active"
	WorkspaceAPIKeyNewResponseStatusRotating WorkspaceAPIKeyNewResponseStatus = "rotating"
	WorkspaceAPIKeyNewResponseStatusExpired  WorkspaceAPIKeyNewResponseStatus = "expired"
)

func (r WorkspaceAPIKeyNewResponseStatus) IsKnown() bool {
	switch r {
	case WorkspaceAPIKeyNewResponseStatusActive, WorkspaceAPIKeyNewResponseStatusRotating, WorkspaceAPIKeyNewResponseStatusExpired:
		return true
	}
	return false
}

type WorkspaceAPIKeyGetResponse struct {
	// The API key id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The API key creation date.
	DateCreated time.Time `json:"dateCreated" api:"required" format:"date-time"`
	// The API key last use date.
	DateLastUsed time.Time `json:"dateLastUsed" api:"required,nullable" format:"date-time"`
	// The API key last update date.
	DateUpdated time.Time `json:"dateUpdated" api:"required" format:"date-time"`
	// An obfuscated hint of the API key value. When a key is created or rotated this
	// also holds the full secret, for backward compatibility; prefer `secret`.
	SecureKey string `json:"secureKey" api:"required"`
	// The key's lifecycle state. `active`: the current secret authenticates.
	// `rotating`: the key was rotated and the previous secret still authenticates
	// until `previousKeyExpiresAt`. `expired`: `expiresAt` has passed and no secret
	// authenticates.
	Status WorkspaceAPIKeyGetResponseStatus `json:"status" api:"required"`
	// When the key stops authenticating. `null` means the key never expires. Set when
	// the key is created or rotated, and must be in the future. When the request is
	// authenticated with an API key that expires, the result can't be later than that
	// key's expiry.
	ExpiresAt time.Time `json:"expiresAt" api:"nullable" format:"date-time"`
	// When the key was last rotated.
	LastRotatedAt time.Time `json:"lastRotatedAt" api:"nullable" format:"date-time"`
	// The API key name.
	Name string `json:"name" api:"nullable"`
	// While `status` is `rotating`, when the previous secret stops authenticating.
	PreviousKeyExpiresAt time.Time `json:"previousKeyExpiresAt" api:"nullable" format:"date-time"`
	// The full API key. Only present in the response that creates or rotates the key,
	// and never shown again.
	Secret string                         `json:"secret"`
	JSON   workspaceAPIKeyGetResponseJSON `json:"-"`
}

// workspaceAPIKeyGetResponseJSON contains the JSON metadata for the struct
// [WorkspaceAPIKeyGetResponse]
type workspaceAPIKeyGetResponseJSON struct {
	ID                   apijson.Field
	DateCreated          apijson.Field
	DateLastUsed         apijson.Field
	DateUpdated          apijson.Field
	SecureKey            apijson.Field
	Status               apijson.Field
	ExpiresAt            apijson.Field
	LastRotatedAt        apijson.Field
	Name                 apijson.Field
	PreviousKeyExpiresAt apijson.Field
	Secret               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *WorkspaceAPIKeyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceAPIKeyGetResponseJSON) RawJSON() string {
	return r.raw
}

// The key's lifecycle state. `active`: the current secret authenticates.
// `rotating`: the key was rotated and the previous secret still authenticates
// until `previousKeyExpiresAt`. `expired`: `expiresAt` has passed and no secret
// authenticates.
type WorkspaceAPIKeyGetResponseStatus string

const (
	WorkspaceAPIKeyGetResponseStatusActive   WorkspaceAPIKeyGetResponseStatus = "active"
	WorkspaceAPIKeyGetResponseStatusRotating WorkspaceAPIKeyGetResponseStatus = "rotating"
	WorkspaceAPIKeyGetResponseStatusExpired  WorkspaceAPIKeyGetResponseStatus = "expired"
)

func (r WorkspaceAPIKeyGetResponseStatus) IsKnown() bool {
	switch r {
	case WorkspaceAPIKeyGetResponseStatusActive, WorkspaceAPIKeyGetResponseStatusRotating, WorkspaceAPIKeyGetResponseStatusExpired:
		return true
	}
	return false
}

type WorkspaceAPIKeyUpdateResponse struct {
	// The API key id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The API key creation date.
	DateCreated time.Time `json:"dateCreated" api:"required" format:"date-time"`
	// The API key last use date.
	DateLastUsed time.Time `json:"dateLastUsed" api:"required,nullable" format:"date-time"`
	// The API key last update date.
	DateUpdated time.Time `json:"dateUpdated" api:"required" format:"date-time"`
	// An obfuscated hint of the API key value. When a key is created or rotated this
	// also holds the full secret, for backward compatibility; prefer `secret`.
	SecureKey string `json:"secureKey" api:"required"`
	// The key's lifecycle state. `active`: the current secret authenticates.
	// `rotating`: the key was rotated and the previous secret still authenticates
	// until `previousKeyExpiresAt`. `expired`: `expiresAt` has passed and no secret
	// authenticates.
	Status WorkspaceAPIKeyUpdateResponseStatus `json:"status" api:"required"`
	// When the key stops authenticating. `null` means the key never expires. Set when
	// the key is created or rotated, and must be in the future. When the request is
	// authenticated with an API key that expires, the result can't be later than that
	// key's expiry.
	ExpiresAt time.Time `json:"expiresAt" api:"nullable" format:"date-time"`
	// When the key was last rotated.
	LastRotatedAt time.Time `json:"lastRotatedAt" api:"nullable" format:"date-time"`
	// The API key name.
	Name string `json:"name" api:"nullable"`
	// While `status` is `rotating`, when the previous secret stops authenticating.
	PreviousKeyExpiresAt time.Time `json:"previousKeyExpiresAt" api:"nullable" format:"date-time"`
	// The full API key. Only present in the response that creates or rotates the key,
	// and never shown again.
	Secret string                            `json:"secret"`
	JSON   workspaceAPIKeyUpdateResponseJSON `json:"-"`
}

// workspaceAPIKeyUpdateResponseJSON contains the JSON metadata for the struct
// [WorkspaceAPIKeyUpdateResponse]
type workspaceAPIKeyUpdateResponseJSON struct {
	ID                   apijson.Field
	DateCreated          apijson.Field
	DateLastUsed         apijson.Field
	DateUpdated          apijson.Field
	SecureKey            apijson.Field
	Status               apijson.Field
	ExpiresAt            apijson.Field
	LastRotatedAt        apijson.Field
	Name                 apijson.Field
	PreviousKeyExpiresAt apijson.Field
	Secret               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *WorkspaceAPIKeyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceAPIKeyUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The key's lifecycle state. `active`: the current secret authenticates.
// `rotating`: the key was rotated and the previous secret still authenticates
// until `previousKeyExpiresAt`. `expired`: `expiresAt` has passed and no secret
// authenticates.
type WorkspaceAPIKeyUpdateResponseStatus string

const (
	WorkspaceAPIKeyUpdateResponseStatusActive   WorkspaceAPIKeyUpdateResponseStatus = "active"
	WorkspaceAPIKeyUpdateResponseStatusRotating WorkspaceAPIKeyUpdateResponseStatus = "rotating"
	WorkspaceAPIKeyUpdateResponseStatusExpired  WorkspaceAPIKeyUpdateResponseStatus = "expired"
)

func (r WorkspaceAPIKeyUpdateResponseStatus) IsKnown() bool {
	switch r {
	case WorkspaceAPIKeyUpdateResponseStatusActive, WorkspaceAPIKeyUpdateResponseStatusRotating, WorkspaceAPIKeyUpdateResponseStatusExpired:
		return true
	}
	return false
}

type WorkspaceAPIKeyListResponse struct {
	// The API key id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The API key creation date.
	DateCreated time.Time `json:"dateCreated" api:"required" format:"date-time"`
	// The API key last use date.
	DateLastUsed time.Time `json:"dateLastUsed" api:"required,nullable" format:"date-time"`
	// The API key last update date.
	DateUpdated time.Time `json:"dateUpdated" api:"required" format:"date-time"`
	// An obfuscated hint of the API key value. When a key is created or rotated this
	// also holds the full secret, for backward compatibility; prefer `secret`.
	SecureKey string `json:"secureKey" api:"required"`
	// The key's lifecycle state. `active`: the current secret authenticates.
	// `rotating`: the key was rotated and the previous secret still authenticates
	// until `previousKeyExpiresAt`. `expired`: `expiresAt` has passed and no secret
	// authenticates.
	Status WorkspaceAPIKeyListResponseStatus `json:"status" api:"required"`
	// When the key stops authenticating. `null` means the key never expires. Set when
	// the key is created or rotated, and must be in the future. When the request is
	// authenticated with an API key that expires, the result can't be later than that
	// key's expiry.
	ExpiresAt time.Time `json:"expiresAt" api:"nullable" format:"date-time"`
	// When the key was last rotated.
	LastRotatedAt time.Time `json:"lastRotatedAt" api:"nullable" format:"date-time"`
	// The API key name.
	Name string `json:"name" api:"nullable"`
	// While `status` is `rotating`, when the previous secret stops authenticating.
	PreviousKeyExpiresAt time.Time `json:"previousKeyExpiresAt" api:"nullable" format:"date-time"`
	// The full API key. Only present in the response that creates or rotates the key,
	// and never shown again.
	Secret string                          `json:"secret"`
	JSON   workspaceAPIKeyListResponseJSON `json:"-"`
}

// workspaceAPIKeyListResponseJSON contains the JSON metadata for the struct
// [WorkspaceAPIKeyListResponse]
type workspaceAPIKeyListResponseJSON struct {
	ID                   apijson.Field
	DateCreated          apijson.Field
	DateLastUsed         apijson.Field
	DateUpdated          apijson.Field
	SecureKey            apijson.Field
	Status               apijson.Field
	ExpiresAt            apijson.Field
	LastRotatedAt        apijson.Field
	Name                 apijson.Field
	PreviousKeyExpiresAt apijson.Field
	Secret               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *WorkspaceAPIKeyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceAPIKeyListResponseJSON) RawJSON() string {
	return r.raw
}

// The key's lifecycle state. `active`: the current secret authenticates.
// `rotating`: the key was rotated and the previous secret still authenticates
// until `previousKeyExpiresAt`. `expired`: `expiresAt` has passed and no secret
// authenticates.
type WorkspaceAPIKeyListResponseStatus string

const (
	WorkspaceAPIKeyListResponseStatusActive   WorkspaceAPIKeyListResponseStatus = "active"
	WorkspaceAPIKeyListResponseStatusRotating WorkspaceAPIKeyListResponseStatus = "rotating"
	WorkspaceAPIKeyListResponseStatusExpired  WorkspaceAPIKeyListResponseStatus = "expired"
)

func (r WorkspaceAPIKeyListResponseStatus) IsKnown() bool {
	switch r {
	case WorkspaceAPIKeyListResponseStatusActive, WorkspaceAPIKeyListResponseStatusRotating, WorkspaceAPIKeyListResponseStatusExpired:
		return true
	}
	return false
}

type WorkspaceAPIKeyRotateResponse struct {
	// The API key id.
	ID string `json:"id" api:"required" format:"uuid"`
	// The API key creation date.
	DateCreated time.Time `json:"dateCreated" api:"required" format:"date-time"`
	// The API key last use date.
	DateLastUsed time.Time `json:"dateLastUsed" api:"required,nullable" format:"date-time"`
	// The API key last update date.
	DateUpdated time.Time `json:"dateUpdated" api:"required" format:"date-time"`
	// An obfuscated hint of the API key value. When a key is created or rotated this
	// also holds the full secret, for backward compatibility; prefer `secret`.
	SecureKey string `json:"secureKey" api:"required"`
	// The key's lifecycle state. `active`: the current secret authenticates.
	// `rotating`: the key was rotated and the previous secret still authenticates
	// until `previousKeyExpiresAt`. `expired`: `expiresAt` has passed and no secret
	// authenticates.
	Status WorkspaceAPIKeyRotateResponseStatus `json:"status" api:"required"`
	// When the key stops authenticating. `null` means the key never expires. Set when
	// the key is created or rotated, and must be in the future. When the request is
	// authenticated with an API key that expires, the result can't be later than that
	// key's expiry.
	ExpiresAt time.Time `json:"expiresAt" api:"nullable" format:"date-time"`
	// When the key was last rotated.
	LastRotatedAt time.Time `json:"lastRotatedAt" api:"nullable" format:"date-time"`
	// The API key name.
	Name string `json:"name" api:"nullable"`
	// While `status` is `rotating`, when the previous secret stops authenticating.
	PreviousKeyExpiresAt time.Time `json:"previousKeyExpiresAt" api:"nullable" format:"date-time"`
	// The full API key. Only present in the response that creates or rotates the key,
	// and never shown again.
	Secret string                            `json:"secret"`
	JSON   workspaceAPIKeyRotateResponseJSON `json:"-"`
}

// workspaceAPIKeyRotateResponseJSON contains the JSON metadata for the struct
// [WorkspaceAPIKeyRotateResponse]
type workspaceAPIKeyRotateResponseJSON struct {
	ID                   apijson.Field
	DateCreated          apijson.Field
	DateLastUsed         apijson.Field
	DateUpdated          apijson.Field
	SecureKey            apijson.Field
	Status               apijson.Field
	ExpiresAt            apijson.Field
	LastRotatedAt        apijson.Field
	Name                 apijson.Field
	PreviousKeyExpiresAt apijson.Field
	Secret               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *WorkspaceAPIKeyRotateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceAPIKeyRotateResponseJSON) RawJSON() string {
	return r.raw
}

// The key's lifecycle state. `active`: the current secret authenticates.
// `rotating`: the key was rotated and the previous secret still authenticates
// until `previousKeyExpiresAt`. `expired`: `expiresAt` has passed and no secret
// authenticates.
type WorkspaceAPIKeyRotateResponseStatus string

const (
	WorkspaceAPIKeyRotateResponseStatusActive   WorkspaceAPIKeyRotateResponseStatus = "active"
	WorkspaceAPIKeyRotateResponseStatusRotating WorkspaceAPIKeyRotateResponseStatus = "rotating"
	WorkspaceAPIKeyRotateResponseStatusExpired  WorkspaceAPIKeyRotateResponseStatus = "expired"
)

func (r WorkspaceAPIKeyRotateResponseStatus) IsKnown() bool {
	switch r {
	case WorkspaceAPIKeyRotateResponseStatusActive, WorkspaceAPIKeyRotateResponseStatusRotating, WorkspaceAPIKeyRotateResponseStatusExpired:
		return true
	}
	return false
}

type WorkspaceAPIKeyNewParams struct {
	// When the key stops authenticating. `null` means the key never expires. Set when
	// the key is created or rotated, and must be in the future. When the request is
	// authenticated with an API key that expires, the result can't be later than that
	// key's expiry.
	ExpiresAt param.Field[time.Time] `json:"expiresAt" format:"date-time"`
	// The API key name.
	Name param.Field[string] `json:"name"`
}

func (r WorkspaceAPIKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkspaceAPIKeyUpdateParams struct {
	// The API key name.
	Name param.Field[string] `json:"name"`
}

func (r WorkspaceAPIKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkspaceAPIKeyRotateParams struct {
	// When the key stops authenticating. `null` means the key never expires. Set when
	// the key is created or rotated, and must be in the future. When the request is
	// authenticated with an API key that expires, the result can't be later than that
	// key's expiry.
	ExpiresAt param.Field[time.Time] `json:"expiresAt" format:"date-time"`
	// Hours the previous secret keeps authenticating.
	GracePeriodHours param.Field[int64] `json:"gracePeriodHours"`
}

func (r WorkspaceAPIKeyRotateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
