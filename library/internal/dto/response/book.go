package response

type BookDto struct {
	Id                int32  `json:"id"`
	Title             string `json:"title"`
	PublishedYear     int32  `json:"published_year"`
	AvailableQuantity int32  `json:"available_quantity"`
	BorrowQuantity    int32  `json:"borrow_quantity"`
}
type BookDetailDto struct {
	BookDto
	Authors []AuthorDto `json:"authors"`
	Genres  []GenreDto  `json:"genres"`
}
