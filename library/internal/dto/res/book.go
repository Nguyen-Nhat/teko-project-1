package res

import "library/internal/database"

type BookDto struct {
	Id                int32  `json:"id"`
	Title             string `json:"title"`
	PublishedYear     int32  `json:"published_year"`
	AvailableQuantity int32  `json:"available_quantity"`
	BorrowQuantity    int32  `json:"borrow_quantity"`
}

func (b *BookDto) FromModel(book database.Book) {
	b.Id = book.ID
	b.Title = book.Title
	b.PublishedYear = book.PublishedYear
	b.AvailableQuantity = book.AvailableQuantity
	b.BorrowQuantity = book.BorrowQuantity
}

type BookDetailDto struct {
	BookDto
	Authors []AuthorDto `json:"authors"`
	Genres  []GenreDto  `json:"genres"`
}

func (b *BookDetailDto) FromModel(book database.Book, authors []database.Author, genres []database.Genre) {
	b.BookDto.FromModel(book)
	b.Authors = make([]AuthorDto, len(authors))
	for i, author := range authors {
		b.Authors[i].FromModel(author)
	}

	b.Genres = make([]GenreDto, len(genres))
	for i, genre := range genres {
		b.Genres[i].FromModel(genre)
	}
}
