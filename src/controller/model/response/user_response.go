package response

type UserResponse struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Age int8 `json:"age"`
}