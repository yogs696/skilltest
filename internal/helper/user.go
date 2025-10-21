package helper

import "github.com/yogs696/skilltest/internal/entity"

type UserFormatter struct {
	ID       uint64 `json:"id"`
	Username string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func FormatUser(user *entity.User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}
	return formatter
}
