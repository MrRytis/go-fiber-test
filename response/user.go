package response

type User struct {
	Uid     string  `json:"uid"`
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Email   string  `json:"email"`
	Icon    *string `json:"icon"`
}
