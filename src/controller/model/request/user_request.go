package request

type UserRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Age int8 `json:"age"`
}