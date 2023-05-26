package request

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type VerifyRequest struct {
	UserId int32 `json:"userId"`
}

type ReminderRequest struct {
	Email string `json:"email"`
}

type ResetRequest struct {
	UserId   int32  `json:"userId"`
	Password string `json:"password"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refreshToken"`
}
