package user

type UserFormatter struct {
	ID    int    `json:"user_id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {

	return UserFormatter{
		ID:    user.ID,
		Email: user.Email,
		Token: token,
	}

}
