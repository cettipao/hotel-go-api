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
	GetHotelById(id int) (hotels_dto.HotelDetailDto, e.ApiError)
	GetHotels() (hotels_dto.HotelsDto, e.ApiError)
	InsertHotel(userDto hotels_dto.HotelDto) (hotels_dto.HotelDto, e.ApiError)
	AddAmenitieToHotel(hotelID, amenitieID int) e.ApiError
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetHotelById(id int) (hotels_dto.HotelDetailDto, e.ApiError) {

	var hotel model.Hotel = hotelCliente.GetHotelById(id)
	var hotelDetailDto hotels_dto.HotelDetailDto

	if hotel.Id == 0 {
		return hotelDetailDto, e.NewBadRequestApiError("hotels_dto not found")
	}

	hotelDetailDto.Name = hotel.Name
	hotelDetailDto.Description = hotel.Description
	hotelDetailDto.RoomsAvailable = hotel.RoomsAvailable
	/*
		for _, reservations_dto := range users_dto.Reservations {
			var dtoReservation dto.ReservationDto

			dtoReservation.Id = reservations_dto.Id
			dtoReservation.HotelName = reservations_dto.Name

			userDetailDto.ReservationsDto = append(userDetailDto.ReservationsDto, dtoReservation)
		}*/

	return hotelDetailDto, nil
}

func (s *hotelService) GetHotels() (hotels_dto.HotelsDto, e.ApiError) {
	var hotels model.Hotels = hotelCliente.GetHotels()
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

func (s *hotelService) AddAmenitieToHotel(hotelID, amenitieID int) e.ApiError {
	// Obtener el hotel por su ID
	hotel := hotelCliente.GetHotelById(hotelID)
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
