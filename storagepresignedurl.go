// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openlayer

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/openlayer-ai/openlayer-go/internal/apijson"
	"github.com/openlayer-ai/openlayer-go/internal/apiquery"
	"github.com/openlayer-ai/openlayer-go/internal/param"
	"github.com/openlayer-ai/openlayer-go/internal/requestconfig"
	"github.com/openlayer-ai/openlayer-go/option"
)

// StoragePresignedURLService contains methods and other services that help with
// interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStoragePresignedURLService] method instead.
type StoragePresignedURLService struct {
	Options []option.RequestOption
}

// NewStoragePresignedURLService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStoragePresignedURLService(opts ...option.RequestOption) (r *StoragePresignedURLService) {
	r = &StoragePresignedURLService{}
	r.Options = opts
	return
}

// Retrieve a presigned url to post storage artifacts.
func (r *StoragePresignedURLService) New(ctx context.Context, body StoragePresignedURLNewParams, opts ...option.RequestOption) (res *StoragePresignedURLNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "storage/presigned-url"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Exchange a `storageUri` for a short-lived presigned url you can download the
// object from.
//
// Use it to collect anything the platform stored on your behalf -- for example the
// archive a framework export leaves behind, whose `storageUri` comes back in the
// background task's `outputs`.
//
// The workspace is taken from the API key, so there is nothing else to send. The
// url is only issued for objects your workspace owns, and `404` covers both "no
// such object" and "not yours".
func (r *StoragePresignedURLService) Get(ctx context.Context, query StoragePresignedURLGetParams, opts ...option.RequestOption) (res *StoragePresignedURLGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "storage/presigned-url"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type StoragePresignedURLNewResponse struct {
	// The storage URI to send back to the backend after the upload was completed.
	StorageUri string `json:"storageUri" api:"required"`
	// The presigned url.
	URL string `json:"url" api:"required" format:"url"`
	// Fields to include in the body of the upload. Only needed by s3
	Fields interface{}                        `json:"fields"`
	JSON   storagePresignedURLNewResponseJSON `json:"-"`
}

// storagePresignedURLNewResponseJSON contains the JSON metadata for the struct
// [StoragePresignedURLNewResponse]
type storagePresignedURLNewResponseJSON struct {
	StorageUri  apijson.Field
	URL         apijson.Field
	Fields      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoragePresignedURLNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storagePresignedURLNewResponseJSON) RawJSON() string {
	return r.raw
}

type StoragePresignedURLGetResponse struct {
	// The presigned url. Short-lived -- download promptly.
	URL  string                             `json:"url" api:"required" format:"url"`
	JSON storagePresignedURLGetResponseJSON `json:"-"`
}

// storagePresignedURLGetResponseJSON contains the JSON metadata for the struct
// [StoragePresignedURLGetResponse]
type storagePresignedURLGetResponseJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoragePresignedURLGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storagePresignedURLGetResponseJSON) RawJSON() string {
	return r.raw
}

type StoragePresignedURLNewParams struct {
	// The name of the object.
	ObjectName param.Field[string] `query:"objectName" api:"required"`
}

// URLQuery serializes [StoragePresignedURLNewParams]'s query parameters as
// `url.Values`.
func (r StoragePresignedURLNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StoragePresignedURLGetParams struct {
	// The object's storage uri.
	StorageUri param.Field[string] `query:"storageUri" api:"required"`
}

// URLQuery serializes [StoragePresignedURLGetParams]'s query parameters as
// `url.Values`.
func (r StoragePresignedURLGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
