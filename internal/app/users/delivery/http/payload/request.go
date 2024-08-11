package payload

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type FindUser struct {
	Id uint64 `json:"id" query:"id" param:"id"`
}
