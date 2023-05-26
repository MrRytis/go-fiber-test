package exception

const (
	InternalServerError = 0
	MalformedRequest    = 1

	// 10 - 99: Authentication
	AuthTokenMalformed          = 10
	AuthTokenExpired            = 11
	AuthTokenNotValid           = 12
	AuthTokenBlacklisted        = 13
	AuthMissingHeader           = 14
	AuthUnauthorized            = 15
	AuthPasswordOrEmailMismatch = 16
	AuthUserNotFound            = 17
	AuthEmailNotVerified        = 18
)
