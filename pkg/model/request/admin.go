package request

type AdminUserCreateByDevRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AdminUserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminUserRefreshToken struct {
	RefreshToken string `json:"refreshToken"`
}
