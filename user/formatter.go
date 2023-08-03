package user

type UserFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Nim      string `json:"nim"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Division string `json:"division"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Name:     user.Name,
		Nim:      user.Nim,
		Email:    user.Email,
		Token:    token,
		Division: user.Division,
	}
	return formatter
}
