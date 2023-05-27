package services

import (
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	imageClient "mvc-go/clients/image"
	images_dto "mvc-go/dto/hotels_dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
	"os"
	"path/filepath"
)

type imageService struct{}

type imageServiceInterface interface {
	GetImageById(id int) (images_dto.ImageDto, e.ApiError)
	GetImagesByHotelId(hotelID int) (images_dto.ImagesDto, e.ApiError)
	InsertImage(hotelID int, imageFile *multipart.FileHeader) (images_dto.ImageDto, e.ApiError)
}

var (
	ImageService imageServiceInterface
)

func init() {
	ImageService = &imageService{}
}

func (s *imageService) GetImageById(id int) (images_dto.ImageDto, e.ApiError) {
	image := imageClient.GetImageById(id)
	if image.ID == 0 {
		return images_dto.ImageDto{}, e.NewNotFoundApiError("Image not found")
	}

	imageDto := images_dto.ImageDto{
		ID:   image.ID,
		Path: image.Path,
	}

	return imageDto, nil
}

func (s *imageService) GetImagesByHotelId(hotelID int) (images_dto.ImagesDto, e.ApiError) {
	images := imageClient.GetImagesByHotelId(hotelID)
	imageDtos := make([]images_dto.ImageDto, len(images))

	for i, image := range images {
		imageDto := images_dto.ImageDto{
			ID:   image.ID,
			Path: image.Path,
		}
		imageDtos[i] = imageDto
	}

	return images_dto.ImagesDto{
		Images: imageDtos,
	}, nil
}

func (s *imageService) InsertImage(hotelID int, imageFile *multipart.FileHeader) (images_dto.ImageDto, e.ApiError) {
	// Crear imageDto para el retorno
	var imageDto images_dto.ImageDto

	// Generar un nombre único para el archivo de imagen
	fileName := uuid.New().String()

	// Obtener la extensión del archivo
	fileExt := filepath.Ext(imageFile.Filename)

	// Construir la ruta completa del archivo
	filePath := "images" + "/" + fileName + fileExt

	// Guardar el archivo en el directorio correspondiente
	err := saveImageToFile(imageFile, filePath)
	if err != nil {
		// Manejar el error en caso de fallo al guardar la imagen
		return imageDto, e.NewInternalServerApiError("Failed to save image", err)
	}

	// Crear una nueva instancia de model.Image
	image := model.Image{
		Path:    filePath,
		HotelID: hotelID,
	}

	// Llamar al cliente imageCliente para insertar la imagen
	image = imageClient.InsertImage(image)

	// Actualizar imageDto con el ID generado
	imageDto.ID = image.ID
	imageDto.Path = image.Path
	imageDto.HotelID = image.HotelID

	return imageDto, nil
}

func saveImageToFile(imageFile *multipart.FileHeader, filePath string) error {
	// Abrir el archivo cargado
	file, err := imageFile.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	// Crear el archivo destino en el sistema de archivos
	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copiar el contenido del archivo cargado al archivo destino
	_, err = io.Copy(dst, file)
	if err != nil {
		return err
	}

	return nil
}
