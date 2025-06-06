ARG GOLANG_VERSION=1.24
FROM golang:${GOLANG_VERSION}-alpine AS builder

ARG SERVICE_NAME

RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN cd ./microservices/${SERVICE_NAME} &&  go mod download

RUN CGO_ENABLED=0 go build \
  -ldflags="-w -s" \
  -o /bin/${SERVICE_NAME} \
  ./microservices/${SERVICE_NAME}/cmd/${SERVICE_NAME}


FROM alpine:latest

ARG SERVICE_NAME

RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    echo "Asia/Tokyo" > /etc/timezone

WORKDIR /app

COPY --from=builder /bin/${SERVICE_NAME} /app/server

ENTRYPOINT ["sh", "-c", "exec /app/server"]