package services

import (
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
