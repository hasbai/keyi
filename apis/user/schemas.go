package user

type RegisterModel struct {
	Name        string `json:"name" validate:"max=32"`
	Description string `json:"description" validate:"max=256"`
	Email       string `json:"email" validate:"required,email,max=64"`
	Password    string `json:"password" validate:"required,max=128"`
	Avatar      string `json:"avatar" validate:"max=256"`
	Code        string `json:"code" validate:"required,max=6,min=6"`
}

type EmailModel struct {
	Email string `json:"email" validate:"required,email,max=64"`
}

type ModifyModel struct {
	Name        string `json:"name" validate:"max=32"`
	Description string `json:"description" validate:"max=256"`
	// to change password, either password or code is required
	NewPassword string `json:"new_password" validate:"max=128"`
	Password    string `json:"password" validate:"max=128"`
	Code        string `json:"code" validate:"max=6"`
	Avatar      string `json:"avatar" validate:"max=256"`
}
