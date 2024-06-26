package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	analyticsServices "goHexBoilerplate/src/modules/analytics/domain/services"
	"goHexBoilerplate/src/modules/shortener/application/rest/schemas"
	shortenerServices "goHexBoilerplate/src/modules/shortener/domain/services"
	"io/ioutil"
	"net/http"
	"strings"
)

type ShortenerHandler struct {
	shortenerService *shortenerServices.ShortenerService
	clickService     *analyticsServices.ClickService
}

func NewShortenerHandler(shortenerService *shortenerServices.ShortenerService, clickService *analyticsServices.ClickService) *ShortenerHandler {
	fmt.Println("NewShortenerHandler")
	return &ShortenerHandler{
		shortenerService: shortenerService,
		clickService:     clickService,
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

	newShortener, err := h.shortenerService.Create(shortenerServices.CreateInput{
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

	userAgent := ctx.GetHeader("User-Agent")
	ip := ctx.ClientIP()
	referer := ctx.Request.Header.Get("Referer")

	if isBrowserUserAgent(userAgent) {
		go func() {
			h.clickService.Create(analyticsServices.CreateInput{
				ShortenerId: shortener.Id,
				UserAgent:   &userAgent,
				UserIp:      &ip,
				ReferrerUrl: &referer,
			})
		}()
		ctx.Redirect(http.StatusFound, shortener.Url)
		return
	}

	// Fetch the content of the URL
	resp, err := http.Get(shortener.Url)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to fetch URL content")
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to read URL content")
		return
	}

	// Return the HTML content as a string
	ctx.String(http.StatusOK, string(body))

}

func isBrowserUserAgent(userAgent string) bool {
	validUserAgents := []string{"Mozilla", "Chrome", "Safari", "Opera", "IE", "Edge"}
	for _, ua := range validUserAgents {
		if strings.Contains(userAgent, ua) {
			return true
		}
	}
	return false
}
