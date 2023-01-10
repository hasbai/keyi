package auth

type RegisterBody struct {
	Username     string `json:"username" validate:"max=32"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=8"`
	TEL          string `json:"tel"`
	TenantID     int    `json:"tenant_id" validate:"required"`
	TenantAreaID int    `json:"tenant_area_id"`
}

type LoginBody struct {
	Username string `json:"username" validate:"max=32"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required,min=8"`
}

type RefreshBody struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type ActivateQuery struct {
	Code string `query:"code" validate:"required"`
}

type ValidateModel struct {
	Email string `query:"email" validate:"email"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type MessageResponse struct {
	Message string `json:"message"`
}
