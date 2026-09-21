# Projects

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectNewResponse">ProjectNewResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectUpdateResponse">ProjectUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectListResponse">ProjectListResponse</a>

Methods:

- <code title="post /projects">client.Projects.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectNewParams">ProjectNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectNewResponse">ProjectNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /projects/{projectId}">client.Projects.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectUpdateParams">ProjectUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectUpdateResponse">ProjectUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /projects">client.Projects.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectListParams">ProjectListParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectListResponse">ProjectListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /projects/{projectId}">client.Projects.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Commits

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectCommitNewResponse">ProjectCommitNewResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectCommitListResponse">ProjectCommitListResponse</a>

Methods:

- <code title="post /projects/{projectId}/versions">client.Projects.Commits.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectCommitService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectCommitNewParams">ProjectCommitNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectCommitNewResponse">ProjectCommitNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /projects/{projectId}/versions">client.Projects.Commits.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectCommitService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectCommitListParams">ProjectCommitListParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectCommitListResponse">ProjectCommitListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## InferencePipelines

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectInferencePipelineNewResponse">ProjectInferencePipelineNewResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectInferencePipelineListResponse">ProjectInferencePipelineListResponse</a>

Methods:

- <code title="post /projects/{projectId}/inference-pipelines">client.Projects.InferencePipelines.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectInferencePipelineService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectInferencePipelineNewParams">ProjectInferencePipelineNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectInferencePipelineNewResponse">ProjectInferencePipelineNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /projects/{projectId}/inference-pipelines">client.Projects.InferencePipelines.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectInferencePipelineService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectInferencePipelineListParams">ProjectInferencePipelineListParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectInferencePipelineListResponse">ProjectInferencePipelineListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Tests

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestNewResponse">ProjectTestNewResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestUpdateResponse">ProjectTestUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestListResponse">ProjectTestListResponse</a>

Methods:

- <code title="post /projects/{projectId}/tests">client.Projects.Tests.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestNewParams">ProjectTestNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestNewResponse">ProjectTestNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /projects/{projectId}/tests">client.Projects.Tests.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestUpdateParams">ProjectTestUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestUpdateResponse">ProjectTestUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /projects/{projectId}/tests">client.Projects.Tests.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestListParams">ProjectTestListParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestListResponse">ProjectTestListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Workspaces

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#WorkspaceGetResponse">WorkspaceGetResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#WorkspaceUpdateResponse">WorkspaceUpdateResponse</a>

Methods:

- <code title="get /workspaces/{workspaceId}">client.Workspaces.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#WorkspaceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workspaceID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#WorkspaceGetResponse">WorkspaceGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /workspaces/{workspaceId}">client.Workspaces.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#WorkspaceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workspaceID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#WorkspaceUpdateParams">WorkspaceUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#WorkspaceUpdateResponse">WorkspaceUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Invites

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#WorkspaceInviteNewResponse">WorkspaceInviteNewResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#WorkspaceInviteListResponse">WorkspaceInviteListResponse</a>

Methods:

- <code title="post /workspaces/{workspaceId}/invites">client.Workspaces.Invites.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#WorkspaceInviteService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workspaceID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#WorkspaceInviteNewParams">WorkspaceInviteNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#WorkspaceInviteNewResponse">WorkspaceInviteNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /workspaces/{workspaceId}/invites">client.Workspaces.Invites.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#WorkspaceInviteService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workspaceID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#WorkspaceInviteListParams">WorkspaceInviteListParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#WorkspaceInviteListResponse">WorkspaceInviteListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## APIKeys

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#WorkspaceAPIKeyNewResponse">WorkspaceAPIKeyNewResponse</a>

Methods:

- <code title="post /workspaces/{workspaceId}/api-keys">client.Workspaces.APIKeys.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#WorkspaceAPIKeyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workspaceID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#WorkspaceAPIKeyNewParams">WorkspaceAPIKeyNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#WorkspaceAPIKeyNewResponse">WorkspaceAPIKeyNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Commits

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#CommitGetResponse">CommitGetResponse</a>

Methods:

