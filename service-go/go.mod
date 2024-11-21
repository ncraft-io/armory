module github.com/ncraft-io/armory/service-go

go 1.22

toolchain go1.22.4

replace github.com/ncraft-io/armory/go => ../go

require (
	github.com/etherlabsio/healthcheck/v2 v2.0.0
	github.com/go-kit/kit v0.13.0
	github.com/gorilla/mux v1.8.0
	github.com/iancoleman/strcase v0.2.0
	github.com/json-iterator/go v1.1.12
	github.com/mojo-lang/core/go v0.0.0-20241101072631-6f07aa735416
	github.com/mojo-lang/db/go v0.0.0-20241121074108-50ee5e6736ca
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
	gorm.io/gorm v1.25.12
)

require (
	atomicgo.dev/cursor v0.1.1 // indirect
	atomicgo.dev/keyboard v0.2.9 // indirect
	atomicgo.dev/schedule v0.0.2 // indirect
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/BurntSushi/toml v1.0.0 // indirect
	github.com/ReneKroon/ttlcache/v2 v2.11.0 // indirect
	github.com/alecthomas/chroma v0.10.0 // indirect
	github.com/alibabacloud-go/debug v0.0.0-20190504072949-9472017b5c68 // indirect
	github.com/alibabacloud-go/tea v1.1.17 // indirect
	github.com/alibabacloud-go/tea-utils v1.4.4 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1800 // indirect
	github.com/aliyun/alibabacloud-dkms-gcs-go-sdk v0.2.2 // indirect
	github.com/aliyun/alibabacloud-dkms-transfer-go-sdk v0.1.7 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.0 // indirect
	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/brianvoe/gofakeit/v6 v6.23.2 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/containerd/console v1.0.3 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/daveshanley/vacuum v0.3.7 // indirect
	github.com/dlclark/regexp2 v1.10.0 // indirect
	github.com/dop251/goja v0.0.0-20230707174833-636fdf960de1 // indirect
	github.com/dop251/goja_nodejs v0.0.0-20230602164024-804a84515562 // indirect
	github.com/dprotaso/go-yit v0.0.0-20220510233725-9ba8df137936 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gertd/go-pluralize v0.2.1 // indirect
	github.com/getkin/kin-openapi v0.118.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-kit/log v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/swag v0.19.5 // indirect
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/go-swiss/fonts v0.0.0-20230807175105-90067c2f5042 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/pprof v0.0.0-20230207041349-798e818bf904 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gookit/color v1.5.3 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/invopop/yaml v0.1.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/lithammer/fuzzysearch v1.1.8 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/mattn/go-sqlite3 v1.14.24 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/mojo-lang/document/go v0.0.0-20231009062124-1063661ebfdd // indirect
	github.com/mojo-lang/lang/go v0.0.0-20240726065356-77211e8660f2 // indirect
	github.com/mojo-lang/openapi/go v0.0.0-20240104072634-11b492bf84c9 // indirect
	github.com/mojo-lang/protobuf/go v0.0.0-20231010083050-67d0430d7b2c // indirect
	github.com/mojo-lang/yaml/go v0.0.0-20231010072828-fb7b8987a3f0 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/nacos-group/nacos-sdk-go/v2 v2.2.7 // indirect
	github.com/pb33f/libopenapi v0.9.7 // indirect
	github.com/pb33f/libopenapi-validator v0.0.15 // indirect
	github.com/perimeterx/marshmallow v1.1.4 // indirect
	github.com/phpdave11/gofpdf v1.4.2 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/pterm/pterm v0.12.62 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/santhosh-tekuri/jsonschema/v5 v5.3.0 // indirect
	github.com/smilextay/dm v0.0.0-20230925054544-9a664a54ce66 // indirect
	github.com/smilextay/gorm-dm v0.0.0-20230925055508-bf78af3bbc54 // indirect
	github.com/stephenafamo/goldmark-pdf v0.2.0 // indirect
	github.com/uber/jaeger-client-go v2.25.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	github.com/vmware-labs/yaml-jsonpath v0.3.2 // indirect
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	github.com/yuin/goldmark v1.5.6 // indirect
	go.etcd.io/etcd/api/v3 v3.5.0 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.0 // indirect
	go.etcd.io/etcd/client/v3 v3.5.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sync v0.9.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/term v0.18.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20230913181813-007df8e322eb // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230803162519-f966b187b2e5 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230920204549-e6e6cdab5c13 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/mysql v1.3.3 // indirect
	gorm.io/driver/postgres v1.5.9 // indirect
	gorm.io/driver/sqlite v1.5.6 // indirect
)
