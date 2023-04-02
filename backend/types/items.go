package types

type Items struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string `json:"name,omitempty" bson:"name"`
	Description string `json:"description,omitempty" bson:"description"`
	ImageUrl    string `json:"image_url,omitempty" bson:"image_url"`
	UserID      string `json:"user_id,omitempty" bson:"user_id"`
	User        []User `json:"user,omitempty" bson:"user"`
	Bids        *[]Bid `json:"bids" bson:"bids"`
}

type Bid struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	UserID string  `json:"user_id,omitempty" bson:"user_id"`
	User   []User  `json:"user,omitempty" `
	Price  float32 `json:"price,omitempty" bson:"price"`
}
