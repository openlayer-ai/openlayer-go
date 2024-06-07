// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomopenlayeraiopenlayergo

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/openlayer-go/internal/apijson"
	"github.com/stainless-sdks/openlayer-go/internal/apiquery"
	"github.com/stainless-sdks/openlayer-go/internal/param"
	"github.com/stainless-sdks/openlayer-go/internal/requestconfig"
	"github.com/stainless-sdks/openlayer-go/option"
)

// ProjectService contains methods and other services that help with interacting
// with the openlayer API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectService] method instead.
type ProjectService struct {
	Options            []option.RequestOption
	Commits            *ProjectCommitService
	InferencePipelines *ProjectInferencePipelineService
}

// NewProjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProjectService(opts ...option.RequestOption) (r *ProjectService) {
	r = &ProjectService{}
	r.Options = opts
	r.Commits = NewProjectCommitService(opts...)
	r.InferencePipelines = NewProjectInferencePipelineService(opts...)
	return
}

// List the projects in a user's workspace.
func (r *ProjectService) List(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) (res *ProjectListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ProjectListResponse struct {
	Meta  ProjectListResponseMeta   `json:"_meta,required"`
	Items []ProjectListResponseItem `json:"items,required"`
	JSON  projectListResponseJSON   `json:"-"`
}

// projectListResponseJSON contains the JSON metadata for the struct
// [ProjectListResponse]
type projectListResponseJSON struct {
	Meta        apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectListResponseMeta struct {
	// The current page.
	Page int64 `json:"page,required"`
	// The number of items per page.
	PerPage int64 `json:"perPage,required"`
	// The total number of items.
	TotalItems int64 `json:"totalItems,required"`
	// The total number of pages.
	TotalPages int64                       `json:"totalPages,required"`
	JSON       projectListResponseMetaJSON `json:"-"`
}

// projectListResponseMetaJSON contains the JSON metadata for the struct
// [ProjectListResponseMeta]
type projectListResponseMetaJSON struct {
	Page        apijson.Field
	PerPage     apijson.Field
	TotalItems  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseMetaJSON) RawJSON() string {
	return r.raw
}

type ProjectListResponseItem struct {
	// The project id.
	ID string `json:"id,required" format:"uuid"`
	// The project creator id.
	CreatorID string `json:"creatorId,required,nullable" format:"uuid"`
	// The project creation date.
	DateCreated time.Time `json:"dateCreated,required" format:"date-time"`
	// The project last updated date.
	DateUpdated time.Time `json:"dateUpdated,required" format:"date-time"`
	// The number of tests in the development mode of the project.
	DevelopmentGoalCount int64 `json:"developmentGoalCount,required"`
	// The total number of tests in the project.
	GoalCount int64 `json:"goalCount,required"`
	// The number of inference pipelines in the project.
	InferencePipelineCount int64 `json:"inferencePipelineCount,required"`
	// Links to the project.
	Links ProjectListResponseItemsLinks `json:"links,required"`
	// The number of tests in the monitoring mode of the project.
	MonitoringGoalCount int64 `json:"monitoringGoalCount,required"`
	// The project name.
	Name string `json:"name,required"`
	// Whether the project is a sample project or a user-created project.
	Sample bool `json:"sample,required"`
	// The source of the project.
	Source ProjectListResponseItemsSource `json:"source,required,nullable"`
	// The task type of the project.
	TaskType ProjectListResponseItemsTaskType `json:"taskType,required"`
	// The number of versions (commits) in the project.
	VersionCount int64 `json:"versionCount,required"`
	// The workspace id.
	WorkspaceID string `json:"workspaceId,required,nullable" format:"uuid"`
	// The project description.
	Description string                          `json:"description,nullable"`
	GitRepo     ProjectListResponseItemsGitRepo `json:"gitRepo,nullable"`
	// The slack channel id connected to the project.
	SlackChannelID string `json:"slackChannelId,nullable"`
	// The slack channel connected to the project.
	SlackChannelName string `json:"slackChannelName,nullable"`
	// Whether slack channel notifications are enabled for the project.
	SlackChannelNotificationsEnabled bool `json:"slackChannelNotificationsEnabled"`
	// The number of unread notifications in the project.
	UnreadNotificationCount int64                       `json:"unreadNotificationCount"`
	JSON                    projectListResponseItemJSON `json:"-"`
}

// projectListResponseItemJSON contains the JSON metadata for the struct
// [ProjectListResponseItem]
type projectListResponseItemJSON struct {
	ID                               apijson.Field
	CreatorID                        apijson.Field
	DateCreated                      apijson.Field
	DateUpdated                      apijson.Field
	DevelopmentGoalCount             apijson.Field
	GoalCount                        apijson.Field
	InferencePipelineCount           apijson.Field
	Links                            apijson.Field
	MonitoringGoalCount              apijson.Field
	Name                             apijson.Field
	Sample                           apijson.Field
	Source                           apijson.Field
	TaskType                         apijson.Field
	VersionCount                     apijson.Field
	WorkspaceID                      apijson.Field
	Description                      apijson.Field
	GitRepo                          apijson.Field
	SlackChannelID                   apijson.Field
	SlackChannelName                 apijson.Field
	SlackChannelNotificationsEnabled apijson.Field
	UnreadNotificationCount          apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r *ProjectListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseItemJSON) RawJSON() string {
	return r.raw
}

// Links to the project.
type ProjectListResponseItemsLinks struct {
	App  string                            `json:"app,required"`
	JSON projectListResponseItemsLinksJSON `json:"-"`
}

// projectListResponseItemsLinksJSON contains the JSON metadata for the struct
// [ProjectListResponseItemsLinks]
type projectListResponseItemsLinksJSON struct {
	App         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseItemsLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseItemsLinksJSON) RawJSON() string {
	return r.raw
}

// The source of the project.
type ProjectListResponseItemsSource string

const (
	ProjectListResponseItemsSourceWeb  ProjectListResponseItemsSource = "web"
	ProjectListResponseItemsSourceAPI  ProjectListResponseItemsSource = "api"
	ProjectListResponseItemsSourceNull ProjectListResponseItemsSource = "null"
)

func (r ProjectListResponseItemsSource) IsKnown() bool {
	switch r {
	case ProjectListResponseItemsSourceWeb, ProjectListResponseItemsSourceAPI, ProjectListResponseItemsSourceNull:
		return true
	}
	return false
}

// The task type of the project.
type ProjectListResponseItemsTaskType string

const (
	ProjectListResponseItemsTaskTypeLlmBase               ProjectListResponseItemsTaskType = "llm-base"
	ProjectListResponseItemsTaskTypeTabularClassification ProjectListResponseItemsTaskType = "tabular-classification"
	ProjectListResponseItemsTaskTypeTabularRegression     ProjectListResponseItemsTaskType = "tabular-regression"
	ProjectListResponseItemsTaskTypeTextClassification    ProjectListResponseItemsTaskType = "text-classification"
)

func (r ProjectListResponseItemsTaskType) IsKnown() bool {
	switch r {
	case ProjectListResponseItemsTaskTypeLlmBase, ProjectListResponseItemsTaskTypeTabularClassification, ProjectListResponseItemsTaskTypeTabularRegression, ProjectListResponseItemsTaskTypeTextClassification:
		return true
	}
	return false
}

type ProjectListResponseItemsGitRepo struct {
	ID            string                              `json:"id,required" format:"uuid"`
	DateConnected time.Time                           `json:"dateConnected,required" format:"date-time"`
	DateUpdated   time.Time                           `json:"dateUpdated,required" format:"date-time"`
	GitAccountID  string                              `json:"gitAccountId,required" format:"uuid"`
	GitID         int64                               `json:"gitId,required"`
	Name          string                              `json:"name,required"`
	Private       bool                                `json:"private,required"`
	ProjectID     string                              `json:"projectId,required" format:"uuid"`
	Slug          string                              `json:"slug,required"`
	URL           string                              `json:"url,required" format:"url"`
	Branch        string                              `json:"branch"`
	RootDir       string                              `json:"rootDir"`
	JSON          projectListResponseItemsGitRepoJSON `json:"-"`
}

// projectListResponseItemsGitRepoJSON contains the JSON metadata for the struct
// [ProjectListResponseItemsGitRepo]
type projectListResponseItemsGitRepoJSON struct {
	ID            apijson.Field
	DateConnected apijson.Field
	DateUpdated   apijson.Field
	GitAccountID  apijson.Field
	GitID         apijson.Field
	Name          apijson.Field
	Private       apijson.Field
	ProjectID     apijson.Field
	Slug          apijson.Field
	URL           apijson.Field
	Branch        apijson.Field
	RootDir       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectListResponseItemsGitRepo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseItemsGitRepoJSON) RawJSON() string {
	return r.raw
}

type ProjectListParams struct {
	// Filter list of items by project name.
	Name param.Field[string] `query:"name"`
	// The page to return in a paginated query.
	Page param.Field[int64] `query:"page"`
	// Maximum number of items to return per page.
	PerPage param.Field[int64] `query:"perPage"`
	// Filter list of items by task type.
	TaskType param.Field[ProjectListParamsTaskType] `query:"taskType"`
}

// URLQuery serializes [ProjectListParams]'s query parameters as `url.Values`.
func (r ProjectListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter list of items by task type.
type ProjectListParamsTaskType string

const (
	ProjectListParamsTaskTypeLlmBase               ProjectListParamsTaskType = "llm-base"
	ProjectListParamsTaskTypeTabularClassification ProjectListParamsTaskType = "tabular-classification"
	ProjectListParamsTaskTypeTabularRegression     ProjectListParamsTaskType = "tabular-regression"
	ProjectListParamsTaskTypeTextClassification    ProjectListParamsTaskType = "text-classification"
)

func (r ProjectListParamsTaskType) IsKnown() bool {
	switch r {
	case ProjectListParamsTaskTypeLlmBase, ProjectListParamsTaskTypeTabularClassification, ProjectListParamsTaskTypeTabularRegression, ProjectListParamsTaskTypeTextClassification:
		return true
	}
	return false
}
