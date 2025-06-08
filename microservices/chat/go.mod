module goa-example/microservices/chat

go 1.24.0

replace goa-example/pkg/security => ../../pkg/security

require (
	github.com/google/uuid v1.6.0
	github.com/redis/go-redis/v9 v9.10.0
	goa-example/pkg/security v0.0.0-00010101000000-000000000000
	goa.design/clue v1.2.0
	goa.design/goa/v3 v3.21.1
	google.golang.org/grpc v1.72.1
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/aws/smithy-go v1.22.3 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dimfeld/httppath v0.0.0-20170720192232-ee938bf73598 // indirect
	github.com/go-chi/chi/v5 v5.2.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/gohugoio/hashstructure v0.5.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/manveru/faker v0.0.0-20171103152722-9fbc68a78c4d // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	go.opentelemetry.io/otel v1.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.35.0 // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sync v0.14.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/term v0.32.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	golang.org/x/tools v0.33.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250428153025-10db94c68c34 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
