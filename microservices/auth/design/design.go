package design

import (
	. "goa-example/pkg/security"
	. "goa.design/goa/v3/dsl"
)

var _ = Service("auth", func() {
	Description("The auth service")

	// gRPCを有効化・設定
	GRPC(func() {
		Package("auth.v1")
	})

	// 共通エラーを定義
	Error("permission_denied", String, "権限が不足しています")
	Error("invalid_argument", String, "リクエストされた引数が無効です")
	Error("internal", String, "処理中に重大なエラーが発生しました")

	// エンドポイントの設定
	Method("login", func() {
		Description("Login")

		Payload(func() {
			Field(1, "username", String, "username")
			Field(2, "password", String, "password")
			Required("username", "password")
		})

		Result(func() {
			Field(1, "access_token", String)
			Field(2, "refresh_token", String)
			Required("access_token", "refresh_token")
		})

		Error("unauthenticated", String, "ユーザーが存在しないまたは、パスワードが無効です")

		// メソッド固有のgRPC設定
		GRPC(func() {
			// 成功レスポンス
			Response(CodeOK)

			// エラーレスポンス
			Response("unauthenticated", CodeUnauthenticated)
			Response("invalid_argument", CodeInvalidArgument)
			Response("internal", CodeInternal)
		})
	})

	Method("logout", func() {
		Description("Logout")

		Security(JWTAuth, func() {
			Scope("api:read")
		})

		Payload(func() {
			Token("token", String, "access_token")
			Required("token")
		})

		GRPC(func() {
			Response(CodeOK)

			Response("invalid_argument", CodeInvalidArgument)
		})
	})

	Method("refresh", func() {
		Description("Refresh")

		Payload(func() {
			Field(1, "refresh_token", String, "refresh_token")
			Required("refresh_token")
		})

		Result(func() {
			Field(1, "access_token", String)
			Field(2, "refresh_token", String)
			Required("access_token", "refresh_token")
		})

		GRPC(func() {
			Response(CodeOK)

			Response("permission_denied", CodePermissionDenied)
		})
	})
})
