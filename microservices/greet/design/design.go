package design

import (
	"goa-example/pkg/security"
	. "goa.design/goa/v3/dsl"
)

// サービスを定義
var _ = Service("greet", func() {
	// サービスの説明
	Description("Greet service")

	// gRPCの有効化・サービス共通の設定
	GRPC(func() {
		Package("greet.v1")
	})

	// メソッド(エンドポイント)の定義
	Method("Greet", func() {
		Description("Greet method")

		// 戻り値をStringで定義する
		Result(String)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	// 認証付きエンドポイント
	Method("Hello", func() {
		Description("Hello method")

		// スコープを検証する
		Security(security.JWTAuth, func() {
			Scope("api:read")
		})

		// 認証に使うTokenをペイロードに含める
		Payload(func() {
			Token("token", String, "access_token")
			Required("token")
		})

		Result(String)

		GRPC(func() {
			Response(CodeOK)
		})
	})
})
