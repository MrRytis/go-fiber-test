package request

type UpdateUserDetailsRequest struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
