package services

import (
	amenitieCliente "mvc-go/clients/amenitie"
	hotelCliente "mvc-go/clients/hotel"
	"mvc-go/dto/hotels_dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotelById(id int) (hotels_dto.HotelDto, e.ApiError)
	GetHotels() (hotels_dto.HotelsDto, e.ApiError)
	InsertHotel(userDto hotels_dto.HotelDto) (hotels_dto.HotelDto, e.ApiError)
	AddAmenitieToHotel(hotelID, amenitieID int) e.ApiError
	RemoveAmenitieToHotel(hotelID, amenitieID int) e.ApiError
	DeleteHotelById(id int) e.ApiError
	UpdateHotel(hotelDto hotels_dto.HotelDto, id int) (hotels_dto.HotelDto, e.ApiError)
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetHotelById(id int) (hotels_dto.HotelDto, e.ApiError) {

	var hotel = hotelCliente.MyClient.GetHotelById(id)
	var hotelDto hotels_dto.HotelDto

	if hotel.Id == 0 {
		return hotelDto, e.NewBadRequestApiError("hotels_dto not found")
	}

	hotelDto.Id = hotel.Id
	hotelDto.Name = hotel.Name
	hotelDto.Description = hotel.Description
	hotelDto.RoomsAvailable = hotel.RoomsAvailable

	// Obtener los amenities del hotel
	amenities := make([]string, 0)
	for _, amenity := range hotel.Amenities {
		amenities = append(amenities, amenity.Name)
	}
	hotelDto.Amenities = amenities

	// Obtener las imágenes del hotel
	images := make([]string, 0)
	for _, image := range hotel.Images {
		images = append(images, image.Path)
	}
	hotelDto.Images = images

	return hotelDto, nil
}

func (s *hotelService) DeleteHotelById(id int) e.ApiError {
	// Verificar si el hotel existe
	_, err := s.GetHotelById(id)
	if err != nil {
		return err
	}

	// Lógica para eliminar el hotel por su ID
	err = hotelCliente.DeleteHotelById(id)
	if err != nil {
		// Otros errores de eliminación del hotel
		return e.NewInternalServerApiError("Failed to delete hotel", err)
	}

	return nil // Sin errores, se eliminó el hotel correctamente
}

func (s *hotelService) GetHotels() (hotels_dto.HotelsDto, e.ApiError) {
	var hotels model.Hotels = hotelCliente.MyClient.GetHotels()
	hotelsList := make([]hotels_dto.HotelDto, 0)

	for _, hotel := range hotels {
		var hotelDto hotels_dto.HotelDto
		hotelDto.Id = hotel.Id
		hotelDto.Name = hotel.Name
		hotelDto.Description = hotel.Description
		hotelDto.RoomsAvailable = hotel.RoomsAvailable

		// Obtener los amenities del hotel
		amenities := make([]string, 0)
		for _, amenity := range hotel.Amenities {
			amenities = append(amenities, amenity.Name)
		}
		hotelDto.Amenities = amenities

		// Obtener las imágenes del hotel
		images := make([]string, 0)
		for _, image := range hotel.Images {
			images = append(images, image.Path)
		}
		hotelDto.Images = images

		hotelsList = append(hotelsList, hotelDto)
	}

	return hotels_dto.HotelsDto{
		Hotels: hotelsList,
	}, nil
}

func (s *hotelService) InsertHotel(hotelDto hotels_dto.HotelDto) (hotels_dto.HotelDto, e.ApiError) {

	var hotel model.Hotel

	hotel.Name = hotelDto.Name
	hotel.Description = hotelDto.Description
	hotel.RoomsAvailable = hotelDto.RoomsAvailable

	hotel = hotelCliente.InsertHotel(hotel)

	hotelDto.Id = hotel.Id

	return hotelDto, nil
}

func (s *hotelService) UpdateHotel(hotelDto hotels_dto.HotelDto, id int) (hotels_dto.HotelDto, e.ApiError) {

	var hotel model.Hotel

	hotel.Id = id
	hotel.Name = hotelDto.Name
	hotel.Description = hotelDto.Description
	hotel.RoomsAvailable = hotelDto.RoomsAvailable

	hotelCliente.UpdateHotel(hotel)

	return hotelDto, nil
}

func (s *hotelService) AddAmenitieToHotel(hotelID, amenitieID int) e.ApiError {
	// Obtener el hotel por su ID
	hotel := hotelCliente.MyClient.GetHotelById(hotelID)
	if hotel.Id == 0 {
		return e.NewNotFoundApiError("Hotel not found")
	}

	// Obtener la amenidad por su ID
	amenitie := amenitieCliente.GetAmenitieById(amenitieID)
	if amenitie.Id == 0 {
		return e.NewNotFoundApiError("Amenitie not found")
	}

	// Verificar si la amenidad ya está asociada al hotel
	for _, a := range hotel.Amenities {
		if a.Id == amenitieID {
			return e.NewBadRequestApiError("Amenitie already added to the hotel")
		}
	}

	// Asociar la amenidad al hotel
	hotel.Amenities = append(hotel.Amenities, &amenitie)
	hotelCliente.UpdateHotel(hotel)

	return nil
}

func (s *hotelService) RemoveAmenitieToHotel(hotelID, amenitieID int) e.ApiError {
	// Obtener el hotel por su ID
	hotel := hotelCliente.MyClient.GetHotelById(hotelID)
	if hotel.Id == 0 {
		return e.NewNotFoundApiError("Hotel not found")
	}

	// Obtener la amenidad por su ID
	amenitie := amenitieCliente.GetAmenitieById(amenitieID)
	if amenitie.Id == 0 {
		return e.NewNotFoundApiError("Amenitie not found")
	}

	// Verificar si la amenidad ya está asociada al hotel
	relation := false
	for _, a := range hotel.Amenities {
		if a.Id == amenitieID {
			relation = true
		}
	}
	if !relation {
		return e.NewBadRequestApiError("Amenitie not linked with hotel")
	}
	// Eliminar la amenidad al hotel
	// Encuentra el índice del amenitie que deseas eliminar
	var indexToRemove int = -1
	for i, a := range hotel.Amenities {
		if a.Id == amenitie.Id {
			indexToRemove = i
			break
		}
	}

	// Si se encontró el amenitie, elimínalo de la lista
	if indexToRemove != -1 {
		hotel.Amenities = append(hotel.Amenities[:indexToRemove], hotel.Amenities[indexToRemove+1:]...)
	}

	// Actualiza el hotel en la base de datos
	hotelCliente.UpdateHotel(hotel)
	hotelCliente.DeleteLinkAmenitieHotel(hotelID, amenitieID)

	return nil
}
