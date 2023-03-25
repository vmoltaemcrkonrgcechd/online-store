package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/usecase"
	"log"
	"net/http"
)

type imageRoutes struct {
	uc usecase.ImageUseCase
}

func withImageRoutes(app *fiber.App, uc usecase.ImageUseCase) {
	r := imageRoutes{uc}
	app.Post("/images", r.uploadImages)
}

//	@tags		изображения
//	@accept		multipart/form-data
//	@param		images	formData	[]file	true	"изображения"	collectionFormat(multi)
//	@success	201		{array}		string
//	@router		/images [post]
func (r imageRoutes) uploadImages(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		log.Println(err)
		return fiber.NewError(http.StatusBadRequest,
			"неверные данные")
	}

	var imagePaths []string
	if imagePaths, err = r.uc.SaveImages(form.File["images"]); err != nil {
		return err
	}

	return ctx.JSON(imagePaths)
}