- <code title="get /versions/{projectVersionId}">client.Commits.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#CommitService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectVersionID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#CommitGetResponse">CommitGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## TestResults

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#CommitTestResultListResponse">CommitTestResultListResponse</a>

Methods:

- <code title="get /versions/{projectVersionId}/results">client.Commits.TestResults.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#CommitTestResultService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#CommitTestResultListParams">CommitTestResultListParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#CommitTestResultListResponse">CommitTestResultListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# InferencePipelines

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineGetResponse">InferencePipelineGetResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineUpdateResponse">InferencePipelineUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineGetSessionsResponse">InferencePipelineGetSessionsResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineGetUsersResponse">InferencePipelineGetUsersResponse</a>

Methods:

- <code title="get /inference-pipelines/{inferencePipelineId}">client.InferencePipelines.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inferencePipelineID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineGetParams">InferencePipelineGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineGetResponse">InferencePipelineGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /inference-pipelines/{inferencePipelineId}">client.InferencePipelines.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inferencePipelineID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineUpdateParams">InferencePipelineUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineUpdateResponse">InferencePipelineUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /inference-pipelines/{inferencePipelineId}">client.InferencePipelines.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inferencePipelineID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /inference-pipelines/{inferencePipelineId}/sessions">client.InferencePipelines.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineService.GetSessions">GetSessions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inferencePipelineID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineGetSessionsParams">InferencePipelineGetSessionsParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineGetSessionsResponse">InferencePipelineGetSessionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /inference-pipelines/{inferencePipelineId}/users">client.InferencePipelines.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineService.GetUsers">GetUsers</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inferencePipelineID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineGetUsersParams">InferencePipelineGetUsersParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineGetUsersResponse">InferencePipelineGetUsersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Data

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineDataStreamResponse">InferencePipelineDataStreamResponse</a>

Methods:

- <code title="post /inference-pipelines/{inferencePipelineId}/data-stream">client.InferencePipelines.Data.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineDataService.Stream">Stream</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inferencePipelineID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineDataStreamParams">InferencePipelineDataStreamParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineDataStreamResponse">InferencePipelineDataStreamResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Rows

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineRowGetResponse">InferencePipelineRowGetResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineRowUpdateResponse">InferencePipelineRowUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineRowListResponse">InferencePipelineRowListResponse</a>

Methods:

- <code title="get /inference-pipelines/{inferencePipelineId}/rows/{inferenceId}">client.InferencePipelines.Rows.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineRowService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inferencePipelineID <a href="https://pkg.go.dev/builtin#string">string</a>, inferenceID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineRowGetResponse">InferencePipelineRowGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /inference-pipelines/{inferencePipelineId}/rows">client.InferencePipelines.Rows.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineRowService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inferencePipelineID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineRowUpdateParams">InferencePipelineRowUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineRowUpdateResponse">InferencePipelineRowUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /inference-pipelines/{inferencePipelineId}/rows">client.InferencePipelines.Rows.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineRowService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inferencePipelineID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineRowListParams">InferencePipelineRowListParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineRowListResponse">InferencePipelineRowListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /inference-pipelines/{inferencePipelineId}/rows/{inferenceId}">client.InferencePipelines.Rows.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineRowService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inferencePipelineID <a href="https://pkg.go.dev/builtin#string">string</a>, inferenceID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## TestResults

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineTestResultListResponse">InferencePipelineTestResultListResponse</a>

Methods:

- <code title="get /inference-pipelines/{inferencePipelineId}/results">client.InferencePipelines.TestResults.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineTestResultService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inferencePipelineID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineTestResultListParams">InferencePipelineTestResultListParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineTestResultListResponse">InferencePipelineTestResultListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Storage

## PresignedURL

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#StoragePresignedURLNewResponse">StoragePresignedURLNewResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#StoragePresignedURLGetResponse">StoragePresignedURLGetResponse</a>

Methods:

