module goa-example/microservices/greet

go 1.24.0

replace goa-example/pkg/security => ../../pkg/security

require (
	goa-example/pkg/security v0.0.0-00010101000000-000000000000
	goa.design/goa/v3 v3.21.1
	golang.org/x/net v0.40.0
	google.golang.org/grpc v1.72.1
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dimfeld/httppath v0.0.0-20170720192232-ee938bf73598 // indirect
	github.com/gohugoio/hashstructure v0.5.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/manveru/faker v0.0.0-20171103152722-9fbc68a78c4d // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/sync v0.14.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	golang.org/x/tools v0.33.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250428153025-10db94c68c34 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
