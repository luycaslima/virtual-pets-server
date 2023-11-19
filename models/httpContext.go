package models

// Structure to pass the issuer fromthe JWT auth to the route that requires auth
type HttpContextStruct struct {
	JwtIssuer string
}