- <code title="post /storage/presigned-url">client.Storage.PresignedURL.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#StoragePresignedURLService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#StoragePresignedURLNewParams">StoragePresignedURLNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#StoragePresignedURLNewResponse">StoragePresignedURLNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /storage/presigned-url">client.Storage.PresignedURL.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#StoragePresignedURLService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#StoragePresignedURLGetParams">StoragePresignedURLGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#StoragePresignedURLGetResponse">StoragePresignedURLGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Tests

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#TestEvaluateResponse">TestEvaluateResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#TestListResultsResponse">TestListResultsResponse</a>

Methods:

- <code title="post /tests/{testId}/evaluate">client.Tests.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#TestService.Evaluate">Evaluate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, testID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#TestEvaluateParams">TestEvaluateParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#TestEvaluateResponse">TestEvaluateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /tests/{testId}/results">client.Tests.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#TestService.ListResults">ListResults</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, testID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#TestListResultsParams">TestListResultsParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#TestListResultsResponse">TestListResultsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# BackgroundTasks

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#BackgroundTaskGetResponse">BackgroundTaskGetResponse</a>

Methods:

- <code title="get /background-tasks/{taskId}">client.BackgroundTasks.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#BackgroundTaskService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, taskID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#BackgroundTaskGetResponse">BackgroundTaskGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Governance

## Frameworks

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkNewResponse">GovernanceFrameworkNewResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkGetResponse">GovernanceFrameworkGetResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkUpdateResponse">GovernanceFrameworkUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkListResponse">GovernanceFrameworkListResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkExportResponse">GovernanceFrameworkExportResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkListProjectRuleStatsResponse">GovernanceFrameworkListProjectRuleStatsResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkListProjectsResponse">GovernanceFrameworkListProjectsResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkListRulesResponse">GovernanceFrameworkListRulesResponse</a>

Methods:

- <code title="post /workspaces/{workspaceId}/frameworks">client.Governance.Frameworks.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workspaceID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkNewParams">GovernanceFrameworkNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkNewResponse">GovernanceFrameworkNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /frameworks/{frameworkId}">client.Governance.Frameworks.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, frameworkID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkGetResponse">GovernanceFrameworkGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /frameworks/{frameworkId}">client.Governance.Frameworks.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, frameworkID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkUpdateParams">GovernanceFrameworkUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkUpdateResponse">GovernanceFrameworkUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /workspaces/{workspaceId}/frameworks">client.Governance.Frameworks.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workspaceID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkListParams">GovernanceFrameworkListParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkListResponse">GovernanceFrameworkListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /frameworks/{frameworkId}/export">client.Governance.Frameworks.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkService.Export">Export</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, frameworkID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkExportParams">GovernanceFrameworkExportParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkExportResponse">GovernanceFrameworkExportResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /frameworks/{frameworkId}/project-rule-stats">client.Governance.Frameworks.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkService.ListProjectRuleStats">ListProjectRuleStats</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, frameworkID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkListProjectRuleStatsParams">GovernanceFrameworkListProjectRuleStatsParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkListProjectRuleStatsResponse">GovernanceFrameworkListProjectRuleStatsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /frameworks/{frameworkId}/projects">client.Governance.Frameworks.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkService.ListProjects">ListProjects</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, frameworkID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkListProjectsParams">GovernanceFrameworkListProjectsParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkListProjectsResponse">GovernanceFrameworkListProjectsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /frameworks/{frameworkId}/rules">client.Governance.Frameworks.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkService.ListRules">ListRules</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, frameworkID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkListRulesParams">GovernanceFrameworkListRulesParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkListRulesResponse">GovernanceFrameworkListRulesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Documents

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkDocumentGetResponse">GovernanceFrameworkDocumentGetResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkDocumentListResponse">GovernanceFrameworkDocumentListResponse</a>

Methods:

- <code title="get /frameworks/{frameworkId}/documents/{documentId}">client.Governance.Frameworks.Documents.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkDocumentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, frameworkID <a href="https://pkg.go.dev/builtin#string">string</a>, documentID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkDocumentGetResponse">GovernanceFrameworkDocumentGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /frameworks/{frameworkId}/documents">client.Governance.Frameworks.Documents.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkDocumentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, frameworkID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkDocumentListParams">GovernanceFrameworkDocumentListParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkDocumentListResponse">GovernanceFrameworkDocumentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Sections

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkSectionListRulesResponse">GovernanceFrameworkSectionListRulesResponse</a>

Methods:

