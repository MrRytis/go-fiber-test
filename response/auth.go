package response

type Login struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresAt    string `json:"expiresAt"`
}

type AuthUser struct {
	UserId  string `json:"userId"`
	Message string `json:"message"`
}
