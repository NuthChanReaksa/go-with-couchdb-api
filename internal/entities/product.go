package entities

type Product struct {
	ID    string  `json:"id,omitempty"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type BulkOperationResult struct {
	ID    string `json:"id,omitempty"`
	Rev   string `json:"rev,omitempty"`
	Error error  `json:"error,omitempty"`
}
