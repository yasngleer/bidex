package types

type Item struct {
	ID          int    `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string `json:"name,omitempty" bson:"name"`
	Description string `json:"description,omitempty" bson:"description"`
	ImageUrl    string `json:"image_url,omitempty" bson:"image_url"`
	UserID      int    `json:"user_id,omitempty" bson:"user_id"`
	User        User   `json:"user,omitempty" bson:"user"`
	Bids        []Bid  `json:"bids" bson:"bids" gorm:"foreignKey:ItemID"`
}

type Bid struct {
	ID     int     `json:"id,omitempty" bson:"_id,omitempty"`
	UserID int     `json:"user_id,omitempty" bson:"user_id"`
	User   User    `json:"user,omitempty" `
	Price  float32 `json:"price,omitempty" bson:"price"`
	ItemID int
}

type ItemResponse struct {
	ID           int           `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string        `json:"name,omitempty" bson:"name"`
	Description  string        `json:"description,omitempty" bson:"description"`
	ImageUrl     string        `json:"image_url,omitempty" bson:"image_url"`
	UserResponse UserResponse  `json:"user,omitempty" bson:"user"`
	Bids         []BidResponse `json:"bids" bson:"bids" gorm:"foreignKey:ItemID"`
}

func (i Item) ToResponse() ItemResponse {
	return ItemResponse{
		ID:           i.ID,
		Name:         i.Name,
		Description:  i.Description,
		ImageUrl:     i.ImageUrl,
		UserResponse: i.User.ToResponse(),
		Bids:         bidsToResponse(i.Bids),
	}
}

func ItemsToResponse(is []Item) []ItemResponse {
	r := []ItemResponse{}
	for _, i := range is {
		r = append(r, i.ToResponse())
	}
	return r
}

type BidResponse struct {
	ID           int          `json:"id,omitempty" bson:"_id,omitempty"`
	UserResponse UserResponse `json:"user,omitempty" `
	Price        float32      `json:"price,omitempty" bson:"price"`
	ItemID       int
}

func (b Bid) ToResponse() BidResponse {
	return BidResponse{
		ID:           b.ID,
		UserResponse: b.User.ToResponse(),
		Price:        b.Price,
		ItemID:       b.ItemID,
	}
}

func bidsToResponse(b []Bid) []BidResponse {
	r := []BidResponse{}
	for _, i := range b {
		r = append(r, i.ToResponse())
	}
	return r
}
