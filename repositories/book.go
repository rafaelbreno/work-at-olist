package repositories

import (
	"github.com/rafaelbreno/work-at-olist/cmd/database"
	"github.com/rafaelbreno/work-at-olist/cmd/error_handler"
	"github.com/rafaelbreno/work-at-olist/domain"
	"github.com/rafaelbreno/work-at-olist/dto"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(bookReq dto.BookResponse) (*domain.Book, *error_handler.AppError)
	Find(id uint) (*domain.Book, *error_handler.AppError)
	Update(id uint, bookReq dto.BookResponse) (*domain.Book, *error_handler.AppError)
	Delete(id uint) (*domain.Book, *error_handler.AppError)
}

type BookRepositoryDB struct {
	DB *gorm.DB
}

func NewBookRepositoryDB() BookRepositoryDB {
	return BookRepositoryDB{database.PGConn.Conn}
}

func (b BookRepositoryDB) Create(bookReq dto.BookResponse) (*domain.Book, *error_handler.AppError) {
	book := domain.Book{
		Name:            bookReq.Name,
		PublicationYear: bookReq.PublicationYear,
		Edition:         bookReq.PublicationYear,
		Authors:         bookReq.Authors,
	}

	if err := b.DB.Create(&book).Error; err != nil {
		return &domain.Book{}, error_handler.NewUnexpectedError(err.Error(), error_handler.SetTrace())
	}

	return &book, nil
}

func (b BookRepositoryDB) Find(id uint) (*domain.Book, *error_handler.AppError) {
	var book domain.Book

	if err := b.
		DB.
		Where("id = ?", id).
		First(&book).
		Error; err != nil {
		return &domain.Book{}, error_handler.NewNotFoundError(err.Error(), error_handler.SetTrace())
	}

	if err := b.
		DB.
		Delete(&book).
		Error; err != nil {
		return &domain.Book{}, error_handler.NewUnexpectedError(err.Error(), error_handler.SetTrace())
	}

	return &book, nil
}

func (b BookRepositoryDB) Update(id uint, bookReq dto.BookResponse) (*domain.Book, *error_handler.AppError) {
	var book domain.Book

	if err := b.
		DB.
		Where("id = ?", id).
		First(&book).
		Error; err != nil {
		return &domain.Book{}, error_handler.NewNotFoundError(err.Error(), error_handler.SetTrace())
	}

	book.Name = bookReq.Name
	book.Edition = bookReq.Edition
	book.PublicationYear = bookReq.PublicationYear
	book.Authors = bookReq.Authors

	if err := b.
		DB.
		Save(&book).
		Error; err != nil {
		return &domain.Book{}, error_handler.NewUnexpectedError(err.Error(), error_handler.SetTrace())
	}

	return &book, nil
}

func (b BookRepositoryDB) Delete(id uint) (*domain.Book, *error_handler.AppError) {
	var book domain.Book

	if err := b.
		DB.
		Where("id = ?", id).
		First(&book).
		Error; err != nil {
		return &domain.Book{}, error_handler.NewNotFoundError("Book not found", error_handler.SetTrace())
	}

	if err := b.
		DB.
		Delete(&book).
		Error; err != nil {
		return &domain.Book{}, error_handler.NewUnexpectedError("Couldn't delete Book", error_handler.SetTrace())
	}

	return &book, nil
}
