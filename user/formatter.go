package user

type UserFormatter struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func FormatUser(user User) UserFormatter {
	return UserFormatter{
		Age:      user.Age,
		Email:    user.Email,
		ID:       user.ID,
		Username: user.Username,
	}
}
