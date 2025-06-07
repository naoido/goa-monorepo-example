package design

import . "goa.design/goa/v3/dsl"

var _ = Service("chat", func() {
	Description("The chat service invokes the chat.")

	GRPC(func() {
		Package("chat.v1")
	})

	Method("creat-room", func() {
		Description("Creates a new chat room.")

		Payload(func() {

		})
		Result(String)

		GRPC(func() {})
	})
})
