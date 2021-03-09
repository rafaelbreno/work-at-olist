package dto

type BookResponse struct {
	ID              uint             `json:"id"`
	Name            string           `json:"name"`
	Edition         uint             `json:"edition"`
	PublicationYear uint             `json:"publication_year"`
	Authors         []AuthorResponse `json:"authors"`
}
