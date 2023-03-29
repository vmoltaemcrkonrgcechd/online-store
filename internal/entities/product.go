package entities

type Product struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	User         User     `json:"user"`
	Color        Color    `json:"color"`
	Category     Category `json:"category"`
	UnitPrice    uint     `json:"unitPrice"`
	UnitsInStock uint     `json:"unitsInStock"`
	ImagePaths   []string `json:"imagePaths"`
}

type ProductDTO struct {
	Name         string   `json:"name"`
	UserID       string   `json:"userID"`
	ColorID      string   `json:"colorID"`
	CategoryID   string   `json:"categoryID"`
	UnitPrice    int      `json:"unitPrice"`
	UnitsInStock int      `json:"unitsInStock"`
	ImagePaths   []string `json:"imagePaths"`
}

type AllProductsDTO struct {
	Products []Product `json:"products"`

	Quantity uint16 `json:"quantity"`
}
