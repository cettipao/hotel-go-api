package hotels_dto

type HotelDto struct {
	Id             int      `json:"id"`
	Name           string   `json:"name"`
	RoomsAvailable int      `json:"rooms_available"`
	Description    string   `json:"description"`
	Amenities      []string `json:"amenities"`
	Images         []string `json:"images"`
}

type HotelsDto struct {
	Hotels []HotelDto `json:"hotels"`
}
