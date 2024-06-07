# Projects

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#ProjectListResponse">ProjectListResponse</a>

Methods:

- <code title="get /projects">client.Projects.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#ProjectService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#ProjectListParams">ProjectListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#ProjectListResponse">ProjectListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Commits

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#ProjectCommitListResponse">ProjectCommitListResponse</a>

Methods:

- <code title="get /projects/{id}/versions">client.Projects.Commits.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#ProjectCommitService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#ProjectCommitListParams">ProjectCommitListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#ProjectCommitListResponse">ProjectCommitListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## InferencePipelines

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#ProjectInferencePipelineListResponse">ProjectInferencePipelineListResponse</a>

Methods:

- <code title="get /projects/{id}/inference-pipelines">client.Projects.InferencePipelines.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#ProjectInferencePipelineService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#ProjectInferencePipelineListParams">ProjectInferencePipelineListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#ProjectInferencePipelineListResponse">ProjectInferencePipelineListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Commits

## TestResults

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#CommitTestResultListResponse">CommitTestResultListResponse</a>

Methods:

- <code title="get /versions/{id}/results">client.Commits.TestResults.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#CommitTestResultService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#CommitTestResultListParams">CommitTestResultListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#CommitTestResultListResponse">CommitTestResultListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# InferencePipelines

## Data

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#InferencePipelineDataStreamResponse">InferencePipelineDataStreamResponse</a>

Methods:

- <code title="post /inference-pipelines/{id}/data-stream">client.InferencePipelines.Data.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#InferencePipelineDataService.Stream">Stream</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#InferencePipelineDataStreamParams">InferencePipelineDataStreamParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#InferencePipelineDataStreamResponse">InferencePipelineDataStreamResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## TestResults

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#InferencePipelineTestResultListResponse">InferencePipelineTestResultListResponse</a>

Methods:

- <code title="get /inference-pipelines/{id}/results">client.InferencePipelines.TestResults.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#InferencePipelineTestResultService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#InferencePipelineTestResultListParams">InferencePipelineTestResultListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go">openlayer</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/openlayer-go#InferencePipelineTestResultListResponse">InferencePipelineTestResultListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
