package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// AuthSuccessMedia 認証成功メディア
var AuthSuccessMedia = MediaType("application/vnd.auth.succes+json", func() {
	Description("認証成功")

	Attributes(func() {
		Attribute("id", Integer, "ID", func() {
			Metadata("struct:field:type", "uint64")
			Example(1)
		})
		Attribute("screen_name", String, "コード", func() {
			Example("tinykitten")
		})
		Attribute("token", String, "トークン", func() {
			Example("eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ0aW55a2l0dGVuIiwiZXhwIjoxNTQyMzQ1MTM1LCJpYXQiOjE1NDIzNDE1MzUsImlzcyI6IkNvbWluZyIsImp0aSI6IjI2ZGExM2E5LTAwZjEtNGQ1NS1hNjM2LWNkOWU5Njc5ZTliZSIsIm5iZiI6Miwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIiwiYXBpOmFkbWluIl0sInN1YiI6IkFjY2Vzc1Rva2VuIn0.NbRlJqUDj5eSWI0v-PV87BVcqo-usM3xwgYjHQcAbS454zIqXUUt1bER3XXgQetrKgt1HRc7_6b3qUXvx6XfWZFqUyEHLZY2QBzKiTwo0rQlXOREDiDDunKTa7Z-BPPT2beJiYjhk7_wp6APbObbiUfW9iRNdNgOHKJ1_qy_6QoqRBrbTXEI-n8pgOsalzBYZuU5fIMLPHPKZmJelwpz5w5ag8ARtLahNlUYPT7iAU_XfBbQVZ5WAKwW2XaLsf4VP1XH4XKsCtUMjftpW7Xrn9avVtzHH9SmmN7HY_rje250g5rsz-nh7-ylUSOgxhu_93MkbnICu0DIjCxaKMOXsg")
		})

		Required(
			"screen_name",
			"token",
		)
	})
	View("default", func() {
		Attribute("screen_name")
		Attribute("token")
	})
})

// UserMedia ユーザー
var UserMedia = MediaType("application/vnd.user+json", func() {
	Description("ユーザー")

	Attributes(func() {
		Attribute("id", Integer, "ID", func() {
			Metadata("struct:field:type", "uint64")
			Example(1)
		})
		Attribute("screen_name", String, "スクリーンネーム", func() {
			Example("tinykitten")
		})
		Attribute("created_at", Integer, "作成日", func() {
			Metadata("struct:field:type", "int64")
			Example(1017327600)
		})
		Attribute("updated_at", Integer, "更新日", func() {
			Metadata("struct:field:type", "int64")
			Example(1017327600)
		})

		Required(
			"id",
			"screen_name",
			"created_at",
			"updated_at",
		)
	})
	View("default", func() {
		Attribute("id")
		Attribute("screen_name")
		Attribute("created_at")
		Attribute("updated_at")
	})
})

// PodMedia ポッド
var PodMedia = MediaType("application/vnd.pod+json", func() {
	Description("ポッド")

	Attributes(func() {
		Attribute("id", Integer, "ID", func() {
			Metadata("struct:field:type", "int64")
			Example(1)
		})
		Attribute("code", String, "コード", func() {
			Example("SHIBUYA")
		})
		Attribute("latitude", Number, "緯度", func() {
			Example(35.658034)
		})
		Attribute("longitude", Number, "経度", func() {
			Example(139.701636)
		})
		Attribute("created_at", Number, "作成日", func() {
			Metadata("struct:field:type", "int64")
			Example(1017327600)
		})
		Attribute("updated_at", Integer, "更新日", func() {
			Metadata("struct:field:type", "int64")
			Example(1017327600)
		})

		Required(
			"id",
			"code",
			"latitude",
			"longitude",
			"created_at",
			"updated_at",
		)
	})
	View("default", func() {
		Attribute("id")
		Attribute("code")
		Attribute("latitude")
		Attribute("longitude")
		Attribute("created_at")
		Attribute("updated_at")
	})
})

// PeerMedia ピア
var PeerMedia = MediaType("application/vnd.peer+json", func() {
	Description("ピア")

	Attributes(func() {
		Attribute("id", Integer, "ID", func() {
			Metadata("struct:field:type", "int64")
			Example(1)
		})
		Attribute("code", String, "コード", func() {
			Example("TS")
		})
		Attribute("created_at", Number, "作成日", func() {
			Metadata("struct:field:type", "int64")
			Example(1017327600)
		})
		Attribute("updated_at", Integer, "更新日", func() {
			Metadata("struct:field:type", "int64")
			Example(1017327600)
		})
		Attribute("approaching", Boolean, "ポッド装置に接近している", func() {
			Example(false)
		})

		Required(
			"id",
			"code",
			"created_at",
			"updated_at",
			"approaching",
		)
	})
	View("default", func() {
		Attribute("id")
		Attribute("code")
		Attribute("created_at")
		Attribute("updated_at")
		Attribute("approaching")
	})
})

