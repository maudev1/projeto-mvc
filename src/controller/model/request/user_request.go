package request

type UserRequest struct {
	Name       string `json:"name" binding:"required,min=3,max=20"`
	Password   string `json:"password" binding:"required,min=4"`
	Age        int8   `json:"age" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Newsletter bool   `json:"newsletter" binding:""`
}
