package product_dto

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	MakerID     string `json:"maker_id"`
	CheckerID   string `json:"checker_id"`
	SignerID    string `json:"signer_id"`
	CreatedAT   string `json:"created_at"`
	UpdatedAT   string `json:"updated_at"`
	DeletedAT   string `json:"deleted_at"`
}
