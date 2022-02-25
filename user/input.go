package user

type RegisterUserInput struct {
	Age      int    `json:"age" binding:"required,gte=8"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Username string `json:"username" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UpdateUserInput struct {
	ID       int
	Email    string `json:"email"`
	Username string `json:"username"`
}
