package reservations_dto

type RoomsAvailable struct {
	Rooms int `json:"rooms_available"`
}

type RoomInfo struct {
	Name           string `json:"name"`
	Id             int    `json:"hotel_id"`
	RoomsAvailable int    `json:"rooms_available"`
}

type RoomsResponse struct {
	Rooms []RoomInfo `json:"rooms"`
}
