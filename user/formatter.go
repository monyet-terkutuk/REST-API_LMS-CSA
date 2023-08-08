package user

type UserFormatter struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Avatar       string `json:"avatar"`
	Nim          string `json:"nim"`
	Email        string `json:"email"`
	NoHP         string `json:"no_hp"`
	Token        string `json:"token"`
	RoleID       bool   `json:"role_id"`
	Division     string `json:"division"`
	AlasanDaftar string `json:"alasan_daftar"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:           user.ID,
		Name:         user.Name,
		Avatar:       user.Avatar,
		Nim:          user.Nim,
		Email:        user.Email,
		Token:        token,
		Division:     user.Division,
		RoleID:       user.RoleID,
		NoHP:         user.NoHP,
		AlasanDaftar: user.AlasanDaftar,
	}
	return formatter
}
