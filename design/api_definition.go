package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("ComingServer", func() {
	Title("TinyKitten/ComingServer")
	Description("Coming API Server")
	Contact(func() {
		Name("TinyKitten")
		Email("catwalk8205@gmail.com")
		URL("https://github.com/TinyKitten/ComingServer/issues")
	})
	License(func() {
		Name("MIT")
		URL("https://github.com/TeamKitten/ComingServer/blob/master/LICENSE")
	})
	Docs(func() {
		Description("wiki")
		URL("https://github.com/TeamKitten/ComingServer/wiki")
	})
	Host("localhost:8080")
	Scheme("http")
	BasePath("/v1")

	Origin("*", func() {
		Expose("X-Time")
		Headers("Authorization", "Origin", "X-Requested-With", "Content-Type", "Accept")
		Methods("GET", "POST", "PUT", "DELETE")
		MaxAge(600)
		Credentials()
	})
})