// PodCreatedMedia ポッド作成成功
var PodCreatedMedia = MediaType("application/vnd.pod.created+json", func() {
	Description("ポッド作成成功")

	Attributes(func() {
		Attribute("id", Integer, "ID", func() {
			Metadata("struct:field:type", "int64")
			Example(1)
		})
		Attribute("code", String, "コード", func() {
			Example("SHIBUYA")
		})
		Attribute("latitude", Number, "緯度", func() {
			Example(35.658034)
		})
		Attribute("longitude", Number, "経度", func() {
			Example(139.701636)
		})
		Attribute("token", String, "トークン", func() {
			Example("AHO-AHO-MAN")
		})
		Attribute("created_at", Number, "作成日", func() {
			Metadata("struct:field:type", "int64")
			Example(1017327600)
		})
		Attribute("updated_at", Integer, "更新日", func() {
			Metadata("struct:field:type", "int64")
			Example(1017327600)
		})

		Required(
			"id",
			"code",
			"latitude",
			"longitude",
			"token",
			"created_at",
			"updated_at",
		)
	})
	View("default", func() {
		Attribute("id")
		Attribute("code")
		Attribute("latitude")
		Attribute("longitude")
		Attribute("token")
		Attribute("created_at")
		Attribute("updated_at")
	})
})

// PeerCreatedMedia ピア作成成功
var PeerCreatedMedia = MediaType("application/vnd.peer.created+json", func() {
	Description("ピア作成成功")

	Attributes(func() {
		Attribute("id", Integer, "ID", func() {
			Metadata("struct:field:type", "int64")
			Example(1)
		})
		Attribute("code", String, "コード", func() {
			Example("TS")
		})
		Attribute("token", String, "トークン", func() {
			Example("AHO-AHO-MAN")
		})
		Attribute("created_at", Number, "作成日", func() {
			Metadata("struct:field:type", "int64")
			Example(1017327600)
		})
		Attribute("updated_at", Integer, "更新日", func() {
			Metadata("struct:field:type", "int64")
			Example(1017327600)
		})

		Required(
			"id",
			"code",
			"token",
			"created_at",
			"updated_at",
		)
	})
	View("default", func() {
		Attribute("id")
		Attribute("code")
		Attribute("token")
		Attribute("created_at")
		Attribute("updated_at")
	})
})

// PeerLocationMedia ピア位置
var PeerLocationMedia = MediaType("application/vnd.peer.location+json", func() {
	Description("ピア位置情報")

	Attributes(func() {
		Attribute("latitude", Number, "緯度", func() {
			Example(35.658034)
		})
		Attribute("longitude", Number, "経度", func() {
			Example(139.701636)
		})
		Attribute("created_at", Number, "作成日", func() {
			Metadata("struct:field:type", "int64")
			Example(1017327600)
		})
		Attribute("updated_at", Integer, "更新日", func() {
			Metadata("struct:field:type", "int64")
			Example(1017327600)
		})

		Required(
			"latitude",
			"longitude",
			"created_at",
			"updated_at",
		)
	})
	View("default", func() {
		Attribute("latitude")
		Attribute("longitude")
		Attribute("created_at")
		Attribute("updated_at")
	})
})

// PeerApproachingMedia ピア接近メディア
var PeerApproachingMedia = MediaType("application/vnd.peer.approaching+json", func() {
	Description("ピア接近")

	Attributes(func() {
		Attribute("type", String, "タイプ", func() {
			Example("APPROACHING")
			Default("APPROACHING")
		})
		Attribute("code", String, "ピアコード", func() {
			Example("TS")
		})
		Attribute("latitude", Number, "緯度", func() {
			Example(35.658034)
		})
		Attribute("longitude", Number, "経度", func() {
			Example(139.701636)
		})
		Attribute("created_at", Number, "作成日", func() {
			Metadata("struct:field:type", "int64")
			Example(1017327600)
		})
		Attribute("approaching", Boolean, "ポッド装置に接近している", func() {
			Example(false)
		})

		Required(
			"type",
			"code",
			"latitude",
			"longitude",
			"created_at",
			"approaching",
		)
	})
	View("default", func() {
		Attribute("type")
		Attribute("code")
		Attribute("latitude")
		Attribute("longitude")
		Attribute("created_at")
		Attribute("approaching")
	})
})

// WSErrorMedia WebSocketエラーメディア
var WSErrorMedia = MediaType("application/vnd.ws.error+json", func() {
	Description("WebSocketエラー")

	Attributes(func() {
		Attribute("type", String, "タイプ", func() {
			Example("ERROR")
			Default("ERROR")
		})
		Attribute("code", String, "エラーコード", func() {
			Example("ERR_PEER_NOT_FOUND")
		})
		Attribute("message", String, "エラーメッセージ", func() {
			Example("Peer not found.")
		})

		Required(
			"type",
			"code",
			"message",
		)
	})
	View("default", func() {
		Attribute("type")
		Attribute("code")
		Attribute("message")
	})
})

// TokenMedia トークンメディア
var TokenMedia = MediaType("application/vnd.token+json", func() {
	Description("トークン")

	Attributes(func() {
		Attribute("id", Integer, "対象ID", func() {
			Metadata("struct:field:type", "uint64")
			Example(1)
		})
		Attribute("token", String, "トークン", func() {
			Example("AHO-AHO-MAN")
		})
		Attribute("updated_at", Integer, "更新日", func() {
			Metadata("struct:field:type", "int64")
			Example(1017327600)
		})

		Required(
			"id",
			"token",
			"updated_at",
		)
	})
	View("default", func() {
		Attribute("id")
		Attribute("token")
		Attribute("updated_at")
	})
})
