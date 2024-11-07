package users

type UserCreateInput struct {
	Email     string `json:"email" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type UserCreateOutput struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
}
