package request

type AdminUserCreateByDevRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

type AdminUserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
