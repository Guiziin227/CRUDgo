package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=4,max=32"`
	Username string `json:"username"`
	Age      int8   `json:"age" binding:"required,gt=0"`
}
