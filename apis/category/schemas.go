package category

type CreateModel struct {
	Name        string `json:"name" validate:"max=32,required"`
	Description string `json:"description" validate:"max=256"`
}

type DeleteModel struct {
	To int `json:"to" validate:"min=1,required"`
}
