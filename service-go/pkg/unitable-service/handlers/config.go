package handlers

import "github.com/ncraft-io/armory/go/pkg/armory/unitable"

// DBQueryConfig holds service configuration loaded from dbQuery.queries.
// DbQuery remains the shared API entity; this wrapper is service-local.
type DBQueryConfig struct {
	Queries []*unitable.DbQuery `json:"queries,omitempty"`
}
