package request

type UserRequest struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	Age        int8   `json:"age"`
	Email      string `json:"email"`
	Newsletter bool   `json:"newsletter"`
}
