package services

import (
	"github.com/rafaelbreno/work-at-olist/cmd/error_handler"
	"github.com/rafaelbreno/work-at-olist/domain"
	"github.com/rafaelbreno/work-at-olist/dto"
	"github.com/rafaelbreno/work-at-olist/repositories"
)

type AuthorService interface {
	Create(author dto.AuthorResponse) (dto.AuthorResponse, *error_handler.AppError)
	ImportCSV(authorsReq []dto.AuthorResponse) ([]dto.AuthorResponse, *error_handler.AppError)
	FindAll() ([]*dto.AuthorResponse, *error_handler.AppError)
	FindById(id uint) (*dto.AuthorResponse, *error_handler.AppError)
}

type DefaultAuthorService struct {
	repo repositories.AuthorRepository
}

func NewAuthorService(r repositories.AuthorRepository) DefaultAuthorService {
	return DefaultAuthorService{r}
}

func (s DefaultAuthorService) Create(authorReq dto.AuthorResponse) (dto.AuthorResponse, *error_handler.AppError) {
	author, err := s.repo.Create(authorReq)

	if err != nil {
		return dto.AuthorResponse{}, err
	}

	authorDTO := author.ToDTO()

	return authorDTO, nil
}

func (s DefaultAuthorService) ImportCSV(authorsReq []dto.AuthorResponse) ([]dto.AuthorResponse, *error_handler.AppError) {
	var domainAuthors []domain.Author

	for _, val := range authorsReq {
		domainAuthors = append(domainAuthors, domain.Author{
			Name: val.Name,
		})
	}

	authors, err := s.repo.ImportCSV(domainAuthors)

	if err != nil {
		return nil, err
	}

	var authorsDTO []dto.AuthorResponse

	for _, val := range authors {
		a := val.ToDTO()
		authorsDTO = append(authorsDTO, a)
	}

	return authorsDTO, nil
}

func (s DefaultAuthorService) FindAll() ([]*dto.AuthorResponse, *error_handler.AppError) {
	authors, err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}

	var authorsDTO []*dto.AuthorResponse

	for _, val := range authors {
		a := val.ToDTO()
		authorsDTO = append(authorsDTO, &a)
	}

	return authorsDTO, nil
}

func (s DefaultAuthorService) FindById(id uint) (*dto.AuthorResponse, *error_handler.AppError) {
	author, err := s.repo.FindById(id)

	if err != nil {
		return nil, err
	}

	authorDTO := author.ToDTO()

	return &authorDTO, nil
}
