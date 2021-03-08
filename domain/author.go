package domain

import (
	"github.com/rafaelbreno/work-at-olist/dto"
	"gorm.io/gorm"
)

/* gorm.Model will add the following fields
 * ID        uint           `gorm:"primaryKey"`
 * CreatedAt time.Time
 * UpdatedAt time.Time
 * DeletedAt gorm.DeletedAt `gorm:"index"`
**/

type Author struct {
	gorm.Model
	Name string `gorm:"name"`
}

func (a *Author) ToDTO() dto.AuthorResponse {
	return dto.AuthorResponse{
		Name: a.Name,
		ID:   a.ID,
	}
}
