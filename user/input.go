package user

type RegisterUserInput struct {
	Name         string `json:"name" binding:"required"`
	Nim          string `json:"nim" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Division     string `json:"division" binding:"required"`
	NoHP         string `json:"no_hp" binding:"required"`
	AlasanDaftar string `json:"alasan_daftar" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type UpdateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Nim      string `json:"nim" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	// AvatarFileName string `json:"avatar_file"`
	Division string `json:"division" binding:"required"`
	NoHP     string `json:"no_hp" binding:"required"`
	RoleID   bool   `json:"role_id"`
}
