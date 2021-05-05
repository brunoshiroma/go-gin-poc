package requests

type Client struct {
	Id    uint64 `json:"id" binding:"numeric,gte=0"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}
