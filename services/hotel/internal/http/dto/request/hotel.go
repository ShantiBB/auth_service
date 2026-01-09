package request

type Location struct {
	Latitude  *float64 `json:"latitude" validate:"required,gte=-90,lte=90"`
	Longitude *float64 `json:"longitude" validate:"required,gte=-180,lte=180"`
}

type HotelCreate struct {
	Title       *string  `json:"title" validate:"required,min=3,max=100"`
	OwnerID     *int64   `json:"owner_id" validate:"required,gt=0"`
	Description *string  `json:"description" validate:"omitempty,max=1000"`
	Address     *string  `json:"address" validate:"required,min=5,max=300"`
	Location    Location `json:"location" validate:"required"`
}

type HotelUpdate struct {
	Description *string  `json:"description" validate:"omitempty,max=1000"`
	Address     *string  `json:"address" validate:"required,min=5,max=300"`
	Location    Location `json:"location" validate:"required"`
}

type HotelTitleUpdate struct {
	Title *string `json:"title" validate:"required,min=3,max=100"`
}
