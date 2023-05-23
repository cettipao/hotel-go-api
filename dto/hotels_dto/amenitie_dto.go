package hotels_dto

type AmenitieDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type AmenitiesDto struct {
	Amenities []AmenitieDto `json:"amenities"`
}
