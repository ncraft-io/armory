module github.com/ncraft-io/armory/service-go

go 1.16

replace github.com/ncraft-io/armory/go => ../go

replace github.com/mojo-lang/db/go => ../../../Mojo/db/go

require (
	github.com/etherlabsio/healthcheck/v2 v2.0.0
	github.com/go-kit/kit v0.13.0
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/gorilla/mux v1.7.4
	github.com/iancoleman/strcase v0.2.0
	github.com/json-iterator/go v1.1.12
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	github.com/mojo-lang/core/go v0.0.0-20240205040023-0a270faf3af3
	github.com/mojo-lang/db/go v0.0.0-20231010082619-1822aed7c28f
	github.com/mojo-lang/http/go v0.0.0-20231026054523-2cf45f147a95
	github.com/ncraft-io/armory/go v0.0.0
	github.com/ncraft-io/ncraft-gokit v0.0.0-20231030074420-02d8bebbc396
	github.com/ncraft-io/ncraft/go v0.0.0-20231130060710-4e89e62eb679
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.1
	github.com/rs/cors v1.7.0
	github.com/segmentio/ksuid v1.0.4
	go.uber.org/automaxprocs v1.5.3
	google.golang.org/grpc v1.58.2
	gorm.io/gorm v1.25.4
)
