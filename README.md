# ✨ Monorepo Microservice Example with Goa
## Speaker Deck
説明資料入れてます！
https://speakerdeck.com/naoido/goadeshi-merumaikurosabisu

## Requirements Tools
```sh
brew install go

# Goa CLIのインストール
go install goa.design/goa/v3/cmd/goa@latest

# proto関連
brew install protoc
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## Run microservices on docker
```sh
# localhost:8090  :  Auth Server
# localhost:8091  :  Greet Server
docker compose up --build
```