- <code title="get /frameworks/{frameworkId}/sections/{sectionId}/rules">client.Governance.Frameworks.Sections.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkSectionService.ListRules">ListRules</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, frameworkID <a href="https://pkg.go.dev/builtin#string">string</a>, sectionID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkSectionListRulesParams">GovernanceFrameworkSectionListRulesParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkSectionListRulesResponse">GovernanceFrameworkSectionListRulesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Subsections

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkSubsectionListRulesResponse">GovernanceFrameworkSubsectionListRulesResponse</a>

Methods:

- <code title="get /frameworks/{frameworkId}/subsections/{subsectionId}/rules">client.Governance.Frameworks.Subsections.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkSubsectionService.ListRules">ListRules</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, frameworkID <a href="https://pkg.go.dev/builtin#string">string</a>, subsectionID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkSubsectionListRulesParams">GovernanceFrameworkSubsectionListRulesParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceFrameworkSubsectionListRulesResponse">GovernanceFrameworkSubsectionListRulesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Rules

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleNewResponse">GovernanceRuleNewResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleGetResponse">GovernanceRuleGetResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleUpdateResponse">GovernanceRuleUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleListResponse">GovernanceRuleListResponse</a>

Methods:

- <code title="post /workspaces/{workspaceId}/rules">client.Governance.Rules.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workspaceID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleNewParams">GovernanceRuleNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleNewResponse">GovernanceRuleNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /rules/{ruleId}">client.Governance.Rules.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ruleID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleGetResponse">GovernanceRuleGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /rules/{ruleId}">client.Governance.Rules.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ruleID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleUpdateParams">GovernanceRuleUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleUpdateResponse">GovernanceRuleUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /workspaces/{workspaceId}/rules">client.Governance.Rules.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workspaceID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleListParams">GovernanceRuleListParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleListResponse">GovernanceRuleListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /rules/{ruleId}">client.Governance.Rules.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ruleID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## RuleResults

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleResultGetResponse">GovernanceRuleResultGetResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleResultUpdateResponse">GovernanceRuleResultUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleResultListResponse">GovernanceRuleResultListResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleResultNewEvidenceResponse">GovernanceRuleResultNewEvidenceResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleResultListEvidenceResponse">GovernanceRuleResultListEvidenceResponse</a>

Methods:

- <code title="get /rule-results/{ruleResultId}">client.Governance.RuleResults.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleResultService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ruleResultID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleResultGetResponse">GovernanceRuleResultGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /rule-results/{ruleResultId}">client.Governance.RuleResults.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleResultService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ruleResultID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleResultUpdateParams">GovernanceRuleResultUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleResultUpdateResponse">GovernanceRuleResultUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /workspaces/{workspaceId}/rule-results">client.Governance.RuleResults.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleResultService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workspaceID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleResultListParams">GovernanceRuleResultListParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleResultListResponse">GovernanceRuleResultListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /rule-results/{ruleResultId}/evidence">client.Governance.RuleResults.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleResultService.NewEvidence">NewEvidence</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ruleResultID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleResultNewEvidenceParams">GovernanceRuleResultNewEvidenceParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleResultNewEvidenceResponse">GovernanceRuleResultNewEvidenceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /rule-results/{ruleResultId}/evidence">client.Governance.RuleResults.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleResultService.ListEvidence">ListEvidence</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ruleResultID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleResultListEvidenceParams">GovernanceRuleResultListEvidenceParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleResultListEvidenceResponse">GovernanceRuleResultListEvidenceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## RuleStats

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleStatGetResponse">GovernanceRuleStatGetResponse</a>

Methods:

- <code title="get /workspaces/{workspaceId}/rule-stats">client.Governance.RuleStats.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleStatService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workspaceID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleStatGetParams">GovernanceRuleStatGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleStatGetResponse">GovernanceRuleStatGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## RuleTags

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleTagListResponse">GovernanceRuleTagListResponse</a>

Methods:

- <code title="get /workspaces/{workspaceId}/rule-tags">client.Governance.RuleTags.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleTagService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workspaceID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleTagListParams">GovernanceRuleTagListParams</a>) (\*<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#GovernanceRuleTagListResponse">GovernanceRuleTagListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
