package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=4,containsany=!@#$%*_"`
	Username string `json:"username" binding:"required,min=4,max=100"`
	Age      int8   `json:"age" binding:"required,gt=0"`
}
