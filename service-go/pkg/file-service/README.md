# File service storage

The file service supports local files and S3-compatible storage (including MinIO).
S3 configuration and implementation come from `ncraft/go/pkg/ncraft/storage`;
armory adapts object metadata to the existing file API.

## Local storage

Existing `file.rootUrl` and `file.staticProvider.root` configurations still work.
An omitted `file.provider`, `local`, or `static` selects local storage. The service
creates the root and nested directories as needed. Relative file names cannot
escape the configured root through `..` or symbolic links.

## MinIO / S3

Replace `configs/file.yaml` with the contents of `examples/file-s3.yaml`, then set
credentials in the file service process environment:

```sh
export AWS_ACCESS_KEY_ID='your-access-key'
export AWS_SECRET_ACCESS_KEY='your-secret-key'
go run ./cmd/file-server
```

Use the S3 API endpoint (usually port 9000 for MinIO). Set `rootUrl` to the public
URL of the **file service**, not the MinIO endpoint. Upload responses continue to
return `/armory/file/v1/files/{name}` URLs; buckets can remain private.

`file.provider` accepts `s3` and `minio`. `file.s3Provider` is the shared ncraft
`storage.Config`:

| Field | Meaning |
| --- | --- |
| `vendor` | `s3` or `minio`; omitted means `s3` in the file service |
| `endpoint` | S3 API origin, e.g. `https://s3.example.com` or `http://localhost:9000` |
| `bucketName` | Bucket containing the files |
| `region` | Bucket region; omit to let the SDK discover it |
| `prefix` | Optional object key prefix, e.g. `files/tenant-a` |
| `accessKey`, `secretKey` | Configure both, or omit both to read AWS credential environment variables |
| `sessionToken` | Optional temporary credential token; environment equivalent: `AWS_SESSION_TOKEN` |
| `secure` | TLS for bare host:port endpoints; default false for existing ncraft configurations; URL scheme takes precedence |
| `createBucket` | `false` uses an existing bucket without bucket provisioning calls; omitted/true checks and creates it if needed |

The SDK uses Signature V4 and path-style bucket addressing. An explicit `region`
and `createBucket: false` allow use with credentials limited to object reads and
writes. If automatic creation is enabled, startup fails when bucket provisioning
fails. No public bucket policy or object ACL is applied.

## API behavior

- `CreateFile` stores the bytes and returns name, MIME type, size, service URL and
  upload time. Existing names are overwritten. Empty uploads remain invalid.
- `GetFile` reads bytes from the selected backend for gRPC callers.
- HTTP GET and HEAD use the generated `GetFile` endpoint and its middleware.
  `GetFile` returns a `BinaryFile` descriptor and binds a request-local HTTP
  writer to that exact result. The encoder opens the backend lazily and serves
  byte ranges, HEAD, ETag and modification-time conditions via `http.ServeContent`.
  Storage metadata and open errors are resolved during encoding; the HTTP result
  inside endpoint middleware contains only the name and URL, not file bytes.
- The reader is closed after encoding, including HEAD, 304 and 416 responses.
  Middleware rejection or replacement does not open a reader. A successful
  stream bypasses JSON envelopes; errors before streaming use the normal error
  encoder. Transfer errors after headers are written are logged without appending
  JSON. Endpoint timing excludes encoding; use HTTP-level metrics for downloads.
- Downloads use the main HTTP port. There is no separate Range listener.
  CORS preflight OPTIONS is handled by the host's existing CORS middleware.
- Nested names, spaces, Unicode, `#`, `?` and `%` are escaped in returned URLs.
  GET and HEAD bind `/armory/file/v1/files/{name:.+}`, allowing a non-empty name
  containing multiple path segments. Storage validation still rejects absolute
  paths, dot segments and attempts to escape the configured root.
- Batch uploads validate all inputs before writing, preserve input order, and
  stop at the first storage failure. Earlier successful writes are retained;
  batches are not transactional.
- Switching providers does not migrate existing files.

Example multipart upload and range download:

```sh
curl -H "Authorization: Bearer $TOKEN" -F 'file=@example.txt' \
  http://localhost:20171/armory/file/v1/files
curl -H 'Range: bytes=0-99' \
  http://localhost:20171/armory/file/v1/files/example.txt
```

## Development and validation

The matching ncraft storage changes are developed in the sibling `ncraft/go`
checkout. Use a local Go workspace containing both modules for joint development.
Before building outside that workspace, publish the matching ncraft module and
update the required version in `service-go/go.mod`.

From `armory/service-go`:

```sh
go test ./pkg/file-service/... ./internal/file-server ./cmd/file-server
```

From `ncraft/go`:

```sh
go test ./pkg/ncraft/storage/...
```

Tests exercise the real MinIO SDK against local HTTP S3 protocol fixtures,
including signing headers, prefix handling, bucket creation, metadata, missing
objects, denied access, cancellation and ranges. These fixtures do not replace
an integration test against a deployed MinIO server.

## Regenerating server code

`internal/file-server/run.go` is generated from Mojo's GoKit server template.
File-specific behavior lives in `handlers/server.go`, which generation preserves:

- `Start(nserver.Config) error` reports storage initialization failures before
  the generated server accepts requests. It does not start any listeners.

The file service has no independent background tasks and therefore does not
implement `ErrorSource` or `Shutdowner`. The generated host manages the HTTP/gRPC
listeners and gracefully drains requests on shutdown.

`file.mojo` declares GET and HEAD on `get_file`; the HTTP transport template
installs `nhttp.RequestToContext` and resolves `nhttp.BoundResponseWriter` before
ordinary response writers and JSON encoding. The handwritten `StreamBinaryFile`
writer lives in `handlers/stream_binary_file.go`. No custom download route or
`DownloadHandler` is needed.

Regeneration requires the matching Mojo HTTP template and ncraft response-binding
runtime. The explicit `@http.body` on the batch upload parameter preserves its
existing array request-body format. Do not add application-specific behavior to
generated `run.go` or `transport_http.go`.
