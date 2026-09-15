# Predefined database queries

`DbQuery` supports both configuration and database storage. Execution resolves a
query by database and name in this order:

1. An entry in `dbQuery.queries` with the matching `database` and `name`.
2. An entry with the matching `name` and an empty `database` (shared configuration).
3. A persisted query with the matching `database` and `name`.

A matching configuration entry is authoritative; an execution error does not
fall back to the stored definition. Configuration still loads at service startup.
Database definitions are read on each execution, so updates and deletions take
effect immediately. Definitions live in the service metadata database (`db`),
while their SQL runs against the selected connection in `dataDB.dbs`.

## API

All paths below start with `/armory/unitable/v1/databases/{database}`.

| Method | Path | Operation |
| --- | --- | --- |
| POST | `/queries` | Create a persisted definition; duplicates return AlreadyExists. |
| PUT | `/queries/{id}` | Replace a persisted definition after validation. |
| GET | `/queries/{id}` | Read a persisted definition, or the effective definition with `effective=true`. |
| GET | `/queries` | List persisted definitions, or the merged effective view with `effective=true`. |
| DELETE | `/queries/{id}` | Delete a persisted definition. |
| POST | `/queries/{id}:run` | Execute the effective definition with a JSON parameter object. |

`id` accepts either the query name or `database.name`. New records use
`database.name` as their ID. Names may contain letters, digits, underscores and
hyphens, must start with a letter, and have a maximum length of 128 characters.
The path selects the database; supplied body identifiers must agree with it.
Names and databases cannot be changed by update.

Get and List default to persisted records, including records currently overridden
by configuration. Set the optional `effective` boolean to `true` to read the same
definitions used by execution:

```http
GET /armory/unitable/v1/databases/main/queries/sales-summary?effective=true
GET /armory/unitable/v1/databases/main/queries?effective=true&pageSize=20
```

Go/gRPC callers set `GetDbQueryRequest.Effective` or
`ListDbQueriesRequest.Effective`. Omitting the flag or using `false` preserves
the existing database-only behavior.

Effective List merges by query name within the selected database, applying the
same precedence as execution. It includes configuration-only entries and removes
shadowed database records **before** filtering, sorting, projection and pagination.
`totalCount` and `nextPageToken` describe the merged, filtered view. Normal filter,
order, field-mask and unique parameters remain available. A name tie-breaker
provides stable pagination when distinct projections are not requested.

Configuration-backed results are copies: their `id` is `database.name`, their
`database` is the selected database, and an unspecified JSON style is returned
as `lowerCamel`, matching execution. Nested configuration data is not mutated.
Effective reads never persist configuration entries. List still requires access
to the metadata database to include stored definitions; it does not silently
return an incomplete list when that database is unavailable.

Create, Update and Delete always operate on persisted records. Deleting a record
does not remove its configuration override. The existing `ListRow` API also resolves database-backed
queries through its `query` parameter and binds parameters from the HTTP query
string. The new run endpoint supports HTTP and gRPC parameter objects directly.

Create body example (without a surrounding `query` property):

```json
{
  "name": "sales-summary",
  "sql": "SELECT region, SUM(amount) AS total FROM sales WHERE amount >= ? [AND region IN ?] GROUP BY region",
  "parameters": ["minimum:float", "regions:string:array"],
  "columns": [
    {"name": "region", "type": "string"},
    {"name": "total", "type": "float"}
  ],
  "jsonStyle": "lowerCamel"
}
```

Run body example:

```json
{"minimum": 100, "regions": ["north", "south"]}
```

Omitting `regions` omits the optional `AND` clause. `minimum` remains required.
An optional clause is included when all its parameters are supplied, omitted
when none are supplied, and rejected when only some are supplied. Empty arrays
count as absent; `0`, `false`, and the empty string are valid scalar values.
Nested optional clauses are unsupported. Quoted SQL text, comments and native
array subscripts retain their contents.

Definitions list parameters in SQL placeholder order. Repeated names are allowed
and reuse the same input; their type and array flags must agree. Types support
`string`, `integer`, `float`, and `bool` (with common aliases). Arrays expand
safely for either `IN ?` or `IN (?)`; a string input can use comma-separated
values. PostgreSQL `pg_array` binds string elements as one PostgreSQL text array.
Values are passed to database/sql separately from the SQL text.

Go/gRPC run responses contain `objects`, `totalCount` and `nextPageToken`.
HTTP uses the existing service response encoding and pagination metadata. `pageSize`,
`pageToken` and `skip` select a page from the full result; they do not rewrite the
predefined SQL. Filtering, sorting, projection and uniqueness belong in the SQL
and its declared parameters. For large datasets, include limit/offset parameters
in the SQL to limit database work.

## Save validation

Create and update validate the definition before writing:

- Database, query identity, parameter definitions and placeholder counts.
- Result column names, supported types, duplicate names and JSON style.
- A single SELECT or WITH SELECT statement, excluding mutation keywords and
  executable comments. Use doubled SQL quotes rather than backslash escapes.
- SQL plans with `EXPLAIN` (without `ANALYZE`) on SQLite, PostgreSQL and MySQL.
- Result column names and order through a zero-row wrapper query.
- Every reachable optional-parameter combination; definitions may contain up to
  six independently optional parameter names. Partial optional clauses are
  rejected at execution and are not considered reachable combinations.

Validation uses NULL bound values so an invented sample value cannot incorrectly
fail a valid UUID or date cast. It verifies SQL structure and metadata; actual
parameter values, result conversions and later schema changes can still produce
execution errors. The validation context has a five-second timeout. Unsupported
database dialects return an explicit validation error.

Configuration-only definitions with empty SQL retain the legacy example-result
behavior. Persisted definitions require SQL and result columns. If metadata
migration is disabled, provision the existing `DbQuery` model's table before
using database-backed definitions.
