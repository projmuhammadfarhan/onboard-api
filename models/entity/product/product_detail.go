package product

type ProductDetail struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	MakerID     string `json:"maker"`
	CheckerID   string `json:"checker"`
	SignerID    string `json:"signer"`
	MakerName   string `json:"maker_name"`
	CheckerName string `json:"checker_name"`
	SignerName  string `json:"signer_name"`
}
