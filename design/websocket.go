package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("websocket", func() {
	Action("send current peer location", func() {
		Routing(GET("echo"))
		Scheme("ws")
		Description("echo websocket server")
		Params(func() {
			Param("token", String, "Secret token for send location", func() {
				Example("AHO-AHO-MAN")
			})
			Param("latitude", Number, "緯度", func() {
				Example(35.658034)
			})
			Param("longitude", Number, "スクリーンネーム", func() {
				Example(139.701636)
			})
			Required("token", "latitude", "longitude")
		})
		Response(SwitchingProtocols)
	})

})
