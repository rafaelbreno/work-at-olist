package handler

import (
	"bufio"
	"encoding/csv"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelbreno/work-at-olist/cmd/error_handler"
	"github.com/rafaelbreno/work-at-olist/cmd/helper"
	"github.com/rafaelbreno/work-at-olist/cmd/logger"
	"github.com/rafaelbreno/work-at-olist/dto"
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

func (h *AuthorHandlers) ImportCSV(c *gin.Context) {
	csvFormFile, _, errFile := c.Request.FormFile("contacts")

	defer csvFormFile.Close()

	if errFile != nil {
		errHandler := error_handler.AppError{
			HTTPStatus: http.StatusBadRequest,
			Err:        errFile,
		}
		c.JSON(errHandler.StatusCode(), errHandler.GetJSON())
		logger.Error(errHandler.GetTrace())
		return
	}

	csvReader := csv.NewReader(bufio.NewReader(csvFormFile))

	var authors []dto.AuthorResponse

	for {
		line, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		authors = append(authors, dto.AuthorResponse{
			Name: line[0],
		})
	}

	authors, err := h.service.ImportCSV(authors)

	if err != nil {
		c.JSON(err.StatusCode(), err.GetJSON())
		logger.Error(err.GetTrace())
		return
	}

	c.JSON(http.StatusOK, gin.H{"authors": authors})
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
