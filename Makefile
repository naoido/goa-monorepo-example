gen-auth:
	cd microservices/auth && goa gen goa-example/microservices/auth/design

gen-greet:
	cd microservices/greet && goa gen goa-example/microservices/greet/design

run-auth:
	cd microservices/auth && export JWT_SECRET=secret && go build -o auth cmd/auth/main.go && ./auth

run-greet:
	cd microservices/greet && export JWT_SECRET=secret && go build -o greet cmd/greet/main.go && ./greet
