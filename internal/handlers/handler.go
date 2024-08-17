package handler

import (
	"net/http"
	"url-shortener/internal/handlers/dto"
	"url-shortener/internal/lib/random"
	url "url-shortener/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UrlHandler struct {
	Repo *url.UrlRepository
}

func (handler *UrlHandler) GetUrlByAlias(c *gin.Context) {
	alias := c.Param("alias")
	url, err := handler.Repo.GetUrlByAlias(alias)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, url.Value)
}

func (handler *UrlHandler) PostUrl(c *gin.Context) {
	var newUrl dto.PostUrlDTO

	if err := c.ShouldBindJSON(&newUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.New().Struct(newUrl); err != nil {
		validateErr := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
		return
	}

	if newUrl.Alias == "" {
		newUrl.Alias = random.NewRandomString(6)
	}

	result, err := handler.Repo.SaveUrl(url.UrlDTO{Value: newUrl.Value, Alias: newUrl.Alias})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (handler *UrlHandler) DeleteAlias(c *gin.Context) {
	alias := c.Param("alias")

	if alias == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "alias is required"})
	}

	result, err := handler.Repo.DeleteAlias(alias)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func New(db *gorm.DB) *UrlHandler {
	return &UrlHandler{Repo: &url.UrlRepository{Db: db}}
}
