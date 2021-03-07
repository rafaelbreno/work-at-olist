package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelbreno/work-at-olist/cmd/helper"
	"github.com/rafaelbreno/work-at-olist/cmd/logger"
	"github.com/rafaelbreno/work-at-olist/repositories"
	"github.com/rafaelbreno/work-at-olist/services"
)

type AuthorHandlers struct {
	service services.AuthorService
}

func GetAuthorHandlers() AuthorHandlers {
	return AuthorHandlers{
		service: services.NewAuthorService(repositories.NewAuthorRepositoryDB()),
	}
}

func (h *AuthorHandlers) FindAll(c *gin.Context) {
	authors, err := h.service.FindAll()

	if err != nil {
		c.JSON(err.StatusCode(), err.GetJSON())
		logger.Error(err.GetTrace())
		return
	}

	c.JSON(http.StatusOK, gin.H{"authors": authors})
}

func (h *AuthorHandlers) FindById(c *gin.Context) {
	id := helper.StrToUint(c.Param("id"))
	author, err := h.service.FindById(id)

	if err != nil {
		c.JSON(err.StatusCode(), err.GetJSON())
		logger.Error(err.GetTrace())
		return
	}

	c.JSON(http.StatusOK, gin.H{"author": author})
}
