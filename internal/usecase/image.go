package usecase

import (
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/usecase/repo"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type ImageUseCase struct {
	repo repo.ImageRepo
	dir  string
}

func NewImageUseCase(repo repo.ImageRepo, dir string) ImageUseCase {
	return ImageUseCase{repo, dir}
}

func (uc ImageUseCase) SaveImages(images []*multipart.FileHeader) ([]string, error) {
	if images == nil {
		return nil, fiber.NewError(http.StatusBadRequest,
			"должно быть хотя бы одно изображение")
	}

	var imagePaths []string
	for _, image := range images {
		typ, ext := uc.getTypeAndExt(image.Header.Get("Content-Type"))
		if typ != "image" {
			return nil, fiber.NewError(http.StatusBadRequest,
				"тип файла не разрешен")
		}

		data, err := uc.readHeader(image)
		if err != nil {
			return nil, err
		}

		imagePath := fmt.Sprintf("%s.%s", utils.UUID(), ext)
		if err = os.WriteFile(fmt.Sprintf("%s\\%s", uc.dir, imagePath),
			data, 0666); err != nil {
			log.Println(err)
			return nil, fiber.NewError(http.StatusInternalServerError,
				"произошла ошибка при сохранении изображений")
		}

		imagePaths = append(imagePaths, imagePath)
	}

	if err := uc.repo.SaveImagePaths(imagePaths); err != nil {
		return nil, err
	}

	return imagePaths, nil
}

func (uc ImageUseCase) getTypeAndExt(contentType string) (string, string) {
	t := strings.Split(contentType, "/")

	return t[0], t[1]
}

func (uc ImageUseCase) readHeader(header *multipart.FileHeader) ([]byte, error) {
	userFriendlyErr := fiber.NewError(http.StatusInternalServerError,
		"произошла ошибка при чтении файла")

	image, err := header.Open()
	if err != nil {
		log.Println(err)
		return nil, userFriendlyErr
	}
	defer image.Close()

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(image); err != nil {
		log.Println(err)
		return nil, userFriendlyErr
	}

	return buf.Bytes(), nil
}
