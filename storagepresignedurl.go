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
	return
}

type StoragePresignedURLNewResponse struct {
	// The storage URI to send back to the backend after the upload was completed.
	StorageUri string `json:"storageUri,required"`
	// The presigned url.
	URL string `json:"url,required" format:"url"`
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

type StoragePresignedURLNewParams struct {
	// The name of the object.
	ObjectName param.Field[string] `query:"objectName,required"`
}

// URLQuery serializes [StoragePresignedURLNewParams]'s query parameters as
// `url.Values`.
func (r StoragePresignedURLNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
