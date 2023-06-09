package hotels_dto

type ImageDto struct {
	ID      int    `json:"id"`
	Path    string `json:"path"`
	HotelID int    `json:"hotel_id"`
}

type ImagesDto struct {
	Images []ImageDto `json:"images"`
}
