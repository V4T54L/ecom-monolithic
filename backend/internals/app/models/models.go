package models

type User struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type AuthToken struct {
	ID        int
	Name      string
	Username  string
	Email     string
	Role      string
	CreatedAt string
	UpdatedAt string
}

// type Address struct {
// 	ID      int    `json:"id,omitempty"`
// 	Address string `json:"address"`
// 	City    string `json:"city"`
// 	Zipcode string `json:"zipcode"`
// }

// type Category struct {
// 	ID       int
// 	Name     string
// 	StartIdx int
// 	EndIdx   int
// }

// type Product struct {
// 	ID              int
// 	ThumbnailURL    string
// 	Title           string
// 	Description     string
// 	Rating          float32
// 	Price           float32
// 	DiscountedPrice float32
// 	Category        *Category

// }
