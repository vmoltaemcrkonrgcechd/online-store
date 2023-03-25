package usecase

import (
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/usecase/repo"
	"io"
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

func (uc ImageUseCase) GetImage(path string) ([]byte, error) {
	image, err := os.Open(fmt.Sprintf("%s\\%s", uc.dir, path))
	if err != nil {
		log.Println(err)
		if err == os.ErrNotExist {
			return nil, fiber.NewError(http.StatusBadRequest,
				"изображения не существует")
		}

		return nil, fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при получении изображения")
	}
	defer image.Close()

	return uc.readFile(image)
}

func (uc ImageUseCase) getTypeAndExt(contentType string) (string, string) {
	t := strings.Split(contentType, "/")

	return t[0], t[1]
}

func (uc ImageUseCase) readHeader(header *multipart.FileHeader) ([]byte, error) {
	image, err := header.Open()
	if err != nil {
		log.Println(err)
		return nil, fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при открытии файла")
	}
	defer image.Close()

	return uc.readFile(image)
}

func (uc ImageUseCase) readFile(file io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(file); err != nil {
		log.Println(err)
		return nil, fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при чтении файла")
	}

	return buf.Bytes(), nil
}
