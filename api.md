# Projects

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectNewResponse">ProjectNewResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectListResponse">ProjectListResponse</a>

Methods:

- <code title="post /projects">client.Projects.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectNewParams">ProjectNewParams</a>) (<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectNewResponse">ProjectNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /projects">client.Projects.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectListParams">ProjectListParams</a>) (<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectListResponse">ProjectListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Commits

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectCommitNewResponse">ProjectCommitNewResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectCommitListResponse">ProjectCommitListResponse</a>

Methods:

- <code title="post /projects/{projectId}/versions">client.Projects.Commits.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectCommitService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectCommitNewParams">ProjectCommitNewParams</a>) (<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectCommitNewResponse">ProjectCommitNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /projects/{projectId}/versions">client.Projects.Commits.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectCommitService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectCommitListParams">ProjectCommitListParams</a>) (<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectCommitListResponse">ProjectCommitListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## InferencePipelines

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectInferencePipelineNewResponse">ProjectInferencePipelineNewResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectInferencePipelineListResponse">ProjectInferencePipelineListResponse</a>

Methods:

- <code title="post /projects/{projectId}/inference-pipelines">client.Projects.InferencePipelines.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectInferencePipelineService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectInferencePipelineNewParams">ProjectInferencePipelineNewParams</a>) (<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectInferencePipelineNewResponse">ProjectInferencePipelineNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /projects/{projectId}/inference-pipelines">client.Projects.InferencePipelines.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectInferencePipelineService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectInferencePipelineListParams">ProjectInferencePipelineListParams</a>) (<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectInferencePipelineListResponse">ProjectInferencePipelineListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Tests

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestNewResponse">ProjectTestNewResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestUpdateResponse">ProjectTestUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestListResponse">ProjectTestListResponse</a>

Methods:

- <code title="post /projects/{projectId}/tests">client.Projects.Tests.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestNewParams">ProjectTestNewParams</a>) (<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestNewResponse">ProjectTestNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /projects/{projectId}/tests">client.Projects.Tests.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestUpdateParams">ProjectTestUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestUpdateResponse">ProjectTestUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /projects/{projectId}/tests">client.Projects.Tests.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestListParams">ProjectTestListParams</a>) (<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#ProjectTestListResponse">ProjectTestListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Commits

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#CommitGetResponse">CommitGetResponse</a>

Methods:

- <code title="get /versions/{projectVersionId}">client.Commits.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#CommitService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectVersionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#CommitGetResponse">CommitGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## TestResults

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#CommitTestResultListResponse">CommitTestResultListResponse</a>

Methods:

- <code title="get /versions/{projectVersionId}/results">client.Commits.TestResults.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#CommitTestResultService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#CommitTestResultListParams">CommitTestResultListParams</a>) (<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#CommitTestResultListResponse">CommitTestResultListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# InferencePipelines

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineGetResponse">InferencePipelineGetResponse</a>
- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineUpdateResponse">InferencePipelineUpdateResponse</a>

Methods:

- <code title="get /inference-pipelines/{inferencePipelineId}">client.InferencePipelines.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inferencePipelineID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineGetParams">InferencePipelineGetParams</a>) (<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineGetResponse">InferencePipelineGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /inference-pipelines/{inferencePipelineId}">client.InferencePipelines.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inferencePipelineID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineUpdateParams">InferencePipelineUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineUpdateResponse">InferencePipelineUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /inference-pipelines/{inferencePipelineId}">client.InferencePipelines.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inferencePipelineID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Data

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineDataStreamResponse">InferencePipelineDataStreamResponse</a>

Methods:

- <code title="post /inference-pipelines/{inferencePipelineId}/data-stream">client.InferencePipelines.Data.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineDataService.Stream">Stream</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inferencePipelineID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineDataStreamParams">InferencePipelineDataStreamParams</a>) (<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineDataStreamResponse">InferencePipelineDataStreamResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Rows

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineRowUpdateResponse">InferencePipelineRowUpdateResponse</a>

Methods:

- <code title="put /inference-pipelines/{inferencePipelineId}/rows">client.InferencePipelines.Rows.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineRowService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inferencePipelineID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineRowUpdateParams">InferencePipelineRowUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineRowUpdateResponse">InferencePipelineRowUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## TestResults

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineTestResultListResponse">InferencePipelineTestResultListResponse</a>

Methods:

- <code title="get /inference-pipelines/{inferencePipelineId}/results">client.InferencePipelines.TestResults.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineTestResultService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inferencePipelineID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineTestResultListParams">InferencePipelineTestResultListParams</a>) (<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#InferencePipelineTestResultListResponse">InferencePipelineTestResultListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Storage

## PresignedURL

Response Types:

- <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#StoragePresignedURLNewResponse">StoragePresignedURLNewResponse</a>

Methods:

- <code title="post /storage/presigned-url">client.Storage.PresignedURL.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#StoragePresignedURLService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#StoragePresignedURLNewParams">StoragePresignedURLNewParams</a>) (<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/openlayer-ai/openlayer-go#StoragePresignedURLNewResponse">StoragePresignedURLNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
