package services

import (
	"github.com/rafaelbreno/work-at-olist/cmd/error_handler"
	"github.com/rafaelbreno/work-at-olist/dto"
	"github.com/rafaelbreno/work-at-olist/repositories"
)

type BookService interface {
	Create(bookReq dto.BookResponse) (*dto.BookResponse, *error_handler.AppError)
	Find(bookReq dto.BookResponse) (*dto.BookResponse, *error_handler.AppError)
	Update(id uint, bookReq dto.BookResponse) (*dto.BookResponse, *error_handler.AppError)
	Delete(bookReq dto.BookResponse) (*dto.BookResponse, *error_handler.AppError)
}

type DefaulBookService struct {
	repo repositories.BookRepository
}

func NewBookService(r repositories.BookRepository) DefaulBookService {
	return DefaulBookService{r}
}

func (s DefaulBookService) Create(bookReq dto.BookResponse) (*dto.BookResponse, *error_handler.AppError) {
	book, err := s.repo.Create(bookReq)

	if err != nil {
		return &dto.BookResponse{}, err
	}

	bookDTO := book.ToDTO()

	return &bookDTO, nil
}

func (s DefaulBookService) Find(bookReq dto.BookResponse) (*dto.BookResponse, *error_handler.AppError) {
	book, err := s.repo.Find(bookReq.ID)

	if err != nil {
		return &dto.BookResponse{}, err
	}

	bookDTO := book.ToDTO()

	return &bookDTO, nil
}

func (s DefaulBookService) Update(id uint, bookReq dto.BookResponse) (*dto.BookResponse, *error_handler.AppError) {
	book, err := s.repo.Update(id, bookReq)

	if err != nil {
		return &dto.BookResponse{}, err
	}

	bookDTO := book.ToDTO()

	return &bookDTO, nil
}

func (s DefaulBookService) Delete(bookReq dto.BookResponse) (*dto.BookResponse, *error_handler.AppError) {
	book, err := s.repo.Delete(bookReq.ID)

	if err != nil {
		return &dto.BookResponse{}, err
	}

	bookDTO := book.ToDTO()

	return &bookDTO, nil
}
