package domain

import "time"

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

type UserUpdateFormatter struct {
	ID         int       `json:"id"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	Age        int       `json:"age"`
	Updated_at time.Time `json:"updated_at"`
}

func FormatUserUpdateFormatter(user User) UserUpdateFormatter {
	return UserUpdateFormatter{
		ID:         user.ID,
		Email:      user.Email,
		Username:   user.Username,
		Age:        user.Age,
		Updated_at: user.UpdatedAt,
	}
}
