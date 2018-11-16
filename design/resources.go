package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
	})
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swaggerui/*filepath", "swaggerui/dist")
})

var _ = Resource("auth", func() {
	BasePath("/auth")

	Action("auth", func() {
		Description("JWTトークンを生成する")
		Routing(POST(""))

		Payload(func() {
			Attribute("id", Integer, "ID", func() {
				Example(0)
			})
			Attribute("screen_name", String, "スクリーンネーム", func() {
				Example("tinykitten")
			})
			Attribute("password", String, "パスワード", func() {
				Example("password")
				MinLength(6)
			})
			Required("screen_name", "password")
		})

		Response(OK, AuthSuccessMedia, func() {
			Headers(func() {
				Header("Authorization", String, "Generated JWT")
			})
		})
		Response(Unauthorized, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
})

var _ = Resource("users", func() {
	BasePath("/users")
	Security(JWT, func() {
		Scope("api:read")
	})
	Action("list", func() {
		Description("すべてのユーザーを取得")
		Security(OptionalJWT, func() {
			Scope("api:read")
		})
		Routing(
			GET(""),
		)

		Params(func() {
			Param("limit", Integer, "最大件数", func() {
				Default(100)
			})
			Param("offset", Integer, "取得開始位置", func() {
				Default(0)
			})
		})

		Response(OK, CollectionOf(UserMedia))

		Response(NotFound, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
	Action("show", func() {
		Description("ユーザーをIDで取得")
		Security(OptionalJWT, func() {
			Scope("api:read")
		})
		Routing(
			GET("/:id"),
		)
		Params(func() {
			Param("id", Integer, "ユーザーID", func() {
				Example(0)
			})
		})
		Response(OK, UserMedia)

		Response(NotFound, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
	Action("add", func() {
		Description("ユーザー追加")
		Security(JWT, func() {
			Scope("api:admin")
		})
		Routing(
			POST(""),
		)
		Payload(func() {
			Attribute("privilege_id", Integer, "権限ID", func() {
				Default(2)
				Example(0)
			})
			Attribute("password", String, "パスワード", func() {
				Example("password")
				MinLength(6)
			})
			Attribute("screen_name", String, "スクリーンネーム", func() {
				Example("tinykitten")
			})
			Required("privilege_id", "password", "screen_name")
		})

		Response(Created, UserMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
})

var _ = Resource("pods", func() {
	BasePath("/pods")
	Security(JWT, func() {
		Scope("api:read")
	})
	Action("list", func() {
		Description("すべてのポッドを取得")
		Routing(
			GET(""),
		)

		Params(func() {
			Param("limit", Integer, "最大件数", func() {
				Default(100)
			})
			Param("offset", Integer, "取得開始位置", func() {
				Default(0)
			})
		})

		Response(OK, CollectionOf(PodMedia))

		Response(NotFound, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
	Action("show", func() {
		Description("ポッドをIDで取得")
		Routing(
			GET("/:id"),
		)
		Params(func() {
			Param("id", Integer, "ポッドID", func() {
				Example(0)
			})
		})
		Response(OK, PodMedia)

		Response(NotFound, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
	Action("add", func() {
		Description("ポッド追加")
		Security(JWT, func() {
			Scope("api:write")
		})
		Routing(
			POST(""),
		)
		Payload(func() {
			Attribute("code", String, "ポッド名", func() {
				Example("SHIBUYA")
			})
			Attribute("latitude", Number, "緯度", func() {
				Example(35.658034)
			})
			Attribute("longitude", Number, "スクリーンネーム", func() {
				Example(139.701636)
			})
			Required("code", "latitude", "longitude")
		})

		Response(Created, PodMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
	Action("update", func() {
		Description("ポッド更新")
		Security(JWT, func() { // Use JWT to auth requests to this endpoint
			Scope("api:write")
		})
		Routing(
			PUT("/:id"),
		)
		Params(func() {
			Param("id", Integer, "ポッドID", func() {
				Example(0)
			})
		})
		Payload(func() {
			Attribute("latitude", Number, "緯度", func() {
				Example(35.689592)
			})
			Attribute("longitude", Number, "経度", func() {
				Example(139.700413)
			})
			Attribute("name", String, "名前", func() {
				Example("SHINJUKU")
			})
		})
		Response(NoContent)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
	Action("regenerate token", func() {
		Description("トークン再発行")
		Security(JWT, func() { // Use JWT to auth requests to this endpoint
			Scope("api:write")
		})
		Routing(
			POST("/:id/token"),
		)
		Params(func() {
			Param("id", Integer, "ポッドID", func() {
				Example(0)
			})
		})
		Response(OK, TokenMedia)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
})

var _ = Resource("peers", func() {
	BasePath("/peers")
	Security(JWT, func() {
		Scope("api:read")
	})
	Action("list", func() {
		Description("すべてのピアを取得")
		Routing(
			GET(""),
		)

		Params(func() {
			Param("limit", Integer, "最大件数", func() {
				Default(100)
			})
			Param("offset", Integer, "取得開始位置", func() {
				Default(0)
			})
		})

		Response(OK, CollectionOf(PeerMedia))

		Response(NotFound, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
	Action("show", func() {
		Description("ピアをIDで取得")
		Routing(
			GET("/:id"),
		)
		Params(func() {
			Param("id", Integer, "ピアID", func() {
				Example(0)
			})
		})
		Response(OK, PeerMedia)

		Response(NotFound, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
	Action("locations", func() {
		Description("ピアの位置履歴")
		Routing(
			GET("/:id/locations"),
		)
		Params(func() {
			Param("id", Integer, "ピアID", func() {
				Example(0)
			})
		})
		Response(OK, CollectionOf(PeerLocationMedia))

		Response(NotFound, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
	Action("current location", func() {
		Description("ピアの現在位置")
		Routing(
			GET("/:id/location"),
		)
		Params(func() {
			Param("id", Integer, "ピアID", func() {
				Example(0)
			})
		})
		Response(OK, PeerLocationMedia)

		Response(NotFound, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
	Action("add", func() {
		Description("ピア追加")
		Security(JWT, func() {
			Scope("api:write")
		})
		Routing(
			POST(""),
		)
		Payload(func() {
			Attribute("code", String, "ポッド名", func() {
				Example("TS")
			})
			Required("code")
		})

		Response(Created, PodMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
	Action("update", func() {
		Description("ピア更新")
		Security(JWT, func() { // Use JWT to auth requests to this endpoint
			Scope("api:write")
		})
		Routing(
			PUT("/:id"),
		)
		Params(func() {
			Param("code", Integer, "ピアID", func() {
				Example(0)
			})
		})
		Payload(func() {
			Attribute("name", String, "名前", func() {
				Example("TS-IPHONE")
			})
		})
		Response(NoContent)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
	Action("regenerate token", func() {
		Description("トークン再発行")
		Security(JWT, func() { // Use JWT to auth requests to this endpoint
			Scope("api:write")
		})
		Routing(
			PATCH("/:id/token"),
		)
		Params(func() {
			Param("id", Integer, "ピアID", func() {
				Example(0)
			})
		})
		Response(OK, TokenMedia)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
	Action("send location", func() {
		Description("ピア位置送信")
		Security(OptionalJWT, func() {
			// 基本的にトークンが一致しなければ送信できないが、管理者は無条件で送信できる
			Scope("api:admin")
		})
		Routing(
			POST("/:id"),
		)
		Params(func() {
			Param("id", Integer, "ピアID", func() {
				Example(0)
			})
		})
		Payload(func() {
			Attribute("latitude", Number, "緯度", func() {
				Example(35.658034)
			})
			Attribute("longitude", Number, "スクリーンネーム", func() {
				Example(139.701636)
			})
			Required("latitude", "longitude")
		})
		Response(OK, TokenMedia)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
})