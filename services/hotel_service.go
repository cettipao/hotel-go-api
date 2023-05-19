package services

import (
	hotelCliente "mvc-go/clients/hotel"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotelById(id int) (dto.HotelDetailDto, e.ApiError)
	GetHotels() (dto.HotelsDto, e.ApiError)
	InsertHotel(userDto dto.HotelDto) (dto.HotelDto, e.ApiError)
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetHotelById(id int) (dto.HotelDetailDto, e.ApiError) {

	var hotel model.Hotel = hotelCliente.GetHotelById(id)
	var hotelDetailDto dto.HotelDetailDto

	if hotel.Id == 0 {
		return hotelDetailDto, e.NewBadRequestApiError("hotel not found")
	}

	hotelDetailDto.Name = hotel.Name
	hotelDetailDto.Description = hotel.Description
	hotelDetailDto.RoomsAvailable = hotel.RoomsAvailable
	/*
		for _, reservation := range user.Reservations {
			var dtoReservation dto.ReservationDto

			dtoReservation.Id = reservation.Id
			dtoReservation.HotelName = reservation.Name

			userDetailDto.ReservationsDto = append(userDetailDto.ReservationsDto, dtoReservation)
		}*/

	return hotelDetailDto, nil
}

func (s *hotelService) GetHotels() (dto.HotelsDto, e.ApiError) {

	var hotels model.Hotels = hotelCliente.GetHotels()
	var hotelsDto dto.HotelsDto

	for _, hotel := range hotels {
		var hotelDto dto.HotelDto
		hotelDto.Id = hotel.Id
		hotelDto.Name = hotel.Name
		hotelDto.Description = hotel.Description
		hotelDto.RoomsAvailable = hotel.RoomsAvailable

		hotelsDto = append(hotelsDto, hotelDto)
	}

	return hotelsDto, nil
}

func (s *hotelService) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {

	var hotel model.Hotel

	hotel.Name = hotelDto.Name
	hotel.Description = hotelDto.Description
	hotel.RoomsAvailable = hotelDto.RoomsAvailable

	hotel = hotelCliente.InsertHotel(hotel)

	hotelDto.Id = hotel.Id

	return hotelDto, nil
}
