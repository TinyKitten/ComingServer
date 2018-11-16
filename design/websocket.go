package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("websocket", func() {
	Action("send current peer location", func() {
		Routing(GET("peer"))
		Scheme("ws")
		Description("ピアのWebSocketエンドポイント")
		Params(func() {
			Param("token", String, "接続用トークン", func() {
				Example("AHO-AHO-MAN")
			})
			Required("token")
		})
		Response(SwitchingProtocols)
	})

	Action("receive peer location", func() {
		Routing(GET("pod"))
		Scheme("ws")
		Description("ポッドのWebSocketエンドポイント")
		Params(func() {
			Param("token", String, "接続用トークン", func() {
				Example("AHO-AHO-MAN")
			})
			Required("token")
		})
		Response(SwitchingProtocols)
	})
})
