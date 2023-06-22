package services

import (
	"github.com/stretchr/testify/assert"
	hotelClient "mvc-go/clients/hotel"
	"mvc-go/model"
	"testing"
)

type TestHotels struct {
}

func (testHotels TestHotels) GetHotels() model.Hotels {
	hotel1 := model.Hotel{
		Id:             1,
		Name:           "Hotel A",
		RoomsAvailable: 10,
		Description:    "This is a beautiful hotel with great amenities.",
		Amenities: []*model.Amenitie{
			{Id: 1, Name: "Swimming Pool"},
			{Id: 2, Name: "Gym"},
		},
		Images: []*model.Image{
			{ID: 1, Path: "https://example.com/hotel1_image1.jpg"},
			{ID: 2, Path: "https://example.com/hotel1_image2.jpg"},
		},
	}

	hotel2 := model.Hotel{
		Id:             2,
		Name:           "Hotel B",
		RoomsAvailable: 10,
		Description:    "Experience luxury at its finest in this modern hotel.",
		Amenities: []*model.Amenitie{
			{Id: 3, Name: "Spa"},
			{Id: 4, Name: "Restaurant"},
		},
		Images: []*model.Image{
			{ID: 3, Path: "https://example.com/hotel2_image1.jpg"},
			{ID: 4, Path: "https://example.com/hotel2_image2.jpg"},
		},
	}

	hotels := model.Hotels{hotel1, hotel2}
	return hotels
}

func (testHotels TestHotels) GetHotelById(id int) model.Hotel {
	hotel := model.Hotel{
		Id:             id,
		Name:           "Hotel A",
		RoomsAvailable: 10,
		Description:    "This is a beautiful hotel with great amenities.",
		Amenities: []*model.Amenitie{
			{Id: 1, Name: "Swimming Pool"},
			{Id: 2, Name: "Gym"},
		},
		Images: []*model.Image{
			{ID: 1, Path: "https://example.com/hotel1_image1.jpg"},
			{ID: 2, Path: "https://example.com/hotel1_image2.jpg"},
		},
	}
	return hotel
}

func TestGetHotels(t *testing.T) {
	//prepare
	hotelClient.MyClient = TestHotels{}

	//act
	hotels, _ := HotelService.GetHotels()

	//assert
	assert.Equal(t, hotels.Hotels[0].Name, "Hotel A")
	assert.Equal(t, hotels.Hotels[0].Id, 1)
}

func TestGetHotelById(t *testing.T) {
	//prepare
	hotelClient.MyClient = TestHotels{}

	//act
	hotel, _ := HotelService.GetHotelById(3)

	//assert
	assert.Equal(t, hotel.Id, 3)
}
