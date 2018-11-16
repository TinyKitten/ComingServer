package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var JWT = JWTSecurity("jwt", func() {
	Header("Authorization")
	Scope("api:read", "API READ access")
	Scope("api:write", "API WRITE access")
	Scope("api:admin", "API ADMIN access")
})

var OptionalJWT = JWTSecurity("optional_jwt", func() {
	Header("Authorization")
	Scope("api:read", "API READ access")
	Scope("api:write", "API WRITE access")
	Scope("api:admin", "API ADMIN access")
})
