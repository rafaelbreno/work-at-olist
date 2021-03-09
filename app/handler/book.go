package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelbreno/work-at-olist/cmd/error_handler"
	"github.com/rafaelbreno/work-at-olist/cmd/helper"
	"github.com/rafaelbreno/work-at-olist/cmd/logger"
	"github.com/rafaelbreno/work-at-olist/dto"
	"github.com/rafaelbreno/work-at-olist/repositories"
	"github.com/rafaelbreno/work-at-olist/services"
)

type BookHandlers struct {
	service services.BookService
}

func GetBookHandlers() BookHandlers {
	return BookHandlers{
		service: services.NewBookService(repositories.NewBookRepositoryDB()),
	}
}

func (h *BookHandlers) Create(c *gin.Context) {
	var bookReq dto.BookResponse

	if err := c.ShouldBindJSON(&bookReq); err != nil {
		errHandler := error_handler.AppError{
			HTTPStatus: http.StatusBadRequest,
			Err:        err,
			Trace:      error_handler.SetTrace(),
		}
		c.JSON(errHandler.HTTPStatus, errHandler.GetJSON())
		logger.Error(errHandler.GetTrace())
		return
	}

	book, err := h.service.Create(bookReq)

	if err != nil {
		c.JSON(err.StatusCode(), err.GetJSON())
		logger.Error(err.GetTrace())
	}

	c.JSON(http.StatusOK, gin.H{"book": book})
}

func (h *BookHandlers) Find(c *gin.Context) {
	id := helper.StrToUint(c.Param("id"))
	bookReq := dto.BookResponse{
		ID: id,
	}

	book, err := h.service.Find(bookReq)

	if err != nil {
		c.JSON(err.StatusCode(), err.StatusCode())
		logger.Error(err.GetTrace())
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": book})
}
func (h *BookHandlers) Update(c *gin.Context) {
	id := helper.StrToUint(c.Param("id"))
	var bookReq dto.BookResponse

	if err := c.ShouldBindJSON(&bookReq); err != nil {
		errHandler := error_handler.AppError{
			HTTPStatus: http.StatusBadRequest,
			Err:        err,
			Trace:      error_handler.SetTrace(),
		}
		c.JSON(errHandler.HTTPStatus, errHandler.GetJSON())
		logger.Error(errHandler.GetTrace())
		return
	}

	book, err := h.service.Update(id, bookReq)

	if err != nil {
		c.JSON(err.StatusCode(), err.StatusCode())
		logger.Error(err.GetTrace())
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": book})
}
func (h *BookHandlers) Delete(c *gin.Context) {
	id := helper.StrToUint(c.Param("id"))

	bookReq := dto.BookResponse{
		ID: id,
	}

	book, err := h.service.Delete(bookReq)

	if err != nil {
		c.JSON(err.StatusCode(), err.StatusCode())
		logger.Error(err.GetTrace())
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": book})
}
