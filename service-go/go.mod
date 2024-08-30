module github.com/ncraft-io/armory/service-go

go 1.16

replace github.com/ncraft-io/armory/go => ../go

require (
	github.com/etherlabsio/healthcheck/v2 v2.0.0
	github.com/go-kit/kit v0.13.0
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/iancoleman/strcase v0.2.0
	github.com/json-iterator/go v1.1.12
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	github.com/mojo-lang/core/go v0.0.0-20240808085347-53eed0a2e240
	github.com/mojo-lang/db/go v0.0.0-20240807121857-107d18e42a06
	github.com/mojo-lang/http/go v0.0.0-20240627072315-89ee9c59b807
	github.com/mojo-lang/mojo/go v0.0.0-20240807122300-2fee6f63223b
	github.com/ncraft-io/armory/go v0.0.0
	github.com/ncraft-io/ncraft-gokit v0.0.0-20240808082035-9f3ee3a0659d
	github.com/ncraft-io/ncraft/go v0.0.0-20231130060710-4e89e62eb679
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.12.2
	github.com/rs/cors v1.7.0
	github.com/segmentio/ksuid v1.0.4
	github.com/zeebo/assert v1.3.0
	go.uber.org/automaxprocs v1.5.3
	google.golang.org/grpc v1.58.3
	gorm.io/gorm v1.25.4
)
