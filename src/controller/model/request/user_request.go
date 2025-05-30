package request

type UserRequest struct {
	Name string `json:"name" binding:"required,min=4,max=20"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#S%*"`
	Age int8 `json:"age" binding:"required,numeric,min=1,max=100"`
}