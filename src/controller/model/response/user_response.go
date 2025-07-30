package response

type UserResponse struct {
	Id       string `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username"`
	Age      int8   `json:"age" binding:"required,gt=0"`
}
