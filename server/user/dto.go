package user

type RegisterUserDto struct {
	Name      string `json:"name" binding:"required"`
	Ocupation string `json:"ocupation" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

type LoginUserDto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailDto struct {
	Email string `json:"email" binding:"required,email"`
}
