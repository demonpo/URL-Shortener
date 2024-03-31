package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"goHexBoilerplate/src/modules/shortener/application/rest/schemas"
	"goHexBoilerplate/src/modules/shortener/domain/services"
	"net/http"
)

type ShortenerHandler struct {
	shortenerService *services.ShortenerService
}

func NewShortenerHandler(shortenerService *services.ShortenerService) *ShortenerHandler {
	fmt.Println("NewShortenerHandler")
	return &ShortenerHandler{
		shortenerService: shortenerService,
	}
}

func (h *ShortenerHandler) Create(ctx *gin.Context) {
	var shortener schemas.CreateShortenerSchema
	if err := ctx.ShouldBindJSON(&shortener); err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	validate := validator.New()

	// Validate the User struct
	err := validate.Struct(shortener)
	if err != nil {
		// Validation failed, handle the error
		validationErrors := err.(validator.ValidationErrors)
		HandleError(ctx, http.StatusBadRequest, validationErrors)
		return
	}

	newShortener, err := h.shortenerService.Create(services.CreateInput{
		Url: shortener.Url,
	})
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "New user created successfully",
		"data":    newShortener,
	})
}

func (h *ShortenerHandler) Redirect(ctx *gin.Context) {
	id := ctx.Param("id")
	shortener, err := h.shortenerService.GetById(id)

	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.Redirect(http.StatusFound, shortener.Url)
}
