package design

import (
	"goa-example/pkg/security"
	. "goa.design/goa/v3/dsl"
)

var _ = API("chat", func() {
	Title("Chat Service")
})

var _ = Service("chat", func() {
	Description("The chat service invokes the chat.")

	GRPC(func() {
		Package("chat.v1")
	})

	Error("unauthorized", String)
	Error("permission-denied", String)
	Error("internal", String)

	Method("create-room", func() {
		Description("Creates a new chat room.")

		Security(security.JWTAuth, func() {
			Scope("api:write")
		})

		Payload(func() {
			Token("token", String, "The access token")

			Required("token")
		})
		Result(String)

		GRPC(func() {
			Response(CodeOK)

			Response("internal", CodeInvalidArgument)
		})
	})

	Method("history", func() {
		Description("Get all chat rooms history.")

		Security(security.JWTAuth, func() {
			Scope("api:read")
		})

		Payload(func() {
			Token("token", String, "The access token")

			Field(1, "room_id", String, "The id of the room")

			Required("token", "room_id")
		})

		Result(ArrayOf(Chat))

		GRPC(func() {
			Response(CodeOK)

			Response("unauthorized", CodeInvalidArgument)
			Response("permission-denied", CodePermissionDenied)
		})
	})

	Method("stream-room", func() {
		Description("Streams chat room events on a chat room.")

		Security(security.JWTAuth, func() {
			Scope("api:read")
		})

		Payload(func() {
			Token("token", String, "The access token")
			Field(1, "room_id", String, "The room id")

			Required("token", "room_id")
		})

		StreamingPayload(String)
		StreamingResult(Chat)

		GRPC(func() {
			Response(CodeOK)

			Response("unauthorized", CodeInvalidArgument)
			Response("permission-denied", CodePermissionDenied)
		})
	})
})

var Chat = Type("Chat", func() {
	Description("Chat")

	Field(1, "username", String, "username")
	Field(2, "message", String, "message")
	Field(3, "sent_at", Int64, "sent_at")
	Required("username", "message", "sent_at")
})
