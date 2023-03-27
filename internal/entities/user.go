package entities

type User struct {
	ID        string  `json:"id"`
	Username  string  `json:"username"`
	City      City    `json:"city"`
	ImagePath *string `json:"imagePath"`
	Role      string  `json:"role"`
}

type UserDTO struct {
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	CityID    string  `json:"cityID"`
	ImagePath *string `json:"imagePath"`
	Role      string  `json:"role"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
