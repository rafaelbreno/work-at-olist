package repositories

import (
	"errors"

	"github.com/rafaelbreno/work-at-olist/cmd/database"
	"github.com/rafaelbreno/work-at-olist/cmd/error_handler"
	"github.com/rafaelbreno/work-at-olist/domain"
	"github.com/rafaelbreno/work-at-olist/dto"
	"gorm.io/gorm"
)

// Reference to Authors DB methods
type AuthorRepository interface {
	ImportCSV(authorsReq []dto.AuthorResponse) ([]domain.Author, *error_handler.AppError)
	FindAll() ([]domain.Author, *error_handler.AppError)
	FindById(id uint) (*domain.Author, *error_handler.AppError)
}

type AuthorRepositoryDB struct {
	DB *gorm.DB
}

func (a AuthorRepositoryDB) ImportCSV(authorsReq []dto.AuthorResponse) ([]domain.Author, *error_handler.AppError) {
	var authors []domain.Author

	if err := a.
		DB.
		CreateInBatches(&authors, 100).
		Error; err != nil {
		return []domain.Author{}, error_handler.NewUnexpectedError("Unable to stablish a DB conneciton", error_handler.SetTrace())
	}

	return authors, nil
}

func (a AuthorRepositoryDB) FindAll() ([]domain.Author, *error_handler.AppError) {
	var authors []domain.Author

	if err := a.
		DB.
		Find(&authors).
		Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return []domain.Author{}, error_handler.NewUnexpectedError("Unable to stablish a DB conneciton", error_handler.SetTrace())
	}

	return authors, nil
}

func (a AuthorRepositoryDB) FindById(id uint) (*domain.Author, *error_handler.AppError) {
	var author domain.Author

	if err := a.DB.Where("id = ?", id).First(&author).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Author{}, error_handler.NewNotFoundError("Author not found", error_handler.SetTrace())
	}

	return &author, nil
}

func NewAuthorRepositoryDB() AuthorRepositoryDB {
	return AuthorRepositoryDB{database.PGConn.Conn}
}
