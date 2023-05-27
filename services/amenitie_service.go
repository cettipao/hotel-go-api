package services

import (
	amenitieCliente "mvc-go/clients/amenitie"
	"mvc-go/dto/hotels_dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type amenitieService struct{}

type amenitieServiceInterface interface {
	GetAmenitieById(id int) (hotels_dto.AmenitieDto, e.ApiError)
	GetAmenities() (hotels_dto.AmenitiesDto, e.ApiError)
	InsertAmenitie(amenitieDto hotels_dto.AmenitieDto) (hotels_dto.AmenitieDto, e.ApiError)
	GetAmenitiesByHotelId(hotelId int) (hotels_dto.AmenitiesDto, e.ApiError)
}

var (
	AmenitieService amenitieServiceInterface
)

func init() {
	AmenitieService = &amenitieService{}
}

func (s *amenitieService) GetAmenitieById(id int) (hotels_dto.AmenitieDto, e.ApiError) {
	var amenitie model.Amenitie = amenitieCliente.GetAmenitieById(id)
	var amenitieDto hotels_dto.AmenitieDto

	if amenitie.Id == 0 {
		return amenitieDto, e.NewBadRequestApiError("Amenitie not found")
	}

	amenitieDto.Id = amenitie.Id
	amenitieDto.Name = amenitie.Name

	return amenitieDto, nil
}

func (s *amenitieService) GetAmenities() (hotels_dto.AmenitiesDto, e.ApiError) {
	var amenities model.Amenities = amenitieCliente.GetAmenities()
	amenitiesList := make([]hotels_dto.AmenitieDto, 0)

	for _, amenitie := range amenities {
		var amenitieDto hotels_dto.AmenitieDto
		amenitieDto.Id = amenitie.Id
		amenitieDto.Name = amenitie.Name

		amenitiesList = append(amenitiesList, amenitieDto)
	}

	return hotels_dto.AmenitiesDto{
		Amenities: amenitiesList,
	}, nil
}

func (s *amenitieService) InsertAmenitie(amenitieDto hotels_dto.AmenitieDto) (hotels_dto.AmenitieDto, e.ApiError) {
	var amenitie model.Amenitie

	amenitie.Name = amenitieDto.Name

	amenitie = amenitieCliente.InsertAmenitie(amenitie)

	amenitieDto.Id = amenitie.Id

	return amenitieDto, nil
}

func (s *amenitieService) GetAmenitiesByHotelId(hotelId int) (hotels_dto.AmenitiesDto, e.ApiError) {
	var amenities model.Amenities = amenitieCliente.GetAmenitiesByHotelId(hotelId)
	amenitiesList := make([]hotels_dto.AmenitieDto, 0)

	for _, amenitie := range amenities {
		var amenitieDto hotels_dto.AmenitieDto
		amenitieDto.Id = amenitie.Id
		amenitieDto.Name = amenitie.Name

		amenitiesList = append(amenitiesList, amenitieDto)
	}

	return hotels_dto.AmenitiesDto{
		Amenities: amenitiesList,
	}, nil
}
