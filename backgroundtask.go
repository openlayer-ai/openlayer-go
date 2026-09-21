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
	"github.com/openlayer-ai/openlayer-go/internal/requestconfig"
	"github.com/openlayer-ai/openlayer-go/option"
)

// BackgroundTaskService contains methods and other services that help with
// interacting with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBackgroundTaskService] method instead.
type BackgroundTaskService struct {
	Options []option.RequestOption
}

// NewBackgroundTaskService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBackgroundTaskService(opts ...option.RequestOption) (r *BackgroundTaskService) {
	r = &BackgroundTaskService{}
	r.Options = opts
	return
}

// Retrieve a background task's status, progress and results.
//
// Endpoints that cannot answer within one request queue a task and hand back its
// id -- for example `POST /frameworks/{frameworkId}/export`. Poll this endpoint
// until `complete` is `true`, then read what the task produced from `outputs`.
func (r *BackgroundTaskService) Get(ctx context.Context, taskID string, opts ...option.RequestOption) (res *BackgroundTaskGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if taskID == "" {
		err = errors.New("missing required taskId parameter")
		return nil, err
	}
	path := fmt.Sprintf("background-tasks/%s", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type BackgroundTaskGetResponse struct {
	// The background task id.
	ID string `json:"id" api:"required" format:"uuid"`
	// Whether the task has finished. Check this before reading `outputs`.
	Complete bool `json:"complete" api:"required"`
	// When the task was queued.
	DateCreated time.Time `json:"dateCreated" api:"required" format:"date-time"`
	// When the task last reported progress.
	DateUpdated time.Time `json:"dateUpdated" api:"required" format:"date-time"`
	// The task's internal name, including the arguments it was queued with.
	Name string `json:"name" api:"required"`
	// How far along the task is, from 0 to 100.
	Progress float64 `json:"progress" api:"required"`
	// Why the task failed, or `null` if it has not failed.
	Error string `json:"error" api:"nullable"`
	// Whatever the task produced, keyed by name. `null` until the task completes. A
	// framework export returns `storageUri` -- pass it to `GET /storage/presigned-url`
	// to download the archive -- along with `filename`, `controlCount`,
	// `evidenceCount` and `missingEvidenceCount`.
	Outputs interface{}                   `json:"outputs" api:"nullable"`
	JSON    backgroundTaskGetResponseJSON `json:"-"`
}

// backgroundTaskGetResponseJSON contains the JSON metadata for the struct
// [BackgroundTaskGetResponse]
type backgroundTaskGetResponseJSON struct {
	ID          apijson.Field
	Complete    apijson.Field
	DateCreated apijson.Field
	DateUpdated apijson.Field
	Name        apijson.Field
	Progress    apijson.Field
	Error       apijson.Field
	Outputs     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BackgroundTaskGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r backgroundTaskGetResponseJSON) RawJSON() string {
	return r.raw
}
