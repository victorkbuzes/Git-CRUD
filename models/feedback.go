package models

type Feedback struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateFeadbackInput struct {
	Name        string `json:"name"`
	Description string `json:"description" binding:"required"`
}

type UpdateFeadbackInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
