package domain

import (
	"github.com/rafaelbreno/work-at-olist/dto"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name            string   `gorm:"name"`
	Edition         uint     `gorm:"edition"`
	PublicationYear uint     `gorm:"publication_year"`
	Authors         []Author `gorm:"authors"`
}

func (b *Book) ToDTO() dto.BookResponse {
	var authorsDTO []dto.AuthorResponse
	for _, val := range b.Authors {
		authorsDTO = append(authorsDTO, val.ToDTO())
	}

	return dto.BookResponse{
		Name:            b.Name,
		Edition:         b.Edition,
		PublicationYear: b.PublicationYear,
		Authors:         authorsDTO,
	}
}
