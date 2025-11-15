package request

type CreateExerciseRequest struct {
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description"`
	Items       []string `json:"items" validate:"required,min=5"` // minimum 5 soal
}

type RecordAttemptRequest struct {
	ItemID uint `json:"item_id" validate:"required"`
}
